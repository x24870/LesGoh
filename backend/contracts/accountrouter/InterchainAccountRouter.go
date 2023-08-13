// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountrouter

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

// CallLibCall is an auto generated low-level Go binding around an user-defined struct.
type CallLibCall struct {
	To    [32]byte
	Value *big.Int
	Data  []byte
}

// AccountrouterMetaData contains all meta data concerning the Accountrouter contract.
var AccountrouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"owner\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ism\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InterchainAccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"interchainGasPaymaster\",\"type\":\"address\"}],\"name\":\"InterchainGasPaymasterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainSecurityModuleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mailbox\",\"type\":\"address\"}],\"name\":\"MailboxSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"router\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ism\",\"type\":\"bytes32\"}],\"name\":\"RemoteCallDispatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ism\",\"type\":\"bytes32\"}],\"name\":\"RemoteIsmEnrolled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"router\",\"type\":\"bytes32\"}],\"name\":\"RemoteRouterEnrolled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"callRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCallLib.Call[]\",\"name\":\"_calls\",\"type\":\"tuple[]\"}],\"name\":\"callRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_ism\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCallLib.Call[]\",\"name\":\"_calls\",\"type\":\"tuple[]\"}],\"name\":\"callRemoteWithOverrides\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"}],\"name\":\"enrollRemoteRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_ism\",\"type\":\"bytes32\"}],\"name\":\"enrollRemoteRouterAndIsm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_destinations\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_routers\",\"type\":\"bytes32[]\"}],\"name\":\"enrollRemoteRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ism\",\"type\":\"address\"}],\"name\":\"getDeployedInterchainAccount\",\"outputs\":[{\"internalType\":\"contractOwnableMulticall\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_owner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_ism\",\"type\":\"address\"}],\"name\":\"getDeployedInterchainAccount\",\"outputs\":[{\"internalType\":\"contractOwnableMulticall\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ism\",\"type\":\"address\"}],\"name\":\"getLocalInterchainAccount\",\"outputs\":[{\"internalType\":\"contractOwnableMulticall\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_owner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_ism\",\"type\":\"address\"}],\"name\":\"getLocalInterchainAccount\",\"outputs\":[{\"internalType\":\"contractOwnableMulticall\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getRemoteInterchainAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ism\",\"type\":\"address\"}],\"name\":\"getRemoteInterchainAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mailbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_interchainGasPaymaster\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_interchainSecurityModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainGasPaymaster\",\"outputs\":[{\"internalType\":\"contractIInterchainGasPaymaster\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainSecurityModule\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"isms\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mailbox\",\"outputs\":[{\"internalType\":\"contractIMailbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"routers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchainGasPaymaster\",\"type\":\"address\"}],\"name\":\"setInterchainGasPaymaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"setInterchainSecurityModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mailbox\",\"type\":\"address\"}],\"name\":\"setMailbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AccountrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountrouterMetaData.ABI instead.
var AccountrouterABI = AccountrouterMetaData.ABI

// Accountrouter is an auto generated Go binding around an Ethereum contract.
type Accountrouter struct {
	AccountrouterCaller     // Read-only binding to the contract
	AccountrouterTransactor // Write-only binding to the contract
	AccountrouterFilterer   // Log filterer for contract events
}

// AccountrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountrouterSession struct {
	Contract     *Accountrouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountrouterCallerSession struct {
	Contract *AccountrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccountrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountrouterTransactorSession struct {
	Contract     *AccountrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccountrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountrouterRaw struct {
	Contract *Accountrouter // Generic contract binding to access the raw methods on
}

// AccountrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountrouterCallerRaw struct {
	Contract *AccountrouterCaller // Generic read-only contract binding to access the raw methods on
}

// AccountrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountrouterTransactorRaw struct {
	Contract *AccountrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountrouter creates a new instance of Accountrouter, bound to a specific deployed contract.
func NewAccountrouter(address common.Address, backend bind.ContractBackend) (*Accountrouter, error) {
	contract, err := bindAccountrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accountrouter{AccountrouterCaller: AccountrouterCaller{contract: contract}, AccountrouterTransactor: AccountrouterTransactor{contract: contract}, AccountrouterFilterer: AccountrouterFilterer{contract: contract}}, nil
}

// NewAccountrouterCaller creates a new read-only instance of Accountrouter, bound to a specific deployed contract.
func NewAccountrouterCaller(address common.Address, caller bind.ContractCaller) (*AccountrouterCaller, error) {
	contract, err := bindAccountrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountrouterCaller{contract: contract}, nil
}

// NewAccountrouterTransactor creates a new write-only instance of Accountrouter, bound to a specific deployed contract.
func NewAccountrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountrouterTransactor, error) {
	contract, err := bindAccountrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountrouterTransactor{contract: contract}, nil
}

// NewAccountrouterFilterer creates a new log filterer instance of Accountrouter, bound to a specific deployed contract.
func NewAccountrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountrouterFilterer, error) {
	contract, err := bindAccountrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountrouterFilterer{contract: contract}, nil
}

// bindAccountrouter binds a generic wrapper to an already deployed contract.
func bindAccountrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountrouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountrouter *AccountrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accountrouter.Contract.AccountrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountrouter *AccountrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountrouter.Contract.AccountrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountrouter *AccountrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountrouter.Contract.AccountrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accountrouter *AccountrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accountrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accountrouter *AccountrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accountrouter *AccountrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accountrouter.Contract.contract.Transact(opts, method, params...)
}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_Accountrouter *AccountrouterCaller) Domains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "domains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_Accountrouter *AccountrouterSession) Domains() ([]uint32, error) {
	return _Accountrouter.Contract.Domains(&_Accountrouter.CallOpts)
}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_Accountrouter *AccountrouterCallerSession) Domains() ([]uint32, error) {
	return _Accountrouter.Contract.Domains(&_Accountrouter.CallOpts)
}

// GetLocalInterchainAccount is a free data retrieval call binding the contract method 0x9662e8ac.
//
// Solidity: function getLocalInterchainAccount(uint32 _origin, address _owner, address _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterCaller) GetLocalInterchainAccount(opts *bind.CallOpts, _origin uint32, _owner common.Address, _router common.Address, _ism common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "getLocalInterchainAccount", _origin, _owner, _router, _ism)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocalInterchainAccount is a free data retrieval call binding the contract method 0x9662e8ac.
//
// Solidity: function getLocalInterchainAccount(uint32 _origin, address _owner, address _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterSession) GetLocalInterchainAccount(_origin uint32, _owner common.Address, _router common.Address, _ism common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetLocalInterchainAccount(&_Accountrouter.CallOpts, _origin, _owner, _router, _ism)
}

// GetLocalInterchainAccount is a free data retrieval call binding the contract method 0x9662e8ac.
//
// Solidity: function getLocalInterchainAccount(uint32 _origin, address _owner, address _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterCallerSession) GetLocalInterchainAccount(_origin uint32, _owner common.Address, _router common.Address, _ism common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetLocalInterchainAccount(&_Accountrouter.CallOpts, _origin, _owner, _router, _ism)
}

// GetLocalInterchainAccount0 is a free data retrieval call binding the contract method 0xa5da3611.
//
// Solidity: function getLocalInterchainAccount(uint32 _origin, bytes32 _owner, bytes32 _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterCaller) GetLocalInterchainAccount0(opts *bind.CallOpts, _origin uint32, _owner [32]byte, _router [32]byte, _ism common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "getLocalInterchainAccount0", _origin, _owner, _router, _ism)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocalInterchainAccount0 is a free data retrieval call binding the contract method 0xa5da3611.
//
// Solidity: function getLocalInterchainAccount(uint32 _origin, bytes32 _owner, bytes32 _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterSession) GetLocalInterchainAccount0(_origin uint32, _owner [32]byte, _router [32]byte, _ism common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetLocalInterchainAccount0(&_Accountrouter.CallOpts, _origin, _owner, _router, _ism)
}

// GetLocalInterchainAccount0 is a free data retrieval call binding the contract method 0xa5da3611.
//
// Solidity: function getLocalInterchainAccount(uint32 _origin, bytes32 _owner, bytes32 _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterCallerSession) GetLocalInterchainAccount0(_origin uint32, _owner [32]byte, _router [32]byte, _ism common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetLocalInterchainAccount0(&_Accountrouter.CallOpts, _origin, _owner, _router, _ism)
}

// GetRemoteInterchainAccount is a free data retrieval call binding the contract method 0x0e258285.
//
// Solidity: function getRemoteInterchainAccount(uint32 _destination, address _owner) view returns(address)
func (_Accountrouter *AccountrouterCaller) GetRemoteInterchainAccount(opts *bind.CallOpts, _destination uint32, _owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "getRemoteInterchainAccount", _destination, _owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRemoteInterchainAccount is a free data retrieval call binding the contract method 0x0e258285.
//
// Solidity: function getRemoteInterchainAccount(uint32 _destination, address _owner) view returns(address)
func (_Accountrouter *AccountrouterSession) GetRemoteInterchainAccount(_destination uint32, _owner common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetRemoteInterchainAccount(&_Accountrouter.CallOpts, _destination, _owner)
}

// GetRemoteInterchainAccount is a free data retrieval call binding the contract method 0x0e258285.
//
// Solidity: function getRemoteInterchainAccount(uint32 _destination, address _owner) view returns(address)
func (_Accountrouter *AccountrouterCallerSession) GetRemoteInterchainAccount(_destination uint32, _owner common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetRemoteInterchainAccount(&_Accountrouter.CallOpts, _destination, _owner)
}

// GetRemoteInterchainAccount0 is a free data retrieval call binding the contract method 0xef96bb4b.
//
// Solidity: function getRemoteInterchainAccount(address _owner, address _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterCaller) GetRemoteInterchainAccount0(opts *bind.CallOpts, _owner common.Address, _router common.Address, _ism common.Address) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "getRemoteInterchainAccount0", _owner, _router, _ism)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRemoteInterchainAccount0 is a free data retrieval call binding the contract method 0xef96bb4b.
//
// Solidity: function getRemoteInterchainAccount(address _owner, address _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterSession) GetRemoteInterchainAccount0(_owner common.Address, _router common.Address, _ism common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetRemoteInterchainAccount0(&_Accountrouter.CallOpts, _owner, _router, _ism)
}

// GetRemoteInterchainAccount0 is a free data retrieval call binding the contract method 0xef96bb4b.
//
// Solidity: function getRemoteInterchainAccount(address _owner, address _router, address _ism) view returns(address)
func (_Accountrouter *AccountrouterCallerSession) GetRemoteInterchainAccount0(_owner common.Address, _router common.Address, _ism common.Address) (common.Address, error) {
	return _Accountrouter.Contract.GetRemoteInterchainAccount0(&_Accountrouter.CallOpts, _owner, _router, _ism)
}

// InterchainGasPaymaster is a free data retrieval call binding the contract method 0x39bb4ad9.
//
// Solidity: function interchainGasPaymaster() view returns(address)
func (_Accountrouter *AccountrouterCaller) InterchainGasPaymaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "interchainGasPaymaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainGasPaymaster is a free data retrieval call binding the contract method 0x39bb4ad9.
//
// Solidity: function interchainGasPaymaster() view returns(address)
func (_Accountrouter *AccountrouterSession) InterchainGasPaymaster() (common.Address, error) {
	return _Accountrouter.Contract.InterchainGasPaymaster(&_Accountrouter.CallOpts)
}

// InterchainGasPaymaster is a free data retrieval call binding the contract method 0x39bb4ad9.
//
// Solidity: function interchainGasPaymaster() view returns(address)
func (_Accountrouter *AccountrouterCallerSession) InterchainGasPaymaster() (common.Address, error) {
	return _Accountrouter.Contract.InterchainGasPaymaster(&_Accountrouter.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_Accountrouter *AccountrouterCaller) InterchainSecurityModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "interchainSecurityModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_Accountrouter *AccountrouterSession) InterchainSecurityModule() (common.Address, error) {
	return _Accountrouter.Contract.InterchainSecurityModule(&_Accountrouter.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_Accountrouter *AccountrouterCallerSession) InterchainSecurityModule() (common.Address, error) {
	return _Accountrouter.Contract.InterchainSecurityModule(&_Accountrouter.CallOpts)
}

// Isms is a free data retrieval call binding the contract method 0x5a007b7e.
//
// Solidity: function isms(uint32 ) view returns(bytes32)
func (_Accountrouter *AccountrouterCaller) Isms(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "isms", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Isms is a free data retrieval call binding the contract method 0x5a007b7e.
//
// Solidity: function isms(uint32 ) view returns(bytes32)
func (_Accountrouter *AccountrouterSession) Isms(arg0 uint32) ([32]byte, error) {
	return _Accountrouter.Contract.Isms(&_Accountrouter.CallOpts, arg0)
}

// Isms is a free data retrieval call binding the contract method 0x5a007b7e.
//
// Solidity: function isms(uint32 ) view returns(bytes32)
func (_Accountrouter *AccountrouterCallerSession) Isms(arg0 uint32) ([32]byte, error) {
	return _Accountrouter.Contract.Isms(&_Accountrouter.CallOpts, arg0)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_Accountrouter *AccountrouterCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_Accountrouter *AccountrouterSession) Mailbox() (common.Address, error) {
	return _Accountrouter.Contract.Mailbox(&_Accountrouter.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_Accountrouter *AccountrouterCallerSession) Mailbox() (common.Address, error) {
	return _Accountrouter.Contract.Mailbox(&_Accountrouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accountrouter *AccountrouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accountrouter *AccountrouterSession) Owner() (common.Address, error) {
	return _Accountrouter.Contract.Owner(&_Accountrouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Accountrouter *AccountrouterCallerSession) Owner() (common.Address, error) {
	return _Accountrouter.Contract.Owner(&_Accountrouter.CallOpts)
}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 ) view returns(bytes32)
func (_Accountrouter *AccountrouterCaller) Routers(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Accountrouter.contract.Call(opts, &out, "routers", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 ) view returns(bytes32)
func (_Accountrouter *AccountrouterSession) Routers(arg0 uint32) ([32]byte, error) {
	return _Accountrouter.Contract.Routers(&_Accountrouter.CallOpts, arg0)
}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 ) view returns(bytes32)
func (_Accountrouter *AccountrouterCallerSession) Routers(arg0 uint32) ([32]byte, error) {
	return _Accountrouter.Contract.Routers(&_Accountrouter.CallOpts, arg0)
}

// CallRemote is a paid mutator transaction binding the contract method 0x1ddce514.
//
// Solidity: function callRemote(uint32 _destination, address _to, uint256 _value, bytes _data) returns(bytes32)
func (_Accountrouter *AccountrouterTransactor) CallRemote(opts *bind.TransactOpts, _destination uint32, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "callRemote", _destination, _to, _value, _data)
}

// CallRemote is a paid mutator transaction binding the contract method 0x1ddce514.
//
// Solidity: function callRemote(uint32 _destination, address _to, uint256 _value, bytes _data) returns(bytes32)
func (_Accountrouter *AccountrouterSession) CallRemote(_destination uint32, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.CallRemote(&_Accountrouter.TransactOpts, _destination, _to, _value, _data)
}

// CallRemote is a paid mutator transaction binding the contract method 0x1ddce514.
//
// Solidity: function callRemote(uint32 _destination, address _to, uint256 _value, bytes _data) returns(bytes32)
func (_Accountrouter *AccountrouterTransactorSession) CallRemote(_destination uint32, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.CallRemote(&_Accountrouter.TransactOpts, _destination, _to, _value, _data)
}

// CallRemote0 is a paid mutator transaction binding the contract method 0xdd91cc87.
//
// Solidity: function callRemote(uint32 _destination, (bytes32,uint256,bytes)[] _calls) returns(bytes32)
func (_Accountrouter *AccountrouterTransactor) CallRemote0(opts *bind.TransactOpts, _destination uint32, _calls []CallLibCall) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "callRemote0", _destination, _calls)
}

// CallRemote0 is a paid mutator transaction binding the contract method 0xdd91cc87.
//
// Solidity: function callRemote(uint32 _destination, (bytes32,uint256,bytes)[] _calls) returns(bytes32)
func (_Accountrouter *AccountrouterSession) CallRemote0(_destination uint32, _calls []CallLibCall) (*types.Transaction, error) {
	return _Accountrouter.Contract.CallRemote0(&_Accountrouter.TransactOpts, _destination, _calls)
}

// CallRemote0 is a paid mutator transaction binding the contract method 0xdd91cc87.
//
// Solidity: function callRemote(uint32 _destination, (bytes32,uint256,bytes)[] _calls) returns(bytes32)
func (_Accountrouter *AccountrouterTransactorSession) CallRemote0(_destination uint32, _calls []CallLibCall) (*types.Transaction, error) {
	return _Accountrouter.Contract.CallRemote0(&_Accountrouter.TransactOpts, _destination, _calls)
}

// CallRemoteWithOverrides is a paid mutator transaction binding the contract method 0x061197d5.
//
// Solidity: function callRemoteWithOverrides(uint32 _destination, bytes32 _router, bytes32 _ism, (bytes32,uint256,bytes)[] _calls) returns(bytes32)
func (_Accountrouter *AccountrouterTransactor) CallRemoteWithOverrides(opts *bind.TransactOpts, _destination uint32, _router [32]byte, _ism [32]byte, _calls []CallLibCall) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "callRemoteWithOverrides", _destination, _router, _ism, _calls)
}

// CallRemoteWithOverrides is a paid mutator transaction binding the contract method 0x061197d5.
//
// Solidity: function callRemoteWithOverrides(uint32 _destination, bytes32 _router, bytes32 _ism, (bytes32,uint256,bytes)[] _calls) returns(bytes32)
func (_Accountrouter *AccountrouterSession) CallRemoteWithOverrides(_destination uint32, _router [32]byte, _ism [32]byte, _calls []CallLibCall) (*types.Transaction, error) {
	return _Accountrouter.Contract.CallRemoteWithOverrides(&_Accountrouter.TransactOpts, _destination, _router, _ism, _calls)
}

// CallRemoteWithOverrides is a paid mutator transaction binding the contract method 0x061197d5.
//
// Solidity: function callRemoteWithOverrides(uint32 _destination, bytes32 _router, bytes32 _ism, (bytes32,uint256,bytes)[] _calls) returns(bytes32)
func (_Accountrouter *AccountrouterTransactorSession) CallRemoteWithOverrides(_destination uint32, _router [32]byte, _ism [32]byte, _calls []CallLibCall) (*types.Transaction, error) {
	return _Accountrouter.Contract.CallRemoteWithOverrides(&_Accountrouter.TransactOpts, _destination, _router, _ism, _calls)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _destination, bytes32 _router) returns()
func (_Accountrouter *AccountrouterTransactor) EnrollRemoteRouter(opts *bind.TransactOpts, _destination uint32, _router [32]byte) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "enrollRemoteRouter", _destination, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _destination, bytes32 _router) returns()
func (_Accountrouter *AccountrouterSession) EnrollRemoteRouter(_destination uint32, _router [32]byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.EnrollRemoteRouter(&_Accountrouter.TransactOpts, _destination, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _destination, bytes32 _router) returns()
func (_Accountrouter *AccountrouterTransactorSession) EnrollRemoteRouter(_destination uint32, _router [32]byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.EnrollRemoteRouter(&_Accountrouter.TransactOpts, _destination, _router)
}

// EnrollRemoteRouterAndIsm is a paid mutator transaction binding the contract method 0x09643ff7.
//
// Solidity: function enrollRemoteRouterAndIsm(uint32 _destination, bytes32 _router, bytes32 _ism) returns()
func (_Accountrouter *AccountrouterTransactor) EnrollRemoteRouterAndIsm(opts *bind.TransactOpts, _destination uint32, _router [32]byte, _ism [32]byte) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "enrollRemoteRouterAndIsm", _destination, _router, _ism)
}

// EnrollRemoteRouterAndIsm is a paid mutator transaction binding the contract method 0x09643ff7.
//
// Solidity: function enrollRemoteRouterAndIsm(uint32 _destination, bytes32 _router, bytes32 _ism) returns()
func (_Accountrouter *AccountrouterSession) EnrollRemoteRouterAndIsm(_destination uint32, _router [32]byte, _ism [32]byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.EnrollRemoteRouterAndIsm(&_Accountrouter.TransactOpts, _destination, _router, _ism)
}

// EnrollRemoteRouterAndIsm is a paid mutator transaction binding the contract method 0x09643ff7.
//
// Solidity: function enrollRemoteRouterAndIsm(uint32 _destination, bytes32 _router, bytes32 _ism) returns()
func (_Accountrouter *AccountrouterTransactorSession) EnrollRemoteRouterAndIsm(_destination uint32, _router [32]byte, _ism [32]byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.EnrollRemoteRouterAndIsm(&_Accountrouter.TransactOpts, _destination, _router, _ism)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _destinations, bytes32[] _routers) returns()
func (_Accountrouter *AccountrouterTransactor) EnrollRemoteRouters(opts *bind.TransactOpts, _destinations []uint32, _routers [][32]byte) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "enrollRemoteRouters", _destinations, _routers)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _destinations, bytes32[] _routers) returns()
func (_Accountrouter *AccountrouterSession) EnrollRemoteRouters(_destinations []uint32, _routers [][32]byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.EnrollRemoteRouters(&_Accountrouter.TransactOpts, _destinations, _routers)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _destinations, bytes32[] _routers) returns()
func (_Accountrouter *AccountrouterTransactorSession) EnrollRemoteRouters(_destinations []uint32, _routers [][32]byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.EnrollRemoteRouters(&_Accountrouter.TransactOpts, _destinations, _routers)
}

// GetDeployedInterchainAccount is a paid mutator transaction binding the contract method 0x8e45f8fc.
//
// Solidity: function getDeployedInterchainAccount(uint32 _origin, address _owner, address _router, address _ism) returns(address)
func (_Accountrouter *AccountrouterTransactor) GetDeployedInterchainAccount(opts *bind.TransactOpts, _origin uint32, _owner common.Address, _router common.Address, _ism common.Address) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "getDeployedInterchainAccount", _origin, _owner, _router, _ism)
}

// GetDeployedInterchainAccount is a paid mutator transaction binding the contract method 0x8e45f8fc.
//
// Solidity: function getDeployedInterchainAccount(uint32 _origin, address _owner, address _router, address _ism) returns(address)
func (_Accountrouter *AccountrouterSession) GetDeployedInterchainAccount(_origin uint32, _owner common.Address, _router common.Address, _ism common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.GetDeployedInterchainAccount(&_Accountrouter.TransactOpts, _origin, _owner, _router, _ism)
}

// GetDeployedInterchainAccount is a paid mutator transaction binding the contract method 0x8e45f8fc.
//
// Solidity: function getDeployedInterchainAccount(uint32 _origin, address _owner, address _router, address _ism) returns(address)
func (_Accountrouter *AccountrouterTransactorSession) GetDeployedInterchainAccount(_origin uint32, _owner common.Address, _router common.Address, _ism common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.GetDeployedInterchainAccount(&_Accountrouter.TransactOpts, _origin, _owner, _router, _ism)
}

// GetDeployedInterchainAccount0 is a paid mutator transaction binding the contract method 0x940c14dd.
//
// Solidity: function getDeployedInterchainAccount(uint32 _origin, bytes32 _owner, bytes32 _router, address _ism) returns(address)
func (_Accountrouter *AccountrouterTransactor) GetDeployedInterchainAccount0(opts *bind.TransactOpts, _origin uint32, _owner [32]byte, _router [32]byte, _ism common.Address) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "getDeployedInterchainAccount0", _origin, _owner, _router, _ism)
}

// GetDeployedInterchainAccount0 is a paid mutator transaction binding the contract method 0x940c14dd.
//
// Solidity: function getDeployedInterchainAccount(uint32 _origin, bytes32 _owner, bytes32 _router, address _ism) returns(address)
func (_Accountrouter *AccountrouterSession) GetDeployedInterchainAccount0(_origin uint32, _owner [32]byte, _router [32]byte, _ism common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.GetDeployedInterchainAccount0(&_Accountrouter.TransactOpts, _origin, _owner, _router, _ism)
}

// GetDeployedInterchainAccount0 is a paid mutator transaction binding the contract method 0x940c14dd.
//
// Solidity: function getDeployedInterchainAccount(uint32 _origin, bytes32 _owner, bytes32 _router, address _ism) returns(address)
func (_Accountrouter *AccountrouterTransactorSession) GetDeployedInterchainAccount0(_origin uint32, _owner [32]byte, _router [32]byte, _ism common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.GetDeployedInterchainAccount0(&_Accountrouter.TransactOpts, _origin, _owner, _router, _ism)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) returns()
func (_Accountrouter *AccountrouterTransactor) Handle(opts *bind.TransactOpts, _origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "handle", _origin, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) returns()
func (_Accountrouter *AccountrouterSession) Handle(_origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.Handle(&_Accountrouter.TransactOpts, _origin, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) returns()
func (_Accountrouter *AccountrouterTransactorSession) Handle(_origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _Accountrouter.Contract.Handle(&_Accountrouter.TransactOpts, _origin, _sender, _message)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _mailbox, address _interchainGasPaymaster, address _interchainSecurityModule, address _owner) returns()
func (_Accountrouter *AccountrouterTransactor) Initialize(opts *bind.TransactOpts, _mailbox common.Address, _interchainGasPaymaster common.Address, _interchainSecurityModule common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "initialize", _mailbox, _interchainGasPaymaster, _interchainSecurityModule, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _mailbox, address _interchainGasPaymaster, address _interchainSecurityModule, address _owner) returns()
func (_Accountrouter *AccountrouterSession) Initialize(_mailbox common.Address, _interchainGasPaymaster common.Address, _interchainSecurityModule common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.Initialize(&_Accountrouter.TransactOpts, _mailbox, _interchainGasPaymaster, _interchainSecurityModule, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _mailbox, address _interchainGasPaymaster, address _interchainSecurityModule, address _owner) returns()
func (_Accountrouter *AccountrouterTransactorSession) Initialize(_mailbox common.Address, _interchainGasPaymaster common.Address, _interchainSecurityModule common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.Initialize(&_Accountrouter.TransactOpts, _mailbox, _interchainGasPaymaster, _interchainSecurityModule, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accountrouter *AccountrouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accountrouter *AccountrouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accountrouter.Contract.RenounceOwnership(&_Accountrouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accountrouter *AccountrouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accountrouter.Contract.RenounceOwnership(&_Accountrouter.TransactOpts)
}

// SetInterchainGasPaymaster is a paid mutator transaction binding the contract method 0xf1bd6f0a.
//
// Solidity: function setInterchainGasPaymaster(address _interchainGasPaymaster) returns()
func (_Accountrouter *AccountrouterTransactor) SetInterchainGasPaymaster(opts *bind.TransactOpts, _interchainGasPaymaster common.Address) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "setInterchainGasPaymaster", _interchainGasPaymaster)
}

// SetInterchainGasPaymaster is a paid mutator transaction binding the contract method 0xf1bd6f0a.
//
// Solidity: function setInterchainGasPaymaster(address _interchainGasPaymaster) returns()
func (_Accountrouter *AccountrouterSession) SetInterchainGasPaymaster(_interchainGasPaymaster common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.SetInterchainGasPaymaster(&_Accountrouter.TransactOpts, _interchainGasPaymaster)
}

// SetInterchainGasPaymaster is a paid mutator transaction binding the contract method 0xf1bd6f0a.
//
// Solidity: function setInterchainGasPaymaster(address _interchainGasPaymaster) returns()
func (_Accountrouter *AccountrouterTransactorSession) SetInterchainGasPaymaster(_interchainGasPaymaster common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.SetInterchainGasPaymaster(&_Accountrouter.TransactOpts, _interchainGasPaymaster)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_Accountrouter *AccountrouterTransactor) SetInterchainSecurityModule(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "setInterchainSecurityModule", _module)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_Accountrouter *AccountrouterSession) SetInterchainSecurityModule(_module common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.SetInterchainSecurityModule(&_Accountrouter.TransactOpts, _module)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_Accountrouter *AccountrouterTransactorSession) SetInterchainSecurityModule(_module common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.SetInterchainSecurityModule(&_Accountrouter.TransactOpts, _module)
}

// SetMailbox is a paid mutator transaction binding the contract method 0xf3c61d6b.
//
// Solidity: function setMailbox(address _mailbox) returns()
func (_Accountrouter *AccountrouterTransactor) SetMailbox(opts *bind.TransactOpts, _mailbox common.Address) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "setMailbox", _mailbox)
}

// SetMailbox is a paid mutator transaction binding the contract method 0xf3c61d6b.
//
// Solidity: function setMailbox(address _mailbox) returns()
func (_Accountrouter *AccountrouterSession) SetMailbox(_mailbox common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.SetMailbox(&_Accountrouter.TransactOpts, _mailbox)
}

// SetMailbox is a paid mutator transaction binding the contract method 0xf3c61d6b.
//
// Solidity: function setMailbox(address _mailbox) returns()
func (_Accountrouter *AccountrouterTransactorSession) SetMailbox(_mailbox common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.SetMailbox(&_Accountrouter.TransactOpts, _mailbox)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accountrouter *AccountrouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Accountrouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accountrouter *AccountrouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.TransferOwnership(&_Accountrouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accountrouter *AccountrouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accountrouter.Contract.TransferOwnership(&_Accountrouter.TransactOpts, newOwner)
}

// AccountrouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Accountrouter contract.
type AccountrouterInitializedIterator struct {
	Event *AccountrouterInitialized // Event containing the contract specifics and raw log

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
func (it *AccountrouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterInitialized)
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
		it.Event = new(AccountrouterInitialized)
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
func (it *AccountrouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterInitialized represents a Initialized event raised by the Accountrouter contract.
type AccountrouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Accountrouter *AccountrouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountrouterInitializedIterator, error) {

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountrouterInitializedIterator{contract: _Accountrouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Accountrouter *AccountrouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountrouterInitialized) (event.Subscription, error) {

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterInitialized)
				if err := _Accountrouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Accountrouter *AccountrouterFilterer) ParseInitialized(log types.Log) (*AccountrouterInitialized, error) {
	event := new(AccountrouterInitialized)
	if err := _Accountrouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterInterchainAccountCreatedIterator is returned from FilterInterchainAccountCreated and is used to iterate over the raw logs and unpacked data for InterchainAccountCreated events raised by the Accountrouter contract.
type AccountrouterInterchainAccountCreatedIterator struct {
	Event *AccountrouterInterchainAccountCreated // Event containing the contract specifics and raw log

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
func (it *AccountrouterInterchainAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterInterchainAccountCreated)
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
		it.Event = new(AccountrouterInterchainAccountCreated)
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
func (it *AccountrouterInterchainAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterInterchainAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterInterchainAccountCreated represents a InterchainAccountCreated event raised by the Accountrouter contract.
type AccountrouterInterchainAccountCreated struct {
	Origin  uint32
	Owner   [32]byte
	Ism     common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInterchainAccountCreated is a free log retrieval operation binding the contract event 0xee33e8a4a721d99ee890b32b36b48bc5a14d276b7e2510763a52f839d178edf6.
//
// Solidity: event InterchainAccountCreated(uint32 indexed origin, bytes32 indexed owner, address ism, address account)
func (_Accountrouter *AccountrouterFilterer) FilterInterchainAccountCreated(opts *bind.FilterOpts, origin []uint32, owner [][32]byte) (*AccountrouterInterchainAccountCreatedIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "InterchainAccountCreated", originRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterInterchainAccountCreatedIterator{contract: _Accountrouter.contract, event: "InterchainAccountCreated", logs: logs, sub: sub}, nil
}

// WatchInterchainAccountCreated is a free log subscription operation binding the contract event 0xee33e8a4a721d99ee890b32b36b48bc5a14d276b7e2510763a52f839d178edf6.
//
// Solidity: event InterchainAccountCreated(uint32 indexed origin, bytes32 indexed owner, address ism, address account)
func (_Accountrouter *AccountrouterFilterer) WatchInterchainAccountCreated(opts *bind.WatchOpts, sink chan<- *AccountrouterInterchainAccountCreated, origin []uint32, owner [][32]byte) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "InterchainAccountCreated", originRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterInterchainAccountCreated)
				if err := _Accountrouter.contract.UnpackLog(event, "InterchainAccountCreated", log); err != nil {
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

// ParseInterchainAccountCreated is a log parse operation binding the contract event 0xee33e8a4a721d99ee890b32b36b48bc5a14d276b7e2510763a52f839d178edf6.
//
// Solidity: event InterchainAccountCreated(uint32 indexed origin, bytes32 indexed owner, address ism, address account)
func (_Accountrouter *AccountrouterFilterer) ParseInterchainAccountCreated(log types.Log) (*AccountrouterInterchainAccountCreated, error) {
	event := new(AccountrouterInterchainAccountCreated)
	if err := _Accountrouter.contract.UnpackLog(event, "InterchainAccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterInterchainGasPaymasterSetIterator is returned from FilterInterchainGasPaymasterSet and is used to iterate over the raw logs and unpacked data for InterchainGasPaymasterSet events raised by the Accountrouter contract.
type AccountrouterInterchainGasPaymasterSetIterator struct {
	Event *AccountrouterInterchainGasPaymasterSet // Event containing the contract specifics and raw log

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
func (it *AccountrouterInterchainGasPaymasterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterInterchainGasPaymasterSet)
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
		it.Event = new(AccountrouterInterchainGasPaymasterSet)
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
func (it *AccountrouterInterchainGasPaymasterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterInterchainGasPaymasterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterInterchainGasPaymasterSet represents a InterchainGasPaymasterSet event raised by the Accountrouter contract.
type AccountrouterInterchainGasPaymasterSet struct {
	InterchainGasPaymaster common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterInterchainGasPaymasterSet is a free log retrieval operation binding the contract event 0xb87f5a0bface22cde699143559fb99f990ae54eea236995fc54178af91af859f.
//
// Solidity: event InterchainGasPaymasterSet(address indexed interchainGasPaymaster)
func (_Accountrouter *AccountrouterFilterer) FilterInterchainGasPaymasterSet(opts *bind.FilterOpts, interchainGasPaymaster []common.Address) (*AccountrouterInterchainGasPaymasterSetIterator, error) {

	var interchainGasPaymasterRule []interface{}
	for _, interchainGasPaymasterItem := range interchainGasPaymaster {
		interchainGasPaymasterRule = append(interchainGasPaymasterRule, interchainGasPaymasterItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "InterchainGasPaymasterSet", interchainGasPaymasterRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterInterchainGasPaymasterSetIterator{contract: _Accountrouter.contract, event: "InterchainGasPaymasterSet", logs: logs, sub: sub}, nil
}

// WatchInterchainGasPaymasterSet is a free log subscription operation binding the contract event 0xb87f5a0bface22cde699143559fb99f990ae54eea236995fc54178af91af859f.
//
// Solidity: event InterchainGasPaymasterSet(address indexed interchainGasPaymaster)
func (_Accountrouter *AccountrouterFilterer) WatchInterchainGasPaymasterSet(opts *bind.WatchOpts, sink chan<- *AccountrouterInterchainGasPaymasterSet, interchainGasPaymaster []common.Address) (event.Subscription, error) {

	var interchainGasPaymasterRule []interface{}
	for _, interchainGasPaymasterItem := range interchainGasPaymaster {
		interchainGasPaymasterRule = append(interchainGasPaymasterRule, interchainGasPaymasterItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "InterchainGasPaymasterSet", interchainGasPaymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterInterchainGasPaymasterSet)
				if err := _Accountrouter.contract.UnpackLog(event, "InterchainGasPaymasterSet", log); err != nil {
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

// ParseInterchainGasPaymasterSet is a log parse operation binding the contract event 0xb87f5a0bface22cde699143559fb99f990ae54eea236995fc54178af91af859f.
//
// Solidity: event InterchainGasPaymasterSet(address indexed interchainGasPaymaster)
func (_Accountrouter *AccountrouterFilterer) ParseInterchainGasPaymasterSet(log types.Log) (*AccountrouterInterchainGasPaymasterSet, error) {
	event := new(AccountrouterInterchainGasPaymasterSet)
	if err := _Accountrouter.contract.UnpackLog(event, "InterchainGasPaymasterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterInterchainSecurityModuleSetIterator is returned from FilterInterchainSecurityModuleSet and is used to iterate over the raw logs and unpacked data for InterchainSecurityModuleSet events raised by the Accountrouter contract.
type AccountrouterInterchainSecurityModuleSetIterator struct {
	Event *AccountrouterInterchainSecurityModuleSet // Event containing the contract specifics and raw log

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
func (it *AccountrouterInterchainSecurityModuleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterInterchainSecurityModuleSet)
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
		it.Event = new(AccountrouterInterchainSecurityModuleSet)
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
func (it *AccountrouterInterchainSecurityModuleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterInterchainSecurityModuleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterInterchainSecurityModuleSet represents a InterchainSecurityModuleSet event raised by the Accountrouter contract.
type AccountrouterInterchainSecurityModuleSet struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainSecurityModuleSet is a free log retrieval operation binding the contract event 0xfec811ed4e60aebdaf7a79cad8a97196bf56e35362f039705598226d30c9d248.
//
// Solidity: event InterchainSecurityModuleSet(address indexed module)
func (_Accountrouter *AccountrouterFilterer) FilterInterchainSecurityModuleSet(opts *bind.FilterOpts, module []common.Address) (*AccountrouterInterchainSecurityModuleSetIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "InterchainSecurityModuleSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterInterchainSecurityModuleSetIterator{contract: _Accountrouter.contract, event: "InterchainSecurityModuleSet", logs: logs, sub: sub}, nil
}

// WatchInterchainSecurityModuleSet is a free log subscription operation binding the contract event 0xfec811ed4e60aebdaf7a79cad8a97196bf56e35362f039705598226d30c9d248.
//
// Solidity: event InterchainSecurityModuleSet(address indexed module)
func (_Accountrouter *AccountrouterFilterer) WatchInterchainSecurityModuleSet(opts *bind.WatchOpts, sink chan<- *AccountrouterInterchainSecurityModuleSet, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "InterchainSecurityModuleSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterInterchainSecurityModuleSet)
				if err := _Accountrouter.contract.UnpackLog(event, "InterchainSecurityModuleSet", log); err != nil {
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

// ParseInterchainSecurityModuleSet is a log parse operation binding the contract event 0xfec811ed4e60aebdaf7a79cad8a97196bf56e35362f039705598226d30c9d248.
//
// Solidity: event InterchainSecurityModuleSet(address indexed module)
func (_Accountrouter *AccountrouterFilterer) ParseInterchainSecurityModuleSet(log types.Log) (*AccountrouterInterchainSecurityModuleSet, error) {
	event := new(AccountrouterInterchainSecurityModuleSet)
	if err := _Accountrouter.contract.UnpackLog(event, "InterchainSecurityModuleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterMailboxSetIterator is returned from FilterMailboxSet and is used to iterate over the raw logs and unpacked data for MailboxSet events raised by the Accountrouter contract.
type AccountrouterMailboxSetIterator struct {
	Event *AccountrouterMailboxSet // Event containing the contract specifics and raw log

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
func (it *AccountrouterMailboxSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterMailboxSet)
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
		it.Event = new(AccountrouterMailboxSet)
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
func (it *AccountrouterMailboxSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterMailboxSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterMailboxSet represents a MailboxSet event raised by the Accountrouter contract.
type AccountrouterMailboxSet struct {
	Mailbox common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMailboxSet is a free log retrieval operation binding the contract event 0x7a61e573722ff8b205c8962b59d37e7d30573f368965597a844a94872204ebd7.
//
// Solidity: event MailboxSet(address indexed mailbox)
func (_Accountrouter *AccountrouterFilterer) FilterMailboxSet(opts *bind.FilterOpts, mailbox []common.Address) (*AccountrouterMailboxSetIterator, error) {

	var mailboxRule []interface{}
	for _, mailboxItem := range mailbox {
		mailboxRule = append(mailboxRule, mailboxItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "MailboxSet", mailboxRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterMailboxSetIterator{contract: _Accountrouter.contract, event: "MailboxSet", logs: logs, sub: sub}, nil
}

// WatchMailboxSet is a free log subscription operation binding the contract event 0x7a61e573722ff8b205c8962b59d37e7d30573f368965597a844a94872204ebd7.
//
// Solidity: event MailboxSet(address indexed mailbox)
func (_Accountrouter *AccountrouterFilterer) WatchMailboxSet(opts *bind.WatchOpts, sink chan<- *AccountrouterMailboxSet, mailbox []common.Address) (event.Subscription, error) {

	var mailboxRule []interface{}
	for _, mailboxItem := range mailbox {
		mailboxRule = append(mailboxRule, mailboxItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "MailboxSet", mailboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterMailboxSet)
				if err := _Accountrouter.contract.UnpackLog(event, "MailboxSet", log); err != nil {
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

// ParseMailboxSet is a log parse operation binding the contract event 0x7a61e573722ff8b205c8962b59d37e7d30573f368965597a844a94872204ebd7.
//
// Solidity: event MailboxSet(address indexed mailbox)
func (_Accountrouter *AccountrouterFilterer) ParseMailboxSet(log types.Log) (*AccountrouterMailboxSet, error) {
	event := new(AccountrouterMailboxSet)
	if err := _Accountrouter.contract.UnpackLog(event, "MailboxSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Accountrouter contract.
type AccountrouterOwnershipTransferredIterator struct {
	Event *AccountrouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountrouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterOwnershipTransferred)
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
		it.Event = new(AccountrouterOwnershipTransferred)
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
func (it *AccountrouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterOwnershipTransferred represents a OwnershipTransferred event raised by the Accountrouter contract.
type AccountrouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accountrouter *AccountrouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountrouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterOwnershipTransferredIterator{contract: _Accountrouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accountrouter *AccountrouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountrouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterOwnershipTransferred)
				if err := _Accountrouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Accountrouter *AccountrouterFilterer) ParseOwnershipTransferred(log types.Log) (*AccountrouterOwnershipTransferred, error) {
	event := new(AccountrouterOwnershipTransferred)
	if err := _Accountrouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterRemoteCallDispatchedIterator is returned from FilterRemoteCallDispatched and is used to iterate over the raw logs and unpacked data for RemoteCallDispatched events raised by the Accountrouter contract.
type AccountrouterRemoteCallDispatchedIterator struct {
	Event *AccountrouterRemoteCallDispatched // Event containing the contract specifics and raw log

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
func (it *AccountrouterRemoteCallDispatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterRemoteCallDispatched)
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
		it.Event = new(AccountrouterRemoteCallDispatched)
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
func (it *AccountrouterRemoteCallDispatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterRemoteCallDispatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterRemoteCallDispatched represents a RemoteCallDispatched event raised by the Accountrouter contract.
type AccountrouterRemoteCallDispatched struct {
	Destination uint32
	Owner       common.Address
	Router      [32]byte
	Ism         [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoteCallDispatched is a free log retrieval operation binding the contract event 0x4fd2d122031a8a1e6c1aed1e9128109dff5ce6b929edfa90cc52abd95f086e73.
//
// Solidity: event RemoteCallDispatched(uint32 indexed destination, address indexed owner, bytes32 router, bytes32 ism)
func (_Accountrouter *AccountrouterFilterer) FilterRemoteCallDispatched(opts *bind.FilterOpts, destination []uint32, owner []common.Address) (*AccountrouterRemoteCallDispatchedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "RemoteCallDispatched", destinationRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterRemoteCallDispatchedIterator{contract: _Accountrouter.contract, event: "RemoteCallDispatched", logs: logs, sub: sub}, nil
}

// WatchRemoteCallDispatched is a free log subscription operation binding the contract event 0x4fd2d122031a8a1e6c1aed1e9128109dff5ce6b929edfa90cc52abd95f086e73.
//
// Solidity: event RemoteCallDispatched(uint32 indexed destination, address indexed owner, bytes32 router, bytes32 ism)
func (_Accountrouter *AccountrouterFilterer) WatchRemoteCallDispatched(opts *bind.WatchOpts, sink chan<- *AccountrouterRemoteCallDispatched, destination []uint32, owner []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "RemoteCallDispatched", destinationRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterRemoteCallDispatched)
				if err := _Accountrouter.contract.UnpackLog(event, "RemoteCallDispatched", log); err != nil {
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

// ParseRemoteCallDispatched is a log parse operation binding the contract event 0x4fd2d122031a8a1e6c1aed1e9128109dff5ce6b929edfa90cc52abd95f086e73.
//
// Solidity: event RemoteCallDispatched(uint32 indexed destination, address indexed owner, bytes32 router, bytes32 ism)
func (_Accountrouter *AccountrouterFilterer) ParseRemoteCallDispatched(log types.Log) (*AccountrouterRemoteCallDispatched, error) {
	event := new(AccountrouterRemoteCallDispatched)
	if err := _Accountrouter.contract.UnpackLog(event, "RemoteCallDispatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterRemoteIsmEnrolledIterator is returned from FilterRemoteIsmEnrolled and is used to iterate over the raw logs and unpacked data for RemoteIsmEnrolled events raised by the Accountrouter contract.
type AccountrouterRemoteIsmEnrolledIterator struct {
	Event *AccountrouterRemoteIsmEnrolled // Event containing the contract specifics and raw log

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
func (it *AccountrouterRemoteIsmEnrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterRemoteIsmEnrolled)
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
		it.Event = new(AccountrouterRemoteIsmEnrolled)
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
func (it *AccountrouterRemoteIsmEnrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterRemoteIsmEnrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterRemoteIsmEnrolled represents a RemoteIsmEnrolled event raised by the Accountrouter contract.
type AccountrouterRemoteIsmEnrolled struct {
	Domain uint32
	Ism    [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoteIsmEnrolled is a free log retrieval operation binding the contract event 0xba9685b280b1e8f64909e2dfd7465463e6a0eabbaafd732402e7a619699a3821.
//
// Solidity: event RemoteIsmEnrolled(uint32 indexed domain, bytes32 ism)
func (_Accountrouter *AccountrouterFilterer) FilterRemoteIsmEnrolled(opts *bind.FilterOpts, domain []uint32) (*AccountrouterRemoteIsmEnrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "RemoteIsmEnrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterRemoteIsmEnrolledIterator{contract: _Accountrouter.contract, event: "RemoteIsmEnrolled", logs: logs, sub: sub}, nil
}

// WatchRemoteIsmEnrolled is a free log subscription operation binding the contract event 0xba9685b280b1e8f64909e2dfd7465463e6a0eabbaafd732402e7a619699a3821.
//
// Solidity: event RemoteIsmEnrolled(uint32 indexed domain, bytes32 ism)
func (_Accountrouter *AccountrouterFilterer) WatchRemoteIsmEnrolled(opts *bind.WatchOpts, sink chan<- *AccountrouterRemoteIsmEnrolled, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "RemoteIsmEnrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterRemoteIsmEnrolled)
				if err := _Accountrouter.contract.UnpackLog(event, "RemoteIsmEnrolled", log); err != nil {
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

// ParseRemoteIsmEnrolled is a log parse operation binding the contract event 0xba9685b280b1e8f64909e2dfd7465463e6a0eabbaafd732402e7a619699a3821.
//
// Solidity: event RemoteIsmEnrolled(uint32 indexed domain, bytes32 ism)
func (_Accountrouter *AccountrouterFilterer) ParseRemoteIsmEnrolled(log types.Log) (*AccountrouterRemoteIsmEnrolled, error) {
	event := new(AccountrouterRemoteIsmEnrolled)
	if err := _Accountrouter.contract.UnpackLog(event, "RemoteIsmEnrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountrouterRemoteRouterEnrolledIterator is returned from FilterRemoteRouterEnrolled and is used to iterate over the raw logs and unpacked data for RemoteRouterEnrolled events raised by the Accountrouter contract.
type AccountrouterRemoteRouterEnrolledIterator struct {
	Event *AccountrouterRemoteRouterEnrolled // Event containing the contract specifics and raw log

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
func (it *AccountrouterRemoteRouterEnrolledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountrouterRemoteRouterEnrolled)
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
		it.Event = new(AccountrouterRemoteRouterEnrolled)
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
func (it *AccountrouterRemoteRouterEnrolledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountrouterRemoteRouterEnrolledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountrouterRemoteRouterEnrolled represents a RemoteRouterEnrolled event raised by the Accountrouter contract.
type AccountrouterRemoteRouterEnrolled struct {
	Domain uint32
	Router [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoteRouterEnrolled is a free log retrieval operation binding the contract event 0x97fd51fef1f80c854cbd0150c248e8b2ac4ecdc97cc05e742b5f61c9f4dc6458.
//
// Solidity: event RemoteRouterEnrolled(uint32 indexed domain, bytes32 router)
func (_Accountrouter *AccountrouterFilterer) FilterRemoteRouterEnrolled(opts *bind.FilterOpts, domain []uint32) (*AccountrouterRemoteRouterEnrolledIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Accountrouter.contract.FilterLogs(opts, "RemoteRouterEnrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return &AccountrouterRemoteRouterEnrolledIterator{contract: _Accountrouter.contract, event: "RemoteRouterEnrolled", logs: logs, sub: sub}, nil
}

// WatchRemoteRouterEnrolled is a free log subscription operation binding the contract event 0x97fd51fef1f80c854cbd0150c248e8b2ac4ecdc97cc05e742b5f61c9f4dc6458.
//
// Solidity: event RemoteRouterEnrolled(uint32 indexed domain, bytes32 router)
func (_Accountrouter *AccountrouterFilterer) WatchRemoteRouterEnrolled(opts *bind.WatchOpts, sink chan<- *AccountrouterRemoteRouterEnrolled, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Accountrouter.contract.WatchLogs(opts, "RemoteRouterEnrolled", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountrouterRemoteRouterEnrolled)
				if err := _Accountrouter.contract.UnpackLog(event, "RemoteRouterEnrolled", log); err != nil {
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

// ParseRemoteRouterEnrolled is a log parse operation binding the contract event 0x97fd51fef1f80c854cbd0150c248e8b2ac4ecdc97cc05e742b5f61c9f4dc6458.
//
// Solidity: event RemoteRouterEnrolled(uint32 indexed domain, bytes32 router)
func (_Accountrouter *AccountrouterFilterer) ParseRemoteRouterEnrolled(log types.Log) (*AccountrouterRemoteRouterEnrolled, error) {
	event := new(AccountrouterRemoteRouterEnrolled)
	if err := _Accountrouter.contract.UnpackLog(event, "RemoteRouterEnrolled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
