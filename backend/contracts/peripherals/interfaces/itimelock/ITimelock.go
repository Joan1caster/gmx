// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itimelock

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

// ITimelockMetaData contains all meta data concerning the ITimelock contract.
var ITimelockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"disableLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"enableLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"}],\"name\":\"requestGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLeverageEnabled\",\"type\":\"bool\"}],\"name\":\"setIsLeverageEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"signalSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use ITimelockMetaData.ABI instead.
var ITimelockABI = ITimelockMetaData.ABI

// ITimelock is an auto generated Go binding around an Ethereum contract.
type ITimelock struct {
	ITimelockCaller     // Read-only binding to the contract
	ITimelockTransactor // Write-only binding to the contract
	ITimelockFilterer   // Log filterer for contract events
}

// ITimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITimelockSession struct {
	Contract     *ITimelock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITimelockCallerSession struct {
	Contract *ITimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ITimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITimelockTransactorSession struct {
	Contract     *ITimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITimelockRaw struct {
	Contract *ITimelock // Generic contract binding to access the raw methods on
}

// ITimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITimelockCallerRaw struct {
	Contract *ITimelockCaller // Generic read-only contract binding to access the raw methods on
}

// ITimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITimelockTransactorRaw struct {
	Contract *ITimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITimelock creates a new instance of ITimelock, bound to a specific deployed contract.
func NewITimelock(address common.Address, backend bind.ContractBackend) (*ITimelock, error) {
	contract, err := bindITimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITimelock{ITimelockCaller: ITimelockCaller{contract: contract}, ITimelockTransactor: ITimelockTransactor{contract: contract}, ITimelockFilterer: ITimelockFilterer{contract: contract}}, nil
}

// NewITimelockCaller creates a new read-only instance of ITimelock, bound to a specific deployed contract.
func NewITimelockCaller(address common.Address, caller bind.ContractCaller) (*ITimelockCaller, error) {
	contract, err := bindITimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITimelockCaller{contract: contract}, nil
}

// NewITimelockTransactor creates a new write-only instance of ITimelock, bound to a specific deployed contract.
func NewITimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*ITimelockTransactor, error) {
	contract, err := bindITimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITimelockTransactor{contract: contract}, nil
}

// NewITimelockFilterer creates a new log filterer instance of ITimelock, bound to a specific deployed contract.
func NewITimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*ITimelockFilterer, error) {
	contract, err := bindITimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITimelockFilterer{contract: contract}, nil
}

// bindITimelock binds a generic wrapper to an already deployed contract.
func bindITimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITimelockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimelock *ITimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimelock.Contract.ITimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimelock *ITimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimelock.Contract.ITimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimelock *ITimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimelock.Contract.ITimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimelock *ITimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimelock *ITimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimelock *ITimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimelock.Contract.contract.Transact(opts, method, params...)
}

// DisableLeverage is a paid mutator transaction binding the contract method 0xd3c87bbb.
//
// Solidity: function disableLeverage(address _vault) returns()
func (_ITimelock *ITimelockTransactor) DisableLeverage(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "disableLeverage", _vault)
}

// DisableLeverage is a paid mutator transaction binding the contract method 0xd3c87bbb.
//
// Solidity: function disableLeverage(address _vault) returns()
func (_ITimelock *ITimelockSession) DisableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.DisableLeverage(&_ITimelock.TransactOpts, _vault)
}

// DisableLeverage is a paid mutator transaction binding the contract method 0xd3c87bbb.
//
// Solidity: function disableLeverage(address _vault) returns()
func (_ITimelock *ITimelockTransactorSession) DisableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.DisableLeverage(&_ITimelock.TransactOpts, _vault)
}

// EnableLeverage is a paid mutator transaction binding the contract method 0x6d63c1d0.
//
// Solidity: function enableLeverage(address _vault) returns()
func (_ITimelock *ITimelockTransactor) EnableLeverage(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "enableLeverage", _vault)
}

// EnableLeverage is a paid mutator transaction binding the contract method 0x6d63c1d0.
//
// Solidity: function enableLeverage(address _vault) returns()
func (_ITimelock *ITimelockSession) EnableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.EnableLeverage(&_ITimelock.TransactOpts, _vault)
}

// EnableLeverage is a paid mutator transaction binding the contract method 0x6d63c1d0.
//
// Solidity: function enableLeverage(address _vault) returns()
func (_ITimelock *ITimelockTransactorSession) EnableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.EnableLeverage(&_ITimelock.TransactOpts, _vault)
}

// MarginFeeBasisPoints is a paid mutator transaction binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() returns(uint256)
func (_ITimelock *ITimelockTransactor) MarginFeeBasisPoints(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "marginFeeBasisPoints")
}

// MarginFeeBasisPoints is a paid mutator transaction binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() returns(uint256)
func (_ITimelock *ITimelockSession) MarginFeeBasisPoints() (*types.Transaction, error) {
	return _ITimelock.Contract.MarginFeeBasisPoints(&_ITimelock.TransactOpts)
}

// MarginFeeBasisPoints is a paid mutator transaction binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() returns(uint256)
func (_ITimelock *ITimelockTransactorSession) MarginFeeBasisPoints() (*types.Transaction, error) {
	return _ITimelock.Contract.MarginFeeBasisPoints(&_ITimelock.TransactOpts)
}

// RequestGov is a paid mutator transaction binding the contract method 0xde8616ff.
//
// Solidity: function requestGov(address[] _targets) returns()
func (_ITimelock *ITimelockTransactor) RequestGov(opts *bind.TransactOpts, _targets []common.Address) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "requestGov", _targets)
}

// RequestGov is a paid mutator transaction binding the contract method 0xde8616ff.
//
// Solidity: function requestGov(address[] _targets) returns()
func (_ITimelock *ITimelockSession) RequestGov(_targets []common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.RequestGov(&_ITimelock.TransactOpts, _targets)
}

// RequestGov is a paid mutator transaction binding the contract method 0xde8616ff.
//
// Solidity: function requestGov(address[] _targets) returns()
func (_ITimelock *ITimelockTransactorSession) RequestGov(_targets []common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.RequestGov(&_ITimelock.TransactOpts, _targets)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_ITimelock *ITimelockTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_ITimelock *ITimelockSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.SetAdmin(&_ITimelock.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_ITimelock *ITimelockTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.SetAdmin(&_ITimelock.TransactOpts, _admin)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _target) returns()
func (_ITimelock *ITimelockTransactor) SetGov(opts *bind.TransactOpts, _target common.Address) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "setGov", _target)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _target) returns()
func (_ITimelock *ITimelockSession) SetGov(_target common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.SetGov(&_ITimelock.TransactOpts, _target)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _target) returns()
func (_ITimelock *ITimelockTransactorSession) SetGov(_target common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.SetGov(&_ITimelock.TransactOpts, _target)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_ITimelock *ITimelockTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "setIsLeverageEnabled", _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_ITimelock *ITimelockSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _ITimelock.Contract.SetIsLeverageEnabled(&_ITimelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_ITimelock *ITimelockTransactorSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _ITimelock.Contract.SetIsLeverageEnabled(&_ITimelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_ITimelock *ITimelockTransactor) SignalSetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ITimelock.contract.Transact(opts, "signalSetGov", _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_ITimelock *ITimelockSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.SignalSetGov(&_ITimelock.TransactOpts, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_ITimelock *ITimelockTransactorSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ITimelock.Contract.SignalSetGov(&_ITimelock.TransactOpts, _target, _gov)
}
