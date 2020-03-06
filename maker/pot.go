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

// PotABI is the input ABI used to generate the binding from.
const PotABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Pie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tmp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dsr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rho\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PotBin is the compiled bytecode used for deploying new contracts.
var PotBin = "0x608060405234801561001057600080fd5b506040516115853803806115858339818101604052602081101561003357600080fd5b810190808051906020019092919050505060016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506b033b2e3c9fd0803ce80000006003819055506b033b2e3c9fd0803ce8000000600481905550426007819055506001600881905550506114778061010e6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806365fae35e116100a25780639c52a7f1116100715780639c52a7f1146103565780639f678cca1461039a578063bf353dbb146103b8578063c92aecc414610410578063d4e8be831461042e5761010b565b806365fae35e146102bc57806369245009146103005780637f8661a11461030a578063957aa58c146103385761010b565b80632c69ed58116100de5780632c69ed58146101ec57806336569e771461020a578063487bf08214610254578063626cb3c5146102725761010b565b8063049878f3146101105780630bebac861461013e57806320aba08b1461019657806329ae8114146101b4575b600080fd5b61013c6004803603602081101561012657600080fd5b810190808035906020019092919050505061047c565b005b6101806004803603602081101561015457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106c7565b6040518082815260200191505060405180910390f35b61019e6106df565b6040518082815260200191505060405180910390f35b6101ea600480360360408110156101ca57600080fd5b8101908080359060200190929190803590602001909291905050506106e5565b005b6101f4610961565b6040518082815260200191505060405180910390f35b610212610967565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61025c61098d565b6040518082815260200191505060405180910390f35b61027a610993565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102fe600480360360208110156102d257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109b9565b005b610308610ae7565b005b6103366004803603602081101561032057600080fd5b8101908080359060200190929190505050610beb565b005b610340610dbf565b6040518082815260200191505060405180910390f35b6103986004803603602081101561036c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610dc5565b005b6103a2610ef3565b6040518082815260200191505060405180910390f35b6103fa600480360360208110156103ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061110e565b6040518082815260200191505060405180910390f35b610418611126565b6040518082815260200191505060405180910390f35b61047a6004803603604081101561044457600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061112c565b005b60075442146104f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f506f742f72686f2d6e6f742d757064617465640000000000000000000000000081525060200191505060405180910390fd5b61053c600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826112f3565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061058b600254826112f3565b600281905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b33306105dd6004548661130d565b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15801561067957600080fd5b505af115801561068d573d6000803e3d6000fd5b505050505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016020528060005260406000206000915090505481565b60075481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610799576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f506f742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b600160085414610811576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f506f742f6e6f742d6c697665000000000000000000000000000000000000000081525060200191505060405180910390fd5b6007544214610888576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f506f742f72686f2d6e6f742d757064617465640000000000000000000000000081525060200191505060405180910390fd5b7f64737200000000000000000000000000000000000000000000000000000000008214156108bc578060038190555061092a565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f506f742f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b60025481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610a6d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f506f742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610b9b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f506f742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60006008819055506b033b2e3c9fd0803ce80000006003819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a450565b610c34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482611339565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610c8360025482611339565b600281905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb35783b3033610cd56004548661130d565b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610d7157600080fd5b505af1158015610d85573d6000803e3d6000fd5b505050505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b60085481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610e79576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f506f742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b6000600754421015610f6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f506f742f696e76616c69642d6e6f77000000000000000000000000000000000081525060200191505060405180910390fd5b610f95610f8d60035460075442036b033b2e3c9fd0803ce8000000611353565b600454611419565b90506000610fa582600454611339565b90508160048190555042600781905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f24e23eb600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16306110236002548661130d565b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b1580156110bf57600080fd5b505af11580156110d3573d6000803e3d6000fd5b50505050505961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45090565b60006020528060005260406000206000915090505481565b60045481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054146111e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f506f742f6e6f742d617574686f72697a6564000000000000000000000000000081525060200191505060405180910390fd5b7f766f77000000000000000000000000000000000000000000000000000000000082141561124e5780600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112bc565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f506f742f66696c652d756e7265636f676e697a65642d706172616d000000000081525060200191505060405180910390fd5b5961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a4505050565b600082828401915081101561130757600080fd5b92915050565b60008082148061132a575082828385029250828161132757fe5b04145b61133357600080fd5b92915050565b600082828403915081111561134d57600080fd5b92915050565b600083600081146113f957600284066000811461137257859250611376565b8392505b50600283046002850494505b84156113f357858602868782041461139957600080fd5b818101818110156113a957600080fd5b858104975060028706156113e65787850285898204141589151516156113ce57600080fd5b838101818110156113de57600080fd5b878104965050505b5050600285049450611382565b50611411565b836000811461140b576000925061140f565b8392505b505b509392505050565b60006b033b2e3c9fd0803ce8000000611432848461130d565b8161143957fe5b0490509291505056fea265627a7a72315820ff8ba90017c254f3f5592e4f8b086f8d871188602073605ddf8b6854b682c05764736f6c63430005100032"

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
