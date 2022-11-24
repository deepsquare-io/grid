// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jobmanager

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
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Timestamp              *big.Int
	JobName                [32]byte
}

// JobDefinition is an auto generated low-level Go binding around an user-defined struct.
type JobDefinition struct {
	GpuPerTask        uint64
	MemPerCpu         uint64
	CpuPerTask        uint64
	Ntasks            uint64
	BatchLocationHash string
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

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220be095c725f60eb1fdd761cc34620e0b5ff5f510f328d0a14257c0a293b2e9ccf64736f6c63430008110033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
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
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122028f5b1dea628652b41df96b585ea461940b245dbfcf5e5b84cf6b9e3ffe4317c64736f6c63430008110033",
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

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobManagerMetaData contains all meta data concerning the JobManager contract.
var JobManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimJobTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimNextJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobFromId\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"}],\"internalType\":\"structJob\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotJobList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metaQueue\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimFails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimableJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerScheduledJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerStartFails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"setJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateJobsStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallet2JobId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50600060018054906101000a900460ff1615905080801562000044575060018060009054906101000a900460ff1660ff16105b8062000081575062000061306200016360201b62002e0a1760201c565b15801562000080575060018060009054906101000a900460ff1660ff16145b5b620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba906200020d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015620001005760018060016101000a81548160ff0219169083151502179055505b80156200015c5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405162000153919062000289565b60405180910390a15b50620002a6565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620001f5602e8362000186565b9150620002028262000197565b604082019050919050565b600060208201905081810360008301526200022881620001e6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620002716200026b62000265846200022f565b62000246565b62000239565b9050919050565b620002838162000250565b82525050565b6000602082019050620002a0600083018462000278565b92915050565b6156d980620002b66000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80637ff67e2911610104578063b3130fba116100a2578063d547741f11610071578063d547741f14610597578063d6aa37a6146105b3578063e1908676146105d1578063e88fe8ca146105ed576101cf565b8063b3130fba146104fe578063b67644b91461052e578063c58467b01461054a578063d1cee5461461057b576101cf565b806391d14854116100de57806391d1485414610450578063a217fddf14610480578063ade197b41461049e578063aef3276f146104ce576101cf565b80637ff67e29146103fa5780638129fc1c146104165780638fb70f6314610420576101cf565b80632f2ff15d1161017157806342308ee41161014b57806342308ee41461035e5780634609ca501461038f57806346200b6b146103ab57806347ce0a26146103ca576101cf565b80632f2ff15d146102ed57806336568abe1461030957806338ed7cfc14610325576101cf565b80631a3cbef4116101ad5780631a3cbef4146102535780631df4e61014610283578063248a9ca3146102b35780632a242a76146102e3576101cf565b806301ffc9a7146101d45780630f5a11ca14610204578063110e87a614610222575b600080fd5b6101ee60048036038101906101e99190613805565b61061d565b6040516101fb919061384d565b60405180910390f35b61020c610697565b6040516102199190613881565b60405180910390f35b61023c600480360381019061023791906138fa565b61069d565b60405161024a929190613943565b60405180910390f35b61026d600480360381019061026891906138fa565b6106db565b60405161027a9190613a34565b60405180910390f35b61029d60048036038101906102989190613ce6565b610772565b6040516102aa9190613d78565b60405180910390f35b6102cd60048036038101906102c89190613d93565b610bcf565b6040516102da9190613d78565b60405180910390f35b6102eb610bee565b005b61030760048036038101906103029190613dc0565b611088565b005b610323600480360381019061031e9190613dc0565b6110a9565b005b61033f600480360381019061033a9190613d93565b61112c565b6040516103559a99989796959493929190613f8a565b60405180910390f35b610378600480360381019061037391906138fa565b61133e565b604051610386929190613943565b60405180910390f35b6103a960048036038101906103a49190613dc0565b61137c565b005b6103b3611499565b6040516103c1929190613943565b60405180910390f35b6103e460048036038101906103df91906138fa565b6114c5565b6040516103f19190613881565b60405180910390f35b610414600480360381019061040f9190614052565b6114dd565b005b61041e611c8e565b005b61043a60048036038101906104359190613dc0565b611dfb565b60405161044791906140a5565b60405180910390f35b61046a60048036038101906104659190613dc0565b61243a565b604051610477919061384d565b60405180910390f35b6104886124a4565b6040516104959190613d78565b60405180910390f35b6104b860048036038101906104b391906138fa565b6124ab565b6040516104c59190613881565b60405180910390f35b6104e860048036038101906104e391906140c7565b6124c3565b6040516104f59190613d78565b60405180910390f35b61051860048036038101906105139190613d93565b6124e7565b6040516105259190614280565b60405180910390f35b610548600480360381019061054391906142a2565b612801565b005b610564600480360381019061055f91906138fa565b6129d0565b6040516105729291906142e2565b60405180910390f35b61059560048036038101906105909190613dc0565b612ace565b005b6105b160048036038101906105ac9190613dc0565b612cec565b005b6105bb612d0d565b6040516105c89190614321565b60405180910390f35b6105eb60048036038101906105e6919061433c565b612d27565b005b6106076004803603810190610602919061437c565b612dd9565b6040516106149190613d78565b60405180910390f35b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610690575061068f82612e2d565b5b9050919050565b600c5481565b60046020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6060600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561076657602002820191906000526020600020905b815481526020019060010190808311610752575b50505050509050919050565b600080846060015167ffffffffffffffff16116107c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bb90614419565b60405180910390fd5b6000846040015167ffffffffffffffff1611610815576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080c906144ab565b60405180910390fd5b6000846020015167ffffffffffffffff1611610866576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085d90614517565b60405180910390fd5b6108736000801b33612e97565b600c600081548092919061088690614566565b91905055506000600c5460001b9050604051806101400160405280828152602001600060088111156108bb576108ba613e00565b5b81526020018581526020018773ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016001151581526020016000815260200142815260200184815250600260008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff0219169083600881111561096757610966613e00565b5b02179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816001019081610ade91906147ba565b50505060c08201518160070160006101000a81548160ff02191690831515021790555060e08201518160080155610100820151816009015561012082015181600a0155905050600860008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150557f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada38187604051610bbb92919061488c565b60405180910390a180915050949350505050565b6000806000838152602001908152602001600020600101549050919050565b610bfb6000801b33612e97565b60005b6003805490508167ffffffffffffffff16101561108557600060038267ffffffffffffffff1681548110610c3557610c346148b5565b5b906000526020600020015490506000801b8103610c525750611072565b60016008811115610c6657610c65613e00565b5b6002600083815260200190815260200160002060010160009054906101000a900460ff166008811115610c9c57610c9b613e00565b5b14158015610cf2575060026008811115610cb957610cb8613e00565b5b6002600083815260200190815260200160002060010160009054906101000a900460ff166008811115610cef57610cee613e00565b5b14155b15610d2d576000801b60038367ffffffffffffffff1681548110610d1957610d186148b5565b5b906000526020600020018190555050611072565b610d4d81600b60009054906101000a900467ffffffffffffffff16612f1c565b156110705760006002600083815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000801b60038467ffffffffffffffff1681548110610daf57610dae6148b5565b5b90600052602060002001819055505b610e05600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612fb8565b610e99576000610e52600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612fed565b905060006002600083815260200190815260200160002060010160006101000a81548160ff02191690836008811115610e8e57610e8d613e00565b5b021790555050610dbe565b5b610ee1600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612fb8565b610f75576000610f2e600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612fed565b905060006002600083815260200190815260200160002060010160006101000a81548160ff02191690836008811115610f6a57610f69613e00565b5b021790555050610e9a565b60006008811115610f8957610f88613e00565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff166008811115610fbf57610fbe613e00565b5b14610fff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff69061497c565b60405180910390fd5b7f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3826002600085815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161106692919061488c565b60405180910390a1505b505b808061107d9061499c565b915050610bfe565b50565b61109182610bcf565b61109a816130c9565b6110a483836130dd565b505050565b6110b16131bd565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461111e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111590614a3e565b60405180910390fd5b61112882826131c5565b5050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054611292906145dd565b80601f01602080910402602001604051908101604052809291908181526020018280546112be906145dd565b801561130b5780601f106112e05761010080835404028352916020019161130b565b820191906000526020600020905b8154815290600101906020018083116112ee57829003601f168201915b505050505081525050908060070160009054906101000a900460ff169080600801549080600901549080600a015490508a565b60056020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6113896000801b33612e97565b61139382826132a6565b600260088111156113a7576113a6613e00565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff1660088111156113dd576113dc613e00565b5b1461141d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141490614aaa565b60405180910390fd5b60006002600084815260200190815260200160002060010160006101000a81548160ff0219169083600881111561145757611456613e00565b5b02179055507f5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0828260405161148d92919061488c565b60405180910390a15050565b60098060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b60076020528060005260406000206000915090505481565b600060088111156114f1576114f0613e00565b5b81600881111561150457611503613e00565b5b03611544576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161153b90614b3c565b60405180910390fd5b6002600881111561155857611557613e00565b5b81600881111561156b5761156a613e00565b5b036115ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115a290614b3c565b60405180910390fd5b600460088111156115bf576115be613e00565b5b8160088111156115d2576115d1613e00565b5b03611612576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160990614bce565b60405180910390fd5b61161f6000801b33612e97565b61162983836132a6565b6003600881111561163d5761163c613e00565b5b8160088111156116505761164f613e00565b5b03611723576002600881111561166957611668613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff16600881111561169f5761169e613e00565b5b146116df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d690614c60565b60405180910390fd5b60036002600085815260200190815260200160002060010160006101000a81548160ff0219169083600881111561171957611718613e00565b5b0217905550611c00565b6005600881111561173757611736613e00565b5b81600881111561174a57611749613e00565b5b0361181d576004600881111561176357611762613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff16600881111561179957611798613e00565b5b146117d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117d090614cf2565b60405180910390fd5b60056002600085815260200190815260200160002060010160006101000a81548160ff0219169083600881111561181357611812613e00565b5b0217905550611bff565b6006600881111561183157611830613e00565b5b81600881111561184457611843613e00565b5b03611969576003600881111561185d5761185c613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff16600881111561189357611892613e00565b5b14806118e65750600260088111156118ae576118ad613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff1660088111156118e4576118e3613e00565b5b145b611925576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161191c90614d84565b60405180910390fd5b60066002600085815260200190815260200160002060010160006101000a81548160ff0219169083600881111561195f5761195e613e00565b5b0217905550611bfe565b6007600881111561197d5761197c613e00565b5b8160088111156119905761198f613e00565b5b03611ab557600360088111156119a9576119a8613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff1660088111156119df576119de613e00565b5b1480611a325750600260088111156119fa576119f9613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff166008811115611a3057611a2f613e00565b5b145b611a71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a6890614e16565b60405180910390fd5b60076002600085815260200190815260200160002060010160006101000a81548160ff02191690836008811115611aab57611aaa613e00565b5b0217905550611bfd565b600880811115611ac857611ac7613e00565b5b816008811115611adb57611ada613e00565b5b03611bfc5760036008811115611af457611af3613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff166008811115611b2a57611b29613e00565b5b1480611b7d575060026008811115611b4557611b44613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff166008811115611b7b57611b7a613e00565b5b145b611bbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bb390614ea8565b60405180910390fd5b60086002600085815260200190815260200160002060010160006101000a81548160ff02191690836008811115611bf657611bf5613e00565b5b02179055505b5b5b5b5b806008811115611c1357611c12613e00565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff166008811115611c4957611c48613e00565b5b14611c89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c8090614f14565b60405180910390fd5b505050565b600060018054906101000a900460ff16159050808015611cbf575060018060009054906101000a900460ff1660ff16105b80611ced5750611cce30612e0a565b158015611cec575060018060009054906101000a900460ff1660ff16145b5b611d2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d2390614fa6565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015611d685760018060016101000a81548160ff0219169083151502179055505b611d756000801b336130dd565b600a600b60006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508015611df85760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051611def919061500e565b60405180910390a15b50565b611e036136a3565b611e106000801b33612e97565b6002600084815260200190815260200160002060070160009054906101000a900460ff16611e73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e6a90615075565b60405180910390fd5b60006002600085815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff166008811115611ec357611ec2613e00565b5b6008811115611ed557611ed4613e00565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054612077906145dd565b80601f01602080910402602001604051908101604052809291908181526020018280546120a3906145dd565b80156120f05780601f106120c5576101008083540402835291602001916120f0565b820191906000526020600020905b8154815290600101906020018083116120d357829003601f168201915b50505050508152505081526020016007820160009054906101000a900460ff161515151581526020016008820154815260200160098201548152602001600a8201548152505090508273ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16146121aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121a190615107565b60405180910390fd5b600160088111156121be576121bd613e00565b5b816020015160088111156121d5576121d4613e00565b5b14612215576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161220c90615199565b60405180910390fd5b600280600086815260200190815260200160002060010160006101000a81548160ff0219169083600881111561224e5761224d613e00565b5b02179055504360026000868152602001908152602001600020600801819055506122b6600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020856133be565b600260008581526020019081526020016000206005016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546123af906145dd565b80601f01602080910402602001604051908101604052809291908181526020018280546123db906145dd565b80156124285780601f106123fd57610100808354040283529160200191612428565b820191906000526020600020905b81548152906001019060200180831161240b57829003601f168201915b50505050508152505091505092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b60066020528060005260406000206000915090505481565b600381815481106124d357600080fd5b906000526020600020016000915090505481565b6124ef6136fa565b60006002600084815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff16600881111561253f5761253e613e00565b5b600881111561255157612550613e00565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546126f3906145dd565b80601f016020809104026020016040519081016040528092919081815260200182805461271f906145dd565b801561276c5780601f106127415761010080835404028352916020019161276c565b820191906000526020600020905b81548152906001019060200180831161274f57829003601f168201915b50505050508152505081526020016007820160009054906101000a900460ff161515151581526020016008820154815260200160098201548152602001600a8201548152505090508060c001516127f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127ef90615075565b60405180910390fd5b80915050919050565b61280e6000801b33612e97565b6002600083815260200190815260200160002060070160009054906101000a900460ff16612871576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161286890615075565b60405180910390fd5b6000600881111561288557612884613e00565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff1660088111156128bb576128ba613e00565b5b148061290e5750600160088111156128d6576128d5613e00565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600881111561290c5761290b613e00565b5b145b8061296057506002600881111561292857612927613e00565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600881111561295e5761295d613e00565b5b145b61299f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161299690615205565b60405180910390fd5b806002600084815260200190815260200160002060020160008282546129c59190615225565b925050819055505050565b60006129da6136a3565b6129e76000801b33612e97565b612a2e600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612fb8565b15612a6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a65906152a5565b60405180910390fd5b6000612ab7600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020612fed565b905080612ac48286611dfb565b9250925050915091565b612adb6000801b33612e97565b6002600083815260200190815260200160002060070160009054906101000a900460ff16612b3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b3590615311565b60405180910390fd5b60006008811115612b5257612b51613e00565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff166008811115612b8857612b87613e00565b5b14612bc8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bbf9061537d565b60405180910390fd5b806002600084815260200190815260200160002060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016002600084815260200190815260200160002060010160006101000a81548160ff02191690836008811115612c5757612c56613e00565b5b02179055504360026000848152602001908152602001600020600801819055506003829080600181540180825580915050600190039060005260206000200160009091909190915055612ce8600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020836133be565b5050565b612cf582610bcf565b612cfe816130c9565b612d0883836131c5565b505050565b600b60009054906101000a900467ffffffffffffffff1681565b612d346000801b33612e97565b6002600083815260200190815260200160002060070160009054906101000a900460ff16612d97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d8e90615075565b60405180910390fd5b806002600084815260200190815260200160002060010160006101000a81548160ff02191690836008811115612dd057612dcf613e00565b5b02179055505050565b60086020528160005260406000208181548110612df557600080fd5b90600052602060002001600091509150505481565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b612ea1828261243a565b612f1857612eae8161343a565b612ebc8360001c6020613467565b604051602001612ecd929190615471565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f0f91906154e4565b60405180910390fd5b5050565b60006002600084815260200190815260200160002060070160009054906101000a900460ff16612f81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f7890615075565b60405180910390fd5b8167ffffffffffffffff16600260008581526020019081526020016000206008015443612fae9190615506565b1015905092915050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b13159050919050565b6000612ff882612fb8565b1561302f576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260000160009054906101000a9004600f0b905082600101600082600f0b600f0b815260200190815260200160002054915082600101600082600f0b600f0b815260200190815260200160002060009055600181018360000160006101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff16021790555050919050565b6130da816130d56131bd565b612e97565b50565b6130e7828261243a565b6131b957600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061315e6131bd565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6131cf828261243a565b156132a257600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506132476131bd565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6132b36000801b33612e97565b6002600083815260200190815260200160002060070160009054906101000a900460ff16613316576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161330d90615075565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146133ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133b1906155ac565b60405180910390fd5b5050565b60008260000160109054906101000a9004600f0b90508183600101600083600f0b600f0b815260200190815260200160002081905550600181018360000160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff160217905550505050565b60606134608273ffffffffffffffffffffffffffffffffffffffff16601460ff16613467565b9050919050565b60606000600283600261347a91906155cc565b6134849190615225565b67ffffffffffffffff81111561349d5761349c613a6c565b5b6040519080825280601f01601f1916602001820160405280156134cf5781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110613507576135066148b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061356b5761356a6148b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026135ab91906155cc565b6135b59190615225565b90505b6001811115613655577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106135f7576135f66148b5565b5b1a60f81b82828151811061360e5761360d6148b5565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061364e9061560e565b90506135b8565b5060008414613699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161369090615683565b60405180910390fd5b8091505092915050565b6040518060a00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b604051806101400160405280600080191681526020016000600881111561372457613723613e00565b5b815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016137726136a3565b81526020016000151581526020016000815260200160008152602001600080191681525090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6137e2816137ad565b81146137ed57600080fd5b50565b6000813590506137ff816137d9565b92915050565b60006020828403121561381b5761381a6137a3565b5b6000613829848285016137f0565b91505092915050565b60008115159050919050565b61384781613832565b82525050565b6000602082019050613862600083018461383e565b92915050565b6000819050919050565b61387b81613868565b82525050565b60006020820190506138966000830184613872565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006138c78261389c565b9050919050565b6138d7816138bc565b81146138e257600080fd5b50565b6000813590506138f4816138ce565b92915050565b6000602082840312156139105761390f6137a3565b5b600061391e848285016138e5565b91505092915050565b600081600f0b9050919050565b61393d81613927565b82525050565b60006040820190506139586000830185613934565b6139656020830184613934565b9392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b6139ab81613998565b82525050565b60006139bd83836139a2565b60208301905092915050565b6000602082019050919050565b60006139e18261396c565b6139eb8185613977565b93506139f683613988565b8060005b83811015613a27578151613a0e88826139b1565b9750613a19836139c9565b9250506001810190506139fa565b5085935050505092915050565b60006020820190508181036000830152613a4e81846139d6565b905092915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b613aa482613a5b565b810181811067ffffffffffffffff82111715613ac357613ac2613a6c565b5b80604052505050565b6000613ad6613799565b9050613ae28282613a9b565b919050565b600080fd5b600067ffffffffffffffff82169050919050565b613b0981613aec565b8114613b1457600080fd5b50565b600081359050613b2681613b00565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff821115613b5157613b50613a6c565b5b613b5a82613a5b565b9050602081019050919050565b82818337600083830152505050565b6000613b89613b8484613b36565b613acc565b905082815260208101848484011115613ba557613ba4613b31565b5b613bb0848285613b67565b509392505050565b600082601f830112613bcd57613bcc613b2c565b5b8135613bdd848260208601613b76565b91505092915050565b600060a08284031215613bfc57613bfb613a56565b5b613c0660a0613acc565b90506000613c1684828501613b17565b6000830152506020613c2a84828501613b17565b6020830152506040613c3e84828501613b17565b6040830152506060613c5284828501613b17565b606083015250608082013567ffffffffffffffff811115613c7657613c75613ae7565b5b613c8284828501613bb8565b60808301525092915050565b613c9781613868565b8114613ca257600080fd5b50565b600081359050613cb481613c8e565b92915050565b613cc381613998565b8114613cce57600080fd5b50565b600081359050613ce081613cba565b92915050565b60008060008060808587031215613d0057613cff6137a3565b5b6000613d0e878288016138e5565b945050602085013567ffffffffffffffff811115613d2f57613d2e6137a8565b5b613d3b87828801613be6565b9350506040613d4c87828801613ca5565b9250506060613d5d87828801613cd1565b91505092959194509250565b613d7281613998565b82525050565b6000602082019050613d8d6000830184613d69565b92915050565b600060208284031215613da957613da86137a3565b5b6000613db784828501613cd1565b91505092915050565b60008060408385031215613dd757613dd66137a3565b5b6000613de585828601613cd1565b9250506020613df6858286016138e5565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60098110613e4057613e3f613e00565b5b50565b6000819050613e5182613e2f565b919050565b6000613e6182613e43565b9050919050565b613e7181613e56565b82525050565b613e80816138bc565b82525050565b613e8f81613aec565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613ecf578082015181840152602081019050613eb4565b60008484015250505050565b6000613ee682613e95565b613ef08185613ea0565b9350613f00818560208601613eb1565b613f0981613a5b565b840191505092915050565b600060a083016000830151613f2c6000860182613e86565b506020830151613f3f6020860182613e86565b506040830151613f526040860182613e86565b506060830151613f656060860182613e86565b5060808301518482036080860152613f7d8282613edb565b9150508091505092915050565b600061014082019050613fa0600083018d613d69565b613fad602083018c613e68565b613fba604083018b613872565b613fc7606083018a613e77565b613fd46080830189613e77565b81810360a0830152613fe68188613f14565b9050613ff560c083018761383e565b61400260e0830186613872565b614010610100830185613872565b61401e610120830184613d69565b9b9a5050505050505050505050565b6009811061403a57600080fd5b50565b60008135905061404c8161402d565b92915050565b60008060006060848603121561406b5761406a6137a3565b5b600061407986828701613cd1565b935050602061408a868287016138e5565b925050604061409b8682870161403d565b9150509250925092565b600060208201905081810360008301526140bf8184613f14565b905092915050565b6000602082840312156140dd576140dc6137a3565b5b60006140eb84828501613ca5565b91505092915050565b6140fd81613e56565b82525050565b61410c81613868565b82525050565b61411b816138bc565b82525050565b600060a0830160008301516141396000860182613e86565b50602083015161414c6020860182613e86565b50604083015161415f6040860182613e86565b5060608301516141726060860182613e86565b506080830151848203608086015261418a8282613edb565b9150508091505092915050565b6141a081613832565b82525050565b6000610140830160008301516141bf60008601826139a2565b5060208301516141d260208601826140f4565b5060408301516141e56040860182614103565b5060608301516141f86060860182614112565b50608083015161420b6080860182614112565b5060a083015184820360a08601526142238282614121565b91505060c083015161423860c0860182614197565b5060e083015161424b60e0860182614103565b50610100830151614260610100860182614103565b506101208301516142756101208601826139a2565b508091505092915050565b6000602082019050818103600083015261429a81846141a6565b905092915050565b600080604083850312156142b9576142b86137a3565b5b60006142c785828601613cd1565b92505060206142d885828601613ca5565b9150509250929050565b60006040820190506142f76000830185613d69565b81810360208301526143098184613f14565b90509392505050565b61431b81613aec565b82525050565b60006020820190506143366000830184614312565b92915050565b60008060408385031215614353576143526137a3565b5b600061436185828601613cd1565b92505060206143728582860161403d565b9150509250929050565b60008060408385031215614393576143926137a3565b5b60006143a1858286016138e5565b92505060206143b285828601613ca5565b9150509250929050565b600082825260208201905092915050565b7f6e7461736b73206d7573742062652067726561746572207468616e2030000000600082015250565b6000614403601d836143bc565b915061440e826143cd565b602082019050919050565b60006020820190508181036000830152614432816143f6565b9050919050565b7f6370755065725461736b206d7573742062652067726561746572207468616e2060008201527f3000000000000000000000000000000000000000000000000000000000000000602082015250565b60006144956021836143bc565b91506144a082614439565b604082019050919050565b600060208201905081810360008301526144c481614488565b9050919050565b7f6d656d506572437075206d7573742062652067726561746572207468616e2030600082015250565b60006145016020836143bc565b915061450c826144cb565b602082019050919050565b60006020820190508181036000830152614530816144f4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061457182613868565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036145a3576145a2614537565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806145f557607f821691505b602082108103614608576146076145ae565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026146707fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614633565b61467a8683614633565b95508019841693508086168417925050509392505050565b6000819050919050565b60006146b76146b26146ad84613868565b614692565b613868565b9050919050565b6000819050919050565b6146d18361469c565b6146e56146dd826146be565b848454614640565b825550505050565b600090565b6146fa6146ed565b6147058184846146c8565b505050565b5b818110156147295761471e6000826146f2565b60018101905061470b565b5050565b601f82111561476e5761473f8161460e565b61474884614623565b81016020851015614757578190505b61476b61476385614623565b83018261470a565b50505b505050565b600082821c905092915050565b600061479160001984600802614773565b1980831691505092915050565b60006147aa8383614780565b9150826002028217905092915050565b6147c382613e95565b67ffffffffffffffff8111156147dc576147db613a6c565b5b6147e682546145dd565b6147f182828561472d565b600060209050601f8311600181146148245760008415614812578287015190505b61481c858261479e565b865550614884565b601f1984166148328661460e565b60005b8281101561485a57848901518255600182019150602085019450602081019050614835565b868310156148775784890151614873601f891682614780565b8355505b6001600288020188555050505b505050505050565b60006040820190506148a16000830185613d69565b6148ae6020830184613e77565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f54686973206a6f62206d7573742062652050454e44494e47207468657265206960008201527f732061206d69736d61746368206265747765656e2070726f76696465724a6f6260208201527f7320616e6420686f744a6f624c69737400000000000000000000000000000000604082015250565b60006149666050836143bc565b9150614971826148e4565b606082019050919050565b6000602082019050818103600083015261499581614959565b9050919050565b60006149a782613aec565b915067ffffffffffffffff82036149c1576149c0614537565b5b600182019050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000614a28602f836143bc565b9150614a33826149cc565b604082019050919050565b60006020820190508181036000830152614a5781614a1b565b9050919050565b7f4f6e6c7920717565756564206a6f622063616e20626520737461727465640000600082015250565b6000614a94601e836143bc565b9150614a9f82614a5e565b602082019050919050565b60006020820190508181036000830152614ac381614a87565b9050919050565b7f43616e6e6f74207075742061206a6f6220696e2070656e64696e67207374617460008201527f7573000000000000000000000000000000000000000000000000000000000000602082015250565b6000614b266022836143bc565b9150614b3182614aca565b604082019050919050565b60006020820190508181036000830152614b5581614b19565b9050919050565b7f43616e6e6f74207075742061206a6f6220696e2063616e63656c6c696e67207360008201527f7461747573000000000000000000000000000000000000000000000000000000602082015250565b6000614bb86025836143bc565b9150614bc382614b5c565b604082019050919050565b60006020820190508181036000830152614be781614bab565b9050919050565b7f4f6e6c79207363686564756c6564206a6f62732063616e20626520737461727460008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b6000614c4a6022836143bc565b9150614c5582614bee565b604082019050919050565b60006020820190508181036000830152614c7981614c3d565b9050919050565b7f4f6e6c792063616e63656c6c696e67206a6f62732063616e2062652063616e6360008201527f656c6c6564000000000000000000000000000000000000000000000000000000602082015250565b6000614cdc6025836143bc565b9150614ce782614c80565b604082019050919050565b60006020820190508181036000830152614d0b81614ccf565b9050919050565b7f4f6e6c792072756e6e696e6720616e64207363686564756c6564206a6f62732060008201527f63616e2062652066696e69736865640000000000000000000000000000000000602082015250565b6000614d6e602f836143bc565b9150614d7982614d12565b604082019050919050565b60006020820190508181036000830152614d9d81614d61565b9050919050565b7f4f6e6c79207363686564756c656420616e642072756e6e696e67206a6f62206360008201527f616e206661696c00000000000000000000000000000000000000000000000000602082015250565b6000614e006027836143bc565b9150614e0b82614da4565b604082019050919050565b60006020820190508181036000830152614e2f81614df3565b9050919050565b7f4f6e6c79207363686564756c656420616e642072756e6e696e67206a6f62732060008201527f63616e2072756e206f7574206f66206d6f6e6579000000000000000000000000602082015250565b6000614e926034836143bc565b9150614e9d82614e36565b604082019050919050565b60006020820190508181036000830152614ec181614e85565b9050919050565b7f4661696c656420746f20757064617465206a6f62207374617475730000000000600082015250565b6000614efe601b836143bc565b9150614f0982614ec8565b602082019050919050565b60006020820190508181036000830152614f2d81614ef1565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000614f90602e836143bc565b9150614f9b82614f34565b604082019050919050565b60006020820190508181036000830152614fbf81614f83565b9050919050565b6000819050919050565b600060ff82169050919050565b6000614ff8614ff3614fee84614fc6565b614692565b614fd0565b9050919050565b61500881614fdd565b82525050565b60006020820190506150236000830184614fff565b92915050565b7f4a6f62206e6f7420666f756e6400000000000000000000000000000000000000600082015250565b600061505f600d836143bc565b915061506a82615029565b602082019050919050565b6000602082019050818103600083015261508e81615052565b9050919050565b7f50726f7669646572732063616e206f6e6c7920636c61696d207468656972206a60008201527f6f62730000000000000000000000000000000000000000000000000000000000602082015250565b60006150f16023836143bc565b91506150fc82615095565b604082019050919050565b60006020820190508181036000830152615120816150e4565b9050919050565b7f4f6e6c79206d6574612d717565756564206a6f622063616e206265207175657560008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b60006151836022836143bc565b915061518e82615127565b604082019050919050565b600060208201905081810360008301526151b281615176565b9050919050565b7f4a6f62206d7573742062652072756e6e696e6700000000000000000000000000600082015250565b60006151ef6013836143bc565b91506151fa826151b9565b602082019050919050565b6000602082019050818103600083015261521e816151e2565b9050919050565b600061523082613868565b915061523b83613868565b925082820190508082111561525357615252614537565b5b92915050565b7f4e6f20617661696c61626c65206a6f6200000000000000000000000000000000600082015250565b600061528f6010836143bc565b915061529a82615259565b602082019050919050565b600060208201905081810360008301526152be81615282565b9050919050565b7f4a6f62206d7573742062652076616c6964000000000000000000000000000000600082015250565b60006152fb6011836143bc565b9150615306826152c5565b602082019050919050565b6000602082019050818103600083015261532a816152ee565b9050919050565b7f4a6f62206d7573742062652050454e44494e4700000000000000000000000000600082015250565b60006153676013836143bc565b915061537282615331565b602082019050919050565b600060208201905081810360008301526153968161535a565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b60006153de60178361539d565b91506153e9826153a8565b601782019050919050565b60006153ff82613e95565b615409818561539d565b9350615419818560208601613eb1565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b600061545b60118361539d565b915061546682615425565b601182019050919050565b600061547c826153d1565b915061548882856153f4565b91506154938261544e565b915061549f82846153f4565b91508190509392505050565b60006154b682613e95565b6154c081856143bc565b93506154d0818560208601613eb1565b6154d981613a5b565b840191505092915050565b600060208201905081810360008301526154fe81846154ab565b905092915050565b600061551182613868565b915061551c83613868565b925082820390508181111561553457615533614537565b5b92915050565b7f50726f7669646572732063616e206f6e6c79207374617274207468656972206a60008201527f6f62730000000000000000000000000000000000000000000000000000000000602082015250565b60006155966023836143bc565b91506155a18261553a565b604082019050919050565b600060208201905081810360008301526155c581615589565b9050919050565b60006155d782613868565b91506155e283613868565b92508282026155f081613868565b9150828204841483151761560757615606614537565b5b5092915050565b600061561982613868565b91506000820361562c5761562b614537565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b600061566d6020836143bc565b915061567882615637565b602082019050919050565b6000602082019050818103600083015261569c81615660565b905091905056fea26469706673582212205de4cc2bd39c8a2d93b51e271ceba4c4629ef69b2c1dc19284dbfe1bcbf3c3c464736f6c63430008110033",
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

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_JobManager *JobManagerCaller) ClaimJobTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "claimJobTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_JobManager *JobManagerSession) ClaimJobTimeout() (uint64, error) {
	return _JobManager.Contract.ClaimJobTimeout(&_JobManager.CallOpts)
}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_JobManager *JobManagerCallerSession) ClaimJobTimeout() (uint64, error) {
	return _JobManager.Contract.ClaimJobTimeout(&_JobManager.CallOpts)
}

// GetJobFromId is a free data retrieval call binding the contract method 0xb3130fba.
//
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,string),bool,uint256,uint256,bytes32))
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
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,string),bool,uint256,uint256,bytes32))
func (_JobManager *JobManagerSession) GetJobFromId(_jobId [32]byte) (Job, error) {
	return _JobManager.Contract.GetJobFromId(&_JobManager.CallOpts, _jobId)
}

// GetJobFromId is a free data retrieval call binding the contract method 0xb3130fba.
//
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,string),bool,uint256,uint256,bytes32))
func (_JobManager *JobManagerCallerSession) GetJobFromId(_jobId [32]byte) (Job, error) {
	return _JobManager.Contract.GetJobFromId(&_JobManager.CallOpts, _jobId)
}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address walletAddr) view returns(bytes32[])
func (_JobManager *JobManagerCaller) GetJobs(opts *bind.CallOpts, walletAddr common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "getJobs", walletAddr)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address walletAddr) view returns(bytes32[])
func (_JobManager *JobManagerSession) GetJobs(walletAddr common.Address) ([][32]byte, error) {
	return _JobManager.Contract.GetJobs(&_JobManager.CallOpts, walletAddr)
}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address walletAddr) view returns(bytes32[])
func (_JobManager *JobManagerCallerSession) GetJobs(walletAddr common.Address) ([][32]byte, error) {
	return _JobManager.Contract.GetJobs(&_JobManager.CallOpts, walletAddr)
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

// HotJobList is a free data retrieval call binding the contract method 0xaef3276f.
//
// Solidity: function hotJobList(uint256 ) view returns(bytes32)
func (_JobManager *JobManagerCaller) HotJobList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "hotJobList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HotJobList is a free data retrieval call binding the contract method 0xaef3276f.
//
// Solidity: function hotJobList(uint256 ) view returns(bytes32)
func (_JobManager *JobManagerSession) HotJobList(arg0 *big.Int) ([32]byte, error) {
	return _JobManager.Contract.HotJobList(&_JobManager.CallOpts, arg0)
}

// HotJobList is a free data retrieval call binding the contract method 0xaef3276f.
//
// Solidity: function hotJobList(uint256 ) view returns(bytes32)
func (_JobManager *JobManagerCallerSession) HotJobList(arg0 *big.Int) ([32]byte, error) {
	return _JobManager.Contract.HotJobList(&_JobManager.CallOpts, arg0)
}

// JobIdCounter is a free data retrieval call binding the contract method 0x0f5a11ca.
//
// Solidity: function jobIdCounter() view returns(uint256)
func (_JobManager *JobManagerCaller) JobIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "jobIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JobIdCounter is a free data retrieval call binding the contract method 0x0f5a11ca.
//
// Solidity: function jobIdCounter() view returns(uint256)
func (_JobManager *JobManagerSession) JobIdCounter() (*big.Int, error) {
	return _JobManager.Contract.JobIdCounter(&_JobManager.CallOpts)
}

// JobIdCounter is a free data retrieval call binding the contract method 0x0f5a11ca.
//
// Solidity: function jobIdCounter() view returns(uint256)
func (_JobManager *JobManagerCallerSession) JobIdCounter() (*big.Int, error) {
	return _JobManager.Contract.JobIdCounter(&_JobManager.CallOpts)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange, uint256 timestamp, bytes32 _jobName)
func (_JobManager *JobManagerCaller) Jobs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Timestamp              *big.Int
	JobName                [32]byte
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		JobId                  [32]byte
		Status                 uint8
		AmountLocked           *big.Int
		CustomerAddr           common.Address
		ProviderAddr           common.Address
		Definition             JobDefinition
		Valid                  bool
		BlockNumberStateChange *big.Int
		Timestamp              *big.Int
		JobName                [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.JobId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.AmountLocked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CustomerAddr = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ProviderAddr = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Definition = *abi.ConvertType(out[5], new(JobDefinition)).(*JobDefinition)
	outstruct.Valid = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.BlockNumberStateChange = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.JobName = *abi.ConvertType(out[9], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange, uint256 timestamp, bytes32 _jobName)
func (_JobManager *JobManagerSession) Jobs(arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Timestamp              *big.Int
	JobName                [32]byte
}, error) {
	return _JobManager.Contract.Jobs(&_JobManager.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange, uint256 timestamp, bytes32 _jobName)
func (_JobManager *JobManagerCallerSession) Jobs(arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Timestamp              *big.Int
	JobName                [32]byte
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

// ProviderClaimFails is a free data retrieval call binding the contract method 0xade197b4.
//
// Solidity: function providerClaimFails(address ) view returns(uint256)
func (_JobManager *JobManagerCaller) ProviderClaimFails(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "providerClaimFails", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderClaimFails is a free data retrieval call binding the contract method 0xade197b4.
//
// Solidity: function providerClaimFails(address ) view returns(uint256)
func (_JobManager *JobManagerSession) ProviderClaimFails(arg0 common.Address) (*big.Int, error) {
	return _JobManager.Contract.ProviderClaimFails(&_JobManager.CallOpts, arg0)
}

// ProviderClaimFails is a free data retrieval call binding the contract method 0xade197b4.
//
// Solidity: function providerClaimFails(address ) view returns(uint256)
func (_JobManager *JobManagerCallerSession) ProviderClaimFails(arg0 common.Address) (*big.Int, error) {
	return _JobManager.Contract.ProviderClaimFails(&_JobManager.CallOpts, arg0)
}

// ProviderClaimableJobsQueues is a free data retrieval call binding the contract method 0x110e87a6.
//
// Solidity: function providerClaimableJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCaller) ProviderClaimableJobsQueues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "providerClaimableJobsQueues", arg0)

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

// ProviderClaimableJobsQueues is a free data retrieval call binding the contract method 0x110e87a6.
//
// Solidity: function providerClaimableJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerSession) ProviderClaimableJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderClaimableJobsQueues(&_JobManager.CallOpts, arg0)
}

// ProviderClaimableJobsQueues is a free data retrieval call binding the contract method 0x110e87a6.
//
// Solidity: function providerClaimableJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCallerSession) ProviderClaimableJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderClaimableJobsQueues(&_JobManager.CallOpts, arg0)
}

// ProviderScheduledJobsQueues is a free data retrieval call binding the contract method 0x42308ee4.
//
// Solidity: function providerScheduledJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCaller) ProviderScheduledJobsQueues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "providerScheduledJobsQueues", arg0)

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

// ProviderScheduledJobsQueues is a free data retrieval call binding the contract method 0x42308ee4.
//
// Solidity: function providerScheduledJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerSession) ProviderScheduledJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderScheduledJobsQueues(&_JobManager.CallOpts, arg0)
}

// ProviderScheduledJobsQueues is a free data retrieval call binding the contract method 0x42308ee4.
//
// Solidity: function providerScheduledJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCallerSession) ProviderScheduledJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderScheduledJobsQueues(&_JobManager.CallOpts, arg0)
}

// ProviderStartFails is a free data retrieval call binding the contract method 0x47ce0a26.
//
// Solidity: function providerStartFails(address ) view returns(uint256)
func (_JobManager *JobManagerCaller) ProviderStartFails(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "providerStartFails", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderStartFails is a free data retrieval call binding the contract method 0x47ce0a26.
//
// Solidity: function providerStartFails(address ) view returns(uint256)
func (_JobManager *JobManagerSession) ProviderStartFails(arg0 common.Address) (*big.Int, error) {
	return _JobManager.Contract.ProviderStartFails(&_JobManager.CallOpts, arg0)
}

// ProviderStartFails is a free data retrieval call binding the contract method 0x47ce0a26.
//
// Solidity: function providerStartFails(address ) view returns(uint256)
func (_JobManager *JobManagerCallerSession) ProviderStartFails(arg0 common.Address) (*big.Int, error) {
	return _JobManager.Contract.ProviderStartFails(&_JobManager.CallOpts, arg0)
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

// Wallet2JobId is a free data retrieval call binding the contract method 0xe88fe8ca.
//
// Solidity: function wallet2JobId(address , uint256 ) view returns(bytes32)
func (_JobManager *JobManagerCaller) Wallet2JobId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "wallet2JobId", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Wallet2JobId is a free data retrieval call binding the contract method 0xe88fe8ca.
//
// Solidity: function wallet2JobId(address , uint256 ) view returns(bytes32)
func (_JobManager *JobManagerSession) Wallet2JobId(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _JobManager.Contract.Wallet2JobId(&_JobManager.CallOpts, arg0, arg1)
}

// Wallet2JobId is a free data retrieval call binding the contract method 0xe88fe8ca.
//
// Solidity: function wallet2JobId(address , uint256 ) view returns(bytes32)
func (_JobManager *JobManagerCallerSession) Wallet2JobId(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _JobManager.Contract.Wallet2JobId(&_JobManager.CallOpts, arg0, arg1)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns((uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactor) ClaimJob(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "claimJob", _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns((uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns((uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactorSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimJob(&_JobManager.TransactOpts, _jobId, _providerAddr)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0xc58467b0.
//
// Solidity: function claimNextJob(address _providerAddr) returns(bytes32 jobId, (uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactor) ClaimNextJob(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "claimNextJob", _providerAddr)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0xc58467b0.
//
// Solidity: function claimNextJob(address _providerAddr) returns(bytes32 jobId, (uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerSession) ClaimNextJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimNextJob(&_JobManager.TransactOpts, _providerAddr)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0xc58467b0.
//
// Solidity: function claimNextJob(address _providerAddr) returns(bytes32 jobId, (uint64,uint64,uint64,uint64,string))
func (_JobManager *JobManagerTransactorSession) ClaimNextJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimNextJob(&_JobManager.TransactOpts, _providerAddr)
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

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_JobManager *JobManagerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_JobManager *JobManagerSession) Initialize() (*types.Transaction, error) {
	return _JobManager.Contract.Initialize(&_JobManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_JobManager *JobManagerTransactorSession) Initialize() (*types.Transaction, error) {
	return _JobManager.Contract.Initialize(&_JobManager.TransactOpts)
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

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x7ff67e29.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, address _providerAddr, uint8 _jobStatus) returns()
func (_JobManager *JobManagerTransactor) ProviderSetJobStatus(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address, _jobStatus uint8) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "providerSetJobStatus", _jobId, _providerAddr, _jobStatus)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x7ff67e29.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, address _providerAddr, uint8 _jobStatus) returns()
func (_JobManager *JobManagerSession) ProviderSetJobStatus(_jobId [32]byte, _providerAddr common.Address, _jobStatus uint8) (*types.Transaction, error) {
	return _JobManager.Contract.ProviderSetJobStatus(&_JobManager.TransactOpts, _jobId, _providerAddr, _jobStatus)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x7ff67e29.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, address _providerAddr, uint8 _jobStatus) returns()
func (_JobManager *JobManagerTransactorSession) ProviderSetJobStatus(_jobId [32]byte, _providerAddr common.Address, _jobStatus uint8) (*types.Transaction, error) {
	return _JobManager.Contract.ProviderSetJobStatus(&_JobManager.TransactOpts, _jobId, _providerAddr, _jobStatus)
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

// RequestNewJob is a paid mutator transaction binding the contract method 0x1df4e610.
//
// Solidity: function requestNewJob(address _customerAddr, (uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked, bytes32 _jobName) returns(bytes32)
func (_JobManager *JobManagerTransactor) RequestNewJob(opts *bind.TransactOpts, _customerAddr common.Address, _definition JobDefinition, _amountLocked *big.Int, _jobName [32]byte) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "requestNewJob", _customerAddr, _definition, _amountLocked, _jobName)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x1df4e610.
//
// Solidity: function requestNewJob(address _customerAddr, (uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked, bytes32 _jobName) returns(bytes32)
func (_JobManager *JobManagerSession) RequestNewJob(_customerAddr common.Address, _definition JobDefinition, _amountLocked *big.Int, _jobName [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.RequestNewJob(&_JobManager.TransactOpts, _customerAddr, _definition, _amountLocked, _jobName)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x1df4e610.
//
// Solidity: function requestNewJob(address _customerAddr, (uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked, bytes32 _jobName) returns(bytes32)
func (_JobManager *JobManagerTransactorSession) RequestNewJob(_customerAddr common.Address, _definition JobDefinition, _amountLocked *big.Int, _jobName [32]byte) (*types.Transaction, error) {
	return _JobManager.Contract.RequestNewJob(&_JobManager.TransactOpts, _customerAddr, _definition, _amountLocked, _jobName)
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

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobId, uint8 _jobStatus) returns()
func (_JobManager *JobManagerTransactor) SetJobStatus(opts *bind.TransactOpts, _jobId [32]byte, _jobStatus uint8) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "setJobStatus", _jobId, _jobStatus)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobId, uint8 _jobStatus) returns()
func (_JobManager *JobManagerSession) SetJobStatus(_jobId [32]byte, _jobStatus uint8) (*types.Transaction, error) {
	return _JobManager.Contract.SetJobStatus(&_JobManager.TransactOpts, _jobId, _jobStatus)
}

// SetJobStatus is a paid mutator transaction binding the contract method 0xe1908676.
//
// Solidity: function setJobStatus(bytes32 _jobId, uint8 _jobStatus) returns()
func (_JobManager *JobManagerTransactorSession) SetJobStatus(_jobId [32]byte, _jobStatus uint8) (*types.Transaction, error) {
	return _JobManager.Contract.SetJobStatus(&_JobManager.TransactOpts, _jobId, _jobStatus)
}

// TopUp is a paid mutator transaction binding the contract method 0xb67644b9.
//
// Solidity: function topUp(bytes32 _jobId, uint256 _amount) returns()
func (_JobManager *JobManagerTransactor) TopUp(opts *bind.TransactOpts, _jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "topUp", _jobId, _amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xb67644b9.
//
// Solidity: function topUp(bytes32 _jobId, uint256 _amount) returns()
func (_JobManager *JobManagerSession) TopUp(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _JobManager.Contract.TopUp(&_JobManager.TransactOpts, _jobId, _amount)
}

// TopUp is a paid mutator transaction binding the contract method 0xb67644b9.
//
// Solidity: function topUp(bytes32 _jobId, uint256 _amount) returns()
func (_JobManager *JobManagerTransactorSession) TopUp(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _JobManager.Contract.TopUp(&_JobManager.TransactOpts, _jobId, _amount)
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

// JobManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the JobManager contract.
type JobManagerInitializedIterator struct {
	Event *JobManagerInitialized // Event containing the contract specifics and raw log

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
func (it *JobManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobManagerInitialized)
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
		it.Event = new(JobManagerInitialized)
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
func (it *JobManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobManagerInitialized represents a Initialized event raised by the JobManager contract.
type JobManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_JobManager *JobManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*JobManagerInitializedIterator, error) {

	logs, sub, err := _JobManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JobManagerInitializedIterator{contract: _JobManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_JobManager *JobManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JobManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _JobManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobManagerInitialized)
				if err := _JobManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_JobManager *JobManagerFilterer) ParseInitialized(log types.Log) (*JobManagerInitialized, error) {
	event := new(JobManagerInitialized)
	if err := _JobManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200deb2d155cc73fbf04f3651b35566c36121307684f46e0e4f511e6d6a66133ad64736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208e68354f14c9d7678f10a06961d6826f3c027320cf756c8ebb9e5426777ee2a864736f6c63430008110033",
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
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205377adc79bb987de03049c655529acbf51a3f7d36bb14c89a2d2f790fee54d6064736f6c63430008110033",
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

// TimeoutManagementMetaData contains all meta data concerning the TimeoutManagement contract.
var TimeoutManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockOrigin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"stillAlive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101cd610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806315a945ce1461003a575b600080fd5b610054600480360381019061004f91906100be565b61006a565b6040516100619190610119565b60405180910390f35b60008183436100799190610163565b1015905092915050565b600080fd5b6000819050919050565b61009b81610088565b81146100a657600080fd5b50565b6000813590506100b881610092565b92915050565b600080604083850312156100d5576100d4610083565b5b60006100e3858286016100a9565b92505060206100f4858286016100a9565b9150509250929050565b60008115159050919050565b610113816100fe565b82525050565b600060208201905061012e600083018461010a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061016e82610088565b915061017983610088565b925082820390508181111561019157610190610134565b5b9291505056fea2646970667358221220f0d73e2d15874b4eb4fc286b938b837d2e1b6a279aeff5c5729ffd1ca92caaea64736f6c63430008110033",
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
