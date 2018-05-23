// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unionpay

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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101268061003b6000396000f30060606040526004361060485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114604d578063f2fde38b146079575b600080fd5b3415605757600080fd5b605d6097565b604051600160a060020a03909116815260200160405180910390f35b3415608357600080fd5b6095600160a060020a036004351660a6565b005b600054600160a060020a031681565b60005433600160a060020a0390811691161460c057600080fd5b600160a060020a0381161560f7576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a723058203a3130576341cb1bd0fd0e03cab2b4d26ce2a0a1d3215cfe397ff9cc56c8ac7f0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// UnionPayABI is the input ABI used to generate the binding from.
const UnionPayABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_feePercentage\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"payCash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platform\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_checker\",\"type\":\"address\"}],\"name\":\"setPlatform\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userReceipts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"plainPay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amountIndeed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"UserPay\",\"type\":\"event\"}]"

// UnionPayBin is the compiled bytecode used for deploying new contracts.
const UnionPayBin = `0x606060405260008054600160a060020a033316600160a060020a03199091161790556107ce806100306000396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe0811461009d57806319d38d71146100c25780633ccfd60b146101385780634bde38c81461014d5780636945c5ea1461017c5780638da5cb5b1461019b578063d072ebe9146101ae578063eb3f2427146101d3578063f2fde38b146101db575b600080fd5b34156100a857600080fd5b6100b06101fa565b60405190815260200160405180910390f35b61012460048035906024803591600160a060020a036044351691606435919060a49060843590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061020895505050505050565b604051901515815260200160405180910390f35b341561014357600080fd5b61014b6104c8565b005b341561015857600080fd5b610160610535565b604051600160a060020a03909116815260200160405180910390f35b341561018757600080fd5b61014b600160a060020a0360043516610544565b34156101a657600080fd5b6101606105a3565b34156101b957600080fd5b6100b0600160a060020a03600435811690602435166105b2565b6101246105cf565b34156101e657600080fd5b61014b600160a060020a0360043516610640565b600160a060020a0330163190565b600080600080871015801561021e575060648711155b151561022957600080fd5b600160a060020a038616151561023e57600080fd5b600160a060020a033381166000908152600260209081526040808320938a1683529290522054600101881461027257600080fd5b600154600160a060020a0316151561028957600080fd5b6102e53387348a8c8a6040516c01000000000000000000000000600160a060020a0397881681028252959096169094026014860152602885019290925260488401526068830152608882015260a8016040518091039020610696565b600154909250600160a060020a03166102fe83866106d9565b600160a060020a03161461031157600080fd5b600160a060020a033381166000908152600260209081526040808320938a16835292905220805460010190558615156103f057600034111561037f57600160a060020a0386163480156108fc0290604051600060405180830381858888f19350505050151561037f57600080fd5b7f660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da502338734348c8a604051600160a060020a0396871681529490951660208501526040808501939093526060840191909152608083015260a082019290925260c001905180910390a1600192506104bd565b348702905034878281151561040157fe5b041461040957fe5b60649004600034829003111561044e5785600160a060020a03166108fc8234039081150290604051600060405180830381858888f19350505050151561044e57600080fd5b7f660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da5023387348434038c8a604051600160a060020a0396871681529490951660208501526040808501939093526060840191909152608083015260a082019290925260c001905180910390a1600192505b505095945050505050565b60005433600160a060020a039081169116146104e357600080fd5b600154600160a060020a031615156104fa57600080fd5b600154600160a060020a039081169030163180156108fc0290604051600060405180830381858888f19350505050151561053357600080fd5b565b600154600160a060020a031681565b60005433600160a060020a0390811691161461055f57600080fd5b600160a060020a038116151561057457600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b600260209081526000928352604080842090915290825290205481565b60007f660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da50233303434600080604051600160a060020a0396871681529490951660208501526040808501939093526060840191909152608083015260a082019290925260c001905180910390a150600190565b60005433600160a060020a0390811691161461065b57600080fd5b600160a060020a03811615610693576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b6000816040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0160405180910390209050919050565b6000806000806106e885610767565b919450925090506001868484846040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561075457600080fd5b5050602060405103519695505050505050565b600080600080600080865160411461077e57600080fd5b6020870151925060408701519150606087015160001a9792965090945090925050505600a165627a7a7230582035b96eaa6f2435dbd8cfc8304356c253930652ded413ab3863136a404235c06d0029`

// DeployUnionPay deploys a new Ethereum contract, binding an instance of UnionPay to it.
func DeployUnionPay(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnionPay, error) {
	parsed, err := abi.JSON(strings.NewReader(UnionPayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnionPayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnionPay{UnionPayCaller: UnionPayCaller{contract: contract}, UnionPayTransactor: UnionPayTransactor{contract: contract}, UnionPayFilterer: UnionPayFilterer{contract: contract}}, nil
}

// UnionPay is an auto generated Go binding around an Ethereum contract.
type UnionPay struct {
	UnionPayCaller     // Read-only binding to the contract
	UnionPayTransactor // Write-only binding to the contract
	UnionPayFilterer   // Log filterer for contract events
}

// UnionPayCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnionPayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionPayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnionPayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionPayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnionPayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionPaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnionPaySession struct {
	Contract     *UnionPay         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnionPayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnionPayCallerSession struct {
	Contract *UnionPayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UnionPayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnionPayTransactorSession struct {
	Contract     *UnionPayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnionPayRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnionPayRaw struct {
	Contract *UnionPay // Generic contract binding to access the raw methods on
}

// UnionPayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnionPayCallerRaw struct {
	Contract *UnionPayCaller // Generic read-only contract binding to access the raw methods on
}

// UnionPayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnionPayTransactorRaw struct {
	Contract *UnionPayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnionPay creates a new instance of UnionPay, bound to a specific deployed contract.
func NewUnionPay(address common.Address, backend bind.ContractBackend) (*UnionPay, error) {
	contract, err := bindUnionPay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnionPay{UnionPayCaller: UnionPayCaller{contract: contract}, UnionPayTransactor: UnionPayTransactor{contract: contract}, UnionPayFilterer: UnionPayFilterer{contract: contract}}, nil
}

// NewUnionPayCaller creates a new read-only instance of UnionPay, bound to a specific deployed contract.
func NewUnionPayCaller(address common.Address, caller bind.ContractCaller) (*UnionPayCaller, error) {
	contract, err := bindUnionPay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnionPayCaller{contract: contract}, nil
}

// NewUnionPayTransactor creates a new write-only instance of UnionPay, bound to a specific deployed contract.
func NewUnionPayTransactor(address common.Address, transactor bind.ContractTransactor) (*UnionPayTransactor, error) {
	contract, err := bindUnionPay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnionPayTransactor{contract: contract}, nil
}

// NewUnionPayFilterer creates a new log filterer instance of UnionPay, bound to a specific deployed contract.
func NewUnionPayFilterer(address common.Address, filterer bind.ContractFilterer) (*UnionPayFilterer, error) {
	contract, err := bindUnionPay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnionPayFilterer{contract: contract}, nil
}

// bindUnionPay binds a generic wrapper to an already deployed contract.
func bindUnionPay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnionPayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnionPay *UnionPayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnionPay.Contract.UnionPayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnionPay *UnionPayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.Contract.UnionPayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnionPay *UnionPayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnionPay.Contract.UnionPayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnionPay *UnionPayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UnionPay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnionPay *UnionPayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnionPay *UnionPayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnionPay.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_UnionPay *UnionPayCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_UnionPay *UnionPaySession) GetBalance() (*big.Int, error) {
	return _UnionPay.Contract.GetBalance(&_UnionPay.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_UnionPay *UnionPayCallerSession) GetBalance() (*big.Int, error) {
	return _UnionPay.Contract.GetBalance(&_UnionPay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnionPay *UnionPayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnionPay *UnionPaySession) Owner() (common.Address, error) {
	return _UnionPay.Contract.Owner(&_UnionPay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UnionPay *UnionPayCallerSession) Owner() (common.Address, error) {
	return _UnionPay.Contract.Owner(&_UnionPay.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_UnionPay *UnionPayCaller) Platform(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "platform")
	return *ret0, err
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_UnionPay *UnionPaySession) Platform() (common.Address, error) {
	return _UnionPay.Contract.Platform(&_UnionPay.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_UnionPay *UnionPayCallerSession) Platform() (common.Address, error) {
	return _UnionPay.Contract.Platform(&_UnionPay.CallOpts)
}

// UserReceipts is a free data retrieval call binding the contract method 0xd072ebe9.
//
// Solidity: function userReceipts( address,  address) constant returns(uint256)
func (_UnionPay *UnionPayCaller) UserReceipts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "userReceipts", arg0, arg1)
	return *ret0, err
}

// UserReceipts is a free data retrieval call binding the contract method 0xd072ebe9.
//
// Solidity: function userReceipts( address,  address) constant returns(uint256)
func (_UnionPay *UnionPaySession) UserReceipts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _UnionPay.Contract.UserReceipts(&_UnionPay.CallOpts, arg0, arg1)
}

// UserReceipts is a free data retrieval call binding the contract method 0xd072ebe9.
//
// Solidity: function userReceipts( address,  address) constant returns(uint256)
func (_UnionPay *UnionPayCallerSession) UserReceipts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _UnionPay.Contract.UserReceipts(&_UnionPay.CallOpts, arg0, arg1)
}

// PayCash is a paid mutator transaction binding the contract method 0x19d38d71.
//
// Solidity: function payCash(_nonce uint256, _feePercentage uint256, _to address, _state uint256, _sig bytes) returns(bool)
func (_UnionPay *UnionPayTransactor) PayCash(opts *bind.TransactOpts, _nonce *big.Int, _feePercentage *big.Int, _to common.Address, _state *big.Int, _sig []byte) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "payCash", _nonce, _feePercentage, _to, _state, _sig)
}

// PayCash is a paid mutator transaction binding the contract method 0x19d38d71.
//
// Solidity: function payCash(_nonce uint256, _feePercentage uint256, _to address, _state uint256, _sig bytes) returns(bool)
func (_UnionPay *UnionPaySession) PayCash(_nonce *big.Int, _feePercentage *big.Int, _to common.Address, _state *big.Int, _sig []byte) (*types.Transaction, error) {
	return _UnionPay.Contract.PayCash(&_UnionPay.TransactOpts, _nonce, _feePercentage, _to, _state, _sig)
}

// PayCash is a paid mutator transaction binding the contract method 0x19d38d71.
//
// Solidity: function payCash(_nonce uint256, _feePercentage uint256, _to address, _state uint256, _sig bytes) returns(bool)
func (_UnionPay *UnionPayTransactorSession) PayCash(_nonce *big.Int, _feePercentage *big.Int, _to common.Address, _state *big.Int, _sig []byte) (*types.Transaction, error) {
	return _UnionPay.Contract.PayCash(&_UnionPay.TransactOpts, _nonce, _feePercentage, _to, _state, _sig)
}

// PlainPay is a paid mutator transaction binding the contract method 0xeb3f2427.
//
// Solidity: function plainPay() returns(bool)
func (_UnionPay *UnionPayTransactor) PlainPay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "plainPay")
}

// PlainPay is a paid mutator transaction binding the contract method 0xeb3f2427.
//
// Solidity: function plainPay() returns(bool)
func (_UnionPay *UnionPaySession) PlainPay() (*types.Transaction, error) {
	return _UnionPay.Contract.PlainPay(&_UnionPay.TransactOpts)
}

// PlainPay is a paid mutator transaction binding the contract method 0xeb3f2427.
//
// Solidity: function plainPay() returns(bool)
func (_UnionPay *UnionPayTransactorSession) PlainPay() (*types.Transaction, error) {
	return _UnionPay.Contract.PlainPay(&_UnionPay.TransactOpts)
}

// SetPlatform is a paid mutator transaction binding the contract method 0x6945c5ea.
//
// Solidity: function setPlatform(_checker address) returns()
func (_UnionPay *UnionPayTransactor) SetPlatform(opts *bind.TransactOpts, _checker common.Address) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "setPlatform", _checker)
}

// SetPlatform is a paid mutator transaction binding the contract method 0x6945c5ea.
//
// Solidity: function setPlatform(_checker address) returns()
func (_UnionPay *UnionPaySession) SetPlatform(_checker common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.SetPlatform(&_UnionPay.TransactOpts, _checker)
}

// SetPlatform is a paid mutator transaction binding the contract method 0x6945c5ea.
//
// Solidity: function setPlatform(_checker address) returns()
func (_UnionPay *UnionPayTransactorSession) SetPlatform(_checker common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.SetPlatform(&_UnionPay.TransactOpts, _checker)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnionPay *UnionPayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnionPay *UnionPaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.TransferOwnership(&_UnionPay.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_UnionPay *UnionPayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnionPay.Contract.TransferOwnership(&_UnionPay.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_UnionPay *UnionPayTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnionPay.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_UnionPay *UnionPaySession) Withdraw() (*types.Transaction, error) {
	return _UnionPay.Contract.Withdraw(&_UnionPay.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_UnionPay *UnionPayTransactorSession) Withdraw() (*types.Transaction, error) {
	return _UnionPay.Contract.Withdraw(&_UnionPay.TransactOpts)
}

// UnionPayUserPayIterator is returned from FilterUserPay and is used to iterate over the raw logs and unpacked data for UserPay events raised by the UnionPay contract.
type UnionPayUserPayIterator struct {
	Event *UnionPayUserPay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnionPayUserPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnionPayUserPay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnionPayUserPay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnionPayUserPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnionPayUserPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnionPayUserPay represents a UserPay event raised by the UnionPay contract.
type UnionPayUserPay struct {
	From         common.Address
	To           common.Address
	Amount       *big.Int
	AmountIndeed *big.Int
	Nonce        *big.Int
	State        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserPay is a free log retrieval operation binding the contract event 0x660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da502.
//
// Solidity: event UserPay(_from address, _to address, _amount uint256, _amountIndeed uint256, _nonce uint256, _state uint256)
func (_UnionPay *UnionPayFilterer) FilterUserPay(opts *bind.FilterOpts) (*UnionPayUserPayIterator, error) {

	logs, sub, err := _UnionPay.contract.FilterLogs(opts, "UserPay")
	if err != nil {
		return nil, err
	}
	return &UnionPayUserPayIterator{contract: _UnionPay.contract, event: "UserPay", logs: logs, sub: sub}, nil
}

// WatchUserPay is a free log subscription operation binding the contract event 0x660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da502.
//
// Solidity: event UserPay(_from address, _to address, _amount uint256, _amountIndeed uint256, _nonce uint256, _state uint256)
func (_UnionPay *UnionPayFilterer) WatchUserPay(opts *bind.WatchOpts, sink chan<- *UnionPayUserPay) (event.Subscription, error) {

	logs, sub, err := _UnionPay.contract.WatchLogs(opts, "UserPay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnionPayUserPay)
				if err := _UnionPay.contract.UnpackLog(event, "UserPay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
