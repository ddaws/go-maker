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

// PotABI is the input ABI used to generate the binding from.
const PotABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Pie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tmp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dsr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rho\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PotFuncSigs maps the 4-byte function signature to its string representation.
var PotFuncSigs = map[string]string{
	"2c69ed58": "Pie()",
	"69245009": "cage()",
	"c92aecc4": "chi()",
	"9c52a7f1": "deny(address)",
	"9f678cca": "drip()",
	"487bf082": "dsr()",
	"7f8661a1": "exit(uint256)",
	"d4e8be83": "file(bytes32,address)",
	"29ae8114": "file(bytes32,uint256)",
	"049878f3": "join(uint256)",
	"957aa58c": "live()",
	"0bebac86": "pie(address)",
	"65fae35e": "rely(address)",
	"20aba08b": "rho()",
	"36569e77": "vat()",
	"626cb3c5": "vow()",
	"bf353dbb": "wards(address)",
}

// PotBin is the compiled bytecode used for deploying new contracts.
var PotBin = "0x608060405234801561001057600080fd5b50604051610c23380380610c238339818101604052602081101561003357600080fd5b5051336000908152602081905260409020600190819055600580546001600160a01b039093166001600160a01b0319909316929092179091556b033b2e3c9fd0803ce8000000600381905560045542600755600855610b8c806100976000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806365fae35e116100a25780639c52a7f1116100715780639c52a7f1146102215780639f678cca14610247578063bf353dbb1461024f578063c92aecc414610275578063d4e8be831461027d5761010b565b806365fae35e146101ce57806369245009146101f45780637f8661a1146101fc578063957aa58c146102195761010b565b80632c69ed58116100de5780632c69ed581461019257806336569e771461019a578063487bf082146101be578063626cb3c5146101c65761010b565b8063049878f3146101105780630bebac861461012f57806320aba08b1461016757806329ae81141461016f575b600080fd5b61012d6004803603602081101561012657600080fd5b50356102a9565b005b6101556004803603602081101561014557600080fd5b50356001600160a01b03166103f8565b60405190815260200160405180910390f35b61015561040c565b61012d6004803603604081101561018557600080fd5b5080359060200135610412565b610155610599565b6101a261059f565b6040516001600160a01b03909116815260200160405180910390f35b6101556105ae565b6101a26105b4565b61012d600480360360208110156101e457600080fd5b50356001600160a01b03166105c3565b61012d610671565b61012d6004803603602081101561021257600080fd5b5035610717565b610155610777565b61012d6004803603602081101561023757600080fd5b50356001600160a01b031661077d565b610155610828565b6101556004803603602081101561026557600080fd5b50356001600160a01b0316610986565b61015561099a565b61012d6004803603604081101561029357600080fd5b50803590602001356001600160a01b03166109a0565b60075442146102f45760405162461bcd60e51b8152602060048201526013602482015272141bdd0bdc9a1bcb5b9bdd0b5d5c19185d1959606a1b604482015260640160405180910390fd5b336000908152600160205261030e90604090205482610a26565b3360009081526001602052604090205560025461032b9082610a26565b6002556005546004546001600160a01b039091169063bb35783b90339030906103549086610a3c565b6040516001600160e01b031960e086901b1681526001600160a01b0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b1580156103a657600080fd5b505af11580156103ba573d6000803e3d6000fd5b505050505961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b600160205280600052604060002054905081565b60075481565b3360009081526020819052604090205460011461046a5760405162461bcd60e51b8152602060048201526012602482015271141bdd0bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6008546001146104af5760405162461bcd60e51b815260206004820152600c60248201526b506f742f6e6f742d6c69766560a01b604482015260640160405180910390fd5b60075442146104fa5760405162461bcd60e51b8152602060048201526013602482015272141bdd0bdc9a1bcb5b9bdd0b5d5c19185d1959606a1b604482015260640160405180910390fd5b81623239b960e91b141561051257600381905561055e565b60405162461bcd60e51b815260206004820152601b60248201527f506f742f66696c652d756e7265636f676e697a65642d706172616d0000000000604482015260640160405180910390fd5b5961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a4505050565b60025481565b6005546001600160a01b031681565b60035481565b6006546001600160a01b031681565b3360009081526020819052604090205460011461061b5760405162461bcd60e51b8152602060048201526012602482015271141bdd0bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260019060409020555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b336000908152602081905260409020546001146106c95760405162461bcd60e51b8152602060048201526012602482015271141bdd0bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b60006008556b033b2e3c9fd0803ce80000006003555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a450565b336000908152600160205261073190604090205482610a60565b3360009081526001602052604090205560025461074e9082610a60565b6002556005546004546001600160a01b039091169063bb35783b90309033906103549086610a3c565b60085481565b336000908152602081905260409020546001146107d55760405162461bcd60e51b8152602060048201526012602482015271141bdd0bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b6001600160a01b0381166000908152602081905260408120555961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45050565b60006007544210156108725760405162461bcd60e51b815260206004820152600f60248201526e506f742f696e76616c69642d6e6f7760881b604482015260640160405180910390fd5b61089a61089260035460075442036b033b2e3c9fd0803ce8000000610a70565b600454610b2f565b905060006108aa82600454610a60565b6004839055426007556005546006546002549293506001600160a01b039182169263f24e23eb929091169030906108e19086610a3c565b6040516001600160e01b031960e086901b1681526001600160a01b0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b15801561093357600080fd5b505af1158015610947573d6000803e3d6000fd5b50505050505961012081016040526020815260e0602082015260e060006040830137602435600435336001600160e01b03196000351661012085a45090565b600060205280600052604060002054905081565b60045481565b336000908152602081905260409020546001146109f85760405162461bcd60e51b8152602060048201526012602482015271141bdd0bdb9bdd0b585d5d1a1bdc9a5e995960721b604482015260640160405180910390fd5b8162766f7760e81b141561051257600680546001600160a01b0319166001600160a01b03831617905561055e565b80820182811015610a3657600080fd5b92915050565b6000811580610a5757505080820282828281610a5457fe5b04145b610a3657600080fd5b80820382811115610a3657600080fd5b6000838015610b1157600184168015610a8b57859250610a8f565b8392505b50600283046002850494505b8415610b0b578586028687820414610ab257600080fd5b81810181811015610ac257600080fd5b85810497506002870615610afe578785028589820414158915151615610ae757600080fd5b83810181811015610af757600080fd5b8790049550505b5050600285049450610a9b565b50610b27565b838015610b215760009250610b25565b8392505b505b509392505050565b60006b033b2e3c9fd0803ce8000000610b488484610a3c565b81610b4f57fe5b04939250505056fea265627a7a7231582042a4a30441c55387fd212de2fb3ed381ada38ee965a56c228fb942b92bc83c9064736f6c63430005100032"

// DeployPot deploys a new Ethereum contract, binding an instance of Pot to it.
func DeployPot(auth *bind.TransactOpts, backend bind.ContractBackend, vat_ common.Address) (common.Address, *types.Transaction, *Pot, error) {
	parsed, err := abi.JSON(strings.NewReader(PotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PotBin), backend, vat_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pot{PotCaller: PotCaller{contract: contract}, PotTransactor: PotTransactor{contract: contract}, PotFilterer: PotFilterer{contract: contract}}, nil
}

// Pot is an auto generated Go binding around an Ethereum contract.
type Pot struct {
	PotCaller     // Read-only binding to the contract
	PotTransactor // Write-only binding to the contract
	PotFilterer   // Log filterer for contract events
}

// PotCaller is an auto generated read-only Go binding around an Ethereum contract.
type PotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PotSession struct {
	Contract     *Pot              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PotCallerSession struct {
	Contract *PotCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PotTransactorSession struct {
	Contract     *PotTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PotRaw is an auto generated low-level Go binding around an Ethereum contract.
type PotRaw struct {
	Contract *Pot // Generic contract binding to access the raw methods on
}

// PotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PotCallerRaw struct {
	Contract *PotCaller // Generic read-only contract binding to access the raw methods on
}

// PotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PotTransactorRaw struct {
	Contract *PotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPot creates a new instance of Pot, bound to a specific deployed contract.
func NewPot(address common.Address, backend bind.ContractBackend) (*Pot, error) {
	contract, err := bindPot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pot{PotCaller: PotCaller{contract: contract}, PotTransactor: PotTransactor{contract: contract}, PotFilterer: PotFilterer{contract: contract}}, nil
}

// NewPotCaller creates a new read-only instance of Pot, bound to a specific deployed contract.
func NewPotCaller(address common.Address, caller bind.ContractCaller) (*PotCaller, error) {
	contract, err := bindPot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PotCaller{contract: contract}, nil
}

// NewPotTransactor creates a new write-only instance of Pot, bound to a specific deployed contract.
func NewPotTransactor(address common.Address, transactor bind.ContractTransactor) (*PotTransactor, error) {
	contract, err := bindPot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PotTransactor{contract: contract}, nil
}

// NewPotFilterer creates a new log filterer instance of Pot, bound to a specific deployed contract.
func NewPotFilterer(address common.Address, filterer bind.ContractFilterer) (*PotFilterer, error) {
	contract, err := bindPot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PotFilterer{contract: contract}, nil
}

// bindPot binds a generic wrapper to an already deployed contract.
func bindPot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pot *PotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pot.Contract.PotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pot *PotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pot.Contract.PotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pot *PotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pot.Contract.PotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pot *PotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pot *PotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pot *PotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pot.Contract.contract.Transact(opts, method, params...)
}

// TotalPie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() constant returns(uint256)
func (_Pot *PotCaller) TotalPie(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "Pie")
	return *ret0, err
}

// TotalPie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() constant returns(uint256)
func (_Pot *PotSession) TotalPie() (*big.Int, error) {
	return _Pot.Contract.TotalPie(&_Pot.CallOpts)
}

// TotalPie is a free data retrieval call binding the contract method 0x2c69ed58.
//
// Solidity: function Pie() constant returns(uint256)
func (_Pot *PotCallerSession) TotalPie() (*big.Int, error) {
	return _Pot.Contract.TotalPie(&_Pot.CallOpts)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() constant returns(uint256)
func (_Pot *PotCaller) Chi(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "chi")
	return *ret0, err
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() constant returns(uint256)
func (_Pot *PotSession) Chi() (*big.Int, error) {
	return _Pot.Contract.Chi(&_Pot.CallOpts)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() constant returns(uint256)
func (_Pot *PotCallerSession) Chi() (*big.Int, error) {
	return _Pot.Contract.Chi(&_Pot.CallOpts)
}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() constant returns(uint256)
func (_Pot *PotCaller) Dsr(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "dsr")
	return *ret0, err
}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() constant returns(uint256)
func (_Pot *PotSession) Dsr() (*big.Int, error) {
	return _Pot.Contract.Dsr(&_Pot.CallOpts)
}

// Dsr is a free data retrieval call binding the contract method 0x487bf082.
//
// Solidity: function dsr() constant returns(uint256)
func (_Pot *PotCallerSession) Dsr() (*big.Int, error) {
	return _Pot.Contract.Dsr(&_Pot.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Pot *PotCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "live")
	return *ret0, err
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Pot *PotSession) Live() (*big.Int, error) {
	return _Pot.Contract.Live(&_Pot.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() constant returns(uint256)
func (_Pot *PotCallerSession) Live() (*big.Int, error) {
	return _Pot.Contract.Live(&_Pot.CallOpts)
}

// Pie is a free data retrieval call binding the contract method 0x0bebac86.
//
// Solidity: function pie(address ) constant returns(uint256)
func (_Pot *PotCaller) Pie(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "pie", arg0)
	return *ret0, err
}

// Pie is a free data retrieval call binding the contract method 0x0bebac86.
//
// Solidity: function pie(address ) constant returns(uint256)
func (_Pot *PotSession) Pie(arg0 common.Address) (*big.Int, error) {
	return _Pot.Contract.Pie(&_Pot.CallOpts, arg0)
}

// Pie is a free data retrieval call binding the contract method 0x0bebac86.
//
// Solidity: function pie(address ) constant returns(uint256)
func (_Pot *PotCallerSession) Pie(arg0 common.Address) (*big.Int, error) {
	return _Pot.Contract.Pie(&_Pot.CallOpts, arg0)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() constant returns(uint256)
func (_Pot *PotCaller) Rho(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "rho")
	return *ret0, err
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() constant returns(uint256)
func (_Pot *PotSession) Rho() (*big.Int, error) {
	return _Pot.Contract.Rho(&_Pot.CallOpts)
}

// Rho is a free data retrieval call binding the contract method 0x20aba08b.
//
// Solidity: function rho() constant returns(uint256)
func (_Pot *PotCallerSession) Rho() (*big.Int, error) {
	return _Pot.Contract.Rho(&_Pot.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Pot *PotCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "vat")
	return *ret0, err
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Pot *PotSession) Vat() (common.Address, error) {
	return _Pot.Contract.Vat(&_Pot.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() constant returns(address)
func (_Pot *PotCallerSession) Vat() (common.Address, error) {
	return _Pot.Contract.Vat(&_Pot.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Pot *PotCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "vow")
	return *ret0, err
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Pot *PotSession) Vow() (common.Address, error) {
	return _Pot.Contract.Vow(&_Pot.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() constant returns(address)
func (_Pot *PotCallerSession) Vow() (common.Address, error) {
	return _Pot.Contract.Vow(&_Pot.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Pot *PotCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pot.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Pot *PotSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Pot.Contract.Wards(&_Pot.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) constant returns(uint256)
func (_Pot *PotCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Pot.Contract.Wards(&_Pot.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Pot *PotTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Pot *PotSession) Cage() (*types.Transaction, error) {
	return _Pot.Contract.Cage(&_Pot.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Pot *PotTransactorSession) Cage() (*types.Transaction, error) {
	return _Pot.Contract.Cage(&_Pot.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Pot *PotTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Pot *PotSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Pot.Contract.Deny(&_Pot.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Pot *PotTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Pot.Contract.Deny(&_Pot.TransactOpts, guy)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_Pot *PotTransactor) Drip(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "drip")
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_Pot *PotSession) Drip() (*types.Transaction, error) {
	return _Pot.Contract.Drip(&_Pot.TransactOpts)
}

// Drip is a paid mutator transaction binding the contract method 0x9f678cca.
//
// Solidity: function drip() returns(uint256 tmp)
func (_Pot *PotTransactorSession) Drip() (*types.Transaction, error) {
	return _Pot.Contract.Drip(&_Pot.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_Pot *PotTransactor) Exit(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "exit", wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_Pot *PotSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _Pot.Contract.Exit(&_Pot.TransactOpts, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_Pot *PotTransactorSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _Pot.Contract.Exit(&_Pot.TransactOpts, wad)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Pot *PotTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Pot *PotSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Pot.Contract.File(&_Pot.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Pot *PotTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Pot.Contract.File(&_Pot.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_Pot *PotTransactor) File0(opts *bind.TransactOpts, what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "file0", what, addr)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_Pot *PotSession) File0(what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Pot.Contract.File0(&_Pot.TransactOpts, what, addr)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address addr) returns()
func (_Pot *PotTransactorSession) File0(what [32]byte, addr common.Address) (*types.Transaction, error) {
	return _Pot.Contract.File0(&_Pot.TransactOpts, what, addr)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_Pot *PotTransactor) Join(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "join", wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_Pot *PotSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _Pot.Contract.Join(&_Pot.TransactOpts, wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_Pot *PotTransactorSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _Pot.Contract.Join(&_Pot.TransactOpts, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Pot *PotTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Pot.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Pot *PotSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Pot.Contract.Rely(&_Pot.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Pot *PotTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Pot.Contract.Rely(&_Pot.TransactOpts, guy)
}

// VatLikeABI is the input ABI used to generate the binding from.
const VatLikeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"suck\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VatLikeFuncSigs maps the 4-byte function signature to its string representation.
var VatLikeFuncSigs = map[string]string{
	"bb35783b": "move(address,address,uint256)",
	"f24e23eb": "suck(address,address,uint256)",
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

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address , address , uint256 ) returns()
func (_VatLike *VatLikeTransactor) Move(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.contract.Transact(opts, "move", arg0, arg1, arg2)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address , address , uint256 ) returns()
func (_VatLike *VatLikeSession) Move(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.Contract.Move(&_VatLike.TransactOpts, arg0, arg1, arg2)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address , address , uint256 ) returns()
func (_VatLike *VatLikeTransactorSession) Move(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.Contract.Move(&_VatLike.TransactOpts, arg0, arg1, arg2)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address , address , uint256 ) returns()
func (_VatLike *VatLikeTransactor) Suck(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.contract.Transact(opts, "suck", arg0, arg1, arg2)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address , address , uint256 ) returns()
func (_VatLike *VatLikeSession) Suck(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.Contract.Suck(&_VatLike.TransactOpts, arg0, arg1, arg2)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address , address , uint256 ) returns()
func (_VatLike *VatLikeTransactorSession) Suck(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _VatLike.Contract.Suck(&_VatLike.TransactOpts, arg0, arg1, arg2)
}
