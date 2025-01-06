// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockpyth

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

// PythStructsPrice is an auto generated low-level Go binding around an user-defined struct.
type PythStructsPrice struct {
	Price       int64
	Conf        uint64
	Expo        int32
	PublishTime *big.Int
}

// MockPythMetaData contains all meta data concerning the MockPyth contract.
var MockPythMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"exponents\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"conf\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"internalType\":\"structPythStructs.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"name\":\"getUpdateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"publishTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"int64\",\"name\":\"price\",\"type\":\"int64\"},{\"internalType\":\"int32\",\"name\":\"expo\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"publishTime\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"name\":\"updatePriceFeeds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MockPythABI is the input ABI used to generate the binding from.
// Deprecated: Use MockPythMetaData.ABI instead.
var MockPythABI = MockPythMetaData.ABI

// MockPyth is an auto generated Go binding around an Ethereum contract.
type MockPyth struct {
	MockPythCaller     // Read-only binding to the contract
	MockPythTransactor // Write-only binding to the contract
	MockPythFilterer   // Log filterer for contract events
}

// MockPythCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockPythCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPythTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockPythTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPythFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockPythFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPythSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockPythSession struct {
	Contract     *MockPyth         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockPythCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockPythCallerSession struct {
	Contract *MockPythCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MockPythTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockPythTransactorSession struct {
	Contract     *MockPythTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockPythRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockPythRaw struct {
	Contract *MockPyth // Generic contract binding to access the raw methods on
}

// MockPythCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockPythCallerRaw struct {
	Contract *MockPythCaller // Generic read-only contract binding to access the raw methods on
}

// MockPythTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockPythTransactorRaw struct {
	Contract *MockPythTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockPyth creates a new instance of MockPyth, bound to a specific deployed contract.
func NewMockPyth(address common.Address, backend bind.ContractBackend) (*MockPyth, error) {
	contract, err := bindMockPyth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockPyth{MockPythCaller: MockPythCaller{contract: contract}, MockPythTransactor: MockPythTransactor{contract: contract}, MockPythFilterer: MockPythFilterer{contract: contract}}, nil
}

// NewMockPythCaller creates a new read-only instance of MockPyth, bound to a specific deployed contract.
func NewMockPythCaller(address common.Address, caller bind.ContractCaller) (*MockPythCaller, error) {
	contract, err := bindMockPyth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockPythCaller{contract: contract}, nil
}

// NewMockPythTransactor creates a new write-only instance of MockPyth, bound to a specific deployed contract.
func NewMockPythTransactor(address common.Address, transactor bind.ContractTransactor) (*MockPythTransactor, error) {
	contract, err := bindMockPyth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockPythTransactor{contract: contract}, nil
}

// NewMockPythFilterer creates a new log filterer instance of MockPyth, bound to a specific deployed contract.
func NewMockPythFilterer(address common.Address, filterer bind.ContractFilterer) (*MockPythFilterer, error) {
	contract, err := bindMockPyth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockPythFilterer{contract: contract}, nil
}

// bindMockPyth binds a generic wrapper to an already deployed contract.
func bindMockPyth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockPythMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPyth *MockPythRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPyth.Contract.MockPythCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPyth *MockPythRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPyth.Contract.MockPythTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPyth *MockPythRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPyth.Contract.MockPythTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPyth *MockPythCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPyth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPyth *MockPythTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPyth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPyth *MockPythTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPyth.Contract.contract.Transact(opts, method, params...)
}

// Exponents is a free data retrieval call binding the contract method 0xbe0b0a6c.
//
// Solidity: function exponents(bytes32 ) view returns(int32)
func (_MockPyth *MockPythCaller) Exponents(opts *bind.CallOpts, arg0 [32]byte) (int32, error) {
	var out []interface{}
	err := _MockPyth.contract.Call(opts, &out, "exponents", arg0)

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// Exponents is a free data retrieval call binding the contract method 0xbe0b0a6c.
//
// Solidity: function exponents(bytes32 ) view returns(int32)
func (_MockPyth *MockPythSession) Exponents(arg0 [32]byte) (int32, error) {
	return _MockPyth.Contract.Exponents(&_MockPyth.CallOpts, arg0)
}

// Exponents is a free data retrieval call binding the contract method 0xbe0b0a6c.
//
// Solidity: function exponents(bytes32 ) view returns(int32)
func (_MockPyth *MockPythCallerSession) Exponents(arg0 [32]byte) (int32, error) {
	return _MockPyth.Contract.Exponents(&_MockPyth.CallOpts, arg0)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_MockPyth *MockPythCaller) GetPrice(opts *bind.CallOpts, id [32]byte) (PythStructsPrice, error) {
	var out []interface{}
	err := _MockPyth.contract.Call(opts, &out, "getPrice", id)

	if err != nil {
		return *new(PythStructsPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(PythStructsPrice)).(*PythStructsPrice)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_MockPyth *MockPythSession) GetPrice(id [32]byte) (PythStructsPrice, error) {
	return _MockPyth.Contract.GetPrice(&_MockPyth.CallOpts, id)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 id) view returns((int64,uint64,int32,uint256) price)
func (_MockPyth *MockPythCallerSession) GetPrice(id [32]byte) (PythStructsPrice, error) {
	return _MockPyth.Contract.GetPrice(&_MockPyth.CallOpts, id)
}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] ) pure returns(uint256 feeAmount)
func (_MockPyth *MockPythCaller) GetUpdateFee(opts *bind.CallOpts, arg0 [][]byte) (*big.Int, error) {
	var out []interface{}
	err := _MockPyth.contract.Call(opts, &out, "getUpdateFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] ) pure returns(uint256 feeAmount)
func (_MockPyth *MockPythSession) GetUpdateFee(arg0 [][]byte) (*big.Int, error) {
	return _MockPyth.Contract.GetUpdateFee(&_MockPyth.CallOpts, arg0)
}

// GetUpdateFee is a free data retrieval call binding the contract method 0xd47eed45.
//
// Solidity: function getUpdateFee(bytes[] ) pure returns(uint256 feeAmount)
func (_MockPyth *MockPythCallerSession) GetUpdateFee(arg0 [][]byte) (*big.Int, error) {
	return _MockPyth.Contract.GetUpdateFee(&_MockPyth.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(int64)
func (_MockPyth *MockPythCaller) Prices(opts *bind.CallOpts, arg0 [32]byte) (int64, error) {
	var out []interface{}
	err := _MockPyth.contract.Call(opts, &out, "prices", arg0)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(int64)
func (_MockPyth *MockPythSession) Prices(arg0 [32]byte) (int64, error) {
	return _MockPyth.Contract.Prices(&_MockPyth.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(int64)
func (_MockPyth *MockPythCallerSession) Prices(arg0 [32]byte) (int64, error) {
	return _MockPyth.Contract.Prices(&_MockPyth.CallOpts, arg0)
}

// PublishTimes is a free data retrieval call binding the contract method 0xf8c7a5ce.
//
// Solidity: function publishTimes(bytes32 ) view returns(uint256)
func (_MockPyth *MockPythCaller) PublishTimes(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MockPyth.contract.Call(opts, &out, "publishTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublishTimes is a free data retrieval call binding the contract method 0xf8c7a5ce.
//
// Solidity: function publishTimes(bytes32 ) view returns(uint256)
func (_MockPyth *MockPythSession) PublishTimes(arg0 [32]byte) (*big.Int, error) {
	return _MockPyth.Contract.PublishTimes(&_MockPyth.CallOpts, arg0)
}

// PublishTimes is a free data retrieval call binding the contract method 0xf8c7a5ce.
//
// Solidity: function publishTimes(bytes32 ) view returns(uint256)
func (_MockPyth *MockPythCallerSession) PublishTimes(arg0 [32]byte) (*big.Int, error) {
	return _MockPyth.Contract.PublishTimes(&_MockPyth.CallOpts, arg0)
}

// SetPrice is a paid mutator transaction binding the contract method 0x2413d378.
//
// Solidity: function setPrice(bytes32 id, int64 price, int32 expo, uint256 publishTime) returns()
func (_MockPyth *MockPythTransactor) SetPrice(opts *bind.TransactOpts, id [32]byte, price int64, expo int32, publishTime *big.Int) (*types.Transaction, error) {
	return _MockPyth.contract.Transact(opts, "setPrice", id, price, expo, publishTime)
}

// SetPrice is a paid mutator transaction binding the contract method 0x2413d378.
//
// Solidity: function setPrice(bytes32 id, int64 price, int32 expo, uint256 publishTime) returns()
func (_MockPyth *MockPythSession) SetPrice(id [32]byte, price int64, expo int32, publishTime *big.Int) (*types.Transaction, error) {
	return _MockPyth.Contract.SetPrice(&_MockPyth.TransactOpts, id, price, expo, publishTime)
}

// SetPrice is a paid mutator transaction binding the contract method 0x2413d378.
//
// Solidity: function setPrice(bytes32 id, int64 price, int32 expo, uint256 publishTime) returns()
func (_MockPyth *MockPythTransactorSession) SetPrice(id [32]byte, price int64, expo int32, publishTime *big.Int) (*types.Transaction, error) {
	return _MockPyth.Contract.SetPrice(&_MockPyth.TransactOpts, id, price, expo, publishTime)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] ) payable returns()
func (_MockPyth *MockPythTransactor) UpdatePriceFeeds(opts *bind.TransactOpts, arg0 [][]byte) (*types.Transaction, error) {
	return _MockPyth.contract.Transact(opts, "updatePriceFeeds", arg0)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] ) payable returns()
func (_MockPyth *MockPythSession) UpdatePriceFeeds(arg0 [][]byte) (*types.Transaction, error) {
	return _MockPyth.Contract.UpdatePriceFeeds(&_MockPyth.TransactOpts, arg0)
}

// UpdatePriceFeeds is a paid mutator transaction binding the contract method 0xef9e5e28.
//
// Solidity: function updatePriceFeeds(bytes[] ) payable returns()
func (_MockPyth *MockPythTransactorSession) UpdatePriceFeeds(arg0 [][]byte) (*types.Transaction, error) {
	return _MockPyth.Contract.UpdatePriceFeeds(&_MockPyth.TransactOpts, arg0)
}
