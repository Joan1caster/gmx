// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irewardtracker

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

// IRewardTrackerMetaData contains all meta data concerning the IRewardTracker contract.
var IRewardTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"averageStakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"cumulativeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"}],\"name\":\"depositBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fundingAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"stakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRewardTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardTrackerMetaData.ABI instead.
var IRewardTrackerABI = IRewardTrackerMetaData.ABI

// IRewardTracker is an auto generated Go binding around an Ethereum contract.
type IRewardTracker struct {
	IRewardTrackerCaller     // Read-only binding to the contract
	IRewardTrackerTransactor // Write-only binding to the contract
	IRewardTrackerFilterer   // Log filterer for contract events
}

// IRewardTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardTrackerSession struct {
	Contract     *IRewardTracker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardTrackerCallerSession struct {
	Contract *IRewardTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IRewardTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardTrackerTransactorSession struct {
	Contract     *IRewardTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IRewardTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardTrackerRaw struct {
	Contract *IRewardTracker // Generic contract binding to access the raw methods on
}

// IRewardTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardTrackerCallerRaw struct {
	Contract *IRewardTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardTrackerTransactorRaw struct {
	Contract *IRewardTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardTracker creates a new instance of IRewardTracker, bound to a specific deployed contract.
func NewIRewardTracker(address common.Address, backend bind.ContractBackend) (*IRewardTracker, error) {
	contract, err := bindIRewardTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewardTracker{IRewardTrackerCaller: IRewardTrackerCaller{contract: contract}, IRewardTrackerTransactor: IRewardTrackerTransactor{contract: contract}, IRewardTrackerFilterer: IRewardTrackerFilterer{contract: contract}}, nil
}

// NewIRewardTrackerCaller creates a new read-only instance of IRewardTracker, bound to a specific deployed contract.
func NewIRewardTrackerCaller(address common.Address, caller bind.ContractCaller) (*IRewardTrackerCaller, error) {
	contract, err := bindIRewardTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardTrackerCaller{contract: contract}, nil
}

// NewIRewardTrackerTransactor creates a new write-only instance of IRewardTracker, bound to a specific deployed contract.
func NewIRewardTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardTrackerTransactor, error) {
	contract, err := bindIRewardTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardTrackerTransactor{contract: contract}, nil
}

// NewIRewardTrackerFilterer creates a new log filterer instance of IRewardTracker, bound to a specific deployed contract.
func NewIRewardTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardTrackerFilterer, error) {
	contract, err := bindIRewardTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardTrackerFilterer{contract: contract}, nil
}

// bindIRewardTracker binds a generic wrapper to an already deployed contract.
func bindIRewardTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRewardTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardTracker *IRewardTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardTracker.Contract.IRewardTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardTracker *IRewardTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardTracker.Contract.IRewardTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardTracker *IRewardTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardTracker.Contract.IRewardTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardTracker *IRewardTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardTracker *IRewardTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardTracker *IRewardTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardTracker.Contract.contract.Transact(opts, method, params...)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCaller) AverageStakedAmounts(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRewardTracker.contract.Call(opts, &out, "averageStakedAmounts", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) AverageStakedAmounts(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.AverageStakedAmounts(&_IRewardTracker.CallOpts, _account)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCallerSession) AverageStakedAmounts(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.AverageStakedAmounts(&_IRewardTracker.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRewardTracker.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.Claimable(&_IRewardTracker.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.Claimable(&_IRewardTracker.CallOpts, _account)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCaller) CumulativeRewards(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRewardTracker.contract.Call(opts, &out, "cumulativeRewards", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) CumulativeRewards(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.CumulativeRewards(&_IRewardTracker.CallOpts, _account)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCallerSession) CumulativeRewards(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.CumulativeRewards(&_IRewardTracker.CallOpts, _account)
}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address _account, address _depositToken) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCaller) DepositBalances(opts *bind.CallOpts, _account common.Address, _depositToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRewardTracker.contract.Call(opts, &out, "depositBalances", _account, _depositToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address _account, address _depositToken) view returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) DepositBalances(_account common.Address, _depositToken common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.DepositBalances(&_IRewardTracker.CallOpts, _account, _depositToken)
}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address _account, address _depositToken) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCallerSession) DepositBalances(_account common.Address, _depositToken common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.DepositBalances(&_IRewardTracker.CallOpts, _account, _depositToken)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCaller) StakedAmounts(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IRewardTracker.contract.Call(opts, &out, "stakedAmounts", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) StakedAmounts(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.StakedAmounts(&_IRewardTracker.CallOpts, _account)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address _account) view returns(uint256)
func (_IRewardTracker *IRewardTrackerCallerSession) StakedAmounts(_account common.Address) (*big.Int, error) {
	return _IRewardTracker.Contract.StakedAmounts(&_IRewardTracker.CallOpts, _account)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_IRewardTracker *IRewardTrackerCaller) TokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRewardTracker.contract.Call(opts, &out, "tokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) TokensPerInterval() (*big.Int, error) {
	return _IRewardTracker.Contract.TokensPerInterval(&_IRewardTracker.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_IRewardTracker *IRewardTrackerCallerSession) TokensPerInterval() (*big.Int, error) {
	return _IRewardTracker.Contract.TokensPerInterval(&_IRewardTracker.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_IRewardTracker *IRewardTrackerTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.Contract.Claim(&_IRewardTracker.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_IRewardTracker *IRewardTrackerTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.Contract.Claim(&_IRewardTracker.TransactOpts, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_IRewardTracker *IRewardTrackerTransactor) ClaimForAccount(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.contract.Transact(opts, "claimForAccount", _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_IRewardTracker *IRewardTrackerSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.Contract.ClaimForAccount(&_IRewardTracker.TransactOpts, _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_IRewardTracker *IRewardTrackerTransactorSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.Contract.ClaimForAccount(&_IRewardTracker.TransactOpts, _account, _receiver)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerTransactor) Stake(opts *bind.TransactOpts, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.contract.Transact(opts, "stake", _depositToken, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerSession) Stake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.Contract.Stake(&_IRewardTracker.TransactOpts, _depositToken, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerTransactorSession) Stake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.Contract.Stake(&_IRewardTracker.TransactOpts, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerTransactor) StakeForAccount(opts *bind.TransactOpts, _fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.contract.Transact(opts, "stakeForAccount", _fundingAccount, _account, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerSession) StakeForAccount(_fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.Contract.StakeForAccount(&_IRewardTracker.TransactOpts, _fundingAccount, _account, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerTransactorSession) StakeForAccount(_fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.Contract.StakeForAccount(&_IRewardTracker.TransactOpts, _fundingAccount, _account, _depositToken, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerTransactor) Unstake(opts *bind.TransactOpts, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.contract.Transact(opts, "unstake", _depositToken, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerSession) Unstake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.Contract.Unstake(&_IRewardTracker.TransactOpts, _depositToken, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_IRewardTracker *IRewardTrackerTransactorSession) Unstake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardTracker.Contract.Unstake(&_IRewardTracker.TransactOpts, _depositToken, _amount)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_IRewardTracker *IRewardTrackerTransactor) UnstakeForAccount(opts *bind.TransactOpts, _account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.contract.Transact(opts, "unstakeForAccount", _account, _depositToken, _amount, _receiver)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_IRewardTracker *IRewardTrackerSession) UnstakeForAccount(_account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.Contract.UnstakeForAccount(&_IRewardTracker.TransactOpts, _account, _depositToken, _amount, _receiver)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_IRewardTracker *IRewardTrackerTransactorSession) UnstakeForAccount(_account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IRewardTracker.Contract.UnstakeForAccount(&_IRewardTracker.TransactOpts, _account, _depositToken, _amount, _receiver)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_IRewardTracker *IRewardTrackerTransactor) UpdateRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardTracker.contract.Transact(opts, "updateRewards")
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_IRewardTracker *IRewardTrackerSession) UpdateRewards() (*types.Transaction, error) {
	return _IRewardTracker.Contract.UpdateRewards(&_IRewardTracker.TransactOpts)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_IRewardTracker *IRewardTrackerTransactorSession) UpdateRewards() (*types.Transaction, error) {
	return _IRewardTracker.Contract.UpdateRewards(&_IRewardTracker.TransactOpts)
}
