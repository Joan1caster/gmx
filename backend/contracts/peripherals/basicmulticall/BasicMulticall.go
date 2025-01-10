// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basicmulticall

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

// BasicMulticallMetaData contains all meta data concerning the BasicMulticall contract.
var BasicMulticallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BasicMulticallABI is the input ABI used to generate the binding from.
// Deprecated: Use BasicMulticallMetaData.ABI instead.
var BasicMulticallABI = BasicMulticallMetaData.ABI

// BasicMulticall is an auto generated Go binding around an Ethereum contract.
type BasicMulticall struct {
	BasicMulticallCaller     // Read-only binding to the contract
	BasicMulticallTransactor // Write-only binding to the contract
	BasicMulticallFilterer   // Log filterer for contract events
}

// BasicMulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicMulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicMulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicMulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicMulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicMulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicMulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicMulticallSession struct {
	Contract     *BasicMulticall   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicMulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicMulticallCallerSession struct {
	Contract *BasicMulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BasicMulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicMulticallTransactorSession struct {
	Contract     *BasicMulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BasicMulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicMulticallRaw struct {
	Contract *BasicMulticall // Generic contract binding to access the raw methods on
}

// BasicMulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicMulticallCallerRaw struct {
	Contract *BasicMulticallCaller // Generic read-only contract binding to access the raw methods on
}

// BasicMulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicMulticallTransactorRaw struct {
	Contract *BasicMulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicMulticall creates a new instance of BasicMulticall, bound to a specific deployed contract.
func NewBasicMulticall(address common.Address, backend bind.ContractBackend) (*BasicMulticall, error) {
	contract, err := bindBasicMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicMulticall{BasicMulticallCaller: BasicMulticallCaller{contract: contract}, BasicMulticallTransactor: BasicMulticallTransactor{contract: contract}, BasicMulticallFilterer: BasicMulticallFilterer{contract: contract}}, nil
}

// NewBasicMulticallCaller creates a new read-only instance of BasicMulticall, bound to a specific deployed contract.
func NewBasicMulticallCaller(address common.Address, caller bind.ContractCaller) (*BasicMulticallCaller, error) {
	contract, err := bindBasicMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicMulticallCaller{contract: contract}, nil
}

// NewBasicMulticallTransactor creates a new write-only instance of BasicMulticall, bound to a specific deployed contract.
func NewBasicMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicMulticallTransactor, error) {
	contract, err := bindBasicMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicMulticallTransactor{contract: contract}, nil
}

// NewBasicMulticallFilterer creates a new log filterer instance of BasicMulticall, bound to a specific deployed contract.
func NewBasicMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicMulticallFilterer, error) {
	contract, err := bindBasicMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicMulticallFilterer{contract: contract}, nil
}

// bindBasicMulticall binds a generic wrapper to an already deployed contract.
func bindBasicMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasicMulticallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicMulticall *BasicMulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicMulticall.Contract.BasicMulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicMulticall *BasicMulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicMulticall.Contract.BasicMulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicMulticall *BasicMulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicMulticall.Contract.BasicMulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicMulticall *BasicMulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicMulticall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicMulticall *BasicMulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicMulticall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicMulticall *BasicMulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicMulticall.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BasicMulticall *BasicMulticallTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _BasicMulticall.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BasicMulticall *BasicMulticallSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BasicMulticall.Contract.Multicall(&_BasicMulticall.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BasicMulticall *BasicMulticallTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BasicMulticall.Contract.Multicall(&_BasicMulticall.TransactOpts, data)
}
