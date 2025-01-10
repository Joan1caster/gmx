// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basepositionmanager

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

// BasePositionManagerMetaData contains all meta data concerning the BasePositionManager contract.
var BasePositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortsTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_depositFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marginFeeBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"referralCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"DecreasePositionReferral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marginFeeBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"referralCode\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"IncreasePositionReferral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"SetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFee\",\"type\":\"uint256\"}],\"name\":\"SetDepositFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethTransferGasLimit\",\"type\":\"uint256\"}],\"name\":\"SetEthTransferGasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"increasePositionBufferBps\",\"type\":\"uint256\"}],\"name\":\"SetIncreasePositionBufferBps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"longSizes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"shortSizes\",\"type\":\"uint256[]\"}],\"name\":\"SetMaxGlobalSizes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referralStorage\",\"type\":\"address\"}],\"name\":\"SetReferralStorage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFees\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethTransferGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increasePositionBufferBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxGlobalLongSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxGlobalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referralStorage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_depositFee\",\"type\":\"uint256\"}],\"name\":\"setDepositFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethTransferGasLimit\",\"type\":\"uint256\"}],\"name\":\"setEthTransferGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_increasePositionBufferBps\",\"type\":\"uint256\"}],\"name\":\"setIncreasePositionBufferBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_longSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_shortSizes\",\"type\":\"uint256[]\"}],\"name\":\"setMaxGlobalSizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referralStorage\",\"type\":\"address\"}],\"name\":\"setReferralStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shortsTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BasePositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BasePositionManagerMetaData.ABI instead.
var BasePositionManagerABI = BasePositionManagerMetaData.ABI

// BasePositionManager is an auto generated Go binding around an Ethereum contract.
type BasePositionManager struct {
	BasePositionManagerCaller     // Read-only binding to the contract
	BasePositionManagerTransactor // Write-only binding to the contract
	BasePositionManagerFilterer   // Log filterer for contract events
}

// BasePositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasePositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasePositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasePositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasePositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasePositionManagerSession struct {
	Contract     *BasePositionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BasePositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasePositionManagerCallerSession struct {
	Contract *BasePositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BasePositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasePositionManagerTransactorSession struct {
	Contract     *BasePositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BasePositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasePositionManagerRaw struct {
	Contract *BasePositionManager // Generic contract binding to access the raw methods on
}

// BasePositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasePositionManagerCallerRaw struct {
	Contract *BasePositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BasePositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasePositionManagerTransactorRaw struct {
	Contract *BasePositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasePositionManager creates a new instance of BasePositionManager, bound to a specific deployed contract.
func NewBasePositionManager(address common.Address, backend bind.ContractBackend) (*BasePositionManager, error) {
	contract, err := bindBasePositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasePositionManager{BasePositionManagerCaller: BasePositionManagerCaller{contract: contract}, BasePositionManagerTransactor: BasePositionManagerTransactor{contract: contract}, BasePositionManagerFilterer: BasePositionManagerFilterer{contract: contract}}, nil
}

// NewBasePositionManagerCaller creates a new read-only instance of BasePositionManager, bound to a specific deployed contract.
func NewBasePositionManagerCaller(address common.Address, caller bind.ContractCaller) (*BasePositionManagerCaller, error) {
	contract, err := bindBasePositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerCaller{contract: contract}, nil
}

// NewBasePositionManagerTransactor creates a new write-only instance of BasePositionManager, bound to a specific deployed contract.
func NewBasePositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BasePositionManagerTransactor, error) {
	contract, err := bindBasePositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerTransactor{contract: contract}, nil
}

// NewBasePositionManagerFilterer creates a new log filterer instance of BasePositionManager, bound to a specific deployed contract.
func NewBasePositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BasePositionManagerFilterer, error) {
	contract, err := bindBasePositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerFilterer{contract: contract}, nil
}

// bindBasePositionManager binds a generic wrapper to an already deployed contract.
func bindBasePositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasePositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasePositionManager *BasePositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasePositionManager.Contract.BasePositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasePositionManager *BasePositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePositionManager.Contract.BasePositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasePositionManager *BasePositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasePositionManager.Contract.BasePositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasePositionManager *BasePositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasePositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasePositionManager *BasePositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasePositionManager *BasePositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasePositionManager.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_BasePositionManager *BasePositionManagerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _BasePositionManager.Contract.BASISPOINTSDIVISOR(&_BasePositionManager.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _BasePositionManager.Contract.BASISPOINTSDIVISOR(&_BasePositionManager.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BasePositionManager *BasePositionManagerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BasePositionManager *BasePositionManagerSession) Admin() (common.Address, error) {
	return _BasePositionManager.Contract.Admin(&_BasePositionManager.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BasePositionManager *BasePositionManagerCallerSession) Admin() (common.Address, error) {
	return _BasePositionManager.Contract.Admin(&_BasePositionManager.CallOpts)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCaller) DepositFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "depositFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_BasePositionManager *BasePositionManagerSession) DepositFee() (*big.Int, error) {
	return _BasePositionManager.Contract.DepositFee(&_BasePositionManager.CallOpts)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCallerSession) DepositFee() (*big.Int, error) {
	return _BasePositionManager.Contract.DepositFee(&_BasePositionManager.CallOpts)
}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCaller) EthTransferGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "ethTransferGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_BasePositionManager *BasePositionManagerSession) EthTransferGasLimit() (*big.Int, error) {
	return _BasePositionManager.Contract.EthTransferGasLimit(&_BasePositionManager.CallOpts)
}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCallerSession) EthTransferGasLimit() (*big.Int, error) {
	return _BasePositionManager.Contract.EthTransferGasLimit(&_BasePositionManager.CallOpts)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerCaller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _BasePositionManager.Contract.FeeReserves(&_BasePositionManager.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerCallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _BasePositionManager.Contract.FeeReserves(&_BasePositionManager.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BasePositionManager *BasePositionManagerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BasePositionManager *BasePositionManagerSession) Gov() (common.Address, error) {
	return _BasePositionManager.Contract.Gov(&_BasePositionManager.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BasePositionManager *BasePositionManagerCallerSession) Gov() (common.Address, error) {
	return _BasePositionManager.Contract.Gov(&_BasePositionManager.CallOpts)
}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCaller) IncreasePositionBufferBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "increasePositionBufferBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_BasePositionManager *BasePositionManagerSession) IncreasePositionBufferBps() (*big.Int, error) {
	return _BasePositionManager.Contract.IncreasePositionBufferBps(&_BasePositionManager.CallOpts)
}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_BasePositionManager *BasePositionManagerCallerSession) IncreasePositionBufferBps() (*big.Int, error) {
	return _BasePositionManager.Contract.IncreasePositionBufferBps(&_BasePositionManager.CallOpts)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerCaller) MaxGlobalLongSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "maxGlobalLongSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerSession) MaxGlobalLongSizes(arg0 common.Address) (*big.Int, error) {
	return _BasePositionManager.Contract.MaxGlobalLongSizes(&_BasePositionManager.CallOpts, arg0)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerCallerSession) MaxGlobalLongSizes(arg0 common.Address) (*big.Int, error) {
	return _BasePositionManager.Contract.MaxGlobalLongSizes(&_BasePositionManager.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerCaller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _BasePositionManager.Contract.MaxGlobalShortSizes(&_BasePositionManager.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_BasePositionManager *BasePositionManagerCallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _BasePositionManager.Contract.MaxGlobalShortSizes(&_BasePositionManager.CallOpts, arg0)
}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_BasePositionManager *BasePositionManagerCaller) ReferralStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "referralStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_BasePositionManager *BasePositionManagerSession) ReferralStorage() (common.Address, error) {
	return _BasePositionManager.Contract.ReferralStorage(&_BasePositionManager.CallOpts)
}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_BasePositionManager *BasePositionManagerCallerSession) ReferralStorage() (common.Address, error) {
	return _BasePositionManager.Contract.ReferralStorage(&_BasePositionManager.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BasePositionManager *BasePositionManagerCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BasePositionManager *BasePositionManagerSession) Router() (common.Address, error) {
	return _BasePositionManager.Contract.Router(&_BasePositionManager.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BasePositionManager *BasePositionManagerCallerSession) Router() (common.Address, error) {
	return _BasePositionManager.Contract.Router(&_BasePositionManager.CallOpts)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_BasePositionManager *BasePositionManagerCaller) ShortsTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "shortsTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_BasePositionManager *BasePositionManagerSession) ShortsTracker() (common.Address, error) {
	return _BasePositionManager.Contract.ShortsTracker(&_BasePositionManager.CallOpts)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_BasePositionManager *BasePositionManagerCallerSession) ShortsTracker() (common.Address, error) {
	return _BasePositionManager.Contract.ShortsTracker(&_BasePositionManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BasePositionManager *BasePositionManagerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BasePositionManager *BasePositionManagerSession) Vault() (common.Address, error) {
	return _BasePositionManager.Contract.Vault(&_BasePositionManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_BasePositionManager *BasePositionManagerCallerSession) Vault() (common.Address, error) {
	return _BasePositionManager.Contract.Vault(&_BasePositionManager.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BasePositionManager *BasePositionManagerCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasePositionManager.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BasePositionManager *BasePositionManagerSession) Weth() (common.Address, error) {
	return _BasePositionManager.Contract.Weth(&_BasePositionManager.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_BasePositionManager *BasePositionManagerCallerSession) Weth() (common.Address, error) {
	return _BasePositionManager.Contract.Weth(&_BasePositionManager.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_BasePositionManager *BasePositionManagerTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "approve", _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_BasePositionManager *BasePositionManagerSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.Approve(&_BasePositionManager.TransactOpts, _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.Approve(&_BasePositionManager.TransactOpts, _token, _spender, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SendValue(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "sendValue", _receiver, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_BasePositionManager *BasePositionManagerSession) SendValue(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SendValue(&_BasePositionManager.TransactOpts, _receiver, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SendValue(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SendValue(&_BasePositionManager.TransactOpts, _receiver, _amount)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_BasePositionManager *BasePositionManagerSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetAdmin(&_BasePositionManager.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetAdmin(&_BasePositionManager.TransactOpts, _admin)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SetDepositFee(opts *bind.TransactOpts, _depositFee *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "setDepositFee", _depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_BasePositionManager *BasePositionManagerSession) SetDepositFee(_depositFee *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetDepositFee(&_BasePositionManager.TransactOpts, _depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SetDepositFee(_depositFee *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetDepositFee(&_BasePositionManager.TransactOpts, _depositFee)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SetEthTransferGasLimit(opts *bind.TransactOpts, _ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "setEthTransferGasLimit", _ethTransferGasLimit)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_BasePositionManager *BasePositionManagerSession) SetEthTransferGasLimit(_ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetEthTransferGasLimit(&_BasePositionManager.TransactOpts, _ethTransferGasLimit)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SetEthTransferGasLimit(_ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetEthTransferGasLimit(&_BasePositionManager.TransactOpts, _ethTransferGasLimit)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BasePositionManager *BasePositionManagerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetGov(&_BasePositionManager.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetGov(&_BasePositionManager.TransactOpts, _gov)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SetIncreasePositionBufferBps(opts *bind.TransactOpts, _increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "setIncreasePositionBufferBps", _increasePositionBufferBps)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_BasePositionManager *BasePositionManagerSession) SetIncreasePositionBufferBps(_increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetIncreasePositionBufferBps(&_BasePositionManager.TransactOpts, _increasePositionBufferBps)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SetIncreasePositionBufferBps(_increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetIncreasePositionBufferBps(&_BasePositionManager.TransactOpts, _increasePositionBufferBps)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SetMaxGlobalSizes(opts *bind.TransactOpts, _tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "setMaxGlobalSizes", _tokens, _longSizes, _shortSizes)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_BasePositionManager *BasePositionManagerSession) SetMaxGlobalSizes(_tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetMaxGlobalSizes(&_BasePositionManager.TransactOpts, _tokens, _longSizes, _shortSizes)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SetMaxGlobalSizes(_tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetMaxGlobalSizes(&_BasePositionManager.TransactOpts, _tokens, _longSizes, _shortSizes)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_BasePositionManager *BasePositionManagerTransactor) SetReferralStorage(opts *bind.TransactOpts, _referralStorage common.Address) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "setReferralStorage", _referralStorage)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_BasePositionManager *BasePositionManagerSession) SetReferralStorage(_referralStorage common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetReferralStorage(&_BasePositionManager.TransactOpts, _referralStorage)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) SetReferralStorage(_referralStorage common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.SetReferralStorage(&_BasePositionManager.TransactOpts, _referralStorage)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_BasePositionManager *BasePositionManagerTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _BasePositionManager.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_BasePositionManager *BasePositionManagerSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.WithdrawFees(&_BasePositionManager.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _BasePositionManager.Contract.WithdrawFees(&_BasePositionManager.TransactOpts, _token, _receiver)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasePositionManager *BasePositionManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasePositionManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasePositionManager *BasePositionManagerSession) Receive() (*types.Transaction, error) {
	return _BasePositionManager.Contract.Receive(&_BasePositionManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasePositionManager *BasePositionManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _BasePositionManager.Contract.Receive(&_BasePositionManager.TransactOpts)
}

// BasePositionManagerDecreasePositionReferralIterator is returned from FilterDecreasePositionReferral and is used to iterate over the raw logs and unpacked data for DecreasePositionReferral events raised by the BasePositionManager contract.
type BasePositionManagerDecreasePositionReferralIterator struct {
	Event *BasePositionManagerDecreasePositionReferral // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerDecreasePositionReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerDecreasePositionReferral)
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
		it.Event = new(BasePositionManagerDecreasePositionReferral)
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
func (it *BasePositionManagerDecreasePositionReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerDecreasePositionReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerDecreasePositionReferral represents a DecreasePositionReferral event raised by the BasePositionManager contract.
type BasePositionManagerDecreasePositionReferral struct {
	Account              common.Address
	SizeDelta            *big.Int
	MarginFeeBasisPoints *big.Int
	ReferralCode         [32]byte
	Referrer             common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDecreasePositionReferral is a free log retrieval operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_BasePositionManager *BasePositionManagerFilterer) FilterDecreasePositionReferral(opts *bind.FilterOpts) (*BasePositionManagerDecreasePositionReferralIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "DecreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerDecreasePositionReferralIterator{contract: _BasePositionManager.contract, event: "DecreasePositionReferral", logs: logs, sub: sub}, nil
}

// WatchDecreasePositionReferral is a free log subscription operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_BasePositionManager *BasePositionManagerFilterer) WatchDecreasePositionReferral(opts *bind.WatchOpts, sink chan<- *BasePositionManagerDecreasePositionReferral) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "DecreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerDecreasePositionReferral)
				if err := _BasePositionManager.contract.UnpackLog(event, "DecreasePositionReferral", log); err != nil {
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

// ParseDecreasePositionReferral is a log parse operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_BasePositionManager *BasePositionManagerFilterer) ParseDecreasePositionReferral(log types.Log) (*BasePositionManagerDecreasePositionReferral, error) {
	event := new(BasePositionManagerDecreasePositionReferral)
	if err := _BasePositionManager.contract.UnpackLog(event, "DecreasePositionReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerIncreasePositionReferralIterator is returned from FilterIncreasePositionReferral and is used to iterate over the raw logs and unpacked data for IncreasePositionReferral events raised by the BasePositionManager contract.
type BasePositionManagerIncreasePositionReferralIterator struct {
	Event *BasePositionManagerIncreasePositionReferral // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerIncreasePositionReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerIncreasePositionReferral)
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
		it.Event = new(BasePositionManagerIncreasePositionReferral)
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
func (it *BasePositionManagerIncreasePositionReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerIncreasePositionReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerIncreasePositionReferral represents a IncreasePositionReferral event raised by the BasePositionManager contract.
type BasePositionManagerIncreasePositionReferral struct {
	Account              common.Address
	SizeDelta            *big.Int
	MarginFeeBasisPoints *big.Int
	ReferralCode         [32]byte
	Referrer             common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIncreasePositionReferral is a free log retrieval operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_BasePositionManager *BasePositionManagerFilterer) FilterIncreasePositionReferral(opts *bind.FilterOpts) (*BasePositionManagerIncreasePositionReferralIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "IncreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerIncreasePositionReferralIterator{contract: _BasePositionManager.contract, event: "IncreasePositionReferral", logs: logs, sub: sub}, nil
}

// WatchIncreasePositionReferral is a free log subscription operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_BasePositionManager *BasePositionManagerFilterer) WatchIncreasePositionReferral(opts *bind.WatchOpts, sink chan<- *BasePositionManagerIncreasePositionReferral) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "IncreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerIncreasePositionReferral)
				if err := _BasePositionManager.contract.UnpackLog(event, "IncreasePositionReferral", log); err != nil {
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

// ParseIncreasePositionReferral is a log parse operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_BasePositionManager *BasePositionManagerFilterer) ParseIncreasePositionReferral(log types.Log) (*BasePositionManagerIncreasePositionReferral, error) {
	event := new(BasePositionManagerIncreasePositionReferral)
	if err := _BasePositionManager.contract.UnpackLog(event, "IncreasePositionReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the BasePositionManager contract.
type BasePositionManagerSetAdminIterator struct {
	Event *BasePositionManagerSetAdmin // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerSetAdmin)
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
		it.Event = new(BasePositionManagerSetAdmin)
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
func (it *BasePositionManagerSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerSetAdmin represents a SetAdmin event raised by the BasePositionManager contract.
type BasePositionManagerSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_BasePositionManager *BasePositionManagerFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*BasePositionManagerSetAdminIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerSetAdminIterator{contract: _BasePositionManager.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_BasePositionManager *BasePositionManagerFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *BasePositionManagerSetAdmin) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerSetAdmin)
				if err := _BasePositionManager.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_BasePositionManager *BasePositionManagerFilterer) ParseSetAdmin(log types.Log) (*BasePositionManagerSetAdmin, error) {
	event := new(BasePositionManagerSetAdmin)
	if err := _BasePositionManager.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerSetDepositFeeIterator is returned from FilterSetDepositFee and is used to iterate over the raw logs and unpacked data for SetDepositFee events raised by the BasePositionManager contract.
type BasePositionManagerSetDepositFeeIterator struct {
	Event *BasePositionManagerSetDepositFee // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerSetDepositFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerSetDepositFee)
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
		it.Event = new(BasePositionManagerSetDepositFee)
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
func (it *BasePositionManagerSetDepositFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerSetDepositFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerSetDepositFee represents a SetDepositFee event raised by the BasePositionManager contract.
type BasePositionManagerSetDepositFee struct {
	DepositFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetDepositFee is a free log retrieval operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_BasePositionManager *BasePositionManagerFilterer) FilterSetDepositFee(opts *bind.FilterOpts) (*BasePositionManagerSetDepositFeeIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerSetDepositFeeIterator{contract: _BasePositionManager.contract, event: "SetDepositFee", logs: logs, sub: sub}, nil
}

// WatchSetDepositFee is a free log subscription operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_BasePositionManager *BasePositionManagerFilterer) WatchSetDepositFee(opts *bind.WatchOpts, sink chan<- *BasePositionManagerSetDepositFee) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerSetDepositFee)
				if err := _BasePositionManager.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
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

// ParseSetDepositFee is a log parse operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_BasePositionManager *BasePositionManagerFilterer) ParseSetDepositFee(log types.Log) (*BasePositionManagerSetDepositFee, error) {
	event := new(BasePositionManagerSetDepositFee)
	if err := _BasePositionManager.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerSetEthTransferGasLimitIterator is returned from FilterSetEthTransferGasLimit and is used to iterate over the raw logs and unpacked data for SetEthTransferGasLimit events raised by the BasePositionManager contract.
type BasePositionManagerSetEthTransferGasLimitIterator struct {
	Event *BasePositionManagerSetEthTransferGasLimit // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerSetEthTransferGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerSetEthTransferGasLimit)
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
		it.Event = new(BasePositionManagerSetEthTransferGasLimit)
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
func (it *BasePositionManagerSetEthTransferGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerSetEthTransferGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerSetEthTransferGasLimit represents a SetEthTransferGasLimit event raised by the BasePositionManager contract.
type BasePositionManagerSetEthTransferGasLimit struct {
	EthTransferGasLimit *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetEthTransferGasLimit is a free log retrieval operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_BasePositionManager *BasePositionManagerFilterer) FilterSetEthTransferGasLimit(opts *bind.FilterOpts) (*BasePositionManagerSetEthTransferGasLimitIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "SetEthTransferGasLimit")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerSetEthTransferGasLimitIterator{contract: _BasePositionManager.contract, event: "SetEthTransferGasLimit", logs: logs, sub: sub}, nil
}

// WatchSetEthTransferGasLimit is a free log subscription operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_BasePositionManager *BasePositionManagerFilterer) WatchSetEthTransferGasLimit(opts *bind.WatchOpts, sink chan<- *BasePositionManagerSetEthTransferGasLimit) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "SetEthTransferGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerSetEthTransferGasLimit)
				if err := _BasePositionManager.contract.UnpackLog(event, "SetEthTransferGasLimit", log); err != nil {
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

// ParseSetEthTransferGasLimit is a log parse operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_BasePositionManager *BasePositionManagerFilterer) ParseSetEthTransferGasLimit(log types.Log) (*BasePositionManagerSetEthTransferGasLimit, error) {
	event := new(BasePositionManagerSetEthTransferGasLimit)
	if err := _BasePositionManager.contract.UnpackLog(event, "SetEthTransferGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerSetIncreasePositionBufferBpsIterator is returned from FilterSetIncreasePositionBufferBps and is used to iterate over the raw logs and unpacked data for SetIncreasePositionBufferBps events raised by the BasePositionManager contract.
type BasePositionManagerSetIncreasePositionBufferBpsIterator struct {
	Event *BasePositionManagerSetIncreasePositionBufferBps // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerSetIncreasePositionBufferBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerSetIncreasePositionBufferBps)
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
		it.Event = new(BasePositionManagerSetIncreasePositionBufferBps)
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
func (it *BasePositionManagerSetIncreasePositionBufferBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerSetIncreasePositionBufferBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerSetIncreasePositionBufferBps represents a SetIncreasePositionBufferBps event raised by the BasePositionManager contract.
type BasePositionManagerSetIncreasePositionBufferBps struct {
	IncreasePositionBufferBps *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetIncreasePositionBufferBps is a free log retrieval operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_BasePositionManager *BasePositionManagerFilterer) FilterSetIncreasePositionBufferBps(opts *bind.FilterOpts) (*BasePositionManagerSetIncreasePositionBufferBpsIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "SetIncreasePositionBufferBps")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerSetIncreasePositionBufferBpsIterator{contract: _BasePositionManager.contract, event: "SetIncreasePositionBufferBps", logs: logs, sub: sub}, nil
}

// WatchSetIncreasePositionBufferBps is a free log subscription operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_BasePositionManager *BasePositionManagerFilterer) WatchSetIncreasePositionBufferBps(opts *bind.WatchOpts, sink chan<- *BasePositionManagerSetIncreasePositionBufferBps) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "SetIncreasePositionBufferBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerSetIncreasePositionBufferBps)
				if err := _BasePositionManager.contract.UnpackLog(event, "SetIncreasePositionBufferBps", log); err != nil {
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

// ParseSetIncreasePositionBufferBps is a log parse operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_BasePositionManager *BasePositionManagerFilterer) ParseSetIncreasePositionBufferBps(log types.Log) (*BasePositionManagerSetIncreasePositionBufferBps, error) {
	event := new(BasePositionManagerSetIncreasePositionBufferBps)
	if err := _BasePositionManager.contract.UnpackLog(event, "SetIncreasePositionBufferBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerSetMaxGlobalSizesIterator is returned from FilterSetMaxGlobalSizes and is used to iterate over the raw logs and unpacked data for SetMaxGlobalSizes events raised by the BasePositionManager contract.
type BasePositionManagerSetMaxGlobalSizesIterator struct {
	Event *BasePositionManagerSetMaxGlobalSizes // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerSetMaxGlobalSizesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerSetMaxGlobalSizes)
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
		it.Event = new(BasePositionManagerSetMaxGlobalSizes)
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
func (it *BasePositionManagerSetMaxGlobalSizesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerSetMaxGlobalSizesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerSetMaxGlobalSizes represents a SetMaxGlobalSizes event raised by the BasePositionManager contract.
type BasePositionManagerSetMaxGlobalSizes struct {
	Tokens     []common.Address
	LongSizes  []*big.Int
	ShortSizes []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMaxGlobalSizes is a free log retrieval operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_BasePositionManager *BasePositionManagerFilterer) FilterSetMaxGlobalSizes(opts *bind.FilterOpts) (*BasePositionManagerSetMaxGlobalSizesIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "SetMaxGlobalSizes")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerSetMaxGlobalSizesIterator{contract: _BasePositionManager.contract, event: "SetMaxGlobalSizes", logs: logs, sub: sub}, nil
}

// WatchSetMaxGlobalSizes is a free log subscription operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_BasePositionManager *BasePositionManagerFilterer) WatchSetMaxGlobalSizes(opts *bind.WatchOpts, sink chan<- *BasePositionManagerSetMaxGlobalSizes) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "SetMaxGlobalSizes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerSetMaxGlobalSizes)
				if err := _BasePositionManager.contract.UnpackLog(event, "SetMaxGlobalSizes", log); err != nil {
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

// ParseSetMaxGlobalSizes is a log parse operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_BasePositionManager *BasePositionManagerFilterer) ParseSetMaxGlobalSizes(log types.Log) (*BasePositionManagerSetMaxGlobalSizes, error) {
	event := new(BasePositionManagerSetMaxGlobalSizes)
	if err := _BasePositionManager.contract.UnpackLog(event, "SetMaxGlobalSizes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerSetReferralStorageIterator is returned from FilterSetReferralStorage and is used to iterate over the raw logs and unpacked data for SetReferralStorage events raised by the BasePositionManager contract.
type BasePositionManagerSetReferralStorageIterator struct {
	Event *BasePositionManagerSetReferralStorage // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerSetReferralStorageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerSetReferralStorage)
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
		it.Event = new(BasePositionManagerSetReferralStorage)
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
func (it *BasePositionManagerSetReferralStorageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerSetReferralStorageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerSetReferralStorage represents a SetReferralStorage event raised by the BasePositionManager contract.
type BasePositionManagerSetReferralStorage struct {
	ReferralStorage common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReferralStorage is a free log retrieval operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_BasePositionManager *BasePositionManagerFilterer) FilterSetReferralStorage(opts *bind.FilterOpts) (*BasePositionManagerSetReferralStorageIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "SetReferralStorage")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerSetReferralStorageIterator{contract: _BasePositionManager.contract, event: "SetReferralStorage", logs: logs, sub: sub}, nil
}

// WatchSetReferralStorage is a free log subscription operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_BasePositionManager *BasePositionManagerFilterer) WatchSetReferralStorage(opts *bind.WatchOpts, sink chan<- *BasePositionManagerSetReferralStorage) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "SetReferralStorage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerSetReferralStorage)
				if err := _BasePositionManager.contract.UnpackLog(event, "SetReferralStorage", log); err != nil {
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

// ParseSetReferralStorage is a log parse operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_BasePositionManager *BasePositionManagerFilterer) ParseSetReferralStorage(log types.Log) (*BasePositionManagerSetReferralStorage, error) {
	event := new(BasePositionManagerSetReferralStorage)
	if err := _BasePositionManager.contract.UnpackLog(event, "SetReferralStorage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasePositionManagerWithdrawFeesIterator is returned from FilterWithdrawFees and is used to iterate over the raw logs and unpacked data for WithdrawFees events raised by the BasePositionManager contract.
type BasePositionManagerWithdrawFeesIterator struct {
	Event *BasePositionManagerWithdrawFees // Event containing the contract specifics and raw log

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
func (it *BasePositionManagerWithdrawFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasePositionManagerWithdrawFees)
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
		it.Event = new(BasePositionManagerWithdrawFees)
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
func (it *BasePositionManagerWithdrawFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasePositionManagerWithdrawFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasePositionManagerWithdrawFees represents a WithdrawFees event raised by the BasePositionManager contract.
type BasePositionManagerWithdrawFees struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFees is a free log retrieval operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_BasePositionManager *BasePositionManagerFilterer) FilterWithdrawFees(opts *bind.FilterOpts) (*BasePositionManagerWithdrawFeesIterator, error) {

	logs, sub, err := _BasePositionManager.contract.FilterLogs(opts, "WithdrawFees")
	if err != nil {
		return nil, err
	}
	return &BasePositionManagerWithdrawFeesIterator{contract: _BasePositionManager.contract, event: "WithdrawFees", logs: logs, sub: sub}, nil
}

// WatchWithdrawFees is a free log subscription operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_BasePositionManager *BasePositionManagerFilterer) WatchWithdrawFees(opts *bind.WatchOpts, sink chan<- *BasePositionManagerWithdrawFees) (event.Subscription, error) {

	logs, sub, err := _BasePositionManager.contract.WatchLogs(opts, "WithdrawFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasePositionManagerWithdrawFees)
				if err := _BasePositionManager.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
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

// ParseWithdrawFees is a log parse operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_BasePositionManager *BasePositionManagerFilterer) ParseWithdrawFees(log types.Log) (*BasePositionManagerWithdrawFees, error) {
	event := new(BasePositionManagerWithdrawFees)
	if err := _BasePositionManager.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
