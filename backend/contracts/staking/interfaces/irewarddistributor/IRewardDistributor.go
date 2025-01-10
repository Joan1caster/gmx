// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irewarddistributor

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

// IRewardDistributorMetaData contains all meta data concerning the IRewardDistributor contract.
var IRewardDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IRewardDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardDistributorMetaData.ABI instead.
var IRewardDistributorABI = IRewardDistributorMetaData.ABI

// IRewardDistributor is an auto generated Go binding around an Ethereum contract.
type IRewardDistributor struct {
	IRewardDistributorCaller     // Read-only binding to the contract
	IRewardDistributorTransactor // Write-only binding to the contract
	IRewardDistributorFilterer   // Log filterer for contract events
}

// IRewardDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardDistributorSession struct {
	Contract     *IRewardDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IRewardDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardDistributorCallerSession struct {
	Contract *IRewardDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IRewardDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardDistributorTransactorSession struct {
	Contract     *IRewardDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IRewardDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardDistributorRaw struct {
	Contract *IRewardDistributor // Generic contract binding to access the raw methods on
}

// IRewardDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardDistributorCallerRaw struct {
	Contract *IRewardDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardDistributorTransactorRaw struct {
	Contract *IRewardDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardDistributor creates a new instance of IRewardDistributor, bound to a specific deployed contract.
func NewIRewardDistributor(address common.Address, backend bind.ContractBackend) (*IRewardDistributor, error) {
	contract, err := bindIRewardDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewardDistributor{IRewardDistributorCaller: IRewardDistributorCaller{contract: contract}, IRewardDistributorTransactor: IRewardDistributorTransactor{contract: contract}, IRewardDistributorFilterer: IRewardDistributorFilterer{contract: contract}}, nil
}

// NewIRewardDistributorCaller creates a new read-only instance of IRewardDistributor, bound to a specific deployed contract.
func NewIRewardDistributorCaller(address common.Address, caller bind.ContractCaller) (*IRewardDistributorCaller, error) {
	contract, err := bindIRewardDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardDistributorCaller{contract: contract}, nil
}

// NewIRewardDistributorTransactor creates a new write-only instance of IRewardDistributor, bound to a specific deployed contract.
func NewIRewardDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardDistributorTransactor, error) {
	contract, err := bindIRewardDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardDistributorTransactor{contract: contract}, nil
}

// NewIRewardDistributorFilterer creates a new log filterer instance of IRewardDistributor, bound to a specific deployed contract.
func NewIRewardDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardDistributorFilterer, error) {
	contract, err := bindIRewardDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardDistributorFilterer{contract: contract}, nil
}

// bindIRewardDistributor binds a generic wrapper to an already deployed contract.
func bindIRewardDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRewardDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardDistributor *IRewardDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardDistributor.Contract.IRewardDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardDistributor *IRewardDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardDistributor.Contract.IRewardDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardDistributor *IRewardDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardDistributor.Contract.IRewardDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardDistributor *IRewardDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardDistributor *IRewardDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardDistributor *IRewardDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardDistributor.Contract.contract.Transact(opts, method, params...)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_IRewardDistributor *IRewardDistributorCaller) PendingRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRewardDistributor.contract.Call(opts, &out, "pendingRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_IRewardDistributor *IRewardDistributorSession) PendingRewards() (*big.Int, error) {
	return _IRewardDistributor.Contract.PendingRewards(&_IRewardDistributor.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_IRewardDistributor *IRewardDistributorCallerSession) PendingRewards() (*big.Int, error) {
	return _IRewardDistributor.Contract.PendingRewards(&_IRewardDistributor.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_IRewardDistributor *IRewardDistributorCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRewardDistributor.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_IRewardDistributor *IRewardDistributorSession) RewardToken() (common.Address, error) {
	return _IRewardDistributor.Contract.RewardToken(&_IRewardDistributor.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_IRewardDistributor *IRewardDistributorCallerSession) RewardToken() (common.Address, error) {
	return _IRewardDistributor.Contract.RewardToken(&_IRewardDistributor.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_IRewardDistributor *IRewardDistributorCaller) TokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRewardDistributor.contract.Call(opts, &out, "tokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_IRewardDistributor *IRewardDistributorSession) TokensPerInterval() (*big.Int, error) {
	return _IRewardDistributor.Contract.TokensPerInterval(&_IRewardDistributor.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_IRewardDistributor *IRewardDistributorCallerSession) TokensPerInterval() (*big.Int, error) {
	return _IRewardDistributor.Contract.TokensPerInterval(&_IRewardDistributor.CallOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_IRewardDistributor *IRewardDistributorTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardDistributor.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_IRewardDistributor *IRewardDistributorSession) Distribute() (*types.Transaction, error) {
	return _IRewardDistributor.Contract.Distribute(&_IRewardDistributor.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_IRewardDistributor *IRewardDistributorTransactorSession) Distribute() (*types.Transaction, error) {
	return _IRewardDistributor.Contract.Distribute(&_IRewardDistributor.TransactOpts)
}
