/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type { BaseMigrator, BaseMigratorInterface } from "../BaseMigrator";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_admin",
        type: "address",
      },
      {
        internalType: "address",
        name: "_stakedGmxTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_bonusGmxTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_feeGmxTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_stakedGlpTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_feeGlpTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_gmxVester",
        type: "address",
      },
      {
        internalType: "address",
        name: "_glpVester",
        type: "address",
      },
      {
        internalType: "address",
        name: "_esGmx",
        type: "address",
      },
      {
        internalType: "address",
        name: "_bnGmx",
        type: "address",
      },
      {
        internalType: "address",
        name: "_rewardRouter",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "admin",
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
    name: "afterGovGranted",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "bnGmx",
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
    name: "bonusGmxTracker",
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
    name: "esGmx",
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
    name: "expectedGovGrantedCaller",
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
    name: "feeGlpTracker",
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
    name: "feeGmxTracker",
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
    name: "glpVester",
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
    name: "gmxVester",
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
    name: "migrate",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "rewardRouter",
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
    name: "stakedGlpTracker",
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
    name: "stakedGmxTracker",
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
];

const _bytecode =
  "0x6101e060405234801561001157600080fd5b50604051611665380380611665833981810160405261016081101561003557600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291905050508a6001600160a01b03166080816001600160a01b031660601b81525050896001600160a01b031660a0816001600160a01b031660601b81525050886001600160a01b031660c0816001600160a01b031660601b81525050876001600160a01b031660e0816001600160a01b031660601b81525050866001600160a01b0316610100816001600160a01b031660601b81525050856001600160a01b0316610120816001600160a01b031660601b81525050846001600160a01b0316610140816001600160a01b031660601b81525050836001600160a01b0316610160816001600160a01b031660601b81525050826001600160a01b0316610180816001600160a01b031660601b81525050816001600160a01b03166101a0816001600160a01b031660601b81525050806001600160a01b03166101c0816001600160a01b031660601b81525050505050505050505050505060805160601c60a05160601c60c05160601c60e05160601c6101005160601c6101205160601c6101405160601c6101605160601c6101805160601c6101a05160601c6101c05160601c611310610355600039806102035280610ccd5280610d7e5280610e2f5280610ee05280610f91528061104252806110f352806111a452806112555250806105ef5280610b985280610c32528061122652508061022752806105a15280610b11528061117552508061019752806105535280610a8a52806110c45250806101bb52806105055280610a0352806110135250806104b7528061097c5280610c565280610f6252508061046952806106db52806108f55280610eb15250806101df528061041b528061086e5280610e0052508061017352806103cd52806107e75280610d4f52508061014f52806102b6528061037f52806107865280610c9e5250806102545280610c7a52506113106000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c80630ce4018a146100bf5780631fcd60e5146100e35780633671df25146100eb57806341d315b4146100f357806351c3e3b4146100fb5780635a3bb989146101035780636a192a781461010b5780638fd3ab80146101135780639a2e08941461011d578063af394d0014610125578063c7cfeb591461012d578063e102f56414610135578063e1c363b71461013d578063f851a44014610145575b600080fd5b6100c761014d565b604080516001600160a01b039092168252519081900360200190f35b6100c7610171565b6100c7610195565b6100c76101b9565b6100c76101dd565b6100c7610201565b6100c7610225565b61011b610249565b005b6100c76106ca565b6100c76106d9565b61011b6106fd565b6100c7610c30565b6100c7610c54565b6100c7610c78565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102b2576040805162461bcd60e51b81526020600482015260096024820152683337b93134b23232b760b91b604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166312d43a516040518163ffffffff1660e01b815260040160206040518083038186803b15801561030d57600080fd5b505afa158015610321573d6000803e3d6000fd5b505050506040513d602081101561033757600080fd5b5051600080546001600160a01b0319166001600160a01b03831617905560408051600980825261014082019092529192506060919060208201610120803683370190505090507f0000000000000000000000000000000000000000000000000000000000000000816000815181106103ab57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106103f957fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160028151811061044757fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160038151811061049557fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f0000000000000000000000000000000000000000000000000000000000000000816004815181106104e357fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160058151811061053157fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160068151811061057f57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f0000000000000000000000000000000000000000000000000000000000000000816007815181106105cd57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160088151811061061b57fe5b6001600160a01b0392831660209182029290920181019190915260405163de8616ff60e01b8152600481018281528451602483015284519386169363de8616ff9386938392604490910191858101910280838360005b83811015610689578181015183820152602001610671565b5050505090500192505050600060405180830381600087803b1580156106ae57600080fd5b505af11580156106c2573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b03163314610748576040805162461bcd60e51b81526020600482015260096024820152683337b93134b23232b760b91b604482015290519081900360640190fd5b6107526001610c9c565b61075a6112d8565b6107646000610c9c565b604080516367d6abd160e11b8152336004820181905291516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163cfad57a291602480830192600092919082900301818387803b1580156107cd57600080fd5b505af11580156107e1573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561085457600080fd5b505af1158015610868573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156108db57600080fd5b505af11580156108ef573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561096257600080fd5b505af1158015610976573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156109e957600080fd5b505af11580156109fd573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015610a7057600080fd5b505af1158015610a84573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015610af757600080fd5b505af1158015610b0b573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015610b7e57600080fd5b505af1158015610b92573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015610c0557600080fd5b505af1158015610c19573d6000803e3d6000fd5b5050600080546001600160a01b0319169055505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610d3557600080fd5b505af1158015610d49573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610de657600080fd5b505af1158015610dfa573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610e9757600080fd5b505af1158015610eab573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610f4857600080fd5b505af1158015610f5c573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610ff957600080fd5b505af115801561100d573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b1580156110aa57600080fd5b505af11580156110be573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b15801561115b57600080fd5b505af115801561116f573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b15801561120c57600080fd5b505af1158015611220573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cf456ae77f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b1580156112bd57600080fd5b505af11580156112d1573d6000803e3d6000fd5b5050505050565b56fea2646970667358221220eb79df5c89d2e60bb8e2894faefb2163665137e0650e6e9d92cafa660ffe7e3364736f6c634300060c0033";

export class BaseMigrator__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    _admin: string,
    _stakedGmxTracker: string,
    _bonusGmxTracker: string,
    _feeGmxTracker: string,
    _stakedGlpTracker: string,
    _feeGlpTracker: string,
    _gmxVester: string,
    _glpVester: string,
    _esGmx: string,
    _bnGmx: string,
    _rewardRouter: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<BaseMigrator> {
    return super.deploy(
      _admin,
      _stakedGmxTracker,
      _bonusGmxTracker,
      _feeGmxTracker,
      _stakedGlpTracker,
      _feeGlpTracker,
      _gmxVester,
      _glpVester,
      _esGmx,
      _bnGmx,
      _rewardRouter,
      overrides || {}
    ) as Promise<BaseMigrator>;
  }
  getDeployTransaction(
    _admin: string,
    _stakedGmxTracker: string,
    _bonusGmxTracker: string,
    _feeGmxTracker: string,
    _stakedGlpTracker: string,
    _feeGlpTracker: string,
    _gmxVester: string,
    _glpVester: string,
    _esGmx: string,
    _bnGmx: string,
    _rewardRouter: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      _admin,
      _stakedGmxTracker,
      _bonusGmxTracker,
      _feeGmxTracker,
      _stakedGlpTracker,
      _feeGlpTracker,
      _gmxVester,
      _glpVester,
      _esGmx,
      _bnGmx,
      _rewardRouter,
      overrides || {}
    );
  }
  attach(address: string): BaseMigrator {
    return super.attach(address) as BaseMigrator;
  }
  connect(signer: Signer): BaseMigrator__factory {
    return super.connect(signer) as BaseMigrator__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): BaseMigratorInterface {
    return new utils.Interface(_abi) as BaseMigratorInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): BaseMigrator {
    return new Contract(address, _abi, signerOrProvider) as BaseMigrator;
  }
}