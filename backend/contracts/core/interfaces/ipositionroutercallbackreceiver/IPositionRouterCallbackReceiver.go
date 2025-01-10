// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipositionroutercallbackreceiver

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

// IPositionRouterCallbackReceiverMetaData contains all meta data concerning the IPositionRouterCallbackReceiver contract.
var IPositionRouterCallbackReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isIncrease\",\"type\":\"bool\"}],\"name\":\"gmxPositionCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPositionRouterCallbackReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IPositionRouterCallbackReceiverMetaData.ABI instead.
var IPositionRouterCallbackReceiverABI = IPositionRouterCallbackReceiverMetaData.ABI

// IPositionRouterCallbackReceiver is an auto generated Go binding around an Ethereum contract.
type IPositionRouterCallbackReceiver struct {
	IPositionRouterCallbackReceiverCaller     // Read-only binding to the contract
	IPositionRouterCallbackReceiverTransactor // Write-only binding to the contract
	IPositionRouterCallbackReceiverFilterer   // Log filterer for contract events
}

// IPositionRouterCallbackReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPositionRouterCallbackReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPositionRouterCallbackReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPositionRouterCallbackReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPositionRouterCallbackReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPositionRouterCallbackReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPositionRouterCallbackReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPositionRouterCallbackReceiverSession struct {
	Contract     *IPositionRouterCallbackReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IPositionRouterCallbackReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPositionRouterCallbackReceiverCallerSession struct {
	Contract *IPositionRouterCallbackReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// IPositionRouterCallbackReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPositionRouterCallbackReceiverTransactorSession struct {
	Contract     *IPositionRouterCallbackReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// IPositionRouterCallbackReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPositionRouterCallbackReceiverRaw struct {
	Contract *IPositionRouterCallbackReceiver // Generic contract binding to access the raw methods on
}

// IPositionRouterCallbackReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPositionRouterCallbackReceiverCallerRaw struct {
	Contract *IPositionRouterCallbackReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IPositionRouterCallbackReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPositionRouterCallbackReceiverTransactorRaw struct {
	Contract *IPositionRouterCallbackReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPositionRouterCallbackReceiver creates a new instance of IPositionRouterCallbackReceiver, bound to a specific deployed contract.
func NewIPositionRouterCallbackReceiver(address common.Address, backend bind.ContractBackend) (*IPositionRouterCallbackReceiver, error) {
	contract, err := bindIPositionRouterCallbackReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPositionRouterCallbackReceiver{IPositionRouterCallbackReceiverCaller: IPositionRouterCallbackReceiverCaller{contract: contract}, IPositionRouterCallbackReceiverTransactor: IPositionRouterCallbackReceiverTransactor{contract: contract}, IPositionRouterCallbackReceiverFilterer: IPositionRouterCallbackReceiverFilterer{contract: contract}}, nil
}

// NewIPositionRouterCallbackReceiverCaller creates a new read-only instance of IPositionRouterCallbackReceiver, bound to a specific deployed contract.
func NewIPositionRouterCallbackReceiverCaller(address common.Address, caller bind.ContractCaller) (*IPositionRouterCallbackReceiverCaller, error) {
	contract, err := bindIPositionRouterCallbackReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPositionRouterCallbackReceiverCaller{contract: contract}, nil
}

// NewIPositionRouterCallbackReceiverTransactor creates a new write-only instance of IPositionRouterCallbackReceiver, bound to a specific deployed contract.
func NewIPositionRouterCallbackReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IPositionRouterCallbackReceiverTransactor, error) {
	contract, err := bindIPositionRouterCallbackReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPositionRouterCallbackReceiverTransactor{contract: contract}, nil
}

// NewIPositionRouterCallbackReceiverFilterer creates a new log filterer instance of IPositionRouterCallbackReceiver, bound to a specific deployed contract.
func NewIPositionRouterCallbackReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IPositionRouterCallbackReceiverFilterer, error) {
	contract, err := bindIPositionRouterCallbackReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPositionRouterCallbackReceiverFilterer{contract: contract}, nil
}

// bindIPositionRouterCallbackReceiver binds a generic wrapper to an already deployed contract.
func bindIPositionRouterCallbackReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPositionRouterCallbackReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPositionRouterCallbackReceiver.Contract.IPositionRouterCallbackReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPositionRouterCallbackReceiver.Contract.IPositionRouterCallbackReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPositionRouterCallbackReceiver.Contract.IPositionRouterCallbackReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPositionRouterCallbackReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPositionRouterCallbackReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPositionRouterCallbackReceiver.Contract.contract.Transact(opts, method, params...)
}

// GmxPositionCallback is a paid mutator transaction binding the contract method 0xedf3daec.
//
// Solidity: function gmxPositionCallback(bytes32 positionKey, bool isExecuted, bool isIncrease) returns()
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverTransactor) GmxPositionCallback(opts *bind.TransactOpts, positionKey [32]byte, isExecuted bool, isIncrease bool) (*types.Transaction, error) {
	return _IPositionRouterCallbackReceiver.contract.Transact(opts, "gmxPositionCallback", positionKey, isExecuted, isIncrease)
}

// GmxPositionCallback is a paid mutator transaction binding the contract method 0xedf3daec.
//
// Solidity: function gmxPositionCallback(bytes32 positionKey, bool isExecuted, bool isIncrease) returns()
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverSession) GmxPositionCallback(positionKey [32]byte, isExecuted bool, isIncrease bool) (*types.Transaction, error) {
	return _IPositionRouterCallbackReceiver.Contract.GmxPositionCallback(&_IPositionRouterCallbackReceiver.TransactOpts, positionKey, isExecuted, isIncrease)
}

// GmxPositionCallback is a paid mutator transaction binding the contract method 0xedf3daec.
//
// Solidity: function gmxPositionCallback(bytes32 positionKey, bool isExecuted, bool isIncrease) returns()
func (_IPositionRouterCallbackReceiver *IPositionRouterCallbackReceiverTransactorSession) GmxPositionCallback(positionKey [32]byte, isExecuted bool, isIncrease bool) (*types.Transaction, error) {
	return _IPositionRouterCallbackReceiver.Contract.GmxPositionCallback(&_IPositionRouterCallbackReceiver.TransactOpts, positionKey, isExecuted, isIncrease)
}
