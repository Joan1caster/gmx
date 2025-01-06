// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iorderbook

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

// IOrderBookMetaData contains all meta data concerning the IOrderBook contract.
var IOrderBookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeDecreaseOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeIncreaseOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeSwapOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderIndex\",\"type\":\"uint256\"}],\"name\":\"getDecreaseOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"triggerAboveThreshold\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderIndex\",\"type\":\"uint256\"}],\"name\":\"getIncreaseOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"purchaseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"purchaseTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"triggerAboveThreshold\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_orderIndex\",\"type\":\"uint256\"}],\"name\":\"getSwapOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"path0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"path1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"path2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerRatio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"triggerAboveThreshold\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldUnwrap\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use IOrderBookMetaData.ABI instead.
var IOrderBookABI = IOrderBookMetaData.ABI

// IOrderBook is an auto generated Go binding around an Ethereum contract.
type IOrderBook struct {
	IOrderBookCaller     // Read-only binding to the contract
	IOrderBookTransactor // Write-only binding to the contract
	IOrderBookFilterer   // Log filterer for contract events
}

// IOrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOrderBookSession struct {
	Contract     *IOrderBook       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOrderBookCallerSession struct {
	Contract *IOrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IOrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOrderBookTransactorSession struct {
	Contract     *IOrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IOrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOrderBookRaw struct {
	Contract *IOrderBook // Generic contract binding to access the raw methods on
}

// IOrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOrderBookCallerRaw struct {
	Contract *IOrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// IOrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOrderBookTransactorRaw struct {
	Contract *IOrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOrderBook creates a new instance of IOrderBook, bound to a specific deployed contract.
func NewIOrderBook(address common.Address, backend bind.ContractBackend) (*IOrderBook, error) {
	contract, err := bindIOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOrderBook{IOrderBookCaller: IOrderBookCaller{contract: contract}, IOrderBookTransactor: IOrderBookTransactor{contract: contract}, IOrderBookFilterer: IOrderBookFilterer{contract: contract}}, nil
}

// NewIOrderBookCaller creates a new read-only instance of IOrderBook, bound to a specific deployed contract.
func NewIOrderBookCaller(address common.Address, caller bind.ContractCaller) (*IOrderBookCaller, error) {
	contract, err := bindIOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOrderBookCaller{contract: contract}, nil
}

// NewIOrderBookTransactor creates a new write-only instance of IOrderBook, bound to a specific deployed contract.
func NewIOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*IOrderBookTransactor, error) {
	contract, err := bindIOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOrderBookTransactor{contract: contract}, nil
}

// NewIOrderBookFilterer creates a new log filterer instance of IOrderBook, bound to a specific deployed contract.
func NewIOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*IOrderBookFilterer, error) {
	contract, err := bindIOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOrderBookFilterer{contract: contract}, nil
}

// bindIOrderBook binds a generic wrapper to an already deployed contract.
func bindIOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOrderBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrderBook *IOrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOrderBook.Contract.IOrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrderBook *IOrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrderBook.Contract.IOrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrderBook *IOrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrderBook.Contract.IOrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrderBook *IOrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrderBook *IOrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrderBook *IOrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrderBook.Contract.contract.Transact(opts, method, params...)
}

// GetDecreaseOrder is a free data retrieval call binding the contract method 0x026032ee.
//
// Solidity: function getDecreaseOrder(address _account, uint256 _orderIndex) view returns(address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_IOrderBook *IOrderBookCaller) GetDecreaseOrder(opts *bind.CallOpts, _account common.Address, _orderIndex *big.Int) (struct {
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _IOrderBook.contract.Call(opts, &out, "getDecreaseOrder", _account, _orderIndex)

	outstruct := new(struct {
		CollateralToken       common.Address
		CollateralDelta       *big.Int
		IndexToken            common.Address
		SizeDelta             *big.Int
		IsLong                bool
		TriggerPrice          *big.Int
		TriggerAboveThreshold bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CollateralDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IndexToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SizeDelta = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.TriggerPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDecreaseOrder is a free data retrieval call binding the contract method 0x026032ee.
//
// Solidity: function getDecreaseOrder(address _account, uint256 _orderIndex) view returns(address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_IOrderBook *IOrderBookSession) GetDecreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _IOrderBook.Contract.GetDecreaseOrder(&_IOrderBook.CallOpts, _account, _orderIndex)
}

// GetDecreaseOrder is a free data retrieval call binding the contract method 0x026032ee.
//
// Solidity: function getDecreaseOrder(address _account, uint256 _orderIndex) view returns(address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_IOrderBook *IOrderBookCallerSession) GetDecreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _IOrderBook.Contract.GetDecreaseOrder(&_IOrderBook.CallOpts, _account, _orderIndex)
}

// GetIncreaseOrder is a free data retrieval call binding the contract method 0xd3bab1d1.
//
// Solidity: function getIncreaseOrder(address _account, uint256 _orderIndex) view returns(address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_IOrderBook *IOrderBookCaller) GetIncreaseOrder(opts *bind.CallOpts, _account common.Address, _orderIndex *big.Int) (struct {
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _IOrderBook.contract.Call(opts, &out, "getIncreaseOrder", _account, _orderIndex)

	outstruct := new(struct {
		PurchaseToken         common.Address
		PurchaseTokenAmount   *big.Int
		CollateralToken       common.Address
		IndexToken            common.Address
		SizeDelta             *big.Int
		IsLong                bool
		TriggerPrice          *big.Int
		TriggerAboveThreshold bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PurchaseToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PurchaseTokenAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CollateralToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IndexToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.SizeDelta = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.TriggerPrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIncreaseOrder is a free data retrieval call binding the contract method 0xd3bab1d1.
//
// Solidity: function getIncreaseOrder(address _account, uint256 _orderIndex) view returns(address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_IOrderBook *IOrderBookSession) GetIncreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _IOrderBook.Contract.GetIncreaseOrder(&_IOrderBook.CallOpts, _account, _orderIndex)
}

// GetIncreaseOrder is a free data retrieval call binding the contract method 0xd3bab1d1.
//
// Solidity: function getIncreaseOrder(address _account, uint256 _orderIndex) view returns(address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_IOrderBook *IOrderBookCallerSession) GetIncreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _IOrderBook.Contract.GetIncreaseOrder(&_IOrderBook.CallOpts, _account, _orderIndex)
}

// GetSwapOrder is a free data retrieval call binding the contract method 0xd0d40cd6.
//
// Solidity: function getSwapOrder(address _account, uint256 _orderIndex) view returns(address path0, address path1, address path2, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_IOrderBook *IOrderBookCaller) GetSwapOrder(opts *bind.CallOpts, _account common.Address, _orderIndex *big.Int) (struct {
	Path0                 common.Address
	Path1                 common.Address
	Path2                 common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _IOrderBook.contract.Call(opts, &out, "getSwapOrder", _account, _orderIndex)

	outstruct := new(struct {
		Path0                 common.Address
		Path1                 common.Address
		Path2                 common.Address
		AmountIn              *big.Int
		MinOut                *big.Int
		TriggerRatio          *big.Int
		TriggerAboveThreshold bool
		ShouldUnwrap          bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Path0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Path1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Path2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AmountIn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinOut = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TriggerRatio = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.ShouldUnwrap = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapOrder is a free data retrieval call binding the contract method 0xd0d40cd6.
//
// Solidity: function getSwapOrder(address _account, uint256 _orderIndex) view returns(address path0, address path1, address path2, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_IOrderBook *IOrderBookSession) GetSwapOrder(_account common.Address, _orderIndex *big.Int) (struct {
	Path0                 common.Address
	Path1                 common.Address
	Path2                 common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	return _IOrderBook.Contract.GetSwapOrder(&_IOrderBook.CallOpts, _account, _orderIndex)
}

// GetSwapOrder is a free data retrieval call binding the contract method 0xd0d40cd6.
//
// Solidity: function getSwapOrder(address _account, uint256 _orderIndex) view returns(address path0, address path1, address path2, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_IOrderBook *IOrderBookCallerSession) GetSwapOrder(_account common.Address, _orderIndex *big.Int) (struct {
	Path0                 common.Address
	Path1                 common.Address
	Path2                 common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	return _IOrderBook.Contract.GetSwapOrder(&_IOrderBook.CallOpts, _account, _orderIndex)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookTransactor) ExecuteDecreaseOrder(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.contract.Transact(opts, "executeDecreaseOrder", arg0, arg1, arg2)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookSession) ExecuteDecreaseOrder(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.Contract.ExecuteDecreaseOrder(&_IOrderBook.TransactOpts, arg0, arg1, arg2)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookTransactorSession) ExecuteDecreaseOrder(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.Contract.ExecuteDecreaseOrder(&_IOrderBook.TransactOpts, arg0, arg1, arg2)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookTransactor) ExecuteIncreaseOrder(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.contract.Transact(opts, "executeIncreaseOrder", arg0, arg1, arg2)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookSession) ExecuteIncreaseOrder(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.Contract.ExecuteIncreaseOrder(&_IOrderBook.TransactOpts, arg0, arg1, arg2)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookTransactorSession) ExecuteIncreaseOrder(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.Contract.ExecuteIncreaseOrder(&_IOrderBook.TransactOpts, arg0, arg1, arg2)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookTransactor) ExecuteSwapOrder(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.contract.Transact(opts, "executeSwapOrder", arg0, arg1, arg2)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookSession) ExecuteSwapOrder(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.Contract.ExecuteSwapOrder(&_IOrderBook.TransactOpts, arg0, arg1, arg2)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address , uint256 , address ) returns()
func (_IOrderBook *IOrderBookTransactorSession) ExecuteSwapOrder(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _IOrderBook.Contract.ExecuteSwapOrder(&_IOrderBook.TransactOpts, arg0, arg1, arg2)
}
