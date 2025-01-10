// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipositionrouter

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

// IPositionRouterMetaData contains all meta data concerning the IPositionRouter contract.
var IPositionRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"decreasePositionRequestKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decreasePositionRequestKeysStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_executionFeeReceiver\",\"type\":\"address\"}],\"name\":\"executeDecreasePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_executionFeeReceiver\",\"type\":\"address\"}],\"name\":\"executeIncreasePositions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getDecreasePositionRequestPath\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getIncreasePositionRequestPath\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestQueueLengths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"increasePositionRequestKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increasePositionRequestKeysStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPositionRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IPositionRouterMetaData.ABI instead.
var IPositionRouterABI = IPositionRouterMetaData.ABI

// IPositionRouter is an auto generated Go binding around an Ethereum contract.
type IPositionRouter struct {
	IPositionRouterCaller     // Read-only binding to the contract
	IPositionRouterTransactor // Write-only binding to the contract
	IPositionRouterFilterer   // Log filterer for contract events
}

// IPositionRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPositionRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPositionRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPositionRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPositionRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPositionRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPositionRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPositionRouterSession struct {
	Contract     *IPositionRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPositionRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPositionRouterCallerSession struct {
	Contract *IPositionRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IPositionRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPositionRouterTransactorSession struct {
	Contract     *IPositionRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IPositionRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPositionRouterRaw struct {
	Contract *IPositionRouter // Generic contract binding to access the raw methods on
}

// IPositionRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPositionRouterCallerRaw struct {
	Contract *IPositionRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IPositionRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPositionRouterTransactorRaw struct {
	Contract *IPositionRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPositionRouter creates a new instance of IPositionRouter, bound to a specific deployed contract.
func NewIPositionRouter(address common.Address, backend bind.ContractBackend) (*IPositionRouter, error) {
	contract, err := bindIPositionRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPositionRouter{IPositionRouterCaller: IPositionRouterCaller{contract: contract}, IPositionRouterTransactor: IPositionRouterTransactor{contract: contract}, IPositionRouterFilterer: IPositionRouterFilterer{contract: contract}}, nil
}

// NewIPositionRouterCaller creates a new read-only instance of IPositionRouter, bound to a specific deployed contract.
func NewIPositionRouterCaller(address common.Address, caller bind.ContractCaller) (*IPositionRouterCaller, error) {
	contract, err := bindIPositionRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPositionRouterCaller{contract: contract}, nil
}

// NewIPositionRouterTransactor creates a new write-only instance of IPositionRouter, bound to a specific deployed contract.
func NewIPositionRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPositionRouterTransactor, error) {
	contract, err := bindIPositionRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPositionRouterTransactor{contract: contract}, nil
}

// NewIPositionRouterFilterer creates a new log filterer instance of IPositionRouter, bound to a specific deployed contract.
func NewIPositionRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPositionRouterFilterer, error) {
	contract, err := bindIPositionRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPositionRouterFilterer{contract: contract}, nil
}

// bindIPositionRouter binds a generic wrapper to an already deployed contract.
func bindIPositionRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPositionRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPositionRouter *IPositionRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPositionRouter.Contract.IPositionRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPositionRouter *IPositionRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPositionRouter.Contract.IPositionRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPositionRouter *IPositionRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPositionRouter.Contract.IPositionRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPositionRouter *IPositionRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPositionRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPositionRouter *IPositionRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPositionRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPositionRouter *IPositionRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPositionRouter.Contract.contract.Transact(opts, method, params...)
}

// DecreasePositionRequestKeys is a free data retrieval call binding the contract method 0x4278555f.
//
// Solidity: function decreasePositionRequestKeys(uint256 index) view returns(bytes32)
func (_IPositionRouter *IPositionRouterCaller) DecreasePositionRequestKeys(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IPositionRouter.contract.Call(opts, &out, "decreasePositionRequestKeys", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DecreasePositionRequestKeys is a free data retrieval call binding the contract method 0x4278555f.
//
// Solidity: function decreasePositionRequestKeys(uint256 index) view returns(bytes32)
func (_IPositionRouter *IPositionRouterSession) DecreasePositionRequestKeys(index *big.Int) ([32]byte, error) {
	return _IPositionRouter.Contract.DecreasePositionRequestKeys(&_IPositionRouter.CallOpts, index)
}

// DecreasePositionRequestKeys is a free data retrieval call binding the contract method 0x4278555f.
//
// Solidity: function decreasePositionRequestKeys(uint256 index) view returns(bytes32)
func (_IPositionRouter *IPositionRouterCallerSession) DecreasePositionRequestKeys(index *big.Int) ([32]byte, error) {
	return _IPositionRouter.Contract.DecreasePositionRequestKeys(&_IPositionRouter.CallOpts, index)
}

// DecreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x1bca8cf0.
//
// Solidity: function decreasePositionRequestKeysStart() view returns(uint256)
func (_IPositionRouter *IPositionRouterCaller) DecreasePositionRequestKeysStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPositionRouter.contract.Call(opts, &out, "decreasePositionRequestKeysStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x1bca8cf0.
//
// Solidity: function decreasePositionRequestKeysStart() view returns(uint256)
func (_IPositionRouter *IPositionRouterSession) DecreasePositionRequestKeysStart() (*big.Int, error) {
	return _IPositionRouter.Contract.DecreasePositionRequestKeysStart(&_IPositionRouter.CallOpts)
}

// DecreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x1bca8cf0.
//
// Solidity: function decreasePositionRequestKeysStart() view returns(uint256)
func (_IPositionRouter *IPositionRouterCallerSession) DecreasePositionRequestKeysStart() (*big.Int, error) {
	return _IPositionRouter.Contract.DecreasePositionRequestKeysStart(&_IPositionRouter.CallOpts)
}

// GetDecreasePositionRequestPath is a free data retrieval call binding the contract method 0x5d5c22e8.
//
// Solidity: function getDecreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_IPositionRouter *IPositionRouterCaller) GetDecreasePositionRequestPath(opts *bind.CallOpts, _key [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _IPositionRouter.contract.Call(opts, &out, "getDecreasePositionRequestPath", _key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDecreasePositionRequestPath is a free data retrieval call binding the contract method 0x5d5c22e8.
//
// Solidity: function getDecreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_IPositionRouter *IPositionRouterSession) GetDecreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _IPositionRouter.Contract.GetDecreasePositionRequestPath(&_IPositionRouter.CallOpts, _key)
}

// GetDecreasePositionRequestPath is a free data retrieval call binding the contract method 0x5d5c22e8.
//
// Solidity: function getDecreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_IPositionRouter *IPositionRouterCallerSession) GetDecreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _IPositionRouter.Contract.GetDecreasePositionRequestPath(&_IPositionRouter.CallOpts, _key)
}

// GetIncreasePositionRequestPath is a free data retrieval call binding the contract method 0x95e9bbd7.
//
// Solidity: function getIncreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_IPositionRouter *IPositionRouterCaller) GetIncreasePositionRequestPath(opts *bind.CallOpts, _key [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _IPositionRouter.contract.Call(opts, &out, "getIncreasePositionRequestPath", _key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetIncreasePositionRequestPath is a free data retrieval call binding the contract method 0x95e9bbd7.
//
// Solidity: function getIncreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_IPositionRouter *IPositionRouterSession) GetIncreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _IPositionRouter.Contract.GetIncreasePositionRequestPath(&_IPositionRouter.CallOpts, _key)
}

// GetIncreasePositionRequestPath is a free data retrieval call binding the contract method 0x95e9bbd7.
//
// Solidity: function getIncreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_IPositionRouter *IPositionRouterCallerSession) GetIncreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _IPositionRouter.Contract.GetIncreasePositionRequestPath(&_IPositionRouter.CallOpts, _key)
}

// GetRequestQueueLengths is a free data retrieval call binding the contract method 0xf2cea6a5.
//
// Solidity: function getRequestQueueLengths() view returns(uint256, uint256, uint256, uint256)
func (_IPositionRouter *IPositionRouterCaller) GetRequestQueueLengths(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IPositionRouter.contract.Call(opts, &out, "getRequestQueueLengths")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetRequestQueueLengths is a free data retrieval call binding the contract method 0xf2cea6a5.
//
// Solidity: function getRequestQueueLengths() view returns(uint256, uint256, uint256, uint256)
func (_IPositionRouter *IPositionRouterSession) GetRequestQueueLengths() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPositionRouter.Contract.GetRequestQueueLengths(&_IPositionRouter.CallOpts)
}

// GetRequestQueueLengths is a free data retrieval call binding the contract method 0xf2cea6a5.
//
// Solidity: function getRequestQueueLengths() view returns(uint256, uint256, uint256, uint256)
func (_IPositionRouter *IPositionRouterCallerSession) GetRequestQueueLengths() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IPositionRouter.Contract.GetRequestQueueLengths(&_IPositionRouter.CallOpts)
}

// IncreasePositionRequestKeys is a free data retrieval call binding the contract method 0x04225954.
//
// Solidity: function increasePositionRequestKeys(uint256 index) view returns(bytes32)
func (_IPositionRouter *IPositionRouterCaller) IncreasePositionRequestKeys(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IPositionRouter.contract.Call(opts, &out, "increasePositionRequestKeys", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IncreasePositionRequestKeys is a free data retrieval call binding the contract method 0x04225954.
//
// Solidity: function increasePositionRequestKeys(uint256 index) view returns(bytes32)
func (_IPositionRouter *IPositionRouterSession) IncreasePositionRequestKeys(index *big.Int) ([32]byte, error) {
	return _IPositionRouter.Contract.IncreasePositionRequestKeys(&_IPositionRouter.CallOpts, index)
}

// IncreasePositionRequestKeys is a free data retrieval call binding the contract method 0x04225954.
//
// Solidity: function increasePositionRequestKeys(uint256 index) view returns(bytes32)
func (_IPositionRouter *IPositionRouterCallerSession) IncreasePositionRequestKeys(index *big.Int) ([32]byte, error) {
	return _IPositionRouter.Contract.IncreasePositionRequestKeys(&_IPositionRouter.CallOpts, index)
}

// IncreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x9b578620.
//
// Solidity: function increasePositionRequestKeysStart() view returns(uint256)
func (_IPositionRouter *IPositionRouterCaller) IncreasePositionRequestKeysStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPositionRouter.contract.Call(opts, &out, "increasePositionRequestKeysStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x9b578620.
//
// Solidity: function increasePositionRequestKeysStart() view returns(uint256)
func (_IPositionRouter *IPositionRouterSession) IncreasePositionRequestKeysStart() (*big.Int, error) {
	return _IPositionRouter.Contract.IncreasePositionRequestKeysStart(&_IPositionRouter.CallOpts)
}

// IncreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x9b578620.
//
// Solidity: function increasePositionRequestKeysStart() view returns(uint256)
func (_IPositionRouter *IPositionRouterCallerSession) IncreasePositionRequestKeysStart() (*big.Int, error) {
	return _IPositionRouter.Contract.IncreasePositionRequestKeysStart(&_IPositionRouter.CallOpts)
}

// ExecuteDecreasePositions is a paid mutator transaction binding the contract method 0xf3883d8b.
//
// Solidity: function executeDecreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_IPositionRouter *IPositionRouterTransactor) ExecuteDecreasePositions(opts *bind.TransactOpts, _endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _IPositionRouter.contract.Transact(opts, "executeDecreasePositions", _endIndex, _executionFeeReceiver)
}

// ExecuteDecreasePositions is a paid mutator transaction binding the contract method 0xf3883d8b.
//
// Solidity: function executeDecreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_IPositionRouter *IPositionRouterSession) ExecuteDecreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _IPositionRouter.Contract.ExecuteDecreasePositions(&_IPositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}

// ExecuteDecreasePositions is a paid mutator transaction binding the contract method 0xf3883d8b.
//
// Solidity: function executeDecreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_IPositionRouter *IPositionRouterTransactorSession) ExecuteDecreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _IPositionRouter.Contract.ExecuteDecreasePositions(&_IPositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}

// ExecuteIncreasePositions is a paid mutator transaction binding the contract method 0x9a208100.
//
// Solidity: function executeIncreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_IPositionRouter *IPositionRouterTransactor) ExecuteIncreasePositions(opts *bind.TransactOpts, _endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _IPositionRouter.contract.Transact(opts, "executeIncreasePositions", _endIndex, _executionFeeReceiver)
}

// ExecuteIncreasePositions is a paid mutator transaction binding the contract method 0x9a208100.
//
// Solidity: function executeIncreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_IPositionRouter *IPositionRouterSession) ExecuteIncreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _IPositionRouter.Contract.ExecuteIncreasePositions(&_IPositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}

// ExecuteIncreasePositions is a paid mutator transaction binding the contract method 0x9a208100.
//
// Solidity: function executeIncreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_IPositionRouter *IPositionRouterTransactorSession) ExecuteIncreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _IPositionRouter.Contract.ExecuteIncreasePositions(&_IPositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}
