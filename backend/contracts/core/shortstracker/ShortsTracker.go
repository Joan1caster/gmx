// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shortstracker

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

// ShortsTrackerMetaData contains all meta data concerning the ShortsTracker contract.
var ShortsTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalShortSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalShortAveragePrice\",\"type\":\"uint256\"}],\"name\":\"GlobalShortDataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_INT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delta\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_realisedPnl\",\"type\":\"int256\"}],\"name\":\"_getNextGlobalAveragePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"data\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getGlobalShortDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"getNextGlobalShortData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"getRealisedPnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"globalShortAveragePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isGlobalShortDataReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_averagePrices\",\"type\":\"uint256[]\"}],\"name\":\"setInitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setIsGlobalShortDataReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_markPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"updateGlobalShortData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ShortsTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use ShortsTrackerMetaData.ABI instead.
var ShortsTrackerABI = ShortsTrackerMetaData.ABI

// ShortsTracker is an auto generated Go binding around an Ethereum contract.
type ShortsTracker struct {
	ShortsTrackerCaller     // Read-only binding to the contract
	ShortsTrackerTransactor // Write-only binding to the contract
	ShortsTrackerFilterer   // Log filterer for contract events
}

// ShortsTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShortsTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShortsTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShortsTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShortsTrackerSession struct {
	Contract     *ShortsTracker    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShortsTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShortsTrackerCallerSession struct {
	Contract *ShortsTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ShortsTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShortsTrackerTransactorSession struct {
	Contract     *ShortsTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ShortsTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShortsTrackerRaw struct {
	Contract *ShortsTracker // Generic contract binding to access the raw methods on
}

// ShortsTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShortsTrackerCallerRaw struct {
	Contract *ShortsTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// ShortsTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShortsTrackerTransactorRaw struct {
	Contract *ShortsTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShortsTracker creates a new instance of ShortsTracker, bound to a specific deployed contract.
func NewShortsTracker(address common.Address, backend bind.ContractBackend) (*ShortsTracker, error) {
	contract, err := bindShortsTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShortsTracker{ShortsTrackerCaller: ShortsTrackerCaller{contract: contract}, ShortsTrackerTransactor: ShortsTrackerTransactor{contract: contract}, ShortsTrackerFilterer: ShortsTrackerFilterer{contract: contract}}, nil
}

// NewShortsTrackerCaller creates a new read-only instance of ShortsTracker, bound to a specific deployed contract.
func NewShortsTrackerCaller(address common.Address, caller bind.ContractCaller) (*ShortsTrackerCaller, error) {
	contract, err := bindShortsTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerCaller{contract: contract}, nil
}

// NewShortsTrackerTransactor creates a new write-only instance of ShortsTracker, bound to a specific deployed contract.
func NewShortsTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*ShortsTrackerTransactor, error) {
	contract, err := bindShortsTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTransactor{contract: contract}, nil
}

// NewShortsTrackerFilterer creates a new log filterer instance of ShortsTracker, bound to a specific deployed contract.
func NewShortsTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*ShortsTrackerFilterer, error) {
	contract, err := bindShortsTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerFilterer{contract: contract}, nil
}

// bindShortsTracker binds a generic wrapper to an already deployed contract.
func bindShortsTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShortsTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShortsTracker *ShortsTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShortsTracker.Contract.ShortsTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShortsTracker *ShortsTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShortsTracker.Contract.ShortsTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShortsTracker *ShortsTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShortsTracker.Contract.ShortsTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShortsTracker *ShortsTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShortsTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShortsTracker *ShortsTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShortsTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShortsTracker *ShortsTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShortsTracker.Contract.contract.Transact(opts, method, params...)
}

// MAXINT256 is a free data retrieval call binding the contract method 0x122e7b07.
//
// Solidity: function MAX_INT256() view returns(uint256)
func (_ShortsTracker *ShortsTrackerCaller) MAXINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "MAX_INT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINT256 is a free data retrieval call binding the contract method 0x122e7b07.
//
// Solidity: function MAX_INT256() view returns(uint256)
func (_ShortsTracker *ShortsTrackerSession) MAXINT256() (*big.Int, error) {
	return _ShortsTracker.Contract.MAXINT256(&_ShortsTracker.CallOpts)
}

// MAXINT256 is a free data retrieval call binding the contract method 0x122e7b07.
//
// Solidity: function MAX_INT256() view returns(uint256)
func (_ShortsTracker *ShortsTrackerCallerSession) MAXINT256() (*big.Int, error) {
	return _ShortsTracker.Contract.MAXINT256(&_ShortsTracker.CallOpts)
}

// GetNextGlobalAveragePrice is a free data retrieval call binding the contract method 0x5886b711.
//
// Solidity: function _getNextGlobalAveragePrice(uint256 _averagePrice, uint256 _nextPrice, uint256 _nextSize, uint256 _delta, int256 _realisedPnl) pure returns(uint256)
func (_ShortsTracker *ShortsTrackerCaller) GetNextGlobalAveragePrice(opts *bind.CallOpts, _averagePrice *big.Int, _nextPrice *big.Int, _nextSize *big.Int, _delta *big.Int, _realisedPnl *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "_getNextGlobalAveragePrice", _averagePrice, _nextPrice, _nextSize, _delta, _realisedPnl)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextGlobalAveragePrice is a free data retrieval call binding the contract method 0x5886b711.
//
// Solidity: function _getNextGlobalAveragePrice(uint256 _averagePrice, uint256 _nextPrice, uint256 _nextSize, uint256 _delta, int256 _realisedPnl) pure returns(uint256)
func (_ShortsTracker *ShortsTrackerSession) GetNextGlobalAveragePrice(_averagePrice *big.Int, _nextPrice *big.Int, _nextSize *big.Int, _delta *big.Int, _realisedPnl *big.Int) (*big.Int, error) {
	return _ShortsTracker.Contract.GetNextGlobalAveragePrice(&_ShortsTracker.CallOpts, _averagePrice, _nextPrice, _nextSize, _delta, _realisedPnl)
}

// GetNextGlobalAveragePrice is a free data retrieval call binding the contract method 0x5886b711.
//
// Solidity: function _getNextGlobalAveragePrice(uint256 _averagePrice, uint256 _nextPrice, uint256 _nextSize, uint256 _delta, int256 _realisedPnl) pure returns(uint256)
func (_ShortsTracker *ShortsTrackerCallerSession) GetNextGlobalAveragePrice(_averagePrice *big.Int, _nextPrice *big.Int, _nextSize *big.Int, _delta *big.Int, _realisedPnl *big.Int) (*big.Int, error) {
	return _ShortsTracker.Contract.GetNextGlobalAveragePrice(&_ShortsTracker.CallOpts, _averagePrice, _nextPrice, _nextSize, _delta, _realisedPnl)
}

// Data is a free data retrieval call binding the contract method 0x0147fb0c.
//
// Solidity: function data(bytes32 ) view returns(bytes32)
func (_ShortsTracker *ShortsTrackerCaller) Data(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "data", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x0147fb0c.
//
// Solidity: function data(bytes32 ) view returns(bytes32)
func (_ShortsTracker *ShortsTrackerSession) Data(arg0 [32]byte) ([32]byte, error) {
	return _ShortsTracker.Contract.Data(&_ShortsTracker.CallOpts, arg0)
}

// Data is a free data retrieval call binding the contract method 0x0147fb0c.
//
// Solidity: function data(bytes32 ) view returns(bytes32)
func (_ShortsTracker *ShortsTrackerCallerSession) Data(arg0 [32]byte) ([32]byte, error) {
	return _ShortsTracker.Contract.Data(&_ShortsTracker.CallOpts, arg0)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_ShortsTracker *ShortsTrackerCaller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "getGlobalShortDelta", _token)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_ShortsTracker *ShortsTrackerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _ShortsTracker.Contract.GetGlobalShortDelta(&_ShortsTracker.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_ShortsTracker *ShortsTrackerCallerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _ShortsTracker.Contract.GetGlobalShortDelta(&_ShortsTracker.CallOpts, _token)
}

// GetNextGlobalShortData is a free data retrieval call binding the contract method 0x9cdeb593.
//
// Solidity: function getNextGlobalShortData(address _account, address _collateralToken, address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTracker *ShortsTrackerCaller) GetNextGlobalShortData(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "getNextGlobalShortData", _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetNextGlobalShortData is a free data retrieval call binding the contract method 0x9cdeb593.
//
// Solidity: function getNextGlobalShortData(address _account, address _collateralToken, address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTracker *ShortsTrackerSession) GetNextGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _ShortsTracker.Contract.GetNextGlobalShortData(&_ShortsTracker.CallOpts, _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)
}

// GetNextGlobalShortData is a free data retrieval call binding the contract method 0x9cdeb593.
//
// Solidity: function getNextGlobalShortData(address _account, address _collateralToken, address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTracker *ShortsTrackerCallerSession) GetNextGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _ShortsTracker.Contract.GetNextGlobalShortData(&_ShortsTracker.CallOpts, _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)
}

// GetRealisedPnl is a free data retrieval call binding the contract method 0xa83b75fd.
//
// Solidity: function getRealisedPnl(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isIncrease) view returns(int256)
func (_ShortsTracker *ShortsTrackerCaller) GetRealisedPnl(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "getRealisedPnl", _account, _collateralToken, _indexToken, _sizeDelta, _isIncrease)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRealisedPnl is a free data retrieval call binding the contract method 0xa83b75fd.
//
// Solidity: function getRealisedPnl(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isIncrease) view returns(int256)
func (_ShortsTracker *ShortsTrackerSession) GetRealisedPnl(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, error) {
	return _ShortsTracker.Contract.GetRealisedPnl(&_ShortsTracker.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isIncrease)
}

// GetRealisedPnl is a free data retrieval call binding the contract method 0xa83b75fd.
//
// Solidity: function getRealisedPnl(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isIncrease) view returns(int256)
func (_ShortsTracker *ShortsTrackerCallerSession) GetRealisedPnl(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, error) {
	return _ShortsTracker.Contract.GetRealisedPnl(&_ShortsTracker.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isIncrease)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_ShortsTracker *ShortsTrackerCaller) GlobalShortAveragePrices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "globalShortAveragePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_ShortsTracker *ShortsTrackerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _ShortsTracker.Contract.GlobalShortAveragePrices(&_ShortsTracker.CallOpts, arg0)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_ShortsTracker *ShortsTrackerCallerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _ShortsTracker.Contract.GlobalShortAveragePrices(&_ShortsTracker.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ShortsTracker *ShortsTrackerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ShortsTracker *ShortsTrackerSession) Gov() (common.Address, error) {
	return _ShortsTracker.Contract.Gov(&_ShortsTracker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ShortsTracker *ShortsTrackerCallerSession) Gov() (common.Address, error) {
	return _ShortsTracker.Contract.Gov(&_ShortsTracker.CallOpts)
}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_ShortsTracker *ShortsTrackerCaller) IsGlobalShortDataReady(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "isGlobalShortDataReady")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_ShortsTracker *ShortsTrackerSession) IsGlobalShortDataReady() (bool, error) {
	return _ShortsTracker.Contract.IsGlobalShortDataReady(&_ShortsTracker.CallOpts)
}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_ShortsTracker *ShortsTrackerCallerSession) IsGlobalShortDataReady() (bool, error) {
	return _ShortsTracker.Contract.IsGlobalShortDataReady(&_ShortsTracker.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTracker *ShortsTrackerCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTracker *ShortsTrackerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ShortsTracker.Contract.IsHandler(&_ShortsTracker.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTracker *ShortsTrackerCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ShortsTracker.Contract.IsHandler(&_ShortsTracker.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ShortsTracker *ShortsTrackerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShortsTracker.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ShortsTracker *ShortsTrackerSession) Vault() (common.Address, error) {
	return _ShortsTracker.Contract.Vault(&_ShortsTracker.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ShortsTracker *ShortsTrackerCallerSession) Vault() (common.Address, error) {
	return _ShortsTracker.Contract.Vault(&_ShortsTracker.CallOpts)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ShortsTracker *ShortsTrackerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTracker.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ShortsTracker *ShortsTrackerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetGov(&_ShortsTracker.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ShortsTracker *ShortsTrackerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetGov(&_ShortsTracker.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ShortsTracker *ShortsTrackerTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTracker.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ShortsTracker *ShortsTrackerSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetHandler(&_ShortsTracker.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ShortsTracker *ShortsTrackerTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetHandler(&_ShortsTracker.TransactOpts, _handler, _isActive)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTracker *ShortsTrackerTransactor) SetInitData(opts *bind.TransactOpts, _tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTracker.contract.Transact(opts, "setInitData", _tokens, _averagePrices)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTracker *ShortsTrackerSession) SetInitData(_tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetInitData(&_ShortsTracker.TransactOpts, _tokens, _averagePrices)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTracker *ShortsTrackerTransactorSession) SetInitData(_tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetInitData(&_ShortsTracker.TransactOpts, _tokens, _averagePrices)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_ShortsTracker *ShortsTrackerTransactor) SetIsGlobalShortDataReady(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ShortsTracker.contract.Transact(opts, "setIsGlobalShortDataReady", value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_ShortsTracker *ShortsTrackerSession) SetIsGlobalShortDataReady(value bool) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetIsGlobalShortDataReady(&_ShortsTracker.TransactOpts, value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_ShortsTracker *ShortsTrackerTransactorSession) SetIsGlobalShortDataReady(value bool) (*types.Transaction, error) {
	return _ShortsTracker.Contract.SetIsGlobalShortDataReady(&_ShortsTracker.TransactOpts, value)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_ShortsTracker *ShortsTrackerTransactor) UpdateGlobalShortData(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _ShortsTracker.contract.Transact(opts, "updateGlobalShortData", _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_ShortsTracker *ShortsTrackerSession) UpdateGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _ShortsTracker.Contract.UpdateGlobalShortData(&_ShortsTracker.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_ShortsTracker *ShortsTrackerTransactorSession) UpdateGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _ShortsTracker.Contract.UpdateGlobalShortData(&_ShortsTracker.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// ShortsTrackerGlobalShortDataUpdatedIterator is returned from FilterGlobalShortDataUpdated and is used to iterate over the raw logs and unpacked data for GlobalShortDataUpdated events raised by the ShortsTracker contract.
type ShortsTrackerGlobalShortDataUpdatedIterator struct {
	Event *ShortsTrackerGlobalShortDataUpdated // Event containing the contract specifics and raw log

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
func (it *ShortsTrackerGlobalShortDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerGlobalShortDataUpdated)
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
		it.Event = new(ShortsTrackerGlobalShortDataUpdated)
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
func (it *ShortsTrackerGlobalShortDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerGlobalShortDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerGlobalShortDataUpdated represents a GlobalShortDataUpdated event raised by the ShortsTracker contract.
type ShortsTrackerGlobalShortDataUpdated struct {
	Token                   common.Address
	GlobalShortSize         *big.Int
	GlobalShortAveragePrice *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalShortDataUpdated is a free log retrieval operation binding the contract event 0xd6137be44db128ffcf1ea1821dbe8f889f67f949be7656c2d8acba2a4a891a02.
//
// Solidity: event GlobalShortDataUpdated(address indexed token, uint256 globalShortSize, uint256 globalShortAveragePrice)
func (_ShortsTracker *ShortsTrackerFilterer) FilterGlobalShortDataUpdated(opts *bind.FilterOpts, token []common.Address) (*ShortsTrackerGlobalShortDataUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ShortsTracker.contract.FilterLogs(opts, "GlobalShortDataUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerGlobalShortDataUpdatedIterator{contract: _ShortsTracker.contract, event: "GlobalShortDataUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalShortDataUpdated is a free log subscription operation binding the contract event 0xd6137be44db128ffcf1ea1821dbe8f889f67f949be7656c2d8acba2a4a891a02.
//
// Solidity: event GlobalShortDataUpdated(address indexed token, uint256 globalShortSize, uint256 globalShortAveragePrice)
func (_ShortsTracker *ShortsTrackerFilterer) WatchGlobalShortDataUpdated(opts *bind.WatchOpts, sink chan<- *ShortsTrackerGlobalShortDataUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ShortsTracker.contract.WatchLogs(opts, "GlobalShortDataUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerGlobalShortDataUpdated)
				if err := _ShortsTracker.contract.UnpackLog(event, "GlobalShortDataUpdated", log); err != nil {
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

// ParseGlobalShortDataUpdated is a log parse operation binding the contract event 0xd6137be44db128ffcf1ea1821dbe8f889f67f949be7656c2d8acba2a4a891a02.
//
// Solidity: event GlobalShortDataUpdated(address indexed token, uint256 globalShortSize, uint256 globalShortAveragePrice)
func (_ShortsTracker *ShortsTrackerFilterer) ParseGlobalShortDataUpdated(log types.Log) (*ShortsTrackerGlobalShortDataUpdated, error) {
	event := new(ShortsTrackerGlobalShortDataUpdated)
	if err := _ShortsTracker.contract.UnpackLog(event, "GlobalShortDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
