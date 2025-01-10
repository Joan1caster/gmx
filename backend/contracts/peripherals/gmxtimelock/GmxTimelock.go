// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmxtimelock

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

// GmxTimelockMetaData contains all meta data concerning the GmxTimelock contract.
var GmxTimelockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_buffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_longBuffer\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mintReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxTokenSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"ClearAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"plugin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalAddPlugin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalPendingAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultPriceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceDecimals\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isStrictStable\",\"type\":\"bool\"}],\"name\":\"SignalPriceFeedSetTokenConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SignalRedeemUsdg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalSetGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalSetPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenDecimals\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minProfitBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxUsdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isStable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isShortable\",\"type\":\"bool\"}],\"name\":\"SignalVaultSetTokenConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"SignalWithdrawToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BUFFER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FUNDING_RATE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LEVERAGE_VALIDATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"addExcludedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_plugin\",\"type\":\"address\"}],\"name\":\"addPlugin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_action\",\"type\":\"bytes32\"}],\"name\":\"cancelAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"excludedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"longBuffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTokenSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrictStable\",\"type\":\"bool\"}],\"name\":\"priceFeedSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"processMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemUsdg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdditive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_adjustmentBps\",\"type\":\"uint256\"}],\"name\":\"setAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_buffer\",\"type\":\"uint256\"}],\"name\":\"setBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setContractHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setExternalAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_favorPrimaryPrice\",\"type\":\"bool\"}],\"name\":\"setFavorPrimaryPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_taxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_hasDynamicFees\",\"type\":\"bool\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fundingInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fundingRateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\"}],\"name\":\"setFundingRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateLiquidationMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateTransferMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsAmmEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLeverageEnabled\",\"type\":\"bool\"}],\"name\":\"setIsLeverageEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSecondaryPriceEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isSwapEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSwapEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGasPrice\",\"type\":\"uint256\"}],\"name\":\"setMaxGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMaxGlobalShortSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxLeverage\",\"type\":\"uint256\"}],\"name\":\"setMaxLeverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxStrictPriceDeviation\",\"type\":\"uint256\"}],\"name\":\"setMaxStrictPriceDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceSampleSpace\",\"type\":\"uint256\"}],\"name\":\"setPriceSampleSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_spreadThresholdBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setSpreadThresholdBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bufferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useV2Pricing\",\"type\":\"bool\"}],\"name\":\"setUseV2Pricing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"contractIVaultUtils\",\"name\":\"_vaultUtils\",\"type\":\"address\"}],\"name\":\"setVaultUtils\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_plugin\",\"type\":\"address\"}],\"name\":\"signalAddPlugin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrictStable\",\"type\":\"bool\"}],\"name\":\"signalPriceFeedSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalRedeemUsdg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"signalSetGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"signalSetPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isShortable\",\"type\":\"bool\"}],\"name\":\"signalVaultSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"signalWithdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfitBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isShortable\",\"type\":\"bool\"}],\"name\":\"vaultSetTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GmxTimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use GmxTimelockMetaData.ABI instead.
var GmxTimelockABI = GmxTimelockMetaData.ABI

// GmxTimelock is an auto generated Go binding around an Ethereum contract.
type GmxTimelock struct {
	GmxTimelockCaller     // Read-only binding to the contract
	GmxTimelockTransactor // Write-only binding to the contract
	GmxTimelockFilterer   // Log filterer for contract events
}

// GmxTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmxTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmxTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmxTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmxTimelockSession struct {
	Contract     *GmxTimelock      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmxTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmxTimelockCallerSession struct {
	Contract *GmxTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GmxTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmxTimelockTransactorSession struct {
	Contract     *GmxTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GmxTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmxTimelockRaw struct {
	Contract *GmxTimelock // Generic contract binding to access the raw methods on
}

// GmxTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmxTimelockCallerRaw struct {
	Contract *GmxTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// GmxTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmxTimelockTransactorRaw struct {
	Contract *GmxTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmxTimelock creates a new instance of GmxTimelock, bound to a specific deployed contract.
func NewGmxTimelock(address common.Address, backend bind.ContractBackend) (*GmxTimelock, error) {
	contract, err := bindGmxTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GmxTimelock{GmxTimelockCaller: GmxTimelockCaller{contract: contract}, GmxTimelockTransactor: GmxTimelockTransactor{contract: contract}, GmxTimelockFilterer: GmxTimelockFilterer{contract: contract}}, nil
}

// NewGmxTimelockCaller creates a new read-only instance of GmxTimelock, bound to a specific deployed contract.
func NewGmxTimelockCaller(address common.Address, caller bind.ContractCaller) (*GmxTimelockCaller, error) {
	contract, err := bindGmxTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmxTimelockCaller{contract: contract}, nil
}

// NewGmxTimelockTransactor creates a new write-only instance of GmxTimelock, bound to a specific deployed contract.
func NewGmxTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*GmxTimelockTransactor, error) {
	contract, err := bindGmxTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmxTimelockTransactor{contract: contract}, nil
}

// NewGmxTimelockFilterer creates a new log filterer instance of GmxTimelock, bound to a specific deployed contract.
func NewGmxTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*GmxTimelockFilterer, error) {
	contract, err := bindGmxTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmxTimelockFilterer{contract: contract}, nil
}

// bindGmxTimelock binds a generic wrapper to an already deployed contract.
func bindGmxTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GmxTimelockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxTimelock *GmxTimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxTimelock.Contract.GmxTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxTimelock *GmxTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxTimelock.Contract.GmxTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxTimelock *GmxTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxTimelock.Contract.GmxTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxTimelock *GmxTimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxTimelock *GmxTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxTimelock *GmxTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxTimelock.Contract.contract.Transact(opts, method, params...)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) MAXBUFFER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "MAX_BUFFER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) MAXBUFFER() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXBUFFER(&_GmxTimelock.CallOpts)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) MAXBUFFER() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXBUFFER(&_GmxTimelock.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) MAXFEEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "MAX_FEE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXFEEBASISPOINTS(&_GmxTimelock.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXFEEBASISPOINTS(&_GmxTimelock.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) MAXFUNDINGRATEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "MAX_FUNDING_RATE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXFUNDINGRATEFACTOR(&_GmxTimelock.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXFUNDINGRATEFACTOR(&_GmxTimelock.CallOpts)
}

// MAXLEVERAGEVALIDATION is a free data retrieval call binding the contract method 0x2ba3725a.
//
// Solidity: function MAX_LEVERAGE_VALIDATION() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) MAXLEVERAGEVALIDATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "MAX_LEVERAGE_VALIDATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLEVERAGEVALIDATION is a free data retrieval call binding the contract method 0x2ba3725a.
//
// Solidity: function MAX_LEVERAGE_VALIDATION() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) MAXLEVERAGEVALIDATION() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXLEVERAGEVALIDATION(&_GmxTimelock.CallOpts)
}

// MAXLEVERAGEVALIDATION is a free data retrieval call binding the contract method 0x2ba3725a.
//
// Solidity: function MAX_LEVERAGE_VALIDATION() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) MAXLEVERAGEVALIDATION() (*big.Int, error) {
	return _GmxTimelock.Contract.MAXLEVERAGEVALIDATION(&_GmxTimelock.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) PRICEPRECISION() (*big.Int, error) {
	return _GmxTimelock.Contract.PRICEPRECISION(&_GmxTimelock.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _GmxTimelock.Contract.PRICEPRECISION(&_GmxTimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxTimelock *GmxTimelockCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxTimelock *GmxTimelockSession) Admin() (common.Address, error) {
	return _GmxTimelock.Contract.Admin(&_GmxTimelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxTimelock *GmxTimelockCallerSession) Admin() (common.Address, error) {
	return _GmxTimelock.Contract.Admin(&_GmxTimelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) Buffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "buffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) Buffer() (*big.Int, error) {
	return _GmxTimelock.Contract.Buffer(&_GmxTimelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) Buffer() (*big.Int, error) {
	return _GmxTimelock.Contract.Buffer(&_GmxTimelock.CallOpts)
}

// ExcludedTokens is a free data retrieval call binding the contract method 0x5dae8841.
//
// Solidity: function excludedTokens(address ) view returns(bool)
func (_GmxTimelock *GmxTimelockCaller) ExcludedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "excludedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExcludedTokens is a free data retrieval call binding the contract method 0x5dae8841.
//
// Solidity: function excludedTokens(address ) view returns(bool)
func (_GmxTimelock *GmxTimelockSession) ExcludedTokens(arg0 common.Address) (bool, error) {
	return _GmxTimelock.Contract.ExcludedTokens(&_GmxTimelock.CallOpts, arg0)
}

// ExcludedTokens is a free data retrieval call binding the contract method 0x5dae8841.
//
// Solidity: function excludedTokens(address ) view returns(bool)
func (_GmxTimelock *GmxTimelockCallerSession) ExcludedTokens(arg0 common.Address) (bool, error) {
	return _GmxTimelock.Contract.ExcludedTokens(&_GmxTimelock.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxTimelock *GmxTimelockCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxTimelock *GmxTimelockSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxTimelock.Contract.IsHandler(&_GmxTimelock.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxTimelock *GmxTimelockCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxTimelock.Contract.IsHandler(&_GmxTimelock.CallOpts, arg0)
}

// LongBuffer is a free data retrieval call binding the contract method 0x5fa7955c.
//
// Solidity: function longBuffer() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) LongBuffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "longBuffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LongBuffer is a free data retrieval call binding the contract method 0x5fa7955c.
//
// Solidity: function longBuffer() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) LongBuffer() (*big.Int, error) {
	return _GmxTimelock.Contract.LongBuffer(&_GmxTimelock.CallOpts)
}

// LongBuffer is a free data retrieval call binding the contract method 0x5fa7955c.
//
// Solidity: function longBuffer() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) LongBuffer() (*big.Int, error) {
	return _GmxTimelock.Contract.LongBuffer(&_GmxTimelock.CallOpts)
}

// MaxTokenSupply is a free data retrieval call binding the contract method 0x50f7c204.
//
// Solidity: function maxTokenSupply() view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) MaxTokenSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "maxTokenSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTokenSupply is a free data retrieval call binding the contract method 0x50f7c204.
//
// Solidity: function maxTokenSupply() view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) MaxTokenSupply() (*big.Int, error) {
	return _GmxTimelock.Contract.MaxTokenSupply(&_GmxTimelock.CallOpts)
}

// MaxTokenSupply is a free data retrieval call binding the contract method 0x50f7c204.
//
// Solidity: function maxTokenSupply() view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) MaxTokenSupply() (*big.Int, error) {
	return _GmxTimelock.Contract.MaxTokenSupply(&_GmxTimelock.CallOpts)
}

// MintReceiver is a free data retrieval call binding the contract method 0xc7bb26a0.
//
// Solidity: function mintReceiver() view returns(address)
func (_GmxTimelock *GmxTimelockCaller) MintReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "mintReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintReceiver is a free data retrieval call binding the contract method 0xc7bb26a0.
//
// Solidity: function mintReceiver() view returns(address)
func (_GmxTimelock *GmxTimelockSession) MintReceiver() (common.Address, error) {
	return _GmxTimelock.Contract.MintReceiver(&_GmxTimelock.CallOpts)
}

// MintReceiver is a free data retrieval call binding the contract method 0xc7bb26a0.
//
// Solidity: function mintReceiver() view returns(address)
func (_GmxTimelock *GmxTimelockCallerSession) MintReceiver() (common.Address, error) {
	return _GmxTimelock.Contract.MintReceiver(&_GmxTimelock.CallOpts)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_GmxTimelock *GmxTimelockCaller) PendingActions(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "pendingActions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_GmxTimelock *GmxTimelockSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _GmxTimelock.Contract.PendingActions(&_GmxTimelock.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_GmxTimelock *GmxTimelockCallerSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _GmxTimelock.Contract.PendingActions(&_GmxTimelock.CallOpts, arg0)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_GmxTimelock *GmxTimelockCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_GmxTimelock *GmxTimelockSession) RewardManager() (common.Address, error) {
	return _GmxTimelock.Contract.RewardManager(&_GmxTimelock.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_GmxTimelock *GmxTimelockCallerSession) RewardManager() (common.Address, error) {
	return _GmxTimelock.Contract.RewardManager(&_GmxTimelock.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_GmxTimelock *GmxTimelockCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxTimelock.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_GmxTimelock *GmxTimelockSession) TokenManager() (common.Address, error) {
	return _GmxTimelock.Contract.TokenManager(&_GmxTimelock.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_GmxTimelock *GmxTimelockCallerSession) TokenManager() (common.Address, error) {
	return _GmxTimelock.Contract.TokenManager(&_GmxTimelock.CallOpts)
}

// AddExcludedToken is a paid mutator transaction binding the contract method 0xd7c2e92f.
//
// Solidity: function addExcludedToken(address _token) returns()
func (_GmxTimelock *GmxTimelockTransactor) AddExcludedToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "addExcludedToken", _token)
}

// AddExcludedToken is a paid mutator transaction binding the contract method 0xd7c2e92f.
//
// Solidity: function addExcludedToken(address _token) returns()
func (_GmxTimelock *GmxTimelockSession) AddExcludedToken(_token common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.AddExcludedToken(&_GmxTimelock.TransactOpts, _token)
}

// AddExcludedToken is a paid mutator transaction binding the contract method 0xd7c2e92f.
//
// Solidity: function addExcludedToken(address _token) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) AddExcludedToken(_token common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.AddExcludedToken(&_GmxTimelock.TransactOpts, _token)
}

// AddPlugin is a paid mutator transaction binding the contract method 0x69623ae2.
//
// Solidity: function addPlugin(address _router, address _plugin) returns()
func (_GmxTimelock *GmxTimelockTransactor) AddPlugin(opts *bind.TransactOpts, _router common.Address, _plugin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "addPlugin", _router, _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0x69623ae2.
//
// Solidity: function addPlugin(address _router, address _plugin) returns()
func (_GmxTimelock *GmxTimelockSession) AddPlugin(_router common.Address, _plugin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.AddPlugin(&_GmxTimelock.TransactOpts, _router, _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0x69623ae2.
//
// Solidity: function addPlugin(address _router, address _plugin) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) AddPlugin(_router common.Address, _plugin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.AddPlugin(&_GmxTimelock.TransactOpts, _router, _plugin)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "approve", _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.Approve(&_GmxTimelock.TransactOpts, _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.Approve(&_GmxTimelock.TransactOpts, _token, _spender, _amount)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_GmxTimelock *GmxTimelockTransactor) CancelAction(opts *bind.TransactOpts, _action [32]byte) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "cancelAction", _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_GmxTimelock *GmxTimelockSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _GmxTimelock.Contract.CancelAction(&_GmxTimelock.TransactOpts, _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _GmxTimelock.Contract.CancelAction(&_GmxTimelock.TransactOpts, _action)
}

// PriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x3335e38a.
//
// Solidity: function priceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_GmxTimelock *GmxTimelockTransactor) PriceFeedSetTokenConfig(opts *bind.TransactOpts, _vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "priceFeedSetTokenConfig", _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// PriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x3335e38a.
//
// Solidity: function priceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_GmxTimelock *GmxTimelockSession) PriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.PriceFeedSetTokenConfig(&_GmxTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// PriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x3335e38a.
//
// Solidity: function priceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) PriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.PriceFeedSetTokenConfig(&_GmxTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// ProcessMint is a paid mutator transaction binding the contract method 0xbc8a8ab9.
//
// Solidity: function processMint(address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) ProcessMint(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "processMint", _token, _receiver, _amount)
}

// ProcessMint is a paid mutator transaction binding the contract method 0xbc8a8ab9.
//
// Solidity: function processMint(address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) ProcessMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.ProcessMint(&_GmxTimelock.TransactOpts, _token, _receiver, _amount)
}

// ProcessMint is a paid mutator transaction binding the contract method 0xbc8a8ab9.
//
// Solidity: function processMint(address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) ProcessMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.ProcessMint(&_GmxTimelock.TransactOpts, _token, _receiver, _amount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) RedeemUsdg(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "redeemUsdg", _vault, _token, _amount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) RedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.RedeemUsdg(&_GmxTimelock.TransactOpts, _vault, _token, _amount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) RedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.RedeemUsdg(&_GmxTimelock.TransactOpts, _vault, _token, _amount)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address _token, address _account) returns()
func (_GmxTimelock *GmxTimelockTransactor) RemoveAdmin(opts *bind.TransactOpts, _token common.Address, _account common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "removeAdmin", _token, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address _token, address _account) returns()
func (_GmxTimelock *GmxTimelockSession) RemoveAdmin(_token common.Address, _account common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.RemoveAdmin(&_GmxTimelock.TransactOpts, _token, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address _token, address _account) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) RemoveAdmin(_token common.Address, _account common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.RemoveAdmin(&_GmxTimelock.TransactOpts, _token, _account)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0x296b07e5.
//
// Solidity: function setAdjustment(address _priceFeed, address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetAdjustment(opts *bind.TransactOpts, _priceFeed common.Address, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setAdjustment", _priceFeed, _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0x296b07e5.
//
// Solidity: function setAdjustment(address _priceFeed, address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_GmxTimelock *GmxTimelockSession) SetAdjustment(_priceFeed common.Address, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetAdjustment(&_GmxTimelock.TransactOpts, _priceFeed, _token, _isAdditive, _adjustmentBps)
}

// SetAdjustment is a paid mutator transaction binding the contract method 0x296b07e5.
//
// Solidity: function setAdjustment(address _priceFeed, address _token, bool _isAdditive, uint256 _adjustmentBps) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetAdjustment(_priceFeed common.Address, _token common.Address, _isAdditive bool, _adjustmentBps *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetAdjustment(&_GmxTimelock.TransactOpts, _priceFeed, _token, _isAdditive, _adjustmentBps)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_GmxTimelock *GmxTimelockSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetAdmin(&_GmxTimelock.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetAdmin(&_GmxTimelock.TransactOpts, _admin)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetBuffer(opts *bind.TransactOpts, _buffer *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setBuffer", _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_GmxTimelock *GmxTimelockSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetBuffer(&_GmxTimelock.TransactOpts, _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetBuffer(&_GmxTimelock.TransactOpts, _buffer)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetContractHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setContractHandler", _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_GmxTimelock *GmxTimelockSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetContractHandler(&_GmxTimelock.TransactOpts, _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetContractHandler(&_GmxTimelock.TransactOpts, _handler, _isActive)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetExternalAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setExternalAdmin", _target, _admin)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_GmxTimelock *GmxTimelockSession) SetExternalAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetExternalAdmin(&_GmxTimelock.TransactOpts, _target, _admin)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetExternalAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetExternalAdmin(&_GmxTimelock.TransactOpts, _target, _admin)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x7b9fb227.
//
// Solidity: function setFavorPrimaryPrice(address _priceFeed, bool _favorPrimaryPrice) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetFavorPrimaryPrice(opts *bind.TransactOpts, _priceFeed common.Address, _favorPrimaryPrice bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setFavorPrimaryPrice", _priceFeed, _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x7b9fb227.
//
// Solidity: function setFavorPrimaryPrice(address _priceFeed, bool _favorPrimaryPrice) returns()
func (_GmxTimelock *GmxTimelockSession) SetFavorPrimaryPrice(_priceFeed common.Address, _favorPrimaryPrice bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetFavorPrimaryPrice(&_GmxTimelock.TransactOpts, _priceFeed, _favorPrimaryPrice)
}

// SetFavorPrimaryPrice is a paid mutator transaction binding the contract method 0x7b9fb227.
//
// Solidity: function setFavorPrimaryPrice(address _priceFeed, bool _favorPrimaryPrice) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetFavorPrimaryPrice(_priceFeed common.Address, _favorPrimaryPrice bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetFavorPrimaryPrice(&_GmxTimelock.TransactOpts, _priceFeed, _favorPrimaryPrice)
}

// SetFees is a paid mutator transaction binding the contract method 0x6e5227d4.
//
// Solidity: function setFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetFees(opts *bind.TransactOpts, _vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setFees", _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x6e5227d4.
//
// Solidity: function setFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_GmxTimelock *GmxTimelockSession) SetFees(_vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetFees(&_GmxTimelock.TransactOpts, _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x6e5227d4.
//
// Solidity: function setFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetFees(_vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetFees(&_GmxTimelock.TransactOpts, _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x227f03eb.
//
// Solidity: function setFundingRate(address _vault, uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetFundingRate(opts *bind.TransactOpts, _vault common.Address, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setFundingRate", _vault, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x227f03eb.
//
// Solidity: function setFundingRate(address _vault, uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_GmxTimelock *GmxTimelockSession) SetFundingRate(_vault common.Address, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetFundingRate(&_GmxTimelock.TransactOpts, _vault, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x227f03eb.
//
// Solidity: function setFundingRate(address _vault, uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetFundingRate(_vault common.Address, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetFundingRate(&_GmxTimelock.TransactOpts, _vault, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setGov", _target, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_GmxTimelock *GmxTimelockSession) SetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetGov(&_GmxTimelock.TransactOpts, _target, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetGov(&_GmxTimelock.TransactOpts, _target, _gov)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0x21bd0592.
//
// Solidity: function setInPrivateLiquidationMode(address _vault, bool _inPrivateLiquidationMode) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _vault common.Address, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setInPrivateLiquidationMode", _vault, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0x21bd0592.
//
// Solidity: function setInPrivateLiquidationMode(address _vault, bool _inPrivateLiquidationMode) returns()
func (_GmxTimelock *GmxTimelockSession) SetInPrivateLiquidationMode(_vault common.Address, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetInPrivateLiquidationMode(&_GmxTimelock.TransactOpts, _vault, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0x21bd0592.
//
// Solidity: function setInPrivateLiquidationMode(address _vault, bool _inPrivateLiquidationMode) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetInPrivateLiquidationMode(_vault common.Address, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetInPrivateLiquidationMode(&_GmxTimelock.TransactOpts, _vault, _inPrivateLiquidationMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x86803c72.
//
// Solidity: function setInPrivateTransferMode(address _token, bool _inPrivateTransferMode) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _token common.Address, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setInPrivateTransferMode", _token, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x86803c72.
//
// Solidity: function setInPrivateTransferMode(address _token, bool _inPrivateTransferMode) returns()
func (_GmxTimelock *GmxTimelockSession) SetInPrivateTransferMode(_token common.Address, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetInPrivateTransferMode(&_GmxTimelock.TransactOpts, _token, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x86803c72.
//
// Solidity: function setInPrivateTransferMode(address _token, bool _inPrivateTransferMode) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetInPrivateTransferMode(_token common.Address, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetInPrivateTransferMode(&_GmxTimelock.TransactOpts, _token, _inPrivateTransferMode)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x5843752f.
//
// Solidity: function setIsAmmEnabled(address _priceFeed, bool _isEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetIsAmmEnabled(opts *bind.TransactOpts, _priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setIsAmmEnabled", _priceFeed, _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x5843752f.
//
// Solidity: function setIsAmmEnabled(address _priceFeed, bool _isEnabled) returns()
func (_GmxTimelock *GmxTimelockSession) SetIsAmmEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsAmmEnabled(&_GmxTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsAmmEnabled is a paid mutator transaction binding the contract method 0x5843752f.
//
// Solidity: function setIsAmmEnabled(address _priceFeed, bool _isEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetIsAmmEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsAmmEnabled(&_GmxTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setIsLeverageEnabled", _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_GmxTimelock *GmxTimelockSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsLeverageEnabled(&_GmxTimelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsLeverageEnabled(&_GmxTimelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0x2965c8c7.
//
// Solidity: function setIsSecondaryPriceEnabled(address _priceFeed, bool _isEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetIsSecondaryPriceEnabled(opts *bind.TransactOpts, _priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setIsSecondaryPriceEnabled", _priceFeed, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0x2965c8c7.
//
// Solidity: function setIsSecondaryPriceEnabled(address _priceFeed, bool _isEnabled) returns()
func (_GmxTimelock *GmxTimelockSession) SetIsSecondaryPriceEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsSecondaryPriceEnabled(&_GmxTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsSecondaryPriceEnabled is a paid mutator transaction binding the contract method 0x2965c8c7.
//
// Solidity: function setIsSecondaryPriceEnabled(address _priceFeed, bool _isEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetIsSecondaryPriceEnabled(_priceFeed common.Address, _isEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsSecondaryPriceEnabled(&_GmxTimelock.TransactOpts, _priceFeed, _isEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x117cf204.
//
// Solidity: function setIsSwapEnabled(address _vault, bool _isSwapEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _vault common.Address, _isSwapEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setIsSwapEnabled", _vault, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x117cf204.
//
// Solidity: function setIsSwapEnabled(address _vault, bool _isSwapEnabled) returns()
func (_GmxTimelock *GmxTimelockSession) SetIsSwapEnabled(_vault common.Address, _isSwapEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsSwapEnabled(&_GmxTimelock.TransactOpts, _vault, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x117cf204.
//
// Solidity: function setIsSwapEnabled(address _vault, bool _isSwapEnabled) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetIsSwapEnabled(_vault common.Address, _isSwapEnabled bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetIsSwapEnabled(&_GmxTimelock.TransactOpts, _vault, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x47de43e2.
//
// Solidity: function setLiquidator(address _vault, address _liquidator, bool _isActive) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetLiquidator(opts *bind.TransactOpts, _vault common.Address, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setLiquidator", _vault, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x47de43e2.
//
// Solidity: function setLiquidator(address _vault, address _liquidator, bool _isActive) returns()
func (_GmxTimelock *GmxTimelockSession) SetLiquidator(_vault common.Address, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetLiquidator(&_GmxTimelock.TransactOpts, _vault, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x47de43e2.
//
// Solidity: function setLiquidator(address _vault, address _liquidator, bool _isActive) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetLiquidator(_vault common.Address, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetLiquidator(&_GmxTimelock.TransactOpts, _vault, _liquidator, _isActive)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0x8bf22c46.
//
// Solidity: function setMaxGasPrice(address _vault, uint256 _maxGasPrice) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _vault common.Address, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setMaxGasPrice", _vault, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0x8bf22c46.
//
// Solidity: function setMaxGasPrice(address _vault, uint256 _maxGasPrice) returns()
func (_GmxTimelock *GmxTimelockSession) SetMaxGasPrice(_vault common.Address, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxGasPrice(&_GmxTimelock.TransactOpts, _vault, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0x8bf22c46.
//
// Solidity: function setMaxGasPrice(address _vault, uint256 _maxGasPrice) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetMaxGasPrice(_vault common.Address, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxGasPrice(&_GmxTimelock.TransactOpts, _vault, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xdf3a66d9.
//
// Solidity: function setMaxGlobalShortSize(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setMaxGlobalShortSize", _vault, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xdf3a66d9.
//
// Solidity: function setMaxGlobalShortSize(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) SetMaxGlobalShortSize(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxGlobalShortSize(&_GmxTimelock.TransactOpts, _vault, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xdf3a66d9.
//
// Solidity: function setMaxGlobalShortSize(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetMaxGlobalShortSize(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxGlobalShortSize(&_GmxTimelock.TransactOpts, _vault, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0x7b6f775a.
//
// Solidity: function setMaxLeverage(address _vault, uint256 _maxLeverage) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetMaxLeverage(opts *bind.TransactOpts, _vault common.Address, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setMaxLeverage", _vault, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0x7b6f775a.
//
// Solidity: function setMaxLeverage(address _vault, uint256 _maxLeverage) returns()
func (_GmxTimelock *GmxTimelockSession) SetMaxLeverage(_vault common.Address, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxLeverage(&_GmxTimelock.TransactOpts, _vault, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0x7b6f775a.
//
// Solidity: function setMaxLeverage(address _vault, uint256 _maxLeverage) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetMaxLeverage(_vault common.Address, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxLeverage(&_GmxTimelock.TransactOpts, _vault, _maxLeverage)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0xa0a316a2.
//
// Solidity: function setMaxStrictPriceDeviation(address _priceFeed, uint256 _maxStrictPriceDeviation) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetMaxStrictPriceDeviation(opts *bind.TransactOpts, _priceFeed common.Address, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setMaxStrictPriceDeviation", _priceFeed, _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0xa0a316a2.
//
// Solidity: function setMaxStrictPriceDeviation(address _priceFeed, uint256 _maxStrictPriceDeviation) returns()
func (_GmxTimelock *GmxTimelockSession) SetMaxStrictPriceDeviation(_priceFeed common.Address, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxStrictPriceDeviation(&_GmxTimelock.TransactOpts, _priceFeed, _maxStrictPriceDeviation)
}

// SetMaxStrictPriceDeviation is a paid mutator transaction binding the contract method 0xa0a316a2.
//
// Solidity: function setMaxStrictPriceDeviation(address _priceFeed, uint256 _maxStrictPriceDeviation) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetMaxStrictPriceDeviation(_priceFeed common.Address, _maxStrictPriceDeviation *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetMaxStrictPriceDeviation(&_GmxTimelock.TransactOpts, _priceFeed, _maxStrictPriceDeviation)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _vault, address _priceFeed) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetPriceFeed(opts *bind.TransactOpts, _vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setPriceFeed", _vault, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _vault, address _priceFeed) returns()
func (_GmxTimelock *GmxTimelockSession) SetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetPriceFeed(&_GmxTimelock.TransactOpts, _vault, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _vault, address _priceFeed) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetPriceFeed(&_GmxTimelock.TransactOpts, _vault, _priceFeed)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2877f4c3.
//
// Solidity: function setPriceSampleSpace(address _priceFeed, uint256 _priceSampleSpace) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetPriceSampleSpace(opts *bind.TransactOpts, _priceFeed common.Address, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setPriceSampleSpace", _priceFeed, _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2877f4c3.
//
// Solidity: function setPriceSampleSpace(address _priceFeed, uint256 _priceSampleSpace) returns()
func (_GmxTimelock *GmxTimelockSession) SetPriceSampleSpace(_priceFeed common.Address, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetPriceSampleSpace(&_GmxTimelock.TransactOpts, _priceFeed, _priceSampleSpace)
}

// SetPriceSampleSpace is a paid mutator transaction binding the contract method 0x2877f4c3.
//
// Solidity: function setPriceSampleSpace(address _priceFeed, uint256 _priceSampleSpace) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetPriceSampleSpace(_priceFeed common.Address, _priceSampleSpace *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetPriceSampleSpace(&_GmxTimelock.TransactOpts, _priceFeed, _priceSampleSpace)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0xe7b0a3a1.
//
// Solidity: function setSpreadBasisPoints(address _priceFeed, address _token, uint256 _spreadBasisPoints) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetSpreadBasisPoints(opts *bind.TransactOpts, _priceFeed common.Address, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setSpreadBasisPoints", _priceFeed, _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0xe7b0a3a1.
//
// Solidity: function setSpreadBasisPoints(address _priceFeed, address _token, uint256 _spreadBasisPoints) returns()
func (_GmxTimelock *GmxTimelockSession) SetSpreadBasisPoints(_priceFeed common.Address, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetSpreadBasisPoints(&_GmxTimelock.TransactOpts, _priceFeed, _token, _spreadBasisPoints)
}

// SetSpreadBasisPoints is a paid mutator transaction binding the contract method 0xe7b0a3a1.
//
// Solidity: function setSpreadBasisPoints(address _priceFeed, address _token, uint256 _spreadBasisPoints) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetSpreadBasisPoints(_priceFeed common.Address, _token common.Address, _spreadBasisPoints *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetSpreadBasisPoints(&_GmxTimelock.TransactOpts, _priceFeed, _token, _spreadBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0x14aea307.
//
// Solidity: function setSpreadThresholdBasisPoints(address _priceFeed, uint256 _spreadThresholdBasisPoints) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetSpreadThresholdBasisPoints(opts *bind.TransactOpts, _priceFeed common.Address, _spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setSpreadThresholdBasisPoints", _priceFeed, _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0x14aea307.
//
// Solidity: function setSpreadThresholdBasisPoints(address _priceFeed, uint256 _spreadThresholdBasisPoints) returns()
func (_GmxTimelock *GmxTimelockSession) SetSpreadThresholdBasisPoints(_priceFeed common.Address, _spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetSpreadThresholdBasisPoints(&_GmxTimelock.TransactOpts, _priceFeed, _spreadThresholdBasisPoints)
}

// SetSpreadThresholdBasisPoints is a paid mutator transaction binding the contract method 0x14aea307.
//
// Solidity: function setSpreadThresholdBasisPoints(address _priceFeed, uint256 _spreadThresholdBasisPoints) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetSpreadThresholdBasisPoints(_priceFeed common.Address, _spreadThresholdBasisPoints *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetSpreadThresholdBasisPoints(&_GmxTimelock.TransactOpts, _priceFeed, _spreadThresholdBasisPoints)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x0e0dc426.
//
// Solidity: function setTokenConfig(address _vault, address _token, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, uint256 _bufferAmount, uint256 _usdgAmount) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetTokenConfig(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _bufferAmount *big.Int, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setTokenConfig", _vault, _token, _tokenWeight, _minProfitBps, _maxUsdgAmount, _bufferAmount, _usdgAmount)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x0e0dc426.
//
// Solidity: function setTokenConfig(address _vault, address _token, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, uint256 _bufferAmount, uint256 _usdgAmount) returns()
func (_GmxTimelock *GmxTimelockSession) SetTokenConfig(_vault common.Address, _token common.Address, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _bufferAmount *big.Int, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetTokenConfig(&_GmxTimelock.TransactOpts, _vault, _token, _tokenWeight, _minProfitBps, _maxUsdgAmount, _bufferAmount, _usdgAmount)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x0e0dc426.
//
// Solidity: function setTokenConfig(address _vault, address _token, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, uint256 _bufferAmount, uint256 _usdgAmount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetTokenConfig(_vault common.Address, _token common.Address, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _bufferAmount *big.Int, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetTokenConfig(&_GmxTimelock.TransactOpts, _vault, _token, _tokenWeight, _minProfitBps, _maxUsdgAmount, _bufferAmount, _usdgAmount)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0x454ffa46.
//
// Solidity: function setUseV2Pricing(address _priceFeed, bool _useV2Pricing) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetUseV2Pricing(opts *bind.TransactOpts, _priceFeed common.Address, _useV2Pricing bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setUseV2Pricing", _priceFeed, _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0x454ffa46.
//
// Solidity: function setUseV2Pricing(address _priceFeed, bool _useV2Pricing) returns()
func (_GmxTimelock *GmxTimelockSession) SetUseV2Pricing(_priceFeed common.Address, _useV2Pricing bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetUseV2Pricing(&_GmxTimelock.TransactOpts, _priceFeed, _useV2Pricing)
}

// SetUseV2Pricing is a paid mutator transaction binding the contract method 0x454ffa46.
//
// Solidity: function setUseV2Pricing(address _priceFeed, bool _useV2Pricing) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetUseV2Pricing(_priceFeed common.Address, _useV2Pricing bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetUseV2Pricing(&_GmxTimelock.TransactOpts, _priceFeed, _useV2Pricing)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0xbc476dfd.
//
// Solidity: function setVaultUtils(address _vault, address _vaultUtils) returns()
func (_GmxTimelock *GmxTimelockTransactor) SetVaultUtils(opts *bind.TransactOpts, _vault common.Address, _vaultUtils common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "setVaultUtils", _vault, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0xbc476dfd.
//
// Solidity: function setVaultUtils(address _vault, address _vaultUtils) returns()
func (_GmxTimelock *GmxTimelockSession) SetVaultUtils(_vault common.Address, _vaultUtils common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetVaultUtils(&_GmxTimelock.TransactOpts, _vault, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0xbc476dfd.
//
// Solidity: function setVaultUtils(address _vault, address _vaultUtils) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SetVaultUtils(_vault common.Address, _vaultUtils common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SetVaultUtils(&_GmxTimelock.TransactOpts, _vault, _vaultUtils)
}

// SignalAddPlugin is a paid mutator transaction binding the contract method 0x53b6bfdd.
//
// Solidity: function signalAddPlugin(address _router, address _plugin) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalAddPlugin(opts *bind.TransactOpts, _router common.Address, _plugin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalAddPlugin", _router, _plugin)
}

// SignalAddPlugin is a paid mutator transaction binding the contract method 0x53b6bfdd.
//
// Solidity: function signalAddPlugin(address _router, address _plugin) returns()
func (_GmxTimelock *GmxTimelockSession) SignalAddPlugin(_router common.Address, _plugin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalAddPlugin(&_GmxTimelock.TransactOpts, _router, _plugin)
}

// SignalAddPlugin is a paid mutator transaction binding the contract method 0x53b6bfdd.
//
// Solidity: function signalAddPlugin(address _router, address _plugin) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalAddPlugin(_router common.Address, _plugin common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalAddPlugin(&_GmxTimelock.TransactOpts, _router, _plugin)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalApprove", _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalApprove(&_GmxTimelock.TransactOpts, _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalApprove(&_GmxTimelock.TransactOpts, _token, _spender, _amount)
}

// SignalMint is a paid mutator transaction binding the contract method 0x09cc9a08.
//
// Solidity: function signalMint(address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalMint(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalMint", _token, _receiver, _amount)
}

// SignalMint is a paid mutator transaction binding the contract method 0x09cc9a08.
//
// Solidity: function signalMint(address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) SignalMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalMint(&_GmxTimelock.TransactOpts, _token, _receiver, _amount)
}

// SignalMint is a paid mutator transaction binding the contract method 0x09cc9a08.
//
// Solidity: function signalMint(address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalMint(&_GmxTimelock.TransactOpts, _token, _receiver, _amount)
}

// SignalPriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x384cae73.
//
// Solidity: function signalPriceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalPriceFeedSetTokenConfig(opts *bind.TransactOpts, _vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalPriceFeedSetTokenConfig", _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SignalPriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x384cae73.
//
// Solidity: function signalPriceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_GmxTimelock *GmxTimelockSession) SignalPriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalPriceFeedSetTokenConfig(&_GmxTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SignalPriceFeedSetTokenConfig is a paid mutator transaction binding the contract method 0x384cae73.
//
// Solidity: function signalPriceFeedSetTokenConfig(address _vaultPriceFeed, address _token, address _priceFeed, uint256 _priceDecimals, bool _isStrictStable) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalPriceFeedSetTokenConfig(_vaultPriceFeed common.Address, _token common.Address, _priceFeed common.Address, _priceDecimals *big.Int, _isStrictStable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalPriceFeedSetTokenConfig(&_GmxTimelock.TransactOpts, _vaultPriceFeed, _token, _priceFeed, _priceDecimals, _isStrictStable)
}

// SignalRedeemUsdg is a paid mutator transaction binding the contract method 0x0191c237.
//
// Solidity: function signalRedeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalRedeemUsdg(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalRedeemUsdg", _vault, _token, _amount)
}

// SignalRedeemUsdg is a paid mutator transaction binding the contract method 0x0191c237.
//
// Solidity: function signalRedeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) SignalRedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalRedeemUsdg(&_GmxTimelock.TransactOpts, _vault, _token, _amount)
}

// SignalRedeemUsdg is a paid mutator transaction binding the contract method 0x0191c237.
//
// Solidity: function signalRedeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalRedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalRedeemUsdg(&_GmxTimelock.TransactOpts, _vault, _token, _amount)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalSetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalSetGov", _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_GmxTimelock *GmxTimelockSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalSetGov(&_GmxTimelock.TransactOpts, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalSetGov(&_GmxTimelock.TransactOpts, _target, _gov)
}

// SignalSetPriceFeed is a paid mutator transaction binding the contract method 0x80894d62.
//
// Solidity: function signalSetPriceFeed(address _vault, address _priceFeed) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalSetPriceFeed(opts *bind.TransactOpts, _vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalSetPriceFeed", _vault, _priceFeed)
}

// SignalSetPriceFeed is a paid mutator transaction binding the contract method 0x80894d62.
//
// Solidity: function signalSetPriceFeed(address _vault, address _priceFeed) returns()
func (_GmxTimelock *GmxTimelockSession) SignalSetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalSetPriceFeed(&_GmxTimelock.TransactOpts, _vault, _priceFeed)
}

// SignalSetPriceFeed is a paid mutator transaction binding the contract method 0x80894d62.
//
// Solidity: function signalSetPriceFeed(address _vault, address _priceFeed) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalSetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalSetPriceFeed(&_GmxTimelock.TransactOpts, _vault, _priceFeed)
}

// SignalVaultSetTokenConfig is a paid mutator transaction binding the contract method 0xdb1c8441.
//
// Solidity: function signalVaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalVaultSetTokenConfig(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalVaultSetTokenConfig", _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SignalVaultSetTokenConfig is a paid mutator transaction binding the contract method 0xdb1c8441.
//
// Solidity: function signalVaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_GmxTimelock *GmxTimelockSession) SignalVaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalVaultSetTokenConfig(&_GmxTimelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SignalVaultSetTokenConfig is a paid mutator transaction binding the contract method 0xdb1c8441.
//
// Solidity: function signalVaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalVaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalVaultSetTokenConfig(&_GmxTimelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) SignalWithdrawToken(opts *bind.TransactOpts, _target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "signalWithdrawToken", _target, _token, _receiver, _amount)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) SignalWithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalWithdrawToken(&_GmxTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) SignalWithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.SignalWithdrawToken(&_GmxTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) TransferIn(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "transferIn", _sender, _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) TransferIn(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.TransferIn(&_GmxTimelock.TransactOpts, _sender, _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) TransferIn(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.TransferIn(&_GmxTimelock.TransactOpts, _sender, _token, _amount)
}

// VaultSetTokenConfig is a paid mutator transaction binding the contract method 0xe3cbeb0f.
//
// Solidity: function vaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_GmxTimelock *GmxTimelockTransactor) VaultSetTokenConfig(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "vaultSetTokenConfig", _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// VaultSetTokenConfig is a paid mutator transaction binding the contract method 0xe3cbeb0f.
//
// Solidity: function vaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_GmxTimelock *GmxTimelockSession) VaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.VaultSetTokenConfig(&_GmxTimelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// VaultSetTokenConfig is a paid mutator transaction binding the contract method 0xe3cbeb0f.
//
// Solidity: function vaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) VaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _GmxTimelock.Contract.VaultSetTokenConfig(&_GmxTimelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x0e9587f3.
//
// Solidity: function withdrawFees(address _vault, address _token, address _receiver) returns()
func (_GmxTimelock *GmxTimelockTransactor) WithdrawFees(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "withdrawFees", _vault, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x0e9587f3.
//
// Solidity: function withdrawFees(address _vault, address _token, address _receiver) returns()
func (_GmxTimelock *GmxTimelockSession) WithdrawFees(_vault common.Address, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.WithdrawFees(&_GmxTimelock.TransactOpts, _vault, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x0e9587f3.
//
// Solidity: function withdrawFees(address _vault, address _token, address _receiver) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) WithdrawFees(_vault common.Address, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxTimelock.Contract.WithdrawFees(&_GmxTimelock.TransactOpts, _vault, _token, _receiver)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactor) WithdrawToken(opts *bind.TransactOpts, _target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.contract.Transact(opts, "withdrawToken", _target, _token, _receiver, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockSession) WithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.WithdrawToken(&_GmxTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_GmxTimelock *GmxTimelockTransactorSession) WithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxTimelock.Contract.WithdrawToken(&_GmxTimelock.TransactOpts, _target, _token, _receiver, _amount)
}

// GmxTimelockClearActionIterator is returned from FilterClearAction and is used to iterate over the raw logs and unpacked data for ClearAction events raised by the GmxTimelock contract.
type GmxTimelockClearActionIterator struct {
	Event *GmxTimelockClearAction // Event containing the contract specifics and raw log

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
func (it *GmxTimelockClearActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockClearAction)
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
		it.Event = new(GmxTimelockClearAction)
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
func (it *GmxTimelockClearActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockClearActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockClearAction represents a ClearAction event raised by the GmxTimelock contract.
type GmxTimelockClearAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClearAction is a free log retrieval operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterClearAction(opts *bind.FilterOpts) (*GmxTimelockClearActionIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockClearActionIterator{contract: _GmxTimelock.contract, event: "ClearAction", logs: logs, sub: sub}, nil
}

// WatchClearAction is a free log subscription operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchClearAction(opts *bind.WatchOpts, sink chan<- *GmxTimelockClearAction) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockClearAction)
				if err := _GmxTimelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
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

// ParseClearAction is a log parse operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseClearAction(log types.Log) (*GmxTimelockClearAction, error) {
	event := new(GmxTimelockClearAction)
	if err := _GmxTimelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalAddPluginIterator is returned from FilterSignalAddPlugin and is used to iterate over the raw logs and unpacked data for SignalAddPlugin events raised by the GmxTimelock contract.
type GmxTimelockSignalAddPluginIterator struct {
	Event *GmxTimelockSignalAddPlugin // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalAddPluginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalAddPlugin)
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
		it.Event = new(GmxTimelockSignalAddPlugin)
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
func (it *GmxTimelockSignalAddPluginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalAddPluginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalAddPlugin represents a SignalAddPlugin event raised by the GmxTimelock contract.
type GmxTimelockSignalAddPlugin struct {
	Router common.Address
	Plugin common.Address
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalAddPlugin is a free log retrieval operation binding the contract event 0xc0ad20f21e0e6b9c02a7ecfb229e39b214cb1914a4d6f202e7f9ec7ffb9445f6.
//
// Solidity: event SignalAddPlugin(address router, address plugin, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalAddPlugin(opts *bind.FilterOpts) (*GmxTimelockSignalAddPluginIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalAddPlugin")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalAddPluginIterator{contract: _GmxTimelock.contract, event: "SignalAddPlugin", logs: logs, sub: sub}, nil
}

// WatchSignalAddPlugin is a free log subscription operation binding the contract event 0xc0ad20f21e0e6b9c02a7ecfb229e39b214cb1914a4d6f202e7f9ec7ffb9445f6.
//
// Solidity: event SignalAddPlugin(address router, address plugin, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalAddPlugin(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalAddPlugin) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalAddPlugin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalAddPlugin)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalAddPlugin", log); err != nil {
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

// ParseSignalAddPlugin is a log parse operation binding the contract event 0xc0ad20f21e0e6b9c02a7ecfb229e39b214cb1914a4d6f202e7f9ec7ffb9445f6.
//
// Solidity: event SignalAddPlugin(address router, address plugin, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalAddPlugin(log types.Log) (*GmxTimelockSignalAddPlugin, error) {
	event := new(GmxTimelockSignalAddPlugin)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalAddPlugin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalApproveIterator is returned from FilterSignalApprove and is used to iterate over the raw logs and unpacked data for SignalApprove events raised by the GmxTimelock contract.
type GmxTimelockSignalApproveIterator struct {
	Event *GmxTimelockSignalApprove // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalApprove)
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
		it.Event = new(GmxTimelockSignalApprove)
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
func (it *GmxTimelockSignalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalApprove represents a SignalApprove event raised by the GmxTimelock contract.
type GmxTimelockSignalApprove struct {
	Token   common.Address
	Spender common.Address
	Amount  *big.Int
	Action  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignalApprove is a free log retrieval operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalApprove(opts *bind.FilterOpts) (*GmxTimelockSignalApproveIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalApproveIterator{contract: _GmxTimelock.contract, event: "SignalApprove", logs: logs, sub: sub}, nil
}

// WatchSignalApprove is a free log subscription operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalApprove(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalApprove) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalApprove)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalApprove", log); err != nil {
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

// ParseSignalApprove is a log parse operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalApprove(log types.Log) (*GmxTimelockSignalApprove, error) {
	event := new(GmxTimelockSignalApprove)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalMintIterator is returned from FilterSignalMint and is used to iterate over the raw logs and unpacked data for SignalMint events raised by the GmxTimelock contract.
type GmxTimelockSignalMintIterator struct {
	Event *GmxTimelockSignalMint // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalMint)
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
		it.Event = new(GmxTimelockSignalMint)
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
func (it *GmxTimelockSignalMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalMint represents a SignalMint event raised by the GmxTimelock contract.
type GmxTimelockSignalMint struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Action   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalMint is a free log retrieval operation binding the contract event 0x23d37bec99db82564427c9bbfe48ad7434bccf413a40fd357fb838c90a0d6828.
//
// Solidity: event SignalMint(address token, address receiver, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalMint(opts *bind.FilterOpts) (*GmxTimelockSignalMintIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalMint")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalMintIterator{contract: _GmxTimelock.contract, event: "SignalMint", logs: logs, sub: sub}, nil
}

// WatchSignalMint is a free log subscription operation binding the contract event 0x23d37bec99db82564427c9bbfe48ad7434bccf413a40fd357fb838c90a0d6828.
//
// Solidity: event SignalMint(address token, address receiver, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalMint(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalMint) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalMint)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalMint", log); err != nil {
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

// ParseSignalMint is a log parse operation binding the contract event 0x23d37bec99db82564427c9bbfe48ad7434bccf413a40fd357fb838c90a0d6828.
//
// Solidity: event SignalMint(address token, address receiver, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalMint(log types.Log) (*GmxTimelockSignalMint, error) {
	event := new(GmxTimelockSignalMint)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalPendingActionIterator is returned from FilterSignalPendingAction and is used to iterate over the raw logs and unpacked data for SignalPendingAction events raised by the GmxTimelock contract.
type GmxTimelockSignalPendingActionIterator struct {
	Event *GmxTimelockSignalPendingAction // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalPendingActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalPendingAction)
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
		it.Event = new(GmxTimelockSignalPendingAction)
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
func (it *GmxTimelockSignalPendingActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalPendingActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalPendingAction represents a SignalPendingAction event raised by the GmxTimelock contract.
type GmxTimelockSignalPendingAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalPendingAction is a free log retrieval operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalPendingAction(opts *bind.FilterOpts) (*GmxTimelockSignalPendingActionIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalPendingActionIterator{contract: _GmxTimelock.contract, event: "SignalPendingAction", logs: logs, sub: sub}, nil
}

// WatchSignalPendingAction is a free log subscription operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalPendingAction(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalPendingAction) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalPendingAction)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
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

// ParseSignalPendingAction is a log parse operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalPendingAction(log types.Log) (*GmxTimelockSignalPendingAction, error) {
	event := new(GmxTimelockSignalPendingAction)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalPriceFeedSetTokenConfigIterator is returned from FilterSignalPriceFeedSetTokenConfig and is used to iterate over the raw logs and unpacked data for SignalPriceFeedSetTokenConfig events raised by the GmxTimelock contract.
type GmxTimelockSignalPriceFeedSetTokenConfigIterator struct {
	Event *GmxTimelockSignalPriceFeedSetTokenConfig // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalPriceFeedSetTokenConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalPriceFeedSetTokenConfig)
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
		it.Event = new(GmxTimelockSignalPriceFeedSetTokenConfig)
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
func (it *GmxTimelockSignalPriceFeedSetTokenConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalPriceFeedSetTokenConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalPriceFeedSetTokenConfig represents a SignalPriceFeedSetTokenConfig event raised by the GmxTimelock contract.
type GmxTimelockSignalPriceFeedSetTokenConfig struct {
	VaultPriceFeed common.Address
	Token          common.Address
	PriceFeed      common.Address
	PriceDecimals  *big.Int
	IsStrictStable bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSignalPriceFeedSetTokenConfig is a free log retrieval operation binding the contract event 0x1b2ddf357ae016d8c127dcd3a73c34744fdeaeeb4b7ef1e04490cebf7f4816fe.
//
// Solidity: event SignalPriceFeedSetTokenConfig(address vaultPriceFeed, address token, address priceFeed, uint256 priceDecimals, bool isStrictStable)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalPriceFeedSetTokenConfig(opts *bind.FilterOpts) (*GmxTimelockSignalPriceFeedSetTokenConfigIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalPriceFeedSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalPriceFeedSetTokenConfigIterator{contract: _GmxTimelock.contract, event: "SignalPriceFeedSetTokenConfig", logs: logs, sub: sub}, nil
}

// WatchSignalPriceFeedSetTokenConfig is a free log subscription operation binding the contract event 0x1b2ddf357ae016d8c127dcd3a73c34744fdeaeeb4b7ef1e04490cebf7f4816fe.
//
// Solidity: event SignalPriceFeedSetTokenConfig(address vaultPriceFeed, address token, address priceFeed, uint256 priceDecimals, bool isStrictStable)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalPriceFeedSetTokenConfig(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalPriceFeedSetTokenConfig) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalPriceFeedSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalPriceFeedSetTokenConfig)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalPriceFeedSetTokenConfig", log); err != nil {
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

// ParseSignalPriceFeedSetTokenConfig is a log parse operation binding the contract event 0x1b2ddf357ae016d8c127dcd3a73c34744fdeaeeb4b7ef1e04490cebf7f4816fe.
//
// Solidity: event SignalPriceFeedSetTokenConfig(address vaultPriceFeed, address token, address priceFeed, uint256 priceDecimals, bool isStrictStable)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalPriceFeedSetTokenConfig(log types.Log) (*GmxTimelockSignalPriceFeedSetTokenConfig, error) {
	event := new(GmxTimelockSignalPriceFeedSetTokenConfig)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalPriceFeedSetTokenConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalRedeemUsdgIterator is returned from FilterSignalRedeemUsdg and is used to iterate over the raw logs and unpacked data for SignalRedeemUsdg events raised by the GmxTimelock contract.
type GmxTimelockSignalRedeemUsdgIterator struct {
	Event *GmxTimelockSignalRedeemUsdg // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalRedeemUsdgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalRedeemUsdg)
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
		it.Event = new(GmxTimelockSignalRedeemUsdg)
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
func (it *GmxTimelockSignalRedeemUsdgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalRedeemUsdgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalRedeemUsdg represents a SignalRedeemUsdg event raised by the GmxTimelock contract.
type GmxTimelockSignalRedeemUsdg struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalRedeemUsdg is a free log retrieval operation binding the contract event 0xe6bd553b6ef21f3a22ebc877b3aaedc30fe15826b8156d4e8c8b373ebf11d78b.
//
// Solidity: event SignalRedeemUsdg(address vault, address token, uint256 amount)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalRedeemUsdg(opts *bind.FilterOpts) (*GmxTimelockSignalRedeemUsdgIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalRedeemUsdg")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalRedeemUsdgIterator{contract: _GmxTimelock.contract, event: "SignalRedeemUsdg", logs: logs, sub: sub}, nil
}

// WatchSignalRedeemUsdg is a free log subscription operation binding the contract event 0xe6bd553b6ef21f3a22ebc877b3aaedc30fe15826b8156d4e8c8b373ebf11d78b.
//
// Solidity: event SignalRedeemUsdg(address vault, address token, uint256 amount)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalRedeemUsdg(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalRedeemUsdg) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalRedeemUsdg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalRedeemUsdg)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalRedeemUsdg", log); err != nil {
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

// ParseSignalRedeemUsdg is a log parse operation binding the contract event 0xe6bd553b6ef21f3a22ebc877b3aaedc30fe15826b8156d4e8c8b373ebf11d78b.
//
// Solidity: event SignalRedeemUsdg(address vault, address token, uint256 amount)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalRedeemUsdg(log types.Log) (*GmxTimelockSignalRedeemUsdg, error) {
	event := new(GmxTimelockSignalRedeemUsdg)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalRedeemUsdg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalSetGovIterator is returned from FilterSignalSetGov and is used to iterate over the raw logs and unpacked data for SignalSetGov events raised by the GmxTimelock contract.
type GmxTimelockSignalSetGovIterator struct {
	Event *GmxTimelockSignalSetGov // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalSetGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalSetGov)
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
		it.Event = new(GmxTimelockSignalSetGov)
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
func (it *GmxTimelockSignalSetGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalSetGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalSetGov represents a SignalSetGov event raised by the GmxTimelock contract.
type GmxTimelockSignalSetGov struct {
	Target common.Address
	Gov    common.Address
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalSetGov is a free log retrieval operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalSetGov(opts *bind.FilterOpts) (*GmxTimelockSignalSetGovIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalSetGovIterator{contract: _GmxTimelock.contract, event: "SignalSetGov", logs: logs, sub: sub}, nil
}

// WatchSignalSetGov is a free log subscription operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalSetGov(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalSetGov) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalSetGov)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
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

// ParseSignalSetGov is a log parse operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalSetGov(log types.Log) (*GmxTimelockSignalSetGov, error) {
	event := new(GmxTimelockSignalSetGov)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalSetPriceFeedIterator is returned from FilterSignalSetPriceFeed and is used to iterate over the raw logs and unpacked data for SignalSetPriceFeed events raised by the GmxTimelock contract.
type GmxTimelockSignalSetPriceFeedIterator struct {
	Event *GmxTimelockSignalSetPriceFeed // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalSetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalSetPriceFeed)
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
		it.Event = new(GmxTimelockSignalSetPriceFeed)
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
func (it *GmxTimelockSignalSetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalSetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalSetPriceFeed represents a SignalSetPriceFeed event raised by the GmxTimelock contract.
type GmxTimelockSignalSetPriceFeed struct {
	Vault     common.Address
	PriceFeed common.Address
	Action    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignalSetPriceFeed is a free log retrieval operation binding the contract event 0xb878dd4b5762f4118ad54995be907dd2bcd915d942e4ac75580fba9b4ee4727f.
//
// Solidity: event SignalSetPriceFeed(address vault, address priceFeed, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalSetPriceFeed(opts *bind.FilterOpts) (*GmxTimelockSignalSetPriceFeedIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalSetPriceFeed")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalSetPriceFeedIterator{contract: _GmxTimelock.contract, event: "SignalSetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSignalSetPriceFeed is a free log subscription operation binding the contract event 0xb878dd4b5762f4118ad54995be907dd2bcd915d942e4ac75580fba9b4ee4727f.
//
// Solidity: event SignalSetPriceFeed(address vault, address priceFeed, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalSetPriceFeed(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalSetPriceFeed) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalSetPriceFeed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalSetPriceFeed)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalSetPriceFeed", log); err != nil {
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

// ParseSignalSetPriceFeed is a log parse operation binding the contract event 0xb878dd4b5762f4118ad54995be907dd2bcd915d942e4ac75580fba9b4ee4727f.
//
// Solidity: event SignalSetPriceFeed(address vault, address priceFeed, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalSetPriceFeed(log types.Log) (*GmxTimelockSignalSetPriceFeed, error) {
	event := new(GmxTimelockSignalSetPriceFeed)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalSetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalVaultSetTokenConfigIterator is returned from FilterSignalVaultSetTokenConfig and is used to iterate over the raw logs and unpacked data for SignalVaultSetTokenConfig events raised by the GmxTimelock contract.
type GmxTimelockSignalVaultSetTokenConfigIterator struct {
	Event *GmxTimelockSignalVaultSetTokenConfig // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalVaultSetTokenConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalVaultSetTokenConfig)
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
		it.Event = new(GmxTimelockSignalVaultSetTokenConfig)
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
func (it *GmxTimelockSignalVaultSetTokenConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalVaultSetTokenConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalVaultSetTokenConfig represents a SignalVaultSetTokenConfig event raised by the GmxTimelock contract.
type GmxTimelockSignalVaultSetTokenConfig struct {
	Vault         common.Address
	Token         common.Address
	TokenDecimals *big.Int
	TokenWeight   *big.Int
	MinProfitBps  *big.Int
	MaxUsdgAmount *big.Int
	IsStable      bool
	IsShortable   bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignalVaultSetTokenConfig is a free log retrieval operation binding the contract event 0x3510e9d8245371c6c1061c33781ce16bd0eafa03cd3d0781865036520af4c743.
//
// Solidity: event SignalVaultSetTokenConfig(address vault, address token, uint256 tokenDecimals, uint256 tokenWeight, uint256 minProfitBps, uint256 maxUsdgAmount, bool isStable, bool isShortable)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalVaultSetTokenConfig(opts *bind.FilterOpts) (*GmxTimelockSignalVaultSetTokenConfigIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalVaultSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalVaultSetTokenConfigIterator{contract: _GmxTimelock.contract, event: "SignalVaultSetTokenConfig", logs: logs, sub: sub}, nil
}

// WatchSignalVaultSetTokenConfig is a free log subscription operation binding the contract event 0x3510e9d8245371c6c1061c33781ce16bd0eafa03cd3d0781865036520af4c743.
//
// Solidity: event SignalVaultSetTokenConfig(address vault, address token, uint256 tokenDecimals, uint256 tokenWeight, uint256 minProfitBps, uint256 maxUsdgAmount, bool isStable, bool isShortable)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalVaultSetTokenConfig(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalVaultSetTokenConfig) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalVaultSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalVaultSetTokenConfig)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalVaultSetTokenConfig", log); err != nil {
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

// ParseSignalVaultSetTokenConfig is a log parse operation binding the contract event 0x3510e9d8245371c6c1061c33781ce16bd0eafa03cd3d0781865036520af4c743.
//
// Solidity: event SignalVaultSetTokenConfig(address vault, address token, uint256 tokenDecimals, uint256 tokenWeight, uint256 minProfitBps, uint256 maxUsdgAmount, bool isStable, bool isShortable)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalVaultSetTokenConfig(log types.Log) (*GmxTimelockSignalVaultSetTokenConfig, error) {
	event := new(GmxTimelockSignalVaultSetTokenConfig)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalVaultSetTokenConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxTimelockSignalWithdrawTokenIterator is returned from FilterSignalWithdrawToken and is used to iterate over the raw logs and unpacked data for SignalWithdrawToken events raised by the GmxTimelock contract.
type GmxTimelockSignalWithdrawTokenIterator struct {
	Event *GmxTimelockSignalWithdrawToken // Event containing the contract specifics and raw log

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
func (it *GmxTimelockSignalWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxTimelockSignalWithdrawToken)
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
		it.Event = new(GmxTimelockSignalWithdrawToken)
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
func (it *GmxTimelockSignalWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxTimelockSignalWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxTimelockSignalWithdrawToken represents a SignalWithdrawToken event raised by the GmxTimelock contract.
type GmxTimelockSignalWithdrawToken struct {
	Target   common.Address
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Action   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalWithdrawToken is a free log retrieval operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) FilterSignalWithdrawToken(opts *bind.FilterOpts) (*GmxTimelockSignalWithdrawTokenIterator, error) {

	logs, sub, err := _GmxTimelock.contract.FilterLogs(opts, "SignalWithdrawToken")
	if err != nil {
		return nil, err
	}
	return &GmxTimelockSignalWithdrawTokenIterator{contract: _GmxTimelock.contract, event: "SignalWithdrawToken", logs: logs, sub: sub}, nil
}

// WatchSignalWithdrawToken is a free log subscription operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) WatchSignalWithdrawToken(opts *bind.WatchOpts, sink chan<- *GmxTimelockSignalWithdrawToken) (event.Subscription, error) {

	logs, sub, err := _GmxTimelock.contract.WatchLogs(opts, "SignalWithdrawToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxTimelockSignalWithdrawToken)
				if err := _GmxTimelock.contract.UnpackLog(event, "SignalWithdrawToken", log); err != nil {
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

// ParseSignalWithdrawToken is a log parse operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_GmxTimelock *GmxTimelockFilterer) ParseSignalWithdrawToken(log types.Log) (*GmxTimelockSignalWithdrawToken, error) {
	event := new(GmxTimelockSignalWithdrawToken)
	if err := _GmxTimelock.contract.UnpackLog(event, "SignalWithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
