// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ichainlinkflags

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

// IChainlinkFlagsMetaData contains all meta data concerning the IChainlinkFlags contract.
var IChainlinkFlagsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getFlag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IChainlinkFlagsABI is the input ABI used to generate the binding from.
// Deprecated: Use IChainlinkFlagsMetaData.ABI instead.
var IChainlinkFlagsABI = IChainlinkFlagsMetaData.ABI

// IChainlinkFlags is an auto generated Go binding around an Ethereum contract.
type IChainlinkFlags struct {
	IChainlinkFlagsCaller     // Read-only binding to the contract
	IChainlinkFlagsTransactor // Write-only binding to the contract
	IChainlinkFlagsFilterer   // Log filterer for contract events
}

// IChainlinkFlagsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChainlinkFlagsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChainlinkFlagsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChainlinkFlagsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChainlinkFlagsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChainlinkFlagsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChainlinkFlagsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChainlinkFlagsSession struct {
	Contract     *IChainlinkFlags  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IChainlinkFlagsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChainlinkFlagsCallerSession struct {
	Contract *IChainlinkFlagsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IChainlinkFlagsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChainlinkFlagsTransactorSession struct {
	Contract     *IChainlinkFlagsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IChainlinkFlagsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChainlinkFlagsRaw struct {
	Contract *IChainlinkFlags // Generic contract binding to access the raw methods on
}

// IChainlinkFlagsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChainlinkFlagsCallerRaw struct {
	Contract *IChainlinkFlagsCaller // Generic read-only contract binding to access the raw methods on
}

// IChainlinkFlagsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChainlinkFlagsTransactorRaw struct {
	Contract *IChainlinkFlagsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChainlinkFlags creates a new instance of IChainlinkFlags, bound to a specific deployed contract.
func NewIChainlinkFlags(address common.Address, backend bind.ContractBackend) (*IChainlinkFlags, error) {
	contract, err := bindIChainlinkFlags(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChainlinkFlags{IChainlinkFlagsCaller: IChainlinkFlagsCaller{contract: contract}, IChainlinkFlagsTransactor: IChainlinkFlagsTransactor{contract: contract}, IChainlinkFlagsFilterer: IChainlinkFlagsFilterer{contract: contract}}, nil
}

// NewIChainlinkFlagsCaller creates a new read-only instance of IChainlinkFlags, bound to a specific deployed contract.
func NewIChainlinkFlagsCaller(address common.Address, caller bind.ContractCaller) (*IChainlinkFlagsCaller, error) {
	contract, err := bindIChainlinkFlags(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChainlinkFlagsCaller{contract: contract}, nil
}

// NewIChainlinkFlagsTransactor creates a new write-only instance of IChainlinkFlags, bound to a specific deployed contract.
func NewIChainlinkFlagsTransactor(address common.Address, transactor bind.ContractTransactor) (*IChainlinkFlagsTransactor, error) {
	contract, err := bindIChainlinkFlags(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChainlinkFlagsTransactor{contract: contract}, nil
}

// NewIChainlinkFlagsFilterer creates a new log filterer instance of IChainlinkFlags, bound to a specific deployed contract.
func NewIChainlinkFlagsFilterer(address common.Address, filterer bind.ContractFilterer) (*IChainlinkFlagsFilterer, error) {
	contract, err := bindIChainlinkFlags(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChainlinkFlagsFilterer{contract: contract}, nil
}

// bindIChainlinkFlags binds a generic wrapper to an already deployed contract.
func bindIChainlinkFlags(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IChainlinkFlagsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChainlinkFlags *IChainlinkFlagsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChainlinkFlags.Contract.IChainlinkFlagsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChainlinkFlags *IChainlinkFlagsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChainlinkFlags.Contract.IChainlinkFlagsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChainlinkFlags *IChainlinkFlagsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChainlinkFlags.Contract.IChainlinkFlagsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChainlinkFlags *IChainlinkFlagsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChainlinkFlags.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChainlinkFlags *IChainlinkFlagsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChainlinkFlags.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChainlinkFlags *IChainlinkFlagsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChainlinkFlags.Contract.contract.Transact(opts, method, params...)
}

// GetFlag is a free data retrieval call binding the contract method 0x357e47fe.
//
// Solidity: function getFlag(address ) view returns(bool)
func (_IChainlinkFlags *IChainlinkFlagsCaller) GetFlag(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IChainlinkFlags.contract.Call(opts, &out, "getFlag", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFlag is a free data retrieval call binding the contract method 0x357e47fe.
//
// Solidity: function getFlag(address ) view returns(bool)
func (_IChainlinkFlags *IChainlinkFlagsSession) GetFlag(arg0 common.Address) (bool, error) {
	return _IChainlinkFlags.Contract.GetFlag(&_IChainlinkFlags.CallOpts, arg0)
}

// GetFlag is a free data retrieval call binding the contract method 0x357e47fe.
//
// Solidity: function getFlag(address ) view returns(bool)
func (_IChainlinkFlags *IChainlinkFlagsCallerSession) GetFlag(arg0 common.Address) (bool, error) {
	return _IChainlinkFlags.Contract.GetFlag(&_IChainlinkFlags.CallOpts, arg0)
}
