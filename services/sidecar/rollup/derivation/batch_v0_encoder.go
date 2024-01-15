package derivation

import (
	"bytes"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/specularL2/specular/services/sidecar/utils/fmt"
	"github.com/specularL2/specular/services/sidecar/utils/log"
)

var (
	errBatchFull     = errors.New("batch full")
	errBatchTooSmall = errors.New("batch too small")
)

type V0Config interface {
	GetTargetBatchSize() uint64
}

type BatchV0Encoder struct {
	cfg        V0Config
	subBatches []*subBatch
	runningLen uint64
}

func NewBatchV0Encoder(cfg V0Config) *BatchV0Encoder {
	return &BatchV0Encoder{cfg, []*subBatch{newSubBatch()}, 0}
}

func (e *BatchV0Encoder) IsEmpty() bool { return e.size() == 0 }

// Flushes data queued to the returned byte-array either if the batch is ready, or if forced.
func (e *BatchV0Encoder) Flush(force bool) ([]byte, error) {
	// Return error if the batch is too small (unless forced).
	if !force && e.size() < e.cfg.GetTargetBatchSize() {
		return nil, errBatchTooSmall
	}
	var lastSubBatchIdx = len(e.subBatches) - 1
	if lastSubBatchIdx > 0 && e.sizeExceedsTarget() {
		// Ignore the open sub-batch if the batch can't fit it.
		lastSubBatchIdx -= 1
	} else {
		// Close the open sub-batch.
		e.closeSubBatch()
	}
	// Encode version.
	buf := bytes.NewBuffer(nil)
	if err := buf.WriteByte(e.getVersion()); err != nil {
		return nil, fmt.Errorf("failed to encode version: %w", err)
	}
	// Encode all sub-batches (except the open one).
	var (
		firstL2BlockNum = e.subBatches[0].FirstL2BlockNum
		lastL2BlockNum  = e.subBatches[lastSubBatchIdx].lastL2BlockNum()
	)
	if err := rlp.Encode(buf, e.subBatches[:lastSubBatchIdx+1]); err != nil {
		return nil, fmt.Errorf("failed to encode batch (l2# %d-%d): %w", firstL2BlockNum, lastL2BlockNum, err)
	}
	log.Info("Flushed batch", "first_l2#", firstL2BlockNum, "last_l2#", lastL2BlockNum, "size (B)", buf.Len())
	// Delete all sub-batches (except the open one).
	e.subBatches = e.subBatches[lastSubBatchIdx+1:]
	e.runningLen = 0
	return buf.Bytes(), nil
}

// Processes a block. If the block is non-empty and fits, add it to the current sub-batch.
// If the block belongs to a new epoch, close the current sub-batch and start a new one.
func (e *BatchV0Encoder) ProcessBlock(block *types.Block, isNewEpoch bool) error {
	var (
		// Block is empty
		shouldSkipBlock = len(block.Transactions()) == 0
		// Batch would exceed the target size with the current sub-batch.
		shouldCloseBatch = e.sizeExceedsTarget()
		// Should close sub-batch if we're closing the batch entirely, OR...
		// the block is empty, OR... the block belongs to a new epoch.
		shouldCloseSubBatch = shouldCloseBatch || shouldSkipBlock || isNewEpoch
	)
	if shouldCloseSubBatch {
		e.closeSubBatch()
	}
	// Enforce soft cap on batch size.
	if shouldCloseBatch {
		return errBatchFull
	}
	// Skip intrinsically-derivable blocks.
	if shouldSkipBlock {
		log.Info("Skipping intrinsically-derivable block", "block#", block.NumberU64())
		return nil
	}
	// Append a block's txs to the current sub-batch.
	var currSubBatch = e.getOpenSubBatch()
	if err := currSubBatch.appendTxBlock(block.NumberU64(), block.Transactions()); err != nil {
		return fmt.Errorf("could not append block of txs: %w", err)
	}
	return nil
}

func (e *BatchV0Encoder) Reset() {
	e.subBatches = []*subBatch{newSubBatch()}
	e.runningLen = 0
}

// Returns the data format version (v0).
func (e *BatchV0Encoder) getVersion() BatchEncoderVersion { return V0 }

// Returns the expected size of the encoded batch.
func (e *BatchV0Encoder) size() uint64 {
	var (
		openSubBatch    = e.getOpenSubBatch()
		openSubBatchLen uint64
	)
	if openSubBatch != nil {
		openSubBatchLen = openSubBatch.size()
	}
	return e.runningLen + openSubBatchLen
}

// Returns the open sub-batch.
func (e *BatchV0Encoder) getOpenSubBatch() *subBatch {
	if len(e.subBatches) == 0 {
		log.Errorf("Unexpected state detected", errors.New("no sub-batches exist"), "running_len", e.runningLen)
		return nil
	}
	return e.subBatches[len(e.subBatches)-1]
}

func (e *BatchV0Encoder) sizeExceedsTarget() bool {
	return e.size() > e.cfg.GetTargetBatchSize()
}

// Closes the current sub-batch, if non-empty.
func (e *BatchV0Encoder) closeSubBatch() {
	var currSubBatch = e.getOpenSubBatch()
	// No need to close if it's empty.
	if currSubBatch.contentSize == 0 {
		return
	}
	e.runningLen += currSubBatch.size()
	log.Info("Closing sub-batch...", "running_len", e.runningLen)
	e.subBatches = append(e.subBatches, newSubBatch())
}

type subBatch struct {
	FirstL2BlockNum uint64
	TxBlocks        []rawTxBlock
	contentSize     uint64 `rlp:"-"` // size of sub-batch content (# of bytes)
}

type rawTxBlock []hexutil.Bytes

func newSubBatch() *subBatch               { return &subBatch{contentSize: 0} }
func (s *subBatch) size() uint64           { return rlp.ListSize(s.contentSize) }
func (s *subBatch) lastL2BlockNum() uint64 { return s.FirstL2BlockNum + uint64(len(s.TxBlocks)) - 1 }

func (s *subBatch) appendTxBlock(blockNum uint64, txs types.Transactions) error {
	// Set the first L2 block number if it hasn't been set yet.
	if len(s.TxBlocks) == 0 {
		s.FirstL2BlockNum = blockNum
		s.contentSize = uint64(rlp.IntSize(blockNum))
	}
	// Append the block of txs to the sub-batch.
	marshalled, numBytes, err := marshallTxs(txs)
	if err != nil {
		return fmt.Errorf("could not marshall txs: %w", err)
	}
	s.TxBlocks = append(s.TxBlocks, marshalled)
	s.contentSize += uint64(numBytes)
	return nil
}

func marshallTxs(txs types.Transactions) (rawTxs rawTxBlock, numBytes int, err error) {
	rawTxs = make([]hexutil.Bytes, 0, len(txs))
	for i, tx := range txs {
		rawTx, err := tx.MarshalBinary()
		if err != nil {
			return nil, 0, fmt.Errorf("could not marshall tx %v in block %v: %w", i, tx.Hash(), err)
		}
		rawTxs = append(rawTxs, rawTx)
		numBytes += len(rawTx)
	}
	return rawTxs, numBytes, err
}

func decodeV0(data []byte) ([]subBatch, error) {
	var decoded []subBatch
	if err := rlp.Decode(bytes.NewReader(data), &decoded); err != nil {
		return nil, err
	}
	return decoded, nil
}
