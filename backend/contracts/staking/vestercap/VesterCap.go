// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vestercap

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

// VesterCapMetaData contains all meta data concerning the VesterCap contract.
var VesterCapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_extendedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxBoostBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bnGmxToEsGmxConversionDivisor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmxToEsGmxConversionDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extendedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isUpdateCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBoostBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"syncFeeGmxTrackerBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"updateBnGmxForAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VesterCapABI is the input ABI used to generate the binding from.
// Deprecated: Use VesterCapMetaData.ABI instead.
var VesterCapABI = VesterCapMetaData.ABI

// VesterCap is an auto generated Go binding around an Ethereum contract.
type VesterCap struct {
	VesterCapCaller     // Read-only binding to the contract
	VesterCapTransactor // Write-only binding to the contract
	VesterCapFilterer   // Log filterer for contract events
}

// VesterCapCaller is an auto generated read-only Go binding around an Ethereum contract.
type VesterCapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VesterCapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VesterCapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VesterCapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VesterCapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VesterCapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VesterCapSession struct {
	Contract     *VesterCap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VesterCapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VesterCapCallerSession struct {
	Contract *VesterCapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VesterCapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VesterCapTransactorSession struct {
	Contract     *VesterCapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VesterCapRaw is an auto generated low-level Go binding around an Ethereum contract.
type VesterCapRaw struct {
	Contract *VesterCap // Generic contract binding to access the raw methods on
}

// VesterCapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VesterCapCallerRaw struct {
	Contract *VesterCapCaller // Generic read-only contract binding to access the raw methods on
}

// VesterCapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VesterCapTransactorRaw struct {
	Contract *VesterCapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesterCap creates a new instance of VesterCap, bound to a specific deployed contract.
func NewVesterCap(address common.Address, backend bind.ContractBackend) (*VesterCap, error) {
	contract, err := bindVesterCap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VesterCap{VesterCapCaller: VesterCapCaller{contract: contract}, VesterCapTransactor: VesterCapTransactor{contract: contract}, VesterCapFilterer: VesterCapFilterer{contract: contract}}, nil
}

// NewVesterCapCaller creates a new read-only instance of VesterCap, bound to a specific deployed contract.
func NewVesterCapCaller(address common.Address, caller bind.ContractCaller) (*VesterCapCaller, error) {
	contract, err := bindVesterCap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VesterCapCaller{contract: contract}, nil
}

// NewVesterCapTransactor creates a new write-only instance of VesterCap, bound to a specific deployed contract.
func NewVesterCapTransactor(address common.Address, transactor bind.ContractTransactor) (*VesterCapTransactor, error) {
	contract, err := bindVesterCap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VesterCapTransactor{contract: contract}, nil
}

// NewVesterCapFilterer creates a new log filterer instance of VesterCap, bound to a specific deployed contract.
func NewVesterCapFilterer(address common.Address, filterer bind.ContractFilterer) (*VesterCapFilterer, error) {
	contract, err := bindVesterCap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VesterCapFilterer{contract: contract}, nil
}

// bindVesterCap binds a generic wrapper to an already deployed contract.
func bindVesterCap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VesterCapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VesterCap *VesterCapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VesterCap.Contract.VesterCapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VesterCap *VesterCapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VesterCap.Contract.VesterCapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VesterCap *VesterCapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VesterCap.Contract.VesterCapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VesterCap *VesterCapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VesterCap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VesterCap *VesterCapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VesterCap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VesterCap *VesterCapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VesterCap.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VesterCap *VesterCapCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VesterCap *VesterCapSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VesterCap.Contract.BASISPOINTSDIVISOR(&_VesterCap.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VesterCap *VesterCapCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VesterCap.Contract.BASISPOINTSDIVISOR(&_VesterCap.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_VesterCap *VesterCapCaller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_VesterCap *VesterCapSession) BnGmx() (common.Address, error) {
	return _VesterCap.Contract.BnGmx(&_VesterCap.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_VesterCap *VesterCapCallerSession) BnGmx() (common.Address, error) {
	return _VesterCap.Contract.BnGmx(&_VesterCap.CallOpts)
}

// BnGmxToEsGmxConversionDivisor is a free data retrieval call binding the contract method 0x40d33ea9.
//
// Solidity: function bnGmxToEsGmxConversionDivisor() view returns(uint256)
func (_VesterCap *VesterCapCaller) BnGmxToEsGmxConversionDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "bnGmxToEsGmxConversionDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BnGmxToEsGmxConversionDivisor is a free data retrieval call binding the contract method 0x40d33ea9.
//
// Solidity: function bnGmxToEsGmxConversionDivisor() view returns(uint256)
func (_VesterCap *VesterCapSession) BnGmxToEsGmxConversionDivisor() (*big.Int, error) {
	return _VesterCap.Contract.BnGmxToEsGmxConversionDivisor(&_VesterCap.CallOpts)
}

// BnGmxToEsGmxConversionDivisor is a free data retrieval call binding the contract method 0x40d33ea9.
//
// Solidity: function bnGmxToEsGmxConversionDivisor() view returns(uint256)
func (_VesterCap *VesterCapCallerSession) BnGmxToEsGmxConversionDivisor() (*big.Int, error) {
	return _VesterCap.Contract.BnGmxToEsGmxConversionDivisor(&_VesterCap.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_VesterCap *VesterCapCaller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_VesterCap *VesterCapSession) BonusGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.BonusGmxTracker(&_VesterCap.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_VesterCap *VesterCapCallerSession) BonusGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.BonusGmxTracker(&_VesterCap.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_VesterCap *VesterCapCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_VesterCap *VesterCapSession) EsGmx() (common.Address, error) {
	return _VesterCap.Contract.EsGmx(&_VesterCap.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_VesterCap *VesterCapCallerSession) EsGmx() (common.Address, error) {
	return _VesterCap.Contract.EsGmx(&_VesterCap.CallOpts)
}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_VesterCap *VesterCapCaller) ExtendedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "extendedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_VesterCap *VesterCapSession) ExtendedGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.ExtendedGmxTracker(&_VesterCap.CallOpts)
}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_VesterCap *VesterCapCallerSession) ExtendedGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.ExtendedGmxTracker(&_VesterCap.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_VesterCap *VesterCapCaller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_VesterCap *VesterCapSession) FeeGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.FeeGmxTracker(&_VesterCap.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_VesterCap *VesterCapCallerSession) FeeGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.FeeGmxTracker(&_VesterCap.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_VesterCap *VesterCapCaller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_VesterCap *VesterCapSession) GmxVester() (common.Address, error) {
	return _VesterCap.Contract.GmxVester(&_VesterCap.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_VesterCap *VesterCapCallerSession) GmxVester() (common.Address, error) {
	return _VesterCap.Contract.GmxVester(&_VesterCap.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VesterCap *VesterCapCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VesterCap *VesterCapSession) Gov() (common.Address, error) {
	return _VesterCap.Contract.Gov(&_VesterCap.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VesterCap *VesterCapCallerSession) Gov() (common.Address, error) {
	return _VesterCap.Contract.Gov(&_VesterCap.CallOpts)
}

// IsUpdateCompleted is a free data retrieval call binding the contract method 0xaeb3cdec.
//
// Solidity: function isUpdateCompleted(address ) view returns(bool)
func (_VesterCap *VesterCapCaller) IsUpdateCompleted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "isUpdateCompleted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpdateCompleted is a free data retrieval call binding the contract method 0xaeb3cdec.
//
// Solidity: function isUpdateCompleted(address ) view returns(bool)
func (_VesterCap *VesterCapSession) IsUpdateCompleted(arg0 common.Address) (bool, error) {
	return _VesterCap.Contract.IsUpdateCompleted(&_VesterCap.CallOpts, arg0)
}

// IsUpdateCompleted is a free data retrieval call binding the contract method 0xaeb3cdec.
//
// Solidity: function isUpdateCompleted(address ) view returns(bool)
func (_VesterCap *VesterCapCallerSession) IsUpdateCompleted(arg0 common.Address) (bool, error) {
	return _VesterCap.Contract.IsUpdateCompleted(&_VesterCap.CallOpts, arg0)
}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_VesterCap *VesterCapCaller) MaxBoostBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "maxBoostBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_VesterCap *VesterCapSession) MaxBoostBasisPoints() (*big.Int, error) {
	return _VesterCap.Contract.MaxBoostBasisPoints(&_VesterCap.CallOpts)
}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_VesterCap *VesterCapCallerSession) MaxBoostBasisPoints() (*big.Int, error) {
	return _VesterCap.Contract.MaxBoostBasisPoints(&_VesterCap.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_VesterCap *VesterCapCaller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VesterCap.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_VesterCap *VesterCapSession) StakedGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.StakedGmxTracker(&_VesterCap.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_VesterCap *VesterCapCallerSession) StakedGmxTracker() (common.Address, error) {
	return _VesterCap.Contract.StakedGmxTracker(&_VesterCap.CallOpts)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VesterCap *VesterCapTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _VesterCap.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VesterCap *VesterCapSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VesterCap.Contract.SetGov(&_VesterCap.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VesterCap *VesterCapTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VesterCap.Contract.SetGov(&_VesterCap.TransactOpts, _gov)
}

// SyncFeeGmxTrackerBalance is a paid mutator transaction binding the contract method 0xe21c36ea.
//
// Solidity: function syncFeeGmxTrackerBalance(address _account) returns()
func (_VesterCap *VesterCapTransactor) SyncFeeGmxTrackerBalance(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _VesterCap.contract.Transact(opts, "syncFeeGmxTrackerBalance", _account)
}

// SyncFeeGmxTrackerBalance is a paid mutator transaction binding the contract method 0xe21c36ea.
//
// Solidity: function syncFeeGmxTrackerBalance(address _account) returns()
func (_VesterCap *VesterCapSession) SyncFeeGmxTrackerBalance(_account common.Address) (*types.Transaction, error) {
	return _VesterCap.Contract.SyncFeeGmxTrackerBalance(&_VesterCap.TransactOpts, _account)
}

// SyncFeeGmxTrackerBalance is a paid mutator transaction binding the contract method 0xe21c36ea.
//
// Solidity: function syncFeeGmxTrackerBalance(address _account) returns()
func (_VesterCap *VesterCapTransactorSession) SyncFeeGmxTrackerBalance(_account common.Address) (*types.Transaction, error) {
	return _VesterCap.Contract.SyncFeeGmxTrackerBalance(&_VesterCap.TransactOpts, _account)
}

// UpdateBnGmxForAccounts is a paid mutator transaction binding the contract method 0x82f94887.
//
// Solidity: function updateBnGmxForAccounts(address[] _accounts) returns()
func (_VesterCap *VesterCapTransactor) UpdateBnGmxForAccounts(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _VesterCap.contract.Transact(opts, "updateBnGmxForAccounts", _accounts)
}

// UpdateBnGmxForAccounts is a paid mutator transaction binding the contract method 0x82f94887.
//
// Solidity: function updateBnGmxForAccounts(address[] _accounts) returns()
func (_VesterCap *VesterCapSession) UpdateBnGmxForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _VesterCap.Contract.UpdateBnGmxForAccounts(&_VesterCap.TransactOpts, _accounts)
}

// UpdateBnGmxForAccounts is a paid mutator transaction binding the contract method 0x82f94887.
//
// Solidity: function updateBnGmxForAccounts(address[] _accounts) returns()
func (_VesterCap *VesterCapTransactorSession) UpdateBnGmxForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _VesterCap.Contract.UpdateBnGmxForAccounts(&_VesterCap.TransactOpts, _accounts)
}
