// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itimelockcontroller

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

// ITimelockControllerMetaData contains all meta data concerning the ITimelockController contract.
var ITimelockControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ITimelockControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITimelockControllerMetaData.ABI instead.
var ITimelockControllerABI = ITimelockControllerMetaData.ABI

// ITimelockController is an auto generated Go binding around an Ethereum contract.
type ITimelockController struct {
	ITimelockControllerCaller     // Read-only binding to the contract
	ITimelockControllerTransactor // Write-only binding to the contract
	ITimelockControllerFilterer   // Log filterer for contract events
}

// ITimelockControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITimelockControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITimelockControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITimelockControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITimelockControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITimelockControllerSession struct {
	Contract     *ITimelockController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITimelockControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITimelockControllerCallerSession struct {
	Contract *ITimelockControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ITimelockControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITimelockControllerTransactorSession struct {
	Contract     *ITimelockControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ITimelockControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITimelockControllerRaw struct {
	Contract *ITimelockController // Generic contract binding to access the raw methods on
}

// ITimelockControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITimelockControllerCallerRaw struct {
	Contract *ITimelockControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ITimelockControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITimelockControllerTransactorRaw struct {
	Contract *ITimelockControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITimelockController creates a new instance of ITimelockController, bound to a specific deployed contract.
func NewITimelockController(address common.Address, backend bind.ContractBackend) (*ITimelockController, error) {
	contract, err := bindITimelockController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITimelockController{ITimelockControllerCaller: ITimelockControllerCaller{contract: contract}, ITimelockControllerTransactor: ITimelockControllerTransactor{contract: contract}, ITimelockControllerFilterer: ITimelockControllerFilterer{contract: contract}}, nil
}

// NewITimelockControllerCaller creates a new read-only instance of ITimelockController, bound to a specific deployed contract.
func NewITimelockControllerCaller(address common.Address, caller bind.ContractCaller) (*ITimelockControllerCaller, error) {
	contract, err := bindITimelockController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITimelockControllerCaller{contract: contract}, nil
}

// NewITimelockControllerTransactor creates a new write-only instance of ITimelockController, bound to a specific deployed contract.
func NewITimelockControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITimelockControllerTransactor, error) {
	contract, err := bindITimelockController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITimelockControllerTransactor{contract: contract}, nil
}

// NewITimelockControllerFilterer creates a new log filterer instance of ITimelockController, bound to a specific deployed contract.
func NewITimelockControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITimelockControllerFilterer, error) {
	contract, err := bindITimelockController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITimelockControllerFilterer{contract: contract}, nil
}

// bindITimelockController binds a generic wrapper to an already deployed contract.
func bindITimelockController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITimelockControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimelockController *ITimelockControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimelockController.Contract.ITimelockControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimelockController *ITimelockControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimelockController.Contract.ITimelockControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimelockController *ITimelockControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimelockController.Contract.ITimelockControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITimelockController *ITimelockControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITimelockController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITimelockController *ITimelockControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITimelockController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITimelockController *ITimelockControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITimelockController.Contract.contract.Transact(opts, method, params...)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_ITimelockController *ITimelockControllerTransactor) ExecuteBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _ITimelockController.contract.Transact(opts, "executeBatch", targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_ITimelockController *ITimelockControllerSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _ITimelockController.Contract.ExecuteBatch(&_ITimelockController.TransactOpts, targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_ITimelockController *ITimelockControllerTransactorSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _ITimelockController.Contract.ExecuteBatch(&_ITimelockController.TransactOpts, targets, values, payloads, predecessor, salt)
}
