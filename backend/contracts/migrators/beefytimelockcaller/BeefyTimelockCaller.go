// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package beefytimelockcaller

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

// BeefyTimelockCallerMetaData contains all meta data concerning the BeefyTimelockCaller contract.
var BeefyTimelockCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executeProposals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_parent\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BeefyTimelockCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use BeefyTimelockCallerMetaData.ABI instead.
var BeefyTimelockCallerABI = BeefyTimelockCallerMetaData.ABI

// BeefyTimelockCaller is an auto generated Go binding around an Ethereum contract.
type BeefyTimelockCaller struct {
	BeefyTimelockCallerCaller     // Read-only binding to the contract
	BeefyTimelockCallerTransactor // Write-only binding to the contract
	BeefyTimelockCallerFilterer   // Log filterer for contract events
}

// BeefyTimelockCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeefyTimelockCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyTimelockCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeefyTimelockCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyTimelockCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeefyTimelockCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyTimelockCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeefyTimelockCallerSession struct {
	Contract     *BeefyTimelockCaller // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BeefyTimelockCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeefyTimelockCallerCallerSession struct {
	Contract *BeefyTimelockCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BeefyTimelockCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeefyTimelockCallerTransactorSession struct {
	Contract     *BeefyTimelockCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BeefyTimelockCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeefyTimelockCallerRaw struct {
	Contract *BeefyTimelockCaller // Generic contract binding to access the raw methods on
}

// BeefyTimelockCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeefyTimelockCallerCallerRaw struct {
	Contract *BeefyTimelockCallerCaller // Generic read-only contract binding to access the raw methods on
}

// BeefyTimelockCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeefyTimelockCallerTransactorRaw struct {
	Contract *BeefyTimelockCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeefyTimelockCaller creates a new instance of BeefyTimelockCaller, bound to a specific deployed contract.
func NewBeefyTimelockCaller(address common.Address, backend bind.ContractBackend) (*BeefyTimelockCaller, error) {
	contract, err := bindBeefyTimelockCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeefyTimelockCaller{BeefyTimelockCallerCaller: BeefyTimelockCallerCaller{contract: contract}, BeefyTimelockCallerTransactor: BeefyTimelockCallerTransactor{contract: contract}, BeefyTimelockCallerFilterer: BeefyTimelockCallerFilterer{contract: contract}}, nil
}

// NewBeefyTimelockCallerCaller creates a new read-only instance of BeefyTimelockCaller, bound to a specific deployed contract.
func NewBeefyTimelockCallerCaller(address common.Address, caller bind.ContractCaller) (*BeefyTimelockCallerCaller, error) {
	contract, err := bindBeefyTimelockCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyTimelockCallerCaller{contract: contract}, nil
}

// NewBeefyTimelockCallerTransactor creates a new write-only instance of BeefyTimelockCaller, bound to a specific deployed contract.
func NewBeefyTimelockCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*BeefyTimelockCallerTransactor, error) {
	contract, err := bindBeefyTimelockCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyTimelockCallerTransactor{contract: contract}, nil
}

// NewBeefyTimelockCallerFilterer creates a new log filterer instance of BeefyTimelockCaller, bound to a specific deployed contract.
func NewBeefyTimelockCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*BeefyTimelockCallerFilterer, error) {
	contract, err := bindBeefyTimelockCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeefyTimelockCallerFilterer{contract: contract}, nil
}

// bindBeefyTimelockCaller binds a generic wrapper to an already deployed contract.
func bindBeefyTimelockCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeefyTimelockCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyTimelockCaller *BeefyTimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyTimelockCaller.Contract.BeefyTimelockCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyTimelockCaller *BeefyTimelockCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.BeefyTimelockCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyTimelockCaller *BeefyTimelockCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.BeefyTimelockCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyTimelockCaller *BeefyTimelockCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyTimelockCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyTimelockCaller *BeefyTimelockCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyTimelockCaller *BeefyTimelockCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_BeefyTimelockCaller *BeefyTimelockCallerCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeefyTimelockCaller.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_BeefyTimelockCaller *BeefyTimelockCallerSession) ChainId() (*big.Int, error) {
	return _BeefyTimelockCaller.Contract.ChainId(&_BeefyTimelockCaller.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_BeefyTimelockCaller *BeefyTimelockCallerCallerSession) ChainId() (*big.Int, error) {
	return _BeefyTimelockCaller.Contract.ChainId(&_BeefyTimelockCaller.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BeefyTimelockCaller *BeefyTimelockCallerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyTimelockCaller.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BeefyTimelockCaller *BeefyTimelockCallerSession) Gov() (common.Address, error) {
	return _BeefyTimelockCaller.Contract.Gov(&_BeefyTimelockCaller.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BeefyTimelockCaller *BeefyTimelockCallerCallerSession) Gov() (common.Address, error) {
	return _BeefyTimelockCaller.Contract.Gov(&_BeefyTimelockCaller.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_BeefyTimelockCaller *BeefyTimelockCallerCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BeefyTimelockCaller.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_BeefyTimelockCaller *BeefyTimelockCallerSession) IsInitialized() (bool, error) {
	return _BeefyTimelockCaller.Contract.IsInitialized(&_BeefyTimelockCaller.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_BeefyTimelockCaller *BeefyTimelockCallerCallerSession) IsInitialized() (bool, error) {
	return _BeefyTimelockCaller.Contract.IsInitialized(&_BeefyTimelockCaller.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_BeefyTimelockCaller *BeefyTimelockCallerCaller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyTimelockCaller.contract.Call(opts, &out, "parent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_BeefyTimelockCaller *BeefyTimelockCallerSession) Parent() (common.Address, error) {
	return _BeefyTimelockCaller.Contract.Parent(&_BeefyTimelockCaller.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_BeefyTimelockCaller *BeefyTimelockCallerCallerSession) Parent() (common.Address, error) {
	return _BeefyTimelockCaller.Contract.Parent(&_BeefyTimelockCaller.CallOpts)
}

// ExecuteProposals is a paid mutator transaction binding the contract method 0x135aaf7f.
//
// Solidity: function executeProposals() returns()
func (_BeefyTimelockCaller *BeefyTimelockCallerTransactor) ExecuteProposals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyTimelockCaller.contract.Transact(opts, "executeProposals")
}

// ExecuteProposals is a paid mutator transaction binding the contract method 0x135aaf7f.
//
// Solidity: function executeProposals() returns()
func (_BeefyTimelockCaller *BeefyTimelockCallerSession) ExecuteProposals() (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.ExecuteProposals(&_BeefyTimelockCaller.TransactOpts)
}

// ExecuteProposals is a paid mutator transaction binding the contract method 0x135aaf7f.
//
// Solidity: function executeProposals() returns()
func (_BeefyTimelockCaller *BeefyTimelockCallerTransactorSession) ExecuteProposals() (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.ExecuteProposals(&_BeefyTimelockCaller.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _chainId, address _parent) returns()
func (_BeefyTimelockCaller *BeefyTimelockCallerTransactor) Initialize(opts *bind.TransactOpts, _chainId *big.Int, _parent common.Address) (*types.Transaction, error) {
	return _BeefyTimelockCaller.contract.Transact(opts, "initialize", _chainId, _parent)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _chainId, address _parent) returns()
func (_BeefyTimelockCaller *BeefyTimelockCallerSession) Initialize(_chainId *big.Int, _parent common.Address) (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.Initialize(&_BeefyTimelockCaller.TransactOpts, _chainId, _parent)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 _chainId, address _parent) returns()
func (_BeefyTimelockCaller *BeefyTimelockCallerTransactorSession) Initialize(_chainId *big.Int, _parent common.Address) (*types.Transaction, error) {
	return _BeefyTimelockCaller.Contract.Initialize(&_BeefyTimelockCaller.TransactOpts, _chainId, _parent)
}
