// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package esgmxbatchsender

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

// EsGmxBatchSenderMetaData contains all meta data concerning the EsGmxBatchSender contract.
var EsGmxBatchSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVester\",\"name\":\"_vester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minRatio\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EsGmxBatchSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use EsGmxBatchSenderMetaData.ABI instead.
var EsGmxBatchSenderABI = EsGmxBatchSenderMetaData.ABI

// EsGmxBatchSender is an auto generated Go binding around an Ethereum contract.
type EsGmxBatchSender struct {
	EsGmxBatchSenderCaller     // Read-only binding to the contract
	EsGmxBatchSenderTransactor // Write-only binding to the contract
	EsGmxBatchSenderFilterer   // Log filterer for contract events
}

// EsGmxBatchSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type EsGmxBatchSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsGmxBatchSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EsGmxBatchSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsGmxBatchSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EsGmxBatchSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EsGmxBatchSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EsGmxBatchSenderSession struct {
	Contract     *EsGmxBatchSender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EsGmxBatchSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EsGmxBatchSenderCallerSession struct {
	Contract *EsGmxBatchSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EsGmxBatchSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EsGmxBatchSenderTransactorSession struct {
	Contract     *EsGmxBatchSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EsGmxBatchSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type EsGmxBatchSenderRaw struct {
	Contract *EsGmxBatchSender // Generic contract binding to access the raw methods on
}

// EsGmxBatchSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EsGmxBatchSenderCallerRaw struct {
	Contract *EsGmxBatchSenderCaller // Generic read-only contract binding to access the raw methods on
}

// EsGmxBatchSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EsGmxBatchSenderTransactorRaw struct {
	Contract *EsGmxBatchSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEsGmxBatchSender creates a new instance of EsGmxBatchSender, bound to a specific deployed contract.
func NewEsGmxBatchSender(address common.Address, backend bind.ContractBackend) (*EsGmxBatchSender, error) {
	contract, err := bindEsGmxBatchSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EsGmxBatchSender{EsGmxBatchSenderCaller: EsGmxBatchSenderCaller{contract: contract}, EsGmxBatchSenderTransactor: EsGmxBatchSenderTransactor{contract: contract}, EsGmxBatchSenderFilterer: EsGmxBatchSenderFilterer{contract: contract}}, nil
}

// NewEsGmxBatchSenderCaller creates a new read-only instance of EsGmxBatchSender, bound to a specific deployed contract.
func NewEsGmxBatchSenderCaller(address common.Address, caller bind.ContractCaller) (*EsGmxBatchSenderCaller, error) {
	contract, err := bindEsGmxBatchSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EsGmxBatchSenderCaller{contract: contract}, nil
}

// NewEsGmxBatchSenderTransactor creates a new write-only instance of EsGmxBatchSender, bound to a specific deployed contract.
func NewEsGmxBatchSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*EsGmxBatchSenderTransactor, error) {
	contract, err := bindEsGmxBatchSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EsGmxBatchSenderTransactor{contract: contract}, nil
}

// NewEsGmxBatchSenderFilterer creates a new log filterer instance of EsGmxBatchSender, bound to a specific deployed contract.
func NewEsGmxBatchSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*EsGmxBatchSenderFilterer, error) {
	contract, err := bindEsGmxBatchSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EsGmxBatchSenderFilterer{contract: contract}, nil
}

// bindEsGmxBatchSender binds a generic wrapper to an already deployed contract.
func bindEsGmxBatchSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EsGmxBatchSenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsGmxBatchSender *EsGmxBatchSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsGmxBatchSender.Contract.EsGmxBatchSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsGmxBatchSender *EsGmxBatchSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsGmxBatchSender.Contract.EsGmxBatchSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsGmxBatchSender *EsGmxBatchSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsGmxBatchSender.Contract.EsGmxBatchSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EsGmxBatchSender *EsGmxBatchSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EsGmxBatchSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EsGmxBatchSender *EsGmxBatchSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EsGmxBatchSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EsGmxBatchSender *EsGmxBatchSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EsGmxBatchSender.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_EsGmxBatchSender *EsGmxBatchSenderCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EsGmxBatchSender.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_EsGmxBatchSender *EsGmxBatchSenderSession) Admin() (common.Address, error) {
	return _EsGmxBatchSender.Contract.Admin(&_EsGmxBatchSender.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_EsGmxBatchSender *EsGmxBatchSenderCallerSession) Admin() (common.Address, error) {
	return _EsGmxBatchSender.Contract.Admin(&_EsGmxBatchSender.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_EsGmxBatchSender *EsGmxBatchSenderCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EsGmxBatchSender.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_EsGmxBatchSender *EsGmxBatchSenderSession) EsGmx() (common.Address, error) {
	return _EsGmxBatchSender.Contract.EsGmx(&_EsGmxBatchSender.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_EsGmxBatchSender *EsGmxBatchSenderCallerSession) EsGmx() (common.Address, error) {
	return _EsGmxBatchSender.Contract.EsGmx(&_EsGmxBatchSender.CallOpts)
}

// Send is a paid mutator transaction binding the contract method 0x8b8f837b.
//
// Solidity: function send(address _vester, uint256 _minRatio, address[] _accounts, uint256[] _amounts) returns()
func (_EsGmxBatchSender *EsGmxBatchSenderTransactor) Send(opts *bind.TransactOpts, _vester common.Address, _minRatio *big.Int, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _EsGmxBatchSender.contract.Transact(opts, "send", _vester, _minRatio, _accounts, _amounts)
}

// Send is a paid mutator transaction binding the contract method 0x8b8f837b.
//
// Solidity: function send(address _vester, uint256 _minRatio, address[] _accounts, uint256[] _amounts) returns()
func (_EsGmxBatchSender *EsGmxBatchSenderSession) Send(_vester common.Address, _minRatio *big.Int, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _EsGmxBatchSender.Contract.Send(&_EsGmxBatchSender.TransactOpts, _vester, _minRatio, _accounts, _amounts)
}

// Send is a paid mutator transaction binding the contract method 0x8b8f837b.
//
// Solidity: function send(address _vester, uint256 _minRatio, address[] _accounts, uint256[] _amounts) returns()
func (_EsGmxBatchSender *EsGmxBatchSenderTransactorSession) Send(_vester common.Address, _minRatio *big.Int, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _EsGmxBatchSender.Contract.Send(&_EsGmxBatchSender.TransactOpts, _vester, _minRatio, _accounts, _amounts)
}
