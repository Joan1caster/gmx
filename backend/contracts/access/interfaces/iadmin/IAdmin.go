// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iadmin

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

// IAdminMetaData contains all meta data concerning the IAdmin contract.
var IAdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use IAdminMetaData.ABI instead.
var IAdminABI = IAdminMetaData.ABI

// IAdmin is an auto generated Go binding around an Ethereum contract.
type IAdmin struct {
	IAdminCaller     // Read-only binding to the contract
	IAdminTransactor // Write-only binding to the contract
	IAdminFilterer   // Log filterer for contract events
}

// IAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAdminSession struct {
	Contract     *IAdmin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAdminCallerSession struct {
	Contract *IAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAdminTransactorSession struct {
	Contract     *IAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAdminRaw struct {
	Contract *IAdmin // Generic contract binding to access the raw methods on
}

// IAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAdminCallerRaw struct {
	Contract *IAdminCaller // Generic read-only contract binding to access the raw methods on
}

// IAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAdminTransactorRaw struct {
	Contract *IAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAdmin creates a new instance of IAdmin, bound to a specific deployed contract.
func NewIAdmin(address common.Address, backend bind.ContractBackend) (*IAdmin, error) {
	contract, err := bindIAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAdmin{IAdminCaller: IAdminCaller{contract: contract}, IAdminTransactor: IAdminTransactor{contract: contract}, IAdminFilterer: IAdminFilterer{contract: contract}}, nil
}

// NewIAdminCaller creates a new read-only instance of IAdmin, bound to a specific deployed contract.
func NewIAdminCaller(address common.Address, caller bind.ContractCaller) (*IAdminCaller, error) {
	contract, err := bindIAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminCaller{contract: contract}, nil
}

// NewIAdminTransactor creates a new write-only instance of IAdmin, bound to a specific deployed contract.
func NewIAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*IAdminTransactor, error) {
	contract, err := bindIAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAdminTransactor{contract: contract}, nil
}

// NewIAdminFilterer creates a new log filterer instance of IAdmin, bound to a specific deployed contract.
func NewIAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*IAdminFilterer, error) {
	contract, err := bindIAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAdminFilterer{contract: contract}, nil
}

// bindIAdmin binds a generic wrapper to an already deployed contract.
func bindIAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAdminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdmin *IAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdmin.Contract.IAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdmin *IAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdmin.Contract.IAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdmin *IAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdmin.Contract.IAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAdmin *IAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAdmin *IAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAdmin *IAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAdmin.Contract.contract.Transact(opts, method, params...)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_IAdmin *IAdminTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _IAdmin.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_IAdmin *IAdminSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.SetAdmin(&_IAdmin.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_IAdmin *IAdminTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _IAdmin.Contract.SetAdmin(&_IAdmin.TransactOpts, _admin)
}
