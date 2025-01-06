// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igovrequester

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

// IGovRequesterMetaData contains all meta data concerning the IGovRequester contract.
var IGovRequesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"afterGovGranted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGovRequesterABI is the input ABI used to generate the binding from.
// Deprecated: Use IGovRequesterMetaData.ABI instead.
var IGovRequesterABI = IGovRequesterMetaData.ABI

// IGovRequester is an auto generated Go binding around an Ethereum contract.
type IGovRequester struct {
	IGovRequesterCaller     // Read-only binding to the contract
	IGovRequesterTransactor // Write-only binding to the contract
	IGovRequesterFilterer   // Log filterer for contract events
}

// IGovRequesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGovRequesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovRequesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGovRequesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovRequesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGovRequesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovRequesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGovRequesterSession struct {
	Contract     *IGovRequester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovRequesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGovRequesterCallerSession struct {
	Contract *IGovRequesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IGovRequesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGovRequesterTransactorSession struct {
	Contract     *IGovRequesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IGovRequesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGovRequesterRaw struct {
	Contract *IGovRequester // Generic contract binding to access the raw methods on
}

// IGovRequesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGovRequesterCallerRaw struct {
	Contract *IGovRequesterCaller // Generic read-only contract binding to access the raw methods on
}

// IGovRequesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGovRequesterTransactorRaw struct {
	Contract *IGovRequesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGovRequester creates a new instance of IGovRequester, bound to a specific deployed contract.
func NewIGovRequester(address common.Address, backend bind.ContractBackend) (*IGovRequester, error) {
	contract, err := bindIGovRequester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGovRequester{IGovRequesterCaller: IGovRequesterCaller{contract: contract}, IGovRequesterTransactor: IGovRequesterTransactor{contract: contract}, IGovRequesterFilterer: IGovRequesterFilterer{contract: contract}}, nil
}

// NewIGovRequesterCaller creates a new read-only instance of IGovRequester, bound to a specific deployed contract.
func NewIGovRequesterCaller(address common.Address, caller bind.ContractCaller) (*IGovRequesterCaller, error) {
	contract, err := bindIGovRequester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGovRequesterCaller{contract: contract}, nil
}

// NewIGovRequesterTransactor creates a new write-only instance of IGovRequester, bound to a specific deployed contract.
func NewIGovRequesterTransactor(address common.Address, transactor bind.ContractTransactor) (*IGovRequesterTransactor, error) {
	contract, err := bindIGovRequester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGovRequesterTransactor{contract: contract}, nil
}

// NewIGovRequesterFilterer creates a new log filterer instance of IGovRequester, bound to a specific deployed contract.
func NewIGovRequesterFilterer(address common.Address, filterer bind.ContractFilterer) (*IGovRequesterFilterer, error) {
	contract, err := bindIGovRequester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGovRequesterFilterer{contract: contract}, nil
}

// bindIGovRequester binds a generic wrapper to an already deployed contract.
func bindIGovRequester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGovRequesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovRequester *IGovRequesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGovRequester.Contract.IGovRequesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovRequester *IGovRequesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovRequester.Contract.IGovRequesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovRequester *IGovRequesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovRequester.Contract.IGovRequesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovRequester *IGovRequesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGovRequester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovRequester *IGovRequesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovRequester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovRequester *IGovRequesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovRequester.Contract.contract.Transact(opts, method, params...)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_IGovRequester *IGovRequesterTransactor) AfterGovGranted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovRequester.contract.Transact(opts, "afterGovGranted")
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_IGovRequester *IGovRequesterSession) AfterGovGranted() (*types.Transaction, error) {
	return _IGovRequester.Contract.AfterGovGranted(&_IGovRequester.TransactOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_IGovRequester *IGovRequesterTransactorSession) AfterGovGranted() (*types.Transaction, error) {
	return _IGovRequester.Contract.AfterGovGranted(&_IGovRequester.TransactOpts)
}
