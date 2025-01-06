// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itimelocktarget

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

// ITimelockTargetMetaData contains all meta data concerning the ITimelockTarget contract.
var ITimelockTargetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITimelockTargetABI is the input ABI used to generate the binding from.
// Deprecated: Use ITimelockTargetMetaData.ABI instead.
var ITimelockTargetABI = ITimelockTargetMetaData.ABI

// ITimelockTarget is an auto generated Go binding around an Ethereum contract.
type ITimelockTarget struct {
	ITimelockTargetCaller     // Read-only binding to the contract
	ITimelockTargetTransactor // Write-only binding to the contract
	ITimelockTargetFilterer   // Log filterer for contract events
}

// ITimelockTargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITimelockTargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockTargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITimelockTargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockTargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITimelockTargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockTargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITimelockTargetSession struct {
	Contract     *ITimelockTarget  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITimelockTargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITimelockTargetCallerSession struct {
	Contract *ITimelockTargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ITimelockTargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITimelockTargetTransactorSession struct {
	Contract     *ITimelockTargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ITimelockTargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITimelockTargetRaw struct {
	Contract *ITimelockTarget // Generic contract binding to access the raw methods on
}

// ITimelockTargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITimelockTargetCallerRaw struct {
	Contract *ITimelockTargetCaller // Generic read-only contract binding to access the raw methods on
}

// ITimelockTargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITimelockTargetTransactorRaw struct {
	Contract *ITimelockTargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITimelockTarget creates a new instance of ITimelockTarget, bound to a specific deployed contract.
func NewITimelockTarget(address common.Address, backend bind.ContractBackend) (*ITimelockTarget, error) {
	contract, err := bindITimelockTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITimelockTarget{ITimelockTargetCaller: ITimelockTargetCaller{contract: contract}, ITimelockTargetTransactor: ITimelockTargetTransactor{contract: contract}, ITimelockTargetFilterer: ITimelockTargetFilterer{contract: contract}}, nil
}

// NewITimelockTargetCaller creates a new read-only instance of ITimelockTarget, bound to a specific deployed contract.
func NewITimelockTargetCaller(address common.Address, caller bind.ContractCaller) (*ITimelockTargetCaller, error) {
	contract, err := bindITimelockTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITimelockTargetCaller{contract: contract}, nil
}

// NewITimelockTargetTransactor creates a new write-only instance of ITimelockTarget, bound to a specific deployed contract.
func NewITimelockTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*ITimelockTargetTransactor, error) {
	contract, err := bindITimelockTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITimelockTargetTransactor{contract: contract}, nil
}

// NewITimelockTargetFilterer creates a new log filterer instance of ITimelockTarget, bound to a specific deployed contract.
func NewITimelockTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*ITimelockTargetFilterer, error) {
	contract, err := bindITimelockTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITimelockTargetFilterer{contract: contract}, nil
}

// bindITimelockTarget binds a generic wrapper to an already deployed contract.
func bindITimelockTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITimelockTargetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimelockTarget *ITimelockTargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimelockTarget.Contract.ITimelockTargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimelockTarget *ITimelockTargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.ITimelockTargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimelockTarget *ITimelockTargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.ITimelockTargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimelockTarget *ITimelockTargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimelockTarget.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimelockTarget *ITimelockTargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimelockTarget *ITimelockTargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.contract.Transact(opts, method, params...)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ITimelockTarget *ITimelockTargetTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _ITimelockTarget.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ITimelockTarget *ITimelockTargetSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.SetGov(&_ITimelockTarget.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ITimelockTarget *ITimelockTargetTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.SetGov(&_ITimelockTarget.TransactOpts, _gov)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_ITimelockTarget *ITimelockTargetTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITimelockTarget.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_ITimelockTarget *ITimelockTargetSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.WithdrawToken(&_ITimelockTarget.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_ITimelockTarget *ITimelockTargetTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITimelockTarget.Contract.WithdrawToken(&_ITimelockTarget.TransactOpts, _token, _account, _amount)
}
