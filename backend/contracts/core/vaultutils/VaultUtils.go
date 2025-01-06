// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultutils

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

// VaultUtilsMetaData contains all meta data concerning the VaultUtils contract.
var VaultUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUNDING_RATE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getBuyUsdgFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"getEntryFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_increment\",\"type\":\"bool\"}],\"name\":\"getFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_entryFundingRate\",\"type\":\"uint256\"}],\"name\":\"getFundingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"}],\"name\":\"getPositionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getSellUsdgFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getSwapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"updateCumulativeFundingRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validateDecreasePosition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"validateIncreasePosition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_raise\",\"type\":\"bool\"}],\"name\":\"validateLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultUtilsMetaData.ABI instead.
var VaultUtilsABI = VaultUtilsMetaData.ABI

// VaultUtils is an auto generated Go binding around an Ethereum contract.
type VaultUtils struct {
	VaultUtilsCaller     // Read-only binding to the contract
	VaultUtilsTransactor // Write-only binding to the contract
	VaultUtilsFilterer   // Log filterer for contract events
}

// VaultUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultUtilsSession struct {
	Contract     *VaultUtils       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultUtilsCallerSession struct {
	Contract *VaultUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VaultUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultUtilsTransactorSession struct {
	Contract     *VaultUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VaultUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultUtilsRaw struct {
	Contract *VaultUtils // Generic contract binding to access the raw methods on
}

// VaultUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultUtilsCallerRaw struct {
	Contract *VaultUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// VaultUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultUtilsTransactorRaw struct {
	Contract *VaultUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultUtils creates a new instance of VaultUtils, bound to a specific deployed contract.
func NewVaultUtils(address common.Address, backend bind.ContractBackend) (*VaultUtils, error) {
	contract, err := bindVaultUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultUtils{VaultUtilsCaller: VaultUtilsCaller{contract: contract}, VaultUtilsTransactor: VaultUtilsTransactor{contract: contract}, VaultUtilsFilterer: VaultUtilsFilterer{contract: contract}}, nil
}

// NewVaultUtilsCaller creates a new read-only instance of VaultUtils, bound to a specific deployed contract.
func NewVaultUtilsCaller(address common.Address, caller bind.ContractCaller) (*VaultUtilsCaller, error) {
	contract, err := bindVaultUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultUtilsCaller{contract: contract}, nil
}

// NewVaultUtilsTransactor creates a new write-only instance of VaultUtils, bound to a specific deployed contract.
func NewVaultUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultUtilsTransactor, error) {
	contract, err := bindVaultUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultUtilsTransactor{contract: contract}, nil
}

// NewVaultUtilsFilterer creates a new log filterer instance of VaultUtils, bound to a specific deployed contract.
func NewVaultUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultUtilsFilterer, error) {
	contract, err := bindVaultUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultUtilsFilterer{contract: contract}, nil
}

// bindVaultUtils binds a generic wrapper to an already deployed contract.
func bindVaultUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultUtils *VaultUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultUtils.Contract.VaultUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultUtils *VaultUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultUtils.Contract.VaultUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultUtils *VaultUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultUtils.Contract.VaultUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultUtils *VaultUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultUtils *VaultUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultUtils *VaultUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultUtils.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultUtils *VaultUtilsSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultUtils.Contract.BASISPOINTSDIVISOR(&_VaultUtils.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultUtils.Contract.BASISPOINTSDIVISOR(&_VaultUtils.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) FUNDINGRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "FUNDING_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultUtils *VaultUtilsSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultUtils.Contract.FUNDINGRATEPRECISION(&_VaultUtils.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultUtils.Contract.FUNDINGRATEPRECISION(&_VaultUtils.CallOpts)
}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) GetBuyUsdgFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "getBuyUsdgFeeBasisPoints", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsSession) GetBuyUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetBuyUsdgFeeBasisPoints(&_VaultUtils.CallOpts, _token, _usdgAmount)
}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) GetBuyUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetBuyUsdgFeeBasisPoints(&_VaultUtils.CallOpts, _token, _usdgAmount)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address , bool ) view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, arg1 common.Address, arg2 bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address , bool ) view returns(uint256)
func (_VaultUtils *VaultUtilsSession) GetEntryFundingRate(_collateralToken common.Address, arg1 common.Address, arg2 bool) (*big.Int, error) {
	return _VaultUtils.Contract.GetEntryFundingRate(&_VaultUtils.CallOpts, _collateralToken, arg1, arg2)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address , bool ) view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) GetEntryFundingRate(_collateralToken common.Address, arg1 common.Address, arg2 bool) (*big.Int, error) {
	return _VaultUtils.Contract.GetEntryFundingRate(&_VaultUtils.CallOpts, _collateralToken, arg1, arg2)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultUtils *VaultUtilsSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultUtils.Contract.GetFeeBasisPoints(&_VaultUtils.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultUtils.Contract.GetFeeBasisPoints(&_VaultUtils.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address , address _collateralToken, address , bool , uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) GetFundingFee(opts *bind.CallOpts, arg0 common.Address, _collateralToken common.Address, arg2 common.Address, arg3 bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "getFundingFee", arg0, _collateralToken, arg2, arg3, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address , address _collateralToken, address , bool , uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultUtils *VaultUtilsSession) GetFundingFee(arg0 common.Address, _collateralToken common.Address, arg2 common.Address, arg3 bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetFundingFee(&_VaultUtils.CallOpts, arg0, _collateralToken, arg2, arg3, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address , address _collateralToken, address , bool , uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) GetFundingFee(arg0 common.Address, _collateralToken common.Address, arg2 common.Address, arg3 bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetFundingFee(&_VaultUtils.CallOpts, arg0, _collateralToken, arg2, arg3, _size, _entryFundingRate)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address , address , address , bool , uint256 _sizeDelta) view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) GetPositionFee(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "getPositionFee", arg0, arg1, arg2, arg3, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address , address , address , bool , uint256 _sizeDelta) view returns(uint256)
func (_VaultUtils *VaultUtilsSession) GetPositionFee(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetPositionFee(&_VaultUtils.CallOpts, arg0, arg1, arg2, arg3, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address , address , address , bool , uint256 _sizeDelta) view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) GetPositionFee(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetPositionFee(&_VaultUtils.CallOpts, arg0, arg1, arg2, arg3, _sizeDelta)
}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) GetSellUsdgFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "getSellUsdgFeeBasisPoints", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsSession) GetSellUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetSellUsdgFeeBasisPoints(&_VaultUtils.CallOpts, _token, _usdgAmount)
}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) GetSellUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetSellUsdgFeeBasisPoints(&_VaultUtils.CallOpts, _token, _usdgAmount)
}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsCaller) GetSwapFeeBasisPoints(opts *bind.CallOpts, _tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "getSwapFeeBasisPoints", _tokenIn, _tokenOut, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsSession) GetSwapFeeBasisPoints(_tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetSwapFeeBasisPoints(&_VaultUtils.CallOpts, _tokenIn, _tokenOut, _usdgAmount)
}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_VaultUtils *VaultUtilsCallerSession) GetSwapFeeBasisPoints(_tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultUtils.Contract.GetSwapFeeBasisPoints(&_VaultUtils.CallOpts, _tokenIn, _tokenOut, _usdgAmount)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultUtils *VaultUtilsCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultUtils *VaultUtilsSession) Gov() (common.Address, error) {
	return _VaultUtils.Contract.Gov(&_VaultUtils.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultUtils *VaultUtilsCallerSession) Gov() (common.Address, error) {
	return _VaultUtils.Contract.Gov(&_VaultUtils.CallOpts)
}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address , address , address , uint256 , uint256 , bool , address ) view returns()
func (_VaultUtils *VaultUtilsCaller) ValidateDecreasePosition(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 bool, arg6 common.Address) error {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "validateDecreasePosition", arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	if err != nil {
		return err
	}

	return err

}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address , address , address , uint256 , uint256 , bool , address ) view returns()
func (_VaultUtils *VaultUtilsSession) ValidateDecreasePosition(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 bool, arg6 common.Address) error {
	return _VaultUtils.Contract.ValidateDecreasePosition(&_VaultUtils.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address , address , address , uint256 , uint256 , bool , address ) view returns()
func (_VaultUtils *VaultUtilsCallerSession) ValidateDecreasePosition(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 bool, arg6 common.Address) error {
	return _VaultUtils.Contract.ValidateDecreasePosition(&_VaultUtils.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address , address , address , uint256 , bool ) view returns()
func (_VaultUtils *VaultUtilsCaller) ValidateIncreasePosition(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 bool) error {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "validateIncreasePosition", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address , address , address , uint256 , bool ) view returns()
func (_VaultUtils *VaultUtilsSession) ValidateIncreasePosition(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 bool) error {
	return _VaultUtils.Contract.ValidateIncreasePosition(&_VaultUtils.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address , address , address , uint256 , bool ) view returns()
func (_VaultUtils *VaultUtilsCallerSession) ValidateIncreasePosition(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 bool) error {
	return _VaultUtils.Contract.ValidateIncreasePosition(&_VaultUtils.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultUtils *VaultUtilsCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultUtils *VaultUtilsSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultUtils.Contract.ValidateLiquidation(&_VaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultUtils *VaultUtilsCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultUtils.Contract.ValidateLiquidation(&_VaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_VaultUtils *VaultUtilsCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultUtils.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_VaultUtils *VaultUtilsSession) Vault() (common.Address, error) {
	return _VaultUtils.Contract.Vault(&_VaultUtils.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_VaultUtils *VaultUtilsCallerSession) Vault() (common.Address, error) {
	return _VaultUtils.Contract.Vault(&_VaultUtils.CallOpts)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultUtils *VaultUtilsTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _VaultUtils.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultUtils *VaultUtilsSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultUtils.Contract.SetGov(&_VaultUtils.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultUtils *VaultUtilsTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultUtils.Contract.SetGov(&_VaultUtils.TransactOpts, _gov)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address , address ) returns(bool)
func (_VaultUtils *VaultUtilsTransactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _VaultUtils.contract.Transact(opts, "updateCumulativeFundingRate", arg0, arg1)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address , address ) returns(bool)
func (_VaultUtils *VaultUtilsSession) UpdateCumulativeFundingRate(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _VaultUtils.Contract.UpdateCumulativeFundingRate(&_VaultUtils.TransactOpts, arg0, arg1)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address , address ) returns(bool)
func (_VaultUtils *VaultUtilsTransactorSession) UpdateCumulativeFundingRate(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _VaultUtils.Contract.UpdateCumulativeFundingRate(&_VaultUtils.TransactOpts, arg0, arg1)
}
