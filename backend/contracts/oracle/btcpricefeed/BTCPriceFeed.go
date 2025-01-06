// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package btcpricefeed

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

// BTCPriceFeedMetaData contains all meta data concerning the BTCPriceFeed contract.
var BTCPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"answer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"answers\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_answer\",\"type\":\"int256\"}],\"name\":\"setLatestAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BTCPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use BTCPriceFeedMetaData.ABI instead.
var BTCPriceFeedABI = BTCPriceFeedMetaData.ABI

// BTCPriceFeed is an auto generated Go binding around an Ethereum contract.
type BTCPriceFeed struct {
	BTCPriceFeedCaller     // Read-only binding to the contract
	BTCPriceFeedTransactor // Write-only binding to the contract
	BTCPriceFeedFilterer   // Log filterer for contract events
}

// BTCPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type BTCPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTCPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BTCPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BTCPriceFeedSession struct {
	Contract     *BTCPriceFeed     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTCPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BTCPriceFeedCallerSession struct {
	Contract *BTCPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BTCPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTCPriceFeedTransactorSession struct {
	Contract     *BTCPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BTCPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type BTCPriceFeedRaw struct {
	Contract *BTCPriceFeed // Generic contract binding to access the raw methods on
}

// BTCPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BTCPriceFeedCallerRaw struct {
	Contract *BTCPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// BTCPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTCPriceFeedTransactorRaw struct {
	Contract *BTCPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBTCPriceFeed creates a new instance of BTCPriceFeed, bound to a specific deployed contract.
func NewBTCPriceFeed(address common.Address, backend bind.ContractBackend) (*BTCPriceFeed, error) {
	contract, err := bindBTCPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BTCPriceFeed{BTCPriceFeedCaller: BTCPriceFeedCaller{contract: contract}, BTCPriceFeedTransactor: BTCPriceFeedTransactor{contract: contract}, BTCPriceFeedFilterer: BTCPriceFeedFilterer{contract: contract}}, nil
}

// NewBTCPriceFeedCaller creates a new read-only instance of BTCPriceFeed, bound to a specific deployed contract.
func NewBTCPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*BTCPriceFeedCaller, error) {
	contract, err := bindBTCPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BTCPriceFeedCaller{contract: contract}, nil
}

// NewBTCPriceFeedTransactor creates a new write-only instance of BTCPriceFeed, bound to a specific deployed contract.
func NewBTCPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*BTCPriceFeedTransactor, error) {
	contract, err := bindBTCPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTCPriceFeedTransactor{contract: contract}, nil
}

// NewBTCPriceFeedFilterer creates a new log filterer instance of BTCPriceFeed, bound to a specific deployed contract.
func NewBTCPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*BTCPriceFeedFilterer, error) {
	contract, err := bindBTCPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BTCPriceFeedFilterer{contract: contract}, nil
}

// bindBTCPriceFeed binds a generic wrapper to an already deployed contract.
func bindBTCPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BTCPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTCPriceFeed *BTCPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTCPriceFeed.Contract.BTCPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTCPriceFeed *BTCPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.BTCPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTCPriceFeed *BTCPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.BTCPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTCPriceFeed *BTCPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTCPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTCPriceFeed *BTCPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTCPriceFeed *BTCPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BTCPriceFeed *BTCPriceFeedCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BTCPriceFeed *BTCPriceFeedSession) Aggregator() (common.Address, error) {
	return _BTCPriceFeed.Contract.Aggregator(&_BTCPriceFeed.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) Aggregator() (common.Address, error) {
	return _BTCPriceFeed.Contract.Aggregator(&_BTCPriceFeed.CallOpts)
}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedCaller) Answer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "answer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedSession) Answer() (*big.Int, error) {
	return _BTCPriceFeed.Contract.Answer(&_BTCPriceFeed.CallOpts)
}

// Answer is a free data retrieval call binding the contract method 0x85bb7d69.
//
// Solidity: function answer() view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) Answer() (*big.Int, error) {
	return _BTCPriceFeed.Contract.Answer(&_BTCPriceFeed.CallOpts)
}

// Answers is a free data retrieval call binding the contract method 0x4c295ca3.
//
// Solidity: function answers(uint80 ) view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedCaller) Answers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "answers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Answers is a free data retrieval call binding the contract method 0x4c295ca3.
//
// Solidity: function answers(uint80 ) view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedSession) Answers(arg0 *big.Int) (*big.Int, error) {
	return _BTCPriceFeed.Contract.Answers(&_BTCPriceFeed.CallOpts, arg0)
}

// Answers is a free data retrieval call binding the contract method 0x4c295ca3.
//
// Solidity: function answers(uint80 ) view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) Answers(arg0 *big.Int) (*big.Int, error) {
	return _BTCPriceFeed.Contract.Answers(&_BTCPriceFeed.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_BTCPriceFeed *BTCPriceFeedCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_BTCPriceFeed *BTCPriceFeedSession) Decimals() (*big.Int, error) {
	return _BTCPriceFeed.Contract.Decimals(&_BTCPriceFeed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) Decimals() (*big.Int, error) {
	return _BTCPriceFeed.Contract.Decimals(&_BTCPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_BTCPriceFeed *BTCPriceFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_BTCPriceFeed *BTCPriceFeedSession) Description() (string, error) {
	return _BTCPriceFeed.Contract.Description(&_BTCPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) Description() (string, error) {
	return _BTCPriceFeed.Contract.Description(&_BTCPriceFeed.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_BTCPriceFeed *BTCPriceFeedCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "getRoundData", _roundId)

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
func (_BTCPriceFeed *BTCPriceFeedSession) GetRoundData(_roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BTCPriceFeed.Contract.GetRoundData(&_BTCPriceFeed.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) GetRoundData(_roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BTCPriceFeed.Contract.GetRoundData(&_BTCPriceFeed.CallOpts, _roundId)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BTCPriceFeed *BTCPriceFeedCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BTCPriceFeed *BTCPriceFeedSession) Gov() (common.Address, error) {
	return _BTCPriceFeed.Contract.Gov(&_BTCPriceFeed.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) Gov() (common.Address, error) {
	return _BTCPriceFeed.Contract.Gov(&_BTCPriceFeed.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_BTCPriceFeed *BTCPriceFeedCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_BTCPriceFeed *BTCPriceFeedSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _BTCPriceFeed.Contract.IsAdmin(&_BTCPriceFeed.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _BTCPriceFeed.Contract.IsAdmin(&_BTCPriceFeed.CallOpts, arg0)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedSession) LatestAnswer() (*big.Int, error) {
	return _BTCPriceFeed.Contract.LatestAnswer(&_BTCPriceFeed.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) LatestAnswer() (*big.Int, error) {
	return _BTCPriceFeed.Contract.LatestAnswer(&_BTCPriceFeed.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_BTCPriceFeed *BTCPriceFeedCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_BTCPriceFeed *BTCPriceFeedSession) LatestRound() (*big.Int, error) {
	return _BTCPriceFeed.Contract.LatestRound(&_BTCPriceFeed.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) LatestRound() (*big.Int, error) {
	return _BTCPriceFeed.Contract.LatestRound(&_BTCPriceFeed.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_BTCPriceFeed *BTCPriceFeedCaller) RoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCPriceFeed.contract.Call(opts, &out, "roundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_BTCPriceFeed *BTCPriceFeedSession) RoundId() (*big.Int, error) {
	return _BTCPriceFeed.Contract.RoundId(&_BTCPriceFeed.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_BTCPriceFeed *BTCPriceFeedCallerSession) RoundId() (*big.Int, error) {
	return _BTCPriceFeed.Contract.RoundId(&_BTCPriceFeed.CallOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _account, bool _isAdmin) returns()
func (_BTCPriceFeed *BTCPriceFeedTransactor) SetAdmin(opts *bind.TransactOpts, _account common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _BTCPriceFeed.contract.Transact(opts, "setAdmin", _account, _isAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _account, bool _isAdmin) returns()
func (_BTCPriceFeed *BTCPriceFeedSession) SetAdmin(_account common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.SetAdmin(&_BTCPriceFeed.TransactOpts, _account, _isAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address _account, bool _isAdmin) returns()
func (_BTCPriceFeed *BTCPriceFeedTransactorSession) SetAdmin(_account common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.SetAdmin(&_BTCPriceFeed.TransactOpts, _account, _isAdmin)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _answer) returns()
func (_BTCPriceFeed *BTCPriceFeedTransactor) SetLatestAnswer(opts *bind.TransactOpts, _answer *big.Int) (*types.Transaction, error) {
	return _BTCPriceFeed.contract.Transact(opts, "setLatestAnswer", _answer)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _answer) returns()
func (_BTCPriceFeed *BTCPriceFeedSession) SetLatestAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.SetLatestAnswer(&_BTCPriceFeed.TransactOpts, _answer)
}

// SetLatestAnswer is a paid mutator transaction binding the contract method 0x04ea97b0.
//
// Solidity: function setLatestAnswer(int256 _answer) returns()
func (_BTCPriceFeed *BTCPriceFeedTransactorSession) SetLatestAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _BTCPriceFeed.Contract.SetLatestAnswer(&_BTCPriceFeed.TransactOpts, _answer)
}
