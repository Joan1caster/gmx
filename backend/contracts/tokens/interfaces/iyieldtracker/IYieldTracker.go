// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iyieldtracker

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

// IYieldTrackerMetaData contains all meta data concerning the IYieldTracker contract.
var IYieldTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IYieldTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use IYieldTrackerMetaData.ABI instead.
var IYieldTrackerABI = IYieldTrackerMetaData.ABI

// IYieldTracker is an auto generated Go binding around an Ethereum contract.
type IYieldTracker struct {
	IYieldTrackerCaller     // Read-only binding to the contract
	IYieldTrackerTransactor // Write-only binding to the contract
	IYieldTrackerFilterer   // Log filterer for contract events
}

// IYieldTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IYieldTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYieldTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IYieldTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYieldTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IYieldTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYieldTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IYieldTrackerSession struct {
	Contract     *IYieldTracker    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IYieldTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IYieldTrackerCallerSession struct {
	Contract *IYieldTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IYieldTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IYieldTrackerTransactorSession struct {
	Contract     *IYieldTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IYieldTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IYieldTrackerRaw struct {
	Contract *IYieldTracker // Generic contract binding to access the raw methods on
}

// IYieldTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IYieldTrackerCallerRaw struct {
	Contract *IYieldTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// IYieldTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IYieldTrackerTransactorRaw struct {
	Contract *IYieldTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIYieldTracker creates a new instance of IYieldTracker, bound to a specific deployed contract.
func NewIYieldTracker(address common.Address, backend bind.ContractBackend) (*IYieldTracker, error) {
	contract, err := bindIYieldTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IYieldTracker{IYieldTrackerCaller: IYieldTrackerCaller{contract: contract}, IYieldTrackerTransactor: IYieldTrackerTransactor{contract: contract}, IYieldTrackerFilterer: IYieldTrackerFilterer{contract: contract}}, nil
}

// NewIYieldTrackerCaller creates a new read-only instance of IYieldTracker, bound to a specific deployed contract.
func NewIYieldTrackerCaller(address common.Address, caller bind.ContractCaller) (*IYieldTrackerCaller, error) {
	contract, err := bindIYieldTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IYieldTrackerCaller{contract: contract}, nil
}

// NewIYieldTrackerTransactor creates a new write-only instance of IYieldTracker, bound to a specific deployed contract.
func NewIYieldTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*IYieldTrackerTransactor, error) {
	contract, err := bindIYieldTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IYieldTrackerTransactor{contract: contract}, nil
}

// NewIYieldTrackerFilterer creates a new log filterer instance of IYieldTracker, bound to a specific deployed contract.
func NewIYieldTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*IYieldTrackerFilterer, error) {
	contract, err := bindIYieldTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IYieldTrackerFilterer{contract: contract}, nil
}

// bindIYieldTracker binds a generic wrapper to an already deployed contract.
func bindIYieldTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IYieldTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYieldTracker *IYieldTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYieldTracker.Contract.IYieldTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYieldTracker *IYieldTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYieldTracker.Contract.IYieldTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYieldTracker *IYieldTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYieldTracker.Contract.IYieldTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYieldTracker *IYieldTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYieldTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYieldTracker *IYieldTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYieldTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYieldTracker *IYieldTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYieldTracker.Contract.contract.Transact(opts, method, params...)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IYieldTracker *IYieldTrackerCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IYieldTracker.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IYieldTracker *IYieldTrackerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _IYieldTracker.Contract.Claimable(&_IYieldTracker.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IYieldTracker *IYieldTrackerCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _IYieldTracker.Contract.Claimable(&_IYieldTracker.CallOpts, _account)
}

// GetTokensPerInterval is a free data retrieval call binding the contract method 0x2459f51d.
//
// Solidity: function getTokensPerInterval() view returns(uint256)
func (_IYieldTracker *IYieldTrackerCaller) GetTokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IYieldTracker.contract.Call(opts, &out, "getTokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokensPerInterval is a free data retrieval call binding the contract method 0x2459f51d.
//
// Solidity: function getTokensPerInterval() view returns(uint256)
func (_IYieldTracker *IYieldTrackerSession) GetTokensPerInterval() (*big.Int, error) {
	return _IYieldTracker.Contract.GetTokensPerInterval(&_IYieldTracker.CallOpts)
}

// GetTokensPerInterval is a free data retrieval call binding the contract method 0x2459f51d.
//
// Solidity: function getTokensPerInterval() view returns(uint256)
func (_IYieldTracker *IYieldTrackerCallerSession) GetTokensPerInterval() (*big.Int, error) {
	return _IYieldTracker.Contract.GetTokensPerInterval(&_IYieldTracker.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address _account, address _receiver) returns(uint256)
func (_IYieldTracker *IYieldTrackerTransactor) Claim(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IYieldTracker.contract.Transact(opts, "claim", _account, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address _account, address _receiver) returns(uint256)
func (_IYieldTracker *IYieldTrackerSession) Claim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IYieldTracker.Contract.Claim(&_IYieldTracker.TransactOpts, _account, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address _account, address _receiver) returns(uint256)
func (_IYieldTracker *IYieldTrackerTransactorSession) Claim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IYieldTracker.Contract.Claim(&_IYieldTracker.TransactOpts, _account, _receiver)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x5fd61965.
//
// Solidity: function updateRewards(address _account) returns()
func (_IYieldTracker *IYieldTrackerTransactor) UpdateRewards(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IYieldTracker.contract.Transact(opts, "updateRewards", _account)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x5fd61965.
//
// Solidity: function updateRewards(address _account) returns()
func (_IYieldTracker *IYieldTrackerSession) UpdateRewards(_account common.Address) (*types.Transaction, error) {
	return _IYieldTracker.Contract.UpdateRewards(&_IYieldTracker.TransactOpts, _account)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x5fd61965.
//
// Solidity: function updateRewards(address _account) returns()
func (_IYieldTracker *IYieldTrackerTransactorSession) UpdateRewards(_account common.Address) (*types.Transaction, error) {
	return _IYieldTracker.Contract.UpdateRewards(&_IYieldTracker.TransactOpts, _account)
}
