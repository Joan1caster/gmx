// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakedglp

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

// StakedGlpMetaData contains all meta data concerning the StakedGlp contract.
var StakedGlpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_glp\",\"type\":\"address\"},{\"internalType\":\"contractIGlpManager\",\"name\":\"_glpManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpManager\",\"outputs\":[{\"internalType\":\"contractIGlpManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakedGlpABI is the input ABI used to generate the binding from.
// Deprecated: Use StakedGlpMetaData.ABI instead.
var StakedGlpABI = StakedGlpMetaData.ABI

// StakedGlp is an auto generated Go binding around an Ethereum contract.
type StakedGlp struct {
	StakedGlpCaller     // Read-only binding to the contract
	StakedGlpTransactor // Write-only binding to the contract
	StakedGlpFilterer   // Log filterer for contract events
}

// StakedGlpCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedGlpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedGlpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedGlpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedGlpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedGlpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedGlpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedGlpSession struct {
	Contract     *StakedGlp        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakedGlpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedGlpCallerSession struct {
	Contract *StakedGlpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StakedGlpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedGlpTransactorSession struct {
	Contract     *StakedGlpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StakedGlpRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedGlpRaw struct {
	Contract *StakedGlp // Generic contract binding to access the raw methods on
}

// StakedGlpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedGlpCallerRaw struct {
	Contract *StakedGlpCaller // Generic read-only contract binding to access the raw methods on
}

// StakedGlpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedGlpTransactorRaw struct {
	Contract *StakedGlpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedGlp creates a new instance of StakedGlp, bound to a specific deployed contract.
func NewStakedGlp(address common.Address, backend bind.ContractBackend) (*StakedGlp, error) {
	contract, err := bindStakedGlp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedGlp{StakedGlpCaller: StakedGlpCaller{contract: contract}, StakedGlpTransactor: StakedGlpTransactor{contract: contract}, StakedGlpFilterer: StakedGlpFilterer{contract: contract}}, nil
}

// NewStakedGlpCaller creates a new read-only instance of StakedGlp, bound to a specific deployed contract.
func NewStakedGlpCaller(address common.Address, caller bind.ContractCaller) (*StakedGlpCaller, error) {
	contract, err := bindStakedGlp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedGlpCaller{contract: contract}, nil
}

// NewStakedGlpTransactor creates a new write-only instance of StakedGlp, bound to a specific deployed contract.
func NewStakedGlpTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedGlpTransactor, error) {
	contract, err := bindStakedGlp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedGlpTransactor{contract: contract}, nil
}

// NewStakedGlpFilterer creates a new log filterer instance of StakedGlp, bound to a specific deployed contract.
func NewStakedGlpFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedGlpFilterer, error) {
	contract, err := bindStakedGlp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedGlpFilterer{contract: contract}, nil
}

// bindStakedGlp binds a generic wrapper to an already deployed contract.
func bindStakedGlp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakedGlpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedGlp *StakedGlpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedGlp.Contract.StakedGlpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedGlp *StakedGlpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedGlp.Contract.StakedGlpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedGlp *StakedGlpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedGlp.Contract.StakedGlpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedGlp *StakedGlpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedGlp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedGlp *StakedGlpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedGlp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedGlp *StakedGlpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedGlp.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StakedGlp *StakedGlpCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StakedGlp *StakedGlpSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StakedGlp.Contract.Allowance(&_StakedGlp.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StakedGlp *StakedGlpCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StakedGlp.Contract.Allowance(&_StakedGlp.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_StakedGlp *StakedGlpCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_StakedGlp *StakedGlpSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StakedGlp.Contract.Allowances(&_StakedGlp.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_StakedGlp *StakedGlpCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StakedGlp.Contract.Allowances(&_StakedGlp.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StakedGlp *StakedGlpCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StakedGlp *StakedGlpSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _StakedGlp.Contract.BalanceOf(&_StakedGlp.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StakedGlp *StakedGlpCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _StakedGlp.Contract.BalanceOf(&_StakedGlp.CallOpts, _account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakedGlp *StakedGlpCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakedGlp *StakedGlpSession) Decimals() (uint8, error) {
	return _StakedGlp.Contract.Decimals(&_StakedGlp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakedGlp *StakedGlpCallerSession) Decimals() (uint8, error) {
	return _StakedGlp.Contract.Decimals(&_StakedGlp.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StakedGlp *StakedGlpCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StakedGlp *StakedGlpSession) FeeGlpTracker() (common.Address, error) {
	return _StakedGlp.Contract.FeeGlpTracker(&_StakedGlp.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StakedGlp *StakedGlpCallerSession) FeeGlpTracker() (common.Address, error) {
	return _StakedGlp.Contract.FeeGlpTracker(&_StakedGlp.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_StakedGlp *StakedGlpCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_StakedGlp *StakedGlpSession) Glp() (common.Address, error) {
	return _StakedGlp.Contract.Glp(&_StakedGlp.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_StakedGlp *StakedGlpCallerSession) Glp() (common.Address, error) {
	return _StakedGlp.Contract.Glp(&_StakedGlp.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_StakedGlp *StakedGlpCaller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_StakedGlp *StakedGlpSession) GlpManager() (common.Address, error) {
	return _StakedGlp.Contract.GlpManager(&_StakedGlp.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_StakedGlp *StakedGlpCallerSession) GlpManager() (common.Address, error) {
	return _StakedGlp.Contract.GlpManager(&_StakedGlp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedGlp *StakedGlpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedGlp *StakedGlpSession) Name() (string, error) {
	return _StakedGlp.Contract.Name(&_StakedGlp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedGlp *StakedGlpCallerSession) Name() (string, error) {
	return _StakedGlp.Contract.Name(&_StakedGlp.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StakedGlp *StakedGlpCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StakedGlp *StakedGlpSession) StakedGlpTracker() (common.Address, error) {
	return _StakedGlp.Contract.StakedGlpTracker(&_StakedGlp.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StakedGlp *StakedGlpCallerSession) StakedGlpTracker() (common.Address, error) {
	return _StakedGlp.Contract.StakedGlpTracker(&_StakedGlp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedGlp *StakedGlpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedGlp *StakedGlpSession) Symbol() (string, error) {
	return _StakedGlp.Contract.Symbol(&_StakedGlp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedGlp *StakedGlpCallerSession) Symbol() (string, error) {
	return _StakedGlp.Contract.Symbol(&_StakedGlp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedGlp *StakedGlpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedGlp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedGlp *StakedGlpSession) TotalSupply() (*big.Int, error) {
	return _StakedGlp.Contract.TotalSupply(&_StakedGlp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedGlp *StakedGlpCallerSession) TotalSupply() (*big.Int, error) {
	return _StakedGlp.Contract.TotalSupply(&_StakedGlp.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.Contract.Approve(&_StakedGlp.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.Contract.Approve(&_StakedGlp.TransactOpts, _spender, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.Contract.Transfer(&_StakedGlp.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.Contract.Transfer(&_StakedGlp.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.Contract.TransferFrom(&_StakedGlp.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StakedGlp *StakedGlpTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlp.Contract.TransferFrom(&_StakedGlp.TransactOpts, _sender, _recipient, _amount)
}

// StakedGlpApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakedGlp contract.
type StakedGlpApprovalIterator struct {
	Event *StakedGlpApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakedGlpApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedGlpApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakedGlpApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakedGlpApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedGlpApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedGlpApproval represents a Approval event raised by the StakedGlp contract.
type StakedGlpApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedGlp *StakedGlpFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StakedGlpApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedGlp.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StakedGlpApprovalIterator{contract: _StakedGlp.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedGlp *StakedGlpFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakedGlpApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedGlp.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedGlpApproval)
				if err := _StakedGlp.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedGlp *StakedGlpFilterer) ParseApproval(log types.Log) (*StakedGlpApproval, error) {
	event := new(StakedGlpApproval)
	if err := _StakedGlp.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
