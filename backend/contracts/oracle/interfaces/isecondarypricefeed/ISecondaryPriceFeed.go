// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package isecondarypricefeed

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

// ISecondaryPriceFeedMetaData contains all meta data concerning the ISecondaryPriceFeed contract.
var ISecondaryPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_referencePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISecondaryPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use ISecondaryPriceFeedMetaData.ABI instead.
var ISecondaryPriceFeedABI = ISecondaryPriceFeedMetaData.ABI

// ISecondaryPriceFeed is an auto generated Go binding around an Ethereum contract.
type ISecondaryPriceFeed struct {
	ISecondaryPriceFeedCaller     // Read-only binding to the contract
	ISecondaryPriceFeedTransactor // Write-only binding to the contract
	ISecondaryPriceFeedFilterer   // Log filterer for contract events
}

// ISecondaryPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISecondaryPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISecondaryPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISecondaryPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISecondaryPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISecondaryPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISecondaryPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISecondaryPriceFeedSession struct {
	Contract     *ISecondaryPriceFeed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISecondaryPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISecondaryPriceFeedCallerSession struct {
	Contract *ISecondaryPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ISecondaryPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISecondaryPriceFeedTransactorSession struct {
	Contract     *ISecondaryPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ISecondaryPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISecondaryPriceFeedRaw struct {
	Contract *ISecondaryPriceFeed // Generic contract binding to access the raw methods on
}

// ISecondaryPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISecondaryPriceFeedCallerRaw struct {
	Contract *ISecondaryPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// ISecondaryPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISecondaryPriceFeedTransactorRaw struct {
	Contract *ISecondaryPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISecondaryPriceFeed creates a new instance of ISecondaryPriceFeed, bound to a specific deployed contract.
func NewISecondaryPriceFeed(address common.Address, backend bind.ContractBackend) (*ISecondaryPriceFeed, error) {
	contract, err := bindISecondaryPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISecondaryPriceFeed{ISecondaryPriceFeedCaller: ISecondaryPriceFeedCaller{contract: contract}, ISecondaryPriceFeedTransactor: ISecondaryPriceFeedTransactor{contract: contract}, ISecondaryPriceFeedFilterer: ISecondaryPriceFeedFilterer{contract: contract}}, nil
}

// NewISecondaryPriceFeedCaller creates a new read-only instance of ISecondaryPriceFeed, bound to a specific deployed contract.
func NewISecondaryPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*ISecondaryPriceFeedCaller, error) {
	contract, err := bindISecondaryPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISecondaryPriceFeedCaller{contract: contract}, nil
}

// NewISecondaryPriceFeedTransactor creates a new write-only instance of ISecondaryPriceFeed, bound to a specific deployed contract.
func NewISecondaryPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*ISecondaryPriceFeedTransactor, error) {
	contract, err := bindISecondaryPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISecondaryPriceFeedTransactor{contract: contract}, nil
}

// NewISecondaryPriceFeedFilterer creates a new log filterer instance of ISecondaryPriceFeed, bound to a specific deployed contract.
func NewISecondaryPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*ISecondaryPriceFeedFilterer, error) {
	contract, err := bindISecondaryPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISecondaryPriceFeedFilterer{contract: contract}, nil
}

// bindISecondaryPriceFeed binds a generic wrapper to an already deployed contract.
func bindISecondaryPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISecondaryPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISecondaryPriceFeed *ISecondaryPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISecondaryPriceFeed.Contract.ISecondaryPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISecondaryPriceFeed *ISecondaryPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISecondaryPriceFeed.Contract.ISecondaryPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISecondaryPriceFeed *ISecondaryPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISecondaryPriceFeed.Contract.ISecondaryPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISecondaryPriceFeed *ISecondaryPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISecondaryPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISecondaryPriceFeed *ISecondaryPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISecondaryPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISecondaryPriceFeed *ISecondaryPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISecondaryPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _referencePrice, bool _maximise) view returns(uint256)
func (_ISecondaryPriceFeed *ISecondaryPriceFeedCaller) GetPrice(opts *bind.CallOpts, _token common.Address, _referencePrice *big.Int, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _ISecondaryPriceFeed.contract.Call(opts, &out, "getPrice", _token, _referencePrice, _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _referencePrice, bool _maximise) view returns(uint256)
func (_ISecondaryPriceFeed *ISecondaryPriceFeedSession) GetPrice(_token common.Address, _referencePrice *big.Int, _maximise bool) (*big.Int, error) {
	return _ISecondaryPriceFeed.Contract.GetPrice(&_ISecondaryPriceFeed.CallOpts, _token, _referencePrice, _maximise)
}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _referencePrice, bool _maximise) view returns(uint256)
func (_ISecondaryPriceFeed *ISecondaryPriceFeedCallerSession) GetPrice(_token common.Address, _referencePrice *big.Int, _maximise bool) (*big.Int, error) {
	return _ISecondaryPriceFeed.Contract.GetPrice(&_ISecondaryPriceFeed.CallOpts, _token, _referencePrice, _maximise)
}
