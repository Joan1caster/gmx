// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewarddistributor

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

// RewardDistributorMetaData contains all meta data concerning the RewardDistributor contract.
var RewardDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardTracker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Distribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPerIntervalChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDistributionTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTokensPerInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateLastDistributionTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RewardDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardDistributorMetaData.ABI instead.
var RewardDistributorABI = RewardDistributorMetaData.ABI

// RewardDistributor is an auto generated Go binding around an Ethereum contract.
type RewardDistributor struct {
	RewardDistributorCaller     // Read-only binding to the contract
	RewardDistributorTransactor // Write-only binding to the contract
	RewardDistributorFilterer   // Log filterer for contract events
}

// RewardDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardDistributorSession struct {
	Contract     *RewardDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardDistributorCallerSession struct {
	Contract *RewardDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RewardDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardDistributorTransactorSession struct {
	Contract     *RewardDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RewardDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardDistributorRaw struct {
	Contract *RewardDistributor // Generic contract binding to access the raw methods on
}

// RewardDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardDistributorCallerRaw struct {
	Contract *RewardDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// RewardDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardDistributorTransactorRaw struct {
	Contract *RewardDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardDistributor creates a new instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributor(address common.Address, backend bind.ContractBackend) (*RewardDistributor, error) {
	contract, err := bindRewardDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardDistributor{RewardDistributorCaller: RewardDistributorCaller{contract: contract}, RewardDistributorTransactor: RewardDistributorTransactor{contract: contract}, RewardDistributorFilterer: RewardDistributorFilterer{contract: contract}}, nil
}

// NewRewardDistributorCaller creates a new read-only instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributorCaller(address common.Address, caller bind.ContractCaller) (*RewardDistributorCaller, error) {
	contract, err := bindRewardDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributorCaller{contract: contract}, nil
}

// NewRewardDistributorTransactor creates a new write-only instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardDistributorTransactor, error) {
	contract, err := bindRewardDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributorTransactor{contract: contract}, nil
}

// NewRewardDistributorFilterer creates a new log filterer instance of RewardDistributor, bound to a specific deployed contract.
func NewRewardDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardDistributorFilterer, error) {
	contract, err := bindRewardDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardDistributorFilterer{contract: contract}, nil
}

// bindRewardDistributor binds a generic wrapper to an already deployed contract.
func bindRewardDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributor *RewardDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributor.Contract.RewardDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributor *RewardDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributor.Contract.RewardDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributor *RewardDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributor.Contract.RewardDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributor *RewardDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributor *RewardDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributor *RewardDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributor.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_RewardDistributor *RewardDistributorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_RewardDistributor *RewardDistributorSession) Admin() (common.Address, error) {
	return _RewardDistributor.Contract.Admin(&_RewardDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_RewardDistributor *RewardDistributorCallerSession) Admin() (common.Address, error) {
	return _RewardDistributor.Contract.Admin(&_RewardDistributor.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardDistributor *RewardDistributorCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardDistributor *RewardDistributorSession) Gov() (common.Address, error) {
	return _RewardDistributor.Contract.Gov(&_RewardDistributor.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardDistributor *RewardDistributorCallerSession) Gov() (common.Address, error) {
	return _RewardDistributor.Contract.Gov(&_RewardDistributor.CallOpts)
}

// LastDistributionTime is a free data retrieval call binding the contract method 0x75b17350.
//
// Solidity: function lastDistributionTime() view returns(uint256)
func (_RewardDistributor *RewardDistributorCaller) LastDistributionTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "lastDistributionTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDistributionTime is a free data retrieval call binding the contract method 0x75b17350.
//
// Solidity: function lastDistributionTime() view returns(uint256)
func (_RewardDistributor *RewardDistributorSession) LastDistributionTime() (*big.Int, error) {
	return _RewardDistributor.Contract.LastDistributionTime(&_RewardDistributor.CallOpts)
}

// LastDistributionTime is a free data retrieval call binding the contract method 0x75b17350.
//
// Solidity: function lastDistributionTime() view returns(uint256)
func (_RewardDistributor *RewardDistributorCallerSession) LastDistributionTime() (*big.Int, error) {
	return _RewardDistributor.Contract.LastDistributionTime(&_RewardDistributor.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_RewardDistributor *RewardDistributorCaller) PendingRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "pendingRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_RewardDistributor *RewardDistributorSession) PendingRewards() (*big.Int, error) {
	return _RewardDistributor.Contract.PendingRewards(&_RewardDistributor.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_RewardDistributor *RewardDistributorCallerSession) PendingRewards() (*big.Int, error) {
	return _RewardDistributor.Contract.PendingRewards(&_RewardDistributor.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardDistributor *RewardDistributorCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardDistributor *RewardDistributorSession) RewardToken() (common.Address, error) {
	return _RewardDistributor.Contract.RewardToken(&_RewardDistributor.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardDistributor *RewardDistributorCallerSession) RewardToken() (common.Address, error) {
	return _RewardDistributor.Contract.RewardToken(&_RewardDistributor.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_RewardDistributor *RewardDistributorCaller) RewardTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "rewardTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_RewardDistributor *RewardDistributorSession) RewardTracker() (common.Address, error) {
	return _RewardDistributor.Contract.RewardTracker(&_RewardDistributor.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_RewardDistributor *RewardDistributorCallerSession) RewardTracker() (common.Address, error) {
	return _RewardDistributor.Contract.RewardTracker(&_RewardDistributor.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_RewardDistributor *RewardDistributorCaller) TokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributor.contract.Call(opts, &out, "tokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_RewardDistributor *RewardDistributorSession) TokensPerInterval() (*big.Int, error) {
	return _RewardDistributor.Contract.TokensPerInterval(&_RewardDistributor.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_RewardDistributor *RewardDistributorCallerSession) TokensPerInterval() (*big.Int, error) {
	return _RewardDistributor.Contract.TokensPerInterval(&_RewardDistributor.CallOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_RewardDistributor *RewardDistributorTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_RewardDistributor *RewardDistributorSession) Distribute() (*types.Transaction, error) {
	return _RewardDistributor.Contract.Distribute(&_RewardDistributor.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_RewardDistributor *RewardDistributorTransactorSession) Distribute() (*types.Transaction, error) {
	return _RewardDistributor.Contract.Distribute(&_RewardDistributor.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_RewardDistributor *RewardDistributorTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_RewardDistributor *RewardDistributorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetAdmin(&_RewardDistributor.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_RewardDistributor *RewardDistributorTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetAdmin(&_RewardDistributor.TransactOpts, _admin)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardDistributor *RewardDistributorTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardDistributor *RewardDistributorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetGov(&_RewardDistributor.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardDistributor *RewardDistributorTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetGov(&_RewardDistributor.TransactOpts, _gov)
}

// SetTokensPerInterval is a paid mutator transaction binding the contract method 0x18e20a03.
//
// Solidity: function setTokensPerInterval(uint256 _amount) returns()
func (_RewardDistributor *RewardDistributorTransactor) SetTokensPerInterval(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "setTokensPerInterval", _amount)
}

// SetTokensPerInterval is a paid mutator transaction binding the contract method 0x18e20a03.
//
// Solidity: function setTokensPerInterval(uint256 _amount) returns()
func (_RewardDistributor *RewardDistributorSession) SetTokensPerInterval(_amount *big.Int) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetTokensPerInterval(&_RewardDistributor.TransactOpts, _amount)
}

// SetTokensPerInterval is a paid mutator transaction binding the contract method 0x18e20a03.
//
// Solidity: function setTokensPerInterval(uint256 _amount) returns()
func (_RewardDistributor *RewardDistributorTransactorSession) SetTokensPerInterval(_amount *big.Int) (*types.Transaction, error) {
	return _RewardDistributor.Contract.SetTokensPerInterval(&_RewardDistributor.TransactOpts, _amount)
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x3ae6d6eb.
//
// Solidity: function updateLastDistributionTime() returns()
func (_RewardDistributor *RewardDistributorTransactor) UpdateLastDistributionTime(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "updateLastDistributionTime")
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x3ae6d6eb.
//
// Solidity: function updateLastDistributionTime() returns()
func (_RewardDistributor *RewardDistributorSession) UpdateLastDistributionTime() (*types.Transaction, error) {
	return _RewardDistributor.Contract.UpdateLastDistributionTime(&_RewardDistributor.TransactOpts)
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x3ae6d6eb.
//
// Solidity: function updateLastDistributionTime() returns()
func (_RewardDistributor *RewardDistributorTransactorSession) UpdateLastDistributionTime() (*types.Transaction, error) {
	return _RewardDistributor.Contract.UpdateLastDistributionTime(&_RewardDistributor.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardDistributor *RewardDistributorTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardDistributor.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardDistributor *RewardDistributorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardDistributor.Contract.WithdrawToken(&_RewardDistributor.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardDistributor *RewardDistributorTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardDistributor.Contract.WithdrawToken(&_RewardDistributor.TransactOpts, _token, _account, _amount)
}

// RewardDistributorDistributeIterator is returned from FilterDistribute and is used to iterate over the raw logs and unpacked data for Distribute events raised by the RewardDistributor contract.
type RewardDistributorDistributeIterator struct {
	Event *RewardDistributorDistribute // Event containing the contract specifics and raw log

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
func (it *RewardDistributorDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDistributorDistribute)
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
		it.Event = new(RewardDistributorDistribute)
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
func (it *RewardDistributorDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDistributorDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDistributorDistribute represents a Distribute event raised by the RewardDistributor contract.
type RewardDistributorDistribute struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistribute is a free log retrieval operation binding the contract event 0x4def474aca53bf221d07d9ab0f675b3f6d8d2494b8427271bcf43c018ef1eead.
//
// Solidity: event Distribute(uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) FilterDistribute(opts *bind.FilterOpts) (*RewardDistributorDistributeIterator, error) {

	logs, sub, err := _RewardDistributor.contract.FilterLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return &RewardDistributorDistributeIterator{contract: _RewardDistributor.contract, event: "Distribute", logs: logs, sub: sub}, nil
}

// WatchDistribute is a free log subscription operation binding the contract event 0x4def474aca53bf221d07d9ab0f675b3f6d8d2494b8427271bcf43c018ef1eead.
//
// Solidity: event Distribute(uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) WatchDistribute(opts *bind.WatchOpts, sink chan<- *RewardDistributorDistribute) (event.Subscription, error) {

	logs, sub, err := _RewardDistributor.contract.WatchLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDistributorDistribute)
				if err := _RewardDistributor.contract.UnpackLog(event, "Distribute", log); err != nil {
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

// ParseDistribute is a log parse operation binding the contract event 0x4def474aca53bf221d07d9ab0f675b3f6d8d2494b8427271bcf43c018ef1eead.
//
// Solidity: event Distribute(uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) ParseDistribute(log types.Log) (*RewardDistributorDistribute, error) {
	event := new(RewardDistributorDistribute)
	if err := _RewardDistributor.contract.UnpackLog(event, "Distribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardDistributorTokensPerIntervalChangeIterator is returned from FilterTokensPerIntervalChange and is used to iterate over the raw logs and unpacked data for TokensPerIntervalChange events raised by the RewardDistributor contract.
type RewardDistributorTokensPerIntervalChangeIterator struct {
	Event *RewardDistributorTokensPerIntervalChange // Event containing the contract specifics and raw log

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
func (it *RewardDistributorTokensPerIntervalChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDistributorTokensPerIntervalChange)
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
		it.Event = new(RewardDistributorTokensPerIntervalChange)
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
func (it *RewardDistributorTokensPerIntervalChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDistributorTokensPerIntervalChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDistributorTokensPerIntervalChange represents a TokensPerIntervalChange event raised by the RewardDistributor contract.
type RewardDistributorTokensPerIntervalChange struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensPerIntervalChange is a free log retrieval operation binding the contract event 0x98dc76c39aa5a5dcb749f8750a65db3dfa1e14bcc1591a9c16a7420e5da748f8.
//
// Solidity: event TokensPerIntervalChange(uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) FilterTokensPerIntervalChange(opts *bind.FilterOpts) (*RewardDistributorTokensPerIntervalChangeIterator, error) {

	logs, sub, err := _RewardDistributor.contract.FilterLogs(opts, "TokensPerIntervalChange")
	if err != nil {
		return nil, err
	}
	return &RewardDistributorTokensPerIntervalChangeIterator{contract: _RewardDistributor.contract, event: "TokensPerIntervalChange", logs: logs, sub: sub}, nil
}

// WatchTokensPerIntervalChange is a free log subscription operation binding the contract event 0x98dc76c39aa5a5dcb749f8750a65db3dfa1e14bcc1591a9c16a7420e5da748f8.
//
// Solidity: event TokensPerIntervalChange(uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) WatchTokensPerIntervalChange(opts *bind.WatchOpts, sink chan<- *RewardDistributorTokensPerIntervalChange) (event.Subscription, error) {

	logs, sub, err := _RewardDistributor.contract.WatchLogs(opts, "TokensPerIntervalChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDistributorTokensPerIntervalChange)
				if err := _RewardDistributor.contract.UnpackLog(event, "TokensPerIntervalChange", log); err != nil {
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

// ParseTokensPerIntervalChange is a log parse operation binding the contract event 0x98dc76c39aa5a5dcb749f8750a65db3dfa1e14bcc1591a9c16a7420e5da748f8.
//
// Solidity: event TokensPerIntervalChange(uint256 amount)
func (_RewardDistributor *RewardDistributorFilterer) ParseTokensPerIntervalChange(log types.Log) (*RewardDistributorTokensPerIntervalChange, error) {
	event := new(RewardDistributorTokensPerIntervalChange)
	if err := _RewardDistributor.contract.UnpackLog(event, "TokensPerIntervalChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
