// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivaultutils

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

// IVaultUtilsMetaData contains all meta data concerning the IVaultUtils contract.
var IVaultUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getBuyUsdgFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getEntryFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_increment\",\"type\":\"bool\"}],\"name\":\"getFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_entryFundingRate\",\"type\":\"uint256\"}],\"name\":\"getFundingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"}],\"name\":\"getPositionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getSellUsdgFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getSwapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"}],\"name\":\"updateCumulativeFundingRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"validateDecreasePosition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"validateIncreasePosition\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_raise\",\"type\":\"bool\"}],\"name\":\"validateLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IVaultUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use IVaultUtilsMetaData.ABI instead.
var IVaultUtilsABI = IVaultUtilsMetaData.ABI

// IVaultUtils is an auto generated Go binding around an Ethereum contract.
type IVaultUtils struct {
	IVaultUtilsCaller     // Read-only binding to the contract
	IVaultUtilsTransactor // Write-only binding to the contract
	IVaultUtilsFilterer   // Log filterer for contract events
}

// IVaultUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVaultUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVaultUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVaultUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVaultUtilsSession struct {
	Contract     *IVaultUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVaultUtilsCallerSession struct {
	Contract *IVaultUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IVaultUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVaultUtilsTransactorSession struct {
	Contract     *IVaultUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IVaultUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVaultUtilsRaw struct {
	Contract *IVaultUtils // Generic contract binding to access the raw methods on
}

// IVaultUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVaultUtilsCallerRaw struct {
	Contract *IVaultUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// IVaultUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVaultUtilsTransactorRaw struct {
	Contract *IVaultUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVaultUtils creates a new instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtils(address common.Address, backend bind.ContractBackend) (*IVaultUtils, error) {
	contract, err := bindIVaultUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVaultUtils{IVaultUtilsCaller: IVaultUtilsCaller{contract: contract}, IVaultUtilsTransactor: IVaultUtilsTransactor{contract: contract}, IVaultUtilsFilterer: IVaultUtilsFilterer{contract: contract}}, nil
}

// NewIVaultUtilsCaller creates a new read-only instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtilsCaller(address common.Address, caller bind.ContractCaller) (*IVaultUtilsCaller, error) {
	contract, err := bindIVaultUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultUtilsCaller{contract: contract}, nil
}

// NewIVaultUtilsTransactor creates a new write-only instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*IVaultUtilsTransactor, error) {
	contract, err := bindIVaultUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultUtilsTransactor{contract: contract}, nil
}

// NewIVaultUtilsFilterer creates a new log filterer instance of IVaultUtils, bound to a specific deployed contract.
func NewIVaultUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*IVaultUtilsFilterer, error) {
	contract, err := bindIVaultUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVaultUtilsFilterer{contract: contract}, nil
}

// bindIVaultUtils binds a generic wrapper to an already deployed contract.
func bindIVaultUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVaultUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultUtils *IVaultUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultUtils.Contract.IVaultUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultUtils *IVaultUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultUtils.Contract.IVaultUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultUtils *IVaultUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultUtils.Contract.IVaultUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVaultUtils *IVaultUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVaultUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVaultUtils *IVaultUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVaultUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVaultUtils *IVaultUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVaultUtils.Contract.contract.Transact(opts, method, params...)
}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetBuyUsdgFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getBuyUsdgFeeBasisPoints", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetBuyUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetBuyUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetBuyUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0x4adeddc6.
//
// Solidity: function getBuyUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetBuyUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetBuyUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetEntryFundingRate(&_IVaultUtils.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetEntryFundingRate(&_IVaultUtils.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetFundingFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getFundingFee", _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFundingFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetFundingFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetPositionFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getPositionFee", _account, _collateralToken, _indexToken, _isLong, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetPositionFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetPositionFee(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetSellUsdgFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getSellUsdgFeeBasisPoints", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetSellUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSellUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetSellUsdgFeeBasisPoints is a free data retrieval call binding the contract method 0xeb0835bf.
//
// Solidity: function getSellUsdgFeeBasisPoints(address _token, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetSellUsdgFeeBasisPoints(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSellUsdgFeeBasisPoints(&_IVaultUtils.CallOpts, _token, _usdgAmount)
}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCaller) GetSwapFeeBasisPoints(opts *bind.CallOpts, _tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "getSwapFeeBasisPoints", _tokenIn, _tokenOut, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsSession) GetSwapFeeBasisPoints(_tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSwapFeeBasisPoints(&_IVaultUtils.CallOpts, _tokenIn, _tokenOut, _usdgAmount)
}

// GetSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xda133816.
//
// Solidity: function getSwapFeeBasisPoints(address _tokenIn, address _tokenOut, uint256 _usdgAmount) view returns(uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) GetSwapFeeBasisPoints(_tokenIn common.Address, _tokenOut common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _IVaultUtils.Contract.GetSwapFeeBasisPoints(&_IVaultUtils.CallOpts, _tokenIn, _tokenOut, _usdgAmount)
}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) view returns()
func (_IVaultUtils *IVaultUtilsCaller) ValidateDecreasePosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) error {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "validateDecreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)

	if err != nil {
		return err
	}

	return err

}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) view returns()
func (_IVaultUtils *IVaultUtilsSession) ValidateDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) error {
	return _IVaultUtils.Contract.ValidateDecreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// ValidateDecreasePosition is a free data retrieval call binding the contract method 0x81d11a23.
//
// Solidity: function validateDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) view returns()
func (_IVaultUtils *IVaultUtilsCallerSession) ValidateDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) error {
	return _IVaultUtils.Contract.ValidateDecreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) view returns()
func (_IVaultUtils *IVaultUtilsCaller) ValidateIncreasePosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) error {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "validateIncreasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)

	if err != nil {
		return err
	}

	return err

}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) view returns()
func (_IVaultUtils *IVaultUtilsSession) ValidateIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) error {
	return _IVaultUtils.Contract.ValidateIncreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// ValidateIncreasePosition is a free data retrieval call binding the contract method 0x9d5c28fa.
//
// Solidity: function validateIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) view returns()
func (_IVaultUtils *IVaultUtilsCallerSession) ValidateIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) error {
	return _IVaultUtils.Contract.ValidateIncreasePosition(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVaultUtils *IVaultUtilsCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IVaultUtils.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

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
func (_IVaultUtils *IVaultUtilsSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _IVaultUtils.Contract.ValidateLiquidation(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_IVaultUtils *IVaultUtilsCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _IVaultUtils.Contract.ValidateLiquidation(&_IVaultUtils.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns(bool)
func (_IVaultUtils *IVaultUtilsTransactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _IVaultUtils.contract.Transact(opts, "updateCumulativeFundingRate", _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns(bool)
func (_IVaultUtils *IVaultUtilsSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _IVaultUtils.Contract.UpdateCumulativeFundingRate(&_IVaultUtils.TransactOpts, _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns(bool)
func (_IVaultUtils *IVaultUtilsTransactorSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _IVaultUtils.Contract.UpdateCumulativeFundingRate(&_IVaultUtils.TransactOpts, _collateralToken, _indexToken)
}
