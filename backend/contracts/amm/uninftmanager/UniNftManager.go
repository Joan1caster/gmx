// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uninftmanager

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

// UniNftManagerCollectParams is an auto generated low-level Go binding around an user-defined struct.
type UniNftManagerCollectParams struct {
	TokenId    *big.Int
	Recipient  common.Address
	Amount0Max *big.Int
	Amount1Max *big.Int
}

// UniNftManagerMetaData contains all meta data concerning the UniNftManager contract.
var UniNftManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Max\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Max\",\"type\":\"uint128\"}],\"internalType\":\"structUniNftManager.CollectParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokensOwed1\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// UniNftManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use UniNftManagerMetaData.ABI instead.
var UniNftManagerABI = UniNftManagerMetaData.ABI

// UniNftManager is an auto generated Go binding around an Ethereum contract.
type UniNftManager struct {
	UniNftManagerCaller     // Read-only binding to the contract
	UniNftManagerTransactor // Write-only binding to the contract
	UniNftManagerFilterer   // Log filterer for contract events
}

// UniNftManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniNftManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniNftManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniNftManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniNftManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniNftManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniNftManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniNftManagerSession struct {
	Contract     *UniNftManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniNftManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniNftManagerCallerSession struct {
	Contract *UniNftManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UniNftManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniNftManagerTransactorSession struct {
	Contract     *UniNftManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniNftManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniNftManagerRaw struct {
	Contract *UniNftManager // Generic contract binding to access the raw methods on
}

// UniNftManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniNftManagerCallerRaw struct {
	Contract *UniNftManagerCaller // Generic read-only contract binding to access the raw methods on
}

// UniNftManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniNftManagerTransactorRaw struct {
	Contract *UniNftManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniNftManager creates a new instance of UniNftManager, bound to a specific deployed contract.
func NewUniNftManager(address common.Address, backend bind.ContractBackend) (*UniNftManager, error) {
	contract, err := bindUniNftManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniNftManager{UniNftManagerCaller: UniNftManagerCaller{contract: contract}, UniNftManagerTransactor: UniNftManagerTransactor{contract: contract}, UniNftManagerFilterer: UniNftManagerFilterer{contract: contract}}, nil
}

// NewUniNftManagerCaller creates a new read-only instance of UniNftManager, bound to a specific deployed contract.
func NewUniNftManagerCaller(address common.Address, caller bind.ContractCaller) (*UniNftManagerCaller, error) {
	contract, err := bindUniNftManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniNftManagerCaller{contract: contract}, nil
}

// NewUniNftManagerTransactor creates a new write-only instance of UniNftManager, bound to a specific deployed contract.
func NewUniNftManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*UniNftManagerTransactor, error) {
	contract, err := bindUniNftManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniNftManagerTransactor{contract: contract}, nil
}

// NewUniNftManagerFilterer creates a new log filterer instance of UniNftManager, bound to a specific deployed contract.
func NewUniNftManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*UniNftManagerFilterer, error) {
	contract, err := bindUniNftManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniNftManagerFilterer{contract: contract}, nil
}

// bindUniNftManager binds a generic wrapper to an already deployed contract.
func bindUniNftManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniNftManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniNftManager *UniNftManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniNftManager.Contract.UniNftManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniNftManager *UniNftManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniNftManager.Contract.UniNftManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniNftManager *UniNftManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniNftManager.Contract.UniNftManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniNftManager *UniNftManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniNftManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniNftManager *UniNftManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniNftManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniNftManager *UniNftManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniNftManager.Contract.contract.Transact(opts, method, params...)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UniNftManager *UniNftManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniNftManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UniNftManager *UniNftManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _UniNftManager.Contract.OwnerOf(&_UniNftManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_UniNftManager *UniNftManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _UniNftManager.Contract.OwnerOf(&_UniNftManager.CallOpts, tokenId)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) pure returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniNftManager *UniNftManagerCaller) Positions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	var out []interface{}
	err := _UniNftManager.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Nonce                    *big.Int
		Operator                 common.Address
		Token0                   common.Address
		Token1                   common.Address
		Fee                      *big.Int
		TickLower                *big.Int
		TickUpper                *big.Int
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TickLower = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TickUpper = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) pure returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniNftManager *UniNftManagerSession) Positions(arg0 *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _UniNftManager.Contract.Positions(&_UniNftManager.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x99fbab88.
//
// Solidity: function positions(uint256 ) pure returns(uint96 nonce, address operator, address token0, address token1, uint24 fee, int24 tickLower, int24 tickUpper, uint128 liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1)
func (_UniNftManager *UniNftManagerCallerSession) Positions(arg0 *big.Int) (struct {
	Nonce                    *big.Int
	Operator                 common.Address
	Token0                   common.Address
	Token1                   common.Address
	Fee                      *big.Int
	TickLower                *big.Int
	TickUpper                *big.Int
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
}, error) {
	return _UniNftManager.Contract.Positions(&_UniNftManager.CallOpts, arg0)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) returns(uint256 amount0, uint256 amount1)
func (_UniNftManager *UniNftManagerTransactor) Collect(opts *bind.TransactOpts, params UniNftManagerCollectParams) (*types.Transaction, error) {
	return _UniNftManager.contract.Transact(opts, "collect", params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) returns(uint256 amount0, uint256 amount1)
func (_UniNftManager *UniNftManagerSession) Collect(params UniNftManagerCollectParams) (*types.Transaction, error) {
	return _UniNftManager.Contract.Collect(&_UniNftManager.TransactOpts, params)
}

// Collect is a paid mutator transaction binding the contract method 0xfc6f7865.
//
// Solidity: function collect((uint256,address,uint128,uint128) params) returns(uint256 amount0, uint256 amount1)
func (_UniNftManager *UniNftManagerTransactorSession) Collect(params UniNftManagerCollectParams) (*types.Transaction, error) {
	return _UniNftManager.Contract.Collect(&_UniNftManager.TransactOpts, params)
}
