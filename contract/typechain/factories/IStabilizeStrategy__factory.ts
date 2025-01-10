/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import { Provider } from "@ethersproject/providers";
import type {
  IStabilizeStrategy,
  IStabilizeStrategyInterface,
} from "../IStabilizeStrategy";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_sender",
        type: "address",
      },
    ],
    name: "governanceFinishMoveEsGMXFromDeprecatedRouter",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

export class IStabilizeStrategy__factory {
  static readonly abi = _abi;
  static createInterface(): IStabilizeStrategyInterface {
    return new utils.Interface(_abi) as IStabilizeStrategyInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IStabilizeStrategy {
    return new Contract(address, _abi, signerOrProvider) as IStabilizeStrategy;
  }
}
