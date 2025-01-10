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

interface USDTPriceFeedInterface extends ethers.utils.Interface {
  functions: {
    "aggregator()": FunctionFragment;
    "answer()": FunctionFragment;
    "answers(uint80)": FunctionFragment;
    "decimals()": FunctionFragment;
    "description()": FunctionFragment;
    "getRoundData(uint80)": FunctionFragment;
    "gov()": FunctionFragment;
    "isAdmin(address)": FunctionFragment;
    "latestAnswer()": FunctionFragment;
    "latestRound()": FunctionFragment;
    "roundId()": FunctionFragment;
    "setAdmin(address,bool)": FunctionFragment;
    "setLatestAnswer(int256)": FunctionFragment;
  };

  encodeFunctionData(
    functionFragment: "aggregator",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "answer", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "answers",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "decimals", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "description",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getRoundData",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "gov", values?: undefined): string;
  encodeFunctionData(functionFragment: "isAdmin", values: [string]): string;
  encodeFunctionData(
    functionFragment: "latestAnswer",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "latestRound",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "roundId", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "setAdmin",
    values: [string, boolean]
  ): string;
  encodeFunctionData(
    functionFragment: "setLatestAnswer",
    values: [BigNumberish]
  ): string;

  decodeFunctionResult(functionFragment: "aggregator", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "answer", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "answers", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "decimals", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "description",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getRoundData",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "gov", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "isAdmin", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "latestAnswer",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "latestRound",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "roundId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "setAdmin", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "setLatestAnswer",
    data: BytesLike
  ): Result;

  events: {};
}

export class USDTPriceFeed extends BaseContract {
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

  interface: USDTPriceFeedInterface;

  functions: {
    aggregator(overrides?: CallOverrides): Promise<[string]>;

    answer(overrides?: CallOverrides): Promise<[BigNumber]>;

    answers(
      arg0: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    decimals(overrides?: CallOverrides): Promise<[BigNumber]>;

    description(overrides?: CallOverrides): Promise<[string]>;

    getRoundData(
      _roundId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber, BigNumber, BigNumber, BigNumber, BigNumber]>;

    gov(overrides?: CallOverrides): Promise<[string]>;

    isAdmin(arg0: string, overrides?: CallOverrides): Promise<[boolean]>;

    latestAnswer(overrides?: CallOverrides): Promise<[BigNumber]>;

    latestRound(overrides?: CallOverrides): Promise<[BigNumber]>;

    roundId(overrides?: CallOverrides): Promise<[BigNumber]>;

    setAdmin(
      _account: string,
      _isAdmin: boolean,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    setLatestAnswer(
      _answer: BigNumberish,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;
  };

  aggregator(overrides?: CallOverrides): Promise<string>;

  answer(overrides?: CallOverrides): Promise<BigNumber>;

  answers(arg0: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

  decimals(overrides?: CallOverrides): Promise<BigNumber>;

  description(overrides?: CallOverrides): Promise<string>;

  getRoundData(
    _roundId: BigNumberish,
    overrides?: CallOverrides
  ): Promise<[BigNumber, BigNumber, BigNumber, BigNumber, BigNumber]>;

  gov(overrides?: CallOverrides): Promise<string>;

  isAdmin(arg0: string, overrides?: CallOverrides): Promise<boolean>;

  latestAnswer(overrides?: CallOverrides): Promise<BigNumber>;

  latestRound(overrides?: CallOverrides): Promise<BigNumber>;

  roundId(overrides?: CallOverrides): Promise<BigNumber>;

  setAdmin(
    _account: string,
    _isAdmin: boolean,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  setLatestAnswer(
    _answer: BigNumberish,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    aggregator(overrides?: CallOverrides): Promise<string>;

    answer(overrides?: CallOverrides): Promise<BigNumber>;

    answers(arg0: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    decimals(overrides?: CallOverrides): Promise<BigNumber>;

    description(overrides?: CallOverrides): Promise<string>;

    getRoundData(
      _roundId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber, BigNumber, BigNumber, BigNumber, BigNumber]>;

    gov(overrides?: CallOverrides): Promise<string>;

    isAdmin(arg0: string, overrides?: CallOverrides): Promise<boolean>;

    latestAnswer(overrides?: CallOverrides): Promise<BigNumber>;

    latestRound(overrides?: CallOverrides): Promise<BigNumber>;

    roundId(overrides?: CallOverrides): Promise<BigNumber>;

    setAdmin(
      _account: string,
      _isAdmin: boolean,
      overrides?: CallOverrides
    ): Promise<void>;

    setLatestAnswer(
      _answer: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {};

  estimateGas: {
    aggregator(overrides?: CallOverrides): Promise<BigNumber>;

    answer(overrides?: CallOverrides): Promise<BigNumber>;

    answers(arg0: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    decimals(overrides?: CallOverrides): Promise<BigNumber>;

    description(overrides?: CallOverrides): Promise<BigNumber>;

    getRoundData(
      _roundId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    gov(overrides?: CallOverrides): Promise<BigNumber>;

    isAdmin(arg0: string, overrides?: CallOverrides): Promise<BigNumber>;

    latestAnswer(overrides?: CallOverrides): Promise<BigNumber>;

    latestRound(overrides?: CallOverrides): Promise<BigNumber>;

    roundId(overrides?: CallOverrides): Promise<BigNumber>;

    setAdmin(
      _account: string,
      _isAdmin: boolean,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    setLatestAnswer(
      _answer: BigNumberish,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    aggregator(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    answer(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    answers(
      arg0: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    decimals(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    description(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    getRoundData(
      _roundId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    gov(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    isAdmin(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    latestAnswer(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    latestRound(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    roundId(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    setAdmin(
      _account: string,
      _isAdmin: boolean,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    setLatestAnswer(
      _answer: BigNumberish,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;
  };
}
