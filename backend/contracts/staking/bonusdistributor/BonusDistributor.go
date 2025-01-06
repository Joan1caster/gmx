// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bonusdistributor

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

// BonusDistributorMetaData contains all meta data concerning the BonusDistributor contract.
var BonusDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardTracker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BonusMultiplierChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Distribute\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BONUS_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusMultiplierBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDistributionTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusMultiplierBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setBonusMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateLastDistributionTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BonusDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use BonusDistributorMetaData.ABI instead.
var BonusDistributorABI = BonusDistributorMetaData.ABI

// BonusDistributor is an auto generated Go binding around an Ethereum contract.
type BonusDistributor struct {
	BonusDistributorCaller     // Read-only binding to the contract
	BonusDistributorTransactor // Write-only binding to the contract
	BonusDistributorFilterer   // Log filterer for contract events
}

// BonusDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BonusDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BonusDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BonusDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BonusDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BonusDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BonusDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BonusDistributorSession struct {
	Contract     *BonusDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BonusDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BonusDistributorCallerSession struct {
	Contract *BonusDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BonusDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BonusDistributorTransactorSession struct {
	Contract     *BonusDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BonusDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BonusDistributorRaw struct {
	Contract *BonusDistributor // Generic contract binding to access the raw methods on
}

// BonusDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BonusDistributorCallerRaw struct {
	Contract *BonusDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// BonusDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BonusDistributorTransactorRaw struct {
	Contract *BonusDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBonusDistributor creates a new instance of BonusDistributor, bound to a specific deployed contract.
func NewBonusDistributor(address common.Address, backend bind.ContractBackend) (*BonusDistributor, error) {
	contract, err := bindBonusDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BonusDistributor{BonusDistributorCaller: BonusDistributorCaller{contract: contract}, BonusDistributorTransactor: BonusDistributorTransactor{contract: contract}, BonusDistributorFilterer: BonusDistributorFilterer{contract: contract}}, nil
}

// NewBonusDistributorCaller creates a new read-only instance of BonusDistributor, bound to a specific deployed contract.
func NewBonusDistributorCaller(address common.Address, caller bind.ContractCaller) (*BonusDistributorCaller, error) {
	contract, err := bindBonusDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BonusDistributorCaller{contract: contract}, nil
}

// NewBonusDistributorTransactor creates a new write-only instance of BonusDistributor, bound to a specific deployed contract.
func NewBonusDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*BonusDistributorTransactor, error) {
	contract, err := bindBonusDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BonusDistributorTransactor{contract: contract}, nil
}

// NewBonusDistributorFilterer creates a new log filterer instance of BonusDistributor, bound to a specific deployed contract.
func NewBonusDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*BonusDistributorFilterer, error) {
	contract, err := bindBonusDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BonusDistributorFilterer{contract: contract}, nil
}

// bindBonusDistributor binds a generic wrapper to an already deployed contract.
func bindBonusDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BonusDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BonusDistributor *BonusDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BonusDistributor.Contract.BonusDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BonusDistributor *BonusDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BonusDistributor.Contract.BonusDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BonusDistributor *BonusDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BonusDistributor.Contract.BonusDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BonusDistributor *BonusDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BonusDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BonusDistributor *BonusDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BonusDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BonusDistributor *BonusDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BonusDistributor.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_BonusDistributor *BonusDistributorCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_BonusDistributor *BonusDistributorSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _BonusDistributor.Contract.BASISPOINTSDIVISOR(&_BonusDistributor.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_BonusDistributor *BonusDistributorCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _BonusDistributor.Contract.BASISPOINTSDIVISOR(&_BonusDistributor.CallOpts)
}

// BONUSDURATION is a free data retrieval call binding the contract method 0xa013ad54.
//
// Solidity: function BONUS_DURATION() view returns(uint256)
func (_BonusDistributor *BonusDistributorCaller) BONUSDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "BONUS_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSDURATION is a free data retrieval call binding the contract method 0xa013ad54.
//
// Solidity: function BONUS_DURATION() view returns(uint256)
func (_BonusDistributor *BonusDistributorSession) BONUSDURATION() (*big.Int, error) {
	return _BonusDistributor.Contract.BONUSDURATION(&_BonusDistributor.CallOpts)
}

// BONUSDURATION is a free data retrieval call binding the contract method 0xa013ad54.
//
// Solidity: function BONUS_DURATION() view returns(uint256)
func (_BonusDistributor *BonusDistributorCallerSession) BONUSDURATION() (*big.Int, error) {
	return _BonusDistributor.Contract.BONUSDURATION(&_BonusDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BonusDistributor *BonusDistributorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BonusDistributor *BonusDistributorSession) Admin() (common.Address, error) {
	return _BonusDistributor.Contract.Admin(&_BonusDistributor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_BonusDistributor *BonusDistributorCallerSession) Admin() (common.Address, error) {
	return _BonusDistributor.Contract.Admin(&_BonusDistributor.CallOpts)
}

// BonusMultiplierBasisPoints is a free data retrieval call binding the contract method 0x2f6beaa5.
//
// Solidity: function bonusMultiplierBasisPoints() view returns(uint256)
func (_BonusDistributor *BonusDistributorCaller) BonusMultiplierBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "bonusMultiplierBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusMultiplierBasisPoints is a free data retrieval call binding the contract method 0x2f6beaa5.
//
// Solidity: function bonusMultiplierBasisPoints() view returns(uint256)
func (_BonusDistributor *BonusDistributorSession) BonusMultiplierBasisPoints() (*big.Int, error) {
	return _BonusDistributor.Contract.BonusMultiplierBasisPoints(&_BonusDistributor.CallOpts)
}

// BonusMultiplierBasisPoints is a free data retrieval call binding the contract method 0x2f6beaa5.
//
// Solidity: function bonusMultiplierBasisPoints() view returns(uint256)
func (_BonusDistributor *BonusDistributorCallerSession) BonusMultiplierBasisPoints() (*big.Int, error) {
	return _BonusDistributor.Contract.BonusMultiplierBasisPoints(&_BonusDistributor.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BonusDistributor *BonusDistributorCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BonusDistributor *BonusDistributorSession) Gov() (common.Address, error) {
	return _BonusDistributor.Contract.Gov(&_BonusDistributor.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_BonusDistributor *BonusDistributorCallerSession) Gov() (common.Address, error) {
	return _BonusDistributor.Contract.Gov(&_BonusDistributor.CallOpts)
}

// LastDistributionTime is a free data retrieval call binding the contract method 0x75b17350.
//
// Solidity: function lastDistributionTime() view returns(uint256)
func (_BonusDistributor *BonusDistributorCaller) LastDistributionTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "lastDistributionTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDistributionTime is a free data retrieval call binding the contract method 0x75b17350.
//
// Solidity: function lastDistributionTime() view returns(uint256)
func (_BonusDistributor *BonusDistributorSession) LastDistributionTime() (*big.Int, error) {
	return _BonusDistributor.Contract.LastDistributionTime(&_BonusDistributor.CallOpts)
}

// LastDistributionTime is a free data retrieval call binding the contract method 0x75b17350.
//
// Solidity: function lastDistributionTime() view returns(uint256)
func (_BonusDistributor *BonusDistributorCallerSession) LastDistributionTime() (*big.Int, error) {
	return _BonusDistributor.Contract.LastDistributionTime(&_BonusDistributor.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_BonusDistributor *BonusDistributorCaller) PendingRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "pendingRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_BonusDistributor *BonusDistributorSession) PendingRewards() (*big.Int, error) {
	return _BonusDistributor.Contract.PendingRewards(&_BonusDistributor.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_BonusDistributor *BonusDistributorCallerSession) PendingRewards() (*big.Int, error) {
	return _BonusDistributor.Contract.PendingRewards(&_BonusDistributor.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_BonusDistributor *BonusDistributorCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_BonusDistributor *BonusDistributorSession) RewardToken() (common.Address, error) {
	return _BonusDistributor.Contract.RewardToken(&_BonusDistributor.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_BonusDistributor *BonusDistributorCallerSession) RewardToken() (common.Address, error) {
	return _BonusDistributor.Contract.RewardToken(&_BonusDistributor.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_BonusDistributor *BonusDistributorCaller) RewardTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "rewardTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_BonusDistributor *BonusDistributorSession) RewardTracker() (common.Address, error) {
	return _BonusDistributor.Contract.RewardTracker(&_BonusDistributor.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_BonusDistributor *BonusDistributorCallerSession) RewardTracker() (common.Address, error) {
	return _BonusDistributor.Contract.RewardTracker(&_BonusDistributor.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_BonusDistributor *BonusDistributorCaller) TokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BonusDistributor.contract.Call(opts, &out, "tokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_BonusDistributor *BonusDistributorSession) TokensPerInterval() (*big.Int, error) {
	return _BonusDistributor.Contract.TokensPerInterval(&_BonusDistributor.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_BonusDistributor *BonusDistributorCallerSession) TokensPerInterval() (*big.Int, error) {
	return _BonusDistributor.Contract.TokensPerInterval(&_BonusDistributor.CallOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_BonusDistributor *BonusDistributorTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BonusDistributor.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_BonusDistributor *BonusDistributorSession) Distribute() (*types.Transaction, error) {
	return _BonusDistributor.Contract.Distribute(&_BonusDistributor.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns(uint256)
func (_BonusDistributor *BonusDistributorTransactorSession) Distribute() (*types.Transaction, error) {
	return _BonusDistributor.Contract.Distribute(&_BonusDistributor.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_BonusDistributor *BonusDistributorTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _BonusDistributor.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_BonusDistributor *BonusDistributorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BonusDistributor.Contract.SetAdmin(&_BonusDistributor.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_BonusDistributor *BonusDistributorTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _BonusDistributor.Contract.SetAdmin(&_BonusDistributor.TransactOpts, _admin)
}

// SetBonusMultiplier is a paid mutator transaction binding the contract method 0xfd58e63a.
//
// Solidity: function setBonusMultiplier(uint256 _bonusMultiplierBasisPoints) returns()
func (_BonusDistributor *BonusDistributorTransactor) SetBonusMultiplier(opts *bind.TransactOpts, _bonusMultiplierBasisPoints *big.Int) (*types.Transaction, error) {
	return _BonusDistributor.contract.Transact(opts, "setBonusMultiplier", _bonusMultiplierBasisPoints)
}

// SetBonusMultiplier is a paid mutator transaction binding the contract method 0xfd58e63a.
//
// Solidity: function setBonusMultiplier(uint256 _bonusMultiplierBasisPoints) returns()
func (_BonusDistributor *BonusDistributorSession) SetBonusMultiplier(_bonusMultiplierBasisPoints *big.Int) (*types.Transaction, error) {
	return _BonusDistributor.Contract.SetBonusMultiplier(&_BonusDistributor.TransactOpts, _bonusMultiplierBasisPoints)
}

// SetBonusMultiplier is a paid mutator transaction binding the contract method 0xfd58e63a.
//
// Solidity: function setBonusMultiplier(uint256 _bonusMultiplierBasisPoints) returns()
func (_BonusDistributor *BonusDistributorTransactorSession) SetBonusMultiplier(_bonusMultiplierBasisPoints *big.Int) (*types.Transaction, error) {
	return _BonusDistributor.Contract.SetBonusMultiplier(&_BonusDistributor.TransactOpts, _bonusMultiplierBasisPoints)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BonusDistributor *BonusDistributorTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _BonusDistributor.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BonusDistributor *BonusDistributorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _BonusDistributor.Contract.SetGov(&_BonusDistributor.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_BonusDistributor *BonusDistributorTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _BonusDistributor.Contract.SetGov(&_BonusDistributor.TransactOpts, _gov)
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x3ae6d6eb.
//
// Solidity: function updateLastDistributionTime() returns()
func (_BonusDistributor *BonusDistributorTransactor) UpdateLastDistributionTime(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BonusDistributor.contract.Transact(opts, "updateLastDistributionTime")
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x3ae6d6eb.
//
// Solidity: function updateLastDistributionTime() returns()
func (_BonusDistributor *BonusDistributorSession) UpdateLastDistributionTime() (*types.Transaction, error) {
	return _BonusDistributor.Contract.UpdateLastDistributionTime(&_BonusDistributor.TransactOpts)
}

// UpdateLastDistributionTime is a paid mutator transaction binding the contract method 0x3ae6d6eb.
//
// Solidity: function updateLastDistributionTime() returns()
func (_BonusDistributor *BonusDistributorTransactorSession) UpdateLastDistributionTime() (*types.Transaction, error) {
	return _BonusDistributor.Contract.UpdateLastDistributionTime(&_BonusDistributor.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_BonusDistributor *BonusDistributorTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BonusDistributor.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_BonusDistributor *BonusDistributorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BonusDistributor.Contract.WithdrawToken(&_BonusDistributor.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_BonusDistributor *BonusDistributorTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BonusDistributor.Contract.WithdrawToken(&_BonusDistributor.TransactOpts, _token, _account, _amount)
}

// BonusDistributorBonusMultiplierChangeIterator is returned from FilterBonusMultiplierChange and is used to iterate over the raw logs and unpacked data for BonusMultiplierChange events raised by the BonusDistributor contract.
type BonusDistributorBonusMultiplierChangeIterator struct {
	Event *BonusDistributorBonusMultiplierChange // Event containing the contract specifics and raw log

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
func (it *BonusDistributorBonusMultiplierChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BonusDistributorBonusMultiplierChange)
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
		it.Event = new(BonusDistributorBonusMultiplierChange)
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
func (it *BonusDistributorBonusMultiplierChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BonusDistributorBonusMultiplierChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BonusDistributorBonusMultiplierChange represents a BonusMultiplierChange event raised by the BonusDistributor contract.
type BonusDistributorBonusMultiplierChange struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBonusMultiplierChange is a free log retrieval operation binding the contract event 0x58585e4edc172f6517d6789356504dc9ad23553d02202c3ac454a8d0733f6b2b.
//
// Solidity: event BonusMultiplierChange(uint256 amount)
func (_BonusDistributor *BonusDistributorFilterer) FilterBonusMultiplierChange(opts *bind.FilterOpts) (*BonusDistributorBonusMultiplierChangeIterator, error) {

	logs, sub, err := _BonusDistributor.contract.FilterLogs(opts, "BonusMultiplierChange")
	if err != nil {
		return nil, err
	}
	return &BonusDistributorBonusMultiplierChangeIterator{contract: _BonusDistributor.contract, event: "BonusMultiplierChange", logs: logs, sub: sub}, nil
}

// WatchBonusMultiplierChange is a free log subscription operation binding the contract event 0x58585e4edc172f6517d6789356504dc9ad23553d02202c3ac454a8d0733f6b2b.
//
// Solidity: event BonusMultiplierChange(uint256 amount)
func (_BonusDistributor *BonusDistributorFilterer) WatchBonusMultiplierChange(opts *bind.WatchOpts, sink chan<- *BonusDistributorBonusMultiplierChange) (event.Subscription, error) {

	logs, sub, err := _BonusDistributor.contract.WatchLogs(opts, "BonusMultiplierChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BonusDistributorBonusMultiplierChange)
				if err := _BonusDistributor.contract.UnpackLog(event, "BonusMultiplierChange", log); err != nil {
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

// ParseBonusMultiplierChange is a log parse operation binding the contract event 0x58585e4edc172f6517d6789356504dc9ad23553d02202c3ac454a8d0733f6b2b.
//
// Solidity: event BonusMultiplierChange(uint256 amount)
func (_BonusDistributor *BonusDistributorFilterer) ParseBonusMultiplierChange(log types.Log) (*BonusDistributorBonusMultiplierChange, error) {
	event := new(BonusDistributorBonusMultiplierChange)
	if err := _BonusDistributor.contract.UnpackLog(event, "BonusMultiplierChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BonusDistributorDistributeIterator is returned from FilterDistribute and is used to iterate over the raw logs and unpacked data for Distribute events raised by the BonusDistributor contract.
type BonusDistributorDistributeIterator struct {
	Event *BonusDistributorDistribute // Event containing the contract specifics and raw log

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
func (it *BonusDistributorDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BonusDistributorDistribute)
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
		it.Event = new(BonusDistributorDistribute)
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
func (it *BonusDistributorDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BonusDistributorDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BonusDistributorDistribute represents a Distribute event raised by the BonusDistributor contract.
type BonusDistributorDistribute struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistribute is a free log retrieval operation binding the contract event 0x4def474aca53bf221d07d9ab0f675b3f6d8d2494b8427271bcf43c018ef1eead.
//
// Solidity: event Distribute(uint256 amount)
func (_BonusDistributor *BonusDistributorFilterer) FilterDistribute(opts *bind.FilterOpts) (*BonusDistributorDistributeIterator, error) {

	logs, sub, err := _BonusDistributor.contract.FilterLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return &BonusDistributorDistributeIterator{contract: _BonusDistributor.contract, event: "Distribute", logs: logs, sub: sub}, nil
}

// WatchDistribute is a free log subscription operation binding the contract event 0x4def474aca53bf221d07d9ab0f675b3f6d8d2494b8427271bcf43c018ef1eead.
//
// Solidity: event Distribute(uint256 amount)
func (_BonusDistributor *BonusDistributorFilterer) WatchDistribute(opts *bind.WatchOpts, sink chan<- *BonusDistributorDistribute) (event.Subscription, error) {

	logs, sub, err := _BonusDistributor.contract.WatchLogs(opts, "Distribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BonusDistributorDistribute)
				if err := _BonusDistributor.contract.UnpackLog(event, "Distribute", log); err != nil {
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

// ParseDistribute is a log parse operation binding the contract event 0x4def474aca53bf221d07d9ab0f675b3f6d8d2494b8427271bcf43c018ef1eead.
//
// Solidity: event Distribute(uint256 amount)
func (_BonusDistributor *BonusDistributorFilterer) ParseDistribute(log types.Log) (*BonusDistributorDistribute, error) {
	event := new(BonusDistributorDistribute)
	if err := _BonusDistributor.contract.UnpackLog(event, "Distribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
