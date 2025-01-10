// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmxfloor

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

// GmxFloorMetaData contains all meta data concerning the GmxFloor contract.
var GmxFloorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_backedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseMintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_multiplierPrecision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAuthorizations\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"ClearAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalApproveNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalApproveNFTs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalPendingAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalSetAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"SignalSetGov\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionsNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"backedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getBurnAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAuthorizations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierPrecision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingActions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"receiveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_backedSupply\",\"type\":\"uint256\"}],\"name\":\"setBackedSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isHandler\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintMultiplier\",\"type\":\"uint256\"}],\"name\":\"setMintMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signApproveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signApproveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signSetAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"signSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"signalApproveNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"signalApproveNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"signalSetAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"signalSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedActions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GmxFloorABI is the input ABI used to generate the binding from.
// Deprecated: Use GmxFloorMetaData.ABI instead.
var GmxFloorABI = GmxFloorMetaData.ABI

// GmxFloor is an auto generated Go binding around an Ethereum contract.
type GmxFloor struct {
	GmxFloorCaller     // Read-only binding to the contract
	GmxFloorTransactor // Write-only binding to the contract
	GmxFloorFilterer   // Log filterer for contract events
}

// GmxFloorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmxFloorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxFloorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmxFloorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxFloorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmxFloorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxFloorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmxFloorSession struct {
	Contract     *GmxFloor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmxFloorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmxFloorCallerSession struct {
	Contract *GmxFloorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GmxFloorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmxFloorTransactorSession struct {
	Contract     *GmxFloorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GmxFloorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmxFloorRaw struct {
	Contract *GmxFloor // Generic contract binding to access the raw methods on
}

// GmxFloorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmxFloorCallerRaw struct {
	Contract *GmxFloorCaller // Generic read-only contract binding to access the raw methods on
}

// GmxFloorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmxFloorTransactorRaw struct {
	Contract *GmxFloorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmxFloor creates a new instance of GmxFloor, bound to a specific deployed contract.
func NewGmxFloor(address common.Address, backend bind.ContractBackend) (*GmxFloor, error) {
	contract, err := bindGmxFloor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GmxFloor{GmxFloorCaller: GmxFloorCaller{contract: contract}, GmxFloorTransactor: GmxFloorTransactor{contract: contract}, GmxFloorFilterer: GmxFloorFilterer{contract: contract}}, nil
}

// NewGmxFloorCaller creates a new read-only instance of GmxFloor, bound to a specific deployed contract.
func NewGmxFloorCaller(address common.Address, caller bind.ContractCaller) (*GmxFloorCaller, error) {
	contract, err := bindGmxFloor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmxFloorCaller{contract: contract}, nil
}

// NewGmxFloorTransactor creates a new write-only instance of GmxFloor, bound to a specific deployed contract.
func NewGmxFloorTransactor(address common.Address, transactor bind.ContractTransactor) (*GmxFloorTransactor, error) {
	contract, err := bindGmxFloor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmxFloorTransactor{contract: contract}, nil
}

// NewGmxFloorFilterer creates a new log filterer instance of GmxFloor, bound to a specific deployed contract.
func NewGmxFloorFilterer(address common.Address, filterer bind.ContractFilterer) (*GmxFloorFilterer, error) {
	contract, err := bindGmxFloor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmxFloorFilterer{contract: contract}, nil
}

// bindGmxFloor binds a generic wrapper to an already deployed contract.
func bindGmxFloor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GmxFloorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxFloor *GmxFloorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxFloor.Contract.GmxFloorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxFloor *GmxFloorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxFloor.Contract.GmxFloorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxFloor *GmxFloorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxFloor.Contract.GmxFloorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxFloor *GmxFloorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxFloor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxFloor *GmxFloorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxFloor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxFloor *GmxFloorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxFloor.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GmxFloor *GmxFloorSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _GmxFloor.Contract.BASISPOINTSDIVISOR(&_GmxFloor.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _GmxFloor.Contract.BASISPOINTSDIVISOR(&_GmxFloor.CallOpts)
}

// BURNBASISPOINTS is a free data retrieval call binding the contract method 0x5fe15ad7.
//
// Solidity: function BURN_BASIS_POINTS() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) BURNBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "BURN_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BURNBASISPOINTS is a free data retrieval call binding the contract method 0x5fe15ad7.
//
// Solidity: function BURN_BASIS_POINTS() view returns(uint256)
func (_GmxFloor *GmxFloorSession) BURNBASISPOINTS() (*big.Int, error) {
	return _GmxFloor.Contract.BURNBASISPOINTS(&_GmxFloor.CallOpts)
}

// BURNBASISPOINTS is a free data retrieval call binding the contract method 0x5fe15ad7.
//
// Solidity: function BURN_BASIS_POINTS() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) BURNBASISPOINTS() (*big.Int, error) {
	return _GmxFloor.Contract.BURNBASISPOINTS(&_GmxFloor.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GmxFloor *GmxFloorSession) PRICEPRECISION() (*big.Int, error) {
	return _GmxFloor.Contract.PRICEPRECISION(&_GmxFloor.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _GmxFloor.Contract.PRICEPRECISION(&_GmxFloor.CallOpts)
}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) ActionsNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "actionsNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_GmxFloor *GmxFloorSession) ActionsNonce() (*big.Int, error) {
	return _GmxFloor.Contract.ActionsNonce(&_GmxFloor.CallOpts)
}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) ActionsNonce() (*big.Int, error) {
	return _GmxFloor.Contract.ActionsNonce(&_GmxFloor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxFloor *GmxFloorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxFloor *GmxFloorSession) Admin() (common.Address, error) {
	return _GmxFloor.Contract.Admin(&_GmxFloor.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxFloor *GmxFloorCallerSession) Admin() (common.Address, error) {
	return _GmxFloor.Contract.Admin(&_GmxFloor.CallOpts)
}

// BackedSupply is a free data retrieval call binding the contract method 0xc616b592.
//
// Solidity: function backedSupply() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) BackedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "backedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BackedSupply is a free data retrieval call binding the contract method 0xc616b592.
//
// Solidity: function backedSupply() view returns(uint256)
func (_GmxFloor *GmxFloorSession) BackedSupply() (*big.Int, error) {
	return _GmxFloor.Contract.BackedSupply(&_GmxFloor.CallOpts)
}

// BackedSupply is a free data retrieval call binding the contract method 0xc616b592.
//
// Solidity: function backedSupply() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) BackedSupply() (*big.Int, error) {
	return _GmxFloor.Contract.BackedSupply(&_GmxFloor.CallOpts)
}

// BaseMintPrice is a free data retrieval call binding the contract method 0x8559dff3.
//
// Solidity: function baseMintPrice() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) BaseMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "baseMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseMintPrice is a free data retrieval call binding the contract method 0x8559dff3.
//
// Solidity: function baseMintPrice() view returns(uint256)
func (_GmxFloor *GmxFloorSession) BaseMintPrice() (*big.Int, error) {
	return _GmxFloor.Contract.BaseMintPrice(&_GmxFloor.CallOpts)
}

// BaseMintPrice is a free data retrieval call binding the contract method 0x8559dff3.
//
// Solidity: function baseMintPrice() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) BaseMintPrice() (*big.Int, error) {
	return _GmxFloor.Contract.BaseMintPrice(&_GmxFloor.CallOpts)
}

// GetBurnAmountOut is a free data retrieval call binding the contract method 0xb82ce261.
//
// Solidity: function getBurnAmountOut(uint256 _amount) view returns(uint256)
func (_GmxFloor *GmxFloorCaller) GetBurnAmountOut(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "getBurnAmountOut", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurnAmountOut is a free data retrieval call binding the contract method 0xb82ce261.
//
// Solidity: function getBurnAmountOut(uint256 _amount) view returns(uint256)
func (_GmxFloor *GmxFloorSession) GetBurnAmountOut(_amount *big.Int) (*big.Int, error) {
	return _GmxFloor.Contract.GetBurnAmountOut(&_GmxFloor.CallOpts, _amount)
}

// GetBurnAmountOut is a free data retrieval call binding the contract method 0xb82ce261.
//
// Solidity: function getBurnAmountOut(uint256 _amount) view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) GetBurnAmountOut(_amount *big.Int) (*big.Int, error) {
	return _GmxFloor.Contract.GetBurnAmountOut(&_GmxFloor.CallOpts, _amount)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) GetMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "getMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_GmxFloor *GmxFloorSession) GetMintPrice() (*big.Int, error) {
	return _GmxFloor.Contract.GetMintPrice(&_GmxFloor.CallOpts)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) GetMintPrice() (*big.Int, error) {
	return _GmxFloor.Contract.GetMintPrice(&_GmxFloor.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_GmxFloor *GmxFloorCaller) Gmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "gmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_GmxFloor *GmxFloorSession) Gmx() (common.Address, error) {
	return _GmxFloor.Contract.Gmx(&_GmxFloor.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_GmxFloor *GmxFloorCallerSession) Gmx() (common.Address, error) {
	return _GmxFloor.Contract.Gmx(&_GmxFloor.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxFloor *GmxFloorCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxFloor *GmxFloorSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxFloor.Contract.IsHandler(&_GmxFloor.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxFloor *GmxFloorCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxFloor.Contract.IsHandler(&_GmxFloor.CallOpts, arg0)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxFloor *GmxFloorCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxFloor *GmxFloorSession) IsInitialized() (bool, error) {
	return _GmxFloor.Contract.IsInitialized(&_GmxFloor.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxFloor *GmxFloorCallerSession) IsInitialized() (bool, error) {
	return _GmxFloor.Contract.IsInitialized(&_GmxFloor.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_GmxFloor *GmxFloorCaller) IsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "isSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_GmxFloor *GmxFloorSession) IsSigner(arg0 common.Address) (bool, error) {
	return _GmxFloor.Contract.IsSigner(&_GmxFloor.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_GmxFloor *GmxFloorCallerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _GmxFloor.Contract.IsSigner(&_GmxFloor.CallOpts, arg0)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) MinAuthorizations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "minAuthorizations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_GmxFloor *GmxFloorSession) MinAuthorizations() (*big.Int, error) {
	return _GmxFloor.Contract.MinAuthorizations(&_GmxFloor.CallOpts)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) MinAuthorizations() (*big.Int, error) {
	return _GmxFloor.Contract.MinAuthorizations(&_GmxFloor.CallOpts)
}

// MintMultiplier is a free data retrieval call binding the contract method 0x3acd81a5.
//
// Solidity: function mintMultiplier() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) MintMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "mintMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintMultiplier is a free data retrieval call binding the contract method 0x3acd81a5.
//
// Solidity: function mintMultiplier() view returns(uint256)
func (_GmxFloor *GmxFloorSession) MintMultiplier() (*big.Int, error) {
	return _GmxFloor.Contract.MintMultiplier(&_GmxFloor.CallOpts)
}

// MintMultiplier is a free data retrieval call binding the contract method 0x3acd81a5.
//
// Solidity: function mintMultiplier() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) MintMultiplier() (*big.Int, error) {
	return _GmxFloor.Contract.MintMultiplier(&_GmxFloor.CallOpts)
}

// MintedSupply is a free data retrieval call binding the contract method 0xc1bd8cf9.
//
// Solidity: function mintedSupply() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) MintedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "mintedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedSupply is a free data retrieval call binding the contract method 0xc1bd8cf9.
//
// Solidity: function mintedSupply() view returns(uint256)
func (_GmxFloor *GmxFloorSession) MintedSupply() (*big.Int, error) {
	return _GmxFloor.Contract.MintedSupply(&_GmxFloor.CallOpts)
}

// MintedSupply is a free data retrieval call binding the contract method 0xc1bd8cf9.
//
// Solidity: function mintedSupply() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) MintedSupply() (*big.Int, error) {
	return _GmxFloor.Contract.MintedSupply(&_GmxFloor.CallOpts)
}

// MultiplierPrecision is a free data retrieval call binding the contract method 0x2415ab6e.
//
// Solidity: function multiplierPrecision() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) MultiplierPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "multiplierPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierPrecision is a free data retrieval call binding the contract method 0x2415ab6e.
//
// Solidity: function multiplierPrecision() view returns(uint256)
func (_GmxFloor *GmxFloorSession) MultiplierPrecision() (*big.Int, error) {
	return _GmxFloor.Contract.MultiplierPrecision(&_GmxFloor.CallOpts)
}

// MultiplierPrecision is a free data retrieval call binding the contract method 0x2415ab6e.
//
// Solidity: function multiplierPrecision() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) MultiplierPrecision() (*big.Int, error) {
	return _GmxFloor.Contract.MultiplierPrecision(&_GmxFloor.CallOpts)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_GmxFloor *GmxFloorCaller) PendingActions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "pendingActions", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_GmxFloor *GmxFloorSession) PendingActions(arg0 [32]byte) (bool, error) {
	return _GmxFloor.Contract.PendingActions(&_GmxFloor.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_GmxFloor *GmxFloorCallerSession) PendingActions(arg0 [32]byte) (bool, error) {
	return _GmxFloor.Contract.PendingActions(&_GmxFloor.CallOpts, arg0)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_GmxFloor *GmxFloorCaller) ReserveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "reserveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_GmxFloor *GmxFloorSession) ReserveToken() (common.Address, error) {
	return _GmxFloor.Contract.ReserveToken(&_GmxFloor.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_GmxFloor *GmxFloorCallerSession) ReserveToken() (common.Address, error) {
	return _GmxFloor.Contract.ReserveToken(&_GmxFloor.CallOpts)
}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_GmxFloor *GmxFloorCaller) SignedActions(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "signedActions", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_GmxFloor *GmxFloorSession) SignedActions(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _GmxFloor.Contract.SignedActions(&_GmxFloor.CallOpts, arg0, arg1)
}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_GmxFloor *GmxFloorCallerSession) SignedActions(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _GmxFloor.Contract.SignedActions(&_GmxFloor.CallOpts, arg0, arg1)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_GmxFloor *GmxFloorCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_GmxFloor *GmxFloorSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _GmxFloor.Contract.Signers(&_GmxFloor.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_GmxFloor *GmxFloorCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _GmxFloor.Contract.Signers(&_GmxFloor.CallOpts, arg0)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() view returns(uint256)
func (_GmxFloor *GmxFloorCaller) SignersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxFloor.contract.Call(opts, &out, "signersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() view returns(uint256)
func (_GmxFloor *GmxFloorSession) SignersLength() (*big.Int, error) {
	return _GmxFloor.Contract.SignersLength(&_GmxFloor.CallOpts)
}

// SignersLength is a free data retrieval call binding the contract method 0x41f684f3.
//
// Solidity: function signersLength() view returns(uint256)
func (_GmxFloor *GmxFloorCallerSession) SignersLength() (*big.Int, error) {
	return _GmxFloor.Contract.SignersLength(&_GmxFloor.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "approve", _token, _spender, _amount, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.Approve(&_GmxFloor.TransactOpts, _token, _spender, _amount, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.Approve(&_GmxFloor.TransactOpts, _token, _spender, _amount, _nonce)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xf23f9775.
//
// Solidity: function approveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) ApproveNFT(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "approveNFT", _token, _spender, _tokenId, _nonce)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xf23f9775.
//
// Solidity: function approveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) ApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.ApproveNFT(&_GmxFloor.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// ApproveNFT is a paid mutator transaction binding the contract method 0xf23f9775.
//
// Solidity: function approveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) ApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.ApproveNFT(&_GmxFloor.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// ApproveNFTs is a paid mutator transaction binding the contract method 0x42a1fcee.
//
// Solidity: function approveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) ApproveNFTs(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "approveNFTs", _token, _spender, _tokenIds, _nonce)
}

// ApproveNFTs is a paid mutator transaction binding the contract method 0x42a1fcee.
//
// Solidity: function approveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) ApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.ApproveNFTs(&_GmxFloor.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// ApproveNFTs is a paid mutator transaction binding the contract method 0x42a1fcee.
//
// Solidity: function approveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) ApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.ApproveNFTs(&_GmxFloor.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// Burn is a paid mutator transaction binding the contract method 0x749388c4.
//
// Solidity: function burn(uint256 _amount, uint256 _minOut, address _receiver) returns(uint256)
func (_GmxFloor *GmxFloorTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "burn", _amount, _minOut, _receiver)
}

// Burn is a paid mutator transaction binding the contract method 0x749388c4.
//
// Solidity: function burn(uint256 _amount, uint256 _minOut, address _receiver) returns(uint256)
func (_GmxFloor *GmxFloorSession) Burn(_amount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.Burn(&_GmxFloor.TransactOpts, _amount, _minOut, _receiver)
}

// Burn is a paid mutator transaction binding the contract method 0x749388c4.
//
// Solidity: function burn(uint256 _amount, uint256 _minOut, address _receiver) returns(uint256)
func (_GmxFloor *GmxFloorTransactorSession) Burn(_amount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.Burn(&_GmxFloor.TransactOpts, _amount, _minOut, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _signers) returns()
func (_GmxFloor *GmxFloorTransactor) Initialize(opts *bind.TransactOpts, _signers []common.Address) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "initialize", _signers)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _signers) returns()
func (_GmxFloor *GmxFloorSession) Initialize(_signers []common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.Initialize(&_GmxFloor.TransactOpts, _signers)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _signers) returns()
func (_GmxFloor *GmxFloorTransactorSession) Initialize(_signers []common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.Initialize(&_GmxFloor.TransactOpts, _signers)
}

// Mint is a paid mutator transaction binding the contract method 0xe7d3fe6b.
//
// Solidity: function mint(uint256 _amount, uint256 _maxCost, address _receiver) returns(uint256)
func (_GmxFloor *GmxFloorTransactor) Mint(opts *bind.TransactOpts, _amount *big.Int, _maxCost *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "mint", _amount, _maxCost, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0xe7d3fe6b.
//
// Solidity: function mint(uint256 _amount, uint256 _maxCost, address _receiver) returns(uint256)
func (_GmxFloor *GmxFloorSession) Mint(_amount *big.Int, _maxCost *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.Mint(&_GmxFloor.TransactOpts, _amount, _maxCost, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0xe7d3fe6b.
//
// Solidity: function mint(uint256 _amount, uint256 _maxCost, address _receiver) returns(uint256)
func (_GmxFloor *GmxFloorTransactorSession) Mint(_amount *big.Int, _maxCost *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.Mint(&_GmxFloor.TransactOpts, _amount, _maxCost, _receiver)
}

// ReceiveNFTs is a paid mutator transaction binding the contract method 0xadb384b6.
//
// Solidity: function receiveNFTs(address _token, address _sender, uint256[] _tokenIds) returns()
func (_GmxFloor *GmxFloorTransactor) ReceiveNFTs(opts *bind.TransactOpts, _token common.Address, _sender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "receiveNFTs", _token, _sender, _tokenIds)
}

// ReceiveNFTs is a paid mutator transaction binding the contract method 0xadb384b6.
//
// Solidity: function receiveNFTs(address _token, address _sender, uint256[] _tokenIds) returns()
func (_GmxFloor *GmxFloorSession) ReceiveNFTs(_token common.Address, _sender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.ReceiveNFTs(&_GmxFloor.TransactOpts, _token, _sender, _tokenIds)
}

// ReceiveNFTs is a paid mutator transaction binding the contract method 0xadb384b6.
//
// Solidity: function receiveNFTs(address _token, address _sender, uint256[] _tokenIds) returns()
func (_GmxFloor *GmxFloorTransactorSession) ReceiveNFTs(_token common.Address, _sender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.ReceiveNFTs(&_GmxFloor.TransactOpts, _token, _sender, _tokenIds)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x9fddaac1.
//
// Solidity: function setAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) SetAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "setAdmin", _target, _admin, _nonce)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x9fddaac1.
//
// Solidity: function setAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) SetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetAdmin(&_GmxFloor.TransactOpts, _target, _admin, _nonce)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x9fddaac1.
//
// Solidity: function setAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) SetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetAdmin(&_GmxFloor.TransactOpts, _target, _admin, _nonce)
}

// SetBackedSupply is a paid mutator transaction binding the contract method 0x2494b7eb.
//
// Solidity: function setBackedSupply(uint256 _backedSupply) returns()
func (_GmxFloor *GmxFloorTransactor) SetBackedSupply(opts *bind.TransactOpts, _backedSupply *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "setBackedSupply", _backedSupply)
}

// SetBackedSupply is a paid mutator transaction binding the contract method 0x2494b7eb.
//
// Solidity: function setBackedSupply(uint256 _backedSupply) returns()
func (_GmxFloor *GmxFloorSession) SetBackedSupply(_backedSupply *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetBackedSupply(&_GmxFloor.TransactOpts, _backedSupply)
}

// SetBackedSupply is a paid mutator transaction binding the contract method 0x2494b7eb.
//
// Solidity: function setBackedSupply(uint256 _backedSupply) returns()
func (_GmxFloor *GmxFloorTransactorSession) SetBackedSupply(_backedSupply *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetBackedSupply(&_GmxFloor.TransactOpts, _backedSupply)
}

// SetGov is a paid mutator transaction binding the contract method 0xe9075621.
//
// Solidity: function setGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) SetGov(opts *bind.TransactOpts, _timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "setGov", _timelock, _target, _gov, _nonce)
}

// SetGov is a paid mutator transaction binding the contract method 0xe9075621.
//
// Solidity: function setGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) SetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetGov(&_GmxFloor.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SetGov is a paid mutator transaction binding the contract method 0xe9075621.
//
// Solidity: function setGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) SetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetGov(&_GmxFloor.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isHandler) returns()
func (_GmxFloor *GmxFloorTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isHandler bool) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "setHandler", _handler, _isHandler)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isHandler) returns()
func (_GmxFloor *GmxFloorSession) SetHandler(_handler common.Address, _isHandler bool) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetHandler(&_GmxFloor.TransactOpts, _handler, _isHandler)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isHandler) returns()
func (_GmxFloor *GmxFloorTransactorSession) SetHandler(_handler common.Address, _isHandler bool) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetHandler(&_GmxFloor.TransactOpts, _handler, _isHandler)
}

// SetMintMultiplier is a paid mutator transaction binding the contract method 0xeb0eec48.
//
// Solidity: function setMintMultiplier(uint256 _mintMultiplier) returns()
func (_GmxFloor *GmxFloorTransactor) SetMintMultiplier(opts *bind.TransactOpts, _mintMultiplier *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "setMintMultiplier", _mintMultiplier)
}

// SetMintMultiplier is a paid mutator transaction binding the contract method 0xeb0eec48.
//
// Solidity: function setMintMultiplier(uint256 _mintMultiplier) returns()
func (_GmxFloor *GmxFloorSession) SetMintMultiplier(_mintMultiplier *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetMintMultiplier(&_GmxFloor.TransactOpts, _mintMultiplier)
}

// SetMintMultiplier is a paid mutator transaction binding the contract method 0xeb0eec48.
//
// Solidity: function setMintMultiplier(uint256 _mintMultiplier) returns()
func (_GmxFloor *GmxFloorTransactorSession) SetMintMultiplier(_mintMultiplier *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SetMintMultiplier(&_GmxFloor.TransactOpts, _mintMultiplier)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) SignApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signApprove", _token, _spender, _amount, _nonce)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) SignApprove(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignApprove(&_GmxFloor.TransactOpts, _token, _spender, _amount, _nonce)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignApprove(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignApprove(&_GmxFloor.TransactOpts, _token, _spender, _amount, _nonce)
}

// SignApproveNFT is a paid mutator transaction binding the contract method 0xddf67a9f.
//
// Solidity: function signApproveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) SignApproveNFT(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signApproveNFT", _token, _spender, _tokenId, _nonce)
}

// SignApproveNFT is a paid mutator transaction binding the contract method 0xddf67a9f.
//
// Solidity: function signApproveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) SignApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignApproveNFT(&_GmxFloor.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// SignApproveNFT is a paid mutator transaction binding the contract method 0xddf67a9f.
//
// Solidity: function signApproveNFT(address _token, address _spender, uint256 _tokenId, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignApproveNFT(&_GmxFloor.TransactOpts, _token, _spender, _tokenId, _nonce)
}

// SignApproveNFTs is a paid mutator transaction binding the contract method 0xf00cb942.
//
// Solidity: function signApproveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) SignApproveNFTs(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signApproveNFTs", _token, _spender, _tokenIds, _nonce)
}

// SignApproveNFTs is a paid mutator transaction binding the contract method 0xf00cb942.
//
// Solidity: function signApproveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) SignApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignApproveNFTs(&_GmxFloor.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// SignApproveNFTs is a paid mutator transaction binding the contract method 0xf00cb942.
//
// Solidity: function signApproveNFTs(address _token, address _spender, uint256[] _tokenIds, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignApproveNFTs(&_GmxFloor.TransactOpts, _token, _spender, _tokenIds, _nonce)
}

// SignSetAdmin is a paid mutator transaction binding the contract method 0xf466634b.
//
// Solidity: function signSetAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) SignSetAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signSetAdmin", _target, _admin, _nonce)
}

// SignSetAdmin is a paid mutator transaction binding the contract method 0xf466634b.
//
// Solidity: function signSetAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) SignSetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignSetAdmin(&_GmxFloor.TransactOpts, _target, _admin, _nonce)
}

// SignSetAdmin is a paid mutator transaction binding the contract method 0xf466634b.
//
// Solidity: function signSetAdmin(address _target, address _admin, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignSetAdmin(_target common.Address, _admin common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignSetAdmin(&_GmxFloor.TransactOpts, _target, _admin, _nonce)
}

// SignSetGov is a paid mutator transaction binding the contract method 0xa62fb170.
//
// Solidity: function signSetGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactor) SignSetGov(opts *bind.TransactOpts, _timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signSetGov", _timelock, _target, _gov, _nonce)
}

// SignSetGov is a paid mutator transaction binding the contract method 0xa62fb170.
//
// Solidity: function signSetGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorSession) SignSetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignSetGov(&_GmxFloor.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SignSetGov is a paid mutator transaction binding the contract method 0xa62fb170.
//
// Solidity: function signSetGov(address _timelock, address _target, address _gov, uint256 _nonce) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignSetGov(_timelock common.Address, _target common.Address, _gov common.Address, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignSetGov(&_GmxFloor.TransactOpts, _timelock, _target, _gov, _nonce)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxFloor *GmxFloorTransactor) SignalApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signalApprove", _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxFloor *GmxFloorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalApprove(&_GmxFloor.TransactOpts, _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalApprove(&_GmxFloor.TransactOpts, _token, _spender, _amount)
}

// SignalApproveNFT is a paid mutator transaction binding the contract method 0x0b13beca.
//
// Solidity: function signalApproveNFT(address _token, address _spender, uint256 _tokenId) returns()
func (_GmxFloor *GmxFloorTransactor) SignalApproveNFT(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signalApproveNFT", _token, _spender, _tokenId)
}

// SignalApproveNFT is a paid mutator transaction binding the contract method 0x0b13beca.
//
// Solidity: function signalApproveNFT(address _token, address _spender, uint256 _tokenId) returns()
func (_GmxFloor *GmxFloorSession) SignalApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalApproveNFT(&_GmxFloor.TransactOpts, _token, _spender, _tokenId)
}

// SignalApproveNFT is a paid mutator transaction binding the contract method 0x0b13beca.
//
// Solidity: function signalApproveNFT(address _token, address _spender, uint256 _tokenId) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignalApproveNFT(_token common.Address, _spender common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalApproveNFT(&_GmxFloor.TransactOpts, _token, _spender, _tokenId)
}

// SignalApproveNFTs is a paid mutator transaction binding the contract method 0x181fcd33.
//
// Solidity: function signalApproveNFTs(address _token, address _spender, uint256[] _tokenIds) returns()
func (_GmxFloor *GmxFloorTransactor) SignalApproveNFTs(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signalApproveNFTs", _token, _spender, _tokenIds)
}

// SignalApproveNFTs is a paid mutator transaction binding the contract method 0x181fcd33.
//
// Solidity: function signalApproveNFTs(address _token, address _spender, uint256[] _tokenIds) returns()
func (_GmxFloor *GmxFloorSession) SignalApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalApproveNFTs(&_GmxFloor.TransactOpts, _token, _spender, _tokenIds)
}

// SignalApproveNFTs is a paid mutator transaction binding the contract method 0x181fcd33.
//
// Solidity: function signalApproveNFTs(address _token, address _spender, uint256[] _tokenIds) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignalApproveNFTs(_token common.Address, _spender common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalApproveNFTs(&_GmxFloor.TransactOpts, _token, _spender, _tokenIds)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x75fd490c.
//
// Solidity: function signalSetAdmin(address _target, address _admin) returns()
func (_GmxFloor *GmxFloorTransactor) SignalSetAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signalSetAdmin", _target, _admin)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x75fd490c.
//
// Solidity: function signalSetAdmin(address _target, address _admin) returns()
func (_GmxFloor *GmxFloorSession) SignalSetAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalSetAdmin(&_GmxFloor.TransactOpts, _target, _admin)
}

// SignalSetAdmin is a paid mutator transaction binding the contract method 0x75fd490c.
//
// Solidity: function signalSetAdmin(address _target, address _admin) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignalSetAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalSetAdmin(&_GmxFloor.TransactOpts, _target, _admin)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0xe2d03cc5.
//
// Solidity: function signalSetGov(address _timelock, address _target, address _gov) returns()
func (_GmxFloor *GmxFloorTransactor) SignalSetGov(opts *bind.TransactOpts, _timelock common.Address, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxFloor.contract.Transact(opts, "signalSetGov", _timelock, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0xe2d03cc5.
//
// Solidity: function signalSetGov(address _timelock, address _target, address _gov) returns()
func (_GmxFloor *GmxFloorSession) SignalSetGov(_timelock common.Address, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalSetGov(&_GmxFloor.TransactOpts, _timelock, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0xe2d03cc5.
//
// Solidity: function signalSetGov(address _timelock, address _target, address _gov) returns()
func (_GmxFloor *GmxFloorTransactorSession) SignalSetGov(_timelock common.Address, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxFloor.Contract.SignalSetGov(&_GmxFloor.TransactOpts, _timelock, _target, _gov)
}

// GmxFloorClearActionIterator is returned from FilterClearAction and is used to iterate over the raw logs and unpacked data for ClearAction events raised by the GmxFloor contract.
type GmxFloorClearActionIterator struct {
	Event *GmxFloorClearAction // Event containing the contract specifics and raw log

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
func (it *GmxFloorClearActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorClearAction)
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
		it.Event = new(GmxFloorClearAction)
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
func (it *GmxFloorClearActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorClearActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorClearAction represents a ClearAction event raised by the GmxFloor contract.
type GmxFloorClearAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClearAction is a free log retrieval operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) FilterClearAction(opts *bind.FilterOpts) (*GmxFloorClearActionIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return &GmxFloorClearActionIterator{contract: _GmxFloor.contract, event: "ClearAction", logs: logs, sub: sub}, nil
}

// WatchClearAction is a free log subscription operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchClearAction(opts *bind.WatchOpts, sink chan<- *GmxFloorClearAction) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorClearAction)
				if err := _GmxFloor.contract.UnpackLog(event, "ClearAction", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseClearAction(log types.Log) (*GmxFloorClearAction, error) {
	event := new(GmxFloorClearAction)
	if err := _GmxFloor.contract.UnpackLog(event, "ClearAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxFloorSignActionIterator is returned from FilterSignAction and is used to iterate over the raw logs and unpacked data for SignAction events raised by the GmxFloor contract.
type GmxFloorSignActionIterator struct {
	Event *GmxFloorSignAction // Event containing the contract specifics and raw log

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
func (it *GmxFloorSignActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorSignAction)
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
		it.Event = new(GmxFloorSignAction)
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
func (it *GmxFloorSignActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorSignActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorSignAction represents a SignAction event raised by the GmxFloor contract.
type GmxFloorSignAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignAction is a free log retrieval operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) FilterSignAction(opts *bind.FilterOpts) (*GmxFloorSignActionIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "SignAction")
	if err != nil {
		return nil, err
	}
	return &GmxFloorSignActionIterator{contract: _GmxFloor.contract, event: "SignAction", logs: logs, sub: sub}, nil
}

// WatchSignAction is a free log subscription operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchSignAction(opts *bind.WatchOpts, sink chan<- *GmxFloorSignAction) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "SignAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorSignAction)
				if err := _GmxFloor.contract.UnpackLog(event, "SignAction", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseSignAction(log types.Log) (*GmxFloorSignAction, error) {
	event := new(GmxFloorSignAction)
	if err := _GmxFloor.contract.UnpackLog(event, "SignAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxFloorSignalApproveIterator is returned from FilterSignalApprove and is used to iterate over the raw logs and unpacked data for SignalApprove events raised by the GmxFloor contract.
type GmxFloorSignalApproveIterator struct {
	Event *GmxFloorSignalApprove // Event containing the contract specifics and raw log

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
func (it *GmxFloorSignalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorSignalApprove)
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
		it.Event = new(GmxFloorSignalApprove)
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
func (it *GmxFloorSignalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorSignalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorSignalApprove represents a SignalApprove event raised by the GmxFloor contract.
type GmxFloorSignalApprove struct {
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
func (_GmxFloor *GmxFloorFilterer) FilterSignalApprove(opts *bind.FilterOpts) (*GmxFloorSignalApproveIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return &GmxFloorSignalApproveIterator{contract: _GmxFloor.contract, event: "SignalApprove", logs: logs, sub: sub}, nil
}

// WatchSignalApprove is a free log subscription operation binding the contract event 0xc19251bf5f704ddc3d5babe6f4e5bde0dded20b19f7844716861821ab3163cd7.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchSignalApprove(opts *bind.WatchOpts, sink chan<- *GmxFloorSignalApprove) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorSignalApprove)
				if err := _GmxFloor.contract.UnpackLog(event, "SignalApprove", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseSignalApprove(log types.Log) (*GmxFloorSignalApprove, error) {
	event := new(GmxFloorSignalApprove)
	if err := _GmxFloor.contract.UnpackLog(event, "SignalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxFloorSignalApproveNFTIterator is returned from FilterSignalApproveNFT and is used to iterate over the raw logs and unpacked data for SignalApproveNFT events raised by the GmxFloor contract.
type GmxFloorSignalApproveNFTIterator struct {
	Event *GmxFloorSignalApproveNFT // Event containing the contract specifics and raw log

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
func (it *GmxFloorSignalApproveNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorSignalApproveNFT)
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
		it.Event = new(GmxFloorSignalApproveNFT)
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
func (it *GmxFloorSignalApproveNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorSignalApproveNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorSignalApproveNFT represents a SignalApproveNFT event raised by the GmxFloor contract.
type GmxFloorSignalApproveNFT struct {
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
func (_GmxFloor *GmxFloorFilterer) FilterSignalApproveNFT(opts *bind.FilterOpts) (*GmxFloorSignalApproveNFTIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "SignalApproveNFT")
	if err != nil {
		return nil, err
	}
	return &GmxFloorSignalApproveNFTIterator{contract: _GmxFloor.contract, event: "SignalApproveNFT", logs: logs, sub: sub}, nil
}

// WatchSignalApproveNFT is a free log subscription operation binding the contract event 0xcd9ba83b63715dc15ac193645d6e925bf4b487c94b73d709b8b6dea608efd4cc.
//
// Solidity: event SignalApproveNFT(address token, address spender, uint256 tokenId, bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchSignalApproveNFT(opts *bind.WatchOpts, sink chan<- *GmxFloorSignalApproveNFT) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "SignalApproveNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorSignalApproveNFT)
				if err := _GmxFloor.contract.UnpackLog(event, "SignalApproveNFT", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseSignalApproveNFT(log types.Log) (*GmxFloorSignalApproveNFT, error) {
	event := new(GmxFloorSignalApproveNFT)
	if err := _GmxFloor.contract.UnpackLog(event, "SignalApproveNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxFloorSignalApproveNFTsIterator is returned from FilterSignalApproveNFTs and is used to iterate over the raw logs and unpacked data for SignalApproveNFTs events raised by the GmxFloor contract.
type GmxFloorSignalApproveNFTsIterator struct {
	Event *GmxFloorSignalApproveNFTs // Event containing the contract specifics and raw log

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
func (it *GmxFloorSignalApproveNFTsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorSignalApproveNFTs)
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
		it.Event = new(GmxFloorSignalApproveNFTs)
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
func (it *GmxFloorSignalApproveNFTsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorSignalApproveNFTsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorSignalApproveNFTs represents a SignalApproveNFTs event raised by the GmxFloor contract.
type GmxFloorSignalApproveNFTs struct {
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
func (_GmxFloor *GmxFloorFilterer) FilterSignalApproveNFTs(opts *bind.FilterOpts) (*GmxFloorSignalApproveNFTsIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "SignalApproveNFTs")
	if err != nil {
		return nil, err
	}
	return &GmxFloorSignalApproveNFTsIterator{contract: _GmxFloor.contract, event: "SignalApproveNFTs", logs: logs, sub: sub}, nil
}

// WatchSignalApproveNFTs is a free log subscription operation binding the contract event 0xf9d0354d71c261982d98abd09b735f3663b2d7275e2569ad5fd907a4092765f9.
//
// Solidity: event SignalApproveNFTs(address token, address spender, uint256[] tokenIds, bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchSignalApproveNFTs(opts *bind.WatchOpts, sink chan<- *GmxFloorSignalApproveNFTs) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "SignalApproveNFTs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorSignalApproveNFTs)
				if err := _GmxFloor.contract.UnpackLog(event, "SignalApproveNFTs", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseSignalApproveNFTs(log types.Log) (*GmxFloorSignalApproveNFTs, error) {
	event := new(GmxFloorSignalApproveNFTs)
	if err := _GmxFloor.contract.UnpackLog(event, "SignalApproveNFTs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxFloorSignalPendingActionIterator is returned from FilterSignalPendingAction and is used to iterate over the raw logs and unpacked data for SignalPendingAction events raised by the GmxFloor contract.
type GmxFloorSignalPendingActionIterator struct {
	Event *GmxFloorSignalPendingAction // Event containing the contract specifics and raw log

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
func (it *GmxFloorSignalPendingActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorSignalPendingAction)
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
		it.Event = new(GmxFloorSignalPendingAction)
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
func (it *GmxFloorSignalPendingActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorSignalPendingActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorSignalPendingAction represents a SignalPendingAction event raised by the GmxFloor contract.
type GmxFloorSignalPendingAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalPendingAction is a free log retrieval operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) FilterSignalPendingAction(opts *bind.FilterOpts) (*GmxFloorSignalPendingActionIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return &GmxFloorSignalPendingActionIterator{contract: _GmxFloor.contract, event: "SignalPendingAction", logs: logs, sub: sub}, nil
}

// WatchSignalPendingAction is a free log subscription operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchSignalPendingAction(opts *bind.WatchOpts, sink chan<- *GmxFloorSignalPendingAction) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorSignalPendingAction)
				if err := _GmxFloor.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseSignalPendingAction(log types.Log) (*GmxFloorSignalPendingAction, error) {
	event := new(GmxFloorSignalPendingAction)
	if err := _GmxFloor.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxFloorSignalSetAdminIterator is returned from FilterSignalSetAdmin and is used to iterate over the raw logs and unpacked data for SignalSetAdmin events raised by the GmxFloor contract.
type GmxFloorSignalSetAdminIterator struct {
	Event *GmxFloorSignalSetAdmin // Event containing the contract specifics and raw log

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
func (it *GmxFloorSignalSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorSignalSetAdmin)
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
		it.Event = new(GmxFloorSignalSetAdmin)
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
func (it *GmxFloorSignalSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorSignalSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorSignalSetAdmin represents a SignalSetAdmin event raised by the GmxFloor contract.
type GmxFloorSignalSetAdmin struct {
	Target common.Address
	Admin  common.Address
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalSetAdmin is a free log retrieval operation binding the contract event 0x4fc9433645aa0a3670e9185496bbd752209fed7a9696fb8a954a0db30ef927b0.
//
// Solidity: event SignalSetAdmin(address target, address admin, bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) FilterSignalSetAdmin(opts *bind.FilterOpts) (*GmxFloorSignalSetAdminIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "SignalSetAdmin")
	if err != nil {
		return nil, err
	}
	return &GmxFloorSignalSetAdminIterator{contract: _GmxFloor.contract, event: "SignalSetAdmin", logs: logs, sub: sub}, nil
}

// WatchSignalSetAdmin is a free log subscription operation binding the contract event 0x4fc9433645aa0a3670e9185496bbd752209fed7a9696fb8a954a0db30ef927b0.
//
// Solidity: event SignalSetAdmin(address target, address admin, bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchSignalSetAdmin(opts *bind.WatchOpts, sink chan<- *GmxFloorSignalSetAdmin) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "SignalSetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorSignalSetAdmin)
				if err := _GmxFloor.contract.UnpackLog(event, "SignalSetAdmin", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseSignalSetAdmin(log types.Log) (*GmxFloorSignalSetAdmin, error) {
	event := new(GmxFloorSignalSetAdmin)
	if err := _GmxFloor.contract.UnpackLog(event, "SignalSetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxFloorSignalSetGovIterator is returned from FilterSignalSetGov and is used to iterate over the raw logs and unpacked data for SignalSetGov events raised by the GmxFloor contract.
type GmxFloorSignalSetGovIterator struct {
	Event *GmxFloorSignalSetGov // Event containing the contract specifics and raw log

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
func (it *GmxFloorSignalSetGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxFloorSignalSetGov)
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
		it.Event = new(GmxFloorSignalSetGov)
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
func (it *GmxFloorSignalSetGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxFloorSignalSetGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxFloorSignalSetGov represents a SignalSetGov event raised by the GmxFloor contract.
type GmxFloorSignalSetGov struct {
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
func (_GmxFloor *GmxFloorFilterer) FilterSignalSetGov(opts *bind.FilterOpts) (*GmxFloorSignalSetGovIterator, error) {

	logs, sub, err := _GmxFloor.contract.FilterLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return &GmxFloorSignalSetGovIterator{contract: _GmxFloor.contract, event: "SignalSetGov", logs: logs, sub: sub}, nil
}

// WatchSignalSetGov is a free log subscription operation binding the contract event 0x634e13057d45400506e3b303913ac59b61e5a8137ea6fed5ed44aa0b8bc3c568.
//
// Solidity: event SignalSetGov(address timelock, address target, address gov, bytes32 action, uint256 nonce)
func (_GmxFloor *GmxFloorFilterer) WatchSignalSetGov(opts *bind.WatchOpts, sink chan<- *GmxFloorSignalSetGov) (event.Subscription, error) {

	logs, sub, err := _GmxFloor.contract.WatchLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxFloorSignalSetGov)
				if err := _GmxFloor.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
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
func (_GmxFloor *GmxFloorFilterer) ParseSignalSetGov(log types.Log) (*GmxFloorSignalSetGov, error) {
	event := new(GmxFloorSignalSetGov)
	if err := _GmxFloor.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
