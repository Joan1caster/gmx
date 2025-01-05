/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type {
  MockGovRequester,
  MockGovRequesterInterface,
} from "../MockGovRequester";

const _abi = [
  {
    inputs: [],
    name: "afterGovGranted",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address",
        name: "timelock",
        type: "address",
      },
      {
        internalType: "address[]",
        name: "targets",
        type: "address[]",
      },
    ],
    name: "migrate",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x608060405234801561001057600080fd5b506101cb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063047701e41461003b578063c7cfeb59146100ee575b600080fd5b6100ec6004803603604081101561005157600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561007b57600080fd5b82018360208201111561008d57600080fd5b803590602001918460208302840111600160201b831117156100ae57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506100f6945050505050565b005b6100ec610193565b60405163de8616ff60e01b81526020600482018181528351602484015283516001600160a01b0386169363de8616ff93869392839260440191808601910280838360005b8381101561015257818101518382015260200161013a565b5050505090500192505050600060405180830381600087803b15801561017757600080fd5b505af115801561018b573d6000803e3d6000fd5b505050505050565b56fea2646970667358221220094f3e5122a61aba05bb36e38554aaa0af9caa8bfe2f5201728c7706b31ce21364736f6c634300060c0033";

export class MockGovRequester__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<MockGovRequester> {
    return super.deploy(overrides || {}) as Promise<MockGovRequester>;
  }
  getDeployTransaction(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  attach(address: string): MockGovRequester {
    return super.attach(address) as MockGovRequester;
  }
  connect(signer: Signer): MockGovRequester__factory {
    return super.connect(signer) as MockGovRequester__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): MockGovRequesterInterface {
    return new utils.Interface(_abi) as MockGovRequesterInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): MockGovRequester {
    return new Contract(address, _abi, signerOrProvider) as MockGovRequester;
  }
}