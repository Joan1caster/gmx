// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irouter

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

// IRouterMetaData contains all meta data concerning the IRouter contract.
var IRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_plugin\",\"type\":\"address\"}],\"name\":\"addPlugin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"pluginDecreasePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"pluginIncreasePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pluginTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IRouterMetaData.ABI instead.
var IRouterABI = IRouterMetaData.ABI

// IRouter is an auto generated Go binding around an Ethereum contract.
type IRouter struct {
	IRouterCaller     // Read-only binding to the contract
	IRouterTransactor // Write-only binding to the contract
	IRouterFilterer   // Log filterer for contract events
}

// IRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRouterSession struct {
	Contract     *IRouter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRouterCallerSession struct {
	Contract *IRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRouterTransactorSession struct {
	Contract     *IRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRouterRaw struct {
	Contract *IRouter // Generic contract binding to access the raw methods on
}

// IRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRouterCallerRaw struct {
	Contract *IRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRouterTransactorRaw struct {
	Contract *IRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRouter creates a new instance of IRouter, bound to a specific deployed contract.
func NewIRouter(address common.Address, backend bind.ContractBackend) (*IRouter, error) {
	contract, err := bindIRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRouter{IRouterCaller: IRouterCaller{contract: contract}, IRouterTransactor: IRouterTransactor{contract: contract}, IRouterFilterer: IRouterFilterer{contract: contract}}, nil
}

// NewIRouterCaller creates a new read-only instance of IRouter, bound to a specific deployed contract.
func NewIRouterCaller(address common.Address, caller bind.ContractCaller) (*IRouterCaller, error) {
	contract, err := bindIRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRouterCaller{contract: contract}, nil
}

// NewIRouterTransactor creates a new write-only instance of IRouter, bound to a specific deployed contract.
func NewIRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IRouterTransactor, error) {
	contract, err := bindIRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRouterTransactor{contract: contract}, nil
}

// NewIRouterFilterer creates a new log filterer instance of IRouter, bound to a specific deployed contract.
func NewIRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IRouterFilterer, error) {
	contract, err := bindIRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRouterFilterer{contract: contract}, nil
}

// bindIRouter binds a generic wrapper to an already deployed contract.
func bindIRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouter *IRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRouter.Contract.IRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouter *IRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRouter.Contract.IRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRouter *IRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRouter.Contract.IRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRouter *IRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRouter *IRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRouter *IRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRouter.Contract.contract.Transact(opts, method, params...)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_IRouter *IRouterTransactor) AddPlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _IRouter.contract.Transact(opts, "addPlugin", _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_IRouter *IRouterSession) AddPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _IRouter.Contract.AddPlugin(&_IRouter.TransactOpts, _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_IRouter *IRouterTransactorSession) AddPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _IRouter.Contract.AddPlugin(&_IRouter.TransactOpts, _plugin)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IRouter *IRouterTransactor) PluginDecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IRouter.contract.Transact(opts, "pluginDecreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IRouter *IRouterSession) PluginDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IRouter.Contract.PluginDecreasePosition(&_IRouter.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_IRouter *IRouterTransactorSession) PluginDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _IRouter.Contract.PluginDecreasePosition(&_IRouter.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IRouter *IRouterTransactor) PluginIncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IRouter.contract.Transact(opts, "pluginIncreasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IRouter *IRouterSession) PluginIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IRouter.Contract.PluginIncreasePosition(&_IRouter.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_IRouter *IRouterTransactorSession) PluginIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _IRouter.Contract.PluginIncreasePosition(&_IRouter.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_IRouter *IRouterTransactor) PluginTransfer(opts *bind.TransactOpts, _token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRouter.contract.Transact(opts, "pluginTransfer", _token, _account, _receiver, _amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_IRouter *IRouterSession) PluginTransfer(_token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRouter.Contract.PluginTransfer(&_IRouter.TransactOpts, _token, _account, _receiver, _amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_IRouter *IRouterTransactorSession) PluginTransfer(_token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRouter.Contract.PluginTransfer(&_IRouter.TransactOpts, _token, _account, _receiver, _amount)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IRouter *IRouterTransactor) Swap(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IRouter.contract.Transact(opts, "swap", _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IRouter *IRouterSession) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IRouter.Contract.Swap(&_IRouter.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_IRouter *IRouterTransactorSession) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IRouter.Contract.Swap(&_IRouter.TransactOpts, _path, _amountIn, _minOut, _receiver)
}
