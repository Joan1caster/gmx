// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package positionrouterreader

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

// PositionRouterReaderMetaData contains all meta data concerning the PositionRouterReader contract.
var PositionRouterReaderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_positionRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getTransferTokenOfIncreasePositionRequests\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PositionRouterReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionRouterReaderMetaData.ABI instead.
var PositionRouterReaderABI = PositionRouterReaderMetaData.ABI

// PositionRouterReader is an auto generated Go binding around an Ethereum contract.
type PositionRouterReader struct {
	PositionRouterReaderCaller     // Read-only binding to the contract
	PositionRouterReaderTransactor // Write-only binding to the contract
	PositionRouterReaderFilterer   // Log filterer for contract events
}

// PositionRouterReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionRouterReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionRouterReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionRouterReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionRouterReaderSession struct {
	Contract     *PositionRouterReader // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PositionRouterReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionRouterReaderCallerSession struct {
	Contract *PositionRouterReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PositionRouterReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionRouterReaderTransactorSession struct {
	Contract     *PositionRouterReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PositionRouterReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionRouterReaderRaw struct {
	Contract *PositionRouterReader // Generic contract binding to access the raw methods on
}

// PositionRouterReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionRouterReaderCallerRaw struct {
	Contract *PositionRouterReaderCaller // Generic read-only contract binding to access the raw methods on
}

// PositionRouterReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionRouterReaderTransactorRaw struct {
	Contract *PositionRouterReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionRouterReader creates a new instance of PositionRouterReader, bound to a specific deployed contract.
func NewPositionRouterReader(address common.Address, backend bind.ContractBackend) (*PositionRouterReader, error) {
	contract, err := bindPositionRouterReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionRouterReader{PositionRouterReaderCaller: PositionRouterReaderCaller{contract: contract}, PositionRouterReaderTransactor: PositionRouterReaderTransactor{contract: contract}, PositionRouterReaderFilterer: PositionRouterReaderFilterer{contract: contract}}, nil
}

// NewPositionRouterReaderCaller creates a new read-only instance of PositionRouterReader, bound to a specific deployed contract.
func NewPositionRouterReaderCaller(address common.Address, caller bind.ContractCaller) (*PositionRouterReaderCaller, error) {
	contract, err := bindPositionRouterReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionRouterReaderCaller{contract: contract}, nil
}

// NewPositionRouterReaderTransactor creates a new write-only instance of PositionRouterReader, bound to a specific deployed contract.
func NewPositionRouterReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionRouterReaderTransactor, error) {
	contract, err := bindPositionRouterReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionRouterReaderTransactor{contract: contract}, nil
}

// NewPositionRouterReaderFilterer creates a new log filterer instance of PositionRouterReader, bound to a specific deployed contract.
func NewPositionRouterReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionRouterReaderFilterer, error) {
	contract, err := bindPositionRouterReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionRouterReaderFilterer{contract: contract}, nil
}

// bindPositionRouterReader binds a generic wrapper to an already deployed contract.
func bindPositionRouterReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionRouterReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionRouterReader *PositionRouterReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionRouterReader.Contract.PositionRouterReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionRouterReader *PositionRouterReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionRouterReader.Contract.PositionRouterReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionRouterReader *PositionRouterReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionRouterReader.Contract.PositionRouterReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionRouterReader *PositionRouterReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionRouterReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionRouterReader *PositionRouterReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionRouterReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionRouterReader *PositionRouterReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionRouterReader.Contract.contract.Transact(opts, method, params...)
}

// GetTransferTokenOfIncreasePositionRequests is a free data retrieval call binding the contract method 0x0e0deb1a.
//
// Solidity: function getTransferTokenOfIncreasePositionRequests(address _positionRouter, uint256 _endIndex) view returns(uint256[], address[])
func (_PositionRouterReader *PositionRouterReaderCaller) GetTransferTokenOfIncreasePositionRequests(opts *bind.CallOpts, _positionRouter common.Address, _endIndex *big.Int) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _PositionRouterReader.contract.Call(opts, &out, "getTransferTokenOfIncreasePositionRequests", _positionRouter, _endIndex)

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetTransferTokenOfIncreasePositionRequests is a free data retrieval call binding the contract method 0x0e0deb1a.
//
// Solidity: function getTransferTokenOfIncreasePositionRequests(address _positionRouter, uint256 _endIndex) view returns(uint256[], address[])
func (_PositionRouterReader *PositionRouterReaderSession) GetTransferTokenOfIncreasePositionRequests(_positionRouter common.Address, _endIndex *big.Int) ([]*big.Int, []common.Address, error) {
	return _PositionRouterReader.Contract.GetTransferTokenOfIncreasePositionRequests(&_PositionRouterReader.CallOpts, _positionRouter, _endIndex)
}

// GetTransferTokenOfIncreasePositionRequests is a free data retrieval call binding the contract method 0x0e0deb1a.
//
// Solidity: function getTransferTokenOfIncreasePositionRequests(address _positionRouter, uint256 _endIndex) view returns(uint256[], address[])
func (_PositionRouterReader *PositionRouterReaderCallerSession) GetTransferTokenOfIncreasePositionRequests(_positionRouter common.Address, _endIndex *big.Int) ([]*big.Int, []common.Address, error) {
	return _PositionRouterReader.Contract.GetTransferTokenOfIncreasePositionRequests(&_PositionRouterReader.CallOpts, _positionRouter, _endIndex)
}
