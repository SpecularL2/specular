package l2

import (
	"context"

	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/specularl2/specular/clients/geth/specular/rollup/rpc/eth"
	"github.com/specularl2/specular/clients/geth/specular/utils/fmt"
)

// TODO: Use EngineClient in place of ExecutionBackend
// TODO: upgrade Geth to use new Engine types

type EngineClient struct{ *eth.EthClient }

type L2Config interface{ Endpoint() string }

type (
	// ForkChoice
	ForkChoiceResponse = beacon.ForkChoiceResponse
	ForkchoiceState    = beacon.ForkchoiceStateV1
	PayloadAttributes  = beacon.PayloadAttributesV1
	PayloadStatus      = beacon.PayloadStatusV1
	PayloadID          = beacon.PayloadID
	// NewPayload
	ExecutionPayload = beacon.ExecutableDataV1
)

func NewEngineClient(c eth.EthClient) *EngineClient { return &EngineClient{&c} }

func DialWithRetry(ctx context.Context, cfg L2Config) (*EngineClient, error) {
	l2Client, err := eth.DialWithRetry(ctx, cfg.Endpoint(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to dial L2 client: %v", err)
	}
	return &EngineClient{l2Client}, nil
}

func (c *EngineClient) ForkchoiceUpdate(
	ctx context.Context,
	forkChoice *ForkchoiceState,
	attributes *PayloadAttributes,
) (*ForkChoiceResponse, error) {
	var result ForkChoiceResponse
	if err := c.C.CallContext(ctx, &result, "engine_forkchoiceUpdatedV1", forkChoice, attributes); err != nil {
		return nil, fmt.Errorf("Failed to update fork-choice: %w", err)
	}
	return &result, nil
}

func (c *EngineClient) NewPayload(ctx context.Context, payload *ExecutionPayload) (*PayloadStatus, error) {
	var status PayloadStatus
	if err := c.C.CallContext(ctx, &status, "engine_newPayloadV1", payload); err != nil {
		return nil, fmt.Errorf("Failed to execute new payload: %w", err)
	}
	return &status, nil
}