// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package positionutils

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

// PositionUtilsMetaData contains all meta data concerning the PositionUtils contract.
var PositionUtilsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevLeverage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextLeverage\",\"type\":\"uint256\"}],\"name\":\"LeverageDecreased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PositionUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionUtilsMetaData.ABI instead.
var PositionUtilsABI = PositionUtilsMetaData.ABI

// PositionUtils is an auto generated Go binding around an Ethereum contract.
type PositionUtils struct {
	PositionUtilsCaller     // Read-only binding to the contract
	PositionUtilsTransactor // Write-only binding to the contract
	PositionUtilsFilterer   // Log filterer for contract events
}

// PositionUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionUtilsSession struct {
	Contract     *PositionUtils    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PositionUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionUtilsCallerSession struct {
	Contract *PositionUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PositionUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionUtilsTransactorSession struct {
	Contract     *PositionUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PositionUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionUtilsRaw struct {
	Contract *PositionUtils // Generic contract binding to access the raw methods on
}

// PositionUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionUtilsCallerRaw struct {
	Contract *PositionUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// PositionUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionUtilsTransactorRaw struct {
	Contract *PositionUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionUtils creates a new instance of PositionUtils, bound to a specific deployed contract.
func NewPositionUtils(address common.Address, backend bind.ContractBackend) (*PositionUtils, error) {
	contract, err := bindPositionUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionUtils{PositionUtilsCaller: PositionUtilsCaller{contract: contract}, PositionUtilsTransactor: PositionUtilsTransactor{contract: contract}, PositionUtilsFilterer: PositionUtilsFilterer{contract: contract}}, nil
}

// NewPositionUtilsCaller creates a new read-only instance of PositionUtils, bound to a specific deployed contract.
func NewPositionUtilsCaller(address common.Address, caller bind.ContractCaller) (*PositionUtilsCaller, error) {
	contract, err := bindPositionUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionUtilsCaller{contract: contract}, nil
}

// NewPositionUtilsTransactor creates a new write-only instance of PositionUtils, bound to a specific deployed contract.
func NewPositionUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionUtilsTransactor, error) {
	contract, err := bindPositionUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionUtilsTransactor{contract: contract}, nil
}

// NewPositionUtilsFilterer creates a new log filterer instance of PositionUtils, bound to a specific deployed contract.
func NewPositionUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionUtilsFilterer, error) {
	contract, err := bindPositionUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionUtilsFilterer{contract: contract}, nil
}

// bindPositionUtils binds a generic wrapper to an already deployed contract.
func bindPositionUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionUtils *PositionUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionUtils.Contract.PositionUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionUtils *PositionUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionUtils.Contract.PositionUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionUtils *PositionUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionUtils.Contract.PositionUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionUtils *PositionUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionUtils *PositionUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionUtils *PositionUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionUtils.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionUtils *PositionUtilsCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionUtils.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionUtils *PositionUtilsSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _PositionUtils.Contract.BASISPOINTSDIVISOR(&_PositionUtils.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionUtils *PositionUtilsCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _PositionUtils.Contract.BASISPOINTSDIVISOR(&_PositionUtils.CallOpts)
}

// PositionUtilsLeverageDecreasedIterator is returned from FilterLeverageDecreased and is used to iterate over the raw logs and unpacked data for LeverageDecreased events raised by the PositionUtils contract.
type PositionUtilsLeverageDecreasedIterator struct {
	Event *PositionUtilsLeverageDecreased // Event containing the contract specifics and raw log

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
func (it *PositionUtilsLeverageDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionUtilsLeverageDecreased)
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
		it.Event = new(PositionUtilsLeverageDecreased)
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
func (it *PositionUtilsLeverageDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionUtilsLeverageDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionUtilsLeverageDecreased represents a LeverageDecreased event raised by the PositionUtils contract.
type PositionUtilsLeverageDecreased struct {
	CollateralDelta *big.Int
	PrevLeverage    *big.Int
	NextLeverage    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLeverageDecreased is a free log retrieval operation binding the contract event 0x3f6e9241514ae172d9872f51274a73fd6b370b2f8fa612669bb17d933078860c.
//
// Solidity: event LeverageDecreased(uint256 collateralDelta, uint256 prevLeverage, uint256 nextLeverage)
func (_PositionUtils *PositionUtilsFilterer) FilterLeverageDecreased(opts *bind.FilterOpts) (*PositionUtilsLeverageDecreasedIterator, error) {

	logs, sub, err := _PositionUtils.contract.FilterLogs(opts, "LeverageDecreased")
	if err != nil {
		return nil, err
	}
	return &PositionUtilsLeverageDecreasedIterator{contract: _PositionUtils.contract, event: "LeverageDecreased", logs: logs, sub: sub}, nil
}

// WatchLeverageDecreased is a free log subscription operation binding the contract event 0x3f6e9241514ae172d9872f51274a73fd6b370b2f8fa612669bb17d933078860c.
//
// Solidity: event LeverageDecreased(uint256 collateralDelta, uint256 prevLeverage, uint256 nextLeverage)
func (_PositionUtils *PositionUtilsFilterer) WatchLeverageDecreased(opts *bind.WatchOpts, sink chan<- *PositionUtilsLeverageDecreased) (event.Subscription, error) {

	logs, sub, err := _PositionUtils.contract.WatchLogs(opts, "LeverageDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionUtilsLeverageDecreased)
				if err := _PositionUtils.contract.UnpackLog(event, "LeverageDecreased", log); err != nil {
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

// ParseLeverageDecreased is a log parse operation binding the contract event 0x3f6e9241514ae172d9872f51274a73fd6b370b2f8fa612669bb17d933078860c.
//
// Solidity: event LeverageDecreased(uint256 collateralDelta, uint256 prevLeverage, uint256 nextLeverage)
func (_PositionUtils *PositionUtilsFilterer) ParseLeverageDecreased(log types.Log) (*PositionUtilsLeverageDecreased, error) {
	event := new(PositionUtilsLeverageDecreased)
	if err := _PositionUtils.contract.UnpackLog(event, "LeverageDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
