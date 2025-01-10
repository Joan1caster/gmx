// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ifastpricefeed

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

// IFastPriceFeedMetaData contains all meta data concerning the IFastPriceFeed contract.
var IFastPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"lastUpdatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSpreadEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSpreadEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_maxCumulativeDeltaDiffs\",\"type\":\"uint256[]\"}],\"name\":\"setMaxCumulativeDeltaDiffs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDeviationBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setMaxDeviationBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"setMaxPriceUpdateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBlockInterval\",\"type\":\"uint256\"}],\"name\":\"setMinBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceDataInterval\",\"type\":\"uint256\"}],\"name\":\"setPriceDataInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceDuration\",\"type\":\"uint256\"}],\"name\":\"setPriceDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfChainError\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfChainError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfInactive\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfInactive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IFastPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastPriceFeedMetaData.ABI instead.
var IFastPriceFeedABI = IFastPriceFeedMetaData.ABI

// IFastPriceFeed is an auto generated Go binding around an Ethereum contract.
type IFastPriceFeed struct {
	IFastPriceFeedCaller     // Read-only binding to the contract
	IFastPriceFeedTransactor // Write-only binding to the contract
	IFastPriceFeedFilterer   // Log filterer for contract events
}

// IFastPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFastPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastPriceFeedSession struct {
	Contract     *IFastPriceFeed   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFastPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastPriceFeedCallerSession struct {
	Contract *IFastPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFastPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastPriceFeedTransactorSession struct {
	Contract     *IFastPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFastPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFastPriceFeedRaw struct {
	Contract *IFastPriceFeed // Generic contract binding to access the raw methods on
}

// IFastPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastPriceFeedCallerRaw struct {
	Contract *IFastPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// IFastPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastPriceFeedTransactorRaw struct {
	Contract *IFastPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastPriceFeed creates a new instance of IFastPriceFeed, bound to a specific deployed contract.
func NewIFastPriceFeed(address common.Address, backend bind.ContractBackend) (*IFastPriceFeed, error) {
	contract, err := bindIFastPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastPriceFeed{IFastPriceFeedCaller: IFastPriceFeedCaller{contract: contract}, IFastPriceFeedTransactor: IFastPriceFeedTransactor{contract: contract}, IFastPriceFeedFilterer: IFastPriceFeedFilterer{contract: contract}}, nil
}

// NewIFastPriceFeedCaller creates a new read-only instance of IFastPriceFeed, bound to a specific deployed contract.
func NewIFastPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*IFastPriceFeedCaller, error) {
	contract, err := bindIFastPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastPriceFeedCaller{contract: contract}, nil
}

// NewIFastPriceFeedTransactor creates a new write-only instance of IFastPriceFeed, bound to a specific deployed contract.
func NewIFastPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*IFastPriceFeedTransactor, error) {
	contract, err := bindIFastPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastPriceFeedTransactor{contract: contract}, nil
}

// NewIFastPriceFeedFilterer creates a new log filterer instance of IFastPriceFeed, bound to a specific deployed contract.
func NewIFastPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*IFastPriceFeedFilterer, error) {
	contract, err := bindIFastPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastPriceFeedFilterer{contract: contract}, nil
}

// bindIFastPriceFeed binds a generic wrapper to an already deployed contract.
func bindIFastPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFastPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastPriceFeed *IFastPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastPriceFeed.Contract.IFastPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastPriceFeed *IFastPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.IFastPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastPriceFeed *IFastPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.IFastPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastPriceFeed *IFastPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastPriceFeed *IFastPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastPriceFeed *IFastPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_IFastPriceFeed *IFastPriceFeedCaller) LastUpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFastPriceFeed.contract.Call(opts, &out, "lastUpdatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_IFastPriceFeed *IFastPriceFeedSession) LastUpdatedAt() (*big.Int, error) {
	return _IFastPriceFeed.Contract.LastUpdatedAt(&_IFastPriceFeed.CallOpts)
}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_IFastPriceFeed *IFastPriceFeedCallerSession) LastUpdatedAt() (*big.Int, error) {
	return _IFastPriceFeed.Contract.LastUpdatedAt(&_IFastPriceFeed.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_IFastPriceFeed *IFastPriceFeedCaller) LastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IFastPriceFeed.contract.Call(opts, &out, "lastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_IFastPriceFeed *IFastPriceFeedSession) LastUpdatedBlock() (*big.Int, error) {
	return _IFastPriceFeed.Contract.LastUpdatedBlock(&_IFastPriceFeed.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_IFastPriceFeed *IFastPriceFeedCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _IFastPriceFeed.Contract.LastUpdatedBlock(&_IFastPriceFeed.CallOpts)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetIsSpreadEnabled(opts *bind.TransactOpts, _isSpreadEnabled bool) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setIsSpreadEnabled", _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetIsSpreadEnabled(_isSpreadEnabled bool) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetIsSpreadEnabled(&_IFastPriceFeed.TransactOpts, _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetIsSpreadEnabled(_isSpreadEnabled bool) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetIsSpreadEnabled(&_IFastPriceFeed.TransactOpts, _isSpreadEnabled)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetMaxCumulativeDeltaDiffs(opts *bind.TransactOpts, _tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setMaxCumulativeDeltaDiffs", _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetMaxCumulativeDeltaDiffs(_tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMaxCumulativeDeltaDiffs(&_IFastPriceFeed.TransactOpts, _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetMaxCumulativeDeltaDiffs(_tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMaxCumulativeDeltaDiffs(&_IFastPriceFeed.TransactOpts, _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetMaxDeviationBasisPoints(opts *bind.TransactOpts, _maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setMaxDeviationBasisPoints", _maxDeviationBasisPoints)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetMaxDeviationBasisPoints(_maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMaxDeviationBasisPoints(&_IFastPriceFeed.TransactOpts, _maxDeviationBasisPoints)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetMaxDeviationBasisPoints(_maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMaxDeviationBasisPoints(&_IFastPriceFeed.TransactOpts, _maxDeviationBasisPoints)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetMaxPriceUpdateDelay(opts *bind.TransactOpts, _maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setMaxPriceUpdateDelay", _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetMaxPriceUpdateDelay(_maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMaxPriceUpdateDelay(&_IFastPriceFeed.TransactOpts, _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetMaxPriceUpdateDelay(_maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMaxPriceUpdateDelay(&_IFastPriceFeed.TransactOpts, _maxPriceUpdateDelay)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetMinBlockInterval(opts *bind.TransactOpts, _minBlockInterval *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setMinBlockInterval", _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetMinBlockInterval(_minBlockInterval *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMinBlockInterval(&_IFastPriceFeed.TransactOpts, _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetMinBlockInterval(_minBlockInterval *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetMinBlockInterval(&_IFastPriceFeed.TransactOpts, _minBlockInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetPriceDataInterval(opts *bind.TransactOpts, _priceDataInterval *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setPriceDataInterval", _priceDataInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetPriceDataInterval(_priceDataInterval *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetPriceDataInterval(&_IFastPriceFeed.TransactOpts, _priceDataInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetPriceDataInterval(_priceDataInterval *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetPriceDataInterval(&_IFastPriceFeed.TransactOpts, _priceDataInterval)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetPriceDuration(opts *bind.TransactOpts, _priceDuration *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setPriceDuration", _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetPriceDuration(_priceDuration *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetPriceDuration(&_IFastPriceFeed.TransactOpts, _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetPriceDuration(_priceDuration *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetPriceDuration(&_IFastPriceFeed.TransactOpts, _priceDuration)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetSigner(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setSigner", _account, _isActive)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetSigner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetSigner(&_IFastPriceFeed.TransactOpts, _account, _isActive)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetSigner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetSigner(&_IFastPriceFeed.TransactOpts, _account, _isActive)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetSpreadBasisPointsIfChainError(opts *bind.TransactOpts, _spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setSpreadBasisPointsIfChainError", _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetSpreadBasisPointsIfChainError(_spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetSpreadBasisPointsIfChainError(&_IFastPriceFeed.TransactOpts, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetSpreadBasisPointsIfChainError(_spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetSpreadBasisPointsIfChainError(&_IFastPriceFeed.TransactOpts, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetSpreadBasisPointsIfInactive(opts *bind.TransactOpts, _spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setSpreadBasisPointsIfInactive", _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetSpreadBasisPointsIfInactive(_spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetSpreadBasisPointsIfInactive(&_IFastPriceFeed.TransactOpts, _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetSpreadBasisPointsIfInactive(_spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetSpreadBasisPointsIfInactive(&_IFastPriceFeed.TransactOpts, _spreadBasisPointsIfInactive)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactor) SetUpdater(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _IFastPriceFeed.contract.Transact(opts, "setUpdater", _account, _isActive)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_IFastPriceFeed *IFastPriceFeedSession) SetUpdater(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetUpdater(&_IFastPriceFeed.TransactOpts, _account, _isActive)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_IFastPriceFeed *IFastPriceFeedTransactorSession) SetUpdater(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _IFastPriceFeed.Contract.SetUpdater(&_IFastPriceFeed.TransactOpts, _account, _isActive)
}
