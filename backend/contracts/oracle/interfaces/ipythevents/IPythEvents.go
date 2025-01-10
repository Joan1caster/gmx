// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipythevents

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

// IPythEventsMetaData contains all meta data concerning the IPythEvents contract.
var IPythEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"publishTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"}],\"name\":\"PriceFeedUpdate\",\"type\":\"event\"}]",
}

// IPythEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IPythEventsMetaData.ABI instead.
var IPythEventsABI = IPythEventsMetaData.ABI

// IPythEvents is an auto generated Go binding around an Ethereum contract.
type IPythEvents struct {
	IPythEventsCaller     // Read-only binding to the contract
	IPythEventsTransactor // Write-only binding to the contract
	IPythEventsFilterer   // Log filterer for contract events
}

// IPythEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPythEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPythEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPythEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPythEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPythEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPythEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPythEventsSession struct {
	Contract     *IPythEvents      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPythEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPythEventsCallerSession struct {
	Contract *IPythEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IPythEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPythEventsTransactorSession struct {
	Contract     *IPythEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IPythEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPythEventsRaw struct {
	Contract *IPythEvents // Generic contract binding to access the raw methods on
}

// IPythEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPythEventsCallerRaw struct {
	Contract *IPythEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IPythEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPythEventsTransactorRaw struct {
	Contract *IPythEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPythEvents creates a new instance of IPythEvents, bound to a specific deployed contract.
func NewIPythEvents(address common.Address, backend bind.ContractBackend) (*IPythEvents, error) {
	contract, err := bindIPythEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPythEvents{IPythEventsCaller: IPythEventsCaller{contract: contract}, IPythEventsTransactor: IPythEventsTransactor{contract: contract}, IPythEventsFilterer: IPythEventsFilterer{contract: contract}}, nil
}

// NewIPythEventsCaller creates a new read-only instance of IPythEvents, bound to a specific deployed contract.
func NewIPythEventsCaller(address common.Address, caller bind.ContractCaller) (*IPythEventsCaller, error) {
	contract, err := bindIPythEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPythEventsCaller{contract: contract}, nil
}

// NewIPythEventsTransactor creates a new write-only instance of IPythEvents, bound to a specific deployed contract.
func NewIPythEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IPythEventsTransactor, error) {
	contract, err := bindIPythEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPythEventsTransactor{contract: contract}, nil
}

// NewIPythEventsFilterer creates a new log filterer instance of IPythEvents, bound to a specific deployed contract.
func NewIPythEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IPythEventsFilterer, error) {
	contract, err := bindIPythEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPythEventsFilterer{contract: contract}, nil
}

// bindIPythEvents binds a generic wrapper to an already deployed contract.
func bindIPythEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPythEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPythEvents *IPythEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPythEvents.Contract.IPythEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPythEvents *IPythEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPythEvents.Contract.IPythEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPythEvents *IPythEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPythEvents.Contract.IPythEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPythEvents *IPythEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPythEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPythEvents *IPythEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPythEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPythEvents *IPythEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPythEvents.Contract.contract.Transact(opts, method, params...)
}

// IPythEventsPriceFeedUpdateIterator is returned from FilterPriceFeedUpdate and is used to iterate over the raw logs and unpacked data for PriceFeedUpdate events raised by the IPythEvents contract.
type IPythEventsPriceFeedUpdateIterator struct {
	Event *IPythEventsPriceFeedUpdate // Event containing the contract specifics and raw log

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
func (it *IPythEventsPriceFeedUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPythEventsPriceFeedUpdate)
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
		it.Event = new(IPythEventsPriceFeedUpdate)
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
func (it *IPythEventsPriceFeedUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPythEventsPriceFeedUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPythEventsPriceFeedUpdate represents a PriceFeedUpdate event raised by the IPythEvents contract.
type IPythEventsPriceFeedUpdate struct {
	Id          [32]byte
	PublishTime uint64
	Price       int64
	Conf        uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedUpdate is a free log retrieval operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_IPythEvents *IPythEventsFilterer) FilterPriceFeedUpdate(opts *bind.FilterOpts, id [][32]byte) (*IPythEventsPriceFeedUpdateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IPythEvents.contract.FilterLogs(opts, "PriceFeedUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return &IPythEventsPriceFeedUpdateIterator{contract: _IPythEvents.contract, event: "PriceFeedUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceFeedUpdate is a free log subscription operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_IPythEvents *IPythEventsFilterer) WatchPriceFeedUpdate(opts *bind.WatchOpts, sink chan<- *IPythEventsPriceFeedUpdate, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IPythEvents.contract.WatchLogs(opts, "PriceFeedUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPythEventsPriceFeedUpdate)
				if err := _IPythEvents.contract.UnpackLog(event, "PriceFeedUpdate", log); err != nil {
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

// ParsePriceFeedUpdate is a log parse operation binding the contract event 0xd06a6b7f4918494b3719217d1802786c1f5112a6c1d88fe2cfec00b4584f6aec.
//
// Solidity: event PriceFeedUpdate(bytes32 indexed id, uint64 publishTime, int64 price, uint64 conf)
func (_IPythEvents *IPythEventsFilterer) ParsePriceFeedUpdate(log types.Log) (*IPythEventsPriceFeedUpdate, error) {
	event := new(IPythEventsPriceFeedUpdate)
	if err := _IPythEvents.contract.UnpackLog(event, "PriceFeedUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
