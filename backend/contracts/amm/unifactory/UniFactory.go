// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unifactory

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

// UniFactoryMetaData contains all meta data concerning the UniFactory contract.
var UniFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniFactoryMetaData.ABI instead.
var UniFactoryABI = UniFactoryMetaData.ABI

// UniFactory is an auto generated Go binding around an Ethereum contract.
type UniFactory struct {
	UniFactoryCaller     // Read-only binding to the contract
	UniFactoryTransactor // Write-only binding to the contract
	UniFactoryFilterer   // Log filterer for contract events
}

// UniFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniFactorySession struct {
	Contract     *UniFactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniFactoryCallerSession struct {
	Contract *UniFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UniFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniFactoryTransactorSession struct {
	Contract     *UniFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UniFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniFactoryRaw struct {
	Contract *UniFactory // Generic contract binding to access the raw methods on
}

// UniFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniFactoryCallerRaw struct {
	Contract *UniFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UniFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniFactoryTransactorRaw struct {
	Contract *UniFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniFactory creates a new instance of UniFactory, bound to a specific deployed contract.
func NewUniFactory(address common.Address, backend bind.ContractBackend) (*UniFactory, error) {
	contract, err := bindUniFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniFactory{UniFactoryCaller: UniFactoryCaller{contract: contract}, UniFactoryTransactor: UniFactoryTransactor{contract: contract}, UniFactoryFilterer: UniFactoryFilterer{contract: contract}}, nil
}

// NewUniFactoryCaller creates a new read-only instance of UniFactory, bound to a specific deployed contract.
func NewUniFactoryCaller(address common.Address, caller bind.ContractCaller) (*UniFactoryCaller, error) {
	contract, err := bindUniFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniFactoryCaller{contract: contract}, nil
}

// NewUniFactoryTransactor creates a new write-only instance of UniFactory, bound to a specific deployed contract.
func NewUniFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniFactoryTransactor, error) {
	contract, err := bindUniFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniFactoryTransactor{contract: contract}, nil
}

// NewUniFactoryFilterer creates a new log filterer instance of UniFactory, bound to a specific deployed contract.
func NewUniFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniFactoryFilterer, error) {
	contract, err := bindUniFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniFactoryFilterer{contract: contract}, nil
}

// bindUniFactory binds a generic wrapper to an already deployed contract.
func bindUniFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniFactory *UniFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniFactory.Contract.UniFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniFactory *UniFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFactory.Contract.UniFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniFactory *UniFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniFactory.Contract.UniFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniFactory *UniFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniFactory *UniFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniFactory *UniFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniFactory *UniFactoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniFactory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniFactory *UniFactorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniFactory.Contract.GetPool(&_UniFactory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniFactory *UniFactoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniFactory.Contract.GetPool(&_UniFactory.CallOpts, arg0, arg1, arg2)
}
