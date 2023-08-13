// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymaster

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OverheadIgpDomainConfig is an auto generated low-level Go binding around an user-defined struct.
type OverheadIgpDomainConfig struct {
	Domain      uint32
	GasOverhead *big.Int
}

// PaymasterMetaData contains all meta data concerning the Paymaster contract.
var PaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_innerIgp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"}],\"name\":\"DestinationGasOverheadSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_gasAmount\",\"type\":\"uint256\"}],\"name\":\"destinationGasAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"destinationGasOverhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"innerIgp\",\"outputs\":[{\"internalType\":\"contractIInterchainGasPaymaster\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"payForGas\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_gasAmount\",\"type\":\"uint256\"}],\"name\":\"quoteGasPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"}],\"internalType\":\"structOverheadIgp.DomainConfig[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"name\":\"setDestinationGasOverheads\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymasterMetaData.ABI instead.
var PaymasterABI = PaymasterMetaData.ABI

// Paymaster is an auto generated Go binding around an Ethereum contract.
type Paymaster struct {
	PaymasterCaller     // Read-only binding to the contract
	PaymasterTransactor // Write-only binding to the contract
	PaymasterFilterer   // Log filterer for contract events
}

// PaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymasterSession struct {
	Contract     *Paymaster        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymasterCallerSession struct {
	Contract *PaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymasterTransactorSession struct {
	Contract     *PaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymasterRaw struct {
	Contract *Paymaster // Generic contract binding to access the raw methods on
}

// PaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymasterCallerRaw struct {
	Contract *PaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// PaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymasterTransactorRaw struct {
	Contract *PaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymaster creates a new instance of Paymaster, bound to a specific deployed contract.
func NewPaymaster(address common.Address, backend bind.ContractBackend) (*Paymaster, error) {
	contract, err := bindPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paymaster{PaymasterCaller: PaymasterCaller{contract: contract}, PaymasterTransactor: PaymasterTransactor{contract: contract}, PaymasterFilterer: PaymasterFilterer{contract: contract}}, nil
}

// NewPaymasterCaller creates a new read-only instance of Paymaster, bound to a specific deployed contract.
func NewPaymasterCaller(address common.Address, caller bind.ContractCaller) (*PaymasterCaller, error) {
	contract, err := bindPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymasterCaller{contract: contract}, nil
}

// NewPaymasterTransactor creates a new write-only instance of Paymaster, bound to a specific deployed contract.
func NewPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymasterTransactor, error) {
	contract, err := bindPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymasterTransactor{contract: contract}, nil
}

// NewPaymasterFilterer creates a new log filterer instance of Paymaster, bound to a specific deployed contract.
func NewPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymasterFilterer, error) {
	contract, err := bindPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymasterFilterer{contract: contract}, nil
}

// bindPaymaster binds a generic wrapper to an already deployed contract.
func bindPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymaster *PaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymaster.Contract.PaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymaster *PaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.Contract.PaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymaster *PaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymaster.Contract.PaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymaster *PaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymaster *PaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymaster *PaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymaster.Contract.contract.Transact(opts, method, params...)
}

// DestinationGasAmount is a free data retrieval call binding the contract method 0x2fd88292.
//
// Solidity: function destinationGasAmount(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_Paymaster *PaymasterCaller) DestinationGasAmount(opts *bind.CallOpts, _destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "destinationGasAmount", _destinationDomain, _gasAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestinationGasAmount is a free data retrieval call binding the contract method 0x2fd88292.
//
// Solidity: function destinationGasAmount(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_Paymaster *PaymasterSession) DestinationGasAmount(_destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	return _Paymaster.Contract.DestinationGasAmount(&_Paymaster.CallOpts, _destinationDomain, _gasAmount)
}

// DestinationGasAmount is a free data retrieval call binding the contract method 0x2fd88292.
//
// Solidity: function destinationGasAmount(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_Paymaster *PaymasterCallerSession) DestinationGasAmount(_destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	return _Paymaster.Contract.DestinationGasAmount(&_Paymaster.CallOpts, _destinationDomain, _gasAmount)
}

// DestinationGasOverhead is a free data retrieval call binding the contract method 0x026c3348.
//
// Solidity: function destinationGasOverhead(uint32 ) view returns(uint256)
func (_Paymaster *PaymasterCaller) DestinationGasOverhead(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "destinationGasOverhead", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestinationGasOverhead is a free data retrieval call binding the contract method 0x026c3348.
//
// Solidity: function destinationGasOverhead(uint32 ) view returns(uint256)
func (_Paymaster *PaymasterSession) DestinationGasOverhead(arg0 uint32) (*big.Int, error) {
	return _Paymaster.Contract.DestinationGasOverhead(&_Paymaster.CallOpts, arg0)
}

// DestinationGasOverhead is a free data retrieval call binding the contract method 0x026c3348.
//
// Solidity: function destinationGasOverhead(uint32 ) view returns(uint256)
func (_Paymaster *PaymasterCallerSession) DestinationGasOverhead(arg0 uint32) (*big.Int, error) {
	return _Paymaster.Contract.DestinationGasOverhead(&_Paymaster.CallOpts, arg0)
}

// InnerIgp is a free data retrieval call binding the contract method 0x78ecc5ac.
//
// Solidity: function innerIgp() view returns(address)
func (_Paymaster *PaymasterCaller) InnerIgp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "innerIgp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InnerIgp is a free data retrieval call binding the contract method 0x78ecc5ac.
//
// Solidity: function innerIgp() view returns(address)
func (_Paymaster *PaymasterSession) InnerIgp() (common.Address, error) {
	return _Paymaster.Contract.InnerIgp(&_Paymaster.CallOpts)
}

// InnerIgp is a free data retrieval call binding the contract method 0x78ecc5ac.
//
// Solidity: function innerIgp() view returns(address)
func (_Paymaster *PaymasterCallerSession) InnerIgp() (common.Address, error) {
	return _Paymaster.Contract.InnerIgp(&_Paymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymaster *PaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymaster *PaymasterSession) Owner() (common.Address, error) {
	return _Paymaster.Contract.Owner(&_Paymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paymaster *PaymasterCallerSession) Owner() (common.Address, error) {
	return _Paymaster.Contract.Owner(&_Paymaster.CallOpts)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xa6929793.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_Paymaster *PaymasterCaller) QuoteGasPayment(opts *bind.CallOpts, _destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paymaster.contract.Call(opts, &out, "quoteGasPayment", _destinationDomain, _gasAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xa6929793.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_Paymaster *PaymasterSession) QuoteGasPayment(_destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	return _Paymaster.Contract.QuoteGasPayment(&_Paymaster.CallOpts, _destinationDomain, _gasAmount)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xa6929793.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain, uint256 _gasAmount) view returns(uint256)
func (_Paymaster *PaymasterCallerSession) QuoteGasPayment(_destinationDomain uint32, _gasAmount *big.Int) (*big.Int, error) {
	return _Paymaster.Contract.QuoteGasPayment(&_Paymaster.CallOpts, _destinationDomain, _gasAmount)
}

// PayForGas is a paid mutator transaction binding the contract method 0x11bf2c18.
//
// Solidity: function payForGas(bytes32 _messageId, uint32 _destinationDomain, uint256 _gasAmount, address _refundAddress) payable returns()
func (_Paymaster *PaymasterTransactor) PayForGas(opts *bind.TransactOpts, _messageId [32]byte, _destinationDomain uint32, _gasAmount *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "payForGas", _messageId, _destinationDomain, _gasAmount, _refundAddress)
}

// PayForGas is a paid mutator transaction binding the contract method 0x11bf2c18.
//
// Solidity: function payForGas(bytes32 _messageId, uint32 _destinationDomain, uint256 _gasAmount, address _refundAddress) payable returns()
func (_Paymaster *PaymasterSession) PayForGas(_messageId [32]byte, _destinationDomain uint32, _gasAmount *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.PayForGas(&_Paymaster.TransactOpts, _messageId, _destinationDomain, _gasAmount, _refundAddress)
}

// PayForGas is a paid mutator transaction binding the contract method 0x11bf2c18.
//
// Solidity: function payForGas(bytes32 _messageId, uint32 _destinationDomain, uint256 _gasAmount, address _refundAddress) payable returns()
func (_Paymaster *PaymasterTransactorSession) PayForGas(_messageId [32]byte, _destinationDomain uint32, _gasAmount *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.PayForGas(&_Paymaster.TransactOpts, _messageId, _destinationDomain, _gasAmount, _refundAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paymaster *PaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paymaster *PaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paymaster.Contract.RenounceOwnership(&_Paymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paymaster *PaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paymaster.Contract.RenounceOwnership(&_Paymaster.TransactOpts)
}

// SetDestinationGasOverheads is a paid mutator transaction binding the contract method 0x94e92298.
//
// Solidity: function setDestinationGasOverheads((uint32,uint256)[] configs) returns()
func (_Paymaster *PaymasterTransactor) SetDestinationGasOverheads(opts *bind.TransactOpts, configs []OverheadIgpDomainConfig) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "setDestinationGasOverheads", configs)
}

// SetDestinationGasOverheads is a paid mutator transaction binding the contract method 0x94e92298.
//
// Solidity: function setDestinationGasOverheads((uint32,uint256)[] configs) returns()
func (_Paymaster *PaymasterSession) SetDestinationGasOverheads(configs []OverheadIgpDomainConfig) (*types.Transaction, error) {
	return _Paymaster.Contract.SetDestinationGasOverheads(&_Paymaster.TransactOpts, configs)
}

// SetDestinationGasOverheads is a paid mutator transaction binding the contract method 0x94e92298.
//
// Solidity: function setDestinationGasOverheads((uint32,uint256)[] configs) returns()
func (_Paymaster *PaymasterTransactorSession) SetDestinationGasOverheads(configs []OverheadIgpDomainConfig) (*types.Transaction, error) {
	return _Paymaster.Contract.SetDestinationGasOverheads(&_Paymaster.TransactOpts, configs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paymaster *PaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Paymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paymaster *PaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.TransferOwnership(&_Paymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paymaster *PaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paymaster.Contract.TransferOwnership(&_Paymaster.TransactOpts, newOwner)
}

// PaymasterDestinationGasOverheadSetIterator is returned from FilterDestinationGasOverheadSet and is used to iterate over the raw logs and unpacked data for DestinationGasOverheadSet events raised by the Paymaster contract.
type PaymasterDestinationGasOverheadSetIterator struct {
	Event *PaymasterDestinationGasOverheadSet // Event containing the contract specifics and raw log

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
func (it *PaymasterDestinationGasOverheadSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymasterDestinationGasOverheadSet)
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
		it.Event = new(PaymasterDestinationGasOverheadSet)
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
func (it *PaymasterDestinationGasOverheadSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymasterDestinationGasOverheadSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymasterDestinationGasOverheadSet represents a DestinationGasOverheadSet event raised by the Paymaster contract.
type PaymasterDestinationGasOverheadSet struct {
	Domain      uint32
	GasOverhead *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDestinationGasOverheadSet is a free log retrieval operation binding the contract event 0x8685397d4fa4489d21ed19c302e3719e5b1f0acd46b0ef39b3775f5bfa85b910.
//
// Solidity: event DestinationGasOverheadSet(uint32 indexed domain, uint256 gasOverhead)
func (_Paymaster *PaymasterFilterer) FilterDestinationGasOverheadSet(opts *bind.FilterOpts, domain []uint32) (*PaymasterDestinationGasOverheadSetIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Paymaster.contract.FilterLogs(opts, "DestinationGasOverheadSet", domainRule)
	if err != nil {
		return nil, err
	}
	return &PaymasterDestinationGasOverheadSetIterator{contract: _Paymaster.contract, event: "DestinationGasOverheadSet", logs: logs, sub: sub}, nil
}

// WatchDestinationGasOverheadSet is a free log subscription operation binding the contract event 0x8685397d4fa4489d21ed19c302e3719e5b1f0acd46b0ef39b3775f5bfa85b910.
//
// Solidity: event DestinationGasOverheadSet(uint32 indexed domain, uint256 gasOverhead)
func (_Paymaster *PaymasterFilterer) WatchDestinationGasOverheadSet(opts *bind.WatchOpts, sink chan<- *PaymasterDestinationGasOverheadSet, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Paymaster.contract.WatchLogs(opts, "DestinationGasOverheadSet", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymasterDestinationGasOverheadSet)
				if err := _Paymaster.contract.UnpackLog(event, "DestinationGasOverheadSet", log); err != nil {
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

// ParseDestinationGasOverheadSet is a log parse operation binding the contract event 0x8685397d4fa4489d21ed19c302e3719e5b1f0acd46b0ef39b3775f5bfa85b910.
//
// Solidity: event DestinationGasOverheadSet(uint32 indexed domain, uint256 gasOverhead)
func (_Paymaster *PaymasterFilterer) ParseDestinationGasOverheadSet(log types.Log) (*PaymasterDestinationGasOverheadSet, error) {
	event := new(PaymasterDestinationGasOverheadSet)
	if err := _Paymaster.contract.UnpackLog(event, "DestinationGasOverheadSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Paymaster contract.
type PaymasterOwnershipTransferredIterator struct {
	Event *PaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymasterOwnershipTransferred)
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
		it.Event = new(PaymasterOwnershipTransferred)
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
func (it *PaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the Paymaster contract.
type PaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paymaster *PaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymasterOwnershipTransferredIterator{contract: _Paymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paymaster *PaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymasterOwnershipTransferred)
				if err := _Paymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paymaster *PaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*PaymasterOwnershipTransferred, error) {
	event := new(PaymasterOwnershipTransferred)
	if err := _Paymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
