// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaulttest

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

// VaultTestMetaData contains all meta data concerning the VaultTest contract.
var VaultTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"BuyUSDG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averagePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryFundingRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"}],\"name\":\"ClosePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeTokens\",\"type\":\"uint256\"}],\"name\":\"CollectMarginFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeTokens\",\"type\":\"uint256\"}],\"name\":\"CollectSwapFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseGuaranteedUsd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreasePoolAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"DecreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseReservedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseUsdgAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DirectPoolDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseGuaranteedUsd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreasePoolAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sizeDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"IncreasePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseReservedAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseUsdgAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"indexToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"markPrice\",\"type\":\"uint256\"}],\"name\":\"LiquidatePosition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"SellUSDG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutAfterFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundingRate\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProfit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delta\",\"type\":\"uint256\"}],\"name\":\"UpdatePnl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"averagePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryFundingRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"markPrice\",\"type\":\"uint256\"}],\"name\":\"UpdatePosition\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FUNDING_RATE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNDING_RATE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LIQUIDATION_FEE_USD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FUNDING_RATE_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LEVERAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDG_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"addRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenDiv\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenMul\",\"type\":\"address\"}],\"name\":\"adjustForDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allWhitelistedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allWhitelistedTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedRouters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bufferAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"buyUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"clearTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeFundingRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"decreasePosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"directPoolDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"errorController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"errors\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\"}],\"name\":\"getDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getEntryFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_increment\",\"type\":\"bool\"}],\"name\":\"getFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_entryFundingRate\",\"type\":\"uint256\"}],\"name\":\"getFundingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getGlobalShortDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMaxPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getMinPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\"}],\"name\":\"getNextAveragePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getNextFundingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nextPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"}],\"name\":\"getNextGlobalShortAveragePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPositionDelta\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"}],\"name\":\"getPositionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPositionKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"getPositionLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"getRedemptionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getRedemptionCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getRedemptionCollateralUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTargetUsdgAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getUtilisation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"globalShortAveragePrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"globalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"guaranteedUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasDynamicFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inManagerMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateLiquidationMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"includeAmmPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseGlobalShortSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sizeDelta\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"}],\"name\":\"increasePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingRateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLeverageEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isLiquidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSwapEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastFundingTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"liquidatePosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationFeeUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marginFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxGlobalShortSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxUsdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minProfitBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minProfitTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBurnFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averagePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryFundingRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveAmount\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"realisedPnl\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"lastIncreasedTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"removeRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reservedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"sellUSDG\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBufferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_errorCode\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_error\",\"type\":\"string\"}],\"name\":\"setError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_errorController\",\"type\":\"address\"}],\"name\":\"setErrorController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_hasDynamicFees\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingRateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\"}],\"name\":\"setFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inManagerMode\",\"type\":\"bool\"}],\"name\":\"setInManagerMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateLiquidationMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isLeverageEnabled\",\"type\":\"bool\"}],\"name\":\"setIsLeverageEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSwapEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSwapEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isManager\",\"type\":\"bool\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"setMaxGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMaxGlobalShortSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLeverage\",\"type\":\"uint256\"}],\"name\":\"setMaxLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isShortable\",\"type\":\"bool\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setUsdgAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVaultUtils\",\"name\":\"_vaultUtils\",\"type\":\"address\"}],\"name\":\"setVaultUtils\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shortableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableFundingRateFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableTaxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stableTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFeeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"tokenToUsdMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"}],\"name\":\"updateCumulativeFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"upgradeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"usdToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdAmount\",\"type\":\"uint256\"}],\"name\":\"usdToTokenMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdAmount\",\"type\":\"uint256\"}],\"name\":\"usdToTokenMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usdgAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useSwapPricing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLong\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_raise\",\"type\":\"bool\"}],\"name\":\"validateLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultUtils\",\"outputs\":[{\"internalType\":\"contractIVaultUtils\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VaultTestABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultTestMetaData.ABI instead.
var VaultTestABI = VaultTestMetaData.ABI

// VaultTest is an auto generated Go binding around an Ethereum contract.
type VaultTest struct {
	VaultTestCaller     // Read-only binding to the contract
	VaultTestTransactor // Write-only binding to the contract
	VaultTestFilterer   // Log filterer for contract events
}

// VaultTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultTestSession struct {
	Contract     *VaultTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultTestCallerSession struct {
	Contract *VaultTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VaultTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTestTransactorSession struct {
	Contract     *VaultTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VaultTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultTestRaw struct {
	Contract *VaultTest // Generic contract binding to access the raw methods on
}

// VaultTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultTestCallerRaw struct {
	Contract *VaultTestCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTestTransactorRaw struct {
	Contract *VaultTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultTest creates a new instance of VaultTest, bound to a specific deployed contract.
func NewVaultTest(address common.Address, backend bind.ContractBackend) (*VaultTest, error) {
	contract, err := bindVaultTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultTest{VaultTestCaller: VaultTestCaller{contract: contract}, VaultTestTransactor: VaultTestTransactor{contract: contract}, VaultTestFilterer: VaultTestFilterer{contract: contract}}, nil
}

// NewVaultTestCaller creates a new read-only instance of VaultTest, bound to a specific deployed contract.
func NewVaultTestCaller(address common.Address, caller bind.ContractCaller) (*VaultTestCaller, error) {
	contract, err := bindVaultTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTestCaller{contract: contract}, nil
}

// NewVaultTestTransactor creates a new write-only instance of VaultTest, bound to a specific deployed contract.
func NewVaultTestTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTestTransactor, error) {
	contract, err := bindVaultTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTestTransactor{contract: contract}, nil
}

// NewVaultTestFilterer creates a new log filterer instance of VaultTest, bound to a specific deployed contract.
func NewVaultTestFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultTestFilterer, error) {
	contract, err := bindVaultTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultTestFilterer{contract: contract}, nil
}

// bindVaultTest binds a generic wrapper to an already deployed contract.
func bindVaultTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTest *VaultTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTest.Contract.VaultTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTest *VaultTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTest.Contract.VaultTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTest *VaultTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTest.Contract.VaultTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTest *VaultTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTest *VaultTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTest *VaultTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTest.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultTest *VaultTestCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultTest *VaultTestSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultTest.Contract.BASISPOINTSDIVISOR(&_VaultTest.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultTest.Contract.BASISPOINTSDIVISOR(&_VaultTest.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultTest *VaultTestCaller) FUNDINGRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "FUNDING_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultTest *VaultTestSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultTest.Contract.FUNDINGRATEPRECISION(&_VaultTest.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultTest.Contract.FUNDINGRATEPRECISION(&_VaultTest.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultTest *VaultTestCaller) MAXFEEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "MAX_FEE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultTest *VaultTestSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _VaultTest.Contract.MAXFEEBASISPOINTS(&_VaultTest.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _VaultTest.Contract.MAXFEEBASISPOINTS(&_VaultTest.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultTest *VaultTestCaller) MAXFUNDINGRATEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "MAX_FUNDING_RATE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultTest *VaultTestSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _VaultTest.Contract.MAXFUNDINGRATEFACTOR(&_VaultTest.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _VaultTest.Contract.MAXFUNDINGRATEFACTOR(&_VaultTest.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultTest *VaultTestCaller) MAXLIQUIDATIONFEEUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "MAX_LIQUIDATION_FEE_USD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultTest *VaultTestSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _VaultTest.Contract.MAXLIQUIDATIONFEEUSD(&_VaultTest.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _VaultTest.Contract.MAXLIQUIDATIONFEEUSD(&_VaultTest.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultTest *VaultTestCaller) MINFUNDINGRATEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "MIN_FUNDING_RATE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultTest *VaultTestSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _VaultTest.Contract.MINFUNDINGRATEINTERVAL(&_VaultTest.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _VaultTest.Contract.MINFUNDINGRATEINTERVAL(&_VaultTest.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultTest *VaultTestCaller) MINLEVERAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "MIN_LEVERAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultTest *VaultTestSession) MINLEVERAGE() (*big.Int, error) {
	return _VaultTest.Contract.MINLEVERAGE(&_VaultTest.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MINLEVERAGE() (*big.Int, error) {
	return _VaultTest.Contract.MINLEVERAGE(&_VaultTest.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultTest *VaultTestCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultTest *VaultTestSession) PRICEPRECISION() (*big.Int, error) {
	return _VaultTest.Contract.PRICEPRECISION(&_VaultTest.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _VaultTest.Contract.PRICEPRECISION(&_VaultTest.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultTest *VaultTestCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultTest *VaultTestSession) USDGDECIMALS() (*big.Int, error) {
	return _VaultTest.Contract.USDGDECIMALS(&_VaultTest.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _VaultTest.Contract.USDGDECIMALS(&_VaultTest.CallOpts)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultTest *VaultTestCaller) AdjustForDecimals(opts *bind.CallOpts, _amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "adjustForDecimals", _amount, _tokenDiv, _tokenMul)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultTest *VaultTestSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _VaultTest.Contract.AdjustForDecimals(&_VaultTest.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _VaultTest.Contract.AdjustForDecimals(&_VaultTest.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultTest *VaultTestCaller) AllWhitelistedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "allWhitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultTest *VaultTestSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _VaultTest.Contract.AllWhitelistedTokens(&_VaultTest.CallOpts, arg0)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultTest *VaultTestCallerSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _VaultTest.Contract.AllWhitelistedTokens(&_VaultTest.CallOpts, arg0)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultTest *VaultTestCaller) AllWhitelistedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "allWhitelistedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultTest *VaultTestSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _VaultTest.Contract.AllWhitelistedTokensLength(&_VaultTest.CallOpts)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _VaultTest.Contract.AllWhitelistedTokensLength(&_VaultTest.CallOpts)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultTest *VaultTestCaller) ApprovedRouters(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "approvedRouters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultTest *VaultTestSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultTest.Contract.ApprovedRouters(&_VaultTest.CallOpts, arg0, arg1)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultTest *VaultTestCallerSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultTest.Contract.ApprovedRouters(&_VaultTest.CallOpts, arg0, arg1)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) BufferAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "bufferAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.BufferAmounts(&_VaultTest.CallOpts, arg0)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.BufferAmounts(&_VaultTest.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) CumulativeFundingRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "cumulativeFundingRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.CumulativeFundingRates(&_VaultTest.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.CumulativeFundingRates(&_VaultTest.CallOpts, arg0)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultTest *VaultTestCaller) ErrorController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "errorController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultTest *VaultTestSession) ErrorController() (common.Address, error) {
	return _VaultTest.Contract.ErrorController(&_VaultTest.CallOpts)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultTest *VaultTestCallerSession) ErrorController() (common.Address, error) {
	return _VaultTest.Contract.ErrorController(&_VaultTest.CallOpts)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultTest *VaultTestCaller) Errors(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "errors", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultTest *VaultTestSession) Errors(arg0 *big.Int) (string, error) {
	return _VaultTest.Contract.Errors(&_VaultTest.CallOpts, arg0)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultTest *VaultTestCallerSession) Errors(arg0 *big.Int) (string, error) {
	return _VaultTest.Contract.Errors(&_VaultTest.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.FeeReserves(&_VaultTest.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.FeeReserves(&_VaultTest.CallOpts, arg0)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultTest *VaultTestCaller) FundingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "fundingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultTest *VaultTestSession) FundingInterval() (*big.Int, error) {
	return _VaultTest.Contract.FundingInterval(&_VaultTest.CallOpts)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) FundingInterval() (*big.Int, error) {
	return _VaultTest.Contract.FundingInterval(&_VaultTest.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultTest *VaultTestCaller) FundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "fundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultTest *VaultTestSession) FundingRateFactor() (*big.Int, error) {
	return _VaultTest.Contract.FundingRateFactor(&_VaultTest.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) FundingRateFactor() (*big.Int, error) {
	return _VaultTest.Contract.FundingRateFactor(&_VaultTest.CallOpts)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_VaultTest *VaultTestCaller) GetDelta(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getDelta", _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_VaultTest *VaultTestSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _VaultTest.Contract.GetDelta(&_VaultTest.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_VaultTest *VaultTestCallerSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _VaultTest.Contract.GetDelta(&_VaultTest.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultTest *VaultTestSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultTest.Contract.GetEntryFundingRate(&_VaultTest.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultTest.Contract.GetEntryFundingRate(&_VaultTest.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultTest *VaultTestSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultTest.Contract.GetFeeBasisPoints(&_VaultTest.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultTest.Contract.GetFeeBasisPoints(&_VaultTest.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetFundingFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getFundingFee", _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultTest *VaultTestSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetFundingFee(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetFundingFee(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_VaultTest *VaultTestCaller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getGlobalShortDelta", _token)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_VaultTest *VaultTestSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _VaultTest.Contract.GetGlobalShortDelta(&_VaultTest.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_VaultTest *VaultTestCallerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _VaultTest.Contract.GetGlobalShortDelta(&_VaultTest.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetMaxPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getMaxPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultTest *VaultTestSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetMaxPrice(&_VaultTest.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetMaxPrice(&_VaultTest.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetMinPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getMinPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultTest *VaultTestSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetMinPrice(&_VaultTest.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetMinPrice(&_VaultTest.CallOpts, _token)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetNextAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getNextAveragePrice", _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultTest *VaultTestSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetNextAveragePrice(&_VaultTest.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetNextAveragePrice(&_VaultTest.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetNextFundingRate(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getNextFundingRate", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultTest *VaultTestSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetNextFundingRate(&_VaultTest.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetNextFundingRate(&_VaultTest.CallOpts, _token)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetNextGlobalShortAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getNextGlobalShortAveragePrice", _indexToken, _nextPrice, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultTest *VaultTestSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetNextGlobalShortAveragePrice(&_VaultTest.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetNextGlobalShortAveragePrice(&_VaultTest.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_VaultTest *VaultTestCaller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_VaultTest *VaultTestSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _VaultTest.Contract.GetPosition(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_VaultTest *VaultTestCallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _VaultTest.Contract.GetPosition(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_VaultTest *VaultTestCaller) GetPositionDelta(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getPositionDelta", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_VaultTest *VaultTestSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _VaultTest.Contract.GetPositionDelta(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_VaultTest *VaultTestCallerSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _VaultTest.Contract.GetPositionDelta(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetPositionFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getPositionFee", _account, _collateralToken, _indexToken, _isLong, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultTest *VaultTestSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetPositionFee(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetPositionFee(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultTest *VaultTestCaller) GetPositionKey(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getPositionKey", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultTest *VaultTestSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _VaultTest.Contract.GetPositionKey(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultTest *VaultTestCallerSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _VaultTest.Contract.GetPositionKey(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetPositionLeverage(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getPositionLeverage", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultTest *VaultTestSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultTest.Contract.GetPositionLeverage(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultTest.Contract.GetPositionLeverage(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetRedemptionAmount(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getRedemptionAmount", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultTest *VaultTestSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetRedemptionAmount(&_VaultTest.CallOpts, _token, _usdgAmount)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.GetRedemptionAmount(&_VaultTest.CallOpts, _token, _usdgAmount)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetRedemptionCollateral(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getRedemptionCollateral", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultTest *VaultTestSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetRedemptionCollateral(&_VaultTest.CallOpts, _token)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetRedemptionCollateral(&_VaultTest.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetRedemptionCollateralUsd(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getRedemptionCollateralUsd", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultTest *VaultTestSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetRedemptionCollateralUsd(&_VaultTest.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetRedemptionCollateralUsd(&_VaultTest.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetTargetUsdgAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getTargetUsdgAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultTest *VaultTestSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetTargetUsdgAmount(&_VaultTest.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetTargetUsdgAmount(&_VaultTest.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultTest *VaultTestCaller) GetUtilisation(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "getUtilisation", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultTest *VaultTestSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetUtilisation(&_VaultTest.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GetUtilisation(&_VaultTest.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) GlobalShortAveragePrices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "globalShortAveragePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GlobalShortAveragePrices(&_VaultTest.CallOpts, arg0)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GlobalShortAveragePrices(&_VaultTest.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) GlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "globalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GlobalShortSizes(&_VaultTest.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GlobalShortSizes(&_VaultTest.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultTest *VaultTestCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultTest *VaultTestSession) Gov() (common.Address, error) {
	return _VaultTest.Contract.Gov(&_VaultTest.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultTest *VaultTestCallerSession) Gov() (common.Address, error) {
	return _VaultTest.Contract.Gov(&_VaultTest.CallOpts)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) GuaranteedUsd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "guaranteedUsd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GuaranteedUsd(&_VaultTest.CallOpts, arg0)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.GuaranteedUsd(&_VaultTest.CallOpts, arg0)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultTest *VaultTestCaller) HasDynamicFees(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "hasDynamicFees")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultTest *VaultTestSession) HasDynamicFees() (bool, error) {
	return _VaultTest.Contract.HasDynamicFees(&_VaultTest.CallOpts)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultTest *VaultTestCallerSession) HasDynamicFees() (bool, error) {
	return _VaultTest.Contract.HasDynamicFees(&_VaultTest.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultTest *VaultTestCaller) InManagerMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "inManagerMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultTest *VaultTestSession) InManagerMode() (bool, error) {
	return _VaultTest.Contract.InManagerMode(&_VaultTest.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultTest *VaultTestCallerSession) InManagerMode() (bool, error) {
	return _VaultTest.Contract.InManagerMode(&_VaultTest.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultTest *VaultTestCaller) InPrivateLiquidationMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "inPrivateLiquidationMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultTest *VaultTestSession) InPrivateLiquidationMode() (bool, error) {
	return _VaultTest.Contract.InPrivateLiquidationMode(&_VaultTest.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultTest *VaultTestCallerSession) InPrivateLiquidationMode() (bool, error) {
	return _VaultTest.Contract.InPrivateLiquidationMode(&_VaultTest.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultTest *VaultTestCaller) IncludeAmmPrice(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "includeAmmPrice")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultTest *VaultTestSession) IncludeAmmPrice() (bool, error) {
	return _VaultTest.Contract.IncludeAmmPrice(&_VaultTest.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultTest *VaultTestCallerSession) IncludeAmmPrice() (bool, error) {
	return _VaultTest.Contract.IncludeAmmPrice(&_VaultTest.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultTest *VaultTestCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultTest *VaultTestSession) IsInitialized() (bool, error) {
	return _VaultTest.Contract.IsInitialized(&_VaultTest.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultTest *VaultTestCallerSession) IsInitialized() (bool, error) {
	return _VaultTest.Contract.IsInitialized(&_VaultTest.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultTest *VaultTestCaller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultTest *VaultTestSession) IsLeverageEnabled() (bool, error) {
	return _VaultTest.Contract.IsLeverageEnabled(&_VaultTest.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultTest *VaultTestCallerSession) IsLeverageEnabled() (bool, error) {
	return _VaultTest.Contract.IsLeverageEnabled(&_VaultTest.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultTest *VaultTestCaller) IsLiquidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "isLiquidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultTest *VaultTestSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.IsLiquidator(&_VaultTest.CallOpts, arg0)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultTest *VaultTestCallerSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.IsLiquidator(&_VaultTest.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultTest *VaultTestCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultTest *VaultTestSession) IsManager(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.IsManager(&_VaultTest.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultTest *VaultTestCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.IsManager(&_VaultTest.CallOpts, arg0)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultTest *VaultTestCaller) IsSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "isSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultTest *VaultTestSession) IsSwapEnabled() (bool, error) {
	return _VaultTest.Contract.IsSwapEnabled(&_VaultTest.CallOpts)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultTest *VaultTestCallerSession) IsSwapEnabled() (bool, error) {
	return _VaultTest.Contract.IsSwapEnabled(&_VaultTest.CallOpts)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) LastFundingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "lastFundingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.LastFundingTimes(&_VaultTest.CallOpts, arg0)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.LastFundingTimes(&_VaultTest.CallOpts, arg0)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultTest *VaultTestCaller) LiquidationFeeUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "liquidationFeeUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultTest *VaultTestSession) LiquidationFeeUsd() (*big.Int, error) {
	return _VaultTest.Contract.LiquidationFeeUsd(&_VaultTest.CallOpts)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) LiquidationFeeUsd() (*big.Int, error) {
	return _VaultTest.Contract.LiquidationFeeUsd(&_VaultTest.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCaller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.MarginFeeBasisPoints(&_VaultTest.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.MarginFeeBasisPoints(&_VaultTest.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultTest *VaultTestCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultTest *VaultTestSession) MaxGasPrice() (*big.Int, error) {
	return _VaultTest.Contract.MaxGasPrice(&_VaultTest.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MaxGasPrice() (*big.Int, error) {
	return _VaultTest.Contract.MaxGasPrice(&_VaultTest.CallOpts)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.MaxGlobalShortSizes(&_VaultTest.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.MaxGlobalShortSizes(&_VaultTest.CallOpts, arg0)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultTest *VaultTestCaller) MaxLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "maxLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultTest *VaultTestSession) MaxLeverage() (*big.Int, error) {
	return _VaultTest.Contract.MaxLeverage(&_VaultTest.CallOpts)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MaxLeverage() (*big.Int, error) {
	return _VaultTest.Contract.MaxLeverage(&_VaultTest.CallOpts)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) MaxUsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "maxUsdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.MaxUsdgAmounts(&_VaultTest.CallOpts, arg0)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.MaxUsdgAmounts(&_VaultTest.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) MinProfitBasisPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "minProfitBasisPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.MinProfitBasisPoints(&_VaultTest.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.MinProfitBasisPoints(&_VaultTest.CallOpts, arg0)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultTest *VaultTestCaller) MinProfitTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "minProfitTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultTest *VaultTestSession) MinProfitTime() (*big.Int, error) {
	return _VaultTest.Contract.MinProfitTime(&_VaultTest.CallOpts)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MinProfitTime() (*big.Int, error) {
	return _VaultTest.Contract.MinProfitTime(&_VaultTest.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCaller) MintBurnFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "mintBurnFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.MintBurnFeeBasisPoints(&_VaultTest.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.MintBurnFeeBasisPoints(&_VaultTest.CallOpts)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) PoolAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "poolAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.PoolAmounts(&_VaultTest.CallOpts, arg0)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.PoolAmounts(&_VaultTest.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_VaultTest *VaultTestCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Size              *big.Int
		Collateral        *big.Int
		AveragePrice      *big.Int
		EntryFundingRate  *big.Int
		ReserveAmount     *big.Int
		RealisedPnl       *big.Int
		LastIncreasedTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Size = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Collateral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AveragePrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EntryFundingRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RealisedPnl = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastIncreasedTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_VaultTest *VaultTestSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _VaultTest.Contract.Positions(&_VaultTest.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_VaultTest *VaultTestCallerSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _VaultTest.Contract.Positions(&_VaultTest.CallOpts, arg0)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultTest *VaultTestCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultTest *VaultTestSession) PriceFeed() (common.Address, error) {
	return _VaultTest.Contract.PriceFeed(&_VaultTest.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultTest *VaultTestCallerSession) PriceFeed() (common.Address, error) {
	return _VaultTest.Contract.PriceFeed(&_VaultTest.CallOpts)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) ReservedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "reservedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.ReservedAmounts(&_VaultTest.CallOpts, arg0)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.ReservedAmounts(&_VaultTest.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultTest *VaultTestCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultTest *VaultTestSession) Router() (common.Address, error) {
	return _VaultTest.Contract.Router(&_VaultTest.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultTest *VaultTestCallerSession) Router() (common.Address, error) {
	return _VaultTest.Contract.Router(&_VaultTest.CallOpts)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultTest *VaultTestCaller) ShortableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "shortableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultTest *VaultTestSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.ShortableTokens(&_VaultTest.CallOpts, arg0)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultTest *VaultTestCallerSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.ShortableTokens(&_VaultTest.CallOpts, arg0)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultTest *VaultTestCaller) StableFundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "stableFundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultTest *VaultTestSession) StableFundingRateFactor() (*big.Int, error) {
	return _VaultTest.Contract.StableFundingRateFactor(&_VaultTest.CallOpts)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) StableFundingRateFactor() (*big.Int, error) {
	return _VaultTest.Contract.StableFundingRateFactor(&_VaultTest.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCaller) StableSwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "stableSwapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.StableSwapFeeBasisPoints(&_VaultTest.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.StableSwapFeeBasisPoints(&_VaultTest.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCaller) StableTaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "stableTaxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestSession) StableTaxBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.StableTaxBasisPoints(&_VaultTest.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) StableTaxBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.StableTaxBasisPoints(&_VaultTest.CallOpts)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultTest *VaultTestCaller) StableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "stableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultTest *VaultTestSession) StableTokens(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.StableTokens(&_VaultTest.CallOpts, arg0)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultTest *VaultTestCallerSession) StableTokens(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.StableTokens(&_VaultTest.CallOpts, arg0)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCaller) SwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "swapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.SwapFeeBasisPoints(&_VaultTest.CallOpts)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.SwapFeeBasisPoints(&_VaultTest.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCaller) TaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "taxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestSession) TaxBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.TaxBasisPoints(&_VaultTest.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) TaxBasisPoints() (*big.Int, error) {
	return _VaultTest.Contract.TaxBasisPoints(&_VaultTest.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.TokenBalances(&_VaultTest.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.TokenBalances(&_VaultTest.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) TokenDecimals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "tokenDecimals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.TokenDecimals(&_VaultTest.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.TokenDecimals(&_VaultTest.CallOpts, arg0)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultTest *VaultTestCaller) TokenToUsdMin(opts *bind.CallOpts, _token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "tokenToUsdMin", _token, _tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultTest *VaultTestSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.TokenToUsdMin(&_VaultTest.CallOpts, _token, _tokenAmount)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.TokenToUsdMin(&_VaultTest.CallOpts, _token, _tokenAmount)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) TokenWeights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "tokenWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.TokenWeights(&_VaultTest.CallOpts, arg0)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.TokenWeights(&_VaultTest.CallOpts, arg0)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultTest *VaultTestCaller) TotalTokenWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "totalTokenWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultTest *VaultTestSession) TotalTokenWeights() (*big.Int, error) {
	return _VaultTest.Contract.TotalTokenWeights(&_VaultTest.CallOpts)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) TotalTokenWeights() (*big.Int, error) {
	return _VaultTest.Contract.TotalTokenWeights(&_VaultTest.CallOpts)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultTest *VaultTestCaller) UsdToToken(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "usdToToken", _token, _usdAmount, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultTest *VaultTestSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.UsdToToken(&_VaultTest.CallOpts, _token, _usdAmount, _price)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.UsdToToken(&_VaultTest.CallOpts, _token, _usdAmount, _price)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultTest *VaultTestCaller) UsdToTokenMax(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "usdToTokenMax", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultTest *VaultTestSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.UsdToTokenMax(&_VaultTest.CallOpts, _token, _usdAmount)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.UsdToTokenMax(&_VaultTest.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultTest *VaultTestCaller) UsdToTokenMin(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "usdToTokenMin", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultTest *VaultTestSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.UsdToTokenMin(&_VaultTest.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultTest.Contract.UsdToTokenMin(&_VaultTest.CallOpts, _token, _usdAmount)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultTest *VaultTestCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultTest *VaultTestSession) Usdg() (common.Address, error) {
	return _VaultTest.Contract.Usdg(&_VaultTest.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultTest *VaultTestCallerSession) Usdg() (common.Address, error) {
	return _VaultTest.Contract.Usdg(&_VaultTest.CallOpts)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCaller) UsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "usdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.UsdgAmounts(&_VaultTest.CallOpts, arg0)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultTest *VaultTestCallerSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultTest.Contract.UsdgAmounts(&_VaultTest.CallOpts, arg0)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultTest *VaultTestCaller) UseSwapPricing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "useSwapPricing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultTest *VaultTestSession) UseSwapPricing() (bool, error) {
	return _VaultTest.Contract.UseSwapPricing(&_VaultTest.CallOpts)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultTest *VaultTestCallerSession) UseSwapPricing() (bool, error) {
	return _VaultTest.Contract.UseSwapPricing(&_VaultTest.CallOpts)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultTest *VaultTestCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

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
func (_VaultTest *VaultTestSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultTest.Contract.ValidateLiquidation(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultTest *VaultTestCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultTest.Contract.ValidateLiquidation(&_VaultTest.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultTest *VaultTestCaller) VaultUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "vaultUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultTest *VaultTestSession) VaultUtils() (common.Address, error) {
	return _VaultTest.Contract.VaultUtils(&_VaultTest.CallOpts)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultTest *VaultTestCallerSession) VaultUtils() (common.Address, error) {
	return _VaultTest.Contract.VaultUtils(&_VaultTest.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultTest *VaultTestCaller) WhitelistedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "whitelistedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultTest *VaultTestSession) WhitelistedTokenCount() (*big.Int, error) {
	return _VaultTest.Contract.WhitelistedTokenCount(&_VaultTest.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultTest *VaultTestCallerSession) WhitelistedTokenCount() (*big.Int, error) {
	return _VaultTest.Contract.WhitelistedTokenCount(&_VaultTest.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultTest *VaultTestCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultTest.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultTest *VaultTestSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.WhitelistedTokens(&_VaultTest.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultTest *VaultTestCallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _VaultTest.Contract.WhitelistedTokens(&_VaultTest.CallOpts, arg0)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultTest *VaultTestTransactor) AddRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "addRouter", _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultTest *VaultTestSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.AddRouter(&_VaultTest.TransactOpts, _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultTest *VaultTestTransactorSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.AddRouter(&_VaultTest.TransactOpts, _router)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactor) BuyUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "buyUSDG", _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.BuyUSDG(&_VaultTest.TransactOpts, _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactorSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.BuyUSDG(&_VaultTest.TransactOpts, _token, _receiver)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultTest *VaultTestTransactor) ClearTokenConfig(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "clearTokenConfig", _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultTest *VaultTestSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.ClearTokenConfig(&_VaultTest.TransactOpts, _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultTest *VaultTestTransactorSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.ClearTokenConfig(&_VaultTest.TransactOpts, _token)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactor) DecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "decreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultTest *VaultTestSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.DecreasePosition(&_VaultTest.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactorSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.DecreasePosition(&_VaultTest.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultTest *VaultTestTransactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "directPoolDeposit", _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultTest *VaultTestSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.DirectPoolDeposit(&_VaultTest.TransactOpts, _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultTest *VaultTestTransactorSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.DirectPoolDeposit(&_VaultTest.TransactOpts, _token)
}

// IncreaseGlobalShortSize is a paid mutator transaction binding the contract method 0x6a6f3424.
//
// Solidity: function increaseGlobalShortSize(address token, uint256 amount) returns()
func (_VaultTest *VaultTestTransactor) IncreaseGlobalShortSize(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "increaseGlobalShortSize", token, amount)
}

// IncreaseGlobalShortSize is a paid mutator transaction binding the contract method 0x6a6f3424.
//
// Solidity: function increaseGlobalShortSize(address token, uint256 amount) returns()
func (_VaultTest *VaultTestSession) IncreaseGlobalShortSize(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.IncreaseGlobalShortSize(&_VaultTest.TransactOpts, token, amount)
}

// IncreaseGlobalShortSize is a paid mutator transaction binding the contract method 0x6a6f3424.
//
// Solidity: function increaseGlobalShortSize(address token, uint256 amount) returns()
func (_VaultTest *VaultTestTransactorSession) IncreaseGlobalShortSize(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.IncreaseGlobalShortSize(&_VaultTest.TransactOpts, token, amount)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultTest *VaultTestTransactor) IncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "increasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultTest *VaultTestSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultTest.Contract.IncreasePosition(&_VaultTest.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultTest *VaultTestTransactorSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultTest.Contract.IncreasePosition(&_VaultTest.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultTest *VaultTestTransactor) Initialize(opts *bind.TransactOpts, _router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "initialize", _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultTest *VaultTestSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.Initialize(&_VaultTest.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultTest *VaultTestTransactorSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.Initialize(&_VaultTest.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultTest *VaultTestTransactor) LiquidatePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "liquidatePosition", _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultTest *VaultTestSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.LiquidatePosition(&_VaultTest.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultTest *VaultTestTransactorSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.LiquidatePosition(&_VaultTest.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultTest *VaultTestTransactor) RemoveRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "removeRouter", _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultTest *VaultTestSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.RemoveRouter(&_VaultTest.TransactOpts, _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultTest *VaultTestTransactorSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.RemoveRouter(&_VaultTest.TransactOpts, _router)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactor) SellUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "sellUSDG", _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SellUSDG(&_VaultTest.TransactOpts, _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactorSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SellUSDG(&_VaultTest.TransactOpts, _token, _receiver)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactor) SetBufferAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setBufferAmount", _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetBufferAmount(&_VaultTest.TransactOpts, _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactorSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetBufferAmount(&_VaultTest.TransactOpts, _token, _amount)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultTest *VaultTestTransactor) SetError(opts *bind.TransactOpts, _errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setError", _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultTest *VaultTestSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultTest.Contract.SetError(&_VaultTest.TransactOpts, _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultTest *VaultTestTransactorSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultTest.Contract.SetError(&_VaultTest.TransactOpts, _errorCode, _error)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultTest *VaultTestTransactor) SetErrorController(opts *bind.TransactOpts, _errorController common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setErrorController", _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultTest *VaultTestSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetErrorController(&_VaultTest.TransactOpts, _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultTest *VaultTestTransactorSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetErrorController(&_VaultTest.TransactOpts, _errorController)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultTest *VaultTestTransactor) SetFees(opts *bind.TransactOpts, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setFees", _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultTest *VaultTestSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetFees(&_VaultTest.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultTest *VaultTestTransactorSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetFees(&_VaultTest.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultTest *VaultTestTransactor) SetFundingRate(opts *bind.TransactOpts, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setFundingRate", _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultTest *VaultTestSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetFundingRate(&_VaultTest.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultTest *VaultTestTransactorSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetFundingRate(&_VaultTest.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultTest *VaultTestTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultTest *VaultTestSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetGov(&_VaultTest.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultTest *VaultTestTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetGov(&_VaultTest.TransactOpts, _gov)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultTest *VaultTestTransactor) SetInManagerMode(opts *bind.TransactOpts, _inManagerMode bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setInManagerMode", _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultTest *VaultTestSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetInManagerMode(&_VaultTest.TransactOpts, _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultTest *VaultTestTransactorSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetInManagerMode(&_VaultTest.TransactOpts, _inManagerMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultTest *VaultTestTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setInPrivateLiquidationMode", _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultTest *VaultTestSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetInPrivateLiquidationMode(&_VaultTest.TransactOpts, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultTest *VaultTestTransactorSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetInPrivateLiquidationMode(&_VaultTest.TransactOpts, _inPrivateLiquidationMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultTest *VaultTestTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultTest *VaultTestSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetIsLeverageEnabled(&_VaultTest.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultTest *VaultTestTransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetIsLeverageEnabled(&_VaultTest.TransactOpts, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultTest *VaultTestTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setIsSwapEnabled", _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultTest *VaultTestSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetIsSwapEnabled(&_VaultTest.TransactOpts, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultTest *VaultTestTransactorSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetIsSwapEnabled(&_VaultTest.TransactOpts, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultTest *VaultTestTransactor) SetLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setLiquidator", _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultTest *VaultTestSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetLiquidator(&_VaultTest.TransactOpts, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultTest *VaultTestTransactorSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetLiquidator(&_VaultTest.TransactOpts, _liquidator, _isActive)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultTest *VaultTestTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setManager", _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultTest *VaultTestSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetManager(&_VaultTest.TransactOpts, _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultTest *VaultTestTransactorSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetManager(&_VaultTest.TransactOpts, _manager, _isManager)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultTest *VaultTestTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setMaxGasPrice", _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultTest *VaultTestSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetMaxGasPrice(&_VaultTest.TransactOpts, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultTest *VaultTestTransactorSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetMaxGasPrice(&_VaultTest.TransactOpts, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setMaxGlobalShortSize", _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetMaxGlobalShortSize(&_VaultTest.TransactOpts, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactorSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetMaxGlobalShortSize(&_VaultTest.TransactOpts, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultTest *VaultTestTransactor) SetMaxLeverage(opts *bind.TransactOpts, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setMaxLeverage", _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultTest *VaultTestSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetMaxLeverage(&_VaultTest.TransactOpts, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultTest *VaultTestTransactorSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetMaxLeverage(&_VaultTest.TransactOpts, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultTest *VaultTestTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultTest *VaultTestSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetPriceFeed(&_VaultTest.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultTest *VaultTestTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetPriceFeed(&_VaultTest.TransactOpts, _priceFeed)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultTest *VaultTestTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setTokenConfig", _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultTest *VaultTestSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetTokenConfig(&_VaultTest.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultTest *VaultTestTransactorSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultTest.Contract.SetTokenConfig(&_VaultTest.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactor) SetUsdgAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setUsdgAmount", _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetUsdgAmount(&_VaultTest.TransactOpts, _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactorSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.SetUsdgAmount(&_VaultTest.TransactOpts, _token, _amount)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultTest *VaultTestTransactor) SetVaultUtils(opts *bind.TransactOpts, _vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "setVaultUtils", _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultTest *VaultTestSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetVaultUtils(&_VaultTest.TransactOpts, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultTest *VaultTestTransactorSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.SetVaultUtils(&_VaultTest.TransactOpts, _vaultUtils)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultTest *VaultTestSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.Swap(&_VaultTest.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.Swap(&_VaultTest.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultTest *VaultTestTransactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "updateCumulativeFundingRate", _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultTest *VaultTestSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.UpdateCumulativeFundingRate(&_VaultTest.TransactOpts, _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultTest *VaultTestTransactorSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.UpdateCumulativeFundingRate(&_VaultTest.TransactOpts, _collateralToken, _indexToken)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactor) UpgradeVault(opts *bind.TransactOpts, _newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "upgradeVault", _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.UpgradeVault(&_VaultTest.TransactOpts, _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultTest *VaultTestTransactorSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultTest.Contract.UpgradeVault(&_VaultTest.TransactOpts, _newVault, _token, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.WithdrawFees(&_VaultTest.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultTest *VaultTestTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultTest.Contract.WithdrawFees(&_VaultTest.TransactOpts, _token, _receiver)
}

// VaultTestBuyUSDGIterator is returned from FilterBuyUSDG and is used to iterate over the raw logs and unpacked data for BuyUSDG events raised by the VaultTest contract.
type VaultTestBuyUSDGIterator struct {
	Event *VaultTestBuyUSDG // Event containing the contract specifics and raw log

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
func (it *VaultTestBuyUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestBuyUSDG)
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
		it.Event = new(VaultTestBuyUSDG)
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
func (it *VaultTestBuyUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestBuyUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestBuyUSDG represents a BuyUSDG event raised by the VaultTest contract.
type VaultTestBuyUSDG struct {
	Account        common.Address
	Token          common.Address
	TokenAmount    *big.Int
	UsdgAmount     *big.Int
	FeeBasisPoints *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBuyUSDG is a free log retrieval operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) FilterBuyUSDG(opts *bind.FilterOpts) (*VaultTestBuyUSDGIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultTestBuyUSDGIterator{contract: _VaultTest.contract, event: "BuyUSDG", logs: logs, sub: sub}, nil
}

// WatchBuyUSDG is a free log subscription operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) WatchBuyUSDG(opts *bind.WatchOpts, sink chan<- *VaultTestBuyUSDG) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestBuyUSDG)
				if err := _VaultTest.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
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

// ParseBuyUSDG is a log parse operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) ParseBuyUSDG(log types.Log) (*VaultTestBuyUSDG, error) {
	event := new(VaultTestBuyUSDG)
	if err := _VaultTest.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestClosePositionIterator is returned from FilterClosePosition and is used to iterate over the raw logs and unpacked data for ClosePosition events raised by the VaultTest contract.
type VaultTestClosePositionIterator struct {
	Event *VaultTestClosePosition // Event containing the contract specifics and raw log

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
func (it *VaultTestClosePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestClosePosition)
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
		it.Event = new(VaultTestClosePosition)
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
func (it *VaultTestClosePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestClosePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestClosePosition represents a ClosePosition event raised by the VaultTest contract.
type VaultTestClosePosition struct {
	Key              [32]byte
	Size             *big.Int
	Collateral       *big.Int
	AveragePrice     *big.Int
	EntryFundingRate *big.Int
	ReserveAmount    *big.Int
	RealisedPnl      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClosePosition is a free log retrieval operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_VaultTest *VaultTestFilterer) FilterClosePosition(opts *bind.FilterOpts) (*VaultTestClosePositionIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return &VaultTestClosePositionIterator{contract: _VaultTest.contract, event: "ClosePosition", logs: logs, sub: sub}, nil
}

// WatchClosePosition is a free log subscription operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_VaultTest *VaultTestFilterer) WatchClosePosition(opts *bind.WatchOpts, sink chan<- *VaultTestClosePosition) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestClosePosition)
				if err := _VaultTest.contract.UnpackLog(event, "ClosePosition", log); err != nil {
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

// ParseClosePosition is a log parse operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_VaultTest *VaultTestFilterer) ParseClosePosition(log types.Log) (*VaultTestClosePosition, error) {
	event := new(VaultTestClosePosition)
	if err := _VaultTest.contract.UnpackLog(event, "ClosePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestCollectMarginFeesIterator is returned from FilterCollectMarginFees and is used to iterate over the raw logs and unpacked data for CollectMarginFees events raised by the VaultTest contract.
type VaultTestCollectMarginFeesIterator struct {
	Event *VaultTestCollectMarginFees // Event containing the contract specifics and raw log

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
func (it *VaultTestCollectMarginFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestCollectMarginFees)
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
		it.Event = new(VaultTestCollectMarginFees)
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
func (it *VaultTestCollectMarginFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestCollectMarginFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestCollectMarginFees represents a CollectMarginFees event raised by the VaultTest contract.
type VaultTestCollectMarginFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectMarginFees is a free log retrieval operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultTest *VaultTestFilterer) FilterCollectMarginFees(opts *bind.FilterOpts) (*VaultTestCollectMarginFeesIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return &VaultTestCollectMarginFeesIterator{contract: _VaultTest.contract, event: "CollectMarginFees", logs: logs, sub: sub}, nil
}

// WatchCollectMarginFees is a free log subscription operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultTest *VaultTestFilterer) WatchCollectMarginFees(opts *bind.WatchOpts, sink chan<- *VaultTestCollectMarginFees) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestCollectMarginFees)
				if err := _VaultTest.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
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

// ParseCollectMarginFees is a log parse operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultTest *VaultTestFilterer) ParseCollectMarginFees(log types.Log) (*VaultTestCollectMarginFees, error) {
	event := new(VaultTestCollectMarginFees)
	if err := _VaultTest.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestCollectSwapFeesIterator is returned from FilterCollectSwapFees and is used to iterate over the raw logs and unpacked data for CollectSwapFees events raised by the VaultTest contract.
type VaultTestCollectSwapFeesIterator struct {
	Event *VaultTestCollectSwapFees // Event containing the contract specifics and raw log

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
func (it *VaultTestCollectSwapFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestCollectSwapFees)
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
		it.Event = new(VaultTestCollectSwapFees)
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
func (it *VaultTestCollectSwapFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestCollectSwapFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestCollectSwapFees represents a CollectSwapFees event raised by the VaultTest contract.
type VaultTestCollectSwapFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectSwapFees is a free log retrieval operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultTest *VaultTestFilterer) FilterCollectSwapFees(opts *bind.FilterOpts) (*VaultTestCollectSwapFeesIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return &VaultTestCollectSwapFeesIterator{contract: _VaultTest.contract, event: "CollectSwapFees", logs: logs, sub: sub}, nil
}

// WatchCollectSwapFees is a free log subscription operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultTest *VaultTestFilterer) WatchCollectSwapFees(opts *bind.WatchOpts, sink chan<- *VaultTestCollectSwapFees) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestCollectSwapFees)
				if err := _VaultTest.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
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

// ParseCollectSwapFees is a log parse operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultTest *VaultTestFilterer) ParseCollectSwapFees(log types.Log) (*VaultTestCollectSwapFees, error) {
	event := new(VaultTestCollectSwapFees)
	if err := _VaultTest.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestDecreaseGuaranteedUsdIterator is returned from FilterDecreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for DecreaseGuaranteedUsd events raised by the VaultTest contract.
type VaultTestDecreaseGuaranteedUsdIterator struct {
	Event *VaultTestDecreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultTestDecreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestDecreaseGuaranteedUsd)
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
		it.Event = new(VaultTestDecreaseGuaranteedUsd)
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
func (it *VaultTestDecreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestDecreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestDecreaseGuaranteedUsd represents a DecreaseGuaranteedUsd event raised by the VaultTest contract.
type VaultTestDecreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterDecreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultTestDecreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultTestDecreaseGuaranteedUsdIterator{contract: _VaultTest.contract, event: "DecreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchDecreaseGuaranteedUsd is a free log subscription operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchDecreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultTestDecreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestDecreaseGuaranteedUsd)
				if err := _VaultTest.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
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

// ParseDecreaseGuaranteedUsd is a log parse operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseDecreaseGuaranteedUsd(log types.Log) (*VaultTestDecreaseGuaranteedUsd, error) {
	event := new(VaultTestDecreaseGuaranteedUsd)
	if err := _VaultTest.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestDecreasePoolAmountIterator is returned from FilterDecreasePoolAmount and is used to iterate over the raw logs and unpacked data for DecreasePoolAmount events raised by the VaultTest contract.
type VaultTestDecreasePoolAmountIterator struct {
	Event *VaultTestDecreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultTestDecreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestDecreasePoolAmount)
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
		it.Event = new(VaultTestDecreasePoolAmount)
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
func (it *VaultTestDecreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestDecreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestDecreasePoolAmount represents a DecreasePoolAmount event raised by the VaultTest contract.
type VaultTestDecreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreasePoolAmount is a free log retrieval operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterDecreasePoolAmount(opts *bind.FilterOpts) (*VaultTestDecreasePoolAmountIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultTestDecreasePoolAmountIterator{contract: _VaultTest.contract, event: "DecreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchDecreasePoolAmount is a free log subscription operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchDecreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultTestDecreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestDecreasePoolAmount)
				if err := _VaultTest.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
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

// ParseDecreasePoolAmount is a log parse operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseDecreasePoolAmount(log types.Log) (*VaultTestDecreasePoolAmount, error) {
	event := new(VaultTestDecreasePoolAmount)
	if err := _VaultTest.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestDecreasePositionIterator is returned from FilterDecreasePosition and is used to iterate over the raw logs and unpacked data for DecreasePosition events raised by the VaultTest contract.
type VaultTestDecreasePositionIterator struct {
	Event *VaultTestDecreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultTestDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestDecreasePosition)
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
		it.Event = new(VaultTestDecreasePosition)
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
func (it *VaultTestDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestDecreasePosition represents a DecreasePosition event raised by the VaultTest contract.
type VaultTestDecreasePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Price           *big.Int
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDecreasePosition is a free log retrieval operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultTest *VaultTestFilterer) FilterDecreasePosition(opts *bind.FilterOpts) (*VaultTestDecreasePositionIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultTestDecreasePositionIterator{contract: _VaultTest.contract, event: "DecreasePosition", logs: logs, sub: sub}, nil
}

// WatchDecreasePosition is a free log subscription operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultTest *VaultTestFilterer) WatchDecreasePosition(opts *bind.WatchOpts, sink chan<- *VaultTestDecreasePosition) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestDecreasePosition)
				if err := _VaultTest.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
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

// ParseDecreasePosition is a log parse operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultTest *VaultTestFilterer) ParseDecreasePosition(log types.Log) (*VaultTestDecreasePosition, error) {
	event := new(VaultTestDecreasePosition)
	if err := _VaultTest.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestDecreaseReservedAmountIterator is returned from FilterDecreaseReservedAmount and is used to iterate over the raw logs and unpacked data for DecreaseReservedAmount events raised by the VaultTest contract.
type VaultTestDecreaseReservedAmountIterator struct {
	Event *VaultTestDecreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultTestDecreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestDecreaseReservedAmount)
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
		it.Event = new(VaultTestDecreaseReservedAmount)
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
func (it *VaultTestDecreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestDecreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestDecreaseReservedAmount represents a DecreaseReservedAmount event raised by the VaultTest contract.
type VaultTestDecreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseReservedAmount is a free log retrieval operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterDecreaseReservedAmount(opts *bind.FilterOpts) (*VaultTestDecreaseReservedAmountIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultTestDecreaseReservedAmountIterator{contract: _VaultTest.contract, event: "DecreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseReservedAmount is a free log subscription operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchDecreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultTestDecreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestDecreaseReservedAmount)
				if err := _VaultTest.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
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

// ParseDecreaseReservedAmount is a log parse operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseDecreaseReservedAmount(log types.Log) (*VaultTestDecreaseReservedAmount, error) {
	event := new(VaultTestDecreaseReservedAmount)
	if err := _VaultTest.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestDecreaseUsdgAmountIterator is returned from FilterDecreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for DecreaseUsdgAmount events raised by the VaultTest contract.
type VaultTestDecreaseUsdgAmountIterator struct {
	Event *VaultTestDecreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultTestDecreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestDecreaseUsdgAmount)
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
		it.Event = new(VaultTestDecreaseUsdgAmount)
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
func (it *VaultTestDecreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestDecreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestDecreaseUsdgAmount represents a DecreaseUsdgAmount event raised by the VaultTest contract.
type VaultTestDecreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseUsdgAmount is a free log retrieval operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterDecreaseUsdgAmount(opts *bind.FilterOpts) (*VaultTestDecreaseUsdgAmountIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultTestDecreaseUsdgAmountIterator{contract: _VaultTest.contract, event: "DecreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseUsdgAmount is a free log subscription operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchDecreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultTestDecreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestDecreaseUsdgAmount)
				if err := _VaultTest.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
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

// ParseDecreaseUsdgAmount is a log parse operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseDecreaseUsdgAmount(log types.Log) (*VaultTestDecreaseUsdgAmount, error) {
	event := new(VaultTestDecreaseUsdgAmount)
	if err := _VaultTest.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestDirectPoolDepositIterator is returned from FilterDirectPoolDeposit and is used to iterate over the raw logs and unpacked data for DirectPoolDeposit events raised by the VaultTest contract.
type VaultTestDirectPoolDepositIterator struct {
	Event *VaultTestDirectPoolDeposit // Event containing the contract specifics and raw log

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
func (it *VaultTestDirectPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestDirectPoolDeposit)
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
		it.Event = new(VaultTestDirectPoolDeposit)
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
func (it *VaultTestDirectPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestDirectPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestDirectPoolDeposit represents a DirectPoolDeposit event raised by the VaultTest contract.
type VaultTestDirectPoolDeposit struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDirectPoolDeposit is a free log retrieval operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterDirectPoolDeposit(opts *bind.FilterOpts) (*VaultTestDirectPoolDepositIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return &VaultTestDirectPoolDepositIterator{contract: _VaultTest.contract, event: "DirectPoolDeposit", logs: logs, sub: sub}, nil
}

// WatchDirectPoolDeposit is a free log subscription operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchDirectPoolDeposit(opts *bind.WatchOpts, sink chan<- *VaultTestDirectPoolDeposit) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestDirectPoolDeposit)
				if err := _VaultTest.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
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

// ParseDirectPoolDeposit is a log parse operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseDirectPoolDeposit(log types.Log) (*VaultTestDirectPoolDeposit, error) {
	event := new(VaultTestDirectPoolDeposit)
	if err := _VaultTest.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestIncreaseGuaranteedUsdIterator is returned from FilterIncreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for IncreaseGuaranteedUsd events raised by the VaultTest contract.
type VaultTestIncreaseGuaranteedUsdIterator struct {
	Event *VaultTestIncreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultTestIncreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestIncreaseGuaranteedUsd)
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
		it.Event = new(VaultTestIncreaseGuaranteedUsd)
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
func (it *VaultTestIncreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestIncreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestIncreaseGuaranteedUsd represents a IncreaseGuaranteedUsd event raised by the VaultTest contract.
type VaultTestIncreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterIncreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultTestIncreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultTestIncreaseGuaranteedUsdIterator{contract: _VaultTest.contract, event: "IncreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchIncreaseGuaranteedUsd is a free log subscription operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchIncreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultTestIncreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestIncreaseGuaranteedUsd)
				if err := _VaultTest.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
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

// ParseIncreaseGuaranteedUsd is a log parse operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseIncreaseGuaranteedUsd(log types.Log) (*VaultTestIncreaseGuaranteedUsd, error) {
	event := new(VaultTestIncreaseGuaranteedUsd)
	if err := _VaultTest.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestIncreasePoolAmountIterator is returned from FilterIncreasePoolAmount and is used to iterate over the raw logs and unpacked data for IncreasePoolAmount events raised by the VaultTest contract.
type VaultTestIncreasePoolAmountIterator struct {
	Event *VaultTestIncreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultTestIncreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestIncreasePoolAmount)
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
		it.Event = new(VaultTestIncreasePoolAmount)
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
func (it *VaultTestIncreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestIncreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestIncreasePoolAmount represents a IncreasePoolAmount event raised by the VaultTest contract.
type VaultTestIncreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreasePoolAmount is a free log retrieval operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterIncreasePoolAmount(opts *bind.FilterOpts) (*VaultTestIncreasePoolAmountIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultTestIncreasePoolAmountIterator{contract: _VaultTest.contract, event: "IncreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchIncreasePoolAmount is a free log subscription operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchIncreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultTestIncreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestIncreasePoolAmount)
				if err := _VaultTest.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
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

// ParseIncreasePoolAmount is a log parse operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseIncreasePoolAmount(log types.Log) (*VaultTestIncreasePoolAmount, error) {
	event := new(VaultTestIncreasePoolAmount)
	if err := _VaultTest.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestIncreasePositionIterator is returned from FilterIncreasePosition and is used to iterate over the raw logs and unpacked data for IncreasePosition events raised by the VaultTest contract.
type VaultTestIncreasePositionIterator struct {
	Event *VaultTestIncreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultTestIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestIncreasePosition)
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
		it.Event = new(VaultTestIncreasePosition)
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
func (it *VaultTestIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestIncreasePosition represents a IncreasePosition event raised by the VaultTest contract.
type VaultTestIncreasePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Price           *big.Int
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIncreasePosition is a free log retrieval operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultTest *VaultTestFilterer) FilterIncreasePosition(opts *bind.FilterOpts) (*VaultTestIncreasePositionIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultTestIncreasePositionIterator{contract: _VaultTest.contract, event: "IncreasePosition", logs: logs, sub: sub}, nil
}

// WatchIncreasePosition is a free log subscription operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultTest *VaultTestFilterer) WatchIncreasePosition(opts *bind.WatchOpts, sink chan<- *VaultTestIncreasePosition) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestIncreasePosition)
				if err := _VaultTest.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
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

// ParseIncreasePosition is a log parse operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultTest *VaultTestFilterer) ParseIncreasePosition(log types.Log) (*VaultTestIncreasePosition, error) {
	event := new(VaultTestIncreasePosition)
	if err := _VaultTest.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestIncreaseReservedAmountIterator is returned from FilterIncreaseReservedAmount and is used to iterate over the raw logs and unpacked data for IncreaseReservedAmount events raised by the VaultTest contract.
type VaultTestIncreaseReservedAmountIterator struct {
	Event *VaultTestIncreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultTestIncreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestIncreaseReservedAmount)
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
		it.Event = new(VaultTestIncreaseReservedAmount)
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
func (it *VaultTestIncreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestIncreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestIncreaseReservedAmount represents a IncreaseReservedAmount event raised by the VaultTest contract.
type VaultTestIncreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseReservedAmount is a free log retrieval operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterIncreaseReservedAmount(opts *bind.FilterOpts) (*VaultTestIncreaseReservedAmountIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultTestIncreaseReservedAmountIterator{contract: _VaultTest.contract, event: "IncreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseReservedAmount is a free log subscription operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchIncreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultTestIncreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestIncreaseReservedAmount)
				if err := _VaultTest.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
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

// ParseIncreaseReservedAmount is a log parse operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseIncreaseReservedAmount(log types.Log) (*VaultTestIncreaseReservedAmount, error) {
	event := new(VaultTestIncreaseReservedAmount)
	if err := _VaultTest.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestIncreaseUsdgAmountIterator is returned from FilterIncreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for IncreaseUsdgAmount events raised by the VaultTest contract.
type VaultTestIncreaseUsdgAmountIterator struct {
	Event *VaultTestIncreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultTestIncreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestIncreaseUsdgAmount)
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
		it.Event = new(VaultTestIncreaseUsdgAmount)
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
func (it *VaultTestIncreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestIncreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestIncreaseUsdgAmount represents a IncreaseUsdgAmount event raised by the VaultTest contract.
type VaultTestIncreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseUsdgAmount is a free log retrieval operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) FilterIncreaseUsdgAmount(opts *bind.FilterOpts) (*VaultTestIncreaseUsdgAmountIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultTestIncreaseUsdgAmountIterator{contract: _VaultTest.contract, event: "IncreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseUsdgAmount is a free log subscription operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) WatchIncreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultTestIncreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestIncreaseUsdgAmount)
				if err := _VaultTest.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
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

// ParseIncreaseUsdgAmount is a log parse operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_VaultTest *VaultTestFilterer) ParseIncreaseUsdgAmount(log types.Log) (*VaultTestIncreaseUsdgAmount, error) {
	event := new(VaultTestIncreaseUsdgAmount)
	if err := _VaultTest.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestLiquidatePositionIterator is returned from FilterLiquidatePosition and is used to iterate over the raw logs and unpacked data for LiquidatePosition events raised by the VaultTest contract.
type VaultTestLiquidatePositionIterator struct {
	Event *VaultTestLiquidatePosition // Event containing the contract specifics and raw log

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
func (it *VaultTestLiquidatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestLiquidatePosition)
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
		it.Event = new(VaultTestLiquidatePosition)
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
func (it *VaultTestLiquidatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestLiquidatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestLiquidatePosition represents a LiquidatePosition event raised by the VaultTest contract.
type VaultTestLiquidatePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	IsLong          bool
	Size            *big.Int
	Collateral      *big.Int
	ReserveAmount   *big.Int
	RealisedPnl     *big.Int
	MarkPrice       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidatePosition is a free log retrieval operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultTest *VaultTestFilterer) FilterLiquidatePosition(opts *bind.FilterOpts) (*VaultTestLiquidatePositionIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultTestLiquidatePositionIterator{contract: _VaultTest.contract, event: "LiquidatePosition", logs: logs, sub: sub}, nil
}

// WatchLiquidatePosition is a free log subscription operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultTest *VaultTestFilterer) WatchLiquidatePosition(opts *bind.WatchOpts, sink chan<- *VaultTestLiquidatePosition) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestLiquidatePosition)
				if err := _VaultTest.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
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

// ParseLiquidatePosition is a log parse operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultTest *VaultTestFilterer) ParseLiquidatePosition(log types.Log) (*VaultTestLiquidatePosition, error) {
	event := new(VaultTestLiquidatePosition)
	if err := _VaultTest.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestSellUSDGIterator is returned from FilterSellUSDG and is used to iterate over the raw logs and unpacked data for SellUSDG events raised by the VaultTest contract.
type VaultTestSellUSDGIterator struct {
	Event *VaultTestSellUSDG // Event containing the contract specifics and raw log

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
func (it *VaultTestSellUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestSellUSDG)
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
		it.Event = new(VaultTestSellUSDG)
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
func (it *VaultTestSellUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestSellUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestSellUSDG represents a SellUSDG event raised by the VaultTest contract.
type VaultTestSellUSDG struct {
	Account        common.Address
	Token          common.Address
	UsdgAmount     *big.Int
	TokenAmount    *big.Int
	FeeBasisPoints *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSellUSDG is a free log retrieval operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) FilterSellUSDG(opts *bind.FilterOpts) (*VaultTestSellUSDGIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultTestSellUSDGIterator{contract: _VaultTest.contract, event: "SellUSDG", logs: logs, sub: sub}, nil
}

// WatchSellUSDG is a free log subscription operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) WatchSellUSDG(opts *bind.WatchOpts, sink chan<- *VaultTestSellUSDG) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestSellUSDG)
				if err := _VaultTest.contract.UnpackLog(event, "SellUSDG", log); err != nil {
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

// ParseSellUSDG is a log parse operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) ParseSellUSDG(log types.Log) (*VaultTestSellUSDG, error) {
	event := new(VaultTestSellUSDG)
	if err := _VaultTest.contract.UnpackLog(event, "SellUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the VaultTest contract.
type VaultTestSwapIterator struct {
	Event *VaultTestSwap // Event containing the contract specifics and raw log

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
func (it *VaultTestSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestSwap)
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
		it.Event = new(VaultTestSwap)
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
func (it *VaultTestSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestSwap represents a Swap event raised by the VaultTest contract.
type VaultTestSwap struct {
	Account            common.Address
	TokenIn            common.Address
	TokenOut           common.Address
	AmountIn           *big.Int
	AmountOut          *big.Int
	AmountOutAfterFees *big.Int
	FeeBasisPoints     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) FilterSwap(opts *bind.FilterOpts) (*VaultTestSwapIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &VaultTestSwapIterator{contract: _VaultTest.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultTestSwap) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestSwap)
				if err := _VaultTest.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_VaultTest *VaultTestFilterer) ParseSwap(log types.Log) (*VaultTestSwap, error) {
	event := new(VaultTestSwap)
	if err := _VaultTest.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestUpdateFundingRateIterator is returned from FilterUpdateFundingRate and is used to iterate over the raw logs and unpacked data for UpdateFundingRate events raised by the VaultTest contract.
type VaultTestUpdateFundingRateIterator struct {
	Event *VaultTestUpdateFundingRate // Event containing the contract specifics and raw log

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
func (it *VaultTestUpdateFundingRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestUpdateFundingRate)
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
		it.Event = new(VaultTestUpdateFundingRate)
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
func (it *VaultTestUpdateFundingRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestUpdateFundingRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestUpdateFundingRate represents a UpdateFundingRate event raised by the VaultTest contract.
type VaultTestUpdateFundingRate struct {
	Token       common.Address
	FundingRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFundingRate is a free log retrieval operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_VaultTest *VaultTestFilterer) FilterUpdateFundingRate(opts *bind.FilterOpts) (*VaultTestUpdateFundingRateIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return &VaultTestUpdateFundingRateIterator{contract: _VaultTest.contract, event: "UpdateFundingRate", logs: logs, sub: sub}, nil
}

// WatchUpdateFundingRate is a free log subscription operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_VaultTest *VaultTestFilterer) WatchUpdateFundingRate(opts *bind.WatchOpts, sink chan<- *VaultTestUpdateFundingRate) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestUpdateFundingRate)
				if err := _VaultTest.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
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

// ParseUpdateFundingRate is a log parse operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_VaultTest *VaultTestFilterer) ParseUpdateFundingRate(log types.Log) (*VaultTestUpdateFundingRate, error) {
	event := new(VaultTestUpdateFundingRate)
	if err := _VaultTest.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestUpdatePnlIterator is returned from FilterUpdatePnl and is used to iterate over the raw logs and unpacked data for UpdatePnl events raised by the VaultTest contract.
type VaultTestUpdatePnlIterator struct {
	Event *VaultTestUpdatePnl // Event containing the contract specifics and raw log

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
func (it *VaultTestUpdatePnlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestUpdatePnl)
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
		it.Event = new(VaultTestUpdatePnl)
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
func (it *VaultTestUpdatePnlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestUpdatePnlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestUpdatePnl represents a UpdatePnl event raised by the VaultTest contract.
type VaultTestUpdatePnl struct {
	Key       [32]byte
	HasProfit bool
	Delta     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatePnl is a free log retrieval operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_VaultTest *VaultTestFilterer) FilterUpdatePnl(opts *bind.FilterOpts) (*VaultTestUpdatePnlIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return &VaultTestUpdatePnlIterator{contract: _VaultTest.contract, event: "UpdatePnl", logs: logs, sub: sub}, nil
}

// WatchUpdatePnl is a free log subscription operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_VaultTest *VaultTestFilterer) WatchUpdatePnl(opts *bind.WatchOpts, sink chan<- *VaultTestUpdatePnl) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestUpdatePnl)
				if err := _VaultTest.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
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

// ParseUpdatePnl is a log parse operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_VaultTest *VaultTestFilterer) ParseUpdatePnl(log types.Log) (*VaultTestUpdatePnl, error) {
	event := new(VaultTestUpdatePnl)
	if err := _VaultTest.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTestUpdatePositionIterator is returned from FilterUpdatePosition and is used to iterate over the raw logs and unpacked data for UpdatePosition events raised by the VaultTest contract.
type VaultTestUpdatePositionIterator struct {
	Event *VaultTestUpdatePosition // Event containing the contract specifics and raw log

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
func (it *VaultTestUpdatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTestUpdatePosition)
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
		it.Event = new(VaultTestUpdatePosition)
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
func (it *VaultTestUpdatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTestUpdatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTestUpdatePosition represents a UpdatePosition event raised by the VaultTest contract.
type VaultTestUpdatePosition struct {
	Key              [32]byte
	Size             *big.Int
	Collateral       *big.Int
	AveragePrice     *big.Int
	EntryFundingRate *big.Int
	ReserveAmount    *big.Int
	RealisedPnl      *big.Int
	MarkPrice        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdatePosition is a free log retrieval operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultTest *VaultTestFilterer) FilterUpdatePosition(opts *bind.FilterOpts) (*VaultTestUpdatePositionIterator, error) {

	logs, sub, err := _VaultTest.contract.FilterLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultTestUpdatePositionIterator{contract: _VaultTest.contract, event: "UpdatePosition", logs: logs, sub: sub}, nil
}

// WatchUpdatePosition is a free log subscription operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultTest *VaultTestFilterer) WatchUpdatePosition(opts *bind.WatchOpts, sink chan<- *VaultTestUpdatePosition) (event.Subscription, error) {

	logs, sub, err := _VaultTest.contract.WatchLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTestUpdatePosition)
				if err := _VaultTest.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
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

// ParseUpdatePosition is a log parse operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultTest *VaultTestFilterer) ParseUpdatePosition(log types.Log) (*VaultTestUpdatePosition, error) {
	event := new(VaultTestUpdatePosition)
	if err := _VaultTest.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
