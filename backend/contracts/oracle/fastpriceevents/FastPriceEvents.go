// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fastpriceevents

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

// FastPriceEventsMetaData contains all meta data concerning the FastPriceEvents contract.
var FastPriceEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"emitPriceEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPriceFeed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isPriceFeed\",\"type\":\"bool\"}],\"name\":\"setIsPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FastPriceEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use FastPriceEventsMetaData.ABI instead.
var FastPriceEventsABI = FastPriceEventsMetaData.ABI

// FastPriceEvents is an auto generated Go binding around an Ethereum contract.
type FastPriceEvents struct {
	FastPriceEventsCaller     // Read-only binding to the contract
	FastPriceEventsTransactor // Write-only binding to the contract
	FastPriceEventsFilterer   // Log filterer for contract events
}

// FastPriceEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastPriceEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastPriceEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastPriceEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastPriceEventsSession struct {
	Contract     *FastPriceEvents  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FastPriceEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastPriceEventsCallerSession struct {
	Contract *FastPriceEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FastPriceEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastPriceEventsTransactorSession struct {
	Contract     *FastPriceEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FastPriceEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastPriceEventsRaw struct {
	Contract *FastPriceEvents // Generic contract binding to access the raw methods on
}

// FastPriceEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastPriceEventsCallerRaw struct {
	Contract *FastPriceEventsCaller // Generic read-only contract binding to access the raw methods on
}

// FastPriceEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastPriceEventsTransactorRaw struct {
	Contract *FastPriceEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastPriceEvents creates a new instance of FastPriceEvents, bound to a specific deployed contract.
func NewFastPriceEvents(address common.Address, backend bind.ContractBackend) (*FastPriceEvents, error) {
	contract, err := bindFastPriceEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastPriceEvents{FastPriceEventsCaller: FastPriceEventsCaller{contract: contract}, FastPriceEventsTransactor: FastPriceEventsTransactor{contract: contract}, FastPriceEventsFilterer: FastPriceEventsFilterer{contract: contract}}, nil
}

// NewFastPriceEventsCaller creates a new read-only instance of FastPriceEvents, bound to a specific deployed contract.
func NewFastPriceEventsCaller(address common.Address, caller bind.ContractCaller) (*FastPriceEventsCaller, error) {
	contract, err := bindFastPriceEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastPriceEventsCaller{contract: contract}, nil
}

// NewFastPriceEventsTransactor creates a new write-only instance of FastPriceEvents, bound to a specific deployed contract.
func NewFastPriceEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*FastPriceEventsTransactor, error) {
	contract, err := bindFastPriceEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastPriceEventsTransactor{contract: contract}, nil
}

// NewFastPriceEventsFilterer creates a new log filterer instance of FastPriceEvents, bound to a specific deployed contract.
func NewFastPriceEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*FastPriceEventsFilterer, error) {
	contract, err := bindFastPriceEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastPriceEventsFilterer{contract: contract}, nil
}

// bindFastPriceEvents binds a generic wrapper to an already deployed contract.
func bindFastPriceEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastPriceEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastPriceEvents *FastPriceEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastPriceEvents.Contract.FastPriceEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastPriceEvents *FastPriceEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.FastPriceEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastPriceEvents *FastPriceEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.FastPriceEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastPriceEvents *FastPriceEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastPriceEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastPriceEvents *FastPriceEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastPriceEvents *FastPriceEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.contract.Transact(opts, method, params...)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceEvents *FastPriceEventsCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceEvents.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceEvents *FastPriceEventsSession) Gov() (common.Address, error) {
	return _FastPriceEvents.Contract.Gov(&_FastPriceEvents.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceEvents *FastPriceEventsCallerSession) Gov() (common.Address, error) {
	return _FastPriceEvents.Contract.Gov(&_FastPriceEvents.CallOpts)
}

// IsPriceFeed is a free data retrieval call binding the contract method 0x57a94beb.
//
// Solidity: function isPriceFeed(address ) view returns(bool)
func (_FastPriceEvents *FastPriceEventsCaller) IsPriceFeed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceEvents.contract.Call(opts, &out, "isPriceFeed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPriceFeed is a free data retrieval call binding the contract method 0x57a94beb.
//
// Solidity: function isPriceFeed(address ) view returns(bool)
func (_FastPriceEvents *FastPriceEventsSession) IsPriceFeed(arg0 common.Address) (bool, error) {
	return _FastPriceEvents.Contract.IsPriceFeed(&_FastPriceEvents.CallOpts, arg0)
}

// IsPriceFeed is a free data retrieval call binding the contract method 0x57a94beb.
//
// Solidity: function isPriceFeed(address ) view returns(bool)
func (_FastPriceEvents *FastPriceEventsCallerSession) IsPriceFeed(arg0 common.Address) (bool, error) {
	return _FastPriceEvents.Contract.IsPriceFeed(&_FastPriceEvents.CallOpts, arg0)
}

// EmitPriceEvent is a paid mutator transaction binding the contract method 0xe0409c71.
//
// Solidity: function emitPriceEvent(address _token, uint256 _price) returns()
func (_FastPriceEvents *FastPriceEventsTransactor) EmitPriceEvent(opts *bind.TransactOpts, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _FastPriceEvents.contract.Transact(opts, "emitPriceEvent", _token, _price)
}

// EmitPriceEvent is a paid mutator transaction binding the contract method 0xe0409c71.
//
// Solidity: function emitPriceEvent(address _token, uint256 _price) returns()
func (_FastPriceEvents *FastPriceEventsSession) EmitPriceEvent(_token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.EmitPriceEvent(&_FastPriceEvents.TransactOpts, _token, _price)
}

// EmitPriceEvent is a paid mutator transaction binding the contract method 0xe0409c71.
//
// Solidity: function emitPriceEvent(address _token, uint256 _price) returns()
func (_FastPriceEvents *FastPriceEventsTransactorSession) EmitPriceEvent(_token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.EmitPriceEvent(&_FastPriceEvents.TransactOpts, _token, _price)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceEvents *FastPriceEventsTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _FastPriceEvents.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceEvents *FastPriceEventsSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.SetGov(&_FastPriceEvents.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceEvents *FastPriceEventsTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.SetGov(&_FastPriceEvents.TransactOpts, _gov)
}

// SetIsPriceFeed is a paid mutator transaction binding the contract method 0x69d4c924.
//
// Solidity: function setIsPriceFeed(address _priceFeed, bool _isPriceFeed) returns()
func (_FastPriceEvents *FastPriceEventsTransactor) SetIsPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address, _isPriceFeed bool) (*types.Transaction, error) {
	return _FastPriceEvents.contract.Transact(opts, "setIsPriceFeed", _priceFeed, _isPriceFeed)
}

// SetIsPriceFeed is a paid mutator transaction binding the contract method 0x69d4c924.
//
// Solidity: function setIsPriceFeed(address _priceFeed, bool _isPriceFeed) returns()
func (_FastPriceEvents *FastPriceEventsSession) SetIsPriceFeed(_priceFeed common.Address, _isPriceFeed bool) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.SetIsPriceFeed(&_FastPriceEvents.TransactOpts, _priceFeed, _isPriceFeed)
}

// SetIsPriceFeed is a paid mutator transaction binding the contract method 0x69d4c924.
//
// Solidity: function setIsPriceFeed(address _priceFeed, bool _isPriceFeed) returns()
func (_FastPriceEvents *FastPriceEventsTransactorSession) SetIsPriceFeed(_priceFeed common.Address, _isPriceFeed bool) (*types.Transaction, error) {
	return _FastPriceEvents.Contract.SetIsPriceFeed(&_FastPriceEvents.TransactOpts, _priceFeed, _isPriceFeed)
}

// FastPriceEventsPriceUpdateIterator is returned from FilterPriceUpdate and is used to iterate over the raw logs and unpacked data for PriceUpdate events raised by the FastPriceEvents contract.
type FastPriceEventsPriceUpdateIterator struct {
	Event *FastPriceEventsPriceUpdate // Event containing the contract specifics and raw log

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
func (it *FastPriceEventsPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceEventsPriceUpdate)
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
		it.Event = new(FastPriceEventsPriceUpdate)
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
func (it *FastPriceEventsPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceEventsPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceEventsPriceUpdate represents a PriceUpdate event raised by the FastPriceEvents contract.
type FastPriceEventsPriceUpdate struct {
	Token     common.Address
	Price     *big.Int
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdate is a free log retrieval operation binding the contract event 0xc37a77b91cc3fc2d0e4b43fd2f347ec67adda10e39215de4742836cc3e42c97a.
//
// Solidity: event PriceUpdate(address token, uint256 price, address priceFeed)
func (_FastPriceEvents *FastPriceEventsFilterer) FilterPriceUpdate(opts *bind.FilterOpts) (*FastPriceEventsPriceUpdateIterator, error) {

	logs, sub, err := _FastPriceEvents.contract.FilterLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return &FastPriceEventsPriceUpdateIterator{contract: _FastPriceEvents.contract, event: "PriceUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceUpdate is a free log subscription operation binding the contract event 0xc37a77b91cc3fc2d0e4b43fd2f347ec67adda10e39215de4742836cc3e42c97a.
//
// Solidity: event PriceUpdate(address token, uint256 price, address priceFeed)
func (_FastPriceEvents *FastPriceEventsFilterer) WatchPriceUpdate(opts *bind.WatchOpts, sink chan<- *FastPriceEventsPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _FastPriceEvents.contract.WatchLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceEventsPriceUpdate)
				if err := _FastPriceEvents.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
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

// ParsePriceUpdate is a log parse operation binding the contract event 0xc37a77b91cc3fc2d0e4b43fd2f347ec67adda10e39215de4742836cc3e42c97a.
//
// Solidity: event PriceUpdate(address token, uint256 price, address priceFeed)
func (_FastPriceEvents *FastPriceEventsFilterer) ParsePriceUpdate(log types.Log) (*FastPriceEventsPriceUpdate, error) {
	event := new(FastPriceEventsPriceUpdate)
	if err := _FastPriceEvents.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
