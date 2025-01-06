// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faucettoken

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

// FaucetTokenMetaData contains all meta data concerning the FaucetToken contract.
var FaucetTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dropletAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DROPLET_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_claimedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_dropletAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_isFaucetEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDroplet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableFaucet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableFaucet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dropletAmount\",\"type\":\"uint256\"}],\"name\":\"setDropletAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FaucetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use FaucetTokenMetaData.ABI instead.
var FaucetTokenABI = FaucetTokenMetaData.ABI

// FaucetToken is an auto generated Go binding around an Ethereum contract.
type FaucetToken struct {
	FaucetTokenCaller     // Read-only binding to the contract
	FaucetTokenTransactor // Write-only binding to the contract
	FaucetTokenFilterer   // Log filterer for contract events
}

// FaucetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetTokenSession struct {
	Contract     *FaucetToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetTokenCallerSession struct {
	Contract *FaucetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FaucetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetTokenTransactorSession struct {
	Contract     *FaucetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FaucetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetTokenRaw struct {
	Contract *FaucetToken // Generic contract binding to access the raw methods on
}

// FaucetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetTokenCallerRaw struct {
	Contract *FaucetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetTokenTransactorRaw struct {
	Contract *FaucetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucetToken creates a new instance of FaucetToken, bound to a specific deployed contract.
func NewFaucetToken(address common.Address, backend bind.ContractBackend) (*FaucetToken, error) {
	contract, err := bindFaucetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FaucetToken{FaucetTokenCaller: FaucetTokenCaller{contract: contract}, FaucetTokenTransactor: FaucetTokenTransactor{contract: contract}, FaucetTokenFilterer: FaucetTokenFilterer{contract: contract}}, nil
}

// NewFaucetTokenCaller creates a new read-only instance of FaucetToken, bound to a specific deployed contract.
func NewFaucetTokenCaller(address common.Address, caller bind.ContractCaller) (*FaucetTokenCaller, error) {
	contract, err := bindFaucetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTokenCaller{contract: contract}, nil
}

// NewFaucetTokenTransactor creates a new write-only instance of FaucetToken, bound to a specific deployed contract.
func NewFaucetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetTokenTransactor, error) {
	contract, err := bindFaucetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTokenTransactor{contract: contract}, nil
}

// NewFaucetTokenFilterer creates a new log filterer instance of FaucetToken, bound to a specific deployed contract.
func NewFaucetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetTokenFilterer, error) {
	contract, err := bindFaucetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetTokenFilterer{contract: contract}, nil
}

// bindFaucetToken binds a generic wrapper to an already deployed contract.
func bindFaucetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FaucetTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaucetToken *FaucetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaucetToken.Contract.FaucetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaucetToken *FaucetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetToken.Contract.FaucetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaucetToken *FaucetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaucetToken.Contract.FaucetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaucetToken *FaucetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaucetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaucetToken *FaucetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaucetToken *FaucetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaucetToken.Contract.contract.Transact(opts, method, params...)
}

// DROPLETINTERVAL is a free data retrieval call binding the contract method 0xbd136448.
//
// Solidity: function DROPLET_INTERVAL() view returns(uint256)
func (_FaucetToken *FaucetTokenCaller) DROPLETINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "DROPLET_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DROPLETINTERVAL is a free data retrieval call binding the contract method 0xbd136448.
//
// Solidity: function DROPLET_INTERVAL() view returns(uint256)
func (_FaucetToken *FaucetTokenSession) DROPLETINTERVAL() (*big.Int, error) {
	return _FaucetToken.Contract.DROPLETINTERVAL(&_FaucetToken.CallOpts)
}

// DROPLETINTERVAL is a free data retrieval call binding the contract method 0xbd136448.
//
// Solidity: function DROPLET_INTERVAL() view returns(uint256)
func (_FaucetToken *FaucetTokenCallerSession) DROPLETINTERVAL() (*big.Int, error) {
	return _FaucetToken.Contract.DROPLETINTERVAL(&_FaucetToken.CallOpts)
}

// ClaimedAt is a free data retrieval call binding the contract method 0x309ef252.
//
// Solidity: function _claimedAt(address ) view returns(uint256)
func (_FaucetToken *FaucetTokenCaller) ClaimedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "_claimedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedAt is a free data retrieval call binding the contract method 0x309ef252.
//
// Solidity: function _claimedAt(address ) view returns(uint256)
func (_FaucetToken *FaucetTokenSession) ClaimedAt(arg0 common.Address) (*big.Int, error) {
	return _FaucetToken.Contract.ClaimedAt(&_FaucetToken.CallOpts, arg0)
}

// ClaimedAt is a free data retrieval call binding the contract method 0x309ef252.
//
// Solidity: function _claimedAt(address ) view returns(uint256)
func (_FaucetToken *FaucetTokenCallerSession) ClaimedAt(arg0 common.Address) (*big.Int, error) {
	return _FaucetToken.Contract.ClaimedAt(&_FaucetToken.CallOpts, arg0)
}

// DropletAmount is a free data retrieval call binding the contract method 0x458a7c36.
//
// Solidity: function _dropletAmount() view returns(uint256)
func (_FaucetToken *FaucetTokenCaller) DropletAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "_dropletAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DropletAmount is a free data retrieval call binding the contract method 0x458a7c36.
//
// Solidity: function _dropletAmount() view returns(uint256)
func (_FaucetToken *FaucetTokenSession) DropletAmount() (*big.Int, error) {
	return _FaucetToken.Contract.DropletAmount(&_FaucetToken.CallOpts)
}

// DropletAmount is a free data retrieval call binding the contract method 0x458a7c36.
//
// Solidity: function _dropletAmount() view returns(uint256)
func (_FaucetToken *FaucetTokenCallerSession) DropletAmount() (*big.Int, error) {
	return _FaucetToken.Contract.DropletAmount(&_FaucetToken.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x63878f16.
//
// Solidity: function _gov() view returns(address)
func (_FaucetToken *FaucetTokenCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "_gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x63878f16.
//
// Solidity: function _gov() view returns(address)
func (_FaucetToken *FaucetTokenSession) Gov() (common.Address, error) {
	return _FaucetToken.Contract.Gov(&_FaucetToken.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x63878f16.
//
// Solidity: function _gov() view returns(address)
func (_FaucetToken *FaucetTokenCallerSession) Gov() (common.Address, error) {
	return _FaucetToken.Contract.Gov(&_FaucetToken.CallOpts)
}

// IsFaucetEnabled is a free data retrieval call binding the contract method 0xa1c2fb0c.
//
// Solidity: function _isFaucetEnabled() view returns(bool)
func (_FaucetToken *FaucetTokenCaller) IsFaucetEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "_isFaucetEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFaucetEnabled is a free data retrieval call binding the contract method 0xa1c2fb0c.
//
// Solidity: function _isFaucetEnabled() view returns(bool)
func (_FaucetToken *FaucetTokenSession) IsFaucetEnabled() (bool, error) {
	return _FaucetToken.Contract.IsFaucetEnabled(&_FaucetToken.CallOpts)
}

// IsFaucetEnabled is a free data retrieval call binding the contract method 0xa1c2fb0c.
//
// Solidity: function _isFaucetEnabled() view returns(bool)
func (_FaucetToken *FaucetTokenCallerSession) IsFaucetEnabled() (bool, error) {
	return _FaucetToken.Contract.IsFaucetEnabled(&_FaucetToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FaucetToken *FaucetTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FaucetToken *FaucetTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FaucetToken.Contract.Allowance(&_FaucetToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FaucetToken *FaucetTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FaucetToken.Contract.Allowance(&_FaucetToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FaucetToken *FaucetTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FaucetToken *FaucetTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FaucetToken.Contract.BalanceOf(&_FaucetToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FaucetToken *FaucetTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FaucetToken.Contract.BalanceOf(&_FaucetToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FaucetToken *FaucetTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FaucetToken *FaucetTokenSession) Decimals() (uint8, error) {
	return _FaucetToken.Contract.Decimals(&_FaucetToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FaucetToken *FaucetTokenCallerSession) Decimals() (uint8, error) {
	return _FaucetToken.Contract.Decimals(&_FaucetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FaucetToken *FaucetTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FaucetToken *FaucetTokenSession) Name() (string, error) {
	return _FaucetToken.Contract.Name(&_FaucetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FaucetToken *FaucetTokenCallerSession) Name() (string, error) {
	return _FaucetToken.Contract.Name(&_FaucetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FaucetToken *FaucetTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FaucetToken *FaucetTokenSession) Symbol() (string, error) {
	return _FaucetToken.Contract.Symbol(&_FaucetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FaucetToken *FaucetTokenCallerSession) Symbol() (string, error) {
	return _FaucetToken.Contract.Symbol(&_FaucetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FaucetToken *FaucetTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaucetToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FaucetToken *FaucetTokenSession) TotalSupply() (*big.Int, error) {
	return _FaucetToken.Contract.TotalSupply(&_FaucetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FaucetToken *FaucetTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FaucetToken.Contract.TotalSupply(&_FaucetToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.Approve(&_FaucetToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.Approve(&_FaucetToken.TransactOpts, spender, amount)
}

// ClaimDroplet is a paid mutator transaction binding the contract method 0xe6511743.
//
// Solidity: function claimDroplet() returns()
func (_FaucetToken *FaucetTokenTransactor) ClaimDroplet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "claimDroplet")
}

// ClaimDroplet is a paid mutator transaction binding the contract method 0xe6511743.
//
// Solidity: function claimDroplet() returns()
func (_FaucetToken *FaucetTokenSession) ClaimDroplet() (*types.Transaction, error) {
	return _FaucetToken.Contract.ClaimDroplet(&_FaucetToken.TransactOpts)
}

// ClaimDroplet is a paid mutator transaction binding the contract method 0xe6511743.
//
// Solidity: function claimDroplet() returns()
func (_FaucetToken *FaucetTokenTransactorSession) ClaimDroplet() (*types.Transaction, error) {
	return _FaucetToken.Contract.ClaimDroplet(&_FaucetToken.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaucetToken *FaucetTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaucetToken *FaucetTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.DecreaseAllowance(&_FaucetToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaucetToken *FaucetTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.DecreaseAllowance(&_FaucetToken.TransactOpts, spender, subtractedValue)
}

// DisableFaucet is a paid mutator transaction binding the contract method 0x944d3c2a.
//
// Solidity: function disableFaucet() returns()
func (_FaucetToken *FaucetTokenTransactor) DisableFaucet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "disableFaucet")
}

// DisableFaucet is a paid mutator transaction binding the contract method 0x944d3c2a.
//
// Solidity: function disableFaucet() returns()
func (_FaucetToken *FaucetTokenSession) DisableFaucet() (*types.Transaction, error) {
	return _FaucetToken.Contract.DisableFaucet(&_FaucetToken.TransactOpts)
}

// DisableFaucet is a paid mutator transaction binding the contract method 0x944d3c2a.
//
// Solidity: function disableFaucet() returns()
func (_FaucetToken *FaucetTokenTransactorSession) DisableFaucet() (*types.Transaction, error) {
	return _FaucetToken.Contract.DisableFaucet(&_FaucetToken.TransactOpts)
}

// EnableFaucet is a paid mutator transaction binding the contract method 0xc7c7dee5.
//
// Solidity: function enableFaucet() returns()
func (_FaucetToken *FaucetTokenTransactor) EnableFaucet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "enableFaucet")
}

// EnableFaucet is a paid mutator transaction binding the contract method 0xc7c7dee5.
//
// Solidity: function enableFaucet() returns()
func (_FaucetToken *FaucetTokenSession) EnableFaucet() (*types.Transaction, error) {
	return _FaucetToken.Contract.EnableFaucet(&_FaucetToken.TransactOpts)
}

// EnableFaucet is a paid mutator transaction binding the contract method 0xc7c7dee5.
//
// Solidity: function enableFaucet() returns()
func (_FaucetToken *FaucetTokenTransactorSession) EnableFaucet() (*types.Transaction, error) {
	return _FaucetToken.Contract.EnableFaucet(&_FaucetToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaucetToken *FaucetTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaucetToken *FaucetTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.IncreaseAllowance(&_FaucetToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaucetToken *FaucetTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.IncreaseAllowance(&_FaucetToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_FaucetToken *FaucetTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_FaucetToken *FaucetTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.Mint(&_FaucetToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_FaucetToken *FaucetTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.Mint(&_FaucetToken.TransactOpts, account, amount)
}

// SetDropletAmount is a paid mutator transaction binding the contract method 0xd08f8b3f.
//
// Solidity: function setDropletAmount(uint256 dropletAmount) returns()
func (_FaucetToken *FaucetTokenTransactor) SetDropletAmount(opts *bind.TransactOpts, dropletAmount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "setDropletAmount", dropletAmount)
}

// SetDropletAmount is a paid mutator transaction binding the contract method 0xd08f8b3f.
//
// Solidity: function setDropletAmount(uint256 dropletAmount) returns()
func (_FaucetToken *FaucetTokenSession) SetDropletAmount(dropletAmount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.SetDropletAmount(&_FaucetToken.TransactOpts, dropletAmount)
}

// SetDropletAmount is a paid mutator transaction binding the contract method 0xd08f8b3f.
//
// Solidity: function setDropletAmount(uint256 dropletAmount) returns()
func (_FaucetToken *FaucetTokenTransactorSession) SetDropletAmount(dropletAmount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.SetDropletAmount(&_FaucetToken.TransactOpts, dropletAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.Transfer(&_FaucetToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.Transfer(&_FaucetToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.TransferFrom(&_FaucetToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaucetToken *FaucetTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaucetToken.Contract.TransferFrom(&_FaucetToken.TransactOpts, sender, recipient, amount)
}

// FaucetTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FaucetToken contract.
type FaucetTokenApprovalIterator struct {
	Event *FaucetTokenApproval // Event containing the contract specifics and raw log

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
func (it *FaucetTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetTokenApproval)
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
		it.Event = new(FaucetTokenApproval)
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
func (it *FaucetTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetTokenApproval represents a Approval event raised by the FaucetToken contract.
type FaucetTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FaucetToken *FaucetTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FaucetTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FaucetToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FaucetTokenApprovalIterator{contract: _FaucetToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FaucetToken *FaucetTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FaucetTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FaucetToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetTokenApproval)
				if err := _FaucetToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FaucetToken *FaucetTokenFilterer) ParseApproval(log types.Log) (*FaucetTokenApproval, error) {
	event := new(FaucetTokenApproval)
	if err := _FaucetToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FaucetTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FaucetToken contract.
type FaucetTokenTransferIterator struct {
	Event *FaucetTokenTransfer // Event containing the contract specifics and raw log

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
func (it *FaucetTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetTokenTransfer)
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
		it.Event = new(FaucetTokenTransfer)
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
func (it *FaucetTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetTokenTransfer represents a Transfer event raised by the FaucetToken contract.
type FaucetTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FaucetToken *FaucetTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FaucetTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FaucetToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FaucetTokenTransferIterator{contract: _FaucetToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FaucetToken *FaucetTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FaucetTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FaucetToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetTokenTransfer)
				if err := _FaucetToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FaucetToken *FaucetTokenFilterer) ParseTransfer(log types.Log) (*FaucetTokenTransfer, error) {
	event := new(FaucetTokenTransfer)
	if err := _FaucetToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
