// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdtpricefeed

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

// USDTPriceFeedMetaData contains all meta data concerning the USDTPriceFeed contract.
var USDTPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"answer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"answers\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_answer\",\"type\":\"int256\"}],\"name\":\"setLatestAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// USDTPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use USDTPriceFeedMetaData.ABI instead.
var USDTPriceFeedABI = USDTPriceFeedMetaData.ABI

// USDTPriceFeed is an auto generated Go binding around an Ethereum contract.
type USDTPriceFeed struct {
	USDTPriceFeedCaller     // Read-only binding to the contract
	USDTPriceFeedTransactor // Write-only binding to the contract
	USDTPriceFeedFilterer   // Log filterer for contract events
}

// USDTPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDTPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDTPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDTPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDTPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDTPriceFeedSession struct {
	Contract     *USDTPriceFeed    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDTPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDTPriceFeedCallerSession struct {
	Contract *USDTPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// USDTPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDTPriceFeedTransactorSession struct {
	Contract     *USDTPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// USDTPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDTPriceFeedRaw struct {
	Contract *USDTPriceFeed // Generic contract binding to access the raw methods on
}

// USDTPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDTPriceFeedCallerRaw struct {
	Contract *USDTPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// USDTPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDTPriceFeedTransactorRaw struct {
	Contract *USDTPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDTPriceFeed creates a new instance of USDTPriceFeed, bound to a specific deployed contract.
func NewUSDTPriceFeed(address common.Address, backend bind.ContractBackend) (*USDTPriceFeed, error) {
	contract, err := bindUSDTPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDTPriceFeed{USDTPriceFeedCaller: USDTPriceFeedCaller{contract: contract}, USDTPriceFeedTransactor: USDTPriceFeedTransactor{contract: contract}, USDTPriceFeedFilterer: USDTPriceFeedFilterer{contract: contract}}, nil
}

// NewUSDTPriceFeedCaller creates a new read-only instance of USDTPriceFeed, bound to a specific deployed contract.
func NewUSDTPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*USDTPriceFeedCaller, error) {
	contract, err := bindUSDTPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDTPriceFeedCaller{contract: contract}, nil
}

// NewUSDTPriceFeedTransactor creates a new write-only instance of USDTPriceFeed, bound to a specific deployed contract.
func NewUSDTPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*USDTPriceFeedTransactor, error) {
	contract, err := bindUSDTPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDTPriceFeedTransactor{contract: contract}, nil
}

// NewUSDTPriceFeedFilterer creates a new log filterer instance of USDTPriceFeed, bound to a specific deployed contract.
func NewUSDTPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*USDTPriceFeedFilterer, error) {
	contract, err := bindUSDTPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDTPriceFeedFilterer{contract: contract}, nil
}

// bindUSDTPriceFeed binds a generic wrapper to an already deployed contract.
func bindUSDTPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDTPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDTPriceFeed *USDTPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDTPriceFeed.Contract.USDTPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDTPriceFeed *USDTPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.USDTPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDTPriceFeed *USDTPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.USDTPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDTPriceFeed *USDTPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDTPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDTPriceFeed *USDTPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDTPriceFeed *USDTPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_USDTPriceFeed *USDTPriceFeedCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_USDTPriceFeed *USDTPriceFeedSession) Aggregator() (common.Address, error) {
	return _USDTPriceFeed.Contract.Aggregator(&_USDTPriceFeed.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) Aggregator() (common.Address, error) {
	return _USDTPriceFeed.Contract.Aggregator(&_USDTPriceFeed.CallOpts)
}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedCaller) Answer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "answer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedSession) Answer() (*big.Int, error) {
	return _USDTPriceFeed.Contract.Answer(&_USDTPriceFeed.CallOpts)
}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) Answer() (*big.Int, error) {
	return _USDTPriceFeed.Contract.Answer(&_USDTPriceFeed.CallOpts)
}

// Answers is a free data retrieval call binding the contract method 0x4c295ca3.
//
// Solidity: function answers(uint80 ) view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedCaller) Answers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "answers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Answers is a free data retrieval call binding the contract method 0x4c295ca3.
//
// Solidity: function answers(uint80 ) view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedSession) Answers(arg0 *big.Int) (*big.Int, error) {
	return _USDTPriceFeed.Contract.Answers(&_USDTPriceFeed.CallOpts, arg0)
}

// Answers is a free data retrieval call binding the contract method 0x4c295ca3.
//
// Solidity: function answers(uint80 ) view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) Answers(arg0 *big.Int) (*big.Int, error) {
	return _USDTPriceFeed.Contract.Answers(&_USDTPriceFeed.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_USDTPriceFeed *USDTPriceFeedCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_USDTPriceFeed *USDTPriceFeedSession) Decimals() (*big.Int, error) {
	return _USDTPriceFeed.Contract.Decimals(&_USDTPriceFeed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) Decimals() (*big.Int, error) {
	return _USDTPriceFeed.Contract.Decimals(&_USDTPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_USDTPriceFeed *USDTPriceFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_USDTPriceFeed *USDTPriceFeedSession) Description() (string, error) {
	return _USDTPriceFeed.Contract.Description(&_USDTPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) Description() (string, error) {
	return _USDTPriceFeed.Contract.Description(&_USDTPriceFeed.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_USDTPriceFeed *USDTPriceFeedCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "getRoundData", _roundId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_USDTPriceFeed *USDTPriceFeedSession) GetRoundData(_roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _USDTPriceFeed.Contract.GetRoundData(&_USDTPriceFeed.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) GetRoundData(_roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _USDTPriceFeed.Contract.GetRoundData(&_USDTPriceFeed.CallOpts, _roundId)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_USDTPriceFeed *USDTPriceFeedCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_USDTPriceFeed *USDTPriceFeedSession) Gov() (common.Address, error) {
	return _USDTPriceFeed.Contract.Gov(&_USDTPriceFeed.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) Gov() (common.Address, error) {
	return _USDTPriceFeed.Contract.Gov(&_USDTPriceFeed.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_USDTPriceFeed *USDTPriceFeedCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_USDTPriceFeed *USDTPriceFeedSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _USDTPriceFeed.Contract.IsAdmin(&_USDTPriceFeed.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _USDTPriceFeed.Contract.IsAdmin(&_USDTPriceFeed.CallOpts, arg0)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedSession) LatestAnswer() (*big.Int, error) {
	return _USDTPriceFeed.Contract.LatestAnswer(&_USDTPriceFeed.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) LatestAnswer() (*big.Int, error) {
	return _USDTPriceFeed.Contract.LatestAnswer(&_USDTPriceFeed.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_USDTPriceFeed *USDTPriceFeedCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_USDTPriceFeed *USDTPriceFeedSession) LatestRound() (*big.Int, error) {
	return _USDTPriceFeed.Contract.LatestRound(&_USDTPriceFeed.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) LatestRound() (*big.Int, error) {
	return _USDTPriceFeed.Contract.LatestRound(&_USDTPriceFeed.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_USDTPriceFeed *USDTPriceFeedCaller) RoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDTPriceFeed.contract.Call(opts, &out, "roundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_USDTPriceFeed *USDTPriceFeedSession) RoundId() (*big.Int, error) {
	return _USDTPriceFeed.Contract.RoundId(&_USDTPriceFeed.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_USDTPriceFeed *USDTPriceFeedCallerSession) RoundId() (*big.Int, error) {
	return _USDTPriceFeed.Contract.RoundId(&_USDTPriceFeed.CallOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _account, bool _isAdmin) returns()
func (_USDTPriceFeed *USDTPriceFeedTransactor) SetAdmin(opts *bind.TransactOpts, _account common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _USDTPriceFeed.contract.Transact(opts, "setAdmin", _account, _isAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _account, bool _isAdmin) returns()
func (_USDTPriceFeed *USDTPriceFeedSession) SetAdmin(_account common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.SetAdmin(&_USDTPriceFeed.TransactOpts, _account, _isAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _account, bool _isAdmin) returns()
func (_USDTPriceFeed *USDTPriceFeedTransactorSession) SetAdmin(_account common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.SetAdmin(&_USDTPriceFeed.TransactOpts, _account, _isAdmin)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _answer) returns()
func (_USDTPriceFeed *USDTPriceFeedTransactor) SetLatestAnswer(opts *bind.TransactOpts, _answer *big.Int) (*types.Transaction, error) {
	return _USDTPriceFeed.contract.Transact(opts, "setLatestAnswer", _answer)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _answer) returns()
func (_USDTPriceFeed *USDTPriceFeedSession) SetLatestAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.SetLatestAnswer(&_USDTPriceFeed.TransactOpts, _answer)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _answer) returns()
func (_USDTPriceFeed *USDTPriceFeedTransactorSession) SetLatestAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _USDTPriceFeed.Contract.SetLatestAnswer(&_USDTPriceFeed.TransactOpts, _answer)
}
