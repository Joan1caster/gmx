// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockgovrequester

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

// MockGovRequesterMetaData contains all meta data concerning the MockGovRequester contract.
var MockGovRequesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"afterGovGranted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MockGovRequesterABI is the input ABI used to generate the binding from.
// Deprecated: Use MockGovRequesterMetaData.ABI instead.
var MockGovRequesterABI = MockGovRequesterMetaData.ABI

// MockGovRequester is an auto generated Go binding around an Ethereum contract.
type MockGovRequester struct {
	MockGovRequesterCaller     // Read-only binding to the contract
	MockGovRequesterTransactor // Write-only binding to the contract
	MockGovRequesterFilterer   // Log filterer for contract events
}

// MockGovRequesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockGovRequesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGovRequesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockGovRequesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGovRequesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockGovRequesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGovRequesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockGovRequesterSession struct {
	Contract     *MockGovRequester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockGovRequesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockGovRequesterCallerSession struct {
	Contract *MockGovRequesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MockGovRequesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockGovRequesterTransactorSession struct {
	Contract     *MockGovRequesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MockGovRequesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockGovRequesterRaw struct {
	Contract *MockGovRequester // Generic contract binding to access the raw methods on
}

// MockGovRequesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockGovRequesterCallerRaw struct {
	Contract *MockGovRequesterCaller // Generic read-only contract binding to access the raw methods on
}

// MockGovRequesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockGovRequesterTransactorRaw struct {
	Contract *MockGovRequesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockGovRequester creates a new instance of MockGovRequester, bound to a specific deployed contract.
func NewMockGovRequester(address common.Address, backend bind.ContractBackend) (*MockGovRequester, error) {
	contract, err := bindMockGovRequester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockGovRequester{MockGovRequesterCaller: MockGovRequesterCaller{contract: contract}, MockGovRequesterTransactor: MockGovRequesterTransactor{contract: contract}, MockGovRequesterFilterer: MockGovRequesterFilterer{contract: contract}}, nil
}

// NewMockGovRequesterCaller creates a new read-only instance of MockGovRequester, bound to a specific deployed contract.
func NewMockGovRequesterCaller(address common.Address, caller bind.ContractCaller) (*MockGovRequesterCaller, error) {
	contract, err := bindMockGovRequester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockGovRequesterCaller{contract: contract}, nil
}

// NewMockGovRequesterTransactor creates a new write-only instance of MockGovRequester, bound to a specific deployed contract.
func NewMockGovRequesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockGovRequesterTransactor, error) {
	contract, err := bindMockGovRequester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockGovRequesterTransactor{contract: contract}, nil
}

// NewMockGovRequesterFilterer creates a new log filterer instance of MockGovRequester, bound to a specific deployed contract.
func NewMockGovRequesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockGovRequesterFilterer, error) {
	contract, err := bindMockGovRequester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockGovRequesterFilterer{contract: contract}, nil
}

// bindMockGovRequester binds a generic wrapper to an already deployed contract.
func bindMockGovRequester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockGovRequesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGovRequester *MockGovRequesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGovRequester.Contract.MockGovRequesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGovRequester *MockGovRequesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGovRequester.Contract.MockGovRequesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGovRequester *MockGovRequesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGovRequester.Contract.MockGovRequesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGovRequester *MockGovRequesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGovRequester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGovRequester *MockGovRequesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGovRequester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGovRequester *MockGovRequesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGovRequester.Contract.contract.Transact(opts, method, params...)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_MockGovRequester *MockGovRequesterTransactor) AfterGovGranted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGovRequester.contract.Transact(opts, "afterGovGranted")
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_MockGovRequester *MockGovRequesterSession) AfterGovGranted() (*types.Transaction, error) {
	return _MockGovRequester.Contract.AfterGovGranted(&_MockGovRequester.TransactOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_MockGovRequester *MockGovRequesterTransactorSession) AfterGovGranted() (*types.Transaction, error) {
	return _MockGovRequester.Contract.AfterGovGranted(&_MockGovRequester.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x047701e4.
//
// Solidity: function migrate(address timelock, address[] targets) returns()
func (_MockGovRequester *MockGovRequesterTransactor) Migrate(opts *bind.TransactOpts, timelock common.Address, targets []common.Address) (*types.Transaction, error) {
	return _MockGovRequester.contract.Transact(opts, "migrate", timelock, targets)
}

// Migrate is a paid mutator transaction binding the contract method 0x047701e4.
//
// Solidity: function migrate(address timelock, address[] targets) returns()
func (_MockGovRequester *MockGovRequesterSession) Migrate(timelock common.Address, targets []common.Address) (*types.Transaction, error) {
	return _MockGovRequester.Contract.Migrate(&_MockGovRequester.TransactOpts, timelock, targets)
}

// Migrate is a paid mutator transaction binding the contract method 0x047701e4.
//
// Solidity: function migrate(address timelock, address[] targets) returns()
func (_MockGovRequester *MockGovRequesterTransactorSession) Migrate(timelock common.Address, targets []common.Address) (*types.Transaction, error) {
	return _MockGovRequester.Contract.Migrate(&_MockGovRequester.TransactOpts, timelock, targets)
}
