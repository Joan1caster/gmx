// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balanceupdater

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

// BalanceUpdaterMetaData contains all meta data concerning the BalanceUpdater contract.
var BalanceUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdg\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"updateBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BalanceUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceUpdaterMetaData.ABI instead.
var BalanceUpdaterABI = BalanceUpdaterMetaData.ABI

// BalanceUpdater is an auto generated Go binding around an Ethereum contract.
type BalanceUpdater struct {
	BalanceUpdaterCaller     // Read-only binding to the contract
	BalanceUpdaterTransactor // Write-only binding to the contract
	BalanceUpdaterFilterer   // Log filterer for contract events
}

// BalanceUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceUpdaterSession struct {
	Contract     *BalanceUpdater   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceUpdaterCallerSession struct {
	Contract *BalanceUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalanceUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceUpdaterTransactorSession struct {
	Contract     *BalanceUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BalanceUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceUpdaterRaw struct {
	Contract *BalanceUpdater // Generic contract binding to access the raw methods on
}

// BalanceUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceUpdaterCallerRaw struct {
	Contract *BalanceUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceUpdaterTransactorRaw struct {
	Contract *BalanceUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceUpdater creates a new instance of BalanceUpdater, bound to a specific deployed contract.
func NewBalanceUpdater(address common.Address, backend bind.ContractBackend) (*BalanceUpdater, error) {
	contract, err := bindBalanceUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceUpdater{BalanceUpdaterCaller: BalanceUpdaterCaller{contract: contract}, BalanceUpdaterTransactor: BalanceUpdaterTransactor{contract: contract}, BalanceUpdaterFilterer: BalanceUpdaterFilterer{contract: contract}}, nil
}

// NewBalanceUpdaterCaller creates a new read-only instance of BalanceUpdater, bound to a specific deployed contract.
func NewBalanceUpdaterCaller(address common.Address, caller bind.ContractCaller) (*BalanceUpdaterCaller, error) {
	contract, err := bindBalanceUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceUpdaterCaller{contract: contract}, nil
}

// NewBalanceUpdaterTransactor creates a new write-only instance of BalanceUpdater, bound to a specific deployed contract.
func NewBalanceUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceUpdaterTransactor, error) {
	contract, err := bindBalanceUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceUpdaterTransactor{contract: contract}, nil
}

// NewBalanceUpdaterFilterer creates a new log filterer instance of BalanceUpdater, bound to a specific deployed contract.
func NewBalanceUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceUpdaterFilterer, error) {
	contract, err := bindBalanceUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceUpdaterFilterer{contract: contract}, nil
}

// bindBalanceUpdater binds a generic wrapper to an already deployed contract.
func bindBalanceUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalanceUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceUpdater *BalanceUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceUpdater.Contract.BalanceUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceUpdater *BalanceUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceUpdater.Contract.BalanceUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceUpdater *BalanceUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceUpdater.Contract.BalanceUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceUpdater *BalanceUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceUpdater *BalanceUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceUpdater *BalanceUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceUpdater.Contract.contract.Transact(opts, method, params...)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x43aaa690.
//
// Solidity: function updateBalance(address _vault, address _token, address _usdg, uint256 _usdgAmount) returns()
func (_BalanceUpdater *BalanceUpdaterTransactor) UpdateBalance(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _usdg common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _BalanceUpdater.contract.Transact(opts, "updateBalance", _vault, _token, _usdg, _usdgAmount)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x43aaa690.
//
// Solidity: function updateBalance(address _vault, address _token, address _usdg, uint256 _usdgAmount) returns()
func (_BalanceUpdater *BalanceUpdaterSession) UpdateBalance(_vault common.Address, _token common.Address, _usdg common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _BalanceUpdater.Contract.UpdateBalance(&_BalanceUpdater.TransactOpts, _vault, _token, _usdg, _usdgAmount)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x43aaa690.
//
// Solidity: function updateBalance(address _vault, address _token, address _usdg, uint256 _usdgAmount) returns()
func (_BalanceUpdater *BalanceUpdaterTransactorSession) UpdateBalance(_vault common.Address, _token common.Address, _usdg common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _BalanceUpdater.Contract.UpdateBalance(&_BalanceUpdater.TransactOpts, _vault, _token, _usdg, _usdgAmount)
}
