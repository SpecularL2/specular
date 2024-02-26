// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MintableERC20FactoryMetaData contains all meta data concerning the MintableERC20Factory contract.
var MintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"MintableERC20Created\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516114c43803806114c483398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b60805161143361009160003960008181606f015261011701526114336000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806350d38c6c1461003b578063ee9a31a21461006a575b600080fd5b61004e61004936600461026d565b610091565b6040516001600160a01b03909116815260200160405180910390f35b61004e7f000000000000000000000000000000000000000000000000000000000000000081565b60006001600160a01b0384166101135760405162461bcd60e51b815260206004820152603760248201527f4d696e7461626c654552433230466163746f72793a206d7573742070726f766960448201527f64652072656d6f746520746f6b656e2061646472657373000000000000000000606482015260840160405180910390fd5b60007f0000000000000000000000000000000000000000000000000000000000000000858585604051610145906101bd565b6101529493929190610335565b604051809103906000f08015801561016e573d6000803e3d6000fd5b506040513381529091506001600160a01b0380871691908316907fd8d73ce8454cb13796d0cbdbbc836fc2251faa966d0ad8a176b933492562e5649060200160405180910390a3949350505050565b61107f8061037f83390190565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126101f157600080fd5b813567ffffffffffffffff8082111561020c5761020c6101ca565b604051601f8301601f19908116603f01168101908282118183101715610234576102346101ca565b8160405283815286602085880101111561024d57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561028257600080fd5b83356001600160a01b038116811461029957600080fd5b9250602084013567ffffffffffffffff808211156102b657600080fd5b6102c2878388016101e0565b935060408601359150808211156102d857600080fd5b506102e5868287016101e0565b9150509250925092565b6000815180845260005b81811015610315576020818501810151868301820152016102f9565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b03858116825284166020820152608060408201819052600090610361908301856102ef565b828103606084015261037381856102ef565b97965050505050505056fe60c06040523480156200001157600080fd5b506040516200107f3803806200107f833981016040819052620000349162000152565b8181600362000044838262000271565b50600462000053828262000271565b5050506001600160a01b0392831660805250501660a0526200033d565b80516001600160a01b03811681146200008857600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000b557600080fd5b81516001600160401b0380821115620000d257620000d26200008d565b604051601f8301601f19908116603f01168101908282118183101715620000fd57620000fd6200008d565b816040528381526020925086838588010111156200011a57600080fd5b600091505b838210156200013e57858201830151818301840152908201906200011f565b600093810190920192909252949350505050565b600080600080608085870312156200016957600080fd5b620001748562000070565b9350620001846020860162000070565b60408601519093506001600160401b0380821115620001a257600080fd5b620001b088838901620000a3565b93506060870151915080821115620001c757600080fd5b50620001d687828801620000a3565b91505092959194509250565b600181811c90821680620001f757607f821691505b6020821081036200021857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200026c57600081815260208120601f850160051c81016020861015620002475750805b601f850160051c820191505b81811015620002685782815560010162000253565b5050505b505050565b81516001600160401b038111156200028d576200028d6200008d565b620002a5816200029e8454620001e2565b846200021e565b602080601f831160018114620002dd5760008415620002c45750858301515b600019600386901b1c1916600185901b17855562000268565b600085815260208120601f198616915b828110156200030e57888601518255948401946001909101908401620002ed565b50858210156200032d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a051610d0e6200037160003960008181610272015281816103cf0152610480015260006101320152610d0e6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806340c10f1911610097578063a457c2d711610066578063a457c2d714610234578063a9059cbb14610247578063dd62ed3e1461025a578063ee9a31a21461026d57600080fd5b806340c10f19146101db57806370a08231146101f057806395d89b41146102195780639dc29fac1461022157600080fd5b806318160ddd116100d357806318160ddd1461019457806323b872dd146101a6578063313ce567146101b957806339509351146101c857600080fd5b806301ffc9a714610105578063033964be1461012d57806306fdde031461016c578063095ea7b314610181575b600080fd5b610118610113366004610ae2565b610294565b60405190151581526020015b60405180910390f35b6101547f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610124565b6101746102d2565b6040516101249190610b13565b61011861018f366004610b7d565b610364565b6002545b604051908152602001610124565b6101186101b4366004610ba7565b61037e565b60405160128152602001610124565b6101186101d6366004610b7d565b6103a2565b6101ee6101e9366004610b7d565b6103c4565b005b6101986101fe366004610be3565b6001600160a01b031660009081526020819052604090205490565b610174610466565b6101ee61022f366004610b7d565b610475565b610118610242366004610b7d565b610502565b610118610255366004610b7d565b61057d565b610198610268366004610bfe565b61058b565b6101547f000000000000000000000000000000000000000000000000000000000000000081565b60006301ffc9a760e01b6330a0c5a960e01b6001600160e01b031984168214806102ca57506001600160e01b0319848116908216145b949350505050565b6060600380546102e190610c31565b80601f016020809104026020016040519081016040528092919081815260200182805461030d90610c31565b801561035a5780601f1061032f5761010080835404028352916020019161035a565b820191906000526020600020905b81548152906001019060200180831161033d57829003601f168201915b5050505050905090565b6000336103728185856105b6565b60019150505b92915050565b60003361038c8582856106db565b610397858585610755565b506001949350505050565b6000336103728185856103b5838361058b565b6103bf9190610c6b565b6105b6565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104155760405162461bcd60e51b815260040161040c90610c8c565b60405180910390fd5b61041f82826108f9565b816001600160a01b03167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161045a91815260200190565b60405180910390a25050565b6060600480546102e190610c31565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104bd5760405162461bcd60e51b815260040161040c90610c8c565b6104c782826109b8565b816001600160a01b03167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161045a91815260200190565b60003381610510828661058b565b9050838110156105705760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b606482015260840161040c565b61039782868684036105b6565b600033610372818585610755565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166106185760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161040c565b6001600160a01b0382166106795760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161040c565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006106e7848461058b565b9050600019811461074f57818110156107425760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161040c565b61074f84848484036105b6565b50505050565b6001600160a01b0383166107b95760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161040c565b6001600160a01b03821661081b5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161040c565b6001600160a01b038316600090815260208190526040902054818110156108935760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161040c565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361074f565b6001600160a01b03821661094f5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161040c565b80600260008282546109619190610c6b565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b038216610a185760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b606482015260840161040c565b6001600160a01b03821660009081526020819052604090205481811015610a8c5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b606482015260840161040c565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91016106ce565b600060208284031215610af457600080fd5b81356001600160e01b031981168114610b0c57600080fd5b9392505050565b600060208083528351808285015260005b81811015610b4057858101830151858201604001528201610b24565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b0381168114610b7857600080fd5b919050565b60008060408385031215610b9057600080fd5b610b9983610b61565b946020939093013593505050565b600080600060608486031215610bbc57600080fd5b610bc584610b61565b9250610bd360208501610b61565b9150604084013590509250925092565b600060208284031215610bf557600080fd5b610b0c82610b61565b60008060408385031215610c1157600080fd5b610c1a83610b61565b9150610c2860208401610b61565b90509250929050565b600181811c90821680610c4557607f821691505b602082108103610c6557634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561037857634e487b7160e01b600052601160045260246000fd5b6020808252602c908201527f4d696e7461626c6545524332303a206f6e6c79206272696467652063616e206d60408201526b34b73a1030b73210313ab93760a11b60608201526080019056fea26469706673582212205b716129da7b34c33b760dfdcea73ceefd02020e32fb68bf767101a6b4ce2d4b64736f6c63430008110033a2646970667358221220881c5062c13c71f41eaabbcc1d241a4a91ba6133d6174a74ff1541e22900fac464736f6c63430008110033",
}

// MintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MintableERC20FactoryMetaData.ABI instead.
var MintableERC20FactoryABI = MintableERC20FactoryMetaData.ABI

// MintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MintableERC20FactoryMetaData.Bin instead.
var MintableERC20FactoryBin = MintableERC20FactoryMetaData.Bin

// DeployMintableERC20Factory deploys a new Ethereum contract, binding an instance of MintableERC20Factory to it.
func DeployMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *MintableERC20Factory, error) {
	parsed, err := MintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MintableERC20FactoryBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintableERC20Factory{MintableERC20FactoryCaller: MintableERC20FactoryCaller{contract: contract}, MintableERC20FactoryTransactor: MintableERC20FactoryTransactor{contract: contract}, MintableERC20FactoryFilterer: MintableERC20FactoryFilterer{contract: contract}}, nil
}

// MintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type MintableERC20Factory struct {
	MintableERC20FactoryCaller     // Read-only binding to the contract
	MintableERC20FactoryTransactor // Write-only binding to the contract
	MintableERC20FactoryFilterer   // Log filterer for contract events
}

// MintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintableERC20FactorySession struct {
	Contract     *MintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintableERC20FactoryCallerSession struct {
	Contract *MintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintableERC20FactoryTransactorSession struct {
	Contract     *MintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintableERC20FactoryRaw struct {
	Contract *MintableERC20Factory // Generic contract binding to access the raw methods on
}

// MintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintableERC20FactoryCallerRaw struct {
	Contract *MintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintableERC20FactoryTransactorRaw struct {
	Contract *MintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintableERC20Factory creates a new instance of MintableERC20Factory, bound to a specific deployed contract.
func NewMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*MintableERC20Factory, error) {
	contract, err := bindMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintableERC20Factory{MintableERC20FactoryCaller: MintableERC20FactoryCaller{contract: contract}, MintableERC20FactoryTransactor: MintableERC20FactoryTransactor{contract: contract}, MintableERC20FactoryFilterer: MintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewMintableERC20FactoryCaller creates a new read-only instance of MintableERC20Factory, bound to a specific deployed contract.
func NewMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*MintableERC20FactoryCaller, error) {
	contract, err := bindMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintableERC20FactoryCaller{contract: contract}, nil
}

// NewMintableERC20FactoryTransactor creates a new write-only instance of MintableERC20Factory, bound to a specific deployed contract.
func NewMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MintableERC20FactoryTransactor, error) {
	contract, err := bindMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintableERC20FactoryTransactor{contract: contract}, nil
}

// NewMintableERC20FactoryFilterer creates a new log filterer instance of MintableERC20Factory, bound to a specific deployed contract.
func NewMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MintableERC20FactoryFilterer, error) {
	contract, err := bindMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintableERC20FactoryFilterer{contract: contract}, nil
}

// bindMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableERC20Factory *MintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableERC20Factory.Contract.MintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableERC20Factory *MintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableERC20Factory.Contract.MintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableERC20Factory *MintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableERC20Factory.Contract.MintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableERC20Factory *MintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableERC20Factory *MintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableERC20Factory *MintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MintableERC20Factory *MintableERC20FactoryCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintableERC20Factory.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MintableERC20Factory *MintableERC20FactorySession) BRIDGE() (common.Address, error) {
	return _MintableERC20Factory.Contract.BRIDGE(&_MintableERC20Factory.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MintableERC20Factory *MintableERC20FactoryCallerSession) BRIDGE() (common.Address, error) {
	return _MintableERC20Factory.Contract.BRIDGE(&_MintableERC20Factory.CallOpts)
}

// CreateMintableERC20 is a paid mutator transaction binding the contract method 0x50d38c6c.
//
// Solidity: function createMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MintableERC20Factory *MintableERC20FactoryTransactor) CreateMintableERC20(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MintableERC20Factory.contract.Transact(opts, "createMintableERC20", _remoteToken, _name, _symbol)
}

// CreateMintableERC20 is a paid mutator transaction binding the contract method 0x50d38c6c.
//
// Solidity: function createMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MintableERC20Factory *MintableERC20FactorySession) CreateMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MintableERC20Factory.Contract.CreateMintableERC20(&_MintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateMintableERC20 is a paid mutator transaction binding the contract method 0x50d38c6c.
//
// Solidity: function createMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MintableERC20Factory *MintableERC20FactoryTransactorSession) CreateMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MintableERC20Factory.Contract.CreateMintableERC20(&_MintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// MintableERC20FactoryMintableERC20CreatedIterator is returned from FilterMintableERC20Created and is used to iterate over the raw logs and unpacked data for MintableERC20Created events raised by the MintableERC20Factory contract.
type MintableERC20FactoryMintableERC20CreatedIterator struct {
	Event *MintableERC20FactoryMintableERC20Created // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MintableERC20FactoryMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableERC20FactoryMintableERC20Created)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MintableERC20FactoryMintableERC20Created)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MintableERC20FactoryMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableERC20FactoryMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableERC20FactoryMintableERC20Created represents a MintableERC20Created event raised by the MintableERC20Factory contract.
type MintableERC20FactoryMintableERC20Created struct {
	LocalToken  common.Address
	RemoteToken common.Address
	Deployer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintableERC20Created is a free log retrieval operation binding the contract event 0xd8d73ce8454cb13796d0cbdbbc836fc2251faa966d0ad8a176b933492562e564.
//
// Solidity: event MintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MintableERC20Factory *MintableERC20FactoryFilterer) FilterMintableERC20Created(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address) (*MintableERC20FactoryMintableERC20CreatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _MintableERC20Factory.contract.FilterLogs(opts, "MintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &MintableERC20FactoryMintableERC20CreatedIterator{contract: _MintableERC20Factory.contract, event: "MintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchMintableERC20Created is a free log subscription operation binding the contract event 0xd8d73ce8454cb13796d0cbdbbc836fc2251faa966d0ad8a176b933492562e564.
//
// Solidity: event MintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MintableERC20Factory *MintableERC20FactoryFilterer) WatchMintableERC20Created(opts *bind.WatchOpts, sink chan<- *MintableERC20FactoryMintableERC20Created, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _MintableERC20Factory.contract.WatchLogs(opts, "MintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableERC20FactoryMintableERC20Created)
				if err := _MintableERC20Factory.contract.UnpackLog(event, "MintableERC20Created", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintableERC20Created is a log parse operation binding the contract event 0xd8d73ce8454cb13796d0cbdbbc836fc2251faa966d0ad8a176b933492562e564.
//
// Solidity: event MintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MintableERC20Factory *MintableERC20FactoryFilterer) ParseMintableERC20Created(log types.Log) (*MintableERC20FactoryMintableERC20Created, error) {
	event := new(MintableERC20FactoryMintableERC20Created)
	if err := _MintableERC20Factory.contract.UnpackLog(event, "MintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
