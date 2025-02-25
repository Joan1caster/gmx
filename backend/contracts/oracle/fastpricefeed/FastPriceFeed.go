// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fastpricefeed

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

// FastPriceFeedMetaData contains all meta data concerning the FastPriceFeed contract.
var FastPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pyth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPriceUpdateDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDeviationBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fastPriceEvents\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"DisableFastPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"EnableFastPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeRefDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFastDelta\",\"type\":\"uint256\"}],\"name\":\"MaxCumulativeDeltaDiffExceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeRefDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFastDelta\",\"type\":\"uint256\"}],\"name\":\"PriceData\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CUMULATIVE_DELTA_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CUMULATIVE_FAST_DELTA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CUMULATIVE_REF_DELTA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PRICE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REF_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableFastPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableFastPriceVoteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"disableFastPriceVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableFastPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastPriceEvents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"favorFastPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_refPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAuthorizations\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_updaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_priceFeedIds\",\"type\":\"bytes32[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSpreadEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isUpdater\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxCumulativeDeltaDiffs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDeviationBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPriceUpdateDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAuthorizations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceData\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"refPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"refTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"cumulativeRefDelta\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"cumulativeFastDelta\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDataInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceFeedIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pyth\",\"outputs\":[{\"internalType\":\"contractIPyth\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSpreadEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSpreadEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastUpdatedAt\",\"type\":\"uint256\"}],\"name\":\"setLastUpdatedAt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_maxCumulativeDeltaDiffs\",\"type\":\"uint256[]\"}],\"name\":\"setMaxCumulativeDeltaDiffs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDeviationBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setMaxDeviationBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"setMaxPriceUpdateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAuthorizations\",\"type\":\"uint256\"}],\"name\":\"setMinAuthorizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBlockInterval\",\"type\":\"uint256\"}],\"name\":\"setMinBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceDataInterval\",\"type\":\"uint256\"}],\"name\":\"setPriceDataInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceDuration\",\"type\":\"uint256\"}],\"name\":\"setPriceDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"priceUpdateData\",\"type\":\"bytes[]\"}],\"name\":\"setPricesWithData\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_positionRouter\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"priceUpdateData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"_endIndexForIncreasePositions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndexForDecreasePositions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxIncreasePositions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDecreasePositions\",\"type\":\"uint256\"}],\"name\":\"setPricesWithDataAndExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfChainError\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfChainError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfInactive\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfInactive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadBasisPointsIfChainError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadBasisPointsIfInactive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FastPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use FastPriceFeedMetaData.ABI instead.
var FastPriceFeedABI = FastPriceFeedMetaData.ABI

// FastPriceFeed is an auto generated Go binding around an Ethereum contract.
type FastPriceFeed struct {
	FastPriceFeedCaller     // Read-only binding to the contract
	FastPriceFeedTransactor // Write-only binding to the contract
	FastPriceFeedFilterer   // Log filterer for contract events
}

// FastPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastPriceFeedSession struct {
	Contract     *FastPriceFeed    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FastPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastPriceFeedCallerSession struct {
	Contract *FastPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FastPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastPriceFeedTransactorSession struct {
	Contract     *FastPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FastPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastPriceFeedRaw struct {
	Contract *FastPriceFeed // Generic contract binding to access the raw methods on
}

// FastPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastPriceFeedCallerRaw struct {
	Contract *FastPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// FastPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastPriceFeedTransactorRaw struct {
	Contract *FastPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastPriceFeed creates a new instance of FastPriceFeed, bound to a specific deployed contract.
func NewFastPriceFeed(address common.Address, backend bind.ContractBackend) (*FastPriceFeed, error) {
	contract, err := bindFastPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastPriceFeed{FastPriceFeedCaller: FastPriceFeedCaller{contract: contract}, FastPriceFeedTransactor: FastPriceFeedTransactor{contract: contract}, FastPriceFeedFilterer: FastPriceFeedFilterer{contract: contract}}, nil
}

// NewFastPriceFeedCaller creates a new read-only instance of FastPriceFeed, bound to a specific deployed contract.
func NewFastPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*FastPriceFeedCaller, error) {
	contract, err := bindFastPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastPriceFeedCaller{contract: contract}, nil
}

// NewFastPriceFeedTransactor creates a new write-only instance of FastPriceFeed, bound to a specific deployed contract.
func NewFastPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*FastPriceFeedTransactor, error) {
	contract, err := bindFastPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastPriceFeedTransactor{contract: contract}, nil
}

// NewFastPriceFeedFilterer creates a new log filterer instance of FastPriceFeed, bound to a specific deployed contract.
func NewFastPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*FastPriceFeedFilterer, error) {
	contract, err := bindFastPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastPriceFeedFilterer{contract: contract}, nil
}

// bindFastPriceFeed binds a generic wrapper to an already deployed contract.
func bindFastPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastPriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastPriceFeed *FastPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastPriceFeed.Contract.FastPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastPriceFeed *FastPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.FastPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastPriceFeed *FastPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.FastPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastPriceFeed *FastPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastPriceFeed *FastPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastPriceFeed *FastPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _FastPriceFeed.Contract.BASISPOINTSDIVISOR(&_FastPriceFeed.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _FastPriceFeed.Contract.BASISPOINTSDIVISOR(&_FastPriceFeed.CallOpts)
}

// CUMULATIVEDELTAPRECISION is a free data retrieval call binding the contract method 0xa2b47c16.
//
// Solidity: function CUMULATIVE_DELTA_PRECISION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) CUMULATIVEDELTAPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "CUMULATIVE_DELTA_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CUMULATIVEDELTAPRECISION is a free data retrieval call binding the contract method 0xa2b47c16.
//
// Solidity: function CUMULATIVE_DELTA_PRECISION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) CUMULATIVEDELTAPRECISION() (*big.Int, error) {
	return _FastPriceFeed.Contract.CUMULATIVEDELTAPRECISION(&_FastPriceFeed.CallOpts)
}

// CUMULATIVEDELTAPRECISION is a free data retrieval call binding the contract method 0xa2b47c16.
//
// Solidity: function CUMULATIVE_DELTA_PRECISION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) CUMULATIVEDELTAPRECISION() (*big.Int, error) {
	return _FastPriceFeed.Contract.CUMULATIVEDELTAPRECISION(&_FastPriceFeed.CallOpts)
}

// MAXCUMULATIVEFASTDELTA is a free data retrieval call binding the contract method 0x4bd66c1c.
//
// Solidity: function MAX_CUMULATIVE_FAST_DELTA() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MAXCUMULATIVEFASTDELTA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "MAX_CUMULATIVE_FAST_DELTA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCUMULATIVEFASTDELTA is a free data retrieval call binding the contract method 0x4bd66c1c.
//
// Solidity: function MAX_CUMULATIVE_FAST_DELTA() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MAXCUMULATIVEFASTDELTA() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXCUMULATIVEFASTDELTA(&_FastPriceFeed.CallOpts)
}

// MAXCUMULATIVEFASTDELTA is a free data retrieval call binding the contract method 0x4bd66c1c.
//
// Solidity: function MAX_CUMULATIVE_FAST_DELTA() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MAXCUMULATIVEFASTDELTA() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXCUMULATIVEFASTDELTA(&_FastPriceFeed.CallOpts)
}

// MAXCUMULATIVEREFDELTA is a free data retrieval call binding the contract method 0x0604ddea.
//
// Solidity: function MAX_CUMULATIVE_REF_DELTA() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MAXCUMULATIVEREFDELTA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "MAX_CUMULATIVE_REF_DELTA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCUMULATIVEREFDELTA is a free data retrieval call binding the contract method 0x0604ddea.
//
// Solidity: function MAX_CUMULATIVE_REF_DELTA() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MAXCUMULATIVEREFDELTA() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXCUMULATIVEREFDELTA(&_FastPriceFeed.CallOpts)
}

// MAXCUMULATIVEREFDELTA is a free data retrieval call binding the contract method 0x0604ddea.
//
// Solidity: function MAX_CUMULATIVE_REF_DELTA() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MAXCUMULATIVEREFDELTA() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXCUMULATIVEREFDELTA(&_FastPriceFeed.CallOpts)
}

// MAXPRICEDURATION is a free data retrieval call binding the contract method 0x668d3d65.
//
// Solidity: function MAX_PRICE_DURATION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MAXPRICEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "MAX_PRICE_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPRICEDURATION is a free data retrieval call binding the contract method 0x668d3d65.
//
// Solidity: function MAX_PRICE_DURATION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MAXPRICEDURATION() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXPRICEDURATION(&_FastPriceFeed.CallOpts)
}

// MAXPRICEDURATION is a free data retrieval call binding the contract method 0x668d3d65.
//
// Solidity: function MAX_PRICE_DURATION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MAXPRICEDURATION() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXPRICEDURATION(&_FastPriceFeed.CallOpts)
}

// MAXREFPRICE is a free data retrieval call binding the contract method 0xe68a22c0.
//
// Solidity: function MAX_REF_PRICE() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MAXREFPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "MAX_REF_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREFPRICE is a free data retrieval call binding the contract method 0xe68a22c0.
//
// Solidity: function MAX_REF_PRICE() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MAXREFPRICE() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXREFPRICE(&_FastPriceFeed.CallOpts)
}

// MAXREFPRICE is a free data retrieval call binding the contract method 0xe68a22c0.
//
// Solidity: function MAX_REF_PRICE() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MAXREFPRICE() (*big.Int, error) {
	return _FastPriceFeed.Contract.MAXREFPRICE(&_FastPriceFeed.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) PRICEPRECISION() (*big.Int, error) {
	return _FastPriceFeed.Contract.PRICEPRECISION(&_FastPriceFeed.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _FastPriceFeed.Contract.PRICEPRECISION(&_FastPriceFeed.CallOpts)
}

// DisableFastPriceVoteCount is a free data retrieval call binding the contract method 0xb0a25666.
//
// Solidity: function disableFastPriceVoteCount() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) DisableFastPriceVoteCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "disableFastPriceVoteCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisableFastPriceVoteCount is a free data retrieval call binding the contract method 0xb0a25666.
//
// Solidity: function disableFastPriceVoteCount() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) DisableFastPriceVoteCount() (*big.Int, error) {
	return _FastPriceFeed.Contract.DisableFastPriceVoteCount(&_FastPriceFeed.CallOpts)
}

// DisableFastPriceVoteCount is a free data retrieval call binding the contract method 0xb0a25666.
//
// Solidity: function disableFastPriceVoteCount() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) DisableFastPriceVoteCount() (*big.Int, error) {
	return _FastPriceFeed.Contract.DisableFastPriceVoteCount(&_FastPriceFeed.CallOpts)
}

// DisableFastPriceVotes is a free data retrieval call binding the contract method 0x03b04936.
//
// Solidity: function disableFastPriceVotes(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCaller) DisableFastPriceVotes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "disableFastPriceVotes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DisableFastPriceVotes is a free data retrieval call binding the contract method 0x03b04936.
//
// Solidity: function disableFastPriceVotes(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedSession) DisableFastPriceVotes(arg0 common.Address) (bool, error) {
	return _FastPriceFeed.Contract.DisableFastPriceVotes(&_FastPriceFeed.CallOpts, arg0)
}

// DisableFastPriceVotes is a free data retrieval call binding the contract method 0x03b04936.
//
// Solidity: function disableFastPriceVotes(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCallerSession) DisableFastPriceVotes(arg0 common.Address) (bool, error) {
	return _FastPriceFeed.Contract.DisableFastPriceVotes(&_FastPriceFeed.CallOpts, arg0)
}

// FastPriceEvents is a free data retrieval call binding the contract method 0x0e9272ea.
//
// Solidity: function fastPriceEvents() view returns(address)
func (_FastPriceFeed *FastPriceFeedCaller) FastPriceEvents(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "fastPriceEvents")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastPriceEvents is a free data retrieval call binding the contract method 0x0e9272ea.
//
// Solidity: function fastPriceEvents() view returns(address)
func (_FastPriceFeed *FastPriceFeedSession) FastPriceEvents() (common.Address, error) {
	return _FastPriceFeed.Contract.FastPriceEvents(&_FastPriceFeed.CallOpts)
}

// FastPriceEvents is a free data retrieval call binding the contract method 0x0e9272ea.
//
// Solidity: function fastPriceEvents() view returns(address)
func (_FastPriceFeed *FastPriceFeedCallerSession) FastPriceEvents() (common.Address, error) {
	return _FastPriceFeed.Contract.FastPriceEvents(&_FastPriceFeed.CallOpts)
}

// FavorFastPrice is a free data retrieval call binding the contract method 0x6c56fd05.
//
// Solidity: function favorFastPrice(address _token) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCaller) FavorFastPrice(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "favorFastPrice", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FavorFastPrice is a free data retrieval call binding the contract method 0x6c56fd05.
//
// Solidity: function favorFastPrice(address _token) view returns(bool)
func (_FastPriceFeed *FastPriceFeedSession) FavorFastPrice(_token common.Address) (bool, error) {
	return _FastPriceFeed.Contract.FavorFastPrice(&_FastPriceFeed.CallOpts, _token)
}

// FavorFastPrice is a free data retrieval call binding the contract method 0x6c56fd05.
//
// Solidity: function favorFastPrice(address _token) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCallerSession) FavorFastPrice(_token common.Address) (bool, error) {
	return _FastPriceFeed.Contract.FavorFastPrice(&_FastPriceFeed.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _refPrice, bool _maximise) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) GetPrice(opts *bind.CallOpts, _token common.Address, _refPrice *big.Int, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "getPrice", _token, _refPrice, _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _refPrice, bool _maximise) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) GetPrice(_token common.Address, _refPrice *big.Int, _maximise bool) (*big.Int, error) {
	return _FastPriceFeed.Contract.GetPrice(&_FastPriceFeed.CallOpts, _token, _refPrice, _maximise)
}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _refPrice, bool _maximise) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) GetPrice(_token common.Address, _refPrice *big.Int, _maximise bool) (*big.Int, error) {
	return _FastPriceFeed.Contract.GetPrice(&_FastPriceFeed.CallOpts, _token, _refPrice, _maximise)
}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address _token) view returns(uint256, uint256, uint256, uint256)
func (_FastPriceFeed *FastPriceFeedCaller) GetPriceData(opts *bind.CallOpts, _token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "getPriceData", _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address _token) view returns(uint256, uint256, uint256, uint256)
func (_FastPriceFeed *FastPriceFeedSession) GetPriceData(_token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FastPriceFeed.Contract.GetPriceData(&_FastPriceFeed.CallOpts, _token)
}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address _token) view returns(uint256, uint256, uint256, uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) GetPriceData(_token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FastPriceFeed.Contract.GetPriceData(&_FastPriceFeed.CallOpts, _token)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceFeed *FastPriceFeedCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceFeed *FastPriceFeedSession) Gov() (common.Address, error) {
	return _FastPriceFeed.Contract.Gov(&_FastPriceFeed.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceFeed *FastPriceFeedCallerSession) Gov() (common.Address, error) {
	return _FastPriceFeed.Contract.Gov(&_FastPriceFeed.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_FastPriceFeed *FastPriceFeedCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_FastPriceFeed *FastPriceFeedSession) IsInitialized() (bool, error) {
	return _FastPriceFeed.Contract.IsInitialized(&_FastPriceFeed.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_FastPriceFeed *FastPriceFeedCallerSession) IsInitialized() (bool, error) {
	return _FastPriceFeed.Contract.IsInitialized(&_FastPriceFeed.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCaller) IsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "isSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedSession) IsSigner(arg0 common.Address) (bool, error) {
	return _FastPriceFeed.Contract.IsSigner(&_FastPriceFeed.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCallerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _FastPriceFeed.Contract.IsSigner(&_FastPriceFeed.CallOpts, arg0)
}

// IsSpreadEnabled is a free data retrieval call binding the contract method 0x695d4184.
//
// Solidity: function isSpreadEnabled() view returns(bool)
func (_FastPriceFeed *FastPriceFeedCaller) IsSpreadEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "isSpreadEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpreadEnabled is a free data retrieval call binding the contract method 0x695d4184.
//
// Solidity: function isSpreadEnabled() view returns(bool)
func (_FastPriceFeed *FastPriceFeedSession) IsSpreadEnabled() (bool, error) {
	return _FastPriceFeed.Contract.IsSpreadEnabled(&_FastPriceFeed.CallOpts)
}

// IsSpreadEnabled is a free data retrieval call binding the contract method 0x695d4184.
//
// Solidity: function isSpreadEnabled() view returns(bool)
func (_FastPriceFeed *FastPriceFeedCallerSession) IsSpreadEnabled() (bool, error) {
	return _FastPriceFeed.Contract.IsSpreadEnabled(&_FastPriceFeed.CallOpts)
}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCaller) IsUpdater(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "isUpdater", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedSession) IsUpdater(arg0 common.Address) (bool, error) {
	return _FastPriceFeed.Contract.IsUpdater(&_FastPriceFeed.CallOpts, arg0)
}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_FastPriceFeed *FastPriceFeedCallerSession) IsUpdater(arg0 common.Address) (bool, error) {
	return _FastPriceFeed.Contract.IsUpdater(&_FastPriceFeed.CallOpts, arg0)
}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) LastUpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "lastUpdatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) LastUpdatedAt() (*big.Int, error) {
	return _FastPriceFeed.Contract.LastUpdatedAt(&_FastPriceFeed.CallOpts)
}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) LastUpdatedAt() (*big.Int, error) {
	return _FastPriceFeed.Contract.LastUpdatedAt(&_FastPriceFeed.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) LastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "lastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) LastUpdatedBlock() (*big.Int, error) {
	return _FastPriceFeed.Contract.LastUpdatedBlock(&_FastPriceFeed.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _FastPriceFeed.Contract.LastUpdatedBlock(&_FastPriceFeed.CallOpts)
}

// MaxCumulativeDeltaDiffs is a free data retrieval call binding the contract method 0xa3742425.
//
// Solidity: function maxCumulativeDeltaDiffs(address ) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MaxCumulativeDeltaDiffs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "maxCumulativeDeltaDiffs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCumulativeDeltaDiffs is a free data retrieval call binding the contract method 0xa3742425.
//
// Solidity: function maxCumulativeDeltaDiffs(address ) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MaxCumulativeDeltaDiffs(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFeed.Contract.MaxCumulativeDeltaDiffs(&_FastPriceFeed.CallOpts, arg0)
}

// MaxCumulativeDeltaDiffs is a free data retrieval call binding the contract method 0xa3742425.
//
// Solidity: function maxCumulativeDeltaDiffs(address ) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MaxCumulativeDeltaDiffs(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFeed.Contract.MaxCumulativeDeltaDiffs(&_FastPriceFeed.CallOpts, arg0)
}

// MaxDeviationBasisPoints is a free data retrieval call binding the contract method 0x715c7536.
//
// Solidity: function maxDeviationBasisPoints() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MaxDeviationBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "maxDeviationBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeviationBasisPoints is a free data retrieval call binding the contract method 0x715c7536.
//
// Solidity: function maxDeviationBasisPoints() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MaxDeviationBasisPoints() (*big.Int, error) {
	return _FastPriceFeed.Contract.MaxDeviationBasisPoints(&_FastPriceFeed.CallOpts)
}

// MaxDeviationBasisPoints is a free data retrieval call binding the contract method 0x715c7536.
//
// Solidity: function maxDeviationBasisPoints() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MaxDeviationBasisPoints() (*big.Int, error) {
	return _FastPriceFeed.Contract.MaxDeviationBasisPoints(&_FastPriceFeed.CallOpts)
}

// MaxPriceUpdateDelay is a free data retrieval call binding the contract method 0xe64559ad.
//
// Solidity: function maxPriceUpdateDelay() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MaxPriceUpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "maxPriceUpdateDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPriceUpdateDelay is a free data retrieval call binding the contract method 0xe64559ad.
//
// Solidity: function maxPriceUpdateDelay() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MaxPriceUpdateDelay() (*big.Int, error) {
	return _FastPriceFeed.Contract.MaxPriceUpdateDelay(&_FastPriceFeed.CallOpts)
}

// MaxPriceUpdateDelay is a free data retrieval call binding the contract method 0xe64559ad.
//
// Solidity: function maxPriceUpdateDelay() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MaxPriceUpdateDelay() (*big.Int, error) {
	return _FastPriceFeed.Contract.MaxPriceUpdateDelay(&_FastPriceFeed.CallOpts)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MinAuthorizations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "minAuthorizations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MinAuthorizations() (*big.Int, error) {
	return _FastPriceFeed.Contract.MinAuthorizations(&_FastPriceFeed.CallOpts)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MinAuthorizations() (*big.Int, error) {
	return _FastPriceFeed.Contract.MinAuthorizations(&_FastPriceFeed.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) MinBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "minBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) MinBlockInterval() (*big.Int, error) {
	return _FastPriceFeed.Contract.MinBlockInterval(&_FastPriceFeed.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) MinBlockInterval() (*big.Int, error) {
	return _FastPriceFeed.Contract.MinBlockInterval(&_FastPriceFeed.CallOpts)
}

// PriceData is a free data retrieval call binding the contract method 0xcab44b76.
//
// Solidity: function priceData(address ) view returns(uint160 refPrice, uint32 refTime, uint32 cumulativeRefDelta, uint32 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedCaller) PriceData(opts *bind.CallOpts, arg0 common.Address) (struct {
	RefPrice            *big.Int
	RefTime             uint32
	CumulativeRefDelta  uint32
	CumulativeFastDelta uint32
}, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "priceData", arg0)

	outstruct := new(struct {
		RefPrice            *big.Int
		RefTime             uint32
		CumulativeRefDelta  uint32
		CumulativeFastDelta uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RefPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RefTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.CumulativeRefDelta = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.CumulativeFastDelta = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// PriceData is a free data retrieval call binding the contract method 0xcab44b76.
//
// Solidity: function priceData(address ) view returns(uint160 refPrice, uint32 refTime, uint32 cumulativeRefDelta, uint32 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedSession) PriceData(arg0 common.Address) (struct {
	RefPrice            *big.Int
	RefTime             uint32
	CumulativeRefDelta  uint32
	CumulativeFastDelta uint32
}, error) {
	return _FastPriceFeed.Contract.PriceData(&_FastPriceFeed.CallOpts, arg0)
}

// PriceData is a free data retrieval call binding the contract method 0xcab44b76.
//
// Solidity: function priceData(address ) view returns(uint160 refPrice, uint32 refTime, uint32 cumulativeRefDelta, uint32 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedCallerSession) PriceData(arg0 common.Address) (struct {
	RefPrice            *big.Int
	RefTime             uint32
	CumulativeRefDelta  uint32
	CumulativeFastDelta uint32
}, error) {
	return _FastPriceFeed.Contract.PriceData(&_FastPriceFeed.CallOpts, arg0)
}

// PriceDataInterval is a free data retrieval call binding the contract method 0xdfb481c9.
//
// Solidity: function priceDataInterval() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) PriceDataInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "priceDataInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDataInterval is a free data retrieval call binding the contract method 0xdfb481c9.
//
// Solidity: function priceDataInterval() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) PriceDataInterval() (*big.Int, error) {
	return _FastPriceFeed.Contract.PriceDataInterval(&_FastPriceFeed.CallOpts)
}

// PriceDataInterval is a free data retrieval call binding the contract method 0xdfb481c9.
//
// Solidity: function priceDataInterval() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) PriceDataInterval() (*big.Int, error) {
	return _FastPriceFeed.Contract.PriceDataInterval(&_FastPriceFeed.CallOpts)
}

// PriceDuration is a free data retrieval call binding the contract method 0x03cd2571.
//
// Solidity: function priceDuration() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) PriceDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "priceDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDuration is a free data retrieval call binding the contract method 0x03cd2571.
//
// Solidity: function priceDuration() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) PriceDuration() (*big.Int, error) {
	return _FastPriceFeed.Contract.PriceDuration(&_FastPriceFeed.CallOpts)
}

// PriceDuration is a free data retrieval call binding the contract method 0x03cd2571.
//
// Solidity: function priceDuration() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) PriceDuration() (*big.Int, error) {
	return _FastPriceFeed.Contract.PriceDuration(&_FastPriceFeed.CallOpts)
}

// PriceFeedIds is a free data retrieval call binding the contract method 0x336eb6ae.
//
// Solidity: function priceFeedIds(address ) view returns(bytes32)
func (_FastPriceFeed *FastPriceFeedCaller) PriceFeedIds(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "priceFeedIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PriceFeedIds is a free data retrieval call binding the contract method 0x336eb6ae.
//
// Solidity: function priceFeedIds(address ) view returns(bytes32)
func (_FastPriceFeed *FastPriceFeedSession) PriceFeedIds(arg0 common.Address) ([32]byte, error) {
	return _FastPriceFeed.Contract.PriceFeedIds(&_FastPriceFeed.CallOpts, arg0)
}

// PriceFeedIds is a free data retrieval call binding the contract method 0x336eb6ae.
//
// Solidity: function priceFeedIds(address ) view returns(bytes32)
func (_FastPriceFeed *FastPriceFeedCallerSession) PriceFeedIds(arg0 common.Address) ([32]byte, error) {
	return _FastPriceFeed.Contract.PriceFeedIds(&_FastPriceFeed.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) Prices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "prices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) Prices(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFeed.Contract.Prices(&_FastPriceFeed.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) Prices(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFeed.Contract.Prices(&_FastPriceFeed.CallOpts, arg0)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_FastPriceFeed *FastPriceFeedCaller) Pyth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "pyth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_FastPriceFeed *FastPriceFeedSession) Pyth() (common.Address, error) {
	return _FastPriceFeed.Contract.Pyth(&_FastPriceFeed.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_FastPriceFeed *FastPriceFeedCallerSession) Pyth() (common.Address, error) {
	return _FastPriceFeed.Contract.Pyth(&_FastPriceFeed.CallOpts)
}

// SpreadBasisPointsIfChainError is a free data retrieval call binding the contract method 0xa6eca896.
//
// Solidity: function spreadBasisPointsIfChainError() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) SpreadBasisPointsIfChainError(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "spreadBasisPointsIfChainError")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadBasisPointsIfChainError is a free data retrieval call binding the contract method 0xa6eca896.
//
// Solidity: function spreadBasisPointsIfChainError() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) SpreadBasisPointsIfChainError() (*big.Int, error) {
	return _FastPriceFeed.Contract.SpreadBasisPointsIfChainError(&_FastPriceFeed.CallOpts)
}

// SpreadBasisPointsIfChainError is a free data retrieval call binding the contract method 0xa6eca896.
//
// Solidity: function spreadBasisPointsIfChainError() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) SpreadBasisPointsIfChainError() (*big.Int, error) {
	return _FastPriceFeed.Contract.SpreadBasisPointsIfChainError(&_FastPriceFeed.CallOpts)
}

// SpreadBasisPointsIfInactive is a free data retrieval call binding the contract method 0x74bfed89.
//
// Solidity: function spreadBasisPointsIfInactive() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCaller) SpreadBasisPointsIfInactive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "spreadBasisPointsIfInactive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadBasisPointsIfInactive is a free data retrieval call binding the contract method 0x74bfed89.
//
// Solidity: function spreadBasisPointsIfInactive() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedSession) SpreadBasisPointsIfInactive() (*big.Int, error) {
	return _FastPriceFeed.Contract.SpreadBasisPointsIfInactive(&_FastPriceFeed.CallOpts)
}

// SpreadBasisPointsIfInactive is a free data retrieval call binding the contract method 0x74bfed89.
//
// Solidity: function spreadBasisPointsIfInactive() view returns(uint256)
func (_FastPriceFeed *FastPriceFeedCallerSession) SpreadBasisPointsIfInactive() (*big.Int, error) {
	return _FastPriceFeed.Contract.SpreadBasisPointsIfInactive(&_FastPriceFeed.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_FastPriceFeed *FastPriceFeedCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_FastPriceFeed *FastPriceFeedSession) TokenManager() (common.Address, error) {
	return _FastPriceFeed.Contract.TokenManager(&_FastPriceFeed.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_FastPriceFeed *FastPriceFeedCallerSession) TokenManager() (common.Address, error) {
	return _FastPriceFeed.Contract.TokenManager(&_FastPriceFeed.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_FastPriceFeed *FastPriceFeedCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_FastPriceFeed *FastPriceFeedSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _FastPriceFeed.Contract.Tokens(&_FastPriceFeed.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_FastPriceFeed *FastPriceFeedCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _FastPriceFeed.Contract.Tokens(&_FastPriceFeed.CallOpts, arg0)
}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_FastPriceFeed *FastPriceFeedCaller) VaultPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFeed.contract.Call(opts, &out, "vaultPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_FastPriceFeed *FastPriceFeedSession) VaultPriceFeed() (common.Address, error) {
	return _FastPriceFeed.Contract.VaultPriceFeed(&_FastPriceFeed.CallOpts)
}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_FastPriceFeed *FastPriceFeedCallerSession) VaultPriceFeed() (common.Address, error) {
	return _FastPriceFeed.Contract.VaultPriceFeed(&_FastPriceFeed.CallOpts)
}

// DisableFastPrice is a paid mutator transaction binding the contract method 0xc84a9124.
//
// Solidity: function disableFastPrice() returns()
func (_FastPriceFeed *FastPriceFeedTransactor) DisableFastPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "disableFastPrice")
}

// DisableFastPrice is a paid mutator transaction binding the contract method 0xc84a9124.
//
// Solidity: function disableFastPrice() returns()
func (_FastPriceFeed *FastPriceFeedSession) DisableFastPrice() (*types.Transaction, error) {
	return _FastPriceFeed.Contract.DisableFastPrice(&_FastPriceFeed.TransactOpts)
}

// DisableFastPrice is a paid mutator transaction binding the contract method 0xc84a9124.
//
// Solidity: function disableFastPrice() returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) DisableFastPrice() (*types.Transaction, error) {
	return _FastPriceFeed.Contract.DisableFastPrice(&_FastPriceFeed.TransactOpts)
}

// EnableFastPrice is a paid mutator transaction binding the contract method 0x6ccd47c4.
//
// Solidity: function enableFastPrice() returns()
func (_FastPriceFeed *FastPriceFeedTransactor) EnableFastPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "enableFastPrice")
}

// EnableFastPrice is a paid mutator transaction binding the contract method 0x6ccd47c4.
//
// Solidity: function enableFastPrice() returns()
func (_FastPriceFeed *FastPriceFeedSession) EnableFastPrice() (*types.Transaction, error) {
	return _FastPriceFeed.Contract.EnableFastPrice(&_FastPriceFeed.TransactOpts)
}

// EnableFastPrice is a paid mutator transaction binding the contract method 0x6ccd47c4.
//
// Solidity: function enableFastPrice() returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) EnableFastPrice() (*types.Transaction, error) {
	return _FastPriceFeed.Contract.EnableFastPrice(&_FastPriceFeed.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x63a22d71.
//
// Solidity: function initialize(uint256 _minAuthorizations, address[] _signers, address[] _updaters, address[] _tokens, bytes32[] _priceFeedIds) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) Initialize(opts *bind.TransactOpts, _minAuthorizations *big.Int, _signers []common.Address, _updaters []common.Address, _tokens []common.Address, _priceFeedIds [][32]byte) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "initialize", _minAuthorizations, _signers, _updaters, _tokens, _priceFeedIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x63a22d71.
//
// Solidity: function initialize(uint256 _minAuthorizations, address[] _signers, address[] _updaters, address[] _tokens, bytes32[] _priceFeedIds) returns()
func (_FastPriceFeed *FastPriceFeedSession) Initialize(_minAuthorizations *big.Int, _signers []common.Address, _updaters []common.Address, _tokens []common.Address, _priceFeedIds [][32]byte) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.Initialize(&_FastPriceFeed.TransactOpts, _minAuthorizations, _signers, _updaters, _tokens, _priceFeedIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x63a22d71.
//
// Solidity: function initialize(uint256 _minAuthorizations, address[] _signers, address[] _updaters, address[] _tokens, bytes32[] _priceFeedIds) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) Initialize(_minAuthorizations *big.Int, _signers []common.Address, _updaters []common.Address, _tokens []common.Address, _priceFeedIds [][32]byte) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.Initialize(&_FastPriceFeed.TransactOpts, _minAuthorizations, _signers, _updaters, _tokens, _priceFeedIds)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetGov(&_FastPriceFeed.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetGov(&_FastPriceFeed.TransactOpts, _gov)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetIsSpreadEnabled(opts *bind.TransactOpts, _isSpreadEnabled bool) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setIsSpreadEnabled", _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetIsSpreadEnabled(_isSpreadEnabled bool) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetIsSpreadEnabled(&_FastPriceFeed.TransactOpts, _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetIsSpreadEnabled(_isSpreadEnabled bool) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetIsSpreadEnabled(&_FastPriceFeed.TransactOpts, _isSpreadEnabled)
}

// SetLastUpdatedAt is a paid mutator transaction binding the contract method 0x14dd2dce.
//
// Solidity: function setLastUpdatedAt(uint256 _lastUpdatedAt) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetLastUpdatedAt(opts *bind.TransactOpts, _lastUpdatedAt *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setLastUpdatedAt", _lastUpdatedAt)
}

// SetLastUpdatedAt is a paid mutator transaction binding the contract method 0x14dd2dce.
//
// Solidity: function setLastUpdatedAt(uint256 _lastUpdatedAt) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetLastUpdatedAt(_lastUpdatedAt *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetLastUpdatedAt(&_FastPriceFeed.TransactOpts, _lastUpdatedAt)
}

// SetLastUpdatedAt is a paid mutator transaction binding the contract method 0x14dd2dce.
//
// Solidity: function setLastUpdatedAt(uint256 _lastUpdatedAt) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetLastUpdatedAt(_lastUpdatedAt *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetLastUpdatedAt(&_FastPriceFeed.TransactOpts, _lastUpdatedAt)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetMaxCumulativeDeltaDiffs(opts *bind.TransactOpts, _tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setMaxCumulativeDeltaDiffs", _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetMaxCumulativeDeltaDiffs(_tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMaxCumulativeDeltaDiffs(&_FastPriceFeed.TransactOpts, _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetMaxCumulativeDeltaDiffs(_tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMaxCumulativeDeltaDiffs(&_FastPriceFeed.TransactOpts, _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetMaxDeviationBasisPoints(opts *bind.TransactOpts, _maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setMaxDeviationBasisPoints", _maxDeviationBasisPoints)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetMaxDeviationBasisPoints(_maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMaxDeviationBasisPoints(&_FastPriceFeed.TransactOpts, _maxDeviationBasisPoints)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetMaxDeviationBasisPoints(_maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMaxDeviationBasisPoints(&_FastPriceFeed.TransactOpts, _maxDeviationBasisPoints)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetMaxPriceUpdateDelay(opts *bind.TransactOpts, _maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setMaxPriceUpdateDelay", _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetMaxPriceUpdateDelay(_maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMaxPriceUpdateDelay(&_FastPriceFeed.TransactOpts, _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetMaxPriceUpdateDelay(_maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMaxPriceUpdateDelay(&_FastPriceFeed.TransactOpts, _maxPriceUpdateDelay)
}

// SetMinAuthorizations is a paid mutator transaction binding the contract method 0xd925351a.
//
// Solidity: function setMinAuthorizations(uint256 _minAuthorizations) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetMinAuthorizations(opts *bind.TransactOpts, _minAuthorizations *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setMinAuthorizations", _minAuthorizations)
}

// SetMinAuthorizations is a paid mutator transaction binding the contract method 0xd925351a.
//
// Solidity: function setMinAuthorizations(uint256 _minAuthorizations) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetMinAuthorizations(_minAuthorizations *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMinAuthorizations(&_FastPriceFeed.TransactOpts, _minAuthorizations)
}

// SetMinAuthorizations is a paid mutator transaction binding the contract method 0xd925351a.
//
// Solidity: function setMinAuthorizations(uint256 _minAuthorizations) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetMinAuthorizations(_minAuthorizations *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMinAuthorizations(&_FastPriceFeed.TransactOpts, _minAuthorizations)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetMinBlockInterval(opts *bind.TransactOpts, _minBlockInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setMinBlockInterval", _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetMinBlockInterval(_minBlockInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMinBlockInterval(&_FastPriceFeed.TransactOpts, _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetMinBlockInterval(_minBlockInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetMinBlockInterval(&_FastPriceFeed.TransactOpts, _minBlockInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetPriceDataInterval(opts *bind.TransactOpts, _priceDataInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setPriceDataInterval", _priceDataInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetPriceDataInterval(_priceDataInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPriceDataInterval(&_FastPriceFeed.TransactOpts, _priceDataInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetPriceDataInterval(_priceDataInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPriceDataInterval(&_FastPriceFeed.TransactOpts, _priceDataInterval)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetPriceDuration(opts *bind.TransactOpts, _priceDuration *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setPriceDuration", _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetPriceDuration(_priceDuration *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPriceDuration(&_FastPriceFeed.TransactOpts, _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetPriceDuration(_priceDuration *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPriceDuration(&_FastPriceFeed.TransactOpts, _priceDuration)
}

// SetPricesWithData is a paid mutator transaction binding the contract method 0x78e8ab17.
//
// Solidity: function setPricesWithData(bytes[] priceUpdateData) payable returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetPricesWithData(opts *bind.TransactOpts, priceUpdateData [][]byte) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setPricesWithData", priceUpdateData)
}

// SetPricesWithData is a paid mutator transaction binding the contract method 0x78e8ab17.
//
// Solidity: function setPricesWithData(bytes[] priceUpdateData) payable returns()
func (_FastPriceFeed *FastPriceFeedSession) SetPricesWithData(priceUpdateData [][]byte) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPricesWithData(&_FastPriceFeed.TransactOpts, priceUpdateData)
}

// SetPricesWithData is a paid mutator transaction binding the contract method 0x78e8ab17.
//
// Solidity: function setPricesWithData(bytes[] priceUpdateData) payable returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetPricesWithData(priceUpdateData [][]byte) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPricesWithData(&_FastPriceFeed.TransactOpts, priceUpdateData)
}

// SetPricesWithDataAndExecute is a paid mutator transaction binding the contract method 0x4576bbd8.
//
// Solidity: function setPricesWithDataAndExecute(address _positionRouter, bytes[] priceUpdateData, uint256 _endIndexForIncreasePositions, uint256 _endIndexForDecreasePositions, uint256 _maxIncreasePositions, uint256 _maxDecreasePositions) payable returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetPricesWithDataAndExecute(opts *bind.TransactOpts, _positionRouter common.Address, priceUpdateData [][]byte, _endIndexForIncreasePositions *big.Int, _endIndexForDecreasePositions *big.Int, _maxIncreasePositions *big.Int, _maxDecreasePositions *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setPricesWithDataAndExecute", _positionRouter, priceUpdateData, _endIndexForIncreasePositions, _endIndexForDecreasePositions, _maxIncreasePositions, _maxDecreasePositions)
}

// SetPricesWithDataAndExecute is a paid mutator transaction binding the contract method 0x4576bbd8.
//
// Solidity: function setPricesWithDataAndExecute(address _positionRouter, bytes[] priceUpdateData, uint256 _endIndexForIncreasePositions, uint256 _endIndexForDecreasePositions, uint256 _maxIncreasePositions, uint256 _maxDecreasePositions) payable returns()
func (_FastPriceFeed *FastPriceFeedSession) SetPricesWithDataAndExecute(_positionRouter common.Address, priceUpdateData [][]byte, _endIndexForIncreasePositions *big.Int, _endIndexForDecreasePositions *big.Int, _maxIncreasePositions *big.Int, _maxDecreasePositions *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPricesWithDataAndExecute(&_FastPriceFeed.TransactOpts, _positionRouter, priceUpdateData, _endIndexForIncreasePositions, _endIndexForDecreasePositions, _maxIncreasePositions, _maxDecreasePositions)
}

// SetPricesWithDataAndExecute is a paid mutator transaction binding the contract method 0x4576bbd8.
//
// Solidity: function setPricesWithDataAndExecute(address _positionRouter, bytes[] priceUpdateData, uint256 _endIndexForIncreasePositions, uint256 _endIndexForDecreasePositions, uint256 _maxIncreasePositions, uint256 _maxDecreasePositions) payable returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetPricesWithDataAndExecute(_positionRouter common.Address, priceUpdateData [][]byte, _endIndexForIncreasePositions *big.Int, _endIndexForDecreasePositions *big.Int, _maxIncreasePositions *big.Int, _maxDecreasePositions *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetPricesWithDataAndExecute(&_FastPriceFeed.TransactOpts, _positionRouter, priceUpdateData, _endIndexForIncreasePositions, _endIndexForDecreasePositions, _maxIncreasePositions, _maxDecreasePositions)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetSigner(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setSigner", _account, _isActive)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetSigner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetSigner(&_FastPriceFeed.TransactOpts, _account, _isActive)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetSigner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetSigner(&_FastPriceFeed.TransactOpts, _account, _isActive)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetSpreadBasisPointsIfChainError(opts *bind.TransactOpts, _spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setSpreadBasisPointsIfChainError", _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetSpreadBasisPointsIfChainError(_spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetSpreadBasisPointsIfChainError(&_FastPriceFeed.TransactOpts, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetSpreadBasisPointsIfChainError(_spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetSpreadBasisPointsIfChainError(&_FastPriceFeed.TransactOpts, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetSpreadBasisPointsIfInactive(opts *bind.TransactOpts, _spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setSpreadBasisPointsIfInactive", _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetSpreadBasisPointsIfInactive(_spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetSpreadBasisPointsIfInactive(&_FastPriceFeed.TransactOpts, _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetSpreadBasisPointsIfInactive(_spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetSpreadBasisPointsIfInactive(&_FastPriceFeed.TransactOpts, _spreadBasisPointsIfInactive)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetTokenManager(opts *bind.TransactOpts, _tokenManager common.Address) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setTokenManager", _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetTokenManager(&_FastPriceFeed.TransactOpts, _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetTokenManager(&_FastPriceFeed.TransactOpts, _tokenManager)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_FastPriceFeed *FastPriceFeedTransactor) SetUpdater(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFeed.contract.Transact(opts, "setUpdater", _account, _isActive)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_FastPriceFeed *FastPriceFeedSession) SetUpdater(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetUpdater(&_FastPriceFeed.TransactOpts, _account, _isActive)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_FastPriceFeed *FastPriceFeedTransactorSession) SetUpdater(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFeed.Contract.SetUpdater(&_FastPriceFeed.TransactOpts, _account, _isActive)
}

// FastPriceFeedDisableFastPriceIterator is returned from FilterDisableFastPrice and is used to iterate over the raw logs and unpacked data for DisableFastPrice events raised by the FastPriceFeed contract.
type FastPriceFeedDisableFastPriceIterator struct {
	Event *FastPriceFeedDisableFastPrice // Event containing the contract specifics and raw log

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
func (it *FastPriceFeedDisableFastPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeedDisableFastPrice)
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
		it.Event = new(FastPriceFeedDisableFastPrice)
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
func (it *FastPriceFeedDisableFastPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeedDisableFastPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeedDisableFastPrice represents a DisableFastPrice event raised by the FastPriceFeed contract.
type FastPriceFeedDisableFastPrice struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisableFastPrice is a free log retrieval operation binding the contract event 0x4c0c5fabf50e808e3bc8d19577d305e3a7163eea7e8a74a50caa8896694cd44b.
//
// Solidity: event DisableFastPrice(address signer)
func (_FastPriceFeed *FastPriceFeedFilterer) FilterDisableFastPrice(opts *bind.FilterOpts) (*FastPriceFeedDisableFastPriceIterator, error) {

	logs, sub, err := _FastPriceFeed.contract.FilterLogs(opts, "DisableFastPrice")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeedDisableFastPriceIterator{contract: _FastPriceFeed.contract, event: "DisableFastPrice", logs: logs, sub: sub}, nil
}

// WatchDisableFastPrice is a free log subscription operation binding the contract event 0x4c0c5fabf50e808e3bc8d19577d305e3a7163eea7e8a74a50caa8896694cd44b.
//
// Solidity: event DisableFastPrice(address signer)
func (_FastPriceFeed *FastPriceFeedFilterer) WatchDisableFastPrice(opts *bind.WatchOpts, sink chan<- *FastPriceFeedDisableFastPrice) (event.Subscription, error) {

	logs, sub, err := _FastPriceFeed.contract.WatchLogs(opts, "DisableFastPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeedDisableFastPrice)
				if err := _FastPriceFeed.contract.UnpackLog(event, "DisableFastPrice", log); err != nil {
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

// ParseDisableFastPrice is a log parse operation binding the contract event 0x4c0c5fabf50e808e3bc8d19577d305e3a7163eea7e8a74a50caa8896694cd44b.
//
// Solidity: event DisableFastPrice(address signer)
func (_FastPriceFeed *FastPriceFeedFilterer) ParseDisableFastPrice(log types.Log) (*FastPriceFeedDisableFastPrice, error) {
	event := new(FastPriceFeedDisableFastPrice)
	if err := _FastPriceFeed.contract.UnpackLog(event, "DisableFastPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastPriceFeedEnableFastPriceIterator is returned from FilterEnableFastPrice and is used to iterate over the raw logs and unpacked data for EnableFastPrice events raised by the FastPriceFeed contract.
type FastPriceFeedEnableFastPriceIterator struct {
	Event *FastPriceFeedEnableFastPrice // Event containing the contract specifics and raw log

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
func (it *FastPriceFeedEnableFastPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeedEnableFastPrice)
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
		it.Event = new(FastPriceFeedEnableFastPrice)
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
func (it *FastPriceFeedEnableFastPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeedEnableFastPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeedEnableFastPrice represents a EnableFastPrice event raised by the FastPriceFeed contract.
type FastPriceFeedEnableFastPrice struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnableFastPrice is a free log retrieval operation binding the contract event 0x9fe0c305c33aa92757a537936872a60be0d91549a4303cc99fd8b7fce8a00275.
//
// Solidity: event EnableFastPrice(address signer)
func (_FastPriceFeed *FastPriceFeedFilterer) FilterEnableFastPrice(opts *bind.FilterOpts) (*FastPriceFeedEnableFastPriceIterator, error) {

	logs, sub, err := _FastPriceFeed.contract.FilterLogs(opts, "EnableFastPrice")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeedEnableFastPriceIterator{contract: _FastPriceFeed.contract, event: "EnableFastPrice", logs: logs, sub: sub}, nil
}

// WatchEnableFastPrice is a free log subscription operation binding the contract event 0x9fe0c305c33aa92757a537936872a60be0d91549a4303cc99fd8b7fce8a00275.
//
// Solidity: event EnableFastPrice(address signer)
func (_FastPriceFeed *FastPriceFeedFilterer) WatchEnableFastPrice(opts *bind.WatchOpts, sink chan<- *FastPriceFeedEnableFastPrice) (event.Subscription, error) {

	logs, sub, err := _FastPriceFeed.contract.WatchLogs(opts, "EnableFastPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeedEnableFastPrice)
				if err := _FastPriceFeed.contract.UnpackLog(event, "EnableFastPrice", log); err != nil {
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

// ParseEnableFastPrice is a log parse operation binding the contract event 0x9fe0c305c33aa92757a537936872a60be0d91549a4303cc99fd8b7fce8a00275.
//
// Solidity: event EnableFastPrice(address signer)
func (_FastPriceFeed *FastPriceFeedFilterer) ParseEnableFastPrice(log types.Log) (*FastPriceFeedEnableFastPrice, error) {
	event := new(FastPriceFeedEnableFastPrice)
	if err := _FastPriceFeed.contract.UnpackLog(event, "EnableFastPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastPriceFeedMaxCumulativeDeltaDiffExceededIterator is returned from FilterMaxCumulativeDeltaDiffExceeded and is used to iterate over the raw logs and unpacked data for MaxCumulativeDeltaDiffExceeded events raised by the FastPriceFeed contract.
type FastPriceFeedMaxCumulativeDeltaDiffExceededIterator struct {
	Event *FastPriceFeedMaxCumulativeDeltaDiffExceeded // Event containing the contract specifics and raw log

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
func (it *FastPriceFeedMaxCumulativeDeltaDiffExceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeedMaxCumulativeDeltaDiffExceeded)
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
		it.Event = new(FastPriceFeedMaxCumulativeDeltaDiffExceeded)
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
func (it *FastPriceFeedMaxCumulativeDeltaDiffExceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeedMaxCumulativeDeltaDiffExceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeedMaxCumulativeDeltaDiffExceeded represents a MaxCumulativeDeltaDiffExceeded event raised by the FastPriceFeed contract.
type FastPriceFeedMaxCumulativeDeltaDiffExceeded struct {
	Token               common.Address
	RefPrice            *big.Int
	FastPrice           *big.Int
	CumulativeRefDelta  *big.Int
	CumulativeFastDelta *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMaxCumulativeDeltaDiffExceeded is a free log retrieval operation binding the contract event 0xe582322b389ad06b2bbf619cd6da3f16a288ec873ea0fa6df4d72f3d9480b447.
//
// Solidity: event MaxCumulativeDeltaDiffExceeded(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedFilterer) FilterMaxCumulativeDeltaDiffExceeded(opts *bind.FilterOpts) (*FastPriceFeedMaxCumulativeDeltaDiffExceededIterator, error) {

	logs, sub, err := _FastPriceFeed.contract.FilterLogs(opts, "MaxCumulativeDeltaDiffExceeded")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeedMaxCumulativeDeltaDiffExceededIterator{contract: _FastPriceFeed.contract, event: "MaxCumulativeDeltaDiffExceeded", logs: logs, sub: sub}, nil
}

// WatchMaxCumulativeDeltaDiffExceeded is a free log subscription operation binding the contract event 0xe582322b389ad06b2bbf619cd6da3f16a288ec873ea0fa6df4d72f3d9480b447.
//
// Solidity: event MaxCumulativeDeltaDiffExceeded(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedFilterer) WatchMaxCumulativeDeltaDiffExceeded(opts *bind.WatchOpts, sink chan<- *FastPriceFeedMaxCumulativeDeltaDiffExceeded) (event.Subscription, error) {

	logs, sub, err := _FastPriceFeed.contract.WatchLogs(opts, "MaxCumulativeDeltaDiffExceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeedMaxCumulativeDeltaDiffExceeded)
				if err := _FastPriceFeed.contract.UnpackLog(event, "MaxCumulativeDeltaDiffExceeded", log); err != nil {
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

// ParseMaxCumulativeDeltaDiffExceeded is a log parse operation binding the contract event 0xe582322b389ad06b2bbf619cd6da3f16a288ec873ea0fa6df4d72f3d9480b447.
//
// Solidity: event MaxCumulativeDeltaDiffExceeded(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedFilterer) ParseMaxCumulativeDeltaDiffExceeded(log types.Log) (*FastPriceFeedMaxCumulativeDeltaDiffExceeded, error) {
	event := new(FastPriceFeedMaxCumulativeDeltaDiffExceeded)
	if err := _FastPriceFeed.contract.UnpackLog(event, "MaxCumulativeDeltaDiffExceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastPriceFeedPriceDataIterator is returned from FilterPriceData and is used to iterate over the raw logs and unpacked data for PriceData events raised by the FastPriceFeed contract.
type FastPriceFeedPriceDataIterator struct {
	Event *FastPriceFeedPriceData // Event containing the contract specifics and raw log

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
func (it *FastPriceFeedPriceDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeedPriceData)
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
		it.Event = new(FastPriceFeedPriceData)
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
func (it *FastPriceFeedPriceDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeedPriceDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeedPriceData represents a PriceData event raised by the FastPriceFeed contract.
type FastPriceFeedPriceData struct {
	Token               common.Address
	RefPrice            *big.Int
	FastPrice           *big.Int
	CumulativeRefDelta  *big.Int
	CumulativeFastDelta *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPriceData is a free log retrieval operation binding the contract event 0x23b9387f81fca646aac1dc4487ede045c65f5f7445482906565f01e05afdb3a8.
//
// Solidity: event PriceData(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedFilterer) FilterPriceData(opts *bind.FilterOpts) (*FastPriceFeedPriceDataIterator, error) {

	logs, sub, err := _FastPriceFeed.contract.FilterLogs(opts, "PriceData")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeedPriceDataIterator{contract: _FastPriceFeed.contract, event: "PriceData", logs: logs, sub: sub}, nil
}

// WatchPriceData is a free log subscription operation binding the contract event 0x23b9387f81fca646aac1dc4487ede045c65f5f7445482906565f01e05afdb3a8.
//
// Solidity: event PriceData(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedFilterer) WatchPriceData(opts *bind.WatchOpts, sink chan<- *FastPriceFeedPriceData) (event.Subscription, error) {

	logs, sub, err := _FastPriceFeed.contract.WatchLogs(opts, "PriceData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeedPriceData)
				if err := _FastPriceFeed.contract.UnpackLog(event, "PriceData", log); err != nil {
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

// ParsePriceData is a log parse operation binding the contract event 0x23b9387f81fca646aac1dc4487ede045c65f5f7445482906565f01e05afdb3a8.
//
// Solidity: event PriceData(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFeed *FastPriceFeedFilterer) ParsePriceData(log types.Log) (*FastPriceFeedPriceData, error) {
	event := new(FastPriceFeedPriceData)
	if err := _FastPriceFeed.contract.UnpackLog(event, "PriceData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
