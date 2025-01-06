// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iglpmanager

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

// IGlpManagerMetaData contains all meta data concerning the IGlpManager contract.
var IGlpManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fundingAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGlp\",\"type\":\"uint256\"}],\"name\":\"addLiquidityForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"maximise\",\"type\":\"bool\"}],\"name\":\"getAumInUsdg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"lastAddedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_glpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeLiquidityForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cooldownDuration\",\"type\":\"uint256\"}],\"name\":\"setCooldownDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shortsTrackerAveragePriceWeight\",\"type\":\"uint256\"}],\"name\":\"setShortsTrackerAveragePriceWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IGlpManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IGlpManagerMetaData.ABI instead.
var IGlpManagerABI = IGlpManagerMetaData.ABI

// IGlpManager is an auto generated Go binding around an Ethereum contract.
type IGlpManager struct {
	IGlpManagerCaller     // Read-only binding to the contract
	IGlpManagerTransactor // Write-only binding to the contract
	IGlpManagerFilterer   // Log filterer for contract events
}

// IGlpManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGlpManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlpManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGlpManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlpManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGlpManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlpManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGlpManagerSession struct {
	Contract     *IGlpManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGlpManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGlpManagerCallerSession struct {
	Contract *IGlpManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IGlpManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGlpManagerTransactorSession struct {
	Contract     *IGlpManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IGlpManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGlpManagerRaw struct {
	Contract *IGlpManager // Generic contract binding to access the raw methods on
}

// IGlpManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGlpManagerCallerRaw struct {
	Contract *IGlpManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IGlpManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGlpManagerTransactorRaw struct {
	Contract *IGlpManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGlpManager creates a new instance of IGlpManager, bound to a specific deployed contract.
func NewIGlpManager(address common.Address, backend bind.ContractBackend) (*IGlpManager, error) {
	contract, err := bindIGlpManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGlpManager{IGlpManagerCaller: IGlpManagerCaller{contract: contract}, IGlpManagerTransactor: IGlpManagerTransactor{contract: contract}, IGlpManagerFilterer: IGlpManagerFilterer{contract: contract}}, nil
}

// NewIGlpManagerCaller creates a new read-only instance of IGlpManager, bound to a specific deployed contract.
func NewIGlpManagerCaller(address common.Address, caller bind.ContractCaller) (*IGlpManagerCaller, error) {
	contract, err := bindIGlpManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGlpManagerCaller{contract: contract}, nil
}

// NewIGlpManagerTransactor creates a new write-only instance of IGlpManager, bound to a specific deployed contract.
func NewIGlpManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IGlpManagerTransactor, error) {
	contract, err := bindIGlpManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGlpManagerTransactor{contract: contract}, nil
}

// NewIGlpManagerFilterer creates a new log filterer instance of IGlpManager, bound to a specific deployed contract.
func NewIGlpManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IGlpManagerFilterer, error) {
	contract, err := bindIGlpManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGlpManagerFilterer{contract: contract}, nil
}

// bindIGlpManager binds a generic wrapper to an already deployed contract.
func bindIGlpManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGlpManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlpManager *IGlpManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGlpManager.Contract.IGlpManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlpManager *IGlpManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlpManager.Contract.IGlpManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlpManager *IGlpManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlpManager.Contract.IGlpManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlpManager *IGlpManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGlpManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlpManager *IGlpManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlpManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlpManager *IGlpManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlpManager.Contract.contract.Transact(opts, method, params...)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_IGlpManager *IGlpManagerCaller) GetAumInUsdg(opts *bind.CallOpts, maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _IGlpManager.contract.Call(opts, &out, "getAumInUsdg", maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_IGlpManager *IGlpManagerSession) GetAumInUsdg(maximise bool) (*big.Int, error) {
	return _IGlpManager.Contract.GetAumInUsdg(&_IGlpManager.CallOpts, maximise)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_IGlpManager *IGlpManagerCallerSession) GetAumInUsdg(maximise bool) (*big.Int, error) {
	return _IGlpManager.Contract.GetAumInUsdg(&_IGlpManager.CallOpts, maximise)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_IGlpManager *IGlpManagerCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGlpManager.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_IGlpManager *IGlpManagerSession) Glp() (common.Address, error) {
	return _IGlpManager.Contract.Glp(&_IGlpManager.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_IGlpManager *IGlpManagerCallerSession) Glp() (common.Address, error) {
	return _IGlpManager.Contract.Glp(&_IGlpManager.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IGlpManager *IGlpManagerCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGlpManager.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IGlpManager *IGlpManagerSession) Usdg() (common.Address, error) {
	return _IGlpManager.Contract.Usdg(&_IGlpManager.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_IGlpManager *IGlpManagerCallerSession) Usdg() (common.Address, error) {
	return _IGlpManager.Contract.Usdg(&_IGlpManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_IGlpManager *IGlpManagerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IGlpManager.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_IGlpManager *IGlpManagerSession) Vault() (common.Address, error) {
	return _IGlpManager.Contract.Vault(&_IGlpManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_IGlpManager *IGlpManagerCallerSession) Vault() (common.Address, error) {
	return _IGlpManager.Contract.Vault(&_IGlpManager.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_IGlpManager *IGlpManagerTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "addLiquidity", _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_IGlpManager *IGlpManagerSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.AddLiquidity(&_IGlpManager.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_IGlpManager *IGlpManagerTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.AddLiquidity(&_IGlpManager.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_IGlpManager *IGlpManagerTransactor) AddLiquidityForAccount(opts *bind.TransactOpts, _fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "addLiquidityForAccount", _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_IGlpManager *IGlpManagerSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.AddLiquidityForAccount(&_IGlpManager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_IGlpManager *IGlpManagerTransactorSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.AddLiquidityForAccount(&_IGlpManager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// CooldownDuration is a paid mutator transaction binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() returns(uint256)
func (_IGlpManager *IGlpManagerTransactor) CooldownDuration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "cooldownDuration")
}

// CooldownDuration is a paid mutator transaction binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() returns(uint256)
func (_IGlpManager *IGlpManagerSession) CooldownDuration() (*types.Transaction, error) {
	return _IGlpManager.Contract.CooldownDuration(&_IGlpManager.TransactOpts)
}

// CooldownDuration is a paid mutator transaction binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() returns(uint256)
func (_IGlpManager *IGlpManagerTransactorSession) CooldownDuration() (*types.Transaction, error) {
	return _IGlpManager.Contract.CooldownDuration(&_IGlpManager.TransactOpts)
}

// LastAddedAt is a paid mutator transaction binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address _account) returns(uint256)
func (_IGlpManager *IGlpManagerTransactor) LastAddedAt(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "lastAddedAt", _account)
}

// LastAddedAt is a paid mutator transaction binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address _account) returns(uint256)
func (_IGlpManager *IGlpManagerSession) LastAddedAt(_account common.Address) (*types.Transaction, error) {
	return _IGlpManager.Contract.LastAddedAt(&_IGlpManager.TransactOpts, _account)
}

// LastAddedAt is a paid mutator transaction binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address _account) returns(uint256)
func (_IGlpManager *IGlpManagerTransactorSession) LastAddedAt(_account common.Address) (*types.Transaction, error) {
	return _IGlpManager.Contract.LastAddedAt(&_IGlpManager.TransactOpts, _account)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_IGlpManager *IGlpManagerTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "removeLiquidity", _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_IGlpManager *IGlpManagerSession) RemoveLiquidity(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGlpManager.Contract.RemoveLiquidity(&_IGlpManager.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_IGlpManager *IGlpManagerTransactorSession) RemoveLiquidity(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGlpManager.Contract.RemoveLiquidity(&_IGlpManager.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_IGlpManager *IGlpManagerTransactor) RemoveLiquidityForAccount(opts *bind.TransactOpts, _account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "removeLiquidityForAccount", _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_IGlpManager *IGlpManagerSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGlpManager.Contract.RemoveLiquidityForAccount(&_IGlpManager.TransactOpts, _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_IGlpManager *IGlpManagerTransactorSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _IGlpManager.Contract.RemoveLiquidityForAccount(&_IGlpManager.TransactOpts, _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_IGlpManager *IGlpManagerTransactor) SetCooldownDuration(opts *bind.TransactOpts, _cooldownDuration *big.Int) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "setCooldownDuration", _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_IGlpManager *IGlpManagerSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.SetCooldownDuration(&_IGlpManager.TransactOpts, _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_IGlpManager *IGlpManagerTransactorSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.SetCooldownDuration(&_IGlpManager.TransactOpts, _cooldownDuration)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_IGlpManager *IGlpManagerTransactor) SetShortsTrackerAveragePriceWeight(opts *bind.TransactOpts, _shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _IGlpManager.contract.Transact(opts, "setShortsTrackerAveragePriceWeight", _shortsTrackerAveragePriceWeight)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_IGlpManager *IGlpManagerSession) SetShortsTrackerAveragePriceWeight(_shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.SetShortsTrackerAveragePriceWeight(&_IGlpManager.TransactOpts, _shortsTrackerAveragePriceWeight)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_IGlpManager *IGlpManagerTransactorSession) SetShortsTrackerAveragePriceWeight(_shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _IGlpManager.Contract.SetShortsTrackerAveragePriceWeight(&_IGlpManager.TransactOpts, _shortsTrackerAveragePriceWeight)
}
