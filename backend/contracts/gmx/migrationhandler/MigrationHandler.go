// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package migrationhandler

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

// MigrationHandlerMetaData contains all meta data concerning the MigrationHandler contract.
var MigrationHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"USDG_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ammRouterV1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ammRouterV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"busd\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ammRouterV1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ammRouterV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xgmt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnb\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_busd\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_redemptionToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUsdg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"refundedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gmtAmountForUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_xgmtAmountForUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gmtAmountForBusd\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xgmt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MigrationHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use MigrationHandlerMetaData.ABI instead.
var MigrationHandlerABI = MigrationHandlerMetaData.ABI

// MigrationHandler is an auto generated Go binding around an Ethereum contract.
type MigrationHandler struct {
	MigrationHandlerCaller     // Read-only binding to the contract
	MigrationHandlerTransactor // Write-only binding to the contract
	MigrationHandlerFilterer   // Log filterer for contract events
}

// MigrationHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MigrationHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MigrationHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MigrationHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MigrationHandlerSession struct {
	Contract     *MigrationHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MigrationHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MigrationHandlerCallerSession struct {
	Contract *MigrationHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MigrationHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MigrationHandlerTransactorSession struct {
	Contract     *MigrationHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MigrationHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MigrationHandlerRaw struct {
	Contract *MigrationHandler // Generic contract binding to access the raw methods on
}

// MigrationHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MigrationHandlerCallerRaw struct {
	Contract *MigrationHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// MigrationHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MigrationHandlerTransactorRaw struct {
	Contract *MigrationHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMigrationHandler creates a new instance of MigrationHandler, bound to a specific deployed contract.
func NewMigrationHandler(address common.Address, backend bind.ContractBackend) (*MigrationHandler, error) {
	contract, err := bindMigrationHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MigrationHandler{MigrationHandlerCaller: MigrationHandlerCaller{contract: contract}, MigrationHandlerTransactor: MigrationHandlerTransactor{contract: contract}, MigrationHandlerFilterer: MigrationHandlerFilterer{contract: contract}}, nil
}

// NewMigrationHandlerCaller creates a new read-only instance of MigrationHandler, bound to a specific deployed contract.
func NewMigrationHandlerCaller(address common.Address, caller bind.ContractCaller) (*MigrationHandlerCaller, error) {
	contract, err := bindMigrationHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationHandlerCaller{contract: contract}, nil
}

// NewMigrationHandlerTransactor creates a new write-only instance of MigrationHandler, bound to a specific deployed contract.
func NewMigrationHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*MigrationHandlerTransactor, error) {
	contract, err := bindMigrationHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationHandlerTransactor{contract: contract}, nil
}

// NewMigrationHandlerFilterer creates a new log filterer instance of MigrationHandler, bound to a specific deployed contract.
func NewMigrationHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*MigrationHandlerFilterer, error) {
	contract, err := bindMigrationHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MigrationHandlerFilterer{contract: contract}, nil
}

// bindMigrationHandler binds a generic wrapper to an already deployed contract.
func bindMigrationHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MigrationHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MigrationHandler *MigrationHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MigrationHandler.Contract.MigrationHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MigrationHandler *MigrationHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MigrationHandler.Contract.MigrationHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MigrationHandler *MigrationHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MigrationHandler.Contract.MigrationHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MigrationHandler *MigrationHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MigrationHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MigrationHandler *MigrationHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MigrationHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MigrationHandler *MigrationHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MigrationHandler.Contract.contract.Transact(opts, method, params...)
}

// USDGPRECISION is a free data retrieval call binding the contract method 0x4a686d67.
//
// Solidity: function USDG_PRECISION() view returns(uint256)
func (_MigrationHandler *MigrationHandlerCaller) USDGPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "USDG_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGPRECISION is a free data retrieval call binding the contract method 0x4a686d67.
//
// Solidity: function USDG_PRECISION() view returns(uint256)
func (_MigrationHandler *MigrationHandlerSession) USDGPRECISION() (*big.Int, error) {
	return _MigrationHandler.Contract.USDGPRECISION(&_MigrationHandler.CallOpts)
}

// USDGPRECISION is a free data retrieval call binding the contract method 0x4a686d67.
//
// Solidity: function USDG_PRECISION() view returns(uint256)
func (_MigrationHandler *MigrationHandlerCallerSession) USDGPRECISION() (*big.Int, error) {
	return _MigrationHandler.Contract.USDGPRECISION(&_MigrationHandler.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) Admin() (common.Address, error) {
	return _MigrationHandler.Contract.Admin(&_MigrationHandler.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) Admin() (common.Address, error) {
	return _MigrationHandler.Contract.Admin(&_MigrationHandler.CallOpts)
}

// AmmRouterV1 is a free data retrieval call binding the contract method 0xa53064d1.
//
// Solidity: function ammRouterV1() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) AmmRouterV1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "ammRouterV1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmRouterV1 is a free data retrieval call binding the contract method 0xa53064d1.
//
// Solidity: function ammRouterV1() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) AmmRouterV1() (common.Address, error) {
	return _MigrationHandler.Contract.AmmRouterV1(&_MigrationHandler.CallOpts)
}

// AmmRouterV1 is a free data retrieval call binding the contract method 0xa53064d1.
//
// Solidity: function ammRouterV1() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) AmmRouterV1() (common.Address, error) {
	return _MigrationHandler.Contract.AmmRouterV1(&_MigrationHandler.CallOpts)
}

// AmmRouterV2 is a free data retrieval call binding the contract method 0xade027f6.
//
// Solidity: function ammRouterV2() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) AmmRouterV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "ammRouterV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmRouterV2 is a free data retrieval call binding the contract method 0xade027f6.
//
// Solidity: function ammRouterV2() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) AmmRouterV2() (common.Address, error) {
	return _MigrationHandler.Contract.AmmRouterV2(&_MigrationHandler.CallOpts)
}

// AmmRouterV2 is a free data retrieval call binding the contract method 0xade027f6.
//
// Solidity: function ammRouterV2() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) AmmRouterV2() (common.Address, error) {
	return _MigrationHandler.Contract.AmmRouterV2(&_MigrationHandler.CallOpts)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) Bnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "bnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) Bnb() (common.Address, error) {
	return _MigrationHandler.Contract.Bnb(&_MigrationHandler.CallOpts)
}

// Bnb is a free data retrieval call binding the contract method 0x49a876e4.
//
// Solidity: function bnb() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) Bnb() (common.Address, error) {
	return _MigrationHandler.Contract.Bnb(&_MigrationHandler.CallOpts)
}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) Busd(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "busd")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) Busd() (common.Address, error) {
	return _MigrationHandler.Contract.Busd(&_MigrationHandler.CallOpts)
}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) Busd() (common.Address, error) {
	return _MigrationHandler.Contract.Busd(&_MigrationHandler.CallOpts)
}

// Gmt is a free data retrieval call binding the contract method 0xf8322d24.
//
// Solidity: function gmt() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) Gmt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "gmt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gmt is a free data retrieval call binding the contract method 0xf8322d24.
//
// Solidity: function gmt() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) Gmt() (common.Address, error) {
	return _MigrationHandler.Contract.Gmt(&_MigrationHandler.CallOpts)
}

// Gmt is a free data retrieval call binding the contract method 0xf8322d24.
//
// Solidity: function gmt() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) Gmt() (common.Address, error) {
	return _MigrationHandler.Contract.Gmt(&_MigrationHandler.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_MigrationHandler *MigrationHandlerCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_MigrationHandler *MigrationHandlerSession) IsInitialized() (bool, error) {
	return _MigrationHandler.Contract.IsInitialized(&_MigrationHandler.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_MigrationHandler *MigrationHandlerCallerSession) IsInitialized() (bool, error) {
	return _MigrationHandler.Contract.IsInitialized(&_MigrationHandler.CallOpts)
}

// RefundedAmounts is a free data retrieval call binding the contract method 0x252c3222.
//
// Solidity: function refundedAmounts(address , address ) view returns(uint256)
func (_MigrationHandler *MigrationHandlerCaller) RefundedAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "refundedAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundedAmounts is a free data retrieval call binding the contract method 0x252c3222.
//
// Solidity: function refundedAmounts(address , address ) view returns(uint256)
func (_MigrationHandler *MigrationHandlerSession) RefundedAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MigrationHandler.Contract.RefundedAmounts(&_MigrationHandler.CallOpts, arg0, arg1)
}

// RefundedAmounts is a free data retrieval call binding the contract method 0x252c3222.
//
// Solidity: function refundedAmounts(address , address ) view returns(uint256)
func (_MigrationHandler *MigrationHandlerCallerSession) RefundedAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MigrationHandler.Contract.RefundedAmounts(&_MigrationHandler.CallOpts, arg0, arg1)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) Usdg() (common.Address, error) {
	return _MigrationHandler.Contract.Usdg(&_MigrationHandler.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) Usdg() (common.Address, error) {
	return _MigrationHandler.Contract.Usdg(&_MigrationHandler.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) Vault() (common.Address, error) {
	return _MigrationHandler.Contract.Vault(&_MigrationHandler.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) Vault() (common.Address, error) {
	return _MigrationHandler.Contract.Vault(&_MigrationHandler.CallOpts)
}

// Xgmt is a free data retrieval call binding the contract method 0xa9eb51c6.
//
// Solidity: function xgmt() view returns(address)
func (_MigrationHandler *MigrationHandlerCaller) Xgmt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MigrationHandler.contract.Call(opts, &out, "xgmt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xgmt is a free data retrieval call binding the contract method 0xa9eb51c6.
//
// Solidity: function xgmt() view returns(address)
func (_MigrationHandler *MigrationHandlerSession) Xgmt() (common.Address, error) {
	return _MigrationHandler.Contract.Xgmt(&_MigrationHandler.CallOpts)
}

// Xgmt is a free data retrieval call binding the contract method 0xa9eb51c6.
//
// Solidity: function xgmt() view returns(address)
func (_MigrationHandler *MigrationHandlerCallerSession) Xgmt() (common.Address, error) {
	return _MigrationHandler.Contract.Xgmt(&_MigrationHandler.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _ammRouterV1, address _ammRouterV2, address _vault, address _gmt, address _xgmt, address _usdg, address _bnb, address _busd) returns()
func (_MigrationHandler *MigrationHandlerTransactor) Initialize(opts *bind.TransactOpts, _ammRouterV1 common.Address, _ammRouterV2 common.Address, _vault common.Address, _gmt common.Address, _xgmt common.Address, _usdg common.Address, _bnb common.Address, _busd common.Address) (*types.Transaction, error) {
	return _MigrationHandler.contract.Transact(opts, "initialize", _ammRouterV1, _ammRouterV2, _vault, _gmt, _xgmt, _usdg, _bnb, _busd)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _ammRouterV1, address _ammRouterV2, address _vault, address _gmt, address _xgmt, address _usdg, address _bnb, address _busd) returns()
func (_MigrationHandler *MigrationHandlerSession) Initialize(_ammRouterV1 common.Address, _ammRouterV2 common.Address, _vault common.Address, _gmt common.Address, _xgmt common.Address, _usdg common.Address, _bnb common.Address, _busd common.Address) (*types.Transaction, error) {
	return _MigrationHandler.Contract.Initialize(&_MigrationHandler.TransactOpts, _ammRouterV1, _ammRouterV2, _vault, _gmt, _xgmt, _usdg, _bnb, _busd)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _ammRouterV1, address _ammRouterV2, address _vault, address _gmt, address _xgmt, address _usdg, address _bnb, address _busd) returns()
func (_MigrationHandler *MigrationHandlerTransactorSession) Initialize(_ammRouterV1 common.Address, _ammRouterV2 common.Address, _vault common.Address, _gmt common.Address, _xgmt common.Address, _usdg common.Address, _bnb common.Address, _busd common.Address) (*types.Transaction, error) {
	return _MigrationHandler.Contract.Initialize(&_MigrationHandler.TransactOpts, _ammRouterV1, _ammRouterV2, _vault, _gmt, _xgmt, _usdg, _bnb, _busd)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _migrator, address _redemptionToken, uint256 _usdgAmount) returns()
func (_MigrationHandler *MigrationHandlerTransactor) RedeemUsdg(opts *bind.TransactOpts, _migrator common.Address, _redemptionToken common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.contract.Transact(opts, "redeemUsdg", _migrator, _redemptionToken, _usdgAmount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _migrator, address _redemptionToken, uint256 _usdgAmount) returns()
func (_MigrationHandler *MigrationHandlerSession) RedeemUsdg(_migrator common.Address, _redemptionToken common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.Contract.RedeemUsdg(&_MigrationHandler.TransactOpts, _migrator, _redemptionToken, _usdgAmount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _migrator, address _redemptionToken, uint256 _usdgAmount) returns()
func (_MigrationHandler *MigrationHandlerTransactorSession) RedeemUsdg(_migrator common.Address, _redemptionToken common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.Contract.RedeemUsdg(&_MigrationHandler.TransactOpts, _migrator, _redemptionToken, _usdgAmount)
}

// Refund is a paid mutator transaction binding the contract method 0xccc362e7.
//
// Solidity: function refund(address _migrator, address _account, address _token, uint256 _usdgAmount) returns()
func (_MigrationHandler *MigrationHandlerTransactor) Refund(opts *bind.TransactOpts, _migrator common.Address, _account common.Address, _token common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.contract.Transact(opts, "refund", _migrator, _account, _token, _usdgAmount)
}

// Refund is a paid mutator transaction binding the contract method 0xccc362e7.
//
// Solidity: function refund(address _migrator, address _account, address _token, uint256 _usdgAmount) returns()
func (_MigrationHandler *MigrationHandlerSession) Refund(_migrator common.Address, _account common.Address, _token common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.Contract.Refund(&_MigrationHandler.TransactOpts, _migrator, _account, _token, _usdgAmount)
}

// Refund is a paid mutator transaction binding the contract method 0xccc362e7.
//
// Solidity: function refund(address _migrator, address _account, address _token, uint256 _usdgAmount) returns()
func (_MigrationHandler *MigrationHandlerTransactorSession) Refund(_migrator common.Address, _account common.Address, _token common.Address, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.Contract.Refund(&_MigrationHandler.TransactOpts, _migrator, _account, _token, _usdgAmount)
}

// Swap is a paid mutator transaction binding the contract method 0xc9ec9248.
//
// Solidity: function swap(address _migrator, uint256 _gmtAmountForUsdg, uint256 _xgmtAmountForUsdg, uint256 _gmtAmountForBusd) returns()
func (_MigrationHandler *MigrationHandlerTransactor) Swap(opts *bind.TransactOpts, _migrator common.Address, _gmtAmountForUsdg *big.Int, _xgmtAmountForUsdg *big.Int, _gmtAmountForBusd *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.contract.Transact(opts, "swap", _migrator, _gmtAmountForUsdg, _xgmtAmountForUsdg, _gmtAmountForBusd)
}

// Swap is a paid mutator transaction binding the contract method 0xc9ec9248.
//
// Solidity: function swap(address _migrator, uint256 _gmtAmountForUsdg, uint256 _xgmtAmountForUsdg, uint256 _gmtAmountForBusd) returns()
func (_MigrationHandler *MigrationHandlerSession) Swap(_migrator common.Address, _gmtAmountForUsdg *big.Int, _xgmtAmountForUsdg *big.Int, _gmtAmountForBusd *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.Contract.Swap(&_MigrationHandler.TransactOpts, _migrator, _gmtAmountForUsdg, _xgmtAmountForUsdg, _gmtAmountForBusd)
}

// Swap is a paid mutator transaction binding the contract method 0xc9ec9248.
//
// Solidity: function swap(address _migrator, uint256 _gmtAmountForUsdg, uint256 _xgmtAmountForUsdg, uint256 _gmtAmountForBusd) returns()
func (_MigrationHandler *MigrationHandlerTransactorSession) Swap(_migrator common.Address, _gmtAmountForUsdg *big.Int, _xgmtAmountForUsdg *big.Int, _gmtAmountForBusd *big.Int) (*types.Transaction, error) {
	return _MigrationHandler.Contract.Swap(&_MigrationHandler.TransactOpts, _migrator, _gmtAmountForUsdg, _xgmtAmountForUsdg, _gmtAmountForBusd)
}
