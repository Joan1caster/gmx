// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batchsender

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

// BatchSenderMetaData contains all meta data concerning the BatchSender contract.
var BatchSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"typeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchSend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_typeId\",\"type\":\"uint256\"}],\"name\":\"sendAndEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BatchSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchSenderMetaData.ABI instead.
var BatchSenderABI = BatchSenderMetaData.ABI

// BatchSender is an auto generated Go binding around an Ethereum contract.
type BatchSender struct {
	BatchSenderCaller     // Read-only binding to the contract
	BatchSenderTransactor // Write-only binding to the contract
	BatchSenderFilterer   // Log filterer for contract events
}

// BatchSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchSenderSession struct {
	Contract     *BatchSender      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchSenderCallerSession struct {
	Contract *BatchSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BatchSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchSenderTransactorSession struct {
	Contract     *BatchSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BatchSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchSenderRaw struct {
	Contract *BatchSender // Generic contract binding to access the raw methods on
}

// BatchSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchSenderCallerRaw struct {
	Contract *BatchSenderCaller // Generic read-only contract binding to access the raw methods on
}

// BatchSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchSenderTransactorRaw struct {
	Contract *BatchSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchSender creates a new instance of BatchSender, bound to a specific deployed contract.
func NewBatchSender(address common.Address, backend bind.ContractBackend) (*BatchSender, error) {
	contract, err := bindBatchSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchSender{BatchSenderCaller: BatchSenderCaller{contract: contract}, BatchSenderTransactor: BatchSenderTransactor{contract: contract}, BatchSenderFilterer: BatchSenderFilterer{contract: contract}}, nil
}

// NewBatchSenderCaller creates a new read-only instance of BatchSender, bound to a specific deployed contract.
func NewBatchSenderCaller(address common.Address, caller bind.ContractCaller) (*BatchSenderCaller, error) {
	contract, err := bindBatchSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchSenderCaller{contract: contract}, nil
}

// NewBatchSenderTransactor creates a new write-only instance of BatchSender, bound to a specific deployed contract.
func NewBatchSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchSenderTransactor, error) {
	contract, err := bindBatchSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchSenderTransactor{contract: contract}, nil
}

// NewBatchSenderFilterer creates a new log filterer instance of BatchSender, bound to a specific deployed contract.
func NewBatchSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchSenderFilterer, error) {
	contract, err := bindBatchSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchSenderFilterer{contract: contract}, nil
}

// bindBatchSender binds a generic wrapper to an already deployed contract.
func bindBatchSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchSenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchSender *BatchSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchSender.Contract.BatchSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchSender *BatchSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchSender.Contract.BatchSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchSender *BatchSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchSender.Contract.BatchSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchSender *BatchSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchSender *BatchSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchSender *BatchSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchSender.Contract.contract.Transact(opts, method, params...)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BatchSender *BatchSenderCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchSender.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BatchSender *BatchSenderSession) Gov() (common.Address, error) {
	return _BatchSender.Contract.Gov(&_BatchSender.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BatchSender *BatchSenderCallerSession) Gov() (common.Address, error) {
	return _BatchSender.Contract.Gov(&_BatchSender.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_BatchSender *BatchSenderCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BatchSender.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_BatchSender *BatchSenderSession) IsHandler(arg0 common.Address) (bool, error) {
	return _BatchSender.Contract.IsHandler(&_BatchSender.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_BatchSender *BatchSenderCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _BatchSender.Contract.IsHandler(&_BatchSender.CallOpts, arg0)
}

// Send is a paid mutator transaction binding the contract method 0xf8129cd2.
//
// Solidity: function send(address _token, address[] _accounts, uint256[] _amounts) returns()
func (_BatchSender *BatchSenderTransactor) Send(opts *bind.TransactOpts, _token common.Address, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _BatchSender.contract.Transact(opts, "send", _token, _accounts, _amounts)
}

// Send is a paid mutator transaction binding the contract method 0xf8129cd2.
//
// Solidity: function send(address _token, address[] _accounts, uint256[] _amounts) returns()
func (_BatchSender *BatchSenderSession) Send(_token common.Address, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _BatchSender.Contract.Send(&_BatchSender.TransactOpts, _token, _accounts, _amounts)
}

// Send is a paid mutator transaction binding the contract method 0xf8129cd2.
//
// Solidity: function send(address _token, address[] _accounts, uint256[] _amounts) returns()
func (_BatchSender *BatchSenderTransactorSession) Send(_token common.Address, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _BatchSender.Contract.Send(&_BatchSender.TransactOpts, _token, _accounts, _amounts)
}

// SendAndEmit is a paid mutator transaction binding the contract method 0x745ae40b.
//
// Solidity: function sendAndEmit(address _token, address[] _accounts, uint256[] _amounts, uint256 _typeId) returns()
func (_BatchSender *BatchSenderTransactor) SendAndEmit(opts *bind.TransactOpts, _token common.Address, _accounts []common.Address, _amounts []*big.Int, _typeId *big.Int) (*types.Transaction, error) {
	return _BatchSender.contract.Transact(opts, "sendAndEmit", _token, _accounts, _amounts, _typeId)
}

// SendAndEmit is a paid mutator transaction binding the contract method 0x745ae40b.
//
// Solidity: function sendAndEmit(address _token, address[] _accounts, uint256[] _amounts, uint256 _typeId) returns()
func (_BatchSender *BatchSenderSession) SendAndEmit(_token common.Address, _accounts []common.Address, _amounts []*big.Int, _typeId *big.Int) (*types.Transaction, error) {
	return _BatchSender.Contract.SendAndEmit(&_BatchSender.TransactOpts, _token, _accounts, _amounts, _typeId)
}

// SendAndEmit is a paid mutator transaction binding the contract method 0x745ae40b.
//
// Solidity: function sendAndEmit(address _token, address[] _accounts, uint256[] _amounts, uint256 _typeId) returns()
func (_BatchSender *BatchSenderTransactorSession) SendAndEmit(_token common.Address, _accounts []common.Address, _amounts []*big.Int, _typeId *big.Int) (*types.Transaction, error) {
	return _BatchSender.Contract.SendAndEmit(&_BatchSender.TransactOpts, _token, _accounts, _amounts, _typeId)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BatchSender *BatchSenderTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _BatchSender.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BatchSender *BatchSenderSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _BatchSender.Contract.SetGov(&_BatchSender.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BatchSender *BatchSenderTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _BatchSender.Contract.SetGov(&_BatchSender.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_BatchSender *BatchSenderTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _BatchSender.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_BatchSender *BatchSenderSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _BatchSender.Contract.SetHandler(&_BatchSender.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_BatchSender *BatchSenderTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _BatchSender.Contract.SetHandler(&_BatchSender.TransactOpts, _handler, _isActive)
}

// BatchSenderBatchSendIterator is returned from FilterBatchSend and is used to iterate over the raw logs and unpacked data for BatchSend events raised by the BatchSender contract.
type BatchSenderBatchSendIterator struct {
	Event *BatchSenderBatchSend // Event containing the contract specifics and raw log

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
func (it *BatchSenderBatchSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchSenderBatchSend)
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
		it.Event = new(BatchSenderBatchSend)
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
func (it *BatchSenderBatchSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchSenderBatchSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchSenderBatchSend represents a BatchSend event raised by the BatchSender contract.
type BatchSenderBatchSend struct {
	TypeId   *big.Int
	Token    common.Address
	Accounts []common.Address
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchSend is a free log retrieval operation binding the contract event 0xa1552778fd4edc5098fd82f614c100bf0f42c781e7926907643e2894679da0f3.
//
// Solidity: event BatchSend(uint256 indexed typeId, address indexed token, address[] accounts, uint256[] amounts)
func (_BatchSender *BatchSenderFilterer) FilterBatchSend(opts *bind.FilterOpts, typeId []*big.Int, token []common.Address) (*BatchSenderBatchSendIterator, error) {

	var typeIdRule []interface{}
	for _, typeIdItem := range typeId {
		typeIdRule = append(typeIdRule, typeIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BatchSender.contract.FilterLogs(opts, "BatchSend", typeIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BatchSenderBatchSendIterator{contract: _BatchSender.contract, event: "BatchSend", logs: logs, sub: sub}, nil
}

// WatchBatchSend is a free log subscription operation binding the contract event 0xa1552778fd4edc5098fd82f614c100bf0f42c781e7926907643e2894679da0f3.
//
// Solidity: event BatchSend(uint256 indexed typeId, address indexed token, address[] accounts, uint256[] amounts)
func (_BatchSender *BatchSenderFilterer) WatchBatchSend(opts *bind.WatchOpts, sink chan<- *BatchSenderBatchSend, typeId []*big.Int, token []common.Address) (event.Subscription, error) {

	var typeIdRule []interface{}
	for _, typeIdItem := range typeId {
		typeIdRule = append(typeIdRule, typeIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BatchSender.contract.WatchLogs(opts, "BatchSend", typeIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchSenderBatchSend)
				if err := _BatchSender.contract.UnpackLog(event, "BatchSend", log); err != nil {
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

// ParseBatchSend is a log parse operation binding the contract event 0xa1552778fd4edc5098fd82f614c100bf0f42c781e7926907643e2894679da0f3.
//
// Solidity: event BatchSend(uint256 indexed typeId, address indexed token, address[] accounts, uint256[] amounts)
func (_BatchSender *BatchSenderFilterer) ParseBatchSend(log types.Log) (*BatchSenderBatchSend, error) {
	event := new(BatchSenderBatchSend)
	if err := _BatchSender.contract.UnpackLog(event, "BatchSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
