// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iyieldtoken

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

// IYieldTokenMetaData contains all meta data concerning the IYieldToken contract.
var IYieldTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IYieldTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IYieldTokenMetaData.ABI instead.
var IYieldTokenABI = IYieldTokenMetaData.ABI

// IYieldToken is an auto generated Go binding around an Ethereum contract.
type IYieldToken struct {
	IYieldTokenCaller     // Read-only binding to the contract
	IYieldTokenTransactor // Write-only binding to the contract
	IYieldTokenFilterer   // Log filterer for contract events
}

// IYieldTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IYieldTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYieldTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IYieldTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYieldTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IYieldTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYieldTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IYieldTokenSession struct {
	Contract     *IYieldToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IYieldTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IYieldTokenCallerSession struct {
	Contract *IYieldTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IYieldTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IYieldTokenTransactorSession struct {
	Contract     *IYieldTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IYieldTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IYieldTokenRaw struct {
	Contract *IYieldToken // Generic contract binding to access the raw methods on
}

// IYieldTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IYieldTokenCallerRaw struct {
	Contract *IYieldTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IYieldTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IYieldTokenTransactorRaw struct {
	Contract *IYieldTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIYieldToken creates a new instance of IYieldToken, bound to a specific deployed contract.
func NewIYieldToken(address common.Address, backend bind.ContractBackend) (*IYieldToken, error) {
	contract, err := bindIYieldToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IYieldToken{IYieldTokenCaller: IYieldTokenCaller{contract: contract}, IYieldTokenTransactor: IYieldTokenTransactor{contract: contract}, IYieldTokenFilterer: IYieldTokenFilterer{contract: contract}}, nil
}

// NewIYieldTokenCaller creates a new read-only instance of IYieldToken, bound to a specific deployed contract.
func NewIYieldTokenCaller(address common.Address, caller bind.ContractCaller) (*IYieldTokenCaller, error) {
	contract, err := bindIYieldToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IYieldTokenCaller{contract: contract}, nil
}

// NewIYieldTokenTransactor creates a new write-only instance of IYieldToken, bound to a specific deployed contract.
func NewIYieldTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IYieldTokenTransactor, error) {
	contract, err := bindIYieldToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IYieldTokenTransactor{contract: contract}, nil
}

// NewIYieldTokenFilterer creates a new log filterer instance of IYieldToken, bound to a specific deployed contract.
func NewIYieldTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IYieldTokenFilterer, error) {
	contract, err := bindIYieldToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IYieldTokenFilterer{contract: contract}, nil
}

// bindIYieldToken binds a generic wrapper to an already deployed contract.
func bindIYieldToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IYieldTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYieldToken *IYieldTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYieldToken.Contract.IYieldTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYieldToken *IYieldTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYieldToken.Contract.IYieldTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYieldToken *IYieldTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYieldToken.Contract.IYieldTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYieldToken *IYieldTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYieldToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYieldToken *IYieldTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYieldToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYieldToken *IYieldTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYieldToken.Contract.contract.Transact(opts, method, params...)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_IYieldToken *IYieldTokenCaller) StakedBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IYieldToken.contract.Call(opts, &out, "stakedBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_IYieldToken *IYieldTokenSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _IYieldToken.Contract.StakedBalance(&_IYieldToken.CallOpts, _account)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_IYieldToken *IYieldTokenCallerSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _IYieldToken.Contract.StakedBalance(&_IYieldToken.CallOpts, _account)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IYieldToken *IYieldTokenCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IYieldToken.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IYieldToken *IYieldTokenSession) TotalStaked() (*big.Int, error) {
	return _IYieldToken.Contract.TotalStaked(&_IYieldToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IYieldToken *IYieldTokenCallerSession) TotalStaked() (*big.Int, error) {
	return _IYieldToken.Contract.TotalStaked(&_IYieldToken.CallOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_IYieldToken *IYieldTokenTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IYieldToken.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_IYieldToken *IYieldTokenSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _IYieldToken.Contract.RemoveAdmin(&_IYieldToken.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_IYieldToken *IYieldTokenTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _IYieldToken.Contract.RemoveAdmin(&_IYieldToken.TransactOpts, _account)
}
