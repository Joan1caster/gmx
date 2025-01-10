// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmx

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

// GMXMetaData contains all meta data concerning the GMX contract.
var GMXMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addNonStakingAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateTransferMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonStakingAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonStakingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"recoverClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeNonStakingAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateTransferMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_yieldTrackers\",\"type\":\"address[]\"}],\"name\":\"setYieldTrackers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"yieldTrackers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GMXABI is the input ABI used to generate the binding from.
// Deprecated: Use GMXMetaData.ABI instead.
var GMXABI = GMXMetaData.ABI

// GMX is an auto generated Go binding around an Ethereum contract.
type GMX struct {
	GMXCaller     // Read-only binding to the contract
	GMXTransactor // Write-only binding to the contract
	GMXFilterer   // Log filterer for contract events
}

// GMXCaller is an auto generated read-only Go binding around an Ethereum contract.
type GMXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GMXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GMXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GMXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GMXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GMXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GMXSession struct {
	Contract     *GMX              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GMXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GMXCallerSession struct {
	Contract *GMXCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GMXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GMXTransactorSession struct {
	Contract     *GMXTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GMXRaw is an auto generated low-level Go binding around an Ethereum contract.
type GMXRaw struct {
	Contract *GMX // Generic contract binding to access the raw methods on
}

// GMXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GMXCallerRaw struct {
	Contract *GMXCaller // Generic read-only contract binding to access the raw methods on
}

// GMXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GMXTransactorRaw struct {
	Contract *GMXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGMX creates a new instance of GMX, bound to a specific deployed contract.
func NewGMX(address common.Address, backend bind.ContractBackend) (*GMX, error) {
	contract, err := bindGMX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GMX{GMXCaller: GMXCaller{contract: contract}, GMXTransactor: GMXTransactor{contract: contract}, GMXFilterer: GMXFilterer{contract: contract}}, nil
}

// NewGMXCaller creates a new read-only instance of GMX, bound to a specific deployed contract.
func NewGMXCaller(address common.Address, caller bind.ContractCaller) (*GMXCaller, error) {
	contract, err := bindGMX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GMXCaller{contract: contract}, nil
}

// NewGMXTransactor creates a new write-only instance of GMX, bound to a specific deployed contract.
func NewGMXTransactor(address common.Address, transactor bind.ContractTransactor) (*GMXTransactor, error) {
	contract, err := bindGMX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GMXTransactor{contract: contract}, nil
}

// NewGMXFilterer creates a new log filterer instance of GMX, bound to a specific deployed contract.
func NewGMXFilterer(address common.Address, filterer bind.ContractFilterer) (*GMXFilterer, error) {
	contract, err := bindGMX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GMXFilterer{contract: contract}, nil
}

// bindGMX binds a generic wrapper to an already deployed contract.
func bindGMX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GMXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GMX *GMXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GMX.Contract.GMXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GMX *GMXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GMX.Contract.GMXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GMX *GMXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GMX.Contract.GMXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GMX *GMXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GMX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GMX *GMXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GMX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GMX *GMXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GMX.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GMX *GMXCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GMX *GMXSession) Admins(arg0 common.Address) (bool, error) {
	return _GMX.Contract.Admins(&_GMX.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GMX *GMXCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _GMX.Contract.Admins(&_GMX.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GMX *GMXCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GMX *GMXSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GMX.Contract.Allowance(&_GMX.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GMX *GMXCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GMX.Contract.Allowance(&_GMX.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GMX *GMXCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GMX *GMXSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GMX.Contract.Allowances(&_GMX.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GMX *GMXCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GMX.Contract.Allowances(&_GMX.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GMX *GMXCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GMX *GMXSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GMX.Contract.BalanceOf(&_GMX.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GMX *GMXCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GMX.Contract.BalanceOf(&_GMX.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GMX *GMXCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GMX *GMXSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GMX.Contract.Balances(&_GMX.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GMX *GMXCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GMX.Contract.Balances(&_GMX.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GMX *GMXCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GMX *GMXSession) Decimals() (uint8, error) {
	return _GMX.Contract.Decimals(&_GMX.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GMX *GMXCallerSession) Decimals() (uint8, error) {
	return _GMX.Contract.Decimals(&_GMX.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GMX *GMXCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GMX *GMXSession) Gov() (common.Address, error) {
	return _GMX.Contract.Gov(&_GMX.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GMX *GMXCallerSession) Gov() (common.Address, error) {
	return _GMX.Contract.Gov(&_GMX.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() pure returns(string _name)
func (_GMX *GMXCaller) Id(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() pure returns(string _name)
func (_GMX *GMXSession) Id() (string, error) {
	return _GMX.Contract.Id(&_GMX.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() pure returns(string _name)
func (_GMX *GMXCallerSession) Id() (string, error) {
	return _GMX.Contract.Id(&_GMX.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_GMX *GMXCaller) InPrivateTransferMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "inPrivateTransferMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_GMX *GMXSession) InPrivateTransferMode() (bool, error) {
	return _GMX.Contract.InPrivateTransferMode(&_GMX.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_GMX *GMXCallerSession) InPrivateTransferMode() (bool, error) {
	return _GMX.Contract.InPrivateTransferMode(&_GMX.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GMX *GMXCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GMX *GMXSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GMX.Contract.IsHandler(&_GMX.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GMX *GMXCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GMX.Contract.IsHandler(&_GMX.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_GMX *GMXCaller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "isMinter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_GMX *GMXSession) IsMinter(arg0 common.Address) (bool, error) {
	return _GMX.Contract.IsMinter(&_GMX.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_GMX *GMXCallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _GMX.Contract.IsMinter(&_GMX.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GMX *GMXCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GMX *GMXSession) Name() (string, error) {
	return _GMX.Contract.Name(&_GMX.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GMX *GMXCallerSession) Name() (string, error) {
	return _GMX.Contract.Name(&_GMX.CallOpts)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_GMX *GMXCaller) NonStakingAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "nonStakingAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_GMX *GMXSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _GMX.Contract.NonStakingAccounts(&_GMX.CallOpts, arg0)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_GMX *GMXCallerSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _GMX.Contract.NonStakingAccounts(&_GMX.CallOpts, arg0)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_GMX *GMXCaller) NonStakingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "nonStakingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_GMX *GMXSession) NonStakingSupply() (*big.Int, error) {
	return _GMX.Contract.NonStakingSupply(&_GMX.CallOpts)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_GMX *GMXCallerSession) NonStakingSupply() (*big.Int, error) {
	return _GMX.Contract.NonStakingSupply(&_GMX.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_GMX *GMXCaller) StakedBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "stakedBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_GMX *GMXSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _GMX.Contract.StakedBalance(&_GMX.CallOpts, _account)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_GMX *GMXCallerSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _GMX.Contract.StakedBalance(&_GMX.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GMX *GMXCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GMX *GMXSession) Symbol() (string, error) {
	return _GMX.Contract.Symbol(&_GMX.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GMX *GMXCallerSession) Symbol() (string, error) {
	return _GMX.Contract.Symbol(&_GMX.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_GMX *GMXCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_GMX *GMXSession) TotalStaked() (*big.Int, error) {
	return _GMX.Contract.TotalStaked(&_GMX.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_GMX *GMXCallerSession) TotalStaked() (*big.Int, error) {
	return _GMX.Contract.TotalStaked(&_GMX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GMX *GMXCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GMX *GMXSession) TotalSupply() (*big.Int, error) {
	return _GMX.Contract.TotalSupply(&_GMX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GMX *GMXCallerSession) TotalSupply() (*big.Int, error) {
	return _GMX.Contract.TotalSupply(&_GMX.CallOpts)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_GMX *GMXCaller) YieldTrackers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GMX.contract.Call(opts, &out, "yieldTrackers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_GMX *GMXSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _GMX.Contract.YieldTrackers(&_GMX.CallOpts, arg0)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_GMX *GMXCallerSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _GMX.Contract.YieldTrackers(&_GMX.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_GMX *GMXTransactor) AddAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "addAdmin", _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_GMX *GMXSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.AddAdmin(&_GMX.TransactOpts, _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_GMX *GMXTransactorSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.AddAdmin(&_GMX.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_GMX *GMXTransactor) AddNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "addNonStakingAccount", _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_GMX *GMXSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.AddNonStakingAccount(&_GMX.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_GMX *GMXTransactorSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.AddNonStakingAccount(&_GMX.TransactOpts, _account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GMX *GMXTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GMX *GMXSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Approve(&_GMX.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GMX *GMXTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Approve(&_GMX.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_GMX *GMXTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_GMX *GMXSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Burn(&_GMX.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_GMX *GMXTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Burn(&_GMX.TransactOpts, _account, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_GMX *GMXTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_GMX *GMXSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _GMX.Contract.Claim(&_GMX.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_GMX *GMXTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _GMX.Contract.Claim(&_GMX.TransactOpts, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_GMX *GMXTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_GMX *GMXSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Mint(&_GMX.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_GMX *GMXTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Mint(&_GMX.TransactOpts, _account, _amount)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_GMX *GMXTransactor) RecoverClaim(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "recoverClaim", _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_GMX *GMXSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GMX.Contract.RecoverClaim(&_GMX.TransactOpts, _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_GMX *GMXTransactorSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GMX.Contract.RecoverClaim(&_GMX.TransactOpts, _account, _receiver)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_GMX *GMXTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_GMX *GMXSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.RemoveAdmin(&_GMX.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_GMX *GMXTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.RemoveAdmin(&_GMX.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_GMX *GMXTransactor) RemoveNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "removeNonStakingAccount", _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_GMX *GMXSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.RemoveNonStakingAccount(&_GMX.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_GMX *GMXTransactorSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _GMX.Contract.RemoveNonStakingAccount(&_GMX.TransactOpts, _account)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GMX *GMXTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GMX *GMXSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GMX.Contract.SetGov(&_GMX.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GMX *GMXTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GMX.Contract.SetGov(&_GMX.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GMX *GMXTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GMX *GMXSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GMX.Contract.SetHandler(&_GMX.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GMX *GMXTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GMX.Contract.SetHandler(&_GMX.TransactOpts, _handler, _isActive)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_GMX *GMXTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "setInPrivateTransferMode", _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_GMX *GMXSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GMX.Contract.SetInPrivateTransferMode(&_GMX.TransactOpts, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_GMX *GMXTransactorSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GMX.Contract.SetInPrivateTransferMode(&_GMX.TransactOpts, _inPrivateTransferMode)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_GMX *GMXTransactor) SetInfo(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "setInfo", _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_GMX *GMXSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _GMX.Contract.SetInfo(&_GMX.TransactOpts, _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_GMX *GMXTransactorSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _GMX.Contract.SetInfo(&_GMX.TransactOpts, _name, _symbol)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_GMX *GMXTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "setMinter", _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_GMX *GMXSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _GMX.Contract.SetMinter(&_GMX.TransactOpts, _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_GMX *GMXTransactorSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _GMX.Contract.SetMinter(&_GMX.TransactOpts, _minter, _isActive)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_GMX *GMXTransactor) SetYieldTrackers(opts *bind.TransactOpts, _yieldTrackers []common.Address) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "setYieldTrackers", _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_GMX *GMXSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _GMX.Contract.SetYieldTrackers(&_GMX.TransactOpts, _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_GMX *GMXTransactorSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _GMX.Contract.SetYieldTrackers(&_GMX.TransactOpts, _yieldTrackers)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GMX *GMXTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GMX *GMXSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Transfer(&_GMX.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GMX *GMXTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.Transfer(&_GMX.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GMX *GMXTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GMX *GMXSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.TransferFrom(&_GMX.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GMX *GMXTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.TransferFrom(&_GMX.TransactOpts, _sender, _recipient, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GMX *GMXTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GMX *GMXSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.WithdrawToken(&_GMX.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GMX *GMXTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMX.Contract.WithdrawToken(&_GMX.TransactOpts, _token, _account, _amount)
}

// GMXApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GMX contract.
type GMXApprovalIterator struct {
	Event *GMXApproval // Event containing the contract specifics and raw log

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
func (it *GMXApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GMXApproval)
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
		it.Event = new(GMXApproval)
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
func (it *GMXApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GMXApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GMXApproval represents a Approval event raised by the GMX contract.
type GMXApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GMX *GMXFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GMXApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GMX.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GMXApprovalIterator{contract: _GMX.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GMX *GMXFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GMXApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GMX.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GMXApproval)
				if err := _GMX.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GMX *GMXFilterer) ParseApproval(log types.Log) (*GMXApproval, error) {
	event := new(GMXApproval)
	if err := _GMX.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GMXTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GMX contract.
type GMXTransferIterator struct {
	Event *GMXTransfer // Event containing the contract specifics and raw log

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
func (it *GMXTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GMXTransfer)
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
		it.Event = new(GMXTransfer)
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
func (it *GMXTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GMXTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GMXTransfer represents a Transfer event raised by the GMX contract.
type GMXTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GMX *GMXFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GMXTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GMX.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GMXTransferIterator{contract: _GMX.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GMX *GMXFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GMXTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GMX.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GMXTransfer)
				if err := _GMX.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GMX *GMXFilterer) ParseTransfer(log types.Log) (*GMXTransfer, error) {
	event := new(GMXTransfer)
	if err := _GMX.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
