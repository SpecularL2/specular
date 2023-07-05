package geth_prover

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/specularl2/specular/clients/geth/specular/prover/state"
	"math/big"
)

type GethBlockContext struct {
	Context vm.BlockContext
}

// implements L2ELClientBlockContextInterface

func SpecularCanTransferFuncFromGeth(f func(vm.StateDB, common.Address, *big.Int) bool) func(state.L2ELClientStateInterface, common.Address, *big.Int) bool {
	return func(s state.L2ELClientStateInterface, address common.Address, amount *big.Int) bool {
		return f(s.(GethState).StateDB, address, amount)
	}
}

func (c GethBlockContext) CanTransfer() state.CanTransferFunc {
	return SpecularCanTransferFuncFromGeth(c.Context.CanTransfer)
}

func (c GethBlockContext) GetHash() state.GetHashFunc {
	return (func(uint64) common.Hash)(c.Context.GetHash)
}

func (c GethBlockContext) Coinbase() common.Address {
	return c.Context.Coinbase
}

func (c GethBlockContext) GasLimit() uint64 {
	return c.Context.GasLimit
}

func (c GethBlockContext) BlockNumber() *big.Int {
	return c.Context.BlockNumber
}

func (c GethBlockContext) Time() *big.Int {
	return c.Context.Time
}

func (c GethBlockContext) Difficulty() *big.Int {
	return c.Context.Difficulty
}

func (c GethBlockContext) BaseFee() *big.Int {
	return c.Context.BaseFee
}

func (c GethBlockContext) Random() *common.Hash {
	return c.Context.Random
}
