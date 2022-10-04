// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metascheduler

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
)

// Job is an auto generated low-level Go binding around an user-defined struct.
type Job struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Schedulable            bool
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
}

// JobDefinition is an auto generated low-level Go binding around an user-defined struct.
type JobDefinition struct {
	GpuPerNode        uint64
	MemPerNode        uint64
	CpuPerTask        uint64
	Nodes             uint64
	Ntasks            uint64
	BatchLocationHash string
}

// Provider is an auto generated low-level Go binding around an user-defined struct.
type Provider struct {
	Definition ProviderDefinition
	Status     uint8
	Valid      bool
	JobCount   uint64
}

// ProviderDefinition is an auto generated low-level Go binding around an user-defined struct.
type ProviderDefinition struct {
	Nodes          uint64
	Gpus           uint64
	GpuPricePerMin uint64
	Cpus           uint64
	CpuPricePerMin uint64
	Mem            uint64
	MemPricePerMin uint64
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
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a68a65a692984f4f4604da149948b16f5fe34139df3f7bc08afdc880e6a5fa3364736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(DoubleEndedQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(IAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// IERC20MetadataMetaData contains all meta data concerning the IERC20Metadata contract.
var IERC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetadataMetaData.ABI instead.
var IERC20MetadataABI = IERC20MetadataMetaData.ABI

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, from, to, amount)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
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
		it.Event = new(IERC20MetadataApproval)
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
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
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
		it.Event = new(IERC20MetadataTransfer)
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
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerMetaData contains all meta data concerning the JobManager contract.
var JobManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimNextJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"customerJobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"finishJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobFromId\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"schedulable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"}],\"internalType\":\"structJob\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobId2CustomerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobId2ProviderAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"schedulable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metaQueue\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerJobs\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"startJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"triggerFailed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateJobsStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_metaSchedulerAddr\",\"type\":\"address\"}],\"name\":\"updateRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50620000276000801b336200002d60201b60201c565b62000190565b6200003f82826200011e60201b60201c565b6200011a57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620000bf6200018860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b614dd380620001a06000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80635792edd4116100de578063b3130fba11610097578063c58467b011610071578063c58467b0146104d3578063d1cee54614610504578063d547741f14610520578063e3401e001461053c5761018e565b8063b3130fba14610457578063b3ba274e14610487578063c4d252f5146104b75761018e565b80635792edd4146103715780636f3de6421461038d5780637d17544b146103bd5780638fb70f63146103d957806391d1485414610409578063a217fddf146104395761018e565b80632f2ff15d1161014b5780633a80760a116101255780633a80760a146102e95780634609ca501461030557806346200b6b146103215780634f1a1106146103405761018e565b80632f2ff15d1461027957806336568abe1461029557806338ed7cfc146102b15761018e565b8063010a71941461019357806301ffc9a7146101c35780630764e1cd146101f35780630cb855271461020f578063248a9ca31461023f5780632a242a761461026f575b600080fd5b6101ad60048036038101906101a891906137ef565b610558565b6040516101ba9190613877565b60405180910390f35b6101dd60048036038101906101d891906138ea565b610903565b6040516101ea9190613932565b60405180910390f35b61020d60048036038101906102089190613979565b61097d565b005b610229600480360381019061022491906139b9565b610b56565b60405161023691906139f5565b60405180910390f35b610259600480360381019061025491906139b9565b610b89565b6040516102669190613877565b60405180910390f35b610277610ba8565b005b610293600480360381019061028e9190613a10565b610f6c565b005b6102af60048036038101906102aa9190613a10565b610f8d565b005b6102cb60048036038101906102c691906139b9565b611010565b6040516102e099989796959493929190613bed565b60405180910390f35b61030360048036038101906102fe9190613c81565b61125b565b005b61031f600480360381019061031a9190613a10565b611278565b005b61032961142f565b604051610337929190613cca565b60405180910390f35b61035a60048036038101906103559190613c81565b61145b565b604051610368929190613cca565b60405180910390f35b61038b60048036038101906103869190613a10565b611499565b005b6103a760048036038101906103a291906139b9565b611617565b6040516103b491906139f5565b60405180910390f35b6103d760048036038101906103d29190613a10565b61164a565b005b6103f360048036038101906103ee9190613a10565b6117c8565b6040516104009190613cf3565b60405180910390f35b610423600480360381019061041e9190613a10565b611dad565b6040516104309190613932565b60405180910390f35b610441611e17565b60405161044e9190613877565b60405180910390f35b610471600480360381019061046c91906139b9565b611e1e565b60405161047e9190613eae565b60405180910390f35b6104a1600480360381019061049c9190613ed0565b61217e565b6040516104ae9190613877565b60405180910390f35b6104d160048036038101906104cc91906139b9565b6121af565b005b6104ed60048036038101906104e89190613c81565b612261565b6040516104fb929190613f10565b60405180910390f35b61051e60048036038101906105199190613a10565b612920565b005b61053a60048036038101906105359190613a10565b612ad2565b005b61055660048036038101906105519190613a10565b612af3565b005b60006105676000801b33612cc2565b6000844260405160200161057c929190613fa9565b6040516020818303038152906040528051906020012090505b6006600082815260200190815260200160002060080160009054906101000a900460ff16806105c657506000801b81145b156105f957806040516020016105dc9190613ff6565b604051602081830303815290604052805190602001209050610595565b6040518061012001604052808281526020016000600681111561061f5761061e613a50565b5b81526020018481526020018673ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020018581526020016001151581526020016000815250600660008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff021916908360068111156106c8576106c7613a50565b5b02179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040160146101000a81548160ff02191690831515021790555060c08201518160050160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a082015181600201908161088e919061421d565b50505060e08201518160080160006101000a81548160ff02191690831515021790555061010082015181600901559050507f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada381866040516108f09291906142ef565b60405180910390a1809150509392505050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610976575061097582612d5f565b5b9050919050565b61098a6000801b33612cc2565b6006600083815260200190815260200160002060080160009054906101000a900460ff166109ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e490614375565b60405180910390fd5b60006006811115610a0157610a00613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff166006811115610a3757610a36613a50565b5b1480610a8a575060016006811115610a5257610a51613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff166006811115610a8857610a87613a50565b5b145b80610adc575060026006811115610aa457610aa3613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff166006811115610ada57610ad9613a50565b5b145b610b1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b12906143e1565b60405180910390fd5b8067ffffffffffffffff16600660008481526020019081526020016000206002016000828254610b4b9190614430565b925050819055505050565b60056020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000838152602001908152602001600020600101549050919050565b610bb56000801b33612cc2565b60005b6007805490508167ffffffffffffffff161015610f69576000801b60078267ffffffffffffffff1681548110610bf157610bf0614464565b5b90600052602060002001540315610f565760016006811115610c1657610c15613a50565b5b6006600060078467ffffffffffffffff1681548110610c3857610c37614464565b5b9060005260206000200154815260200190815260200160002060010160009054906101000a900460ff166006811115610c7457610c73613a50565b5b14610cae576000801b60078267ffffffffffffffff1681548110610c9b57610c9a614464565b5b9060005260206000200181905550610f55565b610cdf60078267ffffffffffffffff1681548110610ccf57610cce614464565b5b9060005260206000200154612dc9565b15610f5457600060078267ffffffffffffffff1681548110610d0457610d03614464565b5b9060005260206000200154905060006006600083815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000801b60078467ffffffffffffffff1681548110610d6e57610d6d614464565b5b90600052602060002001819055505b610dc4600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612e56565b610e58576000610e11600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612e8b565b905060006006600083815260200190815260200160002060010160006101000a81548160ff02191690836006811115610e4d57610e4c613a50565b5b021790555050610d7d565b60006006811115610e6c57610e6b613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff166006811115610ea257610ea1613a50565b5b14610ee2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed99061452b565b60405180910390fd5b7f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3826006600085815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051610f499291906142ef565b60405180910390a150505b5b5b8080610f619061454b565b915050610bb8565b50565b610f7582610b89565b610f7e81612f67565b610f888383612f7b565b505050565b610f9561305b565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611002576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff9906145ed565b60405180910390fd5b61100c8282613063565b5050565b60066020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160149054906101000a900460ff1690806005016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820180546111bb90614040565b80601f01602080910402602001604051908101604052809291908181526020018280546111e790614040565b80156112345780601f1061120957610100808354040283529160200191611234565b820191906000526020600020905b81548152906001019060200180831161121757829003601f168201915b505050505081525050908060080160009054906101000a900460ff16908060090154905089565b6112686000801b33612cc2565b6112756000801b82612f7b565b50565b6112856000801b33612cc2565b8073ffffffffffffffffffffffffffffffffffffffff166006600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611329576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113209061467f565b60405180910390fd5b6002600681111561133d5761133c613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff16600681111561137357611372613a50565b5b146113b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113aa906146eb565b60405180910390fd5b60006006600084815260200190815260200160002060010160006101000a81548160ff021916908360068111156113ed576113ec613a50565b5b02179055507f5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada082826040516114239291906142ef565b60405180910390a15050565b600a8060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b60086020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6114a66000801b33612cc2565b8073ffffffffffffffffffffffffffffffffffffffff166006600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461154a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115419061467f565b60405180910390fd5b6002600681111561155e5761155d613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff16600681111561159457611593613a50565b5b146115d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115cb906146eb565b60405180910390fd5b60036006600084815260200190815260200160002060010160006101000a81548160ff0219169083600681111561160e5761160d613a50565b5b02179055505050565b60046020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6116576000801b33612cc2565b8073ffffffffffffffffffffffffffffffffffffffff166006600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146116fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116f29061477d565b60405180910390fd5b6003600681111561170f5761170e613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff16600681111561174557611744613a50565b5b14611785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161177c906147e9565b60405180910390fd5b60056006600084815260200190815260200160002060010160006101000a81548160ff021916908360068111156117bf576117be613a50565b5b02179055505050565b6117d06133fc565b6117dd6000801b33612cc2565b60006006600085815260200190815260200160002060405180610120016040529081600082015481526020016001820160009054906101000a900460ff16600681111561182d5761182c613a50565b5b600681111561183f5761183e613a50565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460ff16151515158152602001600582016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282018054611a2e90614040565b80601f0160208091040260200160405190810160405280929190818152602001828054611a5a90614040565b8015611aa75780601f10611a7c57610100808354040283529160200191611aa7565b820191906000526020600020905b815481529060010190602001808311611a8a57829003601f168201915b50505050508152505081526020016008820160009054906101000a900460ff1615151515815260200160098201548152505090508273ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff1614611b4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b449061487b565b60405180910390fd5b60016006811115611b6157611b60613a50565b5b81602001516006811115611b7857611b77613a50565b5b14611bb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611baf9061490d565b60405180910390fd5b60026006600086815260200190815260200160002060010160006101000a81548160ff02191690836006811115611bf257611bf1613a50565b5b0217905550600660008581526020019081526020016000206005016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282018054611d2290614040565b80601f0160208091040260200160405190810160405280929190818152602001828054611d4e90614040565b8015611d9b5780601f10611d7057610100808354040283529160200191611d9b565b820191906000526020600020905b815481529060010190602001808311611d7e57829003601f168201915b50505050508152505091505092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b611e26613464565b611e336000801b33612cc2565b60006006600084815260200190815260200160002060405180610120016040529081600082015481526020016001820160009054906101000a900460ff166006811115611e8357611e82613a50565b5b6006811115611e9557611e94613a50565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460ff16151515158152602001600582016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160028201805461208490614040565b80601f01602080910402602001604051908101604052809291908181526020018280546120b090614040565b80156120fd5780601f106120d2576101008083540402835291602001916120fd565b820191906000526020600020905b8154815290600101906020018083116120e057829003601f168201915b50505050508152505081526020016008820160009054906101000a900460ff1615151515815260200160098201548152505090508060e00151612175576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161216c90614979565b60405180910390fd5b80915050919050565b6009602052816000526040600020818154811061219a57600080fd5b90600052602060002001600091509150505481565b6121bc6000801b33612cc2565b6006600082815260200190815260200160002060080160009054906101000a900460ff1661221f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161221690614375565b60405180910390fd5b60046006600083815260200190815260200160002060010160006101000a81548160ff0219169083600681111561225957612258613a50565b5b021790555050565b600061226b6133fc565b6122786000801b33612cc2565b6122bf600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612e56565b156122ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122f6906149e5565b60405180910390fd5b6000612348600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612e8b565b905060006006600083815260200190815260200160002060405180610120016040529081600082015481526020016001820160009054906101000a900460ff16600681111561239a57612399613a50565b5b60068111156123ac576123ab613a50565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460ff16151515158152602001600582016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160028201805461259b90614040565b80601f01602080910402602001604051908101604052809291908181526020018280546125c790614040565b80156126145780601f106125e957610100808354040283529160200191612614565b820191906000526020600020905b8154815290600101906020018083116125f757829003601f168201915b50505050508152505081526020016008820160009054906101000a900460ff1615151515815260200160098201548152505090508473ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16146126ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126b19061487b565b60405180910390fd5b600160068111156126ce576126cd613a50565b5b816020015160068111156126e5576126e4613a50565b5b14612725576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161271c9061490d565b60405180910390fd5b60026006600084815260200190815260200160002060010160006101000a81548160ff0219169083600681111561275f5761275e613a50565b5b02179055508160066000848152602001908152602001600020600501806040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160028201805461289190614040565b80601f01602080910402602001604051908101604052809291908181526020018280546128bd90614040565b801561290a5780601f106128df5761010080835404028352916020019161290a565b820191906000526020600020905b8154815290600101906020018083116128ed57829003601f168201915b5050505050815250509050935093505050915091565b61292d6000801b33612cc2565b6006600083815260200190815260200160002060080160009054906101000a900460ff1661295a57600080fd5b6000600681111561296e5761296d613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff1660068111156129a4576129a3613a50565b5b146129ae57600080fd5b806006600084815260200190815260200160002060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016006600084815260200190815260200160002060010160006101000a81548160ff02191690836006811115612a3d57612a3c613a50565b5b02179055504360066000848152602001908152602001600020600901819055506007829080600181540180825580915050600190039060005260206000200160009091909190915055612ace600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083613144565b5050565b612adb82610b89565b612ae481612f67565b612aee8383613063565b505050565b612b006000801b33612cc2565b8073ffffffffffffffffffffffffffffffffffffffff166006600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612ba4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b9b90614a77565b60405180910390fd5b60036006811115612bb857612bb7613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff166006811115612bee57612bed613a50565b5b1480612c41575060026006811115612c0957612c08613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff166006811115612c3f57612c3e613a50565b5b145b612c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c7790614b09565b60405180910390fd5b600680600084815260200190815260200160002060010160006101000a81548160ff02191690836006811115612cb957612cb8613a50565b5b02179055505050565b612ccc8282611dad565b612d5b57612cf18173ffffffffffffffffffffffffffffffffffffffff1660146131c0565b612cff8360001c60206131c0565b604051602001612d10929190614bfd565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d529190614c70565b60405180910390fd5b5050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600060016006811115612ddf57612dde613a50565b5b6006600084815260200190815260200160002060010160009054906101000a900460ff166006811115612e1557612e14613a50565b5b14612e1f57600080fd5b600a67ffffffffffffffff16436006600085815260200190815260200160002060090154612e4d9190614c92565b10159050919050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b13159050919050565b6000612e9682612e56565b15612ecd576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260000160009054906101000a9004600f0b905082600101600082600f0b600f0b815260200190815260200160002054915082600101600082600f0b600f0b815260200190815260200160002060009055600181018360000160006101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff16021790555050919050565b612f7881612f7361305b565b612cc2565b50565b612f858282611dad565b61305757600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612ffc61305b565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b61306d8282611dad565b1561314057600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506130e561305b565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b60008260000160109054906101000a9004600f0b90508183600101600083600f0b600f0b815260200190815260200160002081905550600181018360000160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff160217905550505050565b6060600060028360026131d39190614cc6565b6131dd9190614430565b67ffffffffffffffff8111156131f6576131f5613583565b5b6040519080825280601f01601f1916602001820160405280156132285781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106132605761325f614464565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106132c4576132c3614464565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026133049190614cc6565b61330e9190614430565b90505b60018111156133ae577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106133505761334f614464565b5b1a60f81b82828151811061336757613366614464565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806133a790614d08565b9050613311565b50600084146133f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133e990614d7d565b60405180910390fd5b8091505092915050565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b604051806101200160405280600080191681526020016000600681111561348e5761348d613a50565b5b815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020016134e56133fc565b8152602001600015158152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061353a8261350f565b9050919050565b61354a8161352f565b811461355557600080fd5b50565b60008135905061356781613541565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6135bb82613572565b810181811067ffffffffffffffff821117156135da576135d9613583565b5b80604052505050565b60006135ed6134fb565b90506135f982826135b2565b919050565b600080fd5b600067ffffffffffffffff82169050919050565b61362081613603565b811461362b57600080fd5b50565b60008135905061363d81613617565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff82111561366857613667613583565b5b61367182613572565b9050602081019050919050565b82818337600083830152505050565b60006136a061369b8461364d565b6135e3565b9050828152602081018484840111156136bc576136bb613648565b5b6136c784828561367e565b509392505050565b600082601f8301126136e4576136e3613643565b5b81356136f484826020860161368d565b91505092915050565b600060c082840312156137135761371261356d565b5b61371d60c06135e3565b9050600061372d8482850161362e565b60008301525060206137418482850161362e565b60208301525060406137558482850161362e565b60408301525060606137698482850161362e565b606083015250608061377d8482850161362e565b60808301525060a082013567ffffffffffffffff8111156137a1576137a06135fe565b5b6137ad848285016136cf565b60a08301525092915050565b6000819050919050565b6137cc816137b9565b81146137d757600080fd5b50565b6000813590506137e9816137c3565b92915050565b60008060006060848603121561380857613807613505565b5b600061381686828701613558565b935050602084013567ffffffffffffffff8111156138375761383661350a565b5b613843868287016136fd565b9250506040613854868287016137da565b9150509250925092565b6000819050919050565b6138718161385e565b82525050565b600060208201905061388c6000830184613868565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6138c781613892565b81146138d257600080fd5b50565b6000813590506138e4816138be565b92915050565b600060208284031215613900576138ff613505565b5b600061390e848285016138d5565b91505092915050565b60008115159050919050565b61392c81613917565b82525050565b60006020820190506139476000830184613923565b92915050565b6139568161385e565b811461396157600080fd5b50565b6000813590506139738161394d565b92915050565b600080604083850312156139905761398f613505565b5b600061399e85828601613964565b92505060206139af8582860161362e565b9150509250929050565b6000602082840312156139cf576139ce613505565b5b60006139dd84828501613964565b91505092915050565b6139ef8161352f565b82525050565b6000602082019050613a0a60008301846139e6565b92915050565b60008060408385031215613a2757613a26613505565b5b6000613a3585828601613964565b9250506020613a4685828601613558565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60078110613a9057613a8f613a50565b5b50565b6000819050613aa182613a7f565b919050565b6000613ab182613a93565b9050919050565b613ac181613aa6565b82525050565b613ad0816137b9565b82525050565b613adf81613603565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613b1f578082015181840152602081019050613b04565b60008484015250505050565b6000613b3682613ae5565b613b408185613af0565b9350613b50818560208601613b01565b613b5981613572565b840191505092915050565b600060c083016000830151613b7c6000860182613ad6565b506020830151613b8f6020860182613ad6565b506040830151613ba26040860182613ad6565b506060830151613bb56060860182613ad6565b506080830151613bc86080860182613ad6565b5060a083015184820360a0860152613be08282613b2b565b9150508091505092915050565b600061012082019050613c03600083018c613868565b613c10602083018b613ab8565b613c1d604083018a613ac7565b613c2a60608301896139e6565b613c3760808301886139e6565b613c4460a0830187613923565b81810360c0830152613c568186613b64565b9050613c6560e0830185613923565b613c73610100830184613ac7565b9a9950505050505050505050565b600060208284031215613c9757613c96613505565b5b6000613ca584828501613558565b91505092915050565b600081600f0b9050919050565b613cc481613cae565b82525050565b6000604082019050613cdf6000830185613cbb565b613cec6020830184613cbb565b9392505050565b60006020820190508181036000830152613d0d8184613b64565b905092915050565b613d1e8161385e565b82525050565b613d2d81613aa6565b82525050565b613d3c816137b9565b82525050565b613d4b8161352f565b82525050565b613d5a81613917565b82525050565b600060c083016000830151613d786000860182613ad6565b506020830151613d8b6020860182613ad6565b506040830151613d9e6040860182613ad6565b506060830151613db16060860182613ad6565b506080830151613dc46080860182613ad6565b5060a083015184820360a0860152613ddc8282613b2b565b9150508091505092915050565b600061012083016000830151613e026000860182613d15565b506020830151613e156020860182613d24565b506040830151613e286040860182613d33565b506060830151613e3b6060860182613d42565b506080830151613e4e6080860182613d42565b5060a0830151613e6160a0860182613d51565b5060c083015184820360c0860152613e798282613d60565b91505060e0830151613e8e60e0860182613d51565b50610100830151613ea3610100860182613d33565b508091505092915050565b60006020820190508181036000830152613ec88184613de9565b905092915050565b60008060408385031215613ee757613ee6613505565b5b6000613ef585828601613558565b9250506020613f06858286016137da565b9150509250929050565b6000604082019050613f256000830185613868565b8181036020830152613f378184613b64565b90509392505050565b60008160601b9050919050565b6000613f5882613f40565b9050919050565b6000613f6a82613f4d565b9050919050565b613f82613f7d8261352f565b613f5f565b82525050565b6000819050919050565b613fa3613f9e826137b9565b613f88565b82525050565b6000613fb58285613f71565b601482019150613fc58284613f92565b6020820191508190509392505050565b6000819050919050565b613ff0613feb8261385e565b613fd5565b82525050565b60006140028284613fdf565b60208201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061405857607f821691505b60208210810361406b5761406a614011565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026140d37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614096565b6140dd8683614096565b95508019841693508086168417925050509392505050565b6000819050919050565b600061411a614115614110846137b9565b6140f5565b6137b9565b9050919050565b6000819050919050565b614134836140ff565b61414861414082614121565b8484546140a3565b825550505050565b600090565b61415d614150565b61416881848461412b565b505050565b5b8181101561418c57614181600082614155565b60018101905061416e565b5050565b601f8211156141d1576141a281614071565b6141ab84614086565b810160208510156141ba578190505b6141ce6141c685614086565b83018261416d565b50505b505050565b600082821c905092915050565b60006141f4600019846008026141d6565b1980831691505092915050565b600061420d83836141e3565b9150826002028217905092915050565b61422682613ae5565b67ffffffffffffffff81111561423f5761423e613583565b5b6142498254614040565b614254828285614190565b600060209050601f8311600181146142875760008415614275578287015190505b61427f8582614201565b8655506142e7565b601f19841661429586614071565b60005b828110156142bd57848901518255600182019150602085019450602081019050614298565b868310156142da57848901516142d6601f8916826141e3565b8355505b6001600288020188555050505b505050505050565b60006040820190506143046000830185613868565b61431160208301846139e6565b9392505050565b600082825260208201905092915050565b7f4a6f62206d757374206578697374730000000000000000000000000000000000600082015250565b600061435f600f83614318565b915061436a82614329565b602082019050919050565b6000602082019050818103600083015261438e81614352565b9050919050565b7f4a6f62206d7573742062652072756e6e696e6700000000000000000000000000600082015250565b60006143cb601383614318565b91506143d682614395565b602082019050919050565b600060208201905081810360008301526143fa816143be565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061443b826137b9565b9150614446836137b9565b925082820190508082111561445e5761445d614401565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f54686973206a6f62206d7573742062652050454e44494e47207468657265206960008201527f732061206d69736d61746368206265747765656e2070726f76696465724a6f6260208201527f7320616e6420686f744a6f624c69737400000000000000000000000000000000604082015250565b6000614515605083614318565b915061452082614493565b606082019050919050565b6000602082019050818103600083015261454481614508565b9050919050565b600061455682613603565b915067ffffffffffffffff82036145705761456f614401565b5b600182019050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b60006145d7602f83614318565b91506145e28261457b565b604082019050919050565b60006020820190508181036000830152614606816145ca565b9050919050565b7f50726f7669646572732063616e206f6e6c79207374617274207468656972206a60008201527f6f62730000000000000000000000000000000000000000000000000000000000602082015250565b6000614669602383614318565b91506146748261460d565b604082019050919050565b600060208201905081810360008301526146988161465c565b9050919050565b7f4f6e6c7920717565756564206a6f622063616e20626520737461727465640000600082015250565b60006146d5601e83614318565b91506146e08261469f565b602082019050919050565b60006020820190508181036000830152614704816146c8565b9050919050565b7f50726f7669646572732063616e206f6e6c792066696e6973682074686569722060008201527f6a6f627300000000000000000000000000000000000000000000000000000000602082015250565b6000614767602483614318565b91506147728261470b565b604082019050919050565b600060208201905081810360008301526147968161475a565b9050919050565b7f4f6e6c792072756e6e696e67206a6f622063616e2062652066696e6973686564600082015250565b60006147d3602083614318565b91506147de8261479d565b602082019050919050565b60006020820190508181036000830152614802816147c6565b9050919050565b7f50726f7669646572732063616e206f6e6c7920636c61696d207468656972206a60008201527f6f62730000000000000000000000000000000000000000000000000000000000602082015250565b6000614865602383614318565b915061487082614809565b604082019050919050565b6000602082019050818103600083015261489481614858565b9050919050565b7f4f6e6c79206d6574612d717565756564206a6f622063616e206265207175657560008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b60006148f7602283614318565b91506149028261489b565b604082019050919050565b60006020820190508181036000830152614926816148ea565b9050919050565b7f4a6f62206e6f7420666f756e6400000000000000000000000000000000000000600082015250565b6000614963600d83614318565b915061496e8261492d565b602082019050919050565b6000602082019050818103600083015261499281614956565b9050919050565b7f4e6f20617661696c61626c65206a6f6200000000000000000000000000000000600082015250565b60006149cf601083614318565b91506149da82614999565b602082019050919050565b600060208201905081810360008301526149fe816149c2565b9050919050565b7f50726f7669646572732063616e206f6e6c792074726967676572206661696c7560008201527f7265206f66206a6f622074686579206c61756e63686564000000000000000000602082015250565b6000614a61603783614318565b9150614a6c82614a05565b604082019050919050565b60006020820190508181036000830152614a9081614a54565b9050919050565b7f4f6e6c79207363686564756c656420616e642072756e6e696e67206a6f62206360008201527f616e206661696c00000000000000000000000000000000000000000000000000602082015250565b6000614af3602783614318565b9150614afe82614a97565b604082019050919050565b60006020820190508181036000830152614b2281614ae6565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000614b6a601783614b29565b9150614b7582614b34565b601782019050919050565b6000614b8b82613ae5565b614b958185614b29565b9350614ba5818560208601613b01565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000614be7601183614b29565b9150614bf282614bb1565b601182019050919050565b6000614c0882614b5d565b9150614c148285614b80565b9150614c1f82614bda565b9150614c2b8284614b80565b91508190509392505050565b6000614c4282613ae5565b614c4c8185614318565b9350614c5c818560208601613b01565b614c6581613572565b840191505092915050565b60006020820190508181036000830152614c8a8184614c37565b905092915050565b6000614c9d826137b9565b9150614ca8836137b9565b9250828203905081811115614cc057614cbf614401565b5b92915050565b6000614cd1826137b9565b9150614cdc836137b9565b9250828202614cea816137b9565b91508282048414831517614d0157614d00614401565b5b5092915050565b6000614d13826137b9565b915060008203614d2657614d25614401565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000614d67602083614318565b9150614d7282614d31565b602082019050919050565b60006020820190508181036000830152614d9681614d5a565b905091905056fea2646970667358221220aec858db16c24828eb8b9bde32803fba5dba42133895fdabe4cd22c9eba2892264736f6c63430008110033",
}

// JobManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use JobManagerMetaData.ABI instead.
var JobManagerABI = JobManagerMetaData.ABI

// JobManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobManagerMetaData.Bin instead.
var JobManagerBin = JobManagerMetaData.Bin

// DeployJobManager deploys a new Ethereum contract, binding an instance of JobManager to it.
func DeployJobManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JobManager, error) {
	parsed, err := JobManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JobManager{JobManagerCaller: JobManagerCaller{contract: contract}, JobManagerTransactor: JobManagerTransactor{contract: contract}, JobManagerFilterer: JobManagerFilterer{contract: contract}}, nil
}

// JobManager is an auto generated Go binding around an Ethereum contract.
type JobManager struct {
	JobManagerCaller     // Read-only binding to the contract
	JobManagerTransactor // Write-only binding to the contract
	JobManagerFilterer   // Log filterer for contract events
}

// JobManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobManagerSession struct {
	Contract     *JobManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobManagerCallerSession struct {
	Contract *JobManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobManagerTransactorSession struct {
	Contract     *JobManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobManagerRaw struct {
	Contract *JobManager // Generic contract binding to access the raw methods on
}

// JobManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobManagerCallerRaw struct {
	Contract *JobManagerCaller // Generic read-only contract binding to access the raw methods on
}

// JobManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobManagerTransactorRaw struct {
	Contract *JobManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobManager creates a new instance of JobManager, bound to a specific deployed contract.
func NewJobManager(address common.Address, backend bind.ContractBackend) (*JobManager, error) {
	contract, err := bindJobManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JobManager{JobManagerCaller: JobManagerCaller{contract: contract}, JobManagerTransactor: JobManagerTransactor{contract: contract}, JobManagerFilterer: JobManagerFilterer{contract: contract}}, nil
}

// NewJobManagerCaller creates a new read-only instance of JobManager, bound to a specific deployed contract.
func NewJobManagerCaller(address common.Address, caller bind.ContractCaller) (*JobManagerCaller, error) {
	contract, err := bindJobManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobManagerCaller{contract: contract}, nil
}

// NewJobManagerTransactor creates a new write-only instance of JobManager, bound to a specific deployed contract.
func NewJobManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*JobManagerTransactor, error) {
	contract, err := bindJobManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobManagerTransactor{contract: contract}, nil
}

// NewJobManagerFilterer creates a new log filterer instance of JobManager, bound to a specific deployed contract.
func NewJobManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*JobManagerFilterer, error) {
	contract, err := bindJobManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobManagerFilterer{contract: contract}, nil
}

// bindJobManager binds a generic wrapper to an already deployed contract.
func bindJobManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobManager *JobManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobManager.Contract.JobManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobManager *JobManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobManager.Contract.JobManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobManager *JobManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobManager.Contract.JobManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobManager *JobManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobManager *JobManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobManager *JobManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_JobManager *JobManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_JobManager *JobManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _JobManager.Contract.DEFAULTADMINROLE(&_JobManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_JobManager *JobManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _JobManager.Contract.DEFAULTADMINROLE(&_JobManager.CallOpts)
}

// CustomerJobs is a free data retrieval call binding the contract method 0xb3ba274e.
//
// Solidity: function customerJobs(address , uint256 ) view returns(bytes32)
func (_JobManager *JobManagerCaller) CustomerJobs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "customerJobs", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CustomerJobs is a free data retrieval call binding the contract method 0xb3ba274e.
//
// Solidity: function customerJobs(address , uint256 ) view returns(bytes32)
func (_JobManager *JobManagerSession) CustomerJobs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _JobManager.Contract.CustomerJobs(&_JobManager.CallOpts, arg0, arg1)
}

// CustomerJobs is a free data retrieval call binding the contract method 0xb3ba274e.
//
// Solidity: function customerJobs(address , uint256 ) view returns(bytes32)
func (_JobManager *JobManagerCallerSession) CustomerJobs(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _JobManager.Contract.CustomerJobs(&_JobManager.CallOpts, arg0, arg1)
}

// GetJobFromId is a free data retrieval call binding the contract method 0xb3130fba.
//
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,bool,(uint64,uint64,uint64,uint64,uint64,string),bool,uint256))
func (_JobManager *JobManagerCaller) GetJobFromId(opts *bind.CallOpts, _jobId [32]byte) (Job, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "getJobFromId", _jobId)

	if err != nil {
		return *new(Job), err
	}

	out0 := *abi.ConvertType(out[0], new(Job)).(*Job)

	return out0, err

}

// GetJobFromId is a free data retrieval call binding the contract method 0xb3130fba.
//
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,bool,(uint64,uint64,uint64,uint64,uint64,string),bool,uint256))
func (_JobManager *JobManagerSession) GetJobFromId(_jobId [32]byte) (Job, error) {
	return _JobManager.Contract.GetJobFromId(&_JobManager.CallOpts, _jobId)
}

// GetJobFromId is a free data retrieval call binding the contract method 0xb3130fba.
//
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,bool,(uint64,uint64,uint64,uint64,uint64,string),bool,uint256))
func (_JobManager *JobManagerCallerSession) GetJobFromId(_jobId [32]byte) (Job, error) {
	return _JobManager.Contract.GetJobFromId(&_JobManager.CallOpts, _jobId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_JobManager *JobManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_JobManager *JobManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _JobManager.Contract.GetRoleAdmin(&_JobManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_JobManager *JobManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _JobManager.Contract.GetRoleAdmin(&_JobManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_JobManager *JobManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_JobManager *JobManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _JobManager.Contract.HasRole(&_JobManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_JobManager *JobManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _JobManager.Contract.HasRole(&_JobManager.CallOpts, role, account)
}

// JobId2CustomerAddr is a free data retrieval call binding the contract method 0x6f3de642.
//
// Solidity: function jobId2CustomerAddr(bytes32 ) view returns(address)
func (_JobManager *JobManagerCaller) JobId2CustomerAddr(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "jobId2CustomerAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobId2CustomerAddr is a free data retrieval call binding the contract method 0x6f3de642.
//
// Solidity: function jobId2CustomerAddr(bytes32 ) view returns(address)
func (_JobManager *JobManagerSession) JobId2CustomerAddr(arg0 [32]byte) (common.Address, error) {
	return _JobManager.Contract.JobId2CustomerAddr(&_JobManager.CallOpts, arg0)
}

// JobId2CustomerAddr is a free data retrieval call binding the contract method 0x6f3de642.
//
// Solidity: function jobId2CustomerAddr(bytes32 ) view returns(address)
func (_JobManager *JobManagerCallerSession) JobId2CustomerAddr(arg0 [32]byte) (common.Address, error) {
	return _JobManager.Contract.JobId2CustomerAddr(&_JobManager.CallOpts, arg0)
}

// JobId2ProviderAddr is a free data retrieval call binding the contract method 0x0cb85527.
//
// Solidity: function jobId2ProviderAddr(bytes32 ) view returns(address)
func (_JobManager *JobManagerCaller) JobId2ProviderAddr(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "jobId2ProviderAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobId2ProviderAddr is a free data retrieval call binding the contract method 0x0cb85527.
//
// Solidity: function jobId2ProviderAddr(bytes32 ) view returns(address)
func (_JobManager *JobManagerSession) JobId2ProviderAddr(arg0 [32]byte) (common.Address, error) {
	return _JobManager.Contract.JobId2ProviderAddr(&_JobManager.CallOpts, arg0)
}

// JobId2ProviderAddr is a free data retrieval call binding the contract method 0x0cb85527.
//
// Solidity: function jobId2ProviderAddr(bytes32 ) view returns(address)
func (_JobManager *JobManagerCallerSession) JobId2ProviderAddr(arg0 [32]byte) (common.Address, error) {
	return _JobManager.Contract.JobId2ProviderAddr(&_JobManager.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, bool schedulable, (uint64,uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange)
func (_JobManager *JobManagerCaller) Jobs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Schedulable            bool
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		JobId                  [32]byte
		Status                 uint8
		AmountLocked           *big.Int
		CustomerAddr           common.Address
		ProviderAddr           common.Address
		Schedulable            bool
		Definition             JobDefinition
		Valid                  bool
		BlockNumberStateChange *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.JobId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.AmountLocked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CustomerAddr = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ProviderAddr = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Schedulable = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Definition = *abi.ConvertType(out[6], new(JobDefinition)).(*JobDefinition)
	outstruct.Valid = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.BlockNumberStateChange = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, bool schedulable, (uint64,uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange)
func (_JobManager *JobManagerSession) Jobs(arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Schedulable            bool
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
}, error) {
	return _JobManager.Contract.Jobs(&_JobManager.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, bool schedulable, (uint64,uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange)
func (_JobManager *JobManagerCallerSession) Jobs(arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Schedulable            bool
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
}, error) {
	return _JobManager.Contract.Jobs(&_JobManager.CallOpts, arg0)
}

// MetaQueue is a free data retrieval call binding the contract method 0x46200b6b.
//
// Solidity: function metaQueue() view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCaller) MetaQueue(opts *bind.CallOpts) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "metaQueue")

	outstruct := new(struct {
		Begin *big.Int
		End   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Begin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MetaQueue is a free data retrieval call binding the contract method 0x46200b6b.
//
// Solidity: function metaQueue() view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerSession) MetaQueue() (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.MetaQueue(&_JobManager.CallOpts)
}

// MetaQueue is a free data retrieval call binding the contract method 0x46200b6b.
//
// Solidity: function metaQueue() view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCallerSession) MetaQueue() (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.MetaQueue(&_JobManager.CallOpts)
}

// ProviderJobs is a free data retrieval call binding the contract method 0x4f1a1106.
//
// Solidity: function providerJobs(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCaller) ProviderJobs(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "providerJobs", arg0)

	outstruct := new(struct {
		Begin *big.Int
		End   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Begin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProviderJobs is a free data retrieval call binding the contract method 0x4f1a1106.
//
// Solidity: function providerJobs(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerSession) ProviderJobs(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderJobs(&_JobManager.CallOpts, arg0)
}

// ProviderJobs is a free data retrieval call binding the contract method 0x4f1a1106.
//
// Solidity: function providerJobs(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCallerSession) ProviderJobs(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderJobs(&_JobManager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JobManager *JobManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JobManager *JobManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JobManager.Contract.SupportsInterface(&_JobManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_JobManager *JobManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _JobManager.Contract.SupportsInterface(&_JobManager.CallOpts, interfaceId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 jobId) returns()
func (_JobManager *JobManagerTransactor) Cancel(opts *bind.TransactOpts, jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "cancel", jobId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 jobId) returns()
func (_JobManager *JobManagerSession) Cancel(jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.Cancel(&_JobManager.TransactOpts, jobId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 jobId) returns()
func (_JobManager *JobManagerTransactorSession) Cancel(jobId [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.Cancel(&_JobManager.TransactOpts, jobId)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns((uint64,uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactor) ClaimJob(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "claimJob", _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns((uint64,uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns((uint64,uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactorSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0xc58467b0.
//
// Solidity: function claimNextJob(address _providerAddr) returns(bytes32 jobId, (uint64,uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactor) ClaimNextJob(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "claimNextJob", _providerAddr)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0xc58467b0.
//
// Solidity: function claimNextJob(address _providerAddr) returns(bytes32 jobId, (uint64,uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerSession) ClaimNextJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimNextJob(&_JobManager.TransactOpts, _providerAddr)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0xc58467b0.
//
// Solidity: function claimNextJob(address _providerAddr) returns(bytes32 jobId, (uint64,uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactorSession) ClaimNextJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimNextJob(&_JobManager.TransactOpts, _providerAddr)
}

// FinishJob is a paid mutator transaction binding the contract method 0x7d17544b.
//
// Solidity: function finishJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactor) FinishJob(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "finishJob", _jobId, _providerAddr)
}

// FinishJob is a paid mutator transaction binding the contract method 0x7d17544b.
//
// Solidity: function finishJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerSession) FinishJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.FinishJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// FinishJob is a paid mutator transaction binding the contract method 0x7d17544b.
//
// Solidity: function finishJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactorSession) FinishJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.FinishJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.GrantRole(&_JobManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.GrantRole(&_JobManager.TransactOpts, role, account)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactor) MetaSchedule(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "metaSchedule", _jobId, _providerAddr)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerSession) MetaSchedule(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.MetaSchedule(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactorSession) MetaSchedule(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.MetaSchedule(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x4609ca50.
//
// Solidity: function refuseJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactor) RefuseJob(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "refuseJob", _jobId, _providerAddr)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x4609ca50.
//
// Solidity: function refuseJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerSession) RefuseJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.RefuseJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x4609ca50.
//
// Solidity: function refuseJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactorSession) RefuseJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.RefuseJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.RenounceRole(&_JobManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.RenounceRole(&_JobManager.TransactOpts, role, account)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x010a7194.
//
// Solidity: function requestNewJob(address _customerAddr, (uint64,uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked) returns(bytes32)
func (_JobManager *JobManagerTransactor) RequestNewJob(opts *bind.TransactOpts, _customerAddr common.Address, _definition JobDefinition, _amountLocked *big.Int) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "requestNewJob", _customerAddr, _definition, _amountLocked)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x010a7194.
//
// Solidity: function requestNewJob(address _customerAddr, (uint64,uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked) returns(bytes32)
func (_JobManager *JobManagerSession) RequestNewJob(_customerAddr common.Address, _definition JobDefinition, _amountLocked *big.Int) (*types.Transaction, error) {
	return _JobManager.Contract.RequestNewJob(&_JobManager.TransactOpts, _customerAddr, _definition, _amountLocked)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x010a7194.
//
// Solidity: function requestNewJob(address _customerAddr, (uint64,uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked) returns(bytes32)
func (_JobManager *JobManagerTransactorSession) RequestNewJob(_customerAddr common.Address, _definition JobDefinition, _amountLocked *big.Int) (*types.Transaction, error) {
	return _JobManager.Contract.RequestNewJob(&_JobManager.TransactOpts, _customerAddr, _definition, _amountLocked)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.RevokeRole(&_JobManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_JobManager *JobManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.RevokeRole(&_JobManager.TransactOpts, role, account)
}

// StartJob is a paid mutator transaction binding the contract method 0x5792edd4.
//
// Solidity: function startJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactor) StartJob(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "startJob", _jobId, _providerAddr)
}

// StartJob is a paid mutator transaction binding the contract method 0x5792edd4.
//
// Solidity: function startJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerSession) StartJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.StartJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// StartJob is a paid mutator transaction binding the contract method 0x5792edd4.
//
// Solidity: function startJob(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactorSession) StartJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.StartJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// TopUp is a paid mutator transaction binding the contract method 0x0764e1cd.
//
// Solidity: function topUp(bytes32 _jobId, uint64 _amount) returns()
func (_JobManager *JobManagerTransactor) TopUp(opts *bind.TransactOpts, _jobId [32]byte, _amount uint64) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "topUp", _jobId, _amount)
}

// TopUp is a paid mutator transaction binding the contract method 0x0764e1cd.
//
// Solidity: function topUp(bytes32 _jobId, uint64 _amount) returns()
func (_JobManager *JobManagerSession) TopUp(_jobId [32]byte, _amount uint64) (*types.Transaction, error) {
	return _JobManager.Contract.TopUp(&_JobManager.TransactOpts, _jobId, _amount)
}

// TopUp is a paid mutator transaction binding the contract method 0x0764e1cd.
//
// Solidity: function topUp(bytes32 _jobId, uint64 _amount) returns()
func (_JobManager *JobManagerTransactorSession) TopUp(_jobId [32]byte, _amount uint64) (*types.Transaction, error) {
	return _JobManager.Contract.TopUp(&_JobManager.TransactOpts, _jobId, _amount)
}

// TriggerFailed is a paid mutator transaction binding the contract method 0xe3401e00.
//
// Solidity: function triggerFailed(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactor) TriggerFailed(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "triggerFailed", _jobId, _providerAddr)
}

// TriggerFailed is a paid mutator transaction binding the contract method 0xe3401e00.
//
// Solidity: function triggerFailed(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerSession) TriggerFailed(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.TriggerFailed(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// TriggerFailed is a paid mutator transaction binding the contract method 0xe3401e00.
//
// Solidity: function triggerFailed(bytes32 _jobId, address _providerAddr) returns()
func (_JobManager *JobManagerTransactorSession) TriggerFailed(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.TriggerFailed(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// UpdateJobsStatus is a paid mutator transaction binding the contract method 0x2a242a76.
//
// Solidity: function updateJobsStatus() returns()
func (_JobManager *JobManagerTransactor) UpdateJobsStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "updateJobsStatus")
}

// UpdateJobsStatus is a paid mutator transaction binding the contract method 0x2a242a76.
//
// Solidity: function updateJobsStatus() returns()
func (_JobManager *JobManagerSession) UpdateJobsStatus() (*types.Transaction, error) {
	return _JobManager.Contract.UpdateJobsStatus(&_JobManager.TransactOpts)
}

// UpdateJobsStatus is a paid mutator transaction binding the contract method 0x2a242a76.
//
// Solidity: function updateJobsStatus() returns()
func (_JobManager *JobManagerTransactorSession) UpdateJobsStatus() (*types.Transaction, error) {
	return _JobManager.Contract.UpdateJobsStatus(&_JobManager.TransactOpts)
}

// UpdateRoles is a paid mutator transaction binding the contract method 0x3a80760a.
//
// Solidity: function updateRoles(address _metaSchedulerAddr) returns()
func (_JobManager *JobManagerTransactor) UpdateRoles(opts *bind.TransactOpts, _metaSchedulerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "updateRoles", _metaSchedulerAddr)
}

// UpdateRoles is a paid mutator transaction binding the contract method 0x3a80760a.
//
// Solidity: function updateRoles(address _metaSchedulerAddr) returns()
func (_JobManager *JobManagerSession) UpdateRoles(_metaSchedulerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.UpdateRoles(&_JobManager.TransactOpts, _metaSchedulerAddr)
}

// UpdateRoles is a paid mutator transaction binding the contract method 0x3a80760a.
//
// Solidity: function updateRoles(address _metaSchedulerAddr) returns()
func (_JobManager *JobManagerTransactorSession) UpdateRoles(_metaSchedulerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.UpdateRoles(&_JobManager.TransactOpts, _metaSchedulerAddr)
}

// JobManagerJobRefusedEventIterator is returned from FilterJobRefusedEvent and is used to iterate over the raw logs and unpacked data for JobRefusedEvent events raised by the JobManager contract.
type JobManagerJobRefusedEventIterator struct {
	Event *JobManagerJobRefusedEvent // Event containing the contract specifics and raw log

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
func (it *JobManagerJobRefusedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerJobRefusedEvent)
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
		it.Event = new(JobManagerJobRefusedEvent)
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
func (it *JobManagerJobRefusedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerJobRefusedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerJobRefusedEvent represents a JobRefusedEvent event raised by the JobManager contract.
type JobManagerJobRefusedEvent struct {
	JobId        [32]byte
	ProviderAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterJobRefusedEvent is a free log retrieval operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
func (_JobManager *JobManagerFilterer) FilterJobRefusedEvent(opts *bind.FilterOpts) (*JobManagerJobRefusedEventIterator, error) {

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "JobRefusedEvent")
	if err != nil {
		return nil, err
	}
	return &JobManagerJobRefusedEventIterator{contract: _JobManager.contract, event: "JobRefusedEvent", logs: logs, sub: sub}, nil
}

// WatchJobRefusedEvent is a free log subscription operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
func (_JobManager *JobManagerFilterer) WatchJobRefusedEvent(opts *bind.WatchOpts, sink chan<- *JobManagerJobRefusedEvent) (event.Subscription, error) {

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "JobRefusedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerJobRefusedEvent)
				if err := _JobManager.contract.UnpackLog(event, "JobRefusedEvent", log); err != nil {
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

// ParseJobRefusedEvent is a log parse operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
func (_JobManager *JobManagerFilterer) ParseJobRefusedEvent(log types.Log) (*JobManagerJobRefusedEvent, error) {
	event := new(JobManagerJobRefusedEvent)
	if err := _JobManager.contract.UnpackLog(event, "JobRefusedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerNewJobRequestEventIterator is returned from FilterNewJobRequestEvent and is used to iterate over the raw logs and unpacked data for NewJobRequestEvent events raised by the JobManager contract.
type JobManagerNewJobRequestEventIterator struct {
	Event *JobManagerNewJobRequestEvent // Event containing the contract specifics and raw log

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
func (it *JobManagerNewJobRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerNewJobRequestEvent)
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
		it.Event = new(JobManagerNewJobRequestEvent)
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
func (it *JobManagerNewJobRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerNewJobRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerNewJobRequestEvent represents a NewJobRequestEvent event raised by the JobManager contract.
type JobManagerNewJobRequestEvent struct {
	JobId        [32]byte
	CustomerAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewJobRequestEvent is a free log retrieval operation binding the contract event 0x1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3.
//
// Solidity: event NewJobRequestEvent(bytes32 _jobId, address _customerAddr)
func (_JobManager *JobManagerFilterer) FilterNewJobRequestEvent(opts *bind.FilterOpts) (*JobManagerNewJobRequestEventIterator, error) {

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "NewJobRequestEvent")
	if err != nil {
		return nil, err
	}
	return &JobManagerNewJobRequestEventIterator{contract: _JobManager.contract, event: "NewJobRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewJobRequestEvent is a free log subscription operation binding the contract event 0x1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3.
//
// Solidity: event NewJobRequestEvent(bytes32 _jobId, address _customerAddr)
func (_JobManager *JobManagerFilterer) WatchNewJobRequestEvent(opts *bind.WatchOpts, sink chan<- *JobManagerNewJobRequestEvent) (event.Subscription, error) {

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "NewJobRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerNewJobRequestEvent)
				if err := _JobManager.contract.UnpackLog(event, "NewJobRequestEvent", log); err != nil {
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
func (_JobManager *JobManagerFilterer) ParseNewJobRequestEvent(log types.Log) (*JobManagerNewJobRequestEvent, error) {
	event := new(JobManagerNewJobRequestEvent)
	if err := _JobManager.contract.UnpackLog(event, "NewJobRequestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the JobManager contract.
type JobManagerRoleAdminChangedIterator struct {
	Event *JobManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *JobManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerRoleAdminChanged)
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
		it.Event = new(JobManagerRoleAdminChanged)
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
func (it *JobManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerRoleAdminChanged represents a RoleAdminChanged event raised by the JobManager contract.
type JobManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_JobManager *JobManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*JobManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &JobManagerRoleAdminChangedIterator{contract: _JobManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_JobManager *JobManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *JobManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerRoleAdminChanged)
				if err := _JobManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_JobManager *JobManagerFilterer) ParseRoleAdminChanged(log types.Log) (*JobManagerRoleAdminChanged, error) {
	event := new(JobManagerRoleAdminChanged)
	if err := _JobManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the JobManager contract.
type JobManagerRoleGrantedIterator struct {
	Event *JobManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *JobManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerRoleGranted)
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
		it.Event = new(JobManagerRoleGranted)
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
func (it *JobManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerRoleGranted represents a RoleGranted event raised by the JobManager contract.
type JobManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_JobManager *JobManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*JobManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &JobManagerRoleGrantedIterator{contract: _JobManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_JobManager *JobManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *JobManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerRoleGranted)
				if err := _JobManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_JobManager *JobManagerFilterer) ParseRoleGranted(log types.Log) (*JobManagerRoleGranted, error) {
	event := new(JobManagerRoleGranted)
	if err := _JobManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the JobManager contract.
type JobManagerRoleRevokedIterator struct {
	Event *JobManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *JobManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerRoleRevoked)
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
		it.Event = new(JobManagerRoleRevoked)
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
func (it *JobManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerRoleRevoked represents a RoleRevoked event raised by the JobManager contract.
type JobManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_JobManager *JobManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*JobManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &JobManagerRoleRevokedIterator{contract: _JobManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_JobManager *JobManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *JobManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerRoleRevoked)
				if err := _JobManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_JobManager *JobManagerFilterer) ParseRoleRevoked(log types.Log) (*JobManagerRoleRevoked, error) {
	event := new(JobManagerRoleRevoked)
	if err := _JobManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerMetaData contains all meta data concerning the MetaScheduler contract.
var MetaSchedulerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_credit\",\"type\":\"address\"},{\"internalType\":\"contractJobManager\",\"name\":\"_jobManager\",\"type\":\"address\"},{\"internalType\":\"contractProviderManager\",\"name\":\"_providerManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDurationMinute\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimNextJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_METASCHEDULER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"jobDuration\",\"type\":\"uint256\"}],\"name\":\"finishJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobStatus\",\"outputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleMetaScheduler\",\"type\":\"address\"}],\"name\":\"grantRoleOracleMetaScheduler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleLiveness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_memPricePerMin\",\"type\":\"uint64\"}],\"name\":\"providerApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerRedemption\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nNodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_mem\",\"type\":\"uint64\"}],\"name\":\"providerRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"startJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"triggerFailedJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620042f6380380620042f683398181016040528101906200003791906200045a565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603620000a9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a0906200053d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200011b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200011290620005d5565b60405180910390fd5b620001306000801b33620001ef60201b60201c565b8273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620005f7565b620002018282620002e060201b60201c565b620002dc57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620002816200034a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003848262000357565b9050919050565b6000620003988262000377565b9050919050565b620003aa816200038b565b8114620003b657600080fd5b50565b600081519050620003ca816200039f565b92915050565b6000620003dd8262000377565b9050919050565b620003ef81620003d0565b8114620003fb57600080fd5b50565b6000815190506200040f81620003e4565b92915050565b6000620004228262000377565b9050919050565b620004348162000415565b81146200044057600080fd5b50565b600081519050620004548162000429565b92915050565b60008060006060848603121562000476576200047562000352565b5b60006200048686828701620003b9565b93505060206200049986828701620003fe565b9250506040620004ac8682870162000443565b9150509250925092565b600082825260208201905092915050565b7f4d6574615363686564756c65723a206372656469742061646472206973207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b600062000525602283620004b6565b91506200053282620004c7565b604082019050919050565b60006020820190508181036000830152620005588162000516565b9050919050565b7f4d6574615363686564756c65723a2070726f76696465724d616e61676572206160008201527f646472206973207a65726f000000000000000000000000000000000000000000602082015250565b6000620005bd602b83620004b6565b9150620005ca826200055f565b604082019050919050565b60006020820190508181036000830152620005f081620005ae565b9050919050565b608051613cc0620006366000396000818161087c0152818161099e01528181611520015281816115fa0152818161183101526119e60152613cc06000f3fe6080604052600436106101405760003560e01c8063750f0acc116100b6578063a217fddf1161006f578063a217fddf1461046b578063ba9c7f1814610496578063d1c48e92146104ad578063d1cee546146104d6578063d547741f146104ff578063fbc3611a1461052857610140565b8063750f0acc1461034957806383ff0b281461037257806389be60441461039d5780638e855488146103c657806391d1485414610403578063a06d083c1461044057610140565b8063236e26ae11610108578063236e26ae1461023a578063248a9ca3146102635780632f2ff15d146102a05780632fcca6fe146102c957806336568abe146102f25780635d3a71801461031b57610140565b806301ffc9a7146101455780630dca0ac414610182578063124b09f3146101ab5780631f92a63f146101d45780632081f4c8146101fd575b600080fd5b34801561015157600080fd5b5061016c600480360381019061016791906124c0565b610532565b6040516101799190612508565b60405180910390f35b34801561018e57600080fd5b506101a960048036038101906101a4919061258f565b6105ac565b005b3480156101b757600080fd5b506101d260048036038101906101cd919061262d565b610b5a565b005b3480156101e057600080fd5b506101fb60048036038101906101f6919061265a565b610b94565b005b34801561020957600080fd5b50610224600480360381019061021f919061265a565b610d01565b60405161023191906126fe565b60405180910390f35b34801561024657600080fd5b50610261600480360381019061025c919061265a565b610daf565b005b34801561026f57600080fd5b5061028a6004803603810190610285919061265a565b610f1c565b6040516102979190612728565b60405180910390f35b3480156102ac57600080fd5b506102c760048036038101906102c29190612743565b610f3b565b005b3480156102d557600080fd5b506102f060048036038101906102eb91906127c3565b610f5c565b005b3480156102fe57600080fd5b5061031960048036038101906103149190612743565b61101f565b005b34801561032757600080fd5b506103306110a2565b6040516103409493929190612970565b60405180910390f35b34801561035557600080fd5b50610370600480360381019061036b91906129bc565b611506565b005b34801561037e57600080fd5b5061038761169c565b6040516103949190612728565b60405180910390f35b3480156103a957600080fd5b506103c460048036038101906103bf919061265a565b6116c0565b005b3480156103d257600080fd5b506103ed60048036038101906103e89190612be4565b61182d565b6040516103fa9190612728565b60405180910390f35b34801561040f57600080fd5b5061042a60048036038101906104259190612743565b61197a565b6040516104379190612508565b60405180910390f35b34801561044c57600080fd5b506104556119e4565b6040516104629190612c9f565b60405180910390f35b34801561047757600080fd5b50610480611a08565b60405161048d9190612728565b60405180910390f35b3480156104a257600080fd5b506104ab611a0f565b005b3480156104b957600080fd5b506104d460048036038101906104cf9190612cba565b611abd565b005b3480156104e257600080fd5b506104fd60048036038101906104f89190612743565b611b83565b005b34801561050b57600080fd5b5061052660048036038101906105219190612743565b611c40565b005b610530611c61565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806105a557506105a482611ecc565b5b9050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b81526004016106079190612d35565b602060405180830381865afa158015610624573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106489190612d7c565b610687576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067e90612e2c565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba846040518263ffffffff1660e01b81526004016106e49190612728565b600060405180830381865afa158015610701573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061072a91906130ed565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663624bc8e3336040518263ffffffff1660e01b81526004016107899190612d35565b61014060405180830381865afa1580156107a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cb919061328a565b9050600073__$d9cf0c7370a9f90abc1cfd27bd34fdbb07$__637316b6ef8460c00151856040015185600001516040518463ffffffff1660e01b815260040161081693929190613437565b602060405180830381865af4158015610833573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108579190613476565b905060008490508482101561086a578190505b6000818361087891906134d2565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb3373__$d9cf0c7370a9f90abc1cfd27bd34fdbb07$__63214ef2828960c001518960000151886040518463ffffffff1660e01b81526004016108fa93929190613506565b602060405180830381865af4158015610917573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093b9190613476565b6040518363ffffffff1660e01b8152600401610958929190613546565b6020604051808303816000875af1158015610977573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099b9190612d7c565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb866060015173__$d9cf0c7370a9f90abc1cfd27bd34fdbb07$__63214ef2828960c001518960000151876040518463ffffffff1660e01b8152600401610a2093929190613506565b602060405180830381865af4158015610a3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a619190613476565b6040518363ffffffff1660e01b8152600401610a7e929190613546565b6020604051808303816000875af1158015610a9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac19190612d7c565b50600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637d17544b88336040518363ffffffff1660e01b8152600401610b1f92919061356f565b600060405180830381600087803b158015610b3957600080fd5b505af1158015610b4d573d6000803e3d6000fd5b5050505050505050505050565b610b676000801b33611f36565b610b917f54c54853c56f9bc984c7720e9a3cf89c2d7e7c065a87e64cbe4fc77aecc423bd82611fd3565b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401610bef9190612d35565b602060405180830381865afa158015610c0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c309190612d7c565b610c6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6690612e2c565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634609ca5082336040518363ffffffff1660e01b8152600401610ccc92919061356f565b600060405180830381600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b5050505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b8152600401610d5e9190612728565b600060405180830381865afa158015610d7b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610da491906130ed565b602001519050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401610e0a9190612d35565b602060405180830381865afa158015610e27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4b9190612d7c565b610e8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8190612e2c565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635792edd482336040518363ffffffff1660e01b8152600401610ee792919061356f565b600060405180830381600087803b158015610f0157600080fd5b505af1158015610f15573d6000803e3d6000fd5b5050505050565b6000806000838152602001908152602001600020600101549050919050565b610f4482610f1c565b610f4d816120b3565b610f578383611fd3565b505050565b610f867f54c54853c56f9bc984c7720e9a3cf89c2d7e7c065a87e64cbe4fc77aecc423bd33611f36565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632af37fcd858585856040518563ffffffff1660e01b8152600401610fe794939291906135a7565b600060405180830381600087803b15801561100157600080fd5b505af1158015611015573d6000803e3d6000fd5b5050505050505050565b6110276120c7565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611094576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108b9061365e565b60405180910390fd5b61109e82826120cf565b5050565b60008060006110af6123ec565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b815260040161110a9190612d35565b602060405180830381865afa158015611127573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061114b9190612d7c565b61118a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118190612e2c565b60405180910390fd5b60006111946123ec565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c58467b0336040518263ffffffff1660e01b81526004016111ef9190612d35565b6000604051808303816000875af115801561120e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611237919061367e565b80925081935050506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba846040518263ffffffff1660e01b815260040161129c9190612728565b600060405180830381865afa1580156112b9573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906112e291906130ed565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f6edb5f82608001516040518263ffffffff1660e01b81526004016113439190612d35565b600060405180830381600087803b15801561135d57600080fd5b505af1158015611371573d6000803e3d6000fd5b505050506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663624bc8e3336040518263ffffffff1660e01b81526004016113d29190612d35565b61014060405180830381865afa1580156113f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611414919061328a565b9050600073__$d9cf0c7370a9f90abc1cfd27bd34fdbb07$__637316b6ef8460c00151856040015185600001516040518463ffffffff1660e01b815260040161145f93929190613437565b602060405180830381865af415801561147c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a09190613476565b90507fbc999b859676a74ea766936a74f9ade3d1f245417fd1afc6a1e5ca29af316335836060015186838660c001516040516114df9493929190612970565b60405180910390a1826060015185828560c001519850985098509850505050505090919293565b6115136000801b33611f36565b8067ffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016115779190612d35565b602060405180830381865afa158015611594573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b89190613476565b116115f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115ef9061374c565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3033846040518463ffffffff1660e01b81526004016116559392919061379d565b6020604051808303816000875af1158015611674573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116989190612d7c565b5050565b7f54c54853c56f9bc984c7720e9a3cf89c2d7e7c065a87e64cbe4fc77aecc423bd81565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b815260040161171b9190612d35565b602060405180830381865afa158015611738573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175c9190612d7c565b61179b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179290612e2c565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3401e0082336040518363ffffffff1660e01b81526004016117f892919061356f565b600060405180830381600087803b15801561181257600080fd5b505af1158015611826573d6000803e3d6000fd5b5050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161188c939291906137d4565b6020604051808303816000875af11580156118ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118cf9190612d7c565b50600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663010a71943385856040518463ffffffff1660e01b815260040161192f9392919061380b565b6020604051808303816000875af115801561194e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119729190613849565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000801b81565b611a397f54c54853c56f9bc984c7720e9a3cf89c2d7e7c065a87e64cbe4fc77aecc423bd33611f36565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a242a766040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611aa357600080fd5b505af1158015611ab7573d6000803e3d6000fd5b50505050565b611ae77f54c54853c56f9bc984c7720e9a3cf89c2d7e7c065a87e64cbe4fc77aecc423bd33611f36565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634ee6f90786868686866040518663ffffffff1660e01b8152600401611b4a959493929190613876565b600060405180830381600087803b158015611b6457600080fd5b505af1158015611b78573d6000803e3d6000fd5b505050505050505050565b611bad7f54c54853c56f9bc984c7720e9a3cf89c2d7e7c065a87e64cbe4fc77aecc423bd33611f36565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d1cee54683836040518363ffffffff1660e01b8152600401611c0a92919061356f565b600060405180830381600087803b158015611c2457600080fd5b505af1158015611c38573d6000803e3d6000fd5b505050505050565b611c4982610f1c565b611c52816120b3565b611c5c83836120cf565b505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b7bb9145336040518263ffffffff1660e01b8152600401611cbc9190612d35565b602060405180830381865afa158015611cd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cfd9190612d7c565b611d3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d3390613915565b60405180910390fd5b670de0b6b3a7640000341015611d87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d7e90613915565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff16683635c9adc5dea00000604051611db690613966565b60006040518083038185875af1925050503d8060008114611df3576040519150601f19603f3d011682016040523d82523d6000602084013e611df8565b606091505b5050905080611e3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e33906139c7565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7692d09336040518263ffffffff1660e01b8152600401611e979190612d35565b600060405180830381600087803b158015611eb157600080fd5b505af1158015611ec5573d6000803e3d6000fd5b5050505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b611f40828261197a565b611fcf57611f658173ffffffffffffffffffffffffffffffffffffffff1660146121b0565b611f738360001c60206121b0565b604051602001611f84929190613abb565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fc69190613b2e565b60405180910390fd5b5050565b611fdd828261197a565b6120af57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506120546120c7565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6120c4816120bf6120c7565b611f36565b50565b600033905090565b6120d9828261197a565b156121ac57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506121516120c7565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6060600060028360026121c39190613b50565b6121cd9190613b92565b67ffffffffffffffff8111156121e6576121e56129ee565b5b6040519080825280601f01601f1916602001820160405280156122185781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106122505761224f613bc6565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106122b4576122b3613bc6565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026122f49190613b50565b6122fe9190613b92565b90505b600181111561239e577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106123405761233f613bc6565b5b1a60f81b82828151811061235757612356613bc6565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061239790613bf5565b9050612301565b50600084146123e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123d990613c6a565b60405180910390fd5b8091505092915050565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61249d81612468565b81146124a857600080fd5b50565b6000813590506124ba81612494565b92915050565b6000602082840312156124d6576124d561245e565b5b60006124e4848285016124ab565b91505092915050565b60008115159050919050565b612502816124ed565b82525050565b600060208201905061251d60008301846124f9565b92915050565b6000819050919050565b61253681612523565b811461254157600080fd5b50565b6000813590506125538161252d565b92915050565b6000819050919050565b61256c81612559565b811461257757600080fd5b50565b60008135905061258981612563565b92915050565b600080604083850312156125a6576125a561245e565b5b60006125b485828601612544565b92505060206125c58582860161257a565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006125fa826125cf565b9050919050565b61260a816125ef565b811461261557600080fd5b50565b60008135905061262781612601565b92915050565b6000602082840312156126435761264261245e565b5b600061265184828501612618565b91505092915050565b6000602082840312156126705761266f61245e565b5b600061267e84828501612544565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600781106126c7576126c6612687565b5b50565b60008190506126d8826126b6565b919050565b60006126e8826126ca565b9050919050565b6126f8816126dd565b82525050565b600060208201905061271360008301846126ef565b92915050565b61272281612523565b82525050565b600060208201905061273d6000830184612719565b92915050565b6000806040838503121561275a5761275961245e565b5b600061276885828601612544565b925050602061277985828601612618565b9150509250929050565b600067ffffffffffffffff82169050919050565b6127a081612783565b81146127ab57600080fd5b50565b6000813590506127bd81612797565b92915050565b600080600080608085870312156127dd576127dc61245e565b5b60006127eb87828801612618565b94505060206127fc878288016127ae565b935050604061280d878288016127ae565b925050606061281e878288016127ae565b91505092959194509250565b612833816125ef565b82525050565b61284281612559565b82525050565b61285181612783565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015612891578082015181840152602081019050612876565b60008484015250505050565b6000601f19601f8301169050919050565b60006128b982612857565b6128c38185612862565b93506128d3818560208601612873565b6128dc8161289d565b840191505092915050565b600060c0830160008301516128ff6000860182612848565b5060208301516129126020860182612848565b5060408301516129256040860182612848565b5060608301516129386060860182612848565b50608083015161294b6080860182612848565b5060a083015184820360a086015261296382826128ae565b9150508091505092915050565b6000608082019050612985600083018761282a565b6129926020830186612719565b61299f6040830185612839565b81810360608301526129b181846128e7565b905095945050505050565b6000602082840312156129d2576129d161245e565b5b60006129e0848285016127ae565b91505092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b612a268261289d565b810181811067ffffffffffffffff82111715612a4557612a446129ee565b5b80604052505050565b6000612a58612454565b9050612a648282612a1d565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115612a9357612a926129ee565b5b612a9c8261289d565b9050602081019050919050565b82818337600083830152505050565b6000612acb612ac684612a78565b612a4e565b905082815260208101848484011115612ae757612ae6612a73565b5b612af2848285612aa9565b509392505050565b600082601f830112612b0f57612b0e612a6e565b5b8135612b1f848260208601612ab8565b91505092915050565b600060c08284031215612b3e57612b3d6129e9565b5b612b4860c0612a4e565b90506000612b58848285016127ae565b6000830152506020612b6c848285016127ae565b6020830152506040612b80848285016127ae565b6040830152506060612b94848285016127ae565b6060830152506080612ba8848285016127ae565b60808301525060a082013567ffffffffffffffff811115612bcc57612bcb612a69565b5b612bd884828501612afa565b60a08301525092915050565b60008060408385031215612bfb57612bfa61245e565b5b600083013567ffffffffffffffff811115612c1957612c18612463565b5b612c2585828601612b28565b9250506020612c368582860161257a565b9150509250929050565b6000819050919050565b6000612c65612c60612c5b846125cf565b612c40565b6125cf565b9050919050565b6000612c7782612c4a565b9050919050565b6000612c8982612c6c565b9050919050565b612c9981612c7e565b82525050565b6000602082019050612cb46000830184612c90565b92915050565b600080600080600060a08688031215612cd657612cd561245e565b5b6000612ce488828901612618565b9550506020612cf5888289016127ae565b9450506040612d06888289016127ae565b9350506060612d17888289016127ae565b9250506080612d28888289016127ae565b9150509295509295909350565b6000602082019050612d4a600083018461282a565b92915050565b612d59816124ed565b8114612d6457600080fd5b50565b600081519050612d7681612d50565b92915050565b600060208284031215612d9257612d9161245e565b5b6000612da084828501612d67565b91505092915050565b600082825260208201905092915050565b7f50726f766964657273206e6f7420666f756e64206f72206e6f74206a6f696e6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b6000612e16602183612da9565b9150612e2182612dba565b604082019050919050565b60006020820190508181036000830152612e4581612e09565b9050919050565b600081519050612e5b8161252d565b92915050565b60078110612e6e57600080fd5b50565b600081519050612e8081612e61565b92915050565b600081519050612e9581612563565b92915050565b600081519050612eaa81612601565b92915050565b600081519050612ebf81612797565b92915050565b6000612ed8612ed384612a78565b612a4e565b905082815260208101848484011115612ef457612ef3612a73565b5b612eff848285612873565b509392505050565b600082601f830112612f1c57612f1b612a6e565b5b8151612f2c848260208601612ec5565b91505092915050565b600060c08284031215612f4b57612f4a6129e9565b5b612f5560c0612a4e565b90506000612f6584828501612eb0565b6000830152506020612f7984828501612eb0565b6020830152506040612f8d84828501612eb0565b6040830152506060612fa184828501612eb0565b6060830152506080612fb584828501612eb0565b60808301525060a082015167ffffffffffffffff811115612fd957612fd8612a69565b5b612fe584828501612f07565b60a08301525092915050565b60006101208284031215613008576130076129e9565b5b613013610120612a4e565b9050600061302384828501612e4c565b600083015250602061303784828501612e71565b602083015250604061304b84828501612e86565b604083015250606061305f84828501612e9b565b606083015250608061307384828501612e9b565b60808301525060a061308784828501612d67565b60a08301525060c082015167ffffffffffffffff8111156130ab576130aa612a69565b5b6130b784828501612f35565b60c08301525060e06130cb84828501612d67565b60e0830152506101006130e084828501612e86565b6101008301525092915050565b6000602082840312156131035761310261245e565b5b600082015167ffffffffffffffff81111561312157613120612463565b5b61312d84828501612ff1565b91505092915050565b600060e0828403121561314c5761314b6129e9565b5b61315660e0612a4e565b9050600061316684828501612eb0565b600083015250602061317a84828501612eb0565b602083015250604061318e84828501612eb0565b60408301525060606131a284828501612eb0565b60608301525060806131b684828501612eb0565b60808301525060a06131ca84828501612eb0565b60a08301525060c06131de84828501612eb0565b60c08301525092915050565b600381106131f757600080fd5b50565b600081519050613209816131ea565b92915050565b60006101408284031215613226576132256129e9565b5b6132306080612a4e565b9050600061324084828501613136565b60008301525060e0613254848285016131fa565b60208301525061010061326984828501612d67565b60408301525061012061327e84828501612eb0565b60608301525092915050565b600061014082840312156132a1576132a061245e565b5b60006132af8482850161320f565b91505092915050565b6132c181612783565b82525050565b600082825260208201905092915050565b60006132e382612857565b6132ed81856132c7565b93506132fd818560208601612873565b6133068161289d565b840191505092915050565b600060c08301600083015161332960008601826132b8565b50602083015161333c60208601826132b8565b50604083015161334f60408601826132b8565b50606083015161336260608601826132b8565b50608083015161337560808601826132b8565b5060a083015184820360a086015261338d82826132d8565b9150508091505092915050565b6133a381612559565b82525050565b60e0820160008201516133bf60008501826132b8565b5060208201516133d260208501826132b8565b5060408201516133e560408501826132b8565b5060608201516133f860608501826132b8565b50608082015161340b60808501826132b8565b5060a082015161341e60a08501826132b8565b5060c082015161343160c08501826132b8565b50505050565b60006101208201905081810360008301526134528186613311565b9050613461602083018561339a565b61346e60408301846133a9565b949350505050565b60006020828403121561348c5761348b61245e565b5b600061349a84828501612e86565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006134dd82612559565b91506134e883612559565b9250828203905081811115613500576134ff6134a3565b5b92915050565b60006101208201905081810360008301526135218186613311565b905061353060208301856133a9565b61353e61010083018461339a565b949350505050565b600060408201905061355b600083018561282a565b6135686020830184612839565b9392505050565b60006040820190506135846000830185612719565b613591602083018461282a565b9392505050565b6135a181612783565b82525050565b60006080820190506135bc600083018761282a565b6135c96020830186613598565b6135d66040830185613598565b6135e36060830184613598565b95945050505050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000613648602f83612da9565b9150613653826135ec565b604082019050919050565b600060208201905081810360008301526136778161363b565b9050919050565b600080604083850312156136955761369461245e565b5b60006136a385828601612e4c565b925050602083015167ffffffffffffffff8111156136c4576136c3612463565b5b6136d085828601612f35565b9150509250929050565b7f4d6574615363686564756c65723a20776974686472617720616d6f756e74206860008201527f6967686572207468616e2062616c616e63652e00000000000000000000000000602082015250565b6000613736603383612da9565b9150613741826136da565b604082019050919050565b6000602082019050818103600083015261376581613729565b9050919050565b600061378761378261377d84612783565b612c40565b612559565b9050919050565b6137978161376c565b82525050565b60006060820190506137b2600083018661282a565b6137bf602083018561282a565b6137cc604083018461378e565b949350505050565b60006060820190506137e9600083018661282a565b6137f6602083018561282a565b6138036040830184612839565b949350505050565b6000606082019050613820600083018661282a565b818103602083015261383281856128e7565b90506138416040830184612839565b949350505050565b60006020828403121561385f5761385e61245e565b5b600061386d84828501612e4c565b91505092915050565b600060a08201905061388b600083018861282a565b6138986020830187613598565b6138a56040830186613598565b6138b26060830185613598565b6138bf6080830184613598565b9695505050505050565b7f50726f7669646572206e6f7420666f756e64206f72206e6f74206b69636b6564600082015250565b60006138ff602083612da9565b915061390a826138c9565b602082019050919050565b6000602082019050818103600083015261392e816138f2565b9050919050565b600081905092915050565b50565b6000613950600083613935565b915061395b82613940565b600082019050919050565b600061397182613943565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b60006139b1601083612da9565b91506139bc8261397b565b602082019050919050565b600060208201905081810360008301526139e0816139a4565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000613a286017836139e7565b9150613a33826139f2565b601782019050919050565b6000613a4982612857565b613a5381856139e7565b9350613a63818560208601612873565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000613aa56011836139e7565b9150613ab082613a6f565b601182019050919050565b6000613ac682613a1b565b9150613ad28285613a3e565b9150613add82613a98565b9150613ae98284613a3e565b91508190509392505050565b6000613b0082612857565b613b0a8185612da9565b9350613b1a818560208601612873565b613b238161289d565b840191505092915050565b60006020820190508181036000830152613b488184613af5565b905092915050565b6000613b5b82612559565b9150613b6683612559565b9250828202613b7481612559565b91508282048414831517613b8b57613b8a6134a3565b5b5092915050565b6000613b9d82612559565b9150613ba883612559565b9250828201905080821115613bc057613bbf6134a3565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000613c0082612559565b915060008203613c1357613c126134a3565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000613c54602083612da9565b9150613c5f82613c1e565b602082019050919050565b60006020820190508181036000830152613c8381613c47565b905091905056fea2646970667358221220fbb546a09d6a3686e03a14818eadf4e1f0adba5d665a61822bc08da3a032d2cc64736f6c63430008110033",
}

// MetaSchedulerABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaSchedulerMetaData.ABI instead.
var MetaSchedulerABI = MetaSchedulerMetaData.ABI

// MetaSchedulerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MetaSchedulerMetaData.Bin instead.
var MetaSchedulerBin = MetaSchedulerMetaData.Bin

// DeployMetaScheduler deploys a new Ethereum contract, binding an instance of MetaScheduler to it.
func DeployMetaScheduler(auth *bind.TransactOpts, backend bind.ContractBackend, _credit common.Address, _jobManager common.Address, _providerManager common.Address) (common.Address, *types.Transaction, *MetaScheduler, error) {
	parsed, err := MetaSchedulerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	timeIsMoneyAddr, _, _, _ := DeployTimeIsMoney(auth, backend)
	MetaSchedulerBin = strings.ReplaceAll(MetaSchedulerBin, "__$d9cf0c7370a9f90abc1cfd27bd34fdbb07$__", timeIsMoneyAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetaSchedulerBin), backend, _credit, _jobManager, _providerManager)
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
	parsed, err := abi.JSON(strings.NewReader(MetaSchedulerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// ORACLEMETASCHEDULER is a free data retrieval call binding the contract method 0x83ff0b28.
//
// Solidity: function ORACLE_METASCHEDULER() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCaller) ORACLEMETASCHEDULER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "ORACLE_METASCHEDULER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEMETASCHEDULER is a free data retrieval call binding the contract method 0x83ff0b28.
//
// Solidity: function ORACLE_METASCHEDULER() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) ORACLEMETASCHEDULER() ([32]byte, error) {
	return _MetaScheduler.Contract.ORACLEMETASCHEDULER(&_MetaScheduler.CallOpts)
}

// ORACLEMETASCHEDULER is a free data retrieval call binding the contract method 0x83ff0b28.
//
// Solidity: function ORACLE_METASCHEDULER() view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCallerSession) ORACLEMETASCHEDULER() ([32]byte, error) {
	return _MetaScheduler.Contract.ORACLEMETASCHEDULER(&_MetaScheduler.CallOpts)
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

// GetJobStatus is a free data retrieval call binding the contract method 0x2081f4c8.
//
// Solidity: function getJobStatus(bytes32 _jobId) view returns(uint8)
func (_MetaScheduler *MetaSchedulerCaller) GetJobStatus(opts *bind.CallOpts, _jobId [32]byte) (uint8, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "getJobStatus", _jobId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetJobStatus is a free data retrieval call binding the contract method 0x2081f4c8.
//
// Solidity: function getJobStatus(bytes32 _jobId) view returns(uint8)
func (_MetaScheduler *MetaSchedulerSession) GetJobStatus(_jobId [32]byte) (uint8, error) {
	return _MetaScheduler.Contract.GetJobStatus(&_MetaScheduler.CallOpts, _jobId)
}

// GetJobStatus is a free data retrieval call binding the contract method 0x2081f4c8.
//
// Solidity: function getJobStatus(bytes32 _jobId) view returns(uint8)
func (_MetaScheduler *MetaSchedulerCallerSession) GetJobStatus(_jobId [32]byte) (uint8, error) {
	return _MetaScheduler.Contract.GetJobStatus(&_MetaScheduler.CallOpts, _jobId)
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

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns(address, bytes32, uint256, (uint64,uint64,uint64,uint64,uint64,string))
func (_MetaScheduler *MetaSchedulerTransactor) ClaimNextJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "claimNextJob")
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns(address, bytes32, uint256, (uint64,uint64,uint64,uint64,uint64,string))
func (_MetaScheduler *MetaSchedulerSession) ClaimNextJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextJob(&_MetaScheduler.TransactOpts)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns(address, bytes32, uint256, (uint64,uint64,uint64,uint64,uint64,string))
func (_MetaScheduler *MetaSchedulerTransactorSession) ClaimNextJob() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimNextJob(&_MetaScheduler.TransactOpts)
}

// FinishJob is a paid mutator transaction binding the contract method 0x0dca0ac4.
//
// Solidity: function finishJob(bytes32 _jobId, uint256 jobDuration) returns()
func (_MetaScheduler *MetaSchedulerTransactor) FinishJob(opts *bind.TransactOpts, _jobId [32]byte, jobDuration *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "finishJob", _jobId, jobDuration)
}

// FinishJob is a paid mutator transaction binding the contract method 0x0dca0ac4.
//
// Solidity: function finishJob(bytes32 _jobId, uint256 jobDuration) returns()
func (_MetaScheduler *MetaSchedulerSession) FinishJob(_jobId [32]byte, jobDuration *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.FinishJob(&_MetaScheduler.TransactOpts, _jobId, jobDuration)
}

// FinishJob is a paid mutator transaction binding the contract method 0x0dca0ac4.
//
// Solidity: function finishJob(bytes32 _jobId, uint256 jobDuration) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) FinishJob(_jobId [32]byte, jobDuration *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.FinishJob(&_MetaScheduler.TransactOpts, _jobId, jobDuration)
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

// GrantRoleOracleMetaScheduler is a paid mutator transaction binding the contract method 0x124b09f3.
//
// Solidity: function grantRoleOracleMetaScheduler(address _oracleMetaScheduler) returns()
func (_MetaScheduler *MetaSchedulerTransactor) GrantRoleOracleMetaScheduler(opts *bind.TransactOpts, _oracleMetaScheduler common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "grantRoleOracleMetaScheduler", _oracleMetaScheduler)
}

// GrantRoleOracleMetaScheduler is a paid mutator transaction binding the contract method 0x124b09f3.
//
// Solidity: function grantRoleOracleMetaScheduler(address _oracleMetaScheduler) returns()
func (_MetaScheduler *MetaSchedulerSession) GrantRoleOracleMetaScheduler(_oracleMetaScheduler common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.GrantRoleOracleMetaScheduler(&_MetaScheduler.TransactOpts, _oracleMetaScheduler)
}

// GrantRoleOracleMetaScheduler is a paid mutator transaction binding the contract method 0x124b09f3.
//
// Solidity: function grantRoleOracleMetaScheduler(address _oracleMetaScheduler) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) GrantRoleOracleMetaScheduler(_oracleMetaScheduler common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.GrantRoleOracleMetaScheduler(&_MetaScheduler.TransactOpts, _oracleMetaScheduler)
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

// OracleLiveness is a paid mutator transaction binding the contract method 0xba9c7f18.
//
// Solidity: function oracleLiveness() returns()
func (_MetaScheduler *MetaSchedulerTransactor) OracleLiveness(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "oracleLiveness")
}

// OracleLiveness is a paid mutator transaction binding the contract method 0xba9c7f18.
//
// Solidity: function oracleLiveness() returns()
func (_MetaScheduler *MetaSchedulerSession) OracleLiveness() (*types.Transaction, error) {
	return _MetaScheduler.Contract.OracleLiveness(&_MetaScheduler.TransactOpts)
}

// OracleLiveness is a paid mutator transaction binding the contract method 0xba9c7f18.
//
// Solidity: function oracleLiveness() returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) OracleLiveness() (*types.Transaction, error) {
	return _MetaScheduler.Contract.OracleLiveness(&_MetaScheduler.TransactOpts)
}

// ProviderApprove is a paid mutator transaction binding the contract method 0x2fcca6fe.
//
// Solidity: function providerApprove(address _providerAddr, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderApprove(opts *bind.TransactOpts, _providerAddr common.Address, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerApprove", _providerAddr, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// ProviderApprove is a paid mutator transaction binding the contract method 0x2fcca6fe.
//
// Solidity: function providerApprove(address _providerAddr, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderApprove(_providerAddr common.Address, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderApprove(&_MetaScheduler.TransactOpts, _providerAddr, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// ProviderApprove is a paid mutator transaction binding the contract method 0x2fcca6fe.
//
// Solidity: function providerApprove(address _providerAddr, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderApprove(_providerAddr common.Address, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderApprove(&_MetaScheduler.TransactOpts, _providerAddr, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// ProviderRedemption is a paid mutator transaction binding the contract method 0xfbc3611a.
//
// Solidity: function providerRedemption() payable returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderRedemption(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerRedemption")
}

// ProviderRedemption is a paid mutator transaction binding the contract method 0xfbc3611a.
//
// Solidity: function providerRedemption() payable returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderRedemption() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderRedemption(&_MetaScheduler.TransactOpts)
}

// ProviderRedemption is a paid mutator transaction binding the contract method 0xfbc3611a.
//
// Solidity: function providerRedemption() payable returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderRedemption() (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderRedemption(&_MetaScheduler.TransactOpts)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0xd1c48e92.
//
// Solidity: function providerRegister(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderRegister(opts *bind.TransactOpts, _providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerRegister", _providerAddr, _nNodes, _gpus, _cpus, _mem)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0xd1c48e92.
//
// Solidity: function providerRegister(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem) returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderRegister(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderRegister(&_MetaScheduler.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0xd1c48e92.
//
// Solidity: function providerRegister(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderRegister(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderRegister(&_MetaScheduler.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem)
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

// RequestNewJob is a paid mutator transaction binding the contract method 0x8e855488.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactor) RequestNewJob(opts *bind.TransactOpts, _definition JobDefinition, _amountLocked *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "requestNewJob", _definition, _amountLocked)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x8e855488.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked) returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) RequestNewJob(_definition JobDefinition, _amountLocked *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _amountLocked)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x8e855488.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactorSession) RequestNewJob(_definition JobDefinition, _amountLocked *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _amountLocked)
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

// StartJob is a paid mutator transaction binding the contract method 0x236e26ae.
//
// Solidity: function startJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactor) StartJob(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "startJob", _jobId)
}

// StartJob is a paid mutator transaction binding the contract method 0x236e26ae.
//
// Solidity: function startJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerSession) StartJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.StartJob(&_MetaScheduler.TransactOpts, _jobId)
}

// StartJob is a paid mutator transaction binding the contract method 0x236e26ae.
//
// Solidity: function startJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) StartJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.StartJob(&_MetaScheduler.TransactOpts, _jobId)
}

// TriggerFailedJob is a paid mutator transaction binding the contract method 0x89be6044.
//
// Solidity: function triggerFailedJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactor) TriggerFailedJob(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "triggerFailedJob", _jobId)
}

// TriggerFailedJob is a paid mutator transaction binding the contract method 0x89be6044.
//
// Solidity: function triggerFailedJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerSession) TriggerFailedJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TriggerFailedJob(&_MetaScheduler.TransactOpts, _jobId)
}

// TriggerFailedJob is a paid mutator transaction binding the contract method 0x89be6044.
//
// Solidity: function triggerFailedJob(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) TriggerFailedJob(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TriggerFailedJob(&_MetaScheduler.TransactOpts, _jobId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactor) Withdraw(opts *bind.TransactOpts, _amount uint64) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 _amount) returns()
func (_MetaScheduler *MetaSchedulerSession) Withdraw(_amount uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Withdraw(&_MetaScheduler.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(uint64 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) Withdraw(_amount uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Withdraw(&_MetaScheduler.TransactOpts, _amount)
}

// MetaSchedulerClaimNextJobEventIterator is returned from FilterClaimNextJobEvent and is used to iterate over the raw logs and unpacked data for ClaimNextJobEvent events raised by the MetaScheduler contract.
type MetaSchedulerClaimNextJobEventIterator struct {
	Event *MetaSchedulerClaimNextJobEvent // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerClaimNextJobEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerClaimNextJobEvent)
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
		it.Event = new(MetaSchedulerClaimNextJobEvent)
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
func (it *MetaSchedulerClaimNextJobEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerClaimNextJobEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerClaimNextJobEvent represents a ClaimNextJobEvent event raised by the MetaScheduler contract.
type MetaSchedulerClaimNextJobEvent struct {
	CustomerAddr      common.Address
	JobId             [32]byte
	MaxDurationMinute *big.Int
	JobDefinition     JobDefinition
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimNextJobEvent is a free log retrieval operation binding the contract event 0xbc999b859676a74ea766936a74f9ade3d1f245417fd1afc6a1e5ca29af316335.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, bytes32 jobId, uint256 maxDurationMinute, (uint64,uint64,uint64,uint64,uint64,string) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) FilterClaimNextJobEvent(opts *bind.FilterOpts) (*MetaSchedulerClaimNextJobEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "ClaimNextJobEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerClaimNextJobEventIterator{contract: _MetaScheduler.contract, event: "ClaimNextJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimNextJobEvent is a free log subscription operation binding the contract event 0xbc999b859676a74ea766936a74f9ade3d1f245417fd1afc6a1e5ca29af316335.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, bytes32 jobId, uint256 maxDurationMinute, (uint64,uint64,uint64,uint64,uint64,string) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) WatchClaimNextJobEvent(opts *bind.WatchOpts, sink chan<- *MetaSchedulerClaimNextJobEvent) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "ClaimNextJobEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerClaimNextJobEvent)
				if err := _MetaScheduler.contract.UnpackLog(event, "ClaimNextJobEvent", log); err != nil {
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

// ParseClaimNextJobEvent is a log parse operation binding the contract event 0xbc999b859676a74ea766936a74f9ade3d1f245417fd1afc6a1e5ca29af316335.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, bytes32 jobId, uint256 maxDurationMinute, (uint64,uint64,uint64,uint64,uint64,string) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) ParseClaimNextJobEvent(log types.Log) (*MetaSchedulerClaimNextJobEvent, error) {
	event := new(MetaSchedulerClaimNextJobEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "ClaimNextJobEvent", log); err != nil {
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

// ProviderManagerMetaData contains all meta data concerning the ProviderManager contract.
var ProviderManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"ToBeApproved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_memPricePerMin\",\"type\":\"uint64\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"decJobCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getProviderFromAddr\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"}],\"internalType\":\"structProvider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasBeenKicked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasJoined\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"incJobCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"joinGrid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nNodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_mem\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalJobCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_metaSchedulerAddr\",\"type\":\"address\"}],\"name\":\"updateRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50620000276000801b336200005760201b60201c565b6000600260006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550620001ba565b6200006982826200014860201b60201c565b6200014457600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620000e9620001b260201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b6127b680620001ca6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c8063624bc8e3116100ad578063a217fddf11610071578063a217fddf1461032b578063b7bb914514610349578063bb26e6e814610379578063d547741f14610395578063e7692d09146103b157610121565b8063624bc8e314610261578063877f4e121461029157806391d14854146102c1578063939daf9c146102f157806396c551751461030f57610121565b80632f2ff15d116100f45780632f2ff15d146101d557806336568abe146101f15780633a80760a1461020d5780633f6edb5f146102295780634ee6f9071461024557610121565b806301ffc9a7146101265780630787bc2714610156578063248a9ca3146101895780632af37fcd146101b9575b600080fd5b610140600480360381019061013b9190611bb7565b6103cd565b60405161014d9190611bff565b60405180910390f35b610170600480360381019061016b9190611c78565b610447565b6040516101809493929190611ddc565b60405180910390f35b6101a3600480360381019061019e9190611e5a565b61060d565b6040516101b09190611e96565b60405180910390f35b6101d360048036038101906101ce9190611edd565b61062c565b005b6101ef60048036038101906101ea9190611f44565b610939565b005b61020b60048036038101906102069190611f44565b61095a565b005b61022760048036038101906102229190611c78565b6109dd565b005b610243600480360381019061023e9190611c78565b6109fa565b005b61025f600480360381019061025a9190611f84565b610ae4565b005b61027b60048036038101906102769190611c78565b610e07565b6040516102889190612103565b60405180910390f35b6102ab60048036038101906102a69190611c78565b611067565b6040516102b89190611bff565b60405180910390f35b6102db60048036038101906102d69190611f44565b6110e7565b6040516102e89190611bff565b60405180910390f35b6102f9611151565b604051610306919061211f565b60405180910390f35b61032960048036038101906103249190611c78565b61116b565b005b61033361126f565b6040516103409190611e96565b60405180910390f35b610363600480360381019061035e9190611c78565b611276565b6040516103709190611bff565b60405180910390f35b610393600480360381019061038e919061215f565b6112f5565b005b6103af60048036038101906103aa9190611f44565b611447565b005b6103cb60048036038101906103c69190611c78565b611468565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610440575061043f8261156b565b5b9050919050565b6001602052806000526040600020600091509050806000016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060020160009054906101000a900460ff16908060020160019054906101000a900460ff16908060020160029054906101000a900467ffffffffffffffff16905084565b6000806000838152602001908152602001600020600101549050919050565b6106396000801b336115d5565b60011515600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160019054906101000a900460ff161515146106cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c6906121fc565b60405180910390fd5b600060028111156106e3576106e2611d56565b5b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff16600281111561074557610744611d56565b5b14610785576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077c90612268565b60405180910390fd5b82600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060018060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff0219169083600281111561092e5761092d611d56565b5b021790555050505050565b6109428261060d565b61094b81611672565b6109558383611686565b505050565b610962611766565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146109cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c6906122fa565b60405180910390fd5b6109d9828261176e565b5050565b6109ea6000801b336115d5565b6109f76000801b82611686565b50565b610a076000801b336115d5565b6002600081819054906101000a900467ffffffffffffffff1680929190610a2d90612349565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201600281819054906101000a900467ffffffffffffffff1680929190610aba90612349565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b610af16000801b336115d5565b60006040518060e001604052808667ffffffffffffffff1681526020018567ffffffffffffffff168152602001600067ffffffffffffffff1681526020018467ffffffffffffffff168152602001600067ffffffffffffffff1681526020018367ffffffffffffffff168152602001600067ffffffffffffffff168152509050604051806080016040528082815260200160006002811115610b9657610b95611d56565b5b8152602001600115158152602001600067ffffffffffffffff16815250600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505060208201518160020160006101000a81548160ff02191690836002811115610d7157610d70611d56565b5b021790555060408201518160020160016101000a81548160ff02191690831515021790555060608201518160020160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050507fc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce287586604051610df79190612388565b60405180910390a1505050505050565b610e0f611a8b565b610e1c6000801b336115d5565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806080016040529081600082016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020016002820160009054906101000a900460ff166002811115610ffd57610ffc611d56565b5b600281111561100f5761100e611d56565b5b81526020016002820160019054906101000a900460ff161515151581526020016002820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250509050919050565b60006001600281111561107d5761107c611d56565b5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff1660028111156110df576110de611d56565b5b149050919050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600260009054906101000a900467ffffffffffffffff1681565b60011515600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160019054906101000a900460ff16151514611201576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f8906121fc565b60405180910390fd5b6002600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff0219169083600281111561126757611266611d56565b5b021790555050565b6000801b81565b600060028081111561128b5761128a611d56565b5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff1660028111156112ed576112ec611d56565b5b149050919050565b6113026000801b336115d5565b6001600681111561131657611315611d56565b5b82600681111561132957611328611d56565b5b14611369576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136090612415565b60405180910390fd5b6002600081819054906101000a900467ffffffffffffffff168092919061138f90612435565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201600281819054906101000a900467ffffffffffffffff168092919061141c90612435565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b6114508261060d565b61145981611672565b611463838361176e565b505050565b60011515600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160019054906101000a900460ff161515146114fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114f5906121fc565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff0219169083600281111561156357611562611d56565b5b021790555050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6115df82826110e7565b61166e576116048173ffffffffffffffffffffffffffffffffffffffff16601461184f565b6116128360001c602061184f565b604051602001611623929190612567565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161166591906125eb565b60405180910390fd5b5050565b6116838161167e611766565b6115d5565b50565b61169082826110e7565b61176257600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550611707611766565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b61177882826110e7565b1561184b57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506117f0611766565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6060600060028360026118629190612617565b61186c9190612659565b67ffffffffffffffff8111156118855761188461268d565b5b6040519080825280601f01601f1916602001820160405280156118b75781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106118ef576118ee6126bc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110611953576119526126bc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026119939190612617565b61199d9190612659565b90505b6001811115611a3d577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106119df576119de6126bc565b5b1a60f81b8282815181106119f6576119f56126bc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080611a36906126eb565b90506119a0565b5060008414611a81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a7890612760565b60405180910390fd5b8091505092915050565b6040518060800160405280611a9e611ad7565b815260200160006002811115611ab757611ab6611d56565b5b8152602001600015158152602001600067ffffffffffffffff1681525090565b6040518060e00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611b9481611b5f565b8114611b9f57600080fd5b50565b600081359050611bb181611b8b565b92915050565b600060208284031215611bcd57611bcc611b5a565b5b6000611bdb84828501611ba2565b91505092915050565b60008115159050919050565b611bf981611be4565b82525050565b6000602082019050611c146000830184611bf0565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611c4582611c1a565b9050919050565b611c5581611c3a565b8114611c6057600080fd5b50565b600081359050611c7281611c4c565b92915050565b600060208284031215611c8e57611c8d611b5a565b5b6000611c9c84828501611c63565b91505092915050565b600067ffffffffffffffff82169050919050565b611cc281611ca5565b82525050565b60e082016000820151611cde6000850182611cb9565b506020820151611cf16020850182611cb9565b506040820151611d046040850182611cb9565b506060820151611d176060850182611cb9565b506080820151611d2a6080850182611cb9565b5060a0820151611d3d60a0850182611cb9565b5060c0820151611d5060c0850182611cb9565b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110611d9657611d95611d56565b5b50565b6000819050611da782611d85565b919050565b6000611db782611d99565b9050919050565b611dc781611dac565b82525050565b611dd681611ca5565b82525050565b600061014082019050611df26000830187611cc8565b611dff60e0830186611dbe565b611e0d610100830185611bf0565b611e1b610120830184611dcd565b95945050505050565b6000819050919050565b611e3781611e24565b8114611e4257600080fd5b50565b600081359050611e5481611e2e565b92915050565b600060208284031215611e7057611e6f611b5a565b5b6000611e7e84828501611e45565b91505092915050565b611e9081611e24565b82525050565b6000602082019050611eab6000830184611e87565b92915050565b611eba81611ca5565b8114611ec557600080fd5b50565b600081359050611ed781611eb1565b92915050565b60008060008060808587031215611ef757611ef6611b5a565b5b6000611f0587828801611c63565b9450506020611f1687828801611ec8565b9350506040611f2787828801611ec8565b9250506060611f3887828801611ec8565b91505092959194509250565b60008060408385031215611f5b57611f5a611b5a565b5b6000611f6985828601611e45565b9250506020611f7a85828601611c63565b9150509250929050565b600080600080600060a08688031215611fa057611f9f611b5a565b5b6000611fae88828901611c63565b9550506020611fbf88828901611ec8565b9450506040611fd088828901611ec8565b9350506060611fe188828901611ec8565b9250506080611ff288828901611ec8565b9150509295509295909350565b60e0820160008201516120156000850182611cb9565b5060208201516120286020850182611cb9565b50604082015161203b6040850182611cb9565b50606082015161204e6060850182611cb9565b5060808201516120616080850182611cb9565b5060a082015161207460a0850182611cb9565b5060c082015161208760c0850182611cb9565b50505050565b61209681611dac565b82525050565b6120a581611be4565b82525050565b610140820160008201516120c26000850182611fff565b5060208201516120d560e085018261208d565b5060408201516120e961010085018261209c565b5060608201516120fd610120850182611cb9565b50505050565b60006101408201905061211960008301846120ab565b92915050565b60006020820190506121346000830184611dcd565b92915050565b6007811061214757600080fd5b50565b6000813590506121598161213a565b92915050565b6000806040838503121561217657612175611b5a565b5b60006121848582860161214a565b925050602061219585828601611c63565b9150509250929050565b600082825260208201905092915050565b7f4e6f2070726f766964657220666f756e64000000000000000000000000000000600082015250565b60006121e660118361219f565b91506121f1826121b0565b602082019050919050565b60006020820190508181036000830152612215816121d9565b9050919050565b7f416c726561647920617070726f76656400000000000000000000000000000000600082015250565b600061225260108361219f565b915061225d8261221c565b602082019050919050565b6000602082019050818103600083015261228181612245565b9050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b60006122e4602f8361219f565b91506122ef82612288565b604082019050919050565b60006020820190508181036000830152612313816122d7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061235482611ca5565b915067ffffffffffffffff820361236e5761236d61231a565b5b600182019050919050565b61238281611c3a565b82525050565b600060208201905061239d6000830184612379565b92915050565b7f4f6e6c79204d4554415f5343484544554c4544206a6f62732063616e2062652060008201527f6465637265617365640000000000000000000000000000000000000000000000602082015250565b60006123ff60298361219f565b915061240a826123a3565b604082019050919050565b6000602082019050818103600083015261242e816123f2565b9050919050565b600061244082611ca5565b9150600082036124535761245261231a565b5b600182039050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b600061249f60178361245e565b91506124aa82612469565b601782019050919050565b600081519050919050565b60005b838110156124de5780820151818401526020810190506124c3565b60008484015250505050565b60006124f5826124b5565b6124ff818561245e565b935061250f8185602086016124c0565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b600061255160118361245e565b915061255c8261251b565b601182019050919050565b600061257282612492565b915061257e82856124ea565b915061258982612544565b915061259582846124ea565b91508190509392505050565b6000601f19601f8301169050919050565b60006125bd826124b5565b6125c7818561219f565b93506125d78185602086016124c0565b6125e0816125a1565b840191505092915050565b6000602082019050818103600083015261260581846125b2565b905092915050565b6000819050919050565b60006126228261260d565b915061262d8361260d565b925082820261263b8161260d565b915082820484148315176126525761265161231a565b5b5092915050565b60006126648261260d565b915061266f8361260d565b92508282019050808211156126875761268661231a565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006126f68261260d565b9150600082036127095761270861231a565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b600061274a60208361219f565b915061275582612714565b602082019050919050565b600060208201905081810360008301526127798161273d565b905091905056fea2646970667358221220b72bb4fa88170d24af74cc2f2e45693c188a6c610ad69320d4e79ad72d43885864736f6c63430008110033",
}

// ProviderManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ProviderManagerMetaData.ABI instead.
var ProviderManagerABI = ProviderManagerMetaData.ABI

// ProviderManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProviderManagerMetaData.Bin instead.
var ProviderManagerBin = ProviderManagerMetaData.Bin

// DeployProviderManager deploys a new Ethereum contract, binding an instance of ProviderManager to it.
func DeployProviderManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProviderManager, error) {
	parsed, err := ProviderManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProviderManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProviderManager{ProviderManagerCaller: ProviderManagerCaller{contract: contract}, ProviderManagerTransactor: ProviderManagerTransactor{contract: contract}, ProviderManagerFilterer: ProviderManagerFilterer{contract: contract}}, nil
}

// ProviderManager is an auto generated Go binding around an Ethereum contract.
type ProviderManager struct {
	ProviderManagerCaller     // Read-only binding to the contract
	ProviderManagerTransactor // Write-only binding to the contract
	ProviderManagerFilterer   // Log filterer for contract events
}

// ProviderManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderManagerSession struct {
	Contract     *ProviderManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderManagerCallerSession struct {
	Contract *ProviderManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProviderManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderManagerTransactorSession struct {
	Contract     *ProviderManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProviderManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderManagerRaw struct {
	Contract *ProviderManager // Generic contract binding to access the raw methods on
}

// ProviderManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderManagerCallerRaw struct {
	Contract *ProviderManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderManagerTransactorRaw struct {
	Contract *ProviderManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProviderManager creates a new instance of ProviderManager, bound to a specific deployed contract.
func NewProviderManager(address common.Address, backend bind.ContractBackend) (*ProviderManager, error) {
	contract, err := bindProviderManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProviderManager{ProviderManagerCaller: ProviderManagerCaller{contract: contract}, ProviderManagerTransactor: ProviderManagerTransactor{contract: contract}, ProviderManagerFilterer: ProviderManagerFilterer{contract: contract}}, nil
}

// NewProviderManagerCaller creates a new read-only instance of ProviderManager, bound to a specific deployed contract.
func NewProviderManagerCaller(address common.Address, caller bind.ContractCaller) (*ProviderManagerCaller, error) {
	contract, err := bindProviderManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderManagerCaller{contract: contract}, nil
}

// NewProviderManagerTransactor creates a new write-only instance of ProviderManager, bound to a specific deployed contract.
func NewProviderManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderManagerTransactor, error) {
	contract, err := bindProviderManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderManagerTransactor{contract: contract}, nil
}

// NewProviderManagerFilterer creates a new log filterer instance of ProviderManager, bound to a specific deployed contract.
func NewProviderManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderManagerFilterer, error) {
	contract, err := bindProviderManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderManagerFilterer{contract: contract}, nil
}

// bindProviderManager binds a generic wrapper to an already deployed contract.
func bindProviderManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderManager *ProviderManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderManager.Contract.ProviderManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderManager *ProviderManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderManager.Contract.ProviderManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderManager *ProviderManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderManager.Contract.ProviderManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderManager *ProviderManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderManager *ProviderManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderManager *ProviderManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProviderManager *ProviderManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProviderManager *ProviderManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProviderManager.Contract.DEFAULTADMINROLE(&_ProviderManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProviderManager *ProviderManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProviderManager.Contract.DEFAULTADMINROLE(&_ProviderManager.CallOpts)
}

// GetProviderFromAddr is a free data retrieval call binding the contract method 0x624bc8e3.
//
// Solidity: function getProviderFromAddr(address _providerAddr) view returns(((uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64))
func (_ProviderManager *ProviderManagerCaller) GetProviderFromAddr(opts *bind.CallOpts, _providerAddr common.Address) (Provider, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "getProviderFromAddr", _providerAddr)

	if err != nil {
		return *new(Provider), err
	}

	out0 := *abi.ConvertType(out[0], new(Provider)).(*Provider)

	return out0, err

}

// GetProviderFromAddr is a free data retrieval call binding the contract method 0x624bc8e3.
//
// Solidity: function getProviderFromAddr(address _providerAddr) view returns(((uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64))
func (_ProviderManager *ProviderManagerSession) GetProviderFromAddr(_providerAddr common.Address) (Provider, error) {
	return _ProviderManager.Contract.GetProviderFromAddr(&_ProviderManager.CallOpts, _providerAddr)
}

// GetProviderFromAddr is a free data retrieval call binding the contract method 0x624bc8e3.
//
// Solidity: function getProviderFromAddr(address _providerAddr) view returns(((uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64))
func (_ProviderManager *ProviderManagerCallerSession) GetProviderFromAddr(_providerAddr common.Address) (Provider, error) {
	return _ProviderManager.Contract.GetProviderFromAddr(&_ProviderManager.CallOpts, _providerAddr)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProviderManager *ProviderManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProviderManager *ProviderManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProviderManager.Contract.GetRoleAdmin(&_ProviderManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProviderManager *ProviderManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProviderManager.Contract.GetRoleAdmin(&_ProviderManager.CallOpts, role)
}

// HasBeenKicked is a free data retrieval call binding the contract method 0xb7bb9145.
//
// Solidity: function hasBeenKicked(address _providerAddr) view returns(bool)
func (_ProviderManager *ProviderManagerCaller) HasBeenKicked(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "hasBeenKicked", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBeenKicked is a free data retrieval call binding the contract method 0xb7bb9145.
//
// Solidity: function hasBeenKicked(address _providerAddr) view returns(bool)
func (_ProviderManager *ProviderManagerSession) HasBeenKicked(_providerAddr common.Address) (bool, error) {
	return _ProviderManager.Contract.HasBeenKicked(&_ProviderManager.CallOpts, _providerAddr)
}

// HasBeenKicked is a free data retrieval call binding the contract method 0xb7bb9145.
//
// Solidity: function hasBeenKicked(address _providerAddr) view returns(bool)
func (_ProviderManager *ProviderManagerCallerSession) HasBeenKicked(_providerAddr common.Address) (bool, error) {
	return _ProviderManager.Contract.HasBeenKicked(&_ProviderManager.CallOpts, _providerAddr)
}

// HasJoined is a free data retrieval call binding the contract method 0x877f4e12.
//
// Solidity: function hasJoined(address _providerAddr) view returns(bool)
func (_ProviderManager *ProviderManagerCaller) HasJoined(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "hasJoined", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasJoined is a free data retrieval call binding the contract method 0x877f4e12.
//
// Solidity: function hasJoined(address _providerAddr) view returns(bool)
func (_ProviderManager *ProviderManagerSession) HasJoined(_providerAddr common.Address) (bool, error) {
	return _ProviderManager.Contract.HasJoined(&_ProviderManager.CallOpts, _providerAddr)
}

// HasJoined is a free data retrieval call binding the contract method 0x877f4e12.
//
// Solidity: function hasJoined(address _providerAddr) view returns(bool)
func (_ProviderManager *ProviderManagerCallerSession) HasJoined(_providerAddr common.Address) (bool, error) {
	return _ProviderManager.Contract.HasJoined(&_ProviderManager.CallOpts, _providerAddr)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProviderManager *ProviderManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProviderManager *ProviderManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProviderManager.Contract.HasRole(&_ProviderManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProviderManager *ProviderManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProviderManager.Contract.HasRole(&_ProviderManager.CallOpts, role, account)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount)
func (_ProviderManager *ProviderManagerCaller) Providers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Definition ProviderDefinition
	Status     uint8
	Valid      bool
	JobCount   uint64
}, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "providers", arg0)

	outstruct := new(struct {
		Definition ProviderDefinition
		Status     uint8
		Valid      bool
		JobCount   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Definition = *abi.ConvertType(out[0], new(ProviderDefinition)).(*ProviderDefinition)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Valid = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.JobCount = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount)
func (_ProviderManager *ProviderManagerSession) Providers(arg0 common.Address) (struct {
	Definition ProviderDefinition
	Status     uint8
	Valid      bool
	JobCount   uint64
}, error) {
	return _ProviderManager.Contract.Providers(&_ProviderManager.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount)
func (_ProviderManager *ProviderManagerCallerSession) Providers(arg0 common.Address) (struct {
	Definition ProviderDefinition
	Status     uint8
	Valid      bool
	JobCount   uint64
}, error) {
	return _ProviderManager.Contract.Providers(&_ProviderManager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProviderManager *ProviderManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProviderManager *ProviderManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProviderManager.Contract.SupportsInterface(&_ProviderManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProviderManager *ProviderManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProviderManager.Contract.SupportsInterface(&_ProviderManager.CallOpts, interfaceId)
}

// TotalJobCount is a free data retrieval call binding the contract method 0x939daf9c.
//
// Solidity: function totalJobCount() view returns(uint64)
func (_ProviderManager *ProviderManagerCaller) TotalJobCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "totalJobCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalJobCount is a free data retrieval call binding the contract method 0x939daf9c.
//
// Solidity: function totalJobCount() view returns(uint64)
func (_ProviderManager *ProviderManagerSession) TotalJobCount() (uint64, error) {
	return _ProviderManager.Contract.TotalJobCount(&_ProviderManager.CallOpts)
}

// TotalJobCount is a free data retrieval call binding the contract method 0x939daf9c.
//
// Solidity: function totalJobCount() view returns(uint64)
func (_ProviderManager *ProviderManagerCallerSession) TotalJobCount() (uint64, error) {
	return _ProviderManager.Contract.TotalJobCount(&_ProviderManager.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x2af37fcd.
//
// Solidity: function approve(address _providerAddr, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerTransactor) Approve(opts *bind.TransactOpts, _providerAddr common.Address, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "approve", _providerAddr, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// Approve is a paid mutator transaction binding the contract method 0x2af37fcd.
//
// Solidity: function approve(address _providerAddr, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerSession) Approve(_providerAddr common.Address, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.Approve(&_ProviderManager.TransactOpts, _providerAddr, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// Approve is a paid mutator transaction binding the contract method 0x2af37fcd.
//
// Solidity: function approve(address _providerAddr, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerTransactorSession) Approve(_providerAddr common.Address, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.Approve(&_ProviderManager.TransactOpts, _providerAddr, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// DecJobCount is a paid mutator transaction binding the contract method 0xbb26e6e8.
//
// Solidity: function decJobCount(uint8 _jobStatus, address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactor) DecJobCount(opts *bind.TransactOpts, _jobStatus uint8, _providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "decJobCount", _jobStatus, _providerAddr)
}

// DecJobCount is a paid mutator transaction binding the contract method 0xbb26e6e8.
//
// Solidity: function decJobCount(uint8 _jobStatus, address _providerAddr) returns()
func (_ProviderManager *ProviderManagerSession) DecJobCount(_jobStatus uint8, _providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.DecJobCount(&_ProviderManager.TransactOpts, _jobStatus, _providerAddr)
}

// DecJobCount is a paid mutator transaction binding the contract method 0xbb26e6e8.
//
// Solidity: function decJobCount(uint8 _jobStatus, address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactorSession) DecJobCount(_jobStatus uint8, _providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.DecJobCount(&_ProviderManager.TransactOpts, _jobStatus, _providerAddr)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.GrantRole(&_ProviderManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.GrantRole(&_ProviderManager.TransactOpts, role, account)
}

// IncJobCount is a paid mutator transaction binding the contract method 0x3f6edb5f.
//
// Solidity: function incJobCount(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactor) IncJobCount(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "incJobCount", _providerAddr)
}

// IncJobCount is a paid mutator transaction binding the contract method 0x3f6edb5f.
//
// Solidity: function incJobCount(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerSession) IncJobCount(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.IncJobCount(&_ProviderManager.TransactOpts, _providerAddr)
}

// IncJobCount is a paid mutator transaction binding the contract method 0x3f6edb5f.
//
// Solidity: function incJobCount(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactorSession) IncJobCount(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.IncJobCount(&_ProviderManager.TransactOpts, _providerAddr)
}

// JoinGrid is a paid mutator transaction binding the contract method 0xe7692d09.
//
// Solidity: function joinGrid(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactor) JoinGrid(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "joinGrid", _providerAddr)
}

// JoinGrid is a paid mutator transaction binding the contract method 0xe7692d09.
//
// Solidity: function joinGrid(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerSession) JoinGrid(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.JoinGrid(&_ProviderManager.TransactOpts, _providerAddr)
}

// JoinGrid is a paid mutator transaction binding the contract method 0xe7692d09.
//
// Solidity: function joinGrid(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactorSession) JoinGrid(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.JoinGrid(&_ProviderManager.TransactOpts, _providerAddr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactor) Kick(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "kick", _providerAddr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerSession) Kick(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.Kick(&_ProviderManager.TransactOpts, _providerAddr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactorSession) Kick(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.Kick(&_ProviderManager.TransactOpts, _providerAddr)
}

// Register is a paid mutator transaction binding the contract method 0x4ee6f907.
//
// Solidity: function register(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem) returns()
func (_ProviderManager *ProviderManagerTransactor) Register(opts *bind.TransactOpts, _providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "register", _providerAddr, _nNodes, _gpus, _cpus, _mem)
}

// Register is a paid mutator transaction binding the contract method 0x4ee6f907.
//
// Solidity: function register(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem) returns()
func (_ProviderManager *ProviderManagerSession) Register(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.Register(&_ProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem)
}

// Register is a paid mutator transaction binding the contract method 0x4ee6f907.
//
// Solidity: function register(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem) returns()
func (_ProviderManager *ProviderManagerTransactorSession) Register(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.Register(&_ProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.RenounceRole(&_ProviderManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.RenounceRole(&_ProviderManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.RevokeRole(&_ProviderManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProviderManager *ProviderManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.RevokeRole(&_ProviderManager.TransactOpts, role, account)
}

// UpdateRoles is a paid mutator transaction binding the contract method 0x3a80760a.
//
// Solidity: function updateRoles(address _metaSchedulerAddr) returns()
func (_ProviderManager *ProviderManagerTransactor) UpdateRoles(opts *bind.TransactOpts, _metaSchedulerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "updateRoles", _metaSchedulerAddr)
}

// UpdateRoles is a paid mutator transaction binding the contract method 0x3a80760a.
//
// Solidity: function updateRoles(address _metaSchedulerAddr) returns()
func (_ProviderManager *ProviderManagerSession) UpdateRoles(_metaSchedulerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.UpdateRoles(&_ProviderManager.TransactOpts, _metaSchedulerAddr)
}

// UpdateRoles is a paid mutator transaction binding the contract method 0x3a80760a.
//
// Solidity: function updateRoles(address _metaSchedulerAddr) returns()
func (_ProviderManager *ProviderManagerTransactorSession) UpdateRoles(_metaSchedulerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.UpdateRoles(&_ProviderManager.TransactOpts, _metaSchedulerAddr)
}

// ProviderManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ProviderManager contract.
type ProviderManagerRoleAdminChangedIterator struct {
	Event *ProviderManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProviderManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderManagerRoleAdminChanged)
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
		it.Event = new(ProviderManagerRoleAdminChanged)
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
func (it *ProviderManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderManagerRoleAdminChanged represents a RoleAdminChanged event raised by the ProviderManager contract.
type ProviderManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProviderManager *ProviderManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ProviderManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ProviderManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ProviderManagerRoleAdminChangedIterator{contract: _ProviderManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProviderManager *ProviderManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ProviderManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ProviderManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderManagerRoleAdminChanged)
				if err := _ProviderManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ProviderManager *ProviderManagerFilterer) ParseRoleAdminChanged(log types.Log) (*ProviderManagerRoleAdminChanged, error) {
	event := new(ProviderManagerRoleAdminChanged)
	if err := _ProviderManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ProviderManager contract.
type ProviderManagerRoleGrantedIterator struct {
	Event *ProviderManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *ProviderManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderManagerRoleGranted)
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
		it.Event = new(ProviderManagerRoleGranted)
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
func (it *ProviderManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderManagerRoleGranted represents a RoleGranted event raised by the ProviderManager contract.
type ProviderManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProviderManager *ProviderManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProviderManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _ProviderManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProviderManagerRoleGrantedIterator{contract: _ProviderManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProviderManager *ProviderManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ProviderManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ProviderManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderManagerRoleGranted)
				if err := _ProviderManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ProviderManager *ProviderManagerFilterer) ParseRoleGranted(log types.Log) (*ProviderManagerRoleGranted, error) {
	event := new(ProviderManagerRoleGranted)
	if err := _ProviderManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ProviderManager contract.
type ProviderManagerRoleRevokedIterator struct {
	Event *ProviderManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ProviderManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderManagerRoleRevoked)
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
		it.Event = new(ProviderManagerRoleRevoked)
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
func (it *ProviderManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderManagerRoleRevoked represents a RoleRevoked event raised by the ProviderManager contract.
type ProviderManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProviderManager *ProviderManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProviderManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _ProviderManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProviderManagerRoleRevokedIterator{contract: _ProviderManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProviderManager *ProviderManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ProviderManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ProviderManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderManagerRoleRevoked)
				if err := _ProviderManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ProviderManager *ProviderManagerFilterer) ParseRoleRevoked(log types.Log) (*ProviderManagerRoleRevoked, error) {
	event := new(ProviderManagerRoleRevoked)
	if err := _ProviderManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderManagerToBeApprovedIterator is returned from FilterToBeApproved and is used to iterate over the raw logs and unpacked data for ToBeApproved events raised by the ProviderManager contract.
type ProviderManagerToBeApprovedIterator struct {
	Event *ProviderManagerToBeApproved // Event containing the contract specifics and raw log

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
func (it *ProviderManagerToBeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderManagerToBeApproved)
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
		it.Event = new(ProviderManagerToBeApproved)
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
func (it *ProviderManagerToBeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderManagerToBeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderManagerToBeApproved represents a ToBeApproved event raised by the ProviderManager contract.
type ProviderManagerToBeApproved struct {
	ProviderAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterToBeApproved is a free log retrieval operation binding the contract event 0xc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce2875.
//
// Solidity: event ToBeApproved(address _providerAddr)
func (_ProviderManager *ProviderManagerFilterer) FilterToBeApproved(opts *bind.FilterOpts) (*ProviderManagerToBeApprovedIterator, error) {

	logs, sub, err := _ProviderManager.contract.FilterLogs(opts, "ToBeApproved")
	if err != nil {
		return nil, err
	}
	return &ProviderManagerToBeApprovedIterator{contract: _ProviderManager.contract, event: "ToBeApproved", logs: logs, sub: sub}, nil
}

// WatchToBeApproved is a free log subscription operation binding the contract event 0xc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce2875.
//
// Solidity: event ToBeApproved(address _providerAddr)
func (_ProviderManager *ProviderManagerFilterer) WatchToBeApproved(opts *bind.WatchOpts, sink chan<- *ProviderManagerToBeApproved) (event.Subscription, error) {

	logs, sub, err := _ProviderManager.contract.WatchLogs(opts, "ToBeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderManagerToBeApproved)
				if err := _ProviderManager.contract.UnpackLog(event, "ToBeApproved", log); err != nil {
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

// ParseToBeApproved is a log parse operation binding the contract event 0xc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce2875.
//
// Solidity: event ToBeApproved(address _providerAddr)
func (_ProviderManager *ProviderManagerFilterer) ParseToBeApproved(log types.Log) (*ProviderManagerToBeApproved, error) {
	event := new(ProviderManagerToBeApproved)
	if err := _ProviderManager.contract.UnpackLog(event, "ToBeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220735ec1f653914c63bc4fafa80f087af90c0b4f7136e3ca338319e624b4f9ff3464736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(SafeCastABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bffded732aebf91122fca0fd45fc9f602ec710ed32beef917ebc45ecbc29a65c64736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// TimeIsMoneyMetaData contains all meta data concerning the TimeIsMoney contract.
var TimeIsMoneyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"job\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"provider\",\"type\":\"tuple\"}],\"name\":\"convertCreditToDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"job\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"provider\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"durationMinute\",\"type\":\"uint256\"}],\"name\":\"convertJobDurationInMoney\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"job\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"provider\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"convertMoneyInJobDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x61082e610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063214ef282146100505780637316b6ef14610080578063a0a56273146100b0575b600080fd5b61006a600480360381019061006591906105e8565b6100e0565b6040516100779190610668565b60405180910390f35b61009a60048036038101906100959190610683565b6101b4565b6040516100a79190610668565b60405180910390f35b6100ca60048036038101906100c591906105e8565b6101ca565b6040516100d79190610668565b60405180910390f35b6000836040015167ffffffffffffffff16846080015167ffffffffffffffff16846080015167ffffffffffffffff166101199190610722565b6101239190610722565b846020015167ffffffffffffffff168460c0015167ffffffffffffffff1661014b9190610722565b856000015167ffffffffffffffff16856040015167ffffffffffffffff166101739190610722565b61017d9190610764565b6101879190610764565b846060015167ffffffffffffffff16836101a19190610722565b6101ab9190610722565b90509392505050565b60006101c18483856101ca565b90509392505050565b6000836040015167ffffffffffffffff16846080015167ffffffffffffffff16846080015167ffffffffffffffff166102039190610722565b61020d9190610722565b846020015167ffffffffffffffff168460c0015167ffffffffffffffff166102359190610722565b856000015167ffffffffffffffff16856040015167ffffffffffffffff1661025d9190610722565b6102679190610764565b6102719190610764565b846060015167ffffffffffffffff1661028a9190610722565b8261029591906107c7565b90509392505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610300826102b7565b810181811067ffffffffffffffff8211171561031f5761031e6102c8565b5b80604052505050565b600061033261029e565b905061033e82826102f7565b919050565b600080fd5b600067ffffffffffffffff82169050919050565b61036581610348565b811461037057600080fd5b50565b6000813590506103828161035c565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff8211156103ad576103ac6102c8565b5b6103b6826102b7565b9050602081019050919050565b82818337600083830152505050565b60006103e56103e084610392565b610328565b9050828152602081018484840111156104015761040061038d565b5b61040c8482856103c3565b509392505050565b600082601f83011261042957610428610388565b5b81356104398482602086016103d2565b91505092915050565b600060c08284031215610458576104576102b2565b5b61046260c0610328565b9050600061047284828501610373565b600083015250602061048684828501610373565b602083015250604061049a84828501610373565b60408301525060606104ae84828501610373565b60608301525060806104c284828501610373565b60808301525060a082013567ffffffffffffffff8111156104e6576104e5610343565b5b6104f284828501610414565b60a08301525092915050565b600060e08284031215610514576105136102b2565b5b61051e60e0610328565b9050600061052e84828501610373565b600083015250602061054284828501610373565b602083015250604061055684828501610373565b604083015250606061056a84828501610373565b606083015250608061057e84828501610373565b60808301525060a061059284828501610373565b60a08301525060c06105a684828501610373565b60c08301525092915050565b6000819050919050565b6105c5816105b2565b81146105d057600080fd5b50565b6000813590506105e2816105bc565b92915050565b60008060006101208486031215610602576106016102a8565b5b600084013567ffffffffffffffff8111156106205761061f6102ad565b5b61062c86828701610442565b935050602061063d868287016104fe565b92505061010061064f868287016105d3565b9150509250925092565b610662816105b2565b82525050565b600060208201905061067d6000830184610659565b92915050565b6000806000610120848603121561069d5761069c6102a8565b5b600084013567ffffffffffffffff8111156106bb576106ba6102ad565b5b6106c786828701610442565b93505060206106d8868287016105d3565b92505060406106e9868287016104fe565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061072d826105b2565b9150610738836105b2565b9250828202610746816105b2565b9150828204841483151761075d5761075c6106f3565b5b5092915050565b600061076f826105b2565b915061077a836105b2565b9250828201905080821115610792576107916106f3565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006107d2826105b2565b91506107dd836105b2565b9250826107ed576107ec610798565b5b82820490509291505056fea264697066735822122021dbed64587c79a3e7fe85bcbda5ae1ca75620db9b7f5c22fc08a28fcd73d79d64736f6c63430008110033",
}

// TimeIsMoneyABI is the input ABI used to generate the binding from.
// Deprecated: Use TimeIsMoneyMetaData.ABI instead.
var TimeIsMoneyABI = TimeIsMoneyMetaData.ABI

// TimeIsMoneyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimeIsMoneyMetaData.Bin instead.
var TimeIsMoneyBin = TimeIsMoneyMetaData.Bin

// DeployTimeIsMoney deploys a new Ethereum contract, binding an instance of TimeIsMoney to it.
func DeployTimeIsMoney(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TimeIsMoney, error) {
	parsed, err := TimeIsMoneyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimeIsMoneyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TimeIsMoney{TimeIsMoneyCaller: TimeIsMoneyCaller{contract: contract}, TimeIsMoneyTransactor: TimeIsMoneyTransactor{contract: contract}, TimeIsMoneyFilterer: TimeIsMoneyFilterer{contract: contract}}, nil
}

// TimeIsMoney is an auto generated Go binding around an Ethereum contract.
type TimeIsMoney struct {
	TimeIsMoneyCaller     // Read-only binding to the contract
	TimeIsMoneyTransactor // Write-only binding to the contract
	TimeIsMoneyFilterer   // Log filterer for contract events
}

// TimeIsMoneyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeIsMoneyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeIsMoneyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeIsMoneyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeIsMoneyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeIsMoneyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeIsMoneySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeIsMoneySession struct {
	Contract     *TimeIsMoney      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeIsMoneyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeIsMoneyCallerSession struct {
	Contract *TimeIsMoneyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TimeIsMoneyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeIsMoneyTransactorSession struct {
	Contract     *TimeIsMoneyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TimeIsMoneyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeIsMoneyRaw struct {
	Contract *TimeIsMoney // Generic contract binding to access the raw methods on
}

// TimeIsMoneyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeIsMoneyCallerRaw struct {
	Contract *TimeIsMoneyCaller // Generic read-only contract binding to access the raw methods on
}

// TimeIsMoneyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeIsMoneyTransactorRaw struct {
	Contract *TimeIsMoneyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimeIsMoney creates a new instance of TimeIsMoney, bound to a specific deployed contract.
func NewTimeIsMoney(address common.Address, backend bind.ContractBackend) (*TimeIsMoney, error) {
	contract, err := bindTimeIsMoney(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimeIsMoney{TimeIsMoneyCaller: TimeIsMoneyCaller{contract: contract}, TimeIsMoneyTransactor: TimeIsMoneyTransactor{contract: contract}, TimeIsMoneyFilterer: TimeIsMoneyFilterer{contract: contract}}, nil
}

// NewTimeIsMoneyCaller creates a new read-only instance of TimeIsMoney, bound to a specific deployed contract.
func NewTimeIsMoneyCaller(address common.Address, caller bind.ContractCaller) (*TimeIsMoneyCaller, error) {
	contract, err := bindTimeIsMoney(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeIsMoneyCaller{contract: contract}, nil
}

// NewTimeIsMoneyTransactor creates a new write-only instance of TimeIsMoney, bound to a specific deployed contract.
func NewTimeIsMoneyTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeIsMoneyTransactor, error) {
	contract, err := bindTimeIsMoney(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeIsMoneyTransactor{contract: contract}, nil
}

// NewTimeIsMoneyFilterer creates a new log filterer instance of TimeIsMoney, bound to a specific deployed contract.
func NewTimeIsMoneyFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeIsMoneyFilterer, error) {
	contract, err := bindTimeIsMoney(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeIsMoneyFilterer{contract: contract}, nil
}

// bindTimeIsMoney binds a generic wrapper to an already deployed contract.
func bindTimeIsMoney(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeIsMoneyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeIsMoney *TimeIsMoneyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeIsMoney.Contract.TimeIsMoneyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeIsMoney *TimeIsMoneyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeIsMoney.Contract.TimeIsMoneyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeIsMoney *TimeIsMoneyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeIsMoney.Contract.TimeIsMoneyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeIsMoney *TimeIsMoneyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeIsMoney.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeIsMoney *TimeIsMoneyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeIsMoney.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeIsMoney *TimeIsMoneyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeIsMoney.Contract.contract.Transact(opts, method, params...)
}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0xf688411b.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,uint64,string) job, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneyCaller) ConvertCreditToDuration(opts *bind.CallOpts, job JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (*big.Int, error) {
	var out []interface{}
	err := _TimeIsMoney.contract.Call(opts, &out, "convertCreditToDuration", job, amountLocked, provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0xf688411b.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,uint64,string) job, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneySession) ConvertCreditToDuration(job JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (*big.Int, error) {
	return _TimeIsMoney.Contract.ConvertCreditToDuration(&_TimeIsMoney.CallOpts, job, amountLocked, provider)
}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0xf688411b.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,uint64,string) job, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneyCallerSession) ConvertCreditToDuration(job JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (*big.Int, error) {
	return _TimeIsMoney.Contract.ConvertCreditToDuration(&_TimeIsMoney.CallOpts, job, amountLocked, provider)
}

// ConvertJobDurationInMoney is a free data retrieval call binding the contract method 0x3a9b3c88.
//
// Solidity: function convertJobDurationInMoney((uint64,uint64,uint64,uint64,uint64,string) job, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint256 durationMinute) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneyCaller) ConvertJobDurationInMoney(opts *bind.CallOpts, job JobDefinition, provider ProviderDefinition, durationMinute *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TimeIsMoney.contract.Call(opts, &out, "convertJobDurationInMoney", job, provider, durationMinute)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertJobDurationInMoney is a free data retrieval call binding the contract method 0x3a9b3c88.
//
// Solidity: function convertJobDurationInMoney((uint64,uint64,uint64,uint64,uint64,string) job, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint256 durationMinute) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneySession) ConvertJobDurationInMoney(job JobDefinition, provider ProviderDefinition, durationMinute *big.Int) (*big.Int, error) {
	return _TimeIsMoney.Contract.ConvertJobDurationInMoney(&_TimeIsMoney.CallOpts, job, provider, durationMinute)
}

// ConvertJobDurationInMoney is a free data retrieval call binding the contract method 0x3a9b3c88.
//
// Solidity: function convertJobDurationInMoney((uint64,uint64,uint64,uint64,uint64,string) job, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint256 durationMinute) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneyCallerSession) ConvertJobDurationInMoney(job JobDefinition, provider ProviderDefinition, durationMinute *big.Int) (*big.Int, error) {
	return _TimeIsMoney.Contract.ConvertJobDurationInMoney(&_TimeIsMoney.CallOpts, job, provider, durationMinute)
}

// ConvertMoneyInJobDuration is a free data retrieval call binding the contract method 0x7bf67420.
//
// Solidity: function convertMoneyInJobDuration((uint64,uint64,uint64,uint64,uint64,string) job, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint256 cost) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneyCaller) ConvertMoneyInJobDuration(opts *bind.CallOpts, job JobDefinition, provider ProviderDefinition, cost *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TimeIsMoney.contract.Call(opts, &out, "convertMoneyInJobDuration", job, provider, cost)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertMoneyInJobDuration is a free data retrieval call binding the contract method 0x7bf67420.
//
// Solidity: function convertMoneyInJobDuration((uint64,uint64,uint64,uint64,uint64,string) job, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint256 cost) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneySession) ConvertMoneyInJobDuration(job JobDefinition, provider ProviderDefinition, cost *big.Int) (*big.Int, error) {
	return _TimeIsMoney.Contract.ConvertMoneyInJobDuration(&_TimeIsMoney.CallOpts, job, provider, cost)
}

// ConvertMoneyInJobDuration is a free data retrieval call binding the contract method 0x7bf67420.
//
// Solidity: function convertMoneyInJobDuration((uint64,uint64,uint64,uint64,uint64,string) job, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint256 cost) pure returns(uint256)
func (_TimeIsMoney *TimeIsMoneyCallerSession) ConvertMoneyInJobDuration(job JobDefinition, provider ProviderDefinition, cost *big.Int) (*big.Int, error) {
	return _TimeIsMoney.Contract.ConvertMoneyInJobDuration(&_TimeIsMoney.CallOpts, job, provider, cost)
}

// TimeoutManagementMetaData contains all meta data concerning the TimeoutManagement contract.
var TimeoutManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockOrigin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"stillAlive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101cd610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806315a945ce1461003a575b600080fd5b610054600480360381019061004f91906100be565b61006a565b6040516100619190610119565b60405180910390f35b60008183436100799190610163565b1015905092915050565b600080fd5b6000819050919050565b61009b81610088565b81146100a657600080fd5b50565b6000813590506100b881610092565b92915050565b600080604083850312156100d5576100d4610083565b5b60006100e3858286016100a9565b92505060206100f4858286016100a9565b9150509250929050565b60008115159050919050565b610113816100fe565b82525050565b600060208201905061012e600083018461010a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061016e82610088565b915061017983610088565b925082820390508181111561019157610190610134565b5b9291505056fea26469706673582212208eee8d3481005d9804b72cf5a70dc8c4e97bb7988a2ec80b846a77017ddc66b564736f6c63430008110033",
}

// TimeoutManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use TimeoutManagementMetaData.ABI instead.
var TimeoutManagementABI = TimeoutManagementMetaData.ABI

// TimeoutManagementBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimeoutManagementMetaData.Bin instead.
var TimeoutManagementBin = TimeoutManagementMetaData.Bin

// DeployTimeoutManagement deploys a new Ethereum contract, binding an instance of TimeoutManagement to it.
func DeployTimeoutManagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TimeoutManagement, error) {
	parsed, err := TimeoutManagementMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimeoutManagementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TimeoutManagement{TimeoutManagementCaller: TimeoutManagementCaller{contract: contract}, TimeoutManagementTransactor: TimeoutManagementTransactor{contract: contract}, TimeoutManagementFilterer: TimeoutManagementFilterer{contract: contract}}, nil
}

// TimeoutManagement is an auto generated Go binding around an Ethereum contract.
type TimeoutManagement struct {
	TimeoutManagementCaller     // Read-only binding to the contract
	TimeoutManagementTransactor // Write-only binding to the contract
	TimeoutManagementFilterer   // Log filterer for contract events
}

// TimeoutManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeoutManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeoutManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeoutManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeoutManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeoutManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeoutManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeoutManagementSession struct {
	Contract     *TimeoutManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TimeoutManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeoutManagementCallerSession struct {
	Contract *TimeoutManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TimeoutManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeoutManagementTransactorSession struct {
	Contract     *TimeoutManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TimeoutManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeoutManagementRaw struct {
	Contract *TimeoutManagement // Generic contract binding to access the raw methods on
}

// TimeoutManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeoutManagementCallerRaw struct {
	Contract *TimeoutManagementCaller // Generic read-only contract binding to access the raw methods on
}

// TimeoutManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeoutManagementTransactorRaw struct {
	Contract *TimeoutManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimeoutManagement creates a new instance of TimeoutManagement, bound to a specific deployed contract.
func NewTimeoutManagement(address common.Address, backend bind.ContractBackend) (*TimeoutManagement, error) {
	contract, err := bindTimeoutManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimeoutManagement{TimeoutManagementCaller: TimeoutManagementCaller{contract: contract}, TimeoutManagementTransactor: TimeoutManagementTransactor{contract: contract}, TimeoutManagementFilterer: TimeoutManagementFilterer{contract: contract}}, nil
}

// NewTimeoutManagementCaller creates a new read-only instance of TimeoutManagement, bound to a specific deployed contract.
func NewTimeoutManagementCaller(address common.Address, caller bind.ContractCaller) (*TimeoutManagementCaller, error) {
	contract, err := bindTimeoutManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeoutManagementCaller{contract: contract}, nil
}

// NewTimeoutManagementTransactor creates a new write-only instance of TimeoutManagement, bound to a specific deployed contract.
func NewTimeoutManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeoutManagementTransactor, error) {
	contract, err := bindTimeoutManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeoutManagementTransactor{contract: contract}, nil
}

// NewTimeoutManagementFilterer creates a new log filterer instance of TimeoutManagement, bound to a specific deployed contract.
func NewTimeoutManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeoutManagementFilterer, error) {
	contract, err := bindTimeoutManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeoutManagementFilterer{contract: contract}, nil
}

// bindTimeoutManagement binds a generic wrapper to an already deployed contract.
func bindTimeoutManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeoutManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeoutManagement *TimeoutManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeoutManagement.Contract.TimeoutManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeoutManagement *TimeoutManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeoutManagement.Contract.TimeoutManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeoutManagement *TimeoutManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeoutManagement.Contract.TimeoutManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeoutManagement *TimeoutManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeoutManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeoutManagement *TimeoutManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeoutManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeoutManagement *TimeoutManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeoutManagement.Contract.contract.Transact(opts, method, params...)
}

// StillAlive is a free data retrieval call binding the contract method 0x15a945ce.
//
// Solidity: function stillAlive(uint256 _blockOrigin, uint256 _timeout) view returns(bool)
func (_TimeoutManagement *TimeoutManagementCaller) StillAlive(opts *bind.CallOpts, _blockOrigin *big.Int, _timeout *big.Int) (bool, error) {
	var out []interface{}
	err := _TimeoutManagement.contract.Call(opts, &out, "stillAlive", _blockOrigin, _timeout)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StillAlive is a free data retrieval call binding the contract method 0x15a945ce.
//
// Solidity: function stillAlive(uint256 _blockOrigin, uint256 _timeout) view returns(bool)
func (_TimeoutManagement *TimeoutManagementSession) StillAlive(_blockOrigin *big.Int, _timeout *big.Int) (bool, error) {
	return _TimeoutManagement.Contract.StillAlive(&_TimeoutManagement.CallOpts, _blockOrigin, _timeout)
}

// StillAlive is a free data retrieval call binding the contract method 0x15a945ce.
//
// Solidity: function stillAlive(uint256 _blockOrigin, uint256 _timeout) view returns(bool)
func (_TimeoutManagement *TimeoutManagementCallerSession) StillAlive(_blockOrigin *big.Int, _timeout *big.Int) (bool, error) {
	return _TimeoutManagement.Contract.StillAlive(&_TimeoutManagement.CallOpts, _blockOrigin, _timeout)
}

// ConsoleMetaData contains all meta data concerning the Console contract.
var ConsoleMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d654291ad24250b03f64650ffb31edab92a812b94caf4e1cc4a3298a906f29b864736f6c63430008110033",
}

// ConsoleABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsoleMetaData.ABI instead.
var ConsoleABI = ConsoleMetaData.ABI

// ConsoleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsoleMetaData.Bin instead.
var ConsoleBin = ConsoleMetaData.Bin

// DeployConsole deploys a new Ethereum contract, binding an instance of Console to it.
func DeployConsole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Console, error) {
	parsed, err := ConsoleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// Console is an auto generated Go binding around an Ethereum contract.
type Console struct {
	ConsoleCaller     // Read-only binding to the contract
	ConsoleTransactor // Write-only binding to the contract
	ConsoleFilterer   // Log filterer for contract events
}

// ConsoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsoleSession struct {
	Contract     *Console          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsoleCallerSession struct {
	Contract *ConsoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ConsoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsoleTransactorSession struct {
	Contract     *ConsoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConsoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsoleRaw struct {
	Contract *Console // Generic contract binding to access the raw methods on
}

// ConsoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsoleCallerRaw struct {
	Contract *ConsoleCaller // Generic read-only contract binding to access the raw methods on
}

// ConsoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsoleTransactorRaw struct {
	Contract *ConsoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsole creates a new instance of Console, bound to a specific deployed contract.
func NewConsole(address common.Address, backend bind.ContractBackend) (*Console, error) {
	contract, err := bindConsole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// NewConsoleCaller creates a new read-only instance of Console, bound to a specific deployed contract.
func NewConsoleCaller(address common.Address, caller bind.ContractCaller) (*ConsoleCaller, error) {
	contract, err := bindConsole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleCaller{contract: contract}, nil
}

// NewConsoleTransactor creates a new write-only instance of Console, bound to a specific deployed contract.
func NewConsoleTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsoleTransactor, error) {
	contract, err := bindConsole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleTransactor{contract: contract}, nil
}

// NewConsoleFilterer creates a new log filterer instance of Console, bound to a specific deployed contract.
func NewConsoleFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsoleFilterer, error) {
	contract, err := bindConsole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsoleFilterer{contract: contract}, nil
}

// bindConsole binds a generic wrapper to an already deployed contract.
func bindConsole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console.Contract.ConsoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.contract.Transact(opts, method, params...)
}
