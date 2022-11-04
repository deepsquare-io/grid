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
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
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

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208107dcb809aaf5fe8a0dbd8dc2252420e6e5d9a0344c1556513b6a69b32ec34e64736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimJobTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimNextJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"finishJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobFromId\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"schedulable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"}],\"internalType\":\"structJob\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotJobList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"schedulable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metaQueue\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimFails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimableJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerScheduledJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerStartFails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"startJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"triggerFailed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateJobsStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallet2JobId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50600060018054906101000a900460ff1615905080801562000044575060018060009054906101000a900460ff1660ff16105b8062000081575062000061306200016360201b62002e4c1760201c565b15801562000080575060018060009054906101000a900460ff1660ff16145b5b620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba906200020d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015620001005760018060016101000a81548160ff0219169083151502179055505b80156200015c5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405162000153919062000289565b60405180910390a15b50620002a6565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620001f5602e8362000186565b9150620002028262000197565b604082019050919050565b600060208201905081810360008301526200022881620001e6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620002716200026b62000265846200022f565b62000246565b62000239565b9050919050565b620002838162000250565b82525050565b6000602082019050620002a0600083018462000278565b92915050565b61528480620002b66000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c80638129fc1c11610104578063b67644b9116100a2578063d547741f11610071578063d547741f146105bb578063d6aa37a6146105d7578063e3401e00146105f5578063e88fe8ca14610611576101da565b8063b67644b914610536578063c4d252f514610552578063c58467b01461056e578063d1cee5461461059f576101da565b8063a217fddf116100de578063a217fddf14610488578063ade197b4146104a6578063aef3276f146104d6578063b3130fba14610506576101da565b80638129fc1c1461041e5780638fb70f631461042857806391d1485414610458576101da565b806336568abe1161017c57806346200b6b1161014b57806346200b6b1461039757806347ce0a26146103b65780635792edd4146103e65780637d17544b14610402576101da565b806336568abe146102f657806338ed7cfc1461031257806342308ee41461034a5780634609ca501461037b576101da565b80631a3cbef4116101b85780631a3cbef414610270578063248a9ca3146102a05780632a242a76146102d05780632f2ff15d146102da576101da565b8063010a7194146101df57806301ffc9a71461020f578063110e87a61461023f575b600080fd5b6101f960048036038101906101f491906139ab565b610641565b6040516102069190613a33565b60405180910390f35b61022960048036038101906102249190613aa6565b610a52565b6040516102369190613aee565b60405180910390f35b61025960048036038101906102549190613b09565b610acc565b604051610267929190613b52565b60405180910390f35b61028a60048036038101906102859190613b09565b610b0a565b6040516102979190613c39565b60405180910390f35b6102ba60048036038101906102b59190613c87565b610ba1565b6040516102c79190613a33565b60405180910390f35b6102d8610bc0565b005b6102f460048036038101906102ef9190613cb4565b61105a565b005b610310600480360381019061030b9190613cb4565b61107b565b005b61032c60048036038101906103279190613c87565b6110fe565b60405161034199989796959493929190613ea0565b60405180910390f35b610364600480360381019061035f9190613b09565b611349565b604051610372929190613b52565b60405180910390f35b61039560048036038101906103909190613cb4565b611387565b005b61039f6115a1565b6040516103ad929190613b52565b60405180910390f35b6103d060048036038101906103cb9190613b09565b6115cd565b6040516103dd9190613f34565b60405180910390f35b61040060048036038101906103fb9190613cb4565b6115e5565b005b61041c60048036038101906104179190613cb4565b6117c6565b005b6104266119f9565b005b610442600480360381019061043d9190613cb4565b611b66565b60405161044f9190613f4f565b60405180910390f35b610472600480360381019061046d9190613cb4565b612210565b60405161047f9190613aee565b60405180910390f35b61049061227a565b60405161049d9190613a33565b60405180910390f35b6104c060048036038101906104bb9190613b09565b612281565b6040516104cd9190613f34565b60405180910390f35b6104f060048036038101906104eb9190613f71565b612299565b6040516104fd9190613a33565b60405180910390f35b610520600480360381019061051b9190613c87565b6122bd565b60405161052d9190614128565b60405180910390f35b610550600480360381019061054b919061414a565b612610565b005b61056c60048036038101906105679190613c87565b6127df565b005b61058860048036038101906105839190613b09565b612891565b60405161059692919061418a565b60405180910390f35b6105b960048036038101906105b49190613cb4565b61298f565b005b6105d560048036038101906105d09190613cb4565b612bad565b005b6105df612bce565b6040516105ec91906141c9565b60405180910390f35b61060f600480360381019061060a9190613cb4565b612be8565b005b61062b600480360381019061062691906141e4565b612e1b565b6040516106389190613a33565b60405180910390f35b60006106506000801b33612e6f565b6000844260405160200161066592919061428d565b6040516020818303038152906040528051906020012090505b6002600082815260200190815260200160002060080160009054906101000a900460ff16806106af57506000801b81145b156106e257806040516020016106c591906142da565b60405160208183030381529060405280519060200120905061067e565b6040518061012001604052808281526020016000600681111561070857610707613cf4565b5b81526020018481526020018673ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020018581526020016001151581526020016000815250600260008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff021916908360068111156107b1576107b0613cf4565b5b02179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040160146101000a81548160ff02191690831515021790555060c08201518160050160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160020190816109779190614501565b50505060e08201518160080160006101000a81548160ff0219169083151502179055506101008201518160090155905050600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150557f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada38186604051610a3f9291906145d3565b60405180910390a1809150509392505050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610ac55750610ac482612f0c565b5b9050919050565b60046020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6060600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610b9557602002820191906000526020600020905b815481526020019060010190808311610b81575b50505050509050919050565b6000806000838152602001908152602001600020600101549050919050565b610bcd6000801b33612e6f565b60005b6003805490508167ffffffffffffffff16101561105757600060038267ffffffffffffffff1681548110610c0757610c066145fc565b5b906000526020600020015490506000801b8103610c245750611044565b60016006811115610c3857610c37613cf4565b5b6002600083815260200190815260200160002060010160009054906101000a900460ff166006811115610c6e57610c6d613cf4565b5b14158015610cc4575060026006811115610c8b57610c8a613cf4565b5b6002600083815260200190815260200160002060010160009054906101000a900460ff166006811115610cc157610cc0613cf4565b5b14155b15610cff576000801b60038367ffffffffffffffff1681548110610ceb57610cea6145fc565b5b906000526020600020018190555050611044565b610d1f81600b60009054906101000a900467ffffffffffffffff16612f76565b156110425760006002600083815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000801b60038467ffffffffffffffff1681548110610d8157610d806145fc565b5b90600052602060002001819055505b610dd7600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613012565b610e6b576000610e24600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613047565b905060006002600083815260200190815260200160002060010160006101000a81548160ff02191690836006811115610e6057610e5f613cf4565b5b021790555050610d90565b5b610eb3600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613012565b610f47576000610f00600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613047565b905060006002600083815260200190815260200160002060010160006101000a81548160ff02191690836006811115610f3c57610f3b613cf4565b5b021790555050610e6c565b60006006811115610f5b57610f5a613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff166006811115610f9157610f90613cf4565b5b14610fd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc8906146d4565b60405180910390fd5b7f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3826002600085815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516110389291906145d3565b60405180910390a1505b505b808061104f90614723565b915050610bd0565b50565b61106382610ba1565b61106c81613123565b6110768383613137565b505050565b611083613217565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146110f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110e7906147c5565b60405180910390fd5b6110fa828261321f565b5050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160149054906101000a900460ff1690806005016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016002820180546112a990614324565b80601f01602080910402602001604051908101604052809291908181526020018280546112d590614324565b80156113225780601f106112f757610100808354040283529160200191611322565b820191906000526020600020905b81548152906001019060200180831161130557829003601f168201915b505050505081525050908060080160009054906101000a900460ff16908060090154905089565b60056020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6113946000801b33612e6f565b6002600083815260200190815260200160002060080160009054906101000a900460ff166113f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ee90614831565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461149b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611492906148c3565b60405180910390fd5b600260068111156114af576114ae613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff1660068111156114e5576114e4613cf4565b5b14611525576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151c9061492f565b60405180910390fd5b60006002600084815260200190815260200160002060010160006101000a81548160ff0219169083600681111561155f5761155e613cf4565b5b02179055507f5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada082826040516115959291906145d3565b60405180910390a15050565b60098060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b60076020528060005260406000206000915090505481565b6115f26000801b33612e6f565b6002600083815260200190815260200160002060080160009054906101000a900460ff16611655576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164c90614831565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146116f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116f0906148c3565b60405180910390fd5b6002600681111561170d5761170c613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600681111561174357611742613cf4565b5b14611783576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161177a9061492f565b60405180910390fd5b60036002600084815260200190815260200160002060010160006101000a81548160ff021916908360068111156117bd576117bc613cf4565b5b02179055505050565b6117d36000801b33612e6f565b6002600083815260200190815260200160002060080160009054906101000a900460ff16611836576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161182d90614831565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146118da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118d1906149c1565b60405180910390fd5b600360068111156118ee576118ed613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600681111561192457611923613cf4565b5b148061197757506002600681111561193f5761193e613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600681111561197557611974613cf4565b5b145b6119b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119ad90614a2d565b60405180910390fd5b60056002600084815260200190815260200160002060010160006101000a81548160ff021916908360068111156119f0576119ef613cf4565b5b02179055505050565b600060018054906101000a900460ff16159050808015611a2a575060018060009054906101000a900460ff1660ff16105b80611a585750611a3930612e4c565b158015611a57575060018060009054906101000a900460ff1660ff16145b5b611a97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a8e90614abf565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015611ad35760018060016101000a81548160ff0219169083151502179055505b611ae06000801b33613137565b600a600b60006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508015611b635760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051611b5a9190614b27565b60405180910390a15b50565b611b6e6135b8565b611b7b6000801b33612e6f565b6002600084815260200190815260200160002060080160009054906101000a900460ff16611bde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bd590614831565b60405180910390fd5b60006002600085815260200190815260200160002060405180610120016040529081600082015481526020016001820160009054906101000a900460ff166006811115611c2e57611c2d613cf4565b5b6006811115611c4057611c3f613cf4565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460ff16151515158152602001600582016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282018054611e2f90614324565b80601f0160208091040260200160405190810160405280929190818152602001828054611e5b90614324565b8015611ea85780601f10611e7d57610100808354040283529160200191611ea8565b820191906000526020600020905b815481529060010190602001808311611e8b57829003601f168201915b50505050508152505081526020016008820160009054906101000a900460ff1615151515815260200160098201548152505090508273ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff1614611f4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f4590614bb4565b60405180910390fd5b60016006811115611f6257611f61613cf4565b5b81602001516006811115611f7957611f78613cf4565b5b14611fb9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fb090614c46565b60405180910390fd5b600280600086815260200190815260200160002060010160006101000a81548160ff02191690836006811115611ff257611ff1613cf4565b5b021790555043600260008681526020019081526020016000206009018190555061205a600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085613300565b600260008581526020019081526020016000206005016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160028201805461218590614324565b80601f01602080910402602001604051908101604052809291908181526020018280546121b190614324565b80156121fe5780601f106121d3576101008083540402835291602001916121fe565b820191906000526020600020905b8154815290600101906020018083116121e157829003601f168201915b50505050508152505091505092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b60066020528060005260406000206000915090505481565b600381815481106122a957600080fd5b906000526020600020016000915090505481565b6122c5613620565b60006002600084815260200190815260200160002060405180610120016040529081600082015481526020016001820160009054906101000a900460ff16600681111561231557612314613cf4565b5b600681111561232757612326613cf4565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460ff16151515158152602001600582016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160028201805461251690614324565b80601f016020809104026020016040519081016040528092919081815260200182805461254290614324565b801561258f5780601f106125645761010080835404028352916020019161258f565b820191906000526020600020905b81548152906001019060200180831161257257829003601f168201915b50505050508152505081526020016008820160009054906101000a900460ff1615151515815260200160098201548152505090508060e00151612607576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125fe90614cb2565b60405180910390fd5b80915050919050565b61261d6000801b33612e6f565b6002600083815260200190815260200160002060080160009054906101000a900460ff16612680576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161267790614831565b60405180910390fd5b6000600681111561269457612693613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff1660068111156126ca576126c9613cf4565b5b148061271d5750600160068111156126e5576126e4613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600681111561271b5761271a613cf4565b5b145b8061276f57506002600681111561273757612736613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600681111561276d5761276c613cf4565b5b145b6127ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127a590614d1e565b60405180910390fd5b806002600084815260200190815260200160002060020160008282546127d49190614d3e565b925050819055505050565b6127ec6000801b33612e6f565b6002600082815260200190815260200160002060080160009054906101000a900460ff1661284f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161284690614831565b60405180910390fd5b60046002600083815260200190815260200160002060010160006101000a81548160ff0219169083600681111561288957612888613cf4565b5b021790555050565b600061289b6135b8565b6128a86000801b33612e6f565b6128ef600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613012565b1561292f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161292690614dbe565b60405180910390fd5b6000612978600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613047565b9050806129858286611b66565b9250925050915091565b61299c6000801b33612e6f565b6002600083815260200190815260200160002060080160009054906101000a900460ff166129ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129f690614e2a565b60405180910390fd5b60006006811115612a1357612a12613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff166006811115612a4957612a48613cf4565b5b14612a89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a8090614e96565b60405180910390fd5b806002600084815260200190815260200160002060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016002600084815260200190815260200160002060010160006101000a81548160ff02191690836006811115612b1857612b17613cf4565b5b02179055504360026000848152602001908152602001600020600901819055506003829080600181540180825580915050600190039060005260206000200160009091909190915055612ba9600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083613300565b5050565b612bb682610ba1565b612bbf81613123565b612bc9838361321f565b505050565b600b60009054906101000a900467ffffffffffffffff1681565b612bf56000801b33612e6f565b6002600083815260200190815260200160002060080160009054906101000a900460ff16612c58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c4f90614831565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612cfc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cf390614f28565b60405180910390fd5b60036006811115612d1057612d0f613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff166006811115612d4657612d45613cf4565b5b1480612d99575060026006811115612d6157612d60613cf4565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff166006811115612d9757612d96613cf4565b5b145b612dd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dcf90614fba565b60405180910390fd5b60066002600084815260200190815260200160002060010160006101000a81548160ff02191690836006811115612e1257612e11613cf4565b5b02179055505050565b60086020528160005260406000208181548110612e3757600080fd5b90600052602060002001600091509150505481565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b612e798282612210565b612f0857612e9e8173ffffffffffffffffffffffffffffffffffffffff16601461337c565b612eac8360001c602061337c565b604051602001612ebd9291906150ae565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612eff9190615121565b60405180910390fd5b5050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60006002600084815260200190815260200160002060080160009054906101000a900460ff16612fdb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fd290614831565b60405180910390fd5b8167ffffffffffffffff166002600085815260200190815260200160002060090154436130089190615143565b1015905092915050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b13159050919050565b600061305282613012565b15613089576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260000160009054906101000a9004600f0b905082600101600082600f0b600f0b815260200190815260200160002054915082600101600082600f0b600f0b815260200190815260200160002060009055600181018360000160006101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff16021790555050919050565b6131348161312f613217565b612e6f565b50565b6131418282612210565b61321357600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506131b8613217565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6132298282612210565b156132fc57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506132a1613217565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b60008260000160109054906101000a9004600f0b90508183600101600083600f0b600f0b815260200190815260200160002081905550600181018360000160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff160217905550505050565b60606000600283600261338f9190615177565b6133999190614d3e565b67ffffffffffffffff8111156133b2576133b161373f565b5b6040519080825280601f01601f1916602001820160405280156133e45781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061341c5761341b6145fc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106134805761347f6145fc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026134c09190615177565b6134ca9190614d3e565b90505b600181111561356a577f3031323334353637383961626364656600000000000000000000000000000000600f86166010811061350c5761350b6145fc565b5b1a60f81b828281518110613523576135226145fc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080613563906151b9565b90506134cd565b50600084146135ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135a59061522e565b60405180910390fd5b8091505092915050565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b604051806101200160405280600080191681526020016000600681111561364a57613649613cf4565b5b815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020016136a16135b8565b8152602001600015158152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006136f6826136cb565b9050919050565b613706816136eb565b811461371157600080fd5b50565b600081359050613723816136fd565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6137778261372e565b810181811067ffffffffffffffff821117156137965761379561373f565b5b80604052505050565b60006137a96136b7565b90506137b5828261376e565b919050565b600080fd5b600067ffffffffffffffff82169050919050565b6137dc816137bf565b81146137e757600080fd5b50565b6000813590506137f9816137d3565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff8211156138245761382361373f565b5b61382d8261372e565b9050602081019050919050565b82818337600083830152505050565b600061385c61385784613809565b61379f565b90508281526020810184848401111561387857613877613804565b5b61388384828561383a565b509392505050565b600082601f8301126138a05761389f6137ff565b5b81356138b0848260208601613849565b91505092915050565b600060c082840312156138cf576138ce613729565b5b6138d960c061379f565b905060006138e9848285016137ea565b60008301525060206138fd848285016137ea565b6020830152506040613911848285016137ea565b6040830152506060613925848285016137ea565b6060830152506080613939848285016137ea565b60808301525060a082013567ffffffffffffffff81111561395d5761395c6137ba565b5b6139698482850161388b565b60a08301525092915050565b6000819050919050565b61398881613975565b811461399357600080fd5b50565b6000813590506139a58161397f565b92915050565b6000806000606084860312156139c4576139c36136c1565b5b60006139d286828701613714565b935050602084013567ffffffffffffffff8111156139f3576139f26136c6565b5b6139ff868287016138b9565b9250506040613a1086828701613996565b9150509250925092565b6000819050919050565b613a2d81613a1a565b82525050565b6000602082019050613a486000830184613a24565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b613a8381613a4e565b8114613a8e57600080fd5b50565b600081359050613aa081613a7a565b92915050565b600060208284031215613abc57613abb6136c1565b5b6000613aca84828501613a91565b91505092915050565b60008115159050919050565b613ae881613ad3565b82525050565b6000602082019050613b036000830184613adf565b92915050565b600060208284031215613b1f57613b1e6136c1565b5b6000613b2d84828501613714565b91505092915050565b600081600f0b9050919050565b613b4c81613b36565b82525050565b6000604082019050613b676000830185613b43565b613b746020830184613b43565b9392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b613bb081613a1a565b82525050565b6000613bc28383613ba7565b60208301905092915050565b6000602082019050919050565b6000613be682613b7b565b613bf08185613b86565b9350613bfb83613b97565b8060005b83811015613c2c578151613c138882613bb6565b9750613c1e83613bce565b925050600181019050613bff565b5085935050505092915050565b60006020820190508181036000830152613c538184613bdb565b905092915050565b613c6481613a1a565b8114613c6f57600080fd5b50565b600081359050613c8181613c5b565b92915050565b600060208284031215613c9d57613c9c6136c1565b5b6000613cab84828501613c72565b91505092915050565b60008060408385031215613ccb57613cca6136c1565b5b6000613cd985828601613c72565b9250506020613cea85828601613714565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60078110613d3457613d33613cf4565b5b50565b6000819050613d4582613d23565b919050565b6000613d5582613d37565b9050919050565b613d6581613d4a565b82525050565b613d7481613975565b82525050565b613d83816136eb565b82525050565b613d92816137bf565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613dd2578082015181840152602081019050613db7565b60008484015250505050565b6000613de982613d98565b613df38185613da3565b9350613e03818560208601613db4565b613e0c8161372e565b840191505092915050565b600060c083016000830151613e2f6000860182613d89565b506020830151613e426020860182613d89565b506040830151613e556040860182613d89565b506060830151613e686060860182613d89565b506080830151613e7b6080860182613d89565b5060a083015184820360a0860152613e938282613dde565b9150508091505092915050565b600061012082019050613eb6600083018c613a24565b613ec3602083018b613d5c565b613ed0604083018a613d6b565b613edd6060830189613d7a565b613eea6080830188613d7a565b613ef760a0830187613adf565b81810360c0830152613f098186613e17565b9050613f1860e0830185613adf565b613f26610100830184613d6b565b9a9950505050505050505050565b6000602082019050613f496000830184613d6b565b92915050565b60006020820190508181036000830152613f698184613e17565b905092915050565b600060208284031215613f8757613f866136c1565b5b6000613f9584828501613996565b91505092915050565b613fa781613d4a565b82525050565b613fb681613975565b82525050565b613fc5816136eb565b82525050565b613fd481613ad3565b82525050565b600060c083016000830151613ff26000860182613d89565b5060208301516140056020860182613d89565b5060408301516140186040860182613d89565b50606083015161402b6060860182613d89565b50608083015161403e6080860182613d89565b5060a083015184820360a08601526140568282613dde565b9150508091505092915050565b60006101208301600083015161407c6000860182613ba7565b50602083015161408f6020860182613f9e565b5060408301516140a26040860182613fad565b5060608301516140b56060860182613fbc565b5060808301516140c86080860182613fbc565b5060a08301516140db60a0860182613fcb565b5060c083015184820360c08601526140f38282613fda565b91505060e083015161410860e0860182613fcb565b5061010083015161411d610100860182613fad565b508091505092915050565b600060208201905081810360008301526141428184614063565b905092915050565b60008060408385031215614161576141606136c1565b5b600061416f85828601613c72565b925050602061418085828601613996565b9150509250929050565b600060408201905061419f6000830185613a24565b81810360208301526141b18184613e17565b90509392505050565b6141c3816137bf565b82525050565b60006020820190506141de60008301846141ba565b92915050565b600080604083850312156141fb576141fa6136c1565b5b600061420985828601613714565b925050602061421a85828601613996565b9150509250929050565b60008160601b9050919050565b600061423c82614224565b9050919050565b600061424e82614231565b9050919050565b614266614261826136eb565b614243565b82525050565b6000819050919050565b61428761428282613975565b61426c565b82525050565b60006142998285614255565b6014820191506142a98284614276565b6020820191508190509392505050565b6000819050919050565b6142d46142cf82613a1a565b6142b9565b82525050565b60006142e682846142c3565b60208201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061433c57607f821691505b60208210810361434f5761434e6142f5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026143b77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261437a565b6143c1868361437a565b95508019841693508086168417925050509392505050565b6000819050919050565b60006143fe6143f96143f484613975565b6143d9565b613975565b9050919050565b6000819050919050565b614418836143e3565b61442c61442482614405565b848454614387565b825550505050565b600090565b614441614434565b61444c81848461440f565b505050565b5b8181101561447057614465600082614439565b600181019050614452565b5050565b601f8211156144b55761448681614355565b61448f8461436a565b8101602085101561449e578190505b6144b26144aa8561436a565b830182614451565b50505b505050565b600082821c905092915050565b60006144d8600019846008026144ba565b1980831691505092915050565b60006144f183836144c7565b9150826002028217905092915050565b61450a82613d98565b67ffffffffffffffff8111156145235761452261373f565b5b61452d8254614324565b614538828285614474565b600060209050601f83116001811461456b5760008415614559578287015190505b61456385826144e5565b8655506145cb565b601f19841661457986614355565b60005b828110156145a15784890151825560018201915060208501945060208101905061457c565b868310156145be57848901516145ba601f8916826144c7565b8355505b6001600288020188555050505b505050505050565b60006040820190506145e86000830185613a24565b6145f56020830184613d7a565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f54686973206a6f62206d7573742062652050454e44494e47207468657265206960008201527f732061206d69736d61746368206265747765656e2070726f76696465724a6f6260208201527f7320616e6420686f744a6f624c69737400000000000000000000000000000000604082015250565b60006146be60508361462b565b91506146c98261463c565b606082019050919050565b600060208201905081810360008301526146ed816146b1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061472e826137bf565b915067ffffffffffffffff8203614748576147476146f4565b5b600182019050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b60006147af602f8361462b565b91506147ba82614753565b604082019050919050565b600060208201905081810360008301526147de816147a2565b9050919050565b7f4a6f62206d757374206578697374000000000000000000000000000000000000600082015250565b600061481b600e8361462b565b9150614826826147e5565b602082019050919050565b6000602082019050818103600083015261484a8161480e565b9050919050565b7f50726f7669646572732063616e206f6e6c79207374617274207468656972206a60008201527f6f62730000000000000000000000000000000000000000000000000000000000602082015250565b60006148ad60238361462b565b91506148b882614851565b604082019050919050565b600060208201905081810360008301526148dc816148a0565b9050919050565b7f4f6e6c7920717565756564206a6f622063616e20626520737461727465640000600082015250565b6000614919601e8361462b565b9150614924826148e3565b602082019050919050565b600060208201905081810360008301526149488161490c565b9050919050565b7f50726f7669646572732063616e206f6e6c792066696e6973682074686569722060008201527f6a6f627300000000000000000000000000000000000000000000000000000000602082015250565b60006149ab60248361462b565b91506149b68261494f565b604082019050919050565b600060208201905081810360008301526149da8161499e565b9050919050565b7f4f6e6c792072756e6e696e67206a6f622063616e2062652066696e6973686564600082015250565b6000614a1760208361462b565b9150614a22826149e1565b602082019050919050565b60006020820190508181036000830152614a4681614a0a565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000614aa9602e8361462b565b9150614ab482614a4d565b604082019050919050565b60006020820190508181036000830152614ad881614a9c565b9050919050565b6000819050919050565b600060ff82169050919050565b6000614b11614b0c614b0784614adf565b6143d9565b614ae9565b9050919050565b614b2181614af6565b82525050565b6000602082019050614b3c6000830184614b18565b92915050565b7f50726f7669646572732063616e206f6e6c7920636c61696d207468656972206a60008201527f6f62730000000000000000000000000000000000000000000000000000000000602082015250565b6000614b9e60238361462b565b9150614ba982614b42565b604082019050919050565b60006020820190508181036000830152614bcd81614b91565b9050919050565b7f4f6e6c79206d6574612d717565756564206a6f622063616e206265207175657560008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b6000614c3060228361462b565b9150614c3b82614bd4565b604082019050919050565b60006020820190508181036000830152614c5f81614c23565b9050919050565b7f4a6f62206e6f7420666f756e6400000000000000000000000000000000000000600082015250565b6000614c9c600d8361462b565b9150614ca782614c66565b602082019050919050565b60006020820190508181036000830152614ccb81614c8f565b9050919050565b7f4a6f62206d7573742062652072756e6e696e6700000000000000000000000000600082015250565b6000614d0860138361462b565b9150614d1382614cd2565b602082019050919050565b60006020820190508181036000830152614d3781614cfb565b9050919050565b6000614d4982613975565b9150614d5483613975565b9250828201905080821115614d6c57614d6b6146f4565b5b92915050565b7f4e6f20617661696c61626c65206a6f6200000000000000000000000000000000600082015250565b6000614da860108361462b565b9150614db382614d72565b602082019050919050565b60006020820190508181036000830152614dd781614d9b565b9050919050565b7f4a6f62206d7573742062652076616c6964000000000000000000000000000000600082015250565b6000614e1460118361462b565b9150614e1f82614dde565b602082019050919050565b60006020820190508181036000830152614e4381614e07565b9050919050565b7f4a6f62206d7573742062652050454e44494e4700000000000000000000000000600082015250565b6000614e8060138361462b565b9150614e8b82614e4a565b602082019050919050565b60006020820190508181036000830152614eaf81614e73565b9050919050565b7f50726f7669646572732063616e206f6e6c792074726967676572206661696c7560008201527f7265206f66206a6f622074686579206c61756e63686564000000000000000000602082015250565b6000614f1260378361462b565b9150614f1d82614eb6565b604082019050919050565b60006020820190508181036000830152614f4181614f05565b9050919050565b7f4f6e6c79207363686564756c656420616e642072756e6e696e67206a6f62206360008201527f616e206661696c00000000000000000000000000000000000000000000000000602082015250565b6000614fa460278361462b565b9150614faf82614f48565b604082019050919050565b60006020820190508181036000830152614fd381614f97565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b600061501b601783614fda565b915061502682614fe5565b601782019050919050565b600061503c82613d98565b6150468185614fda565b9350615056818560208601613db4565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000615098601183614fda565b91506150a382615062565b601182019050919050565b60006150b98261500e565b91506150c58285615031565b91506150d08261508b565b91506150dc8284615031565b91508190509392505050565b60006150f382613d98565b6150fd818561462b565b935061510d818560208601613db4565b6151168161372e565b840191505092915050565b6000602082019050818103600083015261513b81846150e8565b905092915050565b600061514e82613975565b915061515983613975565b9250828203905081811115615171576151706146f4565b5b92915050565b600061518282613975565b915061518d83613975565b925082820261519b81613975565b915082820484148315176151b2576151b16146f4565b5b5092915050565b60006151c482613975565b9150600082036151d7576151d66146f4565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b600061521860208361462b565b9150615223826151e2565b602082019050919050565b600060208201905081810360008301526152478161520b565b905091905056fea2646970667358221220af68d3808851080e7038ca68ae3929ab0f23b8870b7337a2357734e742cd6d2c64736f6c63430008110033",
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

// MetaSchedulerMetaData contains all meta data concerning the MetaScheduler contract.
var MetaSchedulerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimNextJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_AMOUNT_LOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_METASCHEDULER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVIDER_REGISTRATION_TAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"provider\",\"type\":\"tuple\"}],\"name\":\"convertCreditToDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"provider\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"durationMinute\",\"type\":\"uint64\"}],\"name\":\"convertDurationToCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"actualJobDurationMinute\",\"type\":\"uint64\"}],\"name\":\"finishJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobAmountLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobStatus\",\"outputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_credit\",\"type\":\"address\"},{\"internalType\":\"contractJobManager\",\"name\":\"_jobManager\",\"type\":\"address\"},{\"internalType\":\"contractProviderManager\",\"name\":\"_providerManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleLiveness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"providerApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerRedemption\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"_providerDefinition\",\"type\":\"tuple\"}],\"name\":\"providerRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"_providerDefinition\",\"type\":\"tuple\"}],\"name\":\"providerUpdateHardware\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerNode\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"startJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpMyJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"triggerFailedJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateJobsStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50600060018054906101000a900460ff1615905080801562000044575060018060009054906101000a900460ff1660ff16105b8062000081575062000061306200016360201b62002d0f1760201c565b15801562000080575060018060009054906101000a900460ff1660ff16145b5b620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba906200020d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015620001005760018060016101000a81548160ff0219169083151502179055505b80156200015c5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405162000153919062000289565b60405180910390a15b50620002a6565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620001f5602e8362000186565b9150620002028262000197565b604082019050919050565b600060208201905081810360008301526200022881620001e6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620002716200026b62000265846200022f565b62000246565b62000239565b9050919050565b620002838162000250565b82525050565b6000602082019050620002a0600083018462000278565b92915050565b6152c580620002b66000396000f3fe6080604052600436106101ee5760003560e01c806383ff0b281161010d578063b6e5133f116100a0578063d1cee5461161006f578063d1cee546146106d9578063d547741f14610702578063e8a819861461072b578063f688411b14610754578063fbc3611a14610791576101ee565b8063b6e5133f14610640578063ba9c7f181461067d578063c0c53b8b14610694578063cffe070c146106bd576101ee565b8063a06d083c116100dc578063a06d083c14610582578063a217fddf146105ad578063a492811f146105d8578063b35b67ae14610615576101ee565b806383ff0b28146104b457806389be6044146104df5780638e8554881461050857806391d1485414610545576101ee565b80632fecc4f61161018557806362500f491161015457806362500f49146104105780636a6fad0714610439578063750f0acc1461046257806378952c881461048b576101ee565b80632fecc4f61461036757806336568abe146103905780635d3a7180146103b95780635fae1450146103e7576101ee565b8063236e26ae116101c1578063236e26ae146102c1578063248a9ca3146102ea5780632a242a76146103275780632f2ff15d1461033e576101ee565b806301ffc9a7146101f35780631f92a63f146102305780632081f4c814610259578063214a353c14610296575b600080fd5b3480156101ff57600080fd5b5061021a60048036038101906102159190613326565b61079b565b604051610227919061336e565b60405180910390f35b34801561023c57600080fd5b50610257600480360381019061025291906133bf565b610815565b005b34801561026557600080fd5b50610280600480360381019061027b91906133bf565b610982565b60405161028d9190613463565b60405180910390f35b3480156102a257600080fd5b506102ab610a30565b6040516102b89190613497565b60405180910390f35b3480156102cd57600080fd5b506102e860048036038101906102e391906133bf565b610a3d565b005b3480156102f657600080fd5b50610311600480360381019061030c91906133bf565b610baa565b60405161031e91906134c1565b60405180910390f35b34801561033357600080fd5b5061033c610bc9565b005b34801561034a57600080fd5b506103656004803603810190610360919061353a565b610c77565b005b34801561037357600080fd5b5061038e600480360381019061038991906135a6565b610c98565b005b34801561039c57600080fd5b506103b760048036038101906103b2919061353a565b610e11565b005b3480156103c557600080fd5b506103ce610e94565b6040516103de9493929190613731565b60405180910390f35b3480156103f357600080fd5b5061040e600480360381019061040991906133bf565b611296565b005b34801561041c57600080fd5b50610437600480360381019061043291906137a9565b611581565b005b34801561044557600080fd5b50610460600480360381019061045b9190613922565b611911565b005b34801561046e57600080fd5b5061048960048036038101906104849190613963565b6119cc565b005b34801561049757600080fd5b506104b260048036038101906104ad9190613990565b611b66565b005b3480156104c057600080fd5b506104c9611c20565b6040516104d691906134c1565b60405180910390f35b3480156104eb57600080fd5b50610506600480360381019061050191906133bf565b611c44565b005b34801561051457600080fd5b5061052f600480360381019061052a9190613b33565b611db1565b60405161053c91906134c1565b60405180910390f35b34801561055157600080fd5b5061056c6004803603810190610567919061353a565b611f4c565b604051610579919061336e565b60405180910390f35b34801561058e57600080fd5b50610597611fb6565b6040516105a49190613bee565b60405180910390f35b3480156105b957600080fd5b506105c2611fdc565b6040516105cf91906134c1565b60405180910390f35b3480156105e457600080fd5b506105ff60048036038101906105fa9190613c09565b611fe3565b60405161060c9190613497565b60405180910390f35b34801561062157600080fd5b5061062a6120d1565b6040516106379190613497565b60405180910390f35b34801561064c57600080fd5b50610667600480360381019061066291906133bf565b6120dd565b6040516106749190613497565b60405180910390f35b34801561068957600080fd5b5061069261218b565b005b3480156106a057600080fd5b506106bb60048036038101906106b69190613d34565b612239565b005b6106d760048036038101906106d29190613922565b612520565b005b3480156106e557600080fd5b5061070060048036038101906106fb919061353a565b6126db565b005b34801561070e57600080fd5b506107296004803603810190610724919061353a565b612873565b005b34801561073757600080fd5b50610752600480360381019061074d91906135a6565b612894565b005b34801561076057600080fd5b5061077b60048036038101906107769190613d87565b6129b4565b6040516107889190613e06565b60405180910390f35b610799612aa4565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061080e575061080d82612d32565b5b9050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b81526004016108709190613e21565b602060405180830381865afa15801561088d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b19190613e68565b6108f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e790613f18565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634609ca5082336040518363ffffffff1660e01b815260040161094d929190613f38565b600060405180830381600087803b15801561096757600080fd5b505af115801561097b573d6000803e3d6000fd5b5050505050565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b81526004016109df91906134c1565b600060405180830381865afa1580156109fc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610a259190614202565b602001519050919050565b68056bc75e2d6310000081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401610a989190613e21565b602060405180830381865afa158015610ab5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad99190613e68565b610b18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0f90613f18565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635792edd482336040518363ffffffff1660e01b8152600401610b75929190613f38565b600060405180830381600087803b158015610b8f57600080fd5b505af1158015610ba3573d6000803e3d6000fd5b5050505050565b6000806000838152602001908152602001600020600101549050919050565b610bf37f38b7bc8d4a3fe7426a545cf94fef501dcde1f3ee0d0b8583aa00aa82137b191433612d9c565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a242a766040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610c5d57600080fd5b505af1158015610c71573d6000803e3d6000fd5b50505050565b610c8082610baa565b610c8981612e39565b610c938383612e4d565b505050565b60008111610cdb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd290614297565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401610d3a939291906142b7565b6020604051808303816000875af1158015610d59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7d9190613e68565b50600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b67644b983836040518363ffffffff1660e01b8152600401610ddb9291906142ee565b600060405180830381600087803b158015610df557600080fd5b505af1158015610e09573d6000803e3d6000fd5b505050505050565b610e19612f2d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610e86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7d90614389565b60405180910390fd5b610e908282612f35565b5050565b6000806000610ea1613252565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401610efc9190613e21565b602060405180830381865afa158015610f19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3d9190613e68565b610f7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f7390613f18565b60405180910390fd5b6000610f86613252565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c58467b0336040518263ffffffff1660e01b8152600401610fe19190613e21565b6000604051808303816000875af1158015611000573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061102991906143a9565b80925081935050506000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba846040518263ffffffff1660e01b815260040161108e91906134c1565b600060405180830381865afa1580156110ab573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906110d49190614202565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f6edb5f82608001516040518263ffffffff1660e01b81526004016111359190613e21565b600060405180830381600087803b15801561114f57600080fd5b505af1158015611163573d6000803e3d6000fd5b505050506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663624bc8e3336040518263ffffffff1660e01b81526004016111c49190613e21565b6101a060405180830381865afa1580156111e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112069190614598565b905060006112218360c00151846040015184602001516129b4565b90507f01bc8742bee4608213456e8f5ed4490abbedafbd4f7a5df5706fbb255788c9e583606001513387848760c001516040516112629594939291906145c6565b60405180910390a1826060015185828560c001518167ffffffffffffffff1691509850985098509850505050505090919293565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b81526004016112f391906134c1565b600060405180830381865afa158015611310573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906113399190614202565b9050806060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a490614692565b60405180910390fd5b600060068111156113c1576113c06133ec565b5b816020015160068111156113d8576113d76133ec565b5b148061140c5750600160068111156113f3576113f26133ec565b5b8160200151600681111561140a576114096133ec565b5b145b61144b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161144290614724565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c4d252f5836040518263ffffffff1660e01b81526004016114a691906134c1565b600060405180830381600087803b1580156114c057600080fd5b505af11580156114d4573d6000803e3d6000fd5b50505050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb3383604001516040518363ffffffff1660e01b8152600401611539929190614744565b6020604051808303816000875af1158015611558573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061157c9190613e68565b505050565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637d17544b83336040518363ffffffff1660e01b81526004016115de929190613f38565b600060405180830381600087803b1580156115f857600080fd5b505af115801561160c573d6000803e3d6000fd5b505050506000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba846040518263ffffffff1660e01b815260040161166d91906134c1565b600060405180830381865afa15801561168a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906116b39190614202565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663624bc8e3336040518263ffffffff1660e01b81526004016117129190613e21565b6101a060405180830381865afa158015611730573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117549190614598565b9050600061176f8360c00151846040015184602001516129b4565b905060008490508467ffffffffffffffff168267ffffffffffffffff161015611796578190505b60006117ab8560c00151856020015184611fe3565b9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161180a929190614744565b6020604051808303816000875af1158015611829573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184d9190613e68565b506000818660400151611860919061479c565b9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8760600151836040518363ffffffff1660e01b81526004016118c3929190614744565b6020604051808303816000875af11580156118e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119069190613e68565b505050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e9da3ed838360000151846020015185606001518660a00151876040015188608001518960c001516040518963ffffffff1660e01b81526004016119969897969594939291906147d0565b600060405180830381600087803b1580156119b057600080fd5b505af11580156119c4573d6000803e3d6000fd5b505050505050565b6119d96000801b33612d9c565b8067ffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401611a3f9190613e21565b602060405180830381865afa158015611a5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a80919061484e565b11611ac0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ab7906148ed565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3033846040518463ffffffff1660e01b8152600401611b1f9392919061493e565b6020604051808303816000875af1158015611b3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b629190613e68565b5050565b611b907f38b7bc8d4a3fe7426a545cf94fef501dcde1f3ee0d0b8583aa00aa82137b191433612d9c565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663daea85c5826040518263ffffffff1660e01b8152600401611beb9190613e21565b600060405180830381600087803b158015611c0557600080fd5b505af1158015611c19573d6000803e3d6000fd5b5050505050565b7f38b7bc8d4a3fe7426a545cf94fef501dcde1f3ee0d0b8583aa00aa82137b191481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401611c9f9190613e21565b602060405180830381865afa158015611cbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ce09190613e68565b611d1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d1690613f18565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3401e0082336040518363ffffffff1660e01b8152600401611d7c929190613f38565b600060405180830381600087803b158015611d9657600080fd5b505af1158015611daa573d6000803e3d6000fd5b5050505050565b600068056bc75e2d63100000821015611dff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611df6906149e7565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401611e5e939291906142b7565b6020604051808303816000875af1158015611e7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea19190613e68565b50600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663010a71943385856040518463ffffffff1660e01b8152600401611f0193929190614a07565b6020604051808303816000875af1158015611f20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f449190614a45565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000801b81565b600064e8d4a51000846040015167ffffffffffffffff16856080015167ffffffffffffffff16856080015167ffffffffffffffff166120229190614a72565b61202c9190614a72565b856020015167ffffffffffffffff168560c0015167ffffffffffffffff166120549190614a72565b866000015167ffffffffffffffff16866040015167ffffffffffffffff1661207c9190614a72565b6120869190614ab4565b6120909190614ab4565b856060015167ffffffffffffffff168467ffffffffffffffff166120b49190614a72565b6120be9190614a72565b6120c89190614a72565b90509392505050565b670de0b6b3a764000081565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b815260040161213a91906134c1565b600060405180830381865afa158015612157573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906121809190614202565b604001519050919050565b6121b57f38b7bc8d4a3fe7426a545cf94fef501dcde1f3ee0d0b8583aa00aa82137b191433612d9c565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a242a766040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561221f57600080fd5b505af1158015612233573d6000803e3d6000fd5b50505050565b600060018054906101000a900460ff1615905080801561226a575060018060009054906101000a900460ff1660ff16105b80612298575061227930612d0f565b158015612297575060018060009054906101000a900460ff1660ff16145b5b6122d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ce90614b5a565b60405180910390fd5b60018060006101000a81548160ff021916908360ff16021790555080156123135760018060016101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603612382576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161237990614bec565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036123f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123e890614c7e565b60405180910390fd5b6123fe6000801b33612e4d565b83600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561251a5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516125119190614ce6565b60405180910390a15b50505050565b670de0b6b3a764000034101561256b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161256290614d73565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff16670de0b6b3a764000060405161259990614dc4565b60006040518083038185875af1925050503d80600081146125d6576040519150601f19603f3d011682016040523d82523d6000602084013e6125db565b606091505b505090508061261f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161261690614e25565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b7c8116e848460000151856020015186606001518760a00151886040015189608001518a60c001516040518963ffffffff1660e01b81526004016126a49897969594939291906147d0565b600060405180830381600087803b1580156126be57600080fd5b505af11580156126d2573d6000803e3d6000fd5b50505050505050565b6127057f38b7bc8d4a3fe7426a545cf94fef501dcde1f3ee0d0b8583aa00aa82137b191433612d9c565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b81526004016127609190613e21565b602060405180830381865afa15801561277d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127a19190613e68565b6127e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127d790613f18565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d1cee54683836040518363ffffffff1660e01b815260040161283d929190613f38565b600060405180830381600087803b15801561285757600080fd5b505af115801561286b573d6000803e3d6000fd5b505050505050565b61287c82610baa565b61288581612e39565b61288f8383612f35565b505050565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b81526004016128ef91906134c1565b600060405180830381865afa15801561290c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906129359190614202565b6060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146129a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161299d90614e91565b60405180910390fd5b6129b08282610c98565b5050565b60008064e8d4a510008560400151866080015185608001516129d69190614eb1565b6129e09190614eb1565b86602001518560c001516129f49190614eb1565b87600001518660400151612a089190614eb1565b612a129190614eee565b612a1c9190614eee565b8660600151612a2b9190614eb1565b612a359190614eb1565b905060008167ffffffffffffffff1603612a84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a7b90614f76565b60405180910390fd5b8067ffffffffffffffff1684612a9a9190614fc5565b9150509392505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b7bb9145336040518263ffffffff1660e01b8152600401612aff9190613e21565b602060405180830381865afa158015612b1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b409190613e68565b612b7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b7690615042565b60405180910390fd5b670de0b6b3a7640000341015612bca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bc190615042565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff16683635c9adc5dea00000604051612bf990614dc4565b60006040518083038185875af1925050503d8060008114612c36576040519150601f19603f3d011682016040523d82523d6000602084013e612c3b565b606091505b5050905080612c7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c7690614e25565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7692d09336040518263ffffffff1660e01b8152600401612cda9190613e21565b600060405180830381600087803b158015612cf457600080fd5b505af1158015612d08573d6000803e3d6000fd5b5050505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b612da68282611f4c565b612e3557612dcb8173ffffffffffffffffffffffffffffffffffffffff166014613016565b612dd98360001c6020613016565b604051602001612dea929190615136565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e2c91906151a9565b60405180910390fd5b5050565b612e4a81612e45612f2d565b612d9c565b50565b612e578282611f4c565b612f2957600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612ece612f2d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b612f3f8282611f4c565b1561301257600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612fb7612f2d565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6060600060028360026130299190614a72565b6130339190614ab4565b67ffffffffffffffff81111561304c5761304b6137ee565b5b6040519080825280601f01601f19166020018201604052801561307e5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106130b6576130b56151cb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061311a576131196151cb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000600184600261315a9190614a72565b6131649190614ab4565b90505b6001811115613204577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106131a6576131a56151cb565b5b1a60f81b8282815181106131bd576131bc6151cb565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806131fd906151fa565b9050613167565b5060008414613248576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161323f9061526f565b60405180910390fd5b8091505092915050565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b613303816132ce565b811461330e57600080fd5b50565b600081359050613320816132fa565b92915050565b60006020828403121561333c5761333b6132c4565b5b600061334a84828501613311565b91505092915050565b60008115159050919050565b61336881613353565b82525050565b6000602082019050613383600083018461335f565b92915050565b6000819050919050565b61339c81613389565b81146133a757600080fd5b50565b6000813590506133b981613393565b92915050565b6000602082840312156133d5576133d46132c4565b5b60006133e3848285016133aa565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6007811061342c5761342b6133ec565b5b50565b600081905061343d8261341b565b919050565b600061344d8261342f565b9050919050565b61345d81613442565b82525050565b60006020820190506134786000830184613454565b92915050565b6000819050919050565b6134918161347e565b82525050565b60006020820190506134ac6000830184613488565b92915050565b6134bb81613389565b82525050565b60006020820190506134d660008301846134b2565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613507826134dc565b9050919050565b613517816134fc565b811461352257600080fd5b50565b6000813590506135348161350e565b92915050565b60008060408385031215613551576135506132c4565b5b600061355f858286016133aa565b925050602061357085828601613525565b9150509250929050565b6135838161347e565b811461358e57600080fd5b50565b6000813590506135a08161357a565b92915050565b600080604083850312156135bd576135bc6132c4565b5b60006135cb858286016133aa565b92505060206135dc85828601613591565b9150509250929050565b6135ef816134fc565b82525050565b600067ffffffffffffffff82169050919050565b613612816135f5565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613652578082015181840152602081019050613637565b60008484015250505050565b6000601f19601f8301169050919050565b600061367a82613618565b6136848185613623565b9350613694818560208601613634565b61369d8161365e565b840191505092915050565b600060c0830160008301516136c06000860182613609565b5060208301516136d36020860182613609565b5060408301516136e66040860182613609565b5060608301516136f96060860182613609565b50608083015161370c6080860182613609565b5060a083015184820360a0860152613724828261366f565b9150508091505092915050565b600060808201905061374660008301876135e6565b61375360208301866134b2565b6137606040830185613488565b818103606083015261377281846136a8565b905095945050505050565b613786816135f5565b811461379157600080fd5b50565b6000813590506137a38161377d565b92915050565b600080604083850312156137c0576137bf6132c4565b5b60006137ce858286016133aa565b92505060206137df85828601613794565b9150509250929050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6138268261365e565b810181811067ffffffffffffffff82111715613845576138446137ee565b5b80604052505050565b60006138586132ba565b9050613864828261381d565b919050565b600080fd5b600060e08284031215613884576138836137e9565b5b61388e60e061384e565b9050600061389e84828501613794565b60008301525060206138b284828501613794565b60208301525060406138c684828501613794565b60408301525060606138da84828501613794565b60608301525060806138ee84828501613794565b60808301525060a061390284828501613794565b60a08301525060c061391684828501613794565b60c08301525092915050565b600080610100838503121561393a576139396132c4565b5b600061394885828601613525565b92505060206139598582860161386e565b9150509250929050565b600060208284031215613979576139786132c4565b5b600061398784828501613794565b91505092915050565b6000602082840312156139a6576139a56132c4565b5b60006139b484828501613525565b91505092915050565b600080fd5b600080fd5b600067ffffffffffffffff8211156139e2576139e16137ee565b5b6139eb8261365e565b9050602081019050919050565b82818337600083830152505050565b6000613a1a613a15846139c7565b61384e565b905082815260208101848484011115613a3657613a356139c2565b5b613a418482856139f8565b509392505050565b600082601f830112613a5e57613a5d6139bd565b5b8135613a6e848260208601613a07565b91505092915050565b600060c08284031215613a8d57613a8c6137e9565b5b613a9760c061384e565b90506000613aa784828501613794565b6000830152506020613abb84828501613794565b6020830152506040613acf84828501613794565b6040830152506060613ae384828501613794565b6060830152506080613af784828501613794565b60808301525060a082013567ffffffffffffffff811115613b1b57613b1a613869565b5b613b2784828501613a49565b60a08301525092915050565b60008060408385031215613b4a57613b496132c4565b5b600083013567ffffffffffffffff811115613b6857613b676132c9565b5b613b7485828601613a77565b9250506020613b8585828601613591565b9150509250929050565b6000819050919050565b6000613bb4613baf613baa846134dc565b613b8f565b6134dc565b9050919050565b6000613bc682613b99565b9050919050565b6000613bd882613bbb565b9050919050565b613be881613bcd565b82525050565b6000602082019050613c036000830184613bdf565b92915050565b60008060006101208486031215613c2357613c226132c4565b5b600084013567ffffffffffffffff811115613c4157613c406132c9565b5b613c4d86828701613a77565b9350506020613c5e8682870161386e565b925050610100613c7086828701613794565b9150509250925092565b6000613c85826134fc565b9050919050565b613c9581613c7a565b8114613ca057600080fd5b50565b600081359050613cb281613c8c565b92915050565b6000613cc3826134fc565b9050919050565b613cd381613cb8565b8114613cde57600080fd5b50565b600081359050613cf081613cca565b92915050565b6000613d01826134fc565b9050919050565b613d1181613cf6565b8114613d1c57600080fd5b50565b600081359050613d2e81613d08565b92915050565b600080600060608486031215613d4d57613d4c6132c4565b5b6000613d5b86828701613ca3565b9350506020613d6c86828701613ce1565b9250506040613d7d86828701613d1f565b9150509250925092565b60008060006101208486031215613da157613da06132c4565b5b600084013567ffffffffffffffff811115613dbf57613dbe6132c9565b5b613dcb86828701613a77565b9350506020613ddc86828701613591565b9250506040613ded8682870161386e565b9150509250925092565b613e00816135f5565b82525050565b6000602082019050613e1b6000830184613df7565b92915050565b6000602082019050613e3660008301846135e6565b92915050565b613e4581613353565b8114613e5057600080fd5b50565b600081519050613e6281613e3c565b92915050565b600060208284031215613e7e57613e7d6132c4565b5b6000613e8c84828501613e53565b91505092915050565b600082825260208201905092915050565b7f50726f766964657273206e6f7420666f756e64206f72206e6f74206a6f696e6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b6000613f02602183613e95565b9150613f0d82613ea6565b604082019050919050565b60006020820190508181036000830152613f3181613ef5565b9050919050565b6000604082019050613f4d60008301856134b2565b613f5a60208301846135e6565b9392505050565b600081519050613f7081613393565b92915050565b60078110613f8357600080fd5b50565b600081519050613f9581613f76565b92915050565b600081519050613faa8161357a565b92915050565b600081519050613fbf8161350e565b92915050565b600081519050613fd48161377d565b92915050565b6000613fed613fe8846139c7565b61384e565b905082815260208101848484011115614009576140086139c2565b5b614014848285613634565b509392505050565b600082601f830112614031576140306139bd565b5b8151614041848260208601613fda565b91505092915050565b600060c082840312156140605761405f6137e9565b5b61406a60c061384e565b9050600061407a84828501613fc5565b600083015250602061408e84828501613fc5565b60208301525060406140a284828501613fc5565b60408301525060606140b684828501613fc5565b60608301525060806140ca84828501613fc5565b60808301525060a082015167ffffffffffffffff8111156140ee576140ed613869565b5b6140fa8482850161401c565b60a08301525092915050565b6000610120828403121561411d5761411c6137e9565b5b61412861012061384e565b9050600061413884828501613f61565b600083015250602061414c84828501613f86565b602083015250604061416084828501613f9b565b604083015250606061417484828501613fb0565b606083015250608061418884828501613fb0565b60808301525060a061419c84828501613e53565b60a08301525060c082015167ffffffffffffffff8111156141c0576141bf613869565b5b6141cc8482850161404a565b60c08301525060e06141e084828501613e53565b60e0830152506101006141f584828501613f9b565b6101008301525092915050565b600060208284031215614218576142176132c4565b5b600082015167ffffffffffffffff811115614236576142356132c9565b5b61424284828501614106565b91505092915050565b7f616d6f756e74206d757374206e6f74206265206e756c6c000000000000000000600082015250565b6000614281601783613e95565b915061428c8261424b565b602082019050919050565b600060208201905081810360008301526142b081614274565b9050919050565b60006060820190506142cc60008301866135e6565b6142d960208301856135e6565b6142e66040830184613488565b949350505050565b600060408201905061430360008301856134b2565b6143106020830184613488565b9392505050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000614373602f83613e95565b915061437e82614317565b604082019050919050565b600060208201905081810360008301526143a281614366565b9050919050565b600080604083850312156143c0576143bf6132c4565b5b60006143ce85828601613f61565b925050602083015167ffffffffffffffff8111156143ef576143ee6132c9565b5b6143fb8582860161404a565b9150509250929050565b600060e0828403121561441b5761441a6137e9565b5b61442560e061384e565b9050600061443584828501613fc5565b600083015250602061444984828501613fc5565b602083015250604061445d84828501613fc5565b604083015250606061447184828501613fc5565b606083015250608061448584828501613fc5565b60808301525060a061449984828501613fc5565b60a08301525060c06144ad84828501613fc5565b60c08301525092915050565b600381106144c657600080fd5b50565b6000815190506144d8816144b9565b92915050565b60006101a082840312156144f5576144f46137e9565b5b6144ff60e061384e565b9050600061450f84828501613fb0565b600083015250602061452384828501614405565b602083015250610100614538848285016144c9565b60408301525061012061454d84828501613e53565b60608301525061014061456284828501613fc5565b60808301525061016061457784828501613f9b565b60a08301525061018061458c84828501613f9b565b60c08301525092915050565b60006101a082840312156145af576145ae6132c4565b5b60006145bd848285016144de565b91505092915050565b600060a0820190506145db60008301886135e6565b6145e860208301876135e6565b6145f560408301866134b2565b6146026060830185613df7565b818103608083015261461481846136a8565b90509695505050505050565b7f4f6e6c7920746865206a6f62206f776e65722063616e2063616e63656c20697460008201527f73206a6f62000000000000000000000000000000000000000000000000000000602082015250565b600061467c602583613e95565b915061468782614620565b604082019050919050565b600060208201905081810360008301526146ab8161466f565b9050919050565b7f4f6e6c792050454e44494e4720616e64204d4554415f5343484544554c45442060008201527f6a6f62732063616e2062652063616e63656c6c65640000000000000000000000602082015250565b600061470e603583613e95565b9150614719826146b2565b604082019050919050565b6000602082019050818103600083015261473d81614701565b9050919050565b600060408201905061475960008301856135e6565b6147666020830184613488565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006147a78261347e565b91506147b28361347e565b92508282039050818111156147ca576147c961476d565b5b92915050565b6000610100820190506147e6600083018b6135e6565b6147f3602083018a613df7565b6148006040830189613df7565b61480d6060830188613df7565b61481a6080830187613df7565b61482760a0830186613df7565b61483460c0830185613df7565b61484160e0830184613df7565b9998505050505050505050565b600060208284031215614864576148636132c4565b5b600061487284828501613f9b565b91505092915050565b7f4d6574615363686564756c65723a20776974686472617720616d6f756e74206860008201527f6967686572207468616e2062616c616e63652e00000000000000000000000000602082015250565b60006148d7603383613e95565b91506148e28261487b565b604082019050919050565b60006020820190508181036000830152614906816148ca565b9050919050565b600061492861492361491e846135f5565b613b8f565b61347e565b9050919050565b6149388161490d565b82525050565b600060608201905061495360008301866135e6565b61496060208301856135e6565b61496d604083018461492f565b949350505050565b7f5f616d6f756e744c6f636b6564206d757374206265206772656174657220746860008201527f616e204d494e494d554d5f414d4f554e545f4c4f434b00000000000000000000602082015250565b60006149d1603683613e95565b91506149dc82614975565b604082019050919050565b60006020820190508181036000830152614a00816149c4565b9050919050565b6000606082019050614a1c60008301866135e6565b8181036020830152614a2e81856136a8565b9050614a3d6040830184613488565b949350505050565b600060208284031215614a5b57614a5a6132c4565b5b6000614a6984828501613f61565b91505092915050565b6000614a7d8261347e565b9150614a888361347e565b9250828202614a968161347e565b91508282048414831517614aad57614aac61476d565b5b5092915050565b6000614abf8261347e565b9150614aca8361347e565b9250828201905080821115614ae257614ae161476d565b5b92915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000614b44602e83613e95565b9150614b4f82614ae8565b604082019050919050565b60006020820190508181036000830152614b7381614b37565b9050919050565b7f4d6574615363686564756c65723a206372656469742061646472206973207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000614bd6602283613e95565b9150614be182614b7a565b604082019050919050565b60006020820190508181036000830152614c0581614bc9565b9050919050565b7f4d6574615363686564756c65723a2070726f76696465724d616e61676572206160008201527f646472206973207a65726f000000000000000000000000000000000000000000602082015250565b6000614c68602b83613e95565b9150614c7382614c0c565b604082019050919050565b60006020820190508181036000830152614c9781614c5b565b9050919050565b6000819050919050565b600060ff82169050919050565b6000614cd0614ccb614cc684614c9e565b613b8f565b614ca8565b9050919050565b614ce081614cb5565b82525050565b6000602082019050614cfb6000830184614cd7565b92915050565b7f4d696e696d756d20616d6f756e7420746f207265676973746572206e6f74207260008201527f6561636865640000000000000000000000000000000000000000000000000000602082015250565b6000614d5d602683613e95565b9150614d6882614d01565b604082019050919050565b60006020820190508181036000830152614d8c81614d50565b9050919050565b600081905092915050565b50565b6000614dae600083614d93565b9150614db982614d9e565b600082019050919050565b6000614dcf82614da1565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000614e0f601083613e95565b9150614e1a82614dd9565b602082019050919050565b60006020820190508181036000830152614e3e81614e02565b9050919050565b7f4f6e6c79206a6f62206f776e65722063616e2063616c6c207468697300000000600082015250565b6000614e7b601c83613e95565b9150614e8682614e45565b602082019050919050565b60006020820190508181036000830152614eaa81614e6e565b9050919050565b6000614ebc826135f5565b9150614ec7836135f5565b9250828202614ed5816135f5565b9150808214614ee757614ee661476d565b5b5092915050565b6000614ef9826135f5565b9150614f04836135f5565b9250828201905067ffffffffffffffff811115614f2457614f2361476d565b5b92915050565b7f43616e6e6f742064697669646520627920300000000000000000000000000000600082015250565b6000614f60601283613e95565b9150614f6b82614f2a565b602082019050919050565b60006020820190508181036000830152614f8f81614f53565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000614fd08261347e565b9150614fdb8361347e565b925082614feb57614fea614f96565b5b828204905092915050565b7f50726f7669646572206e6f7420666f756e64206f72206e6f74206b69636b6564600082015250565b600061502c602083613e95565b915061503782614ff6565b602082019050919050565b6000602082019050818103600083015261505b8161501f565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b60006150a3601783615062565b91506150ae8261506d565b601782019050919050565b60006150c482613618565b6150ce8185615062565b93506150de818560208601613634565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000615120601183615062565b915061512b826150ea565b601182019050919050565b600061514182615096565b915061514d82856150b9565b915061515882615113565b915061516482846150b9565b91508190509392505050565b600061517b82613618565b6151858185613e95565b9350615195818560208601613634565b61519e8161365e565b840191505092915050565b600060208201905081810360008301526151c38184615170565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006152058261347e565b9150600082036152185761521761476d565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000615259602083613e95565b915061526482615223565b602082019050919050565b600060208201905081810360008301526152888161524c565b905091905056fea264697066735822122012ba6daec114c927e18c86558cb908f4fa0cc1744c3b154df40da8654a9133cf64736f6c63430008110033",
}

// MetaSchedulerABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaSchedulerMetaData.ABI instead.
var MetaSchedulerABI = MetaSchedulerMetaData.ABI

// MetaSchedulerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MetaSchedulerMetaData.Bin instead.
var MetaSchedulerBin = MetaSchedulerMetaData.Bin

// DeployMetaScheduler deploys a new Ethereum contract, binding an instance of MetaScheduler to it.
func DeployMetaScheduler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MetaScheduler, error) {
	parsed, err := MetaSchedulerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MetaSchedulerBin), backend)
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

// MINIMUMAMOUNTLOCK is a free data retrieval call binding the contract method 0x214a353c.
//
// Solidity: function MINIMUM_AMOUNT_LOCK() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) MINIMUMAMOUNTLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "MINIMUM_AMOUNT_LOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMAMOUNTLOCK is a free data retrieval call binding the contract method 0x214a353c.
//
// Solidity: function MINIMUM_AMOUNT_LOCK() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) MINIMUMAMOUNTLOCK() (*big.Int, error) {
	return _MetaScheduler.Contract.MINIMUMAMOUNTLOCK(&_MetaScheduler.CallOpts)
}

// MINIMUMAMOUNTLOCK is a free data retrieval call binding the contract method 0x214a353c.
//
// Solidity: function MINIMUM_AMOUNT_LOCK() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) MINIMUMAMOUNTLOCK() (*big.Int, error) {
	return _MetaScheduler.Contract.MINIMUMAMOUNTLOCK(&_MetaScheduler.CallOpts)
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

// PROVIDERREGISTRATIONTAX is a free data retrieval call binding the contract method 0xb35b67ae.
//
// Solidity: function PROVIDER_REGISTRATION_TAX() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) PROVIDERREGISTRATIONTAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "PROVIDER_REGISTRATION_TAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROVIDERREGISTRATIONTAX is a free data retrieval call binding the contract method 0xb35b67ae.
//
// Solidity: function PROVIDER_REGISTRATION_TAX() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) PROVIDERREGISTRATIONTAX() (*big.Int, error) {
	return _MetaScheduler.Contract.PROVIDERREGISTRATIONTAX(&_MetaScheduler.CallOpts)
}

// PROVIDERREGISTRATIONTAX is a free data retrieval call binding the contract method 0xb35b67ae.
//
// Solidity: function PROVIDER_REGISTRATION_TAX() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) PROVIDERREGISTRATIONTAX() (*big.Int, error) {
	return _MetaScheduler.Contract.PROVIDERREGISTRATIONTAX(&_MetaScheduler.CallOpts)
}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0xf688411b.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,uint64,string) jobDefinition, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint64)
func (_MetaScheduler *MetaSchedulerCaller) ConvertCreditToDuration(opts *bind.CallOpts, jobDefinition JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (uint64, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "convertCreditToDuration", jobDefinition, amountLocked, provider)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0xf688411b.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,uint64,string) jobDefinition, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint64)
func (_MetaScheduler *MetaSchedulerSession) ConvertCreditToDuration(jobDefinition JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (uint64, error) {
	return _MetaScheduler.Contract.ConvertCreditToDuration(&_MetaScheduler.CallOpts, jobDefinition, amountLocked, provider)
}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0xf688411b.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,uint64,string) jobDefinition, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint64)
func (_MetaScheduler *MetaSchedulerCallerSession) ConvertCreditToDuration(jobDefinition JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (uint64, error) {
	return _MetaScheduler.Contract.ConvertCreditToDuration(&_MetaScheduler.CallOpts, jobDefinition, amountLocked, provider)
}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0xa492811f.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,uint64,string) jobDefinition, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint64 durationMinute) pure returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) ConvertDurationToCredit(opts *bind.CallOpts, jobDefinition JobDefinition, provider ProviderDefinition, durationMinute uint64) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "convertDurationToCredit", jobDefinition, provider, durationMinute)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0xa492811f.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,uint64,string) jobDefinition, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint64 durationMinute) pure returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) ConvertDurationToCredit(jobDefinition JobDefinition, provider ProviderDefinition, durationMinute uint64) (*big.Int, error) {
	return _MetaScheduler.Contract.ConvertDurationToCredit(&_MetaScheduler.CallOpts, jobDefinition, provider, durationMinute)
}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0xa492811f.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,uint64,string) jobDefinition, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint64 durationMinute) pure returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) ConvertDurationToCredit(jobDefinition JobDefinition, provider ProviderDefinition, durationMinute uint64) (*big.Int, error) {
	return _MetaScheduler.Contract.ConvertDurationToCredit(&_MetaScheduler.CallOpts, jobDefinition, provider, durationMinute)
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

// GetJobAmountLock is a free data retrieval call binding the contract method 0xb6e5133f.
//
// Solidity: function getJobAmountLock(bytes32 _jobId) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) GetJobAmountLock(opts *bind.CallOpts, _jobId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "getJobAmountLock", _jobId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobAmountLock is a free data retrieval call binding the contract method 0xb6e5133f.
//
// Solidity: function getJobAmountLock(bytes32 _jobId) view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) GetJobAmountLock(_jobId [32]byte) (*big.Int, error) {
	return _MetaScheduler.Contract.GetJobAmountLock(&_MetaScheduler.CallOpts, _jobId)
}

// GetJobAmountLock is a free data retrieval call binding the contract method 0xb6e5133f.
//
// Solidity: function getJobAmountLock(bytes32 _jobId) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) GetJobAmountLock(_jobId [32]byte) (*big.Int, error) {
	return _MetaScheduler.Contract.GetJobAmountLock(&_MetaScheduler.CallOpts, _jobId)
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

// FinishJob is a paid mutator transaction binding the contract method 0x62500f49.
//
// Solidity: function finishJob(bytes32 _jobId, uint64 actualJobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerTransactor) FinishJob(opts *bind.TransactOpts, _jobId [32]byte, actualJobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "finishJob", _jobId, actualJobDurationMinute)
}

// FinishJob is a paid mutator transaction binding the contract method 0x62500f49.
//
// Solidity: function finishJob(bytes32 _jobId, uint64 actualJobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerSession) FinishJob(_jobId [32]byte, actualJobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.FinishJob(&_MetaScheduler.TransactOpts, _jobId, actualJobDurationMinute)
}

// FinishJob is a paid mutator transaction binding the contract method 0x62500f49.
//
// Solidity: function finishJob(bytes32 _jobId, uint64 actualJobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) FinishJob(_jobId [32]byte, actualJobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.FinishJob(&_MetaScheduler.TransactOpts, _jobId, actualJobDurationMinute)
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

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _credit, address _jobManager, address _providerManager) returns()
func (_MetaScheduler *MetaSchedulerTransactor) Initialize(opts *bind.TransactOpts, _credit common.Address, _jobManager common.Address, _providerManager common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "initialize", _credit, _jobManager, _providerManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _credit, address _jobManager, address _providerManager) returns()
func (_MetaScheduler *MetaSchedulerSession) Initialize(_credit common.Address, _jobManager common.Address, _providerManager common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Initialize(&_MetaScheduler.TransactOpts, _credit, _jobManager, _providerManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _credit, address _jobManager, address _providerManager) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) Initialize(_credit common.Address, _jobManager common.Address, _providerManager common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Initialize(&_MetaScheduler.TransactOpts, _credit, _jobManager, _providerManager)
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

// ProviderApprove is a paid mutator transaction binding the contract method 0x78952c88.
//
// Solidity: function providerApprove(address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderApprove(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerApprove", _providerAddr)
}

// ProviderApprove is a paid mutator transaction binding the contract method 0x78952c88.
//
// Solidity: function providerApprove(address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderApprove(_providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderApprove(&_MetaScheduler.TransactOpts, _providerAddr)
}

// ProviderApprove is a paid mutator transaction binding the contract method 0x78952c88.
//
// Solidity: function providerApprove(address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderApprove(_providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderApprove(&_MetaScheduler.TransactOpts, _providerAddr)
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

// ProviderRegister is a paid mutator transaction binding the contract method 0xcffe070c.
//
// Solidity: function providerRegister(address _providerAddr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) _providerDefinition) payable returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderRegister(opts *bind.TransactOpts, _providerAddr common.Address, _providerDefinition ProviderDefinition) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerRegister", _providerAddr, _providerDefinition)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0xcffe070c.
//
// Solidity: function providerRegister(address _providerAddr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) _providerDefinition) payable returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderRegister(_providerAddr common.Address, _providerDefinition ProviderDefinition) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderRegister(&_MetaScheduler.TransactOpts, _providerAddr, _providerDefinition)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0xcffe070c.
//
// Solidity: function providerRegister(address _providerAddr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) _providerDefinition) payable returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderRegister(_providerAddr common.Address, _providerDefinition ProviderDefinition) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderRegister(&_MetaScheduler.TransactOpts, _providerAddr, _providerDefinition)
}

// ProviderUpdateHardware is a paid mutator transaction binding the contract method 0x6a6fad07.
//
// Solidity: function providerUpdateHardware(address _providerAddr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) _providerDefinition) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderUpdateHardware(opts *bind.TransactOpts, _providerAddr common.Address, _providerDefinition ProviderDefinition) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerUpdateHardware", _providerAddr, _providerDefinition)
}

// ProviderUpdateHardware is a paid mutator transaction binding the contract method 0x6a6fad07.
//
// Solidity: function providerUpdateHardware(address _providerAddr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) _providerDefinition) returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderUpdateHardware(_providerAddr common.Address, _providerDefinition ProviderDefinition) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderUpdateHardware(&_MetaScheduler.TransactOpts, _providerAddr, _providerDefinition)
}

// ProviderUpdateHardware is a paid mutator transaction binding the contract method 0x6a6fad07.
//
// Solidity: function providerUpdateHardware(address _providerAddr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) _providerDefinition) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderUpdateHardware(_providerAddr common.Address, _providerDefinition ProviderDefinition) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderUpdateHardware(&_MetaScheduler.TransactOpts, _providerAddr, _providerDefinition)
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

// TopUpMyJob is a paid mutator transaction binding the contract method 0xe8a81986.
//
// Solidity: function topUpMyJob(bytes32 _jobId, uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactor) TopUpMyJob(opts *bind.TransactOpts, _jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "topUpMyJob", _jobId, _amount)
}

// TopUpMyJob is a paid mutator transaction binding the contract method 0xe8a81986.
//
// Solidity: function topUpMyJob(bytes32 _jobId, uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerSession) TopUpMyJob(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpMyJob(&_MetaScheduler.TransactOpts, _jobId, _amount)
}

// TopUpMyJob is a paid mutator transaction binding the contract method 0xe8a81986.
//
// Solidity: function topUpMyJob(bytes32 _jobId, uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) TopUpMyJob(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpMyJob(&_MetaScheduler.TransactOpts, _jobId, _amount)
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

// UpdateJobsStatus is a paid mutator transaction binding the contract method 0x2a242a76.
//
// Solidity: function updateJobsStatus() returns()
func (_MetaScheduler *MetaSchedulerTransactor) UpdateJobsStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "updateJobsStatus")
}

// UpdateJobsStatus is a paid mutator transaction binding the contract method 0x2a242a76.
//
// Solidity: function updateJobsStatus() returns()
func (_MetaScheduler *MetaSchedulerSession) UpdateJobsStatus() (*types.Transaction, error) {
	return _MetaScheduler.Contract.UpdateJobsStatus(&_MetaScheduler.TransactOpts)
}

// UpdateJobsStatus is a paid mutator transaction binding the contract method 0x2a242a76.
//
// Solidity: function updateJobsStatus() returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) UpdateJobsStatus() (*types.Transaction, error) {
	return _MetaScheduler.Contract.UpdateJobsStatus(&_MetaScheduler.TransactOpts)
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
	ProviderAddr      common.Address
	JobId             [32]byte
	MaxDurationMinute uint64
	JobDefinition     JobDefinition
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimNextJobEvent is a free log retrieval operation binding the contract event 0x01bc8742bee4608213456e8f5ed4490abbedafbd4f7a5df5706fbb255788c9e5.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,uint64,string) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) FilterClaimNextJobEvent(opts *bind.FilterOpts) (*MetaSchedulerClaimNextJobEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "ClaimNextJobEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerClaimNextJobEventIterator{contract: _MetaScheduler.contract, event: "ClaimNextJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimNextJobEvent is a free log subscription operation binding the contract event 0x01bc8742bee4608213456e8f5ed4490abbedafbd4f7a5df5706fbb255788c9e5.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,uint64,string) jobDefinition)
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

// ParseClaimNextJobEvent is a log parse operation binding the contract event 0x01bc8742bee4608213456e8f5ed4490abbedafbd4f7a5df5706fbb255788c9e5.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,uint64,string) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) ParseClaimNextJobEvent(log types.Log) (*MetaSchedulerClaimNextJobEvent, error) {
	event := new(MetaSchedulerClaimNextJobEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "ClaimNextJobEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaSchedulerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MetaScheduler contract.
type MetaSchedulerInitializedIterator struct {
	Event *MetaSchedulerInitialized // Event containing the contract specifics and raw log

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
func (it *MetaSchedulerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaSchedulerInitialized)
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
		it.Event = new(MetaSchedulerInitialized)
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
func (it *MetaSchedulerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaSchedulerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaSchedulerInitialized represents a Initialized event raised by the MetaScheduler contract.
type MetaSchedulerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MetaScheduler *MetaSchedulerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MetaSchedulerInitializedIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerInitializedIterator{contract: _MetaScheduler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MetaScheduler *MetaSchedulerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MetaSchedulerInitialized) (event.Subscription, error) {

	logs, sub, err := _MetaScheduler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaSchedulerInitialized)
				if err := _MetaScheduler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MetaScheduler *MetaSchedulerFilterer) ParseInitialized(log types.Log) (*MetaSchedulerInitialized, error) {
	event := new(MetaSchedulerInitialized)
	if err := _MetaScheduler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"HardwareUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"ToBeApproved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"decJobCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"name\":\"getProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pointPrevNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointNextNode\",\"type\":\"uint256\"}],\"internalType\":\"structProvider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getProviderFromAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pointPrevNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointNextNode\",\"type\":\"uint256\"}],\"internalType\":\"structProvider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasBeenKicked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasJoined\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"incJobCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"joinGrid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pointPrevNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointNextNode\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last_used\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pointPrevNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointNextNode\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pointPrevNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointNextNode\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"providersLinkedList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pointPrevNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointNextNode\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nNodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_memPricePerMin\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalJobCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nNodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_memPricePerMin\",\"type\":\"uint64\"}],\"name\":\"updateHardware\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50600060018054906101000a900460ff1615905080801562000044575060018060009054906101000a900460ff1660ff16105b8062000081575062000061306200016360201b62002a9d1760201c565b15801562000080575060018060009054906101000a900460ff1660ff16145b5b620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba906200020d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015620001005760018060016101000a81548160ff0219169083151502179055505b80156200015c5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405162000153919062000289565b60405180910390a15b50620002a6565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620001f5602e8362000186565b9150620002028262000197565b604082019050919050565b600060208201905081810360008301526200022881620001e6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620002716200026b62000265846200022f565b62000246565b62000239565b9050919050565b620002838162000250565b82525050565b6000602082019050620002a0600083018462000278565b92915050565b61418d80620002b66000396000f3fe608060405234801561001057600080fd5b50600436106101a85760003560e01c8063624bc8e3116100f9578063a94aa1a611610097578063bb26e6e811610071578063bb26e6e8146104f9578063d547741f14610515578063daea85c514610531578063e7692d091461054d576101a8565b8063a94aa1a614610477578063b7bb9145146104ad578063b7c8116e146104dd576101a8565b806391d14854116100d357806391d14854146103ef578063939daf9c1461041f57806396c551751461043d578063a217fddf14610459576101a8565b8063624bc8e3146103855780638129fc1c146103b5578063877f4e12146103bf576101a8565b80632f2ff15d1161016657806347799da81161014057806347799da8146102f75780634b939bd41461031b5780635c42d079146103395780635e9da3ed14610369576101a8565b80632f2ff15d146102a357806336568abe146102bf5780633f6edb5f146102db576101a8565b80623052a6146101ad57806301ffc9a7146101cb57806306661abd146101fb5780630787bc27146102195780630fc5d8031461024f578063248a9ca314610273575b600080fd5b6101b5610569565b6040516101c2919061314f565b60405180910390f35b6101e560048036038101906101e091906131c7565b610573565b6040516101f2919061320f565b60405180910390f35b6102036105ed565b604051610210919061314f565b60405180910390f35b610233600480360381019061022e9190613288565b6105f3565b60405161024697969594939291906133fb565b60405180910390f35b6102576107eb565b60405161026a97969594939291906133fb565b60405180910390f35b61028d600480360381019061028891906134a6565b6109d1565b60405161029a91906134e2565b60405180910390f35b6102bd60048036038101906102b891906134fd565b6109f0565b005b6102d960048036038101906102d491906134fd565b610a11565b005b6102f560048036038101906102f09190613288565b610a94565b005b6102ff610b7e565b60405161031297969594939291906133fb565b60405180910390f35b610323610d64565b604051610330919061314f565b60405180910390f35b610353600480360381019061034e9190613569565b610d6a565b60405161036091906136f4565b60405180910390f35b610383600480360381019061037e919061373c565b611065565b005b61039f600480360381019061039a9190613288565b611559565b6040516103ac91906136f4565b60405180910390f35b6103bd61186b565b005b6103d960048036038101906103d49190613288565b6119e8565b6040516103e6919061320f565b60405180910390f35b610409600480360381019061040491906134fd565b611a68565b604051610416919061320f565b60405180910390f35b610427611ad2565b60405161043491906137f2565b60405180910390f35b61045760048036038101906104529190613288565b611aec565b005b610461611bef565b60405161046e91906134e2565b60405180910390f35b610491600480360381019061048c9190613569565b611bf6565b6040516104a497969594939291906133fb565b60405180910390f35b6104c760048036038101906104c29190613288565b611dee565b6040516104d4919061320f565b60405180910390f35b6104f760048036038101906104f2919061373c565b611e6d565b005b610513600480360381019061050e9190613832565b6122ab565b005b61052f600480360381019061052a91906134fd565b6123fd565b005b61054b60048036038101906105469190613288565b61241e565b005b61056760048036038101906105629190613288565b612999565b005b6000601154905090565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806105e657506105e582612ac0565b5b9050919050565b60125481565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b60098060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b6000806000838152602001908152602001600020600101549050919050565b6109f9826109d1565b610a0281612b2a565b610a0c8383612b3e565b505050565b610a19612c1e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7d906138f5565b60405180910390fd5b610a908282612c26565b5050565b610aa16000801b33612d07565b600f600081819054906101000a900467ffffffffffffffff1680929190610ac790613944565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600281819054906101000a900467ffffffffffffffff1680929190610b5490613944565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b60038060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b60115481565b610d7261303c565b600115156010600084815260200190815260200160002060030160019054906101000a900460ff16151514610ddc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd3906139c0565b60405180910390fd5b601060008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020016003820160009054906101000a900460ff166002811115610fe757610fe6613375565b5b6002811115610ff957610ff8613375565b5b81526020016003820160019054906101000a900460ff161515151581526020016003820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b6110726000801b33612d07565b60011515600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff16151514611108576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ff90613a2c565b60405180910390fd5b6001600281111561111c5761111b613375565b5b600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff16600281111561117e5761117d613375565b5b146111be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b590613a98565b60405180910390fd5b60006040518060e001604052808967ffffffffffffffff1681526020018867ffffffffffffffff1681526020018567ffffffffffffffff1681526020018767ffffffffffffffff1681526020018467ffffffffffffffff1681526020018667ffffffffffffffff1681526020018367ffffffffffffffff1681525090506040518060e001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018281526020016000600281111561127c5761127b613375565b5b8152602001600115158152602001600067ffffffffffffffff168152602001600081526020016000815250600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505060408201518160030160006101000a81548160ff021916908360028111156114ac576114ab613375565b5b021790555060608201518160030160016101000a81548160ff02191690831515021790555060808201518160030160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816004015560c082015181600501559050507f3aeb53b0dee89ac04567fa6305e626e8d5246b478acd34d0a217507b9dfd076c896040516115469190613ab8565b60405180910390a1505050505050505050565b61156161303c565b61156e6000801b33612d07565b611577826119e8565b6115b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115ad90613b45565b60405180910390fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020016003820160009054906101000a900460ff1660028111156117ed576117ec613375565b5b60028111156117ff576117fe613375565b5b81526020016003820160019054906101000a900460ff161515151581526020016003820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b600060018054906101000a900460ff1615905080801561189c575060018060009054906101000a900460ff1660ff16105b806118ca57506118ab30612a9d565b1580156118c9575060018060009054906101000a900460ff1660ff16145b5b611909576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161190090613bd7565b60405180910390fd5b60018060006101000a81548160ff021916908360ff16021790555080156119455760018060016101000a81548160ff0219169083151502179055505b6119526000801b33612b3e565b6000600f60006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000601181905550600160128190555080156119e55760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516119dc9190613c49565b60405180910390a15b50565b6000600160028111156119fe576119fd613375565b5b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff166002811115611a6057611a5f613375565b5b149050919050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600f60009054906101000a900467ffffffffffffffff1681565b60011515600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff16151514611b82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b79906139c0565b60405180910390fd5b60028060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548160ff02191690836002811115611be757611be6613375565b5b021790555050565b6000801b81565b60106020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b6000600280811115611e0357611e02613375565b5b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff166002811115611e6557611e64613375565b5b149050919050565b611e7a6000801b33612d07565b60011515600260008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff16151503611f10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f0790613cb0565b60405180910390fd5b60006040518060e001604052808967ffffffffffffffff1681526020018867ffffffffffffffff1681526020018567ffffffffffffffff1681526020018767ffffffffffffffff1681526020018467ffffffffffffffff1681526020018667ffffffffffffffff1681526020018367ffffffffffffffff1681525090506040518060e001604052808a73ffffffffffffffffffffffffffffffffffffffff16815260200182815260200160006002811115611fce57611fcd613375565b5b8152602001600115158152602001600067ffffffffffffffff168152602001600081526020016000815250600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505060408201518160030160006101000a81548160ff021916908360028111156121fe576121fd613375565b5b021790555060608201518160030160016101000a81548160ff02191690831515021790555060808201518160030160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816004015560c082015181600501559050507fc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce2875896040516122989190613ab8565b60405180910390a1505050505050505050565b6122b86000801b33612d07565b600160068111156122cc576122cb613375565b5b8260068111156122df576122de613375565b5b1461231f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161231690613d42565b60405180910390fd5b600f600081819054906101000a900467ffffffffffffffff168092919061234590613d62565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600281819054906101000a900467ffffffffffffffff16809291906123d290613d62565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b612406826109d1565b61240f81612b2a565b6124198383612c26565b505050565b61242b6000801b33612d07565b60011515600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff161515146124c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124b8906139c0565b60405180910390fd5b600060028111156124d5576124d4613375565b5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff16600281111561253757612536613375565b5b14612577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161256e90613dd7565b60405180910390fd5b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548160ff021916908360028111156125dd576125dc613375565b5b02179055506000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600501540361299657600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206010600060125481526020019081526020016000206000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018201816001016000820160009054906101000a900467ffffffffffffffff168160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160089054906101000a900467ffffffffffffffff168160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160109054906101000a900467ffffffffffffffff168160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160189054906101000a900467ffffffffffffffff168160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160089054906101000a900467ffffffffffffffff168160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160109054906101000a900467ffffffffffffffff168160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050506003820160009054906101000a900460ff168160030160006101000a81548160ff021916908360028111156128fe576128fd613375565b5b02179055506003820160019054906101000a900460ff168160030160016101000a81548160ff0219169083151502179055506003820160029054906101000a900467ffffffffffffffff168160030160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820154816004015560058201548160050155905050612995601254612da4565b5b50565b60011515600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff16151514612a2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a26906139c0565b60405180910390fd5b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548160ff02191690836002811115612a9557612a94613375565b5b021790555050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b612b3b81612b36612c1e565b612d07565b50565b612b488282611a68565b612c1a57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612bbf612c1e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b612c308282611a68565b15612d0357600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612ca8612c1e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b612d118282611a68565b612da057612d368173ffffffffffffffffffffffffffffffffffffffff166014612e00565b612d448360001c6020612e00565b604051602001612d55929190613f00565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d979190613f84565b60405180910390fd5b5050565b806010600060115481526020019081526020016000206005018190555060115460106000838152602001908152602001600020600401819055508060118190555060126000815480929190612df890613fa6565b919050555050565b606060006002836002612e139190613fee565b612e1d9190614030565b67ffffffffffffffff811115612e3657612e35614064565b5b6040519080825280601f01601f191660200182016040528015612e685781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110612ea057612e9f614093565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110612f0457612f03614093565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002612f449190613fee565b612f4e9190614030565b90505b6001811115612fee577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110612f9057612f8f614093565b5b1a60f81b828281518110612fa757612fa6614093565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080612fe7906140c2565b9050612f51565b5060008414613032576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161302990614137565b60405180910390fd5b8091505092915050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200161306c6130b3565b81526020016000600281111561308557613084613375565b5b8152602001600015158152602001600067ffffffffffffffff16815260200160008152602001600081525090565b6040518060e00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000819050919050565b61314981613136565b82525050565b60006020820190506131646000830184613140565b92915050565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6131a48161316f565b81146131af57600080fd5b50565b6000813590506131c18161319b565b92915050565b6000602082840312156131dd576131dc61316a565b5b60006131eb848285016131b2565b91505092915050565b60008115159050919050565b613209816131f4565b82525050565b60006020820190506132246000830184613200565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006132558261322a565b9050919050565b6132658161324a565b811461327057600080fd5b50565b6000813590506132828161325c565b92915050565b60006020828403121561329e5761329d61316a565b5b60006132ac84828501613273565b91505092915050565b6132be8161324a565b82525050565b600067ffffffffffffffff82169050919050565b6132e1816132c4565b82525050565b60e0820160008201516132fd60008501826132d8565b50602082015161331060208501826132d8565b50604082015161332360408501826132d8565b50606082015161333660608501826132d8565b50608082015161334960808501826132d8565b5060a082015161335c60a08501826132d8565b5060c082015161336f60c08501826132d8565b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106133b5576133b4613375565b5b50565b60008190506133c6826133a4565b919050565b60006133d6826133b8565b9050919050565b6133e6816133cb565b82525050565b6133f5816132c4565b82525050565b60006101a082019050613411600083018a6132b5565b61341e60208301896132e7565b61342c6101008301886133dd565b61343a610120830187613200565b6134486101408301866133ec565b613456610160830185613140565b613464610180830184613140565b98975050505050505050565b6000819050919050565b61348381613470565b811461348e57600080fd5b50565b6000813590506134a08161347a565b92915050565b6000602082840312156134bc576134bb61316a565b5b60006134ca84828501613491565b91505092915050565b6134dc81613470565b82525050565b60006020820190506134f760008301846134d3565b92915050565b600080604083850312156135145761351361316a565b5b600061352285828601613491565b925050602061353385828601613273565b9150509250929050565b61354681613136565b811461355157600080fd5b50565b6000813590506135638161353d565b92915050565b60006020828403121561357f5761357e61316a565b5b600061358d84828501613554565b91505092915050565b61359f8161324a565b82525050565b60e0820160008201516135bb60008501826132d8565b5060208201516135ce60208501826132d8565b5060408201516135e160408501826132d8565b5060608201516135f460608501826132d8565b50608082015161360760808501826132d8565b5060a082015161361a60a08501826132d8565b5060c082015161362d60c08501826132d8565b50505050565b61363c816133cb565b82525050565b61364b816131f4565b82525050565b61365a81613136565b82525050565b6101a0820160008201516136776000850182613596565b50602082015161368a60208501826135a5565b50604082015161369e610100850182613633565b5060608201516136b2610120850182613642565b5060808201516136c66101408501826132d8565b5060a08201516136da610160850182613651565b5060c08201516136ee610180850182613651565b50505050565b60006101a08201905061370a6000830184613660565b92915050565b613719816132c4565b811461372457600080fd5b50565b60008135905061373681613710565b92915050565b600080600080600080600080610100898b03121561375d5761375c61316a565b5b600061376b8b828c01613273565b985050602061377c8b828c01613727565b975050604061378d8b828c01613727565b965050606061379e8b828c01613727565b95505060806137af8b828c01613727565b94505060a06137c08b828c01613727565b93505060c06137d18b828c01613727565b92505060e06137e28b828c01613727565b9150509295985092959890939650565b600060208201905061380760008301846133ec565b92915050565b6007811061381a57600080fd5b50565b60008135905061382c8161380d565b92915050565b600080604083850312156138495761384861316a565b5b60006138578582860161381d565b925050602061386885828601613273565b9150509250929050565b600082825260208201905092915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b60006138df602f83613872565b91506138ea82613883565b604082019050919050565b6000602082019050818103600083015261390e816138d2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061394f826132c4565b915067ffffffffffffffff820361396957613968613915565b5b600182019050919050565b7f4e6f2070726f766964657220666f756e64000000000000000000000000000000600082015250565b60006139aa601183613872565b91506139b582613974565b602082019050919050565b600060208201905081810360008301526139d98161399d565b9050919050565b7f50726f7669646572206d75737420626520726567697374657265640000000000600082015250565b6000613a16601b83613872565b9150613a21826139e0565b602082019050919050565b60006020820190508181036000830152613a4581613a09565b9050919050565b7f50726f76696465722068617665206265656e20617070726f7665640000000000600082015250565b6000613a82601b83613872565b9150613a8d82613a4c565b602082019050919050565b60006020820190508181036000830152613ab181613a75565b9050919050565b6000602082019050613acd60008301846132b5565b92915050565b7f50726f766964657273206e6f7420666f756e64206f72206e6f74206a6f696e6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b6000613b2f602183613872565b9150613b3a82613ad3565b604082019050919050565b60006020820190508181036000830152613b5e81613b22565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000613bc1602e83613872565b9150613bcc82613b65565b604082019050919050565b60006020820190508181036000830152613bf081613bb4565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000613c33613c2e613c2984613bf7565b613c0e565b613c01565b9050919050565b613c4381613c18565b82525050565b6000602082019050613c5e6000830184613c3a565b92915050565b7f50726f76696465722063616e206e6f7420726567697374657220747769636500600082015250565b6000613c9a601f83613872565b9150613ca582613c64565b602082019050919050565b60006020820190508181036000830152613cc981613c8d565b9050919050565b7f4f6e6c79204d4554415f5343484544554c4544206a6f62732063616e2062652060008201527f6465637265617365640000000000000000000000000000000000000000000000602082015250565b6000613d2c602983613872565b9150613d3782613cd0565b604082019050919050565b60006020820190508181036000830152613d5b81613d1f565b9050919050565b6000613d6d826132c4565b915060008203613d8057613d7f613915565b5b600182039050919050565b7f50726f766964657220737461747573206d75737420626520494e495400000000600082015250565b6000613dc1601c83613872565b9150613dcc82613d8b565b602082019050919050565b60006020820190508181036000830152613df081613db4565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000613e38601783613df7565b9150613e4382613e02565b601782019050919050565b600081519050919050565b60005b83811015613e77578082015181840152602081019050613e5c565b60008484015250505050565b6000613e8e82613e4e565b613e988185613df7565b9350613ea8818560208601613e59565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000613eea601183613df7565b9150613ef582613eb4565b601182019050919050565b6000613f0b82613e2b565b9150613f178285613e83565b9150613f2282613edd565b9150613f2e8284613e83565b91508190509392505050565b6000601f19601f8301169050919050565b6000613f5682613e4e565b613f608185613872565b9350613f70818560208601613e59565b613f7981613f3a565b840191505092915050565b60006020820190508181036000830152613f9e8184613f4b565b905092915050565b6000613fb182613136565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613fe357613fe2613915565b5b600182019050919050565b6000613ff982613136565b915061400483613136565b925082820261401281613136565b9150828204841483151761402957614028613915565b5b5092915050565b600061403b82613136565b915061404683613136565b925082820190508082111561405e5761405d613915565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006140cd82613136565b9150600082036140e0576140df613915565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000614121602083613872565b915061412c826140eb565b602082019050919050565b6000602082019050818103600083015261415081614114565b905091905056fea2646970667358221220a60116a7251f1349fb12c51ceaa74b31e331ec5501421e00355d9979a83349c964736f6c63430008110033",
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

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_ProviderManager *ProviderManagerCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_ProviderManager *ProviderManagerSession) Count() (*big.Int, error) {
	return _ProviderManager.Contract.Count(&_ProviderManager.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_ProviderManager *ProviderManagerCallerSession) Count() (*big.Int, error) {
	return _ProviderManager.Contract.Count(&_ProviderManager.CallOpts)
}

// GetProvider is a free data retrieval call binding the contract method 0x5c42d079.
//
// Solidity: function getProvider(uint256 counter) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
func (_ProviderManager *ProviderManagerCaller) GetProvider(opts *bind.CallOpts, counter *big.Int) (Provider, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "getProvider", counter)

	if err != nil {
		return *new(Provider), err
	}

	out0 := *abi.ConvertType(out[0], new(Provider)).(*Provider)

	return out0, err

}

// GetProvider is a free data retrieval call binding the contract method 0x5c42d079.
//
// Solidity: function getProvider(uint256 counter) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
func (_ProviderManager *ProviderManagerSession) GetProvider(counter *big.Int) (Provider, error) {
	return _ProviderManager.Contract.GetProvider(&_ProviderManager.CallOpts, counter)
}

// GetProvider is a free data retrieval call binding the contract method 0x5c42d079.
//
// Solidity: function getProvider(uint256 counter) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
func (_ProviderManager *ProviderManagerCallerSession) GetProvider(counter *big.Int) (Provider, error) {
	return _ProviderManager.Contract.GetProvider(&_ProviderManager.CallOpts, counter)
}

// GetProviderFromAddr is a free data retrieval call binding the contract method 0x624bc8e3.
//
// Solidity: function getProviderFromAddr(address _providerAddr) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
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
// Solidity: function getProviderFromAddr(address _providerAddr) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
func (_ProviderManager *ProviderManagerSession) GetProviderFromAddr(_providerAddr common.Address) (Provider, error) {
	return _ProviderManager.Contract.GetProviderFromAddr(&_ProviderManager.CallOpts, _providerAddr)
}

// GetProviderFromAddr is a free data retrieval call binding the contract method 0x624bc8e3.
//
// Solidity: function getProviderFromAddr(address _providerAddr) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
func (_ProviderManager *ProviderManagerCallerSession) GetProviderFromAddr(_providerAddr common.Address) (Provider, error) {
	return _ProviderManager.Contract.GetProviderFromAddr(&_ProviderManager.CallOpts, _providerAddr)
}

// GetProviderNumber is a free data retrieval call binding the contract method 0x003052a6.
//
// Solidity: function getProviderNumber() view returns(uint256)
func (_ProviderManager *ProviderManagerCaller) GetProviderNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "getProviderNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderNumber is a free data retrieval call binding the contract method 0x003052a6.
//
// Solidity: function getProviderNumber() view returns(uint256)
func (_ProviderManager *ProviderManagerSession) GetProviderNumber() (*big.Int, error) {
	return _ProviderManager.Contract.GetProviderNumber(&_ProviderManager.CallOpts)
}

// GetProviderNumber is a free data retrieval call binding the contract method 0x003052a6.
//
// Solidity: function getProviderNumber() view returns(uint256)
func (_ProviderManager *ProviderManagerCallerSession) GetProviderNumber() (*big.Int, error) {
	return _ProviderManager.Contract.GetProviderNumber(&_ProviderManager.CallOpts)
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

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCaller) Last(opts *bind.CallOpts) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "last")

	outstruct := new(struct {
		Addr          common.Address
		Definition    ProviderDefinition
		Status        uint8
		Valid         bool
		JobCount      uint64
		PointPrevNode *big.Int
		PointNextNode *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Definition = *abi.ConvertType(out[1], new(ProviderDefinition)).(*ProviderDefinition)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.JobCount = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.PointPrevNode = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PointNextNode = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerSession) Last() (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.Last(&_ProviderManager.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCallerSession) Last() (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.Last(&_ProviderManager.CallOpts)
}

// LastUsed is a free data retrieval call binding the contract method 0x0fc5d803.
//
// Solidity: function last_used() view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCaller) LastUsed(opts *bind.CallOpts) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "last_used")

	outstruct := new(struct {
		Addr          common.Address
		Definition    ProviderDefinition
		Status        uint8
		Valid         bool
		JobCount      uint64
		PointPrevNode *big.Int
		PointNextNode *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Definition = *abi.ConvertType(out[1], new(ProviderDefinition)).(*ProviderDefinition)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.JobCount = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.PointPrevNode = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PointNextNode = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LastUsed is a free data retrieval call binding the contract method 0x0fc5d803.
//
// Solidity: function last_used() view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerSession) LastUsed() (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.LastUsed(&_ProviderManager.CallOpts)
}

// LastUsed is a free data retrieval call binding the contract method 0x0fc5d803.
//
// Solidity: function last_used() view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCallerSession) LastUsed() (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.LastUsed(&_ProviderManager.CallOpts)
}

// ProviderNumber is a free data retrieval call binding the contract method 0x4b939bd4.
//
// Solidity: function providerNumber() view returns(uint256)
func (_ProviderManager *ProviderManagerCaller) ProviderNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "providerNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProviderNumber is a free data retrieval call binding the contract method 0x4b939bd4.
//
// Solidity: function providerNumber() view returns(uint256)
func (_ProviderManager *ProviderManagerSession) ProviderNumber() (*big.Int, error) {
	return _ProviderManager.Contract.ProviderNumber(&_ProviderManager.CallOpts)
}

// ProviderNumber is a free data retrieval call binding the contract method 0x4b939bd4.
//
// Solidity: function providerNumber() view returns(uint256)
func (_ProviderManager *ProviderManagerCallerSession) ProviderNumber() (*big.Int, error) {
	return _ProviderManager.Contract.ProviderNumber(&_ProviderManager.CallOpts)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCaller) Providers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "providers", arg0)

	outstruct := new(struct {
		Addr          common.Address
		Definition    ProviderDefinition
		Status        uint8
		Valid         bool
		JobCount      uint64
		PointPrevNode *big.Int
		PointNextNode *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Definition = *abi.ConvertType(out[1], new(ProviderDefinition)).(*ProviderDefinition)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.JobCount = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.PointPrevNode = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PointNextNode = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerSession) Providers(arg0 common.Address) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.Providers(&_ProviderManager.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCallerSession) Providers(arg0 common.Address) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.Providers(&_ProviderManager.CallOpts, arg0)
}

// ProvidersLinkedList is a free data retrieval call binding the contract method 0xa94aa1a6.
//
// Solidity: function providersLinkedList(uint256 ) view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCaller) ProvidersLinkedList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	var out []interface{}
	err := _ProviderManager.contract.Call(opts, &out, "providersLinkedList", arg0)

	outstruct := new(struct {
		Addr          common.Address
		Definition    ProviderDefinition
		Status        uint8
		Valid         bool
		JobCount      uint64
		PointPrevNode *big.Int
		PointNextNode *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Definition = *abi.ConvertType(out[1], new(ProviderDefinition)).(*ProviderDefinition)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.JobCount = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.PointPrevNode = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PointNextNode = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProvidersLinkedList is a free data retrieval call binding the contract method 0xa94aa1a6.
//
// Solidity: function providersLinkedList(uint256 ) view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerSession) ProvidersLinkedList(arg0 *big.Int) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.ProvidersLinkedList(&_ProviderManager.CallOpts, arg0)
}

// ProvidersLinkedList is a free data retrieval call binding the contract method 0xa94aa1a6.
//
// Solidity: function providersLinkedList(uint256 ) view returns(address addr, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) definition, uint8 status, bool valid, uint64 jobCount, uint256 pointPrevNode, uint256 pointNextNode)
func (_ProviderManager *ProviderManagerCallerSession) ProvidersLinkedList(arg0 *big.Int) (struct {
	Addr          common.Address
	Definition    ProviderDefinition
	Status        uint8
	Valid         bool
	JobCount      uint64
	PointPrevNode *big.Int
	PointNextNode *big.Int
}, error) {
	return _ProviderManager.Contract.ProvidersLinkedList(&_ProviderManager.CallOpts, arg0)
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

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactor) Approve(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "approve", _providerAddr)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerSession) Approve(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.Approve(&_ProviderManager.TransactOpts, _providerAddr)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address _providerAddr) returns()
func (_ProviderManager *ProviderManagerTransactorSession) Approve(_providerAddr common.Address) (*types.Transaction, error) {
	return _ProviderManager.Contract.Approve(&_ProviderManager.TransactOpts, _providerAddr)
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

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProviderManager *ProviderManagerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProviderManager *ProviderManagerSession) Initialize() (*types.Transaction, error) {
	return _ProviderManager.Contract.Initialize(&_ProviderManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProviderManager *ProviderManagerTransactorSession) Initialize() (*types.Transaction, error) {
	return _ProviderManager.Contract.Initialize(&_ProviderManager.TransactOpts)
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

// Register is a paid mutator transaction binding the contract method 0xb7c8116e.
//
// Solidity: function register(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerTransactor) Register(opts *bind.TransactOpts, _providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "register", _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// Register is a paid mutator transaction binding the contract method 0xb7c8116e.
//
// Solidity: function register(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerSession) Register(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.Register(&_ProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// Register is a paid mutator transaction binding the contract method 0xb7c8116e.
//
// Solidity: function register(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerTransactorSession) Register(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.Register(&_ProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
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

// UpdateHardware is a paid mutator transaction binding the contract method 0x5e9da3ed.
//
// Solidity: function updateHardware(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerTransactor) UpdateHardware(opts *bind.TransactOpts, _providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.contract.Transact(opts, "updateHardware", _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// UpdateHardware is a paid mutator transaction binding the contract method 0x5e9da3ed.
//
// Solidity: function updateHardware(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerSession) UpdateHardware(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.UpdateHardware(&_ProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// UpdateHardware is a paid mutator transaction binding the contract method 0x5e9da3ed.
//
// Solidity: function updateHardware(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_ProviderManager *ProviderManagerTransactorSession) UpdateHardware(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _ProviderManager.Contract.UpdateHardware(&_ProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// ProviderManagerHardwareUpdatedEventIterator is returned from FilterHardwareUpdatedEvent and is used to iterate over the raw logs and unpacked data for HardwareUpdatedEvent events raised by the ProviderManager contract.
type ProviderManagerHardwareUpdatedEventIterator struct {
	Event *ProviderManagerHardwareUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *ProviderManagerHardwareUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderManagerHardwareUpdatedEvent)
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
		it.Event = new(ProviderManagerHardwareUpdatedEvent)
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
func (it *ProviderManagerHardwareUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderManagerHardwareUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderManagerHardwareUpdatedEvent represents a HardwareUpdatedEvent event raised by the ProviderManager contract.
type ProviderManagerHardwareUpdatedEvent struct {
	ProviderAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHardwareUpdatedEvent is a free log retrieval operation binding the contract event 0x3aeb53b0dee89ac04567fa6305e626e8d5246b478acd34d0a217507b9dfd076c.
//
// Solidity: event HardwareUpdatedEvent(address _providerAddr)
func (_ProviderManager *ProviderManagerFilterer) FilterHardwareUpdatedEvent(opts *bind.FilterOpts) (*ProviderManagerHardwareUpdatedEventIterator, error) {

	logs, sub, err := _ProviderManager.contract.FilterLogs(opts, "HardwareUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return &ProviderManagerHardwareUpdatedEventIterator{contract: _ProviderManager.contract, event: "HardwareUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchHardwareUpdatedEvent is a free log subscription operation binding the contract event 0x3aeb53b0dee89ac04567fa6305e626e8d5246b478acd34d0a217507b9dfd076c.
//
// Solidity: event HardwareUpdatedEvent(address _providerAddr)
func (_ProviderManager *ProviderManagerFilterer) WatchHardwareUpdatedEvent(opts *bind.WatchOpts, sink chan<- *ProviderManagerHardwareUpdatedEvent) (event.Subscription, error) {

	logs, sub, err := _ProviderManager.contract.WatchLogs(opts, "HardwareUpdatedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderManagerHardwareUpdatedEvent)
				if err := _ProviderManager.contract.UnpackLog(event, "HardwareUpdatedEvent", log); err != nil {
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
func (_ProviderManager *ProviderManagerFilterer) ParseHardwareUpdatedEvent(log types.Log) (*ProviderManagerHardwareUpdatedEvent, error) {
	event := new(ProviderManagerHardwareUpdatedEvent)
	if err := _ProviderManager.contract.UnpackLog(event, "HardwareUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProviderManager contract.
type ProviderManagerInitializedIterator struct {
	Event *ProviderManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ProviderManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderManagerInitialized)
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
		it.Event = new(ProviderManagerInitialized)
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
func (it *ProviderManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderManagerInitialized represents a Initialized event raised by the ProviderManager contract.
type ProviderManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProviderManager *ProviderManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProviderManagerInitializedIterator, error) {

	logs, sub, err := _ProviderManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProviderManagerInitializedIterator{contract: _ProviderManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProviderManager *ProviderManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProviderManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ProviderManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderManagerInitialized)
				if err := _ProviderManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ProviderManager *ProviderManagerFilterer) ParseInitialized(log types.Log) (*ProviderManagerInitialized, error) {
	event := new(ProviderManagerInitialized)
	if err := _ProviderManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220932d324ab357f63d09c477174691a5b74f30abff9d48f5732dfc487f5967b4e964736f6c63430008110033",
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

// TimeoutManagementMetaData contains all meta data concerning the TimeoutManagement contract.
var TimeoutManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockOrigin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"stillAlive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101cd610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806315a945ce1461003a575b600080fd5b610054600480360381019061004f91906100be565b61006a565b6040516100619190610119565b60405180910390f35b60008183436100799190610163565b1015905092915050565b600080fd5b6000819050919050565b61009b81610088565b81146100a657600080fd5b50565b6000813590506100b881610092565b92915050565b600080604083850312156100d5576100d4610083565b5b60006100e3858286016100a9565b92505060206100f4858286016100a9565b9150509250929050565b60008115159050919050565b610113816100fe565b82525050565b600060208201905061012e600083018461010a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061016e82610088565b915061017983610088565b925082820390508181111561019157610190610134565b5b9291505056fea264697066735822122087491dbccae6987c8f076dc29dcaeb4db9bd5cbd3bfc9c1d961de45817cf3b2164736f6c63430008110033",
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
