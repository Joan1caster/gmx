// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igmxmigrator

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

// IGmxMigratorMetaData contains all meta data concerning the IGmxMigrator contract.
var IGmxMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"iouTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IGmxMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use IGmxMigratorMetaData.ABI instead.
var IGmxMigratorABI = IGmxMigratorMetaData.ABI

// IGmxMigrator is an auto generated Go binding around an Ethereum contract.
type IGmxMigrator struct {
	IGmxMigratorCaller     // Read-only binding to the contract
	IGmxMigratorTransactor // Write-only binding to the contract
	IGmxMigratorFilterer   // Log filterer for contract events
}

// IGmxMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGmxMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGmxMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGmxMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGmxMigratorSession struct {
	Contract     *IGmxMigrator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGmxMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGmxMigratorCallerSession struct {
	Contract *IGmxMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGmxMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGmxMigratorTransactorSession struct {
	Contract     *IGmxMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGmxMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGmxMigratorRaw struct {
	Contract *IGmxMigrator // Generic contract binding to access the raw methods on
}

// IGmxMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGmxMigratorCallerRaw struct {
	Contract *IGmxMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// IGmxMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGmxMigratorTransactorRaw struct {
	Contract *IGmxMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGmxMigrator creates a new instance of IGmxMigrator, bound to a specific deployed contract.
func NewIGmxMigrator(address common.Address, backend bind.ContractBackend) (*IGmxMigrator, error) {
	contract, err := bindIGmxMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGmxMigrator{IGmxMigratorCaller: IGmxMigratorCaller{contract: contract}, IGmxMigratorTransactor: IGmxMigratorTransactor{contract: contract}, IGmxMigratorFilterer: IGmxMigratorFilterer{contract: contract}}, nil
}

// NewIGmxMigratorCaller creates a new read-only instance of IGmxMigrator, bound to a specific deployed contract.
func NewIGmxMigratorCaller(address common.Address, caller bind.ContractCaller) (*IGmxMigratorCaller, error) {
	contract, err := bindIGmxMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGmxMigratorCaller{contract: contract}, nil
}

// NewIGmxMigratorTransactor creates a new write-only instance of IGmxMigrator, bound to a specific deployed contract.
func NewIGmxMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*IGmxMigratorTransactor, error) {
	contract, err := bindIGmxMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGmxMigratorTransactor{contract: contract}, nil
}

// NewIGmxMigratorFilterer creates a new log filterer instance of IGmxMigrator, bound to a specific deployed contract.
func NewIGmxMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*IGmxMigratorFilterer, error) {
	contract, err := bindIGmxMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGmxMigratorFilterer{contract: contract}, nil
}

// bindIGmxMigrator binds a generic wrapper to an already deployed contract.
func bindIGmxMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGmxMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGmxMigrator *IGmxMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGmxMigrator.Contract.IGmxMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGmxMigrator *IGmxMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGmxMigrator.Contract.IGmxMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGmxMigrator *IGmxMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGmxMigrator.Contract.IGmxMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGmxMigrator *IGmxMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGmxMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGmxMigrator *IGmxMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGmxMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGmxMigrator *IGmxMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGmxMigrator.Contract.contract.Transact(opts, method, params...)
}

// IouTokens is a free data retrieval call binding the contract method 0x41410d4a.
//
// Solidity: function iouTokens(address _token) view returns(address)
func (_IGmxMigrator *IGmxMigratorCaller) IouTokens(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IGmxMigrator.contract.Call(opts, &out, "iouTokens", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IouTokens is a free data retrieval call binding the contract method 0x41410d4a.
//
// Solidity: function iouTokens(address _token) view returns(address)
func (_IGmxMigrator *IGmxMigratorSession) IouTokens(_token common.Address) (common.Address, error) {
	return _IGmxMigrator.Contract.IouTokens(&_IGmxMigrator.CallOpts, _token)
}

// IouTokens is a free data retrieval call binding the contract method 0x41410d4a.
//
// Solidity: function iouTokens(address _token) view returns(address)
func (_IGmxMigrator *IGmxMigratorCallerSession) IouTokens(_token common.Address) (common.Address, error) {
	return _IGmxMigrator.Contract.IouTokens(&_IGmxMigrator.CallOpts, _token)
}
