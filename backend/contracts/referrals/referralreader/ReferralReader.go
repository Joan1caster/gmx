// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package referralreader

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

// ReferralReaderMetaData contains all meta data concerning the ReferralReader contract.
var ReferralReaderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIReferralStorage\",\"name\":\"_referralStorage\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_codes\",\"type\":\"bytes32[]\"}],\"name\":\"getCodeOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ReferralReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use ReferralReaderMetaData.ABI instead.
var ReferralReaderABI = ReferralReaderMetaData.ABI

// ReferralReader is an auto generated Go binding around an Ethereum contract.
type ReferralReader struct {
	ReferralReaderCaller     // Read-only binding to the contract
	ReferralReaderTransactor // Write-only binding to the contract
	ReferralReaderFilterer   // Log filterer for contract events
}

// ReferralReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReferralReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReferralReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReferralReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReferralReaderSession struct {
	Contract     *ReferralReader   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReferralReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReferralReaderCallerSession struct {
	Contract *ReferralReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReferralReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReferralReaderTransactorSession struct {
	Contract     *ReferralReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReferralReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReferralReaderRaw struct {
	Contract *ReferralReader // Generic contract binding to access the raw methods on
}

// ReferralReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReferralReaderCallerRaw struct {
	Contract *ReferralReaderCaller // Generic read-only contract binding to access the raw methods on
}

// ReferralReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReferralReaderTransactorRaw struct {
	Contract *ReferralReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReferralReader creates a new instance of ReferralReader, bound to a specific deployed contract.
func NewReferralReader(address common.Address, backend bind.ContractBackend) (*ReferralReader, error) {
	contract, err := bindReferralReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReferralReader{ReferralReaderCaller: ReferralReaderCaller{contract: contract}, ReferralReaderTransactor: ReferralReaderTransactor{contract: contract}, ReferralReaderFilterer: ReferralReaderFilterer{contract: contract}}, nil
}

// NewReferralReaderCaller creates a new read-only instance of ReferralReader, bound to a specific deployed contract.
func NewReferralReaderCaller(address common.Address, caller bind.ContractCaller) (*ReferralReaderCaller, error) {
	contract, err := bindReferralReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralReaderCaller{contract: contract}, nil
}

// NewReferralReaderTransactor creates a new write-only instance of ReferralReader, bound to a specific deployed contract.
func NewReferralReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*ReferralReaderTransactor, error) {
	contract, err := bindReferralReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralReaderTransactor{contract: contract}, nil
}

// NewReferralReaderFilterer creates a new log filterer instance of ReferralReader, bound to a specific deployed contract.
func NewReferralReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*ReferralReaderFilterer, error) {
	contract, err := bindReferralReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReferralReaderFilterer{contract: contract}, nil
}

// bindReferralReader binds a generic wrapper to an already deployed contract.
func bindReferralReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReferralReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferralReader *ReferralReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferralReader.Contract.ReferralReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferralReader *ReferralReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralReader.Contract.ReferralReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferralReader *ReferralReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferralReader.Contract.ReferralReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferralReader *ReferralReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferralReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferralReader *ReferralReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferralReader *ReferralReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferralReader.Contract.contract.Transact(opts, method, params...)
}

// GetCodeOwners is a free data retrieval call binding the contract method 0x0f6803c0.
//
// Solidity: function getCodeOwners(address _referralStorage, bytes32[] _codes) view returns(address[])
func (_ReferralReader *ReferralReaderCaller) GetCodeOwners(opts *bind.CallOpts, _referralStorage common.Address, _codes [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ReferralReader.contract.Call(opts, &out, "getCodeOwners", _referralStorage, _codes)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCodeOwners is a free data retrieval call binding the contract method 0x0f6803c0.
//
// Solidity: function getCodeOwners(address _referralStorage, bytes32[] _codes) view returns(address[])
func (_ReferralReader *ReferralReaderSession) GetCodeOwners(_referralStorage common.Address, _codes [][32]byte) ([]common.Address, error) {
	return _ReferralReader.Contract.GetCodeOwners(&_ReferralReader.CallOpts, _referralStorage, _codes)
}

// GetCodeOwners is a free data retrieval call binding the contract method 0x0f6803c0.
//
// Solidity: function getCodeOwners(address _referralStorage, bytes32[] _codes) view returns(address[])
func (_ReferralReader *ReferralReaderCallerSession) GetCodeOwners(_referralStorage common.Address, _codes [][32]byte) ([]common.Address, error) {
	return _ReferralReader.Contract.GetCodeOwners(&_ReferralReader.CallOpts, _referralStorage, _codes)
}
