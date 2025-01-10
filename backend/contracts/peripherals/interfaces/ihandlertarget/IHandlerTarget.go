// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ihandlertarget

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

// IHandlerTargetMetaData contains all meta data concerning the IHandlerTarget contract.
var IHandlerTargetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isDepositToken\",\"type\":\"bool\"}],\"name\":\"setDepositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IHandlerTargetABI is the input ABI used to generate the binding from.
// Deprecated: Use IHandlerTargetMetaData.ABI instead.
var IHandlerTargetABI = IHandlerTargetMetaData.ABI

// IHandlerTarget is an auto generated Go binding around an Ethereum contract.
type IHandlerTarget struct {
	IHandlerTargetCaller     // Read-only binding to the contract
	IHandlerTargetTransactor // Write-only binding to the contract
	IHandlerTargetFilterer   // Log filterer for contract events
}

// IHandlerTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHandlerTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHandlerTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHandlerTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHandlerTargetSession struct {
	Contract     *IHandlerTarget   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHandlerTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHandlerTargetCallerSession struct {
	Contract *IHandlerTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IHandlerTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHandlerTargetTransactorSession struct {
	Contract     *IHandlerTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IHandlerTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHandlerTargetRaw struct {
	Contract *IHandlerTarget // Generic contract binding to access the raw methods on
}

// IHandlerTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHandlerTargetCallerRaw struct {
	Contract *IHandlerTargetCaller // Generic read-only contract binding to access the raw methods on
}

// IHandlerTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHandlerTargetTransactorRaw struct {
	Contract *IHandlerTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHandlerTarget creates a new instance of IHandlerTarget, bound to a specific deployed contract.
func NewIHandlerTarget(address common.Address, backend bind.ContractBackend) (*IHandlerTarget, error) {
	contract, err := bindIHandlerTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHandlerTarget{IHandlerTargetCaller: IHandlerTargetCaller{contract: contract}, IHandlerTargetTransactor: IHandlerTargetTransactor{contract: contract}, IHandlerTargetFilterer: IHandlerTargetFilterer{contract: contract}}, nil
}

// NewIHandlerTargetCaller creates a new read-only instance of IHandlerTarget, bound to a specific deployed contract.
func NewIHandlerTargetCaller(address common.Address, caller bind.ContractCaller) (*IHandlerTargetCaller, error) {
	contract, err := bindIHandlerTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHandlerTargetCaller{contract: contract}, nil
}

// NewIHandlerTargetTransactor creates a new write-only instance of IHandlerTarget, bound to a specific deployed contract.
func NewIHandlerTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*IHandlerTargetTransactor, error) {
	contract, err := bindIHandlerTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHandlerTargetTransactor{contract: contract}, nil
}

// NewIHandlerTargetFilterer creates a new log filterer instance of IHandlerTarget, bound to a specific deployed contract.
func NewIHandlerTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*IHandlerTargetFilterer, error) {
	contract, err := bindIHandlerTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHandlerTargetFilterer{contract: contract}, nil
}

// bindIHandlerTarget binds a generic wrapper to an already deployed contract.
func bindIHandlerTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHandlerTargetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHandlerTarget *IHandlerTargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHandlerTarget.Contract.IHandlerTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHandlerTarget *IHandlerTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.IHandlerTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHandlerTarget *IHandlerTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.IHandlerTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHandlerTarget *IHandlerTargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHandlerTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHandlerTarget *IHandlerTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHandlerTarget *IHandlerTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.contract.Transact(opts, method, params...)
}

// IsHandler is a paid mutator transaction binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address _account) returns(bool)
func (_IHandlerTarget *IHandlerTargetTransactor) IsHandler(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IHandlerTarget.contract.Transact(opts, "isHandler", _account)
}

// IsHandler is a paid mutator transaction binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address _account) returns(bool)
func (_IHandlerTarget *IHandlerTargetSession) IsHandler(_account common.Address) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.IsHandler(&_IHandlerTarget.TransactOpts, _account)
}

// IsHandler is a paid mutator transaction binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address _account) returns(bool)
func (_IHandlerTarget *IHandlerTargetTransactorSession) IsHandler(_account common.Address) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.IsHandler(&_IHandlerTarget.TransactOpts, _account)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_IHandlerTarget *IHandlerTargetTransactor) SetDepositToken(opts *bind.TransactOpts, _depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _IHandlerTarget.contract.Transact(opts, "setDepositToken", _depositToken, _isDepositToken)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_IHandlerTarget *IHandlerTargetSession) SetDepositToken(_depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.SetDepositToken(&_IHandlerTarget.TransactOpts, _depositToken, _isDepositToken)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_IHandlerTarget *IHandlerTargetTransactorSession) SetDepositToken(_depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.SetDepositToken(&_IHandlerTarget.TransactOpts, _depositToken, _isDepositToken)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_IHandlerTarget *IHandlerTargetTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _IHandlerTarget.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_IHandlerTarget *IHandlerTargetSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.SetHandler(&_IHandlerTarget.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_IHandlerTarget *IHandlerTargetTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _IHandlerTarget.Contract.SetHandler(&_IHandlerTarget.TransactOpts, _handler, _isActive)
}
