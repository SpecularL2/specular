package infra

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/specularL2/specular/services/sidecar/rollup/derivation"
	"github.com/specularL2/specular/services/sidecar/rollup/rpc/bridge"
	"github.com/specularL2/specular/services/sidecar/rollup/rpc/eth"
	"github.com/specularL2/specular/services/sidecar/rollup/rpc/eth/txmgr"
	"github.com/specularL2/specular/services/sidecar/rollup/services"
	"github.com/specularL2/specular/services/sidecar/rollup/services/disseminator"
	"github.com/specularL2/specular/services/sidecar/rollup/services/validator"
	"math/big"
)

// TODO: This file is not a proper structure, but I need a bridge
// 	     between the existing code and the proper refactor

func CreateDisseminator(
	ctx context.Context,
	cfg *services.SystemConfig,
	l1State *eth.EthState,
) (*disseminator.BatchDisseminator, error) {
	l1TxMgr, err := createTxManager(ctx, "disseminator", cfg.L1().Endpoint, cfg.Protocol(), cfg.Disseminator())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize l1 tx manager: %w", err)
	}
	var (
		encoder      = derivation.NewBatchV0Encoder(cfg)
		batchBuilder = derivation.NewBatchBuilder(cfg, encoder)
		l2Client     = eth.NewLazilyDialedEthClient(cfg.L2().GetEndpoint())
	)
	return disseminator.NewBatchDisseminator(cfg.Disseminator(), batchBuilder, l1TxMgr, l1State, l2Client), nil
}

func CreateValidator(
	ctx context.Context,
	cfg *services.SystemConfig,
	l1State *eth.EthState,
) (*validator.Validator, error) {
	l1TxMgr, err := createTxManager(ctx, "validator", cfg.L1().Endpoint, cfg.Protocol(), cfg.Validator())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize l1 tx manager: %w", err)
	}
	l1Client, err := eth.DialWithRetry(ctx, cfg.L1().GetEndpoint())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize l1 client: %w", err)
	}
	l1BridgeClient, err := bridge.NewBridgeClient(l1Client, cfg.Protocol())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize l1 bridge client: %w", err)
	}
	l2Client := eth.NewLazilyDialedEthClient(cfg.L2().GetEndpoint())
	return validator.NewValidator(cfg.Validator(), l1TxMgr, l1BridgeClient, l1State, l2Client), nil
}

func CreateTxManager(
	ctx context.Context,
	name string,
	l1RpcUrl string,
	protocolCfg services.ProtocolConfig,
	serCfg serviceCfg,
) (*bridge.TxManager, error) {
	transactor, err := createTransactor(
		serCfg.GetAccountAddr(), serCfg.GetClefEndpoint(), serCfg.GetPrivateKey(), protocolCfg.GetL1ChainID(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize transactor: %w", err)
	}
	l1Client, err := eth.DialWithRetry(ctx, l1RpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize l1 client: %w", err)
	}
	signer := func(ctx context.Context, address common.Address, tx *ethTypes.Transaction) (*ethTypes.Transaction, error) {
		return transactor.Signer(address, tx)
	}
	return bridge.NewTxManager(txmgr.NewTxManager(log.New("service", name), serCfg.GetTxMgrCfg(), l1Client, signer), protocolCfg)
}

// Creates a transactor for the given account address, either using a clef endpoint (preferred) or secret key.
func CreateTransactor(
	accountAddress common.Address,
	clefEndpoint string,
	secretKey *ecdsa.PrivateKey,
	chainID uint64,
) (*bind.TransactOpts, error) {
	if clefEndpoint != "" {
		clef, err := external.NewExternalSigner(clefEndpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize external signer from clef endpoint: %w", err)
		}
		return bind.NewClefTransactor(clef, accounts.Account{Address: accountAddress}), nil
	}
	log.Warn("No external signer specified, using geth signer")
	return bind.NewKeyedTransactorWithChainID(secretKey, new(big.Int).SetUint64(chainID))
}

func createL1State(ctx context.Context, cfg *services.SystemConfig) (*eth.EthState, error) {
	l1Client, err := eth.DialWithRetry(ctx, cfg.L1().GetEndpoint())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize l1 client: %w", err)
	}
	l1State := eth.NewEthState()
	l1Syncer := eth.NewEthSyncer(l1State)
	l1Syncer.Start(ctx, l1Client)
	return l1State, nil
}
