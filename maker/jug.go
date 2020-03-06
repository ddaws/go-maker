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

// JugBin is the compiled bytecode used for deploying new contracts.
var JugBin = "0x608060405234801561001057600080fd5b5060405161131b38038061131b8339818101604052602081101561003357600080fd5b810190808051906020019092919050505060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611242806100d96000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063626cb3c511610071578063626cb3c51461020b57806365fae35e146102555780639c52a7f114610299578063bf353dbb146102dd578063d4e8be8314610335578063d9638d3614610383576100b4565b80631a0b287e146100b957806329ae8114146100fb57806336569e77146101335780633b6631951461017d57806344e2a5a8146101ab5780635001f3b5146101ed575b600080fd5b6100f9600480360360608110156100cf57600080fd5b810190808035906020019092919080359060200190929190803590602001909291905050506103cc565b005b6101316004803603604081101561011157600080fd5b8101908080359060200190929190803590602001909291905050506105f9565b005b61013b610786565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101a96004803603602081101561019357600080fd5b81019080803590602001909291905050506107ac565b005b6101d7600480360360208110156101c157600080fd5b8101908080359060200190929190505050610946565b6040518082815260200191505060405180910390f35b6101f5610c35565b6040518082815260200191505060405180910390f35b610213610c3b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102976004803603602081101561026b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c61565b005b6102db600480360360208110156102af57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d8f565b005b61031f600480360360208110156102f357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ebd565b6040518082815260200191505060405180910390f35b6103816004803603604081101561034b57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ed5565b005b6103af6004803603602081101561039957600080fd5b810190808035906020019092919050505061109c565b604051808381526020018281526020019250505060405180910390f35b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610480576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4a75672f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b6001600084815260200190815260200160002060010154421461050b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f4a75672f72686f2d6e6f742d757064617465640000000000000000000000000081525060200191505060405180910390fd5b7f6475747900000000000000000000000000000000000000000000000000000000821415610553578060016000858152602001908152602001600020600001819055506105c1565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4a75672f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450505050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146106ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4a75672f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b7f62617365000000000000000000000000000000000000000000000000000000008214156106e1578060048190555061074f565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4a75672f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610860576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4a75672f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b600060016000838152602001908152602001600020905060008160000154146108f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f4a75672f696c6b2d616c72656164792d696e697400000000000000000000000081525060200191505060405180910390fd5b6b033b2e3c9fd0803ce80000008160000181905550428160010181905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600060016000838152602001908152602001600020600101544210156109d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f4a75672f696e76616c69642d6e6f77000000000000000000000000000000000081525060200191505060405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d9638d36846040518263ffffffff1660e01b8152600401808281526020019150506040805180830381600087803b158015610a4a57600080fd5b505af1158015610a5e573d6000803e3d6000fd5b505050506040513d6040811015610a7457600080fd5b810190808051906020019092919080519060200190929190505050915050610aeb610ae5610aba60045460016000888152602001908152602001600020600001546110c0565b600160008781526020019081526020016000206001015442036b033b2e3c9fd0803ce80000006110da565b826111a0565b9150600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b65337df84600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610b5986866111e5565b6040518463ffffffff1660e01b8152600401808481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610bc957600080fd5b505af1158015610bdd573d6000803e3d6000fd5b50505050426001600085815260200190815260200160002060010181905550505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450919050565b60045481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610d15576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4a75672f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610e43576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4a75672f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60006020528060005260406000206000915090505481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610f89576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4a75672f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b7f766f770000000000000000000000000000000000000000000000000000000000821415610ff75780600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611065565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f4a75672f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b60016020528060005260406000206000915090508060000154908060010154905082565b60008183019050828110156110d457600080fd5b92915050565b600083600081146111805760028406600081146110f9578592506110fd565b8392505b50600283046002850494505b841561117a57858602868782041461112057600080fd5b8181018181101561113057600080fd5b8581049750600287061561116d57878502858982041415891515161561115557600080fd5b8381018181101561116557600080fd5b878104965050505b5050600285049450611109565b50611198565b83600081146111925760009250611196565b8392505b505b509392505050565b6000818302905060008214806111be5750828282816111bb57fe5b04145b6111c757600080fd5b6b033b2e3c9fd0803ce800000081816111dc57fe5b04905092915050565b60008183039050600083121580156111fe575060008212155b61120757600080fd5b9291505056fea265627a7a72315820014492f997e4f26c7bee172f217606d889b54f464c1cb7107c0701e63824fc8064736f6c63430005100032"

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
