// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stabilizemigrator

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

// StabilizeMigratorMetaData contains all meta data concerning the StabilizeMigrator contract.
var StabilizeMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stabilizeCaller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"afterGovGranted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedGovGrantedCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stabilizeCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StabilizeMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use StabilizeMigratorMetaData.ABI instead.
var StabilizeMigratorABI = StabilizeMigratorMetaData.ABI

// StabilizeMigrator is an auto generated Go binding around an Ethereum contract.
type StabilizeMigrator struct {
	StabilizeMigratorCaller     // Read-only binding to the contract
	StabilizeMigratorTransactor // Write-only binding to the contract
	StabilizeMigratorFilterer   // Log filterer for contract events
}

// StabilizeMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type StabilizeMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizeMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StabilizeMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizeMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StabilizeMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StabilizeMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StabilizeMigratorSession struct {
	Contract     *StabilizeMigrator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StabilizeMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StabilizeMigratorCallerSession struct {
	Contract *StabilizeMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StabilizeMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StabilizeMigratorTransactorSession struct {
	Contract     *StabilizeMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StabilizeMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type StabilizeMigratorRaw struct {
	Contract *StabilizeMigrator // Generic contract binding to access the raw methods on
}

// StabilizeMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StabilizeMigratorCallerRaw struct {
	Contract *StabilizeMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// StabilizeMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StabilizeMigratorTransactorRaw struct {
	Contract *StabilizeMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStabilizeMigrator creates a new instance of StabilizeMigrator, bound to a specific deployed contract.
func NewStabilizeMigrator(address common.Address, backend bind.ContractBackend) (*StabilizeMigrator, error) {
	contract, err := bindStabilizeMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StabilizeMigrator{StabilizeMigratorCaller: StabilizeMigratorCaller{contract: contract}, StabilizeMigratorTransactor: StabilizeMigratorTransactor{contract: contract}, StabilizeMigratorFilterer: StabilizeMigratorFilterer{contract: contract}}, nil
}

// NewStabilizeMigratorCaller creates a new read-only instance of StabilizeMigrator, bound to a specific deployed contract.
func NewStabilizeMigratorCaller(address common.Address, caller bind.ContractCaller) (*StabilizeMigratorCaller, error) {
	contract, err := bindStabilizeMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StabilizeMigratorCaller{contract: contract}, nil
}

// NewStabilizeMigratorTransactor creates a new write-only instance of StabilizeMigrator, bound to a specific deployed contract.
func NewStabilizeMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*StabilizeMigratorTransactor, error) {
	contract, err := bindStabilizeMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StabilizeMigratorTransactor{contract: contract}, nil
}

// NewStabilizeMigratorFilterer creates a new log filterer instance of StabilizeMigrator, bound to a specific deployed contract.
func NewStabilizeMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*StabilizeMigratorFilterer, error) {
	contract, err := bindStabilizeMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StabilizeMigratorFilterer{contract: contract}, nil
}

// bindStabilizeMigrator binds a generic wrapper to an already deployed contract.
func bindStabilizeMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StabilizeMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StabilizeMigrator *StabilizeMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StabilizeMigrator.Contract.StabilizeMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StabilizeMigrator *StabilizeMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.StabilizeMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StabilizeMigrator *StabilizeMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.StabilizeMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StabilizeMigrator *StabilizeMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StabilizeMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StabilizeMigrator *StabilizeMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StabilizeMigrator *StabilizeMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) Admin() (common.Address, error) {
	return _StabilizeMigrator.Contract.Admin(&_StabilizeMigrator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) Admin() (common.Address, error) {
	return _StabilizeMigrator.Contract.Admin(&_StabilizeMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) BnGmx() (common.Address, error) {
	return _StabilizeMigrator.Contract.BnGmx(&_StabilizeMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) BnGmx() (common.Address, error) {
	return _StabilizeMigrator.Contract.BnGmx(&_StabilizeMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) BonusGmxTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.BonusGmxTracker(&_StabilizeMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) BonusGmxTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.BonusGmxTracker(&_StabilizeMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) EsGmx() (common.Address, error) {
	return _StabilizeMigrator.Contract.EsGmx(&_StabilizeMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) EsGmx() (common.Address, error) {
	return _StabilizeMigrator.Contract.EsGmx(&_StabilizeMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) ExpectedGovGrantedCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "expectedGovGrantedCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _StabilizeMigrator.Contract.ExpectedGovGrantedCaller(&_StabilizeMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _StabilizeMigrator.Contract.ExpectedGovGrantedCaller(&_StabilizeMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) FeeGlpTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.FeeGlpTracker(&_StabilizeMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) FeeGlpTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.FeeGlpTracker(&_StabilizeMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) FeeGmxTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.FeeGmxTracker(&_StabilizeMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) FeeGmxTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.FeeGmxTracker(&_StabilizeMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) GlpVester() (common.Address, error) {
	return _StabilizeMigrator.Contract.GlpVester(&_StabilizeMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) GlpVester() (common.Address, error) {
	return _StabilizeMigrator.Contract.GlpVester(&_StabilizeMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) GmxVester() (common.Address, error) {
	return _StabilizeMigrator.Contract.GmxVester(&_StabilizeMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) GmxVester() (common.Address, error) {
	return _StabilizeMigrator.Contract.GmxVester(&_StabilizeMigrator.CallOpts)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) RewardRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "rewardRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) RewardRouter() (common.Address, error) {
	return _StabilizeMigrator.Contract.RewardRouter(&_StabilizeMigrator.CallOpts)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) RewardRouter() (common.Address, error) {
	return _StabilizeMigrator.Contract.RewardRouter(&_StabilizeMigrator.CallOpts)
}

// StabilizeCaller is a free data retrieval call binding the contract method 0xb3ee2d19.
//
// Solidity: function stabilizeCaller() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) StabilizeCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "stabilizeCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StabilizeCaller is a free data retrieval call binding the contract method 0xb3ee2d19.
//
// Solidity: function stabilizeCaller() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) StabilizeCaller() (common.Address, error) {
	return _StabilizeMigrator.Contract.StabilizeCaller(&_StabilizeMigrator.CallOpts)
}

// StabilizeCaller is a free data retrieval call binding the contract method 0xb3ee2d19.
//
// Solidity: function stabilizeCaller() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) StabilizeCaller() (common.Address, error) {
	return _StabilizeMigrator.Contract.StabilizeCaller(&_StabilizeMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) StakedGlpTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.StakedGlpTracker(&_StabilizeMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) StakedGlpTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.StakedGlpTracker(&_StabilizeMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCaller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StabilizeMigrator.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorSession) StakedGmxTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.StakedGmxTracker(&_StabilizeMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_StabilizeMigrator *StabilizeMigratorCallerSession) StakedGmxTracker() (common.Address, error) {
	return _StabilizeMigrator.Contract.StakedGmxTracker(&_StabilizeMigrator.CallOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_StabilizeMigrator *StabilizeMigratorTransactor) AfterGovGranted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StabilizeMigrator.contract.Transact(opts, "afterGovGranted")
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_StabilizeMigrator *StabilizeMigratorSession) AfterGovGranted() (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.AfterGovGranted(&_StabilizeMigrator.TransactOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_StabilizeMigrator *StabilizeMigratorTransactorSession) AfterGovGranted() (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.AfterGovGranted(&_StabilizeMigrator.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_StabilizeMigrator *StabilizeMigratorTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StabilizeMigrator.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_StabilizeMigrator *StabilizeMigratorSession) Migrate() (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.Migrate(&_StabilizeMigrator.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_StabilizeMigrator *StabilizeMigratorTransactorSession) Migrate() (*types.Transaction, error) {
	return _StabilizeMigrator.Contract.Migrate(&_StabilizeMigrator.TransactOpts)
}
