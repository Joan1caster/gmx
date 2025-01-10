/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type {
  MaliciousTraderTest,
  MaliciousTraderTestInterface,
} from "../MaliciousTraderTest";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_positionRouter",
        type: "address",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [],
    name: "Received",
    type: "event",
  },
  {
    inputs: [
      {
        internalType: "address[]",
        name: "_path",
        type: "address[]",
      },
      {
        internalType: "address",
        name: "_indexToken",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "_minOut",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "_sizeDelta",
        type: "uint256",
      },
      {
        internalType: "bool",
        name: "_isLong",
        type: "bool",
      },
      {
        internalType: "uint256",
        name: "_acceptablePrice",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "_executionFee",
        type: "uint256",
      },
      {
        internalType: "bytes32",
        name: "_referralCode",
        type: "bytes32",
      },
      {
        internalType: "address",
        name: "_callbackTarget",
        type: "address",
      },
    ],
    name: "createIncreasePositionETH",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [],
    name: "positionRouter",
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
    stateMutability: "payable",
    type: "receive",
  },
];

const _bytecode =
  "0x608060405234801561001057600080fd5b506040516109243803806109248339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556108bf806100656000396000f3fe60806040526004361061002d5760003560e01c80635b88e8c61461007f57806361ef161f1461015a5761007a565b3661007a576000805b620f424081101561004d5790810290600101610036565b506040517f544c765b33ca411cce832250371569244f765a17fcd217832be093f0fd5fa45b90600090a150005b600080fd5b610158600480360361012081101561009657600080fd5b810190602081018135600160201b8111156100b057600080fd5b8201836020820111156100c257600080fd5b803590602001918460208302840111600160201b831117156100e357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505081356001600160a01b03908116935060208301359260408101359250606081013515159160808201359160a08101359160c08201359160e001351661018b565b005b34801561016657600080fd5b5061016f610464565b604080516001600160a01b039092168252519081900360200190f35b6101b060405180606001604052806026815260200161081a602691398a518a8a610473565b6101d460405180606001604052806029815260200161086160299139878787610540565b6101f7604051806060016040528060218152602001610840602191398483610600565b6000606060008054906101000a90046001600160a01b03166001600160a01b0316348c8c8c8c8c8c8c8c8c60405160240180806020018a6001600160a01b031681526020018981526020018881526020018715158152602001868152602001858152602001848152602001836001600160a01b0316815260200182810382528b818151815260200191508051906020019060200280838360005b838110156102a9578181015183820152602001610291565b505050509050019a5050505050505050505050604051602081830303815290604052632dc4746360e11b6001600160e01b0319166020820180516001600160e01b0383818316178352505050506040518082805190602001908083835b602083106103255780518252601f199092019160209182019101610306565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114610387576040519150601f19603f3d011682016040523d82523d6000602084013e61038c565b606091505b50915091506103cc60405180604001604052806018815260200177737563636573733a20257320726561736f6e3a202725732760401b81525083836106c4565b80826104565760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561041b578181015183820152602001610403565b50505050905090810190601f1680156104485780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050505050505050505050565b6000546001600160a01b031681565b61053a848484846040516024018080602001858152602001846001600160a01b03168152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156104d55781810151838201526020016104bd565b50505050905090810190601f1680156105025780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b03166327827ee360e11b17905295506107db945050505050565b50505050565b61053a8484848460405160240180806020018581526020018415158152602001838152602001828103825286818151815260200191508051906020019080838360005b8381101561059b578181015183820152602001610583565b50505050905090810190601f1680156105c85780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b031663e41b6f6f60e01b17905295506107db945050505050565b6106bf8383836040516024018080602001848152602001836001600160a01b03168152602001828103825285818151815260200191508051906020019080838360005b8381101561065b578181015183820152602001610643565b50505050905090810190601f1680156106885780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b031663038fd88960e31b17905294506107db9350505050565b505050565b6106bf8383836040516024018080602001841515815260200180602001838103835286818151815260200191508051906020019080838360005b838110156107165781810151838201526020016106fe565b50505050905090810190601f1680156107435780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561077657818101518382015260200161075e565b50505050905090810190601f1680156107a35780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b031663e298f47d60e01b17905295506107db945050505050565b6107f2816107ea6107f5610816565b63ffffffff16565b50565b60006a636f6e736f6c652e6c6f679050600080835160208501845afa505050565b9056fe706174682e6c656e67746820257320696e646578546f6b656e202573206d696e4f7574202573657865637574696f6e4665652025732063616c6c6261636b54617267657420257373697a6544656c74612025732069734c6f6e672025732061636365707461626c655072696365202573a264697066735822122007d6d8622705bc64fa46f1ec98d6e60f912f11bfa56fc0c1923f327804d024f764736f6c634300060c0033";

export class MaliciousTraderTest__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    _positionRouter: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<MaliciousTraderTest> {
    return super.deploy(
      _positionRouter,
      overrides || {}
    ) as Promise<MaliciousTraderTest>;
  }
  getDeployTransaction(
    _positionRouter: string,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(_positionRouter, overrides || {});
  }
  attach(address: string): MaliciousTraderTest {
    return super.attach(address) as MaliciousTraderTest;
  }
  connect(signer: Signer): MaliciousTraderTest__factory {
    return super.connect(signer) as MaliciousTraderTest__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): MaliciousTraderTestInterface {
    return new utils.Interface(_abi) as MaliciousTraderTestInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): MaliciousTraderTest {
    return new Contract(address, _abi, signerOrProvider) as MaliciousTraderTest;
  }
}
