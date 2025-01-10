// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governable

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

// GovernableMetaData contains all meta data concerning the Governable contract.
var GovernableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GovernableABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernableMetaData.ABI instead.
var GovernableABI = GovernableMetaData.ABI

// Governable is an auto generated Go binding around an Ethereum contract.
type Governable struct {
	GovernableCaller     // Read-only binding to the contract
	GovernableTransactor // Write-only binding to the contract
	GovernableFilterer   // Log filterer for contract events
}

// GovernableCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernableSession struct {
	Contract     *Governable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernableCallerSession struct {
	Contract *GovernableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernableTransactorSession struct {
	Contract     *GovernableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernableRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernableRaw struct {
	Contract *Governable // Generic contract binding to access the raw methods on
}

// GovernableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernableCallerRaw struct {
	Contract *GovernableCaller // Generic read-only contract binding to access the raw methods on
}

// GovernableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernableTransactorRaw struct {
	Contract *GovernableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernable creates a new instance of Governable, bound to a specific deployed contract.
func NewGovernable(address common.Address, backend bind.ContractBackend) (*Governable, error) {
	contract, err := bindGovernable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governable{GovernableCaller: GovernableCaller{contract: contract}, GovernableTransactor: GovernableTransactor{contract: contract}, GovernableFilterer: GovernableFilterer{contract: contract}}, nil
}

// NewGovernableCaller creates a new read-only instance of Governable, bound to a specific deployed contract.
func NewGovernableCaller(address common.Address, caller bind.ContractCaller) (*GovernableCaller, error) {
	contract, err := bindGovernable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernableCaller{contract: contract}, nil
}

// NewGovernableTransactor creates a new write-only instance of Governable, bound to a specific deployed contract.
func NewGovernableTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernableTransactor, error) {
	contract, err := bindGovernable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernableTransactor{contract: contract}, nil
}

// NewGovernableFilterer creates a new log filterer instance of Governable, bound to a specific deployed contract.
func NewGovernableFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernableFilterer, error) {
	contract, err := bindGovernable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernableFilterer{contract: contract}, nil
}

// bindGovernable binds a generic wrapper to an already deployed contract.
func bindGovernable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovernableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governable *GovernableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governable.Contract.GovernableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governable *GovernableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governable.Contract.GovernableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governable *GovernableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governable.Contract.GovernableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governable *GovernableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governable *GovernableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governable *GovernableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governable.Contract.contract.Transact(opts, method, params...)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Governable *GovernableCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governable.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Governable *GovernableSession) Gov() (common.Address, error) {
	return _Governable.Contract.Gov(&_Governable.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Governable *GovernableCallerSession) Gov() (common.Address, error) {
	return _Governable.Contract.Gov(&_Governable.CallOpts)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Governable *GovernableTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Governable.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Governable *GovernableSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Governable.Contract.SetGov(&_Governable.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Governable *GovernableTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Governable.Contract.SetGov(&_Governable.TransactOpts, _gov)
}
