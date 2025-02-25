/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import { Provider } from "@ethersproject/providers";
import type { IPythEvents, IPythEventsInterface } from "../IPythEvents";

const _abi = [
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "bytes32",
        name: "id",
        type: "bytes32",
      },
      {
        indexed: false,
        internalType: "uint64",
        name: "publishTime",
        type: "uint64",
      },
      {
        indexed: false,
        internalType: "int64",
        name: "price",
        type: "int64",
      },
      {
        indexed: false,
        internalType: "uint64",
        name: "conf",
        type: "uint64",
      },
    ],
    name: "PriceFeedUpdate",
    type: "event",
  },
];

export class IPythEvents__factory {
  static readonly abi = _abi;
  static createInterface(): IPythEventsInterface {
    return new utils.Interface(_abi) as IPythEventsInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): IPythEvents {
    return new Contract(address, _abi, signerOrProvider) as IPythEvents;
  }
}
