// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igmxtimelock

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

// IGmxTimelockMetaData contains all meta data concerning the IGmxTimelock contract.
var IGmxTimelockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLeverageEnabled\",\"type\":\"bool\"}],\"name\":\"setIsLeverageEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"signalSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGmxTimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use IGmxTimelockMetaData.ABI instead.
var IGmxTimelockABI = IGmxTimelockMetaData.ABI

// IGmxTimelock is an auto generated Go binding around an Ethereum contract.
type IGmxTimelock struct {
	IGmxTimelockCaller     // Read-only binding to the contract
	IGmxTimelockTransactor // Write-only binding to the contract
	IGmxTimelockFilterer   // Log filterer for contract events
}

// IGmxTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGmxTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGmxTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGmxTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGmxTimelockSession struct {
	Contract     *IGmxTimelock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGmxTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGmxTimelockCallerSession struct {
	Contract *IGmxTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGmxTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGmxTimelockTransactorSession struct {
	Contract     *IGmxTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGmxTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGmxTimelockRaw struct {
	Contract *IGmxTimelock // Generic contract binding to access the raw methods on
}

// IGmxTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGmxTimelockCallerRaw struct {
	Contract *IGmxTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// IGmxTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGmxTimelockTransactorRaw struct {
	Contract *IGmxTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGmxTimelock creates a new instance of IGmxTimelock, bound to a specific deployed contract.
func NewIGmxTimelock(address common.Address, backend bind.ContractBackend) (*IGmxTimelock, error) {
	contract, err := bindIGmxTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGmxTimelock{IGmxTimelockCaller: IGmxTimelockCaller{contract: contract}, IGmxTimelockTransactor: IGmxTimelockTransactor{contract: contract}, IGmxTimelockFilterer: IGmxTimelockFilterer{contract: contract}}, nil
}

// NewIGmxTimelockCaller creates a new read-only instance of IGmxTimelock, bound to a specific deployed contract.
func NewIGmxTimelockCaller(address common.Address, caller bind.ContractCaller) (*IGmxTimelockCaller, error) {
	contract, err := bindIGmxTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGmxTimelockCaller{contract: contract}, nil
}

// NewIGmxTimelockTransactor creates a new write-only instance of IGmxTimelock, bound to a specific deployed contract.
func NewIGmxTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*IGmxTimelockTransactor, error) {
	contract, err := bindIGmxTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGmxTimelockTransactor{contract: contract}, nil
}

// NewIGmxTimelockFilterer creates a new log filterer instance of IGmxTimelock, bound to a specific deployed contract.
func NewIGmxTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*IGmxTimelockFilterer, error) {
	contract, err := bindIGmxTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGmxTimelockFilterer{contract: contract}, nil
}

// bindIGmxTimelock binds a generic wrapper to an already deployed contract.
func bindIGmxTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGmxTimelockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGmxTimelock *IGmxTimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGmxTimelock.Contract.IGmxTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGmxTimelock *IGmxTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.IGmxTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGmxTimelock *IGmxTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.IGmxTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGmxTimelock *IGmxTimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGmxTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGmxTimelock *IGmxTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGmxTimelock *IGmxTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.contract.Transact(opts, method, params...)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_IGmxTimelock *IGmxTimelockTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _IGmxTimelock.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_IGmxTimelock *IGmxTimelockSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.SetAdmin(&_IGmxTimelock.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_IGmxTimelock *IGmxTimelockTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.SetAdmin(&_IGmxTimelock.TransactOpts, _admin)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_IGmxTimelock *IGmxTimelockTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _IGmxTimelock.contract.Transact(opts, "setIsLeverageEnabled", _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_IGmxTimelock *IGmxTimelockSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.SetIsLeverageEnabled(&_IGmxTimelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_IGmxTimelock *IGmxTimelockTransactorSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.SetIsLeverageEnabled(&_IGmxTimelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_IGmxTimelock *IGmxTimelockTransactor) SignalSetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _IGmxTimelock.contract.Transact(opts, "signalSetGov", _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_IGmxTimelock *IGmxTimelockSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.SignalSetGov(&_IGmxTimelock.TransactOpts, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_IGmxTimelock *IGmxTimelockTransactorSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _IGmxTimelock.Contract.SignalSetGov(&_IGmxTimelock.TransactOpts, _target, _gov)
}
