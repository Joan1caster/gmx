// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pythstructs

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

// PythStructsMetaData contains all meta data concerning the PythStructs contract.
var PythStructsMetaData = &bind.MetaData{
	ABI: "[]",
}

// PythStructsABI is the input ABI used to generate the binding from.
// Deprecated: Use PythStructsMetaData.ABI instead.
var PythStructsABI = PythStructsMetaData.ABI

// PythStructs is an auto generated Go binding around an Ethereum contract.
type PythStructs struct {
	PythStructsCaller     // Read-only binding to the contract
	PythStructsTransactor // Write-only binding to the contract
	PythStructsFilterer   // Log filterer for contract events
}

// PythStructsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PythStructsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythStructsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PythStructsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythStructsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PythStructsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythStructsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PythStructsSession struct {
	Contract     *PythStructs      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PythStructsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PythStructsCallerSession struct {
	Contract *PythStructsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PythStructsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PythStructsTransactorSession struct {
	Contract     *PythStructsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PythStructsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PythStructsRaw struct {
	Contract *PythStructs // Generic contract binding to access the raw methods on
}

// PythStructsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PythStructsCallerRaw struct {
	Contract *PythStructsCaller // Generic read-only contract binding to access the raw methods on
}

// PythStructsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PythStructsTransactorRaw struct {
	Contract *PythStructsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPythStructs creates a new instance of PythStructs, bound to a specific deployed contract.
func NewPythStructs(address common.Address, backend bind.ContractBackend) (*PythStructs, error) {
	contract, err := bindPythStructs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PythStructs{PythStructsCaller: PythStructsCaller{contract: contract}, PythStructsTransactor: PythStructsTransactor{contract: contract}, PythStructsFilterer: PythStructsFilterer{contract: contract}}, nil
}

// NewPythStructsCaller creates a new read-only instance of PythStructs, bound to a specific deployed contract.
func NewPythStructsCaller(address common.Address, caller bind.ContractCaller) (*PythStructsCaller, error) {
	contract, err := bindPythStructs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PythStructsCaller{contract: contract}, nil
}

// NewPythStructsTransactor creates a new write-only instance of PythStructs, bound to a specific deployed contract.
func NewPythStructsTransactor(address common.Address, transactor bind.ContractTransactor) (*PythStructsTransactor, error) {
	contract, err := bindPythStructs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PythStructsTransactor{contract: contract}, nil
}

// NewPythStructsFilterer creates a new log filterer instance of PythStructs, bound to a specific deployed contract.
func NewPythStructsFilterer(address common.Address, filterer bind.ContractFilterer) (*PythStructsFilterer, error) {
	contract, err := bindPythStructs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PythStructsFilterer{contract: contract}, nil
}

// bindPythStructs binds a generic wrapper to an already deployed contract.
func bindPythStructs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PythStructsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PythStructs *PythStructsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PythStructs.Contract.PythStructsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PythStructs *PythStructsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PythStructs.Contract.PythStructsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PythStructs *PythStructsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PythStructs.Contract.PythStructsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PythStructs *PythStructsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PythStructs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PythStructs *PythStructsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PythStructs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PythStructs *PythStructsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PythStructs.Contract.contract.Transact(opts, method, params...)
}
