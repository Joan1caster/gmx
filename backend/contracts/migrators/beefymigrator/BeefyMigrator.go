// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package beefymigrator

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

// BeefyMigratorMetaData contains all meta data concerning the BeefyMigrator contract.
var BeefyMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beefyTimelockCaller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"afterGovGranted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beefyTimelockCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedGovGrantedCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BeefyMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use BeefyMigratorMetaData.ABI instead.
var BeefyMigratorABI = BeefyMigratorMetaData.ABI

// BeefyMigrator is an auto generated Go binding around an Ethereum contract.
type BeefyMigrator struct {
	BeefyMigratorCaller     // Read-only binding to the contract
	BeefyMigratorTransactor // Write-only binding to the contract
	BeefyMigratorFilterer   // Log filterer for contract events
}

// BeefyMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeefyMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeefyMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeefyMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeefyMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeefyMigratorSession struct {
	Contract     *BeefyMigrator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeefyMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeefyMigratorCallerSession struct {
	Contract *BeefyMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BeefyMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeefyMigratorTransactorSession struct {
	Contract     *BeefyMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BeefyMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeefyMigratorRaw struct {
	Contract *BeefyMigrator // Generic contract binding to access the raw methods on
}

// BeefyMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeefyMigratorCallerRaw struct {
	Contract *BeefyMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// BeefyMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeefyMigratorTransactorRaw struct {
	Contract *BeefyMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeefyMigrator creates a new instance of BeefyMigrator, bound to a specific deployed contract.
func NewBeefyMigrator(address common.Address, backend bind.ContractBackend) (*BeefyMigrator, error) {
	contract, err := bindBeefyMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeefyMigrator{BeefyMigratorCaller: BeefyMigratorCaller{contract: contract}, BeefyMigratorTransactor: BeefyMigratorTransactor{contract: contract}, BeefyMigratorFilterer: BeefyMigratorFilterer{contract: contract}}, nil
}

// NewBeefyMigratorCaller creates a new read-only instance of BeefyMigrator, bound to a specific deployed contract.
func NewBeefyMigratorCaller(address common.Address, caller bind.ContractCaller) (*BeefyMigratorCaller, error) {
	contract, err := bindBeefyMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyMigratorCaller{contract: contract}, nil
}

// NewBeefyMigratorTransactor creates a new write-only instance of BeefyMigrator, bound to a specific deployed contract.
func NewBeefyMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*BeefyMigratorTransactor, error) {
	contract, err := bindBeefyMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeefyMigratorTransactor{contract: contract}, nil
}

// NewBeefyMigratorFilterer creates a new log filterer instance of BeefyMigrator, bound to a specific deployed contract.
func NewBeefyMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*BeefyMigratorFilterer, error) {
	contract, err := bindBeefyMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeefyMigratorFilterer{contract: contract}, nil
}

// bindBeefyMigrator binds a generic wrapper to an already deployed contract.
func bindBeefyMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeefyMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyMigrator *BeefyMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyMigrator.Contract.BeefyMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyMigrator *BeefyMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyMigrator.Contract.BeefyMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyMigrator *BeefyMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyMigrator.Contract.BeefyMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeefyMigrator *BeefyMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeefyMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeefyMigrator *BeefyMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeefyMigrator *BeefyMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeefyMigrator.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) Admin() (common.Address, error) {
	return _BeefyMigrator.Contract.Admin(&_BeefyMigrator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) Admin() (common.Address, error) {
	return _BeefyMigrator.Contract.Admin(&_BeefyMigrator.CallOpts)
}

// BeefyTimelockCaller is a free data retrieval call binding the contract method 0xdf53ff5e.
//
// Solidity: function beefyTimelockCaller() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) BeefyTimelockCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "beefyTimelockCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeefyTimelockCaller is a free data retrieval call binding the contract method 0xdf53ff5e.
//
// Solidity: function beefyTimelockCaller() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) BeefyTimelockCaller() (common.Address, error) {
	return _BeefyMigrator.Contract.BeefyTimelockCaller(&_BeefyMigrator.CallOpts)
}

// BeefyTimelockCaller is a free data retrieval call binding the contract method 0xdf53ff5e.
//
// Solidity: function beefyTimelockCaller() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) BeefyTimelockCaller() (common.Address, error) {
	return _BeefyMigrator.Contract.BeefyTimelockCaller(&_BeefyMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) BnGmx() (common.Address, error) {
	return _BeefyMigrator.Contract.BnGmx(&_BeefyMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) BnGmx() (common.Address, error) {
	return _BeefyMigrator.Contract.BnGmx(&_BeefyMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) BonusGmxTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.BonusGmxTracker(&_BeefyMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) BonusGmxTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.BonusGmxTracker(&_BeefyMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) EsGmx() (common.Address, error) {
	return _BeefyMigrator.Contract.EsGmx(&_BeefyMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) EsGmx() (common.Address, error) {
	return _BeefyMigrator.Contract.EsGmx(&_BeefyMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) ExpectedGovGrantedCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "expectedGovGrantedCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _BeefyMigrator.Contract.ExpectedGovGrantedCaller(&_BeefyMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _BeefyMigrator.Contract.ExpectedGovGrantedCaller(&_BeefyMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) FeeGlpTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.FeeGlpTracker(&_BeefyMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) FeeGlpTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.FeeGlpTracker(&_BeefyMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) FeeGmxTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.FeeGmxTracker(&_BeefyMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) FeeGmxTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.FeeGmxTracker(&_BeefyMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) GlpVester() (common.Address, error) {
	return _BeefyMigrator.Contract.GlpVester(&_BeefyMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) GlpVester() (common.Address, error) {
	return _BeefyMigrator.Contract.GlpVester(&_BeefyMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) GmxVester() (common.Address, error) {
	return _BeefyMigrator.Contract.GmxVester(&_BeefyMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) GmxVester() (common.Address, error) {
	return _BeefyMigrator.Contract.GmxVester(&_BeefyMigrator.CallOpts)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) RewardRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "rewardRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) RewardRouter() (common.Address, error) {
	return _BeefyMigrator.Contract.RewardRouter(&_BeefyMigrator.CallOpts)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) RewardRouter() (common.Address, error) {
	return _BeefyMigrator.Contract.RewardRouter(&_BeefyMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) StakedGlpTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.StakedGlpTracker(&_BeefyMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) StakedGlpTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.StakedGlpTracker(&_BeefyMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCaller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeefyMigrator.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorSession) StakedGmxTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.StakedGmxTracker(&_BeefyMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BeefyMigrator *BeefyMigratorCallerSession) StakedGmxTracker() (common.Address, error) {
	return _BeefyMigrator.Contract.StakedGmxTracker(&_BeefyMigrator.CallOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BeefyMigrator *BeefyMigratorTransactor) AfterGovGranted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyMigrator.contract.Transact(opts, "afterGovGranted")
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BeefyMigrator *BeefyMigratorSession) AfterGovGranted() (*types.Transaction, error) {
	return _BeefyMigrator.Contract.AfterGovGranted(&_BeefyMigrator.TransactOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BeefyMigrator *BeefyMigratorTransactorSession) AfterGovGranted() (*types.Transaction, error) {
	return _BeefyMigrator.Contract.AfterGovGranted(&_BeefyMigrator.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_BeefyMigrator *BeefyMigratorTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeefyMigrator.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_BeefyMigrator *BeefyMigratorSession) Migrate() (*types.Transaction, error) {
	return _BeefyMigrator.Contract.Migrate(&_BeefyMigrator.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_BeefyMigrator *BeefyMigratorTransactorSession) Migrate() (*types.Transaction, error) {
	return _BeefyMigrator.Contract.Migrate(&_BeefyMigrator.TransactOpts)
}
