// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iammrouter

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

// IAmmRouterMetaData contains all meta data concerning the IAmmRouter contract.
var IAmmRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAmmRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IAmmRouterMetaData.ABI instead.
var IAmmRouterABI = IAmmRouterMetaData.ABI

// IAmmRouter is an auto generated Go binding around an Ethereum contract.
type IAmmRouter struct {
	IAmmRouterCaller     // Read-only binding to the contract
	IAmmRouterTransactor // Write-only binding to the contract
	IAmmRouterFilterer   // Log filterer for contract events
}

// IAmmRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAmmRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAmmRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAmmRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAmmRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAmmRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAmmRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAmmRouterSession struct {
	Contract     *IAmmRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAmmRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAmmRouterCallerSession struct {
	Contract *IAmmRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IAmmRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAmmRouterTransactorSession struct {
	Contract     *IAmmRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IAmmRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAmmRouterRaw struct {
	Contract *IAmmRouter // Generic contract binding to access the raw methods on
}

// IAmmRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAmmRouterCallerRaw struct {
	Contract *IAmmRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IAmmRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAmmRouterTransactorRaw struct {
	Contract *IAmmRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAmmRouter creates a new instance of IAmmRouter, bound to a specific deployed contract.
func NewIAmmRouter(address common.Address, backend bind.ContractBackend) (*IAmmRouter, error) {
	contract, err := bindIAmmRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAmmRouter{IAmmRouterCaller: IAmmRouterCaller{contract: contract}, IAmmRouterTransactor: IAmmRouterTransactor{contract: contract}, IAmmRouterFilterer: IAmmRouterFilterer{contract: contract}}, nil
}

// NewIAmmRouterCaller creates a new read-only instance of IAmmRouter, bound to a specific deployed contract.
func NewIAmmRouterCaller(address common.Address, caller bind.ContractCaller) (*IAmmRouterCaller, error) {
	contract, err := bindIAmmRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAmmRouterCaller{contract: contract}, nil
}

// NewIAmmRouterTransactor creates a new write-only instance of IAmmRouter, bound to a specific deployed contract.
func NewIAmmRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IAmmRouterTransactor, error) {
	contract, err := bindIAmmRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAmmRouterTransactor{contract: contract}, nil
}

// NewIAmmRouterFilterer creates a new log filterer instance of IAmmRouter, bound to a specific deployed contract.
func NewIAmmRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IAmmRouterFilterer, error) {
	contract, err := bindIAmmRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAmmRouterFilterer{contract: contract}, nil
}

// bindIAmmRouter binds a generic wrapper to an already deployed contract.
func bindIAmmRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAmmRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAmmRouter *IAmmRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAmmRouter.Contract.IAmmRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAmmRouter *IAmmRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAmmRouter.Contract.IAmmRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAmmRouter *IAmmRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAmmRouter.Contract.IAmmRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAmmRouter *IAmmRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAmmRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAmmRouter *IAmmRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAmmRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAmmRouter *IAmmRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAmmRouter.Contract.contract.Transact(opts, method, params...)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IAmmRouter *IAmmRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAmmRouter.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IAmmRouter *IAmmRouterSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAmmRouter.Contract.RemoveLiquidity(&_IAmmRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_IAmmRouter *IAmmRouterTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAmmRouter.Contract.RemoveLiquidity(&_IAmmRouter.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAmmRouter *IAmmRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAmmRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAmmRouter *IAmmRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAmmRouter.Contract.SwapExactTokensForTokens(&_IAmmRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_IAmmRouter *IAmmRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _IAmmRouter.Contract.SwapExactTokensForTokens(&_IAmmRouter.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}
