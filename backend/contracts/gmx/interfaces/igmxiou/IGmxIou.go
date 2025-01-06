// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igmxiou

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

// IGmxIouMetaData contains all meta data concerning the IGmxIou contract.
var IGmxIouMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGmxIouABI is the input ABI used to generate the binding from.
// Deprecated: Use IGmxIouMetaData.ABI instead.
var IGmxIouABI = IGmxIouMetaData.ABI

// IGmxIou is an auto generated Go binding around an Ethereum contract.
type IGmxIou struct {
	IGmxIouCaller     // Read-only binding to the contract
	IGmxIouTransactor // Write-only binding to the contract
	IGmxIouFilterer   // Log filterer for contract events
}

// IGmxIouCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGmxIouCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxIouTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGmxIouTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxIouFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGmxIouFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGmxIouSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGmxIouSession struct {
	Contract     *IGmxIou          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGmxIouCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGmxIouCallerSession struct {
	Contract *IGmxIouCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IGmxIouTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGmxIouTransactorSession struct {
	Contract     *IGmxIouTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IGmxIouRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGmxIouRaw struct {
	Contract *IGmxIou // Generic contract binding to access the raw methods on
}

// IGmxIouCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGmxIouCallerRaw struct {
	Contract *IGmxIouCaller // Generic read-only contract binding to access the raw methods on
}

// IGmxIouTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGmxIouTransactorRaw struct {
	Contract *IGmxIouTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGmxIou creates a new instance of IGmxIou, bound to a specific deployed contract.
func NewIGmxIou(address common.Address, backend bind.ContractBackend) (*IGmxIou, error) {
	contract, err := bindIGmxIou(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGmxIou{IGmxIouCaller: IGmxIouCaller{contract: contract}, IGmxIouTransactor: IGmxIouTransactor{contract: contract}, IGmxIouFilterer: IGmxIouFilterer{contract: contract}}, nil
}

// NewIGmxIouCaller creates a new read-only instance of IGmxIou, bound to a specific deployed contract.
func NewIGmxIouCaller(address common.Address, caller bind.ContractCaller) (*IGmxIouCaller, error) {
	contract, err := bindIGmxIou(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGmxIouCaller{contract: contract}, nil
}

// NewIGmxIouTransactor creates a new write-only instance of IGmxIou, bound to a specific deployed contract.
func NewIGmxIouTransactor(address common.Address, transactor bind.ContractTransactor) (*IGmxIouTransactor, error) {
	contract, err := bindIGmxIou(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGmxIouTransactor{contract: contract}, nil
}

// NewIGmxIouFilterer creates a new log filterer instance of IGmxIou, bound to a specific deployed contract.
func NewIGmxIouFilterer(address common.Address, filterer bind.ContractFilterer) (*IGmxIouFilterer, error) {
	contract, err := bindIGmxIou(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGmxIouFilterer{contract: contract}, nil
}

// bindIGmxIou binds a generic wrapper to an already deployed contract.
func bindIGmxIou(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGmxIouMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGmxIou *IGmxIouRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGmxIou.Contract.IGmxIouCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGmxIou *IGmxIouRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGmxIou.Contract.IGmxIouTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGmxIou *IGmxIouRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGmxIou.Contract.IGmxIouTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGmxIou *IGmxIouCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGmxIou.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGmxIou *IGmxIouTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGmxIou.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGmxIou *IGmxIouTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGmxIou.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_IGmxIou *IGmxIouTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IGmxIou.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_IGmxIou *IGmxIouSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IGmxIou.Contract.Mint(&_IGmxIou.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_IGmxIou *IGmxIouTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IGmxIou.Contract.Mint(&_IGmxIou.TransactOpts, account, amount)
}
