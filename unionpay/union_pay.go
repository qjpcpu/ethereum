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
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a03199091161790556101838061003d6000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008e575b600080fd5b34801561005c57600080fd5b506100656100be565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100da565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161461010257600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615610154576000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b505600a165627a7a723058206908870f8d671b43b3fec717762a3f17c259e82e869775a49e9efb2a646c76f40029`

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
const UnionPayABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_feePercentage\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_state\",\"type\":\"uint256\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"payCash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platform\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_checker\",\"type\":\"address\"}],\"name\":\"setPlatform\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userReceipts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"plainPay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_amountIndeed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_state\",\"type\":\"uint256\"}],\"name\":\"UserPay\",\"type\":\"event\"}]"

// UnionPayBin is the compiled bytecode used for deploying new contracts.
const UnionPayBin = `0x608060405260008054600160a060020a033316600160a060020a03199091161790556107b1806100306000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe0811461009d57806319d38d71146100c45780633ccfd60b1461013f5780634bde38c8146101565780636945c5ea146101875780638da5cb5b146101a8578063a843c5fe146101bd578063eb3f2427146101e1578063f2fde38b146101e9575b600080fd5b3480156100a957600080fd5b506100b261020a565b60408051918252519081900360200190f35b604080516020601f60843560048181013592830184900484028501840190955281845261012b948035946024803595600160a060020a0360443516956064359536959460a49490939101919081908401838280828437509497506102189650505050505050565b604080519115158252519081900360200190f35b34801561014b57600080fd5b506101546104c0565b005b34801561016257600080fd5b5061016b610530565b60408051600160a060020a039092168252519081900360200190f35b34801561019357600080fd5b50610154600160a060020a036004351661053f565b3480156101b457600080fd5b5061016b61059e565b3480156101c957600080fd5b506100b2600160a060020a03600435166024356105ad565b61012b6105ca565b3480156101f557600080fd5b50610154600160a060020a0360043516610630565b600160a060020a0330163190565b600080600080871015801561022e575060648711155b151561023957600080fd5b600160a060020a038616151561024e57600080fd5b600160a060020a03331660009081526002602090815260408083208b84529091529020541561027c57600080fd5b600154600160a060020a0316151561029357600080fd5b604080516c01000000000000000000000000600160a060020a0333811682028352891602601482015234602882015260488101899052606881018a90526088810187905290519081900360a80190206102eb90610687565b600154909250600160a060020a031661030483866106c5565b600160a060020a03161461031757600080fd5b600160a060020a03331660009081526002602090815260408083208b84529091529020600190558615156103ec57600034111561038557604051600160a060020a038716903480156108fc02916000818181858888f19350505050158015610383573d6000803e3d6000fd5b505b60408051600160a060020a03338116825288166020820152348183018190526060820152608081018a905260a0810187905290517f660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da5029181900360c00190a1600192506104b5565b34870290503487828115156103fd57fe5b041461040557fe5b60649004600034829003111561044f57604051600160a060020a038716903483900380156108fc02916000818181858888f1935050505015801561044d573d6000803e3d6000fd5b505b60408051600160a060020a03338116825288166020820152348183018190528390036060820152608081018a905260a0810187905290517f660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da5029181900360c00190a1600192505b505095945050505050565b60005433600160a060020a039081169116146104db57600080fd5b600154600160a060020a031615156104f257600080fd5b600154604051600160a060020a039182169130163180156108fc02916000818181858888f1935050505015801561052d573d6000803e3d6000fd5b50565b600154600160a060020a031681565b60005433600160a060020a0390811691161461055a57600080fd5b600160a060020a038116151561056f57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031681565b600260209081526000928352604080842090915290825290205481565b60408051600160a060020a0333811682523016602082015234818301819052606082015260006080820181905260a0820181905291517f660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da5029181900360c00190a150600190565b60005433600160a060020a0390811691161461064b57600080fd5b600160a060020a0381161561052d5760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101839052905190819003603c019020919050565b6000806000806106d48561074c565b60408051600080825260208083018085528d905260ff8716838501526060830186905260808301859052925195985093965091945060019360a0808401949293601f19830193908390039091019190865af1158015610737573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000806000806000808651604114151561076557600080fd5b505050506020830151604084015160609094015160001a949093925090505600a165627a7a72305820252da380623b01e8c841ef2e4044a66abe015d7ff0b379bcdfadef72412387ed0029`

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

// UserReceipts is a free data retrieval call binding the contract method 0xa843c5fe.
//
// Solidity: function userReceipts( address,  uint256) constant returns(uint256)
func (_UnionPay *UnionPayCaller) UserReceipts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UnionPay.contract.Call(opts, out, "userReceipts", arg0, arg1)
	return *ret0, err
}

// UserReceipts is a free data retrieval call binding the contract method 0xa843c5fe.
//
// Solidity: function userReceipts( address,  uint256) constant returns(uint256)
func (_UnionPay *UnionPaySession) UserReceipts(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _UnionPay.Contract.UserReceipts(&_UnionPay.CallOpts, arg0, arg1)
}

// UserReceipts is a free data retrieval call binding the contract method 0xa843c5fe.
//
// Solidity: function userReceipts( address,  uint256) constant returns(uint256)
func (_UnionPay *UnionPayCallerSession) UserReceipts(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
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
// Solidity: e UserPay(_from address, _to address, _amount uint256, _amountIndeed uint256, _nonce uint256, _state uint256)
func (_UnionPay *UnionPayFilterer) FilterUserPay(opts *bind.FilterOpts) (*UnionPayUserPayIterator, error) {

	logs, sub, err := _UnionPay.contract.FilterLogs(opts, "UserPay")
	if err != nil {
		return nil, err
	}
	return &UnionPayUserPayIterator{contract: _UnionPay.contract, event: "UserPay", logs: logs, sub: sub}, nil
}

// WatchUserPay is a free log subscription operation binding the contract event 0x660d0baf29f0dde0dec971546aa706496532cd4be797578fc27143fbb38da502.
//
// Solidity: e UserPay(_from address, _to address, _amount uint256, _amountIndeed uint256, _nonce uint256, _state uint256)
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
