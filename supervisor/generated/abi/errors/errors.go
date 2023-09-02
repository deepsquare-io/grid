// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package errorsabi

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

// Affinity is an auto generated low-level Go binding around an user-defined struct.
type Affinity struct {
	Label Label
	Op    [2]byte
}

// Job is an auto generated low-level Go binding around an user-defined struct.
type Job struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       JobDefinition
	Cost             JobCost
	Time             JobTime
	JobName          [32]byte
	HasCancelRequest bool
	LastError        string
}

// JobCost is an auto generated low-level Go binding around an user-defined struct.
type JobCost struct {
	MaxCost                   *big.Int
	FinalCost                 *big.Int
	PendingTopUp              *big.Int
	DelegateSpendingAuthority bool
}

// JobDefinition is an auto generated low-level Go binding around an user-defined struct.
type JobDefinition struct {
	GpusPerTask       uint64
	MemPerCpu         uint64
	CpusPerTask       uint64
	Ntasks            uint64
	BatchLocationHash string
	StorageType       uint8
	Uses              []Label
	Affinity          []Affinity
}

// JobTime is an auto generated low-level Go binding around an user-defined struct.
type JobTime struct {
	Start                  *big.Int
	End                    *big.Int
	CancelRequestTimestamp *big.Int
	BlockNumberStateChange *big.Int
	PanicTimestamp         *big.Int
}

// Label is an auto generated low-level Go binding around an user-defined struct.
type Label struct {
	Key   string
	Value string
}

// Provider is an auto generated low-level Go binding around an user-defined struct.
type Provider struct {
	WalletAddr       common.Address
	ProviderHardware ProviderHardware
	ProviderPrices   ProviderPrices
	Status           uint8
	JobCount         uint64
	Labels           []Label
	LinkListed       bool
}

// ProviderHardware is an auto generated low-level Go binding around an user-defined struct.
type ProviderHardware struct {
	Nodes       uint64
	GpusPerNode []uint64
	CpusPerNode []uint64
	MemPerNode  []uint64
}

// ProviderPrices is an auto generated low-level Go binding around an user-defined struct.
type ProviderPrices struct {
	GpuPricePerMin *big.Int
	CpuPricePerMin *big.Int
	MemPricePerMin *big.Int
}

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControl contract.
type AccessControlRoleAdminChangedIterator struct {
	Event *AccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleAdminChanged)
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
		it.Event = new(AccessControlRoleAdminChanged)
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
func (it *AccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControl contract.
type AccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleAdminChangedIterator{contract: _AccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleAdminChanged)
				if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlRoleAdminChanged, error) {
	event := new(AccessControlRoleAdminChanged)
	if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstantsMetaData contains all meta data concerning the Constants contract.
var ConstantsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"billDurationDeltaMinute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancellationFeeMinute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimJobTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deepsquareCut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_billDurationDeltaMinute\",\"type\":\"uint256\"}],\"name\":\"setBillDurationDeltaMinute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cancellationFeeMinute\",\"type\":\"uint256\"}],\"name\":\"setCancellationFeeMinute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_claimJobTimeout\",\"type\":\"uint64\"}],\"name\":\"setClaimJobTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deepsquareCut\",\"type\":\"uint256\"}],\"name\":\"setDeepsquareCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumAmount\",\"type\":\"uint256\"}],\"name\":\"setMinimumAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_topUpSliceDurationMin\",\"type\":\"uint64\"}],\"name\":\"setTopUpSliceDurationMin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"topUpSliceDurationMin\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405268056bc75e2d6310000060015560056002819055600f600355600480546001600160401b0319908116601e17909155601490915560068054909116600a17905534801561005057600080fd5b5061005a3361005f565b6100af565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61048d806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063b959017111610097578063d7f37cc611610066578063d7f37cc6146101da578063eeb4a9c8146101e3578063f2fde38b146101f6578063f64a2d671461020957600080fd5b8063b9590171146101a1578063bb0c8298146101aa578063c670a130146101b3578063d6aa37a6146101c657600080fd5b8063715018a6116100d3578063715018a6146101355780638ce9843b1461013d5780638da5cb5b1461016f578063a234d90f1461018a57600080fd5b80631d84a59d146100fa5780632bb301591461010f5780635e60af5114610122575b600080fd5b61010d6101083660046103e4565b61021c565b005b61010d61011d366004610415565b610248565b61010d610130366004610415565b610255565b61010d610262565b6004546101519067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6000546040516001600160a01b039091168152602001610166565b61019360055481565b604051908152602001610166565b61019360035481565b61019360015481565b61010d6101c13660046103e4565b610276565b6006546101519067ffffffffffffffff1681565b61019360025481565b61010d6101f1366004610415565b6102a2565b61010d61020436600461042e565b6102af565b61010d610217366004610415565b61032d565b61022461033a565b6006805467ffffffffffffffff191667ffffffffffffffff92909216919091179055565b61025061033a565b600255565b61025d61033a565b600355565b61026a61033a565b6102746000610394565b565b61027e61033a565b6004805467ffffffffffffffff191667ffffffffffffffff92909216919091179055565b6102aa61033a565b600155565b6102b761033a565b6001600160a01b0381166103215760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61032a81610394565b50565b61033561033a565b600555565b6000546001600160a01b031633146102745760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610318565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156103f657600080fd5b813567ffffffffffffffff8116811461040e57600080fd5b9392505050565b60006020828403121561042757600080fd5b5035919050565b60006020828403121561044057600080fd5b81356001600160a01b038116811461040e57600080fdfea264697066735822122047a9483953f13dbc3d6b4ed2b2efc47ce1c241dfca0ed000f09021c6eb0d721e64736f6c63430008110033",
}

// ConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstantsMetaData.ABI instead.
var ConstantsABI = ConstantsMetaData.ABI

// ConstantsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConstantsMetaData.Bin instead.
var ConstantsBin = ConstantsMetaData.Bin

// DeployConstants deploys a new Ethereum contract, binding an instance of Constants to it.
func DeployConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Constants, error) {
	parsed, err := ConstantsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}, ConstantsFilterer: ConstantsFilterer{contract: contract}}, nil
}

// Constants is an auto generated Go binding around an Ethereum contract.
type Constants struct {
	ConstantsCaller     // Read-only binding to the contract
	ConstantsTransactor // Write-only binding to the contract
	ConstantsFilterer   // Log filterer for contract events
}

// ConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantsSession struct {
	Contract     *Constants        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantsCallerSession struct {
	Contract *ConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantsTransactorSession struct {
	Contract     *ConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantsRaw struct {
	Contract *Constants // Generic contract binding to access the raw methods on
}

// ConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantsCallerRaw struct {
	Contract *ConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantsTransactorRaw struct {
	Contract *ConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstants creates a new instance of Constants, bound to a specific deployed contract.
func NewConstants(address common.Address, backend bind.ContractBackend) (*Constants, error) {
	contract, err := bindConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}, ConstantsFilterer: ConstantsFilterer{contract: contract}}, nil
}

// NewConstantsCaller creates a new read-only instance of Constants, bound to a specific deployed contract.
func NewConstantsCaller(address common.Address, caller bind.ContractCaller) (*ConstantsCaller, error) {
	contract, err := bindConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsCaller{contract: contract}, nil
}

// NewConstantsTransactor creates a new write-only instance of Constants, bound to a specific deployed contract.
func NewConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantsTransactor, error) {
	contract, err := bindConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsTransactor{contract: contract}, nil
}

// NewConstantsFilterer creates a new log filterer instance of Constants, bound to a specific deployed contract.
func NewConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstantsFilterer, error) {
	contract, err := bindConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstantsFilterer{contract: contract}, nil
}

// bindConstants binds a generic wrapper to an already deployed contract.
func bindConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConstantsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.ConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transact(opts, method, params...)
}

// BillDurationDeltaMinute is a free data retrieval call binding the contract method 0xb9590171.
//
// Solidity: function billDurationDeltaMinute() view returns(uint256)
func (_Constants *ConstantsCaller) BillDurationDeltaMinute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "billDurationDeltaMinute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BillDurationDeltaMinute is a free data retrieval call binding the contract method 0xb9590171.
//
// Solidity: function billDurationDeltaMinute() view returns(uint256)
func (_Constants *ConstantsSession) BillDurationDeltaMinute() (*big.Int, error) {
	return _Constants.Contract.BillDurationDeltaMinute(&_Constants.CallOpts)
}

// BillDurationDeltaMinute is a free data retrieval call binding the contract method 0xb9590171.
//
// Solidity: function billDurationDeltaMinute() view returns(uint256)
func (_Constants *ConstantsCallerSession) BillDurationDeltaMinute() (*big.Int, error) {
	return _Constants.Contract.BillDurationDeltaMinute(&_Constants.CallOpts)
}

// CancellationFeeMinute is a free data retrieval call binding the contract method 0xd7f37cc6.
//
// Solidity: function cancellationFeeMinute() view returns(uint256)
func (_Constants *ConstantsCaller) CancellationFeeMinute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "cancellationFeeMinute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CancellationFeeMinute is a free data retrieval call binding the contract method 0xd7f37cc6.
//
// Solidity: function cancellationFeeMinute() view returns(uint256)
func (_Constants *ConstantsSession) CancellationFeeMinute() (*big.Int, error) {
	return _Constants.Contract.CancellationFeeMinute(&_Constants.CallOpts)
}

// CancellationFeeMinute is a free data retrieval call binding the contract method 0xd7f37cc6.
//
// Solidity: function cancellationFeeMinute() view returns(uint256)
func (_Constants *ConstantsCallerSession) CancellationFeeMinute() (*big.Int, error) {
	return _Constants.Contract.CancellationFeeMinute(&_Constants.CallOpts)
}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_Constants *ConstantsCaller) ClaimJobTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "claimJobTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_Constants *ConstantsSession) ClaimJobTimeout() (uint64, error) {
	return _Constants.Contract.ClaimJobTimeout(&_Constants.CallOpts)
}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_Constants *ConstantsCallerSession) ClaimJobTimeout() (uint64, error) {
	return _Constants.Contract.ClaimJobTimeout(&_Constants.CallOpts)
}

// DeepsquareCut is a free data retrieval call binding the contract method 0xa234d90f.
//
// Solidity: function deepsquareCut() view returns(uint256)
func (_Constants *ConstantsCaller) DeepsquareCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "deepsquareCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeepsquareCut is a free data retrieval call binding the contract method 0xa234d90f.
//
// Solidity: function deepsquareCut() view returns(uint256)
func (_Constants *ConstantsSession) DeepsquareCut() (*big.Int, error) {
	return _Constants.Contract.DeepsquareCut(&_Constants.CallOpts)
}

// DeepsquareCut is a free data retrieval call binding the contract method 0xa234d90f.
//
// Solidity: function deepsquareCut() view returns(uint256)
func (_Constants *ConstantsCallerSession) DeepsquareCut() (*big.Int, error) {
	return _Constants.Contract.DeepsquareCut(&_Constants.CallOpts)
}

// MinimumAmount is a free data retrieval call binding the contract method 0xbb0c8298.
//
// Solidity: function minimumAmount() view returns(uint256)
func (_Constants *ConstantsCaller) MinimumAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "minimumAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAmount is a free data retrieval call binding the contract method 0xbb0c8298.
//
// Solidity: function minimumAmount() view returns(uint256)
func (_Constants *ConstantsSession) MinimumAmount() (*big.Int, error) {
	return _Constants.Contract.MinimumAmount(&_Constants.CallOpts)
}

// MinimumAmount is a free data retrieval call binding the contract method 0xbb0c8298.
//
// Solidity: function minimumAmount() view returns(uint256)
func (_Constants *ConstantsCallerSession) MinimumAmount() (*big.Int, error) {
	return _Constants.Contract.MinimumAmount(&_Constants.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Constants *ConstantsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Constants *ConstantsSession) Owner() (common.Address, error) {
	return _Constants.Contract.Owner(&_Constants.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Constants *ConstantsCallerSession) Owner() (common.Address, error) {
	return _Constants.Contract.Owner(&_Constants.CallOpts)
}

// TopUpSliceDurationMin is a free data retrieval call binding the contract method 0x8ce9843b.
//
// Solidity: function topUpSliceDurationMin() view returns(uint64)
func (_Constants *ConstantsCaller) TopUpSliceDurationMin(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "topUpSliceDurationMin")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TopUpSliceDurationMin is a free data retrieval call binding the contract method 0x8ce9843b.
//
// Solidity: function topUpSliceDurationMin() view returns(uint64)
func (_Constants *ConstantsSession) TopUpSliceDurationMin() (uint64, error) {
	return _Constants.Contract.TopUpSliceDurationMin(&_Constants.CallOpts)
}

// TopUpSliceDurationMin is a free data retrieval call binding the contract method 0x8ce9843b.
//
// Solidity: function topUpSliceDurationMin() view returns(uint64)
func (_Constants *ConstantsCallerSession) TopUpSliceDurationMin() (uint64, error) {
	return _Constants.Contract.TopUpSliceDurationMin(&_Constants.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Constants *ConstantsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Constants *ConstantsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Constants.Contract.RenounceOwnership(&_Constants.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Constants *ConstantsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Constants.Contract.RenounceOwnership(&_Constants.TransactOpts)
}

// SetBillDurationDeltaMinute is a paid mutator transaction binding the contract method 0x5e60af51.
//
// Solidity: function setBillDurationDeltaMinute(uint256 _billDurationDeltaMinute) returns()
func (_Constants *ConstantsTransactor) SetBillDurationDeltaMinute(opts *bind.TransactOpts, _billDurationDeltaMinute *big.Int) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "setBillDurationDeltaMinute", _billDurationDeltaMinute)
}

// SetBillDurationDeltaMinute is a paid mutator transaction binding the contract method 0x5e60af51.
//
// Solidity: function setBillDurationDeltaMinute(uint256 _billDurationDeltaMinute) returns()
func (_Constants *ConstantsSession) SetBillDurationDeltaMinute(_billDurationDeltaMinute *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetBillDurationDeltaMinute(&_Constants.TransactOpts, _billDurationDeltaMinute)
}

// SetBillDurationDeltaMinute is a paid mutator transaction binding the contract method 0x5e60af51.
//
// Solidity: function setBillDurationDeltaMinute(uint256 _billDurationDeltaMinute) returns()
func (_Constants *ConstantsTransactorSession) SetBillDurationDeltaMinute(_billDurationDeltaMinute *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetBillDurationDeltaMinute(&_Constants.TransactOpts, _billDurationDeltaMinute)
}

// SetCancellationFeeMinute is a paid mutator transaction binding the contract method 0x2bb30159.
//
// Solidity: function setCancellationFeeMinute(uint256 _cancellationFeeMinute) returns()
func (_Constants *ConstantsTransactor) SetCancellationFeeMinute(opts *bind.TransactOpts, _cancellationFeeMinute *big.Int) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "setCancellationFeeMinute", _cancellationFeeMinute)
}

// SetCancellationFeeMinute is a paid mutator transaction binding the contract method 0x2bb30159.
//
// Solidity: function setCancellationFeeMinute(uint256 _cancellationFeeMinute) returns()
func (_Constants *ConstantsSession) SetCancellationFeeMinute(_cancellationFeeMinute *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetCancellationFeeMinute(&_Constants.TransactOpts, _cancellationFeeMinute)
}

// SetCancellationFeeMinute is a paid mutator transaction binding the contract method 0x2bb30159.
//
// Solidity: function setCancellationFeeMinute(uint256 _cancellationFeeMinute) returns()
func (_Constants *ConstantsTransactorSession) SetCancellationFeeMinute(_cancellationFeeMinute *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetCancellationFeeMinute(&_Constants.TransactOpts, _cancellationFeeMinute)
}

// SetClaimJobTimeout is a paid mutator transaction binding the contract method 0x1d84a59d.
//
// Solidity: function setClaimJobTimeout(uint64 _claimJobTimeout) returns()
func (_Constants *ConstantsTransactor) SetClaimJobTimeout(opts *bind.TransactOpts, _claimJobTimeout uint64) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "setClaimJobTimeout", _claimJobTimeout)
}

// SetClaimJobTimeout is a paid mutator transaction binding the contract method 0x1d84a59d.
//
// Solidity: function setClaimJobTimeout(uint64 _claimJobTimeout) returns()
func (_Constants *ConstantsSession) SetClaimJobTimeout(_claimJobTimeout uint64) (*types.Transaction, error) {
	return _Constants.Contract.SetClaimJobTimeout(&_Constants.TransactOpts, _claimJobTimeout)
}

// SetClaimJobTimeout is a paid mutator transaction binding the contract method 0x1d84a59d.
//
// Solidity: function setClaimJobTimeout(uint64 _claimJobTimeout) returns()
func (_Constants *ConstantsTransactorSession) SetClaimJobTimeout(_claimJobTimeout uint64) (*types.Transaction, error) {
	return _Constants.Contract.SetClaimJobTimeout(&_Constants.TransactOpts, _claimJobTimeout)
}

// SetDeepsquareCut is a paid mutator transaction binding the contract method 0xf64a2d67.
//
// Solidity: function setDeepsquareCut(uint256 _deepsquareCut) returns()
func (_Constants *ConstantsTransactor) SetDeepsquareCut(opts *bind.TransactOpts, _deepsquareCut *big.Int) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "setDeepsquareCut", _deepsquareCut)
}

// SetDeepsquareCut is a paid mutator transaction binding the contract method 0xf64a2d67.
//
// Solidity: function setDeepsquareCut(uint256 _deepsquareCut) returns()
func (_Constants *ConstantsSession) SetDeepsquareCut(_deepsquareCut *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetDeepsquareCut(&_Constants.TransactOpts, _deepsquareCut)
}

// SetDeepsquareCut is a paid mutator transaction binding the contract method 0xf64a2d67.
//
// Solidity: function setDeepsquareCut(uint256 _deepsquareCut) returns()
func (_Constants *ConstantsTransactorSession) SetDeepsquareCut(_deepsquareCut *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetDeepsquareCut(&_Constants.TransactOpts, _deepsquareCut)
}

// SetMinimumAmount is a paid mutator transaction binding the contract method 0xeeb4a9c8.
//
// Solidity: function setMinimumAmount(uint256 _minimumAmount) returns()
func (_Constants *ConstantsTransactor) SetMinimumAmount(opts *bind.TransactOpts, _minimumAmount *big.Int) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "setMinimumAmount", _minimumAmount)
}

// SetMinimumAmount is a paid mutator transaction binding the contract method 0xeeb4a9c8.
//
// Solidity: function setMinimumAmount(uint256 _minimumAmount) returns()
func (_Constants *ConstantsSession) SetMinimumAmount(_minimumAmount *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetMinimumAmount(&_Constants.TransactOpts, _minimumAmount)
}

// SetMinimumAmount is a paid mutator transaction binding the contract method 0xeeb4a9c8.
//
// Solidity: function setMinimumAmount(uint256 _minimumAmount) returns()
func (_Constants *ConstantsTransactorSession) SetMinimumAmount(_minimumAmount *big.Int) (*types.Transaction, error) {
	return _Constants.Contract.SetMinimumAmount(&_Constants.TransactOpts, _minimumAmount)
}

// SetTopUpSliceDurationMin is a paid mutator transaction binding the contract method 0xc670a130.
//
// Solidity: function setTopUpSliceDurationMin(uint64 _topUpSliceDurationMin) returns()
func (_Constants *ConstantsTransactor) SetTopUpSliceDurationMin(opts *bind.TransactOpts, _topUpSliceDurationMin uint64) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "setTopUpSliceDurationMin", _topUpSliceDurationMin)
}

// SetTopUpSliceDurationMin is a paid mutator transaction binding the contract method 0xc670a130.
//
// Solidity: function setTopUpSliceDurationMin(uint64 _topUpSliceDurationMin) returns()
func (_Constants *ConstantsSession) SetTopUpSliceDurationMin(_topUpSliceDurationMin uint64) (*types.Transaction, error) {
	return _Constants.Contract.SetTopUpSliceDurationMin(&_Constants.TransactOpts, _topUpSliceDurationMin)
}

// SetTopUpSliceDurationMin is a paid mutator transaction binding the contract method 0xc670a130.
//
// Solidity: function setTopUpSliceDurationMin(uint64 _topUpSliceDurationMin) returns()
func (_Constants *ConstantsTransactorSession) SetTopUpSliceDurationMin(_topUpSliceDurationMin uint64) (*types.Transaction, error) {
	return _Constants.Contract.SetTopUpSliceDurationMin(&_Constants.TransactOpts, _topUpSliceDurationMin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Constants *ConstantsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Constants.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Constants *ConstantsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Constants.Contract.TransferOwnership(&_Constants.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Constants *ConstantsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Constants.Contract.TransferOwnership(&_Constants.TransactOpts, newOwner)
}

// ConstantsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Constants contract.
type ConstantsOwnershipTransferredIterator struct {
	Event *ConstantsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConstantsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantsOwnershipTransferred)
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
		it.Event = new(ConstantsOwnershipTransferred)
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
func (it *ConstantsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantsOwnershipTransferred represents a OwnershipTransferred event raised by the Constants contract.
type ConstantsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Constants *ConstantsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConstantsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Constants.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConstantsOwnershipTransferredIterator{contract: _Constants.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Constants *ConstantsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConstantsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Constants.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantsOwnershipTransferred)
				if err := _Constants.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Constants *ConstantsFilterer) ParseOwnershipTransferred(log types.Log) (*ConstantsOwnershipTransferred, error) {
	event := new(ConstantsOwnershipTransferred)
	if err := _Constants.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// DoubleEndedQueueMetaData contains all meta data concerning the DoubleEndedQueue contract.
var DoubleEndedQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122099ff439cc24e6694b324ae0fd5b3a9bce8a051ca752cce630ee448c541084d4264736f6c63430008110033",
}

// DoubleEndedQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use DoubleEndedQueueMetaData.ABI instead.
var DoubleEndedQueueABI = DoubleEndedQueueMetaData.ABI

// DoubleEndedQueueBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DoubleEndedQueueMetaData.Bin instead.
var DoubleEndedQueueBin = DoubleEndedQueueMetaData.Bin

// DeployDoubleEndedQueue deploys a new Ethereum contract, binding an instance of DoubleEndedQueue to it.
func DeployDoubleEndedQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DoubleEndedQueue, error) {
	parsed, err := DoubleEndedQueueMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DoubleEndedQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DoubleEndedQueue{DoubleEndedQueueCaller: DoubleEndedQueueCaller{contract: contract}, DoubleEndedQueueTransactor: DoubleEndedQueueTransactor{contract: contract}, DoubleEndedQueueFilterer: DoubleEndedQueueFilterer{contract: contract}}, nil
}

// DoubleEndedQueue is an auto generated Go binding around an Ethereum contract.
type DoubleEndedQueue struct {
	DoubleEndedQueueCaller     // Read-only binding to the contract
	DoubleEndedQueueTransactor // Write-only binding to the contract
	DoubleEndedQueueFilterer   // Log filterer for contract events
}

// DoubleEndedQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type DoubleEndedQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleEndedQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DoubleEndedQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleEndedQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DoubleEndedQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleEndedQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DoubleEndedQueueSession struct {
	Contract     *DoubleEndedQueue // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoubleEndedQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DoubleEndedQueueCallerSession struct {
	Contract *DoubleEndedQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DoubleEndedQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DoubleEndedQueueTransactorSession struct {
	Contract     *DoubleEndedQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DoubleEndedQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type DoubleEndedQueueRaw struct {
	Contract *DoubleEndedQueue // Generic contract binding to access the raw methods on
}

// DoubleEndedQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DoubleEndedQueueCallerRaw struct {
	Contract *DoubleEndedQueueCaller // Generic read-only contract binding to access the raw methods on
}

// DoubleEndedQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DoubleEndedQueueTransactorRaw struct {
	Contract *DoubleEndedQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoubleEndedQueue creates a new instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueue(address common.Address, backend bind.ContractBackend) (*DoubleEndedQueue, error) {
	contract, err := bindDoubleEndedQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueue{DoubleEndedQueueCaller: DoubleEndedQueueCaller{contract: contract}, DoubleEndedQueueTransactor: DoubleEndedQueueTransactor{contract: contract}, DoubleEndedQueueFilterer: DoubleEndedQueueFilterer{contract: contract}}, nil
}

// NewDoubleEndedQueueCaller creates a new read-only instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueueCaller(address common.Address, caller bind.ContractCaller) (*DoubleEndedQueueCaller, error) {
	contract, err := bindDoubleEndedQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueueCaller{contract: contract}, nil
}

// NewDoubleEndedQueueTransactor creates a new write-only instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*DoubleEndedQueueTransactor, error) {
	contract, err := bindDoubleEndedQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueueTransactor{contract: contract}, nil
}

// NewDoubleEndedQueueFilterer creates a new log filterer instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*DoubleEndedQueueFilterer, error) {
	contract, err := bindDoubleEndedQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueueFilterer{contract: contract}, nil
}

// bindDoubleEndedQueue binds a generic wrapper to an already deployed contract.
func bindDoubleEndedQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DoubleEndedQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleEndedQueue *DoubleEndedQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleEndedQueue.Contract.DoubleEndedQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleEndedQueue *DoubleEndedQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.DoubleEndedQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleEndedQueue *DoubleEndedQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.DoubleEndedQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleEndedQueue *DoubleEndedQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleEndedQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleEndedQueue *DoubleEndedQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleEndedQueue *DoubleEndedQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.contract.Transact(opts, method, params...)
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ErrorContractMetaData contains all meta data concerning the ErrorContract contract.
var ErrorContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Banned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CustomerMetaSchedulerProviderOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"CustomerOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJobDefinition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNodesCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTotalCpus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTotalMem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumJobStatus\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidTransition\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"JobHotStatusOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"JobProviderOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"MetaScheduledScheduledStatusOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewJobRequestDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoJob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpendingAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNotJoined\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"RemainingTimeAboveLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"RunningColdStatusOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"RunningScheduledStatusOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameStatusError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WaitingApprovalOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThrowArrayLengthMismatch\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowBanned\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowCustomerMetaSchedulerProviderOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"ThrowCustomerOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowEmpty\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"ThrowInsufficientFunds\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidJob\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidJobDefinition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidNodesCount\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidTotalCpus\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidTotalMem\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumJobStatus\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"ThrowInvalidTransition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowJobHotStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"ThrowJobProviderOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowMetaScheduledScheduledStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowNewJobRequestDisabled\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowNoJob\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowNoProvider\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowNoSpendingAuthority\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowOutOfBounds\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowProviderNotJoined\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"ThrowRemainingTimeAboveLimit\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowRunningColdStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowRunningScheduledStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowSameStatusError\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowWaitingApprovalOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106ef806100206000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c806361fdbb96116100de578063c8ba784511610097578063ede8052311610071578063ede8052314610292578063f2f22fab146102a5578063f522f385146102b8578063fb2ad369146102c057600080fd5b8063c8ba78451461026f578063cbfbbf8a14610277578063d0980b701461028a57600080fd5b806361fdbb961461021e578063693a22971461022657806392234b111461022e57806393514ad5146102415780639c46b00614610249578063a6b419581461025c57600080fd5b80633e3c27691161014b578063573cc84f11610125578063573cc84f146101fe5780635ad71c45146102065780635c71386a1461020e57806361cb83701461021657600080fd5b80633e3c2769146101db5780634094eedf146101ee57806356d74982146101f657600080fd5b806301b9c30f1461019357806301d72a591461019d57806310047926146101b05780631161fdf0146101b857806317d0a33f146101c057806333c78712146101d3575b600080fd5b61019b6102c8565b005b61019b6101ab36600461059c565b6102e1565b61019b610309565b61019b610322565b61019b6101ce3660046105d2565b61033b565b61019b610356565b61019b6101e93660046105d2565b61036f565b61019b61038a565b61019b6103a3565b61019b6103bc565b61019b6103d5565b61019b6103ee565b61019b610407565b61019b610420565b61019b610439565b61019b61023c3660046105d2565b610452565b61019b61046d565b61019b61025736600461059c565b610486565b61019b61026a36600461060b565b6104a8565b61019b6104d4565b61019b61028536600461060b565b6104ed565b61019b610519565b61019b6102a036600461063e565b610532565b61019b6102b33660046105d2565b61054f565b61019b61056a565b61019b610583565b6040516303c2f8df60e51b815260040160405180910390fd5b604051634801db4560e11b815260048101839052602481018290526044015b60405180910390fd5b604051632d91f2bb60e11b815260040160405180910390fd5b604051633abe75b360e01b815260040160405180910390fd5b8060405163ed3e2aad60e01b8152600401610300919061068a565b60405163700dd2fd60e11b815260040160405180910390fd5b80604051634634126160e11b8152600401610300919061068a565b604051630d208e5960e41b815260040160405180910390fd5b60405163038e47a360e51b815260040160405180910390fd5b604051632a9126eb60e01b815260040160405180910390fd5b604051637064572b60e01b815260040160405180910390fd5b60405163ef341f6d60e01b815260040160405180910390fd5b604051632a856fc960e01b815260040160405180910390fd5b604051639773692760e01b815260040160405180910390fd5b6040516342f9901960e01b815260040160405180910390fd5b806040516314475eb760e01b8152600401610300919061068a565b604051637897ef6d60e01b815260040160405180910390fd5b60405162fae2d560e21b81526004810183905260248101829052604401610300565b604051630cb8c19760e21b81526001600160a01b03808416600483015282166024820152604401610300565b604051631ed9509560e11b815260040160405180910390fd5b604051638942331960e01b81526001600160a01b03808416600483015282166024820152604401610300565b604051633a43ca4160e01b815260040160405180910390fd5b81816040516305fdf05f60e31b815260040161030092919061069e565b8060405163048389ff60e11b8152600401610300919061068a565b604051632d0483c560e21b815260040160405180910390fd5b60405163512509d360e11b815260040160405180910390fd5b600080604083850312156105af57600080fd5b50508035926020909101359150565b8035600981106105cd57600080fd5b919050565b6000602082840312156105e457600080fd5b6105ed826105be565b9392505050565b80356001600160a01b03811681146105cd57600080fd5b6000806040838503121561061e57600080fd5b610627836105f4565b9150610635602084016105f4565b90509250929050565b6000806040838503121561065157600080fd5b61065a836105be565b9150610635602084016105be565b6009811061068657634e487b7160e01b600052602160045260246000fd5b9052565b602081016106988284610668565b92915050565b604081016106ac8285610668565b6105ed602083018461066856fea26469706673582212203a8e238699ed116ca13b6bdb9280f4215a36b72811b1eb934c023825469ea9a164736f6c63430008110033",
}

// ErrorContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ErrorContractMetaData.ABI instead.
var ErrorContractABI = ErrorContractMetaData.ABI

// ErrorContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ErrorContractMetaData.Bin instead.
var ErrorContractBin = ErrorContractMetaData.Bin

// DeployErrorContract deploys a new Ethereum contract, binding an instance of ErrorContract to it.
func DeployErrorContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ErrorContract, error) {
	parsed, err := ErrorContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ErrorContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ErrorContract{ErrorContractCaller: ErrorContractCaller{contract: contract}, ErrorContractTransactor: ErrorContractTransactor{contract: contract}, ErrorContractFilterer: ErrorContractFilterer{contract: contract}}, nil
}

// ErrorContract is an auto generated Go binding around an Ethereum contract.
type ErrorContract struct {
	ErrorContractCaller     // Read-only binding to the contract
	ErrorContractTransactor // Write-only binding to the contract
	ErrorContractFilterer   // Log filterer for contract events
}

// ErrorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErrorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErrorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErrorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErrorContractSession struct {
	Contract     *ErrorContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErrorContractCallerSession struct {
	Contract *ErrorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ErrorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErrorContractTransactorSession struct {
	Contract     *ErrorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ErrorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErrorContractRaw struct {
	Contract *ErrorContract // Generic contract binding to access the raw methods on
}

// ErrorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErrorContractCallerRaw struct {
	Contract *ErrorContractCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErrorContractTransactorRaw struct {
	Contract *ErrorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErrorContract creates a new instance of ErrorContract, bound to a specific deployed contract.
func NewErrorContract(address common.Address, backend bind.ContractBackend) (*ErrorContract, error) {
	contract, err := bindErrorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ErrorContract{ErrorContractCaller: ErrorContractCaller{contract: contract}, ErrorContractTransactor: ErrorContractTransactor{contract: contract}, ErrorContractFilterer: ErrorContractFilterer{contract: contract}}, nil
}

// NewErrorContractCaller creates a new read-only instance of ErrorContract, bound to a specific deployed contract.
func NewErrorContractCaller(address common.Address, caller bind.ContractCaller) (*ErrorContractCaller, error) {
	contract, err := bindErrorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorContractCaller{contract: contract}, nil
}

// NewErrorContractTransactor creates a new write-only instance of ErrorContract, bound to a specific deployed contract.
func NewErrorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorContractTransactor, error) {
	contract, err := bindErrorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorContractTransactor{contract: contract}, nil
}

// NewErrorContractFilterer creates a new log filterer instance of ErrorContract, bound to a specific deployed contract.
func NewErrorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorContractFilterer, error) {
	contract, err := bindErrorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorContractFilterer{contract: contract}, nil
}

// bindErrorContract binds a generic wrapper to an already deployed contract.
func bindErrorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ErrorContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErrorContract *ErrorContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErrorContract.Contract.ErrorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErrorContract *ErrorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErrorContract.Contract.ErrorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErrorContract *ErrorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErrorContract.Contract.ErrorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErrorContract *ErrorContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErrorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErrorContract *ErrorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErrorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErrorContract *ErrorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErrorContract.Contract.contract.Transact(opts, method, params...)
}

// ThrowArrayLengthMismatch is a free data retrieval call binding the contract method 0xfb2ad369.
//
// Solidity: function ThrowArrayLengthMismatch() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowArrayLengthMismatch(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowArrayLengthMismatch")

	if err != nil {
		return err
	}

	return err

}

// ThrowArrayLengthMismatch is a free data retrieval call binding the contract method 0xfb2ad369.
//
// Solidity: function ThrowArrayLengthMismatch() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowArrayLengthMismatch() error {
	return _ErrorContract.Contract.ThrowArrayLengthMismatch(&_ErrorContract.CallOpts)
}

// ThrowArrayLengthMismatch is a free data retrieval call binding the contract method 0xfb2ad369.
//
// Solidity: function ThrowArrayLengthMismatch() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowArrayLengthMismatch() error {
	return _ErrorContract.Contract.ThrowArrayLengthMismatch(&_ErrorContract.CallOpts)
}

// ThrowBanned is a free data retrieval call binding the contract method 0x5ad71c45.
//
// Solidity: function ThrowBanned() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowBanned(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowBanned")

	if err != nil {
		return err
	}

	return err

}

// ThrowBanned is a free data retrieval call binding the contract method 0x5ad71c45.
//
// Solidity: function ThrowBanned() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowBanned() error {
	return _ErrorContract.Contract.ThrowBanned(&_ErrorContract.CallOpts)
}

// ThrowBanned is a free data retrieval call binding the contract method 0x5ad71c45.
//
// Solidity: function ThrowBanned() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowBanned() error {
	return _ErrorContract.Contract.ThrowBanned(&_ErrorContract.CallOpts)
}

// ThrowCustomerMetaSchedulerProviderOnly is a free data retrieval call binding the contract method 0x93514ad5.
//
// Solidity: function ThrowCustomerMetaSchedulerProviderOnly() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowCustomerMetaSchedulerProviderOnly(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowCustomerMetaSchedulerProviderOnly")

	if err != nil {
		return err
	}

	return err

}

// ThrowCustomerMetaSchedulerProviderOnly is a free data retrieval call binding the contract method 0x93514ad5.
//
// Solidity: function ThrowCustomerMetaSchedulerProviderOnly() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowCustomerMetaSchedulerProviderOnly() error {
	return _ErrorContract.Contract.ThrowCustomerMetaSchedulerProviderOnly(&_ErrorContract.CallOpts)
}

// ThrowCustomerMetaSchedulerProviderOnly is a free data retrieval call binding the contract method 0x93514ad5.
//
// Solidity: function ThrowCustomerMetaSchedulerProviderOnly() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowCustomerMetaSchedulerProviderOnly() error {
	return _ErrorContract.Contract.ThrowCustomerMetaSchedulerProviderOnly(&_ErrorContract.CallOpts)
}

// ThrowCustomerOnly is a free data retrieval call binding the contract method 0xcbfbbf8a.
//
// Solidity: function ThrowCustomerOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowCustomerOnly(opts *bind.CallOpts, current common.Address, expected common.Address) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowCustomerOnly", current, expected)

	if err != nil {
		return err
	}

	return err

}

// ThrowCustomerOnly is a free data retrieval call binding the contract method 0xcbfbbf8a.
//
// Solidity: function ThrowCustomerOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowCustomerOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowCustomerOnly(&_ErrorContract.CallOpts, current, expected)
}

// ThrowCustomerOnly is a free data retrieval call binding the contract method 0xcbfbbf8a.
//
// Solidity: function ThrowCustomerOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowCustomerOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowCustomerOnly(&_ErrorContract.CallOpts, current, expected)
}

// ThrowEmpty is a free data retrieval call binding the contract method 0xc8ba7845.
//
// Solidity: function ThrowEmpty() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowEmpty(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowEmpty")

	if err != nil {
		return err
	}

	return err

}

// ThrowEmpty is a free data retrieval call binding the contract method 0xc8ba7845.
//
// Solidity: function ThrowEmpty() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowEmpty() error {
	return _ErrorContract.Contract.ThrowEmpty(&_ErrorContract.CallOpts)
}

// ThrowEmpty is a free data retrieval call binding the contract method 0xc8ba7845.
//
// Solidity: function ThrowEmpty() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowEmpty() error {
	return _ErrorContract.Contract.ThrowEmpty(&_ErrorContract.CallOpts)
}

// ThrowInsufficientFunds is a free data retrieval call binding the contract method 0x9c46b006.
//
// Solidity: function ThrowInsufficientFunds(uint256 available, uint256 required) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInsufficientFunds(opts *bind.CallOpts, available *big.Int, required *big.Int) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInsufficientFunds", available, required)

	if err != nil {
		return err
	}

	return err

}

// ThrowInsufficientFunds is a free data retrieval call binding the contract method 0x9c46b006.
//
// Solidity: function ThrowInsufficientFunds(uint256 available, uint256 required) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInsufficientFunds(available *big.Int, required *big.Int) error {
	return _ErrorContract.Contract.ThrowInsufficientFunds(&_ErrorContract.CallOpts, available, required)
}

// ThrowInsufficientFunds is a free data retrieval call binding the contract method 0x9c46b006.
//
// Solidity: function ThrowInsufficientFunds(uint256 available, uint256 required) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInsufficientFunds(available *big.Int, required *big.Int) error {
	return _ErrorContract.Contract.ThrowInsufficientFunds(&_ErrorContract.CallOpts, available, required)
}

// ThrowInvalidJob is a free data retrieval call binding the contract method 0x56d74982.
//
// Solidity: function ThrowInvalidJob() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidJob(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidJob")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidJob is a free data retrieval call binding the contract method 0x56d74982.
//
// Solidity: function ThrowInvalidJob() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidJob() error {
	return _ErrorContract.Contract.ThrowInvalidJob(&_ErrorContract.CallOpts)
}

// ThrowInvalidJob is a free data retrieval call binding the contract method 0x56d74982.
//
// Solidity: function ThrowInvalidJob() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidJob() error {
	return _ErrorContract.Contract.ThrowInvalidJob(&_ErrorContract.CallOpts)
}

// ThrowInvalidJobDefinition is a free data retrieval call binding the contract method 0x4094eedf.
//
// Solidity: function ThrowInvalidJobDefinition() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidJobDefinition(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidJobDefinition")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidJobDefinition is a free data retrieval call binding the contract method 0x4094eedf.
//
// Solidity: function ThrowInvalidJobDefinition() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidJobDefinition() error {
	return _ErrorContract.Contract.ThrowInvalidJobDefinition(&_ErrorContract.CallOpts)
}

// ThrowInvalidJobDefinition is a free data retrieval call binding the contract method 0x4094eedf.
//
// Solidity: function ThrowInvalidJobDefinition() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidJobDefinition() error {
	return _ErrorContract.Contract.ThrowInvalidJobDefinition(&_ErrorContract.CallOpts)
}

// ThrowInvalidNodesCount is a free data retrieval call binding the contract method 0x33c78712.
//
// Solidity: function ThrowInvalidNodesCount() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidNodesCount(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidNodesCount")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidNodesCount is a free data retrieval call binding the contract method 0x33c78712.
//
// Solidity: function ThrowInvalidNodesCount() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidNodesCount() error {
	return _ErrorContract.Contract.ThrowInvalidNodesCount(&_ErrorContract.CallOpts)
}

// ThrowInvalidNodesCount is a free data retrieval call binding the contract method 0x33c78712.
//
// Solidity: function ThrowInvalidNodesCount() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidNodesCount() error {
	return _ErrorContract.Contract.ThrowInvalidNodesCount(&_ErrorContract.CallOpts)
}

// ThrowInvalidTotalCpus is a free data retrieval call binding the contract method 0x10047926.
//
// Solidity: function ThrowInvalidTotalCpus() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidTotalCpus(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidTotalCpus")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidTotalCpus is a free data retrieval call binding the contract method 0x10047926.
//
// Solidity: function ThrowInvalidTotalCpus() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidTotalCpus() error {
	return _ErrorContract.Contract.ThrowInvalidTotalCpus(&_ErrorContract.CallOpts)
}

// ThrowInvalidTotalCpus is a free data retrieval call binding the contract method 0x10047926.
//
// Solidity: function ThrowInvalidTotalCpus() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidTotalCpus() error {
	return _ErrorContract.Contract.ThrowInvalidTotalCpus(&_ErrorContract.CallOpts)
}

// ThrowInvalidTotalMem is a free data retrieval call binding the contract method 0x01b9c30f.
//
// Solidity: function ThrowInvalidTotalMem() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidTotalMem(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidTotalMem")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidTotalMem is a free data retrieval call binding the contract method 0x01b9c30f.
//
// Solidity: function ThrowInvalidTotalMem() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidTotalMem() error {
	return _ErrorContract.Contract.ThrowInvalidTotalMem(&_ErrorContract.CallOpts)
}

// ThrowInvalidTotalMem is a free data retrieval call binding the contract method 0x01b9c30f.
//
// Solidity: function ThrowInvalidTotalMem() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidTotalMem() error {
	return _ErrorContract.Contract.ThrowInvalidTotalMem(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransition is a free data retrieval call binding the contract method 0xede80523.
//
// Solidity: function ThrowInvalidTransition(uint8 from, uint8 to) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidTransition(opts *bind.CallOpts, from uint8, to uint8) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidTransition", from, to)

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidTransition is a free data retrieval call binding the contract method 0xede80523.
//
// Solidity: function ThrowInvalidTransition(uint8 from, uint8 to) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidTransition(from uint8, to uint8) error {
	return _ErrorContract.Contract.ThrowInvalidTransition(&_ErrorContract.CallOpts, from, to)
}

// ThrowInvalidTransition is a free data retrieval call binding the contract method 0xede80523.
//
// Solidity: function ThrowInvalidTransition(uint8 from, uint8 to) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidTransition(from uint8, to uint8) error {
	return _ErrorContract.Contract.ThrowInvalidTransition(&_ErrorContract.CallOpts, from, to)
}

// ThrowJobHotStatusOnly is a free data retrieval call binding the contract method 0x3e3c2769.
//
// Solidity: function ThrowJobHotStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowJobHotStatusOnly(opts *bind.CallOpts, current uint8) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowJobHotStatusOnly", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowJobHotStatusOnly is a free data retrieval call binding the contract method 0x3e3c2769.
//
// Solidity: function ThrowJobHotStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowJobHotStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowJobHotStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowJobHotStatusOnly is a free data retrieval call binding the contract method 0x3e3c2769.
//
// Solidity: function ThrowJobHotStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowJobHotStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowJobHotStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowJobProviderOnly is a free data retrieval call binding the contract method 0xa6b41958.
//
// Solidity: function ThrowJobProviderOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowJobProviderOnly(opts *bind.CallOpts, current common.Address, expected common.Address) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowJobProviderOnly", current, expected)

	if err != nil {
		return err
	}

	return err

}

// ThrowJobProviderOnly is a free data retrieval call binding the contract method 0xa6b41958.
//
// Solidity: function ThrowJobProviderOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowJobProviderOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowJobProviderOnly(&_ErrorContract.CallOpts, current, expected)
}

// ThrowJobProviderOnly is a free data retrieval call binding the contract method 0xa6b41958.
//
// Solidity: function ThrowJobProviderOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowJobProviderOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowJobProviderOnly(&_ErrorContract.CallOpts, current, expected)
}

// ThrowMetaScheduledScheduledStatusOnly is a free data retrieval call binding the contract method 0xf2f22fab.
//
// Solidity: function ThrowMetaScheduledScheduledStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowMetaScheduledScheduledStatusOnly(opts *bind.CallOpts, current uint8) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowMetaScheduledScheduledStatusOnly", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowMetaScheduledScheduledStatusOnly is a free data retrieval call binding the contract method 0xf2f22fab.
//
// Solidity: function ThrowMetaScheduledScheduledStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowMetaScheduledScheduledStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowMetaScheduledScheduledStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowMetaScheduledScheduledStatusOnly is a free data retrieval call binding the contract method 0xf2f22fab.
//
// Solidity: function ThrowMetaScheduledScheduledStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowMetaScheduledScheduledStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowMetaScheduledScheduledStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowNewJobRequestDisabled is a free data retrieval call binding the contract method 0x1161fdf0.
//
// Solidity: function ThrowNewJobRequestDisabled() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowNewJobRequestDisabled(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowNewJobRequestDisabled")

	if err != nil {
		return err
	}

	return err

}

// ThrowNewJobRequestDisabled is a free data retrieval call binding the contract method 0x1161fdf0.
//
// Solidity: function ThrowNewJobRequestDisabled() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowNewJobRequestDisabled() error {
	return _ErrorContract.Contract.ThrowNewJobRequestDisabled(&_ErrorContract.CallOpts)
}

// ThrowNewJobRequestDisabled is a free data retrieval call binding the contract method 0x1161fdf0.
//
// Solidity: function ThrowNewJobRequestDisabled() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowNewJobRequestDisabled() error {
	return _ErrorContract.Contract.ThrowNewJobRequestDisabled(&_ErrorContract.CallOpts)
}

// ThrowNoJob is a free data retrieval call binding the contract method 0x61fdbb96.
//
// Solidity: function ThrowNoJob() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowNoJob(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowNoJob")

	if err != nil {
		return err
	}

	return err

}

// ThrowNoJob is a free data retrieval call binding the contract method 0x61fdbb96.
//
// Solidity: function ThrowNoJob() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowNoJob() error {
	return _ErrorContract.Contract.ThrowNoJob(&_ErrorContract.CallOpts)
}

// ThrowNoJob is a free data retrieval call binding the contract method 0x61fdbb96.
//
// Solidity: function ThrowNoJob() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowNoJob() error {
	return _ErrorContract.Contract.ThrowNoJob(&_ErrorContract.CallOpts)
}

// ThrowNoProvider is a free data retrieval call binding the contract method 0xd0980b70.
//
// Solidity: function ThrowNoProvider() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowNoProvider(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowNoProvider")

	if err != nil {
		return err
	}

	return err

}

// ThrowNoProvider is a free data retrieval call binding the contract method 0xd0980b70.
//
// Solidity: function ThrowNoProvider() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowNoProvider() error {
	return _ErrorContract.Contract.ThrowNoProvider(&_ErrorContract.CallOpts)
}

// ThrowNoProvider is a free data retrieval call binding the contract method 0xd0980b70.
//
// Solidity: function ThrowNoProvider() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowNoProvider() error {
	return _ErrorContract.Contract.ThrowNoProvider(&_ErrorContract.CallOpts)
}

// ThrowNoSpendingAuthority is a free data retrieval call binding the contract method 0x693a2297.
//
// Solidity: function ThrowNoSpendingAuthority() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowNoSpendingAuthority(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowNoSpendingAuthority")

	if err != nil {
		return err
	}

	return err

}

// ThrowNoSpendingAuthority is a free data retrieval call binding the contract method 0x693a2297.
//
// Solidity: function ThrowNoSpendingAuthority() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowNoSpendingAuthority() error {
	return _ErrorContract.Contract.ThrowNoSpendingAuthority(&_ErrorContract.CallOpts)
}

// ThrowNoSpendingAuthority is a free data retrieval call binding the contract method 0x693a2297.
//
// Solidity: function ThrowNoSpendingAuthority() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowNoSpendingAuthority() error {
	return _ErrorContract.Contract.ThrowNoSpendingAuthority(&_ErrorContract.CallOpts)
}

// ThrowOutOfBounds is a free data retrieval call binding the contract method 0xf522f385.
//
// Solidity: function ThrowOutOfBounds() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowOutOfBounds(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowOutOfBounds")

	if err != nil {
		return err
	}

	return err

}

// ThrowOutOfBounds is a free data retrieval call binding the contract method 0xf522f385.
//
// Solidity: function ThrowOutOfBounds() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowOutOfBounds() error {
	return _ErrorContract.Contract.ThrowOutOfBounds(&_ErrorContract.CallOpts)
}

// ThrowOutOfBounds is a free data retrieval call binding the contract method 0xf522f385.
//
// Solidity: function ThrowOutOfBounds() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowOutOfBounds() error {
	return _ErrorContract.Contract.ThrowOutOfBounds(&_ErrorContract.CallOpts)
}

// ThrowProviderNotJoined is a free data retrieval call binding the contract method 0x5c71386a.
//
// Solidity: function ThrowProviderNotJoined() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowProviderNotJoined(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowProviderNotJoined")

	if err != nil {
		return err
	}

	return err

}

// ThrowProviderNotJoined is a free data retrieval call binding the contract method 0x5c71386a.
//
// Solidity: function ThrowProviderNotJoined() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowProviderNotJoined() error {
	return _ErrorContract.Contract.ThrowProviderNotJoined(&_ErrorContract.CallOpts)
}

// ThrowProviderNotJoined is a free data retrieval call binding the contract method 0x5c71386a.
//
// Solidity: function ThrowProviderNotJoined() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowProviderNotJoined() error {
	return _ErrorContract.Contract.ThrowProviderNotJoined(&_ErrorContract.CallOpts)
}

// ThrowRemainingTimeAboveLimit is a free data retrieval call binding the contract method 0x01d72a59.
//
// Solidity: function ThrowRemainingTimeAboveLimit(uint256 remaining, uint256 limit) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowRemainingTimeAboveLimit(opts *bind.CallOpts, remaining *big.Int, limit *big.Int) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowRemainingTimeAboveLimit", remaining, limit)

	if err != nil {
		return err
	}

	return err

}

// ThrowRemainingTimeAboveLimit is a free data retrieval call binding the contract method 0x01d72a59.
//
// Solidity: function ThrowRemainingTimeAboveLimit(uint256 remaining, uint256 limit) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowRemainingTimeAboveLimit(remaining *big.Int, limit *big.Int) error {
	return _ErrorContract.Contract.ThrowRemainingTimeAboveLimit(&_ErrorContract.CallOpts, remaining, limit)
}

// ThrowRemainingTimeAboveLimit is a free data retrieval call binding the contract method 0x01d72a59.
//
// Solidity: function ThrowRemainingTimeAboveLimit(uint256 remaining, uint256 limit) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowRemainingTimeAboveLimit(remaining *big.Int, limit *big.Int) error {
	return _ErrorContract.Contract.ThrowRemainingTimeAboveLimit(&_ErrorContract.CallOpts, remaining, limit)
}

// ThrowRunningColdStatusOnly is a free data retrieval call binding the contract method 0x92234b11.
//
// Solidity: function ThrowRunningColdStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowRunningColdStatusOnly(opts *bind.CallOpts, current uint8) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowRunningColdStatusOnly", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowRunningColdStatusOnly is a free data retrieval call binding the contract method 0x92234b11.
//
// Solidity: function ThrowRunningColdStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowRunningColdStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowRunningColdStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowRunningColdStatusOnly is a free data retrieval call binding the contract method 0x92234b11.
//
// Solidity: function ThrowRunningColdStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowRunningColdStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowRunningColdStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowRunningScheduledStatusOnly is a free data retrieval call binding the contract method 0x17d0a33f.
//
// Solidity: function ThrowRunningScheduledStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowRunningScheduledStatusOnly(opts *bind.CallOpts, current uint8) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowRunningScheduledStatusOnly", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowRunningScheduledStatusOnly is a free data retrieval call binding the contract method 0x17d0a33f.
//
// Solidity: function ThrowRunningScheduledStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowRunningScheduledStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowRunningScheduledStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowRunningScheduledStatusOnly is a free data retrieval call binding the contract method 0x17d0a33f.
//
// Solidity: function ThrowRunningScheduledStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowRunningScheduledStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowRunningScheduledStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowSameStatusError is a free data retrieval call binding the contract method 0x61cb8370.
//
// Solidity: function ThrowSameStatusError() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowSameStatusError(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowSameStatusError")

	if err != nil {
		return err
	}

	return err

}

// ThrowSameStatusError is a free data retrieval call binding the contract method 0x61cb8370.
//
// Solidity: function ThrowSameStatusError() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowSameStatusError() error {
	return _ErrorContract.Contract.ThrowSameStatusError(&_ErrorContract.CallOpts)
}

// ThrowSameStatusError is a free data retrieval call binding the contract method 0x61cb8370.
//
// Solidity: function ThrowSameStatusError() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowSameStatusError() error {
	return _ErrorContract.Contract.ThrowSameStatusError(&_ErrorContract.CallOpts)
}

// ThrowWaitingApprovalOnly is a free data retrieval call binding the contract method 0x573cc84f.
//
// Solidity: function ThrowWaitingApprovalOnly() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowWaitingApprovalOnly(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowWaitingApprovalOnly")

	if err != nil {
		return err
	}

	return err

}

// ThrowWaitingApprovalOnly is a free data retrieval call binding the contract method 0x573cc84f.
//
// Solidity: function ThrowWaitingApprovalOnly() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowWaitingApprovalOnly() error {
	return _ErrorContract.Contract.ThrowWaitingApprovalOnly(&_ErrorContract.CallOpts)
}

// ThrowWaitingApprovalOnly is a free data retrieval call binding the contract method 0x573cc84f.
//
// Solidity: function ThrowWaitingApprovalOnly() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowWaitingApprovalOnly() error {
	return _ErrorContract.Contract.ThrowWaitingApprovalOnly(&_ErrorContract.CallOpts)
}

// IAccessControlMetaData contains all meta data concerning the IAccessControl contract.
var IAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlMetaData.ABI instead.
var IAccessControlABI = IAccessControlMetaData.ABI

// IAccessControl is an auto generated Go binding around an Ethereum contract.
type IAccessControl struct {
	IAccessControlCaller     // Read-only binding to the contract
	IAccessControlTransactor // Write-only binding to the contract
	IAccessControlFilterer   // Log filterer for contract events
}

// IAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlSession struct {
	Contract     *IAccessControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlCallerSession struct {
	Contract *IAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlTransactorSession struct {
	Contract     *IAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlRaw struct {
	Contract *IAccessControl // Generic contract binding to access the raw methods on
}

// IAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlCallerRaw struct {
	Contract *IAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlTransactorRaw struct {
	Contract *IAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControl creates a new instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControl(address common.Address, backend bind.ContractBackend) (*IAccessControl, error) {
	contract, err := bindIAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControl{IAccessControlCaller: IAccessControlCaller{contract: contract}, IAccessControlTransactor: IAccessControlTransactor{contract: contract}, IAccessControlFilterer: IAccessControlFilterer{contract: contract}}, nil
}

// NewIAccessControlCaller creates a new read-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlCaller, error) {
	contract, err := bindIAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlCaller{contract: contract}, nil
}

// NewIAccessControlTransactor creates a new write-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlTransactor, error) {
	contract, err := bindIAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlTransactor{contract: contract}, nil
}

// NewIAccessControlFilterer creates a new log filterer instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlFilterer, error) {
	contract, err := bindIAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlFilterer{contract: contract}, nil
}

// bindIAccessControl binds a generic wrapper to an already deployed contract.
func bindIAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.IAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// IAccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IAccessControl contract.
type IAccessControlRoleAdminChangedIterator struct {
	Event *IAccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleAdminChanged)
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
		it.Event = new(IAccessControlRoleAdminChanged)
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
func (it *IAccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the IAccessControl contract.
type IAccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IAccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleAdminChangedIterator{contract: _IAccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleAdminChanged)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*IAccessControlRoleAdminChanged, error) {
	event := new(IAccessControlRoleAdminChanged)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IAccessControl contract.
type IAccessControlRoleGrantedIterator struct {
	Event *IAccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleGranted)
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
		it.Event = new(IAccessControlRoleGranted)
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
func (it *IAccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleGranted represents a RoleGranted event raised by the IAccessControl contract.
type IAccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleGrantedIterator{contract: _IAccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleGranted)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleGranted(log types.Log) (*IAccessControlRoleGranted, error) {
	event := new(IAccessControlRoleGranted)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IAccessControl contract.
type IAccessControlRoleRevokedIterator struct {
	Event *IAccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleRevoked)
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
		it.Event = new(IAccessControlRoleRevoked)
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
func (it *IAccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleRevoked represents a RoleRevoked event raised by the IAccessControl contract.
type IAccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleRevokedIterator{contract: _IAccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleRevoked)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleRevoked(log types.Log) (*IAccessControlRoleRevoked, error) {
	event := new(IAccessControlRoleRevoked)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IJobRepositoryMetaData contains all meta data concerning the IJobRepository contract.
var IJobRepositoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingTopUp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateSpendingAuthority\",\"type\":\"bool\"}],\"internalType\":\"structJobCost\",\"name\":\"cost\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelRequestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"panicTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structJobTime\",\"name\":\"time\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"hasCancelRequest\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"lastError\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"_job\",\"type\":\"tuple\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingTopUp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateSpendingAuthority\",\"type\":\"bool\"}],\"internalType\":\"structJobCost\",\"name\":\"cost\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelRequestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"panicTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structJobTime\",\"name\":\"time\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"hasCancelRequest\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"lastError\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"}],\"name\":\"getByCustomer\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingTopUp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateSpendingAuthority\",\"type\":\"bool\"}],\"internalType\":\"structJobCost\",\"name\":\"_cost\",\"type\":\"tuple\"}],\"name\":\"setCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"setCustomerAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"}],\"name\":\"setDefinition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_hasCancelRequest\",\"type\":\"bool\"}],\"name\":\"setHasCancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"}],\"name\":\"setJobName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_newStatus\",\"type\":\"uint8\"}],\"name\":\"setJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_error\",\"type\":\"string\"}],\"name\":\"setLastError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"setProviderAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelRequestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"panicTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structJobTime\",\"name\":\"_time\",\"type\":\"tuple\"}],\"name\":\"setTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingTopUp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateSpendingAuthority\",\"type\":\"bool\"}],\"internalType\":\"structJobCost\",\"name\":\"cost\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelRequestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"panicTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structJobTime\",\"name\":\"time\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"hasCancelRequest\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"lastError\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"_job\",\"type\":\"tuple\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IJobRepositoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IJobRepositoryMetaData.ABI instead.
var IJobRepositoryABI = IJobRepositoryMetaData.ABI

// IJobRepository is an auto generated Go binding around an Ethereum contract.
type IJobRepository struct {
	IJobRepositoryCaller     // Read-only binding to the contract
	IJobRepositoryTransactor // Write-only binding to the contract
	IJobRepositoryFilterer   // Log filterer for contract events
}

// IJobRepositoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJobRepositoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobRepositoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJobRepositoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobRepositoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJobRepositoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobRepositorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJobRepositorySession struct {
	Contract     *IJobRepository   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJobRepositoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJobRepositoryCallerSession struct {
	Contract *IJobRepositoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IJobRepositoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJobRepositoryTransactorSession struct {
	Contract     *IJobRepositoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IJobRepositoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJobRepositoryRaw struct {
	Contract *IJobRepository // Generic contract binding to access the raw methods on
}

// IJobRepositoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJobRepositoryCallerRaw struct {
	Contract *IJobRepositoryCaller // Generic read-only contract binding to access the raw methods on
}

// IJobRepositoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJobRepositoryTransactorRaw struct {
	Contract *IJobRepositoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJobRepository creates a new instance of IJobRepository, bound to a specific deployed contract.
func NewIJobRepository(address common.Address, backend bind.ContractBackend) (*IJobRepository, error) {
	contract, err := bindIJobRepository(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJobRepository{IJobRepositoryCaller: IJobRepositoryCaller{contract: contract}, IJobRepositoryTransactor: IJobRepositoryTransactor{contract: contract}, IJobRepositoryFilterer: IJobRepositoryFilterer{contract: contract}}, nil
}

// NewIJobRepositoryCaller creates a new read-only instance of IJobRepository, bound to a specific deployed contract.
func NewIJobRepositoryCaller(address common.Address, caller bind.ContractCaller) (*IJobRepositoryCaller, error) {
	contract, err := bindIJobRepository(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJobRepositoryCaller{contract: contract}, nil
}

// NewIJobRepositoryTransactor creates a new write-only instance of IJobRepository, bound to a specific deployed contract.
func NewIJobRepositoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IJobRepositoryTransactor, error) {
	contract, err := bindIJobRepository(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJobRepositoryTransactor{contract: contract}, nil
}

// NewIJobRepositoryFilterer creates a new log filterer instance of IJobRepository, bound to a specific deployed contract.
func NewIJobRepositoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IJobRepositoryFilterer, error) {
	contract, err := bindIJobRepository(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJobRepositoryFilterer{contract: contract}, nil
}

// bindIJobRepository binds a generic wrapper to an already deployed contract.
func bindIJobRepository(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IJobRepositoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJobRepository *IJobRepositoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJobRepository.Contract.IJobRepositoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJobRepository *IJobRepositoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJobRepository.Contract.IJobRepositoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJobRepository *IJobRepositoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJobRepository.Contract.IJobRepositoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJobRepository *IJobRepositoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJobRepository.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJobRepository *IJobRepositoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJobRepository.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJobRepository *IJobRepositoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJobRepository.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _jobId) view returns((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string))
func (_IJobRepository *IJobRepositoryCaller) Get(opts *bind.CallOpts, _jobId [32]byte) (Job, error) {
	var out []interface{}
	err := _IJobRepository.contract.Call(opts, &out, "get", _jobId)

	if err != nil {
		return *new(Job), err
	}

	out0 := *abi.ConvertType(out[0], new(Job)).(*Job)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _jobId) view returns((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string))
func (_IJobRepository *IJobRepositorySession) Get(_jobId [32]byte) (Job, error) {
	return _IJobRepository.Contract.Get(&_IJobRepository.CallOpts, _jobId)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _jobId) view returns((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string))
func (_IJobRepository *IJobRepositoryCallerSession) Get(_jobId [32]byte) (Job, error) {
	return _IJobRepository.Contract.Get(&_IJobRepository.CallOpts, _jobId)
}

// GetByCustomer is a free data retrieval call binding the contract method 0x89a33883.
//
// Solidity: function getByCustomer(address customerAddr) view returns(bytes32[])
func (_IJobRepository *IJobRepositoryCaller) GetByCustomer(opts *bind.CallOpts, customerAddr common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _IJobRepository.contract.Call(opts, &out, "getByCustomer", customerAddr)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetByCustomer is a free data retrieval call binding the contract method 0x89a33883.
//
// Solidity: function getByCustomer(address customerAddr) view returns(bytes32[])
func (_IJobRepository *IJobRepositorySession) GetByCustomer(customerAddr common.Address) ([][32]byte, error) {
	return _IJobRepository.Contract.GetByCustomer(&_IJobRepository.CallOpts, customerAddr)
}

// GetByCustomer is a free data retrieval call binding the contract method 0x89a33883.
//
// Solidity: function getByCustomer(address customerAddr) view returns(bytes32[])
func (_IJobRepository *IJobRepositoryCallerSession) GetByCustomer(customerAddr common.Address) ([][32]byte, error) {
	return _IJobRepository.Contract.GetByCustomer(&_IJobRepository.CallOpts, customerAddr)
}

// Create is a paid mutator transaction binding the contract method 0xce1ecc5a.
//
// Solidity: function create((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) _job) returns(bytes32)
func (_IJobRepository *IJobRepositoryTransactor) Create(opts *bind.TransactOpts, _job Job) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "create", _job)
}

// Create is a paid mutator transaction binding the contract method 0xce1ecc5a.
//
// Solidity: function create((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) _job) returns(bytes32)
func (_IJobRepository *IJobRepositorySession) Create(_job Job) (*types.Transaction, error) {
	return _IJobRepository.Contract.Create(&_IJobRepository.TransactOpts, _job)
}

// Create is a paid mutator transaction binding the contract method 0xce1ecc5a.
//
// Solidity: function create((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) _job) returns(bytes32)
func (_IJobRepository *IJobRepositoryTransactorSession) Create(_job Job) (*types.Transaction, error) {
	return _IJobRepository.Contract.Create(&_IJobRepository.TransactOpts, _job)
}

// SetCost is a paid mutator transaction binding the contract method 0xb7090331.
//
// Solidity: function setCost(bytes32 _jobId, (uint256,uint256,uint256,bool) _cost) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetCost(opts *bind.TransactOpts, _jobId [32]byte, _cost JobCost) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setCost", _jobId, _cost)
}

// SetCost is a paid mutator transaction binding the contract method 0xb7090331.
//
// Solidity: function setCost(bytes32 _jobId, (uint256,uint256,uint256,bool) _cost) returns()
func (_IJobRepository *IJobRepositorySession) SetCost(_jobId [32]byte, _cost JobCost) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetCost(&_IJobRepository.TransactOpts, _jobId, _cost)
}

// SetCost is a paid mutator transaction binding the contract method 0xb7090331.
//
// Solidity: function setCost(bytes32 _jobId, (uint256,uint256,uint256,bool) _cost) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetCost(_jobId [32]byte, _cost JobCost) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetCost(&_IJobRepository.TransactOpts, _jobId, _cost)
}

// SetCustomerAddr is a paid mutator transaction binding the contract method 0x49c70dfa.
//
// Solidity: function setCustomerAddr(bytes32 _jobId, address _customerAddr) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetCustomerAddr(opts *bind.TransactOpts, _jobId [32]byte, _customerAddr common.Address) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setCustomerAddr", _jobId, _customerAddr)
}

// SetCustomerAddr is a paid mutator transaction binding the contract method 0x49c70dfa.
//
// Solidity: function setCustomerAddr(bytes32 _jobId, address _customerAddr) returns()
func (_IJobRepository *IJobRepositorySession) SetCustomerAddr(_jobId [32]byte, _customerAddr common.Address) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetCustomerAddr(&_IJobRepository.TransactOpts, _jobId, _customerAddr)
}

// SetCustomerAddr is a paid mutator transaction binding the contract method 0x49c70dfa.
//
// Solidity: function setCustomerAddr(bytes32 _jobId, address _customerAddr) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetCustomerAddr(_jobId [32]byte, _customerAddr common.Address) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetCustomerAddr(&_IJobRepository.TransactOpts, _jobId, _customerAddr)
}

// SetDefinition is a paid mutator transaction binding the contract method 0xfbae3f97.
//
// Solidity: function setDefinition(bytes32 _jobId, (uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _definition) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetDefinition(opts *bind.TransactOpts, _jobId [32]byte, _definition JobDefinition) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setDefinition", _jobId, _definition)
}

// SetDefinition is a paid mutator transaction binding the contract method 0xfbae3f97.
//
// Solidity: function setDefinition(bytes32 _jobId, (uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _definition) returns()
func (_IJobRepository *IJobRepositorySession) SetDefinition(_jobId [32]byte, _definition JobDefinition) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetDefinition(&_IJobRepository.TransactOpts, _jobId, _definition)
}

// SetDefinition is a paid mutator transaction binding the contract method 0xfbae3f97.
//
// Solidity: function setDefinition(bytes32 _jobId, (uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _definition) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetDefinition(_jobId [32]byte, _definition JobDefinition) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetDefinition(&_IJobRepository.TransactOpts, _jobId, _definition)
}

// SetHasCancelRequest is a paid mutator transaction binding the contract method 0x8e4de1ca.
//
// Solidity: function setHasCancelRequest(bytes32 _jobId, bool _hasCancelRequest) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetHasCancelRequest(opts *bind.TransactOpts, _jobId [32]byte, _hasCancelRequest bool) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setHasCancelRequest", _jobId, _hasCancelRequest)
}

// SetHasCancelRequest is a paid mutator transaction binding the contract method 0x8e4de1ca.
//
// Solidity: function setHasCancelRequest(bytes32 _jobId, bool _hasCancelRequest) returns()
func (_IJobRepository *IJobRepositorySession) SetHasCancelRequest(_jobId [32]byte, _hasCancelRequest bool) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetHasCancelRequest(&_IJobRepository.TransactOpts, _jobId, _hasCancelRequest)
}

// SetHasCancelRequest is a paid mutator transaction binding the contract method 0x8e4de1ca.
//
// Solidity: function setHasCancelRequest(bytes32 _jobId, bool _hasCancelRequest) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetHasCancelRequest(_jobId [32]byte, _hasCancelRequest bool) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetHasCancelRequest(&_IJobRepository.TransactOpts, _jobId, _hasCancelRequest)
}

// SetJobName is a paid mutator transaction binding the contract method 0x2074ba14.
//
// Solidity: function setJobName(bytes32 _jobId, bytes32 _jobName) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetJobName(opts *bind.TransactOpts, _jobId [32]byte, _jobName [32]byte) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setJobName", _jobId, _jobName)
}

// SetJobName is a paid mutator transaction binding the contract method 0x2074ba14.
//
// Solidity: function setJobName(bytes32 _jobId, bytes32 _jobName) returns()
func (_IJobRepository *IJobRepositorySession) SetJobName(_jobId [32]byte, _jobName [32]byte) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetJobName(&_IJobRepository.TransactOpts, _jobId, _jobName)
}

// SetJobName is a paid mutator transaction binding the contract method 0x2074ba14.
//
// Solidity: function setJobName(bytes32 _jobId, bytes32 _jobName) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetJobName(_jobId [32]byte, _jobName [32]byte) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetJobName(&_IJobRepository.TransactOpts, _jobId, _jobName)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobId, uint8 _newStatus) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetJobStatus(opts *bind.TransactOpts, _jobId [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setJobStatus", _jobId, _newStatus)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobId, uint8 _newStatus) returns()
func (_IJobRepository *IJobRepositorySession) SetJobStatus(_jobId [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetJobStatus(&_IJobRepository.TransactOpts, _jobId, _newStatus)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobId, uint8 _newStatus) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetJobStatus(_jobId [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetJobStatus(&_IJobRepository.TransactOpts, _jobId, _newStatus)
}

// SetLastError is a paid mutator transaction binding the contract method 0xb613a721.
//
// Solidity: function setLastError(bytes32 _jobId, string _error) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetLastError(opts *bind.TransactOpts, _jobId [32]byte, _error string) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setLastError", _jobId, _error)
}

// SetLastError is a paid mutator transaction binding the contract method 0xb613a721.
//
// Solidity: function setLastError(bytes32 _jobId, string _error) returns()
func (_IJobRepository *IJobRepositorySession) SetLastError(_jobId [32]byte, _error string) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetLastError(&_IJobRepository.TransactOpts, _jobId, _error)
}

// SetLastError is a paid mutator transaction binding the contract method 0xb613a721.
//
// Solidity: function setLastError(bytes32 _jobId, string _error) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetLastError(_jobId [32]byte, _error string) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetLastError(&_IJobRepository.TransactOpts, _jobId, _error)
}

// SetProviderAddr is a paid mutator transaction binding the contract method 0x5aae4bbd.
//
// Solidity: function setProviderAddr(bytes32 _jobId, address _providerAddr) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetProviderAddr(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setProviderAddr", _jobId, _providerAddr)
}

// SetProviderAddr is a paid mutator transaction binding the contract method 0x5aae4bbd.
//
// Solidity: function setProviderAddr(bytes32 _jobId, address _providerAddr) returns()
func (_IJobRepository *IJobRepositorySession) SetProviderAddr(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetProviderAddr(&_IJobRepository.TransactOpts, _jobId, _providerAddr)
}

// SetProviderAddr is a paid mutator transaction binding the contract method 0x5aae4bbd.
//
// Solidity: function setProviderAddr(bytes32 _jobId, address _providerAddr) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetProviderAddr(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetProviderAddr(&_IJobRepository.TransactOpts, _jobId, _providerAddr)
}

// SetTime is a paid mutator transaction binding the contract method 0x3baa6cb5.
//
// Solidity: function setTime(bytes32 _jobId, (uint256,uint256,uint256,uint256,uint256) _time) returns()
func (_IJobRepository *IJobRepositoryTransactor) SetTime(opts *bind.TransactOpts, _jobId [32]byte, _time JobTime) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "setTime", _jobId, _time)
}

// SetTime is a paid mutator transaction binding the contract method 0x3baa6cb5.
//
// Solidity: function setTime(bytes32 _jobId, (uint256,uint256,uint256,uint256,uint256) _time) returns()
func (_IJobRepository *IJobRepositorySession) SetTime(_jobId [32]byte, _time JobTime) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetTime(&_IJobRepository.TransactOpts, _jobId, _time)
}

// SetTime is a paid mutator transaction binding the contract method 0x3baa6cb5.
//
// Solidity: function setTime(bytes32 _jobId, (uint256,uint256,uint256,uint256,uint256) _time) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) SetTime(_jobId [32]byte, _time JobTime) (*types.Transaction, error) {
	return _IJobRepository.Contract.SetTime(&_IJobRepository.TransactOpts, _jobId, _time)
}

// Update is a paid mutator transaction binding the contract method 0x0168467d.
//
// Solidity: function update((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) _job) returns()
func (_IJobRepository *IJobRepositoryTransactor) Update(opts *bind.TransactOpts, _job Job) (*types.Transaction, error) {
	return _IJobRepository.contract.Transact(opts, "update", _job)
}

// Update is a paid mutator transaction binding the contract method 0x0168467d.
//
// Solidity: function update((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) _job) returns()
func (_IJobRepository *IJobRepositorySession) Update(_job Job) (*types.Transaction, error) {
	return _IJobRepository.Contract.Update(&_IJobRepository.TransactOpts, _job)
}

// Update is a paid mutator transaction binding the contract method 0x0168467d.
//
// Solidity: function update((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) _job) returns()
func (_IJobRepository *IJobRepositoryTransactorSession) Update(_job Job) (*types.Transaction, error) {
	return _IJobRepository.Contract.Update(&_IJobRepository.TransactOpts, _job)
}

// IProviderJobQueuesMetaData contains all meta data concerning the IProviderJobQueues contract.
var IProviderJobQueuesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getCancellingJobQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getClaimableJobQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getTopUpJobQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasCancellingJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasNextClaimableJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasTopUpJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"popNextCancellingJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"popNextClaimableJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"popNextTopUpJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"pushCancellingJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"pushClaimableJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"pushTopUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IProviderJobQueuesABI is the input ABI used to generate the binding from.
// Deprecated: Use IProviderJobQueuesMetaData.ABI instead.
var IProviderJobQueuesABI = IProviderJobQueuesMetaData.ABI

// IProviderJobQueues is an auto generated Go binding around an Ethereum contract.
type IProviderJobQueues struct {
	IProviderJobQueuesCaller     // Read-only binding to the contract
	IProviderJobQueuesTransactor // Write-only binding to the contract
	IProviderJobQueuesFilterer   // Log filterer for contract events
}

// IProviderJobQueuesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProviderJobQueuesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProviderJobQueuesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProviderJobQueuesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProviderJobQueuesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProviderJobQueuesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProviderJobQueuesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProviderJobQueuesSession struct {
	Contract     *IProviderJobQueues // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IProviderJobQueuesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProviderJobQueuesCallerSession struct {
	Contract *IProviderJobQueuesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IProviderJobQueuesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProviderJobQueuesTransactorSession struct {
	Contract     *IProviderJobQueuesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IProviderJobQueuesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProviderJobQueuesRaw struct {
	Contract *IProviderJobQueues // Generic contract binding to access the raw methods on
}

// IProviderJobQueuesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProviderJobQueuesCallerRaw struct {
	Contract *IProviderJobQueuesCaller // Generic read-only contract binding to access the raw methods on
}

// IProviderJobQueuesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProviderJobQueuesTransactorRaw struct {
	Contract *IProviderJobQueuesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProviderJobQueues creates a new instance of IProviderJobQueues, bound to a specific deployed contract.
func NewIProviderJobQueues(address common.Address, backend bind.ContractBackend) (*IProviderJobQueues, error) {
	contract, err := bindIProviderJobQueues(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProviderJobQueues{IProviderJobQueuesCaller: IProviderJobQueuesCaller{contract: contract}, IProviderJobQueuesTransactor: IProviderJobQueuesTransactor{contract: contract}, IProviderJobQueuesFilterer: IProviderJobQueuesFilterer{contract: contract}}, nil
}

// NewIProviderJobQueuesCaller creates a new read-only instance of IProviderJobQueues, bound to a specific deployed contract.
func NewIProviderJobQueuesCaller(address common.Address, caller bind.ContractCaller) (*IProviderJobQueuesCaller, error) {
	contract, err := bindIProviderJobQueues(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProviderJobQueuesCaller{contract: contract}, nil
}

// NewIProviderJobQueuesTransactor creates a new write-only instance of IProviderJobQueues, bound to a specific deployed contract.
func NewIProviderJobQueuesTransactor(address common.Address, transactor bind.ContractTransactor) (*IProviderJobQueuesTransactor, error) {
	contract, err := bindIProviderJobQueues(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProviderJobQueuesTransactor{contract: contract}, nil
}

// NewIProviderJobQueuesFilterer creates a new log filterer instance of IProviderJobQueues, bound to a specific deployed contract.
func NewIProviderJobQueuesFilterer(address common.Address, filterer bind.ContractFilterer) (*IProviderJobQueuesFilterer, error) {
	contract, err := bindIProviderJobQueues(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProviderJobQueuesFilterer{contract: contract}, nil
}

// bindIProviderJobQueues binds a generic wrapper to an already deployed contract.
func bindIProviderJobQueues(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IProviderJobQueuesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProviderJobQueues *IProviderJobQueuesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProviderJobQueues.Contract.IProviderJobQueuesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProviderJobQueues *IProviderJobQueuesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.IProviderJobQueuesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProviderJobQueues *IProviderJobQueuesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.IProviderJobQueuesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProviderJobQueues *IProviderJobQueuesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProviderJobQueues.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProviderJobQueues *IProviderJobQueuesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProviderJobQueues *IProviderJobQueuesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.contract.Transact(opts, method, params...)
}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesCaller) HasCancellingJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _IProviderJobQueues.contract.Call(opts, &out, "hasCancellingJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _IProviderJobQueues.Contract.HasCancellingJob(&_IProviderJobQueues.CallOpts, _providerAddr)
}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesCallerSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _IProviderJobQueues.Contract.HasCancellingJob(&_IProviderJobQueues.CallOpts, _providerAddr)
}

// HasNextClaimableJob is a free data retrieval call binding the contract method 0x6502e50b.
//
// Solidity: function hasNextClaimableJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesCaller) HasNextClaimableJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _IProviderJobQueues.contract.Call(opts, &out, "hasNextClaimableJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasNextClaimableJob is a free data retrieval call binding the contract method 0x6502e50b.
//
// Solidity: function hasNextClaimableJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesSession) HasNextClaimableJob(_providerAddr common.Address) (bool, error) {
	return _IProviderJobQueues.Contract.HasNextClaimableJob(&_IProviderJobQueues.CallOpts, _providerAddr)
}

// HasNextClaimableJob is a free data retrieval call binding the contract method 0x6502e50b.
//
// Solidity: function hasNextClaimableJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesCallerSession) HasNextClaimableJob(_providerAddr common.Address) (bool, error) {
	return _IProviderJobQueues.Contract.HasNextClaimableJob(&_IProviderJobQueues.CallOpts, _providerAddr)
}

// HasTopUpJob is a free data retrieval call binding the contract method 0xc7070e2c.
//
// Solidity: function hasTopUpJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesCaller) HasTopUpJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _IProviderJobQueues.contract.Call(opts, &out, "hasTopUpJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasTopUpJob is a free data retrieval call binding the contract method 0xc7070e2c.
//
// Solidity: function hasTopUpJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesSession) HasTopUpJob(_providerAddr common.Address) (bool, error) {
	return _IProviderJobQueues.Contract.HasTopUpJob(&_IProviderJobQueues.CallOpts, _providerAddr)
}

// HasTopUpJob is a free data retrieval call binding the contract method 0xc7070e2c.
//
// Solidity: function hasTopUpJob(address _providerAddr) view returns(bool)
func (_IProviderJobQueues *IProviderJobQueuesCallerSession) HasTopUpJob(_providerAddr common.Address) (bool, error) {
	return _IProviderJobQueues.Contract.HasTopUpJob(&_IProviderJobQueues.CallOpts, _providerAddr)
}

// GetCancellingJobQueueSize is a paid mutator transaction binding the contract method 0x4c2a0b7b.
//
// Solidity: function getCancellingJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesTransactor) GetCancellingJobQueueSize(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "getCancellingJobQueueSize", _providerAddr)
}

// GetCancellingJobQueueSize is a paid mutator transaction binding the contract method 0x4c2a0b7b.
//
// Solidity: function getCancellingJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesSession) GetCancellingJobQueueSize(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.GetCancellingJobQueueSize(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// GetCancellingJobQueueSize is a paid mutator transaction binding the contract method 0x4c2a0b7b.
//
// Solidity: function getCancellingJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) GetCancellingJobQueueSize(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.GetCancellingJobQueueSize(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// GetClaimableJobQueueSize is a paid mutator transaction binding the contract method 0x79490261.
//
// Solidity: function getClaimableJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesTransactor) GetClaimableJobQueueSize(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "getClaimableJobQueueSize", _providerAddr)
}

// GetClaimableJobQueueSize is a paid mutator transaction binding the contract method 0x79490261.
//
// Solidity: function getClaimableJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesSession) GetClaimableJobQueueSize(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.GetClaimableJobQueueSize(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// GetClaimableJobQueueSize is a paid mutator transaction binding the contract method 0x79490261.
//
// Solidity: function getClaimableJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) GetClaimableJobQueueSize(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.GetClaimableJobQueueSize(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// GetTopUpJobQueueSize is a paid mutator transaction binding the contract method 0xc3345ca3.
//
// Solidity: function getTopUpJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesTransactor) GetTopUpJobQueueSize(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "getTopUpJobQueueSize", _providerAddr)
}

// GetTopUpJobQueueSize is a paid mutator transaction binding the contract method 0xc3345ca3.
//
// Solidity: function getTopUpJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesSession) GetTopUpJobQueueSize(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.GetTopUpJobQueueSize(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// GetTopUpJobQueueSize is a paid mutator transaction binding the contract method 0xc3345ca3.
//
// Solidity: function getTopUpJobQueueSize(address _providerAddr) returns(uint256 length)
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) GetTopUpJobQueueSize(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.GetTopUpJobQueueSize(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// PopNextCancellingJob is a paid mutator transaction binding the contract method 0xde1a4d46.
//
// Solidity: function popNextCancellingJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesTransactor) PopNextCancellingJob(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "popNextCancellingJob", _providerAddr)
}

// PopNextCancellingJob is a paid mutator transaction binding the contract method 0xde1a4d46.
//
// Solidity: function popNextCancellingJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesSession) PopNextCancellingJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PopNextCancellingJob(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// PopNextCancellingJob is a paid mutator transaction binding the contract method 0xde1a4d46.
//
// Solidity: function popNextCancellingJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) PopNextCancellingJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PopNextCancellingJob(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// PopNextClaimableJob is a paid mutator transaction binding the contract method 0xbd69abf4.
//
// Solidity: function popNextClaimableJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesTransactor) PopNextClaimableJob(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "popNextClaimableJob", _providerAddr)
}

// PopNextClaimableJob is a paid mutator transaction binding the contract method 0xbd69abf4.
//
// Solidity: function popNextClaimableJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesSession) PopNextClaimableJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PopNextClaimableJob(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// PopNextClaimableJob is a paid mutator transaction binding the contract method 0xbd69abf4.
//
// Solidity: function popNextClaimableJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) PopNextClaimableJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PopNextClaimableJob(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// PopNextTopUpJob is a paid mutator transaction binding the contract method 0x22d398d4.
//
// Solidity: function popNextTopUpJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesTransactor) PopNextTopUpJob(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "popNextTopUpJob", _providerAddr)
}

// PopNextTopUpJob is a paid mutator transaction binding the contract method 0x22d398d4.
//
// Solidity: function popNextTopUpJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesSession) PopNextTopUpJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PopNextTopUpJob(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// PopNextTopUpJob is a paid mutator transaction binding the contract method 0x22d398d4.
//
// Solidity: function popNextTopUpJob(address _providerAddr) returns(bytes32 jobId)
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) PopNextTopUpJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PopNextTopUpJob(&_IProviderJobQueues.TransactOpts, _providerAddr)
}

// PushCancellingJob is a paid mutator transaction binding the contract method 0x2b72522e.
//
// Solidity: function pushCancellingJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesTransactor) PushCancellingJob(opts *bind.TransactOpts, _providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "pushCancellingJob", _providerAddr, _jobId)
}

// PushCancellingJob is a paid mutator transaction binding the contract method 0x2b72522e.
//
// Solidity: function pushCancellingJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesSession) PushCancellingJob(_providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PushCancellingJob(&_IProviderJobQueues.TransactOpts, _providerAddr, _jobId)
}

// PushCancellingJob is a paid mutator transaction binding the contract method 0x2b72522e.
//
// Solidity: function pushCancellingJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) PushCancellingJob(_providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PushCancellingJob(&_IProviderJobQueues.TransactOpts, _providerAddr, _jobId)
}

// PushClaimableJob is a paid mutator transaction binding the contract method 0x65b4fc49.
//
// Solidity: function pushClaimableJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesTransactor) PushClaimableJob(opts *bind.TransactOpts, _providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "pushClaimableJob", _providerAddr, _jobId)
}

// PushClaimableJob is a paid mutator transaction binding the contract method 0x65b4fc49.
//
// Solidity: function pushClaimableJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesSession) PushClaimableJob(_providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PushClaimableJob(&_IProviderJobQueues.TransactOpts, _providerAddr, _jobId)
}

// PushClaimableJob is a paid mutator transaction binding the contract method 0x65b4fc49.
//
// Solidity: function pushClaimableJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) PushClaimableJob(_providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PushClaimableJob(&_IProviderJobQueues.TransactOpts, _providerAddr, _jobId)
}

// PushTopUpJob is a paid mutator transaction binding the contract method 0x27fed931.
//
// Solidity: function pushTopUpJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesTransactor) PushTopUpJob(opts *bind.TransactOpts, _providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.contract.Transact(opts, "pushTopUpJob", _providerAddr, _jobId)
}

// PushTopUpJob is a paid mutator transaction binding the contract method 0x27fed931.
//
// Solidity: function pushTopUpJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesSession) PushTopUpJob(_providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PushTopUpJob(&_IProviderJobQueues.TransactOpts, _providerAddr, _jobId)
}

// PushTopUpJob is a paid mutator transaction binding the contract method 0x27fed931.
//
// Solidity: function pushTopUpJob(address _providerAddr, bytes32 _jobId) returns()
func (_IProviderJobQueues *IProviderJobQueuesTransactorSession) PushTopUpJob(_providerAddr common.Address, _jobId [32]byte) (*types.Transaction, error) {
	return _IProviderJobQueues.Contract.PushTopUpJob(&_IProviderJobQueues.TransactOpts, _providerAddr, _jobId)
}

// IProviderManagerMetaData contains all meta data concerning the IProviderManager contract.
var IProviderManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"HardwareUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ProviderStatusChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getJobCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getLabels\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"gpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"cpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"memPerNode\",\"type\":\"uint64[]\"}],\"internalType\":\"structProviderHardware\",\"name\":\"providerHardware\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPricePerMin\",\"type\":\"uint256\"}],\"internalType\":\"structProviderPrices\",\"name\":\"providerPrices\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"labels\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"linkListed\",\"type\":\"bool\"}],\"internalType\":\"structProvider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getProviderHardware\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"gpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"cpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"memPerNode\",\"type\":\"uint64[]\"}],\"internalType\":\"structProviderHardware\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getProviderPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPricePerMin\",\"type\":\"uint256\"}],\"internalType\":\"structProviderPrices\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getProviderStatus\",\"outputs\":[{\"internalType\":\"enumProviderStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasJoined\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"incJobCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"gpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"cpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"memPerNode\",\"type\":\"uint64[]\"}],\"internalType\":\"structProviderHardware\",\"name\":\"_hardware\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPricePerMin\",\"type\":\"uint256\"}],\"internalType\":\"structProviderPrices\",\"name\":\"_prices\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"_labels\",\"type\":\"tuple[]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"gpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"cpusPerNode\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"memPerNode\",\"type\":\"uint64[]\"}],\"internalType\":\"structProviderHardware\",\"name\":\"_hardware\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPricePerMin\",\"type\":\"uint256\"}],\"internalType\":\"structProviderPrices\",\"name\":\"_prices\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"_labels\",\"type\":\"tuple[]\"}],\"name\":\"registerProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"reinstate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"removeProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IProviderManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IProviderManagerMetaData.ABI instead.
var IProviderManagerABI = IProviderManagerMetaData.ABI

// IProviderManager is an auto generated Go binding around an Ethereum contract.
type IProviderManager struct {
	IProviderManagerCaller     // Read-only binding to the contract
	IProviderManagerTransactor // Write-only binding to the contract
	IProviderManagerFilterer   // Log filterer for contract events
}

// IProviderManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProviderManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProviderManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProviderManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProviderManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProviderManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProviderManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProviderManagerSession struct {
	Contract     *IProviderManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProviderManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProviderManagerCallerSession struct {
	Contract *IProviderManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IProviderManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProviderManagerTransactorSession struct {
	Contract     *IProviderManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IProviderManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProviderManagerRaw struct {
	Contract *IProviderManager // Generic contract binding to access the raw methods on
}

// IProviderManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProviderManagerCallerRaw struct {
	Contract *IProviderManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IProviderManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProviderManagerTransactorRaw struct {
	Contract *IProviderManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProviderManager creates a new instance of IProviderManager, bound to a specific deployed contract.
func NewIProviderManager(address common.Address, backend bind.ContractBackend) (*IProviderManager, error) {
	contract, err := bindIProviderManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProviderManager{IProviderManagerCaller: IProviderManagerCaller{contract: contract}, IProviderManagerTransactor: IProviderManagerTransactor{contract: contract}, IProviderManagerFilterer: IProviderManagerFilterer{contract: contract}}, nil
}

// NewIProviderManagerCaller creates a new read-only instance of IProviderManager, bound to a specific deployed contract.
func NewIProviderManagerCaller(address common.Address, caller bind.ContractCaller) (*IProviderManagerCaller, error) {
	contract, err := bindIProviderManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProviderManagerCaller{contract: contract}, nil
}

// NewIProviderManagerTransactor creates a new write-only instance of IProviderManager, bound to a specific deployed contract.
func NewIProviderManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IProviderManagerTransactor, error) {
	contract, err := bindIProviderManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProviderManagerTransactor{contract: contract}, nil
}

// NewIProviderManagerFilterer creates a new log filterer instance of IProviderManager, bound to a specific deployed contract.
func NewIProviderManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IProviderManagerFilterer, error) {
	contract, err := bindIProviderManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProviderManagerFilterer{contract: contract}, nil
}

// bindIProviderManager binds a generic wrapper to an already deployed contract.
func bindIProviderManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IProviderManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProviderManager *IProviderManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProviderManager.Contract.IProviderManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProviderManager *IProviderManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProviderManager.Contract.IProviderManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProviderManager *IProviderManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProviderManager.Contract.IProviderManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProviderManager *IProviderManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProviderManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProviderManager *IProviderManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProviderManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProviderManager *IProviderManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProviderManager.Contract.contract.Transact(opts, method, params...)
}

// GetJobCount is a free data retrieval call binding the contract method 0x6830cdc4.
//
// Solidity: function getJobCount(address _providerAddr) view returns(uint64)
func (_IProviderManager *IProviderManagerCaller) GetJobCount(opts *bind.CallOpts, _providerAddr common.Address) (uint64, error) {
	var out []interface{}
	err := _IProviderManager.contract.Call(opts, &out, "getJobCount", _providerAddr)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetJobCount is a free data retrieval call binding the contract method 0x6830cdc4.
//
// Solidity: function getJobCount(address _providerAddr) view returns(uint64)
func (_IProviderManager *IProviderManagerSession) GetJobCount(_providerAddr common.Address) (uint64, error) {
	return _IProviderManager.Contract.GetJobCount(&_IProviderManager.CallOpts, _providerAddr)
}

// GetJobCount is a free data retrieval call binding the contract method 0x6830cdc4.
//
// Solidity: function getJobCount(address _providerAddr) view returns(uint64)
func (_IProviderManager *IProviderManagerCallerSession) GetJobCount(_providerAddr common.Address) (uint64, error) {
	return _IProviderManager.Contract.GetJobCount(&_IProviderManager.CallOpts, _providerAddr)
}

// GetLabels is a free data retrieval call binding the contract method 0x95473b2b.
//
// Solidity: function getLabels(address _providerAddr) view returns((string,string)[])
func (_IProviderManager *IProviderManagerCaller) GetLabels(opts *bind.CallOpts, _providerAddr common.Address) ([]Label, error) {
	var out []interface{}
	err := _IProviderManager.contract.Call(opts, &out, "getLabels", _providerAddr)

	if err != nil {
		return *new([]Label), err
	}

	out0 := *abi.ConvertType(out[0], new([]Label)).(*[]Label)

	return out0, err

}

// GetLabels is a free data retrieval call binding the contract method 0x95473b2b.
//
// Solidity: function getLabels(address _providerAddr) view returns((string,string)[])
func (_IProviderManager *IProviderManagerSession) GetLabels(_providerAddr common.Address) ([]Label, error) {
	return _IProviderManager.Contract.GetLabels(&_IProviderManager.CallOpts, _providerAddr)
}

// GetLabels is a free data retrieval call binding the contract method 0x95473b2b.
//
// Solidity: function getLabels(address _providerAddr) view returns((string,string)[])
func (_IProviderManager *IProviderManagerCallerSession) GetLabels(_providerAddr common.Address) ([]Label, error) {
	return _IProviderManager.Contract.GetLabels(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address _providerAddr) view returns((address,(uint64,uint64[],uint64[],uint64[]),(uint256,uint256,uint256),uint8,uint64,(string,string)[],bool))
func (_IProviderManager *IProviderManagerCaller) GetProvider(opts *bind.CallOpts, _providerAddr common.Address) (Provider, error) {
	var out []interface{}
	err := _IProviderManager.contract.Call(opts, &out, "getProvider", _providerAddr)

	if err != nil {
		return *new(Provider), err
	}

	out0 := *abi.ConvertType(out[0], new(Provider)).(*Provider)

	return out0, err

}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address _providerAddr) view returns((address,(uint64,uint64[],uint64[],uint64[]),(uint256,uint256,uint256),uint8,uint64,(string,string)[],bool))
func (_IProviderManager *IProviderManagerSession) GetProvider(_providerAddr common.Address) (Provider, error) {
	return _IProviderManager.Contract.GetProvider(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address _providerAddr) view returns((address,(uint64,uint64[],uint64[],uint64[]),(uint256,uint256,uint256),uint8,uint64,(string,string)[],bool))
func (_IProviderManager *IProviderManagerCallerSession) GetProvider(_providerAddr common.Address) (Provider, error) {
	return _IProviderManager.Contract.GetProvider(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProviderHardware is a free data retrieval call binding the contract method 0xe5500e40.
//
// Solidity: function getProviderHardware(address _providerAddr) view returns((uint64,uint64[],uint64[],uint64[]))
func (_IProviderManager *IProviderManagerCaller) GetProviderHardware(opts *bind.CallOpts, _providerAddr common.Address) (ProviderHardware, error) {
	var out []interface{}
	err := _IProviderManager.contract.Call(opts, &out, "getProviderHardware", _providerAddr)

	if err != nil {
		return *new(ProviderHardware), err
	}

	out0 := *abi.ConvertType(out[0], new(ProviderHardware)).(*ProviderHardware)

	return out0, err

}

// GetProviderHardware is a free data retrieval call binding the contract method 0xe5500e40.
//
// Solidity: function getProviderHardware(address _providerAddr) view returns((uint64,uint64[],uint64[],uint64[]))
func (_IProviderManager *IProviderManagerSession) GetProviderHardware(_providerAddr common.Address) (ProviderHardware, error) {
	return _IProviderManager.Contract.GetProviderHardware(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProviderHardware is a free data retrieval call binding the contract method 0xe5500e40.
//
// Solidity: function getProviderHardware(address _providerAddr) view returns((uint64,uint64[],uint64[],uint64[]))
func (_IProviderManager *IProviderManagerCallerSession) GetProviderHardware(_providerAddr common.Address) (ProviderHardware, error) {
	return _IProviderManager.Contract.GetProviderHardware(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProviderPrices is a free data retrieval call binding the contract method 0x106859b6.
//
// Solidity: function getProviderPrices(address _providerAddr) view returns((uint256,uint256,uint256))
func (_IProviderManager *IProviderManagerCaller) GetProviderPrices(opts *bind.CallOpts, _providerAddr common.Address) (ProviderPrices, error) {
	var out []interface{}
	err := _IProviderManager.contract.Call(opts, &out, "getProviderPrices", _providerAddr)

	if err != nil {
		return *new(ProviderPrices), err
	}

	out0 := *abi.ConvertType(out[0], new(ProviderPrices)).(*ProviderPrices)

	return out0, err

}

// GetProviderPrices is a free data retrieval call binding the contract method 0x106859b6.
//
// Solidity: function getProviderPrices(address _providerAddr) view returns((uint256,uint256,uint256))
func (_IProviderManager *IProviderManagerSession) GetProviderPrices(_providerAddr common.Address) (ProviderPrices, error) {
	return _IProviderManager.Contract.GetProviderPrices(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProviderPrices is a free data retrieval call binding the contract method 0x106859b6.
//
// Solidity: function getProviderPrices(address _providerAddr) view returns((uint256,uint256,uint256))
func (_IProviderManager *IProviderManagerCallerSession) GetProviderPrices(_providerAddr common.Address) (ProviderPrices, error) {
	return _IProviderManager.Contract.GetProviderPrices(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProviderStatus is a free data retrieval call binding the contract method 0xd646a1da.
//
// Solidity: function getProviderStatus(address _providerAddr) view returns(uint8 _status)
func (_IProviderManager *IProviderManagerCaller) GetProviderStatus(opts *bind.CallOpts, _providerAddr common.Address) (uint8, error) {
	var out []interface{}
	err := _IProviderManager.contract.Call(opts, &out, "getProviderStatus", _providerAddr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProviderStatus is a free data retrieval call binding the contract method 0xd646a1da.
//
// Solidity: function getProviderStatus(address _providerAddr) view returns(uint8 _status)
func (_IProviderManager *IProviderManagerSession) GetProviderStatus(_providerAddr common.Address) (uint8, error) {
	return _IProviderManager.Contract.GetProviderStatus(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProviderStatus is a free data retrieval call binding the contract method 0xd646a1da.
//
// Solidity: function getProviderStatus(address _providerAddr) view returns(uint8 _status)
func (_IProviderManager *IProviderManagerCallerSession) GetProviderStatus(_providerAddr common.Address) (uint8, error) {
	return _IProviderManager.Contract.GetProviderStatus(&_IProviderManager.CallOpts, _providerAddr)
}

// HasJoined is a free data retrieval call binding the contract method 0x877f4e12.
//
// Solidity: function hasJoined(address _providerAddr) view returns(bool)
func (_IProviderManager *IProviderManagerCaller) HasJoined(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _IProviderManager.contract.Call(opts, &out, "hasJoined", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasJoined is a free data retrieval call binding the contract method 0x877f4e12.
//
// Solidity: function hasJoined(address _providerAddr) view returns(bool)
func (_IProviderManager *IProviderManagerSession) HasJoined(_providerAddr common.Address) (bool, error) {
	return _IProviderManager.Contract.HasJoined(&_IProviderManager.CallOpts, _providerAddr)
}

// HasJoined is a free data retrieval call binding the contract method 0x877f4e12.
//
// Solidity: function hasJoined(address _providerAddr) view returns(bool)
func (_IProviderManager *IProviderManagerCallerSession) HasJoined(_providerAddr common.Address) (bool, error) {
	return _IProviderManager.Contract.HasJoined(&_IProviderManager.CallOpts, _providerAddr)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactor) Approve(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "approve", _providerAddr)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerSession) Approve(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Approve(&_IProviderManager.TransactOpts, _providerAddr)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactorSession) Approve(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Approve(&_IProviderManager.TransactOpts, _providerAddr)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactor) Ban(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "ban", _providerAddr)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerSession) Ban(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Ban(&_IProviderManager.TransactOpts, _providerAddr)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactorSession) Ban(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Ban(&_IProviderManager.TransactOpts, _providerAddr)
}

// IncJobCount is a paid mutator transaction binding the contract method 0x3f6edb5f.
//
// Solidity: function incJobCount(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactor) IncJobCount(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "incJobCount", _providerAddr)
}

// IncJobCount is a paid mutator transaction binding the contract method 0x3f6edb5f.
//
// Solidity: function incJobCount(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerSession) IncJobCount(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.IncJobCount(&_IProviderManager.TransactOpts, _providerAddr)
}

// IncJobCount is a paid mutator transaction binding the contract method 0x3f6edb5f.
//
// Solidity: function incJobCount(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactorSession) IncJobCount(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.IncJobCount(&_IProviderManager.TransactOpts, _providerAddr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactor) Kick(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "kick", _providerAddr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerSession) Kick(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Kick(&_IProviderManager.TransactOpts, _providerAddr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactorSession) Kick(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Kick(&_IProviderManager.TransactOpts, _providerAddr)
}

// Register is a paid mutator transaction binding the contract method 0x94f9b055.
//
// Solidity: function register((uint64,uint64[],uint64[],uint64[]) _hardware, (uint256,uint256,uint256) _prices, (string,string)[] _labels) returns()
func (_IProviderManager *IProviderManagerTransactor) Register(opts *bind.TransactOpts, _hardware ProviderHardware, _prices ProviderPrices, _labels []Label) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "register", _hardware, _prices, _labels)
}

// Register is a paid mutator transaction binding the contract method 0x94f9b055.
//
// Solidity: function register((uint64,uint64[],uint64[],uint64[]) _hardware, (uint256,uint256,uint256) _prices, (string,string)[] _labels) returns()
func (_IProviderManager *IProviderManagerSession) Register(_hardware ProviderHardware, _prices ProviderPrices, _labels []Label) (*types.Transaction, error) {
	return _IProviderManager.Contract.Register(&_IProviderManager.TransactOpts, _hardware, _prices, _labels)
}

// Register is a paid mutator transaction binding the contract method 0x94f9b055.
//
// Solidity: function register((uint64,uint64[],uint64[],uint64[]) _hardware, (uint256,uint256,uint256) _prices, (string,string)[] _labels) returns()
func (_IProviderManager *IProviderManagerTransactorSession) Register(_hardware ProviderHardware, _prices ProviderPrices, _labels []Label) (*types.Transaction, error) {
	return _IProviderManager.Contract.Register(&_IProviderManager.TransactOpts, _hardware, _prices, _labels)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x4be49776.
//
// Solidity: function registerProvider(address _providerAddr, (uint64,uint64[],uint64[],uint64[]) _hardware, (uint256,uint256,uint256) _prices, (string,string)[] _labels) returns()
func (_IProviderManager *IProviderManagerTransactor) RegisterProvider(opts *bind.TransactOpts, _providerAddr common.Address, _hardware ProviderHardware, _prices ProviderPrices, _labels []Label) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "registerProvider", _providerAddr, _hardware, _prices, _labels)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x4be49776.
//
// Solidity: function registerProvider(address _providerAddr, (uint64,uint64[],uint64[],uint64[]) _hardware, (uint256,uint256,uint256) _prices, (string,string)[] _labels) returns()
func (_IProviderManager *IProviderManagerSession) RegisterProvider(_providerAddr common.Address, _hardware ProviderHardware, _prices ProviderPrices, _labels []Label) (*types.Transaction, error) {
	return _IProviderManager.Contract.RegisterProvider(&_IProviderManager.TransactOpts, _providerAddr, _hardware, _prices, _labels)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x4be49776.
//
// Solidity: function registerProvider(address _providerAddr, (uint64,uint64[],uint64[],uint64[]) _hardware, (uint256,uint256,uint256) _prices, (string,string)[] _labels) returns()
func (_IProviderManager *IProviderManagerTransactorSession) RegisterProvider(_providerAddr common.Address, _hardware ProviderHardware, _prices ProviderPrices, _labels []Label) (*types.Transaction, error) {
	return _IProviderManager.Contract.RegisterProvider(&_IProviderManager.TransactOpts, _providerAddr, _hardware, _prices, _labels)
}

// Reinstate is a paid mutator transaction binding the contract method 0x830aebd9.
//
// Solidity: function reinstate(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactor) Reinstate(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "reinstate", _providerAddr)
}

// Reinstate is a paid mutator transaction binding the contract method 0x830aebd9.
//
// Solidity: function reinstate(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerSession) Reinstate(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Reinstate(&_IProviderManager.TransactOpts, _providerAddr)
}

// Reinstate is a paid mutator transaction binding the contract method 0x830aebd9.
//
// Solidity: function reinstate(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactorSession) Reinstate(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.Reinstate(&_IProviderManager.TransactOpts, _providerAddr)
}

// RemoveProvider is a paid mutator transaction binding the contract method 0x8a355a57.
//
// Solidity: function removeProvider(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactor) RemoveProvider(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "removeProvider", _providerAddr)
}

// RemoveProvider is a paid mutator transaction binding the contract method 0x8a355a57.
//
// Solidity: function removeProvider(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerSession) RemoveProvider(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.RemoveProvider(&_IProviderManager.TransactOpts, _providerAddr)
}

// RemoveProvider is a paid mutator transaction binding the contract method 0x8a355a57.
//
// Solidity: function removeProvider(address _providerAddr) returns()
func (_IProviderManager *IProviderManagerTransactorSession) RemoveProvider(_providerAddr common.Address) (*types.Transaction, error) {
	return _IProviderManager.Contract.RemoveProvider(&_IProviderManager.TransactOpts, _providerAddr)
}

// IProviderManagerHardwareUpdatedEventIterator is returned from FilterHardwareUpdatedEvent and is used to iterate over the raw logs and unpacked data for HardwareUpdatedEvent events raised by the IProviderManager contract.
type IProviderManagerHardwareUpdatedEventIterator struct {
	Event *IProviderManagerHardwareUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *IProviderManagerHardwareUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProviderManagerHardwareUpdatedEvent)
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
		it.Event = new(IProviderManagerHardwareUpdatedEvent)
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
func (it *IProviderManagerHardwareUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProviderManagerHardwareUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProviderManagerHardwareUpdatedEvent represents a HardwareUpdatedEvent event raised by the IProviderManager contract.
type IProviderManagerHardwareUpdatedEvent struct {
	ProviderAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHardwareUpdatedEvent is a free log retrieval operation binding the contract event 0x3aeb53b0dee89ac04567fa6305e626e8d5246b478acd34d0a217507b9dfd076c.
//
// Solidity: event HardwareUpdatedEvent(address _providerAddr)
func (_IProviderManager *IProviderManagerFilterer) FilterHardwareUpdatedEvent(opts *bind.FilterOpts) (*IProviderManagerHardwareUpdatedEventIterator, error) {

	logs, sub, err := _IProviderManager.contract.FilterLogs(opts, "HardwareUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return &IProviderManagerHardwareUpdatedEventIterator{contract: _IProviderManager.contract, event: "HardwareUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchHardwareUpdatedEvent is a free log subscription operation binding the contract event 0x3aeb53b0dee89ac04567fa6305e626e8d5246b478acd34d0a217507b9dfd076c.
//
// Solidity: event HardwareUpdatedEvent(address _providerAddr)
func (_IProviderManager *IProviderManagerFilterer) WatchHardwareUpdatedEvent(opts *bind.WatchOpts, sink chan<- *IProviderManagerHardwareUpdatedEvent) (event.Subscription, error) {

	logs, sub, err := _IProviderManager.contract.WatchLogs(opts, "HardwareUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProviderManagerHardwareUpdatedEvent)
				if err := _IProviderManager.contract.UnpackLog(event, "HardwareUpdatedEvent", log); err != nil {
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

// ParseHardwareUpdatedEvent is a log parse operation binding the contract event 0x3aeb53b0dee89ac04567fa6305e626e8d5246b478acd34d0a217507b9dfd076c.
//
// Solidity: event HardwareUpdatedEvent(address _providerAddr)
func (_IProviderManager *IProviderManagerFilterer) ParseHardwareUpdatedEvent(log types.Log) (*IProviderManagerHardwareUpdatedEvent, error) {
	event := new(IProviderManagerHardwareUpdatedEvent)
	if err := _IProviderManager.contract.UnpackLog(event, "HardwareUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProviderManagerProviderStatusChangedIterator is returned from FilterProviderStatusChanged and is used to iterate over the raw logs and unpacked data for ProviderStatusChanged events raised by the IProviderManager contract.
type IProviderManagerProviderStatusChangedIterator struct {
	Event *IProviderManagerProviderStatusChanged // Event containing the contract specifics and raw log

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
func (it *IProviderManagerProviderStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProviderManagerProviderStatusChanged)
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
		it.Event = new(IProviderManagerProviderStatusChanged)
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
func (it *IProviderManagerProviderStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProviderManagerProviderStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProviderManagerProviderStatusChanged represents a ProviderStatusChanged event raised by the IProviderManager contract.
type IProviderManagerProviderStatusChanged struct {
	ProviderAddr common.Address
	Status       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProviderStatusChanged is a free log retrieval operation binding the contract event 0x824cd9acaf7b71614920a4b2651c4aad10142580262655e42fa5a66833bf5401.
//
// Solidity: event ProviderStatusChanged(address _providerAddr, uint8 status)
func (_IProviderManager *IProviderManagerFilterer) FilterProviderStatusChanged(opts *bind.FilterOpts) (*IProviderManagerProviderStatusChangedIterator, error) {

	logs, sub, err := _IProviderManager.contract.FilterLogs(opts, "ProviderStatusChanged")
	if err != nil {
		return nil, err
	}
	return &IProviderManagerProviderStatusChangedIterator{contract: _IProviderManager.contract, event: "ProviderStatusChanged", logs: logs, sub: sub}, nil
}

// WatchProviderStatusChanged is a free log subscription operation binding the contract event 0x824cd9acaf7b71614920a4b2651c4aad10142580262655e42fa5a66833bf5401.
//
// Solidity: event ProviderStatusChanged(address _providerAddr, uint8 status)
func (_IProviderManager *IProviderManagerFilterer) WatchProviderStatusChanged(opts *bind.WatchOpts, sink chan<- *IProviderManagerProviderStatusChanged) (event.Subscription, error) {

	logs, sub, err := _IProviderManager.contract.WatchLogs(opts, "ProviderStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProviderManagerProviderStatusChanged)
				if err := _IProviderManager.contract.UnpackLog(event, "ProviderStatusChanged", log); err != nil {
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

// ParseProviderStatusChanged is a log parse operation binding the contract event 0x824cd9acaf7b71614920a4b2651c4aad10142580262655e42fa5a66833bf5401.
//
// Solidity: event ProviderStatusChanged(address _providerAddr, uint8 status)
func (_IProviderManager *IProviderManagerFilterer) ParseProviderStatusChanged(log types.Log) (*IProviderManagerProviderStatusChanged, error) {
	event := new(IProviderManagerProviderStatusChanged)
	if err := _IProviderManager.contract.UnpackLog(event, "ProviderStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220feec7b17524d34f3c20d4922777dea9e0c157695b6e02a8bd0951b2ebf853fa864736f6c63430008110033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MetaSchedulerMetaData contains all meta data concerning the MetaScheduler contract.
var MetaSchedulerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_credit\",\"type\":\"address\"},{\"internalType\":\"contractConstants\",\"name\":\"_constants\",\"type\":\"address\"},{\"internalType\":\"contractIProviderManager\",\"name\":\"_providerManager\",\"type\":\"address\"},{\"internalType\":\"contractIProviderJobQueues\",\"name\":\"_providerJobQueues\",\"type\":\"address\"},{\"internalType\":\"contractIJobRepository\",\"name\":\"_jobs\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CustomerMetaSchedulerProviderOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"CustomerOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJobDefinition\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"JobHotStatusOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"JobProviderOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"MetaScheduledScheduledStatusOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewJobRequestDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpendingAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNotJoined\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"RemainingTimeAboveLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"RunningColdStatusOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"RunningScheduledStatusOnly\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_billingAmount\",\"type\":\"uint256\"}],\"name\":\"BilledTooMuchEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"ClaimNextCancellingJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"}],\"name\":\"ClaimNextTopUpJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumJobStatus\",\"name\":\"_from\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumJobStatus\",\"name\":\"_to\",\"type\":\"uint8\"}],\"name\":\"JobTransitionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METASCHEDULER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextCancellingJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextTopUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"constants\",\"outputs\":[{\"internalType\":\"contractConstants\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRequestNewJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"contractIJobRepository\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_lastError\",\"type\":\"string\"}],\"name\":\"panicJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerJobQueues\",\"outputs\":[{\"internalType\":\"contractIProviderJobQueues\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerManager\",\"outputs\":[{\"internalType\":\"contractIProviderManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_nextJobStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_jobDurationMinute\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"_lastError\",\"type\":\"string\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_lockedCredits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_delegateSpendingAuthority\",\"type\":\"bool\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_delegateSpendingAuthority\",\"type\":\"bool\"}],\"name\":\"setDelegateSpendingAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setEnableRequestNewJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"timeoutJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"topUpJobDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040526002805460ff191660011790553480156200001f57600080fd5b506040516200535f3803806200535f833981016040819052620000429162000137565b60018055620000536000336200007d565b6001600160a01b0394851660805291841660a052831660c052908216610100521660e052620001b7565b6000828152602081815260408083206001600160a01b038516845290915290205460ff166200011a576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620000d93390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6001600160a01b03811681146200013457600080fd5b50565b600080600080600060a086880312156200015057600080fd5b85516200015d816200011e565b602087015190955062000170816200011e565b604087015190945062000183816200011e565b606087015190935062000196816200011e565b6080870151909250620001a9816200011e565b809150509295509295909350565b60805160a05160c05160e05161010051614fc6620003996000396000818161036101528181611ba701528181611d3501528181611e0501528181612157015281816121e101528181612a6001526132550152600081816103880152818161059301528181610724015281816107d2015281816108af01528181610a9101528181610c1601528181610d3f01528181610e2f01528181610f50015281816110790152818161110a01528181611190015281816114c401528181611586015281816116af01528181611792015281816118130152818161195801528181612007015281816123d70152818161269c015281816128ae015281816129e501528181612c9701528181612d2701528181612ed40152818161348c01528181613512015281816136280152613751015260008181610452015281816112ca0152818161140c015281816118d10152818161260301528181612fb101526137e3015260008181610229015281816104ea015281816109e8015281816112220152818161136401528181611b3801528181611f570152818161255b01528181612751015281816129470152818161305501526131df0152600081816103c2015281816120ba01528181612489015281816132fc015281816133b5015261386b0152614fc66000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80635e1b2d6511610104578063a217fddf116100a2578063d994378811610071578063d994378814610425578063e052888c14610438578063e2eaf3e71461044d578063ebd4bf001461047457600080fd5b8063a217fddf146103e4578063d1cee546146103ec578063d547741f146103ff578063d77836ce1461041257600080fd5b806372de5b2f116100de57806372de5b2f1461035c5780637c8fce231461038357806391d14854146103aa578063a06d083c146103bd57600080fd5b80635e1b2d651461032e5780635fae14501461033657806369ee1bf91461034957600080fd5b80632f2ff15d1161017157806336568abe1161014b57806336568abe146102f35780634c6dad121461030657806354b4a0d2146103135780635d3a71801461032657600080fd5b80632f2ff15d146102ba5780632fecc4f6146102cd578063329af326146102e057600080fd5b806313151ec9116101ad57806313151ec91461022457806318263c59146102635780631f92a63f14610276578063248a9ca31461028957600080fd5b806301ffc9a7146101d45780630bba4dc5146101fc5780630e4bf0e814610211575b600080fd5b6101e76101e2366004613bac565b61047c565b60405190151581526020015b60405180910390f35b61020f61020a366004613bf4565b6104b3565b005b61020f61021f366004613d6f565b6104d2565b61024b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101f3565b61020f610271366004613dda565b610874565b61020f610284366004613dda565b6109d0565b6102ac610297366004613dda565b60009081526020819052604090206001015490565b6040519081526020016101f3565b61020f6102c8366004613e08565b610bbf565b61020f6102db366004613e38565b610be4565b61020f6102ee366004613e5a565b610bf3565b61020f610301366004613e08565b610e9b565b6002546101e79060ff1681565b61020f610321366004613e7f565b610f15565b61020f61120a565b61020f61134c565b61020f610344366004613dda565b611563565b61020f610357366004613dda565b61193f565b61024b7f000000000000000000000000000000000000000000000000000000000000000081565b61024b7f000000000000000000000000000000000000000000000000000000000000000081565b6101e76103b8366004613e08565b611ef5565b61024b7f000000000000000000000000000000000000000000000000000000000000000081565b6102ac600081565b61020f6103fa366004613e08565b611f1e565b61020f61040d366004613e08565b61206e565b61020f610420366004613dda565b612093565b6102ac6104333660046140db565b61212f565b6102ac600080516020614f7183398151915281565b61024b7f000000000000000000000000000000000000000000000000000000000000000081565b61020f612543565b60006001600160e01b03198216637965db0b60e01b14806104ad57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60006104be81612918565b506002805460ff1916911515919091179055565b6040516343bfa70960e11b81523360048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063877f4e1290602401602060405180830381865afa158015610539573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061055d9190614214565b61057a5760405163ef341f6d60e01b815260040160405180910390fd5b60405163023aa9ab60e61b8152600481018690526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa1580156105e2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261060a919081019061463b565b60608101519091506001600160a01b03163314610656576060810151604051630cb8c19760e21b81523360048201526001600160a01b0390911660248201526044015b60405180910390fd5b600385600881111561066a5761066a614750565b141580156106ed575060405163da498b2960e01b815273__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__9063da498b29906106aa90889060040161477a565b602060405180830381865af41580156106c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106eb9190614214565b155b1561070d57846040516314475eb760e01b815260040161064d919061477a565b60405163b613a72160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b613a7219061075b90899087906004016147b4565b600060405180830381600087803b15801561077557600080fd5b505af1158015610789573d6000803e3d6000fd5b5060089250610796915050565b8560088111156107a8576107a8614750565b0361083a5760c0810180514260809091015251604051633baa6cb560e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691633baa6cb591610807918a916004016147cd565b600060405180830381600087803b15801561082157600080fd5b505af1158015610835573d6000803e3d6000fd5b505050505b60038160200151600881111561085257610852614750565b0361086257610862863386612925565b61086c8686612d06565b505050505050565b600080516020614f7183398151915261088c81612918565b60405163023aa9ab60e61b81526004810183905282906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eaa6ac090602401600060405180830381865afa1580156108f6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261091e919081019061463b565b602081015160405163d55388b960e01b815291925073__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__9163d55388b99161095b9160040161477a565b602060405180830381865af4158015610978573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099c9190614214565b6109bf578060200151604051634634126160e11b815260040161064d919061477a565b6109ca846000612d06565b50505050565b6040516343bfa70960e11b81523360048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063877f4e1290602401602060405180830381865afa158015610a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5b9190614214565b610a785760405163ef341f6d60e01b815260040160405180910390fd5b60405163023aa9ab60e61b8152600481018390526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa158015610ae0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b08919081019061463b565b60608101519091506001600160a01b03163314610b4f576060810151604051630cb8c19760e21b81523360048201526001600160a01b03909116602482015260440161064d565b600281602001516008811115610b6757610b67614750565b14158015610b8b5750600181602001516008811115610b8857610b88614750565b14155b15610baf57806020015160405163048389ff60e11b815260040161064d919061477a565b610bba836000612d06565b505050565b600082815260208190526040902060010154610bda81612918565b610bba838361358b565b610bef82338361360f565b5050565b60405163023aa9ab60e61b81526004810183905282906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eaa6ac090602401600060405180830381865afa158015610c5d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c85919081019061463b565b602081015160405163d55388b960e01b815291925073__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__9163d55388b991610cc29160040161477a565b602060405180830381865af4158015610cdf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d039190614214565b610d26578060200151604051634634126160e11b815260040161064d919061477a565b60405163023aa9ab60e61b8152600481018590526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa158015610d8e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610db6919081019061463b565b905080604001516001600160a01b0316336001600160a01b031614610e05576040808201519051638942331960e01b81523360048201526001600160a01b03909116602482015260440161064d565b60a0810180518515156060909101528151905160405163b709033160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263b709033192610e629260040161480c565b600060405180830381600087803b158015610e7c57600080fd5b505af1158015610e90573d6000803e3d6000fd5b505050505050505050565b6001600160a01b0381163314610f0b5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b606482015260840161064d565b610bef82826138e1565b600080516020614f71833981519152610f2d81612918565b60405163023aa9ab60e61b81526004810184905283906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eaa6ac090602401600060405180830381865afa158015610f97573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fbf919081019061463b565b602081015160405163d55388b960e01b815291925073__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__9163d55388b991610ffc9160040161477a565b602060405180830381865af4158015611019573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103d9190614214565b611060578060200151604051634634126160e11b815260040161064d919061477a565b60405163023aa9ab60e61b8152600481018690526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa1580156110c8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110f0919081019061463b565b60405163b613a72160e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b613a7219061114190899089906004016147b4565b600060405180830381600087803b15801561115b57600080fd5b505af115801561116f573d6000803e3d6000fd5b5050505060c0810180514260809091015251604051633baa6cb560e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691633baa6cb5916111cd918a916004016147cd565b600060405180830381600087803b1580156111e757600080fd5b505af11580156111fb573d6000803e3d6000fd5b5050505061086c866008612d06565b6040516343bfa70960e11b81523360048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063877f4e1290602401602060405180830381865afa158015611271573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112959190614214565b6112b25760405163ef341f6d60e01b815260040160405180910390fd5b604051632f5a6afd60e21b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063bd69abf4906024016020604051808303816000875af115801561131b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061133f9190614843565b9050610bef816002612d06565b6040516343bfa70960e11b81523360048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063877f4e1290602401602060405180830381865afa1580156113b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d79190614214565b6113f45760405163ef341f6d60e01b815260040160405180910390fd5b604051636f0d26a360e11b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063de1a4d46906024016020604051808303816000875af115801561145d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114819190614843565b60405163023aa9ab60e61b8152600481018290529091507f290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eaa6ac090602401600060405180830381865afa15801561150b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611533919081019061463b565b60409081015181516001600160a01b03909116815233602082015290810183905260600160405180910390a15050565b60405163023aa9ab60e61b81526004810182905281906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eaa6ac090602401600060405180830381865afa1580156115cd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115f5919081019061463b565b602081015160405163d55388b960e01b815291925073__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__9163d55388b9916116329160040161477a565b602060405180830381865af415801561164f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116739190614214565b611696578060200151604051634634126160e11b815260040161064d919061477a565b60405163023aa9ab60e61b8152600481018490526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa1580156116fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611726919081019061463b565b905080604001516001600160a01b0316336001600160a01b031614611775576040808201519051638942331960e01b81523360048201526001600160a01b03909116602482015260440161064d565b604051634726f0e560e11b815260048101859052600160248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638e4de1ca90604401600060405180830381600087803b1580156117de57600080fd5b505af11580156117f2573d6000803e3d6000fd5b5050505060c0810180514260409182015290519051633baa6cb560e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691633baa6cb5916118509188916004016147cd565b600060405180830381600087803b15801561186a57600080fd5b505af115801561187e573d6000803e3d6000fd5b506003925061188b915050565b816020015160088111156118a1576118a1614750565b036119345760608101516040516315b9291760e11b81526001600160a01b039182166004820152602481018690527f000000000000000000000000000000000000000000000000000000000000000090911690632b72522e90604401600060405180830381600087803b15801561191757600080fd5b505af115801561192b573d6000803e3d6000fd5b505050506109ca565b6109ca846004612d06565b60405163023aa9ab60e61b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa1580156119a7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119cf919081019061463b565b905080604001516001600160a01b0316336001600160a01b031614158015611a0d575080606001516001600160a01b0316336001600160a01b031614155b8015611a2e5750611a2c600080516020614f7183398151915233611ef5565b155b15611a4c57604051637897ef6d60e01b815260040160405180910390fd5b8060a0015160600151611a72576040516342f9901960e01b815260040160405180910390fd5b6020810151604051630f675b0760e41b815273__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__9163f675b07091611aad919060040161477a565b602060405180830381865af4158015611aca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aee9190614214565b611b1157806020015160405163ed3e2aad60e01b815260040161064d919061477a565b60608101516040516308342cdb60e11b81526001600160a01b0391821660048201526000917f0000000000000000000000000000000000000000000000000000000000000000169063106859b690602401606060405180830381865afa158015611b7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ba3919061485c565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638ce9843b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c2791906148b7565b6001600160401b031673__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__63f4cc70ac84846040518363ffffffff1660e01b8152600401611c69929190614a99565b602060405180830381865af4158015611c86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611caa91906148b7565b6001600160401b03161115611de157604051633d331c2b60e21b815273__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__9063f4cc70ac90611cf29085908590600401614a99565b602060405180830381865af4158015611d0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d3391906148b7565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638ce9843b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611db591906148b7565b604051634801db4560e11b81526001600160401b0392831660048201529116602482015260440161064d565b600073__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__6334d515f58460800151847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638ce9843b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e8591906148b7565b6040518463ffffffff1660e01b8152600401611ea393929190614bc8565b602060405180830381865af4158015611ec0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ee49190614843565b90506109ca8484604001518361360f565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020614f71833981519152611f3681612918565b6040516343bfa70960e11b81526001600160a01b03808416600483015283917f00000000000000000000000000000000000000000000000000000000000000009091169063877f4e1290602401602060405180830381865afa158015611fa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc49190614214565b611fe15760405163ef341f6d60e01b815260040160405180910390fd5b604051635aae4bbd60e01b8152600481018590526001600160a01b0384811660248301527f00000000000000000000000000000000000000000000000000000000000000001690635aae4bbd90604401600060405180830381600087803b15801561204b57600080fd5b505af115801561205f573d6000803e3d6000fd5b505050506109ca846001612d06565b60008281526020819052604090206001015461208981612918565b610bba83836138e1565b600061209e81612918565b60405163a9059cbb60e01b8152336004820152602481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a9059cbb906044016020604051808303816000875af115801561210b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bba9190614214565b60025460009060ff1661215557604051633abe75b360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bb0c82986040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121d79190614843565b84101561228257837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663bb0c82986040518163ffffffff1660e01b8152600401602060405180830381865afa15801561223d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122619190614843565b60405162fae2d560e21b81526004810192909252602482015260440161064d565b604051639e71f7a160e01b815273__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__90639e71f7a1906122b9908890600401614c15565b602060405180830381865af41580156122d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122fa9190614214565b61231757604051630d208e5960e41b815260040160405180910390fd5b60408051610140810190915260008082529060208101828152602001336001600160a01b0316815260200160006001600160a01b031681526020018781526020016040518060800160405280888152602001600081526020016000815260200186151581525081526020016040518060a001604052804281526020014281526020014281526020014381526020016000815250815260200185815260200160001515815260200160405180602001604052806000815250815250905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663ce1ecc5a836040518263ffffffff1660e01b81526004016124219190614c28565b6020604051808303816000875af1158015612440573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124649190614843565b6040516323b872dd60e01b8152336004820152306024820152604481018890529091507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af11580156124da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124fe9190614214565b50604080518281523360208201527f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3910160405180910390a19150505b949350505050565b6040516343bfa70960e11b81523360048201819052907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063877f4e1290602401602060405180830381865afa1580156125aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125ce9190614214565b6125eb5760405163ef341f6d60e01b815260040160405180910390fd5b6040516308b4e63560e21b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906322d398d4906024016020604051808303816000875af1158015612654573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126789190614843565b60405163023aa9ab60e61b8152600481018290529091506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eaa6ac090602401600060405180830381865afa1580156126e3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261270b919081019061463b565b90507fa42f2b4a7ee7f91857a4c98fc71fc48546a284d5db48dd77b7ab81030a494470823373__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__63c4b2bd2c85608001517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663106859b688606001516040518263ffffffff1660e01b81526004016127ae91906001600160a01b0391909116815260200190565b606060405180830381865afa1580156127cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127ef919061485c565b8760a00151604001516040518463ffffffff1660e01b815260040161281693929190614d3a565b602060405180830381865af4158015612833573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061285791906148b7565b604080519384526001600160a01b0390921660208401526001600160401b03169082015260600160405180910390a160a081018051600060409182015282519151905163b709033160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263b7090331926128e19260040161480c565b600060405180830381600087803b1580156128fb57600080fd5b505af115801561290f573d6000803e3d6000fd5b50505050505050565b6129228133613946565b50565b6040516308342cdb60e11b81526001600160a01b0383811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063106859b690602401606060405180830381865afa158015612990573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129b4919061485c565b9050816001600160401b03166000036129cc57600191505b60405163023aa9ab60e61b8152600481018590526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa158015612a34573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612a5c919081019061463b565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b95901716040518163ffffffff1660e01b8152600401602060405180830381865afa158015612abc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ae09190614843565b60c082015151603c90612af39042614d94565b612afd9190614da7565b612b079190614dc9565b836001600160401b03161115612b8757604080518681526001600160a01b03861660208201526001600160401b0385168183015290517f17e65314b087df225f56701d0a66a3f7d9ce0f26077307b4b765a19c60a36d449181900360600190a160c081015151603c90612b7a9042614d94565b612b849190614da7565b92505b612b9283603c614ddc565b6001600160401b03168160c001516000015142612baf9190614d94565b10612bd957612bbf83603c614ddc565b612bd2906001600160401b031642614d94565b60c0820151525b60808101516040516334d515f560e01b815260009173__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__916334d515f591612c1a9187908990600401614bc8565b602060405180830381865af4158015612c37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c5b9190614843565b60a083015151909150811115612c73575060a0810151515b60a0820180516020018290525160405163b709033160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163b709033191612ccc918a9160040161480c565b600060405180830381600087803b158015612ce657600080fd5b505af1158015612cfa573d6000803e3d6000fd5b50505050505050505050565b612d0e61399f565b60405163023aa9ab60e61b8152600481018390526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa158015612d76573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612d9e919081019061463b565b60208101516040516397d2874f60e01b815291925073__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__916397d2874f91612ddd918690600401614e07565b60006040518083038186803b158015612df557600080fd5b505af4158015612e09573d6000803e3d6000fd5b505050507f0bba917f0a1e0fc0d51a75273e7088a4dfecb010699e60ac9c58526429f6c37f83826020015184604051612e4493929190614e22565b60405180910390a16000826008811115612e6057612e60614750565b03612f6d5760608082015160408084015181518781526001600160a01b0393841660208201529216908201527f50d9c3fab9ef0192905beb84254b4ffb6fe086795cc23de484ec65947b6615a2910160405180910390a1604051635aae4bbd60e01b815260048101849052600060248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635aae4bbd90604401600060405180830381600087803b158015612f2057600080fd5b505af1158015612f34573d6000803e3d6000fd5b505050506040518060a0016040528042815260200142815260200142815260200143815260200160008152508160c00181905250613475565b6001826008811115612f8157612f81614750565b036130155760608101516040516365b4fc4960e01b81526001600160a01b039182166004820152602481018590527f0000000000000000000000000000000000000000000000000000000000000000909116906365b4fc49906044015b600060405180830381600087803b158015612ff857600080fd5b505af115801561300c573d6000803e3d6000fd5b50505050613475565b600282600881111561302957613029614750565b0361319a5760608101516040516308342cdb60e11b81526001600160a01b0391821660048201526000917f0000000000000000000000000000000000000000000000000000000000000000169063106859b690602401606060405180830381865afa15801561309c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130c0919061485c565b90507fc8d7c7c2914e1aa1462fe7999f8a18a0f1043d7d94ab692c3ac9af846f1be8c7826040015183606001518673__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__63c4b2bd2c8760800151878960a00151600001516040518463ffffffff1660e01b815260040161313593929190614d3a565b602060405180830381865af4158015613152573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061317691906148b7565b866080015160405161318c959493929190614e43565b60405180910390a150613475565b60038260088111156131ae576131ae614750565b036132105760c08101514290526060810151604051633f6edb5f60e01b81526001600160a01b0391821660048201527f000000000000000000000000000000000000000000000000000000000000000090911690633f6edb5f90602401612fde565b60c08101514260209091015260038160200151600881111561323457613234614750565b146132465760c0810151602081015190525b600060648260a00151602001517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a234d90f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132d59190614843565b6132df9190614e90565b6132e99190614da7565b60a083015160200151909150156133b3577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb8360600151838560a00151602001516133429190614d94565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561338d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133b19190614214565b505b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb83604001518460a00151602001518560a00151600001516134039190614d94565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561344e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134729190614214565b50505b6040516370c8433b60e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063e1908676906134c39086908690600401614ea7565b600060405180830381600087803b1580156134dd57600080fd5b505af11580156134f1573d6000803e3d6000fd5b5050505060c0810180514360609091015251604051633baa6cb560e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031691633baa6cb59161354f9187916004016147cd565b600060405180830381600087803b15801561356957600080fd5b505af115801561357d573d6000803e3d6000fd5b5050505050610bef60018055565b6135958282611ef5565b610bef576000828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556135cb3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60405163023aa9ab60e61b8152600481018490526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638eaa6ac090602401600060405180830381865afa158015613677573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261369f919081019061463b565b90506002816020015160088111156136b9576136b9614750565b141580156136dd57506003816020015160088111156136da576136da614750565b14155b1561370157806020015160405163ed3e2aad60e01b815260040161064d919061477a565b818160a001516000018181516137179190614dc9565b90525060a08101516040018051839190613732908390614dc9565b90525060a081015160405163b709033160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163b70903319161378691889160040161480c565b600060405180830381600087803b1580156137a057600080fd5b505af11580156137b4573d6000803e3d6000fd5b5050505060608101516040516327fed93160e01b81526001600160a01b039182166004820152602481018690527f0000000000000000000000000000000000000000000000000000000000000000909116906327fed93190604401600060405180830381600087803b15801561382957600080fd5b505af115801561383d573d6000803e3d6000fd5b50506040516323b872dd60e01b81526001600160a01b038681166004830152306024830152604482018690527f00000000000000000000000000000000000000000000000000000000000000001692506323b872dd91506064016020604051808303816000875af11580156138b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138da9190614214565b5050505050565b6138eb8282611ef5565b15610bef576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6139508282611ef5565b610bef5761395d816139f8565b613968836020613a0a565b604051602001613979929190614ebb565b60408051601f198184030181529082905262461bcd60e51b825261064d91600401614f30565b6002600154036139f15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161064d565b6002600155565b60606104ad6001600160a01b03831660145b60606000613a19836002614e90565b613a24906002614dc9565b6001600160401b03811115613a3b57613a3b613c3e565b6040519080825280601f01601f191660200182016040528015613a65576020820181803683370190505b509050600360fc1b81600081518110613a8057613a80614f43565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110613aaf57613aaf614f43565b60200101906001600160f81b031916908160001a9053506000613ad3846002614e90565b613ade906001614dc9565b90505b6001811115613b56576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110613b1257613b12614f43565b1a60f81b828281518110613b2857613b28614f43565b60200101906001600160f81b031916908160001a90535060049490941c93613b4f81614f59565b9050613ae1565b508315613ba55760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161064d565b9392505050565b600060208284031215613bbe57600080fd5b81356001600160e01b031981168114613ba557600080fd5b801515811461292257600080fd5b8035613bef81613bd6565b919050565b600060208284031215613c0657600080fd5b8135613ba581613bd6565b6009811061292257600080fd5b6001600160401b038116811461292257600080fd5b8035613bef81613c1e565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715613c7657613c76613c3e565b60405290565b60405161010081016001600160401b0381118282101715613c7657613c76613c3e565b60405161014081016001600160401b0381118282101715613c7657613c76613c3e565b604051601f8201601f191681016001600160401b0381118282101715613cea57613cea613c3e565b604052919050565b60006001600160401b03821115613d0b57613d0b613c3e565b50601f01601f191660200190565b600082601f830112613d2a57600080fd5b8135613d3d613d3882613cf2565b613cc2565b818152846020838601011115613d5257600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215613d8557600080fd5b843593506020850135613d9781613c11565b92506040850135613da781613c1e565b915060608501356001600160401b03811115613dc257600080fd5b613dce87828801613d19565b91505092959194509250565b600060208284031215613dec57600080fd5b5035919050565b6001600160a01b038116811461292257600080fd5b60008060408385031215613e1b57600080fd5b823591506020830135613e2d81613df3565b809150509250929050565b60008060408385031215613e4b57600080fd5b50508035926020909101359150565b60008060408385031215613e6d57600080fd5b823591506020830135613e2d81613bd6565b60008060408385031215613e9257600080fd5b8235915060208301356001600160401b03811115613eaf57600080fd5b613ebb85828601613d19565b9150509250929050565b6005811061292257600080fd5b8035613bef81613ec5565b60006001600160401b03821115613ef657613ef6613c3e565b5060051b60200190565b600060408284031215613f1257600080fd5b613f1a613c54565b905081356001600160401b0380821115613f3357600080fd5b613f3f85838601613d19565b83526020840135915080821115613f5557600080fd5b50613f6284828501613d19565b60208301525092915050565b600082601f830112613f7f57600080fd5b81356020613f8f613d3883613edd565b82815260059290921b84018101918181019086841115613fae57600080fd5b8286015b84811015613fed5780356001600160401b03811115613fd15760008081fd5b613fdf8986838b0101613f00565b845250918301918301613fb2565b509695505050505050565b6001600160f01b03198116811461292257600080fd5b600082601f83011261401f57600080fd5b8135602061402f613d3883613edd565b82815260059290921b8401810191818101908684111561404e57600080fd5b8286015b84811015613fed5780356001600160401b03808211156140725760008081fd5b908801906040828b03601f190181131561408c5760008081fd5b614094613c54565b87840135838111156140a65760008081fd5b6140b48d8a83880101613f00565b82525092810135926140c584613ff8565b8088019390935250508352918301918301614052565b600080600080608085870312156140f157600080fd5b84356001600160401b038082111561410857600080fd5b90860190610100828903121561411d57600080fd5b614125613c7c565b61412e83613c33565b815261413c60208401613c33565b602082015261414d60408401613c33565b604082015261415e60608401613c33565b606082015260808301358281111561417557600080fd5b6141818a828601613d19565b60808301525061419360a08401613ed2565b60a082015260c0830135828111156141aa57600080fd5b6141b68a828601613f6e565b60c08301525060e0830135828111156141ce57600080fd5b6141da8a82860161400e565b60e0830152509550505060208501359250604085013591506141fe60608601613be4565b905092959194509250565b8051613bef81613bd6565b60006020828403121561422657600080fd5b8151613ba581613bd6565b8051613bef81613c11565b8051613bef81613df3565b8051613bef81613c1e565b60005b8381101561426d578181015183820152602001614255565b50506000910152565b600082601f83011261428757600080fd5b8151614295613d3882613cf2565b8181528460208386010111156142aa57600080fd5b61253b826020830160208701614252565b8051613bef81613ec5565b6000604082840312156142d857600080fd5b6142e0613c54565b905081516001600160401b03808211156142f957600080fd5b61430585838601614276565b8352602084015191508082111561431b57600080fd5b50613f6284828501614276565b600082601f83011261433957600080fd5b81516020614349613d3883613edd565b82815260059290921b8401810191818101908684111561436857600080fd5b8286015b84811015613fed5780516001600160401b0381111561438b5760008081fd5b6143998986838b01016142c6565b84525091830191830161436c565b600082601f8301126143b857600080fd5b815160206143c8613d3883613edd565b82815260059290921b840181019181810190868411156143e757600080fd5b8286015b84811015613fed5780516001600160401b038082111561440b5760008081fd5b908801906040828b03601f19018113156144255760008081fd5b61442d613c54565b878401518381111561443f5760008081fd5b61444d8d8a838801016142c6565b825250928101519261445e84613ff8565b80880193909352505083529183019183016143eb565b6000610100828403121561448757600080fd5b61448f613c7c565b905061449a82614247565b81526144a860208301614247565b60208201526144b960408301614247565b60408201526144ca60608301614247565b606082015260808201516001600160401b03808211156144e957600080fd5b6144f585838601614276565b608084015261450660a085016142bb565b60a084015260c084015191508082111561451f57600080fd5b61452b85838601614328565b60c084015260e084015191508082111561454457600080fd5b50614551848285016143a7565b60e08301525092915050565b60006080828403121561456f57600080fd5b604051608081018181106001600160401b038211171561459157614591613c3e565b806040525080915082518152602083015160208201526040830151604082015260608301516145bf81613bd6565b6060919091015292915050565b600060a082840312156145de57600080fd5b60405160a081018181106001600160401b038211171561460057614600613c3e565b806040525080915082518152602083015160208201526040830151604082015260608301516060820152608083015160808201525092915050565b60006020828403121561464d57600080fd5b81516001600160401b038082111561466457600080fd5b90830190610220828603121561467957600080fd5b614681613c9f565b8251815261469160208401614231565b60208201526146a26040840161423c565b60408201526146b36060840161423c565b60608201526080830151828111156146ca57600080fd5b6146d687828601614474565b6080830152506146e98660a0850161455d565b60a08201526101206146fd878286016145cc565b60c08301526101c084015160e083015261471a6101e08501614209565b6101008301526102008401518381111561473357600080fd5b61473f88828701614276565b918301919091525095945050505050565b634e487b7160e01b600052602160045260246000fd5b6009811061477657614776614750565b9052565b602081016104ad8284614766565b600081518084526147a0816020860160208601614252565b601f01601f19169290920160200192915050565b82815260406020820152600061253b6040830184614788565b82815260c08101613ba5602083018480518252602081015160208301526040810151604083015260608101516060830152608081015160808301525050565b82815260a08101613ba560208301848051825260208101516020830152604081015160408301526060810151151560608301525050565b60006020828403121561485557600080fd5b5051919050565b60006060828403121561486e57600080fd5b604051606081018181106001600160401b038211171561489057614890613c3e565b80604052508251815260208301516020820152604083015160408201528091505092915050565b6000602082840312156148c957600080fd5b8151613ba581613c1e565b6005811061477657614776614750565b60008151604084526148f96040850182614788565b9050602083015184820360208601526149128282614788565b95945050505050565b600081518084526020808501808196508360051b8101915082860160005b858110156149635782840389526149518483516148e4565b98850198935090840190600101614939565b5091979650505050505050565b600081518084526020808501808196508360051b8101915082860160005b858110156149635782840389528151604081518187526149b0828801826148e4565b928801516001600160f01b03191696880196909652509885019893509084019060010161498e565b60006101006001600160401b0383511684526020830151614a0460208601826001600160401b03169052565b506040830151614a1f60408601826001600160401b03169052565b506060830151614a3a60608601826001600160401b03169052565b506080830151816080860152614a5282860182614788565b91505060a0830151614a6760a08601826148d4565b5060c083015184820360c0860152614a7f828261491b565b91505060e083015184820360e08601526149128282614970565b608081528251608082015260006020840151614ab860a0840182614766565b5060408401516001600160a01b03811660c08401525060608401516001600160a01b03811660e084015250608084015161010061022081850152614b006102a08501836149d8565b915060a0860151610120614b39818701838051825260208101516020830152604081015160408301526060810151151560608301525050565b60c088015180516101a088015260208101516101c088015260408101516101e088015260608101516102008801526080015161022087015260e088015161024087015291870151151561026086015250850151838203607f1901610280850152614ba38282614788565b92505050613ba560208301848051825260208082015190830152604090810151910152565b60a081526000614bdb60a08301866149d8565b9050614bfe60208301858051825260208082015190830152604090810151910152565b6001600160401b0383166080830152949350505050565b602081526000613ba560208301846149d8565b602081528151602082015260006020830151614c476040840182614766565b5060408301516001600160a01b03811660608401525060608301516001600160a01b03811660808401525060808301516102208060a0850152614c8e6102408501836149d8565b60a0860151805160c0870152602081015160e08701526040810151610100870152606081015115156101208701529092505060c085015180516101408601526020810151610160860152604081015161018086015260608101516101a086015260808101516101c08601525060e08501516101e085015261010085015180151561020086015250610120850151848303601f190182860152614d308382614788565b9695505050505050565b60a081526000614d4d60a08301866149d8565b9050614d7060208301858051825260208082015190830152604090810151910152565b826080830152949350505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156104ad576104ad614d7e565b600082614dc457634e487b7160e01b600052601260045260246000fd5b500490565b808201808211156104ad576104ad614d7e565b6001600160401b03818116838216028082169190828114614dff57614dff614d7e565b505092915050565b60408101614e158285614766565b613ba56020830184614766565b83815260608101614e366020830185614766565b61253b6040830184614766565b6001600160a01b03868116825285166020820152604081018490526001600160401b038316606082015260a060808201819052600090614e85908301846149d8565b979650505050505050565b80820281158282048414176104ad576104ad614d7e565b82815260408101613ba56020830184614766565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351614ef3816017850160208801614252565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351614f24816028840160208801614252565b01602801949350505050565b602081526000613ba56020830184614788565b634e487b7160e01b600052603260045260246000fd5b600081614f6857614f68614d7e565b50600019019056fe34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996ea26469706673582212204f85f526e7cb7d566acaf93134151aa965949b1d71bafcf55d6bc326ad1f0b3b64736f6c63430008110033",
}

// MetaSchedulerABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaSchedulerMetaData.ABI instead.
var MetaSchedulerABI = MetaSchedulerMetaData.ABI

// MetaSchedulerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MetaSchedulerMetaData.Bin instead.
var MetaSchedulerBin = MetaSchedulerMetaData.Bin

// DeployMetaScheduler deploys a new Ethereum contract, binding an instance of MetaScheduler to it.
func DeployMetaScheduler(auth *bind.TransactOpts, backend bind.ContractBackend, _credit common.Address, _constants common.Address, _providerManager common.Address, _providerJobQueues common.Address, _jobs common.Address) (common.Address, *types.Transaction, *MetaScheduler, error) {
	parsed, err := MetaSchedulerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	toolsAddr, _, _, _ := DeployTools(auth, backend)
	MetaSchedulerBin = strings.ReplaceAll(MetaSchedulerBin, "__$5cf7eb674d9adda106ba0f4bb1fb6582ca$__", toolsAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetaSchedulerBin), backend, _credit, _constants, _providerManager, _providerJobQueues, _jobs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MetaScheduler{MetaSchedulerCaller: MetaSchedulerCaller{contract: contract}, MetaSchedulerTransactor: MetaSchedulerTransactor{contract: contract}, MetaSchedulerFilterer: MetaSchedulerFilterer{contract: contract}}, nil
}

// MetaScheduler is an auto generated Go binding around an Ethereum contract.
type MetaScheduler struct {
	MetaSchedulerCaller     // Read-only binding to the contract
	MetaSchedulerTransactor // Write-only binding to the contract
	MetaSchedulerFilterer   // Log filterer for contract events
}

// MetaSchedulerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaSchedulerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaSchedulerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaSchedulerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaSchedulerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaSchedulerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaSchedulerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaSchedulerSession struct {
	Contract     *MetaScheduler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaSchedulerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaSchedulerCallerSession struct {
	Contract *MetaSchedulerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MetaSchedulerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaSchedulerTransactorSession struct {
	Contract     *MetaSchedulerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MetaSchedulerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaSchedulerRaw struct {
	Contract *MetaScheduler // Generic contract binding to access the raw methods on
}

// MetaSchedulerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaSchedulerCallerRaw struct {
	Contract *MetaSchedulerCaller // Generic read-only contract binding to access the raw methods on
}

// MetaSchedulerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaSchedulerTransactorRaw struct {
	Contract *MetaSchedulerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaScheduler creates a new instance of MetaScheduler, bound to a specific deployed contract.
func NewMetaScheduler(address common.Address, backend bind.ContractBackend) (*MetaScheduler, error) {
	contract, err := bindMetaScheduler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaScheduler{MetaSchedulerCaller: MetaSchedulerCaller{contract: contract}, MetaSchedulerTransactor: MetaSchedulerTransactor{contract: contract}, MetaSchedulerFilterer: MetaSchedulerFilterer{contract: contract}}, nil
}

// NewMetaSchedulerCaller creates a new read-only instance of MetaScheduler, bound to a specific deployed contract.
func NewMetaSchedulerCaller(address common.Address, caller bind.ContractCaller) (*MetaSchedulerCaller, error) {
	contract, err := bindMetaScheduler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerCaller{contract: contract}, nil
}

// NewMetaSchedulerTransactor creates a new write-only instance of MetaScheduler, bound to a specific deployed contract.
func NewMetaSchedulerTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaSchedulerTransactor, error) {
	contract, err := bindMetaScheduler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerTransactor{contract: contract}, nil
}

// NewMetaSchedulerFilterer creates a new log filterer instance of MetaScheduler, bound to a specific deployed contract.
func NewMetaSchedulerFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaSchedulerFilterer, error) {
	contract, err := bindMetaScheduler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerFilterer{contract: contract}, nil
}

// bindMetaScheduler binds a generic wrapper to an already deployed contract.
func bindMetaScheduler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetaSchedulerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaScheduler *MetaSchedulerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaScheduler.Contract.MetaSchedulerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaScheduler *MetaSchedulerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.Contract.MetaSchedulerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaScheduler *MetaSchedulerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaScheduler.Contract.MetaSchedulerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaScheduler *MetaSchedulerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaScheduler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaScheduler *MetaSchedulerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaScheduler *MetaSchedulerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaScheduler.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MetaScheduler.Contract.DEFAULTADMINROLE(&_MetaScheduler.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MetaScheduler.Contract.DEFAULTADMINROLE(&_MetaScheduler.CallOpts)
}

// METASCHEDULERROLE is a free data retrieval call binding the contract method 0xe052888c.
//
// Solidity: function METASCHEDULER_ROLE() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCaller) METASCHEDULERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "METASCHEDULER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// METASCHEDULERROLE is a free data retrieval call binding the contract method 0xe052888c.
//
// Solidity: function METASCHEDULER_ROLE() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) METASCHEDULERROLE() ([32]byte, error) {
	return _MetaScheduler.Contract.METASCHEDULERROLE(&_MetaScheduler.CallOpts)
}

// METASCHEDULERROLE is a free data retrieval call binding the contract method 0xe052888c.
//
// Solidity: function METASCHEDULER_ROLE() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCallerSession) METASCHEDULERROLE() ([32]byte, error) {
	return _MetaScheduler.Contract.METASCHEDULERROLE(&_MetaScheduler.CallOpts)
}

// Constants is a free data retrieval call binding the contract method 0x72de5b2f.
//
// Solidity: function constants() view returns(address)
func (_MetaScheduler *MetaSchedulerCaller) Constants(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "constants")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Constants is a free data retrieval call binding the contract method 0x72de5b2f.
//
// Solidity: function constants() view returns(address)
func (_MetaScheduler *MetaSchedulerSession) Constants() (common.Address, error) {
	return _MetaScheduler.Contract.Constants(&_MetaScheduler.CallOpts)
}

// Constants is a free data retrieval call binding the contract method 0x72de5b2f.
//
// Solidity: function constants() view returns(address)
func (_MetaScheduler *MetaSchedulerCallerSession) Constants() (common.Address, error) {
	return _MetaScheduler.Contract.Constants(&_MetaScheduler.CallOpts)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_MetaScheduler *MetaSchedulerCaller) Credit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "credit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_MetaScheduler *MetaSchedulerSession) Credit() (common.Address, error) {
	return _MetaScheduler.Contract.Credit(&_MetaScheduler.CallOpts)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_MetaScheduler *MetaSchedulerCallerSession) Credit() (common.Address, error) {
	return _MetaScheduler.Contract.Credit(&_MetaScheduler.CallOpts)
}

// EnableRequestNewJob is a free data retrieval call binding the contract method 0x4c6dad12.
//
// Solidity: function enableRequestNewJob() view returns(bool)
func (_MetaScheduler *MetaSchedulerCaller) EnableRequestNewJob(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "enableRequestNewJob")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnableRequestNewJob is a free data retrieval call binding the contract method 0x4c6dad12.
//
// Solidity: function enableRequestNewJob() view returns(bool)
func (_MetaScheduler *MetaSchedulerSession) EnableRequestNewJob() (bool, error) {
	return _MetaScheduler.Contract.EnableRequestNewJob(&_MetaScheduler.CallOpts)
}

// EnableRequestNewJob is a free data retrieval call binding the contract method 0x4c6dad12.
//
// Solidity: function enableRequestNewJob() view returns(bool)
func (_MetaScheduler *MetaSchedulerCallerSession) EnableRequestNewJob() (bool, error) {
	return _MetaScheduler.Contract.EnableRequestNewJob(&_MetaScheduler.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MetaScheduler.Contract.GetRoleAdmin(&_MetaScheduler.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MetaScheduler.Contract.GetRoleAdmin(&_MetaScheduler.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MetaScheduler *MetaSchedulerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MetaScheduler *MetaSchedulerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasRole(&_MetaScheduler.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MetaScheduler *MetaSchedulerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasRole(&_MetaScheduler.CallOpts, role, account)
}

// Jobs is a free data retrieval call binding the contract method 0x7c8fce23.
//
// Solidity: function jobs() view returns(address)
func (_MetaScheduler *MetaSchedulerCaller) Jobs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "jobs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Jobs is a free data retrieval call binding the contract method 0x7c8fce23.
//
// Solidity: function jobs() view returns(address)
func (_MetaScheduler *MetaSchedulerSession) Jobs() (common.Address, error) {
	return _MetaScheduler.Contract.Jobs(&_MetaScheduler.CallOpts)
}

// Jobs is a free data retrieval call binding the contract method 0x7c8fce23.
//
// Solidity: function jobs() view returns(address)
func (_MetaScheduler *MetaSchedulerCallerSession) Jobs() (common.Address, error) {
	return _MetaScheduler.Contract.Jobs(&_MetaScheduler.CallOpts)
}

// ProviderJobQueues is a free data retrieval call binding the contract method 0xe2eaf3e7.
//
// Solidity: function providerJobQueues() view returns(address)
func (_MetaScheduler *MetaSchedulerCaller) ProviderJobQueues(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "providerJobQueues")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProviderJobQueues is a free data retrieval call binding the contract method 0xe2eaf3e7.
//
// Solidity: function providerJobQueues() view returns(address)
func (_MetaScheduler *MetaSchedulerSession) ProviderJobQueues() (common.Address, error) {
	return _MetaScheduler.Contract.ProviderJobQueues(&_MetaScheduler.CallOpts)
}

// ProviderJobQueues is a free data retrieval call binding the contract method 0xe2eaf3e7.
//
// Solidity: function providerJobQueues() view returns(address)
func (_MetaScheduler *MetaSchedulerCallerSession) ProviderJobQueues() (common.Address, error) {
	return _MetaScheduler.Contract.ProviderJobQueues(&_MetaScheduler.CallOpts)
}

// ProviderManager is a free data retrieval call binding the contract method 0x13151ec9.
//
// Solidity: function providerManager() view returns(address)
func (_MetaScheduler *MetaSchedulerCaller) ProviderManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "providerManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProviderManager is a free data retrieval call binding the contract method 0x13151ec9.
//
// Solidity: function providerManager() view returns(address)
func (_MetaScheduler *MetaSchedulerSession) ProviderManager() (common.Address, error) {
	return _MetaScheduler.Contract.ProviderManager(&_MetaScheduler.CallOpts)
}

// ProviderManager is a free data retrieval call binding the contract method 0x13151ec9.
//
// Solidity: function providerManager() view returns(address)
func (_MetaScheduler *MetaSchedulerCallerSession) ProviderManager() (common.Address, error) {
	return _MetaScheduler.Contract.ProviderManager(&_MetaScheduler.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaScheduler *MetaSchedulerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaScheduler *MetaSchedulerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaScheduler.Contract.SupportsInterface(&_MetaScheduler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MetaScheduler *MetaSchedulerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MetaScheduler.Contract.SupportsInterface(&_MetaScheduler.CallOpts, interfaceId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x5fae1450.
//
// Solidity: function cancelJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactor) CancelJob(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "cancelJob", _jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x5fae1450.
//
// Solidity: function cancelJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerSession) CancelJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.CancelJob(&_MetaScheduler.TransactOpts, _jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x5fae1450.
//
// Solidity: function cancelJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) CancelJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.CancelJob(&_MetaScheduler.TransactOpts, _jobId)
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0x5e1b2d65.
//
// Solidity: function claimNextCancellingJob() returns()
func (_MetaScheduler *MetaSchedulerTransactor) ClaimNextCancellingJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "claimNextCancellingJob")
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0x5e1b2d65.
//
// Solidity: function claimNextCancellingJob() returns()
func (_MetaScheduler *MetaSchedulerSession) ClaimNextCancellingJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextCancellingJob(&_MetaScheduler.TransactOpts)
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0x5e1b2d65.
//
// Solidity: function claimNextCancellingJob() returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ClaimNextCancellingJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextCancellingJob(&_MetaScheduler.TransactOpts)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns()
func (_MetaScheduler *MetaSchedulerTransactor) ClaimNextJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "claimNextJob")
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns()
func (_MetaScheduler *MetaSchedulerSession) ClaimNextJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextJob(&_MetaScheduler.TransactOpts)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ClaimNextJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextJob(&_MetaScheduler.TransactOpts)
}

// ClaimNextTopUpJob is a paid mutator transaction binding the contract method 0xebd4bf00.
//
// Solidity: function claimNextTopUpJob() returns()
func (_MetaScheduler *MetaSchedulerTransactor) ClaimNextTopUpJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "claimNextTopUpJob")
}

// ClaimNextTopUpJob is a paid mutator transaction binding the contract method 0xebd4bf00.
//
// Solidity: function claimNextTopUpJob() returns()
func (_MetaScheduler *MetaSchedulerSession) ClaimNextTopUpJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextTopUpJob(&_MetaScheduler.TransactOpts)
}

// ClaimNextTopUpJob is a paid mutator transaction binding the contract method 0xebd4bf00.
//
// Solidity: function claimNextTopUpJob() returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ClaimNextTopUpJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextTopUpJob(&_MetaScheduler.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.GrantRole(&_MetaScheduler.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.GrantRole(&_MetaScheduler.TransactOpts, role, account)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerTransactor) MetaSchedule(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "metaSchedule", _jobId, _providerAddr)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerSession) MetaSchedule(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.MetaSchedule(&_MetaScheduler.TransactOpts, _jobId, _providerAddr)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) MetaSchedule(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.MetaSchedule(&_MetaScheduler.TransactOpts, _jobId, _providerAddr)
}

// PanicJob is a paid mutator transaction binding the contract method 0x54b4a0d2.
//
// Solidity: function panicJob(bytes32 _jobId, string _lastError) returns()
func (_MetaScheduler *MetaSchedulerTransactor) PanicJob(opts *bind.TransactOpts, _jobId [32]byte, _lastError string) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "panicJob", _jobId, _lastError)
}

// PanicJob is a paid mutator transaction binding the contract method 0x54b4a0d2.
//
// Solidity: function panicJob(bytes32 _jobId, string _lastError) returns()
func (_MetaScheduler *MetaSchedulerSession) PanicJob(_jobId [32]byte, _lastError string) (*types.Transaction, error) {
	return _MetaScheduler.Contract.PanicJob(&_MetaScheduler.TransactOpts, _jobId, _lastError)
}

// PanicJob is a paid mutator transaction binding the contract method 0x54b4a0d2.
//
// Solidity: function panicJob(bytes32 _jobId, string _lastError) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) PanicJob(_jobId [32]byte, _lastError string) (*types.Transaction, error) {
	return _MetaScheduler.Contract.PanicJob(&_MetaScheduler.TransactOpts, _jobId, _lastError)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x0e4bf0e8.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _nextJobStatus, uint64 _jobDurationMinute, string _lastError) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderSetJobStatus(opts *bind.TransactOpts, _jobId [32]byte, _nextJobStatus uint8, _jobDurationMinute uint64, _lastError string) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerSetJobStatus", _jobId, _nextJobStatus, _jobDurationMinute, _lastError)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x0e4bf0e8.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _nextJobStatus, uint64 _jobDurationMinute, string _lastError) returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderSetJobStatus(_jobId [32]byte, _nextJobStatus uint8, _jobDurationMinute uint64, _lastError string) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderSetJobStatus(&_MetaScheduler.TransactOpts, _jobId, _nextJobStatus, _jobDurationMinute, _lastError)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x0e4bf0e8.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _nextJobStatus, uint64 _jobDurationMinute, string _lastError) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderSetJobStatus(_jobId [32]byte, _nextJobStatus uint8, _jobDurationMinute uint64, _lastError string) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderSetJobStatus(&_MetaScheduler.TransactOpts, _jobId, _nextJobStatus, _jobDurationMinute, _lastError)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x1f92a63f.
//
// Solidity: function refuseJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactor) RefuseJob(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "refuseJob", _jobId)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x1f92a63f.
//
// Solidity: function refuseJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerSession) RefuseJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RefuseJob(&_MetaScheduler.TransactOpts, _jobId)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x1f92a63f.
//
// Solidity: function refuseJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) RefuseJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RefuseJob(&_MetaScheduler.TransactOpts, _jobId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RenounceRole(&_MetaScheduler.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RenounceRole(&_MetaScheduler.TransactOpts, role, account)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0xd9943788.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _definition, uint256 _lockedCredits, bytes32 _jobName, bool _delegateSpendingAuthority) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactor) RequestNewJob(opts *bind.TransactOpts, _definition JobDefinition, _lockedCredits *big.Int, _jobName [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "requestNewJob", _definition, _lockedCredits, _jobName, _delegateSpendingAuthority)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0xd9943788.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _definition, uint256 _lockedCredits, bytes32 _jobName, bool _delegateSpendingAuthority) returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) RequestNewJob(_definition JobDefinition, _lockedCredits *big.Int, _jobName [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _lockedCredits, _jobName, _delegateSpendingAuthority)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0xd9943788.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _definition, uint256 _lockedCredits, bytes32 _jobName, bool _delegateSpendingAuthority) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactorSession) RequestNewJob(_definition JobDefinition, _lockedCredits *big.Int, _jobName [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _lockedCredits, _jobName, _delegateSpendingAuthority)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RevokeRole(&_MetaScheduler.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RevokeRole(&_MetaScheduler.TransactOpts, role, account)
}

// SetDelegateSpendingAuthority is a paid mutator transaction binding the contract method 0x329af326.
//
// Solidity: function setDelegateSpendingAuthority(bytes32 _jobId, bool _delegateSpendingAuthority) returns()
func (_MetaScheduler *MetaSchedulerTransactor) SetDelegateSpendingAuthority(opts *bind.TransactOpts, _jobId [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "setDelegateSpendingAuthority", _jobId, _delegateSpendingAuthority)
}

// SetDelegateSpendingAuthority is a paid mutator transaction binding the contract method 0x329af326.
//
// Solidity: function setDelegateSpendingAuthority(bytes32 _jobId, bool _delegateSpendingAuthority) returns()
func (_MetaScheduler *MetaSchedulerSession) SetDelegateSpendingAuthority(_jobId [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.SetDelegateSpendingAuthority(&_MetaScheduler.TransactOpts, _jobId, _delegateSpendingAuthority)
}

// SetDelegateSpendingAuthority is a paid mutator transaction binding the contract method 0x329af326.
//
// Solidity: function setDelegateSpendingAuthority(bytes32 _jobId, bool _delegateSpendingAuthority) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) SetDelegateSpendingAuthority(_jobId [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.SetDelegateSpendingAuthority(&_MetaScheduler.TransactOpts, _jobId, _delegateSpendingAuthority)
}

// SetEnableRequestNewJob is a paid mutator transaction binding the contract method 0x0bba4dc5.
//
// Solidity: function setEnableRequestNewJob(bool _enable) returns()
func (_MetaScheduler *MetaSchedulerTransactor) SetEnableRequestNewJob(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "setEnableRequestNewJob", _enable)
}

// SetEnableRequestNewJob is a paid mutator transaction binding the contract method 0x0bba4dc5.
//
// Solidity: function setEnableRequestNewJob(bool _enable) returns()
func (_MetaScheduler *MetaSchedulerSession) SetEnableRequestNewJob(_enable bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.SetEnableRequestNewJob(&_MetaScheduler.TransactOpts, _enable)
}

// SetEnableRequestNewJob is a paid mutator transaction binding the contract method 0x0bba4dc5.
//
// Solidity: function setEnableRequestNewJob(bool _enable) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) SetEnableRequestNewJob(_enable bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.SetEnableRequestNewJob(&_MetaScheduler.TransactOpts, _enable)
}

// TimeoutJob is a paid mutator transaction binding the contract method 0x18263c59.
//
// Solidity: function timeoutJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactor) TimeoutJob(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "timeoutJob", _jobId)
}

// TimeoutJob is a paid mutator transaction binding the contract method 0x18263c59.
//
// Solidity: function timeoutJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerSession) TimeoutJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TimeoutJob(&_MetaScheduler.TransactOpts, _jobId)
}

// TimeoutJob is a paid mutator transaction binding the contract method 0x18263c59.
//
// Solidity: function timeoutJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) TimeoutJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TimeoutJob(&_MetaScheduler.TransactOpts, _jobId)
}

// TopUpJob is a paid mutator transaction binding the contract method 0x2fecc4f6.
//
// Solidity: function topUpJob(bytes32 _jobId, uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactor) TopUpJob(opts *bind.TransactOpts, _jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "topUpJob", _jobId, _amount)
}

// TopUpJob is a paid mutator transaction binding the contract method 0x2fecc4f6.
//
// Solidity: function topUpJob(bytes32 _jobId, uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerSession) TopUpJob(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpJob(&_MetaScheduler.TransactOpts, _jobId, _amount)
}

// TopUpJob is a paid mutator transaction binding the contract method 0x2fecc4f6.
//
// Solidity: function topUpJob(bytes32 _jobId, uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) TopUpJob(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpJob(&_MetaScheduler.TransactOpts, _jobId, _amount)
}

// TopUpJobDelegate is a paid mutator transaction binding the contract method 0x69ee1bf9.
//
// Solidity: function topUpJobDelegate(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactor) TopUpJobDelegate(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "topUpJobDelegate", _jobId)
}

// TopUpJobDelegate is a paid mutator transaction binding the contract method 0x69ee1bf9.
//
// Solidity: function topUpJobDelegate(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerSession) TopUpJobDelegate(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpJobDelegate(&_MetaScheduler.TransactOpts, _jobId)
}

// TopUpJobDelegate is a paid mutator transaction binding the contract method 0x69ee1bf9.
//
// Solidity: function topUpJobDelegate(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) TopUpJobDelegate(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpJobDelegate(&_MetaScheduler.TransactOpts, _jobId)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0xd77836ce.
//
// Solidity: function withdrawAdmin(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactor) WithdrawAdmin(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "withdrawAdmin", _amount)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0xd77836ce.
//
// Solidity: function withdrawAdmin(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerSession) WithdrawAdmin(_amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.WithdrawAdmin(&_MetaScheduler.TransactOpts, _amount)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0xd77836ce.
//
// Solidity: function withdrawAdmin(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) WithdrawAdmin(_amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.WithdrawAdmin(&_MetaScheduler.TransactOpts, _amount)
}

// MetaSchedulerBilledTooMuchEventIterator is returned from FilterBilledTooMuchEvent and is used to iterate over the raw logs and unpacked data for BilledTooMuchEvent events raised by the MetaScheduler contract.
type MetaSchedulerBilledTooMuchEventIterator struct {
	Event *MetaSchedulerBilledTooMuchEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerBilledTooMuchEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerBilledTooMuchEvent)
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
		it.Event = new(MetaSchedulerBilledTooMuchEvent)
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
func (it *MetaSchedulerBilledTooMuchEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerBilledTooMuchEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerBilledTooMuchEvent represents a BilledTooMuchEvent event raised by the MetaScheduler contract.
type MetaSchedulerBilledTooMuchEvent struct {
	JobId         [32]byte
	ProviderAddr  common.Address
	BillingAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBilledTooMuchEvent is a free log retrieval operation binding the contract event 0x17e65314b087df225f56701d0a66a3f7d9ce0f26077307b4b765a19c60a36d44.
//
// Solidity: event BilledTooMuchEvent(bytes32 _jobId, address _providerAddr, uint256 _billingAmount)
func (_MetaScheduler *MetaSchedulerFilterer) FilterBilledTooMuchEvent(opts *bind.FilterOpts) (*MetaSchedulerBilledTooMuchEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "BilledTooMuchEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerBilledTooMuchEventIterator{contract: _MetaScheduler.contract, event: "BilledTooMuchEvent", logs: logs, sub: sub}, nil
}

// WatchBilledTooMuchEvent is a free log subscription operation binding the contract event 0x17e65314b087df225f56701d0a66a3f7d9ce0f26077307b4b765a19c60a36d44.
//
// Solidity: event BilledTooMuchEvent(bytes32 _jobId, address _providerAddr, uint256 _billingAmount)
func (_MetaScheduler *MetaSchedulerFilterer) WatchBilledTooMuchEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerBilledTooMuchEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "BilledTooMuchEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerBilledTooMuchEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "BilledTooMuchEvent", log); err != nil {
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

// ParseBilledTooMuchEvent is a log parse operation binding the contract event 0x17e65314b087df225f56701d0a66a3f7d9ce0f26077307b4b765a19c60a36d44.
//
// Solidity: event BilledTooMuchEvent(bytes32 _jobId, address _providerAddr, uint256 _billingAmount)
func (_MetaScheduler *MetaSchedulerFilterer) ParseBilledTooMuchEvent(log types.Log) (*MetaSchedulerBilledTooMuchEvent, error) {
	event := new(MetaSchedulerBilledTooMuchEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "BilledTooMuchEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerClaimJobEventIterator is returned from FilterClaimJobEvent and is used to iterate over the raw logs and unpacked data for ClaimJobEvent events raised by the MetaScheduler contract.
type MetaSchedulerClaimJobEventIterator struct {
	Event *MetaSchedulerClaimJobEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerClaimJobEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerClaimJobEvent)
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
		it.Event = new(MetaSchedulerClaimJobEvent)
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
func (it *MetaSchedulerClaimJobEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerClaimJobEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerClaimJobEvent represents a ClaimJobEvent event raised by the MetaScheduler contract.
type MetaSchedulerClaimJobEvent struct {
	CustomerAddr      common.Address
	ProviderAddr      common.Address
	JobId             [32]byte
	MaxDurationMinute uint64
	JobDefinition     JobDefinition
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimJobEvent is a free log retrieval operation binding the contract event 0xc8d7c7c2914e1aa1462fe7999f8a18a0f1043d7d94ab692c3ac9af846f1be8c7.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) FilterClaimJobEvent(opts *bind.FilterOpts) (*MetaSchedulerClaimJobEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "ClaimJobEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerClaimJobEventIterator{contract: _MetaScheduler.contract, event: "ClaimJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimJobEvent is a free log subscription operation binding the contract event 0xc8d7c7c2914e1aa1462fe7999f8a18a0f1043d7d94ab692c3ac9af846f1be8c7.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) WatchClaimJobEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerClaimJobEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "ClaimJobEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerClaimJobEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "ClaimJobEvent", log); err != nil {
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

// ParseClaimJobEvent is a log parse operation binding the contract event 0xc8d7c7c2914e1aa1462fe7999f8a18a0f1043d7d94ab692c3ac9af846f1be8c7.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) ParseClaimJobEvent(log types.Log) (*MetaSchedulerClaimJobEvent, error) {
	event := new(MetaSchedulerClaimJobEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "ClaimJobEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerClaimNextCancellingJobEventIterator is returned from FilterClaimNextCancellingJobEvent and is used to iterate over the raw logs and unpacked data for ClaimNextCancellingJobEvent events raised by the MetaScheduler contract.
type MetaSchedulerClaimNextCancellingJobEventIterator struct {
	Event *MetaSchedulerClaimNextCancellingJobEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerClaimNextCancellingJobEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerClaimNextCancellingJobEvent)
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
		it.Event = new(MetaSchedulerClaimNextCancellingJobEvent)
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
func (it *MetaSchedulerClaimNextCancellingJobEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerClaimNextCancellingJobEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerClaimNextCancellingJobEvent represents a ClaimNextCancellingJobEvent event raised by the MetaScheduler contract.
type MetaSchedulerClaimNextCancellingJobEvent struct {
	CustomerAddr common.Address
	ProviderAddr common.Address
	JobId        [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimNextCancellingJobEvent is a free log retrieval operation binding the contract event 0x290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d.
//
// Solidity: event ClaimNextCancellingJobEvent(address customerAddr, address providerAddr, bytes32 jobId)
func (_MetaScheduler *MetaSchedulerFilterer) FilterClaimNextCancellingJobEvent(opts *bind.FilterOpts) (*MetaSchedulerClaimNextCancellingJobEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "ClaimNextCancellingJobEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerClaimNextCancellingJobEventIterator{contract: _MetaScheduler.contract, event: "ClaimNextCancellingJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimNextCancellingJobEvent is a free log subscription operation binding the contract event 0x290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d.
//
// Solidity: event ClaimNextCancellingJobEvent(address customerAddr, address providerAddr, bytes32 jobId)
func (_MetaScheduler *MetaSchedulerFilterer) WatchClaimNextCancellingJobEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerClaimNextCancellingJobEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "ClaimNextCancellingJobEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerClaimNextCancellingJobEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "ClaimNextCancellingJobEvent", log); err != nil {
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

// ParseClaimNextCancellingJobEvent is a log parse operation binding the contract event 0x290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d.
//
// Solidity: event ClaimNextCancellingJobEvent(address customerAddr, address providerAddr, bytes32 jobId)
func (_MetaScheduler *MetaSchedulerFilterer) ParseClaimNextCancellingJobEvent(log types.Log) (*MetaSchedulerClaimNextCancellingJobEvent, error) {
	event := new(MetaSchedulerClaimNextCancellingJobEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "ClaimNextCancellingJobEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerClaimNextTopUpJobEventIterator is returned from FilterClaimNextTopUpJobEvent and is used to iterate over the raw logs and unpacked data for ClaimNextTopUpJobEvent events raised by the MetaScheduler contract.
type MetaSchedulerClaimNextTopUpJobEventIterator struct {
	Event *MetaSchedulerClaimNextTopUpJobEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerClaimNextTopUpJobEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerClaimNextTopUpJobEvent)
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
		it.Event = new(MetaSchedulerClaimNextTopUpJobEvent)
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
func (it *MetaSchedulerClaimNextTopUpJobEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerClaimNextTopUpJobEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerClaimNextTopUpJobEvent represents a ClaimNextTopUpJobEvent event raised by the MetaScheduler contract.
type MetaSchedulerClaimNextTopUpJobEvent struct {
	JobId             [32]byte
	ProviderAddr      common.Address
	MaxDurationMinute uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimNextTopUpJobEvent is a free log retrieval operation binding the contract event 0xa42f2b4a7ee7f91857a4c98fc71fc48546a284d5db48dd77b7ab81030a494470.
//
// Solidity: event ClaimNextTopUpJobEvent(bytes32 _jobId, address _providerAddr, uint64 maxDurationMinute)
func (_MetaScheduler *MetaSchedulerFilterer) FilterClaimNextTopUpJobEvent(opts *bind.FilterOpts) (*MetaSchedulerClaimNextTopUpJobEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "ClaimNextTopUpJobEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerClaimNextTopUpJobEventIterator{contract: _MetaScheduler.contract, event: "ClaimNextTopUpJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimNextTopUpJobEvent is a free log subscription operation binding the contract event 0xa42f2b4a7ee7f91857a4c98fc71fc48546a284d5db48dd77b7ab81030a494470.
//
// Solidity: event ClaimNextTopUpJobEvent(bytes32 _jobId, address _providerAddr, uint64 maxDurationMinute)
func (_MetaScheduler *MetaSchedulerFilterer) WatchClaimNextTopUpJobEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerClaimNextTopUpJobEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "ClaimNextTopUpJobEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerClaimNextTopUpJobEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "ClaimNextTopUpJobEvent", log); err != nil {
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

// ParseClaimNextTopUpJobEvent is a log parse operation binding the contract event 0xa42f2b4a7ee7f91857a4c98fc71fc48546a284d5db48dd77b7ab81030a494470.
//
// Solidity: event ClaimNextTopUpJobEvent(bytes32 _jobId, address _providerAddr, uint64 maxDurationMinute)
func (_MetaScheduler *MetaSchedulerFilterer) ParseClaimNextTopUpJobEvent(log types.Log) (*MetaSchedulerClaimNextTopUpJobEvent, error) {
	event := new(MetaSchedulerClaimNextTopUpJobEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "ClaimNextTopUpJobEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerJobRefusedEventIterator is returned from FilterJobRefusedEvent and is used to iterate over the raw logs and unpacked data for JobRefusedEvent events raised by the MetaScheduler contract.
type MetaSchedulerJobRefusedEventIterator struct {
	Event *MetaSchedulerJobRefusedEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerJobRefusedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerJobRefusedEvent)
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
		it.Event = new(MetaSchedulerJobRefusedEvent)
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
func (it *MetaSchedulerJobRefusedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerJobRefusedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerJobRefusedEvent represents a JobRefusedEvent event raised by the MetaScheduler contract.
type MetaSchedulerJobRefusedEvent struct {
	JobId        [32]byte
	ProviderAddr common.Address
	CustomerAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterJobRefusedEvent is a free log retrieval operation binding the contract event 0x50d9c3fab9ef0192905beb84254b4ffb6fe086795cc23de484ec65947b6615a2.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr, address _customerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) FilterJobRefusedEvent(opts *bind.FilterOpts) (*MetaSchedulerJobRefusedEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "JobRefusedEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerJobRefusedEventIterator{contract: _MetaScheduler.contract, event: "JobRefusedEvent", logs: logs, sub: sub}, nil
}

// WatchJobRefusedEvent is a free log subscription operation binding the contract event 0x50d9c3fab9ef0192905beb84254b4ffb6fe086795cc23de484ec65947b6615a2.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr, address _customerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) WatchJobRefusedEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerJobRefusedEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "JobRefusedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerJobRefusedEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "JobRefusedEvent", log); err != nil {
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

// ParseJobRefusedEvent is a log parse operation binding the contract event 0x50d9c3fab9ef0192905beb84254b4ffb6fe086795cc23de484ec65947b6615a2.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr, address _customerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) ParseJobRefusedEvent(log types.Log) (*MetaSchedulerJobRefusedEvent, error) {
	event := new(MetaSchedulerJobRefusedEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "JobRefusedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerJobTransitionEventIterator is returned from FilterJobTransitionEvent and is used to iterate over the raw logs and unpacked data for JobTransitionEvent events raised by the MetaScheduler contract.
type MetaSchedulerJobTransitionEventIterator struct {
	Event *MetaSchedulerJobTransitionEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerJobTransitionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerJobTransitionEvent)
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
		it.Event = new(MetaSchedulerJobTransitionEvent)
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
func (it *MetaSchedulerJobTransitionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerJobTransitionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerJobTransitionEvent represents a JobTransitionEvent event raised by the MetaScheduler contract.
type MetaSchedulerJobTransitionEvent struct {
	JobId [32]byte
	From  uint8
	To    uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJobTransitionEvent is a free log retrieval operation binding the contract event 0x0bba917f0a1e0fc0d51a75273e7088a4dfecb010699e60ac9c58526429f6c37f.
//
// Solidity: event JobTransitionEvent(bytes32 _jobId, uint8 _from, uint8 _to)
func (_MetaScheduler *MetaSchedulerFilterer) FilterJobTransitionEvent(opts *bind.FilterOpts) (*MetaSchedulerJobTransitionEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "JobTransitionEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerJobTransitionEventIterator{contract: _MetaScheduler.contract, event: "JobTransitionEvent", logs: logs, sub: sub}, nil
}

// WatchJobTransitionEvent is a free log subscription operation binding the contract event 0x0bba917f0a1e0fc0d51a75273e7088a4dfecb010699e60ac9c58526429f6c37f.
//
// Solidity: event JobTransitionEvent(bytes32 _jobId, uint8 _from, uint8 _to)
func (_MetaScheduler *MetaSchedulerFilterer) WatchJobTransitionEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerJobTransitionEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "JobTransitionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerJobTransitionEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "JobTransitionEvent", log); err != nil {
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

// ParseJobTransitionEvent is a log parse operation binding the contract event 0x0bba917f0a1e0fc0d51a75273e7088a4dfecb010699e60ac9c58526429f6c37f.
//
// Solidity: event JobTransitionEvent(bytes32 _jobId, uint8 _from, uint8 _to)
func (_MetaScheduler *MetaSchedulerFilterer) ParseJobTransitionEvent(log types.Log) (*MetaSchedulerJobTransitionEvent, error) {
	event := new(MetaSchedulerJobTransitionEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "JobTransitionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerNewJobRequestEventIterator is returned from FilterNewJobRequestEvent and is used to iterate over the raw logs and unpacked data for NewJobRequestEvent events raised by the MetaScheduler contract.
type MetaSchedulerNewJobRequestEventIterator struct {
	Event *MetaSchedulerNewJobRequestEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerNewJobRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerNewJobRequestEvent)
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
		it.Event = new(MetaSchedulerNewJobRequestEvent)
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
func (it *MetaSchedulerNewJobRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerNewJobRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerNewJobRequestEvent represents a NewJobRequestEvent event raised by the MetaScheduler contract.
type MetaSchedulerNewJobRequestEvent struct {
	JobId        [32]byte
	CustomerAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewJobRequestEvent is a free log retrieval operation binding the contract event 0x1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3.
//
// Solidity: event NewJobRequestEvent(bytes32 _jobId, address _customerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) FilterNewJobRequestEvent(opts *bind.FilterOpts) (*MetaSchedulerNewJobRequestEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "NewJobRequestEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerNewJobRequestEventIterator{contract: _MetaScheduler.contract, event: "NewJobRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewJobRequestEvent is a free log subscription operation binding the contract event 0x1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3.
//
// Solidity: event NewJobRequestEvent(bytes32 _jobId, address _customerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) WatchNewJobRequestEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerNewJobRequestEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "NewJobRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerNewJobRequestEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "NewJobRequestEvent", log); err != nil {
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

// ParseNewJobRequestEvent is a log parse operation binding the contract event 0x1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3.
//
// Solidity: event NewJobRequestEvent(bytes32 _jobId, address _customerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) ParseNewJobRequestEvent(log types.Log) (*MetaSchedulerNewJobRequestEvent, error) {
	event := new(MetaSchedulerNewJobRequestEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "NewJobRequestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MetaScheduler contract.
type MetaSchedulerRoleAdminChangedIterator struct {
	Event *MetaSchedulerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerRoleAdminChanged)
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
		it.Event = new(MetaSchedulerRoleAdminChanged)
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
func (it *MetaSchedulerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerRoleAdminChanged represents a RoleAdminChanged event raised by the MetaScheduler contract.
type MetaSchedulerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MetaScheduler *MetaSchedulerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MetaSchedulerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerRoleAdminChangedIterator{contract: _MetaScheduler.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MetaScheduler *MetaSchedulerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MetaSchedulerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerRoleAdminChanged)
				if err := _MetaScheduler.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MetaScheduler *MetaSchedulerFilterer) ParseRoleAdminChanged(log types.Log) (*MetaSchedulerRoleAdminChanged, error) {
	event := new(MetaSchedulerRoleAdminChanged)
	if err := _MetaScheduler.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MetaScheduler contract.
type MetaSchedulerRoleGrantedIterator struct {
	Event *MetaSchedulerRoleGranted // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerRoleGranted)
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
		it.Event = new(MetaSchedulerRoleGranted)
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
func (it *MetaSchedulerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerRoleGranted represents a RoleGranted event raised by the MetaScheduler contract.
type MetaSchedulerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaScheduler *MetaSchedulerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MetaSchedulerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerRoleGrantedIterator{contract: _MetaScheduler.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaScheduler *MetaSchedulerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MetaSchedulerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerRoleGranted)
				if err := _MetaScheduler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaScheduler *MetaSchedulerFilterer) ParseRoleGranted(log types.Log) (*MetaSchedulerRoleGranted, error) {
	event := new(MetaSchedulerRoleGranted)
	if err := _MetaScheduler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MetaScheduler contract.
type MetaSchedulerRoleRevokedIterator struct {
	Event *MetaSchedulerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerRoleRevoked)
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
		it.Event = new(MetaSchedulerRoleRevoked)
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
func (it *MetaSchedulerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerRoleRevoked represents a RoleRevoked event raised by the MetaScheduler contract.
type MetaSchedulerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaScheduler *MetaSchedulerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MetaSchedulerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerRoleRevokedIterator{contract: _MetaScheduler.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaScheduler *MetaSchedulerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MetaSchedulerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerRoleRevoked)
				if err := _MetaScheduler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MetaScheduler *MetaSchedulerFilterer) ParseRoleRevoked(log types.Log) (*MetaSchedulerRoleRevoked, error) {
	event := new(MetaSchedulerRoleRevoked)
	if err := _MetaScheduler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

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
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardMetaData contains all meta data concerning the ReentrancyGuard contract.
var ReentrancyGuardMetaData = &bind.MetaData{
	ABI: "[]",
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardMetaData.ABI instead.
var ReentrancyGuardABI = ReentrancyGuardMetaData.ABI

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203eb101fc2ade6ef50025d3f0b05fe85efb589be6d1878cd5495a717906bf2de364736f6c63430008110033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c5d6c734e147510a4cba68515ad7f9014cd66d2f51dc998c071d660dc98eca8564736f6c63430008110033",
}

// SignedMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SignedMathMetaData.ABI instead.
var SignedMathABI = SignedMathMetaData.ABI

// SignedMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignedMathMetaData.Bin instead.
var SignedMathBin = SignedMathMetaData.Bin

// DeploySignedMath deploys a new Ethereum contract, binding an instance of SignedMath to it.
func DeploySignedMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignedMath, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignedMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// SignedMath is an auto generated Go binding around an Ethereum contract.
type SignedMath struct {
	SignedMathCaller     // Read-only binding to the contract
	SignedMathTransactor // Write-only binding to the contract
	SignedMathFilterer   // Log filterer for contract events
}

// SignedMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedMathSession struct {
	Contract     *SignedMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedMathCallerSession struct {
	Contract *SignedMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SignedMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedMathTransactorSession struct {
	Contract     *SignedMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SignedMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedMathRaw struct {
	Contract *SignedMath // Generic contract binding to access the raw methods on
}

// SignedMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedMathCallerRaw struct {
	Contract *SignedMathCaller // Generic read-only contract binding to access the raw methods on
}

// SignedMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedMathTransactorRaw struct {
	Contract *SignedMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedMath creates a new instance of SignedMath, bound to a specific deployed contract.
func NewSignedMath(address common.Address, backend bind.ContractBackend) (*SignedMath, error) {
	contract, err := bindSignedMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// NewSignedMathCaller creates a new read-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathCaller(address common.Address, caller bind.ContractCaller) (*SignedMathCaller, error) {
	contract, err := bindSignedMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathCaller{contract: contract}, nil
}

// NewSignedMathTransactor creates a new write-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedMathTransactor, error) {
	contract, err := bindSignedMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathTransactor{contract: contract}, nil
}

// NewSignedMathFilterer creates a new log filterer instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedMathFilterer, error) {
	contract, err := bindSignedMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedMathFilterer{contract: contract}, nil
}

// bindSignedMath binds a generic wrapper to an already deployed contract.
func bindSignedMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.SignedMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transact(opts, method, params...)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220554120e6795de1095fd54992294a43b7bbd9d054b9da4b5204aea17e03b6254e64736f6c63430008110033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// ToolsMetaData contains all meta data concerning the Tools contract.
var ToolsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"from\",\"type\":\"uint8\"},{\"internalType\":\"enumJobStatus\",\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"InvalidTransition\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"JobHotStatusOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameStatusError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"_currentJobStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_nextJobStatus\",\"type\":\"uint8\"}],\"name\":\"checkNewJobStatus\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"convertAddressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes\",\"type\":\"bytes32\"}],\"name\":\"convertBytes32ToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPricePerMin\",\"type\":\"uint256\"}],\"internalType\":\"structProviderPrices\",\"name\":\"providerPrices\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"}],\"name\":\"convertCreditToDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPricePerMin\",\"type\":\"uint256\"}],\"internalType\":\"structProviderPrices\",\"name\":\"providerPrices\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"durationMinute\",\"type\":\"uint64\"}],\"name\":\"convertDurationToCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingTopUp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateSpendingAuthority\",\"type\":\"bool\"}],\"internalType\":\"structJobCost\",\"name\":\"cost\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelRequestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"panicTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structJobTime\",\"name\":\"time\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"hasCancelRequest\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"lastError\",\"type\":\"string\"}],\"internalType\":\"structJob\",\"name\":\"job\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuPricePerMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memPricePerMin\",\"type\":\"uint256\"}],\"internalType\":\"structProviderPrices\",\"name\":\"providerPrices\",\"type\":\"tuple\"}],\"name\":\"getRemainingTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"isDelegateTopable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"isJobCold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpusPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel[]\",\"name\":\"uses\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structLabel\",\"name\":\"label\",\"type\":\"tuple\"},{\"internalType\":\"bytes2\",\"name\":\"op\",\"type\":\"bytes2\"}],\"internalType\":\"structAffinity[]\",\"name\":\"affinity\",\"type\":\"tuple[]\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_jobDefinition\",\"type\":\"tuple\"}],\"name\":\"isJobDefinitionValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"isJobHot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x61101361003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c8063c4b2bd2c11610070578063c4b2bd2c1461014e578063d55388b914610179578063da498b291461018c578063f4cc70ac1461019f578063f675b070146101b257600080fd5b80631efa2220146100ad57806334d515f5146100db5780635893740e146100fc57806397d2874f146101165780639e71f7a11461012b575b600080fd5b6100be6100bb3660046106d6565b90565b6040516001600160a01b0390911681526020015b60405180910390f35b6100ee6100e9366004610b7e565b6101c5565b6040519081526020016100d2565b6100ee61010a366004610bf3565b6001600160a01b031690565b610129610124366004610c1d565b61027a565b005b61013e610139366004610c50565b6104c9565b60405190151581526020016100d2565b61016161015c366004610c8c565b610513565b6040516001600160401b0390911681526020016100d2565b61013e610187366004610ce3565b6105c0565b61013e61019a366004610ce3565b6105d2565b6101616101ad366004610de9565b61065f565b61013e6101c0366004610ce3565b6106b2565b600083604001516001600160401b031683602001516101e49190610f1d565b84604001516001600160401b031685602001516001600160401b0316856040015161020f9190610f1d565b6102199190610f1d565b85518551610230916001600160401b031690610f1d565b61023a9190610f34565b6102449190610f34565b84606001516001600160401b0316836001600160401b03166102669190610f1d565b6102709190610f1d565b90505b9392505050565b80600881111561028c5761028c610f47565b82600881111561029e5761029e610f47565b036102bc57604051632a856fc960e01b815260040160405180910390fd5b6102c5826105c0565b6102ed5781604051634634126160e11b81526004016102e49190610f7f565b60405180910390fd5b600881600881111561030157610301610f47565b0361030a575050565b600082600881111561031e5761031e610f47565b0361037f5760015b81600881111561033857610338610f47565b14158015610359575060045b81600881111561035657610356610f47565b14155b1561037b5781816040516305fdf05f60e31b81526004016102e4929190610f8d565b5050565b600182600881111561039357610393610f47565b036103ca5760008160088111156103ac576103ac610f47565b141580156103bc57506002610326565b801561035957506004610344565b60028260088111156103de576103de610f47565b036104435760008160088111156103f7576103f7610f47565b141580156104175750600381600881111561041457610414610f47565b14155b80156104355750600481600881111561043257610432610f47565b14155b801561035957506006610344565b600382600881111561045757610457610f47565b0361037b57600581600881111561047057610470610f47565b141580156104905750600681600881111561048d5761048d610f47565b14155b80156104ae575060048160088111156104ab576104ab610f47565b14155b80156103595750600781600881111561035657610356610f47565b60008082606001516001600160401b03161180156104f45750600082604001516001600160401b0316115b801561050d5750600082602001516001600160401b0316115b92915050565b60008084604001516001600160401b031684602001516105339190610f1d565b85604001516001600160401b031686602001516001600160401b0316866040015161055e9190610f1d565b6105689190610f1d565b8651865161057f916001600160401b031690610f1d565b6105899190610f34565b6105939190610f34565b85606001516001600160401b03166105ab9190610f1d565b90506105b78184610fa8565b95945050505050565b60006105cb826105d2565b1592915050565b600060048260088111156105e8576105e8610f47565b14806106055750600582600881111561060357610603610f47565b145b806106215750600782600881111561061f5761061f610f47565b145b8061063d5750600682600881111561063b5761063b610f47565b145b8061050d575060085b82600881111561065857610658610f47565b1492915050565b6000806106798460800151848660a0015160000151610513565b90506000603c8560c0015160000151426106939190610fca565b61069d9190610fa8565b90506105b7816001600160401b038416610fca565b600060038260088111156106c8576106c8610f47565b148061050d57506002610646565b6000602082840312156106e857600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715610727576107276106ef565b60405290565b60405161010081016001600160401b0381118282101715610727576107276106ef565b60405161014081016001600160401b0381118282101715610727576107276106ef565b604051601f8201601f191681016001600160401b038111828210171561079b5761079b6106ef565b604052919050565b80356001600160401b03811681146107ba57600080fd5b919050565b600082601f8301126107d057600080fd5b81356001600160401b038111156107e9576107e96106ef565b6107fc601f8201601f1916602001610773565b81815284602083860101111561081157600080fd5b816020850160208301376000918101602001919091529392505050565b8035600581106107ba57600080fd5b60006001600160401b03821115610856576108566106ef565b5060051b60200190565b60006040828403121561087257600080fd5b61087a610705565b905081356001600160401b038082111561089357600080fd5b61089f858386016107bf565b835260208401359150808211156108b557600080fd5b506108c2848285016107bf565b60208301525092915050565b600082601f8301126108df57600080fd5b813560206108f46108ef8361083d565b610773565b82815260059290921b8401810191818101908684111561091357600080fd5b8286015b848110156109525780356001600160401b038111156109365760008081fd5b6109448986838b0101610860565b845250918301918301610917565b509695505050505050565b600082601f83011261096e57600080fd5b8135602061097e6108ef8361083d565b82815260059290921b8401810191818101908684111561099d57600080fd5b8286015b848110156109525780356001600160401b03808211156109c15760008081fd5b908801906040828b03601f19018113156109db5760008081fd5b6109e3610705565b87840135838111156109f55760008081fd5b610a038d8a83880101610860565b82525092810135926001600160f01b031984168414610a2457600092508283fd5b80880193909352505083529183019183016109a1565b60006101008284031215610a4d57600080fd5b610a5561072d565b9050610a60826107a3565b8152610a6e602083016107a3565b6020820152610a7f604083016107a3565b6040820152610a90606083016107a3565b606082015260808201356001600160401b0380821115610aaf57600080fd5b610abb858386016107bf565b6080840152610acc60a0850161082e565b60a084015260c0840135915080821115610ae557600080fd5b610af1858386016108ce565b60c084015260e0840135915080821115610b0a57600080fd5b50610b178482850161095d565b60e08301525092915050565b600060608284031215610b3557600080fd5b604051606081018181106001600160401b0382111715610b5757610b576106ef565b80604052508091508235815260208301356020820152604083013560408201525092915050565b600080600060a08486031215610b9357600080fd5b83356001600160401b03811115610ba957600080fd5b610bb586828701610a3a565b935050610bc58560208601610b23565b9150610bd3608085016107a3565b90509250925092565b80356001600160a01b03811681146107ba57600080fd5b600060208284031215610c0557600080fd5b61027382610bdc565b8035600981106107ba57600080fd5b60008060408385031215610c3057600080fd5b610c3983610c0e565b9150610c4760208401610c0e565b90509250929050565b600060208284031215610c6257600080fd5b81356001600160401b03811115610c7857600080fd5b610c8484828501610a3a565b949350505050565b600080600060a08486031215610ca157600080fd5b83356001600160401b03811115610cb757600080fd5b610cc386828701610a3a565b935050610cd38560208601610b23565b9150608084013590509250925092565b600060208284031215610cf557600080fd5b61027382610c0e565b803580151581146107ba57600080fd5b600060808284031215610d2057600080fd5b604051608081018181106001600160401b0382111715610d4257610d426106ef565b8060405250809150823581526020830135602082015260408301356040820152610d6e60608401610cfe565b60608201525092915050565b600060a08284031215610d8c57600080fd5b60405160a081018181106001600160401b0382111715610dae57610dae6106ef565b806040525080915082358152602083013560208201526040830135604082015260608301356060820152608083013560808201525092915050565b60008060808385031215610dfc57600080fd5b82356001600160401b0380821115610e1357600080fd5b908401906102208287031215610e2857600080fd5b610e30610750565b82358152610e4060208401610c0e565b6020820152610e5160408401610bdc565b6040820152610e6260608401610bdc565b6060820152608083013582811115610e7957600080fd5b610e8588828601610a3a565b608083015250610e988760a08501610d0e565b60a0820152610120610eac88828601610d7a565b60c08301526101c084013560e0830152610ec96101e08501610cfe565b61010083015261020084013583811115610ee257600080fd5b610eee898287016107bf565b828401525050809450505050610c478460208501610b23565b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761050d5761050d610f07565b8082018082111561050d5761050d610f07565b634e487b7160e01b600052602160045260246000fd5b60098110610f7b57634e487b7160e01b600052602160045260246000fd5b9052565b6020810161050d8284610f5d565b60408101610f9b8285610f5d565b6102736020830184610f5d565b600082610fc557634e487b7160e01b600052601260045260246000fd5b500490565b8181038181111561050d5761050d610f0756fea264697066735822122060f855b9bef4b97e4fa95e561f4142bbcba3953bcb08fd13cc07a8dae5764a3f64736f6c63430008110033",
}

// ToolsABI is the input ABI used to generate the binding from.
// Deprecated: Use ToolsMetaData.ABI instead.
var ToolsABI = ToolsMetaData.ABI

// ToolsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ToolsMetaData.Bin instead.
var ToolsBin = ToolsMetaData.Bin

// DeployTools deploys a new Ethereum contract, binding an instance of Tools to it.
func DeployTools(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tools, error) {
	parsed, err := ToolsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ToolsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tools{ToolsCaller: ToolsCaller{contract: contract}, ToolsTransactor: ToolsTransactor{contract: contract}, ToolsFilterer: ToolsFilterer{contract: contract}}, nil
}

// Tools is an auto generated Go binding around an Ethereum contract.
type Tools struct {
	ToolsCaller     // Read-only binding to the contract
	ToolsTransactor // Write-only binding to the contract
	ToolsFilterer   // Log filterer for contract events
}

// ToolsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ToolsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToolsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ToolsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToolsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ToolsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ToolsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ToolsSession struct {
	Contract     *Tools            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ToolsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ToolsCallerSession struct {
	Contract *ToolsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ToolsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ToolsTransactorSession struct {
	Contract     *ToolsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ToolsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ToolsRaw struct {
	Contract *Tools // Generic contract binding to access the raw methods on
}

// ToolsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ToolsCallerRaw struct {
	Contract *ToolsCaller // Generic read-only contract binding to access the raw methods on
}

// ToolsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ToolsTransactorRaw struct {
	Contract *ToolsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTools creates a new instance of Tools, bound to a specific deployed contract.
func NewTools(address common.Address, backend bind.ContractBackend) (*Tools, error) {
	contract, err := bindTools(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tools{ToolsCaller: ToolsCaller{contract: contract}, ToolsTransactor: ToolsTransactor{contract: contract}, ToolsFilterer: ToolsFilterer{contract: contract}}, nil
}

// NewToolsCaller creates a new read-only instance of Tools, bound to a specific deployed contract.
func NewToolsCaller(address common.Address, caller bind.ContractCaller) (*ToolsCaller, error) {
	contract, err := bindTools(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ToolsCaller{contract: contract}, nil
}

// NewToolsTransactor creates a new write-only instance of Tools, bound to a specific deployed contract.
func NewToolsTransactor(address common.Address, transactor bind.ContractTransactor) (*ToolsTransactor, error) {
	contract, err := bindTools(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ToolsTransactor{contract: contract}, nil
}

// NewToolsFilterer creates a new log filterer instance of Tools, bound to a specific deployed contract.
func NewToolsFilterer(address common.Address, filterer bind.ContractFilterer) (*ToolsFilterer, error) {
	contract, err := bindTools(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ToolsFilterer{contract: contract}, nil
}

// bindTools binds a generic wrapper to an already deployed contract.
func bindTools(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ToolsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tools *ToolsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tools.Contract.ToolsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tools *ToolsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tools.Contract.ToolsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tools *ToolsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tools.Contract.ToolsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tools *ToolsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tools.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tools *ToolsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tools.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tools *ToolsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tools.Contract.contract.Transact(opts, method, params...)
}

// CheckNewJobStatus is a free data retrieval call binding the contract method 0xb42f5513.
//
// Solidity: function checkNewJobStatus(uint8 _currentJobStatus, uint8 _nextJobStatus) pure returns()
func (_Tools *ToolsCaller) CheckNewJobStatus(opts *bind.CallOpts, _currentJobStatus uint8, _nextJobStatus uint8) error {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "checkNewJobStatus", _currentJobStatus, _nextJobStatus)

	if err != nil {
		return err
	}

	return err

}

// CheckNewJobStatus is a free data retrieval call binding the contract method 0xb42f5513.
//
// Solidity: function checkNewJobStatus(uint8 _currentJobStatus, uint8 _nextJobStatus) pure returns()
func (_Tools *ToolsSession) CheckNewJobStatus(_currentJobStatus uint8, _nextJobStatus uint8) error {
	return _Tools.Contract.CheckNewJobStatus(&_Tools.CallOpts, _currentJobStatus, _nextJobStatus)
}

// CheckNewJobStatus is a free data retrieval call binding the contract method 0xb42f5513.
//
// Solidity: function checkNewJobStatus(uint8 _currentJobStatus, uint8 _nextJobStatus) pure returns()
func (_Tools *ToolsCallerSession) CheckNewJobStatus(_currentJobStatus uint8, _nextJobStatus uint8) error {
	return _Tools.Contract.CheckNewJobStatus(&_Tools.CallOpts, _currentJobStatus, _nextJobStatus)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_Tools *ToolsCaller) ConvertAddressToBytes32(opts *bind.CallOpts, _address common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "convertAddressToBytes32", _address)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_Tools *ToolsSession) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _Tools.Contract.ConvertAddressToBytes32(&_Tools.CallOpts, _address)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_Tools *ToolsCallerSession) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _Tools.Contract.ConvertAddressToBytes32(&_Tools.CallOpts, _address)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes) pure returns(address)
func (_Tools *ToolsCaller) ConvertBytes32ToAddress(opts *bind.CallOpts, _bytes [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "convertBytes32ToAddress", _bytes)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes) pure returns(address)
func (_Tools *ToolsSession) ConvertBytes32ToAddress(_bytes [32]byte) (common.Address, error) {
	return _Tools.Contract.ConvertBytes32ToAddress(&_Tools.CallOpts, _bytes)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes) pure returns(address)
func (_Tools *ToolsCallerSession) ConvertBytes32ToAddress(_bytes [32]byte) (common.Address, error) {
	return _Tools.Contract.ConvertBytes32ToAddress(&_Tools.CallOpts, _bytes)
}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0x3df1c0ff.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition, (uint256,uint256,uint256) providerPrices, uint256 amountLocked) pure returns(uint64)
func (_Tools *ToolsCaller) ConvertCreditToDuration(opts *bind.CallOpts, jobDefinition JobDefinition, providerPrices ProviderPrices, amountLocked *big.Int) (uint64, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "convertCreditToDuration", jobDefinition, providerPrices, amountLocked)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0x3df1c0ff.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition, (uint256,uint256,uint256) providerPrices, uint256 amountLocked) pure returns(uint64)
func (_Tools *ToolsSession) ConvertCreditToDuration(jobDefinition JobDefinition, providerPrices ProviderPrices, amountLocked *big.Int) (uint64, error) {
	return _Tools.Contract.ConvertCreditToDuration(&_Tools.CallOpts, jobDefinition, providerPrices, amountLocked)
}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0x3df1c0ff.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition, (uint256,uint256,uint256) providerPrices, uint256 amountLocked) pure returns(uint64)
func (_Tools *ToolsCallerSession) ConvertCreditToDuration(jobDefinition JobDefinition, providerPrices ProviderPrices, amountLocked *big.Int) (uint64, error) {
	return _Tools.Contract.ConvertCreditToDuration(&_Tools.CallOpts, jobDefinition, providerPrices, amountLocked)
}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0xb22d1f1f.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition, (uint256,uint256,uint256) providerPrices, uint64 durationMinute) pure returns(uint256)
func (_Tools *ToolsCaller) ConvertDurationToCredit(opts *bind.CallOpts, jobDefinition JobDefinition, providerPrices ProviderPrices, durationMinute uint64) (*big.Int, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "convertDurationToCredit", jobDefinition, providerPrices, durationMinute)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0xb22d1f1f.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition, (uint256,uint256,uint256) providerPrices, uint64 durationMinute) pure returns(uint256)
func (_Tools *ToolsSession) ConvertDurationToCredit(jobDefinition JobDefinition, providerPrices ProviderPrices, durationMinute uint64) (*big.Int, error) {
	return _Tools.Contract.ConvertDurationToCredit(&_Tools.CallOpts, jobDefinition, providerPrices, durationMinute)
}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0xb22d1f1f.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) jobDefinition, (uint256,uint256,uint256) providerPrices, uint64 durationMinute) pure returns(uint256)
func (_Tools *ToolsCallerSession) ConvertDurationToCredit(jobDefinition JobDefinition, providerPrices ProviderPrices, durationMinute uint64) (*big.Int, error) {
	return _Tools.Contract.ConvertDurationToCredit(&_Tools.CallOpts, jobDefinition, providerPrices, durationMinute)
}

// GetRemainingTime is a free data retrieval call binding the contract method 0xc686be48.
//
// Solidity: function getRemainingTime((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) job, (uint256,uint256,uint256) providerPrices) view returns(uint64)
func (_Tools *ToolsCaller) GetRemainingTime(opts *bind.CallOpts, job Job, providerPrices ProviderPrices) (uint64, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "getRemainingTime", job, providerPrices)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRemainingTime is a free data retrieval call binding the contract method 0xc686be48.
//
// Solidity: function getRemainingTime((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) job, (uint256,uint256,uint256) providerPrices) view returns(uint64)
func (_Tools *ToolsSession) GetRemainingTime(job Job, providerPrices ProviderPrices) (uint64, error) {
	return _Tools.Contract.GetRemainingTime(&_Tools.CallOpts, job, providerPrices)
}

// GetRemainingTime is a free data retrieval call binding the contract method 0xc686be48.
//
// Solidity: function getRemainingTime((bytes32,uint8,address,address,(uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]),(uint256,uint256,uint256,bool),(uint256,uint256,uint256,uint256,uint256),bytes32,bool,string) job, (uint256,uint256,uint256) providerPrices) view returns(uint64)
func (_Tools *ToolsCallerSession) GetRemainingTime(job Job, providerPrices ProviderPrices) (uint64, error) {
	return _Tools.Contract.GetRemainingTime(&_Tools.CallOpts, job, providerPrices)
}

// IsDelegateTopable is a free data retrieval call binding the contract method 0x6e404ce0.
//
// Solidity: function isDelegateTopable(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsCaller) IsDelegateTopable(opts *bind.CallOpts, _jobStatus uint8) (bool, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "isDelegateTopable", _jobStatus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegateTopable is a free data retrieval call binding the contract method 0x6e404ce0.
//
// Solidity: function isDelegateTopable(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsSession) IsDelegateTopable(_jobStatus uint8) (bool, error) {
	return _Tools.Contract.IsDelegateTopable(&_Tools.CallOpts, _jobStatus)
}

// IsDelegateTopable is a free data retrieval call binding the contract method 0x6e404ce0.
//
// Solidity: function isDelegateTopable(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsCallerSession) IsDelegateTopable(_jobStatus uint8) (bool, error) {
	return _Tools.Contract.IsDelegateTopable(&_Tools.CallOpts, _jobStatus)
}

// IsJobCold is a free data retrieval call binding the contract method 0x85521276.
//
// Solidity: function isJobCold(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsCaller) IsJobCold(opts *bind.CallOpts, _jobStatus uint8) (bool, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "isJobCold", _jobStatus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsJobCold is a free data retrieval call binding the contract method 0x85521276.
//
// Solidity: function isJobCold(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsSession) IsJobCold(_jobStatus uint8) (bool, error) {
	return _Tools.Contract.IsJobCold(&_Tools.CallOpts, _jobStatus)
}

// IsJobCold is a free data retrieval call binding the contract method 0x85521276.
//
// Solidity: function isJobCold(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsCallerSession) IsJobCold(_jobStatus uint8) (bool, error) {
	return _Tools.Contract.IsJobCold(&_Tools.CallOpts, _jobStatus)
}

// IsJobDefinitionValid is a free data retrieval call binding the contract method 0xceb9b3cb.
//
// Solidity: function isJobDefinitionValid((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _jobDefinition) pure returns(bool)
func (_Tools *ToolsCaller) IsJobDefinitionValid(opts *bind.CallOpts, _jobDefinition JobDefinition) (bool, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "isJobDefinitionValid", _jobDefinition)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsJobDefinitionValid is a free data retrieval call binding the contract method 0xceb9b3cb.
//
// Solidity: function isJobDefinitionValid((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _jobDefinition) pure returns(bool)
func (_Tools *ToolsSession) IsJobDefinitionValid(_jobDefinition JobDefinition) (bool, error) {
	return _Tools.Contract.IsJobDefinitionValid(&_Tools.CallOpts, _jobDefinition)
}

// IsJobDefinitionValid is a free data retrieval call binding the contract method 0xceb9b3cb.
//
// Solidity: function isJobDefinitionValid((uint64,uint64,uint64,uint64,string,uint8,(string,string)[],((string,string),bytes2)[]) _jobDefinition) pure returns(bool)
func (_Tools *ToolsCallerSession) IsJobDefinitionValid(_jobDefinition JobDefinition) (bool, error) {
	return _Tools.Contract.IsJobDefinitionValid(&_Tools.CallOpts, _jobDefinition)
}

// IsJobHot is a free data retrieval call binding the contract method 0xff75c979.
//
// Solidity: function isJobHot(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsCaller) IsJobHot(opts *bind.CallOpts, _jobStatus uint8) (bool, error) {
	var out []interface{}
	err := _Tools.contract.Call(opts, &out, "isJobHot", _jobStatus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsJobHot is a free data retrieval call binding the contract method 0xff75c979.
//
// Solidity: function isJobHot(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsSession) IsJobHot(_jobStatus uint8) (bool, error) {
	return _Tools.Contract.IsJobHot(&_Tools.CallOpts, _jobStatus)
}

// IsJobHot is a free data retrieval call binding the contract method 0xff75c979.
//
// Solidity: function isJobHot(uint8 _jobStatus) pure returns(bool)
func (_Tools *ToolsCallerSession) IsJobHot(_jobStatus uint8) (bool, error) {
	return _Tools.Contract.IsJobHot(&_Tools.CallOpts, _jobStatus)
}
