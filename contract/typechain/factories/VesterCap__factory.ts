/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import {
  Signer,
  utils,
  BigNumberish,
  Contract,
  ContractFactory,
  Overrides,
} from "ethers";
import { Provider, TransactionRequest } from "@ethersproject/providers";
import type { VesterCap, VesterCapInterface } from "../VesterCap";

const _abi = [
  {
    inputs: [
      {
        internalType: "address",
        name: "_gmxVester",
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
        name: "_bnGmx",
        type: "address",
      },
      {
        internalType: "address",
        name: "_esGmx",
        type: "address",
      },
      {
        internalType: "uint256",
        name: "_maxBoostBasisPoints",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "_bnGmxToEsGmxConversionDivisor",
        type: "uint256",
      },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    inputs: [],
    name: "BASIS_POINTS_DIVISOR",
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
    name: "bnGmxToEsGmxConversionDivisor",
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
    name: "isUpdateCompleted",
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
    name: "maxBoostBasisPoints",
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
        name: "_gov",
        type: "address",
      },
    ],
    name: "setGov",
    outputs: [],
    stateMutability: "nonpayable",
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
  {
    inputs: [
      {
        internalType: "address",
        name: "_account",
        type: "address",
      },
    ],
    name: "syncFeeGmxTrackerBalance",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "address[]",
        name: "_accounts",
        type: "address[]",
      },
    ],
    name: "updateBnGmxForAccounts",
    outputs: [],
    stateMutability: "nonpayable",
    type: "function",
  },
];

const _bytecode =
  "0x6101a060405234801561001157600080fd5b50604051611702380380611702833981810160405261012081101561003557600080fd5b5080516020820151604083015160608085015160808087015160a08089015160c0808b015160e0808d01516101009d8e01516001600081905580546001600160a01b031916331790556001600160601b03199d8b1b8e169098529a891b8c1690945297871b8a16905293851b88169052831b8616909652811b8416610120521b90911661014052610160526101805260805160601c60a05160601c60c05160601c60e05160601c6101005160601c6101205160601c6101405160601c61016051610180516115596101a9600039806102fc52806109e85250806104f55280610c895250806103685280610a1a52508061051952806107f752806109555280610eba52508061034452806105e152806106a052806107405280610cda5280610d965280610e385250806102d852806108225280610e025280610ef05250806102b452806108b852508061027b5280610be252508061032052806107635280610aab5280610b4e5280610db852506115596000f3fe608060405234801561001057600080fd5b50600436106100c55760003560e01c80630ce4018a146100ca578063126082cf146100ee57806312d43a51146101085780631fcd60e514610110578063204b4c6f1461011857806340d33ea91461012057806341d315b41461012857806351c3e3b4146101305780636a192a781461013857806382f9488714610140578063aeb3cdec146101e3578063cfad57a21461021d578063dddfa31414610243578063e102f5641461024b578063e21c36ea14610253575b600080fd5b6100d2610279565b604080516001600160a01b039092168252519081900360200190f35b6100f661029d565b60408051918252519081900360200190f35b6100d26102a3565b6100d26102b2565b6100d26102d6565b6100f66102fa565b6100d261031e565b6100d2610342565b6100d2610366565b6101e16004803603602081101561015657600080fd5b810190602081018135600160201b81111561017057600080fd5b82018360208201111561018257600080fd5b803590602001918460208302840111600160201b831117156101a357600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061038a945050505050565b005b610209600480360360208110156101f957600080fd5b50356001600160a01b0316610465565b604080519115158252519081900360200190f35b6101e16004803603602081101561023357600080fd5b50356001600160a01b031661047a565b6100f66104f3565b6100d2610517565b6101e16004803603602081101561026957600080fd5b50356001600160a01b031661053b565b7f000000000000000000000000000000000000000000000000000000000000000081565b61271081565b6001546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600260005414156103d0576040805162461bcd60e51b815260206004820152601f6024820152600080516020611493833981519152604482015290519081900360640190fd5b60026000556001546001600160a01b0316331461042c576040805162461bcd60e51b815260206004820152601560248201527423b7bb32b93730b136329d103337b93134b23232b760591b604482015290519081900360640190fd5b60005b815181101561045c5761045482828151811061044757fe5b6020026020010151610794565b60010161042f565b50506001600055565b60026020526000908152604090205460ff1681565b6001546001600160a01b031633146104d1576040805162461bcd60e51b815260206004820152601560248201527423b7bb32b93730b136329d103337b93134b23232b760591b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60026000541415610581576040805162461bcd60e51b815260206004820152601f6024820152600080516020611493833981519152604482015290519081900360640190fd5b60026000556001546001600160a01b031633146105dd576040805162461bcd60e51b815260206004820152601560248201527423b7bb32b93730b136329d103337b93134b23232b760591b604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166310c1c103836040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561064c57600080fd5b505afa158015610660573d6000803e3d6000fd5b505050506040513d602081101561067657600080fd5b5051604080516370a0823160e01b81526001600160a01b03858116600483015291519293506000927f0000000000000000000000000000000000000000000000000000000000000000909216916370a0823191602480820192602092909190829003018186803b1580156106e957600080fd5b505afa1580156106fd573d6000803e3d6000fd5b505050506040513d602081101561071357600080fd5b5051905081811161072557505061078c565b60006107318284610f61565b90506107886001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016857f000000000000000000000000000000000000000000000000000000000000000084610fac565b5050505b506001600055565b6001600160a01b03811660009081526002602052604090205460ff16156107ba57610f5e565b6001600160a01b038082166000818152600260209081526040808320805460ff191660011790558051637aeceb1f60e11b815260048101949094527f0000000000000000000000000000000000000000000000000000000000000000851660248501525191937f0000000000000000000000000000000000000000000000000000000000000000169263f5d9d63e9260448083019392829003018186803b15801561086457600080fd5b505afa158015610878573d6000803e3d6000fd5b505050506040513d602081101561088e57600080fd5b50516040805163402914f560e01b81526001600160a01b03858116600483015291519293506000927f00000000000000000000000000000000000000000000000000000000000000009092169163402914f591602480820192602092909190829003018186803b15801561090157600080fd5b505afa158015610915573d6000803e3d6000fd5b505050506040513d602081101561092b57600080fd5b5051604080516370a0823160e01b81526001600160a01b03868116600483015291519293506000927f0000000000000000000000000000000000000000000000000000000000000000909216916370a0823191602480820192602092909190829003018186803b15801561099e57600080fd5b505afa1580156109b2573d6000803e3d6000fd5b505050506040513d60208110156109c857600080fd5b5051905060006109e2826109dc868661100c565b9061100c565b905060007f00000000000000000000000000000000000000000000000000000000000000008281610a0f57fe5b0490508015610bde577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166340c10f1987836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015610a8f57600080fd5b505af1158015610aa3573d6000803e3d6000fd5b5050505060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a2545fa5886040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610b1657600080fd5b505afa158015610b2a573d6000803e3d6000fd5b505050506040513d6020811015610b4057600080fd5b505190506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166341f2272488610b7e848661100c565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015610bc457600080fd5b505af1158015610bd8573d6000803e3d6000fd5b50505050505b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166310c1c103886040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610c4d57600080fd5b505afa158015610c61573d6000803e3d6000fd5b505050506040513d6020811015610c7757600080fd5b505190506000610cb3612710610cad847f0000000000000000000000000000000000000000000000000000000000000000611064565b906110bd565b9050808711610cc85750505050505050610f5e565b6000610cd48883610f61565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a082318b6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610d4557600080fd5b505afa158015610d59573d6000803e3d6000fd5b505050506040513d6020811015610d6f57600080fd5b5051905080821115610de0576000610d878383610f61565b9050610dde6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167f00000000000000000000000000000000000000000000000000000000000000008d84610fac565b505b6040805163098bf59d60e01b81526001600160a01b038c8116600483018190527f00000000000000000000000000000000000000000000000000000000000000008216602484015260448301869052606483015291517f00000000000000000000000000000000000000000000000000000000000000009092169163098bf59d9160848082019260009290919082900301818387803b158015610e8257600080fd5b505af1158015610e96573d6000803e3d6000fd5b50506040805163098bf59d60e01b81526001600160a01b038e8116600483018190527f00000000000000000000000000000000000000000000000000000000000000008216602484015260448301889052606483015291517f0000000000000000000000000000000000000000000000000000000000000000909216935063098bf59d925060848082019260009290919082900301818387803b158015610f3c57600080fd5b505af1158015610f50573d6000803e3d6000fd5b505050505050505050505050505b50565b6000610fa383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506110fc565b90505b92915050565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052611006908590611198565b50505050565b600082820183811015610fa3576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b60008261107357506000610fa6565b8282028284828161108057fe5b0414610fa35760405162461bcd60e51b81526004018080602001828103825260218152602001806114d96021913960400191505060405180910390fd5b6000610fa383836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b81525061124e565b6000818484111561118b5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611150578181015183820152602001611138565b50505050905090810190601f16801561117d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50508183035b9392505050565b60606111ed826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166112b39092919063ffffffff16565b8051909150156112495780806020019051602081101561120c57600080fd5b50516112495760405162461bcd60e51b815260040180806020018281038252602a8152602001806114fa602a913960400191505060405180910390fd5b505050565b6000818361129d5760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315611150578181015183820152602001611138565b5060008385816112a957fe5b0495945050505050565b60606112c284846000856112ca565b949350505050565b60608247101561130b5760405162461bcd60e51b81526004018080602001828103825260268152602001806114b36026913960400191505060405180910390fd5b61131485611426565b611365576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b602083106113a45780518252601f199092019160209182019101611385565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611406576040519150601f19603f3d011682016040523d82523d6000602084013e61140b565b606091505b509150915061141b82828661142c565b979650505050505050565b3b151590565b6060831561143b575081611191565b82511561144b5782518084602001fd5b60405162461bcd60e51b815260206004820181815284516024840152845185939192839260440191908501908083836000831561115057818101518382015260200161113856fe5265656e7472616e637947756172643a207265656e7472616e742063616c6c00416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a26469706673582212200bde9989ffb092fae77e6a3f35e1c2e73aa95786bd9f773ab4892dd76f9570d964736f6c634300060c0033";

export class VesterCap__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer);
  }

  deploy(
    _gmxVester: string,
    _stakedGmxTracker: string,
    _bonusGmxTracker: string,
    _extendedGmxTracker: string,
    _feeGmxTracker: string,
    _bnGmx: string,
    _esGmx: string,
    _maxBoostBasisPoints: BigNumberish,
    _bnGmxToEsGmxConversionDivisor: BigNumberish,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<VesterCap> {
    return super.deploy(
      _gmxVester,
      _stakedGmxTracker,
      _bonusGmxTracker,
      _extendedGmxTracker,
      _feeGmxTracker,
      _bnGmx,
      _esGmx,
      _maxBoostBasisPoints,
      _bnGmxToEsGmxConversionDivisor,
      overrides || {}
    ) as Promise<VesterCap>;
  }
  getDeployTransaction(
    _gmxVester: string,
    _stakedGmxTracker: string,
    _bonusGmxTracker: string,
    _extendedGmxTracker: string,
    _feeGmxTracker: string,
    _bnGmx: string,
    _esGmx: string,
    _maxBoostBasisPoints: BigNumberish,
    _bnGmxToEsGmxConversionDivisor: BigNumberish,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): TransactionRequest {
    return super.getDeployTransaction(
      _gmxVester,
      _stakedGmxTracker,
      _bonusGmxTracker,
      _extendedGmxTracker,
      _feeGmxTracker,
      _bnGmx,
      _esGmx,
      _maxBoostBasisPoints,
      _bnGmxToEsGmxConversionDivisor,
      overrides || {}
    );
  }
  attach(address: string): VesterCap {
    return super.attach(address) as VesterCap;
  }
  connect(signer: Signer): VesterCap__factory {
    return super.connect(signer) as VesterCap__factory;
  }
  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): VesterCapInterface {
    return new utils.Interface(_abi) as VesterCapInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): VesterCap {
    return new Contract(address, _abi, signerOrProvider) as VesterCap;
  }
}
