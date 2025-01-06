// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package istabilizestrategy

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

// IStabilizeStrategyMetaData contains all meta data concerning the IStabilizeStrategy contract.
var IStabilizeStrategyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"governanceFinishMoveEsGMXFromDeprecatedRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStabilizeStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use IStabilizeStrategyMetaData.ABI instead.
var IStabilizeStrategyABI = IStabilizeStrategyMetaData.ABI

// IStabilizeStrategy is an auto generated Go binding around an Ethereum contract.
type IStabilizeStrategy struct {
	IStabilizeStrategyCaller     // Read-only binding to the contract
	IStabilizeStrategyTransactor // Write-only binding to the contract
	IStabilizeStrategyFilterer   // Log filterer for contract events
}

// IStabilizeStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStabilizeStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStabilizeStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStabilizeStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStabilizeStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStabilizeStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStabilizeStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStabilizeStrategySession struct {
	Contract     *IStabilizeStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStabilizeStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStabilizeStrategyCallerSession struct {
	Contract *IStabilizeStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IStabilizeStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStabilizeStrategyTransactorSession struct {
	Contract     *IStabilizeStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IStabilizeStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStabilizeStrategyRaw struct {
	Contract *IStabilizeStrategy // Generic contract binding to access the raw methods on
}

// IStabilizeStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStabilizeStrategyCallerRaw struct {
	Contract *IStabilizeStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// IStabilizeStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStabilizeStrategyTransactorRaw struct {
	Contract *IStabilizeStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStabilizeStrategy creates a new instance of IStabilizeStrategy, bound to a specific deployed contract.
func NewIStabilizeStrategy(address common.Address, backend bind.ContractBackend) (*IStabilizeStrategy, error) {
	contract, err := bindIStabilizeStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStabilizeStrategy{IStabilizeStrategyCaller: IStabilizeStrategyCaller{contract: contract}, IStabilizeStrategyTransactor: IStabilizeStrategyTransactor{contract: contract}, IStabilizeStrategyFilterer: IStabilizeStrategyFilterer{contract: contract}}, nil
}

// NewIStabilizeStrategyCaller creates a new read-only instance of IStabilizeStrategy, bound to a specific deployed contract.
func NewIStabilizeStrategyCaller(address common.Address, caller bind.ContractCaller) (*IStabilizeStrategyCaller, error) {
	contract, err := bindIStabilizeStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStabilizeStrategyCaller{contract: contract}, nil
}

// NewIStabilizeStrategyTransactor creates a new write-only instance of IStabilizeStrategy, bound to a specific deployed contract.
func NewIStabilizeStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*IStabilizeStrategyTransactor, error) {
	contract, err := bindIStabilizeStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStabilizeStrategyTransactor{contract: contract}, nil
}

// NewIStabilizeStrategyFilterer creates a new log filterer instance of IStabilizeStrategy, bound to a specific deployed contract.
func NewIStabilizeStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*IStabilizeStrategyFilterer, error) {
	contract, err := bindIStabilizeStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStabilizeStrategyFilterer{contract: contract}, nil
}

// bindIStabilizeStrategy binds a generic wrapper to an already deployed contract.
func bindIStabilizeStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStabilizeStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStabilizeStrategy *IStabilizeStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStabilizeStrategy.Contract.IStabilizeStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStabilizeStrategy *IStabilizeStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStabilizeStrategy.Contract.IStabilizeStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStabilizeStrategy *IStabilizeStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStabilizeStrategy.Contract.IStabilizeStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStabilizeStrategy *IStabilizeStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStabilizeStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStabilizeStrategy *IStabilizeStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStabilizeStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStabilizeStrategy *IStabilizeStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStabilizeStrategy.Contract.contract.Transact(opts, method, params...)
}

// GovernanceFinishMoveEsGMXFromDeprecatedRouter is a paid mutator transaction binding the contract method 0x6ba5f6ea.
//
// Solidity: function governanceFinishMoveEsGMXFromDeprecatedRouter(address _sender) returns()
func (_IStabilizeStrategy *IStabilizeStrategyTransactor) GovernanceFinishMoveEsGMXFromDeprecatedRouter(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _IStabilizeStrategy.contract.Transact(opts, "governanceFinishMoveEsGMXFromDeprecatedRouter", _sender)
}

// GovernanceFinishMoveEsGMXFromDeprecatedRouter is a paid mutator transaction binding the contract method 0x6ba5f6ea.
//
// Solidity: function governanceFinishMoveEsGMXFromDeprecatedRouter(address _sender) returns()
func (_IStabilizeStrategy *IStabilizeStrategySession) GovernanceFinishMoveEsGMXFromDeprecatedRouter(_sender common.Address) (*types.Transaction, error) {
	return _IStabilizeStrategy.Contract.GovernanceFinishMoveEsGMXFromDeprecatedRouter(&_IStabilizeStrategy.TransactOpts, _sender)
}

// GovernanceFinishMoveEsGMXFromDeprecatedRouter is a paid mutator transaction binding the contract method 0x6ba5f6ea.
//
// Solidity: function governanceFinishMoveEsGMXFromDeprecatedRouter(address _sender) returns()
func (_IStabilizeStrategy *IStabilizeStrategyTransactorSession) GovernanceFinishMoveEsGMXFromDeprecatedRouter(_sender common.Address) (*types.Transaction, error) {
	return _IStabilizeStrategy.Contract.GovernanceFinishMoveEsGMXFromDeprecatedRouter(&_IStabilizeStrategy.TransactOpts, _sender)
}
