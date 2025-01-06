// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivaultpricefeed

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

// IVaultPriceFeedMetaData contains all meta data concerning the IVaultPriceFeed contract.
var IVaultPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"adjustmentBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getAmmPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getLatestPrimaryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_includeAmmPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_useSwapPricing\",\"type\":\"bool\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getPrimaryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isAdjustmentAdditive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdditive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_adjustmentBps\",\"type\":\"uint256\"}],\"name\":\"setAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_favorPrimaryPrice\",\"type\":\"bool\"}],\"name\":\"setFavorPrimaryPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsAmmEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSecondaryPriceEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxStrictPriceDeviation\",\"type\":\"uint256\"}],\"name\":\"setMaxStrictPriceDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceSampleSpace\",\"type\":\"uint256\"}],\"name\":\"setPriceSampleSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadThresholdBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadThresholdBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrictStable\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_useV2Pricing\",\"type\":\"bool\"}],\"name\":\"setUseV2Pricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IVaultPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use IVaultPriceFeedMetaData.ABI instead.
var IVaultPriceFeedABI = IVaultPriceFeedMetaData.ABI

// IVaultPriceFeed is an auto generated Go binding around an Ethereum contract.
type IVaultPriceFeed struct {
	IVaultPriceFeedCaller     // Read-only binding to the contract
	IVaultPriceFeedTransactor // Write-only binding to the contract
	IVaultPriceFeedFilterer   // Log filterer for contract events
}

// IVaultPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVaultPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVaultPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVaultPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVaultPriceFeedSession struct {
	Contract     *IVaultPriceFeed  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVaultPriceFeedCallerSession struct {
	Contract *IVaultPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IVaultPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVaultPriceFeedTransactorSession struct {
	Contract     *IVaultPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IVaultPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVaultPriceFeedRaw struct {
	Contract *IVaultPriceFeed // Generic contract binding to access the raw methods on
}

// IVaultPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVaultPriceFeedCallerRaw struct {
	Contract *IVaultPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// IVaultPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVaultPriceFeedTransactorRaw struct {
	Contract *IVaultPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVaultPriceFeed creates a new instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeed(address common.Address, backend bind.ContractBackend) (*IVaultPriceFeed, error) {
	contract, err := bindIVaultPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeed{IVaultPriceFeedCaller: IVaultPriceFeedCaller{contract: contract}, IVaultPriceFeedTransactor: IVaultPriceFeedTransactor{contract: contract}, IVaultPriceFeedFilterer: IVaultPriceFeedFilterer{contract: contract}}, nil
}

// NewIVaultPriceFeedCaller creates a new read-only instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*IVaultPriceFeedCaller, error) {
	contract, err := bindIVaultPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeedCaller{contract: contract}, nil
}

// NewIVaultPriceFeedTransactor creates a new write-only instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*IVaultPriceFeedTransactor, error) {
	contract, err := bindIVaultPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeedTransactor{contract: contract}, nil
}

// NewIVaultPriceFeedFilterer creates a new log filterer instance of IVaultPriceFeed, bound to a specific deployed contract.
func NewIVaultPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*IVaultPriceFeedFilterer, error) {
	contract, err := bindIVaultPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVaultPriceFeedFilterer{contract: contract}, nil
}

// bindIVaultPriceFeed binds a generic wrapper to an already deployed contract.
func bindIVaultPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVaultPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultPriceFeed *IVaultPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultPriceFeed.Contract.IVaultPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultPriceFeed *IVaultPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.IVaultPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultPriceFeed *IVaultPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.IVaultPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultPriceFeed *IVaultPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultPriceFeed *IVaultPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultPriceFeed *IVaultPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) AdjustmentBasisPoints(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "adjustmentBasisPoints", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) AdjustmentBasisPoints(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.AdjustmentBasisPoints(&_IVaultPriceFeed.CallOpts, _token)
}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) AdjustmentBasisPoints(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.AdjustmentBasisPoints(&_IVaultPriceFeed.CallOpts, _token)
}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetAmmPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getAmmPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetAmmPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetAmmPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetAmmPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetAmmPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetLatestPrimaryPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getLatestPrimaryPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetLatestPrimaryPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetLatestPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetLatestPrimaryPrice(_token common.Address) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetLatestPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool _useSwapPricing) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetPrice(opts *bind.CallOpts, _token common.Address, _maximise bool, _includeAmmPrice bool, _useSwapPricing bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getPrice", _token, _maximise, _includeAmmPrice, _useSwapPricing)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool _useSwapPricing) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetPrice(_token common.Address, _maximise bool, _includeAmmPrice bool, _useSwapPricing bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice, _useSwapPricing)
}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool _useSwapPricing) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetPrice(_token common.Address, _maximise bool, _includeAmmPrice bool, _useSwapPricing bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice, _useSwapPricing)
}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) GetPrimaryPrice(opts *bind.CallOpts, _token common.Address, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "getPrimaryPrice", _token, _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedSession) GetPrimaryPrice(_token common.Address, _maximise bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise)
}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) GetPrimaryPrice(_token common.Address, _maximise bool) (*big.Int, error) {
	return _IVaultPriceFeed.Contract.GetPrimaryPrice(&_IVaultPriceFeed.CallOpts, _token, _maximise)
}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address _token) view returns(bool)
func (_IVaultPriceFeed *IVaultPriceFeedCaller) IsAdjustmentAdditive(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _IVaultPriceFeed.contract.Call(opts, &out, "isAdjustmentAdditive", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address _token) view returns(bool)
func (_IVaultPriceFeed *IVaultPriceFeedSession) IsAdjustmentAdditive(_token common.Address) (bool, error) {
	return _IVaultPriceFeed.Contract.IsAdjustmentAdditive(&_IVaultPriceFeed.CallOpts, _token)
}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address _token) view returns(bool)
func (_IVaultPriceFeed *IVaultPriceFeedCallerSession) IsAdjustmentAdditive(_token common.Address) (bool, error) {
	return _IVaultPriceFeed.Contract.IsAdjustmentAdditive(&_IVaultPriceFeed.CallOpts, _token)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetAdjustment(opts *bind.TransactOpts, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setAdjustment", _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetAdjustment(_token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetAdjustment(&_IVaultPriceFeed.TransactOpts, _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetAdjustment(_token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetAdjustment(&_IVaultPriceFeed.TransactOpts, _token, _isAdditive, _adjustmentBps)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetFavorPrimaryPrice(opts *bind.TransactOpts, _favorPrimaryPrice bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setFavorPrimaryPrice", _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetFavorPrimaryPrice(_favorPrimaryPrice bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetFavorPrimaryPrice(&_IVaultPriceFeed.TransactOpts, _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetFavorPrimaryPrice(_favorPrimaryPrice bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetFavorPrimaryPrice(&_IVaultPriceFeed.TransactOpts, _favorPrimaryPrice)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetIsAmmEnabled(opts *bind.TransactOpts, _isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setIsAmmEnabled", _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetIsAmmEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsAmmEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetIsAmmEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsAmmEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetIsSecondaryPriceEnabled(opts *bind.TransactOpts, _isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setIsSecondaryPriceEnabled", _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetIsSecondaryPriceEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsSecondaryPriceEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetIsSecondaryPriceEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetIsSecondaryPriceEnabled(&_IVaultPriceFeed.TransactOpts, _isEnabled)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetMaxStrictPriceDeviation(opts *bind.TransactOpts, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setMaxStrictPriceDeviation", _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetMaxStrictPriceDeviation(_maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetMaxStrictPriceDeviation(&_IVaultPriceFeed.TransactOpts, _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetMaxStrictPriceDeviation(_maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetMaxStrictPriceDeviation(&_IVaultPriceFeed.TransactOpts, _maxStrictPriceDeviation)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetPriceSampleSpace(opts *bind.TransactOpts, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setPriceSampleSpace", _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetPriceSampleSpace(_priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetPriceSampleSpace(&_IVaultPriceFeed.TransactOpts, _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetPriceSampleSpace(_priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetPriceSampleSpace(&_IVaultPriceFeed.TransactOpts, _priceSampleSpace)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetSpreadBasisPoints(opts *bind.TransactOpts, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setSpreadBasisPoints", _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetSpreadBasisPoints(_token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadBasisPoints(&_IVaultPriceFeed.TransactOpts, _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetSpreadBasisPoints(_token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadBasisPoints(&_IVaultPriceFeed.TransactOpts, _token, _spreadBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetSpreadThresholdBasisPoints(opts *bind.TransactOpts, _spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setSpreadThresholdBasisPoints", _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetSpreadThresholdBasisPoints(_spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadThresholdBasisPoints(&_IVaultPriceFeed.TransactOpts, _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetSpreadThresholdBasisPoints(_spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetSpreadThresholdBasisPoints(&_IVaultPriceFeed.TransactOpts, _spreadThresholdBasisPoints)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setTokenConfig", _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetTokenConfig(_token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetTokenConfig(&_IVaultPriceFeed.TransactOpts, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetTokenConfig(_token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetTokenConfig(&_IVaultPriceFeed.TransactOpts, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactor) SetUseV2Pricing(opts *bind.TransactOpts, _useV2Pricing bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.contract.Transact(opts, "setUseV2Pricing", _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_IVaultPriceFeed *IVaultPriceFeedSession) SetUseV2Pricing(_useV2Pricing bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetUseV2Pricing(&_IVaultPriceFeed.TransactOpts, _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_IVaultPriceFeed *IVaultPriceFeedTransactorSession) SetUseV2Pricing(_useV2Pricing bool) (*types.Transaction, error) {
	return _IVaultPriceFeed.Contract.SetUseV2Pricing(&_IVaultPriceFeed.TransactOpts, _useV2Pricing)
}
