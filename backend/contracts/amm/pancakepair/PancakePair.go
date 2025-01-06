// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakepair

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

// PancakePairMetaData contains all meta data concerning the PancakePair contract.
var PancakePairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance1\",\"type\":\"uint256\"}],\"name\":\"setReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PancakePairABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakePairMetaData.ABI instead.
var PancakePairABI = PancakePairMetaData.ABI

// PancakePair is an auto generated Go binding around an Ethereum contract.
type PancakePair struct {
	PancakePairCaller     // Read-only binding to the contract
	PancakePairTransactor // Write-only binding to the contract
	PancakePairFilterer   // Log filterer for contract events
}

// PancakePairCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakePairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakePairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakePairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakePairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakePairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakePairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakePairSession struct {
	Contract     *PancakePair      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakePairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakePairCallerSession struct {
	Contract *PancakePairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PancakePairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakePairTransactorSession struct {
	Contract     *PancakePairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PancakePairRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakePairRaw struct {
	Contract *PancakePair // Generic contract binding to access the raw methods on
}

// PancakePairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakePairCallerRaw struct {
	Contract *PancakePairCaller // Generic read-only contract binding to access the raw methods on
}

// PancakePairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakePairTransactorRaw struct {
	Contract *PancakePairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakePair creates a new instance of PancakePair, bound to a specific deployed contract.
func NewPancakePair(address common.Address, backend bind.ContractBackend) (*PancakePair, error) {
	contract, err := bindPancakePair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakePair{PancakePairCaller: PancakePairCaller{contract: contract}, PancakePairTransactor: PancakePairTransactor{contract: contract}, PancakePairFilterer: PancakePairFilterer{contract: contract}}, nil
}

// NewPancakePairCaller creates a new read-only instance of PancakePair, bound to a specific deployed contract.
func NewPancakePairCaller(address common.Address, caller bind.ContractCaller) (*PancakePairCaller, error) {
	contract, err := bindPancakePair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakePairCaller{contract: contract}, nil
}

// NewPancakePairTransactor creates a new write-only instance of PancakePair, bound to a specific deployed contract.
func NewPancakePairTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakePairTransactor, error) {
	contract, err := bindPancakePair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakePairTransactor{contract: contract}, nil
}

// NewPancakePairFilterer creates a new log filterer instance of PancakePair, bound to a specific deployed contract.
func NewPancakePairFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakePairFilterer, error) {
	contract, err := bindPancakePair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakePairFilterer{contract: contract}, nil
}

// bindPancakePair binds a generic wrapper to an already deployed contract.
func bindPancakePair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakePairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakePair *PancakePairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakePair.Contract.PancakePairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakePair *PancakePairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePair.Contract.PancakePairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakePair *PancakePairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakePair.Contract.PancakePairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakePair *PancakePairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakePair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakePair *PancakePairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakePair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakePair *PancakePairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakePair.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PancakePair *PancakePairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _PancakePair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PancakePair *PancakePairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PancakePair.Contract.GetReserves(&_PancakePair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_PancakePair *PancakePairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _PancakePair.Contract.GetReserves(&_PancakePair.CallOpts)
}

// SetReserves is a paid mutator transaction binding the contract method 0x8392b8c0.
//
// Solidity: function setReserves(uint256 balance0, uint256 balance1) returns()
func (_PancakePair *PancakePairTransactor) SetReserves(opts *bind.TransactOpts, balance0 *big.Int, balance1 *big.Int) (*types.Transaction, error) {
	return _PancakePair.contract.Transact(opts, "setReserves", balance0, balance1)
}

// SetReserves is a paid mutator transaction binding the contract method 0x8392b8c0.
//
// Solidity: function setReserves(uint256 balance0, uint256 balance1) returns()
func (_PancakePair *PancakePairSession) SetReserves(balance0 *big.Int, balance1 *big.Int) (*types.Transaction, error) {
	return _PancakePair.Contract.SetReserves(&_PancakePair.TransactOpts, balance0, balance1)
}

// SetReserves is a paid mutator transaction binding the contract method 0x8392b8c0.
//
// Solidity: function setReserves(uint256 balance0, uint256 balance1) returns()
func (_PancakePair *PancakePairTransactorSession) SetReserves(balance0 *big.Int, balance1 *big.Int) (*types.Transaction, error) {
	return _PancakePair.Contract.SetReserves(&_PancakePair.TransactOpts, balance0, balance1)
}
