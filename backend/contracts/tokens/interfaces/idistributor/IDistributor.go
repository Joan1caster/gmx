// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package idistributor

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

// IDistributorMetaData contains all meta data concerning the IDistributor contract.
var IDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getDistributionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getRewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"tokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use IDistributorMetaData.ABI instead.
var IDistributorABI = IDistributorMetaData.ABI

// IDistributor is an auto generated Go binding around an Ethereum contract.
type IDistributor struct {
	IDistributorCaller     // Read-only binding to the contract
	IDistributorTransactor // Write-only binding to the contract
	IDistributorFilterer   // Log filterer for contract events
}

// IDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDistributorSession struct {
	Contract     *IDistributor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDistributorCallerSession struct {
	Contract *IDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDistributorTransactorSession struct {
	Contract     *IDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDistributorRaw struct {
	Contract *IDistributor // Generic contract binding to access the raw methods on
}

// IDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDistributorCallerRaw struct {
	Contract *IDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// IDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDistributorTransactorRaw struct {
	Contract *IDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDistributor creates a new instance of IDistributor, bound to a specific deployed contract.
func NewIDistributor(address common.Address, backend bind.ContractBackend) (*IDistributor, error) {
	contract, err := bindIDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDistributor{IDistributorCaller: IDistributorCaller{contract: contract}, IDistributorTransactor: IDistributorTransactor{contract: contract}, IDistributorFilterer: IDistributorFilterer{contract: contract}}, nil
}

// NewIDistributorCaller creates a new read-only instance of IDistributor, bound to a specific deployed contract.
func NewIDistributorCaller(address common.Address, caller bind.ContractCaller) (*IDistributorCaller, error) {
	contract, err := bindIDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDistributorCaller{contract: contract}, nil
}

// NewIDistributorTransactor creates a new write-only instance of IDistributor, bound to a specific deployed contract.
func NewIDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*IDistributorTransactor, error) {
	contract, err := bindIDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDistributorTransactor{contract: contract}, nil
}

// NewIDistributorFilterer creates a new log filterer instance of IDistributor, bound to a specific deployed contract.
func NewIDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*IDistributorFilterer, error) {
	contract, err := bindIDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDistributorFilterer{contract: contract}, nil
}

// bindIDistributor binds a generic wrapper to an already deployed contract.
func bindIDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDistributor *IDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDistributor.Contract.IDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDistributor *IDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistributor.Contract.IDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDistributor *IDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDistributor.Contract.IDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDistributor *IDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDistributor *IDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDistributor *IDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDistributor.Contract.contract.Transact(opts, method, params...)
}

// GetDistributionAmount is a free data retrieval call binding the contract method 0xb7fba9f6.
//
// Solidity: function getDistributionAmount(address _receiver) view returns(uint256)
func (_IDistributor *IDistributorCaller) GetDistributionAmount(opts *bind.CallOpts, _receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDistributor.contract.Call(opts, &out, "getDistributionAmount", _receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionAmount is a free data retrieval call binding the contract method 0xb7fba9f6.
//
// Solidity: function getDistributionAmount(address _receiver) view returns(uint256)
func (_IDistributor *IDistributorSession) GetDistributionAmount(_receiver common.Address) (*big.Int, error) {
	return _IDistributor.Contract.GetDistributionAmount(&_IDistributor.CallOpts, _receiver)
}

// GetDistributionAmount is a free data retrieval call binding the contract method 0xb7fba9f6.
//
// Solidity: function getDistributionAmount(address _receiver) view returns(uint256)
func (_IDistributor *IDistributorCallerSession) GetDistributionAmount(_receiver common.Address) (*big.Int, error) {
	return _IDistributor.Contract.GetDistributionAmount(&_IDistributor.CallOpts, _receiver)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address _receiver) view returns(address)
func (_IDistributor *IDistributorCaller) GetRewardToken(opts *bind.CallOpts, _receiver common.Address) (common.Address, error) {
	var out []interface{}
	err := _IDistributor.contract.Call(opts, &out, "getRewardToken", _receiver)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address _receiver) view returns(address)
func (_IDistributor *IDistributorSession) GetRewardToken(_receiver common.Address) (common.Address, error) {
	return _IDistributor.Contract.GetRewardToken(&_IDistributor.CallOpts, _receiver)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address _receiver) view returns(address)
func (_IDistributor *IDistributorCallerSession) GetRewardToken(_receiver common.Address) (common.Address, error) {
	return _IDistributor.Contract.GetRewardToken(&_IDistributor.CallOpts, _receiver)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0x270da240.
//
// Solidity: function tokensPerInterval(address _receiver) view returns(uint256)
func (_IDistributor *IDistributorCaller) TokensPerInterval(opts *bind.CallOpts, _receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IDistributor.contract.Call(opts, &out, "tokensPerInterval", _receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0x270da240.
//
// Solidity: function tokensPerInterval(address _receiver) view returns(uint256)
func (_IDistributor *IDistributorSession) TokensPerInterval(_receiver common.Address) (*big.Int, error) {
	return _IDistributor.Contract.TokensPerInterval(&_IDistributor.CallOpts, _receiver)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0x270da240.
//
// Solidity: function tokensPerInterval(address _receiver) view returns(uint256)
func (_IDistributor *IDistributorCallerSession) TokensPerInterval(_receiver common.Address) (*big.Int, error) {
	return _IDistributor.Contract.TokensPerInterval(&_IDistributor.CallOpts, _receiver)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_IDistributor *IDistributorTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistributor.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_IDistributor *IDistributorSession) Distribute() (*types.Transaction, error) {
	return _IDistributor.Contract.Distribute(&_IDistributor.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_IDistributor *IDistributorTransactorSession) Distribute() (*types.Transaction, error) {
	return _IDistributor.Contract.Distribute(&_IDistributor.TransactOpts)
}
