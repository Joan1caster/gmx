/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type {
  BuybackMigrator,
  BuybackMigratorInterface,
} from "../BuybackMigrator";

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
        name: "_extendedGmxTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_feeGmxTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_feeGlpTracker",
        type: "address",
      },
      {
        internalType: "address",
        name: "_stakedGlpTracker",
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
        name: "_oldRewardRouter",
        type: "address",
      },
      {
        internalType: "address",
        name: "_newRewardRouter",
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
    name: "disableOldRewardRouter",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [],
    name: "enableNewRewardRouter",
    outputs: [],
    stateMutability: "nonpayable",
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
    name: "extendedGmxTracker",
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
    name: "isEnabled",
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
    name: "isRestakingCompleted",
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
    name: "newRewardRouter",
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
    name: "oldRewardRouter",
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
    name: "rewardRouterTarget",
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
    name: "setHandlerAndDepositToken",
    outputs: [],
    stateMutability: "nonpayable",
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
  "0x61022060405234801561001157600080fd5b5060405161254638038061254683398181016040526101a081101561003557600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291905050508c6001600160a01b03166080816001600160a01b031660601b815250508b6001600160a01b031660a0816001600160a01b031660601b815250508a6001600160a01b031660c0816001600160a01b031660601b81525050896001600160a01b031660e0816001600160a01b031660601b81525050886001600160a01b0316610100816001600160a01b031660601b81525050876001600160a01b0316610120816001600160a01b031660601b81525050866001600160a01b0316610140816001600160a01b031660601b81525050856001600160a01b0316610160816001600160a01b031660601b81525050846001600160a01b0316610180816001600160a01b031660601b81525050836001600160a01b03166101a0816001600160a01b031660601b81525050826001600160a01b03166101c0816001600160a01b031660601b81525050816001600160a01b03166101e0816001600160a01b031660601b81525050806001600160a01b0316610200816001600160a01b031660601b815250505050505050505050505050505060805160601c60a05160601c60c05160601c60e05160601c6101005160601c6101205160601c6101405160601c6101605160601c6101805160601c6101a05160601c6101c05160601c6101e05160601c6102005160601c61211761042f6000398061040052806108a0525080610805528061178b5250806107285280610ad75280610c645280610df75280610f01528061153452806115df5280611d7e5280611eaa5250806106da528061086c52806114ad5280611ce5525080610254528061068c52806114265280611c4c525080610278528061063e528061139f5280611bb35250806105f052806108d352806113185280611b1a5250806105a2528061129152806116035280611a8152508061055452806108295280610b255280610be15280610c935280610d165280610dc85280610f88528061120a52806119e85280611f8b528061200e52508061023052806105065280611183528061194f5280611e275280611ed95280611f5c528061203d52508061020c52806104b85280610a895280610bb25280610d455280610e7a52806110fc52806118b65280611df85250806101e85280610353528061046a52806109ae528061107552806116de528061181f5250806102a55280610900528061163052806117c652506121176000f3fe608060405234801561001057600080fd5b50600436106101075760003560e01c80630ce4018a1461010c5780631fcd60e514610130578063204b4c6f146101385780633671df251461014057806341d315b41461014857806345163ae81461015057806345914f601461015a57806351c3e3b4146101625780635b9c37e01461016a5780635f0175bf146101865780636a192a781461018e5780636aa633b61461019657806398b730011461019e5780639a2e0894146101a6578063af394d00146101ae578063be282afc146101b6578063c7cfeb59146101be578063e102f564146101c6578063e1c363b7146101ce578063eb6a305a146101d6578063f851a440146101de575b600080fd5b6101146101e6565b604080516001600160a01b039092168252519081900360200190f35b61011461020a565b61011461022e565b610114610252565b610114610276565b61015861029a565b005b610114610803565b610114610827565b61017261084b565b604080519115158252519081900360200190f35b61011461085b565b61011461086a565b61017261088e565b61011461089e565b6101146108c2565b6101146108d1565b6101586108f5565b610158610b51565b6101146115dd565b610114611601565b610158611625565b6101146117c4565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610303576040805162461bcd60e51b81526020600482015260096024820152683337b93134b23232b760b91b604482015290519081900360640190fd5b6001546001600160a01b03161561034f576040805162461bcd60e51b815260206004820152601a60248201526000805160206120c2833981519152604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166312d43a516040518163ffffffff1660e01b815260040160206040518083038186803b1580156103aa57600080fd5b505afa1580156103be573d6000803e3d6000fd5b505050506040513d60208110156103d457600080fd5b5051600080546001600160a01b038084166001600160a01b03199283161783556001805460ff60a01b197f000000000000000000000000000000000000000000000000000000000000000090931693169290921716600160a01b179055909150606090600a905b50604051908082528060200260200182016040528015610465578160200160208202803683370190505b5090507f00000000000000000000000000000000000000000000000000000000000000008160008151811061049657fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106104e457fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160028151811061053257fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160038151811061058057fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f0000000000000000000000000000000000000000000000000000000000000000816004815181106105ce57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160058151811061061c57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160068151811061066a57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f0000000000000000000000000000000000000000000000000000000000000000816007815181106106b857fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160088151811061070657fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160098151811061075457fe5b6001600160a01b0392831660209182029290920181019190915260405163de8616ff60e01b8152600481018281528451602483015284519386169363de8616ff9386938392604490910191858101910280838360005b838110156107c25781810151838201526020016107aa565b5050505090500192505050600060405180830381600087803b1580156107e757600080fd5b505af11580156107fb573d6000803e3d6000fd5b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600154600160a81b900460ff1681565b6001546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600154600160a01b900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461095e576040805162461bcd60e51b81526020600482015260096024820152683337b93134b23232b760b91b604482015290519081900360640190fd5b6001546001600160a01b0316156109aa576040805162461bcd60e51b815260206004820152601a60248201526000805160206120c2833981519152604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166312d43a516040518163ffffffff1660e01b815260040160206040518083038186803b158015610a0557600080fd5b505afa158015610a19573d6000803e3d6000fd5b505050506040513d6020811015610a2f57600080fd5b5051600080546001600160a01b0383166001600160a01b03199091161790556001805460ff60a81b1916600160a81b17905560408051600380825260808201909252919250606091906020820183803683370190505090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610ab557fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f000000000000000000000000000000000000000000000000000000000000000081600181518110610b0357fe5b60200260200101906001600160a01b031690816001600160a01b0316815250507f00000000000000000000000000000000000000000000000000000000000000008160028151811061075457fe5b60005433906001600160a01b03168114610b9e576040805162461bcd60e51b81526020600482015260096024820152683337b93134b23232b760b91b604482015290519081900360640190fd5b600154600160a81b900460ff1615611020577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f000000000000000000000000000000000000000000000000000000000000000060006040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610c4a57600080fd5b505af1158015610c5e573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f000000000000000000000000000000000000000000000000000000000000000060006040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610cfc57600080fd5b505af1158015610d10573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e44b75587f000000000000000000000000000000000000000000000000000000000000000060006040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610dae57600080fd5b505af1158015610dc2573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e44b75587f000000000000000000000000000000000000000000000000000000000000000060006040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015610e6057600080fd5b505af1158015610e74573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015610ee757600080fd5b505af1158015610efb573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015610f6e57600080fd5b505af1158015610f82573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015610ff557600080fd5b505af1158015611009573d6000803e3d6000fd5b50506001805460ff60a81b19169055506115ca9050565b6001546001600160a01b031661106b576040805162461bcd60e51b815260206004820152601a60248201526000805160206120c2833981519152604482015290519081900360640190fd5b6110736117e8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156110e257600080fd5b505af11580156110f6573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561116957600080fd5b505af115801561117d573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156111f057600080fd5b505af1158015611204573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561127757600080fd5b505af115801561128b573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156112fe57600080fd5b505af1158015611312573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561138557600080fd5b505af1158015611399573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561140c57600080fd5b505af1158015611420573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561149357600080fd5b505af11580156114a7573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561151a57600080fd5b505af115801561152e573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663cfad57a2826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156115a157600080fd5b505af11580156115b5573d6000803e3d6000fd5b5050600180546001600160a81b031916905550505b50600080546001600160a01b0319169055565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461168e576040805162461bcd60e51b81526020600482015260096024820152683337b93134b23232b760b91b604482015290519081900360640190fd5b6001546001600160a01b0316156116da576040805162461bcd60e51b815260206004820152601a60248201526000805160206120c2833981519152604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166312d43a516040518163ffffffff1660e01b815260040160206040518083038186803b15801561173557600080fd5b505afa158015611749573d6000803e3d6000fd5b505050506040513d602081101561175f57600080fd5b5051600080546001600160a01b038084166001600160a01b03199283161783556001805460ff60a01b197f0000000000000000000000000000000000000000000000000000000000000000909316931692909217169055909150606090600a9061043b565b7f000000000000000000000000000000000000000000000000000000000000000081565b60015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f000000000000000000000000000000000000000000000000000000000000000090921691639cb7de4b9160448082019260009290919082900301818387803b15801561186957600080fd5b505af115801561187d573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b15801561190257600080fd5b505af1158015611916573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b15801561199b57600080fd5b505af11580156119af573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b158015611a3457600080fd5b505af1158015611a48573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b158015611acd57600080fd5b505af1158015611ae1573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b158015611b6657600080fd5b505af1158015611b7a573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b158015611bff57600080fd5b505af1158015611c13573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b158015611c9857600080fd5b505af1158015611cac573d6000803e3d6000fd5b505060015460408051639cb7de4b60e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f00000000000000000000000000000000000000000000000000000000000000009092169350639cb7de4b925060448082019260009290919082900301818387803b158015611d3157600080fd5b505af1158015611d45573d6000803e3d6000fd5b50506001546040805163cf456ae760e01b81526001600160a01b038084166004830152600160a01b90930460ff161515602482015290517f0000000000000000000000000000000000000000000000000000000000000000909216935063cf456ae7925060448082019260009290919082900301818387803b158015611dca57600080fd5b505af1158015611dde573d6000803e3d6000fd5b5050600154600160a01b900460ff161591506120bf9050577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f000000000000000000000000000000000000000000000000000000000000000060016040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015611e9057600080fd5b505af1158015611ea4573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f000000000000000000000000000000000000000000000000000000000000000060016040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015611f4257600080fd5b505af1158015611f56573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639cb7de4b7f000000000000000000000000000000000000000000000000000000000000000060016040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b158015611ff457600080fd5b505af1158015612008573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e44b75587f000000000000000000000000000000000000000000000000000000000000000060016040518363ffffffff1660e01b815260040180836001600160a01b03168152602001821515815260200192505050600060405180830381600087803b1580156120a657600080fd5b505af11580156120ba573d6000803e3d6000fd5b505050505b56fe696e76616c696420726577617264526f75746572546172676574000000000000a26469706673582212204de707c1090c1279016b17515b0921f234b8056e289c7cd64a3da31be8efd8e264736f6c634300060c0033";

export class BuybackMigrator__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    _admin: string,
    _stakedGmxTracker: string,
    _bonusGmxTracker: string,
    _extendedGmxTracker: string,
    _feeGmxTracker: string,
    _feeGlpTracker: string,
    _stakedGlpTracker: string,
    _gmxVester: string,
    _glpVester: string,
    _esGmx: string,
    _bnGmx: string,
    _oldRewardRouter: string,
    _newRewardRouter: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<BuybackMigrator> {
    return super.deploy(
      _admin,
      _stakedGmxTracker,
      _bonusGmxTracker,
      _extendedGmxTracker,
      _feeGmxTracker,
      _feeGlpTracker,
      _stakedGlpTracker,
      _gmxVester,
      _glpVester,
      _esGmx,
      _bnGmx,
      _oldRewardRouter,
      _newRewardRouter,
      overrides || {}
    ) as Promise<BuybackMigrator>;
  }
  getDeployTransaction(
    _admin: string,
    _stakedGmxTracker: string,
    _bonusGmxTracker: string,
    _extendedGmxTracker: string,
    _feeGmxTracker: string,
    _feeGlpTracker: string,
    _stakedGlpTracker: string,
    _gmxVester: string,
    _glpVester: string,
    _esGmx: string,
    _bnGmx: string,
    _oldRewardRouter: string,
    _newRewardRouter: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      _admin,
      _stakedGmxTracker,
      _bonusGmxTracker,
      _extendedGmxTracker,
      _feeGmxTracker,
      _feeGlpTracker,
      _stakedGlpTracker,
      _gmxVester,
      _glpVester,
      _esGmx,
      _bnGmx,
      _oldRewardRouter,
      _newRewardRouter,
      overrides || {}
    );
  }
  attach(address: string): BuybackMigrator {
    return super.attach(address) as BuybackMigrator;
  }
  connect(signer: Signer): BuybackMigrator__factory {
    return super.connect(signer) as BuybackMigrator__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): BuybackMigratorInterface {
    return new utils.Interface(_abi) as BuybackMigratorInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): BuybackMigrator {
    return new Contract(address, _abi, signerOrProvider) as BuybackMigrator;
  }
}
