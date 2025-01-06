// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipancakefactory

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

// IPancakeFactoryMetaData contains all meta data concerning the IPancakeFactory contract.
var IPancakeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPancakeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakeFactoryMetaData.ABI instead.
var IPancakeFactoryABI = IPancakeFactoryMetaData.ABI

// IPancakeFactory is an auto generated Go binding around an Ethereum contract.
type IPancakeFactory struct {
	IPancakeFactoryCaller     // Read-only binding to the contract
	IPancakeFactoryTransactor // Write-only binding to the contract
	IPancakeFactoryFilterer   // Log filterer for contract events
}

// IPancakeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeFactorySession struct {
	Contract     *IPancakeFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeFactoryCallerSession struct {
	Contract *IPancakeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPancakeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeFactoryTransactorSession struct {
	Contract     *IPancakeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPancakeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeFactoryRaw struct {
	Contract *IPancakeFactory // Generic contract binding to access the raw methods on
}

// IPancakeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeFactoryCallerRaw struct {
	Contract *IPancakeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IPancakeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeFactoryTransactorRaw struct {
	Contract *IPancakeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeFactory creates a new instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactory(address common.Address, backend bind.ContractBackend) (*IPancakeFactory, error) {
	contract, err := bindIPancakeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactory{IPancakeFactoryCaller: IPancakeFactoryCaller{contract: contract}, IPancakeFactoryTransactor: IPancakeFactoryTransactor{contract: contract}, IPancakeFactoryFilterer: IPancakeFactoryFilterer{contract: contract}}, nil
}

// NewIPancakeFactoryCaller creates a new read-only instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryCaller(address common.Address, caller bind.ContractCaller) (*IPancakeFactoryCaller, error) {
	contract, err := bindIPancakeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryCaller{contract: contract}, nil
}

// NewIPancakeFactoryTransactor creates a new write-only instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeFactoryTransactor, error) {
	contract, err := bindIPancakeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryTransactor{contract: contract}, nil
}

// NewIPancakeFactoryFilterer creates a new log filterer instance of IPancakeFactory, bound to a specific deployed contract.
func NewIPancakeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeFactoryFilterer, error) {
	contract, err := bindIPancakeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeFactoryFilterer{contract: contract}, nil
}

// bindIPancakeFactory binds a generic wrapper to an already deployed contract.
func bindIPancakeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPancakeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeFactory *IPancakeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeFactory.Contract.IPancakeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeFactory *IPancakeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.IPancakeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeFactory *IPancakeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.IPancakeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeFactory *IPancakeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeFactory *IPancakeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeFactory *IPancakeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _IPancakeFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPancakeFactory.Contract.GetPair(&_IPancakeFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IPancakeFactory *IPancakeFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IPancakeFactory.Contract.GetPair(&_IPancakeFactory.CallOpts, tokenA, tokenB)
}
