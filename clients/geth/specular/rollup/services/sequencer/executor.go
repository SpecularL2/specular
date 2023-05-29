package sequencer

import (
	"context"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/specularl2/specular/clients/geth/specular/rollup/utils"
	"github.com/specularl2/specular/clients/geth/specular/rollup/utils/fmt"
	"github.com/specularl2/specular/clients/geth/specular/rollup/utils/log"
)

// Responsible for executing transactions.
type executor struct {
	cfg     Config
	backend ExecutionBackend
	orderer orderer
}

// Responsible for ordering transactions (prior to their execution).
// TODO: Support:
// - PBS-style ordering: publicize current mempool and call remote engine API.
// - remote ordering + weak DA in single call (some systems conflate these roles -- e.g. Espresso)
type orderer interface {
	OrderTransactions(ctx context.Context, txs []*types.Transaction) ([]*types.Transaction, error)
	RegisterL2Client(l2Client L2Client)
}

// This goroutine fetches, orders and executes txs from the tx pool.
// Commits an empty block if no txs are received within an interval
// TODO: handle reorgs in the decentralized sequencer case.
// TODO: commit a msg-passing tx in empty block.
func (e *executor) start(ctx context.Context, l2Client L2Client) error {
	e.orderer.RegisterL2Client(l2Client)
	// Watch transactions in TxPool
	txsCh := make(chan core.NewTxsEvent, 4096)
	txsSub := e.backend.SubscribeNewTxsEvent(txsCh)
	defer txsSub.Unsubscribe()
	batchCh := utils.SubscribeBatched(ctx, txsCh, e.cfg.MinExecutionInterval(), e.cfg.MaxExecutionInterval())
	for {
		select {
		case evs := <-batchCh:
			var txs []*types.Transaction
			for _, ev := range evs {
				txs = append(txs, ev.Txs...)
			}
			if len(txs) == 0 {
				log.Trace("No txs received in last execution window.")
				continue
			} else {
				var err error
				txs, err = e.orderer.OrderTransactions(ctx, txs)
				if err != nil {
					return fmt.Errorf("Failed to order txs: %w", err)
				}
			}
			if len(txs) == 0 {
				log.Info("No txs to execute post-ordering.")
				continue
			}
			log.Info("Committing txs", "#txs", len(txs))
			err := e.backend.CommitTransactions(txs)
			if err != nil {
				return fmt.Errorf("Failed to commit txs: %w", err)
			}
			log.Info("Committed txs", "#txs", len(txs))
		case <-ctx.Done():
			log.Info("Aborting.")
			return nil
		}
	}
}
