// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package timedistributor

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

// TimeDistributorMetaData contains all meta data concerning the TimeDistributor contract.
var TimeDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Distribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"name\":\"DistributionChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPerIntervalChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getDistributionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getIntervals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getRewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastDistributionTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_rewardTokens\",\"type\":\"address[]\"}],\"name\":\"setDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTokensPerInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"updateLastDistributionTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TimeDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use TimeDistributorMetaData.ABI instead.
var TimeDistributorABI = TimeDistributorMetaData.ABI

// TimeDistributor is an auto generated Go binding around an Ethereum contract.
type TimeDistributor struct {
	TimeDistributorCaller     // Read-only binding to the contract
	TimeDistributorTransactor // Write-only binding to the contract
	TimeDistributorFilterer   // Log filterer for contract events
}

// TimeDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeDistributorSession struct {
	Contract     *TimeDistributor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeDistributorCallerSession struct {
	Contract *TimeDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TimeDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeDistributorTransactorSession struct {
	Contract     *TimeDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TimeDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeDistributorRaw struct {
	Contract *TimeDistributor // Generic contract binding to access the raw methods on
}

// TimeDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeDistributorCallerRaw struct {
	Contract *TimeDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// TimeDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeDistributorTransactorRaw struct {
	Contract *TimeDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimeDistributor creates a new instance of TimeDistributor, bound to a specific deployed contract.
func NewTimeDistributor(address common.Address, backend bind.ContractBackend) (*TimeDistributor, error) {
	contract, err := bindTimeDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimeDistributor{TimeDistributorCaller: TimeDistributorCaller{contract: contract}, TimeDistributorTransactor: TimeDistributorTransactor{contract: contract}, TimeDistributorFilterer: TimeDistributorFilterer{contract: contract}}, nil
}

// NewTimeDistributorCaller creates a new read-only instance of TimeDistributor, bound to a specific deployed contract.
func NewTimeDistributorCaller(address common.Address, caller bind.ContractCaller) (*TimeDistributorCaller, error) {
	contract, err := bindTimeDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeDistributorCaller{contract: contract}, nil
}

// NewTimeDistributorTransactor creates a new write-only instance of TimeDistributor, bound to a specific deployed contract.
func NewTimeDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeDistributorTransactor, error) {
	contract, err := bindTimeDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeDistributorTransactor{contract: contract}, nil
}

// NewTimeDistributorFilterer creates a new log filterer instance of TimeDistributor, bound to a specific deployed contract.
func NewTimeDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeDistributorFilterer, error) {
	contract, err := bindTimeDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeDistributorFilterer{contract: contract}, nil
}

// bindTimeDistributor binds a generic wrapper to an already deployed contract.
func bindTimeDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimeDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeDistributor *TimeDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeDistributor.Contract.TimeDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeDistributor *TimeDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeDistributor.Contract.TimeDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeDistributor *TimeDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeDistributor.Contract.TimeDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeDistributor *TimeDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeDistributor *TimeDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeDistributor *TimeDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeDistributor.Contract.contract.Transact(opts, method, params...)
}

// DISTRIBUTIONINTERVAL is a free data retrieval call binding the contract method 0x08b26b75.
//
// Solidity: function DISTRIBUTION_INTERVAL() view returns(uint256)
func (_TimeDistributor *TimeDistributorCaller) DISTRIBUTIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "DISTRIBUTION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONINTERVAL is a free data retrieval call binding the contract method 0x08b26b75.
//
// Solidity: function DISTRIBUTION_INTERVAL() view returns(uint256)
func (_TimeDistributor *TimeDistributorSession) DISTRIBUTIONINTERVAL() (*big.Int, error) {
	return _TimeDistributor.Contract.DISTRIBUTIONINTERVAL(&_TimeDistributor.CallOpts)
}

// DISTRIBUTIONINTERVAL is a free data retrieval call binding the contract method 0x08b26b75.
//
// Solidity: function DISTRIBUTION_INTERVAL() view returns(uint256)
func (_TimeDistributor *TimeDistributorCallerSession) DISTRIBUTIONINTERVAL() (*big.Int, error) {
	return _TimeDistributor.Contract.DISTRIBUTIONINTERVAL(&_TimeDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TimeDistributor *TimeDistributorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TimeDistributor *TimeDistributorSession) Admin() (common.Address, error) {
	return _TimeDistributor.Contract.Admin(&_TimeDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TimeDistributor *TimeDistributorCallerSession) Admin() (common.Address, error) {
	return _TimeDistributor.Contract.Admin(&_TimeDistributor.CallOpts)
}

// GetDistributionAmount is a free data retrieval call binding the contract method 0xb7fba9f6.
//
// Solidity: function getDistributionAmount(address _receiver) view returns(uint256)
func (_TimeDistributor *TimeDistributorCaller) GetDistributionAmount(opts *bind.CallOpts, _receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "getDistributionAmount", _receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionAmount is a free data retrieval call binding the contract method 0xb7fba9f6.
//
// Solidity: function getDistributionAmount(address _receiver) view returns(uint256)
func (_TimeDistributor *TimeDistributorSession) GetDistributionAmount(_receiver common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.GetDistributionAmount(&_TimeDistributor.CallOpts, _receiver)
}

// GetDistributionAmount is a free data retrieval call binding the contract method 0xb7fba9f6.
//
// Solidity: function getDistributionAmount(address _receiver) view returns(uint256)
func (_TimeDistributor *TimeDistributorCallerSession) GetDistributionAmount(_receiver common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.GetDistributionAmount(&_TimeDistributor.CallOpts, _receiver)
}

// GetIntervals is a free data retrieval call binding the contract method 0xc27a1acd.
//
// Solidity: function getIntervals(address _receiver) view returns(uint256)
func (_TimeDistributor *TimeDistributorCaller) GetIntervals(opts *bind.CallOpts, _receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "getIntervals", _receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIntervals is a free data retrieval call binding the contract method 0xc27a1acd.
//
// Solidity: function getIntervals(address _receiver) view returns(uint256)
func (_TimeDistributor *TimeDistributorSession) GetIntervals(_receiver common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.GetIntervals(&_TimeDistributor.CallOpts, _receiver)
}

// GetIntervals is a free data retrieval call binding the contract method 0xc27a1acd.
//
// Solidity: function getIntervals(address _receiver) view returns(uint256)
func (_TimeDistributor *TimeDistributorCallerSession) GetIntervals(_receiver common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.GetIntervals(&_TimeDistributor.CallOpts, _receiver)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address _receiver) view returns(address)
func (_TimeDistributor *TimeDistributorCaller) GetRewardToken(opts *bind.CallOpts, _receiver common.Address) (common.Address, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "getRewardToken", _receiver)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address _receiver) view returns(address)
func (_TimeDistributor *TimeDistributorSession) GetRewardToken(_receiver common.Address) (common.Address, error) {
	return _TimeDistributor.Contract.GetRewardToken(&_TimeDistributor.CallOpts, _receiver)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x77078872.
//
// Solidity: function getRewardToken(address _receiver) view returns(address)
func (_TimeDistributor *TimeDistributorCallerSession) GetRewardToken(_receiver common.Address) (common.Address, error) {
	return _TimeDistributor.Contract.GetRewardToken(&_TimeDistributor.CallOpts, _receiver)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_TimeDistributor *TimeDistributorCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_TimeDistributor *TimeDistributorSession) Gov() (common.Address, error) {
	return _TimeDistributor.Contract.Gov(&_TimeDistributor.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_TimeDistributor *TimeDistributorCallerSession) Gov() (common.Address, error) {
	return _TimeDistributor.Contract.Gov(&_TimeDistributor.CallOpts)
}

// LastDistributionTime is a free data retrieval call binding the contract method 0x847ec40c.
//
// Solidity: function lastDistributionTime(address ) view returns(uint256)
func (_TimeDistributor *TimeDistributorCaller) LastDistributionTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "lastDistributionTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDistributionTime is a free data retrieval call binding the contract method 0x847ec40c.
//
// Solidity: function lastDistributionTime(address ) view returns(uint256)
func (_TimeDistributor *TimeDistributorSession) LastDistributionTime(arg0 common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.LastDistributionTime(&_TimeDistributor.CallOpts, arg0)
}

// LastDistributionTime is a free data retrieval call binding the contract method 0x847ec40c.
//
// Solidity: function lastDistributionTime(address ) view returns(uint256)
func (_TimeDistributor *TimeDistributorCallerSession) LastDistributionTime(arg0 common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.LastDistributionTime(&_TimeDistributor.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0xf5ab16cc.
//
// Solidity: function rewardTokens(address ) view returns(address)
func (_TimeDistributor *TimeDistributorCaller) RewardTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0xf5ab16cc.
//
// Solidity: function rewardTokens(address ) view returns(address)
func (_TimeDistributor *TimeDistributorSession) RewardTokens(arg0 common.Address) (common.Address, error) {
	return _TimeDistributor.Contract.RewardTokens(&_TimeDistributor.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0xf5ab16cc.
//
// Solidity: function rewardTokens(address ) view returns(address)
func (_TimeDistributor *TimeDistributorCallerSession) RewardTokens(arg0 common.Address) (common.Address, error) {
	return _TimeDistributor.Contract.RewardTokens(&_TimeDistributor.CallOpts, arg0)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0x270da240.
//
// Solidity: function tokensPerInterval(address ) view returns(uint256)
func (_TimeDistributor *TimeDistributorCaller) TokensPerInterval(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TimeDistributor.contract.Call(opts, &out, "tokensPerInterval", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0x270da240.
//
// Solidity: function tokensPerInterval(address ) view returns(uint256)
func (_TimeDistributor *TimeDistributorSession) TokensPerInterval(arg0 common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.TokensPerInterval(&_TimeDistributor.CallOpts, arg0)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0x270da240.
//
// Solidity: function tokensPerInterval(address ) view returns(uint256)
func (_TimeDistributor *TimeDistributorCallerSession) TokensPerInterval(arg0 common.Address) (*big.Int, error) {
	return _TimeDistributor.Contract.TokensPerInterval(&_TimeDistributor.CallOpts, arg0)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_TimeDistributor *TimeDistributorTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeDistributor.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_TimeDistributor *TimeDistributorSession) Distribute() (*types.Transaction, error) {
	return _TimeDistributor.Contract.Distribute(&_TimeDistributor.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_TimeDistributor *TimeDistributorTransactorSession) Distribute() (*types.Transaction, error) {
	return _TimeDistributor.Contract.Distribute(&_TimeDistributor.TransactOpts)
}

// SetDistribution is a paid mutator transaction binding the contract method 0x41d0204b.
//
// Solidity: function setDistribution(address[] _receivers, uint256[] _amounts, address[] _rewardTokens) returns()
func (_TimeDistributor *TimeDistributorTransactor) SetDistribution(opts *bind.TransactOpts, _receivers []common.Address, _amounts []*big.Int, _rewardTokens []common.Address) (*types.Transaction, error) {
	return _TimeDistributor.contract.Transact(opts, "setDistribution", _receivers, _amounts, _rewardTokens)
}

// SetDistribution is a paid mutator transaction binding the contract method 0x41d0204b.
//
// Solidity: function setDistribution(address[] _receivers, uint256[] _amounts, address[] _rewardTokens) returns()
func (_TimeDistributor *TimeDistributorSession) SetDistribution(_receivers []common.Address, _amounts []*big.Int, _rewardTokens []common.Address) (*types.Transaction, error) {
	return _TimeDistributor.Contract.SetDistribution(&_TimeDistributor.TransactOpts, _receivers, _amounts, _rewardTokens)
}

// SetDistribution is a paid mutator transaction binding the contract method 0x41d0204b.
//
// Solidity: function setDistribution(address[] _receivers, uint256[] _amounts, address[] _rewardTokens) returns()
func (_TimeDistributor *TimeDistributorTransactorSession) SetDistribution(_receivers []common.Address, _amounts []*big.Int, _rewardTokens []common.Address) (*types.Transaction, error) {
	return _TimeDistributor.Contract.SetDistribution(&_TimeDistributor.TransactOpts, _receivers, _amounts, _rewardTokens)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_TimeDistributor *TimeDistributorTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _TimeDistributor.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_TimeDistributor *TimeDistributorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _TimeDistributor.Contract.SetGov(&_TimeDistributor.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_TimeDistributor *TimeDistributorTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _TimeDistributor.Contract.SetGov(&_TimeDistributor.TransactOpts, _gov)
}

// SetTokensPerInterval is a paid mutator transaction binding the contract method 0x83b3dd24.
//
// Solidity: function setTokensPerInterval(address _receiver, uint256 _amount) returns()
func (_TimeDistributor *TimeDistributorTransactor) SetTokensPerInterval(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TimeDistributor.contract.Transact(opts, "setTokensPerInterval", _receiver, _amount)
}

// SetTokensPerInterval is a paid mutator transaction binding the contract method 0x83b3dd24.
//
// Solidity: function setTokensPerInterval(address _receiver, uint256 _amount) returns()
func (_TimeDistributor *TimeDistributorSession) SetTokensPerInterval(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TimeDistributor.Contract.SetTokensPerInterval(&_TimeDistributor.TransactOpts, _receiver, _amount)
}

// SetTokensPerInterval is a paid mutator transaction binding the contract method 0x83b3dd24.
//
// Solidity: function setTokensPerInterval(address _receiver, uint256 _amount) returns()
func (_TimeDistributor *TimeDistributorTransactorSession) SetTokensPerInterval(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TimeDistributor.Contract.SetTokensPerInterval(&_TimeDistributor.TransactOpts, _receiver, _amount)
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x19f7ab45.
//
// Solidity: function updateLastDistributionTime(address _receiver) returns()
func (_TimeDistributor *TimeDistributorTransactor) UpdateLastDistributionTime(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _TimeDistributor.contract.Transact(opts, "updateLastDistributionTime", _receiver)
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x19f7ab45.
//
// Solidity: function updateLastDistributionTime(address _receiver) returns()
func (_TimeDistributor *TimeDistributorSession) UpdateLastDistributionTime(_receiver common.Address) (*types.Transaction, error) {
	return _TimeDistributor.Contract.UpdateLastDistributionTime(&_TimeDistributor.TransactOpts, _receiver)
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x19f7ab45.
//
// Solidity: function updateLastDistributionTime(address _receiver) returns()
func (_TimeDistributor *TimeDistributorTransactorSession) UpdateLastDistributionTime(_receiver common.Address) (*types.Transaction, error) {
	return _TimeDistributor.Contract.UpdateLastDistributionTime(&_TimeDistributor.TransactOpts, _receiver)
}

// TimeDistributorDistributeIterator is returned from FilterDistribute and is used to iterate over the raw logs and unpacked data for Distribute events raised by the TimeDistributor contract.
type TimeDistributorDistributeIterator struct {
	Event *TimeDistributorDistribute // Event containing the contract specifics and raw log

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
func (it *TimeDistributorDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeDistributorDistribute)
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
		it.Event = new(TimeDistributorDistribute)
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
func (it *TimeDistributorDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeDistributorDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeDistributorDistribute represents a Distribute event raised by the TimeDistributor contract.
type TimeDistributorDistribute struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDistribute is a free log retrieval operation binding the contract event 0xc1d32ad5cca423e7dda2123dbf8c482f8e77d00b631c06e903a47f2cec1334df.
//
// Solidity: event Distribute(address receiver, uint256 amount)
func (_TimeDistributor *TimeDistributorFilterer) FilterDistribute(opts *bind.FilterOpts) (*TimeDistributorDistributeIterator, error) {

	logs, sub, err := _TimeDistributor.contract.FilterLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return &TimeDistributorDistributeIterator{contract: _TimeDistributor.contract, event: "Distribute", logs: logs, sub: sub}, nil
}

// WatchDistribute is a free log subscription operation binding the contract event 0xc1d32ad5cca423e7dda2123dbf8c482f8e77d00b631c06e903a47f2cec1334df.
//
// Solidity: event Distribute(address receiver, uint256 amount)
func (_TimeDistributor *TimeDistributorFilterer) WatchDistribute(opts *bind.WatchOpts, sink chan<- *TimeDistributorDistribute) (event.Subscription, error) {

	logs, sub, err := _TimeDistributor.contract.WatchLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeDistributorDistribute)
				if err := _TimeDistributor.contract.UnpackLog(event, "Distribute", log); err != nil {
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

// ParseDistribute is a log parse operation binding the contract event 0xc1d32ad5cca423e7dda2123dbf8c482f8e77d00b631c06e903a47f2cec1334df.
//
// Solidity: event Distribute(address receiver, uint256 amount)
func (_TimeDistributor *TimeDistributorFilterer) ParseDistribute(log types.Log) (*TimeDistributorDistribute, error) {
	event := new(TimeDistributorDistribute)
	if err := _TimeDistributor.contract.UnpackLog(event, "Distribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeDistributorDistributionChangeIterator is returned from FilterDistributionChange and is used to iterate over the raw logs and unpacked data for DistributionChange events raised by the TimeDistributor contract.
type TimeDistributorDistributionChangeIterator struct {
	Event *TimeDistributorDistributionChange // Event containing the contract specifics and raw log

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
func (it *TimeDistributorDistributionChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeDistributorDistributionChange)
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
		it.Event = new(TimeDistributorDistributionChange)
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
func (it *TimeDistributorDistributionChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeDistributorDistributionChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeDistributorDistributionChange represents a DistributionChange event raised by the TimeDistributor contract.
type TimeDistributorDistributionChange struct {
	Receiver    common.Address
	Amount      *big.Int
	RewardToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDistributionChange is a free log retrieval operation binding the contract event 0xabc51c16142cb4f85965fa5b3cfdeed4475e4751d9201718a2bfc334c152f614.
//
// Solidity: event DistributionChange(address receiver, uint256 amount, address rewardToken)
func (_TimeDistributor *TimeDistributorFilterer) FilterDistributionChange(opts *bind.FilterOpts) (*TimeDistributorDistributionChangeIterator, error) {

	logs, sub, err := _TimeDistributor.contract.FilterLogs(opts, "DistributionChange")
	if err != nil {
		return nil, err
	}
	return &TimeDistributorDistributionChangeIterator{contract: _TimeDistributor.contract, event: "DistributionChange", logs: logs, sub: sub}, nil
}

// WatchDistributionChange is a free log subscription operation binding the contract event 0xabc51c16142cb4f85965fa5b3cfdeed4475e4751d9201718a2bfc334c152f614.
//
// Solidity: event DistributionChange(address receiver, uint256 amount, address rewardToken)
func (_TimeDistributor *TimeDistributorFilterer) WatchDistributionChange(opts *bind.WatchOpts, sink chan<- *TimeDistributorDistributionChange) (event.Subscription, error) {

	logs, sub, err := _TimeDistributor.contract.WatchLogs(opts, "DistributionChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeDistributorDistributionChange)
				if err := _TimeDistributor.contract.UnpackLog(event, "DistributionChange", log); err != nil {
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

// ParseDistributionChange is a log parse operation binding the contract event 0xabc51c16142cb4f85965fa5b3cfdeed4475e4751d9201718a2bfc334c152f614.
//
// Solidity: event DistributionChange(address receiver, uint256 amount, address rewardToken)
func (_TimeDistributor *TimeDistributorFilterer) ParseDistributionChange(log types.Log) (*TimeDistributorDistributionChange, error) {
	event := new(TimeDistributorDistributionChange)
	if err := _TimeDistributor.contract.UnpackLog(event, "DistributionChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeDistributorTokensPerIntervalChangeIterator is returned from FilterTokensPerIntervalChange and is used to iterate over the raw logs and unpacked data for TokensPerIntervalChange events raised by the TimeDistributor contract.
type TimeDistributorTokensPerIntervalChangeIterator struct {
	Event *TimeDistributorTokensPerIntervalChange // Event containing the contract specifics and raw log

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
func (it *TimeDistributorTokensPerIntervalChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeDistributorTokensPerIntervalChange)
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
		it.Event = new(TimeDistributorTokensPerIntervalChange)
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
func (it *TimeDistributorTokensPerIntervalChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeDistributorTokensPerIntervalChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeDistributorTokensPerIntervalChange represents a TokensPerIntervalChange event raised by the TimeDistributor contract.
type TimeDistributorTokensPerIntervalChange struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokensPerIntervalChange is a free log retrieval operation binding the contract event 0x83a9322775b48550bbef762b82384c8ed0b747ff4c37188bd3bb72e3c7cca823.
//
// Solidity: event TokensPerIntervalChange(address receiver, uint256 amount)
func (_TimeDistributor *TimeDistributorFilterer) FilterTokensPerIntervalChange(opts *bind.FilterOpts) (*TimeDistributorTokensPerIntervalChangeIterator, error) {

	logs, sub, err := _TimeDistributor.contract.FilterLogs(opts, "TokensPerIntervalChange")
	if err != nil {
		return nil, err
	}
	return &TimeDistributorTokensPerIntervalChangeIterator{contract: _TimeDistributor.contract, event: "TokensPerIntervalChange", logs: logs, sub: sub}, nil
}

// WatchTokensPerIntervalChange is a free log subscription operation binding the contract event 0x83a9322775b48550bbef762b82384c8ed0b747ff4c37188bd3bb72e3c7cca823.
//
// Solidity: event TokensPerIntervalChange(address receiver, uint256 amount)
func (_TimeDistributor *TimeDistributorFilterer) WatchTokensPerIntervalChange(opts *bind.WatchOpts, sink chan<- *TimeDistributorTokensPerIntervalChange) (event.Subscription, error) {

	logs, sub, err := _TimeDistributor.contract.WatchLogs(opts, "TokensPerIntervalChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeDistributorTokensPerIntervalChange)
				if err := _TimeDistributor.contract.UnpackLog(event, "TokensPerIntervalChange", log); err != nil {
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

// ParseTokensPerIntervalChange is a log parse operation binding the contract event 0x83a9322775b48550bbef762b82384c8ed0b747ff4c37188bd3bb72e3c7cca823.
//
// Solidity: event TokensPerIntervalChange(address receiver, uint256 amount)
func (_TimeDistributor *TimeDistributorFilterer) ParseTokensPerIntervalChange(log types.Log) (*TimeDistributorTokensPerIntervalChange, error) {
	event := new(TimeDistributorTokensPerIntervalChange)
	if err := _TimeDistributor.contract.UnpackLog(event, "TokensPerIntervalChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
