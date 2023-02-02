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
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Start                  *big.Int
	End                    *big.Int
	JobName                [32]byte
	Cost                   *big.Int
}

// JobDefinition is an auto generated low-level Go binding around an user-defined struct.
type JobDefinition struct {
	GpuPerTask        uint64
	MemPerCpu         uint64
	CpuPerTask        uint64
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJobError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"_previousStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"checkNewJobStatus\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimJobTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimNextCancellingJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimNextJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobFromId\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"internalType\":\"structJob\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasCancellingJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasNextJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotJobList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metaQueue\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerCancelledJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimFails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimableJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerScheduledJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerStartFails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_jobCost\",\"type\":\"uint256\"}],\"name\":\"setJobCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"}],\"name\":\"setJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateJobsStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallet2JobId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50600060018054906101000a900460ff1615905080801562000044575060018060009054906101000a900460ff1660ff16105b8062000081575062000061306200016360201b620030651760201c565b15801562000080575060018060009054906101000a900460ff1660ff16145b5b620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba906200020d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015620001005760018060016101000a81548160ff0219169083151502179055505b80156200015c5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405162000153919062000289565b60405180910390a15b50620002a6565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620001f5602e8362000186565b9150620002028262000197565b604082019050919050565b600060208201905081810360008301526200022881620001e6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620002716200026b62000265846200022f565b62000246565b62000239565b9050919050565b620002838162000250565b82525050565b6000602082019050620002a0600083018462000278565b92915050565b615ca380620002b66000396000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80638129fc1c11610125578063c58467b0116100ad578063d6aa37a61161007c578063d6aa37a6146106a4578063d8e3aac7146106c2578063de2418b3146106de578063e19086761461070e578063e88fe8ca1461072a57610211565b8063c58467b01461060a578063d1cee5461461063b578063d43a802714610657578063d547741f1461068857610211565b8063ade197b4116100f4578063ade197b414610542578063aef3276f14610572578063b3130fba146105a2578063b42f5513146105d2578063b67644b9146105ee57610211565b80638129fc1c146104ba5780638fb70f63146104c457806391d14854146104f4578063a217fddf1461052457610211565b80632a242a76116101a857806342308ee41161017757806342308ee4146104025780634609ca501461043357806346200b6b1461044f57806347ce0a261461046e5780637ff67e291461049e57610211565b80632a242a76146103855780632f2ff15d1461038f57806336568abe146103ab57806338ed7cfc146103c757610211565b80631a3cbef4116101e45780631a3cbef4146102c55780631df4e610146102f557806320a5f91914610325578063248a9ca31461035557610211565b806301ffc9a7146102165780630797094e146102465780630f5a11ca14610276578063110e87a614610294575b600080fd5b610230600480360381019061022b9190613e2f565b61075a565b60405161023d9190613e77565b60405180910390f35b610260600480360381019061025b9190613ef0565b6107d4565b60405161026d9190613e77565b60405180910390f35b61027e610826565b60405161028b9190613f36565b60405180910390f35b6102ae60048036038101906102a99190613ef0565b61082c565b6040516102bc929190613f6d565b60405180910390f35b6102df60048036038101906102da9190613ef0565b61086a565b6040516102ec919061405e565b60405180910390f35b61030f600480360381019061030a9190614310565b610901565b60405161031c91906143a2565b60405180910390f35b61033f600480360381019061033a9190613ef0565b610d83565b60405161034c9190613e77565b60405180910390f35b61036f600480360381019061036a91906143bd565b610dd5565b60405161037c91906143a2565b60405180910390f35b61038d610df4565b005b6103a960048036038101906103a491906143ea565b611169565b005b6103c560048036038101906103c091906143ea565b61118a565b005b6103e160048036038101906103dc91906143bd565b61120d565b6040516103f99c9b9a999897969594939291906145b4565b60405180910390f35b61041c60048036038101906104179190613ef0565b61142b565b60405161042a929190613f6d565b60405180910390f35b61044d600480360381019061044891906143ea565b611469565b005b6104576115c0565b604051610465929190613f6d565b60405180910390f35b61048860048036038101906104839190613ef0565b6115ec565b6040516104959190613f36565b60405180910390f35b6104b860048036038101906104b3919061469a565b611604565b005b6104c26118d5565b005b6104de60048036038101906104d991906143ea565b611a42565b6040516104eb91906146ed565b60405180910390f35b61050e600480360381019061050991906143ea565b61208f565b60405161051b9190613e77565b60405180910390f35b61052c6120f9565b60405161053991906143a2565b60405180910390f35b61055c60048036038101906105579190613ef0565b612100565b6040516105699190613f36565b60405180910390f35b61058c6004803603810190610587919061470f565b612118565b60405161059991906143a2565b60405180910390f35b6105bc60048036038101906105b791906143bd565b61213c565b6040516105c991906148f2565b60405180910390f35b6105ec60048036038101906105e79190614914565b612482565b005b61060860048036038101906106039190614954565b61280a565b005b610624600480360381019061061f9190613ef0565b6129d3565b604051610632929190614994565b60405180910390f35b610655600480360381019061065091906143ea565b612a4d565b005b610671600480360381019061066c9190613ef0565b612c65565b60405161067f929190613f6d565b60405180910390f35b6106a2600480360381019061069d91906143ea565b612ca3565b005b6106ac612cc4565b6040516106b991906149d3565b60405180910390f35b6106dc60048036038101906106d79190614954565b612cde565b005b6106f860048036038101906106f39190613ef0565b612d59565b60405161070591906143a2565b60405180910390f35b610728600480360381019061072391906149ee565b612db9565b005b610744600480360381019061073f9190614a2e565b613034565b60405161075191906143a2565b60405180910390f35b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806107cd57506107cc82613088565b5b9050919050565b600061081f6001600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206130f2565b9050919050565b600d5481565b60046020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6060600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156108f557602002820191906000526020600020905b8154815260200190600101908083116108e1575b50505050509050919050565b60008060001b61091081613192565b6000856060015167ffffffffffffffff1611610961576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095890614acb565b60405180910390fd5b6000856040015167ffffffffffffffff16116109b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a990614b37565b60405180910390fd5b6000856020015167ffffffffffffffff1611610a03576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109fa90614ba3565b60405180910390fd5b600d6000815480929190610a1690614bf2565b91905055506000600d5460001b905060405180610180016040528082815260200160006008811115610a4b57610a4a61442a565b5b81526020018681526020018873ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001878152602001600115158152602001600081526020014281526020014281526020018581526020016000815250600260008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff02191690836008811115610b0457610b0361442a565b5b02179055506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816001019081610c7b9190614e46565b50505060c08201518160070160006101000a81548160ff02191690831515021790555060e08201518160080155610100820151816009015561012082015181600a015561014082015181600b015561016082015181600c0155905050600960008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150557f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada38188604051610d6e929190614f18565b60405180910390a18092505050949350505050565b6000610dce6004600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206130f2565b9050919050565b6000806000838152602001908152602001600020600101549050919050565b6000801b610e0181613192565b60005b6003805490508167ffffffffffffffff16101561116557600060038267ffffffffffffffff1681548110610e3b57610e3a614f41565b5b906000526020600020015490506000801b8103610e585750611152565b60016008811115610e6c57610e6b61442a565b5b6002600083815260200190815260200160002060010160009054906101000a900460ff166008811115610ea257610ea161442a565b5b14158015610ef8575060026008811115610ebf57610ebe61442a565b5b6002600083815260200190815260200160002060010160009054906101000a900460ff166008811115610ef557610ef461442a565b5b14155b15610f33576000801b60038367ffffffffffffffff1681548110610f1f57610f1e614f41565b5b906000526020600020018190555050611152565b610f5381600c60009054906101000a900467ffffffffffffffff166131a6565b156111505760006002600083815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000801b60038467ffffffffffffffff1681548110610fb557610fb4614f41565b5b906000526020600020018190555061100c6000600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061323b565b6110556000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061323b565b600060088111156110695761106861442a565b5b6002600084815260200190815260200160002060010160009054906101000a900460ff16600881111561109f5761109e61442a565b5b146110df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d690614fbc565b60405180910390fd5b7f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3826002600085815260200190815260200160002060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051611146929190614f18565b60405180910390a1505b505b808061115d90614fdc565b915050610e04565b5050565b61117282610dd5565b61117b81613192565b611185838361329e565b505050565b61119261337e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146111ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f69061507e565b60405180910390fd5b6112098282613386565b5050565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806005016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461137390614c69565b80601f016020809104026020016040519081016040528092919081815260200182805461139f90614c69565b80156113ec5780601f106113c1576101008083540402835291602001916113ec565b820191906000526020600020905b8154815290600101906020018083116113cf57829003601f168201915b505050505081525050908060070160009054906101000a900460ff169080600801549080600901549080600a01549080600b01549080600c015490508c565b60056020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6000801b61147681613192565b6114808383613467565b600260088111156114945761149361442a565b5b6002600085815260200190815260200160002060010160009054906101000a900460ff1660088111156114ca576114c961442a565b5b1461150a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611501906150ea565b60405180910390fd5b60006002600085815260200190815260200160002060010160006101000a81548160ff021916908360088111156115445761154361442a565b5b02179055507f5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0838360405161157a929190614f18565b60405180910390a17f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada383836040516115b3929190614f18565b60405180910390a1505050565b600a8060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b60086020528060005260406000206000915090505481565b6000801b61161181613192565b600060088111156116255761162461442a565b5b8260088111156116385761163761442a565b5b03611678576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161166f90615156565b60405180910390fd5b6002600881111561168c5761168b61442a565b5b82600881111561169f5761169e61442a565b5b036116df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d6906151c2565b60405180910390fd5b600460088111156116f3576116f261442a565b5b8260088111156117065761170561442a565b5b03611746576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161173d9061522e565b60405180910390fd5b6117508484613467565b61177d6002600086815260200190815260200160002060010160009054906101000a900460ff1683612482565b816002600086815260200190815260200160002060010160006101000a81548160ff021916908360088111156117b6576117b561442a565b5b0217905550600260088111156117cf576117ce61442a565b5b6002600086815260200190815260200160002060010160009054906101000a900460ff1660088111156118055761180461442a565b5b0361182a57426002600086815260200190815260200160002060090181905550611846565b4260026000868152602001908152602001600020600a01819055505b8160088111156118595761185861442a565b5b6002600086815260200190815260200160002060010160009054906101000a900460ff16600881111561188f5761188e61442a565b5b146118cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118c69061529a565b60405180910390fd5b50505050565b600060018054906101000a900460ff16159050808015611906575060018060009054906101000a900460ff1660ff16105b80611934575061191530613065565b158015611933575060018060009054906101000a900460ff1660ff16145b5b611973576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196a9061532c565b60405180910390fd5b60018060006101000a81548160ff021916908360ff16021790555080156119af5760018060016101000a81548160ff0219169083151502179055505b6119bc6000801b3361329e565b600a600c60006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508015611a3f5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051611a369190615394565b60405180910390a15b50565b611a4a613cbf565b826002600082815260200190815260200160002060070160009054906101000a900460ff16611aa5576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b611ab281613192565b60006002600087815260200190815260200160002060405180610180016040529081600082015481526020016001820160009054906101000a900460ff166008811115611b0257611b0161442a565b5b6008811115611b1457611b1361442a565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054611cb690614c69565b80601f0160208091040260200160405190810160405280929190818152602001828054611ce290614c69565b8015611d2f5780601f10611d0457610100808354040283529160200191611d2f565b820191906000526020600020905b815481529060010190602001808311611d1257829003601f168201915b50505050508152505081526020016007820160009054906101000a900460ff161515151581526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152505090508473ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff1614611dfd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611df4906153fb565b60405180910390fd5b60016008811115611e1157611e1061442a565b5b81602001516008811115611e2857611e2761442a565b5b14611e68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e5f90615467565b60405180910390fd5b600280600088815260200190815260200160002060010160006101000a81548160ff02191690836008811115611ea157611ea061442a565b5b0217905550436002600088815260200190815260200160002060080181905550611f09600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002087613579565b600260008781526020019081526020016000206005016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461200290614c69565b80601f016020809104026020016040519081016040528092919081815260200182805461202e90614c69565b801561207b5780601f106120505761010080835404028352916020019161207b565b820191906000526020600020905b81548152906001019060200180831161205e57829003601f168201915b505050505081525050935050505092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b60076020528060005260406000206000915090505481565b6003818154811061212857600080fd5b906000526020600020016000915090505481565b612144613d16565b816002600082815260200190815260200160002060070160009054906101000a900460ff1661219f576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006002600085815260200190815260200160002060405180610180016040529081600082015481526020016001820160009054906101000a900460ff1660088111156121ef576121ee61442a565b5b60088111156122015761220061442a565b5b8152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546123a390614c69565b80601f01602080910402602001604051908101604052809291908181526020018280546123cf90614c69565b801561241c5780601f106123f15761010080835404028352916020019161241c565b820191906000526020600020905b8154815290600101906020018083116123ff57829003601f168201915b50505050508152505081526020016007820160009054906101000a900460ff161515151581526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152505090508092505050919050565b600360088111156124965761249561442a565b5b8160088111156124a9576124a861442a565b5b0361251a57600260088111156124c2576124c161442a565b5b8260088111156124d5576124d461442a565b5b14612515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161250c906154d3565b60405180910390fd5b612806565b6005600881111561252e5761252d61442a565b5b8160088111156125415761254061442a565b5b036125b2576004600881111561255a5761255961442a565b5b82600881111561256d5761256c61442a565b5b146125ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125a49061553f565b60405180910390fd5b612805565b600660088111156125c6576125c561442a565b5b8160088111156125d9576125d861442a565b5b0361267957600360088111156125f2576125f161442a565b5b8260088111156126055761260461442a565b5b14806126355750600260088111156126205761261f61442a565b5b8260088111156126335761263261442a565b5b145b612674576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161266b906155ab565b60405180910390fd5b612804565b6007600881111561268d5761268c61442a565b5b8160088111156126a05761269f61442a565b5b0361274057600360088111156126b9576126b861442a565b5b8260088111156126cc576126cb61442a565b5b14806126fc5750600260088111156126e7576126e661442a565b5b8260088111156126fa576126f961442a565b5b145b61273b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161273290615617565b60405180910390fd5b612803565b6008808111156127535761275261442a565b5b8160088111156127665761276561442a565b5b03612802576003600881111561277f5761277e61442a565b5b8260088111156127925761279161442a565b5b14806127c25750600260088111156127ad576127ac61442a565b5b8260088111156127c0576127bf61442a565b5b145b612801576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127f890615683565b60405180910390fd5b5b5b5b5b5b5050565b816002600082815260200190815260200160002060070160009054906101000a900460ff16612865576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b61287281613192565b600060088111156128865761288561442a565b5b6002600086815260200190815260200160002060010160009054906101000a900460ff1660088111156128bc576128bb61442a565b5b148061290f5750600160088111156128d7576128d661442a565b5b6002600086815260200190815260200160002060010160009054906101000a900460ff16600881111561290d5761290c61442a565b5b145b806129615750600260088111156129295761292861442a565b5b6002600086815260200190815260200160002060010160009054906101000a900460ff16600881111561295f5761295e61442a565b5b145b6129a0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612997906156ef565b60405180910390fd5b826002600086815260200190815260200160002060020160008282546129c6919061570f565b9250508190555050505050565b60006129dd613cbf565b6000801b6129ea81613192565b6000612a356001600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206135f5565b905080612a428287611a42565b935093505050915091565b816002600082815260200190815260200160002060070160009054906101000a900460ff16612aa8576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b612ab581613192565b60006008811115612ac957612ac861442a565b5b6002600086815260200190815260200160002060010160009054906101000a900460ff166008811115612aff57612afe61442a565b5b14612b3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b369061578f565b60405180910390fd5b826002600086815260200190815260200160002060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016002600086815260200190815260200160002060010160006101000a81548160ff02191690836008811115612bce57612bcd61442a565b5b02179055504360026000868152602001908152602001600020600801819055506003849080600181540180825580915050600190039060005260206000200160009091909190915055612c5f600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085613579565b50505050565b60066020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b612cac82610dd5565b612cb581613192565b612cbf8383613386565b505050565b600c60009054906101000a900467ffffffffffffffff1681565b816002600082815260200190815260200160002060070160009054906101000a900460ff16612d39576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160026000858152602001908152602001600020600c0181905550505050565b60008060001b612d6881613192565b612db16004600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206135f5565b915050919050565b816002600082815260200190815260200160002060070160009054906101000a900460ff16612e14576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b612e2181613192565b826008811115612e3457612e3361442a565b5b6002600086815260200190815260200160002060010160009054906101000a900460ff166008811115612e6a57612e6961442a565b5b1461302e57826002600086815260200190815260200160002060010160006101000a81548160ff02191690836008811115612ea857612ea761442a565b5b021790555060046008811115612ec157612ec061442a565b5b836008811115612ed457612ed361442a565b5b03612f5857612f57600660006002600088815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085613579565b5b60056008811115612f6c57612f6b61442a565b5b836008811115612f7f57612f7e61442a565b5b1480612faf575060066008811115612f9a57612f9961442a565b5b836008811115612fad57612fac61442a565b5b145b80612fde575060076008811115612fc957612fc861442a565b5b836008811115612fdc57612fdb61442a565b5b145b8061300c5750600880811115612ff757612ff661442a565b5b83600881111561300a5761300961442a565b5b145b1561302d574260026000868152602001908152602001600020600a01819055505b5b50505050565b6009602052816000526040600020818154811061305057600080fd5b90600052602060002001600091509150505481565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008060001b61310181613192565b600061310b613d16565b60005b61311786613726565b81101561318357613128868261375a565b92506131338361213c565b91508660088111156131485761314761442a565b5b8260200151600881111561315f5761315e61442a565b5b03613170576001945050505061318b565b808061317b90614bf2565b91505061310e565b506000935050505b5092915050565b6131a38161319e61337e565b613804565b50565b6000826002600082815260200190815260200160002060070160009054906101000a900460ff16613203576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8267ffffffffffffffff1660026000868152602001908152602001600020600801544361323091906157af565b101591505092915050565b5b61324581613889565b61329a576000613254826138be565b9050826002600083815260200190815260200160002060010160006101000a81548160ff0219169083600881111561328f5761328e61442a565b5b02179055505061323c565b5050565b6132a8828261208f565b61337a57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061331f61337e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b613390828261208f565b1561346357600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061340861337e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b816002600082815260200190815260200160002060070160009054906101000a900460ff166134c2576040517f7d02dddc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b6134cf81613192565b8273ffffffffffffffffffffffffffffffffffffffff166002600086815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161356a9061582f565b60405180910390fd5b50505050565b60008260000160109054906101000a9004600f0b90508183600101600083600f0b600f0b815260200190815260200160002081905550600181018360000160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff160217905550505050565b60008060001b61360481613192565b61360d83613889565b1561364d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136449061589b565b60405180910390fd5b6000613657613d16565b5b613661856138be565b915061366c8261213c565b905061367785613889565b1580156136ac57508560088111156136925761369161442a565b5b816020015160088111156136a9576136a861442a565b5b14155b613658578560088111156136c3576136c261442a565b5b816020015160088111156136da576136d961442a565b5b1461371a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137119061589b565b60405180910390fd5b81935050505092915050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b039050919050565b60008061378d6137698461399a565b8560000160009054906101000a9004600f0b600f0b61378891906158c5565b613a07565b90508360000160109054906101000a9004600f0b600f0b81600f0b126137df576040517fb4120f1400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600101600082600f0b600f0b81526020019081526020016000205491505092915050565b61380e828261208f565b6138855761381b81613a56565b6138298360001c6020613a83565b60405160200161383a9291906159dd565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161387c9190615a50565b60405180910390fd5b5050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b13159050919050565b60006138c982613889565b15613900576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260000160009054906101000a9004600f0b905082600101600082600f0b600f0b815260200190815260200160002054915082600101600082600f0b600f0b815260200190815260200160002060009055600181018360000160006101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff16021790555050919050565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211156139ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139f690615ae4565b60405180910390fd5b819050919050565b60008190508181600f0b14613a51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a4890615b76565b60405180910390fd5b919050565b6060613a7c8273ffffffffffffffffffffffffffffffffffffffff16601460ff16613a83565b9050919050565b606060006002836002613a969190615b96565b613aa0919061570f565b67ffffffffffffffff811115613ab957613ab8614096565b5b6040519080825280601f01601f191660200182016040528015613aeb5781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110613b2357613b22614f41565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110613b8757613b86614f41565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002613bc79190615b96565b613bd1919061570f565b90505b6001811115613c71577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110613c1357613c12614f41565b5b1a60f81b828281518110613c2a57613c29614f41565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080613c6a90615bd8565b9050613bd4565b5060008414613cb5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613cac90615c4d565b60405180910390fd5b8091505092915050565b6040518060a00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b6040518061018001604052806000801916815260200160006008811115613d4057613d3f61442a565b5b815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001613d8e613cbf565b815260200160001515815260200160008152602001600081526020016000815260200160008019168152602001600081525090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b613e0c81613dd7565b8114613e1757600080fd5b50565b600081359050613e2981613e03565b92915050565b600060208284031215613e4557613e44613dcd565b5b6000613e5384828501613e1a565b91505092915050565b60008115159050919050565b613e7181613e5c565b82525050565b6000602082019050613e8c6000830184613e68565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613ebd82613e92565b9050919050565b613ecd81613eb2565b8114613ed857600080fd5b50565b600081359050613eea81613ec4565b92915050565b600060208284031215613f0657613f05613dcd565b5b6000613f1484828501613edb565b91505092915050565b6000819050919050565b613f3081613f1d565b82525050565b6000602082019050613f4b6000830184613f27565b92915050565b600081600f0b9050919050565b613f6781613f51565b82525050565b6000604082019050613f826000830185613f5e565b613f8f6020830184613f5e565b9392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b613fd581613fc2565b82525050565b6000613fe78383613fcc565b60208301905092915050565b6000602082019050919050565b600061400b82613f96565b6140158185613fa1565b935061402083613fb2565b8060005b838110156140515781516140388882613fdb565b975061404383613ff3565b925050600181019050614024565b5085935050505092915050565b600060208201905081810360008301526140788184614000565b905092915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6140ce82614085565b810181811067ffffffffffffffff821117156140ed576140ec614096565b5b80604052505050565b6000614100613dc3565b905061410c82826140c5565b919050565b600080fd5b600067ffffffffffffffff82169050919050565b61413381614116565b811461413e57600080fd5b50565b6000813590506141508161412a565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff82111561417b5761417a614096565b5b61418482614085565b9050602081019050919050565b82818337600083830152505050565b60006141b36141ae84614160565b6140f6565b9050828152602081018484840111156141cf576141ce61415b565b5b6141da848285614191565b509392505050565b600082601f8301126141f7576141f6614156565b5b81356142078482602086016141a0565b91505092915050565b600060a0828403121561422657614225614080565b5b61423060a06140f6565b9050600061424084828501614141565b600083015250602061425484828501614141565b602083015250604061426884828501614141565b604083015250606061427c84828501614141565b606083015250608082013567ffffffffffffffff8111156142a05761429f614111565b5b6142ac848285016141e2565b60808301525092915050565b6142c181613f1d565b81146142cc57600080fd5b50565b6000813590506142de816142b8565b92915050565b6142ed81613fc2565b81146142f857600080fd5b50565b60008135905061430a816142e4565b92915050565b6000806000806080858703121561432a57614329613dcd565b5b600061433887828801613edb565b945050602085013567ffffffffffffffff81111561435957614358613dd2565b5b61436587828801614210565b9350506040614376878288016142cf565b9250506060614387878288016142fb565b91505092959194509250565b61439c81613fc2565b82525050565b60006020820190506143b76000830184614393565b92915050565b6000602082840312156143d3576143d2613dcd565b5b60006143e1848285016142fb565b91505092915050565b6000806040838503121561440157614400613dcd565b5b600061440f858286016142fb565b925050602061442085828601613edb565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6009811061446a5761446961442a565b5b50565b600081905061447b82614459565b919050565b600061448b8261446d565b9050919050565b61449b81614480565b82525050565b6144aa81613eb2565b82525050565b6144b981614116565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156144f95780820151818401526020810190506144de565b60008484015250505050565b6000614510826144bf565b61451a81856144ca565b935061452a8185602086016144db565b61453381614085565b840191505092915050565b600060a08301600083015161455660008601826144b0565b50602083015161456960208601826144b0565b50604083015161457c60408601826144b0565b50606083015161458f60608601826144b0565b50608083015184820360808601526145a78282614505565b9150508091505092915050565b6000610180820190506145ca600083018f614393565b6145d7602083018e614492565b6145e4604083018d613f27565b6145f1606083018c6144a1565b6145fe608083018b6144a1565b81810360a0830152614610818a61453e565b905061461f60c0830189613e68565b61462c60e0830188613f27565b61463a610100830187613f27565b614648610120830186613f27565b614656610140830185614393565b614664610160830184613f27565b9d9c50505050505050505050505050565b6009811061468257600080fd5b50565b60008135905061469481614675565b92915050565b6000806000606084860312156146b3576146b2613dcd565b5b60006146c1868287016142fb565b93505060206146d286828701613edb565b92505060406146e386828701614685565b9150509250925092565b60006020820190508181036000830152614707818461453e565b905092915050565b60006020828403121561472557614724613dcd565b5b6000614733848285016142cf565b91505092915050565b61474581614480565b82525050565b61475481613f1d565b82525050565b61476381613eb2565b82525050565b600060a08301600083015161478160008601826144b0565b50602083015161479460208601826144b0565b5060408301516147a760408601826144b0565b5060608301516147ba60608601826144b0565b50608083015184820360808601526147d28282614505565b9150508091505092915050565b6147e881613e5c565b82525050565b6000610180830160008301516148076000860182613fcc565b50602083015161481a602086018261473c565b50604083015161482d604086018261474b565b506060830151614840606086018261475a565b506080830151614853608086018261475a565b5060a083015184820360a086015261486b8282614769565b91505060c083015161488060c08601826147df565b5060e083015161489360e086018261474b565b506101008301516148a861010086018261474b565b506101208301516148bd61012086018261474b565b506101408301516148d2610140860182613fcc565b506101608301516148e761016086018261474b565b508091505092915050565b6000602082019050818103600083015261490c81846147ee565b905092915050565b6000806040838503121561492b5761492a613dcd565b5b600061493985828601614685565b925050602061494a85828601614685565b9150509250929050565b6000806040838503121561496b5761496a613dcd565b5b6000614979858286016142fb565b925050602061498a858286016142cf565b9150509250929050565b60006040820190506149a96000830185614393565b81810360208301526149bb818461453e565b90509392505050565b6149cd81614116565b82525050565b60006020820190506149e860008301846149c4565b92915050565b60008060408385031215614a0557614a04613dcd565b5b6000614a13858286016142fb565b9250506020614a2485828601614685565b9150509250929050565b60008060408385031215614a4557614a44613dcd565b5b6000614a5385828601613edb565b9250506020614a64858286016142cf565b9150509250929050565b600082825260208201905092915050565b7f4a42310000000000000000000000000000000000000000000000000000000000600082015250565b6000614ab5600383614a6e565b9150614ac082614a7f565b602082019050919050565b60006020820190508181036000830152614ae481614aa8565b9050919050565b7f4a42320000000000000000000000000000000000000000000000000000000000600082015250565b6000614b21600383614a6e565b9150614b2c82614aeb565b602082019050919050565b60006020820190508181036000830152614b5081614b14565b9050919050565b7f4a42330000000000000000000000000000000000000000000000000000000000600082015250565b6000614b8d600383614a6e565b9150614b9882614b57565b602082019050919050565b60006020820190508181036000830152614bbc81614b80565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000614bfd82613f1d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614c2f57614c2e614bc3565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680614c8157607f821691505b602082108103614c9457614c93614c3a565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302614cfc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614cbf565b614d068683614cbf565b95508019841693508086168417925050509392505050565b6000819050919050565b6000614d43614d3e614d3984613f1d565b614d1e565b613f1d565b9050919050565b6000819050919050565b614d5d83614d28565b614d71614d6982614d4a565b848454614ccc565b825550505050565b600090565b614d86614d79565b614d91818484614d54565b505050565b5b81811015614db557614daa600082614d7e565b600181019050614d97565b5050565b601f821115614dfa57614dcb81614c9a565b614dd484614caf565b81016020851015614de3578190505b614df7614def85614caf565b830182614d96565b50505b505050565b600082821c905092915050565b6000614e1d60001984600802614dff565b1980831691505092915050565b6000614e368383614e0c565b9150826002028217905092915050565b614e4f826144bf565b67ffffffffffffffff811115614e6857614e67614096565b5b614e728254614c69565b614e7d828285614db9565b600060209050601f831160018114614eb05760008415614e9e578287015190505b614ea88582614e2a565b865550614f10565b601f198416614ebe86614c9a565b60005b82811015614ee657848901518255600182019150602085019450602081019050614ec1565b86831015614f035784890151614eff601f891682614e0c565b8355505b6001600288020188555050505b505050505050565b6000604082019050614f2d6000830185614393565b614f3a60208301846144a1565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4a42370000000000000000000000000000000000000000000000000000000000600082015250565b6000614fa6600383614a6e565b9150614fb182614f70565b602082019050919050565b60006020820190508181036000830152614fd581614f99565b9050919050565b6000614fe782614116565b915067ffffffffffffffff820361500157615000614bc3565b5b600182019050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000615068602f83614a6e565b91506150738261500c565b604082019050919050565b600060208201905081810360008301526150978161505b565b9050919050565b7f4a42313200000000000000000000000000000000000000000000000000000000600082015250565b60006150d4600483614a6e565b91506150df8261509e565b602082019050919050565b60006020820190508181036000830152615103816150c7565b9050919050565b7f4a42313300000000000000000000000000000000000000000000000000000000600082015250565b6000615140600483614a6e565b915061514b8261510a565b602082019050919050565b6000602082019050818103600083015261516f81615133565b9050919050565b7f4a42313400000000000000000000000000000000000000000000000000000000600082015250565b60006151ac600483614a6e565b91506151b782615176565b602082019050919050565b600060208201905081810360008301526151db8161519f565b9050919050565b7f4a42313500000000000000000000000000000000000000000000000000000000600082015250565b6000615218600483614a6e565b9150615223826151e2565b602082019050919050565b600060208201905081810360008301526152478161520b565b9050919050565b7f4a42323100000000000000000000000000000000000000000000000000000000600082015250565b6000615284600483614a6e565b915061528f8261524e565b602082019050919050565b600060208201905081810360008301526152b381615277565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000615316602e83614a6e565b9150615321826152ba565b604082019050919050565b6000602082019050818103600083015261534581615309565b9050919050565b6000819050919050565b600060ff82169050919050565b600061537e6153796153748461534c565b614d1e565b615356565b9050919050565b61538e81615363565b82525050565b60006020820190506153a96000830184615385565b92915050565b7f4a42380000000000000000000000000000000000000000000000000000000000600082015250565b60006153e5600383614a6e565b91506153f0826153af565b602082019050919050565b60006020820190508181036000830152615414816153d8565b9050919050565b7f4a42390000000000000000000000000000000000000000000000000000000000600082015250565b6000615451600383614a6e565b915061545c8261541b565b602082019050919050565b6000602082019050818103600083015261548081615444565b9050919050565b7f4a42313600000000000000000000000000000000000000000000000000000000600082015250565b60006154bd600483614a6e565b91506154c882615487565b602082019050919050565b600060208201905081810360008301526154ec816154b0565b9050919050565b7f4a42313700000000000000000000000000000000000000000000000000000000600082015250565b6000615529600483614a6e565b9150615534826154f3565b602082019050919050565b600060208201905081810360008301526155588161551c565b9050919050565b7f4a42313800000000000000000000000000000000000000000000000000000000600082015250565b6000615595600483614a6e565b91506155a08261555f565b602082019050919050565b600060208201905081810360008301526155c481615588565b9050919050565b7f4a42313900000000000000000000000000000000000000000000000000000000600082015250565b6000615601600483614a6e565b915061560c826155cb565b602082019050919050565b60006020820190508181036000830152615630816155f4565b9050919050565b7f4a42323000000000000000000000000000000000000000000000000000000000600082015250565b600061566d600483614a6e565b915061567882615637565b602082019050919050565b6000602082019050818103600083015261569c81615660565b9050919050565b7f4a42323200000000000000000000000000000000000000000000000000000000600082015250565b60006156d9600483614a6e565b91506156e4826156a3565b602082019050919050565b60006020820190508181036000830152615708816156cc565b9050919050565b600061571a82613f1d565b915061572583613f1d565b925082820190508082111561573d5761573c614bc3565b5b92915050565b7f4a42360000000000000000000000000000000000000000000000000000000000600082015250565b6000615779600383614a6e565b915061578482615743565b602082019050919050565b600060208201905081810360008301526157a88161576c565b9050919050565b60006157ba82613f1d565b91506157c583613f1d565b92508282039050818111156157dd576157dc614bc3565b5b92915050565b7f4a42313100000000000000000000000000000000000000000000000000000000600082015250565b6000615819600483614a6e565b9150615824826157e3565b602082019050919050565b600060208201905081810360008301526158488161580c565b9050919050565b7f4a42313000000000000000000000000000000000000000000000000000000000600082015250565b6000615885600483614a6e565b91506158908261584f565b602082019050919050565b600060208201905081810360008301526158b481615878565b9050919050565b6000819050919050565b60006158d0826158bb565b91506158db836158bb565b92508282019050828112156000831216838212600084121516171561590357615902614bc3565b5b92915050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b600061594a601783615909565b915061595582615914565b601782019050919050565b600061596b826144bf565b6159758185615909565b93506159858185602086016144db565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b60006159c7601183615909565b91506159d282615991565b601182019050919050565b60006159e88261593d565b91506159f48285615960565b91506159ff826159ba565b9150615a0b8284615960565b91508190509392505050565b6000615a22826144bf565b615a2c8185614a6e565b9350615a3c8185602086016144db565b615a4581614085565b840191505092915050565b60006020820190508181036000830152615a6a8184615a17565b905092915050565b7f53616665436173743a2076616c756520646f65736e27742066697420696e206160008201527f6e20696e74323536000000000000000000000000000000000000000000000000602082015250565b6000615ace602883614a6e565b9150615ad982615a72565b604082019050919050565b60006020820190508181036000830152615afd81615ac1565b9050919050565b7f53616665436173743a2076616c756520646f65736e27742066697420696e203160008201527f3238206269747300000000000000000000000000000000000000000000000000602082015250565b6000615b60602783614a6e565b9150615b6b82615b04565b604082019050919050565b60006020820190508181036000830152615b8f81615b53565b9050919050565b6000615ba182613f1d565b9150615bac83613f1d565b9250828202615bba81613f1d565b91508282048414831517615bd157615bd0614bc3565b5b5092915050565b6000615be382613f1d565b915060008203615bf657615bf5614bc3565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000615c37602083614a6e565b9150615c4282615c01565b602082019050919050565b60006020820190508181036000830152615c6681615c2a565b905091905056fea26469706673582212200c3989c5aaab2555f234ff0311d2c2fad817ed3014a8c083ecefc944cf3983ae64736f6c63430008110033",
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

// CheckNewJobStatus is a free data retrieval call binding the contract method 0xb42f5513.
//
// Solidity: function checkNewJobStatus(uint8 _previousStatus, uint8 _jobStatus) pure returns()
func (_JobManager *JobManagerCaller) CheckNewJobStatus(opts *bind.CallOpts, _previousStatus uint8, _jobStatus uint8) error {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "checkNewJobStatus", _previousStatus, _jobStatus)

	if err != nil {
		return err
	}

	return err

}

// CheckNewJobStatus is a free data retrieval call binding the contract method 0xb42f5513.
//
// Solidity: function checkNewJobStatus(uint8 _previousStatus, uint8 _jobStatus) pure returns()
func (_JobManager *JobManagerSession) CheckNewJobStatus(_previousStatus uint8, _jobStatus uint8) error {
	return _JobManager.Contract.CheckNewJobStatus(&_JobManager.CallOpts, _previousStatus, _jobStatus)
}

// CheckNewJobStatus is a free data retrieval call binding the contract method 0xb42f5513.
//
// Solidity: function checkNewJobStatus(uint8 _previousStatus, uint8 _jobStatus) pure returns()
func (_JobManager *JobManagerCallerSession) CheckNewJobStatus(_previousStatus uint8, _jobStatus uint8) error {
	return _JobManager.Contract.CheckNewJobStatus(&_JobManager.CallOpts, _previousStatus, _jobStatus)
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
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,string),bool,uint256,uint256,uint256,bytes32,uint256))
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
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,string),bool,uint256,uint256,uint256,bytes32,uint256))
func (_JobManager *JobManagerSession) GetJobFromId(_jobId [32]byte) (Job, error) {
	return _JobManager.Contract.GetJobFromId(&_JobManager.CallOpts, _jobId)
}

// GetJobFromId is a free data retrieval call binding the contract method 0xb3130fba.
//
// Solidity: function getJobFromId(bytes32 _jobId) view returns((bytes32,uint8,uint256,address,address,(uint64,uint64,uint64,uint64,string),bool,uint256,uint256,uint256,bytes32,uint256))
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

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_JobManager *JobManagerCaller) HasCancellingJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "hasCancellingJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_JobManager *JobManagerSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _JobManager.Contract.HasCancellingJob(&_JobManager.CallOpts, _providerAddr)
}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_JobManager *JobManagerCallerSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _JobManager.Contract.HasCancellingJob(&_JobManager.CallOpts, _providerAddr)
}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_JobManager *JobManagerCaller) HasNextJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "hasNextJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_JobManager *JobManagerSession) HasNextJob(_providerAddr common.Address) (bool, error) {
	return _JobManager.Contract.HasNextJob(&_JobManager.CallOpts, _providerAddr)
}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_JobManager *JobManagerCallerSession) HasNextJob(_providerAddr common.Address) (bool, error) {
	return _JobManager.Contract.HasNextJob(&_JobManager.CallOpts, _providerAddr)
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
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange, uint256 start, uint256 end, bytes32 jobName, uint256 cost)
func (_JobManager *JobManagerCaller) Jobs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Start                  *big.Int
	End                    *big.Int
	JobName                [32]byte
	Cost                   *big.Int
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
		Start                  *big.Int
		End                    *big.Int
		JobName                [32]byte
		Cost                   *big.Int
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
	outstruct.Start = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.JobName = *abi.ConvertType(out[10], new([32]byte)).(*[32]byte)
	outstruct.Cost = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange, uint256 start, uint256 end, bytes32 jobName, uint256 cost)
func (_JobManager *JobManagerSession) Jobs(arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Start                  *big.Int
	End                    *big.Int
	JobName                [32]byte
	Cost                   *big.Int
}, error) {
	return _JobManager.Contract.Jobs(&_JobManager.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, uint256 amountLocked, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string) definition, bool valid, uint256 blockNumberStateChange, uint256 start, uint256 end, bytes32 jobName, uint256 cost)
func (_JobManager *JobManagerCallerSession) Jobs(arg0 [32]byte) (struct {
	JobId                  [32]byte
	Status                 uint8
	AmountLocked           *big.Int
	CustomerAddr           common.Address
	ProviderAddr           common.Address
	Definition             JobDefinition
	Valid                  bool
	BlockNumberStateChange *big.Int
	Start                  *big.Int
	End                    *big.Int
	JobName                [32]byte
	Cost                   *big.Int
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

// ProviderCancelledJobsQueues is a free data retrieval call binding the contract method 0xd43a8027.
//
// Solidity: function providerCancelledJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCaller) ProviderCancelledJobsQueues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _JobManager.contract.Call(opts, &out, "providerCancelledJobsQueues", arg0)

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

// ProviderCancelledJobsQueues is a free data retrieval call binding the contract method 0xd43a8027.
//
// Solidity: function providerCancelledJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerSession) ProviderCancelledJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderCancelledJobsQueues(&_JobManager.CallOpts, arg0)
}

// ProviderCancelledJobsQueues is a free data retrieval call binding the contract method 0xd43a8027.
//
// Solidity: function providerCancelledJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_JobManager *JobManagerCallerSession) ProviderCancelledJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _JobManager.Contract.ProviderCancelledJobsQueues(&_JobManager.CallOpts, arg0)
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

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0xde2418b3.
//
// Solidity: function claimNextCancellingJob(address _providerAddr) returns(bytes32 _jobId)
func (_JobManager *JobManagerTransactor) ClaimNextCancellingJob(opts *bind.TransactOpts, _providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "claimNextCancellingJob", _providerAddr)
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0xde2418b3.
//
// Solidity: function claimNextCancellingJob(address _providerAddr) returns(bytes32 _jobId)
func (_JobManager *JobManagerSession) ClaimNextCancellingJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimNextCancellingJob(&_JobManager.TransactOpts, _providerAddr)
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0xde2418b3.
//
// Solidity: function claimNextCancellingJob(address _providerAddr) returns(bytes32 _jobId)
func (_JobManager *JobManagerTransactorSession) ClaimNextCancellingJob(_providerAddr common.Address) (*types.Transaction, error) {
	return _JobManager.Contract.ClaimNextCancellingJob(&_JobManager.TransactOpts, _providerAddr)
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

// SetJobCost is a paid mutator transaction binding the contract method 0xd8e3aac7.
//
// Solidity: function setJobCost(bytes32 _jobId, uint256 _jobCost) returns()
func (_JobManager *JobManagerTransactor) SetJobCost(opts *bind.TransactOpts, _jobId [32]byte, _jobCost *big.Int) (*types.Transaction, error) {
	return _JobManager.contract.Transact(opts, "setJobCost", _jobId, _jobCost)
}

// SetJobCost is a paid mutator transaction binding the contract method 0xd8e3aac7.
//
// Solidity: function setJobCost(bytes32 _jobId, uint256 _jobCost) returns()
func (_JobManager *JobManagerSession) SetJobCost(_jobId [32]byte, _jobCost *big.Int) (*types.Transaction, error) {
	return _JobManager.Contract.SetJobCost(&_JobManager.TransactOpts, _jobId, _jobCost)
}

// SetJobCost is a paid mutator transaction binding the contract method 0xd8e3aac7.
//
// Solidity: function setJobCost(bytes32 _jobId, uint256 _jobCost) returns()
func (_JobManager *JobManagerTransactorSession) SetJobCost(_jobId [32]byte, _jobCost *big.Int) (*types.Transaction, error) {
	return _JobManager.Contract.SetJobCost(&_JobManager.TransactOpts, _jobId, _jobCost)
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

// MetaSchedulerMetaData contains all meta data concerning the MetaScheduler contract.
var MetaSchedulerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyHasJoinedError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"ClaimNextCancellingJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimNextJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METASCHEDULER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_AMOUNT_LOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVIDER_REDEMPTION_TAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVIDER_REGISTRATION_TAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextCancellingJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amountLocked\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"provider\",\"type\":\"tuple\"}],\"name\":\"convertCreditToDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"provider\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"durationMinute\",\"type\":\"uint64\"}],\"name\":\"convertDurationToCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobAmountLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"getJobStatus\",\"outputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasCancellingJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasNextJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_credit\",\"type\":\"address\"},{\"internalType\":\"contractJobManager\",\"name\":\"_jobManager\",\"type\":\"address\"},{\"internalType\":\"contractProviderManager\",\"name\":\"_providerManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"providerApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerRedemption\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"_providerDefinition\",\"type\":\"tuple\"}],\"name\":\"providerRegister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_jobDurationMinute\",\"type\":\"uint64\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"_providerDefinition\",\"type\":\"tuple\"}],\"name\":\"providerUpdateHardware\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amountLocked\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpMyJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateJobsStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_amount\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50600060018054906101000a900460ff1615905080801562000044575060018060009054906101000a900460ff1660ff16105b8062000081575062000061306200016360201b62002f4d1760201c565b15801562000080575060018060009054906101000a900460ff1660ff16145b5b620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba906200020d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015620001005760018060016101000a81548160ff0219169083151502179055505b80156200015c5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405162000153919062000289565b60405180910390a15b50620002a6565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620001f5602e8362000186565b9150620002028262000197565b604082019050919050565b600060208201905081810360008301526200022881620001e6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620002716200026b62000265846200022f565b62000246565b62000239565b9050919050565b620002838162000250565b82525050565b6000602082019050620002a0600083018462000278565b92915050565b615b2380620002b66000396000f3fe6080604052600436106101f95760003560e01c80636a6fad071161010d578063b6e5133f116100a0578063d1cee5461161006f578063d1cee54614610732578063d547741f1461075b578063e052888c14610784578063e8a81986146107af578063fbc3611a146107d8576101f9565b8063b6e5133f14610673578063c0c53b8b146106b0578063c4b6e0cc146106d9578063cffe070c14610716576101f9565b806391d14854116100dc57806391d14854146105b5578063a06d083c146105f2578063a217fddf1461061d578063b35b67ae14610648576101f9565b80636a6fad07146104fd5780636df90c9d14610526578063750f0acc1461056357806378952c881461058c576101f9565b80632a242a761161019057806336568abe1161015f57806336568abe1461045457806348841b9c1461047d5780635d3a7180146104a65780635e1b2d65146104bd5780635fae1450146104d4576101f9565b80632a242a76146103ae5780632f2ff15d146103c55780632fecc4f6146103ee57806333474c7114610417576101f9565b80632081f4c8116101cc5780632081f4c8146102cc57806320a5f91914610309578063214a353c14610346578063248a9ca314610371576101f9565b806301ffc9a7146101fe5780630797094e1461023b5780630948ee63146102785780631f92a63f146102a3575b600080fd5b34801561020a57600080fd5b50610225600480360381019061022091906138ae565b6107e2565b60405161023291906138f6565b60405180910390f35b34801561024757600080fd5b50610262600480360381019061025d919061396f565b61085c565b60405161026f91906138f6565b60405180910390f35b34801561028457600080fd5b5061028d610901565b60405161029a91906139b5565b60405180910390f35b3480156102af57600080fd5b506102ca60048036038101906102c59190613a06565b61090d565b005b3480156102d857600080fd5b506102f360048036038101906102ee9190613a06565b610a7a565b6040516103009190613aaa565b60405180910390f35b34801561031557600080fd5b50610330600480360381019061032b919061396f565b610b28565b60405161033d91906138f6565b60405180910390f35b34801561035257600080fd5b5061035b610bcd565b60405161036891906139b5565b60405180910390f35b34801561037d57600080fd5b5061039860048036038101906103939190613a06565b610bda565b6040516103a59190613ad4565b60405180910390f35b3480156103ba57600080fd5b506103c3610bf9565b005b3480156103d157600080fd5b506103ec60048036038101906103e79190613aef565b610ca8565b005b3480156103fa57600080fd5b5061041560048036038101906104109190613b5b565b610cc9565b005b34801561042357600080fd5b5061043e60048036038101906104399190613e87565b610e86565b60405161044b91906139b5565b60405180910390f35b34801561046057600080fd5b5061047b60048036038101906104769190613aef565b610f74565b005b34801561048957600080fd5b506104a4600480360381019061049f9190613f1d565b610ff7565b005b3480156104b257600080fd5b506104bb611345565b005b3480156104c957600080fd5b506104d2611715565b005b3480156104e057600080fd5b506104fb60048036038101906104f69190613a06565b61197a565b005b34801561050957600080fd5b50610524600480360381019061051f9190613f70565b611ef9565b005b34801561053257600080fd5b5061054d60048036038101906105489190613fb1565b611fb4565b60405161055a9190614030565b60405180910390f35b34801561056f57600080fd5b5061058a6004803603810190610585919061404b565b6120e0565b005b34801561059857600080fd5b506105b360048036038101906105ae919061396f565b6122bf565b005b3480156105c157600080fd5b506105dc60048036038101906105d79190613aef565b61237a565b6040516105e991906138f6565b60405180910390f35b3480156105fe57600080fd5b506106076123e4565b60405161061491906140d7565b60405180910390f35b34801561062957600080fd5b5061063261240a565b60405161063f9190613ad4565b60405180910390f35b34801561065457600080fd5b5061065d612411565b60405161066a91906139b5565b60405180910390f35b34801561067f57600080fd5b5061069a60048036038101906106959190613a06565b61241d565b6040516106a791906139b5565b60405180910390f35b3480156106bc57600080fd5b506106d760048036038101906106d291906141ac565b6124cb565b005b3480156106e557600080fd5b5061070060048036038101906106fb91906141ff565b6127b2565b60405161070d9190613ad4565b60405180910390f35b610730600480360381019061072b9190613f70565b612994565b005b34801561073e57600080fd5b5061075960048036038101906107549190613aef565b612a9a565b005b34801561076757600080fd5b50610782600480360381019061077d9190613aef565b612c33565b005b34801561079057600080fd5b50610799612c54565b6040516107a69190613ad4565b60405180910390f35b3480156107bb57600080fd5b506107d660048036038101906107d19190613b5b565b612c78565b005b6107e0612d98565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610855575061085482612f70565b5b9050919050565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630797094e836040518263ffffffff1660e01b81526004016108b9919061427d565b602060405180830381865afa1580156108d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fa91906142c4565b9050919050565b670de0b6b3a764000081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401610968919061427d565b602060405180830381865afa158015610985573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a991906142c4565b6109e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109df9061434e565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634609ca5082336040518363ffffffff1660e01b8152600401610a4592919061436e565b600060405180830381600087803b158015610a5f57600080fd5b505af1158015610a73573d6000803e3d6000fd5b5050505050565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b8152600401610ad79190613ad4565b600060405180830381865afa158015610af4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610b1d9190614680565b602001519050919050565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166320a5f919836040518263ffffffff1660e01b8152600401610b85919061427d565b602060405180830381865afa158015610ba2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc691906142c4565b9050919050565b68056bc75e2d6310000081565b6000806000838152602001908152602001600020600101549050919050565b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e610c2381612fda565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a242a766040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610c8d57600080fd5b505af1158015610ca1573d6000803e3d6000fd5b5050505050565b610cb182610bda565b610cba81612fda565b610cc48383612fee565b505050565b60008111610d0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0390614715565b60405180910390fd5b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401610d6d93929190614735565b6020604051808303816000875af1158015610d8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db091906142c4565b905080610df2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de9906147b8565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b67644b984846040518363ffffffff1660e01b8152600401610e4f9291906147d8565b600060405180830381600087803b158015610e6957600080fd5b505af1158015610e7d573d6000803e3d6000fd5b50505050505050565b600064e8d4a51000846040015167ffffffffffffffff16846080015167ffffffffffffffff16610eb69190614830565b856040015167ffffffffffffffff16866020015167ffffffffffffffff168660c0015167ffffffffffffffff16610eed9190614830565b610ef79190614830565b866000015167ffffffffffffffff16866040015167ffffffffffffffff16610f1f9190614830565b610f299190614872565b610f339190614872565b856060015167ffffffffffffffff168467ffffffffffffffff16610f579190614830565b610f619190614830565b610f6b9190614830565b90509392505050565b610f7c6130ce565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610fe9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe090614918565b60405180910390fd5b610ff382826130d6565b5050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401611052919061427d565b602060405180830381865afa15801561106f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109391906142c4565b6110d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c99061434e565b60405180910390fd5b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba856040518263ffffffff1660e01b815260040161112f9190613ad4565b600060405180830381865afa15801561114c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906111759190614680565b90503373ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff16146111e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e0906149aa565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637ff67e298533866040518463ffffffff1660e01b8152600401611248939291906149ca565b600060405180830381600087803b15801561126257600080fd5b505af1158015611276573d6000803e3d6000fd5b505050506006600881111561128e5761128d613a33565b5b8360088111156112a1576112a0613a33565b5b14806112d15750600560088111156112bc576112bb613a33565b5b8360088111156112cf576112ce613a33565b5b145b806112ff57506008808111156112ea576112e9613a33565b5b8360088111156112fd576112fc613a33565b5b145b8061132e57506007600881111561131957611318613a33565b5b83600881111561132c5761132b613a33565b5b145b1561133f5761133e8133846131b7565b5b50505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b81526004016113a0919061427d565b602060405180830381865afa1580156113bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e191906142c4565b611420576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114179061434e565b60405180910390fd5b600061142a6137eb565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c58467b0336040518263ffffffff1660e01b8152600401611485919061427d565b6000604051808303816000875af11580156114a4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906114cd9190614a01565b80925081935050506000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba846040518263ffffffff1660e01b81526004016115329190613ad4565b600060405180830381865afa15801561154f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906115789190614680565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f6edb5f82608001516040518263ffffffff1660e01b81526004016115d9919061427d565b600060405180830381600087803b1580156115f357600080fd5b505af1158015611607573d6000803e3d6000fd5b505050506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663624bc8e3336040518263ffffffff1660e01b8152600401611668919061427d565b6101a060405180830381865afa158015611686573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116aa9190614bf0565b905060006116c58360a0015184604001518460200151611fb4565b90507f6103c1800236e5e6e1833bde789240a0577a25b448a539ae39471a9e75d0422583606001513387848760a00151604051611706959493929190614cf8565b60405180910390a15050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12336040518263ffffffff1660e01b8152600401611770919061427d565b602060405180830381865afa15801561178d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117b191906142c4565b6117f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117e79061434e565b60405180910390fd5b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663de2418b3336040518263ffffffff1660e01b815260040161184d919061427d565b6020604051808303816000875af115801561186c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118909190614d52565b90506000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b81526004016118ef9190613ad4565b600060405180830381865afa15801561190c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906119359190614680565b90507f290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d8160600151338460405161196e93929190614d7f565b60405180910390a15050565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b81526004016119d79190613ad4565b600060405180830381865afa1580156119f4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611a1d9190614680565b90508060c00151611a63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a5a90614e02565b60405180910390fd5b806060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611ad5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611acc90614e94565b60405180910390fd5b60056008811115611ae957611ae8613a33565b5b81602001516008811115611b0057611aff613a33565b5b14158015611b37575060046008811115611b1d57611b1c613a33565b5b81602001516008811115611b3457611b33613a33565b5b14155b8015611b6c575060066008811115611b5257611b51613a33565b5b81602001516008811115611b6957611b68613a33565b5b14155b8015611ba05750600880811115611b8657611b85613a33565b5b81602001516008811115611b9d57611b9c613a33565b5b14155b8015611bd5575060076008811115611bbb57611bba613a33565b5b81602001516008811115611bd257611bd1613a33565b5b14155b611c14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c0b90614f4c565b60405180910390fd5b60006008811115611c2857611c27613a33565b5b81602001516008811115611c3f57611c3e613a33565b5b1480611c73575060016008811115611c5a57611c59613a33565b5b81602001516008811115611c7157611c70613a33565b5b145b15611dfa57600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e19086768360056040518363ffffffff1660e01b8152600401611cd6929190614f6c565b600060405180830381600087803b158015611cf057600080fd5b505af1158015611d04573d6000803e3d6000fd5b505050506000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb836060015184604001516040518363ffffffff1660e01b8152600401611d6f929190614f95565b6020604051808303816000875af1158015611d8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611db291906142c4565b905080611df4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611deb9061500a565b60405180910390fd5b50611ef5565b60026008811115611e0e57611e0d613a33565b5b81602001516008811115611e2557611e24613a33565b5b1480611e59575060036008811115611e4057611e3f613a33565b5b81602001516008811115611e5757611e56613a33565b5b145b15611ef357600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e19086768360046040518363ffffffff1660e01b8152600401611ebc929190614f6c565b600060405180830381600087803b158015611ed657600080fd5b505af1158015611eea573d6000803e3d6000fd5b50505050611ef4565b5b5b5050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635e9da3ed838360000151846020015185606001518660a00151876040015188608001518960c001516040518963ffffffff1660e01b8152600401611f7e98979695949392919061502a565b600060405180830381600087803b158015611f9857600080fd5b505af1158015611fac573d6000803e3d6000fd5b505050505050565b60008064e8d4a51000856040015167ffffffffffffffff16846080015167ffffffffffffffff16611fe59190614830565b866040015167ffffffffffffffff16876020015167ffffffffffffffff168660c0015167ffffffffffffffff1661201c9190614830565b6120269190614830565b876000015167ffffffffffffffff16866040015167ffffffffffffffff1661204e9190614830565b6120589190614872565b6120629190614872565b866060015167ffffffffffffffff1661207b9190614830565b6120859190614830565b9050600081036120ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120c1906150f4565b60405180910390fd5b80846120d69190615143565b9150509392505050565b6000801b6120ed81612fda565b8167ffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401612153919061427d565b602060405180830381865afa158015612170573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121949190615174565b116121d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121cb90615213565b60405180910390fd5b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3033866040518463ffffffff1660e01b815260040161223593929190615264565b6020604051808303816000875af1158015612254573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061227891906142c4565b9050806122ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122b1906147b8565b60405180910390fd5b505050565b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e6122e981612fda565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663daea85c5836040518263ffffffff1660e01b8152600401612344919061427d565b600060405180830381600087803b15801561235e57600080fd5b505af1158015612372573d6000803e3d6000fd5b505050505050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000801b81565b670de0b6b3a764000081565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b815260040161247a9190613ad4565b600060405180830381865afa158015612497573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906124c09190614680565b604001519050919050565b600060018054906101000a900460ff161590508080156124fc575060018060009054906101000a900460ff1660ff16105b8061252a575061250b30612f4d565b158015612529575060018060009054906101000a900460ff1660ff16145b5b612569576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125609061530d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff16021790555080156125a55760018060016101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603612614576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161260b9061539f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161267a90615431565b60405180910390fd5b6126906000801b33612fee565b83600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156127ac5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516127a39190615499565b60405180910390a15b50505050565b600068056bc75e2d63100000831015612800576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127f790615526565b60405180910390fd5b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b815260040161286193929190614735565b6020604051808303816000875af1158015612880573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128a491906142c4565b9050806128e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128dd906147b8565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631df4e610338787876040518563ffffffff1660e01b81526004016129479493929190615546565b6020604051808303816000875af1158015612966573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061298a9190614d52565b9150509392505050565b670de0b6b3a76400003410156129df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129d690615604565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b7c8116e838360000151846020015185606001518660a00151876040015188608001518960c001516040518963ffffffff1660e01b8152600401612a6498979695949392919061502a565b600060405180830381600087803b158015612a7e57600080fd5b505af1158015612a92573d6000803e3d6000fd5b505050505050565b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e612ac481612fda565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12836040518263ffffffff1660e01b8152600401612b1f919061427d565b602060405180830381865afa158015612b3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b6091906142c4565b612b9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b9690615696565b60405180910390fd5b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d1cee54684846040518363ffffffff1660e01b8152600401612bfc92919061436e565b600060405180830381600087803b158015612c1657600080fd5b505af1158015612c2a573d6000803e3d6000fd5b50505050505050565b612c3c82610bda565b612c4581612fda565b612c4f83836130d6565b505050565b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e81565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b3130fba836040518263ffffffff1660e01b8152600401612cd39190613ad4565b600060405180830381865afa158015612cf0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190612d199190614680565b6060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612d8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d8190615728565b60405180910390fd5b612d948282610cc9565b5050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b7bb9145336040518263ffffffff1660e01b8152600401612df3919061427d565b602060405180830381865afa158015612e10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e3491906142c4565b612e73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e6a90615794565b60405180910390fd5b670de0b6b3a7640000341015612ebe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612eb590615794565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7692d09336040518263ffffffff1660e01b8152600401612f19919061427d565b600060405180830381600087803b158015612f3357600080fd5b505af1158015612f47573d6000803e3d6000fd5b50505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b612feb81612fe66130ce565b6134fd565b50565b612ff8828261237a565b6130ca57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061306f6130ce565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6130e0828261237a565b156131b357600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506131586130ce565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663624bc8e3846040518263ffffffff1660e01b8152600401613214919061427d565b6101a060405180830381865afa158015613232573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132569190614bf0565b9050600061326d8560a00151836020015185610e86565b9050808560400151101561328357846040015190505b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d8e3aac78660000151836040518363ffffffff1660e01b81526004016132e49291906147d8565b600060405180830381600087803b1580156132fe57600080fd5b505af1158015613312573d6000803e3d6000fd5b505050506000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86846040518363ffffffff1660e01b8152600401613375929190614f95565b6020604051808303816000875af1158015613394573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133b891906142c4565b9050806133fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133f190615800565b60405180910390fd5b600082876040015161340c9190615820565b9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8860600151836040518363ffffffff1660e01b815260040161346f929190614f95565b6020604051808303816000875af115801561348e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134b291906142c4565b9150816134f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134eb906158a0565b60405180910390fd5b50505050505050565b613507828261237a565b61357e5761351481613582565b6135228360001c60206135af565b604051602001613533929190615994565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135759190615a07565b60405180910390fd5b5050565b60606135a88273ffffffffffffffffffffffffffffffffffffffff16601460ff166135af565b9050919050565b6060600060028360026135c29190614830565b6135cc9190614872565b67ffffffffffffffff8111156135e5576135e4613bb1565b5b6040519080825280601f01601f1916602001820160405280156136175781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061364f5761364e615a29565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106136b3576136b2615a29565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026136f39190614830565b6136fd9190614872565b90505b600181111561379d577f3031323334353637383961626364656600000000000000000000000000000000600f86166010811061373f5761373e615a29565b5b1a60f81b82828151811061375657613755615a29565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061379690615a58565b9050613700565b50600084146137e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137d890615acd565b60405180910390fd5b8091505092915050565b6040518060a00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61388b81613856565b811461389657600080fd5b50565b6000813590506138a881613882565b92915050565b6000602082840312156138c4576138c361384c565b5b60006138d284828501613899565b91505092915050565b60008115159050919050565b6138f0816138db565b82525050565b600060208201905061390b60008301846138e7565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061393c82613911565b9050919050565b61394c81613931565b811461395757600080fd5b50565b60008135905061396981613943565b92915050565b6000602082840312156139855761398461384c565b5b60006139938482850161395a565b91505092915050565b6000819050919050565b6139af8161399c565b82525050565b60006020820190506139ca60008301846139a6565b92915050565b6000819050919050565b6139e3816139d0565b81146139ee57600080fd5b50565b600081359050613a00816139da565b92915050565b600060208284031215613a1c57613a1b61384c565b5b6000613a2a848285016139f1565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60098110613a7357613a72613a33565b5b50565b6000819050613a8482613a62565b919050565b6000613a9482613a76565b9050919050565b613aa481613a89565b82525050565b6000602082019050613abf6000830184613a9b565b92915050565b613ace816139d0565b82525050565b6000602082019050613ae96000830184613ac5565b92915050565b60008060408385031215613b0657613b0561384c565b5b6000613b14858286016139f1565b9250506020613b258582860161395a565b9150509250929050565b613b388161399c565b8114613b4357600080fd5b50565b600081359050613b5581613b2f565b92915050565b60008060408385031215613b7257613b7161384c565b5b6000613b80858286016139f1565b9250506020613b9185828601613b46565b9150509250929050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b613be982613ba0565b810181811067ffffffffffffffff82111715613c0857613c07613bb1565b5b80604052505050565b6000613c1b613842565b9050613c278282613be0565b919050565b600080fd5b600067ffffffffffffffff82169050919050565b613c4e81613c31565b8114613c5957600080fd5b50565b600081359050613c6b81613c45565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff821115613c9657613c95613bb1565b5b613c9f82613ba0565b9050602081019050919050565b82818337600083830152505050565b6000613cce613cc984613c7b565b613c11565b905082815260208101848484011115613cea57613ce9613c76565b5b613cf5848285613cac565b509392505050565b600082601f830112613d1257613d11613c71565b5b8135613d22848260208601613cbb565b91505092915050565b600060a08284031215613d4157613d40613b9b565b5b613d4b60a0613c11565b90506000613d5b84828501613c5c565b6000830152506020613d6f84828501613c5c565b6020830152506040613d8384828501613c5c565b6040830152506060613d9784828501613c5c565b606083015250608082013567ffffffffffffffff811115613dbb57613dba613c2c565b5b613dc784828501613cfd565b60808301525092915050565b600060e08284031215613de957613de8613b9b565b5b613df360e0613c11565b90506000613e0384828501613c5c565b6000830152506020613e1784828501613c5c565b6020830152506040613e2b84828501613c5c565b6040830152506060613e3f84828501613c5c565b6060830152506080613e5384828501613c5c565b60808301525060a0613e6784828501613c5c565b60a08301525060c0613e7b84828501613c5c565b60c08301525092915050565b60008060006101208486031215613ea157613ea061384c565b5b600084013567ffffffffffffffff811115613ebf57613ebe613851565b5b613ecb86828701613d2b565b9350506020613edc86828701613dd3565b925050610100613eee86828701613c5c565b9150509250925092565b60098110613f0557600080fd5b50565b600081359050613f1781613ef8565b92915050565b600080600060608486031215613f3657613f3561384c565b5b6000613f44868287016139f1565b9350506020613f5586828701613f08565b9250506040613f6686828701613c5c565b9150509250925092565b6000806101008385031215613f8857613f8761384c565b5b6000613f968582860161395a565b9250506020613fa785828601613dd3565b9150509250929050565b60008060006101208486031215613fcb57613fca61384c565b5b600084013567ffffffffffffffff811115613fe957613fe8613851565b5b613ff586828701613d2b565b935050602061400686828701613b46565b925050604061401786828701613dd3565b9150509250925092565b61402a81613c31565b82525050565b60006020820190506140456000830184614021565b92915050565b6000602082840312156140615761406061384c565b5b600061406f84828501613c5c565b91505092915050565b6000819050919050565b600061409d61409861409384613911565b614078565b613911565b9050919050565b60006140af82614082565b9050919050565b60006140c1826140a4565b9050919050565b6140d1816140b6565b82525050565b60006020820190506140ec60008301846140c8565b92915050565b60006140fd82613931565b9050919050565b61410d816140f2565b811461411857600080fd5b50565b60008135905061412a81614104565b92915050565b600061413b82613931565b9050919050565b61414b81614130565b811461415657600080fd5b50565b60008135905061416881614142565b92915050565b600061417982613931565b9050919050565b6141898161416e565b811461419457600080fd5b50565b6000813590506141a681614180565b92915050565b6000806000606084860312156141c5576141c461384c565b5b60006141d38682870161411b565b93505060206141e486828701614159565b92505060406141f586828701614197565b9150509250925092565b6000806000606084860312156142185761421761384c565b5b600084013567ffffffffffffffff81111561423657614235613851565b5b61424286828701613d2b565b935050602061425386828701613b46565b9250506040614264868287016139f1565b9150509250925092565b61427781613931565b82525050565b6000602082019050614292600083018461426e565b92915050565b6142a1816138db565b81146142ac57600080fd5b50565b6000815190506142be81614298565b92915050565b6000602082840312156142da576142d961384c565b5b60006142e8848285016142af565b91505092915050565b600082825260208201905092915050565b7f70726f766964657220686173206e6f74206a6f696e6564000000000000000000600082015250565b60006143386017836142f1565b915061434382614302565b602082019050919050565b600060208201905081810360008301526143678161432b565b9050919050565b60006040820190506143836000830185613ac5565b614390602083018461426e565b9392505050565b6000815190506143a6816139da565b92915050565b6000815190506143bb81613ef8565b92915050565b6000815190506143d081613b2f565b92915050565b6000815190506143e581613943565b92915050565b6000815190506143fa81613c45565b92915050565b60005b8381101561441e578082015181840152602081019050614403565b60008484015250505050565b600061443d61443884613c7b565b613c11565b90508281526020810184848401111561445957614458613c76565b5b614464848285614400565b509392505050565b600082601f83011261448157614480613c71565b5b815161449184826020860161442a565b91505092915050565b600060a082840312156144b0576144af613b9b565b5b6144ba60a0613c11565b905060006144ca848285016143eb565b60008301525060206144de848285016143eb565b60208301525060406144f2848285016143eb565b6040830152506060614506848285016143eb565b606083015250608082015167ffffffffffffffff81111561452a57614529613c2c565b5b6145368482850161446c565b60808301525092915050565b6000610180828403121561455957614558613b9b565b5b614564610180613c11565b9050600061457484828501614397565b6000830152506020614588848285016143ac565b602083015250604061459c848285016143c1565b60408301525060606145b0848285016143d6565b60608301525060806145c4848285016143d6565b60808301525060a082015167ffffffffffffffff8111156145e8576145e7613c2c565b5b6145f48482850161449a565b60a08301525060c0614608848285016142af565b60c08301525060e061461c848285016143c1565b60e083015250610100614631848285016143c1565b61010083015250610120614647848285016143c1565b6101208301525061014061465d84828501614397565b61014083015250610160614673848285016143c1565b6101608301525092915050565b6000602082840312156146965761469561384c565b5b600082015167ffffffffffffffff8111156146b4576146b3613851565b5b6146c084828501614542565b91505092915050565b7f616d6f756e74206d757374206e6f74206265206e756c6c000000000000000000600082015250565b60006146ff6017836142f1565b915061470a826146c9565b602082019050919050565b6000602082019050818103600083015261472e816146f2565b9050919050565b600060608201905061474a600083018661426e565b614757602083018561426e565b61476460408301846139a6565b949350505050565b7f7472616e7366657246726f6d206661696c656400000000000000000000000000600082015250565b60006147a26013836142f1565b91506147ad8261476c565b602082019050919050565b600060208201905081810360008301526147d181614795565b9050919050565b60006040820190506147ed6000830185613ac5565b6147fa60208301846139a6565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061483b8261399c565b91506148468361399c565b92508282026148548161399c565b9150828204841483151761486b5761486a614801565b5b5092915050565b600061487d8261399c565b91506148888361399c565b92508282019050808211156148a05761489f614801565b5b92915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000614902602f836142f1565b915061490d826148a6565b604082019050919050565b60006020820190508181036000830152614931816148f5565b9050919050565b7f4f6e6c7920746865206a6f622070726f76696465722063616e206368616e676560008201527f2069747320737461747573000000000000000000000000000000000000000000602082015250565b6000614994602b836142f1565b915061499f82614938565b604082019050919050565b600060208201905081810360008301526149c381614987565b9050919050565b60006060820190506149df6000830186613ac5565b6149ec602083018561426e565b6149f96040830184613a9b565b949350505050565b60008060408385031215614a1857614a1761384c565b5b6000614a2685828601614397565b925050602083015167ffffffffffffffff811115614a4757614a46613851565b5b614a538582860161449a565b9150509250929050565b600060e08284031215614a7357614a72613b9b565b5b614a7d60e0613c11565b90506000614a8d848285016143eb565b6000830152506020614aa1848285016143eb565b6020830152506040614ab5848285016143eb565b6040830152506060614ac9848285016143eb565b6060830152506080614add848285016143eb565b60808301525060a0614af1848285016143eb565b60a08301525060c0614b05848285016143eb565b60c08301525092915050565b60038110614b1e57600080fd5b50565b600081519050614b3081614b11565b92915050565b60006101a08284031215614b4d57614b4c613b9b565b5b614b5760e0613c11565b90506000614b67848285016143d6565b6000830152506020614b7b84828501614a5d565b602083015250610100614b9084828501614b21565b604083015250610120614ba5848285016142af565b606083015250610140614bba848285016143eb565b608083015250610160614bcf848285016143c1565b60a083015250610180614be4848285016143c1565b60c08301525092915050565b60006101a08284031215614c0757614c0661384c565b5b6000614c1584828501614b36565b91505092915050565b614c2781613c31565b82525050565b600081519050919050565b600082825260208201905092915050565b6000614c5482614c2d565b614c5e8185614c38565b9350614c6e818560208601614400565b614c7781613ba0565b840191505092915050565b600060a083016000830151614c9a6000860182614c1e565b506020830151614cad6020860182614c1e565b506040830151614cc06040860182614c1e565b506060830151614cd36060860182614c1e565b5060808301518482036080860152614ceb8282614c49565b9150508091505092915050565b600060a082019050614d0d600083018861426e565b614d1a602083018761426e565b614d276040830186613ac5565b614d346060830185614021565b8181036080830152614d468184614c82565b90509695505050505050565b600060208284031215614d6857614d6761384c565b5b6000614d7684828501614397565b91505092915050565b6000606082019050614d94600083018661426e565b614da1602083018561426e565b614dae6040830184613ac5565b949350505050565b7f496e76616c6964206a6f62000000000000000000000000000000000000000000600082015250565b6000614dec600b836142f1565b9150614df782614db6565b602082019050919050565b60006020820190508181036000830152614e1b81614ddf565b9050919050565b7f4f6e6c79206a6f62206f776e6572732063616e2063616e63656c20746865697260008201527f206a6f6273000000000000000000000000000000000000000000000000000000602082015250565b6000614e7e6025836142f1565b9150614e8982614e22565b604082019050919050565b60006020820190508181036000830152614ead81614e71565b9050919050565b7f4a6f6220737461747573206d757374206e6f742062652043414e43454c4c454460008201527f2c2043414e43454c4c494e472c2046494e49534845442c204f55545f4f465f4360208201527f524544495453206f72204641494c454400000000000000000000000000000000604082015250565b6000614f366050836142f1565b9150614f4182614eb4565b606082019050919050565b60006020820190508181036000830152614f6581614f29565b9050919050565b6000604082019050614f816000830185613ac5565b614f8e6020830184613a9b565b9392505050565b6000604082019050614faa600083018561426e565b614fb760208301846139a6565b9392505050565b7f7472616e73666572206661696c65640000000000000000000000000000000000600082015250565b6000614ff4600f836142f1565b9150614fff82614fbe565b602082019050919050565b6000602082019050818103600083015261502381614fe7565b9050919050565b600061010082019050615040600083018b61426e565b61504d602083018a614021565b61505a6040830189614021565b6150676060830188614021565b6150746080830187614021565b61508160a0830186614021565b61508e60c0830185614021565b61509b60e0830184614021565b9998505050505050505050565b7f43616e6e6f742064697669646520627920300000000000000000000000000000600082015250565b60006150de6012836142f1565b91506150e9826150a8565b602082019050919050565b6000602082019050818103600083015261510d816150d1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061514e8261399c565b91506151598361399c565b92508261516957615168615114565b5b828204905092915050565b60006020828403121561518a5761518961384c565b5b6000615198848285016143c1565b91505092915050565b7f4d6574615363686564756c65723a20776974686472617720616d6f756e74206860008201527f6967686572207468616e2062616c616e63652e00000000000000000000000000602082015250565b60006151fd6033836142f1565b9150615208826151a1565b604082019050919050565b6000602082019050818103600083015261522c816151f0565b9050919050565b600061524e61524961524484613c31565b614078565b61399c565b9050919050565b61525e81615233565b82525050565b6000606082019050615279600083018661426e565b615286602083018561426e565b6152936040830184615255565b949350505050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60006152f7602e836142f1565b91506153028261529b565b604082019050919050565b60006020820190508181036000830152615326816152ea565b9050919050565b7f4d6574615363686564756c65723a206372656469742061646472206973207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b60006153896022836142f1565b91506153948261532d565b604082019050919050565b600060208201905081810360008301526153b88161537c565b9050919050565b7f4d6574615363686564756c65723a2070726f76696465724d616e61676572206160008201527f646472206973207a65726f000000000000000000000000000000000000000000602082015250565b600061541b602b836142f1565b9150615426826153bf565b604082019050919050565b6000602082019050818103600083015261544a8161540e565b9050919050565b6000819050919050565b600060ff82169050919050565b600061548361547e61547984615451565b614078565b61545b565b9050919050565b61549381615468565b82525050565b60006020820190506154ae600083018461548a565b92915050565b7f5f616d6f756e744c6f636b6564206d757374206265206772656174657220746860008201527f616e204d494e494d554d5f414d4f554e545f4c4f434b00000000000000000000602082015250565b60006155106036836142f1565b915061551b826154b4565b604082019050919050565b6000602082019050818103600083015261553f81615503565b9050919050565b600060808201905061555b600083018761426e565b818103602083015261556d8186614c82565b905061557c60408301856139a6565b6155896060830184613ac5565b95945050505050565b7f4d696e696d756d20616d6f756e7420746f207265676973746572206e6f74207260008201527f6561636865640000000000000000000000000000000000000000000000000000602082015250565b60006155ee6026836142f1565b91506155f982615592565b604082019050919050565b6000602082019050818103600083015261561d816155e1565b9050919050565b7f50726f766964657273206e6f7420666f756e64206f72206e6f74206a6f696e6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b60006156806021836142f1565b915061568b82615624565b604082019050919050565b600060208201905081810360008301526156af81615673565b9050919050565b7f4f6e6c7920746865206a6f62206f776e65722063616e20746f7020757020746860008201527f6973206a6f620000000000000000000000000000000000000000000000000000602082015250565b60006157126026836142f1565b915061571d826156b6565b604082019050919050565b6000602082019050818103600083015261574181615705565b9050919050565b7f50726f7669646572206e6f7420666f756e64206f72206e6f74206b69636b6564600082015250565b600061577e6020836142f1565b915061578982615748565b602082019050919050565b600060208201905081810360008301526157ad81615771565b9050919050565b7f70726f7669646572207061796261636b207472616e73666572206661696c6564600082015250565b60006157ea6020836142f1565b91506157f5826157b4565b602082019050919050565b60006020820190508181036000830152615819816157dd565b9050919050565b600061582b8261399c565b91506158368361399c565b925082820390508181111561584e5761584d614801565b5b92915050565b7f637573746f6d6572207061796261636b207472616e73666572206661696c6564600082015250565b600061588a6020836142f1565b915061589582615854565b602082019050919050565b600060208201905081810360008301526158b98161587d565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b60006159016017836158c0565b915061590c826158cb565b601782019050919050565b600061592282614c2d565b61592c81856158c0565b935061593c818560208601614400565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b600061597e6011836158c0565b915061598982615948565b601182019050919050565b600061599f826158f4565b91506159ab8285615917565b91506159b682615971565b91506159c28284615917565b91508190509392505050565b60006159d982614c2d565b6159e381856142f1565b93506159f3818560208601614400565b6159fc81613ba0565b840191505092915050565b60006020820190508181036000830152615a2181846159ce565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000615a638261399c565b915060008203615a7657615a75614801565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000615ab76020836142f1565b9150615ac282615a81565b602082019050919050565b60006020820190508181036000830152615ae681615aaa565b905091905056fea26469706673582212202cd3d8df996360b60469ce170aff60ab01025b96822a958bc7a1c562854b6a7d64736f6c63430008110033",
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

// PROVIDERREDEMPTIONTAX is a free data retrieval call binding the contract method 0x0948ee63.
//
// Solidity: function PROVIDER_REDEMPTION_TAX() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) PROVIDERREDEMPTIONTAX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "PROVIDER_REDEMPTION_TAX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROVIDERREDEMPTIONTAX is a free data retrieval call binding the contract method 0x0948ee63.
//
// Solidity: function PROVIDER_REDEMPTION_TAX() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) PROVIDERREDEMPTIONTAX() (*big.Int, error) {
	return _MetaScheduler.Contract.PROVIDERREDEMPTIONTAX(&_MetaScheduler.CallOpts)
}

// PROVIDERREDEMPTIONTAX is a free data retrieval call binding the contract method 0x0948ee63.
//
// Solidity: function PROVIDER_REDEMPTION_TAX() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) PROVIDERREDEMPTIONTAX() (*big.Int, error) {
	return _MetaScheduler.Contract.PROVIDERREDEMPTIONTAX(&_MetaScheduler.CallOpts)
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

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0x6df90c9d.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,string) jobDefinition, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint64)
func (_MetaScheduler *MetaSchedulerCaller) ConvertCreditToDuration(opts *bind.CallOpts, jobDefinition JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (uint64, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "convertCreditToDuration", jobDefinition, amountLocked, provider)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0x6df90c9d.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,string) jobDefinition, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint64)
func (_MetaScheduler *MetaSchedulerSession) ConvertCreditToDuration(jobDefinition JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (uint64, error) {
	return _MetaScheduler.Contract.ConvertCreditToDuration(&_MetaScheduler.CallOpts, jobDefinition, amountLocked, provider)
}

// ConvertCreditToDuration is a free data retrieval call binding the contract method 0x6df90c9d.
//
// Solidity: function convertCreditToDuration((uint64,uint64,uint64,uint64,string) jobDefinition, uint256 amountLocked, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider) pure returns(uint64)
func (_MetaScheduler *MetaSchedulerCallerSession) ConvertCreditToDuration(jobDefinition JobDefinition, amountLocked *big.Int, provider ProviderDefinition) (uint64, error) {
	return _MetaScheduler.Contract.ConvertCreditToDuration(&_MetaScheduler.CallOpts, jobDefinition, amountLocked, provider)
}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0x33474c71.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,string) jobDefinition, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint64 durationMinute) pure returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) ConvertDurationToCredit(opts *bind.CallOpts, jobDefinition JobDefinition, provider ProviderDefinition, durationMinute uint64) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "convertDurationToCredit", jobDefinition, provider, durationMinute)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0x33474c71.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,string) jobDefinition, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint64 durationMinute) pure returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) ConvertDurationToCredit(jobDefinition JobDefinition, provider ProviderDefinition, durationMinute uint64) (*big.Int, error) {
	return _MetaScheduler.Contract.ConvertDurationToCredit(&_MetaScheduler.CallOpts, jobDefinition, provider, durationMinute)
}

// ConvertDurationToCredit is a free data retrieval call binding the contract method 0x33474c71.
//
// Solidity: function convertDurationToCredit((uint64,uint64,uint64,uint64,string) jobDefinition, (uint64,uint64,uint64,uint64,uint64,uint64,uint64) provider, uint64 durationMinute) pure returns(uint256)
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

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerCaller) HasCancellingJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "hasCancellingJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasCancellingJob(&_MetaScheduler.CallOpts, _providerAddr)
}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerCallerSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasCancellingJob(&_MetaScheduler.CallOpts, _providerAddr)
}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerCaller) HasNextJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "hasNextJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerSession) HasNextJob(_providerAddr common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasNextJob(&_MetaScheduler.CallOpts, _providerAddr)
}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerCallerSession) HasNextJob(_providerAddr common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasNextJob(&_MetaScheduler.CallOpts, _providerAddr)
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

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _jobStatus, uint64 _jobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderSetJobStatus(opts *bind.TransactOpts, _jobId [32]byte, _jobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerSetJobStatus", _jobId, _jobStatus, _jobDurationMinute)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _jobStatus, uint64 _jobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderSetJobStatus(_jobId [32]byte, _jobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderSetJobStatus(&_MetaScheduler.TransactOpts, _jobId, _jobStatus, _jobDurationMinute)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _jobStatus, uint64 _jobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderSetJobStatus(_jobId [32]byte, _jobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderSetJobStatus(&_MetaScheduler.TransactOpts, _jobId, _jobStatus, _jobDurationMinute)
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

// RequestNewJob is a paid mutator transaction binding the contract method 0xc4b6e0cc.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked, bytes32 _jobName) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactor) RequestNewJob(opts *bind.TransactOpts, _definition JobDefinition, _amountLocked *big.Int, _jobName [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "requestNewJob", _definition, _amountLocked, _jobName)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0xc4b6e0cc.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked, bytes32 _jobName) returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) RequestNewJob(_definition JobDefinition, _amountLocked *big.Int, _jobName [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _amountLocked, _jobName)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0xc4b6e0cc.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string) _definition, uint256 _amountLocked, bytes32 _jobName) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactorSession) RequestNewJob(_definition JobDefinition, _amountLocked *big.Int, _jobName [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _amountLocked, _jobName)
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

// FilterClaimNextJobEvent is a free log retrieval operation binding the contract event 0x6103c1800236e5e6e1833bde789240a0577a25b448a539ae39471a9e75d04225.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) FilterClaimNextJobEvent(opts *bind.FilterOpts) (*MetaSchedulerClaimNextJobEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "ClaimNextJobEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerClaimNextJobEventIterator{contract: _MetaScheduler.contract, event: "ClaimNextJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimNextJobEvent is a free log subscription operation binding the contract event 0x6103c1800236e5e6e1833bde789240a0577a25b448a539ae39471a9e75d04225.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string) jobDefinition)
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

// ParseClaimNextJobEvent is a log parse operation binding the contract event 0x6103c1800236e5e6e1833bde789240a0577a25b448a539ae39471a9e75d04225.
//
// Solidity: event ClaimNextJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string) jobDefinition)
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
	Bin: "0x60806040523480156200001157600080fd5b50600060018054906101000a900460ff1615905080801562000044575060018060009054906101000a900460ff1660ff16105b8062000081575062000061306200016360201b62002c8d1760201c565b15801562000080575060018060009054906101000a900460ff1660ff16145b5b620000c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ba906200020d565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015620001005760018060016101000a81548160ff0219169083151502179055505b80156200015c5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405162000153919062000289565b60405180910390a15b50620002a6565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620001f5602e8362000186565b9150620002028262000197565b604082019050919050565b600060208201905081810360008301526200022881620001e6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620002716200026b62000265846200022f565b62000246565b62000239565b9050919050565b620002838162000250565b82525050565b6000602082019050620002a0600083018462000278565b92915050565b6144d680620002b66000396000f3fe608060405234801561001057600080fd5b50600436106101a85760003560e01c8063624bc8e3116100f9578063a94aa1a611610097578063bb26e6e811610071578063bb26e6e8146104f9578063d547741f14610515578063daea85c514610531578063e7692d091461054d576101a8565b8063a94aa1a614610477578063b7bb9145146104ad578063b7c8116e146104dd576101a8565b806391d14854116100d357806391d14854146103ef578063939daf9c1461041f57806396c551751461043d578063a217fddf14610459576101a8565b8063624bc8e3146103855780638129fc1c146103b5578063877f4e12146103bf576101a8565b80632f2ff15d1161016657806347799da81161014057806347799da8146102f75780634b939bd41461031b5780635c42d079146103395780635e9da3ed14610369576101a8565b80632f2ff15d146102a357806336568abe146102bf5780633f6edb5f146102db576101a8565b80623052a6146101ad57806301ffc9a7146101cb57806306661abd146101fb5780630787bc27146102195780630fc5d8031461024f578063248a9ca314610273575b600080fd5b6101b5610569565b6040516101c29190613354565b60405180910390f35b6101e560048036038101906101e091906133cc565b610573565b6040516101f29190613414565b60405180910390f35b6102036105ed565b6040516102109190613354565b60405180910390f35b610233600480360381019061022e919061348d565b6105f3565b6040516102469796959493929190613600565b60405180910390f35b6102576107eb565b60405161026a9796959493929190613600565b60405180910390f35b61028d600480360381019061028891906136ab565b6109d1565b60405161029a91906136e7565b60405180910390f35b6102bd60048036038101906102b89190613702565b6109f0565b005b6102d960048036038101906102d49190613702565b610a11565b005b6102f560048036038101906102f0919061348d565b610a94565b005b6102ff610b7f565b6040516103129796959493929190613600565b60405180910390f35b610323610d65565b6040516103309190613354565b60405180910390f35b610353600480360381019061034e919061376e565b610d6b565b60405161036091906138f9565b60405180910390f35b610383600480360381019061037e9190613941565b611066565b005b61039f600480360381019061039a919061348d565b611642565b6040516103ac91906138f9565b60405180910390f35b6103bd611955565b005b6103d960048036038101906103d4919061348d565b611ad2565b6040516103e69190613414565b60405180910390f35b61040960048036038101906104049190613702565b611b52565b6040516104169190613414565b60405180910390f35b610427611bbc565b60405161043491906139f7565b60405180910390f35b6104576004803603810190610452919061348d565b611bd6565b005b610461611ce7565b60405161046e91906136e7565b60405180910390f35b610491600480360381019061048c919061376e565b611cee565b6040516104a49796959493929190613600565b60405180910390f35b6104c760048036038101906104c2919061348d565b611ee6565b6040516104d49190613414565b60405180910390f35b6104f760048036038101906104f29190613941565b611f65565b005b610513600480360381019061050e9190613a37565b61248b565b005b61052f600480360381019061052a9190613702565b6125de565b005b61054b6004803603810190610546919061348d565b6125ff565b005b6105676004803603810190610562919061348d565b612b7b565b005b6000601154905090565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806105e657506105e582612cb0565b5b9050919050565b60125481565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b60098060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b6000806000838152602001908152602001600020600101549050919050565b6109f9826109d1565b610a0281612d1a565b610a0c8383612d2e565b505050565b610a19612e0e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7d90613afa565b60405180910390fd5b610a908282612e16565b5050565b6000801b610aa181612d1a565b600f600081819054906101000a900467ffffffffffffffff1680929190610ac790613b49565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600281819054906101000a900467ffffffffffffffff1680929190610b5490613b49565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b60038060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b60115481565b610d73613241565b600115156010600084815260200190815260200160002060030160019054906101000a900460ff16151514610ddd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd490613bc5565b60405180910390fd5b601060008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020016003820160009054906101000a900460ff166002811115610fe857610fe761357a565b5b6002811115610ffa57610ff961357a565b5b81526020016003820160019054906101000a900460ff161515151581526020016003820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600482015481526020016005820154815250509050919050565b6000801b61107381612d1a565b60008867ffffffffffffffff16116110c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110b790613c31565b60405180910390fd5b60008667ffffffffffffffff161161110d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161110490613c9d565b60405180910390fd5b60008567ffffffffffffffff161161115a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115190613d09565b60405180910390fd5b60011515600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff161515146111f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e790613d75565b60405180910390fd5b600160028111156112045761120361357a565b5b600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff1660028111156112665761126561357a565b5b146112a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161129d90613de1565b60405180910390fd5b60006040518060e001604052808a67ffffffffffffffff1681526020018967ffffffffffffffff1681526020018667ffffffffffffffff1681526020018867ffffffffffffffff1681526020018567ffffffffffffffff1681526020018767ffffffffffffffff1681526020018467ffffffffffffffff1681525090506040518060e001604052808b73ffffffffffffffffffffffffffffffffffffffff168152602001828152602001600060028111156113645761136361357a565b5b8152602001600115158152602001600067ffffffffffffffff168152602001600081526020016000815250600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505060408201518160030160006101000a81548160ff021916908360028111156115945761159361357a565b5b021790555060608201518160030160016101000a81548160ff02191690831515021790555060808201518160030160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816004015560c082015181600501559050507f3aeb53b0dee89ac04567fa6305e626e8d5246b478acd34d0a217507b9dfd076c8a60405161162e9190613e01565b60405180910390a150505050505050505050565b61164a613241565b6000801b61165781612d1a565b61166083611ad2565b61169f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169690613e8e565b60405180910390fd5b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020016003820160009054906101000a900460ff1660028111156118d6576118d561357a565b5b60028111156118e8576118e761357a565b5b81526020016003820160019054906101000a900460ff161515151581526020016003820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160048201548152602001600582015481525050915050919050565b600060018054906101000a900460ff16159050808015611986575060018060009054906101000a900460ff1660ff16105b806119b4575061199530612c8d565b1580156119b3575060018060009054906101000a900460ff1660ff16145b5b6119f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119ea90613f20565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015611a2f5760018060016101000a81548160ff0219169083151502179055505b611a3c6000801b33612d2e565b6000600f60006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600060118190555060016012819055508015611acf5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051611ac69190613f92565b60405180910390a15b50565b600060016002811115611ae857611ae761357a565b5b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff166002811115611b4a57611b4961357a565b5b149050919050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600f60009054906101000a900467ffffffffffffffff1681565b6000801b611be381612d1a565b60011515600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff16151514611c79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c7090613bc5565b60405180910390fd5b60028060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548160ff02191690836002811115611cde57611cdd61357a565b5b02179055505050565b6000801b81565b60106020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16908060030160029054906101000a900467ffffffffffffffff16908060040154908060050154905087565b6000600280811115611efb57611efa61357a565b5b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff166002811115611f5d57611f5c61357a565b5b149050919050565b6000801b611f7281612d1a565b60008867ffffffffffffffff1611611fbf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fb690613c31565b60405180910390fd5b60008667ffffffffffffffff161161200c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161200390613c9d565b60405180910390fd5b60008567ffffffffffffffff1611612059576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161205090613d09565b60405180910390fd5b60011515600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff161515036120ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120e690613ff9565b60405180910390fd5b60006040518060e001604052808a67ffffffffffffffff1681526020018967ffffffffffffffff1681526020018667ffffffffffffffff1681526020018867ffffffffffffffff1681526020018567ffffffffffffffff1681526020018767ffffffffffffffff1681526020018467ffffffffffffffff1681525090506040518060e001604052808b73ffffffffffffffffffffffffffffffffffffffff168152602001828152602001600060028111156121ad576121ac61357a565b5b8152602001600115158152602001600067ffffffffffffffff168152602001600081526020016000815250600260008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505060408201518160030160006101000a81548160ff021916908360028111156123dd576123dc61357a565b5b021790555060608201518160030160016101000a81548160ff02191690831515021790555060808201518160030160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a0820151816004015560c082015181600501559050507fc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce28758a6040516124779190613e01565b60405180910390a150505050505050505050565b6000801b61249881612d1a565b600160088111156124ac576124ab61357a565b5b8360088111156124bf576124be61357a565b5b146124ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124f69061408b565b60405180910390fd5b600f600081819054906101000a900467ffffffffffffffff1680929190612525906140ab565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600281819054906101000a900467ffffffffffffffff16809291906125b2906140ab565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050565b6125e7826109d1565b6125f081612d1a565b6125fa8383612e16565b505050565b6000801b61260c81612d1a565b60011515600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff161515146126a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161269990613bc5565b60405180910390fd5b600060028111156126b6576126b561357a565b5b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff1660028111156127185761271761357a565b5b14612758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161274f90614120565b60405180910390fd5b6001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548160ff021916908360028111156127be576127bd61357a565b5b02179055506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206005015403612b7757600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206010600060125481526020019081526020016000206000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018201816001016000820160009054906101000a900467ffffffffffffffff168160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160089054906101000a900467ffffffffffffffff168160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160109054906101000a900467ffffffffffffffff168160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506000820160189054906101000a900467ffffffffffffffff168160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160089054906101000a900467ffffffffffffffff168160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160109054906101000a900467ffffffffffffffff168160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050506003820160009054906101000a900460ff168160030160006101000a81548160ff02191690836002811115612adf57612ade61357a565b5b02179055506003820160019054906101000a900460ff168160030160016101000a81548160ff0219169083151502179055506003820160029054906101000a900467ffffffffffffffff168160030160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820154816004015560058201548160050155905050612b76601254612ef7565b5b5050565b6000801b612b8881612d1a565b60011515600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160019054906101000a900460ff16151514612c1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c1590613bc5565b60405180910390fd5b6001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160006101000a81548160ff02191690836002811115612c8457612c8361357a565b5b02179055505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b612d2b81612d26612e0e565b612f53565b50565b612d388282611b52565b612e0a57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612daf612e0e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b612e208282611b52565b15612ef357600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612e98612e0e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b806010600060115481526020019081526020016000206005018190555060115460106000838152602001908152602001600020600401819055508060118190555060126000815480929190612f4b90614140565b919050555050565b612f5d8282611b52565b612fd457612f6a81612fd8565b612f788360001c6020613005565b604051602001612f89929190614291565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fcb9190614315565b60405180910390fd5b5050565b6060612ffe8273ffffffffffffffffffffffffffffffffffffffff16601460ff16613005565b9050919050565b6060600060028360026130189190614337565b6130229190614379565b67ffffffffffffffff81111561303b5761303a6143ad565b5b6040519080825280601f01601f19166020018201604052801561306d5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106130a5576130a46143dc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110613109576131086143dc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026131499190614337565b6131539190614379565b90505b60018111156131f3577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110613195576131946143dc565b5b1a60f81b8282815181106131ac576131ab6143dc565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806131ec9061440b565b9050613156565b5060008414613237576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161322e90614480565b60405180910390fd5b8091505092915050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016132716132b8565b81526020016000600281111561328a5761328961357a565b5b8152602001600015158152602001600067ffffffffffffffff16815260200160008152602001600081525090565b6040518060e00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000819050919050565b61334e8161333b565b82525050565b60006020820190506133696000830184613345565b92915050565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6133a981613374565b81146133b457600080fd5b50565b6000813590506133c6816133a0565b92915050565b6000602082840312156133e2576133e161336f565b5b60006133f0848285016133b7565b91505092915050565b60008115159050919050565b61340e816133f9565b82525050565b60006020820190506134296000830184613405565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061345a8261342f565b9050919050565b61346a8161344f565b811461347557600080fd5b50565b60008135905061348781613461565b92915050565b6000602082840312156134a3576134a261336f565b5b60006134b184828501613478565b91505092915050565b6134c38161344f565b82525050565b600067ffffffffffffffff82169050919050565b6134e6816134c9565b82525050565b60e08201600082015161350260008501826134dd565b50602082015161351560208501826134dd565b50604082015161352860408501826134dd565b50606082015161353b60608501826134dd565b50608082015161354e60808501826134dd565b5060a082015161356160a08501826134dd565b5060c082015161357460c08501826134dd565b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106135ba576135b961357a565b5b50565b60008190506135cb826135a9565b919050565b60006135db826135bd565b9050919050565b6135eb816135d0565b82525050565b6135fa816134c9565b82525050565b60006101a082019050613616600083018a6134ba565b61362360208301896134ec565b6136316101008301886135e2565b61363f610120830187613405565b61364d6101408301866135f1565b61365b610160830185613345565b613669610180830184613345565b98975050505050505050565b6000819050919050565b61368881613675565b811461369357600080fd5b50565b6000813590506136a58161367f565b92915050565b6000602082840312156136c1576136c061336f565b5b60006136cf84828501613696565b91505092915050565b6136e181613675565b82525050565b60006020820190506136fc60008301846136d8565b92915050565b600080604083850312156137195761371861336f565b5b600061372785828601613696565b925050602061373885828601613478565b9150509250929050565b61374b8161333b565b811461375657600080fd5b50565b60008135905061376881613742565b92915050565b6000602082840312156137845761378361336f565b5b600061379284828501613759565b91505092915050565b6137a48161344f565b82525050565b60e0820160008201516137c060008501826134dd565b5060208201516137d360208501826134dd565b5060408201516137e660408501826134dd565b5060608201516137f960608501826134dd565b50608082015161380c60808501826134dd565b5060a082015161381f60a08501826134dd565b5060c082015161383260c08501826134dd565b50505050565b613841816135d0565b82525050565b613850816133f9565b82525050565b61385f8161333b565b82525050565b6101a08201600082015161387c600085018261379b565b50602082015161388f60208501826137aa565b5060408201516138a3610100850182613838565b5060608201516138b7610120850182613847565b5060808201516138cb6101408501826134dd565b5060a08201516138df610160850182613856565b5060c08201516138f3610180850182613856565b50505050565b60006101a08201905061390f6000830184613865565b92915050565b61391e816134c9565b811461392957600080fd5b50565b60008135905061393b81613915565b92915050565b600080600080600080600080610100898b0312156139625761396161336f565b5b60006139708b828c01613478565b98505060206139818b828c0161392c565b97505060406139928b828c0161392c565b96505060606139a38b828c0161392c565b95505060806139b48b828c0161392c565b94505060a06139c58b828c0161392c565b93505060c06139d68b828c0161392c565b92505060e06139e78b828c0161392c565b9150509295985092959890939650565b6000602082019050613a0c60008301846135f1565b92915050565b60098110613a1f57600080fd5b50565b600081359050613a3181613a12565b92915050565b60008060408385031215613a4e57613a4d61336f565b5b6000613a5c85828601613a22565b9250506020613a6d85828601613478565b9150509250929050565b600082825260208201905092915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000613ae4602f83613a77565b9150613aef82613a88565b604082019050919050565b60006020820190508181036000830152613b1381613ad7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000613b54826134c9565b915067ffffffffffffffff8203613b6e57613b6d613b1a565b5b600182019050919050565b7f4e6f2070726f766964657220666f756e64000000000000000000000000000000600082015250565b6000613baf601183613a77565b9150613bba82613b79565b602082019050919050565b60006020820190508181036000830152613bde81613ba2565b9050919050565b7f6e6e6f646573206d7573742062652067726561746572207468616e2030000000600082015250565b6000613c1b601d83613a77565b9150613c2682613be5565b602082019050919050565b60006020820190508181036000830152613c4a81613c0e565b9050919050565b7f63707573206d7573742062652067726561746572207468616e20300000000000600082015250565b6000613c87601b83613a77565b9150613c9282613c51565b602082019050919050565b60006020820190508181036000830152613cb681613c7a565b9050919050565b7f6d656d206d7573742062652067726561746572207468616e2030000000000000600082015250565b6000613cf3601a83613a77565b9150613cfe82613cbd565b602082019050919050565b60006020820190508181036000830152613d2281613ce6565b9050919050565b7f50726f7669646572206d75737420626520726567697374657265640000000000600082015250565b6000613d5f601b83613a77565b9150613d6a82613d29565b602082019050919050565b60006020820190508181036000830152613d8e81613d52565b9050919050565b7f50726f76696465722068617665206265656e20617070726f7665640000000000600082015250565b6000613dcb601b83613a77565b9150613dd682613d95565b602082019050919050565b60006020820190508181036000830152613dfa81613dbe565b9050919050565b6000602082019050613e1660008301846134ba565b92915050565b7f50726f766964657273206e6f7420666f756e64206f72206e6f74206a6f696e6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b6000613e78602183613a77565b9150613e8382613e1c565b604082019050919050565b60006020820190508181036000830152613ea781613e6b565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000613f0a602e83613a77565b9150613f1582613eae565b604082019050919050565b60006020820190508181036000830152613f3981613efd565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000613f7c613f77613f7284613f40565b613f57565b613f4a565b9050919050565b613f8c81613f61565b82525050565b6000602082019050613fa76000830184613f83565b92915050565b7f50726f76696465722063616e206e6f7420726567697374657220747769636500600082015250565b6000613fe3601f83613a77565b9150613fee82613fad565b602082019050919050565b6000602082019050818103600083015261401281613fd6565b9050919050565b7f4f6e6c79204d4554415f5343484544554c4544206a6f62732063616e2062652060008201527f6465637265617365640000000000000000000000000000000000000000000000602082015250565b6000614075602983613a77565b915061408082614019565b604082019050919050565b600060208201905081810360008301526140a481614068565b9050919050565b60006140b6826134c9565b9150600082036140c9576140c8613b1a565b5b600182039050919050565b7f50726f766964657220737461747573206d75737420626520494e495400000000600082015250565b600061410a601c83613a77565b9150614115826140d4565b602082019050919050565b60006020820190508181036000830152614139816140fd565b9050919050565b600061414b8261333b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361417d5761417c613b1a565b5b600182019050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b60006141c9601783614188565b91506141d482614193565b601782019050919050565b600081519050919050565b60005b838110156142085780820151818401526020810190506141ed565b60008484015250505050565b600061421f826141df565b6142298185614188565b93506142398185602086016141ea565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b600061427b601183614188565b915061428682614245565b601182019050919050565b600061429c826141bc565b91506142a88285614214565b91506142b38261426e565b91506142bf8284614214565b91508190509392505050565b6000601f19601f8301169050919050565b60006142e7826141df565b6142f18185613a77565b93506143018185602086016141ea565b61430a816142cb565b840191505092915050565b6000602082019050818103600083015261432f81846142dc565b905092915050565b60006143428261333b565b915061434d8361333b565b925082820261435b8161333b565b9150828204841483151761437257614371613b1a565b5b5092915050565b60006143848261333b565b915061438f8361333b565b92508282019050808211156143a7576143a6613b1a565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006144168261333b565b91506000820361442957614428613b1a565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b600061446a602083613a77565b915061447582614434565b602082019050919050565b600060208201905081810360008301526144998161445d565b905091905056fea2646970667358221220a3c39fbf45b8bdfc7e105541429f30747f4c27523bf59ba7bddda8b1035e0b0364736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
	Bin: "0x6101cd610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806315a945ce1461003a575b600080fd5b610054600480360381019061004f91906100be565b61006a565b6040516100619190610119565b60405180910390f35b60008183436100799190610163565b1015905092915050565b600080fd5b6000819050919050565b61009b81610088565b81146100a657600080fd5b50565b6000813590506100b881610092565b92915050565b600080604083850312156100d5576100d4610083565b5b60006100e3858286016100a9565b92505060206100f4858286016100a9565b9150509250929050565b60008115159050919050565b610113816100fe565b82525050565b600060208201905061012e600083018461010a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061016e82610088565b915061017983610088565b925082820390508181111561019157610190610134565b5b9291505056fea26469706673582212208fd0552f68854e3e3d98abdfc92f88a4e02aebfef32f0510b49f2ae333b1105b64736f6c63430008110033",
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
