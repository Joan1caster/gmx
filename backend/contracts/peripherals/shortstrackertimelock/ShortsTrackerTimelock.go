// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shortstrackertimelock

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

// ShortsTrackerTimelockMetaData contains all meta data concerning the ShortsTrackerTimelock contract.
var ShortsTrackerTimelockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_buffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePriceUpdateDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAveragePriceChange\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"ClearAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAveragePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAveragePrice\",\"type\":\"uint256\"}],\"name\":\"GlobalShortAveragePriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"SetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averagePriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"SetAveragePriceUpdateDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isHandler\",\"type\":\"bool\"}],\"name\":\"SetContractHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"}],\"name\":\"SetGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isGlobalShortDataReady\",\"type\":\"bool\"}],\"name\":\"SetIsGlobalShortDataReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAveragePriceChange\",\"type\":\"uint256\"}],\"name\":\"SetMaxAveragePriceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalPendingAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"SignalSetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averagePriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"SignalSetAveragePriceUpdateDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"}],\"name\":\"SignalSetGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalSetHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isGlobalShortDataReady\",\"type\":\"bool\"}],\"name\":\"SignalSetIsGlobalShortDataReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAveragePriceChange\",\"type\":\"uint256\"}],\"name\":\"SignalSetMaxAveragePriceChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BUFFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averagePriceUpdateDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_action\",\"type\":\"bytes32\"}],\"name\":\"cancelAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShortsTracker\",\"name\":\"_shortsTracker\",\"type\":\"address\"}],\"name\":\"disableIsGlobalShortDataReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastUpdated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAveragePriceChange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_averagePriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"setAveragePriceUpdateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_buffer\",\"type\":\"uint256\"}],\"name\":\"setBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setContractHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShortsTracker\",\"name\":\"_shortsTracker\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_averagePrices\",\"type\":\"uint256[]\"}],\"name\":\"setGlobalShortAveragePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_shortsTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShortsTracker\",\"name\":\"_shortsTracker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setIsGlobalShortDataReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAveragePriceChange\",\"type\":\"uint256\"}],\"name\":\"setMaxAveragePriceChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"signalSetAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_averagePriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"signalSetAveragePriceUpdateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_shortsTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"signalSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"signalSetHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIShortsTracker\",\"name\":\"_shortsTracker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"signalSetIsGlobalShortDataReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAveragePriceChange\",\"type\":\"uint256\"}],\"name\":\"signalSetMaxAveragePriceChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ShortsTrackerTimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use ShortsTrackerTimelockMetaData.ABI instead.
var ShortsTrackerTimelockABI = ShortsTrackerTimelockMetaData.ABI

// ShortsTrackerTimelock is an auto generated Go binding around an Ethereum contract.
type ShortsTrackerTimelock struct {
	ShortsTrackerTimelockCaller     // Read-only binding to the contract
	ShortsTrackerTimelockTransactor // Write-only binding to the contract
	ShortsTrackerTimelockFilterer   // Log filterer for contract events
}

// ShortsTrackerTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShortsTrackerTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShortsTrackerTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShortsTrackerTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShortsTrackerTimelockSession struct {
	Contract     *ShortsTrackerTimelock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ShortsTrackerTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShortsTrackerTimelockCallerSession struct {
	Contract *ShortsTrackerTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ShortsTrackerTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShortsTrackerTimelockTransactorSession struct {
	Contract     *ShortsTrackerTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ShortsTrackerTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShortsTrackerTimelockRaw struct {
	Contract *ShortsTrackerTimelock // Generic contract binding to access the raw methods on
}

// ShortsTrackerTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShortsTrackerTimelockCallerRaw struct {
	Contract *ShortsTrackerTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// ShortsTrackerTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShortsTrackerTimelockTransactorRaw struct {
	Contract *ShortsTrackerTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShortsTrackerTimelock creates a new instance of ShortsTrackerTimelock, bound to a specific deployed contract.
func NewShortsTrackerTimelock(address common.Address, backend bind.ContractBackend) (*ShortsTrackerTimelock, error) {
	contract, err := bindShortsTrackerTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelock{ShortsTrackerTimelockCaller: ShortsTrackerTimelockCaller{contract: contract}, ShortsTrackerTimelockTransactor: ShortsTrackerTimelockTransactor{contract: contract}, ShortsTrackerTimelockFilterer: ShortsTrackerTimelockFilterer{contract: contract}}, nil
}

// NewShortsTrackerTimelockCaller creates a new read-only instance of ShortsTrackerTimelock, bound to a specific deployed contract.
func NewShortsTrackerTimelockCaller(address common.Address, caller bind.ContractCaller) (*ShortsTrackerTimelockCaller, error) {
	contract, err := bindShortsTrackerTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockCaller{contract: contract}, nil
}

// NewShortsTrackerTimelockTransactor creates a new write-only instance of ShortsTrackerTimelock, bound to a specific deployed contract.
func NewShortsTrackerTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*ShortsTrackerTimelockTransactor, error) {
	contract, err := bindShortsTrackerTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockTransactor{contract: contract}, nil
}

// NewShortsTrackerTimelockFilterer creates a new log filterer instance of ShortsTrackerTimelock, bound to a specific deployed contract.
func NewShortsTrackerTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*ShortsTrackerTimelockFilterer, error) {
	contract, err := bindShortsTrackerTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockFilterer{contract: contract}, nil
}

// bindShortsTrackerTimelock binds a generic wrapper to an already deployed contract.
func bindShortsTrackerTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShortsTrackerTimelockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShortsTrackerTimelock *ShortsTrackerTimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShortsTrackerTimelock.Contract.ShortsTrackerTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShortsTrackerTimelock *ShortsTrackerTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.ShortsTrackerTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShortsTrackerTimelock *ShortsTrackerTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.ShortsTrackerTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShortsTrackerTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.BASISPOINTSDIVISOR(&_ShortsTrackerTimelock.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.BASISPOINTSDIVISOR(&_ShortsTrackerTimelock.CallOpts)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) MAXBUFFER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "MAX_BUFFER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) MAXBUFFER() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.MAXBUFFER(&_ShortsTrackerTimelock.CallOpts)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) MAXBUFFER() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.MAXBUFFER(&_ShortsTrackerTimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) Admin() (common.Address, error) {
	return _ShortsTrackerTimelock.Contract.Admin(&_ShortsTrackerTimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) Admin() (common.Address, error) {
	return _ShortsTrackerTimelock.Contract.Admin(&_ShortsTrackerTimelock.CallOpts)
}

// AveragePriceUpdateDelay is a free data retrieval call binding the contract method 0x1ab8fe04.
//
// Solidity: function averagePriceUpdateDelay() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) AveragePriceUpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "averagePriceUpdateDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AveragePriceUpdateDelay is a free data retrieval call binding the contract method 0x1ab8fe04.
//
// Solidity: function averagePriceUpdateDelay() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) AveragePriceUpdateDelay() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.AveragePriceUpdateDelay(&_ShortsTrackerTimelock.CallOpts)
}

// AveragePriceUpdateDelay is a free data retrieval call binding the contract method 0x1ab8fe04.
//
// Solidity: function averagePriceUpdateDelay() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) AveragePriceUpdateDelay() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.AveragePriceUpdateDelay(&_ShortsTrackerTimelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) Buffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "buffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) Buffer() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.Buffer(&_ShortsTrackerTimelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) Buffer() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.Buffer(&_ShortsTrackerTimelock.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ShortsTrackerTimelock.Contract.IsHandler(&_ShortsTrackerTimelock.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ShortsTrackerTimelock.Contract.IsHandler(&_ShortsTrackerTimelock.CallOpts, arg0)
}

// LastUpdated is a free data retrieval call binding the contract method 0x0a6f93e6.
//
// Solidity: function lastUpdated(address ) view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) LastUpdated(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "lastUpdated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdated is a free data retrieval call binding the contract method 0x0a6f93e6.
//
// Solidity: function lastUpdated(address ) view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) LastUpdated(arg0 common.Address) (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.LastUpdated(&_ShortsTrackerTimelock.CallOpts, arg0)
}

// LastUpdated is a free data retrieval call binding the contract method 0x0a6f93e6.
//
// Solidity: function lastUpdated(address ) view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) LastUpdated(arg0 common.Address) (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.LastUpdated(&_ShortsTrackerTimelock.CallOpts, arg0)
}

// MaxAveragePriceChange is a free data retrieval call binding the contract method 0x42773c2c.
//
// Solidity: function maxAveragePriceChange() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) MaxAveragePriceChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "maxAveragePriceChange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAveragePriceChange is a free data retrieval call binding the contract method 0x42773c2c.
//
// Solidity: function maxAveragePriceChange() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) MaxAveragePriceChange() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.MaxAveragePriceChange(&_ShortsTrackerTimelock.CallOpts)
}

// MaxAveragePriceChange is a free data retrieval call binding the contract method 0x42773c2c.
//
// Solidity: function maxAveragePriceChange() view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) MaxAveragePriceChange() (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.MaxAveragePriceChange(&_ShortsTrackerTimelock.CallOpts)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCaller) PendingActions(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTimelock.contract.Call(opts, &out, "pendingActions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.PendingActions(&_ShortsTrackerTimelock.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockCallerSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _ShortsTrackerTimelock.Contract.PendingActions(&_ShortsTrackerTimelock.CallOpts, arg0)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) CancelAction(opts *bind.TransactOpts, _action [32]byte) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "cancelAction", _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.CancelAction(&_ShortsTrackerTimelock.TransactOpts, _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.CancelAction(&_ShortsTrackerTimelock.TransactOpts, _action)
}

// DisableIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x627395eb.
//
// Solidity: function disableIsGlobalShortDataReady(address _shortsTracker) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) DisableIsGlobalShortDataReady(opts *bind.TransactOpts, _shortsTracker common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "disableIsGlobalShortDataReady", _shortsTracker)
}

// DisableIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x627395eb.
//
// Solidity: function disableIsGlobalShortDataReady(address _shortsTracker) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) DisableIsGlobalShortDataReady(_shortsTracker common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.DisableIsGlobalShortDataReady(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker)
}

// DisableIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x627395eb.
//
// Solidity: function disableIsGlobalShortDataReady(address _shortsTracker) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) DisableIsGlobalShortDataReady(_shortsTracker common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.DisableIsGlobalShortDataReady(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetAdmin(&_ShortsTrackerTimelock.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetAdmin(&_ShortsTrackerTimelock.TransactOpts, _admin)
}

// SetAveragePriceUpdateDelay is a paid mutator transaction binding the contract method 0x33156639.
//
// Solidity: function setAveragePriceUpdateDelay(uint256 _averagePriceUpdateDelay) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetAveragePriceUpdateDelay(opts *bind.TransactOpts, _averagePriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setAveragePriceUpdateDelay", _averagePriceUpdateDelay)
}

// SetAveragePriceUpdateDelay is a paid mutator transaction binding the contract method 0x33156639.
//
// Solidity: function setAveragePriceUpdateDelay(uint256 _averagePriceUpdateDelay) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetAveragePriceUpdateDelay(_averagePriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetAveragePriceUpdateDelay(&_ShortsTrackerTimelock.TransactOpts, _averagePriceUpdateDelay)
}

// SetAveragePriceUpdateDelay is a paid mutator transaction binding the contract method 0x33156639.
//
// Solidity: function setAveragePriceUpdateDelay(uint256 _averagePriceUpdateDelay) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetAveragePriceUpdateDelay(_averagePriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetAveragePriceUpdateDelay(&_ShortsTrackerTimelock.TransactOpts, _averagePriceUpdateDelay)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetBuffer(opts *bind.TransactOpts, _buffer *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setBuffer", _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetBuffer(&_ShortsTrackerTimelock.TransactOpts, _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetBuffer(&_ShortsTrackerTimelock.TransactOpts, _buffer)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetContractHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setContractHandler", _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetContractHandler(&_ShortsTrackerTimelock.TransactOpts, _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetContractHandler(&_ShortsTrackerTimelock.TransactOpts, _handler, _isActive)
}

// SetGlobalShortAveragePrices is a paid mutator transaction binding the contract method 0x3976e566.
//
// Solidity: function setGlobalShortAveragePrices(address _shortsTracker, address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetGlobalShortAveragePrices(opts *bind.TransactOpts, _shortsTracker common.Address, _tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setGlobalShortAveragePrices", _shortsTracker, _tokens, _averagePrices)
}

// SetGlobalShortAveragePrices is a paid mutator transaction binding the contract method 0x3976e566.
//
// Solidity: function setGlobalShortAveragePrices(address _shortsTracker, address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetGlobalShortAveragePrices(_shortsTracker common.Address, _tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetGlobalShortAveragePrices(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _tokens, _averagePrices)
}

// SetGlobalShortAveragePrices is a paid mutator transaction binding the contract method 0x3976e566.
//
// Solidity: function setGlobalShortAveragePrices(address _shortsTracker, address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetGlobalShortAveragePrices(_shortsTracker common.Address, _tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetGlobalShortAveragePrices(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _tokens, _averagePrices)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _shortsTracker, address _gov) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetGov(opts *bind.TransactOpts, _shortsTracker common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setGov", _shortsTracker, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _shortsTracker, address _gov) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetGov(_shortsTracker common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetGov(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _shortsTracker, address _gov) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetGov(_shortsTracker common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetGov(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x1154e808.
//
// Solidity: function setHandler(address _target, address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetHandler(opts *bind.TransactOpts, _target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setHandler", _target, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x1154e808.
//
// Solidity: function setHandler(address _target, address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetHandler(&_ShortsTrackerTimelock.TransactOpts, _target, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x1154e808.
//
// Solidity: function setHandler(address _target, address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetHandler(&_ShortsTrackerTimelock.TransactOpts, _target, _handler, _isActive)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x2591d483.
//
// Solidity: function setIsGlobalShortDataReady(address _shortsTracker, bool _value) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetIsGlobalShortDataReady(opts *bind.TransactOpts, _shortsTracker common.Address, _value bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setIsGlobalShortDataReady", _shortsTracker, _value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x2591d483.
//
// Solidity: function setIsGlobalShortDataReady(address _shortsTracker, bool _value) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetIsGlobalShortDataReady(_shortsTracker common.Address, _value bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetIsGlobalShortDataReady(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x2591d483.
//
// Solidity: function setIsGlobalShortDataReady(address _shortsTracker, bool _value) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetIsGlobalShortDataReady(_shortsTracker common.Address, _value bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetIsGlobalShortDataReady(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _value)
}

// SetMaxAveragePriceChange is a paid mutator transaction binding the contract method 0x2c166c41.
//
// Solidity: function setMaxAveragePriceChange(uint256 _maxAveragePriceChange) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SetMaxAveragePriceChange(opts *bind.TransactOpts, _maxAveragePriceChange *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "setMaxAveragePriceChange", _maxAveragePriceChange)
}

// SetMaxAveragePriceChange is a paid mutator transaction binding the contract method 0x2c166c41.
//
// Solidity: function setMaxAveragePriceChange(uint256 _maxAveragePriceChange) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SetMaxAveragePriceChange(_maxAveragePriceChange *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetMaxAveragePriceChange(&_ShortsTrackerTimelock.TransactOpts, _maxAveragePriceChange)
}

// SetMaxAveragePriceChange is a paid mutator transaction binding the contract method 0x2c166c41.
//
// Solidity: function setMaxAveragePriceChange(uint256 _maxAveragePriceChange) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SetMaxAveragePriceChange(_maxAveragePriceChange *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SetMaxAveragePriceChange(&_ShortsTrackerTimelock.TransactOpts, _maxAveragePriceChange)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x717cf5d6.
//
// Solidity: function signalSetAdmin(address _admin) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SignalSetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "signalSetAdmin", _admin)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x717cf5d6.
//
// Solidity: function signalSetAdmin(address _admin) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SignalSetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetAdmin(&_ShortsTrackerTimelock.TransactOpts, _admin)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x717cf5d6.
//
// Solidity: function signalSetAdmin(address _admin) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SignalSetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetAdmin(&_ShortsTrackerTimelock.TransactOpts, _admin)
}

// SignalSetAveragePriceUpdateDelay is a paid mutator transaction binding the contract method 0x42588a67.
//
// Solidity: function signalSetAveragePriceUpdateDelay(uint256 _averagePriceUpdateDelay) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SignalSetAveragePriceUpdateDelay(opts *bind.TransactOpts, _averagePriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "signalSetAveragePriceUpdateDelay", _averagePriceUpdateDelay)
}

// SignalSetAveragePriceUpdateDelay is a paid mutator transaction binding the contract method 0x42588a67.
//
// Solidity: function signalSetAveragePriceUpdateDelay(uint256 _averagePriceUpdateDelay) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SignalSetAveragePriceUpdateDelay(_averagePriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetAveragePriceUpdateDelay(&_ShortsTrackerTimelock.TransactOpts, _averagePriceUpdateDelay)
}

// SignalSetAveragePriceUpdateDelay is a paid mutator transaction binding the contract method 0x42588a67.
//
// Solidity: function signalSetAveragePriceUpdateDelay(uint256 _averagePriceUpdateDelay) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SignalSetAveragePriceUpdateDelay(_averagePriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetAveragePriceUpdateDelay(&_ShortsTrackerTimelock.TransactOpts, _averagePriceUpdateDelay)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _shortsTracker, address _gov) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SignalSetGov(opts *bind.TransactOpts, _shortsTracker common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "signalSetGov", _shortsTracker, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _shortsTracker, address _gov) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SignalSetGov(_shortsTracker common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetGov(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _shortsTracker, address _gov) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SignalSetGov(_shortsTracker common.Address, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetGov(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _gov)
}

// SignalSetHandler is a paid mutator transaction binding the contract method 0x24ccbe30.
//
// Solidity: function signalSetHandler(address _target, address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SignalSetHandler(opts *bind.TransactOpts, _target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "signalSetHandler", _target, _handler, _isActive)
}

// SignalSetHandler is a paid mutator transaction binding the contract method 0x24ccbe30.
//
// Solidity: function signalSetHandler(address _target, address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SignalSetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetHandler(&_ShortsTrackerTimelock.TransactOpts, _target, _handler, _isActive)
}

// SignalSetHandler is a paid mutator transaction binding the contract method 0x24ccbe30.
//
// Solidity: function signalSetHandler(address _target, address _handler, bool _isActive) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SignalSetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetHandler(&_ShortsTrackerTimelock.TransactOpts, _target, _handler, _isActive)
}

// SignalSetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x4882fa13.
//
// Solidity: function signalSetIsGlobalShortDataReady(address _shortsTracker, bool _value) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SignalSetIsGlobalShortDataReady(opts *bind.TransactOpts, _shortsTracker common.Address, _value bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "signalSetIsGlobalShortDataReady", _shortsTracker, _value)
}

// SignalSetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x4882fa13.
//
// Solidity: function signalSetIsGlobalShortDataReady(address _shortsTracker, bool _value) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SignalSetIsGlobalShortDataReady(_shortsTracker common.Address, _value bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetIsGlobalShortDataReady(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _value)
}

// SignalSetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x4882fa13.
//
// Solidity: function signalSetIsGlobalShortDataReady(address _shortsTracker, bool _value) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SignalSetIsGlobalShortDataReady(_shortsTracker common.Address, _value bool) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetIsGlobalShortDataReady(&_ShortsTrackerTimelock.TransactOpts, _shortsTracker, _value)
}

// SignalSetMaxAveragePriceChange is a paid mutator transaction binding the contract method 0x5c3bca36.
//
// Solidity: function signalSetMaxAveragePriceChange(uint256 _maxAveragePriceChange) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactor) SignalSetMaxAveragePriceChange(opts *bind.TransactOpts, _maxAveragePriceChange *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.contract.Transact(opts, "signalSetMaxAveragePriceChange", _maxAveragePriceChange)
}

// SignalSetMaxAveragePriceChange is a paid mutator transaction binding the contract method 0x5c3bca36.
//
// Solidity: function signalSetMaxAveragePriceChange(uint256 _maxAveragePriceChange) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockSession) SignalSetMaxAveragePriceChange(_maxAveragePriceChange *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetMaxAveragePriceChange(&_ShortsTrackerTimelock.TransactOpts, _maxAveragePriceChange)
}

// SignalSetMaxAveragePriceChange is a paid mutator transaction binding the contract method 0x5c3bca36.
//
// Solidity: function signalSetMaxAveragePriceChange(uint256 _maxAveragePriceChange) returns()
func (_ShortsTrackerTimelock *ShortsTrackerTimelockTransactorSession) SignalSetMaxAveragePriceChange(_maxAveragePriceChange *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTimelock.Contract.SignalSetMaxAveragePriceChange(&_ShortsTrackerTimelock.TransactOpts, _maxAveragePriceChange)
}

// ShortsTrackerTimelockClearActionIterator is returned from FilterClearAction and is used to iterate over the raw logs and unpacked data for ClearAction events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockClearActionIterator struct {
	Event *ShortsTrackerTimelockClearAction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockClearActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockClearAction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockClearAction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockClearActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockClearActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockClearAction represents a ClearAction event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockClearAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClearAction is a free log retrieval operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterClearAction(opts *bind.FilterOpts) (*ShortsTrackerTimelockClearActionIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockClearActionIterator{contract: _ShortsTrackerTimelock.contract, event: "ClearAction", logs: logs, sub: sub}, nil
}

// WatchClearAction is a free log subscription operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchClearAction(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockClearAction) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockClearAction)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClearAction is a log parse operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseClearAction(log types.Log) (*ShortsTrackerTimelockClearAction, error) {
	event := new(ShortsTrackerTimelockClearAction)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockGlobalShortAveragePriceUpdatedIterator is returned from FilterGlobalShortAveragePriceUpdated and is used to iterate over the raw logs and unpacked data for GlobalShortAveragePriceUpdated events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockGlobalShortAveragePriceUpdatedIterator struct {
	Event *ShortsTrackerTimelockGlobalShortAveragePriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockGlobalShortAveragePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockGlobalShortAveragePriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockGlobalShortAveragePriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockGlobalShortAveragePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockGlobalShortAveragePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockGlobalShortAveragePriceUpdated represents a GlobalShortAveragePriceUpdated event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockGlobalShortAveragePriceUpdated struct {
	Token           common.Address
	OldAveragePrice *big.Int
	NewAveragePrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGlobalShortAveragePriceUpdated is a free log retrieval operation binding the contract event 0xa6a58ff784b1e07988de90568885d7bbd8e5485c3ec4344195c7306108553256.
//
// Solidity: event GlobalShortAveragePriceUpdated(address indexed token, uint256 oldAveragePrice, uint256 newAveragePrice)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterGlobalShortAveragePriceUpdated(opts *bind.FilterOpts, token []common.Address) (*ShortsTrackerTimelockGlobalShortAveragePriceUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "GlobalShortAveragePriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockGlobalShortAveragePriceUpdatedIterator{contract: _ShortsTrackerTimelock.contract, event: "GlobalShortAveragePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalShortAveragePriceUpdated is a free log subscription operation binding the contract event 0xa6a58ff784b1e07988de90568885d7bbd8e5485c3ec4344195c7306108553256.
//
// Solidity: event GlobalShortAveragePriceUpdated(address indexed token, uint256 oldAveragePrice, uint256 newAveragePrice)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchGlobalShortAveragePriceUpdated(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockGlobalShortAveragePriceUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "GlobalShortAveragePriceUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockGlobalShortAveragePriceUpdated)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "GlobalShortAveragePriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGlobalShortAveragePriceUpdated is a log parse operation binding the contract event 0xa6a58ff784b1e07988de90568885d7bbd8e5485c3ec4344195c7306108553256.
//
// Solidity: event GlobalShortAveragePriceUpdated(address indexed token, uint256 oldAveragePrice, uint256 newAveragePrice)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseGlobalShortAveragePriceUpdated(log types.Log) (*ShortsTrackerTimelockGlobalShortAveragePriceUpdated, error) {
	event := new(ShortsTrackerTimelockGlobalShortAveragePriceUpdated)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "GlobalShortAveragePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetAdminIterator struct {
	Event *ShortsTrackerTimelockSetAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSetAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSetAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSetAdmin represents a SetAdmin event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*ShortsTrackerTimelockSetAdminIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSetAdminIterator{contract: _ShortsTrackerTimelock.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSetAdmin) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSetAdmin)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSetAdmin(log types.Log) (*ShortsTrackerTimelockSetAdmin, error) {
	event := new(ShortsTrackerTimelockSetAdmin)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSetAveragePriceUpdateDelayIterator is returned from FilterSetAveragePriceUpdateDelay and is used to iterate over the raw logs and unpacked data for SetAveragePriceUpdateDelay events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetAveragePriceUpdateDelayIterator struct {
	Event *ShortsTrackerTimelockSetAveragePriceUpdateDelay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSetAveragePriceUpdateDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSetAveragePriceUpdateDelay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSetAveragePriceUpdateDelay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSetAveragePriceUpdateDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSetAveragePriceUpdateDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSetAveragePriceUpdateDelay represents a SetAveragePriceUpdateDelay event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetAveragePriceUpdateDelay struct {
	AveragePriceUpdateDelay *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSetAveragePriceUpdateDelay is a free log retrieval operation binding the contract event 0xa043b6b847902bb85c2a2575e9bce3a077bcaad8ec888ac941f50a6f4acdbb3d.
//
// Solidity: event SetAveragePriceUpdateDelay(uint256 averagePriceUpdateDelay)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSetAveragePriceUpdateDelay(opts *bind.FilterOpts) (*ShortsTrackerTimelockSetAveragePriceUpdateDelayIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SetAveragePriceUpdateDelay")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSetAveragePriceUpdateDelayIterator{contract: _ShortsTrackerTimelock.contract, event: "SetAveragePriceUpdateDelay", logs: logs, sub: sub}, nil
}

// WatchSetAveragePriceUpdateDelay is a free log subscription operation binding the contract event 0xa043b6b847902bb85c2a2575e9bce3a077bcaad8ec888ac941f50a6f4acdbb3d.
//
// Solidity: event SetAveragePriceUpdateDelay(uint256 averagePriceUpdateDelay)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSetAveragePriceUpdateDelay(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSetAveragePriceUpdateDelay) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SetAveragePriceUpdateDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSetAveragePriceUpdateDelay)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetAveragePriceUpdateDelay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAveragePriceUpdateDelay is a log parse operation binding the contract event 0xa043b6b847902bb85c2a2575e9bce3a077bcaad8ec888ac941f50a6f4acdbb3d.
//
// Solidity: event SetAveragePriceUpdateDelay(uint256 averagePriceUpdateDelay)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSetAveragePriceUpdateDelay(log types.Log) (*ShortsTrackerTimelockSetAveragePriceUpdateDelay, error) {
	event := new(ShortsTrackerTimelockSetAveragePriceUpdateDelay)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetAveragePriceUpdateDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSetContractHandlerIterator is returned from FilterSetContractHandler and is used to iterate over the raw logs and unpacked data for SetContractHandler events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetContractHandlerIterator struct {
	Event *ShortsTrackerTimelockSetContractHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSetContractHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSetContractHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSetContractHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSetContractHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSetContractHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSetContractHandler represents a SetContractHandler event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetContractHandler struct {
	Handler   common.Address
	IsHandler bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetContractHandler is a free log retrieval operation binding the contract event 0x0c7e57410cc008b0b1abd8a03240abc3921d41be4ca09611330825622d5866cf.
//
// Solidity: event SetContractHandler(address indexed handler, bool isHandler)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSetContractHandler(opts *bind.FilterOpts, handler []common.Address) (*ShortsTrackerTimelockSetContractHandlerIterator, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SetContractHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSetContractHandlerIterator{contract: _ShortsTrackerTimelock.contract, event: "SetContractHandler", logs: logs, sub: sub}, nil
}

// WatchSetContractHandler is a free log subscription operation binding the contract event 0x0c7e57410cc008b0b1abd8a03240abc3921d41be4ca09611330825622d5866cf.
//
// Solidity: event SetContractHandler(address indexed handler, bool isHandler)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSetContractHandler(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSetContractHandler, handler []common.Address) (event.Subscription, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SetContractHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSetContractHandler)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetContractHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetContractHandler is a log parse operation binding the contract event 0x0c7e57410cc008b0b1abd8a03240abc3921d41be4ca09611330825622d5866cf.
//
// Solidity: event SetContractHandler(address indexed handler, bool isHandler)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSetContractHandler(log types.Log) (*ShortsTrackerTimelockSetContractHandler, error) {
	event := new(ShortsTrackerTimelockSetContractHandler)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetContractHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSetGovIterator is returned from FilterSetGov and is used to iterate over the raw logs and unpacked data for SetGov events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetGovIterator struct {
	Event *ShortsTrackerTimelockSetGov // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSetGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSetGov)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSetGov)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSetGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSetGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSetGov represents a SetGov event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetGov struct {
	Target common.Address
	Gov    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetGov is a free log retrieval operation binding the contract event 0x53351836099c03ffc3b1727d8abd4b0222afa87d4ed76ae3102d51369ef7f785.
//
// Solidity: event SetGov(address target, address gov)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSetGov(opts *bind.FilterOpts) (*ShortsTrackerTimelockSetGovIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SetGov")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSetGovIterator{contract: _ShortsTrackerTimelock.contract, event: "SetGov", logs: logs, sub: sub}, nil
}

// WatchSetGov is a free log subscription operation binding the contract event 0x53351836099c03ffc3b1727d8abd4b0222afa87d4ed76ae3102d51369ef7f785.
//
// Solidity: event SetGov(address target, address gov)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSetGov(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSetGov) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SetGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSetGov)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetGov", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetGov is a log parse operation binding the contract event 0x53351836099c03ffc3b1727d8abd4b0222afa87d4ed76ae3102d51369ef7f785.
//
// Solidity: event SetGov(address target, address gov)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSetGov(log types.Log) (*ShortsTrackerTimelockSetGov, error) {
	event := new(ShortsTrackerTimelockSetGov)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSetIsGlobalShortDataReadyIterator is returned from FilterSetIsGlobalShortDataReady and is used to iterate over the raw logs and unpacked data for SetIsGlobalShortDataReady events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetIsGlobalShortDataReadyIterator struct {
	Event *ShortsTrackerTimelockSetIsGlobalShortDataReady // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSetIsGlobalShortDataReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSetIsGlobalShortDataReady)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSetIsGlobalShortDataReady)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSetIsGlobalShortDataReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSetIsGlobalShortDataReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSetIsGlobalShortDataReady represents a SetIsGlobalShortDataReady event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetIsGlobalShortDataReady struct {
	Target                 common.Address
	IsGlobalShortDataReady bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetIsGlobalShortDataReady is a free log retrieval operation binding the contract event 0x177de4361fc37a0e8defe8ac0c685d3472789b507959897eec32350c9d2c3666.
//
// Solidity: event SetIsGlobalShortDataReady(address target, bool isGlobalShortDataReady)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSetIsGlobalShortDataReady(opts *bind.FilterOpts) (*ShortsTrackerTimelockSetIsGlobalShortDataReadyIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SetIsGlobalShortDataReady")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSetIsGlobalShortDataReadyIterator{contract: _ShortsTrackerTimelock.contract, event: "SetIsGlobalShortDataReady", logs: logs, sub: sub}, nil
}

// WatchSetIsGlobalShortDataReady is a free log subscription operation binding the contract event 0x177de4361fc37a0e8defe8ac0c685d3472789b507959897eec32350c9d2c3666.
//
// Solidity: event SetIsGlobalShortDataReady(address target, bool isGlobalShortDataReady)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSetIsGlobalShortDataReady(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSetIsGlobalShortDataReady) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SetIsGlobalShortDataReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSetIsGlobalShortDataReady)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetIsGlobalShortDataReady", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetIsGlobalShortDataReady is a log parse operation binding the contract event 0x177de4361fc37a0e8defe8ac0c685d3472789b507959897eec32350c9d2c3666.
//
// Solidity: event SetIsGlobalShortDataReady(address target, bool isGlobalShortDataReady)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSetIsGlobalShortDataReady(log types.Log) (*ShortsTrackerTimelockSetIsGlobalShortDataReady, error) {
	event := new(ShortsTrackerTimelockSetIsGlobalShortDataReady)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetIsGlobalShortDataReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSetMaxAveragePriceChangeIterator is returned from FilterSetMaxAveragePriceChange and is used to iterate over the raw logs and unpacked data for SetMaxAveragePriceChange events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetMaxAveragePriceChangeIterator struct {
	Event *ShortsTrackerTimelockSetMaxAveragePriceChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSetMaxAveragePriceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSetMaxAveragePriceChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSetMaxAveragePriceChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSetMaxAveragePriceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSetMaxAveragePriceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSetMaxAveragePriceChange represents a SetMaxAveragePriceChange event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSetMaxAveragePriceChange struct {
	MaxAveragePriceChange *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMaxAveragePriceChange is a free log retrieval operation binding the contract event 0xe846c28c10eef345d8d4fea8b403bfac624c32c3e64ce0b666f23404f2edb514.
//
// Solidity: event SetMaxAveragePriceChange(uint256 maxAveragePriceChange)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSetMaxAveragePriceChange(opts *bind.FilterOpts) (*ShortsTrackerTimelockSetMaxAveragePriceChangeIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SetMaxAveragePriceChange")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSetMaxAveragePriceChangeIterator{contract: _ShortsTrackerTimelock.contract, event: "SetMaxAveragePriceChange", logs: logs, sub: sub}, nil
}

// WatchSetMaxAveragePriceChange is a free log subscription operation binding the contract event 0xe846c28c10eef345d8d4fea8b403bfac624c32c3e64ce0b666f23404f2edb514.
//
// Solidity: event SetMaxAveragePriceChange(uint256 maxAveragePriceChange)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSetMaxAveragePriceChange(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSetMaxAveragePriceChange) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SetMaxAveragePriceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSetMaxAveragePriceChange)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetMaxAveragePriceChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMaxAveragePriceChange is a log parse operation binding the contract event 0xe846c28c10eef345d8d4fea8b403bfac624c32c3e64ce0b666f23404f2edb514.
//
// Solidity: event SetMaxAveragePriceChange(uint256 maxAveragePriceChange)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSetMaxAveragePriceChange(log types.Log) (*ShortsTrackerTimelockSetMaxAveragePriceChange, error) {
	event := new(ShortsTrackerTimelockSetMaxAveragePriceChange)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SetMaxAveragePriceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSignalPendingActionIterator is returned from FilterSignalPendingAction and is used to iterate over the raw logs and unpacked data for SignalPendingAction events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalPendingActionIterator struct {
	Event *ShortsTrackerTimelockSignalPendingAction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSignalPendingActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSignalPendingAction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSignalPendingAction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSignalPendingActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSignalPendingActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSignalPendingAction represents a SignalPendingAction event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalPendingAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalPendingAction is a free log retrieval operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSignalPendingAction(opts *bind.FilterOpts) (*ShortsTrackerTimelockSignalPendingActionIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSignalPendingActionIterator{contract: _ShortsTrackerTimelock.contract, event: "SignalPendingAction", logs: logs, sub: sub}, nil
}

// WatchSignalPendingAction is a free log subscription operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSignalPendingAction(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSignalPendingAction) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSignalPendingAction)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignalPendingAction is a log parse operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSignalPendingAction(log types.Log) (*ShortsTrackerTimelockSignalPendingAction, error) {
	event := new(ShortsTrackerTimelockSignalPendingAction)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSignalSetAdminIterator is returned from FilterSignalSetAdmin and is used to iterate over the raw logs and unpacked data for SignalSetAdmin events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetAdminIterator struct {
	Event *ShortsTrackerTimelockSignalSetAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSignalSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSignalSetAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSignalSetAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSignalSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSignalSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSignalSetAdmin represents a SignalSetAdmin event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSignalSetAdmin is a free log retrieval operation binding the contract event 0xbccacda50126e9efa231593d2a9d4394d7cda7909ab2e14d2ae878ab2a306b3f.
//
// Solidity: event SignalSetAdmin(address admin)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSignalSetAdmin(opts *bind.FilterOpts) (*ShortsTrackerTimelockSignalSetAdminIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SignalSetAdmin")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSignalSetAdminIterator{contract: _ShortsTrackerTimelock.contract, event: "SignalSetAdmin", logs: logs, sub: sub}, nil
}

// WatchSignalSetAdmin is a free log subscription operation binding the contract event 0xbccacda50126e9efa231593d2a9d4394d7cda7909ab2e14d2ae878ab2a306b3f.
//
// Solidity: event SignalSetAdmin(address admin)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSignalSetAdmin(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSignalSetAdmin) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SignalSetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSignalSetAdmin)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignalSetAdmin is a log parse operation binding the contract event 0xbccacda50126e9efa231593d2a9d4394d7cda7909ab2e14d2ae878ab2a306b3f.
//
// Solidity: event SignalSetAdmin(address admin)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSignalSetAdmin(log types.Log) (*ShortsTrackerTimelockSignalSetAdmin, error) {
	event := new(ShortsTrackerTimelockSignalSetAdmin)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSignalSetAveragePriceUpdateDelayIterator is returned from FilterSignalSetAveragePriceUpdateDelay and is used to iterate over the raw logs and unpacked data for SignalSetAveragePriceUpdateDelay events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetAveragePriceUpdateDelayIterator struct {
	Event *ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSignalSetAveragePriceUpdateDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSignalSetAveragePriceUpdateDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSignalSetAveragePriceUpdateDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay represents a SignalSetAveragePriceUpdateDelay event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay struct {
	AveragePriceUpdateDelay *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSignalSetAveragePriceUpdateDelay is a free log retrieval operation binding the contract event 0xc226c48cee6a025ab009e01969a2e89f1ce4d455cfa2900745b2237cb0bb8d61.
//
// Solidity: event SignalSetAveragePriceUpdateDelay(uint256 averagePriceUpdateDelay)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSignalSetAveragePriceUpdateDelay(opts *bind.FilterOpts) (*ShortsTrackerTimelockSignalSetAveragePriceUpdateDelayIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SignalSetAveragePriceUpdateDelay")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSignalSetAveragePriceUpdateDelayIterator{contract: _ShortsTrackerTimelock.contract, event: "SignalSetAveragePriceUpdateDelay", logs: logs, sub: sub}, nil
}

// WatchSignalSetAveragePriceUpdateDelay is a free log subscription operation binding the contract event 0xc226c48cee6a025ab009e01969a2e89f1ce4d455cfa2900745b2237cb0bb8d61.
//
// Solidity: event SignalSetAveragePriceUpdateDelay(uint256 averagePriceUpdateDelay)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSignalSetAveragePriceUpdateDelay(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SignalSetAveragePriceUpdateDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetAveragePriceUpdateDelay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignalSetAveragePriceUpdateDelay is a log parse operation binding the contract event 0xc226c48cee6a025ab009e01969a2e89f1ce4d455cfa2900745b2237cb0bb8d61.
//
// Solidity: event SignalSetAveragePriceUpdateDelay(uint256 averagePriceUpdateDelay)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSignalSetAveragePriceUpdateDelay(log types.Log) (*ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay, error) {
	event := new(ShortsTrackerTimelockSignalSetAveragePriceUpdateDelay)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetAveragePriceUpdateDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSignalSetGovIterator is returned from FilterSignalSetGov and is used to iterate over the raw logs and unpacked data for SignalSetGov events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetGovIterator struct {
	Event *ShortsTrackerTimelockSignalSetGov // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSignalSetGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSignalSetGov)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSignalSetGov)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSignalSetGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSignalSetGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSignalSetGov represents a SignalSetGov event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetGov struct {
	Target common.Address
	Gov    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalSetGov is a free log retrieval operation binding the contract event 0x84f5d90471bb5811423318895e49ef4de52f5a6f4dac67af85e4202410bb66a4.
//
// Solidity: event SignalSetGov(address target, address gov)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSignalSetGov(opts *bind.FilterOpts) (*ShortsTrackerTimelockSignalSetGovIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSignalSetGovIterator{contract: _ShortsTrackerTimelock.contract, event: "SignalSetGov", logs: logs, sub: sub}, nil
}

// WatchSignalSetGov is a free log subscription operation binding the contract event 0x84f5d90471bb5811423318895e49ef4de52f5a6f4dac67af85e4202410bb66a4.
//
// Solidity: event SignalSetGov(address target, address gov)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSignalSetGov(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSignalSetGov) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSignalSetGov)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignalSetGov is a log parse operation binding the contract event 0x84f5d90471bb5811423318895e49ef4de52f5a6f4dac67af85e4202410bb66a4.
//
// Solidity: event SignalSetGov(address target, address gov)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSignalSetGov(log types.Log) (*ShortsTrackerTimelockSignalSetGov, error) {
	event := new(ShortsTrackerTimelockSignalSetGov)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSignalSetHandlerIterator is returned from FilterSignalSetHandler and is used to iterate over the raw logs and unpacked data for SignalSetHandler events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetHandlerIterator struct {
	Event *ShortsTrackerTimelockSignalSetHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSignalSetHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSignalSetHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSignalSetHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSignalSetHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSignalSetHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSignalSetHandler represents a SignalSetHandler event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetHandler struct {
	Target   common.Address
	Handler  common.Address
	IsActive bool
	Action   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalSetHandler is a free log retrieval operation binding the contract event 0x1929c4e13b0dbbad7856b9ce1fc9dca98c7bf7cedd56e22c04dd60ad1d34fe4b.
//
// Solidity: event SignalSetHandler(address target, address handler, bool isActive, bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSignalSetHandler(opts *bind.FilterOpts) (*ShortsTrackerTimelockSignalSetHandlerIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SignalSetHandler")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSignalSetHandlerIterator{contract: _ShortsTrackerTimelock.contract, event: "SignalSetHandler", logs: logs, sub: sub}, nil
}

// WatchSignalSetHandler is a free log subscription operation binding the contract event 0x1929c4e13b0dbbad7856b9ce1fc9dca98c7bf7cedd56e22c04dd60ad1d34fe4b.
//
// Solidity: event SignalSetHandler(address target, address handler, bool isActive, bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSignalSetHandler(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSignalSetHandler) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SignalSetHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSignalSetHandler)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignalSetHandler is a log parse operation binding the contract event 0x1929c4e13b0dbbad7856b9ce1fc9dca98c7bf7cedd56e22c04dd60ad1d34fe4b.
//
// Solidity: event SignalSetHandler(address target, address handler, bool isActive, bytes32 action)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSignalSetHandler(log types.Log) (*ShortsTrackerTimelockSignalSetHandler, error) {
	event := new(ShortsTrackerTimelockSignalSetHandler)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSignalSetIsGlobalShortDataReadyIterator is returned from FilterSignalSetIsGlobalShortDataReady and is used to iterate over the raw logs and unpacked data for SignalSetIsGlobalShortDataReady events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetIsGlobalShortDataReadyIterator struct {
	Event *ShortsTrackerTimelockSignalSetIsGlobalShortDataReady // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSignalSetIsGlobalShortDataReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSignalSetIsGlobalShortDataReady)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSignalSetIsGlobalShortDataReady)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSignalSetIsGlobalShortDataReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSignalSetIsGlobalShortDataReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSignalSetIsGlobalShortDataReady represents a SignalSetIsGlobalShortDataReady event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetIsGlobalShortDataReady struct {
	Target                 common.Address
	IsGlobalShortDataReady bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignalSetIsGlobalShortDataReady is a free log retrieval operation binding the contract event 0x362e5bba25647700de1a7562fc07672df837bef17125e07a6f79abea31ee5a48.
//
// Solidity: event SignalSetIsGlobalShortDataReady(address target, bool isGlobalShortDataReady)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSignalSetIsGlobalShortDataReady(opts *bind.FilterOpts) (*ShortsTrackerTimelockSignalSetIsGlobalShortDataReadyIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SignalSetIsGlobalShortDataReady")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSignalSetIsGlobalShortDataReadyIterator{contract: _ShortsTrackerTimelock.contract, event: "SignalSetIsGlobalShortDataReady", logs: logs, sub: sub}, nil
}

// WatchSignalSetIsGlobalShortDataReady is a free log subscription operation binding the contract event 0x362e5bba25647700de1a7562fc07672df837bef17125e07a6f79abea31ee5a48.
//
// Solidity: event SignalSetIsGlobalShortDataReady(address target, bool isGlobalShortDataReady)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSignalSetIsGlobalShortDataReady(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSignalSetIsGlobalShortDataReady) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SignalSetIsGlobalShortDataReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSignalSetIsGlobalShortDataReady)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetIsGlobalShortDataReady", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignalSetIsGlobalShortDataReady is a log parse operation binding the contract event 0x362e5bba25647700de1a7562fc07672df837bef17125e07a6f79abea31ee5a48.
//
// Solidity: event SignalSetIsGlobalShortDataReady(address target, bool isGlobalShortDataReady)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSignalSetIsGlobalShortDataReady(log types.Log) (*ShortsTrackerTimelockSignalSetIsGlobalShortDataReady, error) {
	event := new(ShortsTrackerTimelockSignalSetIsGlobalShortDataReady)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetIsGlobalShortDataReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShortsTrackerTimelockSignalSetMaxAveragePriceChangeIterator is returned from FilterSignalSetMaxAveragePriceChange and is used to iterate over the raw logs and unpacked data for SignalSetMaxAveragePriceChange events raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetMaxAveragePriceChangeIterator struct {
	Event *ShortsTrackerTimelockSignalSetMaxAveragePriceChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ShortsTrackerTimelockSignalSetMaxAveragePriceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTimelockSignalSetMaxAveragePriceChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ShortsTrackerTimelockSignalSetMaxAveragePriceChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ShortsTrackerTimelockSignalSetMaxAveragePriceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTimelockSignalSetMaxAveragePriceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTimelockSignalSetMaxAveragePriceChange represents a SignalSetMaxAveragePriceChange event raised by the ShortsTrackerTimelock contract.
type ShortsTrackerTimelockSignalSetMaxAveragePriceChange struct {
	MaxAveragePriceChange *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSignalSetMaxAveragePriceChange is a free log retrieval operation binding the contract event 0x21f68305d262b9ce144019f5d9f755a2dd1b5c2d35996cd4b649c285d7525cd1.
//
// Solidity: event SignalSetMaxAveragePriceChange(uint256 maxAveragePriceChange)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) FilterSignalSetMaxAveragePriceChange(opts *bind.FilterOpts) (*ShortsTrackerTimelockSignalSetMaxAveragePriceChangeIterator, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.FilterLogs(opts, "SignalSetMaxAveragePriceChange")
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTimelockSignalSetMaxAveragePriceChangeIterator{contract: _ShortsTrackerTimelock.contract, event: "SignalSetMaxAveragePriceChange", logs: logs, sub: sub}, nil
}

// WatchSignalSetMaxAveragePriceChange is a free log subscription operation binding the contract event 0x21f68305d262b9ce144019f5d9f755a2dd1b5c2d35996cd4b649c285d7525cd1.
//
// Solidity: event SignalSetMaxAveragePriceChange(uint256 maxAveragePriceChange)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) WatchSignalSetMaxAveragePriceChange(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTimelockSignalSetMaxAveragePriceChange) (event.Subscription, error) {

	logs, sub, err := _ShortsTrackerTimelock.contract.WatchLogs(opts, "SignalSetMaxAveragePriceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTimelockSignalSetMaxAveragePriceChange)
				if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetMaxAveragePriceChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignalSetMaxAveragePriceChange is a log parse operation binding the contract event 0x21f68305d262b9ce144019f5d9f755a2dd1b5c2d35996cd4b649c285d7525cd1.
//
// Solidity: event SignalSetMaxAveragePriceChange(uint256 maxAveragePriceChange)
func (_ShortsTrackerTimelock *ShortsTrackerTimelockFilterer) ParseSignalSetMaxAveragePriceChange(log types.Log) (*ShortsTrackerTimelockSignalSetMaxAveragePriceChange, error) {
	event := new(ShortsTrackerTimelockSignalSetMaxAveragePriceChange)
	if err := _ShortsTrackerTimelock.contract.UnpackLog(event, "SignalSetMaxAveragePriceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
