/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import {
  ethers,
  EventFilter,
  Signer,
  BigNumber,
  BigNumberish,
  PopulatedTransaction,
  BaseContract,
  ContractTransaction,
  Overrides,
  CallOverrides,
} from "ethers";
import { BytesLike } from "@ethersproject/bytes";
import { Listener, Provider } from "@ethersproject/providers";
import { FunctionFragment, EventFragment, Result } from "@ethersproject/abi";
import { TypedEventFilter, TypedEvent, TypedListener } from "./commons";

interface BuybackMigratorInterface extends ethers.utils.Interface {
  functions: {
    "admin()": FunctionFragment;
    "afterGovGranted()": FunctionFragment;
    "bnGmx()": FunctionFragment;
    "bonusGmxTracker()": FunctionFragment;
    "disableOldRewardRouter()": FunctionFragment;
    "enableNewRewardRouter()": FunctionFragment;
    "esGmx()": FunctionFragment;
    "expectedGovGrantedCaller()": FunctionFragment;
    "extendedGmxTracker()": FunctionFragment;
    "feeGlpTracker()": FunctionFragment;
    "feeGmxTracker()": FunctionFragment;
    "glpVester()": FunctionFragment;
    "gmxVester()": FunctionFragment;
    "isEnabled()": FunctionFragment;
    "isRestakingCompleted()": FunctionFragment;
    "newRewardRouter()": FunctionFragment;
    "oldRewardRouter()": FunctionFragment;
    "rewardRouterTarget()": FunctionFragment;
    "setHandlerAndDepositToken()": FunctionFragment;
    "stakedGlpTracker()": FunctionFragment;
    "stakedGmxTracker()": FunctionFragment;
  };

  encodeFunctionData(functionFragment: "admin", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "afterGovGranted",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "bnGmx", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "bonusGmxTracker",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "disableOldRewardRouter",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "enableNewRewardRouter",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "esGmx", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "expectedGovGrantedCaller",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "extendedGmxTracker",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "feeGlpTracker",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "feeGmxTracker",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "glpVester", values?: undefined): string;
  encodeFunctionData(functionFragment: "gmxVester", values?: undefined): string;
  encodeFunctionData(functionFragment: "isEnabled", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "isRestakingCompleted",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "newRewardRouter",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "oldRewardRouter",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "rewardRouterTarget",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "setHandlerAndDepositToken",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "stakedGlpTracker",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "stakedGmxTracker",
    values?: undefined
  ): string;

  decodeFunctionResult(functionFragment: "admin", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "afterGovGranted",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "bnGmx", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "bonusGmxTracker",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "disableOldRewardRouter",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "enableNewRewardRouter",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "esGmx", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "expectedGovGrantedCaller",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "extendedGmxTracker",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "feeGlpTracker",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "feeGmxTracker",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "glpVester", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "gmxVester", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "isEnabled", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "isRestakingCompleted",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "newRewardRouter",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "oldRewardRouter",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "rewardRouterTarget",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setHandlerAndDepositToken",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "stakedGlpTracker",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "stakedGmxTracker",
    data: BytesLike
  ): Result;

  events: {};
}

export class BuybackMigrator extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  listeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter?: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): Array<TypedListener<EventArgsArray, EventArgsObject>>;
  off<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  on<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  once<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  removeListener<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  removeAllListeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): this;

  listeners(eventName?: string): Array<Listener>;
  off(eventName: string, listener: Listener): this;
  on(eventName: string, listener: Listener): this;
  once(eventName: string, listener: Listener): this;
  removeListener(eventName: string, listener: Listener): this;
  removeAllListeners(eventName?: string): this;

  queryFilter<EventArgsArray extends Array<any>, EventArgsObject>(
    event: TypedEventFilter<EventArgsArray, EventArgsObject>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEvent<EventArgsArray & EventArgsObject>>>;

  interface: BuybackMigratorInterface;

  functions: {
    admin(overrides?: CallOverrides): Promise<[string]>;

    afterGovGranted(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    bnGmx(overrides?: CallOverrides): Promise<[string]>;

    bonusGmxTracker(overrides?: CallOverrides): Promise<[string]>;

    disableOldRewardRouter(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    enableNewRewardRouter(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    esGmx(overrides?: CallOverrides): Promise<[string]>;

    expectedGovGrantedCaller(overrides?: CallOverrides): Promise<[string]>;

    extendedGmxTracker(overrides?: CallOverrides): Promise<[string]>;

    feeGlpTracker(overrides?: CallOverrides): Promise<[string]>;

    feeGmxTracker(overrides?: CallOverrides): Promise<[string]>;

    glpVester(overrides?: CallOverrides): Promise<[string]>;

    gmxVester(overrides?: CallOverrides): Promise<[string]>;

    isEnabled(overrides?: CallOverrides): Promise<[boolean]>;

    isRestakingCompleted(overrides?: CallOverrides): Promise<[boolean]>;

    newRewardRouter(overrides?: CallOverrides): Promise<[string]>;

    oldRewardRouter(overrides?: CallOverrides): Promise<[string]>;

    rewardRouterTarget(overrides?: CallOverrides): Promise<[string]>;

    setHandlerAndDepositToken(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    stakedGlpTracker(overrides?: CallOverrides): Promise<[string]>;

    stakedGmxTracker(overrides?: CallOverrides): Promise<[string]>;
  };

  admin(overrides?: CallOverrides): Promise<string>;

  afterGovGranted(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  bnGmx(overrides?: CallOverrides): Promise<string>;

  bonusGmxTracker(overrides?: CallOverrides): Promise<string>;

  disableOldRewardRouter(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  enableNewRewardRouter(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  esGmx(overrides?: CallOverrides): Promise<string>;

  expectedGovGrantedCaller(overrides?: CallOverrides): Promise<string>;

  extendedGmxTracker(overrides?: CallOverrides): Promise<string>;

  feeGlpTracker(overrides?: CallOverrides): Promise<string>;

  feeGmxTracker(overrides?: CallOverrides): Promise<string>;

  glpVester(overrides?: CallOverrides): Promise<string>;

  gmxVester(overrides?: CallOverrides): Promise<string>;

  isEnabled(overrides?: CallOverrides): Promise<boolean>;

  isRestakingCompleted(overrides?: CallOverrides): Promise<boolean>;

  newRewardRouter(overrides?: CallOverrides): Promise<string>;

  oldRewardRouter(overrides?: CallOverrides): Promise<string>;

  rewardRouterTarget(overrides?: CallOverrides): Promise<string>;

  setHandlerAndDepositToken(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  stakedGlpTracker(overrides?: CallOverrides): Promise<string>;

  stakedGmxTracker(overrides?: CallOverrides): Promise<string>;

  callStatic: {
    admin(overrides?: CallOverrides): Promise<string>;

    afterGovGranted(overrides?: CallOverrides): Promise<void>;

    bnGmx(overrides?: CallOverrides): Promise<string>;

    bonusGmxTracker(overrides?: CallOverrides): Promise<string>;

    disableOldRewardRouter(overrides?: CallOverrides): Promise<void>;

    enableNewRewardRouter(overrides?: CallOverrides): Promise<void>;

    esGmx(overrides?: CallOverrides): Promise<string>;

    expectedGovGrantedCaller(overrides?: CallOverrides): Promise<string>;

    extendedGmxTracker(overrides?: CallOverrides): Promise<string>;

    feeGlpTracker(overrides?: CallOverrides): Promise<string>;

    feeGmxTracker(overrides?: CallOverrides): Promise<string>;

    glpVester(overrides?: CallOverrides): Promise<string>;

    gmxVester(overrides?: CallOverrides): Promise<string>;

    isEnabled(overrides?: CallOverrides): Promise<boolean>;

    isRestakingCompleted(overrides?: CallOverrides): Promise<boolean>;

    newRewardRouter(overrides?: CallOverrides): Promise<string>;

    oldRewardRouter(overrides?: CallOverrides): Promise<string>;

    rewardRouterTarget(overrides?: CallOverrides): Promise<string>;

    setHandlerAndDepositToken(overrides?: CallOverrides): Promise<void>;

    stakedGlpTracker(overrides?: CallOverrides): Promise<string>;

    stakedGmxTracker(overrides?: CallOverrides): Promise<string>;
  };

  filters: {};

  estimateGas: {
    admin(overrides?: CallOverrides): Promise<BigNumber>;

    afterGovGranted(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    bnGmx(overrides?: CallOverrides): Promise<BigNumber>;

    bonusGmxTracker(overrides?: CallOverrides): Promise<BigNumber>;

    disableOldRewardRouter(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    enableNewRewardRouter(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    esGmx(overrides?: CallOverrides): Promise<BigNumber>;

    expectedGovGrantedCaller(overrides?: CallOverrides): Promise<BigNumber>;

    extendedGmxTracker(overrides?: CallOverrides): Promise<BigNumber>;

    feeGlpTracker(overrides?: CallOverrides): Promise<BigNumber>;

    feeGmxTracker(overrides?: CallOverrides): Promise<BigNumber>;

    glpVester(overrides?: CallOverrides): Promise<BigNumber>;

    gmxVester(overrides?: CallOverrides): Promise<BigNumber>;

    isEnabled(overrides?: CallOverrides): Promise<BigNumber>;

    isRestakingCompleted(overrides?: CallOverrides): Promise<BigNumber>;

    newRewardRouter(overrides?: CallOverrides): Promise<BigNumber>;

    oldRewardRouter(overrides?: CallOverrides): Promise<BigNumber>;

    rewardRouterTarget(overrides?: CallOverrides): Promise<BigNumber>;

    setHandlerAndDepositToken(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    stakedGlpTracker(overrides?: CallOverrides): Promise<BigNumber>;

    stakedGmxTracker(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    admin(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    afterGovGranted(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    bnGmx(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    bonusGmxTracker(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    disableOldRewardRouter(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    enableNewRewardRouter(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    esGmx(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    expectedGovGrantedCaller(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    extendedGmxTracker(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    feeGlpTracker(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    feeGmxTracker(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    glpVester(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    gmxVester(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    isEnabled(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    isRestakingCompleted(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    newRewardRouter(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    oldRewardRouter(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    rewardRouterTarget(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    setHandlerAndDepositToken(
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    stakedGlpTracker(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    stakedGmxTracker(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}
