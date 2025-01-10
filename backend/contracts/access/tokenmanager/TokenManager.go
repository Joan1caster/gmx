// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenmanager

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

// TokenManagerMetaData contains all meta data concerning the TokenManager contract.
var TokenManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAuthorizations\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ClearAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalApproveNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalApproveNFTs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalPendingAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalSetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalSetGov\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionsNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAuthorizations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingActions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"receiveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signApproveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signApproveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signSetAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"signalApproveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"signalApproveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"signalSetAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"signalSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedActions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TokenManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenManagerMetaData.ABI instead.
var TokenManagerABI = TokenManagerMetaData.ABI

// TokenManager is an auto generated Go binding around an Ethereum contract.
type TokenManager struct {
	TokenManagerCaller     // Read-only binding to the contract
	TokenManagerTransactor // Write-only binding to the contract
	TokenManagerFilterer   // Log filterer for contract events
}

// TokenManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenManagerSession struct {
	Contract     *TokenManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenManagerCallerSession struct {
	Contract *TokenManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenManagerTransactorSession struct {
	Contract     *TokenManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenManagerRaw struct {
	Contract *TokenManager // Generic contract binding to access the raw methods on
}

// TokenManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenManagerCallerRaw struct {
	Contract *TokenManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenManagerTransactorRaw struct {
	Contract *TokenManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenManager creates a new instance of TokenManager, bound to a specific deployed contract.
func NewTokenManager(address common.Address, backend bind.ContractBackend) (*TokenManager, error) {
	contract, err := bindTokenManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenManager{TokenManagerCaller: TokenManagerCaller{contract: contract}, TokenManagerTransactor: TokenManagerTransactor{contract: contract}, TokenManagerFilterer: TokenManagerFilterer{contract: contract}}, nil
}

// NewTokenManagerCaller creates a new read-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerCaller(address common.Address, caller bind.ContractCaller) (*TokenManagerCaller, error) {
	contract, err := bindTokenManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerCaller{contract: contract}, nil
}

// NewTokenManagerTransactor creates a new write-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenManagerTransactor, error) {
	contract, err := bindTokenManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerTransactor{contract: contract}, nil
}

// NewTokenManagerFilterer creates a new log filterer instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenManagerFilterer, error) {
	contract, err := bindTokenManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenManagerFilterer{contract: contract}, nil
}

// bindTokenManager binds a generic wrapper to an already deployed contract.
func bindTokenManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.TokenManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transact(opts, method, params...)
}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_TokenManager *TokenManagerCaller) ActionsNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "actionsNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_TokenManager *TokenManagerSession) ActionsNonce() (*big.Int, error) {
	return _TokenManager.Contract.ActionsNonce(&_TokenManager.CallOpts)
}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_TokenManager *TokenManagerCallerSession) ActionsNonce() (*big.Int, error) {
	return _TokenManager.Contract.ActionsNonce(&_TokenManager.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TokenManager *TokenManagerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TokenManager *TokenManagerSession) Admin() (common.Address, error) {
	return _TokenManager.Contract.Admin(&_TokenManager.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TokenManager *TokenManagerCallerSession) Admin() (common.Address, error) {
	return _TokenManager.Contract.Admin(&_TokenManager.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_TokenManager *TokenManagerCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_TokenManager *TokenManagerSession) IsInitialized() (bool, error) {
	return _TokenManager.Contract.IsInitialized(&_TokenManager.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_TokenManager *TokenManagerCallerSession) IsInitialized() (bool, error) {
	return _TokenManager.Contract.IsInitialized(&_TokenManager.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_TokenManager *TokenManagerCaller) IsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "isSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_TokenManager *TokenManagerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _TokenManager.Contract.IsSigner(&_TokenManager.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_TokenManager *TokenManagerCallerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _TokenManager.Contract.IsSigner(&_TokenManager.CallOpts, arg0)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_TokenManager *TokenManagerCaller) MinAuthorizations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "minAuthorizations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_TokenManager *TokenManagerSession) MinAuthorizations() (*big.Int, error) {
	return _TokenManager.Contract.MinAuthorizations(&_TokenManager.CallOpts)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_TokenManager *TokenManagerCallerSession) MinAuthorizations() (*big.Int, error) {
	return _TokenManager.Contract.MinAuthorizations(&_TokenManager.CallOpts)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_TokenManager *TokenManagerCaller) PendingActions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "pendingActions", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_TokenManager *TokenManagerSession) PendingActions(arg0 [32]byte) (bool, error) {
	return _TokenManager.Contract.PendingActions(&_TokenManager.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_TokenManager *TokenManagerCallerSession) PendingActions(arg0 [32]byte) (bool, error) {
	return _TokenManager.Contract.PendingActions(&_TokenManager.CallOpts, arg0)
}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_TokenManager *TokenManagerCaller) SignedActions(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "signedActions", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_TokenManager *TokenManagerSession) SignedActions(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _TokenManager.Contract.SignedActions(&_TokenManager.CallOpts, arg0, arg1)
}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_TokenManager *TokenManagerCallerSession) SignedActions(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _TokenManager.Contract.SignedActions(&_TokenManager.CallOpts, arg0, arg1)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_TokenManager *TokenManagerCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_TokenManager *TokenManagerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _TokenManager.Contract.Signers(&_TokenManager.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_TokenManager *TokenManagerCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _TokenManager.Contract.Signers(&_TokenManager.CallOpts, arg0)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() view returns(uint256)
func (_TokenManager *TokenManagerCaller) SignersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenManager.contract.Call(opts, &out, "signersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() view returns(uint256)
func (_TokenManager *TokenManagerSession) SignersLength() (*big.Int, error) {
	return _TokenManager.Contract.SignersLength(&_TokenManager.CallOpts)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() view returns(uint256)
func (_TokenManager *TokenManagerCallerSession) SignersLength() (*big.Int, error) {
	return _TokenManager.Contract.SignersLength(&_TokenManager.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "approve", _token, _spender, _amount, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.Approve(&_TokenManager.TransactOpts, _token, _spender, _amount, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.Approve(&_TokenManager.TransactOpts, _token, _spender, _amount, _nonce)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xf23f9775.
//
// Solidity: function approveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) ApproveNFT(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "approveNFT", _token, _spender, _tokenId, _nonce)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xf23f9775.
//
// Solidity: function approveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) ApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ApproveNFT(&_TokenManager.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xf23f9775.
//
// Solidity: function approveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) ApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ApproveNFT(&_TokenManager.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// ApproveNFTs is a paid mutator transaction binding the contract method 0x42a1fcee.
//
// Solidity: function approveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) ApproveNFTs(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "approveNFTs", _token, _spender, _tokenIds, _nonce)
}

// ApproveNFTs is a paid mutator transaction binding the contract method 0x42a1fcee.
//
// Solidity: function approveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) ApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ApproveNFTs(&_TokenManager.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// ApproveNFTs is a paid mutator transaction binding the contract method 0x42a1fcee.
//
// Solidity: function approveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) ApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ApproveNFTs(&_TokenManager.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _signers) returns()
func (_TokenManager *TokenManagerTransactor) Initialize(opts *bind.TransactOpts, _signers []common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "initialize", _signers)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _signers) returns()
func (_TokenManager *TokenManagerSession) Initialize(_signers []common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Initialize(&_TokenManager.TransactOpts, _signers)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _signers) returns()
func (_TokenManager *TokenManagerTransactorSession) Initialize(_signers []common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Initialize(&_TokenManager.TransactOpts, _signers)
}

// ReceiveNFTs is a paid mutator transaction binding the contract method 0xadb384b6.
//
// Solidity: function receiveNFTs(address _token, address _sender, uint256[] _tokenIds) returns()
func (_TokenManager *TokenManagerTransactor) ReceiveNFTs(opts *bind.TransactOpts, _token common.Address, _sender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "receiveNFTs", _token, _sender, _tokenIds)
}

// ReceiveNFTs is a paid mutator transaction binding the contract method 0xadb384b6.
//
// Solidity: function receiveNFTs(address _token, address _sender, uint256[] _tokenIds) returns()
func (_TokenManager *TokenManagerSession) ReceiveNFTs(_token common.Address, _sender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ReceiveNFTs(&_TokenManager.TransactOpts, _token, _sender, _tokenIds)
}

// ReceiveNFTs is a paid mutator transaction binding the contract method 0xadb384b6.
//
// Solidity: function receiveNFTs(address _token, address _sender, uint256[] _tokenIds) returns()
func (_TokenManager *TokenManagerTransactorSession) ReceiveNFTs(_token common.Address, _sender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.ReceiveNFTs(&_TokenManager.TransactOpts, _token, _sender, _tokenIds)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x9fddaac1.
//
// Solidity: function setAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) SetAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "setAdmin", _target, _admin, _nonce)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x9fddaac1.
//
// Solidity: function setAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) SetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SetAdmin(&_TokenManager.TransactOpts, _target, _admin, _nonce)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x9fddaac1.
//
// Solidity: function setAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) SetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SetAdmin(&_TokenManager.TransactOpts, _target, _admin, _nonce)
}

// SetGov is a paid mutator transaction binding the contract method 0xe9075621.
//
// Solidity: function setGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) SetGov(opts *bind.TransactOpts, _timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "setGov", _timelock, _target, _gov, _nonce)
}

// SetGov is a paid mutator transaction binding the contract method 0xe9075621.
//
// Solidity: function setGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) SetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SetGov(&_TokenManager.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SetGov is a paid mutator transaction binding the contract method 0xe9075621.
//
// Solidity: function setGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) SetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SetGov(&_TokenManager.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) SignApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signApprove", _token, _spender, _amount, _nonce)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) SignApprove(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignApprove(&_TokenManager.TransactOpts, _token, _spender, _amount, _nonce)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) SignApprove(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignApprove(&_TokenManager.TransactOpts, _token, _spender, _amount, _nonce)
}

// SignApproveNFT is a paid mutator transaction binding the contract method 0xddf67a9f.
//
// Solidity: function signApproveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) SignApproveNFT(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signApproveNFT", _token, _spender, _tokenId, _nonce)
}

// SignApproveNFT is a paid mutator transaction binding the contract method 0xddf67a9f.
//
// Solidity: function signApproveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) SignApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignApproveNFT(&_TokenManager.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// SignApproveNFT is a paid mutator transaction binding the contract method 0xddf67a9f.
//
// Solidity: function signApproveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) SignApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignApproveNFT(&_TokenManager.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// SignApproveNFTs is a paid mutator transaction binding the contract method 0xf00cb942.
//
// Solidity: function signApproveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) SignApproveNFTs(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signApproveNFTs", _token, _spender, _tokenIds, _nonce)
}

// SignApproveNFTs is a paid mutator transaction binding the contract method 0xf00cb942.
//
// Solidity: function signApproveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) SignApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignApproveNFTs(&_TokenManager.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// SignApproveNFTs is a paid mutator transaction binding the contract method 0xf00cb942.
//
// Solidity: function signApproveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) SignApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignApproveNFTs(&_TokenManager.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// SignSetAdmin is a paid mutator transaction binding the contract method 0xf466634b.
//
// Solidity: function signSetAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) SignSetAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signSetAdmin", _target, _admin, _nonce)
}

// SignSetAdmin is a paid mutator transaction binding the contract method 0xf466634b.
//
// Solidity: function signSetAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) SignSetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignSetAdmin(&_TokenManager.TransactOpts, _target, _admin, _nonce)
}

// SignSetAdmin is a paid mutator transaction binding the contract method 0xf466634b.
//
// Solidity: function signSetAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) SignSetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignSetAdmin(&_TokenManager.TransactOpts, _target, _admin, _nonce)
}

// SignSetGov is a paid mutator transaction binding the contract method 0xa62fb170.
//
// Solidity: function signSetGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactor) SignSetGov(opts *bind.TransactOpts, _timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signSetGov", _timelock, _target, _gov, _nonce)
}

// SignSetGov is a paid mutator transaction binding the contract method 0xa62fb170.
//
// Solidity: function signSetGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_TokenManager *TokenManagerSession) SignSetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignSetGov(&_TokenManager.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SignSetGov is a paid mutator transaction binding the contract method 0xa62fb170.
//
// Solidity: function signSetGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_TokenManager *TokenManagerTransactorSession) SignSetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignSetGov(&_TokenManager.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_TokenManager *TokenManagerTransactor) SignalApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signalApprove", _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_TokenManager *TokenManagerSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalApprove(&_TokenManager.TransactOpts, _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_TokenManager *TokenManagerTransactorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalApprove(&_TokenManager.TransactOpts, _token, _spender, _amount)
}

// SignalApproveNFT is a paid mutator transaction binding the contract method 0x0b13beca.
//
// Solidity: function signalApproveNFT(address _token, address _spender, uint256 _tokenId) returns()
func (_TokenManager *TokenManagerTransactor) SignalApproveNFT(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signalApproveNFT", _token, _spender, _tokenId)
}

// SignalApproveNFT is a paid mutator transaction binding the contract method 0x0b13beca.
//
// Solidity: function signalApproveNFT(address _token, address _spender, uint256 _tokenId) returns()
func (_TokenManager *TokenManagerSession) SignalApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalApproveNFT(&_TokenManager.TransactOpts, _token, _spender, _tokenId)
}

// SignalApproveNFT is a paid mutator transaction binding the contract method 0x0b13beca.
//
// Solidity: function signalApproveNFT(address _token, address _spender, uint256 _tokenId) returns()
func (_TokenManager *TokenManagerTransactorSession) SignalApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalApproveNFT(&_TokenManager.TransactOpts, _token, _spender, _tokenId)
}

// SignalApproveNFTs is a paid mutator transaction binding the contract method 0x181fcd33.
//
// Solidity: function signalApproveNFTs(address _token, address _spender, uint256[] _tokenIds) returns()
func (_TokenManager *TokenManagerTransactor) SignalApproveNFTs(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signalApproveNFTs", _token, _spender, _tokenIds)
}

// SignalApproveNFTs is a paid mutator transaction binding the contract method 0x181fcd33.
//
// Solidity: function signalApproveNFTs(address _token, address _spender, uint256[] _tokenIds) returns()
func (_TokenManager *TokenManagerSession) SignalApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalApproveNFTs(&_TokenManager.TransactOpts, _token, _spender, _tokenIds)
}

// SignalApproveNFTs is a paid mutator transaction binding the contract method 0x181fcd33.
//
// Solidity: function signalApproveNFTs(address _token, address _spender, uint256[] _tokenIds) returns()
func (_TokenManager *TokenManagerTransactorSession) SignalApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalApproveNFTs(&_TokenManager.TransactOpts, _token, _spender, _tokenIds)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x75fd490c.
//
// Solidity: function signalSetAdmin(address _target, address _admin) returns()
func (_TokenManager *TokenManagerTransactor) SignalSetAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signalSetAdmin", _target, _admin)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x75fd490c.
//
// Solidity: function signalSetAdmin(address _target, address _admin) returns()
func (_TokenManager *TokenManagerSession) SignalSetAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalSetAdmin(&_TokenManager.TransactOpts, _target, _admin)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x75fd490c.
//
// Solidity: function signalSetAdmin(address _target, address _admin) returns()
func (_TokenManager *TokenManagerTransactorSession) SignalSetAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalSetAdmin(&_TokenManager.TransactOpts, _target, _admin)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0xe2d03cc5.
//
// Solidity: function signalSetGov(address _timelock, address _target, address _gov) returns()
func (_TokenManager *TokenManagerTransactor) SignalSetGov(opts *bind.TransactOpts, _timelock common.Address, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "signalSetGov", _timelock, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0xe2d03cc5.
//
// Solidity: function signalSetGov(address _timelock, address _target, address _gov) returns()
func (_TokenManager *TokenManagerSession) SignalSetGov(_timelock common.Address, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalSetGov(&_TokenManager.TransactOpts, _timelock, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0xe2d03cc5.
//
// Solidity: function signalSetGov(address _timelock, address _target, address _gov) returns()
func (_TokenManager *TokenManagerTransactorSession) SignalSetGov(_timelock common.Address, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.SignalSetGov(&_TokenManager.TransactOpts, _timelock, _target, _gov)
}

// TokenManagerClearActionIterator is returned from FilterClearAction and is used to iterate over the raw logs and unpacked data for ClearAction events raised by the TokenManager contract.
type TokenManagerClearActionIterator struct {
	Event *TokenManagerClearAction // Event containing the contract specifics and raw log

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
func (it *TokenManagerClearActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerClearAction)
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
		it.Event = new(TokenManagerClearAction)
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
func (it *TokenManagerClearActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerClearActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerClearAction represents a ClearAction event raised by the TokenManager contract.
type TokenManagerClearAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClearAction is a free log retrieval operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterClearAction(opts *bind.FilterOpts) (*TokenManagerClearActionIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return &TokenManagerClearActionIterator{contract: _TokenManager.contract, event: "ClearAction", logs: logs, sub: sub}, nil
}

// WatchClearAction is a free log subscription operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchClearAction(opts *bind.WatchOpts, sink chan<- *TokenManagerClearAction) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerClearAction)
				if err := _TokenManager.contract.UnpackLog(event, "ClearAction", log); err != nil {
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

// ParseClearAction is a log parse operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseClearAction(log types.Log) (*TokenManagerClearAction, error) {
	event := new(TokenManagerClearAction)
	if err := _TokenManager.contract.UnpackLog(event, "ClearAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerSignActionIterator is returned from FilterSignAction and is used to iterate over the raw logs and unpacked data for SignAction events raised by the TokenManager contract.
type TokenManagerSignActionIterator struct {
	Event *TokenManagerSignAction // Event containing the contract specifics and raw log

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
func (it *TokenManagerSignActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerSignAction)
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
		it.Event = new(TokenManagerSignAction)
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
func (it *TokenManagerSignActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerSignActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerSignAction represents a SignAction event raised by the TokenManager contract.
type TokenManagerSignAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignAction is a free log retrieval operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterSignAction(opts *bind.FilterOpts) (*TokenManagerSignActionIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "SignAction")
	if err != nil {
		return nil, err
	}
	return &TokenManagerSignActionIterator{contract: _TokenManager.contract, event: "SignAction", logs: logs, sub: sub}, nil
}

// WatchSignAction is a free log subscription operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchSignAction(opts *bind.WatchOpts, sink chan<- *TokenManagerSignAction) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "SignAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerSignAction)
				if err := _TokenManager.contract.UnpackLog(event, "SignAction", log); err != nil {
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

// ParseSignAction is a log parse operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseSignAction(log types.Log) (*TokenManagerSignAction, error) {
	event := new(TokenManagerSignAction)
	if err := _TokenManager.contract.UnpackLog(event, "SignAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerSignalApproveIterator is returned from FilterSignalApprove and is used to iterate over the raw logs and unpacked data for SignalApprove events raised by the TokenManager contract.
type TokenManagerSignalApproveIterator struct {
	Event *TokenManagerSignalApprove // Event containing the contract specifics and raw log

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
func (it *TokenManagerSignalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerSignalApprove)
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
		it.Event = new(TokenManagerSignalApprove)
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
func (it *TokenManagerSignalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerSignalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerSignalApprove represents a SignalApprove event raised by the TokenManager contract.
type TokenManagerSignalApprove struct {
	Token   common.Address
	Spender common.Address
	Amount  *big.Int
	Action  [32]byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignalApprove is a free log retrieval operation binding the contract event 0xc19251bf5f704ddc3d5babe6f4e5bde0dded20b19f7844716861821ab3163cd7.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterSignalApprove(opts *bind.FilterOpts) (*TokenManagerSignalApproveIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return &TokenManagerSignalApproveIterator{contract: _TokenManager.contract, event: "SignalApprove", logs: logs, sub: sub}, nil
}

// WatchSignalApprove is a free log subscription operation binding the contract event 0xc19251bf5f704ddc3d5babe6f4e5bde0dded20b19f7844716861821ab3163cd7.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchSignalApprove(opts *bind.WatchOpts, sink chan<- *TokenManagerSignalApprove) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerSignalApprove)
				if err := _TokenManager.contract.UnpackLog(event, "SignalApprove", log); err != nil {
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

// ParseSignalApprove is a log parse operation binding the contract event 0xc19251bf5f704ddc3d5babe6f4e5bde0dded20b19f7844716861821ab3163cd7.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseSignalApprove(log types.Log) (*TokenManagerSignalApprove, error) {
	event := new(TokenManagerSignalApprove)
	if err := _TokenManager.contract.UnpackLog(event, "SignalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerSignalApproveNFTIterator is returned from FilterSignalApproveNFT and is used to iterate over the raw logs and unpacked data for SignalApproveNFT events raised by the TokenManager contract.
type TokenManagerSignalApproveNFTIterator struct {
	Event *TokenManagerSignalApproveNFT // Event containing the contract specifics and raw log

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
func (it *TokenManagerSignalApproveNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerSignalApproveNFT)
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
		it.Event = new(TokenManagerSignalApproveNFT)
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
func (it *TokenManagerSignalApproveNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerSignalApproveNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerSignalApproveNFT represents a SignalApproveNFT event raised by the TokenManager contract.
type TokenManagerSignalApproveNFT struct {
	Token   common.Address
	Spender common.Address
	TokenId *big.Int
	Action  [32]byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignalApproveNFT is a free log retrieval operation binding the contract event 0xcd9ba83b63715dc15ac193645d6e925bf4b487c94b73d709b8b6dea608efd4cc.
//
// Solidity: event SignalApproveNFT(address token, address spender, uint256 tokenId, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterSignalApproveNFT(opts *bind.FilterOpts) (*TokenManagerSignalApproveNFTIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "SignalApproveNFT")
	if err != nil {
		return nil, err
	}
	return &TokenManagerSignalApproveNFTIterator{contract: _TokenManager.contract, event: "SignalApproveNFT", logs: logs, sub: sub}, nil
}

// WatchSignalApproveNFT is a free log subscription operation binding the contract event 0xcd9ba83b63715dc15ac193645d6e925bf4b487c94b73d709b8b6dea608efd4cc.
//
// Solidity: event SignalApproveNFT(address token, address spender, uint256 tokenId, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchSignalApproveNFT(opts *bind.WatchOpts, sink chan<- *TokenManagerSignalApproveNFT) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "SignalApproveNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerSignalApproveNFT)
				if err := _TokenManager.contract.UnpackLog(event, "SignalApproveNFT", log); err != nil {
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

// ParseSignalApproveNFT is a log parse operation binding the contract event 0xcd9ba83b63715dc15ac193645d6e925bf4b487c94b73d709b8b6dea608efd4cc.
//
// Solidity: event SignalApproveNFT(address token, address spender, uint256 tokenId, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseSignalApproveNFT(log types.Log) (*TokenManagerSignalApproveNFT, error) {
	event := new(TokenManagerSignalApproveNFT)
	if err := _TokenManager.contract.UnpackLog(event, "SignalApproveNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerSignalApproveNFTsIterator is returned from FilterSignalApproveNFTs and is used to iterate over the raw logs and unpacked data for SignalApproveNFTs events raised by the TokenManager contract.
type TokenManagerSignalApproveNFTsIterator struct {
	Event *TokenManagerSignalApproveNFTs // Event containing the contract specifics and raw log

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
func (it *TokenManagerSignalApproveNFTsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerSignalApproveNFTs)
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
		it.Event = new(TokenManagerSignalApproveNFTs)
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
func (it *TokenManagerSignalApproveNFTsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerSignalApproveNFTsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerSignalApproveNFTs represents a SignalApproveNFTs event raised by the TokenManager contract.
type TokenManagerSignalApproveNFTs struct {
	Token    common.Address
	Spender  common.Address
	TokenIds []*big.Int
	Action   [32]byte
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalApproveNFTs is a free log retrieval operation binding the contract event 0xf9d0354d71c261982d98abd09b735f3663b2d7275e2569ad5fd907a4092765f9.
//
// Solidity: event SignalApproveNFTs(address token, address spender, uint256[] tokenIds, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterSignalApproveNFTs(opts *bind.FilterOpts) (*TokenManagerSignalApproveNFTsIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "SignalApproveNFTs")
	if err != nil {
		return nil, err
	}
	return &TokenManagerSignalApproveNFTsIterator{contract: _TokenManager.contract, event: "SignalApproveNFTs", logs: logs, sub: sub}, nil
}

// WatchSignalApproveNFTs is a free log subscription operation binding the contract event 0xf9d0354d71c261982d98abd09b735f3663b2d7275e2569ad5fd907a4092765f9.
//
// Solidity: event SignalApproveNFTs(address token, address spender, uint256[] tokenIds, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchSignalApproveNFTs(opts *bind.WatchOpts, sink chan<- *TokenManagerSignalApproveNFTs) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "SignalApproveNFTs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerSignalApproveNFTs)
				if err := _TokenManager.contract.UnpackLog(event, "SignalApproveNFTs", log); err != nil {
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

// ParseSignalApproveNFTs is a log parse operation binding the contract event 0xf9d0354d71c261982d98abd09b735f3663b2d7275e2569ad5fd907a4092765f9.
//
// Solidity: event SignalApproveNFTs(address token, address spender, uint256[] tokenIds, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseSignalApproveNFTs(log types.Log) (*TokenManagerSignalApproveNFTs, error) {
	event := new(TokenManagerSignalApproveNFTs)
	if err := _TokenManager.contract.UnpackLog(event, "SignalApproveNFTs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerSignalPendingActionIterator is returned from FilterSignalPendingAction and is used to iterate over the raw logs and unpacked data for SignalPendingAction events raised by the TokenManager contract.
type TokenManagerSignalPendingActionIterator struct {
	Event *TokenManagerSignalPendingAction // Event containing the contract specifics and raw log

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
func (it *TokenManagerSignalPendingActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerSignalPendingAction)
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
		it.Event = new(TokenManagerSignalPendingAction)
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
func (it *TokenManagerSignalPendingActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerSignalPendingActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerSignalPendingAction represents a SignalPendingAction event raised by the TokenManager contract.
type TokenManagerSignalPendingAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalPendingAction is a free log retrieval operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterSignalPendingAction(opts *bind.FilterOpts) (*TokenManagerSignalPendingActionIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return &TokenManagerSignalPendingActionIterator{contract: _TokenManager.contract, event: "SignalPendingAction", logs: logs, sub: sub}, nil
}

// WatchSignalPendingAction is a free log subscription operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchSignalPendingAction(opts *bind.WatchOpts, sink chan<- *TokenManagerSignalPendingAction) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerSignalPendingAction)
				if err := _TokenManager.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
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

// ParseSignalPendingAction is a log parse operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseSignalPendingAction(log types.Log) (*TokenManagerSignalPendingAction, error) {
	event := new(TokenManagerSignalPendingAction)
	if err := _TokenManager.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerSignalSetAdminIterator is returned from FilterSignalSetAdmin and is used to iterate over the raw logs and unpacked data for SignalSetAdmin events raised by the TokenManager contract.
type TokenManagerSignalSetAdminIterator struct {
	Event *TokenManagerSignalSetAdmin // Event containing the contract specifics and raw log

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
func (it *TokenManagerSignalSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerSignalSetAdmin)
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
		it.Event = new(TokenManagerSignalSetAdmin)
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
func (it *TokenManagerSignalSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerSignalSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerSignalSetAdmin represents a SignalSetAdmin event raised by the TokenManager contract.
type TokenManagerSignalSetAdmin struct {
	Target common.Address
	Admin  common.Address
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalSetAdmin is a free log retrieval operation binding the contract event 0x4fc9433645aa0a3670e9185496bbd752209fed7a9696fb8a954a0db30ef927b0.
//
// Solidity: event SignalSetAdmin(address target, address admin, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterSignalSetAdmin(opts *bind.FilterOpts) (*TokenManagerSignalSetAdminIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "SignalSetAdmin")
	if err != nil {
		return nil, err
	}
	return &TokenManagerSignalSetAdminIterator{contract: _TokenManager.contract, event: "SignalSetAdmin", logs: logs, sub: sub}, nil
}

// WatchSignalSetAdmin is a free log subscription operation binding the contract event 0x4fc9433645aa0a3670e9185496bbd752209fed7a9696fb8a954a0db30ef927b0.
//
// Solidity: event SignalSetAdmin(address target, address admin, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchSignalSetAdmin(opts *bind.WatchOpts, sink chan<- *TokenManagerSignalSetAdmin) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "SignalSetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerSignalSetAdmin)
				if err := _TokenManager.contract.UnpackLog(event, "SignalSetAdmin", log); err != nil {
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

// ParseSignalSetAdmin is a log parse operation binding the contract event 0x4fc9433645aa0a3670e9185496bbd752209fed7a9696fb8a954a0db30ef927b0.
//
// Solidity: event SignalSetAdmin(address target, address admin, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseSignalSetAdmin(log types.Log) (*TokenManagerSignalSetAdmin, error) {
	event := new(TokenManagerSignalSetAdmin)
	if err := _TokenManager.contract.UnpackLog(event, "SignalSetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenManagerSignalSetGovIterator is returned from FilterSignalSetGov and is used to iterate over the raw logs and unpacked data for SignalSetGov events raised by the TokenManager contract.
type TokenManagerSignalSetGovIterator struct {
	Event *TokenManagerSignalSetGov // Event containing the contract specifics and raw log

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
func (it *TokenManagerSignalSetGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerSignalSetGov)
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
		it.Event = new(TokenManagerSignalSetGov)
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
func (it *TokenManagerSignalSetGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerSignalSetGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerSignalSetGov represents a SignalSetGov event raised by the TokenManager contract.
type TokenManagerSignalSetGov struct {
	Timelock common.Address
	Target   common.Address
	Gov      common.Address
	Action   [32]byte
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalSetGov is a free log retrieval operation binding the contract event 0x634e13057d45400506e3b303913ac59b61e5a8137ea6fed5ed44aa0b8bc3c568.
//
// Solidity: event SignalSetGov(address timelock, address target, address gov, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) FilterSignalSetGov(opts *bind.FilterOpts) (*TokenManagerSignalSetGovIterator, error) {

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return &TokenManagerSignalSetGovIterator{contract: _TokenManager.contract, event: "SignalSetGov", logs: logs, sub: sub}, nil
}

// WatchSignalSetGov is a free log subscription operation binding the contract event 0x634e13057d45400506e3b303913ac59b61e5a8137ea6fed5ed44aa0b8bc3c568.
//
// Solidity: event SignalSetGov(address timelock, address target, address gov, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) WatchSignalSetGov(opts *bind.WatchOpts, sink chan<- *TokenManagerSignalSetGov) (event.Subscription, error) {

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerSignalSetGov)
				if err := _TokenManager.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
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

// ParseSignalSetGov is a log parse operation binding the contract event 0x634e13057d45400506e3b303913ac59b61e5a8137ea6fed5ed44aa0b8bc3c568.
//
// Solidity: event SignalSetGov(address timelock, address target, address gov, bytes32 action, uint256 nonce)
func (_TokenManager *TokenManagerFilterer) ParseSignalSetGov(log types.Log) (*TokenManagerSignalSetGov, error) {
	event := new(TokenManagerSignalSetGov)
	if err := _TokenManager.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
