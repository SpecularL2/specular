package services

import (
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/specularL2/specular/services/sidecar/rollup/rpc/eth/txmgr"
	"github.com/specularL2/specular/services/sidecar/utils/fmt"
	"github.com/urfave/cli/v2"
)

// TODO: rename due to naming conflict
type SystemConfig struct {
	ProtocolConfig     `toml:"protocol,omitempty"`
	L1Config           `toml:"l1,omitempty"`
	L2Config           `toml:"l2,omitempty"`
	DisseminatorConfig `toml:"disseminator,omitempty"`
	ValidatorConfig    `toml:"validator,omitempty"`
}

func (c *SystemConfig) Protocol() ProtocolConfig         { return c.ProtocolConfig }
func (c *SystemConfig) L1() L1Config                     { return c.L1Config }
func (c *SystemConfig) L2() L2Config                     { return c.L2Config }
func (c *SystemConfig) Disseminator() DisseminatorConfig { return c.DisseminatorConfig }
func (c *SystemConfig) Validator() ValidatorConfig       { return c.ValidatorConfig }

// Parses all CLI flags and returns a full system config.
func ParseSystemConfig(cliCtx *cli.Context) (*SystemConfig, error) {
	return parseFlags(cliCtx)
}

func parseFlags(cliCtx *cli.Context) (*SystemConfig, error) {
	protocolCfg, err := newProtocolConfigFromCLI(cliCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to parse protocol config: %w", err)
	}
	var disseminatorAddr common.Address
	if cliCtx.String(disseminatorEnableFlag.Name) != "" {
		disseminatorAddr = protocolCfg.Rollup.Genesis.SystemConfig.BatcherAddr
	}
	var validatorAddr common.Address
	if cliCtx.String(validatorEnableFlag.Name) != "" {
		validatorAddr = common.HexToAddress(cliCtx.String(validatorAddrFlag.Name))
	}
	var (
		l1ChainID            = protocolCfg.GetRollup().L1ChainID
		disseminatorTxMgrCfg = txmgr.NewConfigFromCLI(cliCtx, disseminatorTxMgrNamespace, l1ChainID, disseminatorAddr)
		validatorTxMgrCfg    = txmgr.NewConfigFromCLI(cliCtx, validatorTxMgrNamespace, l1ChainID, validatorAddr)
	)
	cfg := &SystemConfig{
		ProtocolConfig:     protocolCfg,
		L1Config:           newL1ConfigFromCLI(cliCtx),
		L2Config:           newL2ConfigFromCLI(cliCtx),
		DisseminatorConfig: newDisseminatorConfigFromCLI(cliCtx, disseminatorTxMgrCfg),
		ValidatorConfig:    newValidatorConfigFromCLI(cliCtx, validatorTxMgrCfg),
	}
	// Validate.
	if err := cfg.DisseminatorConfig.validate(); err != nil {
		return nil, fmt.Errorf("disseminator config invalid: %w", err)
	}
	if err := cfg.ValidatorConfig.validate(); err != nil {
		return nil, fmt.Errorf("validator config invalid: %w", err)
	}
	return cfg, nil
}

// Protocol configuration
type ProtocolConfig struct {
	Rollup          RollupConfig   `toml:"rollup,omitempty"`
	RollupAddr      common.Address `toml:"rollup_addr,omitempty"`       // L1 Rollup contract address
	L1OracleAddress common.Address `toml:"l1_oracle_address,omitempty"` // L2 Address of the L1Oracle
}

func newProtocolConfigFromCLI(cliCtx *cli.Context) (ProtocolConfig, error) {
	rollupCfg, err := NewRollupConfig(cliCtx.String(protocolRollupCfgPathFlag.Name))
	if err != nil {
		return ProtocolConfig{}, err
	}
	return ProtocolConfig{
		Rollup:          *rollupCfg,
		RollupAddr:      common.HexToAddress(cliCtx.String(protocolRollupAddrFlag.Name)),
		L1OracleAddress: common.HexToAddress(protocolL1OracleAddressFlag.Value),
	}, nil
}

func (c ProtocolConfig) GetSequencerInboxAddr() common.Address { return c.Rollup.BatchInboxAddress }
func (c ProtocolConfig) GetRollup() RollupConfig               { return c.Rollup }
func (c ProtocolConfig) GetRollupAddr() common.Address         { return c.RollupAddr }
func (c ProtocolConfig) GetL1ChainID() uint64                  { return c.Rollup.L2ChainID.Uint64() }
func (c ProtocolConfig) GetL2ChainID() uint64                  { return c.Rollup.L1ChainID.Uint64() }
func (c ProtocolConfig) GetL1OracleAddress() common.Address    { return c.L1OracleAddress }

// L1 configuration
type L1Config struct {
	Endpoint string `toml:"endpoint,omitempty"` // L1 API endpoint
}

func newL1ConfigFromCLI(cliCtx *cli.Context) L1Config {
	return L1Config{Endpoint: cliCtx.String(l1EndpointFlag.Name)}
}

func (c L1Config) GetEndpoint() string { return c.Endpoint }

// L2 configuration
type L2Config struct {
	Endpoint string `toml:"endpoint,omitempty"` // L2 API endpoint
	ChainID  uint64 `toml:"chainid,omitempty"`  // L2 chain ID
}

func newL2ConfigFromCLI(cliCtx *cli.Context) L2Config {
	return L2Config{Endpoint: cliCtx.String(l2EndpointFlag.Name)}
}

func (c L2Config) GetEndpoint() string { return c.Endpoint }

// Sequencer node configuration
type DisseminatorConfig struct {
	// Whether this node is a sequencer
	IsEnabled bool `toml:"enabled,omitempty"`
	// The address of this sequencer
	AccountAddr common.Address `toml:"account_addr,omitempty"`
	// The private key for AccountAddr
	PrivateKey *ecdsa.PrivateKey
	// The Clef Endpoint used for signing txs
	ClefEndpoint string `toml:"clef_endpoint,omitempty"`
	// Time between batch dissemination (DA) steps
	DisseminationInterval time.Duration `toml:"dissemination_interval,omitempty"`
	// Transaction manager configuration
	TxMgrCfg txmgr.Config `toml:"txmgr,omitempty"`
}

func (c DisseminatorConfig) GetIsEnabled() bool                      { return c.IsEnabled }
func (c DisseminatorConfig) GetAccountAddr() common.Address          { return c.AccountAddr }
func (c DisseminatorConfig) GetPrivateKey() *ecdsa.PrivateKey        { return c.PrivateKey }
func (c DisseminatorConfig) GetClefEndpoint() string                 { return c.ClefEndpoint }
func (c DisseminatorConfig) GetDisseminationInterval() time.Duration { return c.DisseminationInterval }
func (c DisseminatorConfig) GetTxMgrCfg() txmgr.Config               { return c.TxMgrCfg }

// Validates the configuration.
func (c DisseminatorConfig) validate() error {
	if !c.IsEnabled {
		return nil
	}
	if c.PrivateKey == nil && c.ClefEndpoint == "" {
		return fmt.Errorf("missing both private key and clef endpoint (require at least one)")
	}
	return nil
}

func newDisseminatorConfigFromCLI(
	cliCtx *cli.Context,
	txMgrCfg txmgr.Config,
) DisseminatorConfig {
	return DisseminatorConfig{
		IsEnabled:             cliCtx.Bool(disseminatorEnableFlag.Name),
		AccountAddr:           txMgrCfg.From,
		PrivateKey:            toPrivateKey(cliCtx.String(disseminatorPrivateKeyFlag.Name)),
		ClefEndpoint:          cliCtx.String(disseminatorClefEndpointFlag.Name),
		DisseminationInterval: time.Duration(cliCtx.Uint(disseminatorIntervalFlag.Name)) * time.Second,
		TxMgrCfg:              txMgrCfg,
	}
}

type ValidatorConfig struct {
	// Whether this node is a validator
	IsEnabled bool `toml:"enabled,omitempty"`
	// The address of this validator
	AccountAddr common.Address `toml:"account_addr,omitempty"`
	// The private key for AccountAddr
	PrivateKey *ecdsa.PrivateKey
	// The Clef Endpoint used for signing txs
	ClefEndpoint string `toml:"clef_endpoint,omitempty"`
	// Time between validation steps
	ValidationInterval time.Duration `toml:"validation_interval,omitempty"`
	// Transaction manager configuration
	TxMgrCfg txmgr.Config `toml:"txmgr,omitempty"`
}

func (c ValidatorConfig) GetIsEnabled() bool                   { return c.IsEnabled }
func (c ValidatorConfig) GetAccountAddr() common.Address       { return c.AccountAddr }
func (c ValidatorConfig) GetPrivateKey() *ecdsa.PrivateKey     { return c.PrivateKey }
func (c ValidatorConfig) GetClefEndpoint() string              { return c.ClefEndpoint }
func (c ValidatorConfig) GetValidationInterval() time.Duration { return c.ValidationInterval }
func (c ValidatorConfig) GetTxMgrCfg() txmgr.Config            { return c.TxMgrCfg }

// Validates the configuration.
func (c ValidatorConfig) validate() error {
	if !c.IsEnabled {
		return nil
	}
	if c.PrivateKey == nil && c.ClefEndpoint == "" {
		return fmt.Errorf("missing both private key and clef endpoint (require at least one)")
	}
	return nil
}

func newValidatorConfigFromCLI(
	cliCtx *cli.Context,
	txMgrCfg txmgr.Config,
) ValidatorConfig {
	return ValidatorConfig{
		IsEnabled:          cliCtx.Bool(validatorEnableFlag.Name),
		AccountAddr:        txMgrCfg.From,
		PrivateKey:         toPrivateKey(cliCtx.String(validatorPrivateKeyFlag.Name)),
		ClefEndpoint:       cliCtx.String(validatorClefEndpointFlag.Name),
		ValidationInterval: time.Duration(cliCtx.Uint(validatorValidationIntervalFlag.Name)) * time.Second,
		TxMgrCfg:           txMgrCfg,
	}
}

func toPrivateKey(keyStr string) *ecdsa.PrivateKey {
	if keyStr == "" {
		return nil
	}
	secretKey, err := crypto.HexToECDSA(keyStr[2:])
	if err != nil {
		panic("failed to parse secret key: " + err.Error())
	}
	return secretKey
}
