// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irewardrouterv2

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

// IRewardRouterV2MetaData contains all meta data concerning the IRewardRouterV2 contract.
var IRewardRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IRewardRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardRouterV2MetaData.ABI instead.
var IRewardRouterV2ABI = IRewardRouterV2MetaData.ABI

// IRewardRouterV2 is an auto generated Go binding around an Ethereum contract.
type IRewardRouterV2 struct {
	IRewardRouterV2Caller     // Read-only binding to the contract
	IRewardRouterV2Transactor // Write-only binding to the contract
	IRewardRouterV2Filterer   // Log filterer for contract events
}

// IRewardRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardRouterV2Session struct {
	Contract     *IRewardRouterV2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardRouterV2CallerSession struct {
	Contract *IRewardRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IRewardRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardRouterV2TransactorSession struct {
	Contract     *IRewardRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IRewardRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardRouterV2Raw struct {
	Contract *IRewardRouterV2 // Generic contract binding to access the raw methods on
}

// IRewardRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardRouterV2CallerRaw struct {
	Contract *IRewardRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// IRewardRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardRouterV2TransactorRaw struct {
	Contract *IRewardRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardRouterV2 creates a new instance of IRewardRouterV2, bound to a specific deployed contract.
func NewIRewardRouterV2(address common.Address, backend bind.ContractBackend) (*IRewardRouterV2, error) {
	contract, err := bindIRewardRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewardRouterV2{IRewardRouterV2Caller: IRewardRouterV2Caller{contract: contract}, IRewardRouterV2Transactor: IRewardRouterV2Transactor{contract: contract}, IRewardRouterV2Filterer: IRewardRouterV2Filterer{contract: contract}}, nil
}

// NewIRewardRouterV2Caller creates a new read-only instance of IRewardRouterV2, bound to a specific deployed contract.
func NewIRewardRouterV2Caller(address common.Address, caller bind.ContractCaller) (*IRewardRouterV2Caller, error) {
	contract, err := bindIRewardRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardRouterV2Caller{contract: contract}, nil
}

// NewIRewardRouterV2Transactor creates a new write-only instance of IRewardRouterV2, bound to a specific deployed contract.
func NewIRewardRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*IRewardRouterV2Transactor, error) {
	contract, err := bindIRewardRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardRouterV2Transactor{contract: contract}, nil
}

// NewIRewardRouterV2Filterer creates a new log filterer instance of IRewardRouterV2, bound to a specific deployed contract.
func NewIRewardRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*IRewardRouterV2Filterer, error) {
	contract, err := bindIRewardRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardRouterV2Filterer{contract: contract}, nil
}

// bindIRewardRouterV2 binds a generic wrapper to an already deployed contract.
func bindIRewardRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRewardRouterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardRouterV2 *IRewardRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardRouterV2.Contract.IRewardRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardRouterV2 *IRewardRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardRouterV2.Contract.IRewardRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardRouterV2 *IRewardRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardRouterV2.Contract.IRewardRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardRouterV2 *IRewardRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardRouterV2 *IRewardRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardRouterV2 *IRewardRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardRouterV2.Contract.contract.Transact(opts, method, params...)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_IRewardRouterV2 *IRewardRouterV2Caller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRewardRouterV2.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_IRewardRouterV2 *IRewardRouterV2Session) FeeGlpTracker() (common.Address, error) {
	return _IRewardRouterV2.Contract.FeeGlpTracker(&_IRewardRouterV2.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_IRewardRouterV2 *IRewardRouterV2CallerSession) FeeGlpTracker() (common.Address, error) {
	return _IRewardRouterV2.Contract.FeeGlpTracker(&_IRewardRouterV2.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_IRewardRouterV2 *IRewardRouterV2Caller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRewardRouterV2.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_IRewardRouterV2 *IRewardRouterV2Session) StakedGlpTracker() (common.Address, error) {
	return _IRewardRouterV2.Contract.StakedGlpTracker(&_IRewardRouterV2.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_IRewardRouterV2 *IRewardRouterV2CallerSession) StakedGlpTracker() (common.Address, error) {
	return _IRewardRouterV2.Contract.StakedGlpTracker(&_IRewardRouterV2.CallOpts)
}
