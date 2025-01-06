// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockoldrewardrouterv2

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

// MockOldRewardRouterV2MetaData contains all meta data concerning the MockOldRewardRouterV2 contract.
var MockOldRewardRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeGlp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeGmx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeGlp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeGmx\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"acceptTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"batchCompoundForAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchStakeGmxForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"compoundForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_shouldClaimGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeMultiplierPoints\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimWeth\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldConvertWethToEth\",\"type\":\"bool\"}],\"name\":\"handleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inStrictTransferMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_esGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_glpVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_govToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBoostBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"mintAndStakeGlp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"mintAndStakeGlpETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingReceivers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inStrictTransferMode\",\"type\":\"bool\"}],\"name\":\"setInStrictTransferMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxBoostBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setMaxBoostBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMockOldRewardRouterV2.VotingPowerType\",\"name\":\"_votingPowerType\",\"type\":\"uint8\"}],\"name\":\"setVotingPowerType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"signalTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeGmxForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeAndRedeemGlp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeAndRedeemGlpETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstakeEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstakeGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPowerType\",\"outputs\":[{\"internalType\":\"enumMockOldRewardRouterV2.VotingPowerType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MockOldRewardRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockOldRewardRouterV2MetaData.ABI instead.
var MockOldRewardRouterV2ABI = MockOldRewardRouterV2MetaData.ABI

// MockOldRewardRouterV2 is an auto generated Go binding around an Ethereum contract.
type MockOldRewardRouterV2 struct {
	MockOldRewardRouterV2Caller     // Read-only binding to the contract
	MockOldRewardRouterV2Transactor // Write-only binding to the contract
	MockOldRewardRouterV2Filterer   // Log filterer for contract events
}

// MockOldRewardRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockOldRewardRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOldRewardRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockOldRewardRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOldRewardRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockOldRewardRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockOldRewardRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockOldRewardRouterV2Session struct {
	Contract     *MockOldRewardRouterV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockOldRewardRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockOldRewardRouterV2CallerSession struct {
	Contract *MockOldRewardRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// MockOldRewardRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockOldRewardRouterV2TransactorSession struct {
	Contract     *MockOldRewardRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// MockOldRewardRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockOldRewardRouterV2Raw struct {
	Contract *MockOldRewardRouterV2 // Generic contract binding to access the raw methods on
}

// MockOldRewardRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockOldRewardRouterV2CallerRaw struct {
	Contract *MockOldRewardRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// MockOldRewardRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockOldRewardRouterV2TransactorRaw struct {
	Contract *MockOldRewardRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockOldRewardRouterV2 creates a new instance of MockOldRewardRouterV2, bound to a specific deployed contract.
func NewMockOldRewardRouterV2(address common.Address, backend bind.ContractBackend) (*MockOldRewardRouterV2, error) {
	contract, err := bindMockOldRewardRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2{MockOldRewardRouterV2Caller: MockOldRewardRouterV2Caller{contract: contract}, MockOldRewardRouterV2Transactor: MockOldRewardRouterV2Transactor{contract: contract}, MockOldRewardRouterV2Filterer: MockOldRewardRouterV2Filterer{contract: contract}}, nil
}

// NewMockOldRewardRouterV2Caller creates a new read-only instance of MockOldRewardRouterV2, bound to a specific deployed contract.
func NewMockOldRewardRouterV2Caller(address common.Address, caller bind.ContractCaller) (*MockOldRewardRouterV2Caller, error) {
	contract, err := bindMockOldRewardRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2Caller{contract: contract}, nil
}

// NewMockOldRewardRouterV2Transactor creates a new write-only instance of MockOldRewardRouterV2, bound to a specific deployed contract.
func NewMockOldRewardRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MockOldRewardRouterV2Transactor, error) {
	contract, err := bindMockOldRewardRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2Transactor{contract: contract}, nil
}

// NewMockOldRewardRouterV2Filterer creates a new log filterer instance of MockOldRewardRouterV2, bound to a specific deployed contract.
func NewMockOldRewardRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MockOldRewardRouterV2Filterer, error) {
	contract, err := bindMockOldRewardRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2Filterer{contract: contract}, nil
}

// bindMockOldRewardRouterV2 binds a generic wrapper to an already deployed contract.
func bindMockOldRewardRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockOldRewardRouterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOldRewardRouterV2.Contract.MockOldRewardRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.MockOldRewardRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.MockOldRewardRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockOldRewardRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _MockOldRewardRouterV2.Contract.BASISPOINTSDIVISOR(&_MockOldRewardRouterV2.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _MockOldRewardRouterV2.Contract.BASISPOINTSDIVISOR(&_MockOldRewardRouterV2.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) BnGmx() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.BnGmx(&_MockOldRewardRouterV2.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) BnGmx() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.BnGmx(&_MockOldRewardRouterV2.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) BonusGmxTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.BonusGmxTracker(&_MockOldRewardRouterV2.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) BonusGmxTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.BonusGmxTracker(&_MockOldRewardRouterV2.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) EsGmx() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.EsGmx(&_MockOldRewardRouterV2.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) EsGmx() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.EsGmx(&_MockOldRewardRouterV2.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) FeeGlpTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.FeeGlpTracker(&_MockOldRewardRouterV2.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) FeeGlpTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.FeeGlpTracker(&_MockOldRewardRouterV2.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) FeeGmxTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.FeeGmxTracker(&_MockOldRewardRouterV2.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) FeeGmxTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.FeeGmxTracker(&_MockOldRewardRouterV2.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Glp() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Glp(&_MockOldRewardRouterV2.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) Glp() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Glp(&_MockOldRewardRouterV2.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) GlpManager() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GlpManager(&_MockOldRewardRouterV2.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) GlpManager() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GlpManager(&_MockOldRewardRouterV2.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) GlpVester() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GlpVester(&_MockOldRewardRouterV2.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) GlpVester() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GlpVester(&_MockOldRewardRouterV2.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) Gmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "gmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Gmx() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Gmx(&_MockOldRewardRouterV2.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) Gmx() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Gmx(&_MockOldRewardRouterV2.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) GmxVester() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GmxVester(&_MockOldRewardRouterV2.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) GmxVester() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GmxVester(&_MockOldRewardRouterV2.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Gov() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Gov(&_MockOldRewardRouterV2.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) Gov() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Gov(&_MockOldRewardRouterV2.CallOpts)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) GovToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "govToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) GovToken() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GovToken(&_MockOldRewardRouterV2.CallOpts)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) GovToken() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.GovToken(&_MockOldRewardRouterV2.CallOpts)
}

// InStrictTransferMode is a free data retrieval call binding the contract method 0x743655ab.
//
// Solidity: function inStrictTransferMode() view returns(bool)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) InStrictTransferMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "inStrictTransferMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InStrictTransferMode is a free data retrieval call binding the contract method 0x743655ab.
//
// Solidity: function inStrictTransferMode() view returns(bool)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) InStrictTransferMode() (bool, error) {
	return _MockOldRewardRouterV2.Contract.InStrictTransferMode(&_MockOldRewardRouterV2.CallOpts)
}

// InStrictTransferMode is a free data retrieval call binding the contract method 0x743655ab.
//
// Solidity: function inStrictTransferMode() view returns(bool)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) InStrictTransferMode() (bool, error) {
	return _MockOldRewardRouterV2.Contract.InStrictTransferMode(&_MockOldRewardRouterV2.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) IsInitialized() (bool, error) {
	return _MockOldRewardRouterV2.Contract.IsInitialized(&_MockOldRewardRouterV2.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) IsInitialized() (bool, error) {
	return _MockOldRewardRouterV2.Contract.IsInitialized(&_MockOldRewardRouterV2.CallOpts)
}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) MaxBoostBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "maxBoostBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) MaxBoostBasisPoints() (*big.Int, error) {
	return _MockOldRewardRouterV2.Contract.MaxBoostBasisPoints(&_MockOldRewardRouterV2.CallOpts)
}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) MaxBoostBasisPoints() (*big.Int, error) {
	return _MockOldRewardRouterV2.Contract.MaxBoostBasisPoints(&_MockOldRewardRouterV2.CallOpts)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) PendingReceivers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "pendingReceivers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.PendingReceivers(&_MockOldRewardRouterV2.CallOpts, arg0)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.PendingReceivers(&_MockOldRewardRouterV2.CallOpts, arg0)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) StakedGlpTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.StakedGlpTracker(&_MockOldRewardRouterV2.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) StakedGlpTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.StakedGlpTracker(&_MockOldRewardRouterV2.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) StakedGmxTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.StakedGmxTracker(&_MockOldRewardRouterV2.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) StakedGmxTracker() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.StakedGmxTracker(&_MockOldRewardRouterV2.CallOpts)
}

// VotingPowerType is a free data retrieval call binding the contract method 0x08694e63.
//
// Solidity: function votingPowerType() view returns(uint8)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) VotingPowerType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "votingPowerType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VotingPowerType is a free data retrieval call binding the contract method 0x08694e63.
//
// Solidity: function votingPowerType() view returns(uint8)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) VotingPowerType() (uint8, error) {
	return _MockOldRewardRouterV2.Contract.VotingPowerType(&_MockOldRewardRouterV2.CallOpts)
}

// VotingPowerType is a free data retrieval call binding the contract method 0x08694e63.
//
// Solidity: function votingPowerType() view returns(uint8)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) VotingPowerType() (uint8, error) {
	return _MockOldRewardRouterV2.Contract.VotingPowerType(&_MockOldRewardRouterV2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockOldRewardRouterV2.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Weth() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Weth(&_MockOldRewardRouterV2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2CallerSession) Weth() (common.Address, error) {
	return _MockOldRewardRouterV2.Contract.Weth(&_MockOldRewardRouterV2.CallOpts)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) AcceptTransfer(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "acceptTransfer", _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.AcceptTransfer(&_MockOldRewardRouterV2.TransactOpts, _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.AcceptTransfer(&_MockOldRewardRouterV2.TransactOpts, _sender)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) BatchCompoundForAccounts(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "batchCompoundForAccounts", _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.BatchCompoundForAccounts(&_MockOldRewardRouterV2.TransactOpts, _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.BatchCompoundForAccounts(&_MockOldRewardRouterV2.TransactOpts, _accounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) BatchStakeGmxForAccount(opts *bind.TransactOpts, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "batchStakeGmxForAccount", _accounts, _amounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) BatchStakeGmxForAccount(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.BatchStakeGmxForAccount(&_MockOldRewardRouterV2.TransactOpts, _accounts, _amounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) BatchStakeGmxForAccount(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.BatchStakeGmxForAccount(&_MockOldRewardRouterV2.TransactOpts, _accounts, _amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Claim() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Claim(&_MockOldRewardRouterV2.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) Claim() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Claim(&_MockOldRewardRouterV2.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) ClaimEsGmx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "claimEsGmx")
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) ClaimEsGmx() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.ClaimEsGmx(&_MockOldRewardRouterV2.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) ClaimEsGmx() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.ClaimEsGmx(&_MockOldRewardRouterV2.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) ClaimFees() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.ClaimFees(&_MockOldRewardRouterV2.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) ClaimFees() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.ClaimFees(&_MockOldRewardRouterV2.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) Compound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "compound")
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Compound() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Compound(&_MockOldRewardRouterV2.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) Compound() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Compound(&_MockOldRewardRouterV2.TransactOpts)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) CompoundForAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "compoundForAccount", _account)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) CompoundForAccount(_account common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.CompoundForAccount(&_MockOldRewardRouterV2.TransactOpts, _account)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) CompoundForAccount(_account common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.CompoundForAccount(&_MockOldRewardRouterV2.TransactOpts, _account)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) HandleRewards(opts *bind.TransactOpts, _shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "handleRewards", _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.HandleRewards(&_MockOldRewardRouterV2.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.HandleRewards(&_MockOldRewardRouterV2.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// Initialize is a paid mutator transaction binding the contract method 0x71302627.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester, address _govToken) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) Initialize(opts *bind.TransactOpts, _weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address, _govToken common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "initialize", _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester, _govToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x71302627.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester, address _govToken) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Initialize(_weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address, _govToken common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Initialize(&_MockOldRewardRouterV2.TransactOpts, _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester, _govToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x71302627.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester, address _govToken) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) Initialize(_weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address, _govToken common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Initialize(&_MockOldRewardRouterV2.TransactOpts, _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester, _govToken)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) MintAndStakeGlp(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "mintAndStakeGlp", _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.MintAndStakeGlp(&_MockOldRewardRouterV2.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.MintAndStakeGlp(&_MockOldRewardRouterV2.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) MintAndStakeGlpETH(opts *bind.TransactOpts, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "mintAndStakeGlpETH", _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.MintAndStakeGlpETH(&_MockOldRewardRouterV2.TransactOpts, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.MintAndStakeGlpETH(&_MockOldRewardRouterV2.TransactOpts, _minUsdg, _minGlp)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetGov(&_MockOldRewardRouterV2.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetGov(&_MockOldRewardRouterV2.TransactOpts, _gov)
}

// SetInStrictTransferMode is a paid mutator transaction binding the contract method 0x65f8f4d0.
//
// Solidity: function setInStrictTransferMode(bool _inStrictTransferMode) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) SetInStrictTransferMode(opts *bind.TransactOpts, _inStrictTransferMode bool) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "setInStrictTransferMode", _inStrictTransferMode)
}

// SetInStrictTransferMode is a paid mutator transaction binding the contract method 0x65f8f4d0.
//
// Solidity: function setInStrictTransferMode(bool _inStrictTransferMode) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) SetInStrictTransferMode(_inStrictTransferMode bool) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetInStrictTransferMode(&_MockOldRewardRouterV2.TransactOpts, _inStrictTransferMode)
}

// SetInStrictTransferMode is a paid mutator transaction binding the contract method 0x65f8f4d0.
//
// Solidity: function setInStrictTransferMode(bool _inStrictTransferMode) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) SetInStrictTransferMode(_inStrictTransferMode bool) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetInStrictTransferMode(&_MockOldRewardRouterV2.TransactOpts, _inStrictTransferMode)
}

// SetMaxBoostBasisPoints is a paid mutator transaction binding the contract method 0x35c1b188.
//
// Solidity: function setMaxBoostBasisPoints(uint256 _maxBoostBasisPoints) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) SetMaxBoostBasisPoints(opts *bind.TransactOpts, _maxBoostBasisPoints *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "setMaxBoostBasisPoints", _maxBoostBasisPoints)
}

// SetMaxBoostBasisPoints is a paid mutator transaction binding the contract method 0x35c1b188.
//
// Solidity: function setMaxBoostBasisPoints(uint256 _maxBoostBasisPoints) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) SetMaxBoostBasisPoints(_maxBoostBasisPoints *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetMaxBoostBasisPoints(&_MockOldRewardRouterV2.TransactOpts, _maxBoostBasisPoints)
}

// SetMaxBoostBasisPoints is a paid mutator transaction binding the contract method 0x35c1b188.
//
// Solidity: function setMaxBoostBasisPoints(uint256 _maxBoostBasisPoints) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) SetMaxBoostBasisPoints(_maxBoostBasisPoints *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetMaxBoostBasisPoints(&_MockOldRewardRouterV2.TransactOpts, _maxBoostBasisPoints)
}

// SetVotingPowerType is a paid mutator transaction binding the contract method 0xba2db4f5.
//
// Solidity: function setVotingPowerType(uint8 _votingPowerType) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) SetVotingPowerType(opts *bind.TransactOpts, _votingPowerType uint8) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "setVotingPowerType", _votingPowerType)
}

// SetVotingPowerType is a paid mutator transaction binding the contract method 0xba2db4f5.
//
// Solidity: function setVotingPowerType(uint8 _votingPowerType) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) SetVotingPowerType(_votingPowerType uint8) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetVotingPowerType(&_MockOldRewardRouterV2.TransactOpts, _votingPowerType)
}

// SetVotingPowerType is a paid mutator transaction binding the contract method 0xba2db4f5.
//
// Solidity: function setVotingPowerType(uint8 _votingPowerType) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) SetVotingPowerType(_votingPowerType uint8) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SetVotingPowerType(&_MockOldRewardRouterV2.TransactOpts, _votingPowerType)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) SignalTransfer(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "signalTransfer", _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SignalTransfer(&_MockOldRewardRouterV2.TransactOpts, _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.SignalTransfer(&_MockOldRewardRouterV2.TransactOpts, _receiver)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) StakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "stakeEsGmx", _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.StakeEsGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.StakeEsGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) StakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "stakeGmx", _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.StakeGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.StakeGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) StakeGmxForAccount(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "stakeGmxForAccount", _account, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) StakeGmxForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.StakeGmxForAccount(&_MockOldRewardRouterV2.TransactOpts, _account, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) StakeGmxForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.StakeGmxForAccount(&_MockOldRewardRouterV2.TransactOpts, _account, _amount)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) UnstakeAndRedeemGlp(opts *bind.TransactOpts, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "unstakeAndRedeemGlp", _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeAndRedeemGlp(&_MockOldRewardRouterV2.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeAndRedeemGlp(&_MockOldRewardRouterV2.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) UnstakeAndRedeemGlpETH(opts *bind.TransactOpts, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "unstakeAndRedeemGlpETH", _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeAndRedeemGlpETH(&_MockOldRewardRouterV2.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeAndRedeemGlpETH(&_MockOldRewardRouterV2.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) UnstakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "unstakeEsGmx", _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeEsGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeEsGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) UnstakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "unstakeGmx", _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.UnstakeGmx(&_MockOldRewardRouterV2.TransactOpts, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.WithdrawToken(&_MockOldRewardRouterV2.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.WithdrawToken(&_MockOldRewardRouterV2.TransactOpts, _token, _account, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockOldRewardRouterV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Session) Receive() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Receive(&_MockOldRewardRouterV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2TransactorSession) Receive() (*types.Transaction, error) {
	return _MockOldRewardRouterV2.Contract.Receive(&_MockOldRewardRouterV2.TransactOpts)
}

// MockOldRewardRouterV2StakeGlpIterator is returned from FilterStakeGlp and is used to iterate over the raw logs and unpacked data for StakeGlp events raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2StakeGlpIterator struct {
	Event *MockOldRewardRouterV2StakeGlp // Event containing the contract specifics and raw log

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
func (it *MockOldRewardRouterV2StakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOldRewardRouterV2StakeGlp)
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
		it.Event = new(MockOldRewardRouterV2StakeGlp)
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
func (it *MockOldRewardRouterV2StakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOldRewardRouterV2StakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOldRewardRouterV2StakeGlp represents a StakeGlp event raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2StakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGlp is a free log retrieval operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) FilterStakeGlp(opts *bind.FilterOpts) (*MockOldRewardRouterV2StakeGlpIterator, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.FilterLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2StakeGlpIterator{contract: _MockOldRewardRouterV2.contract, event: "StakeGlp", logs: logs, sub: sub}, nil
}

// WatchStakeGlp is a free log subscription operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) WatchStakeGlp(opts *bind.WatchOpts, sink chan<- *MockOldRewardRouterV2StakeGlp) (event.Subscription, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.WatchLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOldRewardRouterV2StakeGlp)
				if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "StakeGlp", log); err != nil {
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

// ParseStakeGlp is a log parse operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) ParseStakeGlp(log types.Log) (*MockOldRewardRouterV2StakeGlp, error) {
	event := new(MockOldRewardRouterV2StakeGlp)
	if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "StakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockOldRewardRouterV2StakeGmxIterator is returned from FilterStakeGmx and is used to iterate over the raw logs and unpacked data for StakeGmx events raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2StakeGmxIterator struct {
	Event *MockOldRewardRouterV2StakeGmx // Event containing the contract specifics and raw log

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
func (it *MockOldRewardRouterV2StakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOldRewardRouterV2StakeGmx)
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
		it.Event = new(MockOldRewardRouterV2StakeGmx)
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
func (it *MockOldRewardRouterV2StakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOldRewardRouterV2StakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOldRewardRouterV2StakeGmx represents a StakeGmx event raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2StakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGmx is a free log retrieval operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) FilterStakeGmx(opts *bind.FilterOpts) (*MockOldRewardRouterV2StakeGmxIterator, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.FilterLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2StakeGmxIterator{contract: _MockOldRewardRouterV2.contract, event: "StakeGmx", logs: logs, sub: sub}, nil
}

// WatchStakeGmx is a free log subscription operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) WatchStakeGmx(opts *bind.WatchOpts, sink chan<- *MockOldRewardRouterV2StakeGmx) (event.Subscription, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.WatchLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOldRewardRouterV2StakeGmx)
				if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "StakeGmx", log); err != nil {
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

// ParseStakeGmx is a log parse operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) ParseStakeGmx(log types.Log) (*MockOldRewardRouterV2StakeGmx, error) {
	event := new(MockOldRewardRouterV2StakeGmx)
	if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "StakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockOldRewardRouterV2UnstakeGlpIterator is returned from FilterUnstakeGlp and is used to iterate over the raw logs and unpacked data for UnstakeGlp events raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2UnstakeGlpIterator struct {
	Event *MockOldRewardRouterV2UnstakeGlp // Event containing the contract specifics and raw log

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
func (it *MockOldRewardRouterV2UnstakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOldRewardRouterV2UnstakeGlp)
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
		it.Event = new(MockOldRewardRouterV2UnstakeGlp)
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
func (it *MockOldRewardRouterV2UnstakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOldRewardRouterV2UnstakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOldRewardRouterV2UnstakeGlp represents a UnstakeGlp event raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2UnstakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGlp is a free log retrieval operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) FilterUnstakeGlp(opts *bind.FilterOpts) (*MockOldRewardRouterV2UnstakeGlpIterator, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.FilterLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2UnstakeGlpIterator{contract: _MockOldRewardRouterV2.contract, event: "UnstakeGlp", logs: logs, sub: sub}, nil
}

// WatchUnstakeGlp is a free log subscription operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) WatchUnstakeGlp(opts *bind.WatchOpts, sink chan<- *MockOldRewardRouterV2UnstakeGlp) (event.Subscription, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.WatchLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOldRewardRouterV2UnstakeGlp)
				if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
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

// ParseUnstakeGlp is a log parse operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) ParseUnstakeGlp(log types.Log) (*MockOldRewardRouterV2UnstakeGlp, error) {
	event := new(MockOldRewardRouterV2UnstakeGlp)
	if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockOldRewardRouterV2UnstakeGmxIterator is returned from FilterUnstakeGmx and is used to iterate over the raw logs and unpacked data for UnstakeGmx events raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2UnstakeGmxIterator struct {
	Event *MockOldRewardRouterV2UnstakeGmx // Event containing the contract specifics and raw log

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
func (it *MockOldRewardRouterV2UnstakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockOldRewardRouterV2UnstakeGmx)
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
		it.Event = new(MockOldRewardRouterV2UnstakeGmx)
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
func (it *MockOldRewardRouterV2UnstakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockOldRewardRouterV2UnstakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockOldRewardRouterV2UnstakeGmx represents a UnstakeGmx event raised by the MockOldRewardRouterV2 contract.
type MockOldRewardRouterV2UnstakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGmx is a free log retrieval operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) FilterUnstakeGmx(opts *bind.FilterOpts) (*MockOldRewardRouterV2UnstakeGmxIterator, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.FilterLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return &MockOldRewardRouterV2UnstakeGmxIterator{contract: _MockOldRewardRouterV2.contract, event: "UnstakeGmx", logs: logs, sub: sub}, nil
}

// WatchUnstakeGmx is a free log subscription operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) WatchUnstakeGmx(opts *bind.WatchOpts, sink chan<- *MockOldRewardRouterV2UnstakeGmx) (event.Subscription, error) {

	logs, sub, err := _MockOldRewardRouterV2.contract.WatchLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockOldRewardRouterV2UnstakeGmx)
				if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
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

// ParseUnstakeGmx is a log parse operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_MockOldRewardRouterV2 *MockOldRewardRouterV2Filterer) ParseUnstakeGmx(log types.Log) (*MockOldRewardRouterV2UnstakeGmx, error) {
	event := new(MockOldRewardRouterV2UnstakeGmx)
	if err := _MockOldRewardRouterV2.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
