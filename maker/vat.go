// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maker

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VatABI is the input ABI used to generate the binding from.
const VatABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg3\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Line\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"rate\",\"type\":\"int256\"}],\"name\":\"fold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"fork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"grab\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"Art\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"wad\",\"type\":\"int256\"}],\"name\":\"slip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"suck\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"urns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VatFuncSigs maps the 4-byte function signature to its string representation.
var VatFuncSigs = map[string]string{
	"babe8a3f": "Line()",
	"69245009": "cage()",
	"4538c4eb": "can(address,address)",
	"6c25b346": "dai(address)",
	"0dca59c1": "debt()",
	"9c52a7f1": "deny(address)",
	"1a0b287e": "file(bytes32,bytes32,uint256)",
	"29ae8114": "file(bytes32,uint256)",
	"6111be2e": "flux(bytes32,address,address,uint256)",
	"b65337df": "fold(bytes32,address,int256)",
	"870c616d": "fork(bytes32,address,address,int256,int256)",
	"76088703": "frob(bytes32,address,address,address,int256,int256)",
	"214414d5": "gem(bytes32,address)",
	"7bab3f40": "grab(bytes32,address,address,address,int256,int256)",
	"f37ac61c": "heal(uint256)",
	"a3b22fc4": "hope(address)",
	"d9638d36": "ilks(bytes32)",
	"3b663195": "init(bytes32)",
	"957aa58c": "live()",
	"bb35783b": "move(address,address,uint256)",
	"dc4d20fa": "nope(address)",
	"65fae35e": "rely(address)",
	"f059212a": "sin(address)",
	"7cdd3fde": "slip(bytes32,address,int256)",
	"f24e23eb": "suck(address,address,uint256)",
	"2424be5c": "urns(bytes32,address)",
	"2d61a355": "vice()",
	"bf353dbb": "wards(address)",
}

// VatBin is the compiled bytecode used for deploying new contracts.
var VatBin = "0x608060405234801561001057600080fd5b50336000908152602081905260409020600190819055600a55612005806100386000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80637cdd3fde116100f9578063bb35783b11610097578063dc4d20fa11610071578063dc4d20fa146105c1578063f059212a146105e7578063f24e23eb1461060d578063f37ac61c14610643576101c4565b8063bb35783b14610516578063bf353dbb1461054c578063d9638d3614610572576101c4565b80639c52a7f1116100d35780639c52a7f114610490578063a3b22fc4146104b6578063b65337df146104dc578063babe8a3f1461050e576101c4565b80637cdd3fde14610414578063870c616d14610446578063957aa58c14610488576101c4565b80634538c4eb11610166578063692450091161014057806369245009146103565780636c25b3461461035e57806376088703146103845780637bab3f40146103cc576101c4565b80634538c4eb146102c65780636111be2e146102f457806365fae35e14610330576101c4565b80632424be5c116101a25780632424be5c1461023a57806329ae81141461027e5780632d61a355146102a15780633b663195146102a9576101c4565b80630dca59c1146101c95780631a0b287e146101e3578063214414d51461020e575b600080fd5b6101d1610660565b60405190815260200160405180910390f35b61020c600480360360608110156101f957600080fd5b5080359060208101359060400135610666565b005b6101d16004803603604081101561022457600080fd5b50803590602001356001600160a01b03166107fe565b6102666004803603604081101561025057600080fd5b50803590602001356001600160a01b0316610820565b60405191825260208201526040908101905180910390f35b61020c6004803603604081101561029457600080fd5b508035906020013561084a565b6101d1610931565b61020c600480360360208110156102bf57600080fd5b5035610937565b6101d1600480360360408110156102dc57600080fd5b506001600160a01b0381358116916020013516610a3e565b61020c6004803603608081101561030a57600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610a60565b61020c6004803603602081101561034657600080fd5b50356001600160a01b0316610b9f565b61020c610c8d565b6101d16004803603602081101561037457600080fd5b50356001600160a01b0316610d1e565b61020c600480360360c081101561039a57600080fd5b508035906001600160a01b03602082013581169160408101358216916060820135169060808101359060a00135610d32565b61020c600480360360c08110156103e257600080fd5b508035906001600160a01b03602082013581169160408101358216916060820135169060808101359060a00135611255565b61020c6004803603606081101561042a57600080fd5b508035906001600160a01b036020820135169060400135611411565b61020c600480360360a081101561045c57600080fd5b508035906001600160a01b036020820135811691604081013590911690606081013590608001356114fa565b6101d16117d4565b61020c600480360360208110156104a657600080fd5b50356001600160a01b03166117da565b61020c600480360360208110156104cc57600080fd5b50356001600160a01b03166118c5565b61020c600480360360608110156104f257600080fd5b508035906001600160a01b03602082013516906040013561192d565b6101d1611a84565b61020c6004803603606081101561052c57600080fd5b506001600160a01b03813581169160208101359091169060400135611a8a565b6101d16004803603602081101561056257600080fd5b50356001600160a01b0316611b8c565b61058f6004803603602081101561058857600080fd5b5035611ba0565b604051808681526020018581526020018481526020018381526020018281526020019550505050505060405180910390f35b61020c600480360360208110156105d757600080fd5b50356001600160a01b0316611bcf565b6101d1600480360360208110156105fd57600080fd5b50356001600160a01b0316611c34565b61020c6004803603606081101561062357600080fd5b506001600160a01b03813581169160208101359091169060400135611c48565b61020c6004803603602081101561065957600080fd5b5035611d6f565b60075481565b336000908152602081905260409020546001146106b75760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b600a546001146106fc5760405162461bcd60e51b815260206004820152600c60248201526b5661742f6e6f742d6c69766560a01b604482015260640160405180910390fd5b81631cdc1bdd60e21b14156107245760008381526002602052819060409020600201556107c0565b81636c696e6560e01b141561074c5760008381526002602052819060409020600301556107c0565b8163191d5cdd60e21b14156107745760008381526002602052819060409020600401556107c0565b60405162461bcd60e51b815260206004820152601b60248201527f5661742f66696c652d756e7265636f676e697a65642d706172616d0000000000604482015260640160405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450505050565b6004602052816000526040600020602052806000526040600020549150829050565b60036020528160005260406000206020528060005260406000208054600190910154909250905082565b3360009081526020819052604090205460011461089b5760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b600a546001146108e05760405162461bcd60e51b815260206004820152600c60248201526b5661742f6e6f742d6c69766560a01b604482015260640160405180910390fd5b81634c696e6560e01b14156107745760098190555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a4505050565b60085481565b336000908152602081905260409020546001146109885760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b600081815260026020526040902060010154156109e25760405162461bcd60e51b815260206004820152601460248201527315985d0bda5b1acb585b1c9958591e4b5a5b9a5d60621b604482015260640160405180910390fd5b600081815260026020526b033b2e3c9fd0803ce80000009060409020600101555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a45050565b6001602052816000526040600020602052806000526040600020549150829050565b610a6a8333611e3b565b610aac5760405162461bcd60e51b815260206004820152600f60248201526e15985d0bdb9bdd0b585b1b1bddd959608a1b604482015260640160405180910390fd5b60008481526004602052610ade90604090206001600160a01b0385166000908152602091909152604090205482611e88565b60008581526004602052604090206001600160a01b0385166000908152602091909152604090205560008481526004602052610b3890604090206001600160a01b0384166000908152602091909152604090205482611e98565b60008581526004602052604090206001600160a01b038416600090815260209190915260409020555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a45050505050565b33600090815260208190526040902054600114610bf05760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b600a54600114610c355760405162461bcd60e51b815260206004820152600c60248201526b5661742f6e6f742d6c69766560a01b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260019060409020555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a45050565b33600090815260208190526040902054600114610cde5760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b6000600a555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450565b600560205280600052604060002054905081565b600a54600114610d775760405162461bcd60e51b815260206004820152600c60248201526b5661742f6e6f742d6c69766560a01b604482015260640160405180910390fd5b610d7f611f6a565b60008781526003602052604090206001600160a01b038716600090815260209190915260409020604051604080820190528154815260019091015460208201529050610dc9611f81565b600088815260026020526040902060405160a08101604090815282548252600183015460208301908152600284015491830191909152600383015460608301526004909201546080820152915051610e5a5760405162461bcd60e51b815260206004820152601060248201526f15985d0bda5b1acb5b9bdd0b5a5b9a5d60821b604482015260640160405180910390fd5b610e65825185611ea8565b8252610e75602083015184611ea8565b6020830152610e85815184611ea8565b81526000610e97602083015185611edd565b90506000610ead83602001518560200151611f0b565b9050610ebb60075483611ea8565b600755610ef16000861315610eec6060860151610edd87518860200151611f0b565b11156009546007541115611f2e565b611f32565b610f385760405162461bcd60e51b815260206004820152601460248201527315985d0bd8d95a5b1a5b99cb595e18d95959195960621b604482015260640160405180910390fd5b610f64610f4d60008713156000891215611f2e565b610f5c86518660400151611f0b565b831115611f32565b610fa35760405162461bcd60e51b815260206004820152600c60248201526b5661742f6e6f742d7361666560a01b604482015260640160405180910390fd5b610fc2610fb860008713156000891215611f2e565b610eec8b33611e3b565b6110065760405162461bcd60e51b81526020600482015260116024820152705661742f6e6f742d616c6c6f7765642d7560781b604482015260640160405180910390fd5b6110186000871315610eec8a33611e3b565b61105c5760405162461bcd60e51b81526020600482015260116024820152702b30ba17b737ba16b0b63637bbb2b216bb60791b604482015260640160405180910390fd5b61106e6000861215610eec8933611e3b565b6110b25760405162461bcd60e51b81526020600482015260116024820152705661742f6e6f742d616c6c6f7765642d7760781b604482015260640160405180910390fd5b6110c88460200151156080850151831015611f32565b6111035760405162461bcd60e51b815260206004820152600860248201526715985d0bd91d5cdd60c21b604482015260640160405180910390fd5b60008a8152600460205261113590604090206001600160a01b038a166000908152602091909152604090205487611f36565b60008b81526004602052604090206001600160a01b038a16600090815260209190915260409020556001600160a01b0387166000908152600560205261118090604090205483611ea8565b6001600160a01b03881660009081526005602052604090205560008a815260036020528490604090206001600160a01b038b166000908152602091909152604090208151815560208201516001909101555060008a8152600260205283906040902081518155602082015181600101556040820151816002015560608201518160030155608082015160049091015550505050505961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450505050505050565b336000908152602081905260409020546001146112a65760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b60008681526003602052604081206001600160a01b038716600090815260209190915260409020600088815260026020529091506040812090506112ee826000015485611ea8565b825560018201546112ff9084611ea8565b600183015580546113109084611ea8565b815560018101546000906113249085611edd565b60008a8152600460205290915061135990604090206001600160a01b0389166000908152602091909152604090205486611f36565b60008a81526004602052604090206001600160a01b038916600090815260209190915260409020556001600160a01b038616600090815260066020526113a490604090205482611f36565b6001600160a01b0387166000908152600660205260409020556008546113ca9082611f36565b6008555050505961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450505050505050565b336000908152602081905260409020546001146114625760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b6000838152600460205261149490604090206001600160a01b0384166000908152602091909152604090205482611ea8565b60008481526004602052604090206001600160a01b038416600090815260209190915260409020555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450505050565b60008581526003602052604081206001600160a01b03861660009081526020919091526040902060008781526003602052909150604081206001600160a01b0386166000908152602091909152604090206000888152600260205290915060408120905061156c836000015486611f36565b8355600183015461157d9085611f36565b6001840155815461158e9086611ea8565b8255600182015461159f9085611ea8565b826001018190555060006115bb84600101548360010154611f0b565b905060006115d184600101548460010154611f0b565b90506115ef6115e08a33611e3b565b6115ea8a33611e3b565b611f2e565b6116315760405162461bcd60e51b815260206004820152600f60248201526e15985d0bdb9bdd0b585b1b1bddd959608a1b604482015260640160405180910390fd5b61164385600001548460020154611f0b565b8211156116895760405162461bcd60e51b815260206004820152601060248201526f5661742f6e6f742d736166652d73726360801b604482015260640160405180910390fd5b61169b84600001548460020154611f0b565b8111156116e15760405162461bcd60e51b815260206004820152601060248201526f15985d0bdb9bdd0b5cd859994b591cdd60821b604482015260640160405180910390fd5b6116f983600401548310158660010154600014611f32565b6117385760405162461bcd60e51b815260206004820152600c60248201526b5661742f647573742d73726360a01b604482015260640160405180910390fd5b61175083600401548210158560010154600014611f32565b61178f5760405162461bcd60e51b815260206004820152600c60248201526b15985d0bd91d5cdd0b591cdd60a21b604482015260640160405180910390fd5b50505050505961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a4505050505050565b600a5481565b3360009081526020819052604090205460011461182b5760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b600a546001146118705760405162461bcd60e51b815260206004820152600c60248201526b5661742f6e6f742d6c69766560a01b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260408120555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a45050565b3360009081526001602081905290604090206001600160a01b038316600090815260209190915260409020555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a45050565b3360009081526020819052604090205460011461197e5760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b600a546001146119c35760405162461bcd60e51b815260206004820152600c60248201526b5661742f6e6f742d6c69766560a01b604482015260640160405180910390fd5b600083815260026020526040812090506119e1816001015483611ea8565b600182015580546000906119f59084611edd565b6001600160a01b03851660009081526005602052909150611a1b90604090205482611ea8565b6001600160a01b038516600090815260056020526040902055600754611a419082611ea8565b60075550505961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450505050565b60095481565b611a948333611e3b565b611ad65760405162461bcd60e51b815260206004820152600f60248201526e15985d0bdb9bdd0b585b1b1bddd959608a1b604482015260640160405180910390fd5b6001600160a01b03831660009081526005602052611af990604090205482611e88565b6001600160a01b0384166000908152600560205260409020556001600160a01b03821660009081526005602052611b3590604090205482611e98565b6001600160a01b0383166000908152600560205260409020555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450505050565b600060205280600052604060002054905081565b600260205280600052604060002080546001820154600283015460038401546004909401549294509092909185565b3360009081526001602052604081206001600160a01b038316600090815260209190915260409020555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a45050565b600660205280600052604060002054905081565b33600090815260208190526040902054600114611c995760405162461bcd60e51b81526020600482015260126024820152600080516020611fb1833981519152604482015260640160405180910390fd5b6001600160a01b03831660009081526006602052611cbc90604090205482611e98565b6001600160a01b0384166000908152600660205260409020556001600160a01b03821660009081526005602052611cf890604090205482611e98565b6001600160a01b038316600090815260056020526040902055600854611d1e9082611e98565b600855600754611d2e9082611e98565b6007555961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a450505050565b3360008181526006602052611d8990604090205483611e88565b6001600160a01b0382166000908152600660205260409020556001600160a01b03811660009081526005602052611dc590604090205483611e88565b6001600160a01b038216600090815260056020526040902055600854611deb9083611e88565b600855600754611dfb9083611e88565b600755505961012081016040526020815260e0602082015260e0600060408301376044356024356004356001600160e01b03196000351661012085a45050565b6001600160a01b038281166000818152600160205291611e7f9190841614604083206001600160a01b03851660009081526020919091526040902054600114611f32565b90505b92915050565b80820382811115611e8257600080fd5b80820182811015611e8257600080fd5b818101600082121580611ebb5750828111155b611ec457600080fd5b600082131580611ed45750828110155b611e8257600080fd5b8181026000831215611eee57600080fd5b811580611ed4575082828281611f0057fe5b0514611e8257600080fd5b6000811580611ed457505080820282828281611f2357fe5b0414611e8257600080fd5b1690565b1790565b808203600082131580611f495750828111155b611f5257600080fd5b600082121580611ed4575082811015611e8257600080fd5b604051604080820190526000808252602082015290565b6040518060a001604052806000815260200160008152602001600081526020016000815260200160008152509056fe5661742f6e6f742d617574686f72697a65640000000000000000000000000000a265627a7a7231582000423ac5381a6e9b388cad15a2b8aeb63b3bfc2930e4f833967a698c370fd48164736f6c63430005100032"

// DeployVat deploys a new Ethereum contract, binding an instance of Vat to it.
func DeployVat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vat, error) {
	parsed, err := abi.JSON(strings.NewReader(VatABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vat{VatCaller: VatCaller{contract: contract}, VatTransactor: VatTransactor{contract: contract}, VatFilterer: VatFilterer{contract: contract}}, nil
}

// Vat is an auto generated Go binding around an Ethereum contract.
type Vat struct {
	VatCaller     // Read-only binding to the contract
	VatTransactor // Write-only binding to the contract
	VatFilterer   // Log filterer for contract events
}

// VatCaller is an auto generated read-only Go binding around an Ethereum contract.
type VatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VatSession struct {
	Contract     *Vat              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VatCallerSession struct {
	Contract *VatCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VatTransactorSession struct {
	Contract     *VatTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatRaw is an auto generated low-level Go binding around an Ethereum contract.
type VatRaw struct {
	Contract *Vat // Generic contract binding to access the raw methods on
}

// VatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VatCallerRaw struct {
	Contract *VatCaller // Generic read-only contract binding to access the raw methods on
}

// VatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VatTransactorRaw struct {
	Contract *VatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVat creates a new instance of Vat, bound to a specific deployed contract.
func NewVat(address common.Address, backend bind.ContractBackend) (*Vat, error) {
	contract, err := bindVat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vat{VatCaller: VatCaller{contract: contract}, VatTransactor: VatTransactor{contract: contract}, VatFilterer: VatFilterer{contract: contract}}, nil
}

// NewVatCaller creates a new read-only instance of Vat, bound to a specific deployed contract.
func NewVatCaller(address common.Address, caller bind.ContractCaller) (*VatCaller, error) {
	contract, err := bindVat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VatCaller{contract: contract}, nil
}

// NewVatTransactor creates a new write-only instance of Vat, bound to a specific deployed contract.
func NewVatTransactor(address common.Address, transactor bind.ContractTransactor) (*VatTransactor, error) {
	contract, err := bindVat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VatTransactor{contract: contract}, nil
}

// NewVatFilterer creates a new log filterer instance of Vat, bound to a specific deployed contract.
func NewVatFilterer(address common.Address, filterer bind.ContractFilterer) (*VatFilterer, error) {
	contract, err := bindVat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VatFilterer{contract: contract}, nil
}

// bindVat binds a generic wrapper to an already deployed contract.
func bindVat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vat *VatRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vat.Contract.VatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vat *VatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.Contract.VatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vat *VatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vat.Contract.VatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vat *VatCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vat *VatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vat *VatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vat.Contract.contract.Transact(opts, method, params...)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() constant returns(uint256)
func (_Vat *VatCaller) Line(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "Line")
	return *ret0, err
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() constant returns(uint256)
func (_Vat *VatSession) Line() (*big.Int, error) {
	return _Vat.Contract.Line(&_Vat.CallOpts)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() constant returns(uint256)
func (_Vat *VatCallerSession) Line() (*big.Int, error) {
	return _Vat.Contract.Line(&_Vat.CallOpts)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) constant returns(uint256)
func (_Vat *VatCaller) Can(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "can", arg0, arg1)
	return *ret0, err
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) constant returns(uint256)
func (_Vat *VatSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Can(&_Vat.CallOpts, arg0, arg1)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) constant returns(uint256)
func (_Vat *VatCallerSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Can(&_Vat.CallOpts, arg0, arg1)
}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) constant returns(uint256)
func (_Vat *VatCaller) Dai(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "dai", arg0)
	return *ret0, err
}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) constant returns(uint256)
func (_Vat *VatSession) Dai(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Dai(&_Vat.CallOpts, arg0)
}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) constant returns(uint256)
func (_Vat *VatCallerSession) Dai(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Dai(&_Vat.CallOpts, arg0)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() constant returns(uint256)
func (_Vat *VatCaller) Debt(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "debt")
	return *ret0, err
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() constant returns(uint256)
func (_Vat *VatSession) Debt() (*big.Int, error) {
	return _Vat.Contract.Debt(&_Vat.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() constant returns(uint256)
func (_Vat *VatCallerSession) Debt() (*big.Int, error) {
	return _Vat.Contract.Debt(&_Vat.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) constant returns(uint256)
func (_Vat *VatCaller) Gem(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "gem", arg0, arg1)
	return *ret0, err
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) constant returns(uint256)
func (_Vat *VatSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Gem(&_Vat.CallOpts, arg0, arg1)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) constant returns(uint256)
func (_Vat *VatCallerSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Gem(&_Vat.CallOpts, arg0, arg1)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	ret := new(struct {
		Art  *big.Int
		Rate *big.Int
		Spot *big.Int
		Line *big.Int
		Dust *big.Int
	})
	out := ret
	err := _Vat.contract.Call(opts, out, "ilks", arg0)
	return *ret, err
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vat.Contract.Ilks(&_Vat.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatCallerSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vat.Contract.Ilks(&_Vat.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Vat *VatCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "live")
	return *ret0, err
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Vat *VatSession) Live() (*big.Int, error) {
	return _Vat.Contract.Live(&_Vat.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Vat *VatCallerSession) Live() (*big.Int, error) {
	return _Vat.Contract.Live(&_Vat.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) constant returns(uint256)
func (_Vat *VatCaller) Sin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "sin", arg0)
	return *ret0, err
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) constant returns(uint256)
func (_Vat *VatSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sin(&_Vat.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) constant returns(uint256)
func (_Vat *VatCallerSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sin(&_Vat.CallOpts, arg0)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) constant returns(uint256 ink, uint256 art)
func (_Vat *VatCaller) Urns(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	ret := new(struct {
		Ink *big.Int
		Art *big.Int
	})
	out := ret
	err := _Vat.contract.Call(opts, out, "urns", arg0, arg1)
	return *ret, err
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) constant returns(uint256 ink, uint256 art)
func (_Vat *VatSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vat.Contract.Urns(&_Vat.CallOpts, arg0, arg1)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) constant returns(uint256 ink, uint256 art)
func (_Vat *VatCallerSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vat.Contract.Urns(&_Vat.CallOpts, arg0, arg1)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() constant returns(uint256)
func (_Vat *VatCaller) Vice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "vice")
	return *ret0, err
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() constant returns(uint256)
func (_Vat *VatSession) Vice() (*big.Int, error) {
	return _Vat.Contract.Vice(&_Vat.CallOpts)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() constant returns(uint256)
func (_Vat *VatCallerSession) Vice() (*big.Int, error) {
	return _Vat.Contract.Vice(&_Vat.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Vat *VatCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vat.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Vat *VatSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Wards(&_Vat.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Vat *VatCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Wards(&_Vat.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatSession) Cage() (*types.Transaction, error) {
	return _Vat.Contract.Cage(&_Vat.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatTransactorSession) Cage() (*types.Transaction, error) {
	return _Vat.Contract.Cage(&_Vat.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Deny(&_Vat.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Deny(&_Vat.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File(&_Vat.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File(&_Vat.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File0(&_Vat.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File0(&_Vat.TransactOpts, what, data)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatTransactor) Flux(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "flux", ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Flux(&_Vat.TransactOpts, ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatTransactorSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Flux(&_Vat.TransactOpts, ilk, src, dst, wad)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatTransactor) Fold(opts *bind.TransactOpts, i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "fold", i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fold(&_Vat.TransactOpts, i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatTransactorSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fold(&_Vat.TransactOpts, i, u, rate)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Fork(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "fork", ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fork(&_Vat.TransactOpts, ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fork(&_Vat.TransactOpts, ilk, src, dst, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Frob(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "frob", i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Frob(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Frob(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Grab(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "grab", i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Grab(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Grab(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Heal(&_Vat.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Heal(&_Vat.TransactOpts, rad)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatTransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Hope(&_Vat.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatTransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Hope(&_Vat.TransactOpts, usr)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vat.Contract.Init(&_Vat.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vat.Contract.Init(&_Vat.TransactOpts, ilk)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "move", src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Move(&_Vat.TransactOpts, src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatTransactorSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Move(&_Vat.TransactOpts, src, dst, rad)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatTransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Nope(&_Vat.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatTransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Nope(&_Vat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Rely(&_Vat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Rely(&_Vat.TransactOpts, usr)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatTransactor) Slip(opts *bind.TransactOpts, ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "slip", ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Slip(&_Vat.TransactOpts, ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatTransactorSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Slip(&_Vat.TransactOpts, ilk, usr, wad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatTransactor) Suck(opts *bind.TransactOpts, u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "suck", u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Suck(&_Vat.TransactOpts, u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatTransactorSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Suck(&_Vat.TransactOpts, u, v, rad)
}
