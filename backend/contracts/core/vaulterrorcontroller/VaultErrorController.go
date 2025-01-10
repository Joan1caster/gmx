// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaulterrorcontroller

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

// VaultErrorControllerMetaData contains all meta data concerning the VaultErrorController contract.
var VaultErrorControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_errors\",\"type\":\"string[]\"}],\"name\":\"setErrors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VaultErrorControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultErrorControllerMetaData.ABI instead.
var VaultErrorControllerABI = VaultErrorControllerMetaData.ABI

// VaultErrorController is an auto generated Go binding around an Ethereum contract.
type VaultErrorController struct {
	VaultErrorControllerCaller     // Read-only binding to the contract
	VaultErrorControllerTransactor // Write-only binding to the contract
	VaultErrorControllerFilterer   // Log filterer for contract events
}

// VaultErrorControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultErrorControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultErrorControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultErrorControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultErrorControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultErrorControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultErrorControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultErrorControllerSession struct {
	Contract     *VaultErrorController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VaultErrorControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultErrorControllerCallerSession struct {
	Contract *VaultErrorControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VaultErrorControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultErrorControllerTransactorSession struct {
	Contract     *VaultErrorControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VaultErrorControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultErrorControllerRaw struct {
	Contract *VaultErrorController // Generic contract binding to access the raw methods on
}

// VaultErrorControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultErrorControllerCallerRaw struct {
	Contract *VaultErrorControllerCaller // Generic read-only contract binding to access the raw methods on
}

// VaultErrorControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultErrorControllerTransactorRaw struct {
	Contract *VaultErrorControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultErrorController creates a new instance of VaultErrorController, bound to a specific deployed contract.
func NewVaultErrorController(address common.Address, backend bind.ContractBackend) (*VaultErrorController, error) {
	contract, err := bindVaultErrorController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultErrorController{VaultErrorControllerCaller: VaultErrorControllerCaller{contract: contract}, VaultErrorControllerTransactor: VaultErrorControllerTransactor{contract: contract}, VaultErrorControllerFilterer: VaultErrorControllerFilterer{contract: contract}}, nil
}

// NewVaultErrorControllerCaller creates a new read-only instance of VaultErrorController, bound to a specific deployed contract.
func NewVaultErrorControllerCaller(address common.Address, caller bind.ContractCaller) (*VaultErrorControllerCaller, error) {
	contract, err := bindVaultErrorController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultErrorControllerCaller{contract: contract}, nil
}

// NewVaultErrorControllerTransactor creates a new write-only instance of VaultErrorController, bound to a specific deployed contract.
func NewVaultErrorControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultErrorControllerTransactor, error) {
	contract, err := bindVaultErrorController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultErrorControllerTransactor{contract: contract}, nil
}

// NewVaultErrorControllerFilterer creates a new log filterer instance of VaultErrorController, bound to a specific deployed contract.
func NewVaultErrorControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultErrorControllerFilterer, error) {
	contract, err := bindVaultErrorController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultErrorControllerFilterer{contract: contract}, nil
}

// bindVaultErrorController binds a generic wrapper to an already deployed contract.
func bindVaultErrorController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultErrorControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultErrorController *VaultErrorControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultErrorController.Contract.VaultErrorControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultErrorController *VaultErrorControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultErrorController.Contract.VaultErrorControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultErrorController *VaultErrorControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultErrorController.Contract.VaultErrorControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultErrorController *VaultErrorControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultErrorController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultErrorController *VaultErrorControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultErrorController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultErrorController *VaultErrorControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultErrorController.Contract.contract.Transact(opts, method, params...)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultErrorController *VaultErrorControllerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultErrorController.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultErrorController *VaultErrorControllerSession) Gov() (common.Address, error) {
	return _VaultErrorController.Contract.Gov(&_VaultErrorController.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultErrorController *VaultErrorControllerCallerSession) Gov() (common.Address, error) {
	return _VaultErrorController.Contract.Gov(&_VaultErrorController.CallOpts)
}

// SetErrors is a paid mutator transaction binding the contract method 0x6216c01b.
//
// Solidity: function setErrors(address _vault, string[] _errors) returns()
func (_VaultErrorController *VaultErrorControllerTransactor) SetErrors(opts *bind.TransactOpts, _vault common.Address, _errors []string) (*types.Transaction, error) {
	return _VaultErrorController.contract.Transact(opts, "setErrors", _vault, _errors)
}

// SetErrors is a paid mutator transaction binding the contract method 0x6216c01b.
//
// Solidity: function setErrors(address _vault, string[] _errors) returns()
func (_VaultErrorController *VaultErrorControllerSession) SetErrors(_vault common.Address, _errors []string) (*types.Transaction, error) {
	return _VaultErrorController.Contract.SetErrors(&_VaultErrorController.TransactOpts, _vault, _errors)
}

// SetErrors is a paid mutator transaction binding the contract method 0x6216c01b.
//
// Solidity: function setErrors(address _vault, string[] _errors) returns()
func (_VaultErrorController *VaultErrorControllerTransactorSession) SetErrors(_vault common.Address, _errors []string) (*types.Transaction, error) {
	return _VaultErrorController.Contract.SetErrors(&_VaultErrorController.TransactOpts, _vault, _errors)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultErrorController *VaultErrorControllerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _VaultErrorController.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultErrorController *VaultErrorControllerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultErrorController.Contract.SetGov(&_VaultErrorController.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultErrorController *VaultErrorControllerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultErrorController.Contract.SetGov(&_VaultErrorController.TransactOpts, _gov)
}
