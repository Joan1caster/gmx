// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package positionroutercallbackreceivertest

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

// PositionRouterCallbackReceiverTestMetaData contains all meta data concerning the PositionRouterCallbackReceiverTest contract.
var PositionRouterCallbackReceiverTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"positionKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isIncrease\",\"type\":\"bool\"}],\"name\":\"CallbackCalled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isIncrease\",\"type\":\"bool\"}],\"name\":\"gmxPositionCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PositionRouterCallbackReceiverTestABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionRouterCallbackReceiverTestMetaData.ABI instead.
var PositionRouterCallbackReceiverTestABI = PositionRouterCallbackReceiverTestMetaData.ABI

// PositionRouterCallbackReceiverTest is an auto generated Go binding around an Ethereum contract.
type PositionRouterCallbackReceiverTest struct {
	PositionRouterCallbackReceiverTestCaller     // Read-only binding to the contract
	PositionRouterCallbackReceiverTestTransactor // Write-only binding to the contract
	PositionRouterCallbackReceiverTestFilterer   // Log filterer for contract events
}

// PositionRouterCallbackReceiverTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionRouterCallbackReceiverTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterCallbackReceiverTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionRouterCallbackReceiverTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterCallbackReceiverTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionRouterCallbackReceiverTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterCallbackReceiverTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionRouterCallbackReceiverTestSession struct {
	Contract     *PositionRouterCallbackReceiverTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PositionRouterCallbackReceiverTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionRouterCallbackReceiverTestCallerSession struct {
	Contract *PositionRouterCallbackReceiverTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// PositionRouterCallbackReceiverTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionRouterCallbackReceiverTestTransactorSession struct {
	Contract     *PositionRouterCallbackReceiverTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// PositionRouterCallbackReceiverTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionRouterCallbackReceiverTestRaw struct {
	Contract *PositionRouterCallbackReceiverTest // Generic contract binding to access the raw methods on
}

// PositionRouterCallbackReceiverTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionRouterCallbackReceiverTestCallerRaw struct {
	Contract *PositionRouterCallbackReceiverTestCaller // Generic read-only contract binding to access the raw methods on
}

// PositionRouterCallbackReceiverTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionRouterCallbackReceiverTestTransactorRaw struct {
	Contract *PositionRouterCallbackReceiverTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionRouterCallbackReceiverTest creates a new instance of PositionRouterCallbackReceiverTest, bound to a specific deployed contract.
func NewPositionRouterCallbackReceiverTest(address common.Address, backend bind.ContractBackend) (*PositionRouterCallbackReceiverTest, error) {
	contract, err := bindPositionRouterCallbackReceiverTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCallbackReceiverTest{PositionRouterCallbackReceiverTestCaller: PositionRouterCallbackReceiverTestCaller{contract: contract}, PositionRouterCallbackReceiverTestTransactor: PositionRouterCallbackReceiverTestTransactor{contract: contract}, PositionRouterCallbackReceiverTestFilterer: PositionRouterCallbackReceiverTestFilterer{contract: contract}}, nil
}

// NewPositionRouterCallbackReceiverTestCaller creates a new read-only instance of PositionRouterCallbackReceiverTest, bound to a specific deployed contract.
func NewPositionRouterCallbackReceiverTestCaller(address common.Address, caller bind.ContractCaller) (*PositionRouterCallbackReceiverTestCaller, error) {
	contract, err := bindPositionRouterCallbackReceiverTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCallbackReceiverTestCaller{contract: contract}, nil
}

// NewPositionRouterCallbackReceiverTestTransactor creates a new write-only instance of PositionRouterCallbackReceiverTest, bound to a specific deployed contract.
func NewPositionRouterCallbackReceiverTestTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionRouterCallbackReceiverTestTransactor, error) {
	contract, err := bindPositionRouterCallbackReceiverTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCallbackReceiverTestTransactor{contract: contract}, nil
}

// NewPositionRouterCallbackReceiverTestFilterer creates a new log filterer instance of PositionRouterCallbackReceiverTest, bound to a specific deployed contract.
func NewPositionRouterCallbackReceiverTestFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionRouterCallbackReceiverTestFilterer, error) {
	contract, err := bindPositionRouterCallbackReceiverTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCallbackReceiverTestFilterer{contract: contract}, nil
}

// bindPositionRouterCallbackReceiverTest binds a generic wrapper to an already deployed contract.
func bindPositionRouterCallbackReceiverTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionRouterCallbackReceiverTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionRouterCallbackReceiverTest.Contract.PositionRouterCallbackReceiverTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionRouterCallbackReceiverTest.Contract.PositionRouterCallbackReceiverTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionRouterCallbackReceiverTest.Contract.PositionRouterCallbackReceiverTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionRouterCallbackReceiverTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionRouterCallbackReceiverTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionRouterCallbackReceiverTest.Contract.contract.Transact(opts, method, params...)
}

// GmxPositionCallback is a paid mutator transaction binding the contract method 0xedf3daec.
//
// Solidity: function gmxPositionCallback(bytes32 positionKey, bool isExecuted, bool isIncrease) returns()
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestTransactor) GmxPositionCallback(opts *bind.TransactOpts, positionKey [32]byte, isExecuted bool, isIncrease bool) (*types.Transaction, error) {
	return _PositionRouterCallbackReceiverTest.contract.Transact(opts, "gmxPositionCallback", positionKey, isExecuted, isIncrease)
}

// GmxPositionCallback is a paid mutator transaction binding the contract method 0xedf3daec.
//
// Solidity: function gmxPositionCallback(bytes32 positionKey, bool isExecuted, bool isIncrease) returns()
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestSession) GmxPositionCallback(positionKey [32]byte, isExecuted bool, isIncrease bool) (*types.Transaction, error) {
	return _PositionRouterCallbackReceiverTest.Contract.GmxPositionCallback(&_PositionRouterCallbackReceiverTest.TransactOpts, positionKey, isExecuted, isIncrease)
}

// GmxPositionCallback is a paid mutator transaction binding the contract method 0xedf3daec.
//
// Solidity: function gmxPositionCallback(bytes32 positionKey, bool isExecuted, bool isIncrease) returns()
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestTransactorSession) GmxPositionCallback(positionKey [32]byte, isExecuted bool, isIncrease bool) (*types.Transaction, error) {
	return _PositionRouterCallbackReceiverTest.Contract.GmxPositionCallback(&_PositionRouterCallbackReceiverTest.TransactOpts, positionKey, isExecuted, isIncrease)
}

// PositionRouterCallbackReceiverTestCallbackCalledIterator is returned from FilterCallbackCalled and is used to iterate over the raw logs and unpacked data for CallbackCalled events raised by the PositionRouterCallbackReceiverTest contract.
type PositionRouterCallbackReceiverTestCallbackCalledIterator struct {
	Event *PositionRouterCallbackReceiverTestCallbackCalled // Event containing the contract specifics and raw log

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
func (it *PositionRouterCallbackReceiverTestCallbackCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterCallbackReceiverTestCallbackCalled)
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
		it.Event = new(PositionRouterCallbackReceiverTestCallbackCalled)
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
func (it *PositionRouterCallbackReceiverTestCallbackCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterCallbackReceiverTestCallbackCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterCallbackReceiverTestCallbackCalled represents a CallbackCalled event raised by the PositionRouterCallbackReceiverTest contract.
type PositionRouterCallbackReceiverTestCallbackCalled struct {
	PositionKey [32]byte
	IsExecuted  bool
	IsIncrease  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallbackCalled is a free log retrieval operation binding the contract event 0x4c5b183c9c52e97ac4369ee0a4685a3af880b74946655c4b6e7366b6af173380.
//
// Solidity: event CallbackCalled(bytes32 positionKey, bool isExecuted, bool isIncrease)
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestFilterer) FilterCallbackCalled(opts *bind.FilterOpts) (*PositionRouterCallbackReceiverTestCallbackCalledIterator, error) {

	logs, sub, err := _PositionRouterCallbackReceiverTest.contract.FilterLogs(opts, "CallbackCalled")
	if err != nil {
		return nil, err
	}
	return &PositionRouterCallbackReceiverTestCallbackCalledIterator{contract: _PositionRouterCallbackReceiverTest.contract, event: "CallbackCalled", logs: logs, sub: sub}, nil
}

// WatchCallbackCalled is a free log subscription operation binding the contract event 0x4c5b183c9c52e97ac4369ee0a4685a3af880b74946655c4b6e7366b6af173380.
//
// Solidity: event CallbackCalled(bytes32 positionKey, bool isExecuted, bool isIncrease)
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestFilterer) WatchCallbackCalled(opts *bind.WatchOpts, sink chan<- *PositionRouterCallbackReceiverTestCallbackCalled) (event.Subscription, error) {

	logs, sub, err := _PositionRouterCallbackReceiverTest.contract.WatchLogs(opts, "CallbackCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterCallbackReceiverTestCallbackCalled)
				if err := _PositionRouterCallbackReceiverTest.contract.UnpackLog(event, "CallbackCalled", log); err != nil {
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

// ParseCallbackCalled is a log parse operation binding the contract event 0x4c5b183c9c52e97ac4369ee0a4685a3af880b74946655c4b6e7366b6af173380.
//
// Solidity: event CallbackCalled(bytes32 positionKey, bool isExecuted, bool isIncrease)
func (_PositionRouterCallbackReceiverTest *PositionRouterCallbackReceiverTestFilterer) ParseCallbackCalled(log types.Log) (*PositionRouterCallbackReceiverTestCallbackCalled, error) {
	event := new(PositionRouterCallbackReceiverTestCallbackCalled)
	if err := _PositionRouterCallbackReceiverTest.contract.UnpackLog(event, "CallbackCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
