/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer, utils, Contract, ContractFactory, Overrides } from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type {
  MockExternalHandler,
  MockExternalHandlerInterface,
} from "../MockExternalHandler";

const _abi = [
  {
    inputs: [
      {
        internalType: "address[]",
        name: "targets",
        type: "address[]",
      },
      {
        internalType: "bytes[]",
        name: "dataList",
        type: "bytes[]",
      },
      {
        internalType: "address[]",
        name: "refundTokens",
        type: "address[]",
      },
      {
        internalType: "address[]",
        name: "refundReceivers",
        type: "address[]",
      },
    ],
    name: "makeExternalCalls",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x608060405234801561001057600080fd5b5060016000556109a8806100256000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063d59922b014610030575b600080fd5b61004361003e366004610588565b610045565b005b600260005414156100715760405162461bcd60e51b8152600401610068906108c6565b60405180910390fd5b600260005582518451146100975760405162461bcd60e51b815260040161006890610810565b80518251146100b85760405162461bcd60e51b8152600401610068906107db565b60005b84518110156100fc576100f48582815181106100d357fe5b60200260200101518583815181106100e757fe5b60200260200101516101ee565b6001016100bb565b5060005b82518110156101e257600083828151811061011757fe5b602002602001015190506000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040161014f91906106d1565b60206040518083038186803b15801561016757600080fd5b505afa15801561017b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061019f919061069d565b905080156101d8576101d88484815181106101b657fe5b602002602001015182846001600160a01b031661029e9092919063ffffffff16565b5050600101610100565b50506001600055505050565b610200826001600160a01b03166102f4565b61021c5760405162461bcd60e51b815260040161006890610731565b6000826001600160a01b03168260405161023691906106b5565b6000604051808303816000865af19150503d8060008114610273576040519150601f19603f3d011682016040523d82523d6000602084013e610278565b606091505b50509050806102995760405162461bcd60e51b815260040161006890610767565b505050565b6102998363a9059cbb60e01b84846040516024016102bd9291906106e5565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526102fa565b3b151590565b606061034f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166103899092919063ffffffff16565b805190915015610299578080602001905181019061036d919061067d565b6102995760405162461bcd60e51b81526004016100689061087c565b606061039884846000856103a2565b90505b9392505050565b6060824710156103c45760405162461bcd60e51b815260040161006890610795565b6103cd856102f4565b6103e95760405162461bcd60e51b815260040161006890610845565b60006060866001600160a01b0316858760405161040691906106b5565b60006040518083038185875af1925050503d8060008114610443576040519150601f19603f3d011682016040523d82523d6000602084013e610448565b606091505b5091509150610458828286610463565b979650505050505050565b6060831561047257508161039b565b8251156104825782518084602001fd5b8160405162461bcd60e51b815260040161006891906106fe565b600082601f8301126104ac578081fd5b81356104bf6104ba82610923565b6108fd565b8181529150602080830190848101818402860182018710156104e057600080fd5b6000805b858110156105145782356001600160a01b0381168114610502578283fd5b855293830193918301916001016104e4565b50505050505092915050565b600082601f830112610530578081fd5b81356001600160401b03811115610545578182fd5b610558601f8201601f19166020016108fd565b915080825283602082850101111561056f57600080fd5b8060208401602084013760009082016020015292915050565b6000806000806080858703121561059d578384fd5b84356001600160401b03808211156105b3578586fd5b6105bf8883890161049c565b95506020915081870135818111156105d5578586fd5b8701601f810189136105e5578586fd5b80356105f36104ba82610923565b81815284810190838601895b84811015610628576106168e898435890101610520565b845292870192908701906001016105ff565b50909850505050604088013592505080821115610643578384fd5b61064f8883890161049c565b93506060870135915080821115610664578283fd5b506106718782880161049c565b91505092959194509250565b60006020828403121561068e578081fd5b8151801515811461039b578182fd5b6000602082840312156106ae578081fd5b5051919050565b600082516106c7818460208701610942565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b600060208252825180602084015261071d816040850160208701610942565b601f01601f19169190910160400192915050565b6020808252601c908201527b125b9d985b1a5908115e1d195c9b985b0810d85b1b0815185c99d95d60221b604082015260600190565b602080825260149082015273115e1d195c9b985b0810d85b1b0811985a5b195960621b604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b6020808252818101527f496e76616c69642045787465726e616c2052656365697665727320496e707574604082015260600190565b6020808252601b908201527a125b9d985b1a5908115e1d195c9b985b0810d85b1b08125b9c1d5d602a1b604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b6040518181016001600160401b038111828210171561091b57600080fd5b604052919050565b60006001600160401b03821115610938578081fd5b5060209081020190565b60005b8381101561095d578181015183820152602001610945565b8381111561096c576000848401525b5050505056fea26469706673582212201edd24db2c32d7ac6bf2ff0ef07659dbe8f60da8ba60c3f007e6aca7039523e964736f6c634300060c0033";

export class MockExternalHandler__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<MockExternalHandler> {
    return super.deploy(overrides || {}) as Promise<MockExternalHandler>;
  }
  getDeployTransaction(
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(overrides || {});
  }
  attach(address: string): MockExternalHandler {
    return super.attach(address) as MockExternalHandler;
  }
  connect(signer: Signer): MockExternalHandler__factory {
    return super.connect(signer) as MockExternalHandler__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): MockExternalHandlerInterface {
    return new utils.Interface(_abi) as MockExternalHandlerInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): MockExternalHandler {
    return new Contract(address, _abi, signerOrProvider) as MockExternalHandler;
  }
}