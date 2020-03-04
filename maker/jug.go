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

// JugABI is the input ABI used to generate the binding from.
const JugABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rho\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// JugFuncSigs maps the 4-byte function signature to its string representation.
var JugFuncSigs = map[string]string{
	"5001f3b5": "base()",
	"9c52a7f1": "deny(address)",
	"44e2a5a8": "drip(bytes32)",
	"d4e8be83": "file(bytes32,address)",
	"1a0b287e": "file(bytes32,bytes32,uint256)",
	"29ae8114": "file(bytes32,uint256)",
	"d9638d36": "ilks(bytes32)",
	"3b663195": "init(bytes32)",
	"65fae35e": "rely(address)",
	"36569e77": "vat()",
	"626cb3c5": "vow()",
	"bf353dbb": "wards(address)",
}

// JugBin is the compiled bytecode used for deploying new contracts.
var JugBin = "0x608060405234801561001057600080fd5b50604051610b7f380380610b7f8339818101604052602081101561003357600080fd5b505133600090815260208190526040902060019055600280546001600160a01b039092166001600160a01b0319909216919091179055610b07806100786000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063626cb3c511610071578063626cb3c51461017f57806365fae35e146101875780639c52a7f1146101ad578063bf353dbb146101d3578063d4e8be83146101f9578063d9638d3614610225576100b4565b80631a0b287e146100b957806329ae8114146100e457806336569e77146101075780633b6631951461012b57806344e2a5a8146101485780635001f3b514610177575b600080fd5b6100e2600480360360608110156100cf57600080fd5b508035906020810135906040013561025a565b005b6100e2600480360360408110156100fa57600080fd5b50803590602001356103b9565b61010f610461565b6040516001600160a01b03909116815260200160405180910390f35b6100e26004803603602081101561014157600080fd5b5035610470565b6101656004803603602081101561015e57600080fd5b5035610573565b60405190815260200160405180910390f35b610165610785565b61010f61078b565b6100e26004803603602081101561019d57600080fd5b50356001600160a01b031661079a565b6100e2600480360360208110156101c357600080fd5b50356001600160a01b0316610848565b610165600480360360208110156101e957600080fd5b50356001600160a01b03166108f3565b6100e26004803603604081101561020f57600080fd5b50803590602001356001600160a01b0316610907565b6102426004803603602081101561023b57600080fd5b503561098d565b60405191825260208201526040908101905180910390f35b336000908152602081905260409020546001146102b25760405162461bcd60e51b8152602060048201526012602482015271129d59cbdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b600083815260016020526040902060010154421461030c5760405162461bcd60e51b8152602060048201526013602482015272129d59cbdc9a1bcb5b9bdd0b5d5c19185d1959606a1b604482015260640160405180910390fd5b81636475747960e01b141561033157600083815260016020528190604090205561037d565b60405162461bcd60e51b815260206004820152601b60248201527f4a75672f66696c652d756e7265636f676e697a65642d706172616d0000000000604482015260640160405180910390fd5b5961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a450505050565b336000908152602081905260409020546001146104115760405162461bcd60e51b8152602060048201526012602482015271129d59cbdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b81636261736560e01b14156103315760048190555b5961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a4505050565b6002546001600160a01b031681565b336000908152602081905260409020546001146104c85760405162461bcd60e51b8152602060048201526012602482015271129d59cbdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b60008181526001602052604081208054909150156105235760405162461bcd60e51b8152602060048201526014602482015273129d59cbda5b1acb585b1c9958591e4b5a5b9a5d60621b604482015260640160405180910390fd5b6b033b2e3c9fd0803ce80000008155426001909101555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b6000818152600160205260408120600101544210156105ca5760405162461bcd60e51b815260206004820152600f60248201526e4a75672f696e76616c69642d6e6f7760881b604482015260640160405180910390fd5b6002546000906001600160a01b031663d9638d36846040516001600160e01b031960e084901b16815260048101919091526024016040805180830381600087803b15801561061757600080fd5b505af115801561062b573d6000803e3d6000fd5b505050506040513d604081101561064157600080fd5b8101908080519291906020018051600454600089815260016020529196506106a695506106a0945061067a9350915060409020546109a9565b60008681526001602052604090206001015442036b033b2e3c9fd0803ce80000006109bf565b82610a7e565b6002546003549193506001600160a01b039081169163b65337df918691166106ce8686610ab4565b6040516001600160e01b031960e086901b16815260048101939093526001600160a01b0390911660248301526044820152606401600060405180830381600087803b15801561071c57600080fd5b505af1158015610730573d6000803e3d6000fd5b505050600084815260016020524291506040902060010155505961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a450919050565b60045481565b6003546001600160a01b031681565b336000908152602081905260409020546001146107f25760405162461bcd60e51b8152602060048201526012602482015271129d59cbdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260019060409020555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b336000908152602081905260409020546001146108a05760405162461bcd60e51b8152602060048201526012602482015271129d59cbdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260408120555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b600060205280600052604060002054905081565b3360009081526020819052604090205460011461095f5760405162461bcd60e51b8152602060048201526012602482015271129d59cbdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b8162766f7760e81b141561033157600380546001600160a01b0319166001600160a01b038316179055610426565b6001602052806000526040600020805460019091015490915082565b818101828110156109b957600080fd5b92915050565b6000838015610a60576001841680156109da578592506109de565b8392505b50600283046002850494505b8415610a5a578586028687820414610a0157600080fd5b81810181811015610a1157600080fd5b85810497506002870615610a4d578785028589820414158915151615610a3657600080fd5b83810181811015610a4657600080fd5b8790049550505b50506002850494506109ea565b50610a76565b838015610a705760009250610a74565b8392505b505b509392505050565b818102811580610a96575082828281610a9357fe5b04145b610a9f57600080fd5b6b033b2e3c9fd0803ce8000000900492915050565b80820360008312801590610ac9575060008212155b6109b957600080fdfea265627a7a7231582021208e93a064a968090088b7fac6524a5380b0de3c015fec109aa26ba92cae8c64736f6c63430005100032"

// DeployJug deploys a new Ethereum contract, binding an instance of Jug to it.
func DeployJug(auth *bind.TransactOpts, backend bind.ContractBackend, vat_ common.Address) (common.Address, *types.Transaction, *Jug, error) {
	parsed, err := abi.JSON(strings.NewReader(JugABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(JugBin), backend, vat_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jug{JugCaller: JugCaller{contract: contract}, JugTransactor: JugTransactor{contract: contract}, JugFilterer: JugFilterer{contract: contract}}, nil
}

// Jug is an auto generated Go binding around an Ethereum contract.
type Jug struct {
	JugCaller     // Read-only binding to the contract
	JugTransactor // Write-only binding to the contract
	JugFilterer   // Log filterer for contract events
}

// JugCaller is an auto generated read-only Go binding around an Ethereum contract.
type JugCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JugTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JugFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JugSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JugSession struct {
	Contract     *Jug              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JugCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JugCallerSession struct {
	Contract *JugCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JugTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JugTransactorSession struct {
	Contract     *JugTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JugRaw is an auto generated low-level Go binding around an Ethereum contract.
type JugRaw struct {
	Contract *Jug // Generic contract binding to access the raw methods on
}

// JugCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JugCallerRaw struct {
	Contract *JugCaller // Generic read-only contract binding to access the raw methods on
}

// JugTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JugTransactorRaw struct {
	Contract *JugTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJug creates a new instance of Jug, bound to a specific deployed contract.
func NewJug(address common.Address, backend bind.ContractBackend) (*Jug, error) {
	contract, err := bindJug(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jug{JugCaller: JugCaller{contract: contract}, JugTransactor: JugTransactor{contract: contract}, JugFilterer: JugFilterer{contract: contract}}, nil
}

// NewJugCaller creates a new read-only instance of Jug, bound to a specific deployed contract.
func NewJugCaller(address common.Address, caller bind.ContractCaller) (*JugCaller, error) {
	contract, err := bindJug(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JugCaller{contract: contract}, nil
}

// NewJugTransactor creates a new write-only instance of Jug, bound to a specific deployed contract.
func NewJugTransactor(address common.Address, transactor bind.ContractTransactor) (*JugTransactor, error) {
	contract, err := bindJug(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JugTransactor{contract: contract}, nil
}

// NewJugFilterer creates a new log filterer instance of Jug, bound to a specific deployed contract.
func NewJugFilterer(address common.Address, filterer bind.ContractFilterer) (*JugFilterer, error) {
	contract, err := bindJug(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JugFilterer{contract: contract}, nil
}

// bindJug binds a generic wrapper to an already deployed contract.
func bindJug(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JugABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jug *JugRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Jug.Contract.JugCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jug *JugRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jug.Contract.JugTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jug *JugRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jug.Contract.JugTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jug *JugCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Jug.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jug *JugTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jug.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jug *JugTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jug.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() constant returns(uint256)
func (_Jug *JugCaller) Base(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Jug.contract.Call(opts, out, "base")
	return *ret0, err
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() constant returns(uint256)
func (_Jug *JugSession) Base() (*big.Int, error) {
	return _Jug.Contract.Base(&_Jug.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() constant returns(uint256)
func (_Jug *JugCallerSession) Base() (*big.Int, error) {
	return _Jug.Contract.Base(&_Jug.CallOpts)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(uint256 duty, uint256 rho)
func (_Jug *JugCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	ret := new(struct {
		Duty *big.Int
		Rho  *big.Int
	})
	out := ret
	err := _Jug.contract.Call(opts, out, "ilks", arg0)
	return *ret, err
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(uint256 duty, uint256 rho)
func (_Jug *JugSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _Jug.Contract.Ilks(&_Jug.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) constant returns(uint256 duty, uint256 rho)
func (_Jug *JugCallerSession) Ilks(arg0 [32]byte) (struct {
	Duty *big.Int
	Rho  *big.Int
}, error) {
	return _Jug.Contract.Ilks(&_Jug.CallOpts, arg0)
}

// VatAddress is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Jug *JugCaller) VatAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Jug.contract.Call(opts, out, "vat")
	return *ret0, err
}

// VatAddress is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Jug *JugSession) VatAddress() (common.Address, error) {
	return _Jug.Contract.VatAddress(&_Jug.CallOpts)
}

// VatAddress is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Jug *JugCallerSession) VatAddress() (common.Address, error) {
	return _Jug.Contract.VatAddress(&_Jug.CallOpts)
}

// VowAddress is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Jug *JugCaller) VowAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Jug.contract.Call(opts, out, "vow")
	return *ret0, err
}

// VowAddress is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Jug *JugSession) VowAddress() (common.Address, error) {
	return _Jug.Contract.VowAddress(&_Jug.CallOpts)
}

// VowAddress is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Jug *JugCallerSession) VowAddress() (common.Address, error) {
	return _Jug.Contract.VowAddress(&_Jug.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Jug *JugCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Jug.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Jug *JugSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Jug.Contract.Wards(&_Jug.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Jug *JugCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Jug.Contract.Wards(&_Jug.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jug *JugTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jug *JugSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Deny(&_Jug.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Jug *JugTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Deny(&_Jug.TransactOpts, usr)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jug *JugTransactor) Drip(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "drip", ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jug *JugSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Drip(&_Jug.TransactOpts, ilk)
}

// Drip is a paid mutator transaction binding the contract method 0x44e2a5a8.
//
// Solidity: function drip(bytes32 ilk) returns(uint256 rate)
func (_Jug *JugTransactorSession) Drip(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Drip(&_Jug.TransactOpts, ilk)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jug *JugTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jug *JugSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File(&_Jug.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Jug *JugTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File(&_Jug.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jug *JugTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jug *JugSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File0(&_Jug.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Jug *JugTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Jug.Contract.File0(&_Jug.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jug *JugTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jug *JugSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jug.Contract.File1(&_Jug.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Jug *JugTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Jug.Contract.File1(&_Jug.TransactOpts, what, data)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jug *JugTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jug *JugSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Init(&_Jug.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Jug *JugTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Jug.Contract.Init(&_Jug.TransactOpts, ilk)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jug *JugTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Jug.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jug *JugSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Rely(&_Jug.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Jug *JugTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Jug.Contract.Rely(&_Jug.TransactOpts, usr)
}

// LibNoteABI is the input ABI used to generate the binding from.
const LibNoteABI = "[{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"}]"

// LibNoteBin is the compiled bytecode used for deploying new contracts.
var LibNoteBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723158205641d71df588d7386fb7c3d047fdaeb07e01dfbf338abb84e5e041249ea361b264736f6c63430005100032"

// DeployLibNote deploys a new Ethereum contract, binding an instance of LibNote to it.
func DeployLibNote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibNote, error) {
	parsed, err := abi.JSON(strings.NewReader(LibNoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LibNoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibNote{LibNoteCaller: LibNoteCaller{contract: contract}, LibNoteTransactor: LibNoteTransactor{contract: contract}, LibNoteFilterer: LibNoteFilterer{contract: contract}}, nil
}

// LibNote is an auto generated Go binding around an Ethereum contract.
type LibNote struct {
	LibNoteCaller     // Read-only binding to the contract
	LibNoteTransactor // Write-only binding to the contract
	LibNoteFilterer   // Log filterer for contract events
}

// LibNoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibNoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibNoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibNoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibNoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibNoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibNoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibNoteSession struct {
	Contract     *LibNote          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibNoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibNoteCallerSession struct {
	Contract *LibNoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LibNoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibNoteTransactorSession struct {
	Contract     *LibNoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LibNoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibNoteRaw struct {
	Contract *LibNote // Generic contract binding to access the raw methods on
}

// LibNoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibNoteCallerRaw struct {
	Contract *LibNoteCaller // Generic read-only contract binding to access the raw methods on
}

// LibNoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibNoteTransactorRaw struct {
	Contract *LibNoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibNote creates a new instance of LibNote, bound to a specific deployed contract.
func NewLibNote(address common.Address, backend bind.ContractBackend) (*LibNote, error) {
	contract, err := bindLibNote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibNote{LibNoteCaller: LibNoteCaller{contract: contract}, LibNoteTransactor: LibNoteTransactor{contract: contract}, LibNoteFilterer: LibNoteFilterer{contract: contract}}, nil
}

// NewLibNoteCaller creates a new read-only instance of LibNote, bound to a specific deployed contract.
func NewLibNoteCaller(address common.Address, caller bind.ContractCaller) (*LibNoteCaller, error) {
	contract, err := bindLibNote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibNoteCaller{contract: contract}, nil
}

// NewLibNoteTransactor creates a new write-only instance of LibNote, bound to a specific deployed contract.
func NewLibNoteTransactor(address common.Address, transactor bind.ContractTransactor) (*LibNoteTransactor, error) {
	contract, err := bindLibNote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibNoteTransactor{contract: contract}, nil
}

// NewLibNoteFilterer creates a new log filterer instance of LibNote, bound to a specific deployed contract.
func NewLibNoteFilterer(address common.Address, filterer bind.ContractFilterer) (*LibNoteFilterer, error) {
	contract, err := bindLibNote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibNoteFilterer{contract: contract}, nil
}

// bindLibNote binds a generic wrapper to an already deployed contract.
func bindLibNote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibNoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibNote *LibNoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LibNote.Contract.LibNoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibNote *LibNoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibNote.Contract.LibNoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibNote *LibNoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibNote.Contract.LibNoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibNote *LibNoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LibNote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibNote *LibNoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibNote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibNote *LibNoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibNote.Contract.contract.Transact(opts, method, params...)
}

// VatLikeABI is the input ABI used to generate the binding from.
const VatLikeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"fold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"Art\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VatLikeFuncSigs maps the 4-byte function signature to its string representation.
var VatLikeFuncSigs = map[string]string{
	"b65337df": "fold(bytes32,address,int256)",
	"d9638d36": "ilks(bytes32)",
}

// VatLike is an auto generated Go binding around an Ethereum contract.
type VatLike struct {
	VatLikeCaller     // Read-only binding to the contract
	VatLikeTransactor // Write-only binding to the contract
	VatLikeFilterer   // Log filterer for contract events
}

// VatLikeCaller is an auto generated read-only Go binding around an Ethereum contract.
type VatLikeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatLikeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VatLikeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatLikeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VatLikeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatLikeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VatLikeSession struct {
	Contract     *VatLike          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatLikeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VatLikeCallerSession struct {
	Contract *VatLikeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VatLikeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VatLikeTransactorSession struct {
	Contract     *VatLikeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VatLikeRaw is an auto generated low-level Go binding around an Ethereum contract.
type VatLikeRaw struct {
	Contract *VatLike // Generic contract binding to access the raw methods on
}

// VatLikeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VatLikeCallerRaw struct {
	Contract *VatLikeCaller // Generic read-only contract binding to access the raw methods on
}

// VatLikeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VatLikeTransactorRaw struct {
	Contract *VatLikeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVatLike creates a new instance of VatLike, bound to a specific deployed contract.
func NewVatLike(address common.Address, backend bind.ContractBackend) (*VatLike, error) {
	contract, err := bindVatLike(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VatLike{VatLikeCaller: VatLikeCaller{contract: contract}, VatLikeTransactor: VatLikeTransactor{contract: contract}, VatLikeFilterer: VatLikeFilterer{contract: contract}}, nil
}

// NewVatLikeCaller creates a new read-only instance of VatLike, bound to a specific deployed contract.
func NewVatLikeCaller(address common.Address, caller bind.ContractCaller) (*VatLikeCaller, error) {
	contract, err := bindVatLike(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VatLikeCaller{contract: contract}, nil
}

// NewVatLikeTransactor creates a new write-only instance of VatLike, bound to a specific deployed contract.
func NewVatLikeTransactor(address common.Address, transactor bind.ContractTransactor) (*VatLikeTransactor, error) {
	contract, err := bindVatLike(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VatLikeTransactor{contract: contract}, nil
}

// NewVatLikeFilterer creates a new log filterer instance of VatLike, bound to a specific deployed contract.
func NewVatLikeFilterer(address common.Address, filterer bind.ContractFilterer) (*VatLikeFilterer, error) {
	contract, err := bindVatLike(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VatLikeFilterer{contract: contract}, nil
}

// bindVatLike binds a generic wrapper to an already deployed contract.
func bindVatLike(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VatLikeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VatLike *VatLikeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VatLike.Contract.VatLikeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VatLike *VatLikeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VatLike.Contract.VatLikeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VatLike *VatLikeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VatLike.Contract.VatLikeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VatLike *VatLikeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VatLike.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VatLike *VatLikeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VatLike.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VatLike *VatLikeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VatLike.Contract.contract.Transact(opts, method, params...)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 , address , int256 ) returns()
func (_VatLike *VatLikeTransactor) Fold(opts *bind.TransactOpts, arg0 [32]byte, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.contract.Transact(opts, "fold", arg0, arg1, arg2)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 , address , int256 ) returns()
func (_VatLike *VatLikeSession) Fold(arg0 [32]byte, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.Contract.Fold(&_VatLike.TransactOpts, arg0, arg1, arg2)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 , address , int256 ) returns()
func (_VatLike *VatLikeTransactorSession) Fold(arg0 [32]byte, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.Contract.Fold(&_VatLike.TransactOpts, arg0, arg1, arg2)
}

// Ilks is a paid mutator transaction binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) returns(uint256 Art, uint256 rate)
func (_VatLike *VatLikeTransactor) Ilks(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _VatLike.contract.Transact(opts, "ilks", arg0)
}

// Ilks is a paid mutator transaction binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) returns(uint256 Art, uint256 rate)
func (_VatLike *VatLikeSession) Ilks(arg0 [32]byte) (*types.Transaction, error) {
	return _VatLike.Contract.Ilks(&_VatLike.TransactOpts, arg0)
}

// Ilks is a paid mutator transaction binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) returns(uint256 Art, uint256 rate)
func (_VatLike *VatLikeTransactorSession) Ilks(arg0 [32]byte) (*types.Transaction, error) {
	return _VatLike.Contract.Ilks(&_VatLike.TransactOpts, arg0)
}
