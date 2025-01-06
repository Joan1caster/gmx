// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basemigrator

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

// BaseMigratorMetaData contains all meta data concerning the BaseMigrator contract.
var BaseMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"afterGovGranted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedGovGrantedCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BaseMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMigratorMetaData.ABI instead.
var BaseMigratorABI = BaseMigratorMetaData.ABI

// BaseMigrator is an auto generated Go binding around an Ethereum contract.
type BaseMigrator struct {
	BaseMigratorCaller     // Read-only binding to the contract
	BaseMigratorTransactor // Write-only binding to the contract
	BaseMigratorFilterer   // Log filterer for contract events
}

// BaseMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseMigratorSession struct {
	Contract     *BaseMigrator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseMigratorCallerSession struct {
	Contract *BaseMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseMigratorTransactorSession struct {
	Contract     *BaseMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseMigratorRaw struct {
	Contract *BaseMigrator // Generic contract binding to access the raw methods on
}

// BaseMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseMigratorCallerRaw struct {
	Contract *BaseMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// BaseMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseMigratorTransactorRaw struct {
	Contract *BaseMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseMigrator creates a new instance of BaseMigrator, bound to a specific deployed contract.
func NewBaseMigrator(address common.Address, backend bind.ContractBackend) (*BaseMigrator, error) {
	contract, err := bindBaseMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseMigrator{BaseMigratorCaller: BaseMigratorCaller{contract: contract}, BaseMigratorTransactor: BaseMigratorTransactor{contract: contract}, BaseMigratorFilterer: BaseMigratorFilterer{contract: contract}}, nil
}

// NewBaseMigratorCaller creates a new read-only instance of BaseMigrator, bound to a specific deployed contract.
func NewBaseMigratorCaller(address common.Address, caller bind.ContractCaller) (*BaseMigratorCaller, error) {
	contract, err := bindBaseMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMigratorCaller{contract: contract}, nil
}

// NewBaseMigratorTransactor creates a new write-only instance of BaseMigrator, bound to a specific deployed contract.
func NewBaseMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseMigratorTransactor, error) {
	contract, err := bindBaseMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMigratorTransactor{contract: contract}, nil
}

// NewBaseMigratorFilterer creates a new log filterer instance of BaseMigrator, bound to a specific deployed contract.
func NewBaseMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseMigratorFilterer, error) {
	contract, err := bindBaseMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseMigratorFilterer{contract: contract}, nil
}

// bindBaseMigrator binds a generic wrapper to an already deployed contract.
func bindBaseMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMigrator *BaseMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMigrator.Contract.BaseMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMigrator *BaseMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMigrator.Contract.BaseMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMigrator *BaseMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMigrator.Contract.BaseMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMigrator *BaseMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMigrator *BaseMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMigrator *BaseMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMigrator.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BaseMigrator *BaseMigratorSession) Admin() (common.Address, error) {
	return _BaseMigrator.Contract.Admin(&_BaseMigrator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) Admin() (common.Address, error) {
	return _BaseMigrator.Contract.Admin(&_BaseMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BaseMigrator *BaseMigratorSession) BnGmx() (common.Address, error) {
	return _BaseMigrator.Contract.BnGmx(&_BaseMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) BnGmx() (common.Address, error) {
	return _BaseMigrator.Contract.BnGmx(&_BaseMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorSession) BonusGmxTracker() (common.Address, error) {
	return _BaseMigrator.Contract.BonusGmxTracker(&_BaseMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) BonusGmxTracker() (common.Address, error) {
	return _BaseMigrator.Contract.BonusGmxTracker(&_BaseMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BaseMigrator *BaseMigratorSession) EsGmx() (common.Address, error) {
	return _BaseMigrator.Contract.EsGmx(&_BaseMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) EsGmx() (common.Address, error) {
	return _BaseMigrator.Contract.EsGmx(&_BaseMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) ExpectedGovGrantedCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "expectedGovGrantedCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BaseMigrator *BaseMigratorSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _BaseMigrator.Contract.ExpectedGovGrantedCaller(&_BaseMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _BaseMigrator.Contract.ExpectedGovGrantedCaller(&_BaseMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BaseMigrator *BaseMigratorSession) FeeGlpTracker() (common.Address, error) {
	return _BaseMigrator.Contract.FeeGlpTracker(&_BaseMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) FeeGlpTracker() (common.Address, error) {
	return _BaseMigrator.Contract.FeeGlpTracker(&_BaseMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorSession) FeeGmxTracker() (common.Address, error) {
	return _BaseMigrator.Contract.FeeGmxTracker(&_BaseMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) FeeGmxTracker() (common.Address, error) {
	return _BaseMigrator.Contract.FeeGmxTracker(&_BaseMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BaseMigrator *BaseMigratorSession) GlpVester() (common.Address, error) {
	return _BaseMigrator.Contract.GlpVester(&_BaseMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) GlpVester() (common.Address, error) {
	return _BaseMigrator.Contract.GlpVester(&_BaseMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BaseMigrator *BaseMigratorSession) GmxVester() (common.Address, error) {
	return _BaseMigrator.Contract.GmxVester(&_BaseMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) GmxVester() (common.Address, error) {
	return _BaseMigrator.Contract.GmxVester(&_BaseMigrator.CallOpts)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) RewardRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "rewardRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_BaseMigrator *BaseMigratorSession) RewardRouter() (common.Address, error) {
	return _BaseMigrator.Contract.RewardRouter(&_BaseMigrator.CallOpts)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) RewardRouter() (common.Address, error) {
	return _BaseMigrator.Contract.RewardRouter(&_BaseMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BaseMigrator *BaseMigratorSession) StakedGlpTracker() (common.Address, error) {
	return _BaseMigrator.Contract.StakedGlpTracker(&_BaseMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) StakedGlpTracker() (common.Address, error) {
	return _BaseMigrator.Contract.StakedGlpTracker(&_BaseMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCaller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseMigrator.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorSession) StakedGmxTracker() (common.Address, error) {
	return _BaseMigrator.Contract.StakedGmxTracker(&_BaseMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BaseMigrator *BaseMigratorCallerSession) StakedGmxTracker() (common.Address, error) {
	return _BaseMigrator.Contract.StakedGmxTracker(&_BaseMigrator.CallOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BaseMigrator *BaseMigratorTransactor) AfterGovGranted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMigrator.contract.Transact(opts, "afterGovGranted")
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BaseMigrator *BaseMigratorSession) AfterGovGranted() (*types.Transaction, error) {
	return _BaseMigrator.Contract.AfterGovGranted(&_BaseMigrator.TransactOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BaseMigrator *BaseMigratorTransactorSession) AfterGovGranted() (*types.Transaction, error) {
	return _BaseMigrator.Contract.AfterGovGranted(&_BaseMigrator.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_BaseMigrator *BaseMigratorTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMigrator.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_BaseMigrator *BaseMigratorSession) Migrate() (*types.Transaction, error) {
	return _BaseMigrator.Contract.Migrate(&_BaseMigrator.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_BaseMigrator *BaseMigratorTransactorSession) Migrate() (*types.Transaction, error) {
	return _BaseMigrator.Contract.Migrate(&_BaseMigrator.TransactOpts)
}
