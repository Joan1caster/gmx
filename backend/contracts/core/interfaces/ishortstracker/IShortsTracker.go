// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ishortstracker

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

// IShortsTrackerMetaData contains all meta data concerning the IShortsTracker contract.
var IShortsTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"getNextGlobalShortData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"globalShortAveragePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isGlobalShortDataReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_averagePrices\",\"type\":\"uint256[]\"}],\"name\":\"setInitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setIsGlobalShortDataReady\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_markPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isIncrease\",\"type\":\"bool\"}],\"name\":\"updateGlobalShortData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IShortsTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use IShortsTrackerMetaData.ABI instead.
var IShortsTrackerABI = IShortsTrackerMetaData.ABI

// IShortsTracker is an auto generated Go binding around an Ethereum contract.
type IShortsTracker struct {
	IShortsTrackerCaller     // Read-only binding to the contract
	IShortsTrackerTransactor // Write-only binding to the contract
	IShortsTrackerFilterer   // Log filterer for contract events
}

// IShortsTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IShortsTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShortsTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IShortsTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShortsTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IShortsTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IShortsTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IShortsTrackerSession struct {
	Contract     *IShortsTracker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IShortsTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IShortsTrackerCallerSession struct {
	Contract *IShortsTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IShortsTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IShortsTrackerTransactorSession struct {
	Contract     *IShortsTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IShortsTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IShortsTrackerRaw struct {
	Contract *IShortsTracker // Generic contract binding to access the raw methods on
}

// IShortsTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IShortsTrackerCallerRaw struct {
	Contract *IShortsTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// IShortsTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IShortsTrackerTransactorRaw struct {
	Contract *IShortsTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIShortsTracker creates a new instance of IShortsTracker, bound to a specific deployed contract.
func NewIShortsTracker(address common.Address, backend bind.ContractBackend) (*IShortsTracker, error) {
	contract, err := bindIShortsTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IShortsTracker{IShortsTrackerCaller: IShortsTrackerCaller{contract: contract}, IShortsTrackerTransactor: IShortsTrackerTransactor{contract: contract}, IShortsTrackerFilterer: IShortsTrackerFilterer{contract: contract}}, nil
}

// NewIShortsTrackerCaller creates a new read-only instance of IShortsTracker, bound to a specific deployed contract.
func NewIShortsTrackerCaller(address common.Address, caller bind.ContractCaller) (*IShortsTrackerCaller, error) {
	contract, err := bindIShortsTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IShortsTrackerCaller{contract: contract}, nil
}

// NewIShortsTrackerTransactor creates a new write-only instance of IShortsTracker, bound to a specific deployed contract.
func NewIShortsTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*IShortsTrackerTransactor, error) {
	contract, err := bindIShortsTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IShortsTrackerTransactor{contract: contract}, nil
}

// NewIShortsTrackerFilterer creates a new log filterer instance of IShortsTracker, bound to a specific deployed contract.
func NewIShortsTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*IShortsTrackerFilterer, error) {
	contract, err := bindIShortsTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IShortsTrackerFilterer{contract: contract}, nil
}

// bindIShortsTracker binds a generic wrapper to an already deployed contract.
func bindIShortsTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IShortsTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShortsTracker *IShortsTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IShortsTracker.Contract.IShortsTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShortsTracker *IShortsTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShortsTracker.Contract.IShortsTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShortsTracker *IShortsTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShortsTracker.Contract.IShortsTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IShortsTracker *IShortsTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IShortsTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IShortsTracker *IShortsTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IShortsTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IShortsTracker *IShortsTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IShortsTracker.Contract.contract.Transact(opts, method, params...)
}

// GetNextGlobalShortData is a free data retrieval call binding the contract method 0x9cdeb593.
//
// Solidity: function getNextGlobalShortData(address _account, address _collateralToken, address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, bool _isIncrease) view returns(uint256, uint256)
func (_IShortsTracker *IShortsTrackerCaller) GetNextGlobalShortData(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IShortsTracker.contract.Call(opts, &out, "getNextGlobalShortData", _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)

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
func (_IShortsTracker *IShortsTrackerSession) GetNextGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _IShortsTracker.Contract.GetNextGlobalShortData(&_IShortsTracker.CallOpts, _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)
}

// GetNextGlobalShortData is a free data retrieval call binding the contract method 0x9cdeb593.
//
// Solidity: function getNextGlobalShortData(address _account, address _collateralToken, address _indexToken, uint256 _nextPrice, uint256 _sizeDelta, bool _isIncrease) view returns(uint256, uint256)
func (_IShortsTracker *IShortsTrackerCallerSession) GetNextGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int, _isIncrease bool) (*big.Int, *big.Int, error) {
	return _IShortsTracker.Contract.GetNextGlobalShortData(&_IShortsTracker.CallOpts, _account, _collateralToken, _indexToken, _nextPrice, _sizeDelta, _isIncrease)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IShortsTracker *IShortsTrackerCaller) GlobalShortAveragePrices(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IShortsTracker.contract.Call(opts, &out, "globalShortAveragePrices", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IShortsTracker *IShortsTrackerSession) GlobalShortAveragePrices(_token common.Address) (*big.Int, error) {
	return _IShortsTracker.Contract.GlobalShortAveragePrices(&_IShortsTracker.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address _token) view returns(uint256)
func (_IShortsTracker *IShortsTrackerCallerSession) GlobalShortAveragePrices(_token common.Address) (*big.Int, error) {
	return _IShortsTracker.Contract.GlobalShortAveragePrices(&_IShortsTracker.CallOpts, _token)
}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_IShortsTracker *IShortsTrackerCaller) IsGlobalShortDataReady(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IShortsTracker.contract.Call(opts, &out, "isGlobalShortDataReady")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_IShortsTracker *IShortsTrackerSession) IsGlobalShortDataReady() (bool, error) {
	return _IShortsTracker.Contract.IsGlobalShortDataReady(&_IShortsTracker.CallOpts)
}

// IsGlobalShortDataReady is a free data retrieval call binding the contract method 0x9a11178f.
//
// Solidity: function isGlobalShortDataReady() view returns(bool)
func (_IShortsTracker *IShortsTrackerCallerSession) IsGlobalShortDataReady() (bool, error) {
	return _IShortsTracker.Contract.IsGlobalShortDataReady(&_IShortsTracker.CallOpts)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_IShortsTracker *IShortsTrackerTransactor) SetInitData(opts *bind.TransactOpts, _tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _IShortsTracker.contract.Transact(opts, "setInitData", _tokens, _averagePrices)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_IShortsTracker *IShortsTrackerSession) SetInitData(_tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _IShortsTracker.Contract.SetInitData(&_IShortsTracker.TransactOpts, _tokens, _averagePrices)
}

// SetInitData is a paid mutator transaction binding the contract method 0xbbd97187.
//
// Solidity: function setInitData(address[] _tokens, uint256[] _averagePrices) returns()
func (_IShortsTracker *IShortsTrackerTransactorSession) SetInitData(_tokens []common.Address, _averagePrices []*big.Int) (*types.Transaction, error) {
	return _IShortsTracker.Contract.SetInitData(&_IShortsTracker.TransactOpts, _tokens, _averagePrices)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_IShortsTracker *IShortsTrackerTransactor) SetIsGlobalShortDataReady(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _IShortsTracker.contract.Transact(opts, "setIsGlobalShortDataReady", value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_IShortsTracker *IShortsTrackerSession) SetIsGlobalShortDataReady(value bool) (*types.Transaction, error) {
	return _IShortsTracker.Contract.SetIsGlobalShortDataReady(&_IShortsTracker.TransactOpts, value)
}

// SetIsGlobalShortDataReady is a paid mutator transaction binding the contract method 0x3d30cabf.
//
// Solidity: function setIsGlobalShortDataReady(bool value) returns()
func (_IShortsTracker *IShortsTrackerTransactorSession) SetIsGlobalShortDataReady(value bool) (*types.Transaction, error) {
	return _IShortsTracker.Contract.SetIsGlobalShortDataReady(&_IShortsTracker.TransactOpts, value)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_IShortsTracker *IShortsTrackerTransactor) UpdateGlobalShortData(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _IShortsTracker.contract.Transact(opts, "updateGlobalShortData", _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_IShortsTracker *IShortsTrackerSession) UpdateGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _IShortsTracker.Contract.UpdateGlobalShortData(&_IShortsTracker.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}

// UpdateGlobalShortData is a paid mutator transaction binding the contract method 0xf3238cec.
//
// Solidity: function updateGlobalShortData(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta, uint256 _markPrice, bool _isIncrease) returns()
func (_IShortsTracker *IShortsTrackerTransactorSession) UpdateGlobalShortData(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int, _markPrice *big.Int, _isIncrease bool) (*types.Transaction, error) {
	return _IShortsTracker.Contract.UpdateGlobalShortData(&_IShortsTracker.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta, _markPrice, _isIncrease)
}
