// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package buybackmigrator

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

// BuybackMigratorMetaData contains all meta data concerning the BuybackMigrator contract.
var BuybackMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_extendedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oldRewardRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newRewardRouter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"afterGovGranted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableOldRewardRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableNewRewardRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedGovGrantedCaller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extendedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRestakingCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRewardRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldRewardRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRouterTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setHandlerAndDepositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BuybackMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use BuybackMigratorMetaData.ABI instead.
var BuybackMigratorABI = BuybackMigratorMetaData.ABI

// BuybackMigrator is an auto generated Go binding around an Ethereum contract.
type BuybackMigrator struct {
	BuybackMigratorCaller     // Read-only binding to the contract
	BuybackMigratorTransactor // Write-only binding to the contract
	BuybackMigratorFilterer   // Log filterer for contract events
}

// BuybackMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuybackMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuybackMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuybackMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuybackMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuybackMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuybackMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuybackMigratorSession struct {
	Contract     *BuybackMigrator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuybackMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuybackMigratorCallerSession struct {
	Contract *BuybackMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BuybackMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuybackMigratorTransactorSession struct {
	Contract     *BuybackMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BuybackMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuybackMigratorRaw struct {
	Contract *BuybackMigrator // Generic contract binding to access the raw methods on
}

// BuybackMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuybackMigratorCallerRaw struct {
	Contract *BuybackMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// BuybackMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuybackMigratorTransactorRaw struct {
	Contract *BuybackMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuybackMigrator creates a new instance of BuybackMigrator, bound to a specific deployed contract.
func NewBuybackMigrator(address common.Address, backend bind.ContractBackend) (*BuybackMigrator, error) {
	contract, err := bindBuybackMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuybackMigrator{BuybackMigratorCaller: BuybackMigratorCaller{contract: contract}, BuybackMigratorTransactor: BuybackMigratorTransactor{contract: contract}, BuybackMigratorFilterer: BuybackMigratorFilterer{contract: contract}}, nil
}

// NewBuybackMigratorCaller creates a new read-only instance of BuybackMigrator, bound to a specific deployed contract.
func NewBuybackMigratorCaller(address common.Address, caller bind.ContractCaller) (*BuybackMigratorCaller, error) {
	contract, err := bindBuybackMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuybackMigratorCaller{contract: contract}, nil
}

// NewBuybackMigratorTransactor creates a new write-only instance of BuybackMigrator, bound to a specific deployed contract.
func NewBuybackMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*BuybackMigratorTransactor, error) {
	contract, err := bindBuybackMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuybackMigratorTransactor{contract: contract}, nil
}

// NewBuybackMigratorFilterer creates a new log filterer instance of BuybackMigrator, bound to a specific deployed contract.
func NewBuybackMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*BuybackMigratorFilterer, error) {
	contract, err := bindBuybackMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuybackMigratorFilterer{contract: contract}, nil
}

// bindBuybackMigrator binds a generic wrapper to an already deployed contract.
func bindBuybackMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuybackMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuybackMigrator *BuybackMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuybackMigrator.Contract.BuybackMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuybackMigrator *BuybackMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuybackMigrator.Contract.BuybackMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuybackMigrator *BuybackMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuybackMigrator.Contract.BuybackMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuybackMigrator *BuybackMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuybackMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuybackMigrator *BuybackMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuybackMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuybackMigrator *BuybackMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuybackMigrator.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) Admin() (common.Address, error) {
	return _BuybackMigrator.Contract.Admin(&_BuybackMigrator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) Admin() (common.Address, error) {
	return _BuybackMigrator.Contract.Admin(&_BuybackMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) BnGmx() (common.Address, error) {
	return _BuybackMigrator.Contract.BnGmx(&_BuybackMigrator.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) BnGmx() (common.Address, error) {
	return _BuybackMigrator.Contract.BnGmx(&_BuybackMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) BonusGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.BonusGmxTracker(&_BuybackMigrator.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) BonusGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.BonusGmxTracker(&_BuybackMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) EsGmx() (common.Address, error) {
	return _BuybackMigrator.Contract.EsGmx(&_BuybackMigrator.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) EsGmx() (common.Address, error) {
	return _BuybackMigrator.Contract.EsGmx(&_BuybackMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) ExpectedGovGrantedCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "expectedGovGrantedCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _BuybackMigrator.Contract.ExpectedGovGrantedCaller(&_BuybackMigrator.CallOpts)
}

// ExpectedGovGrantedCaller is a free data retrieval call binding the contract method 0x9a2e0894.
//
// Solidity: function expectedGovGrantedCaller() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) ExpectedGovGrantedCaller() (common.Address, error) {
	return _BuybackMigrator.Contract.ExpectedGovGrantedCaller(&_BuybackMigrator.CallOpts)
}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) ExtendedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "extendedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) ExtendedGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.ExtendedGmxTracker(&_BuybackMigrator.CallOpts)
}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) ExtendedGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.ExtendedGmxTracker(&_BuybackMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) FeeGlpTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.FeeGlpTracker(&_BuybackMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) FeeGlpTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.FeeGlpTracker(&_BuybackMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) FeeGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.FeeGmxTracker(&_BuybackMigrator.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) FeeGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.FeeGmxTracker(&_BuybackMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) GlpVester() (common.Address, error) {
	return _BuybackMigrator.Contract.GlpVester(&_BuybackMigrator.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) GlpVester() (common.Address, error) {
	return _BuybackMigrator.Contract.GlpVester(&_BuybackMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) GmxVester() (common.Address, error) {
	return _BuybackMigrator.Contract.GmxVester(&_BuybackMigrator.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) GmxVester() (common.Address, error) {
	return _BuybackMigrator.Contract.GmxVester(&_BuybackMigrator.CallOpts)
}

// IsEnabled is a free data retrieval call binding the contract method 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (_BuybackMigrator *BuybackMigratorCaller) IsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "isEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnabled is a free data retrieval call binding the contract method 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (_BuybackMigrator *BuybackMigratorSession) IsEnabled() (bool, error) {
	return _BuybackMigrator.Contract.IsEnabled(&_BuybackMigrator.CallOpts)
}

// IsEnabled is a free data retrieval call binding the contract method 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (_BuybackMigrator *BuybackMigratorCallerSession) IsEnabled() (bool, error) {
	return _BuybackMigrator.Contract.IsEnabled(&_BuybackMigrator.CallOpts)
}

// IsRestakingCompleted is a free data retrieval call binding the contract method 0x5b9c37e0.
//
// Solidity: function isRestakingCompleted() view returns(bool)
func (_BuybackMigrator *BuybackMigratorCaller) IsRestakingCompleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "isRestakingCompleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRestakingCompleted is a free data retrieval call binding the contract method 0x5b9c37e0.
//
// Solidity: function isRestakingCompleted() view returns(bool)
func (_BuybackMigrator *BuybackMigratorSession) IsRestakingCompleted() (bool, error) {
	return _BuybackMigrator.Contract.IsRestakingCompleted(&_BuybackMigrator.CallOpts)
}

// IsRestakingCompleted is a free data retrieval call binding the contract method 0x5b9c37e0.
//
// Solidity: function isRestakingCompleted() view returns(bool)
func (_BuybackMigrator *BuybackMigratorCallerSession) IsRestakingCompleted() (bool, error) {
	return _BuybackMigrator.Contract.IsRestakingCompleted(&_BuybackMigrator.CallOpts)
}

// NewRewardRouter is a free data retrieval call binding the contract method 0x98b73001.
//
// Solidity: function newRewardRouter() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) NewRewardRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "newRewardRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewRewardRouter is a free data retrieval call binding the contract method 0x98b73001.
//
// Solidity: function newRewardRouter() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) NewRewardRouter() (common.Address, error) {
	return _BuybackMigrator.Contract.NewRewardRouter(&_BuybackMigrator.CallOpts)
}

// NewRewardRouter is a free data retrieval call binding the contract method 0x98b73001.
//
// Solidity: function newRewardRouter() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) NewRewardRouter() (common.Address, error) {
	return _BuybackMigrator.Contract.NewRewardRouter(&_BuybackMigrator.CallOpts)
}

// OldRewardRouter is a free data retrieval call binding the contract method 0x45914f60.
//
// Solidity: function oldRewardRouter() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) OldRewardRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "oldRewardRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OldRewardRouter is a free data retrieval call binding the contract method 0x45914f60.
//
// Solidity: function oldRewardRouter() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) OldRewardRouter() (common.Address, error) {
	return _BuybackMigrator.Contract.OldRewardRouter(&_BuybackMigrator.CallOpts)
}

// OldRewardRouter is a free data retrieval call binding the contract method 0x45914f60.
//
// Solidity: function oldRewardRouter() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) OldRewardRouter() (common.Address, error) {
	return _BuybackMigrator.Contract.OldRewardRouter(&_BuybackMigrator.CallOpts)
}

// RewardRouterTarget is a free data retrieval call binding the contract method 0x5f0175bf.
//
// Solidity: function rewardRouterTarget() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) RewardRouterTarget(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "rewardRouterTarget")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardRouterTarget is a free data retrieval call binding the contract method 0x5f0175bf.
//
// Solidity: function rewardRouterTarget() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) RewardRouterTarget() (common.Address, error) {
	return _BuybackMigrator.Contract.RewardRouterTarget(&_BuybackMigrator.CallOpts)
}

// RewardRouterTarget is a free data retrieval call binding the contract method 0x5f0175bf.
//
// Solidity: function rewardRouterTarget() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) RewardRouterTarget() (common.Address, error) {
	return _BuybackMigrator.Contract.RewardRouterTarget(&_BuybackMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) StakedGlpTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.StakedGlpTracker(&_BuybackMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) StakedGlpTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.StakedGlpTracker(&_BuybackMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCaller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BuybackMigrator.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorSession) StakedGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.StakedGmxTracker(&_BuybackMigrator.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_BuybackMigrator *BuybackMigratorCallerSession) StakedGmxTracker() (common.Address, error) {
	return _BuybackMigrator.Contract.StakedGmxTracker(&_BuybackMigrator.CallOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BuybackMigrator *BuybackMigratorTransactor) AfterGovGranted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuybackMigrator.contract.Transact(opts, "afterGovGranted")
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BuybackMigrator *BuybackMigratorSession) AfterGovGranted() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.AfterGovGranted(&_BuybackMigrator.TransactOpts)
}

// AfterGovGranted is a paid mutator transaction binding the contract method 0xc7cfeb59.
//
// Solidity: function afterGovGranted() returns()
func (_BuybackMigrator *BuybackMigratorTransactorSession) AfterGovGranted() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.AfterGovGranted(&_BuybackMigrator.TransactOpts)
}

// DisableOldRewardRouter is a paid mutator transaction binding the contract method 0xeb6a305a.
//
// Solidity: function disableOldRewardRouter() returns()
func (_BuybackMigrator *BuybackMigratorTransactor) DisableOldRewardRouter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuybackMigrator.contract.Transact(opts, "disableOldRewardRouter")
}

// DisableOldRewardRouter is a paid mutator transaction binding the contract method 0xeb6a305a.
//
// Solidity: function disableOldRewardRouter() returns()
func (_BuybackMigrator *BuybackMigratorSession) DisableOldRewardRouter() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.DisableOldRewardRouter(&_BuybackMigrator.TransactOpts)
}

// DisableOldRewardRouter is a paid mutator transaction binding the contract method 0xeb6a305a.
//
// Solidity: function disableOldRewardRouter() returns()
func (_BuybackMigrator *BuybackMigratorTransactorSession) DisableOldRewardRouter() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.DisableOldRewardRouter(&_BuybackMigrator.TransactOpts)
}

// EnableNewRewardRouter is a paid mutator transaction binding the contract method 0x45163ae8.
//
// Solidity: function enableNewRewardRouter() returns()
func (_BuybackMigrator *BuybackMigratorTransactor) EnableNewRewardRouter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuybackMigrator.contract.Transact(opts, "enableNewRewardRouter")
}

// EnableNewRewardRouter is a paid mutator transaction binding the contract method 0x45163ae8.
//
// Solidity: function enableNewRewardRouter() returns()
func (_BuybackMigrator *BuybackMigratorSession) EnableNewRewardRouter() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.EnableNewRewardRouter(&_BuybackMigrator.TransactOpts)
}

// EnableNewRewardRouter is a paid mutator transaction binding the contract method 0x45163ae8.
//
// Solidity: function enableNewRewardRouter() returns()
func (_BuybackMigrator *BuybackMigratorTransactorSession) EnableNewRewardRouter() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.EnableNewRewardRouter(&_BuybackMigrator.TransactOpts)
}

// SetHandlerAndDepositToken is a paid mutator transaction binding the contract method 0xbe282afc.
//
// Solidity: function setHandlerAndDepositToken() returns()
func (_BuybackMigrator *BuybackMigratorTransactor) SetHandlerAndDepositToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuybackMigrator.contract.Transact(opts, "setHandlerAndDepositToken")
}

// SetHandlerAndDepositToken is a paid mutator transaction binding the contract method 0xbe282afc.
//
// Solidity: function setHandlerAndDepositToken() returns()
func (_BuybackMigrator *BuybackMigratorSession) SetHandlerAndDepositToken() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.SetHandlerAndDepositToken(&_BuybackMigrator.TransactOpts)
}

// SetHandlerAndDepositToken is a paid mutator transaction binding the contract method 0xbe282afc.
//
// Solidity: function setHandlerAndDepositToken() returns()
func (_BuybackMigrator *BuybackMigratorTransactorSession) SetHandlerAndDepositToken() (*types.Transaction, error) {
	return _BuybackMigrator.Contract.SetHandlerAndDepositToken(&_BuybackMigrator.TransactOpts)
}
