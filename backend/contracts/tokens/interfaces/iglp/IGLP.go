// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iglp

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

// IGLPMetaData contains all meta data concerning the IGLP contract.
var IGLPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IGLPABI is the input ABI used to generate the binding from.
// Deprecated: Use IGLPMetaData.ABI instead.
var IGLPABI = IGLPMetaData.ABI

// IGLP is an auto generated Go binding around an Ethereum contract.
type IGLP struct {
	IGLPCaller     // Read-only binding to the contract
	IGLPTransactor // Write-only binding to the contract
	IGLPFilterer   // Log filterer for contract events
}

// IGLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGLPSession struct {
	Contract     *IGLP             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGLPCallerSession struct {
	Contract *IGLPCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGLPTransactorSession struct {
	Contract     *IGLPTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGLPRaw struct {
	Contract *IGLP // Generic contract binding to access the raw methods on
}

// IGLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGLPCallerRaw struct {
	Contract *IGLPCaller // Generic read-only contract binding to access the raw methods on
}

// IGLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGLPTransactorRaw struct {
	Contract *IGLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGLP creates a new instance of IGLP, bound to a specific deployed contract.
func NewIGLP(address common.Address, backend bind.ContractBackend) (*IGLP, error) {
	contract, err := bindIGLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGLP{IGLPCaller: IGLPCaller{contract: contract}, IGLPTransactor: IGLPTransactor{contract: contract}, IGLPFilterer: IGLPFilterer{contract: contract}}, nil
}

// NewIGLPCaller creates a new read-only instance of IGLP, bound to a specific deployed contract.
func NewIGLPCaller(address common.Address, caller bind.ContractCaller) (*IGLPCaller, error) {
	contract, err := bindIGLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGLPCaller{contract: contract}, nil
}

// NewIGLPTransactor creates a new write-only instance of IGLP, bound to a specific deployed contract.
func NewIGLPTransactor(address common.Address, transactor bind.ContractTransactor) (*IGLPTransactor, error) {
	contract, err := bindIGLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGLPTransactor{contract: contract}, nil
}

// NewIGLPFilterer creates a new log filterer instance of IGLP, bound to a specific deployed contract.
func NewIGLPFilterer(address common.Address, filterer bind.ContractFilterer) (*IGLPFilterer, error) {
	contract, err := bindIGLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGLPFilterer{contract: contract}, nil
}

// bindIGLP binds a generic wrapper to an already deployed contract.
func bindIGLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGLPMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGLP *IGLPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGLP.Contract.IGLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGLP *IGLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGLP.Contract.IGLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGLP *IGLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGLP.Contract.IGLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGLP *IGLPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGLP *IGLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGLP *IGLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGLP.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IGLP *IGLPTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGLP.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IGLP *IGLPSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGLP.Contract.Burn(&_IGLP.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_IGLP *IGLPTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGLP.Contract.Burn(&_IGLP.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IGLP *IGLPTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGLP.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IGLP *IGLPSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGLP.Contract.Mint(&_IGLP.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IGLP *IGLPTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IGLP.Contract.Mint(&_IGLP.TransactOpts, _account, _amount)
}
