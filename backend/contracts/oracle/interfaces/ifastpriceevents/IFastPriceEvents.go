// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ifastpriceevents

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

// IFastPriceEventsMetaData contains all meta data concerning the IFastPriceEvents contract.
var IFastPriceEventsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"emitPriceEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IFastPriceEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastPriceEventsMetaData.ABI instead.
var IFastPriceEventsABI = IFastPriceEventsMetaData.ABI

// IFastPriceEvents is an auto generated Go binding around an Ethereum contract.
type IFastPriceEvents struct {
	IFastPriceEventsCaller     // Read-only binding to the contract
	IFastPriceEventsTransactor // Write-only binding to the contract
	IFastPriceEventsFilterer   // Log filterer for contract events
}

// IFastPriceEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFastPriceEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastPriceEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastPriceEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastPriceEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastPriceEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastPriceEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastPriceEventsSession struct {
	Contract     *IFastPriceEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFastPriceEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastPriceEventsCallerSession struct {
	Contract *IFastPriceEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IFastPriceEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastPriceEventsTransactorSession struct {
	Contract     *IFastPriceEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IFastPriceEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFastPriceEventsRaw struct {
	Contract *IFastPriceEvents // Generic contract binding to access the raw methods on
}

// IFastPriceEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastPriceEventsCallerRaw struct {
	Contract *IFastPriceEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IFastPriceEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastPriceEventsTransactorRaw struct {
	Contract *IFastPriceEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastPriceEvents creates a new instance of IFastPriceEvents, bound to a specific deployed contract.
func NewIFastPriceEvents(address common.Address, backend bind.ContractBackend) (*IFastPriceEvents, error) {
	contract, err := bindIFastPriceEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastPriceEvents{IFastPriceEventsCaller: IFastPriceEventsCaller{contract: contract}, IFastPriceEventsTransactor: IFastPriceEventsTransactor{contract: contract}, IFastPriceEventsFilterer: IFastPriceEventsFilterer{contract: contract}}, nil
}

// NewIFastPriceEventsCaller creates a new read-only instance of IFastPriceEvents, bound to a specific deployed contract.
func NewIFastPriceEventsCaller(address common.Address, caller bind.ContractCaller) (*IFastPriceEventsCaller, error) {
	contract, err := bindIFastPriceEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastPriceEventsCaller{contract: contract}, nil
}

// NewIFastPriceEventsTransactor creates a new write-only instance of IFastPriceEvents, bound to a specific deployed contract.
func NewIFastPriceEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IFastPriceEventsTransactor, error) {
	contract, err := bindIFastPriceEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastPriceEventsTransactor{contract: contract}, nil
}

// NewIFastPriceEventsFilterer creates a new log filterer instance of IFastPriceEvents, bound to a specific deployed contract.
func NewIFastPriceEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IFastPriceEventsFilterer, error) {
	contract, err := bindIFastPriceEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastPriceEventsFilterer{contract: contract}, nil
}

// bindIFastPriceEvents binds a generic wrapper to an already deployed contract.
func bindIFastPriceEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFastPriceEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastPriceEvents *IFastPriceEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastPriceEvents.Contract.IFastPriceEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastPriceEvents *IFastPriceEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastPriceEvents.Contract.IFastPriceEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastPriceEvents *IFastPriceEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastPriceEvents.Contract.IFastPriceEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastPriceEvents *IFastPriceEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastPriceEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastPriceEvents *IFastPriceEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastPriceEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastPriceEvents *IFastPriceEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastPriceEvents.Contract.contract.Transact(opts, method, params...)
}

// EmitPriceEvent is a paid mutator transaction binding the contract method 0xe0409c71.
//
// Solidity: function emitPriceEvent(address _token, uint256 _price) returns()
func (_IFastPriceEvents *IFastPriceEventsTransactor) EmitPriceEvent(opts *bind.TransactOpts, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _IFastPriceEvents.contract.Transact(opts, "emitPriceEvent", _token, _price)
}

// EmitPriceEvent is a paid mutator transaction binding the contract method 0xe0409c71.
//
// Solidity: function emitPriceEvent(address _token, uint256 _price) returns()
func (_IFastPriceEvents *IFastPriceEventsSession) EmitPriceEvent(_token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _IFastPriceEvents.Contract.EmitPriceEvent(&_IFastPriceEvents.TransactOpts, _token, _price)
}

// EmitPriceEvent is a paid mutator transaction binding the contract method 0xe0409c71.
//
// Solidity: function emitPriceEvent(address _token, uint256 _price) returns()
func (_IFastPriceEvents *IFastPriceEventsTransactorSession) EmitPriceEvent(_token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _IFastPriceEvents.Contract.EmitPriceEvent(&_IFastPriceEvents.TransactOpts, _token, _price)
}
