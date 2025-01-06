// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakedglpmigrator

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

// StakedGlpMigratorMetaData contains all meta data concerning the StakedGlpMigrator contract.
var StakedGlpMigratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"disable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakedGlpMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use StakedGlpMigratorMetaData.ABI instead.
var StakedGlpMigratorABI = StakedGlpMigratorMetaData.ABI

// StakedGlpMigrator is an auto generated Go binding around an Ethereum contract.
type StakedGlpMigrator struct {
	StakedGlpMigratorCaller     // Read-only binding to the contract
	StakedGlpMigratorTransactor // Write-only binding to the contract
	StakedGlpMigratorFilterer   // Log filterer for contract events
}

// StakedGlpMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedGlpMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedGlpMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedGlpMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedGlpMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedGlpMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedGlpMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedGlpMigratorSession struct {
	Contract     *StakedGlpMigrator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakedGlpMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedGlpMigratorCallerSession struct {
	Contract *StakedGlpMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StakedGlpMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedGlpMigratorTransactorSession struct {
	Contract     *StakedGlpMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StakedGlpMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedGlpMigratorRaw struct {
	Contract *StakedGlpMigrator // Generic contract binding to access the raw methods on
}

// StakedGlpMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedGlpMigratorCallerRaw struct {
	Contract *StakedGlpMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// StakedGlpMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedGlpMigratorTransactorRaw struct {
	Contract *StakedGlpMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedGlpMigrator creates a new instance of StakedGlpMigrator, bound to a specific deployed contract.
func NewStakedGlpMigrator(address common.Address, backend bind.ContractBackend) (*StakedGlpMigrator, error) {
	contract, err := bindStakedGlpMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedGlpMigrator{StakedGlpMigratorCaller: StakedGlpMigratorCaller{contract: contract}, StakedGlpMigratorTransactor: StakedGlpMigratorTransactor{contract: contract}, StakedGlpMigratorFilterer: StakedGlpMigratorFilterer{contract: contract}}, nil
}

// NewStakedGlpMigratorCaller creates a new read-only instance of StakedGlpMigrator, bound to a specific deployed contract.
func NewStakedGlpMigratorCaller(address common.Address, caller bind.ContractCaller) (*StakedGlpMigratorCaller, error) {
	contract, err := bindStakedGlpMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedGlpMigratorCaller{contract: contract}, nil
}

// NewStakedGlpMigratorTransactor creates a new write-only instance of StakedGlpMigrator, bound to a specific deployed contract.
func NewStakedGlpMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedGlpMigratorTransactor, error) {
	contract, err := bindStakedGlpMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedGlpMigratorTransactor{contract: contract}, nil
}

// NewStakedGlpMigratorFilterer creates a new log filterer instance of StakedGlpMigrator, bound to a specific deployed contract.
func NewStakedGlpMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedGlpMigratorFilterer, error) {
	contract, err := bindStakedGlpMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedGlpMigratorFilterer{contract: contract}, nil
}

// bindStakedGlpMigrator binds a generic wrapper to an already deployed contract.
func bindStakedGlpMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakedGlpMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedGlpMigrator *StakedGlpMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedGlpMigrator.Contract.StakedGlpMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedGlpMigrator *StakedGlpMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.StakedGlpMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedGlpMigrator *StakedGlpMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.StakedGlpMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedGlpMigrator *StakedGlpMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedGlpMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedGlpMigrator *StakedGlpMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedGlpMigrator *StakedGlpMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.contract.Transact(opts, method, params...)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlpMigrator.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorSession) FeeGlpTracker() (common.Address, error) {
	return _StakedGlpMigrator.Contract.FeeGlpTracker(&_StakedGlpMigrator.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCallerSession) FeeGlpTracker() (common.Address, error) {
	return _StakedGlpMigrator.Contract.FeeGlpTracker(&_StakedGlpMigrator.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlpMigrator.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorSession) Glp() (common.Address, error) {
	return _StakedGlpMigrator.Contract.Glp(&_StakedGlpMigrator.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCallerSession) Glp() (common.Address, error) {
	return _StakedGlpMigrator.Contract.Glp(&_StakedGlpMigrator.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlpMigrator.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorSession) Gov() (common.Address, error) {
	return _StakedGlpMigrator.Contract.Gov(&_StakedGlpMigrator.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCallerSession) Gov() (common.Address, error) {
	return _StakedGlpMigrator.Contract.Gov(&_StakedGlpMigrator.CallOpts)
}

// IsEnabled is a free data retrieval call binding the contract method 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (_StakedGlpMigrator *StakedGlpMigratorCaller) IsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedGlpMigrator.contract.Call(opts, &out, "isEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnabled is a free data retrieval call binding the contract method 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (_StakedGlpMigrator *StakedGlpMigratorSession) IsEnabled() (bool, error) {
	return _StakedGlpMigrator.Contract.IsEnabled(&_StakedGlpMigrator.CallOpts)
}

// IsEnabled is a free data retrieval call binding the contract method 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (_StakedGlpMigrator *StakedGlpMigratorCallerSession) IsEnabled() (bool, error) {
	return _StakedGlpMigrator.Contract.IsEnabled(&_StakedGlpMigrator.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCaller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlpMigrator.contract.Call(opts, &out, "sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorSession) Sender() (common.Address, error) {
	return _StakedGlpMigrator.Contract.Sender(&_StakedGlpMigrator.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCallerSession) Sender() (common.Address, error) {
	return _StakedGlpMigrator.Contract.Sender(&_StakedGlpMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedGlpMigrator.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorSession) StakedGlpTracker() (common.Address, error) {
	return _StakedGlpMigrator.Contract.StakedGlpTracker(&_StakedGlpMigrator.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_StakedGlpMigrator *StakedGlpMigratorCallerSession) StakedGlpTracker() (common.Address, error) {
	return _StakedGlpMigrator.Contract.StakedGlpTracker(&_StakedGlpMigrator.CallOpts)
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_StakedGlpMigrator *StakedGlpMigratorTransactor) Disable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedGlpMigrator.contract.Transact(opts, "disable")
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_StakedGlpMigrator *StakedGlpMigratorSession) Disable() (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.Disable(&_StakedGlpMigrator.TransactOpts)
}

// Disable is a paid mutator transaction binding the contract method 0x2f2770db.
//
// Solidity: function disable() returns()
func (_StakedGlpMigrator *StakedGlpMigratorTransactorSession) Disable() (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.Disable(&_StakedGlpMigrator.TransactOpts)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_StakedGlpMigrator *StakedGlpMigratorTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _StakedGlpMigrator.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_StakedGlpMigrator *StakedGlpMigratorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.SetGov(&_StakedGlpMigrator.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_StakedGlpMigrator *StakedGlpMigratorTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.SetGov(&_StakedGlpMigrator.TransactOpts, _gov)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns()
func (_StakedGlpMigrator *StakedGlpMigratorTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlpMigrator.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns()
func (_StakedGlpMigrator *StakedGlpMigratorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.Transfer(&_StakedGlpMigrator.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns()
func (_StakedGlpMigrator *StakedGlpMigratorTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedGlpMigrator.Contract.Transfer(&_StakedGlpMigrator.TransactOpts, _recipient, _amount)
}
