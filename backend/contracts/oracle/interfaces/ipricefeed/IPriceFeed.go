// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipricefeed

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

// IPriceFeedMetaData contains all meta data concerning the IPriceFeed contract.
var IPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use IPriceFeedMetaData.ABI instead.
var IPriceFeedABI = IPriceFeedMetaData.ABI

// IPriceFeed is an auto generated Go binding around an Ethereum contract.
type IPriceFeed struct {
	IPriceFeedCaller     // Read-only binding to the contract
	IPriceFeedTransactor // Write-only binding to the contract
	IPriceFeedFilterer   // Log filterer for contract events
}

// IPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPriceFeedSession struct {
	Contract     *IPriceFeed       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPriceFeedCallerSession struct {
	Contract *IPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPriceFeedTransactorSession struct {
	Contract     *IPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPriceFeedRaw struct {
	Contract *IPriceFeed // Generic contract binding to access the raw methods on
}

// IPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPriceFeedCallerRaw struct {
	Contract *IPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// IPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPriceFeedTransactorRaw struct {
	Contract *IPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPriceFeed creates a new instance of IPriceFeed, bound to a specific deployed contract.
func NewIPriceFeed(address common.Address, backend bind.ContractBackend) (*IPriceFeed, error) {
	contract, err := bindIPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPriceFeed{IPriceFeedCaller: IPriceFeedCaller{contract: contract}, IPriceFeedTransactor: IPriceFeedTransactor{contract: contract}, IPriceFeedFilterer: IPriceFeedFilterer{contract: contract}}, nil
}

// NewIPriceFeedCaller creates a new read-only instance of IPriceFeed, bound to a specific deployed contract.
func NewIPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*IPriceFeedCaller, error) {
	contract, err := bindIPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPriceFeedCaller{contract: contract}, nil
}

// NewIPriceFeedTransactor creates a new write-only instance of IPriceFeed, bound to a specific deployed contract.
func NewIPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*IPriceFeedTransactor, error) {
	contract, err := bindIPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPriceFeedTransactor{contract: contract}, nil
}

// NewIPriceFeedFilterer creates a new log filterer instance of IPriceFeed, bound to a specific deployed contract.
func NewIPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*IPriceFeedFilterer, error) {
	contract, err := bindIPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPriceFeedFilterer{contract: contract}, nil
}

// bindIPriceFeed binds a generic wrapper to an already deployed contract.
func bindIPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPriceFeed *IPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPriceFeed.Contract.IPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPriceFeed *IPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPriceFeed.Contract.IPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPriceFeed *IPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPriceFeed.Contract.IPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPriceFeed *IPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPriceFeed *IPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPriceFeed *IPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_IPriceFeed *IPriceFeedCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IPriceFeed.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_IPriceFeed *IPriceFeedSession) Aggregator() (common.Address, error) {
	return _IPriceFeed.Contract.Aggregator(&_IPriceFeed.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_IPriceFeed *IPriceFeedCallerSession) Aggregator() (common.Address, error) {
	return _IPriceFeed.Contract.Aggregator(&_IPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_IPriceFeed *IPriceFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IPriceFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_IPriceFeed *IPriceFeedSession) Description() (string, error) {
	return _IPriceFeed.Contract.Description(&_IPriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_IPriceFeed *IPriceFeedCallerSession) Description() (string, error) {
	return _IPriceFeed.Contract.Description(&_IPriceFeed.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_IPriceFeed *IPriceFeedCaller) GetRoundData(opts *bind.CallOpts, roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPriceFeed.contract.Call(opts, &out, "getRoundData", roundId)

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
// Solidity: function getRoundData(uint80 roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_IPriceFeed *IPriceFeedSession) GetRoundData(roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPriceFeed.Contract.GetRoundData(&_IPriceFeed.CallOpts, roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 roundId) view returns(uint80, int256, uint256, uint256, uint80)
func (_IPriceFeed *IPriceFeedCallerSession) GetRoundData(roundId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPriceFeed.Contract.GetRoundData(&_IPriceFeed.CallOpts, roundId)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_IPriceFeed *IPriceFeedCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPriceFeed.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_IPriceFeed *IPriceFeedSession) LatestAnswer() (*big.Int, error) {
	return _IPriceFeed.Contract.LatestAnswer(&_IPriceFeed.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_IPriceFeed *IPriceFeedCallerSession) LatestAnswer() (*big.Int, error) {
	return _IPriceFeed.Contract.LatestAnswer(&_IPriceFeed.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_IPriceFeed *IPriceFeedCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPriceFeed.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_IPriceFeed *IPriceFeedSession) LatestRound() (*big.Int, error) {
	return _IPriceFeed.Contract.LatestRound(&_IPriceFeed.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint80)
func (_IPriceFeed *IPriceFeedCallerSession) LatestRound() (*big.Int, error) {
	return _IPriceFeed.Contract.LatestRound(&_IPriceFeed.CallOpts)
}
