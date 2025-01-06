// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibasepositionmanager

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

// IBasePositionManagerMetaData contains all meta data concerning the IBasePositionManager contract.
var IBasePositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxGlobalLongSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxGlobalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBasePositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IBasePositionManagerMetaData.ABI instead.
var IBasePositionManagerABI = IBasePositionManagerMetaData.ABI

// IBasePositionManager is an auto generated Go binding around an Ethereum contract.
type IBasePositionManager struct {
	IBasePositionManagerCaller     // Read-only binding to the contract
	IBasePositionManagerTransactor // Write-only binding to the contract
	IBasePositionManagerFilterer   // Log filterer for contract events
}

// IBasePositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBasePositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBasePositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBasePositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBasePositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBasePositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBasePositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBasePositionManagerSession struct {
	Contract     *IBasePositionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IBasePositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBasePositionManagerCallerSession struct {
	Contract *IBasePositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IBasePositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBasePositionManagerTransactorSession struct {
	Contract     *IBasePositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IBasePositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBasePositionManagerRaw struct {
	Contract *IBasePositionManager // Generic contract binding to access the raw methods on
}

// IBasePositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBasePositionManagerCallerRaw struct {
	Contract *IBasePositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IBasePositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBasePositionManagerTransactorRaw struct {
	Contract *IBasePositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBasePositionManager creates a new instance of IBasePositionManager, bound to a specific deployed contract.
func NewIBasePositionManager(address common.Address, backend bind.ContractBackend) (*IBasePositionManager, error) {
	contract, err := bindIBasePositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBasePositionManager{IBasePositionManagerCaller: IBasePositionManagerCaller{contract: contract}, IBasePositionManagerTransactor: IBasePositionManagerTransactor{contract: contract}, IBasePositionManagerFilterer: IBasePositionManagerFilterer{contract: contract}}, nil
}

// NewIBasePositionManagerCaller creates a new read-only instance of IBasePositionManager, bound to a specific deployed contract.
func NewIBasePositionManagerCaller(address common.Address, caller bind.ContractCaller) (*IBasePositionManagerCaller, error) {
	contract, err := bindIBasePositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBasePositionManagerCaller{contract: contract}, nil
}

// NewIBasePositionManagerTransactor creates a new write-only instance of IBasePositionManager, bound to a specific deployed contract.
func NewIBasePositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IBasePositionManagerTransactor, error) {
	contract, err := bindIBasePositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBasePositionManagerTransactor{contract: contract}, nil
}

// NewIBasePositionManagerFilterer creates a new log filterer instance of IBasePositionManager, bound to a specific deployed contract.
func NewIBasePositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IBasePositionManagerFilterer, error) {
	contract, err := bindIBasePositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBasePositionManagerFilterer{contract: contract}, nil
}

// bindIBasePositionManager binds a generic wrapper to an already deployed contract.
func bindIBasePositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBasePositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBasePositionManager *IBasePositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBasePositionManager.Contract.IBasePositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBasePositionManager *IBasePositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBasePositionManager.Contract.IBasePositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBasePositionManager *IBasePositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBasePositionManager.Contract.IBasePositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBasePositionManager *IBasePositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBasePositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBasePositionManager *IBasePositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBasePositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBasePositionManager *IBasePositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBasePositionManager.Contract.contract.Transact(opts, method, params...)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address _token) view returns(uint256)
func (_IBasePositionManager *IBasePositionManagerCaller) MaxGlobalLongSizes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBasePositionManager.contract.Call(opts, &out, "maxGlobalLongSizes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address _token) view returns(uint256)
func (_IBasePositionManager *IBasePositionManagerSession) MaxGlobalLongSizes(_token common.Address) (*big.Int, error) {
	return _IBasePositionManager.Contract.MaxGlobalLongSizes(&_IBasePositionManager.CallOpts, _token)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address _token) view returns(uint256)
func (_IBasePositionManager *IBasePositionManagerCallerSession) MaxGlobalLongSizes(_token common.Address) (*big.Int, error) {
	return _IBasePositionManager.Contract.MaxGlobalLongSizes(&_IBasePositionManager.CallOpts, _token)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IBasePositionManager *IBasePositionManagerCaller) MaxGlobalShortSizes(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBasePositionManager.contract.Call(opts, &out, "maxGlobalShortSizes", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IBasePositionManager *IBasePositionManagerSession) MaxGlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IBasePositionManager.Contract.MaxGlobalShortSizes(&_IBasePositionManager.CallOpts, _token)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address _token) view returns(uint256)
func (_IBasePositionManager *IBasePositionManagerCallerSession) MaxGlobalShortSizes(_token common.Address) (*big.Int, error) {
	return _IBasePositionManager.Contract.MaxGlobalShortSizes(&_IBasePositionManager.CallOpts, _token)
}
