/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type { USDTPriceFeed, USDTPriceFeedInterface } from "../USDTPriceFeed";

const _abi = [
  {
    inputs: [],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "aggregator",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "answer",
    outputs: [
      {
        internalType: "int256",
        name: "",
        type: "int256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint80",
        name: "",
        type: "uint80",
      },
    ],
    name: "answers",
    outputs: [
      {
        internalType: "int256",
        name: "",
        type: "int256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "decimals",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "description",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint80",
        name: "_roundId",
        type: "uint80",
      },
    ],
    name: "getRoundData",
    outputs: [
      {
        internalType: "uint80",
        name: "",
        type: "uint80",
      },
      {
        internalType: "int256",
        name: "",
        type: "int256",
      },
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
      {
        internalType: "uint80",
        name: "",
        type: "uint80",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "gov",
    outputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "isAdmin",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "latestAnswer",
    outputs: [
      {
        internalType: "int256",
        name: "",
        type: "int256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "latestRound",
    outputs: [
      {
        internalType: "uint80",
        name: "",
        type: "uint80",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "roundId",
    outputs: [
      {
        internalType: "uint80",
        name: "",
        type: "uint80",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_account",
        type: "address",
      },
      {
        internalType: "bool",
        name: "_isAdmin",
        type: "bool",
      },
    ],
    name: "setAdmin",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "int256",
        name: "_answer",
        type: "int256",
      },
    ],
    name: "setLatestAnswer",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x60c06040526009608081905268141c9a58d95199595960ba1b60a090815261002a916002919061006b565b5034801561003757600080fd5b50600580546001600160a01b031916339081179091556000908152600760205260409020805460ff191660011790556100fe565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ac57805160ff19168380011785556100d9565b828001600101855582156100d9579182015b828111156100d95782518255916020019190600101906100be565b506100e59291506100e9565b5090565b5b808211156100e557600081556001016100ea565b6105358061010d6000396000f3fe608060405234801561001057600080fd5b50600436106100af5760003560e01c806304ea97b0146100b457806312d43a51146100d3578063245a7bfc146100f757806324d7806c146100ff578063313ce567146101395780634b0bddd2146101535780634c295ca31461018157806350d25bcd146101a7578063668a0f02146101af5780637284e416146101d357806385bb7d69146102505780638cd221c9146102585780639a6fc8f514610260575b600080fd5b6100d1600480360360208110156100ca57600080fd5b50356102ca565b005b6100db61035d565b604080516001600160a01b039092168252519081900360200190f35b6100db61036c565b6101256004803603602081101561011557600080fd5b50356001600160a01b031661037b565b604080519115158252519081900360200190f35b610141610390565b60408051918252519081900360200190f35b6100d16004803603604081101561016957600080fd5b506001600160a01b0381351690602001351515610396565b6101416004803603602081101561019757600080fd5b50356001600160501b0316610417565b610141610429565b6101b761042f565b604080516001600160501b039092168252519081900360200190f35b6101db61043e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102155781810151838201526020016101fd565b50505050905090810190601f1680156102425780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101416104c9565b6101b76104cf565b6102866004803603602081101561027657600080fd5b50356001600160501b03166104de565b60405180866001600160501b03168152602001858152602001848152602001838152602001826001600160501b031681526020019550505050505060405180910390f35b3360009081526007602052604090205460ff16610325576040805162461bcd60e51b8152602060048201526014602482015273283934b1b2a332b2b21d103337b93134b23232b760611b604482015290519081900360640190fd5b600180546001600160501b031981166001600160501b0391821683018216179182905560008381559116815260066020526040902055565b6005546001600160a01b031681565b6003546001600160a01b031681565b60076020526000908152604090205460ff1681565b60045481565b6005546001600160a01b031633146103ec576040805162461bcd60e51b8152602060048201526014602482015273283934b1b2a332b2b21d103337b93134b23232b760611b604482015290519081900360640190fd5b6001600160a01b03919091166000908152600760205260409020805460ff1916911515919091179055565b60066020526000908152604090205481565b60005490565b6001546001600160501b031690565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156104c15780601f10610496576101008083540402835291602001916104c1565b820191906000526020600020905b8154815290600101906020018083116104a457829003601f168201915b505050505081565b60005481565b6001546001600160501b031681565b6001600160501b03811660009081526006602052604081205491928190819056fea264697066735822122046bf25c5c6ed11fbd81ba17cca25bfb69f4cb182ee3ee80ff9285ee4c224a3e064736f6c634300060c0033";

export class USDTPriceFeed__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<USDTPriceFeed> {
    return super.deploy(overrides || {}) as Promise<USDTPriceFeed>;
  }
  getDeployTransaction(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  attach(address: string): USDTPriceFeed {
    return super.attach(address) as USDTPriceFeed;
  }
  connect(signer: Signer): USDTPriceFeed__factory {
    return super.connect(signer) as USDTPriceFeed__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): USDTPriceFeedInterface {
    return new utils.Interface(_abi) as USDTPriceFeedInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): USDTPriceFeed {
    return new Contract(address, _abi, signerOrProvider) as USDTPriceFeed;
  }
}
