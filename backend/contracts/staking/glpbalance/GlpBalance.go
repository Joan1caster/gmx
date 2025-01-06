// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package glpbalance

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

// GlpBalanceMetaData contains all meta data concerning the GlpBalance contract.
var GlpBalanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGlpManager\",\"name\":\"_glpManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpManager\",\"outputs\":[{\"internalType\":\"contractIGlpManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GlpBalanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GlpBalanceMetaData.ABI instead.
var GlpBalanceABI = GlpBalanceMetaData.ABI

// GlpBalance is an auto generated Go binding around an Ethereum contract.
type GlpBalance struct {
	GlpBalanceCaller     // Read-only binding to the contract
	GlpBalanceTransactor // Write-only binding to the contract
	GlpBalanceFilterer   // Log filterer for contract events
}

// GlpBalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlpBalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpBalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlpBalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpBalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlpBalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpBalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlpBalanceSession struct {
	Contract     *GlpBalance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlpBalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlpBalanceCallerSession struct {
	Contract *GlpBalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GlpBalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlpBalanceTransactorSession struct {
	Contract     *GlpBalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GlpBalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlpBalanceRaw struct {
	Contract *GlpBalance // Generic contract binding to access the raw methods on
}

// GlpBalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlpBalanceCallerRaw struct {
	Contract *GlpBalanceCaller // Generic read-only contract binding to access the raw methods on
}

// GlpBalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlpBalanceTransactorRaw struct {
	Contract *GlpBalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlpBalance creates a new instance of GlpBalance, bound to a specific deployed contract.
func NewGlpBalance(address common.Address, backend bind.ContractBackend) (*GlpBalance, error) {
	contract, err := bindGlpBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlpBalance{GlpBalanceCaller: GlpBalanceCaller{contract: contract}, GlpBalanceTransactor: GlpBalanceTransactor{contract: contract}, GlpBalanceFilterer: GlpBalanceFilterer{contract: contract}}, nil
}

// NewGlpBalanceCaller creates a new read-only instance of GlpBalance, bound to a specific deployed contract.
func NewGlpBalanceCaller(address common.Address, caller bind.ContractCaller) (*GlpBalanceCaller, error) {
	contract, err := bindGlpBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlpBalanceCaller{contract: contract}, nil
}

// NewGlpBalanceTransactor creates a new write-only instance of GlpBalance, bound to a specific deployed contract.
func NewGlpBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GlpBalanceTransactor, error) {
	contract, err := bindGlpBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlpBalanceTransactor{contract: contract}, nil
}

// NewGlpBalanceFilterer creates a new log filterer instance of GlpBalance, bound to a specific deployed contract.
func NewGlpBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GlpBalanceFilterer, error) {
	contract, err := bindGlpBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlpBalanceFilterer{contract: contract}, nil
}

// bindGlpBalance binds a generic wrapper to an already deployed contract.
func bindGlpBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlpBalanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlpBalance *GlpBalanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlpBalance.Contract.GlpBalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlpBalance *GlpBalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlpBalance.Contract.GlpBalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlpBalance *GlpBalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlpBalance.Contract.GlpBalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlpBalance *GlpBalanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlpBalance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlpBalance *GlpBalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlpBalance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlpBalance *GlpBalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlpBalance.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GlpBalance *GlpBalanceCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlpBalance.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GlpBalance *GlpBalanceSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GlpBalance.Contract.Allowance(&_GlpBalance.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GlpBalance *GlpBalanceCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GlpBalance.Contract.Allowance(&_GlpBalance.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GlpBalance *GlpBalanceCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlpBalance.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GlpBalance *GlpBalanceSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GlpBalance.Contract.Allowances(&_GlpBalance.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GlpBalance *GlpBalanceCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GlpBalance.Contract.Allowances(&_GlpBalance.CallOpts, arg0, arg1)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_GlpBalance *GlpBalanceCaller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlpBalance.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_GlpBalance *GlpBalanceSession) GlpManager() (common.Address, error) {
	return _GlpBalance.Contract.GlpManager(&_GlpBalance.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_GlpBalance *GlpBalanceCallerSession) GlpManager() (common.Address, error) {
	return _GlpBalance.Contract.GlpManager(&_GlpBalance.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_GlpBalance *GlpBalanceCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlpBalance.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_GlpBalance *GlpBalanceSession) StakedGlpTracker() (common.Address, error) {
	return _GlpBalance.Contract.StakedGlpTracker(&_GlpBalance.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_GlpBalance *GlpBalanceCallerSession) StakedGlpTracker() (common.Address, error) {
	return _GlpBalance.Contract.StakedGlpTracker(&_GlpBalance.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.Contract.Approve(&_GlpBalance.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.Contract.Approve(&_GlpBalance.TransactOpts, _spender, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.Contract.Transfer(&_GlpBalance.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.Contract.Transfer(&_GlpBalance.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.Contract.TransferFrom(&_GlpBalance.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GlpBalance *GlpBalanceTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GlpBalance.Contract.TransferFrom(&_GlpBalance.TransactOpts, _sender, _recipient, _amount)
}

// GlpBalanceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GlpBalance contract.
type GlpBalanceApprovalIterator struct {
	Event *GlpBalanceApproval // Event containing the contract specifics and raw log

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
func (it *GlpBalanceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlpBalanceApproval)
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
		it.Event = new(GlpBalanceApproval)
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
func (it *GlpBalanceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlpBalanceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlpBalanceApproval represents a Approval event raised by the GlpBalance contract.
type GlpBalanceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GlpBalance *GlpBalanceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GlpBalanceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GlpBalance.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GlpBalanceApprovalIterator{contract: _GlpBalance.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GlpBalance *GlpBalanceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GlpBalanceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GlpBalance.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlpBalanceApproval)
				if err := _GlpBalance.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GlpBalance *GlpBalanceFilterer) ParseApproval(log types.Log) (*GlpBalanceApproval, error) {
	event := new(GlpBalanceApproval)
	if err := _GlpBalance.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
