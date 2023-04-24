package services

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type SystemConfig struct {
	L1Config
	SequencerConfig
	ValidatorConfig
	IndexerConfig
}

func (c *SystemConfig) L1() *L1Config               { return &c.L1Config }
func (c *SystemConfig) Sequencer() *SequencerConfig { return &c.SequencerConfig }
func (c *SystemConfig) Validator() *ValidatorConfig { return &c.ValidatorConfig }
func (c *SystemConfig) Indexer() *IndexerConfig     { return &c.IndexerConfig }

type L1Config struct {
	L1Endpoint           string         // L1 API endpoint
	L1ChainID            uint64         // L1 chain ID
	L1RollupGenesisBlock uint64         // L1 Rollup genesis block
	SequencerInboxAddr   common.Address // L1 SequencerInbox contract address
	RollupAddr           common.Address // L1 Rollup contract address
}

type IndexerConfig struct {
	IndexerAccountAddr common.Address
	IndexerPassphrase  string // The passphrase of the indexer account
}

type SequencerConfig struct {
	SequencerAccountAddr common.Address
	SequencerPassphrase  string        // The passphrase of the sequencer account
	ExecutionInterval    time.Duration // Maximum time between block execution
	SequencingInterval   time.Duration // Time between batch sequencing attempts
}

type ValidatorConfig struct {
	ValidatorAccountAddr common.Address
	ValidatorPassphrase  string // The passphrase of the validator account
	RollupStakeAmount    uint64 // Size of stake to deposit to rollup contract
	IsActiveCreator      bool   // Iff true, actively tries to create new assertions (not just for a challenge).
	IsActiveChallenger   bool   // Iff true, actively issues challenges as challenger. *Defends* against challenges regardless.
	IsResolver           bool   // Iff true, attempts to resolve assertions (by confirming or rejecting)
}
