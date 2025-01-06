// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakerouter

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

// PancakeRouterMetaData contains all meta data concerning the PancakeRouter contract.
var PancakeRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PancakeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeRouterMetaData.ABI instead.
var PancakeRouterABI = PancakeRouterMetaData.ABI

// PancakeRouter is an auto generated Go binding around an Ethereum contract.
type PancakeRouter struct {
	PancakeRouterCaller     // Read-only binding to the contract
	PancakeRouterTransactor // Write-only binding to the contract
	PancakeRouterFilterer   // Log filterer for contract events
}

// PancakeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeRouterSession struct {
	Contract     *PancakeRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeRouterCallerSession struct {
	Contract *PancakeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PancakeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeRouterTransactorSession struct {
	Contract     *PancakeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PancakeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeRouterRaw struct {
	Contract *PancakeRouter // Generic contract binding to access the raw methods on
}

// PancakeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeRouterCallerRaw struct {
	Contract *PancakeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeRouterTransactorRaw struct {
	Contract *PancakeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeRouter creates a new instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouter(address common.Address, backend bind.ContractBackend) (*PancakeRouter, error) {
	contract, err := bindPancakeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeRouter{PancakeRouterCaller: PancakeRouterCaller{contract: contract}, PancakeRouterTransactor: PancakeRouterTransactor{contract: contract}, PancakeRouterFilterer: PancakeRouterFilterer{contract: contract}}, nil
}

// NewPancakeRouterCaller creates a new read-only instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouterCaller(address common.Address, caller bind.ContractCaller) (*PancakeRouterCaller, error) {
	contract, err := bindPancakeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeRouterCaller{contract: contract}, nil
}

// NewPancakeRouterTransactor creates a new write-only instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeRouterTransactor, error) {
	contract, err := bindPancakeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeRouterTransactor{contract: contract}, nil
}

// NewPancakeRouterFilterer creates a new log filterer instance of PancakeRouter, bound to a specific deployed contract.
func NewPancakeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeRouterFilterer, error) {
	contract, err := bindPancakeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeRouterFilterer{contract: contract}, nil
}

// bindPancakeRouter binds a generic wrapper to an already deployed contract.
func bindPancakeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeRouter *PancakeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeRouter.Contract.PancakeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeRouter *PancakeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeRouter.Contract.PancakeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeRouter *PancakeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeRouter.Contract.PancakeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeRouter *PancakeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeRouter *PancakeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeRouter *PancakeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeRouter.Contract.contract.Transact(opts, method, params...)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_PancakeRouter *PancakeRouterCaller) Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeRouter.contract.Call(opts, &out, "pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_PancakeRouter *PancakeRouterSession) Pair() (common.Address, error) {
	return _PancakeRouter.Contract.Pair(&_PancakeRouter.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_PancakeRouter *PancakeRouterCallerSession) Pair() (common.Address, error) {
	return _PancakeRouter.Contract.Pair(&_PancakeRouter.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 , uint256 , address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_PancakeRouter *PancakeRouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, arg4 *big.Int, arg5 *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, arg4, arg5, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 , uint256 , address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_PancakeRouter *PancakeRouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, arg4 *big.Int, arg5 *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.AddLiquidity(&_PancakeRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, arg4, arg5, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 , uint256 , address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_PancakeRouter *PancakeRouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, arg4 *big.Int, arg5 *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _PancakeRouter.Contract.AddLiquidity(&_PancakeRouter.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, arg4, arg5, to, deadline)
}
