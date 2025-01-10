// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iexternalhandler

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

// IExternalHandlerMetaData contains all meta data concerning the IExternalHandler contract.
var IExternalHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataList\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"refundTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"refundReceivers\",\"type\":\"address[]\"}],\"name\":\"makeExternalCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IExternalHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IExternalHandlerMetaData.ABI instead.
var IExternalHandlerABI = IExternalHandlerMetaData.ABI

// IExternalHandler is an auto generated Go binding around an Ethereum contract.
type IExternalHandler struct {
	IExternalHandlerCaller     // Read-only binding to the contract
	IExternalHandlerTransactor // Write-only binding to the contract
	IExternalHandlerFilterer   // Log filterer for contract events
}

// IExternalHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExternalHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExternalHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExternalHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExternalHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExternalHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExternalHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExternalHandlerSession struct {
	Contract     *IExternalHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IExternalHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExternalHandlerCallerSession struct {
	Contract *IExternalHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IExternalHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExternalHandlerTransactorSession struct {
	Contract     *IExternalHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IExternalHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExternalHandlerRaw struct {
	Contract *IExternalHandler // Generic contract binding to access the raw methods on
}

// IExternalHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExternalHandlerCallerRaw struct {
	Contract *IExternalHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IExternalHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExternalHandlerTransactorRaw struct {
	Contract *IExternalHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExternalHandler creates a new instance of IExternalHandler, bound to a specific deployed contract.
func NewIExternalHandler(address common.Address, backend bind.ContractBackend) (*IExternalHandler, error) {
	contract, err := bindIExternalHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExternalHandler{IExternalHandlerCaller: IExternalHandlerCaller{contract: contract}, IExternalHandlerTransactor: IExternalHandlerTransactor{contract: contract}, IExternalHandlerFilterer: IExternalHandlerFilterer{contract: contract}}, nil
}

// NewIExternalHandlerCaller creates a new read-only instance of IExternalHandler, bound to a specific deployed contract.
func NewIExternalHandlerCaller(address common.Address, caller bind.ContractCaller) (*IExternalHandlerCaller, error) {
	contract, err := bindIExternalHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExternalHandlerCaller{contract: contract}, nil
}

// NewIExternalHandlerTransactor creates a new write-only instance of IExternalHandler, bound to a specific deployed contract.
func NewIExternalHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IExternalHandlerTransactor, error) {
	contract, err := bindIExternalHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExternalHandlerTransactor{contract: contract}, nil
}

// NewIExternalHandlerFilterer creates a new log filterer instance of IExternalHandler, bound to a specific deployed contract.
func NewIExternalHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IExternalHandlerFilterer, error) {
	contract, err := bindIExternalHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExternalHandlerFilterer{contract: contract}, nil
}

// bindIExternalHandler binds a generic wrapper to an already deployed contract.
func bindIExternalHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExternalHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExternalHandler *IExternalHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExternalHandler.Contract.IExternalHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExternalHandler *IExternalHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExternalHandler.Contract.IExternalHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExternalHandler *IExternalHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExternalHandler.Contract.IExternalHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExternalHandler *IExternalHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExternalHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExternalHandler *IExternalHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExternalHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExternalHandler *IExternalHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExternalHandler.Contract.contract.Transact(opts, method, params...)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] targets, bytes[] dataList, address[] refundTokens, address[] refundReceivers) returns()
func (_IExternalHandler *IExternalHandlerTransactor) MakeExternalCalls(opts *bind.TransactOpts, targets []common.Address, dataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _IExternalHandler.contract.Transact(opts, "makeExternalCalls", targets, dataList, refundTokens, refundReceivers)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] targets, bytes[] dataList, address[] refundTokens, address[] refundReceivers) returns()
func (_IExternalHandler *IExternalHandlerSession) MakeExternalCalls(targets []common.Address, dataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _IExternalHandler.Contract.MakeExternalCalls(&_IExternalHandler.TransactOpts, targets, dataList, refundTokens, refundReceivers)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] targets, bytes[] dataList, address[] refundTokens, address[] refundReceivers) returns()
func (_IExternalHandler *IExternalHandlerTransactorSession) MakeExternalCalls(targets []common.Address, dataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _IExternalHandler.Contract.MakeExternalCalls(&_IExternalHandler.TransactOpts, targets, dataList, refundTokens, refundReceivers)
}
