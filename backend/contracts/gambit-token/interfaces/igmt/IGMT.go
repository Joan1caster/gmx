// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igmt

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

// IGMTMetaData contains all meta data concerning the IGMT contract.
var IGMTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"beginMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGMTABI is the input ABI used to generate the binding from.
// Deprecated: Use IGMTMetaData.ABI instead.
var IGMTABI = IGMTMetaData.ABI

// IGMT is an auto generated Go binding around an Ethereum contract.
type IGMT struct {
	IGMTCaller     // Read-only binding to the contract
	IGMTTransactor // Write-only binding to the contract
	IGMTFilterer   // Log filterer for contract events
}

// IGMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGMTSession struct {
	Contract     *IGMT             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGMTCallerSession struct {
	Contract *IGMTCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGMTTransactorSession struct {
	Contract     *IGMTTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGMTRaw struct {
	Contract *IGMT // Generic contract binding to access the raw methods on
}

// IGMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGMTCallerRaw struct {
	Contract *IGMTCaller // Generic read-only contract binding to access the raw methods on
}

// IGMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGMTTransactorRaw struct {
	Contract *IGMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGMT creates a new instance of IGMT, bound to a specific deployed contract.
func NewIGMT(address common.Address, backend bind.ContractBackend) (*IGMT, error) {
	contract, err := bindIGMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGMT{IGMTCaller: IGMTCaller{contract: contract}, IGMTTransactor: IGMTTransactor{contract: contract}, IGMTFilterer: IGMTFilterer{contract: contract}}, nil
}

// NewIGMTCaller creates a new read-only instance of IGMT, bound to a specific deployed contract.
func NewIGMTCaller(address common.Address, caller bind.ContractCaller) (*IGMTCaller, error) {
	contract, err := bindIGMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGMTCaller{contract: contract}, nil
}

// NewIGMTTransactor creates a new write-only instance of IGMT, bound to a specific deployed contract.
func NewIGMTTransactor(address common.Address, transactor bind.ContractTransactor) (*IGMTTransactor, error) {
	contract, err := bindIGMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGMTTransactor{contract: contract}, nil
}

// NewIGMTFilterer creates a new log filterer instance of IGMT, bound to a specific deployed contract.
func NewIGMTFilterer(address common.Address, filterer bind.ContractFilterer) (*IGMTFilterer, error) {
	contract, err := bindIGMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGMTFilterer{contract: contract}, nil
}

// bindIGMT binds a generic wrapper to an already deployed contract.
func bindIGMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMT *IGMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMT.Contract.IGMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMT *IGMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMT.Contract.IGMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMT *IGMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMT.Contract.IGMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGMT *IGMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGMT *IGMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGMT *IGMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGMT.Contract.contract.Transact(opts, method, params...)
}

// BeginMigration is a paid mutator transaction binding the contract method 0xfe0194f2.
//
// Solidity: function beginMigration() returns()
func (_IGMT *IGMTTransactor) BeginMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMT.contract.Transact(opts, "beginMigration")
}

// BeginMigration is a paid mutator transaction binding the contract method 0xfe0194f2.
//
// Solidity: function beginMigration() returns()
func (_IGMT *IGMTSession) BeginMigration() (*types.Transaction, error) {
	return _IGMT.Contract.BeginMigration(&_IGMT.TransactOpts)
}

// BeginMigration is a paid mutator transaction binding the contract method 0xfe0194f2.
//
// Solidity: function beginMigration() returns()
func (_IGMT *IGMTTransactorSession) BeginMigration() (*types.Transaction, error) {
	return _IGMT.Contract.BeginMigration(&_IGMT.TransactOpts)
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_IGMT *IGMTTransactor) EndMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGMT.contract.Transact(opts, "endMigration")
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_IGMT *IGMTSession) EndMigration() (*types.Transaction, error) {
	return _IGMT.Contract.EndMigration(&_IGMT.TransactOpts)
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_IGMT *IGMTTransactorSession) EndMigration() (*types.Transaction, error) {
	return _IGMT.Contract.EndMigration(&_IGMT.TransactOpts)
}
