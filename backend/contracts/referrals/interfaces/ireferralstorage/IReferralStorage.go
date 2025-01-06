// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ireferralstorage

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

// IReferralStorageMetaData contains all meta data concerning the IReferralStorage contract.
var IReferralStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_code\",\"type\":\"bytes32\"}],\"name\":\"codeOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getTraderReferralInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_code\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_newAccount\",\"type\":\"address\"}],\"name\":\"govSetCodeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"referrerDiscountShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"referrerTiers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tierId\",\"type\":\"uint256\"}],\"name\":\"setReferrerTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tierId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalRebate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_discountShare\",\"type\":\"uint256\"}],\"name\":\"setTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_code\",\"type\":\"bytes32\"}],\"name\":\"setTraderReferralCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"traderReferralCodes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IReferralStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use IReferralStorageMetaData.ABI instead.
var IReferralStorageABI = IReferralStorageMetaData.ABI

// IReferralStorage is an auto generated Go binding around an Ethereum contract.
type IReferralStorage struct {
	IReferralStorageCaller     // Read-only binding to the contract
	IReferralStorageTransactor // Write-only binding to the contract
	IReferralStorageFilterer   // Log filterer for contract events
}

// IReferralStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReferralStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReferralStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReferralStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReferralStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReferralStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReferralStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReferralStorageSession struct {
	Contract     *IReferralStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IReferralStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReferralStorageCallerSession struct {
	Contract *IReferralStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IReferralStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReferralStorageTransactorSession struct {
	Contract     *IReferralStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IReferralStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReferralStorageRaw struct {
	Contract *IReferralStorage // Generic contract binding to access the raw methods on
}

// IReferralStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReferralStorageCallerRaw struct {
	Contract *IReferralStorageCaller // Generic read-only contract binding to access the raw methods on
}

// IReferralStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReferralStorageTransactorRaw struct {
	Contract *IReferralStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReferralStorage creates a new instance of IReferralStorage, bound to a specific deployed contract.
func NewIReferralStorage(address common.Address, backend bind.ContractBackend) (*IReferralStorage, error) {
	contract, err := bindIReferralStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReferralStorage{IReferralStorageCaller: IReferralStorageCaller{contract: contract}, IReferralStorageTransactor: IReferralStorageTransactor{contract: contract}, IReferralStorageFilterer: IReferralStorageFilterer{contract: contract}}, nil
}

// NewIReferralStorageCaller creates a new read-only instance of IReferralStorage, bound to a specific deployed contract.
func NewIReferralStorageCaller(address common.Address, caller bind.ContractCaller) (*IReferralStorageCaller, error) {
	contract, err := bindIReferralStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReferralStorageCaller{contract: contract}, nil
}

// NewIReferralStorageTransactor creates a new write-only instance of IReferralStorage, bound to a specific deployed contract.
func NewIReferralStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IReferralStorageTransactor, error) {
	contract, err := bindIReferralStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReferralStorageTransactor{contract: contract}, nil
}

// NewIReferralStorageFilterer creates a new log filterer instance of IReferralStorage, bound to a specific deployed contract.
func NewIReferralStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IReferralStorageFilterer, error) {
	contract, err := bindIReferralStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReferralStorageFilterer{contract: contract}, nil
}

// bindIReferralStorage binds a generic wrapper to an already deployed contract.
func bindIReferralStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IReferralStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReferralStorage *IReferralStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReferralStorage.Contract.IReferralStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReferralStorage *IReferralStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReferralStorage.Contract.IReferralStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReferralStorage *IReferralStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReferralStorage.Contract.IReferralStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReferralStorage *IReferralStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReferralStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReferralStorage *IReferralStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReferralStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReferralStorage *IReferralStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReferralStorage.Contract.contract.Transact(opts, method, params...)
}

// CodeOwners is a free data retrieval call binding the contract method 0xc8b3c460.
//
// Solidity: function codeOwners(bytes32 _code) view returns(address)
func (_IReferralStorage *IReferralStorageCaller) CodeOwners(opts *bind.CallOpts, _code [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IReferralStorage.contract.Call(opts, &out, "codeOwners", _code)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CodeOwners is a free data retrieval call binding the contract method 0xc8b3c460.
//
// Solidity: function codeOwners(bytes32 _code) view returns(address)
func (_IReferralStorage *IReferralStorageSession) CodeOwners(_code [32]byte) (common.Address, error) {
	return _IReferralStorage.Contract.CodeOwners(&_IReferralStorage.CallOpts, _code)
}

// CodeOwners is a free data retrieval call binding the contract method 0xc8b3c460.
//
// Solidity: function codeOwners(bytes32 _code) view returns(address)
func (_IReferralStorage *IReferralStorageCallerSession) CodeOwners(_code [32]byte) (common.Address, error) {
	return _IReferralStorage.Contract.CodeOwners(&_IReferralStorage.CallOpts, _code)
}

// GetTraderReferralInfo is a free data retrieval call binding the contract method 0x534ef883.
//
// Solidity: function getTraderReferralInfo(address _account) view returns(bytes32, address)
func (_IReferralStorage *IReferralStorageCaller) GetTraderReferralInfo(opts *bind.CallOpts, _account common.Address) ([32]byte, common.Address, error) {
	var out []interface{}
	err := _IReferralStorage.contract.Call(opts, &out, "getTraderReferralInfo", _account)

	if err != nil {
		return *new([32]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetTraderReferralInfo is a free data retrieval call binding the contract method 0x534ef883.
//
// Solidity: function getTraderReferralInfo(address _account) view returns(bytes32, address)
func (_IReferralStorage *IReferralStorageSession) GetTraderReferralInfo(_account common.Address) ([32]byte, common.Address, error) {
	return _IReferralStorage.Contract.GetTraderReferralInfo(&_IReferralStorage.CallOpts, _account)
}

// GetTraderReferralInfo is a free data retrieval call binding the contract method 0x534ef883.
//
// Solidity: function getTraderReferralInfo(address _account) view returns(bytes32, address)
func (_IReferralStorage *IReferralStorageCallerSession) GetTraderReferralInfo(_account common.Address) ([32]byte, common.Address, error) {
	return _IReferralStorage.Contract.GetTraderReferralInfo(&_IReferralStorage.CallOpts, _account)
}

// ReferrerDiscountShares is a free data retrieval call binding the contract method 0x71a6a790.
//
// Solidity: function referrerDiscountShares(address _account) view returns(uint256)
func (_IReferralStorage *IReferralStorageCaller) ReferrerDiscountShares(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IReferralStorage.contract.Call(opts, &out, "referrerDiscountShares", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerDiscountShares is a free data retrieval call binding the contract method 0x71a6a790.
//
// Solidity: function referrerDiscountShares(address _account) view returns(uint256)
func (_IReferralStorage *IReferralStorageSession) ReferrerDiscountShares(_account common.Address) (*big.Int, error) {
	return _IReferralStorage.Contract.ReferrerDiscountShares(&_IReferralStorage.CallOpts, _account)
}

// ReferrerDiscountShares is a free data retrieval call binding the contract method 0x71a6a790.
//
// Solidity: function referrerDiscountShares(address _account) view returns(uint256)
func (_IReferralStorage *IReferralStorageCallerSession) ReferrerDiscountShares(_account common.Address) (*big.Int, error) {
	return _IReferralStorage.Contract.ReferrerDiscountShares(&_IReferralStorage.CallOpts, _account)
}

// ReferrerTiers is a free data retrieval call binding the contract method 0x1582a018.
//
// Solidity: function referrerTiers(address _account) view returns(uint256)
func (_IReferralStorage *IReferralStorageCaller) ReferrerTiers(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IReferralStorage.contract.Call(opts, &out, "referrerTiers", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerTiers is a free data retrieval call binding the contract method 0x1582a018.
//
// Solidity: function referrerTiers(address _account) view returns(uint256)
func (_IReferralStorage *IReferralStorageSession) ReferrerTiers(_account common.Address) (*big.Int, error) {
	return _IReferralStorage.Contract.ReferrerTiers(&_IReferralStorage.CallOpts, _account)
}

// ReferrerTiers is a free data retrieval call binding the contract method 0x1582a018.
//
// Solidity: function referrerTiers(address _account) view returns(uint256)
func (_IReferralStorage *IReferralStorageCallerSession) ReferrerTiers(_account common.Address) (*big.Int, error) {
	return _IReferralStorage.Contract.ReferrerTiers(&_IReferralStorage.CallOpts, _account)
}

// TraderReferralCodes is a free data retrieval call binding the contract method 0x85725b58.
//
// Solidity: function traderReferralCodes(address _account) view returns(bytes32)
func (_IReferralStorage *IReferralStorageCaller) TraderReferralCodes(opts *bind.CallOpts, _account common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IReferralStorage.contract.Call(opts, &out, "traderReferralCodes", _account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TraderReferralCodes is a free data retrieval call binding the contract method 0x85725b58.
//
// Solidity: function traderReferralCodes(address _account) view returns(bytes32)
func (_IReferralStorage *IReferralStorageSession) TraderReferralCodes(_account common.Address) ([32]byte, error) {
	return _IReferralStorage.Contract.TraderReferralCodes(&_IReferralStorage.CallOpts, _account)
}

// TraderReferralCodes is a free data retrieval call binding the contract method 0x85725b58.
//
// Solidity: function traderReferralCodes(address _account) view returns(bytes32)
func (_IReferralStorage *IReferralStorageCallerSession) TraderReferralCodes(_account common.Address) ([32]byte, error) {
	return _IReferralStorage.Contract.TraderReferralCodes(&_IReferralStorage.CallOpts, _account)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0xdfcfa250.
//
// Solidity: function govSetCodeOwner(bytes32 _code, address _newAccount) returns()
func (_IReferralStorage *IReferralStorageTransactor) GovSetCodeOwner(opts *bind.TransactOpts, _code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _IReferralStorage.contract.Transact(opts, "govSetCodeOwner", _code, _newAccount)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0xdfcfa250.
//
// Solidity: function govSetCodeOwner(bytes32 _code, address _newAccount) returns()
func (_IReferralStorage *IReferralStorageSession) GovSetCodeOwner(_code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _IReferralStorage.Contract.GovSetCodeOwner(&_IReferralStorage.TransactOpts, _code, _newAccount)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0xdfcfa250.
//
// Solidity: function govSetCodeOwner(bytes32 _code, address _newAccount) returns()
func (_IReferralStorage *IReferralStorageTransactorSession) GovSetCodeOwner(_code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _IReferralStorage.Contract.GovSetCodeOwner(&_IReferralStorage.TransactOpts, _code, _newAccount)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0x3fb8b323.
//
// Solidity: function setReferrerTier(address _referrer, uint256 _tierId) returns()
func (_IReferralStorage *IReferralStorageTransactor) SetReferrerTier(opts *bind.TransactOpts, _referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _IReferralStorage.contract.Transact(opts, "setReferrerTier", _referrer, _tierId)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0x3fb8b323.
//
// Solidity: function setReferrerTier(address _referrer, uint256 _tierId) returns()
func (_IReferralStorage *IReferralStorageSession) SetReferrerTier(_referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _IReferralStorage.Contract.SetReferrerTier(&_IReferralStorage.TransactOpts, _referrer, _tierId)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0x3fb8b323.
//
// Solidity: function setReferrerTier(address _referrer, uint256 _tierId) returns()
func (_IReferralStorage *IReferralStorageTransactorSession) SetReferrerTier(_referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _IReferralStorage.Contract.SetReferrerTier(&_IReferralStorage.TransactOpts, _referrer, _tierId)
}

// SetTier is a paid mutator transaction binding the contract method 0x836a0187.
//
// Solidity: function setTier(uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_IReferralStorage *IReferralStorageTransactor) SetTier(opts *bind.TransactOpts, _tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _IReferralStorage.contract.Transact(opts, "setTier", _tierId, _totalRebate, _discountShare)
}

// SetTier is a paid mutator transaction binding the contract method 0x836a0187.
//
// Solidity: function setTier(uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_IReferralStorage *IReferralStorageSession) SetTier(_tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _IReferralStorage.Contract.SetTier(&_IReferralStorage.TransactOpts, _tierId, _totalRebate, _discountShare)
}

// SetTier is a paid mutator transaction binding the contract method 0x836a0187.
//
// Solidity: function setTier(uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_IReferralStorage *IReferralStorageTransactorSession) SetTier(_tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _IReferralStorage.Contract.SetTier(&_IReferralStorage.TransactOpts, _tierId, _totalRebate, _discountShare)
}

// SetTraderReferralCode is a paid mutator transaction binding the contract method 0x56b4b2ad.
//
// Solidity: function setTraderReferralCode(address _account, bytes32 _code) returns()
func (_IReferralStorage *IReferralStorageTransactor) SetTraderReferralCode(opts *bind.TransactOpts, _account common.Address, _code [32]byte) (*types.Transaction, error) {
	return _IReferralStorage.contract.Transact(opts, "setTraderReferralCode", _account, _code)
}

// SetTraderReferralCode is a paid mutator transaction binding the contract method 0x56b4b2ad.
//
// Solidity: function setTraderReferralCode(address _account, bytes32 _code) returns()
func (_IReferralStorage *IReferralStorageSession) SetTraderReferralCode(_account common.Address, _code [32]byte) (*types.Transaction, error) {
	return _IReferralStorage.Contract.SetTraderReferralCode(&_IReferralStorage.TransactOpts, _account, _code)
}

// SetTraderReferralCode is a paid mutator transaction binding the contract method 0x56b4b2ad.
//
// Solidity: function setTraderReferralCode(address _account, bytes32 _code) returns()
func (_IReferralStorage *IReferralStorageTransactorSession) SetTraderReferralCode(_account common.Address, _code [32]byte) (*types.Transaction, error) {
	return _IReferralStorage.Contract.SetTraderReferralCode(&_IReferralStorage.TransactOpts, _account, _code)
}
