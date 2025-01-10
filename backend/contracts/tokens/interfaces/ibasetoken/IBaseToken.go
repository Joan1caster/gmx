// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibasetoken

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

// IBaseTokenMetaData contains all meta data concerning the IBaseToken contract.
var IBaseTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateTransferMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBaseTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IBaseTokenMetaData.ABI instead.
var IBaseTokenABI = IBaseTokenMetaData.ABI

// IBaseToken is an auto generated Go binding around an Ethereum contract.
type IBaseToken struct {
	IBaseTokenCaller     // Read-only binding to the contract
	IBaseTokenTransactor // Write-only binding to the contract
	IBaseTokenFilterer   // Log filterer for contract events
}

// IBaseTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBaseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBaseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBaseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBaseTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBaseTokenSession struct {
	Contract     *IBaseToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBaseTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBaseTokenCallerSession struct {
	Contract *IBaseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IBaseTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBaseTokenTransactorSession struct {
	Contract     *IBaseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IBaseTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBaseTokenRaw struct {
	Contract *IBaseToken // Generic contract binding to access the raw methods on
}

// IBaseTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBaseTokenCallerRaw struct {
	Contract *IBaseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IBaseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBaseTokenTransactorRaw struct {
	Contract *IBaseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBaseToken creates a new instance of IBaseToken, bound to a specific deployed contract.
func NewIBaseToken(address common.Address, backend bind.ContractBackend) (*IBaseToken, error) {
	contract, err := bindIBaseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBaseToken{IBaseTokenCaller: IBaseTokenCaller{contract: contract}, IBaseTokenTransactor: IBaseTokenTransactor{contract: contract}, IBaseTokenFilterer: IBaseTokenFilterer{contract: contract}}, nil
}

// NewIBaseTokenCaller creates a new read-only instance of IBaseToken, bound to a specific deployed contract.
func NewIBaseTokenCaller(address common.Address, caller bind.ContractCaller) (*IBaseTokenCaller, error) {
	contract, err := bindIBaseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBaseTokenCaller{contract: contract}, nil
}

// NewIBaseTokenTransactor creates a new write-only instance of IBaseToken, bound to a specific deployed contract.
func NewIBaseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IBaseTokenTransactor, error) {
	contract, err := bindIBaseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBaseTokenTransactor{contract: contract}, nil
}

// NewIBaseTokenFilterer creates a new log filterer instance of IBaseToken, bound to a specific deployed contract.
func NewIBaseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IBaseTokenFilterer, error) {
	contract, err := bindIBaseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBaseTokenFilterer{contract: contract}, nil
}

// bindIBaseToken binds a generic wrapper to an already deployed contract.
func bindIBaseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBaseTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBaseToken *IBaseTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBaseToken.Contract.IBaseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBaseToken *IBaseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBaseToken.Contract.IBaseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBaseToken *IBaseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBaseToken.Contract.IBaseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBaseToken *IBaseTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBaseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBaseToken *IBaseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBaseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBaseToken *IBaseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBaseToken.Contract.contract.Transact(opts, method, params...)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_IBaseToken *IBaseTokenCaller) StakedBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBaseToken.contract.Call(opts, &out, "stakedBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_IBaseToken *IBaseTokenSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _IBaseToken.Contract.StakedBalance(&_IBaseToken.CallOpts, _account)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_IBaseToken *IBaseTokenCallerSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _IBaseToken.Contract.StakedBalance(&_IBaseToken.CallOpts, _account)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IBaseToken *IBaseTokenCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBaseToken.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IBaseToken *IBaseTokenSession) TotalStaked() (*big.Int, error) {
	return _IBaseToken.Contract.TotalStaked(&_IBaseToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_IBaseToken *IBaseTokenCallerSession) TotalStaked() (*big.Int, error) {
	return _IBaseToken.Contract.TotalStaked(&_IBaseToken.CallOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_IBaseToken *IBaseTokenTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IBaseToken.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_IBaseToken *IBaseTokenSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _IBaseToken.Contract.RemoveAdmin(&_IBaseToken.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_IBaseToken *IBaseTokenTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _IBaseToken.Contract.RemoveAdmin(&_IBaseToken.TransactOpts, _account)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_IBaseToken *IBaseTokenTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _IBaseToken.contract.Transact(opts, "setInPrivateTransferMode", _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_IBaseToken *IBaseTokenSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _IBaseToken.Contract.SetInPrivateTransferMode(&_IBaseToken.TransactOpts, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_IBaseToken *IBaseTokenTransactorSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _IBaseToken.Contract.SetInPrivateTransferMode(&_IBaseToken.TransactOpts, _inPrivateTransferMode)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_IBaseToken *IBaseTokenTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBaseToken.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_IBaseToken *IBaseTokenSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBaseToken.Contract.WithdrawToken(&_IBaseToken.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_IBaseToken *IBaseTokenTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IBaseToken.Contract.WithdrawToken(&_IBaseToken.TransactOpts, _token, _account, _amount)
}
