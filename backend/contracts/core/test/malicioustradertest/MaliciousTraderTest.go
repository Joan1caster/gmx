// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package malicioustradertest

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

// MaliciousTraderTestMetaData contains all meta data concerning the MaliciousTraderTest contract.
var MaliciousTraderTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_positionRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_acceptablePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_referralCode\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callbackTarget\",\"type\":\"address\"}],\"name\":\"createIncreasePositionETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MaliciousTraderTestABI is the input ABI used to generate the binding from.
// Deprecated: Use MaliciousTraderTestMetaData.ABI instead.
var MaliciousTraderTestABI = MaliciousTraderTestMetaData.ABI

// MaliciousTraderTest is an auto generated Go binding around an Ethereum contract.
type MaliciousTraderTest struct {
	MaliciousTraderTestCaller     // Read-only binding to the contract
	MaliciousTraderTestTransactor // Write-only binding to the contract
	MaliciousTraderTestFilterer   // Log filterer for contract events
}

// MaliciousTraderTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type MaliciousTraderTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaliciousTraderTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MaliciousTraderTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaliciousTraderTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MaliciousTraderTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaliciousTraderTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MaliciousTraderTestSession struct {
	Contract     *MaliciousTraderTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MaliciousTraderTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MaliciousTraderTestCallerSession struct {
	Contract *MaliciousTraderTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MaliciousTraderTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MaliciousTraderTestTransactorSession struct {
	Contract     *MaliciousTraderTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MaliciousTraderTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type MaliciousTraderTestRaw struct {
	Contract *MaliciousTraderTest // Generic contract binding to access the raw methods on
}

// MaliciousTraderTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MaliciousTraderTestCallerRaw struct {
	Contract *MaliciousTraderTestCaller // Generic read-only contract binding to access the raw methods on
}

// MaliciousTraderTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MaliciousTraderTestTransactorRaw struct {
	Contract *MaliciousTraderTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaliciousTraderTest creates a new instance of MaliciousTraderTest, bound to a specific deployed contract.
func NewMaliciousTraderTest(address common.Address, backend bind.ContractBackend) (*MaliciousTraderTest, error) {
	contract, err := bindMaliciousTraderTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MaliciousTraderTest{MaliciousTraderTestCaller: MaliciousTraderTestCaller{contract: contract}, MaliciousTraderTestTransactor: MaliciousTraderTestTransactor{contract: contract}, MaliciousTraderTestFilterer: MaliciousTraderTestFilterer{contract: contract}}, nil
}

// NewMaliciousTraderTestCaller creates a new read-only instance of MaliciousTraderTest, bound to a specific deployed contract.
func NewMaliciousTraderTestCaller(address common.Address, caller bind.ContractCaller) (*MaliciousTraderTestCaller, error) {
	contract, err := bindMaliciousTraderTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaliciousTraderTestCaller{contract: contract}, nil
}

// NewMaliciousTraderTestTransactor creates a new write-only instance of MaliciousTraderTest, bound to a specific deployed contract.
func NewMaliciousTraderTestTransactor(address common.Address, transactor bind.ContractTransactor) (*MaliciousTraderTestTransactor, error) {
	contract, err := bindMaliciousTraderTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaliciousTraderTestTransactor{contract: contract}, nil
}

// NewMaliciousTraderTestFilterer creates a new log filterer instance of MaliciousTraderTest, bound to a specific deployed contract.
func NewMaliciousTraderTestFilterer(address common.Address, filterer bind.ContractFilterer) (*MaliciousTraderTestFilterer, error) {
	contract, err := bindMaliciousTraderTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaliciousTraderTestFilterer{contract: contract}, nil
}

// bindMaliciousTraderTest binds a generic wrapper to an already deployed contract.
func bindMaliciousTraderTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MaliciousTraderTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaliciousTraderTest *MaliciousTraderTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaliciousTraderTest.Contract.MaliciousTraderTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaliciousTraderTest *MaliciousTraderTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.MaliciousTraderTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaliciousTraderTest *MaliciousTraderTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.MaliciousTraderTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaliciousTraderTest *MaliciousTraderTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaliciousTraderTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaliciousTraderTest *MaliciousTraderTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaliciousTraderTest *MaliciousTraderTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.contract.Transact(opts, method, params...)
}

// PositionRouter is a free data retrieval call binding the contract method 0x61ef161f.
//
// Solidity: function positionRouter() view returns(address)
func (_MaliciousTraderTest *MaliciousTraderTestCaller) PositionRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MaliciousTraderTest.contract.Call(opts, &out, "positionRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionRouter is a free data retrieval call binding the contract method 0x61ef161f.
//
// Solidity: function positionRouter() view returns(address)
func (_MaliciousTraderTest *MaliciousTraderTestSession) PositionRouter() (common.Address, error) {
	return _MaliciousTraderTest.Contract.PositionRouter(&_MaliciousTraderTest.CallOpts)
}

// PositionRouter is a free data retrieval call binding the contract method 0x61ef161f.
//
// Solidity: function positionRouter() view returns(address)
func (_MaliciousTraderTest *MaliciousTraderTestCallerSession) PositionRouter() (common.Address, error) {
	return _MaliciousTraderTest.Contract.PositionRouter(&_MaliciousTraderTest.CallOpts)
}

// CreateIncreasePositionETH is a paid mutator transaction binding the contract method 0x5b88e8c6.
//
// Solidity: function createIncreasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns()
func (_MaliciousTraderTest *MaliciousTraderTestTransactor) CreateIncreasePositionETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _MaliciousTraderTest.contract.Transact(opts, "createIncreasePositionETH", _path, _indexToken, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// CreateIncreasePositionETH is a paid mutator transaction binding the contract method 0x5b88e8c6.
//
// Solidity: function createIncreasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns()
func (_MaliciousTraderTest *MaliciousTraderTestSession) CreateIncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.CreateIncreasePositionETH(&_MaliciousTraderTest.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// CreateIncreasePositionETH is a paid mutator transaction binding the contract method 0x5b88e8c6.
//
// Solidity: function createIncreasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns()
func (_MaliciousTraderTest *MaliciousTraderTestTransactorSession) CreateIncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.CreateIncreasePositionETH(&_MaliciousTraderTest.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MaliciousTraderTest *MaliciousTraderTestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaliciousTraderTest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MaliciousTraderTest *MaliciousTraderTestSession) Receive() (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.Receive(&_MaliciousTraderTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MaliciousTraderTest *MaliciousTraderTestTransactorSession) Receive() (*types.Transaction, error) {
	return _MaliciousTraderTest.Contract.Receive(&_MaliciousTraderTest.TransactOpts)
}

// MaliciousTraderTestReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the MaliciousTraderTest contract.
type MaliciousTraderTestReceivedIterator struct {
	Event *MaliciousTraderTestReceived // Event containing the contract specifics and raw log

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
func (it *MaliciousTraderTestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaliciousTraderTestReceived)
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
		it.Event = new(MaliciousTraderTestReceived)
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
func (it *MaliciousTraderTestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaliciousTraderTestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaliciousTraderTestReceived represents a Received event raised by the MaliciousTraderTest contract.
type MaliciousTraderTestReceived struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x544c765b33ca411cce832250371569244f765a17fcd217832be093f0fd5fa45b.
//
// Solidity: event Received()
func (_MaliciousTraderTest *MaliciousTraderTestFilterer) FilterReceived(opts *bind.FilterOpts) (*MaliciousTraderTestReceivedIterator, error) {

	logs, sub, err := _MaliciousTraderTest.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &MaliciousTraderTestReceivedIterator{contract: _MaliciousTraderTest.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x544c765b33ca411cce832250371569244f765a17fcd217832be093f0fd5fa45b.
//
// Solidity: event Received()
func (_MaliciousTraderTest *MaliciousTraderTestFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *MaliciousTraderTestReceived) (event.Subscription, error) {

	logs, sub, err := _MaliciousTraderTest.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaliciousTraderTestReceived)
				if err := _MaliciousTraderTest.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x544c765b33ca411cce832250371569244f765a17fcd217832be093f0fd5fa45b.
//
// Solidity: event Received()
func (_MaliciousTraderTest *MaliciousTraderTestFilterer) ParseReceived(log types.Log) (*MaliciousTraderTestReceived, error) {
	event := new(MaliciousTraderTestReceived)
	if err := _MaliciousTraderTest.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
