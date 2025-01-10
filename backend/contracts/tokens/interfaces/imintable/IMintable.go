// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package imintable

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

// IMintableMetaData contains all meta data concerning the IMintable contract.
var IMintableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMintableABI is the input ABI used to generate the binding from.
// Deprecated: Use IMintableMetaData.ABI instead.
var IMintableABI = IMintableMetaData.ABI

// IMintable is an auto generated Go binding around an Ethereum contract.
type IMintable struct {
	IMintableCaller     // Read-only binding to the contract
	IMintableTransactor // Write-only binding to the contract
	IMintableFilterer   // Log filterer for contract events
}

// IMintableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMintableSession struct {
	Contract     *IMintable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMintableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMintableCallerSession struct {
	Contract *IMintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IMintableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMintableTransactorSession struct {
	Contract     *IMintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMintableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMintableRaw struct {
	Contract *IMintable // Generic contract binding to access the raw methods on
}

// IMintableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMintableCallerRaw struct {
	Contract *IMintableCaller // Generic read-only contract binding to access the raw methods on
}

// IMintableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMintableTransactorRaw struct {
	Contract *IMintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMintable creates a new instance of IMintable, bound to a specific deployed contract.
func NewIMintable(address common.Address, backend bind.ContractBackend) (*IMintable, error) {
	contract, err := bindIMintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMintable{IMintableCaller: IMintableCaller{contract: contract}, IMintableTransactor: IMintableTransactor{contract: contract}, IMintableFilterer: IMintableFilterer{contract: contract}}, nil
}

// NewIMintableCaller creates a new read-only instance of IMintable, bound to a specific deployed contract.
func NewIMintableCaller(address common.Address, caller bind.ContractCaller) (*IMintableCaller, error) {
	contract, err := bindIMintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMintableCaller{contract: contract}, nil
}

// NewIMintableTransactor creates a new write-only instance of IMintable, bound to a specific deployed contract.
func NewIMintableTransactor(address common.Address, transactor bind.ContractTransactor) (*IMintableTransactor, error) {
	contract, err := bindIMintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMintableTransactor{contract: contract}, nil
}

// NewIMintableFilterer creates a new log filterer instance of IMintable, bound to a specific deployed contract.
func NewIMintableFilterer(address common.Address, filterer bind.ContractFilterer) (*IMintableFilterer, error) {
	contract, err := bindIMintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMintableFilterer{contract: contract}, nil
}

// bindIMintable binds a generic wrapper to an already deployed contract.
func bindIMintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMintableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintable *IMintableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintable.Contract.IMintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintable *IMintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintable.Contract.IMintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintable *IMintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintable.Contract.IMintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintable *IMintableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintable *IMintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintable *IMintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintable.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IMintable *IMintableTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IMintable.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IMintable *IMintableSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IMintable.Contract.Burn(&_IMintable.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IMintable *IMintableTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IMintable.Contract.Burn(&_IMintable.TransactOpts, _account, _amount)
}

// IsMinter is a paid mutator transaction binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _account) returns(bool)
func (_IMintable *IMintableTransactor) IsMinter(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IMintable.contract.Transact(opts, "isMinter", _account)
}

// IsMinter is a paid mutator transaction binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _account) returns(bool)
func (_IMintable *IMintableSession) IsMinter(_account common.Address) (*types.Transaction, error) {
	return _IMintable.Contract.IsMinter(&_IMintable.TransactOpts, _account)
}

// IsMinter is a paid mutator transaction binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _account) returns(bool)
func (_IMintable *IMintableTransactorSession) IsMinter(_account common.Address) (*types.Transaction, error) {
	return _IMintable.Contract.IsMinter(&_IMintable.TransactOpts, _account)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IMintable *IMintableTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IMintable.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IMintable *IMintableSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IMintable.Contract.Mint(&_IMintable.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IMintable *IMintableTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IMintable.Contract.Mint(&_IMintable.TransactOpts, _account, _amount)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_IMintable *IMintableTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _IMintable.contract.Transact(opts, "setMinter", _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_IMintable *IMintableSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _IMintable.Contract.SetMinter(&_IMintable.TransactOpts, _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_IMintable *IMintableTransactorSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _IMintable.Contract.SetMinter(&_IMintable.TransactOpts, _minter, _isActive)
}
