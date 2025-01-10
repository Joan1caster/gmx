// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivester

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

// IVesterMetaData contains all meta data concerning the IVester contract.
var IVesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"bonusRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"cumulativeClaimAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"cumulativeRewardDeductions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getCombinedAverageStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getMaxVestableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getVestedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"pairAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBonusRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setCumulativeRewardDeductions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTransferredAverageStakedAmounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTransferredCumulativeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"transferStakeValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"transferredAverageStakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"transferredCumulativeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IVesterABI is the input ABI used to generate the binding from.
// Deprecated: Use IVesterMetaData.ABI instead.
var IVesterABI = IVesterMetaData.ABI

// IVester is an auto generated Go binding around an Ethereum contract.
type IVester struct {
	IVesterCaller     // Read-only binding to the contract
	IVesterTransactor // Write-only binding to the contract
	IVesterFilterer   // Log filterer for contract events
}

// IVesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVesterSession struct {
	Contract     *IVester          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVesterCallerSession struct {
	Contract *IVesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IVesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVesterTransactorSession struct {
	Contract     *IVesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IVesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVesterRaw struct {
	Contract *IVester // Generic contract binding to access the raw methods on
}

// IVesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVesterCallerRaw struct {
	Contract *IVesterCaller // Generic read-only contract binding to access the raw methods on
}

// IVesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVesterTransactorRaw struct {
	Contract *IVesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVester creates a new instance of IVester, bound to a specific deployed contract.
func NewIVester(address common.Address, backend bind.ContractBackend) (*IVester, error) {
	contract, err := bindIVester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVester{IVesterCaller: IVesterCaller{contract: contract}, IVesterTransactor: IVesterTransactor{contract: contract}, IVesterFilterer: IVesterFilterer{contract: contract}}, nil
}

// NewIVesterCaller creates a new read-only instance of IVester, bound to a specific deployed contract.
func NewIVesterCaller(address common.Address, caller bind.ContractCaller) (*IVesterCaller, error) {
	contract, err := bindIVester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVesterCaller{contract: contract}, nil
}

// NewIVesterTransactor creates a new write-only instance of IVester, bound to a specific deployed contract.
func NewIVesterTransactor(address common.Address, transactor bind.ContractTransactor) (*IVesterTransactor, error) {
	contract, err := bindIVester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVesterTransactor{contract: contract}, nil
}

// NewIVesterFilterer creates a new log filterer instance of IVester, bound to a specific deployed contract.
func NewIVesterFilterer(address common.Address, filterer bind.ContractFilterer) (*IVesterFilterer, error) {
	contract, err := bindIVester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVesterFilterer{contract: contract}, nil
}

// bindIVester binds a generic wrapper to an already deployed contract.
func bindIVester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVester *IVesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVester.Contract.IVesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVester *IVesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVester.Contract.IVesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVester *IVesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVester.Contract.IVesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVester *IVesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVester *IVesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVester *IVesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVester.Contract.contract.Transact(opts, method, params...)
}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address _account) view returns(uint256)
func (_IVester *IVesterCaller) BonusRewards(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "bonusRewards", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address _account) view returns(uint256)
func (_IVester *IVesterSession) BonusRewards(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.BonusRewards(&_IVester.CallOpts, _account)
}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) BonusRewards(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.BonusRewards(&_IVester.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IVester *IVesterCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IVester *IVesterSession) Claimable(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.Claimable(&_IVester.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.Claimable(&_IVester.CallOpts, _account)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCaller) ClaimedAmounts(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "claimedAmounts", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address _account) view returns(uint256)
func (_IVester *IVesterSession) ClaimedAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.ClaimedAmounts(&_IVester.CallOpts, _account)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) ClaimedAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.ClaimedAmounts(&_IVester.CallOpts, _account)
}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCaller) CumulativeClaimAmounts(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "cumulativeClaimAmounts", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address _account) view returns(uint256)
func (_IVester *IVesterSession) CumulativeClaimAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.CumulativeClaimAmounts(&_IVester.CallOpts, _account)
}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) CumulativeClaimAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.CumulativeClaimAmounts(&_IVester.CallOpts, _account)
}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address _account) view returns(uint256)
func (_IVester *IVesterCaller) CumulativeRewardDeductions(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "cumulativeRewardDeductions", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address _account) view returns(uint256)
func (_IVester *IVesterSession) CumulativeRewardDeductions(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.CumulativeRewardDeductions(&_IVester.CallOpts, _account)
}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) CumulativeRewardDeductions(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.CumulativeRewardDeductions(&_IVester.CallOpts, _account)
}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_IVester *IVesterCaller) GetCombinedAverageStakedAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "getCombinedAverageStakedAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_IVester *IVesterSession) GetCombinedAverageStakedAmount(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.GetCombinedAverageStakedAmount(&_IVester.CallOpts, _account)
}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) GetCombinedAverageStakedAmount(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.GetCombinedAverageStakedAmount(&_IVester.CallOpts, _account)
}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_IVester *IVesterCaller) GetMaxVestableAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "getMaxVestableAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_IVester *IVesterSession) GetMaxVestableAmount(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.GetMaxVestableAmount(&_IVester.CallOpts, _account)
}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) GetMaxVestableAmount(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.GetMaxVestableAmount(&_IVester.CallOpts, _account)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_IVester *IVesterCaller) GetVestedAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "getVestedAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_IVester *IVesterSession) GetVestedAmount(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.GetVestedAmount(&_IVester.CallOpts, _account)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) GetVestedAmount(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.GetVestedAmount(&_IVester.CallOpts, _account)
}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCaller) PairAmounts(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "pairAmounts", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address _account) view returns(uint256)
func (_IVester *IVesterSession) PairAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.PairAmounts(&_IVester.CallOpts, _account)
}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) PairAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.PairAmounts(&_IVester.CallOpts, _account)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_IVester *IVesterCaller) RewardTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "rewardTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_IVester *IVesterSession) RewardTracker() (common.Address, error) {
	return _IVester.Contract.RewardTracker(&_IVester.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_IVester *IVesterCallerSession) RewardTracker() (common.Address, error) {
	return _IVester.Contract.RewardTracker(&_IVester.CallOpts)
}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCaller) TransferredAverageStakedAmounts(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "transferredAverageStakedAmounts", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address _account) view returns(uint256)
func (_IVester *IVesterSession) TransferredAverageStakedAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.TransferredAverageStakedAmounts(&_IVester.CallOpts, _account)
}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) TransferredAverageStakedAmounts(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.TransferredAverageStakedAmounts(&_IVester.CallOpts, _account)
}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address _account) view returns(uint256)
func (_IVester *IVesterCaller) TransferredCumulativeRewards(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVester.contract.Call(opts, &out, "transferredCumulativeRewards", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address _account) view returns(uint256)
func (_IVester *IVesterSession) TransferredCumulativeRewards(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.TransferredCumulativeRewards(&_IVester.CallOpts, _account)
}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address _account) view returns(uint256)
func (_IVester *IVesterCallerSession) TransferredCumulativeRewards(_account common.Address) (*big.Int, error) {
	return _IVester.Contract.TransferredCumulativeRewards(&_IVester.CallOpts, _account)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_IVester *IVesterTransactor) ClaimForAccount(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVester.contract.Transact(opts, "claimForAccount", _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_IVester *IVesterSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVester.Contract.ClaimForAccount(&_IVester.TransactOpts, _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_IVester *IVesterTransactorSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVester.Contract.ClaimForAccount(&_IVester.TransactOpts, _account, _receiver)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactor) SetBonusRewards(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.contract.Transact(opts, "setBonusRewards", _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_IVester *IVesterSession) SetBonusRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetBonusRewards(&_IVester.TransactOpts, _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactorSession) SetBonusRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetBonusRewards(&_IVester.TransactOpts, _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactor) SetCumulativeRewardDeductions(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.contract.Transact(opts, "setCumulativeRewardDeductions", _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_IVester *IVesterSession) SetCumulativeRewardDeductions(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetCumulativeRewardDeductions(&_IVester.TransactOpts, _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactorSession) SetCumulativeRewardDeductions(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetCumulativeRewardDeductions(&_IVester.TransactOpts, _account, _amount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactor) SetTransferredAverageStakedAmounts(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.contract.Transact(opts, "setTransferredAverageStakedAmounts", _account, _amount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_IVester *IVesterSession) SetTransferredAverageStakedAmounts(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetTransferredAverageStakedAmounts(&_IVester.TransactOpts, _account, _amount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactorSession) SetTransferredAverageStakedAmounts(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetTransferredAverageStakedAmounts(&_IVester.TransactOpts, _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactor) SetTransferredCumulativeRewards(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.contract.Transact(opts, "setTransferredCumulativeRewards", _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_IVester *IVesterSession) SetTransferredCumulativeRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetTransferredCumulativeRewards(&_IVester.TransactOpts, _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_IVester *IVesterTransactorSession) SetTransferredCumulativeRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IVester.Contract.SetTransferredCumulativeRewards(&_IVester.TransactOpts, _account, _amount)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_IVester *IVesterTransactor) TransferStakeValues(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVester.contract.Transact(opts, "transferStakeValues", _sender, _receiver)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_IVester *IVesterSession) TransferStakeValues(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVester.Contract.TransferStakeValues(&_IVester.TransactOpts, _sender, _receiver)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_IVester *IVesterTransactorSession) TransferStakeValues(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _IVester.Contract.TransferStakeValues(&_IVester.TransactOpts, _sender, _receiver)
}
