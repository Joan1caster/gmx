// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shortstrackertest

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

// ShortsTrackerTestMetaData contains all meta data concerning the ShortsTrackerTest contract.
var ShortsTrackerTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalShortSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalShortAveragePrice\",\"type\":\"uint256\"}],\"name\":\"GlobalShortDataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_INT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delta\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_realisedPnl\",\"type\":\"int256\"}],\"name\":\"_getNextGlobalAveragePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"data\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getGlobalShortDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"getNextGlobalShortData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_realisedPnl\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"getNextGlobalShortDataWithRealisedPnl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"getRealisedPnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"globalShortAveragePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isGlobalShortDataReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"}],\"name\":\"setGlobalShortAveragePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_averagePrices\",\"type\":\"uint256[]\"}],\"name\":\"setInitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setIsGlobalShortDataReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_markPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"updateGlobalShortData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ShortsTrackerTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ShortsTrackerTestMetaData.ABI instead.
var ShortsTrackerTestABI = ShortsTrackerTestMetaData.ABI

// ShortsTrackerTest is an auto generated Go binding around an Ethereum contract.
type ShortsTrackerTest struct {
	ShortsTrackerTestCaller     // Read-only binding to the contract
	ShortsTrackerTestTransactor // Write-only binding to the contract
	ShortsTrackerTestFilterer   // Log filterer for contract events
}

// ShortsTrackerTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShortsTrackerTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShortsTrackerTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShortsTrackerTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShortsTrackerTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShortsTrackerTestSession struct {
	Contract     *ShortsTrackerTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ShortsTrackerTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShortsTrackerTestCallerSession struct {
	Contract *ShortsTrackerTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ShortsTrackerTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShortsTrackerTestTransactorSession struct {
	Contract     *ShortsTrackerTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ShortsTrackerTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShortsTrackerTestRaw struct {
	Contract *ShortsTrackerTest // Generic contract binding to access the raw methods on
}

// ShortsTrackerTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShortsTrackerTestCallerRaw struct {
	Contract *ShortsTrackerTestCaller // Generic read-only contract binding to access the raw methods on
}

// ShortsTrackerTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShortsTrackerTestTransactorRaw struct {
	Contract *ShortsTrackerTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShortsTrackerTest creates a new instance of ShortsTrackerTest, bound to a specific deployed contract.
func NewShortsTrackerTest(address common.Address, backend bind.ContractBackend) (*ShortsTrackerTest, error) {
	contract, err := bindShortsTrackerTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTest{ShortsTrackerTestCaller: ShortsTrackerTestCaller{contract: contract}, ShortsTrackerTestTransactor: ShortsTrackerTestTransactor{contract: contract}, ShortsTrackerTestFilterer: ShortsTrackerTestFilterer{contract: contract}}, nil
}

// NewShortsTrackerTestCaller creates a new read-only instance of ShortsTrackerTest, bound to a specific deployed contract.
func NewShortsTrackerTestCaller(address common.Address, caller bind.ContractCaller) (*ShortsTrackerTestCaller, error) {
	contract, err := bindShortsTrackerTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTestCaller{contract: contract}, nil
}

// NewShortsTrackerTestTransactor creates a new write-only instance of ShortsTrackerTest, bound to a specific deployed contract.
func NewShortsTrackerTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ShortsTrackerTestTransactor, error) {
	contract, err := bindShortsTrackerTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTestTransactor{contract: contract}, nil
}

// NewShortsTrackerTestFilterer creates a new log filterer instance of ShortsTrackerTest, bound to a specific deployed contract.
func NewShortsTrackerTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ShortsTrackerTestFilterer, error) {
	contract, err := bindShortsTrackerTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTestFilterer{contract: contract}, nil
}

// bindShortsTrackerTest binds a generic wrapper to an already deployed contract.
func bindShortsTrackerTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShortsTrackerTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShortsTrackerTest *ShortsTrackerTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShortsTrackerTest.Contract.ShortsTrackerTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShortsTrackerTest *ShortsTrackerTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.ShortsTrackerTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShortsTrackerTest *ShortsTrackerTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.ShortsTrackerTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShortsTrackerTest *ShortsTrackerTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShortsTrackerTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShortsTrackerTest *ShortsTrackerTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShortsTrackerTest *ShortsTrackerTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.contract.Transact(opts, method, params...)
}

// MAXINT256 is a free data retrieval call binding the contract method 0x122e7b07.
//
// Solidity: function MAX_INT256() view returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) MAXINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "MAX_INT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINT256 is a free data retrieval call binding the contract method 0x122e7b07.
//
// Solidity: function MAX_INT256() view returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestSession) MAXINT256() (*big.Int, error) {
	return _ShortsTrackerTest.Contract.MAXINT256(&_ShortsTrackerTest.CallOpts)
}

// MAXINT256 is a free data retrieval call binding the contract method 0x122e7b07.
//
// Solidity: function MAX_INT256() view returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) MAXINT256() (*big.Int, error) {
	return _ShortsTrackerTest.Contract.MAXINT256(&_ShortsTrackerTest.CallOpts)
}

// GetNextGlobalAveragePrice is a free data retrieval call binding the contract method 0x5886b711.
//
// Solidity: function _getNextGlobalAveragePrice(uint256 _averagePrice, uint256 _nextPrice, uint256 _nextSize, uint256 _delta, int256 _realisedPnl) pure returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) GetNextGlobalAveragePrice(opts *bind.CallOpts, _averagePrice *big.Int, _nextPrice *big.Int, _nextSize *big.Int, _delta *big.Int, _realisedPnl *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "_getNextGlobalAveragePrice", _averagePrice, _nextPrice, _nextSize, _delta, _realisedPnl)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextGlobalAveragePrice is a free data retrieval call binding the contract method 0x5886b711.
//
// Solidity: function _getNextGlobalAveragePrice(uint256 _averagePrice, uint256 _nextPrice, uint256 _nextSize, uint256 _delta, int256 _realisedPnl) pure returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestSession) GetNextGlobalAveragePrice(_averagePrice *big.Int, _nextPrice *big.Int, _nextSize *big.Int, _delta *big.Int, _realisedPnl *big.Int) (*big.Int, error) {
	return _ShortsTrackerTest.Contract.GetNextGlobalAveragePrice(&_ShortsTrackerTest.CallOpts, _averagePrice, _nextPrice, _nextSize, _delta, _realisedPnl)
}

// GetNextGlobalAveragePrice is a free data retrieval call binding the contract method 0x5886b711.
//
// Solidity: function _getNextGlobalAveragePrice(uint256 _averagePrice, uint256 _nextPrice, uint256 _nextSize, uint256 _delta, int256 _realisedPnl) pure returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) GetNextGlobalAveragePrice(_averagePrice *big.Int, _nextPrice *big.Int, _nextSize *big.Int, _delta *big.Int, _realisedPnl *big.Int) (*big.Int, error) {
	return _ShortsTrackerTest.Contract.GetNextGlobalAveragePrice(&_ShortsTrackerTest.CallOpts, _averagePrice, _nextPrice, _nextSize, _delta, _realisedPnl)
}

// Data is a free data retrieval call binding the contract method 0x0147fb0c.
//
// Solidity: function data(bytes32 ) view returns(bytes32)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) Data(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "data", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x0147fb0c.
//
// Solidity: function data(bytes32 ) view returns(bytes32)
func (_ShortsTrackerTest *ShortsTrackerTestSession) Data(arg0 [32]byte) ([32]byte, error) {
	return _ShortsTrackerTest.Contract.Data(&_ShortsTrackerTest.CallOpts, arg0)
}

// Data is a free data retrieval call binding the contract method 0x0147fb0c.
//
// Solidity: function data(bytes32 ) view returns(bytes32)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) Data(arg0 [32]byte) ([32]byte, error) {
	return _ShortsTrackerTest.Contract.Data(&_ShortsTrackerTest.CallOpts, arg0)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "getGlobalShortDelta", _token)

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
func (_ShortsTrackerTest *ShortsTrackerTestSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _ShortsTrackerTest.Contract.GetGlobalShortDelta(&_ShortsTrackerTest.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _ShortsTrackerTest.Contract.GetGlobalShortDelta(&_ShortsTrackerTest.CallOpts, _token)
}

// GetNextGlobalShortData is a free data retrieval call binding the contract method 0x9cdeb593.
//
// Solidity: function getNextGlobalShortData(address _account, address _collateralToken, address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) GetNextGlobalShortData(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "getNextGlobalShortData", _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)

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
func (_ShortsTrackerTest *ShortsTrackerTestSession) GetNextGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _ShortsTrackerTest.Contract.GetNextGlobalShortData(&_ShortsTrackerTest.CallOpts, _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)
}

// GetNextGlobalShortData is a free data retrieval call binding the contract method 0x9cdeb593.
//
// Solidity: function getNextGlobalShortData(address _account, address _collateralToken, address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) GetNextGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _ShortsTrackerTest.Contract.GetNextGlobalShortData(&_ShortsTrackerTest.CallOpts, _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)
}

// GetNextGlobalShortDataWithRealisedPnl is a free data retrieval call binding the contract method 0x83d7048d.
//
// Solidity: function getNextGlobalShortDataWithRealisedPnl(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, int256 _realisedPnl, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) GetNextGlobalShortDataWithRealisedPnl(opts *bind.CallOpts, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _realisedPnl *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "getNextGlobalShortDataWithRealisedPnl", _indexToken, _nextPrice, _sizeDelta, _realisedPnl, _isIncrease)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetNextGlobalShortDataWithRealisedPnl is a free data retrieval call binding the contract method 0x83d7048d.
//
// Solidity: function getNextGlobalShortDataWithRealisedPnl(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, int256 _realisedPnl, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTrackerTest *ShortsTrackerTestSession) GetNextGlobalShortDataWithRealisedPnl(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _realisedPnl *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _ShortsTrackerTest.Contract.GetNextGlobalShortDataWithRealisedPnl(&_ShortsTrackerTest.CallOpts, _indexToken, _nextPrice, _sizeDelta, _realisedPnl, _isIncrease)
}

// GetNextGlobalShortDataWithRealisedPnl is a free data retrieval call binding the contract method 0x83d7048d.
//
// Solidity: function getNextGlobalShortDataWithRealisedPnl(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, int256 _realisedPnl, bool _isIncrease) view returns(uint256, uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) GetNextGlobalShortDataWithRealisedPnl(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _realisedPnl *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _ShortsTrackerTest.Contract.GetNextGlobalShortDataWithRealisedPnl(&_ShortsTrackerTest.CallOpts, _indexToken, _nextPrice, _sizeDelta, _realisedPnl, _isIncrease)
}

// GetRealisedPnl is a free data retrieval call binding the contract method 0xa83b75fd.
//
// Solidity: function getRealisedPnl(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isIncrease) view returns(int256)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) GetRealisedPnl(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "getRealisedPnl", _account, _collateralToken, _indexToken, _sizeDelta, _isIncrease)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRealisedPnl is a free data retrieval call binding the contract method 0xa83b75fd.
//
// Solidity: function getRealisedPnl(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isIncrease) view returns(int256)
func (_ShortsTrackerTest *ShortsTrackerTestSession) GetRealisedPnl(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, error) {
	return _ShortsTrackerTest.Contract.GetRealisedPnl(&_ShortsTrackerTest.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isIncrease)
}

// GetRealisedPnl is a free data retrieval call binding the contract method 0xa83b75fd.
//
// Solidity: function getRealisedPnl(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isIncrease) view returns(int256)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) GetRealisedPnl(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, error) {
	return _ShortsTrackerTest.Contract.GetRealisedPnl(&_ShortsTrackerTest.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isIncrease)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) GlobalShortAveragePrices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "globalShortAveragePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _ShortsTrackerTest.Contract.GlobalShortAveragePrices(&_ShortsTrackerTest.CallOpts, arg0)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _ShortsTrackerTest.Contract.GlobalShortAveragePrices(&_ShortsTrackerTest.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ShortsTrackerTest *ShortsTrackerTestSession) Gov() (common.Address, error) {
	return _ShortsTrackerTest.Contract.Gov(&_ShortsTrackerTest.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) Gov() (common.Address, error) {
	return _ShortsTrackerTest.Contract.Gov(&_ShortsTrackerTest.CallOpts)
}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) IsGlobalShortDataReady(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "isGlobalShortDataReady")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_ShortsTrackerTest *ShortsTrackerTestSession) IsGlobalShortDataReady() (bool, error) {
	return _ShortsTrackerTest.Contract.IsGlobalShortDataReady(&_ShortsTrackerTest.CallOpts)
}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) IsGlobalShortDataReady() (bool, error) {
	return _ShortsTrackerTest.Contract.IsGlobalShortDataReady(&_ShortsTrackerTest.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTrackerTest *ShortsTrackerTestSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ShortsTrackerTest.Contract.IsHandler(&_ShortsTrackerTest.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ShortsTrackerTest.Contract.IsHandler(&_ShortsTrackerTest.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ShortsTrackerTest *ShortsTrackerTestCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShortsTrackerTest.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ShortsTrackerTest *ShortsTrackerTestSession) Vault() (common.Address, error) {
	return _ShortsTrackerTest.Contract.Vault(&_ShortsTrackerTest.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ShortsTrackerTest *ShortsTrackerTestCallerSession) Vault() (common.Address, error) {
	return _ShortsTrackerTest.Contract.Vault(&_ShortsTrackerTest.CallOpts)
}

// SetGlobalShortAveragePrice is a paid mutator transaction binding the contract method 0x776edefc.
//
// Solidity: function setGlobalShortAveragePrice(address _token, uint256 _averagePrice) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactor) SetGlobalShortAveragePrice(opts *bind.TransactOpts, _token common.Address, _averagePrice *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTest.contract.Transact(opts, "setGlobalShortAveragePrice", _token, _averagePrice)
}

// SetGlobalShortAveragePrice is a paid mutator transaction binding the contract method 0x776edefc.
//
// Solidity: function setGlobalShortAveragePrice(address _token, uint256 _averagePrice) returns()
func (_ShortsTrackerTest *ShortsTrackerTestSession) SetGlobalShortAveragePrice(_token common.Address, _averagePrice *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetGlobalShortAveragePrice(&_ShortsTrackerTest.TransactOpts, _token, _averagePrice)
}

// SetGlobalShortAveragePrice is a paid mutator transaction binding the contract method 0x776edefc.
//
// Solidity: function setGlobalShortAveragePrice(address _token, uint256 _averagePrice) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactorSession) SetGlobalShortAveragePrice(_token common.Address, _averagePrice *big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetGlobalShortAveragePrice(&_ShortsTrackerTest.TransactOpts, _token, _averagePrice)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTest.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ShortsTrackerTest *ShortsTrackerTestSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetGov(&_ShortsTrackerTest.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetGov(&_ShortsTrackerTest.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ShortsTrackerTest *ShortsTrackerTestSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetHandler(&_ShortsTrackerTest.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetHandler(&_ShortsTrackerTest.TransactOpts, _handler, _isActive)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactor) SetInitData(opts *bind.TransactOpts, _tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTest.contract.Transact(opts, "setInitData", _tokens, _averagePrices)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTrackerTest *ShortsTrackerTestSession) SetInitData(_tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetInitData(&_ShortsTrackerTest.TransactOpts, _tokens, _averagePrices)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactorSession) SetInitData(_tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetInitData(&_ShortsTrackerTest.TransactOpts, _tokens, _averagePrices)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactor) SetIsGlobalShortDataReady(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.contract.Transact(opts, "setIsGlobalShortDataReady", value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_ShortsTrackerTest *ShortsTrackerTestSession) SetIsGlobalShortDataReady(value bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetIsGlobalShortDataReady(&_ShortsTrackerTest.TransactOpts, value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactorSession) SetIsGlobalShortDataReady(value bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.SetIsGlobalShortDataReady(&_ShortsTrackerTest.TransactOpts, value)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactor) UpdateGlobalShortData(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.contract.Transact(opts, "updateGlobalShortData", _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_ShortsTrackerTest *ShortsTrackerTestSession) UpdateGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.UpdateGlobalShortData(&_ShortsTrackerTest.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_ShortsTrackerTest *ShortsTrackerTestTransactorSession) UpdateGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _ShortsTrackerTest.Contract.UpdateGlobalShortData(&_ShortsTrackerTest.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// ShortsTrackerTestGlobalShortDataUpdatedIterator is returned from FilterGlobalShortDataUpdated and is used to iterate over the raw logs and unpacked data for GlobalShortDataUpdated events raised by the ShortsTrackerTest contract.
type ShortsTrackerTestGlobalShortDataUpdatedIterator struct {
	Event *ShortsTrackerTestGlobalShortDataUpdated // Event containing the contract specifics and raw log

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
func (it *ShortsTrackerTestGlobalShortDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShortsTrackerTestGlobalShortDataUpdated)
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
		it.Event = new(ShortsTrackerTestGlobalShortDataUpdated)
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
func (it *ShortsTrackerTestGlobalShortDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShortsTrackerTestGlobalShortDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShortsTrackerTestGlobalShortDataUpdated represents a GlobalShortDataUpdated event raised by the ShortsTrackerTest contract.
type ShortsTrackerTestGlobalShortDataUpdated struct {
	Token                   common.Address
	GlobalShortSize         *big.Int
	GlobalShortAveragePrice *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalShortDataUpdated is a free log retrieval operation binding the contract event 0xd6137be44db128ffcf1ea1821dbe8f889f67f949be7656c2d8acba2a4a891a02.
//
// Solidity: event GlobalShortDataUpdated(address indexed token, uint256 globalShortSize, uint256 globalShortAveragePrice)
func (_ShortsTrackerTest *ShortsTrackerTestFilterer) FilterGlobalShortDataUpdated(opts *bind.FilterOpts, token []common.Address) (*ShortsTrackerTestGlobalShortDataUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ShortsTrackerTest.contract.FilterLogs(opts, "GlobalShortDataUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ShortsTrackerTestGlobalShortDataUpdatedIterator{contract: _ShortsTrackerTest.contract, event: "GlobalShortDataUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalShortDataUpdated is a free log subscription operation binding the contract event 0xd6137be44db128ffcf1ea1821dbe8f889f67f949be7656c2d8acba2a4a891a02.
//
// Solidity: event GlobalShortDataUpdated(address indexed token, uint256 globalShortSize, uint256 globalShortAveragePrice)
func (_ShortsTrackerTest *ShortsTrackerTestFilterer) WatchGlobalShortDataUpdated(opts *bind.WatchOpts, sink chan<- *ShortsTrackerTestGlobalShortDataUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ShortsTrackerTest.contract.WatchLogs(opts, "GlobalShortDataUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShortsTrackerTestGlobalShortDataUpdated)
				if err := _ShortsTrackerTest.contract.UnpackLog(event, "GlobalShortDataUpdated", log); err != nil {
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
func (_ShortsTrackerTest *ShortsTrackerTestFilterer) ParseGlobalShortDataUpdated(log types.Log) (*ShortsTrackerTestGlobalShortDataUpdated, error) {
	event := new(ShortsTrackerTestGlobalShortDataUpdated)
	if err := _ShortsTrackerTest.contract.UnpackLog(event, "GlobalShortDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
