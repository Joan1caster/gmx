// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pricefeedtimelock

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

// PriceFeedTimelockMetaData contains all meta data concerning the PriceFeedTimelock contract.
var PriceFeedTimelockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_buffer\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"ClearAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalPendingAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultPriceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceDecimals\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isStrictStable\",\"type\":\"bool\"}],\"name\":\"SignalPriceFeedSetTokenConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalSetGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fastPriceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"SignalSetPriceFeedWatcher\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalWithdrawToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BUFFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_action\",\"type\":\"bytes32\"}],\"name\":\"cancelAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKeeper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrictStable\",\"type\":\"bool\"}],\"name\":\"priceFeedSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdditive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_adjustmentBps\",\"type\":\"uint256\"}],\"name\":\"setAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_buffer\",\"type\":\"uint256\"}],\"name\":\"setBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setContractHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setExternalAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsAmmEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSecondaryPriceEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isSpreadEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSpreadEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"setMaxPriceUpdateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxStrictPriceDeviation\",\"type\":\"uint256\"}],\"name\":\"setMaxStrictPriceDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minBlockInterval\",\"type\":\"uint256\"}],\"name\":\"setMinBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDuration\",\"type\":\"uint256\"}],\"name\":\"setPriceDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setPriceFeedUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setPriceFeedWatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceSampleSpace\",\"type\":\"uint256\"}],\"name\":\"setPriceSampleSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfChainError\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfChainError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfInactive\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfInactive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useV2Pricing\",\"type\":\"bool\"}],\"name\":\"setUseV2Pricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrictStable\",\"type\":\"bool\"}],\"name\":\"signalPriceFeedSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"signalSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"signalSetPriceFeedUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"signalSetPriceFeedWatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalWithdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceFeedTimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceFeedTimelockMetaData.ABI instead.
var PriceFeedTimelockABI = PriceFeedTimelockMetaData.ABI

// PriceFeedTimelock is an auto generated Go binding around an Ethereum contract.
type PriceFeedTimelock struct {
	PriceFeedTimelockCaller     // Read-only binding to the contract
	PriceFeedTimelockTransactor // Write-only binding to the contract
	PriceFeedTimelockFilterer   // Log filterer for contract events
}

// PriceFeedTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedTimelockSession struct {
	Contract     *PriceFeedTimelock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PriceFeedTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedTimelockCallerSession struct {
	Contract *PriceFeedTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PriceFeedTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedTimelockTransactorSession struct {
	Contract     *PriceFeedTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PriceFeedTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedTimelockRaw struct {
	Contract *PriceFeedTimelock // Generic contract binding to access the raw methods on
}

// PriceFeedTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedTimelockCallerRaw struct {
	Contract *PriceFeedTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedTimelockTransactorRaw struct {
	Contract *PriceFeedTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedTimelock creates a new instance of PriceFeedTimelock, bound to a specific deployed contract.
func NewPriceFeedTimelock(address common.Address, backend bind.ContractBackend) (*PriceFeedTimelock, error) {
	contract, err := bindPriceFeedTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelock{PriceFeedTimelockCaller: PriceFeedTimelockCaller{contract: contract}, PriceFeedTimelockTransactor: PriceFeedTimelockTransactor{contract: contract}, PriceFeedTimelockFilterer: PriceFeedTimelockFilterer{contract: contract}}, nil
}

// NewPriceFeedTimelockCaller creates a new read-only instance of PriceFeedTimelock, bound to a specific deployed contract.
func NewPriceFeedTimelockCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedTimelockCaller, error) {
	contract, err := bindPriceFeedTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockCaller{contract: contract}, nil
}

// NewPriceFeedTimelockTransactor creates a new write-only instance of PriceFeedTimelock, bound to a specific deployed contract.
func NewPriceFeedTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedTimelockTransactor, error) {
	contract, err := bindPriceFeedTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockTransactor{contract: contract}, nil
}

// NewPriceFeedTimelockFilterer creates a new log filterer instance of PriceFeedTimelock, bound to a specific deployed contract.
func NewPriceFeedTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedTimelockFilterer, error) {
	contract, err := bindPriceFeedTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockFilterer{contract: contract}, nil
}

// bindPriceFeedTimelock binds a generic wrapper to an already deployed contract.
func bindPriceFeedTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceFeedTimelockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedTimelock *PriceFeedTimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedTimelock.Contract.PriceFeedTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedTimelock *PriceFeedTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.PriceFeedTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedTimelock *PriceFeedTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.PriceFeedTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedTimelock *PriceFeedTimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceFeedTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedTimelock *PriceFeedTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedTimelock *PriceFeedTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.contract.Transact(opts, method, params...)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockCaller) MAXBUFFER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedTimelock.contract.Call(opts, &out, "MAX_BUFFER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockSession) MAXBUFFER() (*big.Int, error) {
	return _PriceFeedTimelock.Contract.MAXBUFFER(&_PriceFeedTimelock.CallOpts)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockCallerSession) MAXBUFFER() (*big.Int, error) {
	return _PriceFeedTimelock.Contract.MAXBUFFER(&_PriceFeedTimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PriceFeedTimelock *PriceFeedTimelockCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedTimelock.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PriceFeedTimelock *PriceFeedTimelockSession) Admin() (common.Address, error) {
	return _PriceFeedTimelock.Contract.Admin(&_PriceFeedTimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PriceFeedTimelock *PriceFeedTimelockCallerSession) Admin() (common.Address, error) {
	return _PriceFeedTimelock.Contract.Admin(&_PriceFeedTimelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockCaller) Buffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedTimelock.contract.Call(opts, &out, "buffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockSession) Buffer() (*big.Int, error) {
	return _PriceFeedTimelock.Contract.Buffer(&_PriceFeedTimelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockCallerSession) Buffer() (*big.Int, error) {
	return _PriceFeedTimelock.Contract.Buffer(&_PriceFeedTimelock.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_PriceFeedTimelock *PriceFeedTimelockCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PriceFeedTimelock.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_PriceFeedTimelock *PriceFeedTimelockSession) IsHandler(arg0 common.Address) (bool, error) {
	return _PriceFeedTimelock.Contract.IsHandler(&_PriceFeedTimelock.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_PriceFeedTimelock *PriceFeedTimelockCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _PriceFeedTimelock.Contract.IsHandler(&_PriceFeedTimelock.CallOpts, arg0)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_PriceFeedTimelock *PriceFeedTimelockCaller) IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PriceFeedTimelock.contract.Call(opts, &out, "isKeeper", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_PriceFeedTimelock *PriceFeedTimelockSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _PriceFeedTimelock.Contract.IsKeeper(&_PriceFeedTimelock.CallOpts, arg0)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_PriceFeedTimelock *PriceFeedTimelockCallerSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _PriceFeedTimelock.Contract.IsKeeper(&_PriceFeedTimelock.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockCaller) PendingActions(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedTimelock.contract.Call(opts, &out, "pendingActions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeedTimelock.Contract.PendingActions(&_PriceFeedTimelock.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_PriceFeedTimelock *PriceFeedTimelockCallerSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _PriceFeedTimelock.Contract.PendingActions(&_PriceFeedTimelock.CallOpts, arg0)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_PriceFeedTimelock *PriceFeedTimelockCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedTimelock.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_PriceFeedTimelock *PriceFeedTimelockSession) TokenManager() (common.Address, error) {
	return _PriceFeedTimelock.Contract.TokenManager(&_PriceFeedTimelock.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_PriceFeedTimelock *PriceFeedTimelockCallerSession) TokenManager() (common.Address, error) {
	return _PriceFeedTimelock.Contract.TokenManager(&_PriceFeedTimelock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "approve", _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.Approve(&_PriceFeedTimelock.TransactOpts, _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.Approve(&_PriceFeedTimelock.TransactOpts, _token, _spender, _amount)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) CancelAction(opts *bind.TransactOpts, _action [32]byte) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "cancelAction", _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.CancelAction(&_PriceFeedTimelock.TransactOpts, _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.CancelAction(&_PriceFeedTimelock.TransactOpts, _action)
}

// PriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x3335e38a.
//
// Solidity: function priceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) PriceFeedSetTokenConfig(opts *bind.TransactOpts, _vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "priceFeedSetTokenConfig", _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// PriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x3335e38a.
//
// Solidity: function priceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) PriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.PriceFeedSetTokenConfig(&_PriceFeedTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// PriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x3335e38a.
//
// Solidity: function priceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) PriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.PriceFeedSetTokenConfig(&_PriceFeedTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0x296b07e5.
//
// Solidity: function setAdjustment(address _priceFeed, address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetAdjustment(opts *bind.TransactOpts, _priceFeed common.Address, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setAdjustment", _priceFeed, _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0x296b07e5.
//
// Solidity: function setAdjustment(address _priceFeed, address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetAdjustment(_priceFeed common.Address, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetAdjustment(&_PriceFeedTimelock.TransactOpts, _priceFeed, _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0x296b07e5.
//
// Solidity: function setAdjustment(address _priceFeed, address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetAdjustment(_priceFeed common.Address, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetAdjustment(&_PriceFeedTimelock.TransactOpts, _priceFeed, _token, _isAdditive, _adjustmentBps)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetAdmin(&_PriceFeedTimelock.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetAdmin(&_PriceFeedTimelock.TransactOpts, _admin)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetBuffer(opts *bind.TransactOpts, _buffer *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setBuffer", _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetBuffer(&_PriceFeedTimelock.TransactOpts, _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetBuffer(&_PriceFeedTimelock.TransactOpts, _buffer)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetContractHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setContractHandler", _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetContractHandler(&_PriceFeedTimelock.TransactOpts, _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetContractHandler(&_PriceFeedTimelock.TransactOpts, _handler, _isActive)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetExternalAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setExternalAdmin", _target, _admin)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetExternalAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetExternalAdmin(&_PriceFeedTimelock.TransactOpts, _target, _admin)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetExternalAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetExternalAdmin(&_PriceFeedTimelock.TransactOpts, _target, _admin)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setGov", _target, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetGov(&_PriceFeedTimelock.TransactOpts, _target, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetGov(&_PriceFeedTimelock.TransactOpts, _target, _gov)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x5843752f.
//
// Solidity: function setIsAmmEnabled(address _priceFeed, bool _isEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetIsAmmEnabled(opts *bind.TransactOpts, _priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setIsAmmEnabled", _priceFeed, _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x5843752f.
//
// Solidity: function setIsAmmEnabled(address _priceFeed, bool _isEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetIsAmmEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetIsAmmEnabled(&_PriceFeedTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x5843752f.
//
// Solidity: function setIsAmmEnabled(address _priceFeed, bool _isEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetIsAmmEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetIsAmmEnabled(&_PriceFeedTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0x2965c8c7.
//
// Solidity: function setIsSecondaryPriceEnabled(address _priceFeed, bool _isEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetIsSecondaryPriceEnabled(opts *bind.TransactOpts, _priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setIsSecondaryPriceEnabled", _priceFeed, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0x2965c8c7.
//
// Solidity: function setIsSecondaryPriceEnabled(address _priceFeed, bool _isEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetIsSecondaryPriceEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetIsSecondaryPriceEnabled(&_PriceFeedTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0x2965c8c7.
//
// Solidity: function setIsSecondaryPriceEnabled(address _priceFeed, bool _isEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetIsSecondaryPriceEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetIsSecondaryPriceEnabled(&_PriceFeedTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0x43ec6619.
//
// Solidity: function setIsSpreadEnabled(address _fastPriceFeed, bool _isSpreadEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetIsSpreadEnabled(opts *bind.TransactOpts, _fastPriceFeed common.Address, _isSpreadEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setIsSpreadEnabled", _fastPriceFeed, _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0x43ec6619.
//
// Solidity: function setIsSpreadEnabled(address _fastPriceFeed, bool _isSpreadEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetIsSpreadEnabled(_fastPriceFeed common.Address, _isSpreadEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetIsSpreadEnabled(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0x43ec6619.
//
// Solidity: function setIsSpreadEnabled(address _fastPriceFeed, bool _isSpreadEnabled) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetIsSpreadEnabled(_fastPriceFeed common.Address, _isSpreadEnabled bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetIsSpreadEnabled(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _isSpreadEnabled)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setKeeper", _keeper, _isActive)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetKeeper(_keeper common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetKeeper(&_PriceFeedTimelock.TransactOpts, _keeper, _isActive)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetKeeper(_keeper common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetKeeper(&_PriceFeedTimelock.TransactOpts, _keeper, _isActive)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x5b6348ac.
//
// Solidity: function setMaxPriceUpdateDelay(address _fastPriceFeed, uint256 _maxPriceUpdateDelay) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetMaxPriceUpdateDelay(opts *bind.TransactOpts, _fastPriceFeed common.Address, _maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setMaxPriceUpdateDelay", _fastPriceFeed, _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x5b6348ac.
//
// Solidity: function setMaxPriceUpdateDelay(address _fastPriceFeed, uint256 _maxPriceUpdateDelay) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetMaxPriceUpdateDelay(_fastPriceFeed common.Address, _maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetMaxPriceUpdateDelay(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x5b6348ac.
//
// Solidity: function setMaxPriceUpdateDelay(address _fastPriceFeed, uint256 _maxPriceUpdateDelay) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetMaxPriceUpdateDelay(_fastPriceFeed common.Address, _maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetMaxPriceUpdateDelay(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _maxPriceUpdateDelay)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0xa0a316a2.
//
// Solidity: function setMaxStrictPriceDeviation(address _priceFeed, uint256 _maxStrictPriceDeviation) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetMaxStrictPriceDeviation(opts *bind.TransactOpts, _priceFeed common.Address, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setMaxStrictPriceDeviation", _priceFeed, _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0xa0a316a2.
//
// Solidity: function setMaxStrictPriceDeviation(address _priceFeed, uint256 _maxStrictPriceDeviation) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetMaxStrictPriceDeviation(_priceFeed common.Address, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetMaxStrictPriceDeviation(&_PriceFeedTimelock.TransactOpts, _priceFeed, _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0xa0a316a2.
//
// Solidity: function setMaxStrictPriceDeviation(address _priceFeed, uint256 _maxStrictPriceDeviation) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetMaxStrictPriceDeviation(_priceFeed common.Address, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetMaxStrictPriceDeviation(&_PriceFeedTimelock.TransactOpts, _priceFeed, _maxStrictPriceDeviation)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xbe03af58.
//
// Solidity: function setMinBlockInterval(address _fastPriceFeed, uint256 _minBlockInterval) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetMinBlockInterval(opts *bind.TransactOpts, _fastPriceFeed common.Address, _minBlockInterval *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setMinBlockInterval", _fastPriceFeed, _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xbe03af58.
//
// Solidity: function setMinBlockInterval(address _fastPriceFeed, uint256 _minBlockInterval) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetMinBlockInterval(_fastPriceFeed common.Address, _minBlockInterval *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetMinBlockInterval(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xbe03af58.
//
// Solidity: function setMinBlockInterval(address _fastPriceFeed, uint256 _minBlockInterval) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetMinBlockInterval(_fastPriceFeed common.Address, _minBlockInterval *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetMinBlockInterval(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _minBlockInterval)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0xb74517ba.
//
// Solidity: function setPriceDuration(address _fastPriceFeed, uint256 _priceDuration) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetPriceDuration(opts *bind.TransactOpts, _fastPriceFeed common.Address, _priceDuration *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setPriceDuration", _fastPriceFeed, _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0xb74517ba.
//
// Solidity: function setPriceDuration(address _fastPriceFeed, uint256 _priceDuration) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetPriceDuration(_fastPriceFeed common.Address, _priceDuration *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceDuration(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0xb74517ba.
//
// Solidity: function setPriceDuration(address _fastPriceFeed, uint256 _priceDuration) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetPriceDuration(_fastPriceFeed common.Address, _priceDuration *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceDuration(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _priceDuration)
}

// SetPriceFeedUpdater is a paid mutator transaction binding the contract method 0x50e32d1d.
//
// Solidity: function setPriceFeedUpdater(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetPriceFeedUpdater(opts *bind.TransactOpts, _fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setPriceFeedUpdater", _fastPriceFeed, _account, _isActive)
}

// SetPriceFeedUpdater is a paid mutator transaction binding the contract method 0x50e32d1d.
//
// Solidity: function setPriceFeedUpdater(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetPriceFeedUpdater(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceFeedUpdater(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SetPriceFeedUpdater is a paid mutator transaction binding the contract method 0x50e32d1d.
//
// Solidity: function setPriceFeedUpdater(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetPriceFeedUpdater(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceFeedUpdater(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SetPriceFeedWatcher is a paid mutator transaction binding the contract method 0x2a72e8ba.
//
// Solidity: function setPriceFeedWatcher(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetPriceFeedWatcher(opts *bind.TransactOpts, _fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setPriceFeedWatcher", _fastPriceFeed, _account, _isActive)
}

// SetPriceFeedWatcher is a paid mutator transaction binding the contract method 0x2a72e8ba.
//
// Solidity: function setPriceFeedWatcher(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetPriceFeedWatcher(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceFeedWatcher(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SetPriceFeedWatcher is a paid mutator transaction binding the contract method 0x2a72e8ba.
//
// Solidity: function setPriceFeedWatcher(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetPriceFeedWatcher(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceFeedWatcher(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2877f4c3.
//
// Solidity: function setPriceSampleSpace(address _priceFeed, uint256 _priceSampleSpace) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetPriceSampleSpace(opts *bind.TransactOpts, _priceFeed common.Address, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setPriceSampleSpace", _priceFeed, _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2877f4c3.
//
// Solidity: function setPriceSampleSpace(address _priceFeed, uint256 _priceSampleSpace) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetPriceSampleSpace(_priceFeed common.Address, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceSampleSpace(&_PriceFeedTimelock.TransactOpts, _priceFeed, _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2877f4c3.
//
// Solidity: function setPriceSampleSpace(address _priceFeed, uint256 _priceSampleSpace) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetPriceSampleSpace(_priceFeed common.Address, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetPriceSampleSpace(&_PriceFeedTimelock.TransactOpts, _priceFeed, _priceSampleSpace)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0xe7b0a3a1.
//
// Solidity: function setSpreadBasisPoints(address _priceFeed, address _token, uint256 _spreadBasisPoints) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetSpreadBasisPoints(opts *bind.TransactOpts, _priceFeed common.Address, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setSpreadBasisPoints", _priceFeed, _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0xe7b0a3a1.
//
// Solidity: function setSpreadBasisPoints(address _priceFeed, address _token, uint256 _spreadBasisPoints) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetSpreadBasisPoints(_priceFeed common.Address, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetSpreadBasisPoints(&_PriceFeedTimelock.TransactOpts, _priceFeed, _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0xe7b0a3a1.
//
// Solidity: function setSpreadBasisPoints(address _priceFeed, address _token, uint256 _spreadBasisPoints) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetSpreadBasisPoints(_priceFeed common.Address, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetSpreadBasisPoints(&_PriceFeedTimelock.TransactOpts, _priceFeed, _token, _spreadBasisPoints)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xbd0f1c45.
//
// Solidity: function setSpreadBasisPointsIfChainError(address _fastPriceFeed, uint256 _spreadBasisPointsIfChainError) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetSpreadBasisPointsIfChainError(opts *bind.TransactOpts, _fastPriceFeed common.Address, _spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setSpreadBasisPointsIfChainError", _fastPriceFeed, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xbd0f1c45.
//
// Solidity: function setSpreadBasisPointsIfChainError(address _fastPriceFeed, uint256 _spreadBasisPointsIfChainError) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetSpreadBasisPointsIfChainError(_fastPriceFeed common.Address, _spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetSpreadBasisPointsIfChainError(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xbd0f1c45.
//
// Solidity: function setSpreadBasisPointsIfChainError(address _fastPriceFeed, uint256 _spreadBasisPointsIfChainError) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetSpreadBasisPointsIfChainError(_fastPriceFeed common.Address, _spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetSpreadBasisPointsIfChainError(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0x395bc794.
//
// Solidity: function setSpreadBasisPointsIfInactive(address _fastPriceFeed, uint256 _spreadBasisPointsIfInactive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetSpreadBasisPointsIfInactive(opts *bind.TransactOpts, _fastPriceFeed common.Address, _spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setSpreadBasisPointsIfInactive", _fastPriceFeed, _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0x395bc794.
//
// Solidity: function setSpreadBasisPointsIfInactive(address _fastPriceFeed, uint256 _spreadBasisPointsIfInactive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetSpreadBasisPointsIfInactive(_fastPriceFeed common.Address, _spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetSpreadBasisPointsIfInactive(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0x395bc794.
//
// Solidity: function setSpreadBasisPointsIfInactive(address _fastPriceFeed, uint256 _spreadBasisPointsIfInactive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetSpreadBasisPointsIfInactive(_fastPriceFeed common.Address, _spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetSpreadBasisPointsIfInactive(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _spreadBasisPointsIfInactive)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0x454ffa46.
//
// Solidity: function setUseV2Pricing(address _priceFeed, bool _useV2Pricing) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SetUseV2Pricing(opts *bind.TransactOpts, _priceFeed common.Address, _useV2Pricing bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "setUseV2Pricing", _priceFeed, _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0x454ffa46.
//
// Solidity: function setUseV2Pricing(address _priceFeed, bool _useV2Pricing) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SetUseV2Pricing(_priceFeed common.Address, _useV2Pricing bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetUseV2Pricing(&_PriceFeedTimelock.TransactOpts, _priceFeed, _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0x454ffa46.
//
// Solidity: function setUseV2Pricing(address _priceFeed, bool _useV2Pricing) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SetUseV2Pricing(_priceFeed common.Address, _useV2Pricing bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SetUseV2Pricing(&_PriceFeedTimelock.TransactOpts, _priceFeed, _useV2Pricing)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SignalApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "signalApprove", _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalApprove(&_PriceFeedTimelock.TransactOpts, _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalApprove(&_PriceFeedTimelock.TransactOpts, _token, _spender, _amount)
}

// SignalPriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x384cae73.
//
// Solidity: function signalPriceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SignalPriceFeedSetTokenConfig(opts *bind.TransactOpts, _vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "signalPriceFeedSetTokenConfig", _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SignalPriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x384cae73.
//
// Solidity: function signalPriceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SignalPriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalPriceFeedSetTokenConfig(&_PriceFeedTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SignalPriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x384cae73.
//
// Solidity: function signalPriceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SignalPriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalPriceFeedSetTokenConfig(&_PriceFeedTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SignalSetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "signalSetGov", _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalSetGov(&_PriceFeedTimelock.TransactOpts, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalSetGov(&_PriceFeedTimelock.TransactOpts, _target, _gov)
}

// SignalSetPriceFeedUpdater is a paid mutator transaction binding the contract method 0xdb5c875f.
//
// Solidity: function signalSetPriceFeedUpdater(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SignalSetPriceFeedUpdater(opts *bind.TransactOpts, _fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "signalSetPriceFeedUpdater", _fastPriceFeed, _account, _isActive)
}

// SignalSetPriceFeedUpdater is a paid mutator transaction binding the contract method 0xdb5c875f.
//
// Solidity: function signalSetPriceFeedUpdater(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SignalSetPriceFeedUpdater(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalSetPriceFeedUpdater(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SignalSetPriceFeedUpdater is a paid mutator transaction binding the contract method 0xdb5c875f.
//
// Solidity: function signalSetPriceFeedUpdater(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SignalSetPriceFeedUpdater(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalSetPriceFeedUpdater(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SignalSetPriceFeedWatcher is a paid mutator transaction binding the contract method 0x3799c618.
//
// Solidity: function signalSetPriceFeedWatcher(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SignalSetPriceFeedWatcher(opts *bind.TransactOpts, _fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "signalSetPriceFeedWatcher", _fastPriceFeed, _account, _isActive)
}

// SignalSetPriceFeedWatcher is a paid mutator transaction binding the contract method 0x3799c618.
//
// Solidity: function signalSetPriceFeedWatcher(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SignalSetPriceFeedWatcher(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalSetPriceFeedWatcher(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SignalSetPriceFeedWatcher is a paid mutator transaction binding the contract method 0x3799c618.
//
// Solidity: function signalSetPriceFeedWatcher(address _fastPriceFeed, address _account, bool _isActive) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SignalSetPriceFeedWatcher(_fastPriceFeed common.Address, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalSetPriceFeedWatcher(&_PriceFeedTimelock.TransactOpts, _fastPriceFeed, _account, _isActive)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) SignalWithdrawToken(opts *bind.TransactOpts, _target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "signalWithdrawToken", _target, _token, _receiver, _amount)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) SignalWithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalWithdrawToken(&_PriceFeedTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) SignalWithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.SignalWithdrawToken(&_PriceFeedTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) TransferIn(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "transferIn", _sender, _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) TransferIn(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.TransferIn(&_PriceFeedTimelock.TransactOpts, _sender, _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) TransferIn(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.TransferIn(&_PriceFeedTimelock.TransactOpts, _sender, _token, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactor) WithdrawToken(opts *bind.TransactOpts, _target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.contract.Transact(opts, "withdrawToken", _target, _token, _receiver, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockSession) WithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.WithdrawToken(&_PriceFeedTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_PriceFeedTimelock *PriceFeedTimelockTransactorSession) WithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PriceFeedTimelock.Contract.WithdrawToken(&_PriceFeedTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// PriceFeedTimelockClearActionIterator is returned from FilterClearAction and is used to iterate over the raw logs and unpacked data for ClearAction events raised by the PriceFeedTimelock contract.
type PriceFeedTimelockClearActionIterator struct {
	Event *PriceFeedTimelockClearAction // Event containing the contract specifics and raw log

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
func (it *PriceFeedTimelockClearActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedTimelockClearAction)
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
		it.Event = new(PriceFeedTimelockClearAction)
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
func (it *PriceFeedTimelockClearActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedTimelockClearActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedTimelockClearAction represents a ClearAction event raised by the PriceFeedTimelock contract.
type PriceFeedTimelockClearAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClearAction is a free log retrieval operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) FilterClearAction(opts *bind.FilterOpts) (*PriceFeedTimelockClearActionIterator, error) {

	logs, sub, err := _PriceFeedTimelock.contract.FilterLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockClearActionIterator{contract: _PriceFeedTimelock.contract, event: "ClearAction", logs: logs, sub: sub}, nil
}

// WatchClearAction is a free log subscription operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) WatchClearAction(opts *bind.WatchOpts, sink chan<- *PriceFeedTimelockClearAction) (event.Subscription, error) {

	logs, sub, err := _PriceFeedTimelock.contract.WatchLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedTimelockClearAction)
				if err := _PriceFeedTimelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
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
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) ParseClearAction(log types.Log) (*PriceFeedTimelockClearAction, error) {
	event := new(PriceFeedTimelockClearAction)
	if err := _PriceFeedTimelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedTimelockSignalApproveIterator is returned from FilterSignalApprove and is used to iterate over the raw logs and unpacked data for SignalApprove events raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalApproveIterator struct {
	Event *PriceFeedTimelockSignalApprove // Event containing the contract specifics and raw log

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
func (it *PriceFeedTimelockSignalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedTimelockSignalApprove)
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
		it.Event = new(PriceFeedTimelockSignalApprove)
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
func (it *PriceFeedTimelockSignalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedTimelockSignalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedTimelockSignalApprove represents a SignalApprove event raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalApprove struct {
	Token   common.Address
	Spender common.Address
	Amount  *big.Int
	Action  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignalApprove is a free log retrieval operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) FilterSignalApprove(opts *bind.FilterOpts) (*PriceFeedTimelockSignalApproveIterator, error) {

	logs, sub, err := _PriceFeedTimelock.contract.FilterLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockSignalApproveIterator{contract: _PriceFeedTimelock.contract, event: "SignalApprove", logs: logs, sub: sub}, nil
}

// WatchSignalApprove is a free log subscription operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) WatchSignalApprove(opts *bind.WatchOpts, sink chan<- *PriceFeedTimelockSignalApprove) (event.Subscription, error) {

	logs, sub, err := _PriceFeedTimelock.contract.WatchLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedTimelockSignalApprove)
				if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalApprove", log); err != nil {
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

// ParseSignalApprove is a log parse operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) ParseSignalApprove(log types.Log) (*PriceFeedTimelockSignalApprove, error) {
	event := new(PriceFeedTimelockSignalApprove)
	if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedTimelockSignalPendingActionIterator is returned from FilterSignalPendingAction and is used to iterate over the raw logs and unpacked data for SignalPendingAction events raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalPendingActionIterator struct {
	Event *PriceFeedTimelockSignalPendingAction // Event containing the contract specifics and raw log

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
func (it *PriceFeedTimelockSignalPendingActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedTimelockSignalPendingAction)
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
		it.Event = new(PriceFeedTimelockSignalPendingAction)
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
func (it *PriceFeedTimelockSignalPendingActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedTimelockSignalPendingActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedTimelockSignalPendingAction represents a SignalPendingAction event raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalPendingAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalPendingAction is a free log retrieval operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) FilterSignalPendingAction(opts *bind.FilterOpts) (*PriceFeedTimelockSignalPendingActionIterator, error) {

	logs, sub, err := _PriceFeedTimelock.contract.FilterLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockSignalPendingActionIterator{contract: _PriceFeedTimelock.contract, event: "SignalPendingAction", logs: logs, sub: sub}, nil
}

// WatchSignalPendingAction is a free log subscription operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) WatchSignalPendingAction(opts *bind.WatchOpts, sink chan<- *PriceFeedTimelockSignalPendingAction) (event.Subscription, error) {

	logs, sub, err := _PriceFeedTimelock.contract.WatchLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedTimelockSignalPendingAction)
				if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
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
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) ParseSignalPendingAction(log types.Log) (*PriceFeedTimelockSignalPendingAction, error) {
	event := new(PriceFeedTimelockSignalPendingAction)
	if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedTimelockSignalPriceFeedSetTokenConfigIterator is returned from FilterSignalPriceFeedSetTokenConfig and is used to iterate over the raw logs and unpacked data for SignalPriceFeedSetTokenConfig events raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalPriceFeedSetTokenConfigIterator struct {
	Event *PriceFeedTimelockSignalPriceFeedSetTokenConfig // Event containing the contract specifics and raw log

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
func (it *PriceFeedTimelockSignalPriceFeedSetTokenConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedTimelockSignalPriceFeedSetTokenConfig)
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
		it.Event = new(PriceFeedTimelockSignalPriceFeedSetTokenConfig)
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
func (it *PriceFeedTimelockSignalPriceFeedSetTokenConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedTimelockSignalPriceFeedSetTokenConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedTimelockSignalPriceFeedSetTokenConfig represents a SignalPriceFeedSetTokenConfig event raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalPriceFeedSetTokenConfig struct {
	VaultPriceFeed common.Address
	Token          common.Address
	PriceFeed      common.Address
	PriceDecimals  *big.Int
	IsStrictStable bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSignalPriceFeedSetTokenConfig is a free log retrieval operation binding the contract event 0x1b2ddf357ae016d8c127dcd3a73c34744fdeaeeb4b7ef1e04490cebf7f4816fe.
//
// Solidity: event SignalPriceFeedSetTokenConfig(address vaultPriceFeed, address token, address priceFeed, uint256 priceDecimals, bool isStrictStable)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) FilterSignalPriceFeedSetTokenConfig(opts *bind.FilterOpts) (*PriceFeedTimelockSignalPriceFeedSetTokenConfigIterator, error) {

	logs, sub, err := _PriceFeedTimelock.contract.FilterLogs(opts, "SignalPriceFeedSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockSignalPriceFeedSetTokenConfigIterator{contract: _PriceFeedTimelock.contract, event: "SignalPriceFeedSetTokenConfig", logs: logs, sub: sub}, nil
}

// WatchSignalPriceFeedSetTokenConfig is a free log subscription operation binding the contract event 0x1b2ddf357ae016d8c127dcd3a73c34744fdeaeeb4b7ef1e04490cebf7f4816fe.
//
// Solidity: event SignalPriceFeedSetTokenConfig(address vaultPriceFeed, address token, address priceFeed, uint256 priceDecimals, bool isStrictStable)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) WatchSignalPriceFeedSetTokenConfig(opts *bind.WatchOpts, sink chan<- *PriceFeedTimelockSignalPriceFeedSetTokenConfig) (event.Subscription, error) {

	logs, sub, err := _PriceFeedTimelock.contract.WatchLogs(opts, "SignalPriceFeedSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedTimelockSignalPriceFeedSetTokenConfig)
				if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalPriceFeedSetTokenConfig", log); err != nil {
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

// ParseSignalPriceFeedSetTokenConfig is a log parse operation binding the contract event 0x1b2ddf357ae016d8c127dcd3a73c34744fdeaeeb4b7ef1e04490cebf7f4816fe.
//
// Solidity: event SignalPriceFeedSetTokenConfig(address vaultPriceFeed, address token, address priceFeed, uint256 priceDecimals, bool isStrictStable)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) ParseSignalPriceFeedSetTokenConfig(log types.Log) (*PriceFeedTimelockSignalPriceFeedSetTokenConfig, error) {
	event := new(PriceFeedTimelockSignalPriceFeedSetTokenConfig)
	if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalPriceFeedSetTokenConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedTimelockSignalSetGovIterator is returned from FilterSignalSetGov and is used to iterate over the raw logs and unpacked data for SignalSetGov events raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalSetGovIterator struct {
	Event *PriceFeedTimelockSignalSetGov // Event containing the contract specifics and raw log

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
func (it *PriceFeedTimelockSignalSetGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedTimelockSignalSetGov)
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
		it.Event = new(PriceFeedTimelockSignalSetGov)
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
func (it *PriceFeedTimelockSignalSetGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedTimelockSignalSetGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedTimelockSignalSetGov represents a SignalSetGov event raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalSetGov struct {
	Target common.Address
	Gov    common.Address
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalSetGov is a free log retrieval operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) FilterSignalSetGov(opts *bind.FilterOpts) (*PriceFeedTimelockSignalSetGovIterator, error) {

	logs, sub, err := _PriceFeedTimelock.contract.FilterLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockSignalSetGovIterator{contract: _PriceFeedTimelock.contract, event: "SignalSetGov", logs: logs, sub: sub}, nil
}

// WatchSignalSetGov is a free log subscription operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) WatchSignalSetGov(opts *bind.WatchOpts, sink chan<- *PriceFeedTimelockSignalSetGov) (event.Subscription, error) {

	logs, sub, err := _PriceFeedTimelock.contract.WatchLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedTimelockSignalSetGov)
				if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
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

// ParseSignalSetGov is a log parse operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) ParseSignalSetGov(log types.Log) (*PriceFeedTimelockSignalSetGov, error) {
	event := new(PriceFeedTimelockSignalSetGov)
	if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedTimelockSignalSetPriceFeedWatcherIterator is returned from FilterSignalSetPriceFeedWatcher and is used to iterate over the raw logs and unpacked data for SignalSetPriceFeedWatcher events raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalSetPriceFeedWatcherIterator struct {
	Event *PriceFeedTimelockSignalSetPriceFeedWatcher // Event containing the contract specifics and raw log

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
func (it *PriceFeedTimelockSignalSetPriceFeedWatcherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedTimelockSignalSetPriceFeedWatcher)
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
		it.Event = new(PriceFeedTimelockSignalSetPriceFeedWatcher)
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
func (it *PriceFeedTimelockSignalSetPriceFeedWatcherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedTimelockSignalSetPriceFeedWatcherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedTimelockSignalSetPriceFeedWatcher represents a SignalSetPriceFeedWatcher event raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalSetPriceFeedWatcher struct {
	FastPriceFeed common.Address
	Account       common.Address
	IsActive      bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignalSetPriceFeedWatcher is a free log retrieval operation binding the contract event 0x6ab3018654d3055eae2cb61d3dffe4cbb30f257d54ec966059b4d00b325a3669.
//
// Solidity: event SignalSetPriceFeedWatcher(address fastPriceFeed, address account, bool isActive)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) FilterSignalSetPriceFeedWatcher(opts *bind.FilterOpts) (*PriceFeedTimelockSignalSetPriceFeedWatcherIterator, error) {

	logs, sub, err := _PriceFeedTimelock.contract.FilterLogs(opts, "SignalSetPriceFeedWatcher")
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockSignalSetPriceFeedWatcherIterator{contract: _PriceFeedTimelock.contract, event: "SignalSetPriceFeedWatcher", logs: logs, sub: sub}, nil
}

// WatchSignalSetPriceFeedWatcher is a free log subscription operation binding the contract event 0x6ab3018654d3055eae2cb61d3dffe4cbb30f257d54ec966059b4d00b325a3669.
//
// Solidity: event SignalSetPriceFeedWatcher(address fastPriceFeed, address account, bool isActive)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) WatchSignalSetPriceFeedWatcher(opts *bind.WatchOpts, sink chan<- *PriceFeedTimelockSignalSetPriceFeedWatcher) (event.Subscription, error) {

	logs, sub, err := _PriceFeedTimelock.contract.WatchLogs(opts, "SignalSetPriceFeedWatcher")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedTimelockSignalSetPriceFeedWatcher)
				if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalSetPriceFeedWatcher", log); err != nil {
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

// ParseSignalSetPriceFeedWatcher is a log parse operation binding the contract event 0x6ab3018654d3055eae2cb61d3dffe4cbb30f257d54ec966059b4d00b325a3669.
//
// Solidity: event SignalSetPriceFeedWatcher(address fastPriceFeed, address account, bool isActive)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) ParseSignalSetPriceFeedWatcher(log types.Log) (*PriceFeedTimelockSignalSetPriceFeedWatcher, error) {
	event := new(PriceFeedTimelockSignalSetPriceFeedWatcher)
	if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalSetPriceFeedWatcher", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceFeedTimelockSignalWithdrawTokenIterator is returned from FilterSignalWithdrawToken and is used to iterate over the raw logs and unpacked data for SignalWithdrawToken events raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalWithdrawTokenIterator struct {
	Event *PriceFeedTimelockSignalWithdrawToken // Event containing the contract specifics and raw log

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
func (it *PriceFeedTimelockSignalWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedTimelockSignalWithdrawToken)
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
		it.Event = new(PriceFeedTimelockSignalWithdrawToken)
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
func (it *PriceFeedTimelockSignalWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedTimelockSignalWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedTimelockSignalWithdrawToken represents a SignalWithdrawToken event raised by the PriceFeedTimelock contract.
type PriceFeedTimelockSignalWithdrawToken struct {
	Target   common.Address
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Action   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalWithdrawToken is a free log retrieval operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) FilterSignalWithdrawToken(opts *bind.FilterOpts) (*PriceFeedTimelockSignalWithdrawTokenIterator, error) {

	logs, sub, err := _PriceFeedTimelock.contract.FilterLogs(opts, "SignalWithdrawToken")
	if err != nil {
		return nil, err
	}
	return &PriceFeedTimelockSignalWithdrawTokenIterator{contract: _PriceFeedTimelock.contract, event: "SignalWithdrawToken", logs: logs, sub: sub}, nil
}

// WatchSignalWithdrawToken is a free log subscription operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) WatchSignalWithdrawToken(opts *bind.WatchOpts, sink chan<- *PriceFeedTimelockSignalWithdrawToken) (event.Subscription, error) {

	logs, sub, err := _PriceFeedTimelock.contract.WatchLogs(opts, "SignalWithdrawToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedTimelockSignalWithdrawToken)
				if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalWithdrawToken", log); err != nil {
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

// ParseSignalWithdrawToken is a log parse operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_PriceFeedTimelock *PriceFeedTimelockFilterer) ParseSignalWithdrawToken(log types.Log) (*PriceFeedTimelockSignalWithdrawToken, error) {
	event := new(PriceFeedTimelockSignalWithdrawToken)
	if err := _PriceFeedTimelock.contract.UnpackLog(event, "SignalWithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
