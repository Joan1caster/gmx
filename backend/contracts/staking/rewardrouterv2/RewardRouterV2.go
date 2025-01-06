// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewardrouterv2

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

// RewardRouterV2InitializeParams is an auto generated low-level Go binding around an user-defined struct.
type RewardRouterV2InitializeParams struct {
	Weth               common.Address
	Gmx                common.Address
	EsGmx              common.Address
	BnGmx              common.Address
	Glp                common.Address
	StakedGmxTracker   common.Address
	BonusGmxTracker    common.Address
	ExtendedGmxTracker common.Address
	FeeGmxTracker      common.Address
	FeeGlpTracker      common.Address
	StakedGlpTracker   common.Address
	GlpManager         common.Address
	GmxVester          common.Address
	GlpVester          common.Address
	ExternalHandler    common.Address
	GovToken           common.Address
}

// RewardRouterV2MetaData contains all meta data concerning the RewardRouterV2 contract.
var RewardRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeGlp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeGmx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeGlp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeGmx\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"acceptTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"batchCompoundForAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"batchRestakeForAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchStakeGmxForAccounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esGmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extendedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"externalHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmxVester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_shouldClaimGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeMultiplierPoints\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimWeth\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldConvertWethToEth\",\"type\":\"bool\"}],\"name\":\"handleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gmxReceiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeEsGmx\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldStakeMultiplierPoints\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldClaimWeth\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_shouldConvertWethToEth\",\"type\":\"bool\"}],\"name\":\"handleRewardsV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRestakingMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inStrictTransferMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"esGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bnGmx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"glp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bonusGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"extendedGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeGmxTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"glpManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gmxVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"glpVester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"externalHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"govToken\",\"type\":\"address\"}],\"internalType\":\"structRewardRouterV2.InitializeParams\",\"name\":\"_initializeParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"externalCallTargets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"externalCallDataList\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"refundTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"refundReceivers\",\"type\":\"address[]\"}],\"name\":\"makeExternalCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBoostBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"mintAndStakeGlp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"mintAndStakeGlpETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingReceivers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_govToken\",\"type\":\"address\"}],\"name\":\"setGovToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inRestakingMode\",\"type\":\"bool\"}],\"name\":\"setInRestakingMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inStrictTransferMode\",\"type\":\"bool\"}],\"name\":\"setInStrictTransferMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxBoostBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setMaxBoostBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRewardRouterV2.VotingPowerType\",\"name\":\"_votingPowerType\",\"type\":\"uint8\"}],\"name\":\"setVotingPowerType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"signalTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGmxTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeAndRedeemGlp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeAndRedeemGlpETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstakeEsGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstakeGmx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPowerType\",\"outputs\":[{\"internalType\":\"enumRewardRouterV2.VotingPowerType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RewardRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardRouterV2MetaData.ABI instead.
var RewardRouterV2ABI = RewardRouterV2MetaData.ABI

// RewardRouterV2 is an auto generated Go binding around an Ethereum contract.
type RewardRouterV2 struct {
	RewardRouterV2Caller     // Read-only binding to the contract
	RewardRouterV2Transactor // Write-only binding to the contract
	RewardRouterV2Filterer   // Log filterer for contract events
}

// RewardRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type RewardRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardRouterV2Session struct {
	Contract     *RewardRouterV2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardRouterV2CallerSession struct {
	Contract *RewardRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RewardRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardRouterV2TransactorSession struct {
	Contract     *RewardRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RewardRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type RewardRouterV2Raw struct {
	Contract *RewardRouterV2 // Generic contract binding to access the raw methods on
}

// RewardRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardRouterV2CallerRaw struct {
	Contract *RewardRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// RewardRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardRouterV2TransactorRaw struct {
	Contract *RewardRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardRouterV2 creates a new instance of RewardRouterV2, bound to a specific deployed contract.
func NewRewardRouterV2(address common.Address, backend bind.ContractBackend) (*RewardRouterV2, error) {
	contract, err := bindRewardRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2{RewardRouterV2Caller: RewardRouterV2Caller{contract: contract}, RewardRouterV2Transactor: RewardRouterV2Transactor{contract: contract}, RewardRouterV2Filterer: RewardRouterV2Filterer{contract: contract}}, nil
}

// NewRewardRouterV2Caller creates a new read-only instance of RewardRouterV2, bound to a specific deployed contract.
func NewRewardRouterV2Caller(address common.Address, caller bind.ContractCaller) (*RewardRouterV2Caller, error) {
	contract, err := bindRewardRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2Caller{contract: contract}, nil
}

// NewRewardRouterV2Transactor creates a new write-only instance of RewardRouterV2, bound to a specific deployed contract.
func NewRewardRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*RewardRouterV2Transactor, error) {
	contract, err := bindRewardRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2Transactor{contract: contract}, nil
}

// NewRewardRouterV2Filterer creates a new log filterer instance of RewardRouterV2, bound to a specific deployed contract.
func NewRewardRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*RewardRouterV2Filterer, error) {
	contract, err := bindRewardRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2Filterer{contract: contract}, nil
}

// bindRewardRouterV2 binds a generic wrapper to an already deployed contract.
func bindRewardRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardRouterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardRouterV2 *RewardRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardRouterV2.Contract.RewardRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardRouterV2 *RewardRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.RewardRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardRouterV2 *RewardRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.RewardRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardRouterV2 *RewardRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardRouterV2 *RewardRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardRouterV2 *RewardRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Caller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Session) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _RewardRouterV2.Contract.BASISPOINTSDIVISOR(&_RewardRouterV2.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_RewardRouterV2 *RewardRouterV2CallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _RewardRouterV2.Contract.BASISPOINTSDIVISOR(&_RewardRouterV2.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) BnGmx() (common.Address, error) {
	return _RewardRouterV2.Contract.BnGmx(&_RewardRouterV2.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) BnGmx() (common.Address, error) {
	return _RewardRouterV2.Contract.BnGmx(&_RewardRouterV2.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) BonusGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.BonusGmxTracker(&_RewardRouterV2.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) BonusGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.BonusGmxTracker(&_RewardRouterV2.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) EsGmx() (common.Address, error) {
	return _RewardRouterV2.Contract.EsGmx(&_RewardRouterV2.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) EsGmx() (common.Address, error) {
	return _RewardRouterV2.Contract.EsGmx(&_RewardRouterV2.CallOpts)
}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) ExtendedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "extendedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) ExtendedGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.ExtendedGmxTracker(&_RewardRouterV2.CallOpts)
}

// ExtendedGmxTracker is a free data retrieval call binding the contract method 0x204b4c6f.
//
// Solidity: function extendedGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) ExtendedGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.ExtendedGmxTracker(&_RewardRouterV2.CallOpts)
}

// ExternalHandler is a free data retrieval call binding the contract method 0x2e944bd6.
//
// Solidity: function externalHandler() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) ExternalHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "externalHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExternalHandler is a free data retrieval call binding the contract method 0x2e944bd6.
//
// Solidity: function externalHandler() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) ExternalHandler() (common.Address, error) {
	return _RewardRouterV2.Contract.ExternalHandler(&_RewardRouterV2.CallOpts)
}

// ExternalHandler is a free data retrieval call binding the contract method 0x2e944bd6.
//
// Solidity: function externalHandler() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) ExternalHandler() (common.Address, error) {
	return _RewardRouterV2.Contract.ExternalHandler(&_RewardRouterV2.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) FeeGlpTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.FeeGlpTracker(&_RewardRouterV2.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) FeeGlpTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.FeeGlpTracker(&_RewardRouterV2.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) FeeGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.FeeGmxTracker(&_RewardRouterV2.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) FeeGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.FeeGmxTracker(&_RewardRouterV2.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) Glp() (common.Address, error) {
	return _RewardRouterV2.Contract.Glp(&_RewardRouterV2.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) Glp() (common.Address, error) {
	return _RewardRouterV2.Contract.Glp(&_RewardRouterV2.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) GlpManager() (common.Address, error) {
	return _RewardRouterV2.Contract.GlpManager(&_RewardRouterV2.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) GlpManager() (common.Address, error) {
	return _RewardRouterV2.Contract.GlpManager(&_RewardRouterV2.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) GlpVester() (common.Address, error) {
	return _RewardRouterV2.Contract.GlpVester(&_RewardRouterV2.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) GlpVester() (common.Address, error) {
	return _RewardRouterV2.Contract.GlpVester(&_RewardRouterV2.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) Gmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "gmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) Gmx() (common.Address, error) {
	return _RewardRouterV2.Contract.Gmx(&_RewardRouterV2.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) Gmx() (common.Address, error) {
	return _RewardRouterV2.Contract.Gmx(&_RewardRouterV2.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) GmxVester() (common.Address, error) {
	return _RewardRouterV2.Contract.GmxVester(&_RewardRouterV2.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) GmxVester() (common.Address, error) {
	return _RewardRouterV2.Contract.GmxVester(&_RewardRouterV2.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) Gov() (common.Address, error) {
	return _RewardRouterV2.Contract.Gov(&_RewardRouterV2.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) Gov() (common.Address, error) {
	return _RewardRouterV2.Contract.Gov(&_RewardRouterV2.CallOpts)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) GovToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "govToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) GovToken() (common.Address, error) {
	return _RewardRouterV2.Contract.GovToken(&_RewardRouterV2.CallOpts)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) GovToken() (common.Address, error) {
	return _RewardRouterV2.Contract.GovToken(&_RewardRouterV2.CallOpts)
}

// InRestakingMode is a free data retrieval call binding the contract method 0x82b6aecc.
//
// Solidity: function inRestakingMode() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2Caller) InRestakingMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "inRestakingMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRestakingMode is a free data retrieval call binding the contract method 0x82b6aecc.
//
// Solidity: function inRestakingMode() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2Session) InRestakingMode() (bool, error) {
	return _RewardRouterV2.Contract.InRestakingMode(&_RewardRouterV2.CallOpts)
}

// InRestakingMode is a free data retrieval call binding the contract method 0x82b6aecc.
//
// Solidity: function inRestakingMode() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2CallerSession) InRestakingMode() (bool, error) {
	return _RewardRouterV2.Contract.InRestakingMode(&_RewardRouterV2.CallOpts)
}

// InStrictTransferMode is a free data retrieval call binding the contract method 0x743655ab.
//
// Solidity: function inStrictTransferMode() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2Caller) InStrictTransferMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "inStrictTransferMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InStrictTransferMode is a free data retrieval call binding the contract method 0x743655ab.
//
// Solidity: function inStrictTransferMode() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2Session) InStrictTransferMode() (bool, error) {
	return _RewardRouterV2.Contract.InStrictTransferMode(&_RewardRouterV2.CallOpts)
}

// InStrictTransferMode is a free data retrieval call binding the contract method 0x743655ab.
//
// Solidity: function inStrictTransferMode() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2CallerSession) InStrictTransferMode() (bool, error) {
	return _RewardRouterV2.Contract.InStrictTransferMode(&_RewardRouterV2.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2Caller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2Session) IsInitialized() (bool, error) {
	return _RewardRouterV2.Contract.IsInitialized(&_RewardRouterV2.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardRouterV2 *RewardRouterV2CallerSession) IsInitialized() (bool, error) {
	return _RewardRouterV2.Contract.IsInitialized(&_RewardRouterV2.CallOpts)
}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Caller) MaxBoostBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "maxBoostBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Session) MaxBoostBasisPoints() (*big.Int, error) {
	return _RewardRouterV2.Contract.MaxBoostBasisPoints(&_RewardRouterV2.CallOpts)
}

// MaxBoostBasisPoints is a free data retrieval call binding the contract method 0xdddfa314.
//
// Solidity: function maxBoostBasisPoints() view returns(uint256)
func (_RewardRouterV2 *RewardRouterV2CallerSession) MaxBoostBasisPoints() (*big.Int, error) {
	return _RewardRouterV2.Contract.MaxBoostBasisPoints(&_RewardRouterV2.CallOpts)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) PendingReceivers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "pendingReceivers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _RewardRouterV2.Contract.PendingReceivers(&_RewardRouterV2.CallOpts, arg0)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _RewardRouterV2.Contract.PendingReceivers(&_RewardRouterV2.CallOpts, arg0)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) StakedGlpTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.StakedGlpTracker(&_RewardRouterV2.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) StakedGlpTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.StakedGlpTracker(&_RewardRouterV2.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) StakedGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.StakedGmxTracker(&_RewardRouterV2.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) StakedGmxTracker() (common.Address, error) {
	return _RewardRouterV2.Contract.StakedGmxTracker(&_RewardRouterV2.CallOpts)
}

// VotingPowerType is a free data retrieval call binding the contract method 0x08694e63.
//
// Solidity: function votingPowerType() view returns(uint8)
func (_RewardRouterV2 *RewardRouterV2Caller) VotingPowerType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "votingPowerType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VotingPowerType is a free data retrieval call binding the contract method 0x08694e63.
//
// Solidity: function votingPowerType() view returns(uint8)
func (_RewardRouterV2 *RewardRouterV2Session) VotingPowerType() (uint8, error) {
	return _RewardRouterV2.Contract.VotingPowerType(&_RewardRouterV2.CallOpts)
}

// VotingPowerType is a free data retrieval call binding the contract method 0x08694e63.
//
// Solidity: function votingPowerType() view returns(uint8)
func (_RewardRouterV2 *RewardRouterV2CallerSession) VotingPowerType() (uint8, error) {
	return _RewardRouterV2.Contract.VotingPowerType(&_RewardRouterV2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouterV2.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_RewardRouterV2 *RewardRouterV2Session) Weth() (common.Address, error) {
	return _RewardRouterV2.Contract.Weth(&_RewardRouterV2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_RewardRouterV2 *RewardRouterV2CallerSession) Weth() (common.Address, error) {
	return _RewardRouterV2.Contract.Weth(&_RewardRouterV2.CallOpts)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) AcceptTransfer(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "acceptTransfer", _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_RewardRouterV2 *RewardRouterV2Session) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.AcceptTransfer(&_RewardRouterV2.TransactOpts, _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.AcceptTransfer(&_RewardRouterV2.TransactOpts, _sender)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) BatchCompoundForAccounts(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "batchCompoundForAccounts", _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_RewardRouterV2 *RewardRouterV2Session) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.BatchCompoundForAccounts(&_RewardRouterV2.TransactOpts, _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.BatchCompoundForAccounts(&_RewardRouterV2.TransactOpts, _accounts)
}

// BatchRestakeForAccounts is a paid mutator transaction binding the contract method 0x44d3a201.
//
// Solidity: function batchRestakeForAccounts(address[] _accounts) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) BatchRestakeForAccounts(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "batchRestakeForAccounts", _accounts)
}

// BatchRestakeForAccounts is a paid mutator transaction binding the contract method 0x44d3a201.
//
// Solidity: function batchRestakeForAccounts(address[] _accounts) returns()
func (_RewardRouterV2 *RewardRouterV2Session) BatchRestakeForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.BatchRestakeForAccounts(&_RewardRouterV2.TransactOpts, _accounts)
}

// BatchRestakeForAccounts is a paid mutator transaction binding the contract method 0x44d3a201.
//
// Solidity: function batchRestakeForAccounts(address[] _accounts) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) BatchRestakeForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.BatchRestakeForAccounts(&_RewardRouterV2.TransactOpts, _accounts)
}

// BatchStakeGmxForAccounts is a paid mutator transaction binding the contract method 0x3172fc23.
//
// Solidity: function batchStakeGmxForAccounts(address[] _accounts, uint256[] _amounts) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) BatchStakeGmxForAccounts(opts *bind.TransactOpts, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "batchStakeGmxForAccounts", _accounts, _amounts)
}

// BatchStakeGmxForAccounts is a paid mutator transaction binding the contract method 0x3172fc23.
//
// Solidity: function batchStakeGmxForAccounts(address[] _accounts, uint256[] _amounts) returns()
func (_RewardRouterV2 *RewardRouterV2Session) BatchStakeGmxForAccounts(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.BatchStakeGmxForAccounts(&_RewardRouterV2.TransactOpts, _accounts, _amounts)
}

// BatchStakeGmxForAccounts is a paid mutator transaction binding the contract method 0x3172fc23.
//
// Solidity: function batchStakeGmxForAccounts(address[] _accounts, uint256[] _amounts) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) BatchStakeGmxForAccounts(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.BatchStakeGmxForAccounts(&_RewardRouterV2.TransactOpts, _accounts, _amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_RewardRouterV2 *RewardRouterV2Session) Claim() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Claim(&_RewardRouterV2.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) Claim() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Claim(&_RewardRouterV2.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) ClaimEsGmx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "claimEsGmx")
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_RewardRouterV2 *RewardRouterV2Session) ClaimEsGmx() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.ClaimEsGmx(&_RewardRouterV2.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) ClaimEsGmx() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.ClaimEsGmx(&_RewardRouterV2.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_RewardRouterV2 *RewardRouterV2Session) ClaimFees() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.ClaimFees(&_RewardRouterV2.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) ClaimFees() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.ClaimFees(&_RewardRouterV2.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) Compound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "compound")
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_RewardRouterV2 *RewardRouterV2Session) Compound() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Compound(&_RewardRouterV2.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) Compound() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Compound(&_RewardRouterV2.TransactOpts)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) HandleRewards(opts *bind.TransactOpts, _shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "handleRewards", _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouterV2 *RewardRouterV2Session) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.HandleRewards(&_RewardRouterV2.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.HandleRewards(&_RewardRouterV2.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewardsV2 is a paid mutator transaction binding the contract method 0x1cf631bb.
//
// Solidity: function handleRewardsV2(address _gmxReceiver, bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) HandleRewardsV2(opts *bind.TransactOpts, _gmxReceiver common.Address, _shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "handleRewardsV2", _gmxReceiver, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewardsV2 is a paid mutator transaction binding the contract method 0x1cf631bb.
//
// Solidity: function handleRewardsV2(address _gmxReceiver, bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouterV2 *RewardRouterV2Session) HandleRewardsV2(_gmxReceiver common.Address, _shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.HandleRewardsV2(&_RewardRouterV2.TransactOpts, _gmxReceiver, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewardsV2 is a paid mutator transaction binding the contract method 0x1cf631bb.
//
// Solidity: function handleRewardsV2(address _gmxReceiver, bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) HandleRewardsV2(_gmxReceiver common.Address, _shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.HandleRewardsV2(&_RewardRouterV2.TransactOpts, _gmxReceiver, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// Initialize is a paid mutator transaction binding the contract method 0x5261aef3.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,address,address,address,address,address,address,address,address) _initializeParams) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) Initialize(opts *bind.TransactOpts, _initializeParams RewardRouterV2InitializeParams) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "initialize", _initializeParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x5261aef3.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,address,address,address,address,address,address,address,address) _initializeParams) returns()
func (_RewardRouterV2 *RewardRouterV2Session) Initialize(_initializeParams RewardRouterV2InitializeParams) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Initialize(&_RewardRouterV2.TransactOpts, _initializeParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x5261aef3.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,address,address,address,address,address,address,address,address) _initializeParams) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) Initialize(_initializeParams RewardRouterV2InitializeParams) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Initialize(&_RewardRouterV2.TransactOpts, _initializeParams)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] externalCallTargets, bytes[] externalCallDataList, address[] refundTokens, address[] refundReceivers) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) MakeExternalCalls(opts *bind.TransactOpts, externalCallTargets []common.Address, externalCallDataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "makeExternalCalls", externalCallTargets, externalCallDataList, refundTokens, refundReceivers)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] externalCallTargets, bytes[] externalCallDataList, address[] refundTokens, address[] refundReceivers) returns()
func (_RewardRouterV2 *RewardRouterV2Session) MakeExternalCalls(externalCallTargets []common.Address, externalCallDataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.MakeExternalCalls(&_RewardRouterV2.TransactOpts, externalCallTargets, externalCallDataList, refundTokens, refundReceivers)
}

// MakeExternalCalls is a paid mutator transaction binding the contract method 0xd59922b0.
//
// Solidity: function makeExternalCalls(address[] externalCallTargets, bytes[] externalCallDataList, address[] refundTokens, address[] refundReceivers) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) MakeExternalCalls(externalCallTargets []common.Address, externalCallDataList [][]byte, refundTokens []common.Address, refundReceivers []common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.MakeExternalCalls(&_RewardRouterV2.TransactOpts, externalCallTargets, externalCallDataList, refundTokens, refundReceivers)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Transactor) MintAndStakeGlp(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "mintAndStakeGlp", _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Session) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.MintAndStakeGlp(&_RewardRouterV2.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2TransactorSession) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.MintAndStakeGlp(&_RewardRouterV2.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Transactor) MintAndStakeGlpETH(opts *bind.TransactOpts, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "mintAndStakeGlpETH", _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Session) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.MintAndStakeGlpETH(&_RewardRouterV2.TransactOpts, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_RewardRouterV2 *RewardRouterV2TransactorSession) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.MintAndStakeGlpETH(&_RewardRouterV2.TransactOpts, _minUsdg, _minGlp)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RewardRouterV2 *RewardRouterV2Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RewardRouterV2 *RewardRouterV2Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Multicall(&_RewardRouterV2.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RewardRouterV2 *RewardRouterV2TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Multicall(&_RewardRouterV2.TransactOpts, data)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardRouterV2 *RewardRouterV2Session) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetGov(&_RewardRouterV2.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetGov(&_RewardRouterV2.TransactOpts, _gov)
}

// SetGovToken is a paid mutator transaction binding the contract method 0xa86ea695.
//
// Solidity: function setGovToken(address _govToken) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) SetGovToken(opts *bind.TransactOpts, _govToken common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "setGovToken", _govToken)
}

// SetGovToken is a paid mutator transaction binding the contract method 0xa86ea695.
//
// Solidity: function setGovToken(address _govToken) returns()
func (_RewardRouterV2 *RewardRouterV2Session) SetGovToken(_govToken common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetGovToken(&_RewardRouterV2.TransactOpts, _govToken)
}

// SetGovToken is a paid mutator transaction binding the contract method 0xa86ea695.
//
// Solidity: function setGovToken(address _govToken) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) SetGovToken(_govToken common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetGovToken(&_RewardRouterV2.TransactOpts, _govToken)
}

// SetInRestakingMode is a paid mutator transaction binding the contract method 0x6d6440f0.
//
// Solidity: function setInRestakingMode(bool _inRestakingMode) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) SetInRestakingMode(opts *bind.TransactOpts, _inRestakingMode bool) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "setInRestakingMode", _inRestakingMode)
}

// SetInRestakingMode is a paid mutator transaction binding the contract method 0x6d6440f0.
//
// Solidity: function setInRestakingMode(bool _inRestakingMode) returns()
func (_RewardRouterV2 *RewardRouterV2Session) SetInRestakingMode(_inRestakingMode bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetInRestakingMode(&_RewardRouterV2.TransactOpts, _inRestakingMode)
}

// SetInRestakingMode is a paid mutator transaction binding the contract method 0x6d6440f0.
//
// Solidity: function setInRestakingMode(bool _inRestakingMode) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) SetInRestakingMode(_inRestakingMode bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetInRestakingMode(&_RewardRouterV2.TransactOpts, _inRestakingMode)
}

// SetInStrictTransferMode is a paid mutator transaction binding the contract method 0x65f8f4d0.
//
// Solidity: function setInStrictTransferMode(bool _inStrictTransferMode) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) SetInStrictTransferMode(opts *bind.TransactOpts, _inStrictTransferMode bool) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "setInStrictTransferMode", _inStrictTransferMode)
}

// SetInStrictTransferMode is a paid mutator transaction binding the contract method 0x65f8f4d0.
//
// Solidity: function setInStrictTransferMode(bool _inStrictTransferMode) returns()
func (_RewardRouterV2 *RewardRouterV2Session) SetInStrictTransferMode(_inStrictTransferMode bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetInStrictTransferMode(&_RewardRouterV2.TransactOpts, _inStrictTransferMode)
}

// SetInStrictTransferMode is a paid mutator transaction binding the contract method 0x65f8f4d0.
//
// Solidity: function setInStrictTransferMode(bool _inStrictTransferMode) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) SetInStrictTransferMode(_inStrictTransferMode bool) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetInStrictTransferMode(&_RewardRouterV2.TransactOpts, _inStrictTransferMode)
}

// SetMaxBoostBasisPoints is a paid mutator transaction binding the contract method 0x35c1b188.
//
// Solidity: function setMaxBoostBasisPoints(uint256 _maxBoostBasisPoints) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) SetMaxBoostBasisPoints(opts *bind.TransactOpts, _maxBoostBasisPoints *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "setMaxBoostBasisPoints", _maxBoostBasisPoints)
}

// SetMaxBoostBasisPoints is a paid mutator transaction binding the contract method 0x35c1b188.
//
// Solidity: function setMaxBoostBasisPoints(uint256 _maxBoostBasisPoints) returns()
func (_RewardRouterV2 *RewardRouterV2Session) SetMaxBoostBasisPoints(_maxBoostBasisPoints *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetMaxBoostBasisPoints(&_RewardRouterV2.TransactOpts, _maxBoostBasisPoints)
}

// SetMaxBoostBasisPoints is a paid mutator transaction binding the contract method 0x35c1b188.
//
// Solidity: function setMaxBoostBasisPoints(uint256 _maxBoostBasisPoints) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) SetMaxBoostBasisPoints(_maxBoostBasisPoints *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetMaxBoostBasisPoints(&_RewardRouterV2.TransactOpts, _maxBoostBasisPoints)
}

// SetVotingPowerType is a paid mutator transaction binding the contract method 0xba2db4f5.
//
// Solidity: function setVotingPowerType(uint8 _votingPowerType) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) SetVotingPowerType(opts *bind.TransactOpts, _votingPowerType uint8) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "setVotingPowerType", _votingPowerType)
}

// SetVotingPowerType is a paid mutator transaction binding the contract method 0xba2db4f5.
//
// Solidity: function setVotingPowerType(uint8 _votingPowerType) returns()
func (_RewardRouterV2 *RewardRouterV2Session) SetVotingPowerType(_votingPowerType uint8) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetVotingPowerType(&_RewardRouterV2.TransactOpts, _votingPowerType)
}

// SetVotingPowerType is a paid mutator transaction binding the contract method 0xba2db4f5.
//
// Solidity: function setVotingPowerType(uint8 _votingPowerType) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) SetVotingPowerType(_votingPowerType uint8) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SetVotingPowerType(&_RewardRouterV2.TransactOpts, _votingPowerType)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) SignalTransfer(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "signalTransfer", _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_RewardRouterV2 *RewardRouterV2Session) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SignalTransfer(&_RewardRouterV2.TransactOpts, _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.SignalTransfer(&_RewardRouterV2.TransactOpts, _receiver)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) StakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "stakeEsGmx", _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Session) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.StakeEsGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.StakeEsGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) StakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "stakeGmx", _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Session) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.StakeGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.StakeGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Transactor) UnstakeAndRedeemGlp(opts *bind.TransactOpts, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "unstakeAndRedeemGlp", _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Session) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeAndRedeemGlp(&_RewardRouterV2.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2TransactorSession) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeAndRedeemGlp(&_RewardRouterV2.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Transactor) UnstakeAndRedeemGlpETH(opts *bind.TransactOpts, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "unstakeAndRedeemGlpETH", _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2Session) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeAndRedeemGlpETH(&_RewardRouterV2.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouterV2 *RewardRouterV2TransactorSession) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeAndRedeemGlpETH(&_RewardRouterV2.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) UnstakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "unstakeEsGmx", _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Session) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeEsGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeEsGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) UnstakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "unstakeGmx", _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Session) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.UnstakeGmx(&_RewardRouterV2.TransactOpts, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2Session) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.WithdrawToken(&_RewardRouterV2.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouterV2.Contract.WithdrawToken(&_RewardRouterV2.TransactOpts, _token, _account, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardRouterV2 *RewardRouterV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouterV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardRouterV2 *RewardRouterV2Session) Receive() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Receive(&_RewardRouterV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardRouterV2 *RewardRouterV2TransactorSession) Receive() (*types.Transaction, error) {
	return _RewardRouterV2.Contract.Receive(&_RewardRouterV2.TransactOpts)
}

// RewardRouterV2StakeGlpIterator is returned from FilterStakeGlp and is used to iterate over the raw logs and unpacked data for StakeGlp events raised by the RewardRouterV2 contract.
type RewardRouterV2StakeGlpIterator struct {
	Event *RewardRouterV2StakeGlp // Event containing the contract specifics and raw log

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
func (it *RewardRouterV2StakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterV2StakeGlp)
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
		it.Event = new(RewardRouterV2StakeGlp)
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
func (it *RewardRouterV2StakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterV2StakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterV2StakeGlp represents a StakeGlp event raised by the RewardRouterV2 contract.
type RewardRouterV2StakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGlp is a free log retrieval operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) FilterStakeGlp(opts *bind.FilterOpts) (*RewardRouterV2StakeGlpIterator, error) {

	logs, sub, err := _RewardRouterV2.contract.FilterLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2StakeGlpIterator{contract: _RewardRouterV2.contract, event: "StakeGlp", logs: logs, sub: sub}, nil
}

// WatchStakeGlp is a free log subscription operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) WatchStakeGlp(opts *bind.WatchOpts, sink chan<- *RewardRouterV2StakeGlp) (event.Subscription, error) {

	logs, sub, err := _RewardRouterV2.contract.WatchLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterV2StakeGlp)
				if err := _RewardRouterV2.contract.UnpackLog(event, "StakeGlp", log); err != nil {
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
func (_RewardRouterV2 *RewardRouterV2Filterer) ParseStakeGlp(log types.Log) (*RewardRouterV2StakeGlp, error) {
	event := new(RewardRouterV2StakeGlp)
	if err := _RewardRouterV2.contract.UnpackLog(event, "StakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardRouterV2StakeGmxIterator is returned from FilterStakeGmx and is used to iterate over the raw logs and unpacked data for StakeGmx events raised by the RewardRouterV2 contract.
type RewardRouterV2StakeGmxIterator struct {
	Event *RewardRouterV2StakeGmx // Event containing the contract specifics and raw log

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
func (it *RewardRouterV2StakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterV2StakeGmx)
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
		it.Event = new(RewardRouterV2StakeGmx)
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
func (it *RewardRouterV2StakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterV2StakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterV2StakeGmx represents a StakeGmx event raised by the RewardRouterV2 contract.
type RewardRouterV2StakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGmx is a free log retrieval operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) FilterStakeGmx(opts *bind.FilterOpts) (*RewardRouterV2StakeGmxIterator, error) {

	logs, sub, err := _RewardRouterV2.contract.FilterLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2StakeGmxIterator{contract: _RewardRouterV2.contract, event: "StakeGmx", logs: logs, sub: sub}, nil
}

// WatchStakeGmx is a free log subscription operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) WatchStakeGmx(opts *bind.WatchOpts, sink chan<- *RewardRouterV2StakeGmx) (event.Subscription, error) {

	logs, sub, err := _RewardRouterV2.contract.WatchLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterV2StakeGmx)
				if err := _RewardRouterV2.contract.UnpackLog(event, "StakeGmx", log); err != nil {
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
func (_RewardRouterV2 *RewardRouterV2Filterer) ParseStakeGmx(log types.Log) (*RewardRouterV2StakeGmx, error) {
	event := new(RewardRouterV2StakeGmx)
	if err := _RewardRouterV2.contract.UnpackLog(event, "StakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardRouterV2UnstakeGlpIterator is returned from FilterUnstakeGlp and is used to iterate over the raw logs and unpacked data for UnstakeGlp events raised by the RewardRouterV2 contract.
type RewardRouterV2UnstakeGlpIterator struct {
	Event *RewardRouterV2UnstakeGlp // Event containing the contract specifics and raw log

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
func (it *RewardRouterV2UnstakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterV2UnstakeGlp)
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
		it.Event = new(RewardRouterV2UnstakeGlp)
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
func (it *RewardRouterV2UnstakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterV2UnstakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterV2UnstakeGlp represents a UnstakeGlp event raised by the RewardRouterV2 contract.
type RewardRouterV2UnstakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGlp is a free log retrieval operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) FilterUnstakeGlp(opts *bind.FilterOpts) (*RewardRouterV2UnstakeGlpIterator, error) {

	logs, sub, err := _RewardRouterV2.contract.FilterLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2UnstakeGlpIterator{contract: _RewardRouterV2.contract, event: "UnstakeGlp", logs: logs, sub: sub}, nil
}

// WatchUnstakeGlp is a free log subscription operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) WatchUnstakeGlp(opts *bind.WatchOpts, sink chan<- *RewardRouterV2UnstakeGlp) (event.Subscription, error) {

	logs, sub, err := _RewardRouterV2.contract.WatchLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterV2UnstakeGlp)
				if err := _RewardRouterV2.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
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
func (_RewardRouterV2 *RewardRouterV2Filterer) ParseUnstakeGlp(log types.Log) (*RewardRouterV2UnstakeGlp, error) {
	event := new(RewardRouterV2UnstakeGlp)
	if err := _RewardRouterV2.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardRouterV2UnstakeGmxIterator is returned from FilterUnstakeGmx and is used to iterate over the raw logs and unpacked data for UnstakeGmx events raised by the RewardRouterV2 contract.
type RewardRouterV2UnstakeGmxIterator struct {
	Event *RewardRouterV2UnstakeGmx // Event containing the contract specifics and raw log

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
func (it *RewardRouterV2UnstakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterV2UnstakeGmx)
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
		it.Event = new(RewardRouterV2UnstakeGmx)
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
func (it *RewardRouterV2UnstakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterV2UnstakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterV2UnstakeGmx represents a UnstakeGmx event raised by the RewardRouterV2 contract.
type RewardRouterV2UnstakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGmx is a free log retrieval operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) FilterUnstakeGmx(opts *bind.FilterOpts) (*RewardRouterV2UnstakeGmxIterator, error) {

	logs, sub, err := _RewardRouterV2.contract.FilterLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return &RewardRouterV2UnstakeGmxIterator{contract: _RewardRouterV2.contract, event: "UnstakeGmx", logs: logs, sub: sub}, nil
}

// WatchUnstakeGmx is a free log subscription operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_RewardRouterV2 *RewardRouterV2Filterer) WatchUnstakeGmx(opts *bind.WatchOpts, sink chan<- *RewardRouterV2UnstakeGmx) (event.Subscription, error) {

	logs, sub, err := _RewardRouterV2.contract.WatchLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterV2UnstakeGmx)
				if err := _RewardRouterV2.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
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
func (_RewardRouterV2 *RewardRouterV2Filterer) ParseUnstakeGmx(log types.Log) (*RewardRouterV2UnstakeGmx, error) {
	event := new(RewardRouterV2UnstakeGmx)
	if err := _RewardRouterV2.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
