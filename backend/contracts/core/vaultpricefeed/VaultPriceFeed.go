// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultpricefeed

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

// VaultPriceFeedMetaData contains all meta data concerning the VaultPriceFeed contract.
var VaultPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ADJUSTMENT_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ADJUSTMENT_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SPREAD_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ONE_USD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"adjustmentBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnbBusd\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btcBnb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainlinkFlags\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBnb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"favorPrimaryPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getAmmPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_primaryPrice\",\"type\":\"uint256\"}],\"name\":\"getAmmPriceV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getLatestPrimaryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_divByReserve0\",\"type\":\"bool\"}],\"name\":\"getPairPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_includeAmmPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_includeAmmPrice\",\"type\":\"bool\"}],\"name\":\"getPriceV1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_includeAmmPrice\",\"type\":\"bool\"}],\"name\":\"getPriceV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getPrimaryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_referencePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getSecondaryPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdjustmentAdditive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAmmEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSecondaryPriceEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastAdjustmentTimings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStrictPriceDeviation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceSampleSpace\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondaryPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdditive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_adjustmentBps\",\"type\":\"uint256\"}],\"name\":\"setAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainlinkFlags\",\"type\":\"address\"}],\"name\":\"setChainlinkFlags\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_favorPrimaryPrice\",\"type\":\"bool\"}],\"name\":\"setFavorPrimaryPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsAmmEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSecondaryPriceEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxStrictPriceDeviation\",\"type\":\"uint256\"}],\"name\":\"setMaxStrictPriceDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bnbBusd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethBnb\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_btcBnb\",\"type\":\"address\"}],\"name\":\"setPairs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceSampleSpace\",\"type\":\"uint256\"}],\"name\":\"setPriceSampleSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_secondaryPriceFeed\",\"type\":\"address\"}],\"name\":\"setSecondaryPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadThresholdBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadThresholdBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrictStable\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_btc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnb\",\"type\":\"address\"}],\"name\":\"setTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_useV2Pricing\",\"type\":\"bool\"}],\"name\":\"setUseV2Pricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"spreadBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadThresholdBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"strictStableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useV2Pricing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultPriceFeedMetaData.ABI instead.
var VaultPriceFeedABI = VaultPriceFeedMetaData.ABI

// VaultPriceFeed is an auto generated Go binding around an Ethereum contract.
type VaultPriceFeed struct {
	VaultPriceFeedCaller     // Read-only binding to the contract
	VaultPriceFeedTransactor // Write-only binding to the contract
	VaultPriceFeedFilterer   // Log filterer for contract events
}

// VaultPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultPriceFeedSession struct {
	Contract     *VaultPriceFeed   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultPriceFeedCallerSession struct {
	Contract *VaultPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VaultPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultPriceFeedTransactorSession struct {
	Contract     *VaultPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VaultPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultPriceFeedRaw struct {
	Contract *VaultPriceFeed // Generic contract binding to access the raw methods on
}

// VaultPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultPriceFeedCallerRaw struct {
	Contract *VaultPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// VaultPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultPriceFeedTransactorRaw struct {
	Contract *VaultPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultPriceFeed creates a new instance of VaultPriceFeed, bound to a specific deployed contract.
func NewVaultPriceFeed(address common.Address, backend bind.ContractBackend) (*VaultPriceFeed, error) {
	contract, err := bindVaultPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultPriceFeed{VaultPriceFeedCaller: VaultPriceFeedCaller{contract: contract}, VaultPriceFeedTransactor: VaultPriceFeedTransactor{contract: contract}, VaultPriceFeedFilterer: VaultPriceFeedFilterer{contract: contract}}, nil
}

// NewVaultPriceFeedCaller creates a new read-only instance of VaultPriceFeed, bound to a specific deployed contract.
func NewVaultPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*VaultPriceFeedCaller, error) {
	contract, err := bindVaultPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultPriceFeedCaller{contract: contract}, nil
}

// NewVaultPriceFeedTransactor creates a new write-only instance of VaultPriceFeed, bound to a specific deployed contract.
func NewVaultPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultPriceFeedTransactor, error) {
	contract, err := bindVaultPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultPriceFeedTransactor{contract: contract}, nil
}

// NewVaultPriceFeedFilterer creates a new log filterer instance of VaultPriceFeed, bound to a specific deployed contract.
func NewVaultPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultPriceFeedFilterer, error) {
	contract, err := bindVaultPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultPriceFeedFilterer{contract: contract}, nil
}

// bindVaultPriceFeed binds a generic wrapper to an already deployed contract.
func bindVaultPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultPriceFeed *VaultPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultPriceFeed.Contract.VaultPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultPriceFeed *VaultPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.VaultPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultPriceFeed *VaultPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.VaultPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultPriceFeed *VaultPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultPriceFeed *VaultPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultPriceFeed *VaultPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultPriceFeed.Contract.BASISPOINTSDIVISOR(&_VaultPriceFeed.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultPriceFeed.Contract.BASISPOINTSDIVISOR(&_VaultPriceFeed.CallOpts)
}

// MAXADJUSTMENTBASISPOINTS is a free data retrieval call binding the contract method 0x4a4b1f4f.
//
// Solidity: function MAX_ADJUSTMENT_BASIS_POINTS() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) MAXADJUSTMENTBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "MAX_ADJUSTMENT_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADJUSTMENTBASISPOINTS is a free data retrieval call binding the contract method 0x4a4b1f4f.
//
// Solidity: function MAX_ADJUSTMENT_BASIS_POINTS() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) MAXADJUSTMENTBASISPOINTS() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MAXADJUSTMENTBASISPOINTS(&_VaultPriceFeed.CallOpts)
}

// MAXADJUSTMENTBASISPOINTS is a free data retrieval call binding the contract method 0x4a4b1f4f.
//
// Solidity: function MAX_ADJUSTMENT_BASIS_POINTS() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) MAXADJUSTMENTBASISPOINTS() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MAXADJUSTMENTBASISPOINTS(&_VaultPriceFeed.CallOpts)
}

// MAXADJUSTMENTINTERVAL is a free data retrieval call binding the contract method 0x9b18dc47.
//
// Solidity: function MAX_ADJUSTMENT_INTERVAL() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) MAXADJUSTMENTINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "MAX_ADJUSTMENT_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADJUSTMENTINTERVAL is a free data retrieval call binding the contract method 0x9b18dc47.
//
// Solidity: function MAX_ADJUSTMENT_INTERVAL() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) MAXADJUSTMENTINTERVAL() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MAXADJUSTMENTINTERVAL(&_VaultPriceFeed.CallOpts)
}

// MAXADJUSTMENTINTERVAL is a free data retrieval call binding the contract method 0x9b18dc47.
//
// Solidity: function MAX_ADJUSTMENT_INTERVAL() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) MAXADJUSTMENTINTERVAL() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MAXADJUSTMENTINTERVAL(&_VaultPriceFeed.CallOpts)
}

// MAXSPREADBASISPOINTS is a free data retrieval call binding the contract method 0x0957aed9.
//
// Solidity: function MAX_SPREAD_BASIS_POINTS() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) MAXSPREADBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "MAX_SPREAD_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSPREADBASISPOINTS is a free data retrieval call binding the contract method 0x0957aed9.
//
// Solidity: function MAX_SPREAD_BASIS_POINTS() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) MAXSPREADBASISPOINTS() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MAXSPREADBASISPOINTS(&_VaultPriceFeed.CallOpts)
}

// MAXSPREADBASISPOINTS is a free data retrieval call binding the contract method 0x0957aed9.
//
// Solidity: function MAX_SPREAD_BASIS_POINTS() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) MAXSPREADBASISPOINTS() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MAXSPREADBASISPOINTS(&_VaultPriceFeed.CallOpts)
}

// ONEUSD is a free data retrieval call binding the contract method 0x67781c0e.
//
// Solidity: function ONE_USD() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) ONEUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "ONE_USD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ONEUSD is a free data retrieval call binding the contract method 0x67781c0e.
//
// Solidity: function ONE_USD() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) ONEUSD() (*big.Int, error) {
	return _VaultPriceFeed.Contract.ONEUSD(&_VaultPriceFeed.CallOpts)
}

// ONEUSD is a free data retrieval call binding the contract method 0x67781c0e.
//
// Solidity: function ONE_USD() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) ONEUSD() (*big.Int, error) {
	return _VaultPriceFeed.Contract.ONEUSD(&_VaultPriceFeed.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) PRICEPRECISION() (*big.Int, error) {
	return _VaultPriceFeed.Contract.PRICEPRECISION(&_VaultPriceFeed.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _VaultPriceFeed.Contract.PRICEPRECISION(&_VaultPriceFeed.CallOpts)
}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) AdjustmentBasisPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "adjustmentBasisPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) AdjustmentBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.AdjustmentBasisPoints(&_VaultPriceFeed.CallOpts, arg0)
}

// AdjustmentBasisPoints is a free data retrieval call binding the contract method 0x48cac277.
//
// Solidity: function adjustmentBasisPoints(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) AdjustmentBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.AdjustmentBasisPoints(&_VaultPriceFeed.CallOpts, arg0)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) Bnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "bnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) Bnb() (common.Address, error) {
	return _VaultPriceFeed.Contract.Bnb(&_VaultPriceFeed.CallOpts)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) Bnb() (common.Address, error) {
	return _VaultPriceFeed.Contract.Bnb(&_VaultPriceFeed.CallOpts)
}

// BnbBusd is a free data retrieval call binding the contract method 0x97dfade7.
//
// Solidity: function bnbBusd() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) BnbBusd(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "bnbBusd")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnbBusd is a free data retrieval call binding the contract method 0x97dfade7.
//
// Solidity: function bnbBusd() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) BnbBusd() (common.Address, error) {
	return _VaultPriceFeed.Contract.BnbBusd(&_VaultPriceFeed.CallOpts)
}

// BnbBusd is a free data retrieval call binding the contract method 0x97dfade7.
//
// Solidity: function bnbBusd() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) BnbBusd() (common.Address, error) {
	return _VaultPriceFeed.Contract.BnbBusd(&_VaultPriceFeed.CallOpts)
}

// Btc is a free data retrieval call binding the contract method 0xa28d57d8.
//
// Solidity: function btc() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) Btc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "btc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Btc is a free data retrieval call binding the contract method 0xa28d57d8.
//
// Solidity: function btc() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) Btc() (common.Address, error) {
	return _VaultPriceFeed.Contract.Btc(&_VaultPriceFeed.CallOpts)
}

// Btc is a free data retrieval call binding the contract method 0xa28d57d8.
//
// Solidity: function btc() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) Btc() (common.Address, error) {
	return _VaultPriceFeed.Contract.Btc(&_VaultPriceFeed.CallOpts)
}

// BtcBnb is a free data retrieval call binding the contract method 0x971bd396.
//
// Solidity: function btcBnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) BtcBnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "btcBnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcBnb is a free data retrieval call binding the contract method 0x971bd396.
//
// Solidity: function btcBnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) BtcBnb() (common.Address, error) {
	return _VaultPriceFeed.Contract.BtcBnb(&_VaultPriceFeed.CallOpts)
}

// BtcBnb is a free data retrieval call binding the contract method 0x971bd396.
//
// Solidity: function btcBnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) BtcBnb() (common.Address, error) {
	return _VaultPriceFeed.Contract.BtcBnb(&_VaultPriceFeed.CallOpts)
}

// ChainlinkFlags is a free data retrieval call binding the contract method 0xe4440e02.
//
// Solidity: function chainlinkFlags() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) ChainlinkFlags(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "chainlinkFlags")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChainlinkFlags is a free data retrieval call binding the contract method 0xe4440e02.
//
// Solidity: function chainlinkFlags() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) ChainlinkFlags() (common.Address, error) {
	return _VaultPriceFeed.Contract.ChainlinkFlags(&_VaultPriceFeed.CallOpts)
}

// ChainlinkFlags is a free data retrieval call binding the contract method 0xe4440e02.
//
// Solidity: function chainlinkFlags() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) ChainlinkFlags() (common.Address, error) {
	return _VaultPriceFeed.Contract.ChainlinkFlags(&_VaultPriceFeed.CallOpts)
}

// Eth is a free data retrieval call binding the contract method 0x8c7c9e0c.
//
// Solidity: function eth() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) Eth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "eth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eth is a free data retrieval call binding the contract method 0x8c7c9e0c.
//
// Solidity: function eth() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) Eth() (common.Address, error) {
	return _VaultPriceFeed.Contract.Eth(&_VaultPriceFeed.CallOpts)
}

// Eth is a free data retrieval call binding the contract method 0x8c7c9e0c.
//
// Solidity: function eth() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) Eth() (common.Address, error) {
	return _VaultPriceFeed.Contract.Eth(&_VaultPriceFeed.CallOpts)
}

// EthBnb is a free data retrieval call binding the contract method 0x1193c809.
//
// Solidity: function ethBnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) EthBnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "ethBnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthBnb is a free data retrieval call binding the contract method 0x1193c809.
//
// Solidity: function ethBnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) EthBnb() (common.Address, error) {
	return _VaultPriceFeed.Contract.EthBnb(&_VaultPriceFeed.CallOpts)
}

// EthBnb is a free data retrieval call binding the contract method 0x1193c809.
//
// Solidity: function ethBnb() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) EthBnb() (common.Address, error) {
	return _VaultPriceFeed.Contract.EthBnb(&_VaultPriceFeed.CallOpts)
}

// FavorPrimaryPrice is a free data retrieval call binding the contract method 0x593d9e80.
//
// Solidity: function favorPrimaryPrice() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCaller) FavorPrimaryPrice(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "favorPrimaryPrice")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FavorPrimaryPrice is a free data retrieval call binding the contract method 0x593d9e80.
//
// Solidity: function favorPrimaryPrice() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedSession) FavorPrimaryPrice() (bool, error) {
	return _VaultPriceFeed.Contract.FavorPrimaryPrice(&_VaultPriceFeed.CallOpts)
}

// FavorPrimaryPrice is a free data retrieval call binding the contract method 0x593d9e80.
//
// Solidity: function favorPrimaryPrice() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) FavorPrimaryPrice() (bool, error) {
	return _VaultPriceFeed.Contract.FavorPrimaryPrice(&_VaultPriceFeed.CallOpts)
}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetAmmPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getAmmPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetAmmPrice(_token common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetAmmPrice(&_VaultPriceFeed.CallOpts, _token)
}

// GetAmmPrice is a free data retrieval call binding the contract method 0xc2138d8c.
//
// Solidity: function getAmmPrice(address _token) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetAmmPrice(_token common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetAmmPrice(&_VaultPriceFeed.CallOpts, _token)
}

// GetAmmPriceV2 is a free data retrieval call binding the contract method 0xb02a2de4.
//
// Solidity: function getAmmPriceV2(address _token, bool _maximise, uint256 _primaryPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetAmmPriceV2(opts *bind.CallOpts, _token common.Address, _maximise bool, _primaryPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getAmmPriceV2", _token, _maximise, _primaryPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmmPriceV2 is a free data retrieval call binding the contract method 0xb02a2de4.
//
// Solidity: function getAmmPriceV2(address _token, bool _maximise, uint256 _primaryPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetAmmPriceV2(_token common.Address, _maximise bool, _primaryPrice *big.Int) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetAmmPriceV2(&_VaultPriceFeed.CallOpts, _token, _maximise, _primaryPrice)
}

// GetAmmPriceV2 is a free data retrieval call binding the contract method 0xb02a2de4.
//
// Solidity: function getAmmPriceV2(address _token, bool _maximise, uint256 _primaryPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetAmmPriceV2(_token common.Address, _maximise bool, _primaryPrice *big.Int) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetAmmPriceV2(&_VaultPriceFeed.CallOpts, _token, _maximise, _primaryPrice)
}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetLatestPrimaryPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getLatestPrimaryPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetLatestPrimaryPrice(_token common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetLatestPrimaryPrice(&_VaultPriceFeed.CallOpts, _token)
}

// GetLatestPrimaryPrice is a free data retrieval call binding the contract method 0x56bf9de4.
//
// Solidity: function getLatestPrimaryPrice(address _token) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetLatestPrimaryPrice(_token common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetLatestPrimaryPrice(&_VaultPriceFeed.CallOpts, _token)
}

// GetPairPrice is a free data retrieval call binding the contract method 0xa2ad7b93.
//
// Solidity: function getPairPrice(address _pair, bool _divByReserve0) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetPairPrice(opts *bind.CallOpts, _pair common.Address, _divByReserve0 bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getPairPrice", _pair, _divByReserve0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPairPrice is a free data retrieval call binding the contract method 0xa2ad7b93.
//
// Solidity: function getPairPrice(address _pair, bool _divByReserve0) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetPairPrice(_pair common.Address, _divByReserve0 bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPairPrice(&_VaultPriceFeed.CallOpts, _pair, _divByReserve0)
}

// GetPairPrice is a free data retrieval call binding the contract method 0xa2ad7b93.
//
// Solidity: function getPairPrice(address _pair, bool _divByReserve0) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetPairPrice(_pair common.Address, _divByReserve0 bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPairPrice(&_VaultPriceFeed.CallOpts, _pair, _divByReserve0)
}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetPrice(opts *bind.CallOpts, _token common.Address, _maximise bool, _includeAmmPrice bool, arg3 bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getPrice", _token, _maximise, _includeAmmPrice, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetPrice(_token common.Address, _maximise bool, _includeAmmPrice bool, arg3 bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPrice(&_VaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice, arg3)
}

// GetPrice is a free data retrieval call binding the contract method 0x2fc3a70a.
//
// Solidity: function getPrice(address _token, bool _maximise, bool _includeAmmPrice, bool ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetPrice(_token common.Address, _maximise bool, _includeAmmPrice bool, arg3 bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPrice(&_VaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice, arg3)
}

// GetPriceV1 is a free data retrieval call binding the contract method 0x3d949c5f.
//
// Solidity: function getPriceV1(address _token, bool _maximise, bool _includeAmmPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetPriceV1(opts *bind.CallOpts, _token common.Address, _maximise bool, _includeAmmPrice bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getPriceV1", _token, _maximise, _includeAmmPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceV1 is a free data retrieval call binding the contract method 0x3d949c5f.
//
// Solidity: function getPriceV1(address _token, bool _maximise, bool _includeAmmPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetPriceV1(_token common.Address, _maximise bool, _includeAmmPrice bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPriceV1(&_VaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice)
}

// GetPriceV1 is a free data retrieval call binding the contract method 0x3d949c5f.
//
// Solidity: function getPriceV1(address _token, bool _maximise, bool _includeAmmPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetPriceV1(_token common.Address, _maximise bool, _includeAmmPrice bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPriceV1(&_VaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice)
}

// GetPriceV2 is a free data retrieval call binding the contract method 0x732391b4.
//
// Solidity: function getPriceV2(address _token, bool _maximise, bool _includeAmmPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetPriceV2(opts *bind.CallOpts, _token common.Address, _maximise bool, _includeAmmPrice bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getPriceV2", _token, _maximise, _includeAmmPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceV2 is a free data retrieval call binding the contract method 0x732391b4.
//
// Solidity: function getPriceV2(address _token, bool _maximise, bool _includeAmmPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetPriceV2(_token common.Address, _maximise bool, _includeAmmPrice bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPriceV2(&_VaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice)
}

// GetPriceV2 is a free data retrieval call binding the contract method 0x732391b4.
//
// Solidity: function getPriceV2(address _token, bool _maximise, bool _includeAmmPrice) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetPriceV2(_token common.Address, _maximise bool, _includeAmmPrice bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPriceV2(&_VaultPriceFeed.CallOpts, _token, _maximise, _includeAmmPrice)
}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetPrimaryPrice(opts *bind.CallOpts, _token common.Address, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getPrimaryPrice", _token, _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetPrimaryPrice(_token common.Address, _maximise bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPrimaryPrice(&_VaultPriceFeed.CallOpts, _token, _maximise)
}

// GetPrimaryPrice is a free data retrieval call binding the contract method 0x56c8c2c1.
//
// Solidity: function getPrimaryPrice(address _token, bool _maximise) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetPrimaryPrice(_token common.Address, _maximise bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetPrimaryPrice(&_VaultPriceFeed.CallOpts, _token, _maximise)
}

// GetSecondaryPrice is a free data retrieval call binding the contract method 0x3eba8d36.
//
// Solidity: function getSecondaryPrice(address _token, uint256 _referencePrice, bool _maximise) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) GetSecondaryPrice(opts *bind.CallOpts, _token common.Address, _referencePrice *big.Int, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "getSecondaryPrice", _token, _referencePrice, _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecondaryPrice is a free data retrieval call binding the contract method 0x3eba8d36.
//
// Solidity: function getSecondaryPrice(address _token, uint256 _referencePrice, bool _maximise) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) GetSecondaryPrice(_token common.Address, _referencePrice *big.Int, _maximise bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetSecondaryPrice(&_VaultPriceFeed.CallOpts, _token, _referencePrice, _maximise)
}

// GetSecondaryPrice is a free data retrieval call binding the contract method 0x3eba8d36.
//
// Solidity: function getSecondaryPrice(address _token, uint256 _referencePrice, bool _maximise) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) GetSecondaryPrice(_token common.Address, _referencePrice *big.Int, _maximise bool) (*big.Int, error) {
	return _VaultPriceFeed.Contract.GetSecondaryPrice(&_VaultPriceFeed.CallOpts, _token, _referencePrice, _maximise)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) Gov() (common.Address, error) {
	return _VaultPriceFeed.Contract.Gov(&_VaultPriceFeed.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) Gov() (common.Address, error) {
	return _VaultPriceFeed.Contract.Gov(&_VaultPriceFeed.CallOpts)
}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address ) view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCaller) IsAdjustmentAdditive(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "isAdjustmentAdditive", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address ) view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedSession) IsAdjustmentAdditive(arg0 common.Address) (bool, error) {
	return _VaultPriceFeed.Contract.IsAdjustmentAdditive(&_VaultPriceFeed.CallOpts, arg0)
}

// IsAdjustmentAdditive is a free data retrieval call binding the contract method 0x6ce8a44b.
//
// Solidity: function isAdjustmentAdditive(address ) view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) IsAdjustmentAdditive(arg0 common.Address) (bool, error) {
	return _VaultPriceFeed.Contract.IsAdjustmentAdditive(&_VaultPriceFeed.CallOpts, arg0)
}

// IsAmmEnabled is a free data retrieval call binding the contract method 0x3f0c3bb7.
//
// Solidity: function isAmmEnabled() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCaller) IsAmmEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "isAmmEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAmmEnabled is a free data retrieval call binding the contract method 0x3f0c3bb7.
//
// Solidity: function isAmmEnabled() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedSession) IsAmmEnabled() (bool, error) {
	return _VaultPriceFeed.Contract.IsAmmEnabled(&_VaultPriceFeed.CallOpts)
}

// IsAmmEnabled is a free data retrieval call binding the contract method 0x3f0c3bb7.
//
// Solidity: function isAmmEnabled() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) IsAmmEnabled() (bool, error) {
	return _VaultPriceFeed.Contract.IsAmmEnabled(&_VaultPriceFeed.CallOpts)
}

// IsSecondaryPriceEnabled is a free data retrieval call binding the contract method 0x3ebbc601.
//
// Solidity: function isSecondaryPriceEnabled() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCaller) IsSecondaryPriceEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "isSecondaryPriceEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSecondaryPriceEnabled is a free data retrieval call binding the contract method 0x3ebbc601.
//
// Solidity: function isSecondaryPriceEnabled() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedSession) IsSecondaryPriceEnabled() (bool, error) {
	return _VaultPriceFeed.Contract.IsSecondaryPriceEnabled(&_VaultPriceFeed.CallOpts)
}

// IsSecondaryPriceEnabled is a free data retrieval call binding the contract method 0x3ebbc601.
//
// Solidity: function isSecondaryPriceEnabled() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) IsSecondaryPriceEnabled() (bool, error) {
	return _VaultPriceFeed.Contract.IsSecondaryPriceEnabled(&_VaultPriceFeed.CallOpts)
}

// LastAdjustmentTimings is a free data retrieval call binding the contract method 0x717cfe7a.
//
// Solidity: function lastAdjustmentTimings(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) LastAdjustmentTimings(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "lastAdjustmentTimings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAdjustmentTimings is a free data retrieval call binding the contract method 0x717cfe7a.
//
// Solidity: function lastAdjustmentTimings(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) LastAdjustmentTimings(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.LastAdjustmentTimings(&_VaultPriceFeed.CallOpts, arg0)
}

// LastAdjustmentTimings is a free data retrieval call binding the contract method 0x717cfe7a.
//
// Solidity: function lastAdjustmentTimings(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) LastAdjustmentTimings(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.LastAdjustmentTimings(&_VaultPriceFeed.CallOpts, arg0)
}

// MaxStrictPriceDeviation is a free data retrieval call binding the contract method 0x378e7bf7.
//
// Solidity: function maxStrictPriceDeviation() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) MaxStrictPriceDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "maxStrictPriceDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStrictPriceDeviation is a free data retrieval call binding the contract method 0x378e7bf7.
//
// Solidity: function maxStrictPriceDeviation() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) MaxStrictPriceDeviation() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MaxStrictPriceDeviation(&_VaultPriceFeed.CallOpts)
}

// MaxStrictPriceDeviation is a free data retrieval call binding the contract method 0x378e7bf7.
//
// Solidity: function maxStrictPriceDeviation() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) MaxStrictPriceDeviation() (*big.Int, error) {
	return _VaultPriceFeed.Contract.MaxStrictPriceDeviation(&_VaultPriceFeed.CallOpts)
}

// PriceDecimals is a free data retrieval call binding the contract method 0xcefe0f21.
//
// Solidity: function priceDecimals(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) PriceDecimals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "priceDecimals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDecimals is a free data retrieval call binding the contract method 0xcefe0f21.
//
// Solidity: function priceDecimals(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) PriceDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.PriceDecimals(&_VaultPriceFeed.CallOpts, arg0)
}

// PriceDecimals is a free data retrieval call binding the contract method 0xcefe0f21.
//
// Solidity: function priceDecimals(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) PriceDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.PriceDecimals(&_VaultPriceFeed.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "priceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _VaultPriceFeed.Contract.PriceFeeds(&_VaultPriceFeed.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _VaultPriceFeed.Contract.PriceFeeds(&_VaultPriceFeed.CallOpts, arg0)
}

// PriceSampleSpace is a free data retrieval call binding the contract method 0x6fc80708.
//
// Solidity: function priceSampleSpace() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) PriceSampleSpace(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "priceSampleSpace")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceSampleSpace is a free data retrieval call binding the contract method 0x6fc80708.
//
// Solidity: function priceSampleSpace() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) PriceSampleSpace() (*big.Int, error) {
	return _VaultPriceFeed.Contract.PriceSampleSpace(&_VaultPriceFeed.CallOpts)
}

// PriceSampleSpace is a free data retrieval call binding the contract method 0x6fc80708.
//
// Solidity: function priceSampleSpace() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) PriceSampleSpace() (*big.Int, error) {
	return _VaultPriceFeed.Contract.PriceSampleSpace(&_VaultPriceFeed.CallOpts)
}

// SecondaryPriceFeed is a free data retrieval call binding the contract method 0x8b86616c.
//
// Solidity: function secondaryPriceFeed() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCaller) SecondaryPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "secondaryPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SecondaryPriceFeed is a free data retrieval call binding the contract method 0x8b86616c.
//
// Solidity: function secondaryPriceFeed() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedSession) SecondaryPriceFeed() (common.Address, error) {
	return _VaultPriceFeed.Contract.SecondaryPriceFeed(&_VaultPriceFeed.CallOpts)
}

// SecondaryPriceFeed is a free data retrieval call binding the contract method 0x8b86616c.
//
// Solidity: function secondaryPriceFeed() view returns(address)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) SecondaryPriceFeed() (common.Address, error) {
	return _VaultPriceFeed.Contract.SecondaryPriceFeed(&_VaultPriceFeed.CallOpts)
}

// SpreadBasisPoints is a free data retrieval call binding the contract method 0xa27ea386.
//
// Solidity: function spreadBasisPoints(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) SpreadBasisPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "spreadBasisPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadBasisPoints is a free data retrieval call binding the contract method 0xa27ea386.
//
// Solidity: function spreadBasisPoints(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) SpreadBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.SpreadBasisPoints(&_VaultPriceFeed.CallOpts, arg0)
}

// SpreadBasisPoints is a free data retrieval call binding the contract method 0xa27ea386.
//
// Solidity: function spreadBasisPoints(address ) view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) SpreadBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultPriceFeed.Contract.SpreadBasisPoints(&_VaultPriceFeed.CallOpts, arg0)
}

// SpreadThresholdBasisPoints is a free data retrieval call binding the contract method 0xa39c73a3.
//
// Solidity: function spreadThresholdBasisPoints() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCaller) SpreadThresholdBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "spreadThresholdBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadThresholdBasisPoints is a free data retrieval call binding the contract method 0xa39c73a3.
//
// Solidity: function spreadThresholdBasisPoints() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedSession) SpreadThresholdBasisPoints() (*big.Int, error) {
	return _VaultPriceFeed.Contract.SpreadThresholdBasisPoints(&_VaultPriceFeed.CallOpts)
}

// SpreadThresholdBasisPoints is a free data retrieval call binding the contract method 0xa39c73a3.
//
// Solidity: function spreadThresholdBasisPoints() view returns(uint256)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) SpreadThresholdBasisPoints() (*big.Int, error) {
	return _VaultPriceFeed.Contract.SpreadThresholdBasisPoints(&_VaultPriceFeed.CallOpts)
}

// StrictStableTokens is a free data retrieval call binding the contract method 0xb8f61105.
//
// Solidity: function strictStableTokens(address ) view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCaller) StrictStableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "strictStableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrictStableTokens is a free data retrieval call binding the contract method 0xb8f61105.
//
// Solidity: function strictStableTokens(address ) view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedSession) StrictStableTokens(arg0 common.Address) (bool, error) {
	return _VaultPriceFeed.Contract.StrictStableTokens(&_VaultPriceFeed.CallOpts, arg0)
}

// StrictStableTokens is a free data retrieval call binding the contract method 0xb8f61105.
//
// Solidity: function strictStableTokens(address ) view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) StrictStableTokens(arg0 common.Address) (bool, error) {
	return _VaultPriceFeed.Contract.StrictStableTokens(&_VaultPriceFeed.CallOpts, arg0)
}

// UseV2Pricing is a free data retrieval call binding the contract method 0x30536ee5.
//
// Solidity: function useV2Pricing() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCaller) UseV2Pricing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultPriceFeed.contract.Call(opts, &out, "useV2Pricing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseV2Pricing is a free data retrieval call binding the contract method 0x30536ee5.
//
// Solidity: function useV2Pricing() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedSession) UseV2Pricing() (bool, error) {
	return _VaultPriceFeed.Contract.UseV2Pricing(&_VaultPriceFeed.CallOpts)
}

// UseV2Pricing is a free data retrieval call binding the contract method 0x30536ee5.
//
// Solidity: function useV2Pricing() view returns(bool)
func (_VaultPriceFeed *VaultPriceFeedCallerSession) UseV2Pricing() (bool, error) {
	return _VaultPriceFeed.Contract.UseV2Pricing(&_VaultPriceFeed.CallOpts)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetAdjustment(opts *bind.TransactOpts, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setAdjustment", _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetAdjustment(_token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetAdjustment(&_VaultPriceFeed.TransactOpts, _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0xd694376c.
//
// Solidity: function setAdjustment(address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetAdjustment(_token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetAdjustment(&_VaultPriceFeed.TransactOpts, _token, _isAdditive, _adjustmentBps)
}

// SetChainlinkFlags is a paid mutator transaction binding the contract method 0x826e055f.
//
// Solidity: function setChainlinkFlags(address _chainlinkFlags) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetChainlinkFlags(opts *bind.TransactOpts, _chainlinkFlags common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setChainlinkFlags", _chainlinkFlags)
}

// SetChainlinkFlags is a paid mutator transaction binding the contract method 0x826e055f.
//
// Solidity: function setChainlinkFlags(address _chainlinkFlags) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetChainlinkFlags(_chainlinkFlags common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetChainlinkFlags(&_VaultPriceFeed.TransactOpts, _chainlinkFlags)
}

// SetChainlinkFlags is a paid mutator transaction binding the contract method 0x826e055f.
//
// Solidity: function setChainlinkFlags(address _chainlinkFlags) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetChainlinkFlags(_chainlinkFlags common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetChainlinkFlags(&_VaultPriceFeed.TransactOpts, _chainlinkFlags)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetFavorPrimaryPrice(opts *bind.TransactOpts, _favorPrimaryPrice bool) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setFavorPrimaryPrice", _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetFavorPrimaryPrice(_favorPrimaryPrice bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetFavorPrimaryPrice(&_VaultPriceFeed.TransactOpts, _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x604f37e9.
//
// Solidity: function setFavorPrimaryPrice(bool _favorPrimaryPrice) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetFavorPrimaryPrice(_favorPrimaryPrice bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetFavorPrimaryPrice(&_VaultPriceFeed.TransactOpts, _favorPrimaryPrice)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetGov(&_VaultPriceFeed.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetGov(&_VaultPriceFeed.TransactOpts, _gov)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetIsAmmEnabled(opts *bind.TransactOpts, _isEnabled bool) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setIsAmmEnabled", _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetIsAmmEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetIsAmmEnabled(&_VaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x9917dc74.
//
// Solidity: function setIsAmmEnabled(bool _isEnabled) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetIsAmmEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetIsAmmEnabled(&_VaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetIsSecondaryPriceEnabled(opts *bind.TransactOpts, _isEnabled bool) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setIsSecondaryPriceEnabled", _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetIsSecondaryPriceEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetIsSecondaryPriceEnabled(&_VaultPriceFeed.TransactOpts, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0xeb1c92a9.
//
// Solidity: function setIsSecondaryPriceEnabled(bool _isEnabled) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetIsSecondaryPriceEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetIsSecondaryPriceEnabled(&_VaultPriceFeed.TransactOpts, _isEnabled)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetMaxStrictPriceDeviation(opts *bind.TransactOpts, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setMaxStrictPriceDeviation", _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetMaxStrictPriceDeviation(_maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetMaxStrictPriceDeviation(&_VaultPriceFeed.TransactOpts, _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0x2fbfe3d3.
//
// Solidity: function setMaxStrictPriceDeviation(uint256 _maxStrictPriceDeviation) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetMaxStrictPriceDeviation(_maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetMaxStrictPriceDeviation(&_VaultPriceFeed.TransactOpts, _maxStrictPriceDeviation)
}

// SetPairs is a paid mutator transaction binding the contract method 0x93f69074.
//
// Solidity: function setPairs(address _bnbBusd, address _ethBnb, address _btcBnb) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetPairs(opts *bind.TransactOpts, _bnbBusd common.Address, _ethBnb common.Address, _btcBnb common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setPairs", _bnbBusd, _ethBnb, _btcBnb)
}

// SetPairs is a paid mutator transaction binding the contract method 0x93f69074.
//
// Solidity: function setPairs(address _bnbBusd, address _ethBnb, address _btcBnb) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetPairs(_bnbBusd common.Address, _ethBnb common.Address, _btcBnb common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetPairs(&_VaultPriceFeed.TransactOpts, _bnbBusd, _ethBnb, _btcBnb)
}

// SetPairs is a paid mutator transaction binding the contract method 0x93f69074.
//
// Solidity: function setPairs(address _bnbBusd, address _ethBnb, address _btcBnb) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetPairs(_bnbBusd common.Address, _ethBnb common.Address, _btcBnb common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetPairs(&_VaultPriceFeed.TransactOpts, _bnbBusd, _ethBnb, _btcBnb)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetPriceSampleSpace(opts *bind.TransactOpts, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setPriceSampleSpace", _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetPriceSampleSpace(_priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetPriceSampleSpace(&_VaultPriceFeed.TransactOpts, _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2fa03b8f.
//
// Solidity: function setPriceSampleSpace(uint256 _priceSampleSpace) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetPriceSampleSpace(_priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetPriceSampleSpace(&_VaultPriceFeed.TransactOpts, _priceSampleSpace)
}

// SetSecondaryPriceFeed is a paid mutator transaction binding the contract method 0x9a0a6635.
//
// Solidity: function setSecondaryPriceFeed(address _secondaryPriceFeed) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetSecondaryPriceFeed(opts *bind.TransactOpts, _secondaryPriceFeed common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setSecondaryPriceFeed", _secondaryPriceFeed)
}

// SetSecondaryPriceFeed is a paid mutator transaction binding the contract method 0x9a0a6635.
//
// Solidity: function setSecondaryPriceFeed(address _secondaryPriceFeed) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetSecondaryPriceFeed(_secondaryPriceFeed common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetSecondaryPriceFeed(&_VaultPriceFeed.TransactOpts, _secondaryPriceFeed)
}

// SetSecondaryPriceFeed is a paid mutator transaction binding the contract method 0x9a0a6635.
//
// Solidity: function setSecondaryPriceFeed(address _secondaryPriceFeed) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetSecondaryPriceFeed(_secondaryPriceFeed common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetSecondaryPriceFeed(&_VaultPriceFeed.TransactOpts, _secondaryPriceFeed)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetSpreadBasisPoints(opts *bind.TransactOpts, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setSpreadBasisPoints", _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetSpreadBasisPoints(_token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetSpreadBasisPoints(&_VaultPriceFeed.TransactOpts, _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0x9b889380.
//
// Solidity: function setSpreadBasisPoints(address _token, uint256 _spreadBasisPoints) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetSpreadBasisPoints(_token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetSpreadBasisPoints(&_VaultPriceFeed.TransactOpts, _token, _spreadBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetSpreadThresholdBasisPoints(opts *bind.TransactOpts, _spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setSpreadThresholdBasisPoints", _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetSpreadThresholdBasisPoints(_spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetSpreadThresholdBasisPoints(&_VaultPriceFeed.TransactOpts, _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0xb731dd87.
//
// Solidity: function setSpreadThresholdBasisPoints(uint256 _spreadThresholdBasisPoints) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetSpreadThresholdBasisPoints(_spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetSpreadThresholdBasisPoints(&_VaultPriceFeed.TransactOpts, _spreadThresholdBasisPoints)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setTokenConfig", _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetTokenConfig(_token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetTokenConfig(&_VaultPriceFeed.TransactOpts, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x4b9ade47.
//
// Solidity: function setTokenConfig(address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetTokenConfig(_token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetTokenConfig(&_VaultPriceFeed.TransactOpts, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetTokens is a paid mutator transaction binding the contract method 0x443be209.
//
// Solidity: function setTokens(address _btc, address _eth, address _bnb) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetTokens(opts *bind.TransactOpts, _btc common.Address, _eth common.Address, _bnb common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setTokens", _btc, _eth, _bnb)
}

// SetTokens is a paid mutator transaction binding the contract method 0x443be209.
//
// Solidity: function setTokens(address _btc, address _eth, address _bnb) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetTokens(_btc common.Address, _eth common.Address, _bnb common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetTokens(&_VaultPriceFeed.TransactOpts, _btc, _eth, _bnb)
}

// SetTokens is a paid mutator transaction binding the contract method 0x443be209.
//
// Solidity: function setTokens(address _btc, address _eth, address _bnb) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetTokens(_btc common.Address, _eth common.Address, _bnb common.Address) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetTokens(&_VaultPriceFeed.TransactOpts, _btc, _eth, _bnb)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactor) SetUseV2Pricing(opts *bind.TransactOpts, _useV2Pricing bool) (*types.Transaction, error) {
	return _VaultPriceFeed.contract.Transact(opts, "setUseV2Pricing", _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_VaultPriceFeed *VaultPriceFeedSession) SetUseV2Pricing(_useV2Pricing bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetUseV2Pricing(&_VaultPriceFeed.TransactOpts, _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0xfd34ec40.
//
// Solidity: function setUseV2Pricing(bool _useV2Pricing) returns()
func (_VaultPriceFeed *VaultPriceFeedTransactorSession) SetUseV2Pricing(_useV2Pricing bool) (*types.Transaction, error) {
	return _VaultPriceFeed.Contract.SetUseV2Pricing(&_VaultPriceFeed.TransactOpts, _useV2Pricing)
}
