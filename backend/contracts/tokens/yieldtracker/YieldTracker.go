// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yieldtracker

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

// YieldTrackerMetaData contains all meta data concerning the YieldTracker contract.
var YieldTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_yieldToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"previousCumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YieldTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use YieldTrackerMetaData.ABI instead.
var YieldTrackerABI = YieldTrackerMetaData.ABI

// YieldTracker is an auto generated Go binding around an Ethereum contract.
type YieldTracker struct {
	YieldTrackerCaller     // Read-only binding to the contract
	YieldTrackerTransactor // Write-only binding to the contract
	YieldTrackerFilterer   // Log filterer for contract events
}

// YieldTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldTrackerSession struct {
	Contract     *YieldTracker     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldTrackerCallerSession struct {
	Contract *YieldTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// YieldTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldTrackerTransactorSession struct {
	Contract     *YieldTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// YieldTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldTrackerRaw struct {
	Contract *YieldTracker // Generic contract binding to access the raw methods on
}

// YieldTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldTrackerCallerRaw struct {
	Contract *YieldTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// YieldTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldTrackerTransactorRaw struct {
	Contract *YieldTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldTracker creates a new instance of YieldTracker, bound to a specific deployed contract.
func NewYieldTracker(address common.Address, backend bind.ContractBackend) (*YieldTracker, error) {
	contract, err := bindYieldTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YieldTracker{YieldTrackerCaller: YieldTrackerCaller{contract: contract}, YieldTrackerTransactor: YieldTrackerTransactor{contract: contract}, YieldTrackerFilterer: YieldTrackerFilterer{contract: contract}}, nil
}

// NewYieldTrackerCaller creates a new read-only instance of YieldTracker, bound to a specific deployed contract.
func NewYieldTrackerCaller(address common.Address, caller bind.ContractCaller) (*YieldTrackerCaller, error) {
	contract, err := bindYieldTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTrackerCaller{contract: contract}, nil
}

// NewYieldTrackerTransactor creates a new write-only instance of YieldTracker, bound to a specific deployed contract.
func NewYieldTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldTrackerTransactor, error) {
	contract, err := bindYieldTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTrackerTransactor{contract: contract}, nil
}

// NewYieldTrackerFilterer creates a new log filterer instance of YieldTracker, bound to a specific deployed contract.
func NewYieldTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldTrackerFilterer, error) {
	contract, err := bindYieldTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldTrackerFilterer{contract: contract}, nil
}

// bindYieldTracker binds a generic wrapper to an already deployed contract.
func bindYieldTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YieldTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldTracker *YieldTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldTracker.Contract.YieldTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldTracker *YieldTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldTracker.Contract.YieldTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldTracker *YieldTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldTracker.Contract.YieldTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldTracker *YieldTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldTracker *YieldTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldTracker *YieldTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldTracker.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_YieldTracker *YieldTrackerCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_YieldTracker *YieldTrackerSession) PRECISION() (*big.Int, error) {
	return _YieldTracker.Contract.PRECISION(&_YieldTracker.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_YieldTracker *YieldTrackerCallerSession) PRECISION() (*big.Int, error) {
	return _YieldTracker.Contract.PRECISION(&_YieldTracker.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_YieldTracker *YieldTrackerCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_YieldTracker *YieldTrackerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _YieldTracker.Contract.Claimable(&_YieldTracker.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_YieldTracker *YieldTrackerCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _YieldTracker.Contract.Claimable(&_YieldTracker.CallOpts, _account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_YieldTracker *YieldTrackerCaller) ClaimableReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "claimableReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_YieldTracker *YieldTrackerSession) ClaimableReward(arg0 common.Address) (*big.Int, error) {
	return _YieldTracker.Contract.ClaimableReward(&_YieldTracker.CallOpts, arg0)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_YieldTracker *YieldTrackerCallerSession) ClaimableReward(arg0 common.Address) (*big.Int, error) {
	return _YieldTracker.Contract.ClaimableReward(&_YieldTracker.CallOpts, arg0)
}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_YieldTracker *YieldTrackerCaller) CumulativeRewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "cumulativeRewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_YieldTracker *YieldTrackerSession) CumulativeRewardPerToken() (*big.Int, error) {
	return _YieldTracker.Contract.CumulativeRewardPerToken(&_YieldTracker.CallOpts)
}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_YieldTracker *YieldTrackerCallerSession) CumulativeRewardPerToken() (*big.Int, error) {
	return _YieldTracker.Contract.CumulativeRewardPerToken(&_YieldTracker.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_YieldTracker *YieldTrackerCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_YieldTracker *YieldTrackerSession) Distributor() (common.Address, error) {
	return _YieldTracker.Contract.Distributor(&_YieldTracker.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_YieldTracker *YieldTrackerCallerSession) Distributor() (common.Address, error) {
	return _YieldTracker.Contract.Distributor(&_YieldTracker.CallOpts)
}

// GetTokensPerInterval is a free data retrieval call binding the contract method 0x2459f51d.
//
// Solidity: function getTokensPerInterval() view returns(uint256)
func (_YieldTracker *YieldTrackerCaller) GetTokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "getTokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokensPerInterval is a free data retrieval call binding the contract method 0x2459f51d.
//
// Solidity: function getTokensPerInterval() view returns(uint256)
func (_YieldTracker *YieldTrackerSession) GetTokensPerInterval() (*big.Int, error) {
	return _YieldTracker.Contract.GetTokensPerInterval(&_YieldTracker.CallOpts)
}

// GetTokensPerInterval is a free data retrieval call binding the contract method 0x2459f51d.
//
// Solidity: function getTokensPerInterval() view returns(uint256)
func (_YieldTracker *YieldTrackerCallerSession) GetTokensPerInterval() (*big.Int, error) {
	return _YieldTracker.Contract.GetTokensPerInterval(&_YieldTracker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldTracker *YieldTrackerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldTracker *YieldTrackerSession) Gov() (common.Address, error) {
	return _YieldTracker.Contract.Gov(&_YieldTracker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldTracker *YieldTrackerCallerSession) Gov() (common.Address, error) {
	return _YieldTracker.Contract.Gov(&_YieldTracker.CallOpts)
}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_YieldTracker *YieldTrackerCaller) PreviousCumulatedRewardPerToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "previousCumulatedRewardPerToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_YieldTracker *YieldTrackerSession) PreviousCumulatedRewardPerToken(arg0 common.Address) (*big.Int, error) {
	return _YieldTracker.Contract.PreviousCumulatedRewardPerToken(&_YieldTracker.CallOpts, arg0)
}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_YieldTracker *YieldTrackerCallerSession) PreviousCumulatedRewardPerToken(arg0 common.Address) (*big.Int, error) {
	return _YieldTracker.Contract.PreviousCumulatedRewardPerToken(&_YieldTracker.CallOpts, arg0)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_YieldTracker *YieldTrackerCaller) YieldToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldTracker.contract.Call(opts, &out, "yieldToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_YieldTracker *YieldTrackerSession) YieldToken() (common.Address, error) {
	return _YieldTracker.Contract.YieldToken(&_YieldTracker.CallOpts)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_YieldTracker *YieldTrackerCallerSession) YieldToken() (common.Address, error) {
	return _YieldTracker.Contract.YieldToken(&_YieldTracker.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address _account, address _receiver) returns(uint256)
func (_YieldTracker *YieldTrackerTransactor) Claim(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldTracker.contract.Transact(opts, "claim", _account, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address _account, address _receiver) returns(uint256)
func (_YieldTracker *YieldTrackerSession) Claim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.Claim(&_YieldTracker.TransactOpts, _account, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address _account, address _receiver) returns(uint256)
func (_YieldTracker *YieldTrackerTransactorSession) Claim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.Claim(&_YieldTracker.TransactOpts, _account, _receiver)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address _distributor) returns()
func (_YieldTracker *YieldTrackerTransactor) SetDistributor(opts *bind.TransactOpts, _distributor common.Address) (*types.Transaction, error) {
	return _YieldTracker.contract.Transact(opts, "setDistributor", _distributor)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address _distributor) returns()
func (_YieldTracker *YieldTrackerSession) SetDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.SetDistributor(&_YieldTracker.TransactOpts, _distributor)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address _distributor) returns()
func (_YieldTracker *YieldTrackerTransactorSession) SetDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.SetDistributor(&_YieldTracker.TransactOpts, _distributor)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldTracker *YieldTrackerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _YieldTracker.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldTracker *YieldTrackerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.SetGov(&_YieldTracker.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldTracker *YieldTrackerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.SetGov(&_YieldTracker.TransactOpts, _gov)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x5fd61965.
//
// Solidity: function updateRewards(address _account) returns()
func (_YieldTracker *YieldTrackerTransactor) UpdateRewards(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldTracker.contract.Transact(opts, "updateRewards", _account)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x5fd61965.
//
// Solidity: function updateRewards(address _account) returns()
func (_YieldTracker *YieldTrackerSession) UpdateRewards(_account common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.UpdateRewards(&_YieldTracker.TransactOpts, _account)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x5fd61965.
//
// Solidity: function updateRewards(address _account) returns()
func (_YieldTracker *YieldTrackerTransactorSession) UpdateRewards(_account common.Address) (*types.Transaction, error) {
	return _YieldTracker.Contract.UpdateRewards(&_YieldTracker.TransactOpts, _account)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldTracker *YieldTrackerTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldTracker.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldTracker *YieldTrackerSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldTracker.Contract.WithdrawToken(&_YieldTracker.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldTracker *YieldTrackerTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldTracker.Contract.WithdrawToken(&_YieldTracker.TransactOpts, _token, _account, _amount)
}

// YieldTrackerClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the YieldTracker contract.
type YieldTrackerClaimIterator struct {
	Event *YieldTrackerClaim // Event containing the contract specifics and raw log

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
func (it *YieldTrackerClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldTrackerClaim)
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
		it.Event = new(YieldTrackerClaim)
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
func (it *YieldTrackerClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldTrackerClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldTrackerClaim represents a Claim event raised by the YieldTracker contract.
type YieldTrackerClaim struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_YieldTracker *YieldTrackerFilterer) FilterClaim(opts *bind.FilterOpts) (*YieldTrackerClaimIterator, error) {

	logs, sub, err := _YieldTracker.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &YieldTrackerClaimIterator{contract: _YieldTracker.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_YieldTracker *YieldTrackerFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *YieldTrackerClaim) (event.Subscription, error) {

	logs, sub, err := _YieldTracker.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldTrackerClaim)
				if err := _YieldTracker.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_YieldTracker *YieldTrackerFilterer) ParseClaim(log types.Log) (*YieldTrackerClaim, error) {
	event := new(YieldTrackerClaim)
	if err := _YieldTracker.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
