// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stabilizecaller

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

// StabilizeCallerMetaData contains all meta data concerning the StabilizeCaller contract.
var StabilizeCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"completeMove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_parent\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StabilizeCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use StabilizeCallerMetaData.ABI instead.
var StabilizeCallerABI = StabilizeCallerMetaData.ABI

// StabilizeCaller is an auto generated Go binding around an Ethereum contract.
type StabilizeCaller struct {
	StabilizeCallerCaller     // Read-only binding to the contract
	StabilizeCallerTransactor // Write-only binding to the contract
	StabilizeCallerFilterer   // Log filterer for contract events
}

// StabilizeCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StabilizeCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizeCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StabilizeCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizeCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StabilizeCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizeCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StabilizeCallerSession struct {
	Contract     *StabilizeCaller  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StabilizeCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StabilizeCallerCallerSession struct {
	Contract *StabilizeCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StabilizeCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StabilizeCallerTransactorSession struct {
	Contract     *StabilizeCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StabilizeCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StabilizeCallerRaw struct {
	Contract *StabilizeCaller // Generic contract binding to access the raw methods on
}

// StabilizeCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StabilizeCallerCallerRaw struct {
	Contract *StabilizeCallerCaller // Generic read-only contract binding to access the raw methods on
}

// StabilizeCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StabilizeCallerTransactorRaw struct {
	Contract *StabilizeCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStabilizeCaller creates a new instance of StabilizeCaller, bound to a specific deployed contract.
func NewStabilizeCaller(address common.Address, backend bind.ContractBackend) (*StabilizeCaller, error) {
	contract, err := bindStabilizeCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StabilizeCaller{StabilizeCallerCaller: StabilizeCallerCaller{contract: contract}, StabilizeCallerTransactor: StabilizeCallerTransactor{contract: contract}, StabilizeCallerFilterer: StabilizeCallerFilterer{contract: contract}}, nil
}

// NewStabilizeCallerCaller creates a new read-only instance of StabilizeCaller, bound to a specific deployed contract.
func NewStabilizeCallerCaller(address common.Address, caller bind.ContractCaller) (*StabilizeCallerCaller, error) {
	contract, err := bindStabilizeCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StabilizeCallerCaller{contract: contract}, nil
}

// NewStabilizeCallerTransactor creates a new write-only instance of StabilizeCaller, bound to a specific deployed contract.
func NewStabilizeCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*StabilizeCallerTransactor, error) {
	contract, err := bindStabilizeCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StabilizeCallerTransactor{contract: contract}, nil
}

// NewStabilizeCallerFilterer creates a new log filterer instance of StabilizeCaller, bound to a specific deployed contract.
func NewStabilizeCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*StabilizeCallerFilterer, error) {
	contract, err := bindStabilizeCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StabilizeCallerFilterer{contract: contract}, nil
}

// bindStabilizeCaller binds a generic wrapper to an already deployed contract.
func bindStabilizeCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StabilizeCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StabilizeCaller *StabilizeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StabilizeCaller.Contract.StabilizeCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StabilizeCaller *StabilizeCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StabilizeCaller.Contract.StabilizeCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StabilizeCaller *StabilizeCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StabilizeCaller.Contract.StabilizeCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StabilizeCaller *StabilizeCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StabilizeCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StabilizeCaller *StabilizeCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StabilizeCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StabilizeCaller *StabilizeCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StabilizeCaller.Contract.contract.Transact(opts, method, params...)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_StabilizeCaller *StabilizeCallerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeCaller.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_StabilizeCaller *StabilizeCallerSession) Gov() (common.Address, error) {
	return _StabilizeCaller.Contract.Gov(&_StabilizeCaller.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_StabilizeCaller *StabilizeCallerCallerSession) Gov() (common.Address, error) {
	return _StabilizeCaller.Contract.Gov(&_StabilizeCaller.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_StabilizeCaller *StabilizeCallerCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StabilizeCaller.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_StabilizeCaller *StabilizeCallerSession) IsInitialized() (bool, error) {
	return _StabilizeCaller.Contract.IsInitialized(&_StabilizeCaller.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_StabilizeCaller *StabilizeCallerCallerSession) IsInitialized() (bool, error) {
	return _StabilizeCaller.Contract.IsInitialized(&_StabilizeCaller.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_StabilizeCaller *StabilizeCallerCaller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeCaller.contract.Call(opts, &out, "parent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_StabilizeCaller *StabilizeCallerSession) Parent() (common.Address, error) {
	return _StabilizeCaller.Contract.Parent(&_StabilizeCaller.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_StabilizeCaller *StabilizeCallerCallerSession) Parent() (common.Address, error) {
	return _StabilizeCaller.Contract.Parent(&_StabilizeCaller.CallOpts)
}

// CompleteMove is a paid mutator transaction binding the contract method 0x822a7c2f.
//
// Solidity: function completeMove() returns()
func (_StabilizeCaller *StabilizeCallerTransactor) CompleteMove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StabilizeCaller.contract.Transact(opts, "completeMove")
}

// CompleteMove is a paid mutator transaction binding the contract method 0x822a7c2f.
//
// Solidity: function completeMove() returns()
func (_StabilizeCaller *StabilizeCallerSession) CompleteMove() (*types.Transaction, error) {
	return _StabilizeCaller.Contract.CompleteMove(&_StabilizeCaller.TransactOpts)
}

// CompleteMove is a paid mutator transaction binding the contract method 0x822a7c2f.
//
// Solidity: function completeMove() returns()
func (_StabilizeCaller *StabilizeCallerTransactorSession) CompleteMove() (*types.Transaction, error) {
	return _StabilizeCaller.Contract.CompleteMove(&_StabilizeCaller.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _parent) returns()
func (_StabilizeCaller *StabilizeCallerTransactor) Initialize(opts *bind.TransactOpts, _parent common.Address) (*types.Transaction, error) {
	return _StabilizeCaller.contract.Transact(opts, "initialize", _parent)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _parent) returns()
func (_StabilizeCaller *StabilizeCallerSession) Initialize(_parent common.Address) (*types.Transaction, error) {
	return _StabilizeCaller.Contract.Initialize(&_StabilizeCaller.TransactOpts, _parent)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _parent) returns()
func (_StabilizeCaller *StabilizeCallerTransactorSession) Initialize(_parent common.Address) (*types.Transaction, error) {
	return _StabilizeCaller.Contract.Initialize(&_StabilizeCaller.TransactOpts, _parent)
}
