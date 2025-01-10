// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iusdg

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

// IUSDGMetaData contains all meta data concerning the IUSDG contract.
var IUSDGMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"addVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"removeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUSDGABI is the input ABI used to generate the binding from.
// Deprecated: Use IUSDGMetaData.ABI instead.
var IUSDGABI = IUSDGMetaData.ABI

// IUSDG is an auto generated Go binding around an Ethereum contract.
type IUSDG struct {
	IUSDGCaller     // Read-only binding to the contract
	IUSDGTransactor // Write-only binding to the contract
	IUSDGFilterer   // Log filterer for contract events
}

// IUSDGCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUSDGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSDGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUSDGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSDGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUSDGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUSDGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUSDGSession struct {
	Contract     *IUSDG            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUSDGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUSDGCallerSession struct {
	Contract *IUSDGCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IUSDGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUSDGTransactorSession struct {
	Contract     *IUSDGTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUSDGRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUSDGRaw struct {
	Contract *IUSDG // Generic contract binding to access the raw methods on
}

// IUSDGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUSDGCallerRaw struct {
	Contract *IUSDGCaller // Generic read-only contract binding to access the raw methods on
}

// IUSDGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUSDGTransactorRaw struct {
	Contract *IUSDGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUSDG creates a new instance of IUSDG, bound to a specific deployed contract.
func NewIUSDG(address common.Address, backend bind.ContractBackend) (*IUSDG, error) {
	contract, err := bindIUSDG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUSDG{IUSDGCaller: IUSDGCaller{contract: contract}, IUSDGTransactor: IUSDGTransactor{contract: contract}, IUSDGFilterer: IUSDGFilterer{contract: contract}}, nil
}

// NewIUSDGCaller creates a new read-only instance of IUSDG, bound to a specific deployed contract.
func NewIUSDGCaller(address common.Address, caller bind.ContractCaller) (*IUSDGCaller, error) {
	contract, err := bindIUSDG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUSDGCaller{contract: contract}, nil
}

// NewIUSDGTransactor creates a new write-only instance of IUSDG, bound to a specific deployed contract.
func NewIUSDGTransactor(address common.Address, transactor bind.ContractTransactor) (*IUSDGTransactor, error) {
	contract, err := bindIUSDG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUSDGTransactor{contract: contract}, nil
}

// NewIUSDGFilterer creates a new log filterer instance of IUSDG, bound to a specific deployed contract.
func NewIUSDGFilterer(address common.Address, filterer bind.ContractFilterer) (*IUSDGFilterer, error) {
	contract, err := bindIUSDG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUSDGFilterer{contract: contract}, nil
}

// bindIUSDG binds a generic wrapper to an already deployed contract.
func bindIUSDG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUSDGMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUSDG *IUSDGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUSDG.Contract.IUSDGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUSDG *IUSDGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUSDG.Contract.IUSDGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUSDG *IUSDGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUSDG.Contract.IUSDGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUSDG *IUSDGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUSDG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUSDG *IUSDGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUSDG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUSDG *IUSDGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUSDG.Contract.contract.Transact(opts, method, params...)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_IUSDG *IUSDGTransactor) AddVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "addVault", _vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_IUSDG *IUSDGSession) AddVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.AddVault(&_IUSDG.TransactOpts, _vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_IUSDG *IUSDGTransactorSession) AddVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.AddVault(&_IUSDG.TransactOpts, _vault)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Burn(&_IUSDG.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Burn(&_IUSDG.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Mint(&_IUSDG.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IUSDG *IUSDGTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IUSDG.Contract.Mint(&_IUSDG.TransactOpts, _account, _amount)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_IUSDG *IUSDGTransactor) RemoveVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _IUSDG.contract.Transact(opts, "removeVault", _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_IUSDG *IUSDGSession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.RemoveVault(&_IUSDG.TransactOpts, _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_IUSDG *IUSDGTransactorSession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _IUSDG.Contract.RemoveVault(&_IUSDG.TransactOpts, _vault)
}
