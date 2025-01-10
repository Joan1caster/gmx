/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type { StakedGlp, StakedGlpInterface } from "../StakedGlp";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_glp",
        type: "address",
      },
      {
        internalType: "contract IGlpManager",
        name: "_glpManager",
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
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: "address",
        name: "owner",
        type: "address",
      },
      {
        indexed: true,
        internalType: "address",
        name: "spender",
        type: "address",
      },
      {
        indexed: false,
        internalType: "uint256",
        name: "value",
        type: "uint256",
      },
    ],
    name: "Approval",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_owner",
        type: "address",
      },
      {
        internalType: "address",
        name: "_spender",
        type: "address",
      },
    ],
    name: "allowance",
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
    inputs: [
      {
        internalType: "address",
        name: "",
        type: "address",
      },
      {
        internalType: "address",
        name: "",
        type: "address",
      },
    ],
    name: "allowances",
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
    inputs: [
      {
        internalType: "address",
        name: "_spender",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "_amount",
        type: "uint256",
      },
    ],
    name: "approve",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_account",
        type: "address",
      },
    ],
    name: "balanceOf",
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
    name: "decimals",
    outputs: [
      {
        internalType: "uint8",
        name: "",
        type: "uint8",
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
    name: "glp",
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
    name: "glpManager",
    outputs: [
      {
        internalType: "contract IGlpManager",
        name: "",
        type: "address",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "name",
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
    name: "symbol",
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
    inputs: [],
    name: "totalSupply",
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
    inputs: [
      {
        internalType: "address",
        name: "_recipient",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "_amount",
        type: "uint256",
      },
    ],
    name: "transfer",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "_sender",
        type: "address",
      },
      {
        internalType: "address",
        name: "_recipient",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "_amount",
        type: "uint256",
      },
    ],
    name: "transferFrom",
    outputs: [
      {
        internalType: "bool",
        name: "",
        type: "bool",
      },
    ],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x608060405234801561001057600080fd5b50604051610cb7380380610cb78339818101604052608081101561003357600080fd5b50805160208201516040830151606090930151600080546001600160a01b039485166001600160a01b0319918216179091556001805493851693821693909317909255600280549484169483169490941790935560038054929093169116179055610c14806100a36000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c806306fdde03146100bf578063095ea7b31461013c57806318160ddd1461017c57806323b872dd14610196578063313ce567146101cc57806355b6ed5c146101ea57806370a082311461021857806378a207ee1461023e57806395d89b4114610262578063a9059cbb1461026a578063af394d0014610296578063dd62ed3e1461029e578063e1c363b7146102cc578063fa6db1bc146102d4575b600080fd5b6100c76102dc565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101015781810151838201526020016100e9565b50505050905090810190601f16801561012e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101686004803603604081101561015257600080fd5b506001600160a01b038135169060200135610301565b604080519115158252519081900360200190f35b610184610317565b60408051918252519081900360200190f35b610168600480360360608110156101ac57600080fd5b506001600160a01b0381358116916020810135909116906040013561038d565b6101d46103fb565b6040805160ff9092168252519081900360200190f35b6101846004803603604081101561020057600080fd5b506001600160a01b0381358116916020013516610400565b6101846004803603602081101561022e57600080fd5b50356001600160a01b031661041d565b6102466104ab565b604080516001600160a01b039092168252519081900360200190f35b6100c76104ba565b6101686004803603604081101561028057600080fd5b506001600160a01b0381351690602001356104da565b6102466104e7565b610184600480360360408110156102b457600080fd5b506001600160a01b03813581169160200135166104f6565b610246610521565b610246610530565b6040518060400160405280600981526020016805374616b6564476c760bc1b81525081565b600061030e33848461053f565b50600192915050565b600254604080516318160ddd60e01b815290516000926001600160a01b0316916318160ddd916004808301926020929190829003018186803b15801561035c57600080fd5b505afa158015610370573d6000803e3d6000fd5b505050506040513d602081101561038657600080fd5b5051905090565b6000806103d8836040518060600160405280602c8152602001610b88602c91396001600160a01b0388166000908152600460209081526040808320338452909152902054919061062b565b90506103e585338361053f565b6103f08585856106c2565b506001949350505050565b601281565b600460209081526000928352604080842090915290825290205481565b6003546000805460408051637aeceb1f60e11b81526001600160a01b038681166004830152928316602482015290519293919091169163f5d9d63e91604480820192602092909190829003018186803b15801561047957600080fd5b505afa15801561048d573d6000803e3d6000fd5b505050506040513d60208110156104a357600080fd5b505192915050565b6000546001600160a01b031681565b60405180604001604052806004815260200163073474c560e41b81525081565b600061030e3384846106c2565b6002546001600160a01b031681565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205490565b6003546001600160a01b031681565b6001546001600160a01b031681565b6001600160a01b0383166105845760405162461bcd60e51b8152600401808060200182810382526028815260200180610b136028913960400191505060405180910390fd5b6001600160a01b0382166105c95760405162461bcd60e51b8152600401808060200182810382526026815260200180610b3b6026913960400191505060405180910390fd5b6001600160a01b03808416600081815260046020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b600081848411156106ba5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561067f578181015183820152602001610667565b50505050905090810190601f1680156106ac5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b0383166107075760405162461bcd60e51b8152600401808060200182810382526029815260200180610aea6029913960400191505060405180910390fd5b6001600160a01b03821661074c5760405162461bcd60e51b8152600401808060200182810382526027815260200180610b616027913960400191505060405180910390fd5b4261084d600160009054906101000a90046001600160a01b03166001600160a01b031663352693156040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156107a057600080fd5b505af11580156107b4573d6000803e3d6000fd5b505050506040513d60208110156107ca57600080fd5b505160015460408051638b770e1160e01b81526001600160a01b03898116600483015291519190921691638b770e119160248083019260209291908290030181600087803b15801561081b57600080fd5b505af115801561082f573d6000803e3d6000fd5b505050506040513d602081101561084557600080fd5b505190610a8a565b111561088a5760405162461bcd60e51b815260040180806020018281038252602b815260200180610bb4602b913960400191505060405180910390fd5b6002546003546040805163098bf59d60e01b81526001600160a01b038781166004830181905293811660248301526044820186905260648201939093529051919092169163098bf59d91608480830192600092919082900301818387803b1580156108f457600080fd5b505af1158015610908573d6000803e3d6000fd5b5050600354600080546040805163098bf59d60e01b81526001600160a01b038a8116600483018190529381166024830152604482018990526064820193909352905191909316945063098bf59d935060848084019382900301818387803b15801561097257600080fd5b505af1158015610986573d6000803e3d6000fd5b50506003546000805460408051631e42d69b60e21b81526001600160a01b038a811660048301528981166024830152928316604482015260648101889052905191909316945063790b5a6c935060848084019382900301818387803b1580156109ee57600080fd5b505af1158015610a02573d6000803e3d6000fd5b505060025460035460408051631e42d69b60e21b81526001600160a01b03888116600483018190526024830152928316604482015260648101879052905191909216935063790b5a6c9250608480830192600092919082900301818387803b158015610a6d57600080fd5b505af1158015610a81573d6000803e3d6000fd5b50505050505050565b600082820183811015610ae2576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b939250505056fe5374616b6564476c703a207472616e736665722066726f6d20746865207a65726f20616464726573735374616b6564476c703a20617070726f76652066726f6d20746865207a65726f20616464726573735374616b6564476c703a20617070726f766520746f20746865207a65726f20616464726573735374616b6564476c703a207472616e7366657220746f20746865207a65726f20616464726573735374616b6564476c703a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63655374616b6564476c703a20636f6f6c646f776e206475726174696f6e206e6f742079657420706173736564a26469706673582212206294a5c23e73a4c8cb95e91bead74ad0fcb264ffac1645675316306d9d1ae6e864736f6c634300060c0033";

export class StakedGlp__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    _glp: string,
    _glpManager: string,
    _stakedGlpTracker: string,
    _feeGlpTracker: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<StakedGlp> {
    return super.deploy(
      _glp,
      _glpManager,
      _stakedGlpTracker,
      _feeGlpTracker,
      overrides || {}
    ) as Promise<StakedGlp>;
  }
  getDeployTransaction(
    _glp: string,
    _glpManager: string,
    _stakedGlpTracker: string,
    _feeGlpTracker: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      _glp,
      _glpManager,
      _stakedGlpTracker,
      _feeGlpTracker,
      overrides || {}
    );
  }
  attach(address: string): StakedGlp {
    return super.attach(address) as StakedGlp;
  }
  connect(signer: Signer): StakedGlp__factory {
    return super.connect(signer) as StakedGlp__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): StakedGlpInterface {
    return new utils.Interface(_abi) as StakedGlpInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): StakedGlp {
    return new Contract(address, _abi, signerOrProvider) as StakedGlp;
  }
}
