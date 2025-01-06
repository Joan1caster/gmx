// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipancakerouter

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

// IPancakeRouterMetaData contains all meta data concerning the IPancakeRouter contract.
var IPancakeRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPancakeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakeRouterMetaData.ABI instead.
var IPancakeRouterABI = IPancakeRouterMetaData.ABI

// IPancakeRouter is an auto generated Go binding around an Ethereum contract.
type IPancakeRouter struct {
	IPancakeRouterCaller     // Read-only binding to the contract
	IPancakeRouterTransactor // Write-only binding to the contract
	IPancakeRouterFilterer   // Log filterer for contract events
}

// IPancakeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeRouterSession struct {
	Contract     *IPancakeRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeRouterCallerSession struct {
	Contract *IPancakeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IPancakeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeRouterTransactorSession struct {
	Contract     *IPancakeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IPancakeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeRouterRaw struct {
	Contract *IPancakeRouter // Generic contract binding to access the raw methods on
}

// IPancakeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeRouterCallerRaw struct {
	Contract *IPancakeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IPancakeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeRouterTransactorRaw struct {
	Contract *IPancakeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeRouter creates a new instance of IPancakeRouter, bound to a specific deployed contract.
func NewIPancakeRouter(address common.Address, backend bind.ContractBackend) (*IPancakeRouter, error) {
	contract, err := bindIPancakeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter{IPancakeRouterCaller: IPancakeRouterCaller{contract: contract}, IPancakeRouterTransactor: IPancakeRouterTransactor{contract: contract}, IPancakeRouterFilterer: IPancakeRouterFilterer{contract: contract}}, nil
}

// NewIPancakeRouterCaller creates a new read-only instance of IPancakeRouter, bound to a specific deployed contract.
func NewIPancakeRouterCaller(address common.Address, caller bind.ContractCaller) (*IPancakeRouterCaller, error) {
	contract, err := bindIPancakeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouterCaller{contract: contract}, nil
}

// NewIPancakeRouterTransactor creates a new write-only instance of IPancakeRouter, bound to a specific deployed contract.
func NewIPancakeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeRouterTransactor, error) {
	contract, err := bindIPancakeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouterTransactor{contract: contract}, nil
}

// NewIPancakeRouterFilterer creates a new log filterer instance of IPancakeRouter, bound to a specific deployed contract.
func NewIPancakeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeRouterFilterer, error) {
	contract, err := bindIPancakeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouterFilterer{contract: contract}, nil
}

// bindIPancakeRouter binds a generic wrapper to an already deployed contract.
func bindIPancakeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPancakeRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter *IPancakeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter.Contract.IPancakeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter *IPancakeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter.Contract.IPancakeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter *IPancakeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter.Contract.IPancakeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter *IPancakeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter *IPancakeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter *IPancakeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter.Contract.contract.Transact(opts, method, params...)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter *IPancakeRouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter *IPancakeRouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter.Contract.AddLiquidity(&_IPancakeRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_IPancakeRouter *IPancakeRouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IPancakeRouter.Contract.AddLiquidity(&_IPancakeRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}
