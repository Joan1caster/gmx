// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakefactory

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

// PancakeFactoryMetaData contains all meta data concerning the PancakeFactory contract.
var PancakeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bnb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnbBusdPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btcBnbPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"busd\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PancakeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeFactoryMetaData.ABI instead.
var PancakeFactoryABI = PancakeFactoryMetaData.ABI

// PancakeFactory is an auto generated Go binding around an Ethereum contract.
type PancakeFactory struct {
	PancakeFactoryCaller     // Read-only binding to the contract
	PancakeFactoryTransactor // Write-only binding to the contract
	PancakeFactoryFilterer   // Log filterer for contract events
}

// PancakeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeFactorySession struct {
	Contract     *PancakeFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeFactoryCallerSession struct {
	Contract *PancakeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PancakeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeFactoryTransactorSession struct {
	Contract     *PancakeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PancakeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeFactoryRaw struct {
	Contract *PancakeFactory // Generic contract binding to access the raw methods on
}

// PancakeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeFactoryCallerRaw struct {
	Contract *PancakeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeFactoryTransactorRaw struct {
	Contract *PancakeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeFactory creates a new instance of PancakeFactory, bound to a specific deployed contract.
func NewPancakeFactory(address common.Address, backend bind.ContractBackend) (*PancakeFactory, error) {
	contract, err := bindPancakeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeFactory{PancakeFactoryCaller: PancakeFactoryCaller{contract: contract}, PancakeFactoryTransactor: PancakeFactoryTransactor{contract: contract}, PancakeFactoryFilterer: PancakeFactoryFilterer{contract: contract}}, nil
}

// NewPancakeFactoryCaller creates a new read-only instance of PancakeFactory, bound to a specific deployed contract.
func NewPancakeFactoryCaller(address common.Address, caller bind.ContractCaller) (*PancakeFactoryCaller, error) {
	contract, err := bindPancakeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeFactoryCaller{contract: contract}, nil
}

// NewPancakeFactoryTransactor creates a new write-only instance of PancakeFactory, bound to a specific deployed contract.
func NewPancakeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeFactoryTransactor, error) {
	contract, err := bindPancakeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeFactoryTransactor{contract: contract}, nil
}

// NewPancakeFactoryFilterer creates a new log filterer instance of PancakeFactory, bound to a specific deployed contract.
func NewPancakeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeFactoryFilterer, error) {
	contract, err := bindPancakeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeFactoryFilterer{contract: contract}, nil
}

// bindPancakeFactory binds a generic wrapper to an already deployed contract.
func bindPancakeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeFactory *PancakeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeFactory.Contract.PancakeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeFactory *PancakeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeFactory.Contract.PancakeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeFactory *PancakeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeFactory.Contract.PancakeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeFactory *PancakeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeFactory *PancakeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeFactory *PancakeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeFactory.Contract.contract.Transact(opts, method, params...)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_PancakeFactory *PancakeFactoryCaller) Bnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFactory.contract.Call(opts, &out, "bnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_PancakeFactory *PancakeFactorySession) Bnb() (common.Address, error) {
	return _PancakeFactory.Contract.Bnb(&_PancakeFactory.CallOpts)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_PancakeFactory *PancakeFactoryCallerSession) Bnb() (common.Address, error) {
	return _PancakeFactory.Contract.Bnb(&_PancakeFactory.CallOpts)
}

// BnbBusdPair is a free data retrieval call binding the contract method 0xfddc93ca.
//
// Solidity: function bnbBusdPair() view returns(address)
func (_PancakeFactory *PancakeFactoryCaller) BnbBusdPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFactory.contract.Call(opts, &out, "bnbBusdPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnbBusdPair is a free data retrieval call binding the contract method 0xfddc93ca.
//
// Solidity: function bnbBusdPair() view returns(address)
func (_PancakeFactory *PancakeFactorySession) BnbBusdPair() (common.Address, error) {
	return _PancakeFactory.Contract.BnbBusdPair(&_PancakeFactory.CallOpts)
}

// BnbBusdPair is a free data retrieval call binding the contract method 0xfddc93ca.
//
// Solidity: function bnbBusdPair() view returns(address)
func (_PancakeFactory *PancakeFactoryCallerSession) BnbBusdPair() (common.Address, error) {
	return _PancakeFactory.Contract.BnbBusdPair(&_PancakeFactory.CallOpts)
}

// Btc is a free data retrieval call binding the contract method 0xa28d57d8.
//
// Solidity: function btc() view returns(address)
func (_PancakeFactory *PancakeFactoryCaller) Btc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFactory.contract.Call(opts, &out, "btc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Btc is a free data retrieval call binding the contract method 0xa28d57d8.
//
// Solidity: function btc() view returns(address)
func (_PancakeFactory *PancakeFactorySession) Btc() (common.Address, error) {
	return _PancakeFactory.Contract.Btc(&_PancakeFactory.CallOpts)
}

// Btc is a free data retrieval call binding the contract method 0xa28d57d8.
//
// Solidity: function btc() view returns(address)
func (_PancakeFactory *PancakeFactoryCallerSession) Btc() (common.Address, error) {
	return _PancakeFactory.Contract.Btc(&_PancakeFactory.CallOpts)
}

// BtcBnbPair is a free data retrieval call binding the contract method 0xf4a68087.
//
// Solidity: function btcBnbPair() view returns(address)
func (_PancakeFactory *PancakeFactoryCaller) BtcBnbPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFactory.contract.Call(opts, &out, "btcBnbPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcBnbPair is a free data retrieval call binding the contract method 0xf4a68087.
//
// Solidity: function btcBnbPair() view returns(address)
func (_PancakeFactory *PancakeFactorySession) BtcBnbPair() (common.Address, error) {
	return _PancakeFactory.Contract.BtcBnbPair(&_PancakeFactory.CallOpts)
}

// BtcBnbPair is a free data retrieval call binding the contract method 0xf4a68087.
//
// Solidity: function btcBnbPair() view returns(address)
func (_PancakeFactory *PancakeFactoryCallerSession) BtcBnbPair() (common.Address, error) {
	return _PancakeFactory.Contract.BtcBnbPair(&_PancakeFactory.CallOpts)
}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_PancakeFactory *PancakeFactoryCaller) Busd(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFactory.contract.Call(opts, &out, "busd")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_PancakeFactory *PancakeFactorySession) Busd() (common.Address, error) {
	return _PancakeFactory.Contract.Busd(&_PancakeFactory.CallOpts)
}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_PancakeFactory *PancakeFactoryCallerSession) Busd() (common.Address, error) {
	return _PancakeFactory.Contract.Busd(&_PancakeFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address)
func (_PancakeFactory *PancakeFactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _PancakeFactory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address)
func (_PancakeFactory *PancakeFactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _PancakeFactory.Contract.GetPair(&_PancakeFactory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address)
func (_PancakeFactory *PancakeFactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _PancakeFactory.Contract.GetPair(&_PancakeFactory.CallOpts, tokenA, tokenB)
}
