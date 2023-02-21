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
	_ = abi.ConvertType
)

// JobCost is an auto generated low-level Go binding around an user-defined struct.
type JobCost struct {
	MaxCost   *big.Int
	FinalCost *big.Int
	AutoTopUp bool
}

// JobDefinition is an auto generated low-level Go binding around an user-defined struct.
type JobDefinition struct {
	GpuPerTask        uint64
	MemPerCpu         uint64
	CpuPerTask        uint64
	Ntasks            uint64
	BatchLocationHash string
	StorageType       uint8
}

// JobTime is an auto generated low-level Go binding around an user-defined struct.
type JobTime struct {
	Start                  *big.Int
	End                    *big.Int
	CancelRequestTimestamp *big.Int
	BlockNumberStateChange *big.Int
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
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// IMetaSchedulerMetaData contains all meta data concerning the IMetaScheduler contract.
var IMetaSchedulerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"ClaimNextCancellingJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextCancellingJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasCancellingJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasNextJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_jobDurationMinute\",\"type\":\"uint64\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"autoTopUp\",\"type\":\"bool\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"topUpJobSlice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMetaSchedulerABI is the input ABI used to generate the binding from.
// Deprecated: Use IMetaSchedulerMetaData.ABI instead.
var IMetaSchedulerABI = IMetaSchedulerMetaData.ABI

// IMetaScheduler is an auto generated Go binding around an Ethereum contract.
type IMetaScheduler struct {
	IMetaSchedulerCaller     // Read-only binding to the contract
	IMetaSchedulerTransactor // Write-only binding to the contract
	IMetaSchedulerFilterer   // Log filterer for contract events
}

// IMetaSchedulerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMetaSchedulerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMetaSchedulerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMetaSchedulerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMetaSchedulerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMetaSchedulerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMetaSchedulerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMetaSchedulerSession struct {
	Contract     *IMetaScheduler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMetaSchedulerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMetaSchedulerCallerSession struct {
	Contract *IMetaSchedulerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IMetaSchedulerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMetaSchedulerTransactorSession struct {
	Contract     *IMetaSchedulerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IMetaSchedulerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMetaSchedulerRaw struct {
	Contract *IMetaScheduler // Generic contract binding to access the raw methods on
}

// IMetaSchedulerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMetaSchedulerCallerRaw struct {
	Contract *IMetaSchedulerCaller // Generic read-only contract binding to access the raw methods on
}

// IMetaSchedulerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMetaSchedulerTransactorRaw struct {
	Contract *IMetaSchedulerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMetaScheduler creates a new instance of IMetaScheduler, bound to a specific deployed contract.
func NewIMetaScheduler(address common.Address, backend bind.ContractBackend) (*IMetaScheduler, error) {
	contract, err := bindIMetaScheduler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMetaScheduler{IMetaSchedulerCaller: IMetaSchedulerCaller{contract: contract}, IMetaSchedulerTransactor: IMetaSchedulerTransactor{contract: contract}, IMetaSchedulerFilterer: IMetaSchedulerFilterer{contract: contract}}, nil
}

// NewIMetaSchedulerCaller creates a new read-only instance of IMetaScheduler, bound to a specific deployed contract.
func NewIMetaSchedulerCaller(address common.Address, caller bind.ContractCaller) (*IMetaSchedulerCaller, error) {
	contract, err := bindIMetaScheduler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerCaller{contract: contract}, nil
}

// NewIMetaSchedulerTransactor creates a new write-only instance of IMetaScheduler, bound to a specific deployed contract.
func NewIMetaSchedulerTransactor(address common.Address, transactor bind.ContractTransactor) (*IMetaSchedulerTransactor, error) {
	contract, err := bindIMetaScheduler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerTransactor{contract: contract}, nil
}

// NewIMetaSchedulerFilterer creates a new log filterer instance of IMetaScheduler, bound to a specific deployed contract.
func NewIMetaSchedulerFilterer(address common.Address, filterer bind.ContractFilterer) (*IMetaSchedulerFilterer, error) {
	contract, err := bindIMetaScheduler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerFilterer{contract: contract}, nil
}

// bindIMetaScheduler binds a generic wrapper to an already deployed contract.
func bindIMetaScheduler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMetaSchedulerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMetaScheduler *IMetaSchedulerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMetaScheduler.Contract.IMetaSchedulerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMetaScheduler *IMetaSchedulerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.IMetaSchedulerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMetaScheduler *IMetaSchedulerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.IMetaSchedulerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMetaScheduler *IMetaSchedulerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMetaScheduler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMetaScheduler *IMetaSchedulerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMetaScheduler *IMetaSchedulerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.contract.Transact(opts, method, params...)
}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_IMetaScheduler *IMetaSchedulerCaller) HasCancellingJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _IMetaScheduler.contract.Call(opts, &out, "hasCancellingJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_IMetaScheduler *IMetaSchedulerSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _IMetaScheduler.Contract.HasCancellingJob(&_IMetaScheduler.CallOpts, _providerAddr)
}

// HasCancellingJob is a free data retrieval call binding the contract method 0x20a5f919.
//
// Solidity: function hasCancellingJob(address _providerAddr) view returns(bool)
func (_IMetaScheduler *IMetaSchedulerCallerSession) HasCancellingJob(_providerAddr common.Address) (bool, error) {
	return _IMetaScheduler.Contract.HasCancellingJob(&_IMetaScheduler.CallOpts, _providerAddr)
}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_IMetaScheduler *IMetaSchedulerCaller) HasNextJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _IMetaScheduler.contract.Call(opts, &out, "hasNextJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_IMetaScheduler *IMetaSchedulerSession) HasNextJob(_providerAddr common.Address) (bool, error) {
	return _IMetaScheduler.Contract.HasNextJob(&_IMetaScheduler.CallOpts, _providerAddr)
}

// HasNextJob is a free data retrieval call binding the contract method 0x0797094e.
//
// Solidity: function hasNextJob(address _providerAddr) view returns(bool)
func (_IMetaScheduler *IMetaSchedulerCallerSession) HasNextJob(_providerAddr common.Address) (bool, error) {
	return _IMetaScheduler.Contract.HasNextJob(&_IMetaScheduler.CallOpts, _providerAddr)
}

// CancelJob is a paid mutator transaction binding the contract method 0x5fae1450.
//
// Solidity: function cancelJob(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) CancelJob(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "cancelJob", _jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x5fae1450.
//
// Solidity: function cancelJob(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerSession) CancelJob(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.CancelJob(&_IMetaScheduler.TransactOpts, _jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x5fae1450.
//
// Solidity: function cancelJob(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) CancelJob(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.CancelJob(&_IMetaScheduler.TransactOpts, _jobId)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) ClaimJob(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "claimJob", _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns()
func (_IMetaScheduler *IMetaSchedulerSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ClaimJob(&_IMetaScheduler.TransactOpts, _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ClaimJob(&_IMetaScheduler.TransactOpts, _jobId, _providerAddr)
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0x5e1b2d65.
//
// Solidity: function claimNextCancellingJob() returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) ClaimNextCancellingJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "claimNextCancellingJob")
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0x5e1b2d65.
//
// Solidity: function claimNextCancellingJob() returns()
func (_IMetaScheduler *IMetaSchedulerSession) ClaimNextCancellingJob() (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ClaimNextCancellingJob(&_IMetaScheduler.TransactOpts)
}

// ClaimNextCancellingJob is a paid mutator transaction binding the contract method 0x5e1b2d65.
//
// Solidity: function claimNextCancellingJob() returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) ClaimNextCancellingJob() (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ClaimNextCancellingJob(&_IMetaScheduler.TransactOpts)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) ClaimNextJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "claimNextJob")
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns()
func (_IMetaScheduler *IMetaSchedulerSession) ClaimNextJob() (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ClaimNextJob(&_IMetaScheduler.TransactOpts)
}

// ClaimNextJob is a paid mutator transaction binding the contract method 0x5d3a7180.
//
// Solidity: function claimNextJob() returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) ClaimNextJob() (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ClaimNextJob(&_IMetaScheduler.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.Deposit(&_IMetaScheduler.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.Deposit(&_IMetaScheduler.TransactOpts, _amount)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) MetaSchedule(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "metaSchedule", _jobId, _providerAddr)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_IMetaScheduler *IMetaSchedulerSession) MetaSchedule(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.MetaSchedule(&_IMetaScheduler.TransactOpts, _jobId, _providerAddr)
}

// MetaSchedule is a paid mutator transaction binding the contract method 0xd1cee546.
//
// Solidity: function metaSchedule(bytes32 _jobId, address _providerAddr) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) MetaSchedule(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.MetaSchedule(&_IMetaScheduler.TransactOpts, _jobId, _providerAddr)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _jobStatus, uint64 _jobDurationMinute) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) ProviderSetJobStatus(opts *bind.TransactOpts, _jobId [32]byte, _jobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "providerSetJobStatus", _jobId, _jobStatus, _jobDurationMinute)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _jobStatus, uint64 _jobDurationMinute) returns()
func (_IMetaScheduler *IMetaSchedulerSession) ProviderSetJobStatus(_jobId [32]byte, _jobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ProviderSetJobStatus(&_IMetaScheduler.TransactOpts, _jobId, _jobStatus, _jobDurationMinute)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _jobStatus, uint64 _jobDurationMinute) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) ProviderSetJobStatus(_jobId [32]byte, _jobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.ProviderSetJobStatus(&_IMetaScheduler.TransactOpts, _jobId, _jobStatus, _jobDurationMinute)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x1f92a63f.
//
// Solidity: function refuseJob(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) RefuseJob(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "refuseJob", _jobId)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x1f92a63f.
//
// Solidity: function refuseJob(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerSession) RefuseJob(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.RefuseJob(&_IMetaScheduler.TransactOpts, _jobId)
}

// RefuseJob is a paid mutator transaction binding the contract method 0x1f92a63f.
//
// Solidity: function refuseJob(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) RefuseJob(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.RefuseJob(&_IMetaScheduler.TransactOpts, _jobId)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool autoTopUp) returns(bytes32)
func (_IMetaScheduler *IMetaSchedulerTransactor) RequestNewJob(opts *bind.TransactOpts, _definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, autoTopUp bool) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "requestNewJob", _definition, _maxCost, _jobName, autoTopUp)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool autoTopUp) returns(bytes32)
func (_IMetaScheduler *IMetaSchedulerSession) RequestNewJob(_definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, autoTopUp bool) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.RequestNewJob(&_IMetaScheduler.TransactOpts, _definition, _maxCost, _jobName, autoTopUp)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool autoTopUp) returns(bytes32)
func (_IMetaScheduler *IMetaSchedulerTransactorSession) RequestNewJob(_definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, autoTopUp bool) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.RequestNewJob(&_IMetaScheduler.TransactOpts, _definition, _maxCost, _jobName, autoTopUp)
}

// TopUpJob is a paid mutator transaction binding the contract method 0x2fecc4f6.
//
// Solidity: function topUpJob(bytes32 _jobId, uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) TopUpJob(opts *bind.TransactOpts, _jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "topUpJob", _jobId, _amount)
}

// TopUpJob is a paid mutator transaction binding the contract method 0x2fecc4f6.
//
// Solidity: function topUpJob(bytes32 _jobId, uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerSession) TopUpJob(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.TopUpJob(&_IMetaScheduler.TransactOpts, _jobId, _amount)
}

// TopUpJob is a paid mutator transaction binding the contract method 0x2fecc4f6.
//
// Solidity: function topUpJob(bytes32 _jobId, uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) TopUpJob(_jobId [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.TopUpJob(&_IMetaScheduler.TransactOpts, _jobId, _amount)
}

// TopUpJobSlice is a paid mutator transaction binding the contract method 0x9cdf8d9e.
//
// Solidity: function topUpJobSlice(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) TopUpJobSlice(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "topUpJobSlice", _jobId)
}

// TopUpJobSlice is a paid mutator transaction binding the contract method 0x9cdf8d9e.
//
// Solidity: function topUpJobSlice(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerSession) TopUpJobSlice(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.TopUpJobSlice(&_IMetaScheduler.TransactOpts, _jobId)
}

// TopUpJobSlice is a paid mutator transaction binding the contract method 0x9cdf8d9e.
//
// Solidity: function topUpJobSlice(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) TopUpJobSlice(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.TopUpJobSlice(&_IMetaScheduler.TransactOpts, _jobId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.Withdraw(&_IMetaScheduler.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.Withdraw(&_IMetaScheduler.TransactOpts, _amount)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0xd77836ce.
//
// Solidity: function withdrawAdmin(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) WithdrawAdmin(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "withdrawAdmin", _amount)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0xd77836ce.
//
// Solidity: function withdrawAdmin(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerSession) WithdrawAdmin(_amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.WithdrawAdmin(&_IMetaScheduler.TransactOpts, _amount)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0xd77836ce.
//
// Solidity: function withdrawAdmin(uint256 _amount) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) WithdrawAdmin(_amount *big.Int) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.WithdrawAdmin(&_IMetaScheduler.TransactOpts, _amount)
}

// IMetaSchedulerClaimJobEventIterator is returned from FilterClaimJobEvent and is used to iterate over the raw logs and unpacked data for ClaimJobEvent events raised by the IMetaScheduler contract.
type IMetaSchedulerClaimJobEventIterator struct {
	Event *IMetaSchedulerClaimJobEvent // Event containing the contract specifics and raw log

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
func (it *IMetaSchedulerClaimJobEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMetaSchedulerClaimJobEvent)
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
		it.Event = new(IMetaSchedulerClaimJobEvent)
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
func (it *IMetaSchedulerClaimJobEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMetaSchedulerClaimJobEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMetaSchedulerClaimJobEvent represents a ClaimJobEvent event raised by the IMetaScheduler contract.
type IMetaSchedulerClaimJobEvent struct {
	CustomerAddr      common.Address
	ProviderAddr      common.Address
	JobId             [32]byte
	MaxDurationMinute uint64
	JobDefinition     JobDefinition
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimJobEvent is a free log retrieval operation binding the contract event 0xc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c39.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8) jobDefinition)
func (_IMetaScheduler *IMetaSchedulerFilterer) FilterClaimJobEvent(opts *bind.FilterOpts) (*IMetaSchedulerClaimJobEventIterator, error) {

	logs, sub, err := _IMetaScheduler.contract.FilterLogs(opts, "ClaimJobEvent")
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerClaimJobEventIterator{contract: _IMetaScheduler.contract, event: "ClaimJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimJobEvent is a free log subscription operation binding the contract event 0xc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c39.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8) jobDefinition)
func (_IMetaScheduler *IMetaSchedulerFilterer) WatchClaimJobEvent(opts *bind.WatchOpts, sink chan<- *IMetaSchedulerClaimJobEvent) (event.Subscription, error) {

	logs, sub, err := _IMetaScheduler.contract.WatchLogs(opts, "ClaimJobEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMetaSchedulerClaimJobEvent)
				if err := _IMetaScheduler.contract.UnpackLog(event, "ClaimJobEvent", log); err != nil {
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

// ParseClaimJobEvent is a log parse operation binding the contract event 0xc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c39.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8) jobDefinition)
func (_IMetaScheduler *IMetaSchedulerFilterer) ParseClaimJobEvent(log types.Log) (*IMetaSchedulerClaimJobEvent, error) {
	event := new(IMetaSchedulerClaimJobEvent)
	if err := _IMetaScheduler.contract.UnpackLog(event, "ClaimJobEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMetaSchedulerClaimNextCancellingJobEventIterator is returned from FilterClaimNextCancellingJobEvent and is used to iterate over the raw logs and unpacked data for ClaimNextCancellingJobEvent events raised by the IMetaScheduler contract.
type IMetaSchedulerClaimNextCancellingJobEventIterator struct {
	Event *IMetaSchedulerClaimNextCancellingJobEvent // Event containing the contract specifics and raw log

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
func (it *IMetaSchedulerClaimNextCancellingJobEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMetaSchedulerClaimNextCancellingJobEvent)
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
		it.Event = new(IMetaSchedulerClaimNextCancellingJobEvent)
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
func (it *IMetaSchedulerClaimNextCancellingJobEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMetaSchedulerClaimNextCancellingJobEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMetaSchedulerClaimNextCancellingJobEvent represents a ClaimNextCancellingJobEvent event raised by the IMetaScheduler contract.
type IMetaSchedulerClaimNextCancellingJobEvent struct {
	CustomerAddr common.Address
	ProviderAddr common.Address
	JobId        [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimNextCancellingJobEvent is a free log retrieval operation binding the contract event 0x290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d.
//
// Solidity: event ClaimNextCancellingJobEvent(address customerAddr, address providerAddr, bytes32 jobId)
func (_IMetaScheduler *IMetaSchedulerFilterer) FilterClaimNextCancellingJobEvent(opts *bind.FilterOpts) (*IMetaSchedulerClaimNextCancellingJobEventIterator, error) {

	logs, sub, err := _IMetaScheduler.contract.FilterLogs(opts, "ClaimNextCancellingJobEvent")
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerClaimNextCancellingJobEventIterator{contract: _IMetaScheduler.contract, event: "ClaimNextCancellingJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimNextCancellingJobEvent is a free log subscription operation binding the contract event 0x290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d.
//
// Solidity: event ClaimNextCancellingJobEvent(address customerAddr, address providerAddr, bytes32 jobId)
func (_IMetaScheduler *IMetaSchedulerFilterer) WatchClaimNextCancellingJobEvent(opts *bind.WatchOpts, sink chan<- *IMetaSchedulerClaimNextCancellingJobEvent) (event.Subscription, error) {

	logs, sub, err := _IMetaScheduler.contract.WatchLogs(opts, "ClaimNextCancellingJobEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMetaSchedulerClaimNextCancellingJobEvent)
				if err := _IMetaScheduler.contract.UnpackLog(event, "ClaimNextCancellingJobEvent", log); err != nil {
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
func (_IMetaScheduler *IMetaSchedulerFilterer) ParseClaimNextCancellingJobEvent(log types.Log) (*IMetaSchedulerClaimNextCancellingJobEvent, error) {
	event := new(IMetaSchedulerClaimNextCancellingJobEvent)
	if err := _IMetaScheduler.contract.UnpackLog(event, "ClaimNextCancellingJobEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMetaSchedulerJobRefusedEventIterator is returned from FilterJobRefusedEvent and is used to iterate over the raw logs and unpacked data for JobRefusedEvent events raised by the IMetaScheduler contract.
type IMetaSchedulerJobRefusedEventIterator struct {
	Event *IMetaSchedulerJobRefusedEvent // Event containing the contract specifics and raw log

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
func (it *IMetaSchedulerJobRefusedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMetaSchedulerJobRefusedEvent)
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
		it.Event = new(IMetaSchedulerJobRefusedEvent)
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
func (it *IMetaSchedulerJobRefusedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMetaSchedulerJobRefusedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMetaSchedulerJobRefusedEvent represents a JobRefusedEvent event raised by the IMetaScheduler contract.
type IMetaSchedulerJobRefusedEvent struct {
	JobId        [32]byte
	ProviderAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterJobRefusedEvent is a free log retrieval operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
func (_IMetaScheduler *IMetaSchedulerFilterer) FilterJobRefusedEvent(opts *bind.FilterOpts) (*IMetaSchedulerJobRefusedEventIterator, error) {

	logs, sub, err := _IMetaScheduler.contract.FilterLogs(opts, "JobRefusedEvent")
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerJobRefusedEventIterator{contract: _IMetaScheduler.contract, event: "JobRefusedEvent", logs: logs, sub: sub}, nil
}

// WatchJobRefusedEvent is a free log subscription operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
func (_IMetaScheduler *IMetaSchedulerFilterer) WatchJobRefusedEvent(opts *bind.WatchOpts, sink chan<- *IMetaSchedulerJobRefusedEvent) (event.Subscription, error) {

	logs, sub, err := _IMetaScheduler.contract.WatchLogs(opts, "JobRefusedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMetaSchedulerJobRefusedEvent)
				if err := _IMetaScheduler.contract.UnpackLog(event, "JobRefusedEvent", log); err != nil {
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
func (_IMetaScheduler *IMetaSchedulerFilterer) ParseJobRefusedEvent(log types.Log) (*IMetaSchedulerJobRefusedEvent, error) {
	event := new(IMetaSchedulerJobRefusedEvent)
	if err := _IMetaScheduler.contract.UnpackLog(event, "JobRefusedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMetaSchedulerNewJobRequestEventIterator is returned from FilterNewJobRequestEvent and is used to iterate over the raw logs and unpacked data for NewJobRequestEvent events raised by the IMetaScheduler contract.
type IMetaSchedulerNewJobRequestEventIterator struct {
	Event *IMetaSchedulerNewJobRequestEvent // Event containing the contract specifics and raw log

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
func (it *IMetaSchedulerNewJobRequestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMetaSchedulerNewJobRequestEvent)
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
		it.Event = new(IMetaSchedulerNewJobRequestEvent)
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
func (it *IMetaSchedulerNewJobRequestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMetaSchedulerNewJobRequestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMetaSchedulerNewJobRequestEvent represents a NewJobRequestEvent event raised by the IMetaScheduler contract.
type IMetaSchedulerNewJobRequestEvent struct {
	JobId        [32]byte
	CustomerAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewJobRequestEvent is a free log retrieval operation binding the contract event 0x1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3.
//
// Solidity: event NewJobRequestEvent(bytes32 _jobId, address _customerAddr)
func (_IMetaScheduler *IMetaSchedulerFilterer) FilterNewJobRequestEvent(opts *bind.FilterOpts) (*IMetaSchedulerNewJobRequestEventIterator, error) {

	logs, sub, err := _IMetaScheduler.contract.FilterLogs(opts, "NewJobRequestEvent")
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerNewJobRequestEventIterator{contract: _IMetaScheduler.contract, event: "NewJobRequestEvent", logs: logs, sub: sub}, nil
}

// WatchNewJobRequestEvent is a free log subscription operation binding the contract event 0x1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3.
//
// Solidity: event NewJobRequestEvent(bytes32 _jobId, address _customerAddr)
func (_IMetaScheduler *IMetaSchedulerFilterer) WatchNewJobRequestEvent(opts *bind.WatchOpts, sink chan<- *IMetaSchedulerNewJobRequestEvent) (event.Subscription, error) {

	logs, sub, err := _IMetaScheduler.contract.WatchLogs(opts, "NewJobRequestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMetaSchedulerNewJobRequestEvent)
				if err := _IMetaScheduler.contract.UnpackLog(event, "NewJobRequestEvent", log); err != nil {
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
func (_IMetaScheduler *IMetaSchedulerFilterer) ParseNewJobRequestEvent(log types.Log) (*IMetaSchedulerNewJobRequestEvent, error) {
	event := new(IMetaSchedulerNewJobRequestEvent)
	if err := _IMetaScheduler.contract.UnpackLog(event, "NewJobRequestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProviderManagerMetaData contains all meta data concerning the IProviderManager contract.
var IProviderManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"HardwareUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"ToBeApproved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"getProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPricePerMin\",\"type\":\"uint64\"}],\"internalType\":\"structProviderDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"enumProviderStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"jobCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pointPrevNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointNextNode\",\"type\":\"uint256\"}],\"internalType\":\"structProvider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasJoined\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"incJobCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_nNodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_memPricePerMin\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nNodes\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpus\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cpuPricePerMin\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_memPricePerMin\",\"type\":\"uint64\"}],\"name\":\"registerProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address _providerAddr) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
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
// Solidity: function getProvider(address _providerAddr) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
func (_IProviderManager *IProviderManagerSession) GetProvider(_providerAddr common.Address) (Provider, error) {
	return _IProviderManager.Contract.GetProvider(&_IProviderManager.CallOpts, _providerAddr)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address _providerAddr) view returns((address,(uint64,uint64,uint64,uint64,uint64,uint64,uint64),uint8,bool,uint64,uint256,uint256))
func (_IProviderManager *IProviderManagerCallerSession) GetProvider(_providerAddr common.Address) (Provider, error) {
	return _IProviderManager.Contract.GetProvider(&_IProviderManager.CallOpts, _providerAddr)
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

// Register is a paid mutator transaction binding the contract method 0x6099a872.
//
// Solidity: function register(uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_IProviderManager *IProviderManagerTransactor) Register(opts *bind.TransactOpts, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "register", _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// Register is a paid mutator transaction binding the contract method 0x6099a872.
//
// Solidity: function register(uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_IProviderManager *IProviderManagerSession) Register(_nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _IProviderManager.Contract.Register(&_IProviderManager.TransactOpts, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// Register is a paid mutator transaction binding the contract method 0x6099a872.
//
// Solidity: function register(uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_IProviderManager *IProviderManagerTransactorSession) Register(_nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _IProviderManager.Contract.Register(&_IProviderManager.TransactOpts, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xc01aeabb.
//
// Solidity: function registerProvider(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_IProviderManager *IProviderManagerTransactor) RegisterProvider(opts *bind.TransactOpts, _providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _IProviderManager.contract.Transact(opts, "registerProvider", _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xc01aeabb.
//
// Solidity: function registerProvider(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_IProviderManager *IProviderManagerSession) RegisterProvider(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _IProviderManager.Contract.RegisterProvider(&_IProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xc01aeabb.
//
// Solidity: function registerProvider(address _providerAddr, uint64 _nNodes, uint64 _gpus, uint64 _cpus, uint64 _mem, uint64 _gpuPricePerMin, uint64 _cpuPricePerMin, uint64 _memPricePerMin) returns()
func (_IProviderManager *IProviderManagerTransactorSession) RegisterProvider(_providerAddr common.Address, _nNodes uint64, _gpus uint64, _cpus uint64, _mem uint64, _gpuPricePerMin uint64, _cpuPricePerMin uint64, _memPricePerMin uint64) (*types.Transaction, error) {
	return _IProviderManager.Contract.RegisterProvider(&_IProviderManager.TransactOpts, _providerAddr, _nNodes, _gpus, _cpus, _mem, _gpuPricePerMin, _cpuPricePerMin, _memPricePerMin)
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

// IProviderManagerToBeApprovedIterator is returned from FilterToBeApproved and is used to iterate over the raw logs and unpacked data for ToBeApproved events raised by the IProviderManager contract.
type IProviderManagerToBeApprovedIterator struct {
	Event *IProviderManagerToBeApproved // Event containing the contract specifics and raw log

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
func (it *IProviderManagerToBeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProviderManagerToBeApproved)
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
		it.Event = new(IProviderManagerToBeApproved)
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
func (it *IProviderManagerToBeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProviderManagerToBeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProviderManagerToBeApproved represents a ToBeApproved event raised by the IProviderManager contract.
type IProviderManagerToBeApproved struct {
	ProviderAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterToBeApproved is a free log retrieval operation binding the contract event 0xc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce2875.
//
// Solidity: event ToBeApproved(address _providerAddr)
func (_IProviderManager *IProviderManagerFilterer) FilterToBeApproved(opts *bind.FilterOpts) (*IProviderManagerToBeApprovedIterator, error) {

	logs, sub, err := _IProviderManager.contract.FilterLogs(opts, "ToBeApproved")
	if err != nil {
		return nil, err
	}
	return &IProviderManagerToBeApprovedIterator{contract: _IProviderManager.contract, event: "ToBeApproved", logs: logs, sub: sub}, nil
}

// WatchToBeApproved is a free log subscription operation binding the contract event 0xc15938fb0a298e8c66c8b204cc5d2f80a91e65feff41efb8d4e09117ddce2875.
//
// Solidity: event ToBeApproved(address _providerAddr)
func (_IProviderManager *IProviderManagerFilterer) WatchToBeApproved(opts *bind.WatchOpts, sink chan<- *IProviderManagerToBeApproved) (event.Subscription, error) {

	logs, sub, err := _IProviderManager.contract.WatchLogs(opts, "ToBeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProviderManagerToBeApproved)
				if err := _IProviderManager.contract.UnpackLog(event, "ToBeApproved", log); err != nil {
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
func (_IProviderManager *IProviderManagerFilterer) ParseToBeApproved(log types.Log) (*IProviderManagerToBeApproved, error) {
	event := new(IProviderManagerToBeApproved)
	if err := _IProviderManager.contract.UnpackLog(event, "ToBeApproved", log); err != nil {
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
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	ABI: "[{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"ClaimNextCancellingJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BILL_DURATION_DELTA_MINUTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BILL_TIME_CONTROL_DELTA_S\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCELLATION_FEE_MINUTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEEPSQUARE_CUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METASCHEDULER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOP_UP_SLICE_DURATION_MIN\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimJobTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextCancellingJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getUnlockBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasCancellingJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasNextJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotJobList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_credit\",\"type\":\"address\"},{\"internalType\":\"contractIProviderManager\",\"name\":\"_providerManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"autoTopUp\",\"type\":\"bool\"}],\"internalType\":\"structJobCost\",\"name\":\"cost\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelRequestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"}],\"internalType\":\"structJobTime\",\"name\":\"time\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"hasCancelRequest\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerCancellingJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimableJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerManager\",\"outputs\":[{\"internalType\":\"contractIProviderManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_nextJobStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_jobDurationMinute\",\"type\":\"uint64\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerTimeoutJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_autoTopUp\",\"type\":\"bool\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_autoTopUp\",\"type\":\"bool\"}],\"name\":\"setAutoTopUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"topUpJobSlice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallet2JobId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wallet2LockedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wallet2TotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50619e7b80620000226000396000f3fe608060405234801561001057600080fd5b506004361061027f5760003560e01c806348841b9c1161015c5780639cdf8d9e116100ce578063d1cee54611610087578063d1cee546146107f0578063d547741f1461080c578063d6aa37a614610828578063d77836ce14610846578063e052888c14610862578063e88fe8ca146108805761027f565b80639cdf8d9e1461072e578063a06d083c1461074a578063a217fddf14610768578063aef3276f14610786578063b6b55f25146107b6578063c51067db146107d25761027f565b80635fae1450116101205780635fae14501461065c5780637a9d0df7146106785780638964b048146106965780638fb70f63146106c657806391d14854146106e25780639b06ecad146107125761027f565b806348841b9c146105cc5780634b845f4c146105e85780635d3a7180146106185780635d76d0d1146106225780635e1b2d65146106525761027f565b8063257d9bb8116101f557806336568abe116101b957806336568abe146104dc578063374c6e0e146104f857806338ed7cfc146105165780633c2fb3da1461054f578063407969ae1461057f578063485cc955146105b05761027f565b8063257d9bb8146104395780632e1a7d4d146104575780632f2ff15d146104735780632fecc4f61461048f57806331c3b874146104ab5761027f565b80631a3cbef4116102475780631a3cbef4146103515780631a91c3e9146103815780631df44a3c1461039f5780631f92a63f146103bd57806320a5f919146103d9578063248a9ca3146104095761027f565b806301ffc9a7146102845780630797094e146102b45780630f5a11ca146102e4578063110e87a61461030257806313151ec914610333575b600080fd5b61029e600480360381019061029991906173e0565b6108b0565b6040516102ab9190617428565b60405180910390f35b6102ce60048036038101906102c991906174a1565b61092a565b6040516102db9190617428565b60405180910390f35b6102ec61097c565b6040516102f991906174e7565b60405180910390f35b61031c600480360381019061031791906174a1565b610982565b60405161032a92919061751e565b60405180910390f35b61033b6109c0565b60405161034891906175a6565b60405180910390f35b61036b600480360381019061036691906174a1565b6109e6565b6040516103789190617689565b60405180910390f35b610389610a7d565b60405161039691906176ce565b60405180910390f35b6103a7610a82565b6040516103b491906174e7565b60405180910390f35b6103d760048036038101906103d29190617715565b610a87565b005b6103f360048036038101906103ee91906174a1565b610d57565b6040516104009190617428565b60405180910390f35b610423600480360381019061041e9190617715565b610da7565b6040516104309190617751565b60405180910390f35b610441610dc6565b60405161044e91906174e7565b60405180910390f35b610471600480360381019061046c9190617798565b610dd3565b005b61048d600480360381019061048891906177c5565b610f18565b005b6104a960048036038101906104a49190617805565b610f39565b005b6104c560048036038101906104c091906174a1565b6112a7565b6040516104d392919061751e565b60405180910390f35b6104f660048036038101906104f191906177c5565b6112e5565b005b610500611368565b60405161050d91906174e7565b60405180910390f35b610530600480360381019061052b9190617715565b61136d565b6040516105469a99989796959493929190617af0565b60405180910390f35b61056960048036038101906105649190617e0c565b611632565b6040516105769190617751565b60405180910390f35b610599600480360381019061059491906174a1565b611c62565b6040516105a792919061751e565b60405180910390f35b6105ca60048036038101906105c59190617f0b565b611ca0565b005b6105e660048036038101906105e19190617f70565b611f7f565b005b61060260048036038101906105fd91906174a1565b612678565b60405161060f91906174e7565b60405180910390f35b61062061270b565b005b61063c600480360381019061063791906174a1565b612842565b60405161064991906174e7565b60405180910390f35b61065a61285a565b005b61067660048036038101906106719190617715565b6129f6565b005b610680612c9a565b60405161068d91906174e7565b60405180910390f35b6106b060048036038101906106ab91906174a1565b612c9f565b6040516106bd91906174e7565b60405180910390f35b6106e060048036038101906106db91906177c5565b612cb7565b005b6106fc60048036038101906106f791906177c5565b612f50565b6040516107099190617428565b60405180910390f35b61072c60048036038101906107279190617fc3565b612fba565b005b61074860048036038101906107439190617715565b613165565b005b610752613c55565b60405161075f9190618024565b60405180910390f35b610770613c7b565b60405161077d9190617751565b60405180910390f35b6107a0600480360381019061079b9190617798565b613c82565b6040516107ad9190617751565b60405180910390f35b6107d060048036038101906107cb9190617798565b613ca6565b005b6107da613da2565b6040516107e791906174e7565b60405180910390f35b61080a600480360381019061080591906177c5565b613da7565b005b610826600480360381019061082191906177c5565b613f78565b005b610830613f99565b60405161083d91906176ce565b60405180910390f35b610860600480360381019061085b9190617798565b613fb3565b005b61086a61407e565b6040516108779190617751565b60405180910390f35b61089a6004803603810190610895919061803f565b6140a2565b6040516108a79190617751565b60405180910390f35b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806109235750610922826140d3565b5b9050919050565b60006109756001603c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061413d565b9050919050565b60365481565b603c6020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060603960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610a7157602002820191906000526020600020905b815481526020019060010190808311610a5d575b50505050509050919050565b601e81565b601481565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401610ae3919061807f565b602060405180830381865afa158015610b00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2491906180af565b610b63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5a90618139565b60405180910390fd5b816038600082815260200190815260200160002060060160009054906101000a900460ff16610bc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bbe906181a5565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166038600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610c6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6290618211565b60405180910390fd5b60026007811115610c7f57610c7e617845565b5b6038600085815260200190815260200160002060010160009054906101000a900460ff166007811115610cb557610cb4617845565b5b1480610d08575060016007811115610cd057610ccf617845565b5b6038600085815260200190815260200160002060010160009054906101000a900460ff166007811115610d0657610d05617845565b5b145b610d47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3e906182a3565b60405180910390fd5b610d5283600061453d565b505050565b6000610da0603e60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061521d565b9050919050565b6000806000838152602001908152602001600020600101549050919050565b68056bc75e2d6310000081565b610ddc33612678565b811115610e1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1590618335565b60405180910390fd5b80603b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e6d9190618384565b92505081905550603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610ed19291906183b8565b6020604051808303816000875af1158015610ef0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1491906180af565b5050565b610f2182610da7565b610f2a816155f8565b610f34838361560c565b505050565b816038600082815260200190815260200160002060060160009054906101000a900460ff16610f9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f94906181a5565b60405180910390fd5b82610fca6038600083815260200190815260200160002060010160009054906101000a900460ff166156ec565b611009576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110009061842d565b60405180910390fd5b6038600085815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a490618499565b60405180910390fd5b603a60006038600087815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054603b60006038600088815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111a39190618384565b8311156111e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111dc90618505565b60405180910390fd5b82603a60006038600088815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461126a9190618525565b925050819055508260386000868152602001908152602001600020600701600001600082825461129a9190618525565b9250508190555050505050565b603d6020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b6112ed6156ff565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461135a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611351906185cb565b60405180910390fd5b6113648282615707565b5050565b600581565b60386020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546114cd9061861a565b80601f01602080910402602001604051908101604052809291908181526020018280546114f99061861a565b80156115465780601f1061151b57610100808354040283529160200191611546565b820191906000526020600020905b81548152906001019060200180831161152957829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561157457611573617845565b5b600481111561158657611585617845565b5b81525050908060060160009054906101000a900460ff16908060070160405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250509080600a016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250509080600e01549080600f0160009054906101000a900460ff1690508a565b600068056bc75e2d63100000841015611680576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161167790618697565b60405180910390fd5b83603a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054603b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461170b9190618384565b101561174c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174390618505565b60405180910390fd5b6000856060015167ffffffffffffffff1611801561177857506000856040015167ffffffffffffffff16115b801561179257506000856020015167ffffffffffffffff16115b6117d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117c890618703565b60405180910390fd5b83603a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546118209190618525565b9250508190555060006036600081548092919061183c90618723565b9190505560001b90506040518061014001604052808281526020016000600781111561186b5761186a617845565b5b81526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020018781526020016001151581526020016040518060600160405280888152602001600081526020018615158152508152602001604051806080016040528042815260200142815260200142815260200143815250815260200185815260200160001515815250603860008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff0219169083600781111561195657611955617845565b5b021790555060408201518160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506080820151816001019081611ac3919061890d565b5060a08201518160020160006101000a81548160ff02191690836004811115611aef57611aee617845565b5b0217905550505060a08201518160060160006101000a81548160ff02191690831515021790555060c082015181600701600082015181600001556020820151816001015560408201518160020160006101000a81548160ff021916908315150217905550505060e082015181600a0160008201518160000155602082015181600101556040820151816002015560608201518160030155505061010082015181600e015561012082015181600f0160006101000a81548160ff021916908315150217905550905050603960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150557f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada38133604051611c4e9291906189df565b60405180910390a180915050949350505050565b603e6020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b600060018054906101000a900460ff16159050808015611cd1575060018060009054906101000a900460ff1660ff16105b80611cff5750611ce0306157e8565b158015611cfe575060018060009054906101000a900460ff1660ff16145b5b611d3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d3590618a7a565b60405180910390fd5b60018060006101000a81548160ff021916908360ff1602179055508015611d7a5760018060016101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611de9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611de090618b0c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611e58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e4f90618b9e565b60405180910390fd5b611e656000801b3361560c565b82603460086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600a603460006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160368190555060006037819055508015611f7a5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051611f719190618c06565b60405180910390a15b505050565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401611fdb919061807f565b602060405180830381865afa158015611ff8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061201c91906180af565b61205b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161205290618139565b60405180910390fd5b836038600082815260200190815260200160002060060160009054906101000a900460ff166120bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b6906181a5565b60405180910390fd5b6120c761580b565b3373ffffffffffffffffffffffffffffffffffffffff166038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461216b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161216290618c93565b60405180910390fd5b6003600781111561217f5761217e617845565b5b84600781111561219257612191617845565b5b14806121c25750600560078111156121ad576121ac617845565b5b8460078111156121c0576121bf617845565b5b145b806121f15750600460078111156121dc576121db617845565b5b8460078111156121ef576121ee617845565b5b145b8061221f575060078081111561220a57612209617845565b5b84600781111561221d5761221c617845565b5b145b8061224e57506006600781111561223957612238617845565b5b84600781111561224c5761224b617845565b5b145b61228d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161228490618cff565b60405180910390fd5b600360078111156122a1576122a0617845565b5b6038600087815260200190815260200160002060010160009054906101000a900460ff1660078111156122d7576122d6617845565b5b0361265f5761265e6038600087815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff16600781111561232d5761232c617845565b5b600781111561233f5761233e617845565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546124d79061861a565b80601f01602080910402602001604051908101604052809291908181526020018280546125039061861a565b80156125505780601f1061252557610100808354040283529160200191612550565b820191906000526020600020905b81548152906001019060200180831161253357829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561257e5761257d617845565b5b60048111156125905761258f617845565b5b8152505081526020016006820160009054906101000a900460ff161515151581526020016007820160405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250508152602001600a82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600e8201548152602001600f820160009054906101000a900460ff1615151515815250503385615858565b5b612669858561453d565b612671615b0d565b5050505050565b6000603a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054603b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546127049190618384565b9050919050565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401612767919061807f565b602060405180830381865afa158015612784573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127a891906180af565b6127e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127de90618139565b60405180910390fd5b60006128326001603c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615b17565b905061283e8133612cb7565b5050565b603a6020528060005260406000206000915090505481565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b81526004016128b6919061807f565b602060405180830381865afa1580156128d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128f791906180af565b612936576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161292d90618139565b60405180910390fd5b600061297f603e60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615fa9565b90507f290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d6038600083815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633836040516129ea93929190618d1f565b60405180910390a15050565b806038600082815260200190815260200160002060060160009054906101000a900460ff16612a5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a51906181a5565b60405180910390fd5b81612a876038600083815260200190815260200160002060010160009054906101000a900460ff166156ec565b612ac6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612abd9061842d565b60405180910390fd5b6038600084815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612b6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b6190618dc8565b60405180910390fd5b600160386000858152602001908152602001600020600f0160006101000a81548160ff0219169083151502179055504260386000858152602001908152602001600020600a016002018190555060036007811115612bcb57612bca617845565b5b6038600085815260200190815260200160002060010160009054906101000a900460ff166007811115612c0157612c00617845565b5b03612c8957612c84603e60006038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846163f0565b612c95565b612c9483600461453d565b5b505050565b600f81565b603b6020528060005260406000206000915090505481565b80603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401612d13919061807f565b602060405180830381865afa158015612d30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d5491906180af565b612d93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d8a90618139565b60405180910390fd5b826038600082815260200190815260200160002060060160009054906101000a900460ff16612df7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dee906181a5565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480612e5c57503073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b612e9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e9290618e34565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166038600086815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612f3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f3690618211565b60405180910390fd5b612f4a84600261453d565b50505050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b816038600082815260200190815260200160002060060160009054906101000a900460ff1661301e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613015906181a5565b60405180910390fd5b8261304b6038600083815260200190815260200160002060010160009054906101000a900460ff166156ec565b61308a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130819061842d565b60405180910390fd5b6038600085815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461312e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161312590618499565b60405180910390fd5b826038600086815260200190815260200160002060070160020160006101000a81548160ff02191690831515021790555050505050565b806038600082815260200190815260200160002060060160009054906101000a900460ff166131c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131c0906181a5565b60405180910390fd5b816131f66038600083815260200190815260200160002060010160009054906101000a900460ff166156ec565b613235576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161322c9061842d565b60405180910390fd5b6038600084815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061330657506038600084815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b8061333757506133367f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e33612f50565b5b613376576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161336d90618ea0565b60405180910390fd5b600115156038600085815260200190815260200160002060070160020160009054906101000a900460ff161515146133e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133da90618f0c565b60405180910390fd5b6000603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355f21eb76038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401613476919061807f565b6101a060405180830381865afa158015613494573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134b891906190fe565b602001519050601e67ffffffffffffffff1661384b6038600087815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff16600781111561351b5761351a617845565b5b600781111561352d5761352c617845565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546136c59061861a565b80601f01602080910402602001604051908101604052809291908181526020018280546136f19061861a565b801561373e5780601f106137135761010080835404028352916020019161373e565b820191906000526020600020905b81548152906001019060200180831161372157829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561376c5761376b617845565b5b600481111561377e5761377d617845565b5b8152505081526020016006820160009054906101000a900460ff161515151581526020016007820160405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250508152602001600a82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600e8201548152602001600f820160009054906101000a900460ff1615151515815250508361646c565b67ffffffffffffffff1610613895576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161388c90619178565b60405180910390fd5b6000613a58603860008781526020019081526020016000206003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546139939061861a565b80601f01602080910402602001604051908101604052809291908181526020018280546139bf9061861a565b8015613a0c5780601f106139e157610100808354040283529160200191613a0c565b820191906000526020600020905b8154815290600101906020018083116139ef57829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115613a3a57613a39617845565b5b6004811115613a4c57613a4b617845565b5b8152505083601e6164cc565b9050603a60006038600088815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054603b60006038600089815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054613b509190618384565b811115613b92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b8990618505565b60405180910390fd5b80603a60006038600089815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254613c179190618525565b9250508190555080603860008781526020019081526020016000206007016000016000828254613c479190618525565b925050819055505050505050565b603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000801b81565b603f8181548110613c9257600080fd5b906000526020600020016000915090505481565b80603b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254613cf59190618525565b92505081905550603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401613d5b93929190619198565b6020604051808303816000875af1158015613d7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d9e91906180af565b5050565b600f81565b80603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401613e03919061807f565b602060405180830381865afa158015613e20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e4491906180af565b613e83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e7a90618139565b60405180910390fd5b826038600082815260200190815260200160002060060160009054906101000a900460ff16613ee7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ede906181a5565b60405180910390fd5b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e613f11816155f8565b836038600087815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613f7185600161453d565b5050505050565b613f8182610da7565b613f8a816155f8565b613f948383615707565b505050565b603460009054906101000a900467ffffffffffffffff1681565b6000801b613fc0816155f8565b8160376000828254613fd29190618384565b92505081905550603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016140369291906183b8565b6020604051808303816000875af1158015614055573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061407991906180af565b505050565b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e81565b603960205281600052604060002081815481106140be57600080fd5b90600052602060002001600091509150505481565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008061414861720c565b60005b614154856165ba565b81101561452f5761416585826165ee565b92506038600084815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff1660078111156141b5576141b4617845565b5b60078111156141c7576141c6617845565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461435f9061861a565b80601f016020809104026020016040519081016040528092919081815260200182805461438b9061861a565b80156143d85780601f106143ad576101008083540402835291602001916143d8565b820191906000526020600020905b8154815290600101906020018083116143bb57829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561440657614405617845565b5b600481111561441857614417617845565b5b8152505081526020016006820160009054906101000a900460ff161515151581526020016007820160405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250508152602001600a82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600e8201548152602001600f820160009054906101000a900460ff16151515158152505091508560078111156144f4576144f3617845565b5b8260200151600781111561450b5761450a617845565b5b0361451c5760019350505050614537565b808061452790618723565b91505061414b565b506000925050505b92915050565b816038600082815260200190815260200160002060060160009054906101000a900460ff166145a1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614598906181a5565b60405180910390fd5b6145ce6038600085815260200190815260200160002060010160009054906101000a900460ff1683616698565b816038600085815260200190815260200160002060010160006101000a81548160ff0219169083600781111561460757614606617845565b5b02179055504360386000858152602001908152602001600020600a01600301819055506000600781111561463e5761463d617845565b5b82600781111561465157614650617845565b5b03614834576038600084815260200190815260200160002060070160000154603a60006038600087815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546146f49190618384565b925050819055507f5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0836038600086815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516147629291906189df565b60405180910390a17f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3836038600086815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516147d19291906189df565b60405180910390a160006038600085815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506151eb565b6001600781111561484857614847617845565b5b82600781111561485b5761485a617845565b5b0361490c57603f839080600181540180825580915050600190039060005260206000200160009091909190915055614907603c60006038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846163f0565b6151ea565b600260078111156149205761491f617845565b5b82600781111561493357614932617845565b5b03614d2f576149b6603d60006038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846163f0565b6000603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355f21eb76038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401614a49919061807f565b6101a060405180830381865afa158015614a67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a8b91906190fe565b90507fc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c396038600086815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686614cfa603860008a81526020019081526020016000206003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054614c199061861a565b80601f0160208091040260200160405190810160405280929190818152602001828054614c459061861a565b8015614c925780601f10614c6757610100808354040283529160200191614c92565b820191906000526020600020905b815481529060010190602001808311614c7557829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115614cc057614cbf617845565b5b6004811115614cd257614cd1617845565b5b81525050603860008b8152602001908152602001600020600701600001548760200151616b4f565b603860008a8152602001908152602001600020600301604051614d219594939291906193d6565b60405180910390a1506151e9565b60036007811115614d4357614d42617845565b5b826007811115614d5657614d55617845565b5b03614e41574260386000858152602001908152602001600020600a0160000181905550603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f6edb5f6038600086815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401614e0a919061807f565b600060405180830381600087803b158015614e2457600080fd5b505af1158015614e38573d6000803e3d6000fd5b505050506151e8565b60046007811115614e5557614e54617845565b5b826007811115614e6857614e67617845565b5b1480614e98575060056007811115614e8357614e82617845565b5b826007811115614e9657614e95617845565b5b145b80614ec7575060066007811115614eb257614eb1617845565b5b826007811115614ec557614ec4617845565b5b145b80614ef55750600780811115614ee057614edf617845565b5b826007811115614ef357614ef2617845565b5b145b156151e7576038600084815260200190815260200160002060070160010154603b60006038600087815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015614fcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614fc29061947c565b60405180910390fd5b6038600084815260200190815260200160002060070160000154603a60006038600087815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561509c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615093906194e8565b60405180910390fd5b6038600084815260200190815260200160002060070160010154603b60006038600087815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461513a9190618384565b925050819055506038600084815260200190815260200160002060070160000154603a60006038600087815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546151df9190618384565b925050819055505b5b5b5b5b6151f482616c7b565b15615218574260386000858152602001908152602001600020600a01600101819055505b505050565b60008061522861720c565b60005b615234856165ba565b8110156155eb5761524585826165ee565b92506038600084815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff16600781111561529557615294617845565b5b60078111156152a7576152a6617845565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461543f9061861a565b80601f016020809104026020016040519081016040528092919081815260200182805461546b9061861a565b80156154b85780601f1061548d576101008083540402835291602001916154b8565b820191906000526020600020905b81548152906001019060200180831161549b57829003601f168201915b505050505081526020016002820160009054906101000a900460ff1660048111156154e6576154e5617845565b5b60048111156154f8576154f7617845565b5b8152505081526020016006820160009054906101000a900460ff161515151581526020016007820160405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250508152602001600a82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600e8201548152602001600f820160009054906101000a900460ff1615151515815250509150816101200151156155d857600193505050506155f3565b80806155e390618723565b91505061522b565b506000925050505b919050565b615609816156046156ff565b616d38565b50565b6156168282612f50565b6156e857600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061568d6156ff565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b60006156f782616c7b565b159050919050565b600033905090565b6157118282612f50565b156157e457600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506157896156ff565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600280540361584f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161584690619554565b60405180910390fd5b60028081905550565b6000603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355f21eb7846040518263ffffffff1660e01b81526004016158b5919061807f565b6101a060405180830381865afa1580156158d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906158f791906190fe565b9050600061590e85608001518360200151856164cc565b9050600f603c8660e0015160000151426159289190618384565b61593291906195a3565b61593c9190618525565b8367ffffffffffffffff161115615988576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161597f90619646565b60405180910390fd5b8460c00151600001518111156159a4578460c001516000015190505b84610120015115615a01576159fe816159f9876080015185602001516005603c8b60e00151600001518c60e00151604001516159e09190618384565b6159ea91906195a3565b6159f49190618525565b6164cc565b616dbd565b90505b8060386000876000015181526020019081526020016000206007016001018190555060006064826014615a349190619666565b615a3e91906195a3565b90508060376000828254615a529190618525565b92505081905550603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb868385615aa49190618384565b6040518363ffffffff1660e01b8152600401615ac19291906183b8565b6020604051808303816000875af1158015615ae0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615b0491906180af565b50505050505050565b6001600281905550565b6000615b2282616dd6565b15615b62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b59906196f4565b60405180910390fd5b6000615b6c61720c565b5b615b7684616e0b565b91506038600083815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff166007811115615bc657615bc5617845565b5b6007811115615bd857615bd7617845565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054615d709061861a565b80601f0160208091040260200160405190810160405280929190818152602001828054615d9c9061861a565b8015615de95780601f10615dbe57610100808354040283529160200191615de9565b820191906000526020600020905b815481529060010190602001808311615dcc57829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115615e1757615e16617845565b5b6004811115615e2957615e28617845565b5b8152505081526020016006820160009054906101000a900460ff161515151581526020016007820160405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250508152602001600a82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600e8201548152602001600f820160009054906101000a900460ff1615151515815250509050615efb84616dd6565b158015615f305750846007811115615f1657615f15617845565b5b81602001516007811115615f2d57615f2c617845565b5b14155b615b6d57846007811115615f4757615f46617845565b5b81602001516007811115615f5e57615f5d617845565b5b14615f9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615f95906196f4565b60405180910390fd5b819250505092915050565b6000615fb482616dd6565b15615ff4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615feb906196f4565b60405180910390fd5b6000615ffe61720c565b5b61600884616e0b565b91506038600083815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff16600781111561605857616057617845565b5b600781111561606a57616069617845565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546162029061861a565b80601f016020809104026020016040519081016040528092919081815260200182805461622e9061861a565b801561627b5780601f106162505761010080835404028352916020019161627b565b820191906000526020600020905b81548152906001019060200180831161625e57829003601f168201915b505050505081526020016002820160009054906101000a900460ff1660048111156162a9576162a8617845565b5b60048111156162bb576162ba617845565b5b8152505081526020016006820160009054906101000a900460ff161515151581526020016007820160405180606001604052908160008201548152602001600182015481526020016002820160009054906101000a900460ff1615151515815250508152602001600a82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600e8201548152602001600f820160009054906101000a900460ff161515151581525050905061638d84616dd6565b15801561639d5750806101200151155b615fff578061012001516163e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016163dd906196f4565b60405180910390fd5b8192505050919050565b60008260000160109054906101000a9004600f0b90508183600101600083600f0b600f0b815260200190815260200160002081905550600181018360000160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff160217905550505050565b60008061648684608001518560c001516000015185616b4f565b90506000603c8560e0015160000151426164a09190618384565b6164aa91906195a3565b9050808267ffffffffffffffff166164c29190618384565b9250505092915050565b600064e8d4a51000846040015167ffffffffffffffff16846080015167ffffffffffffffff166164fc9190619666565b856040015167ffffffffffffffff16866020015167ffffffffffffffff168660c0015167ffffffffffffffff166165339190619666565b61653d9190619666565b866000015167ffffffffffffffff16866040015167ffffffffffffffff166165659190619666565b61656f9190618525565b6165799190618525565b856060015167ffffffffffffffff168467ffffffffffffffff1661659d9190619666565b6165a79190619666565b6165b19190619666565b90509392505050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b039050919050565b6000806166216165fd84616ee7565b8560000160009054906101000a9004600f0b600f0b61661c919061971e565b616f54565b90508360000160109054906101000a9004600f0b600f0b81600f0b12616673576040517fb4120f1400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600101600082600f0b600f0b81526020019081526020016000205491505092915050565b8060078111156166ab576166aa617845565b5b8260078111156166be576166bd617845565b5b036166fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016166f5906197ae565b60405180910390fd5b616707826156ec565b616746576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161673d9061842d565b60405180910390fd5b6000600781111561675a57616759617845565b5b82600781111561676d5761676c617845565b5b0361680d576001600781111561678657616785617845565b5b81600781111561679957616798617845565b5b14806167c95750600460078111156167b4576167b3617845565b5b8160078111156167c7576167c6617845565b5b145b616808576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016167ff90619840565b60405180910390fd5b616b4b565b6001600781111561682157616820617845565b5b82600781111561683457616833617845565b5b03616903576000600781111561684d5761684c617845565b5b8160078111156168605761685f617845565b5b148061689057506002600781111561687b5761687a617845565b5b81600781111561688e5761688d617845565b5b145b806168bf5750600460078111156168aa576168a9617845565b5b8160078111156168bd576168bc617845565b5b145b6168fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016168f5906198f8565b60405180910390fd5b616b4a565b6002600781111561691757616916617845565b5b82600781111561692a57616929617845565b5b03616a28576000600781111561694357616942617845565b5b81600781111561695657616955617845565b5b148061698657506003600781111561697157616970617845565b5b81600781111561698457616983617845565b5b145b806169b55750600460078111156169a05761699f617845565b5b8160078111156169b3576169b2617845565b5b145b806169e45750600660078111156169cf576169ce617845565b5b8160078111156169e2576169e1617845565b5b145b616a23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616a1a906199b0565b60405180910390fd5b616b49565b60036007811115616a3c57616a3b617845565b5b826007811115616a4f57616a4e617845565b5b03616b485760056007811115616a6857616a67617845565b5b816007811115616a7b57616a7a617845565b5b1480616aab575060066007811115616a9657616a95617845565b5b816007811115616aa957616aa8617845565b5b145b80616ada575060046007811115616ac557616ac4617845565b5b816007811115616ad857616ad7617845565b5b145b80616b085750600780811115616af357616af2617845565b5b816007811115616b0657616b05617845565b5b145b616b47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616b3e90619a68565b60405180910390fd5b5b5b5b5b5050565b60008064e8d4a51000856040015167ffffffffffffffff16846080015167ffffffffffffffff16616b809190619666565b866040015167ffffffffffffffff16876020015167ffffffffffffffff168660c0015167ffffffffffffffff16616bb79190619666565b616bc19190619666565b876000015167ffffffffffffffff16866040015167ffffffffffffffff16616be99190619666565b616bf39190618525565b616bfd9190618525565b866060015167ffffffffffffffff16616c169190619666565b616c209190619666565b905060008103616c65576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616c5c90619ad4565b60405180910390fd5b8084616c7191906195a3565b9150509392505050565b600060046007811115616c9157616c90617845565b5b826007811115616ca457616ca3617845565b5b1480616cd4575060056007811115616cbf57616cbe617845565b5b826007811115616cd257616cd1617845565b5b145b80616d025750600780811115616ced57616cec617845565b5b826007811115616d0057616cff617845565b5b145b80616d31575060066007811115616d1c57616d1b617845565b5b826007811115616d2f57616d2e617845565b5b145b9050919050565b616d428282612f50565b616db957616d4f81616fa3565b616d5d8360001c6020616fd0565b604051602001616d6e929190619bc8565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616db09190619c3b565b60405180910390fd5b5050565b6000818310616dcc5781616dce565b825b905092915050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b13159050919050565b6000616e1682616dd6565b15616e4d576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260000160009054906101000a9004600f0b905082600101600082600f0b600f0b815260200190815260200160002054915082600101600082600f0b600f0b815260200190815260200160002060009055600181018360000160006101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff16021790555050919050565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115616f4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616f4390619ccf565b60405180910390fd5b819050919050565b60008190508181600f0b14616f9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616f9590619d61565b60405180910390fd5b919050565b6060616fc98273ffffffffffffffffffffffffffffffffffffffff16601460ff16616fd0565b9050919050565b606060006002836002616fe39190619666565b616fed9190618525565b67ffffffffffffffff81111561700657617005617b99565b5b6040519080825280601f01601f1916602001820160405280156170385781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106170705761706f619d81565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106170d4576170d3619d81565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026171149190619666565b61711e9190618525565b90505b60018111156171be577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106171605761715f619d81565b5b1a60f81b82828151811061717757617176619d81565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806171b790619db0565b9050617121565b5060008414617202576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016171f990619e25565b60405180910390fd5b8091505092915050565b604051806101400160405280600080191681526020016000600781111561723657617235617845565b5b8152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200161727d6172b9565b8152602001600015158152602001617293617329565b81526020016172a061734c565b8152602001600080191681526020016000151581525090565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016000600481111561732357617322617845565b5b81525090565b604051806060016040528060008152602001600081526020016000151581525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6173bd81617388565b81146173c857600080fd5b50565b6000813590506173da816173b4565b92915050565b6000602082840312156173f6576173f561737e565b5b6000617404848285016173cb565b91505092915050565b60008115159050919050565b6174228161740d565b82525050565b600060208201905061743d6000830184617419565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061746e82617443565b9050919050565b61747e81617463565b811461748957600080fd5b50565b60008135905061749b81617475565b92915050565b6000602082840312156174b7576174b661737e565b5b60006174c58482850161748c565b91505092915050565b6000819050919050565b6174e1816174ce565b82525050565b60006020820190506174fc60008301846174d8565b92915050565b600081600f0b9050919050565b61751881617502565b82525050565b6000604082019050617533600083018561750f565b617540602083018461750f565b9392505050565b6000819050919050565b600061756c61756761756284617443565b617547565b617443565b9050919050565b600061757e82617551565b9050919050565b600061759082617573565b9050919050565b6175a081617585565b82525050565b60006020820190506175bb6000830184617597565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b617600816175ed565b82525050565b600061761283836175f7565b60208301905092915050565b6000602082019050919050565b6000617636826175c1565b61764081856175cc565b935061764b836175dd565b8060005b8381101561767c5781516176638882617606565b975061766e8361761e565b92505060018101905061764f565b5085935050505092915050565b600060208201905081810360008301526176a3818461762b565b905092915050565b600067ffffffffffffffff82169050919050565b6176c8816176ab565b82525050565b60006020820190506176e360008301846176bf565b92915050565b6176f2816175ed565b81146176fd57600080fd5b50565b60008135905061770f816176e9565b92915050565b60006020828403121561772b5761772a61737e565b5b600061773984828501617700565b91505092915050565b61774b816175ed565b82525050565b60006020820190506177666000830184617742565b92915050565b617775816174ce565b811461778057600080fd5b50565b6000813590506177928161776c565b92915050565b6000602082840312156177ae576177ad61737e565b5b60006177bc84828501617783565b91505092915050565b600080604083850312156177dc576177db61737e565b5b60006177ea85828601617700565b92505060206177fb8582860161748c565b9150509250929050565b6000806040838503121561781c5761781b61737e565b5b600061782a85828601617700565b925050602061783b85828601617783565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6008811061788557617884617845565b5b50565b600081905061789682617874565b919050565b60006178a682617888565b9050919050565b6178b68161789b565b82525050565b6178c581617463565b82525050565b6178d4816176ab565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156179145780820151818401526020810190506178f9565b60008484015250505050565b6000601f19601f8301169050919050565b600061793c826178da565b61794681856178e5565b93506179568185602086016178f6565b61795f81617920565b840191505092915050565b6005811061797b5761797a617845565b5b50565b600081905061798c8261796a565b919050565b600061799c8261797e565b9050919050565b6179ac81617991565b82525050565b600060c0830160008301516179ca60008601826178cb565b5060208301516179dd60208601826178cb565b5060408301516179f060408601826178cb565b506060830151617a0360608601826178cb565b5060808301518482036080860152617a1b8282617931565b91505060a0830151617a3060a08601826179a3565b508091505092915050565b617a44816174ce565b82525050565b617a538161740d565b82525050565b606082016000820151617a6f6000850182617a3b565b506020820151617a826020850182617a3b565b506040820151617a956040850182617a4a565b50505050565b608082016000820151617ab16000850182617a3b565b506020820151617ac46020850182617a3b565b506040820151617ad76040850182617a3b565b506060820151617aea6060850182617a3b565b50505050565b60006101e082019050617b06600083018d617742565b617b13602083018c6178ad565b617b20604083018b6178bc565b617b2d606083018a6178bc565b8181036080830152617b3f81896179b2565b9050617b4e60a0830188617419565b617b5b60c0830187617a59565b617b69610120830186617a9b565b617b776101a0830185617742565b617b856101c0830184617419565b9b9a5050505050505050505050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b617bd182617920565b810181811067ffffffffffffffff82111715617bf057617bef617b99565b5b80604052505050565b6000617c03617374565b9050617c0f8282617bc8565b919050565b600080fd5b617c22816176ab565b8114617c2d57600080fd5b50565b600081359050617c3f81617c19565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff821115617c6a57617c69617b99565b5b617c7382617920565b9050602081019050919050565b82818337600083830152505050565b6000617ca2617c9d84617c4f565b617bf9565b905082815260208101848484011115617cbe57617cbd617c4a565b5b617cc9848285617c80565b509392505050565b600082601f830112617ce657617ce5617c45565b5b8135617cf6848260208601617c8f565b91505092915050565b60058110617d0c57600080fd5b50565b600081359050617d1e81617cff565b92915050565b600060c08284031215617d3a57617d39617b94565b5b617d4460c0617bf9565b90506000617d5484828501617c30565b6000830152506020617d6884828501617c30565b6020830152506040617d7c84828501617c30565b6040830152506060617d9084828501617c30565b606083015250608082013567ffffffffffffffff811115617db457617db3617c14565b5b617dc084828501617cd1565b60808301525060a0617dd484828501617d0f565b60a08301525092915050565b617de98161740d565b8114617df457600080fd5b50565b600081359050617e0681617de0565b92915050565b60008060008060808587031215617e2657617e2561737e565b5b600085013567ffffffffffffffff811115617e4457617e43617383565b5b617e5087828801617d24565b9450506020617e6187828801617783565b9350506040617e7287828801617700565b9250506060617e8387828801617df7565b91505092959194509250565b6000617e9a82617463565b9050919050565b617eaa81617e8f565b8114617eb557600080fd5b50565b600081359050617ec781617ea1565b92915050565b6000617ed882617463565b9050919050565b617ee881617ecd565b8114617ef357600080fd5b50565b600081359050617f0581617edf565b92915050565b60008060408385031215617f2257617f2161737e565b5b6000617f3085828601617eb8565b9250506020617f4185828601617ef6565b9150509250929050565b60088110617f5857600080fd5b50565b600081359050617f6a81617f4b565b92915050565b600080600060608486031215617f8957617f8861737e565b5b6000617f9786828701617700565b9350506020617fa886828701617f5b565b9250506040617fb986828701617c30565b9150509250925092565b60008060408385031215617fda57617fd961737e565b5b6000617fe885828601617700565b9250506020617ff985828601617df7565b9150509250929050565b600061800e82617573565b9050919050565b61801e81618003565b82525050565b60006020820190506180396000830184618015565b92915050565b600080604083850312156180565761805561737e565b5b60006180648582860161748c565b925050602061807585828601617783565b9150509250929050565b600060208201905061809460008301846178bc565b92915050565b6000815190506180a981617de0565b92915050565b6000602082840312156180c5576180c461737e565b5b60006180d38482850161809a565b91505092915050565b600082825260208201905092915050565b7f50726f7669646572206e6f74206a6f696e656421000000000000000000000000600082015250565b60006181236014836180dc565b915061812e826180ed565b602082019050919050565b6000602082019050818103600083015261815281618116565b9050919050565b7f4a6f62206e6f742076616c696421000000000000000000000000000000000000600082015250565b600061818f600e836180dc565b915061819a82618159565b602082019050919050565b600060208201905081810360008301526181be81618182565b9050919050565b7f4a42380000000000000000000000000000000000000000000000000000000000600082015250565b60006181fb6003836180dc565b9150618206826181c5565b602082019050919050565b6000602082019050818103600083015261822a816181ee565b9050919050565b7f43616e277420726566757365206966206e6f7420696e204d4554415f5343484560008201527f44554c4544206f72205343484544554c45442073746174650000000000000000602082015250565b600061828d6038836180dc565b915061829882618231565b604082019050919050565b600060208201905081810360008301526182bc81618280565b9050919050565b7f43616e6e6f74207769746864726177206d6f7265207468616e20756e6c6f636b60008201527f65642062616c616e636500000000000000000000000000000000000000000000602082015250565b600061831f602a836180dc565b915061832a826182c3565b604082019050919050565b6000602082019050818103600083015261834e81618312565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061838f826174ce565b915061839a836174ce565b92508282039050818111156183b2576183b1618355565b5b92915050565b60006040820190506183cd60008301856178bc565b6183da60208301846174d8565b9392505050565b7f4a6f6220737461747573206d75737420626520686f7400000000000000000000600082015250565b60006184176016836180dc565b9150618422826183e1565b602082019050919050565b600060208201905081810360008301526184468161840a565b9050919050565b7f5045524d31000000000000000000000000000000000000000000000000000000600082015250565b60006184836005836180dc565b915061848e8261844d565b602082019050919050565b600060208201905081810360008301526184b281618476565b9050919050565b7f42414c3200000000000000000000000000000000000000000000000000000000600082015250565b60006184ef6004836180dc565b91506184fa826184b9565b602082019050919050565b6000602082019050818103600083015261851e816184e2565b9050919050565b6000618530826174ce565b915061853b836174ce565b925082820190508082111561855357618552618355565b5b92915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b60006185b5602f836180dc565b91506185c082618559565b604082019050919050565b600060208201905081810360008301526185e4816185a8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061863257607f821691505b602082108103618645576186446185eb565b5b50919050565b7f42414c3100000000000000000000000000000000000000000000000000000000600082015250565b60006186816004836180dc565b915061868c8261864b565b602082019050919050565b600060208201905081810360008301526186b081618674565b9050919050565b7f4a44454600000000000000000000000000000000000000000000000000000000600082015250565b60006186ed6004836180dc565b91506186f8826186b7565b602082019050919050565b6000602082019050818103600083015261871c816186e0565b9050919050565b600061872e826174ce565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036187605761875f618355565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026187cd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82618790565b6187d78683618790565b95508019841693508086168417925050509392505050565b600061880a618805618800846174ce565b617547565b6174ce565b9050919050565b6000819050919050565b618824836187ef565b61883861883082618811565b84845461879d565b825550505050565b600090565b61884d618840565b61885881848461881b565b505050565b5b8181101561887c57618871600082618845565b60018101905061885e565b5050565b601f8211156188c1576188928161876b565b61889b84618780565b810160208510156188aa578190505b6188be6188b685618780565b83018261885d565b50505b505050565b600082821c905092915050565b60006188e4600019846008026188c6565b1980831691505092915050565b60006188fd83836188d3565b9150826002028217905092915050565b618916826178da565b67ffffffffffffffff81111561892f5761892e617b99565b5b618939825461861a565b618944828285618880565b600060209050601f8311600181146189775760008415618965578287015190505b61896f85826188f1565b8655506189d7565b601f1984166189858661876b565b60005b828110156189ad57848901518255600182019150602085019450602081019050618988565b868310156189ca57848901516189c6601f8916826188d3565b8355505b6001600288020188555050505b505050505050565b60006040820190506189f46000830185617742565b618a0160208301846178bc565b9392505050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000618a64602e836180dc565b9150618a6f82618a08565b604082019050919050565b60006020820190508181036000830152618a9381618a57565b9050919050565b7f4d6574615363686564756c65723a204372656469742061646472206973207a6560008201527f726f000000000000000000000000000000000000000000000000000000000000602082015250565b6000618af66022836180dc565b9150618b0182618a9a565b604082019050919050565b60006020820190508181036000830152618b2581618ae9565b9050919050565b7f4d6574615363686564756c65723a2050726f76696465724d616e61676572206160008201527f646472206973207a65726f000000000000000000000000000000000000000000602082015250565b6000618b88602b836180dc565b9150618b9382618b2c565b604082019050919050565b60006020820190508181036000830152618bb781618b7b565b9050919050565b6000819050919050565b600060ff82169050919050565b6000618bf0618beb618be684618bbe565b617547565b618bc8565b9050919050565b618c0081618bd5565b82525050565b6000602082019050618c1b6000830184618bf7565b92915050565b7f4f6e6c7920746865206a6f622070726f76696465722063616e206368616e676560008201527f2069747320737461747573000000000000000000000000000000000000000000602082015250565b6000618c7d602b836180dc565b9150618c8882618c21565b604082019050919050565b60006020820190508181036000830152618cac81618c70565b9050919050565b7f4a42313300000000000000000000000000000000000000000000000000000000600082015250565b6000618ce96004836180dc565b9150618cf482618cb3565b602082019050919050565b60006020820190508181036000830152618d1881618cdc565b9050919050565b6000606082019050618d3460008301866178bc565b618d4160208301856178bc565b618d4e6040830184617742565b949350505050565b7f4f6e6c79206a6f62206f776e6572732063616e2063616e63656c20746865697260008201527f206a6f6273000000000000000000000000000000000000000000000000000000602082015250565b6000618db26025836180dc565b9150618dbd82618d56565b604082019050919050565b60006020820190508181036000830152618de181618da5565b9050919050565b7f4f6e6c792070726f7669646572206f7220746869732063616e2063616c6c0000600082015250565b6000618e1e601e836180dc565b9150618e2982618de8565b602082019050919050565b60006020820190508181036000830152618e4d81618e11565b9050919050565b7f5045524d32000000000000000000000000000000000000000000000000000000600082015250565b6000618e8a6005836180dc565b9150618e9582618e54565b602082019050919050565b60006020820190508181036000830152618eb981618e7d565b9050919050565b7f4a42323600000000000000000000000000000000000000000000000000000000600082015250565b6000618ef66004836180dc565b9150618f0182618ec0565b602082019050919050565b60006020820190508181036000830152618f2581618ee9565b9050919050565b600081519050618f3b81617475565b92915050565b600081519050618f5081617c19565b92915050565b600060e08284031215618f6c57618f6b617b94565b5b618f7660e0617bf9565b90506000618f8684828501618f41565b6000830152506020618f9a84828501618f41565b6020830152506040618fae84828501618f41565b6040830152506060618fc284828501618f41565b6060830152506080618fd684828501618f41565b60808301525060a0618fea84828501618f41565b60a08301525060c0618ffe84828501618f41565b60c08301525092915050565b6003811061901757600080fd5b50565b6000815190506190298161900a565b92915050565b60008151905061903e8161776c565b92915050565b60006101a0828403121561905b5761905a617b94565b5b61906560e0617bf9565b9050600061907584828501618f2c565b600083015250602061908984828501618f56565b60208301525061010061909e8482850161901a565b6040830152506101206190b38482850161809a565b6060830152506101406190c884828501618f41565b6080830152506101606190dd8482850161902f565b60a0830152506101806190f28482850161902f565b60c08301525092915050565b60006101a082840312156191155761911461737e565b5b600061912384828501619044565b91505092915050565b7f54494d4531000000000000000000000000000000000000000000000000000000600082015250565b60006191626005836180dc565b915061916d8261912c565b602082019050919050565b6000602082019050818103600083015261919181619155565b9050919050565b60006060820190506191ad60008301866178bc565b6191ba60208301856178bc565b6191c760408301846174d8565b949350505050565b60008160001c9050919050565b600067ffffffffffffffff82169050919050565b60006192036191fe836191cf565b6191dc565b9050919050565b60008160401c9050919050565b600061922a6192258361920a565b6191dc565b9050919050565b60008160801c9050919050565b600061925161924c83619231565b6191dc565b9050919050565b60008160c01c9050919050565b600061927861927383619258565b6191dc565b9050919050565b6000815461928c8161861a565b61929681866178e5565b945060018216600081146192b157600181146192c7576192fa565b60ff1983168652811515602002860193506192fa565b6192d08561876b565b60005b838110156192f2578154818901526001820191506020810190506192d3565b808801955050505b50505092915050565b600060ff82169050919050565b600061932361931e836191cf565b619303565b9050919050565b600060c083016000808401549050619341816191f0565b61934e60008701826178cb565b5061935881619217565b61936560208701826178cb565b5061936f8161923e565b61937c60408701826178cb565b5061938681619265565b61939360608701826178cb565b506001840185830360808701526193aa838261927f565b925050600284015490506193bd81619310565b6193ca60a08701826179a3565b50819250505092915050565b600060a0820190506193eb60008301886178bc565b6193f860208301876178bc565b6194056040830186617742565b61941260608301856176bf565b8181036080830152619424818461932a565b90509695505050505050565b7f41524954484d3100000000000000000000000000000000000000000000000000600082015250565b60006194666007836180dc565b915061947182619430565b602082019050919050565b6000602082019050818103600083015261949581619459565b9050919050565b7f41524954484d3200000000000000000000000000000000000000000000000000600082015250565b60006194d26007836180dc565b91506194dd8261949c565b602082019050919050565b60006020820190508181036000830152619501816194c5565b9050919050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b600061953e601f836180dc565b915061954982619508565b602082019050919050565b6000602082019050818103600083015261956d81619531565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006195ae826174ce565b91506195b9836174ce565b9250826195c9576195c8619574565b5b828204905092915050565b7f43616e6e6f742062696c6c206d6f7265207468616e2074686520656c6170736560008201527f642074696d650000000000000000000000000000000000000000000000000000602082015250565b60006196306026836180dc565b915061963b826195d4565b604082019050919050565b6000602082019050818103600083015261965f81619623565b9050919050565b6000619671826174ce565b915061967c836174ce565b925082820261968a816174ce565b915082820484148315176196a1576196a0618355565b5b5092915050565b7f4a42313000000000000000000000000000000000000000000000000000000000600082015250565b60006196de6004836180dc565b91506196e9826196a8565b602082019050919050565b6000602082019050818103600083015261970d816196d1565b9050919050565b6000819050919050565b600061972982619714565b915061973483619714565b92508282019050828112156000831216838212600084121516171561975c5761975b618355565b5b92915050565b7f43616e6e6f74206368616e67652073746174757320746f20697473656c660000600082015250565b6000619798601e836180dc565b91506197a382619762565b602082019050919050565b600060208201905081810360008301526197c78161978b565b9050919050565b7f43616e206368616e67652066726f6d2050454e44494e4720746f204d4554415f60008201527f5343484544554c4544206f722043414e43454c4c4544206f6e6c790000000000602082015250565b600061982a603b836180dc565b9150619835826197ce565b604082019050919050565b600060208201905081810360008301526198598161981d565b9050919050565b7f43616e206368616e67652066726f6d204d4554415f5343484544554c4544207460008201527f6f2050454e44494e472c205343484544554c4544206f722043414e43454c4c4560208201527f44206f6e6c790000000000000000000000000000000000000000000000000000604082015250565b60006198e26046836180dc565b91506198ed82619860565b606082019050919050565b60006020820190508181036000830152619911816198d5565b9050919050565b7f43616e206368616e67652066726f6d205343484544554c454420746f2050454e60008201527f44494e472c2052554e4e494e472c2043414e43454c4c4544206f72204641494c60208201527f4544206f6e6c7900000000000000000000000000000000000000000000000000604082015250565b600061999a6047836180dc565b91506199a582619918565b606082019050919050565b600060208201905081810360008301526199c98161998d565b9050919050565b7f43616e206368616e67652066726f6d2052554e4e494e4720746f2046494e495360008201527f4845442c204641494c4544206f72204f55545f4f465f43524544495453206f6e60208201527f6c79000000000000000000000000000000000000000000000000000000000000604082015250565b6000619a526042836180dc565b9150619a5d826199d0565b606082019050919050565b60006020820190508181036000830152619a8181619a45565b9050919050565b7f43616e6e6f742064697669646520627920300000000000000000000000000000600082015250565b6000619abe6012836180dc565b9150619ac982619a88565b602082019050919050565b60006020820190508181036000830152619aed81619ab1565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000619b35601783619af4565b9150619b4082619aff565b601782019050919050565b6000619b56826178da565b619b608185619af4565b9350619b708185602086016178f6565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000619bb2601183619af4565b9150619bbd82619b7c565b601182019050919050565b6000619bd382619b28565b9150619bdf8285619b4b565b9150619bea82619ba5565b9150619bf68284619b4b565b91508190509392505050565b6000619c0d826178da565b619c1781856180dc565b9350619c278185602086016178f6565b619c3081617920565b840191505092915050565b60006020820190508181036000830152619c558184619c02565b905092915050565b7f53616665436173743a2076616c756520646f65736e27742066697420696e206160008201527f6e20696e74323536000000000000000000000000000000000000000000000000602082015250565b6000619cb96028836180dc565b9150619cc482619c5d565b604082019050919050565b60006020820190508181036000830152619ce881619cac565b9050919050565b7f53616665436173743a2076616c756520646f65736e27742066697420696e203160008201527f3238206269747300000000000000000000000000000000000000000000000000602082015250565b6000619d4b6027836180dc565b9150619d5682619cef565b604082019050919050565b60006020820190508181036000830152619d7a81619d3e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000619dbb826174ce565b915060008203619dce57619dcd618355565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000619e0f6020836180dc565b9150619e1a82619dd9565b602082019050919050565b60006020820190508181036000830152619e3e81619e02565b905091905056fea264697066735822122037ca596a0ccbe49fc586607d8c0e50a908203b3bf8eb7fcb9a7ab4206c0ed8b664736f6c63430008110033",
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

// BILLDURATIONDELTAMINUTE is a free data retrieval call binding the contract method 0x7a9d0df7.
//
// Solidity: function BILL_DURATION_DELTA_MINUTE() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) BILLDURATIONDELTAMINUTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "BILL_DURATION_DELTA_MINUTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BILLDURATIONDELTAMINUTE is a free data retrieval call binding the contract method 0x7a9d0df7.
//
// Solidity: function BILL_DURATION_DELTA_MINUTE() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) BILLDURATIONDELTAMINUTE() (*big.Int, error) {
	return _MetaScheduler.Contract.BILLDURATIONDELTAMINUTE(&_MetaScheduler.CallOpts)
}

// BILLDURATIONDELTAMINUTE is a free data retrieval call binding the contract method 0x7a9d0df7.
//
// Solidity: function BILL_DURATION_DELTA_MINUTE() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) BILLDURATIONDELTAMINUTE() (*big.Int, error) {
	return _MetaScheduler.Contract.BILLDURATIONDELTAMINUTE(&_MetaScheduler.CallOpts)
}

// BILLTIMECONTROLDELTAS is a free data retrieval call binding the contract method 0xc51067db.
//
// Solidity: function BILL_TIME_CONTROL_DELTA_S() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) BILLTIMECONTROLDELTAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "BILL_TIME_CONTROL_DELTA_S")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BILLTIMECONTROLDELTAS is a free data retrieval call binding the contract method 0xc51067db.
//
// Solidity: function BILL_TIME_CONTROL_DELTA_S() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) BILLTIMECONTROLDELTAS() (*big.Int, error) {
	return _MetaScheduler.Contract.BILLTIMECONTROLDELTAS(&_MetaScheduler.CallOpts)
}

// BILLTIMECONTROLDELTAS is a free data retrieval call binding the contract method 0xc51067db.
//
// Solidity: function BILL_TIME_CONTROL_DELTA_S() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) BILLTIMECONTROLDELTAS() (*big.Int, error) {
	return _MetaScheduler.Contract.BILLTIMECONTROLDELTAS(&_MetaScheduler.CallOpts)
}

// CANCELLATIONFEEMINUTE is a free data retrieval call binding the contract method 0x374c6e0e.
//
// Solidity: function CANCELLATION_FEE_MINUTE() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) CANCELLATIONFEEMINUTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "CANCELLATION_FEE_MINUTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CANCELLATIONFEEMINUTE is a free data retrieval call binding the contract method 0x374c6e0e.
//
// Solidity: function CANCELLATION_FEE_MINUTE() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) CANCELLATIONFEEMINUTE() (*big.Int, error) {
	return _MetaScheduler.Contract.CANCELLATIONFEEMINUTE(&_MetaScheduler.CallOpts)
}

// CANCELLATIONFEEMINUTE is a free data retrieval call binding the contract method 0x374c6e0e.
//
// Solidity: function CANCELLATION_FEE_MINUTE() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) CANCELLATIONFEEMINUTE() (*big.Int, error) {
	return _MetaScheduler.Contract.CANCELLATIONFEEMINUTE(&_MetaScheduler.CallOpts)
}

// DEEPSQUARECUT is a free data retrieval call binding the contract method 0x1df44a3c.
//
// Solidity: function DEEPSQUARE_CUT() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) DEEPSQUARECUT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "DEEPSQUARE_CUT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEEPSQUARECUT is a free data retrieval call binding the contract method 0x1df44a3c.
//
// Solidity: function DEEPSQUARE_CUT() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) DEEPSQUARECUT() (*big.Int, error) {
	return _MetaScheduler.Contract.DEEPSQUARECUT(&_MetaScheduler.CallOpts)
}

// DEEPSQUARECUT is a free data retrieval call binding the contract method 0x1df44a3c.
//
// Solidity: function DEEPSQUARE_CUT() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) DEEPSQUARECUT() (*big.Int, error) {
	return _MetaScheduler.Contract.DEEPSQUARECUT(&_MetaScheduler.CallOpts)
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

// MINIMUMAMOUNT is a free data retrieval call binding the contract method 0x257d9bb8.
//
// Solidity: function MINIMUM_AMOUNT() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) MINIMUMAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "MINIMUM_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMAMOUNT is a free data retrieval call binding the contract method 0x257d9bb8.
//
// Solidity: function MINIMUM_AMOUNT() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) MINIMUMAMOUNT() (*big.Int, error) {
	return _MetaScheduler.Contract.MINIMUMAMOUNT(&_MetaScheduler.CallOpts)
}

// MINIMUMAMOUNT is a free data retrieval call binding the contract method 0x257d9bb8.
//
// Solidity: function MINIMUM_AMOUNT() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) MINIMUMAMOUNT() (*big.Int, error) {
	return _MetaScheduler.Contract.MINIMUMAMOUNT(&_MetaScheduler.CallOpts)
}

// TOPUPSLICEDURATIONMIN is a free data retrieval call binding the contract method 0x1a91c3e9.
//
// Solidity: function TOP_UP_SLICE_DURATION_MIN() view returns(uint64)
func (_MetaScheduler *MetaSchedulerCaller) TOPUPSLICEDURATIONMIN(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "TOP_UP_SLICE_DURATION_MIN")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TOPUPSLICEDURATIONMIN is a free data retrieval call binding the contract method 0x1a91c3e9.
//
// Solidity: function TOP_UP_SLICE_DURATION_MIN() view returns(uint64)
func (_MetaScheduler *MetaSchedulerSession) TOPUPSLICEDURATIONMIN() (uint64, error) {
	return _MetaScheduler.Contract.TOPUPSLICEDURATIONMIN(&_MetaScheduler.CallOpts)
}

// TOPUPSLICEDURATIONMIN is a free data retrieval call binding the contract method 0x1a91c3e9.
//
// Solidity: function TOP_UP_SLICE_DURATION_MIN() view returns(uint64)
func (_MetaScheduler *MetaSchedulerCallerSession) TOPUPSLICEDURATIONMIN() (uint64, error) {
	return _MetaScheduler.Contract.TOPUPSLICEDURATIONMIN(&_MetaScheduler.CallOpts)
}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_MetaScheduler *MetaSchedulerCaller) ClaimJobTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "claimJobTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_MetaScheduler *MetaSchedulerSession) ClaimJobTimeout() (uint64, error) {
	return _MetaScheduler.Contract.ClaimJobTimeout(&_MetaScheduler.CallOpts)
}

// ClaimJobTimeout is a free data retrieval call binding the contract method 0xd6aa37a6.
//
// Solidity: function claimJobTimeout() view returns(uint64)
func (_MetaScheduler *MetaSchedulerCallerSession) ClaimJobTimeout() (uint64, error) {
	return _MetaScheduler.Contract.ClaimJobTimeout(&_MetaScheduler.CallOpts)
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

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address walletAddr) view returns(bytes32[])
func (_MetaScheduler *MetaSchedulerCaller) GetJobs(opts *bind.CallOpts, walletAddr common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "getJobs", walletAddr)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address walletAddr) view returns(bytes32[])
func (_MetaScheduler *MetaSchedulerSession) GetJobs(walletAddr common.Address) ([][32]byte, error) {
	return _MetaScheduler.Contract.GetJobs(&_MetaScheduler.CallOpts, walletAddr)
}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address walletAddr) view returns(bytes32[])
func (_MetaScheduler *MetaSchedulerCallerSession) GetJobs(walletAddr common.Address) ([][32]byte, error) {
	return _MetaScheduler.Contract.GetJobs(&_MetaScheduler.CallOpts, walletAddr)
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

// GetUnlockBalance is a free data retrieval call binding the contract method 0x4b845f4c.
//
// Solidity: function getUnlockBalance(address _addr) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) GetUnlockBalance(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "getUnlockBalance", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockBalance is a free data retrieval call binding the contract method 0x4b845f4c.
//
// Solidity: function getUnlockBalance(address _addr) view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) GetUnlockBalance(_addr common.Address) (*big.Int, error) {
	return _MetaScheduler.Contract.GetUnlockBalance(&_MetaScheduler.CallOpts, _addr)
}

// GetUnlockBalance is a free data retrieval call binding the contract method 0x4b845f4c.
//
// Solidity: function getUnlockBalance(address _addr) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) GetUnlockBalance(_addr common.Address) (*big.Int, error) {
	return _MetaScheduler.Contract.GetUnlockBalance(&_MetaScheduler.CallOpts, _addr)
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

// HotJobList is a free data retrieval call binding the contract method 0xaef3276f.
//
// Solidity: function hotJobList(uint256 ) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCaller) HotJobList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "hotJobList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HotJobList is a free data retrieval call binding the contract method 0xaef3276f.
//
// Solidity: function hotJobList(uint256 ) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) HotJobList(arg0 *big.Int) ([32]byte, error) {
	return _MetaScheduler.Contract.HotJobList(&_MetaScheduler.CallOpts, arg0)
}

// HotJobList is a free data retrieval call binding the contract method 0xaef3276f.
//
// Solidity: function hotJobList(uint256 ) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCallerSession) HotJobList(arg0 *big.Int) ([32]byte, error) {
	return _MetaScheduler.Contract.HotJobList(&_MetaScheduler.CallOpts, arg0)
}

// JobIdCounter is a free data retrieval call binding the contract method 0x0f5a11ca.
//
// Solidity: function jobIdCounter() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) JobIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "jobIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JobIdCounter is a free data retrieval call binding the contract method 0x0f5a11ca.
//
// Solidity: function jobIdCounter() view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) JobIdCounter() (*big.Int, error) {
	return _MetaScheduler.Contract.JobIdCounter(&_MetaScheduler.CallOpts)
}

// JobIdCounter is a free data retrieval call binding the contract method 0x0f5a11ca.
//
// Solidity: function jobIdCounter() view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) JobIdCounter() (*big.Int, error) {
	return _MetaScheduler.Contract.JobIdCounter(&_MetaScheduler.CallOpts)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string,uint8) definition, bool valid, (uint256,uint256,bool) cost, (uint256,uint256,uint256,uint256) time, bytes32 jobName, bool hasCancelRequest)
func (_MetaScheduler *MetaSchedulerCaller) Jobs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       JobDefinition
	Valid            bool
	Cost             JobCost
	Time             JobTime
	JobName          [32]byte
	HasCancelRequest bool
}, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		JobId            [32]byte
		Status           uint8
		CustomerAddr     common.Address
		ProviderAddr     common.Address
		Definition       JobDefinition
		Valid            bool
		Cost             JobCost
		Time             JobTime
		JobName          [32]byte
		HasCancelRequest bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.JobId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.CustomerAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ProviderAddr = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Definition = *abi.ConvertType(out[4], new(JobDefinition)).(*JobDefinition)
	outstruct.Valid = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Cost = *abi.ConvertType(out[6], new(JobCost)).(*JobCost)
	outstruct.Time = *abi.ConvertType(out[7], new(JobTime)).(*JobTime)
	outstruct.JobName = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.HasCancelRequest = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string,uint8) definition, bool valid, (uint256,uint256,bool) cost, (uint256,uint256,uint256,uint256) time, bytes32 jobName, bool hasCancelRequest)
func (_MetaScheduler *MetaSchedulerSession) Jobs(arg0 [32]byte) (struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       JobDefinition
	Valid            bool
	Cost             JobCost
	Time             JobTime
	JobName          [32]byte
	HasCancelRequest bool
}, error) {
	return _MetaScheduler.Contract.Jobs(&_MetaScheduler.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x38ed7cfc.
//
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string,uint8) definition, bool valid, (uint256,uint256,bool) cost, (uint256,uint256,uint256,uint256) time, bytes32 jobName, bool hasCancelRequest)
func (_MetaScheduler *MetaSchedulerCallerSession) Jobs(arg0 [32]byte) (struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       JobDefinition
	Valid            bool
	Cost             JobCost
	Time             JobTime
	JobName          [32]byte
	HasCancelRequest bool
}, error) {
	return _MetaScheduler.Contract.Jobs(&_MetaScheduler.CallOpts, arg0)
}

// ProviderCancellingJobsQueues is a free data retrieval call binding the contract method 0x407969ae.
//
// Solidity: function providerCancellingJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCaller) ProviderCancellingJobsQueues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "providerCancellingJobsQueues", arg0)

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

// ProviderCancellingJobsQueues is a free data retrieval call binding the contract method 0x407969ae.
//
// Solidity: function providerCancellingJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerSession) ProviderCancellingJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderCancellingJobsQueues(&_MetaScheduler.CallOpts, arg0)
}

// ProviderCancellingJobsQueues is a free data retrieval call binding the contract method 0x407969ae.
//
// Solidity: function providerCancellingJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCallerSession) ProviderCancellingJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderCancellingJobsQueues(&_MetaScheduler.CallOpts, arg0)
}

// ProviderClaimableJobsQueues is a free data retrieval call binding the contract method 0x110e87a6.
//
// Solidity: function providerClaimableJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCaller) ProviderClaimableJobsQueues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "providerClaimableJobsQueues", arg0)

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
func (_MetaScheduler *MetaSchedulerSession) ProviderClaimableJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderClaimableJobsQueues(&_MetaScheduler.CallOpts, arg0)
}

// ProviderClaimableJobsQueues is a free data retrieval call binding the contract method 0x110e87a6.
//
// Solidity: function providerClaimableJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCallerSession) ProviderClaimableJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderClaimableJobsQueues(&_MetaScheduler.CallOpts, arg0)
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

// ProviderTimeoutJobsQueues is a free data retrieval call binding the contract method 0x31c3b874.
//
// Solidity: function providerTimeoutJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCaller) ProviderTimeoutJobsQueues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "providerTimeoutJobsQueues", arg0)

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

// ProviderTimeoutJobsQueues is a free data retrieval call binding the contract method 0x31c3b874.
//
// Solidity: function providerTimeoutJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerSession) ProviderTimeoutJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderTimeoutJobsQueues(&_MetaScheduler.CallOpts, arg0)
}

// ProviderTimeoutJobsQueues is a free data retrieval call binding the contract method 0x31c3b874.
//
// Solidity: function providerTimeoutJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCallerSession) ProviderTimeoutJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderTimeoutJobsQueues(&_MetaScheduler.CallOpts, arg0)
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

// Wallet2JobId is a free data retrieval call binding the contract method 0xe88fe8ca.
//
// Solidity: function wallet2JobId(address , uint256 ) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCaller) Wallet2JobId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "wallet2JobId", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Wallet2JobId is a free data retrieval call binding the contract method 0xe88fe8ca.
//
// Solidity: function wallet2JobId(address , uint256 ) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) Wallet2JobId(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _MetaScheduler.Contract.Wallet2JobId(&_MetaScheduler.CallOpts, arg0, arg1)
}

// Wallet2JobId is a free data retrieval call binding the contract method 0xe88fe8ca.
//
// Solidity: function wallet2JobId(address , uint256 ) view returns(bytes32)
func (_MetaScheduler *MetaSchedulerCallerSession) Wallet2JobId(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _MetaScheduler.Contract.Wallet2JobId(&_MetaScheduler.CallOpts, arg0, arg1)
}

// Wallet2LockedBalance is a free data retrieval call binding the contract method 0x5d76d0d1.
//
// Solidity: function wallet2LockedBalance(address ) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) Wallet2LockedBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "wallet2LockedBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wallet2LockedBalance is a free data retrieval call binding the contract method 0x5d76d0d1.
//
// Solidity: function wallet2LockedBalance(address ) view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) Wallet2LockedBalance(arg0 common.Address) (*big.Int, error) {
	return _MetaScheduler.Contract.Wallet2LockedBalance(&_MetaScheduler.CallOpts, arg0)
}

// Wallet2LockedBalance is a free data retrieval call binding the contract method 0x5d76d0d1.
//
// Solidity: function wallet2LockedBalance(address ) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) Wallet2LockedBalance(arg0 common.Address) (*big.Int, error) {
	return _MetaScheduler.Contract.Wallet2LockedBalance(&_MetaScheduler.CallOpts, arg0)
}

// Wallet2TotalBalance is a free data retrieval call binding the contract method 0x8964b048.
//
// Solidity: function wallet2TotalBalance(address ) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCaller) Wallet2TotalBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "wallet2TotalBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wallet2TotalBalance is a free data retrieval call binding the contract method 0x8964b048.
//
// Solidity: function wallet2TotalBalance(address ) view returns(uint256)
func (_MetaScheduler *MetaSchedulerSession) Wallet2TotalBalance(arg0 common.Address) (*big.Int, error) {
	return _MetaScheduler.Contract.Wallet2TotalBalance(&_MetaScheduler.CallOpts, arg0)
}

// Wallet2TotalBalance is a free data retrieval call binding the contract method 0x8964b048.
//
// Solidity: function wallet2TotalBalance(address ) view returns(uint256)
func (_MetaScheduler *MetaSchedulerCallerSession) Wallet2TotalBalance(arg0 common.Address) (*big.Int, error) {
	return _MetaScheduler.Contract.Wallet2TotalBalance(&_MetaScheduler.CallOpts, arg0)
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

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ClaimJob(opts *bind.TransactOpts, _jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "claimJob", _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimJob(&_MetaScheduler.TransactOpts, _jobId, _providerAddr)
}

// ClaimJob is a paid mutator transaction binding the contract method 0x8fb70f63.
//
// Solidity: function claimJob(bytes32 _jobId, address _providerAddr) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ClaimJob(_jobId [32]byte, _providerAddr common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ClaimJob(&_MetaScheduler.TransactOpts, _jobId, _providerAddr)
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

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Deposit(&_MetaScheduler.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Deposit(&_MetaScheduler.TransactOpts, _amount)
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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _credit, address _providerManager) returns()
func (_MetaScheduler *MetaSchedulerTransactor) Initialize(opts *bind.TransactOpts, _credit common.Address, _providerManager common.Address) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "initialize", _credit, _providerManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _credit, address _providerManager) returns()
func (_MetaScheduler *MetaSchedulerSession) Initialize(_credit common.Address, _providerManager common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Initialize(&_MetaScheduler.TransactOpts, _credit, _providerManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _credit, address _providerManager) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) Initialize(_credit common.Address, _providerManager common.Address) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Initialize(&_MetaScheduler.TransactOpts, _credit, _providerManager)
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

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _nextJobStatus, uint64 _jobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerTransactor) ProviderSetJobStatus(opts *bind.TransactOpts, _jobId [32]byte, _nextJobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "providerSetJobStatus", _jobId, _nextJobStatus, _jobDurationMinute)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _nextJobStatus, uint64 _jobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerSession) ProviderSetJobStatus(_jobId [32]byte, _nextJobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderSetJobStatus(&_MetaScheduler.TransactOpts, _jobId, _nextJobStatus, _jobDurationMinute)
}

// ProviderSetJobStatus is a paid mutator transaction binding the contract method 0x48841b9c.
//
// Solidity: function providerSetJobStatus(bytes32 _jobId, uint8 _nextJobStatus, uint64 _jobDurationMinute) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) ProviderSetJobStatus(_jobId [32]byte, _nextJobStatus uint8, _jobDurationMinute uint64) (*types.Transaction, error) {
	return _MetaScheduler.Contract.ProviderSetJobStatus(&_MetaScheduler.TransactOpts, _jobId, _nextJobStatus, _jobDurationMinute)
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

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool _autoTopUp) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactor) RequestNewJob(opts *bind.TransactOpts, _definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, _autoTopUp bool) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "requestNewJob", _definition, _maxCost, _jobName, _autoTopUp)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool _autoTopUp) returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) RequestNewJob(_definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, _autoTopUp bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _maxCost, _jobName, _autoTopUp)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool _autoTopUp) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactorSession) RequestNewJob(_definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, _autoTopUp bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _maxCost, _jobName, _autoTopUp)
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

// SetAutoTopUpJob is a paid mutator transaction binding the contract method 0x9b06ecad.
//
// Solidity: function setAutoTopUpJob(bytes32 _jobId, bool _autoTopUp) returns()
func (_MetaScheduler *MetaSchedulerTransactor) SetAutoTopUpJob(opts *bind.TransactOpts, _jobId [32]byte, _autoTopUp bool) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "setAutoTopUpJob", _jobId, _autoTopUp)
}

// SetAutoTopUpJob is a paid mutator transaction binding the contract method 0x9b06ecad.
//
// Solidity: function setAutoTopUpJob(bytes32 _jobId, bool _autoTopUp) returns()
func (_MetaScheduler *MetaSchedulerSession) SetAutoTopUpJob(_jobId [32]byte, _autoTopUp bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.SetAutoTopUpJob(&_MetaScheduler.TransactOpts, _jobId, _autoTopUp)
}

// SetAutoTopUpJob is a paid mutator transaction binding the contract method 0x9b06ecad.
//
// Solidity: function setAutoTopUpJob(bytes32 _jobId, bool _autoTopUp) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) SetAutoTopUpJob(_jobId [32]byte, _autoTopUp bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.SetAutoTopUpJob(&_MetaScheduler.TransactOpts, _jobId, _autoTopUp)
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

// TopUpJobSlice is a paid mutator transaction binding the contract method 0x9cdf8d9e.
//
// Solidity: function topUpJobSlice(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactor) TopUpJobSlice(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "topUpJobSlice", _jobId)
}

// TopUpJobSlice is a paid mutator transaction binding the contract method 0x9cdf8d9e.
//
// Solidity: function topUpJobSlice(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerSession) TopUpJobSlice(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpJobSlice(&_MetaScheduler.TransactOpts, _jobId)
}

// TopUpJobSlice is a paid mutator transaction binding the contract method 0x9cdf8d9e.
//
// Solidity: function topUpJobSlice(bytes32 _jobId) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) TopUpJobSlice(_jobId [32]byte) (*types.Transaction, error) {
	return _MetaScheduler.Contract.TopUpJobSlice(&_MetaScheduler.TransactOpts, _jobId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Withdraw(&_MetaScheduler.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_MetaScheduler *MetaSchedulerTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _MetaScheduler.Contract.Withdraw(&_MetaScheduler.TransactOpts, _amount)
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

// FilterClaimJobEvent is a free log retrieval operation binding the contract event 0xc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c39.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8) jobDefinition)
func (_MetaScheduler *MetaSchedulerFilterer) FilterClaimJobEvent(opts *bind.FilterOpts) (*MetaSchedulerClaimJobEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "ClaimJobEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerClaimJobEventIterator{contract: _MetaScheduler.contract, event: "ClaimJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimJobEvent is a free log subscription operation binding the contract event 0xc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c39.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8) jobDefinition)
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

// ParseClaimJobEvent is a log parse operation binding the contract event 0xc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c39.
//
// Solidity: event ClaimJobEvent(address customerAddr, address providerAddr, bytes32 jobId, uint64 maxDurationMinute, (uint64,uint64,uint64,uint64,string,uint8) jobDefinition)
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
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterJobRefusedEvent is a free log retrieval operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) FilterJobRefusedEvent(opts *bind.FilterOpts) (*MetaSchedulerJobRefusedEventIterator, error) {

	logs, sub, err := _MetaScheduler.contract.FilterLogs(opts, "JobRefusedEvent")
	if err != nil {
		return nil, err
	}
	return &MetaSchedulerJobRefusedEventIterator{contract: _MetaScheduler.contract, event: "JobRefusedEvent", logs: logs, sub: sub}, nil
}

// WatchJobRefusedEvent is a free log subscription operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
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

// ParseJobRefusedEvent is a log parse operation binding the contract event 0x5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0.
//
// Solidity: event JobRefusedEvent(bytes32 _jobId, address _providerAddr)
func (_MetaScheduler *MetaSchedulerFilterer) ParseJobRefusedEvent(log types.Log) (*MetaSchedulerJobRefusedEvent, error) {
	event := new(MetaSchedulerJobRefusedEvent)
	if err := _MetaScheduler.contract.UnpackLog(event, "JobRefusedEvent", log); err != nil {
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

// ReentrancyGuardUpgradeableMetaData contains all meta data concerning the ReentrancyGuardUpgradeable contract.
var ReentrancyGuardUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ReentrancyGuardUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardUpgradeableMetaData.ABI instead.
var ReentrancyGuardUpgradeableABI = ReentrancyGuardUpgradeableMetaData.ABI

// ReentrancyGuardUpgradeable is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeable struct {
	ReentrancyGuardUpgradeableCaller     // Read-only binding to the contract
	ReentrancyGuardUpgradeableTransactor // Write-only binding to the contract
	ReentrancyGuardUpgradeableFilterer   // Log filterer for contract events
}

// ReentrancyGuardUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardUpgradeableSession struct {
	Contract     *ReentrancyGuardUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardUpgradeableCallerSession struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ReentrancyGuardUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardUpgradeableTransactorSession struct {
	Contract     *ReentrancyGuardUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableRaw struct {
	Contract *ReentrancyGuardUpgradeable // Generic contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCallerRaw struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactorRaw struct {
	Contract *ReentrancyGuardUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuardUpgradeable creates a new instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeable(address common.Address, backend bind.ContractBackend) (*ReentrancyGuardUpgradeable, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeable{ReentrancyGuardUpgradeableCaller: ReentrancyGuardUpgradeableCaller{contract: contract}, ReentrancyGuardUpgradeableTransactor: ReentrancyGuardUpgradeableTransactor{contract: contract}, ReentrancyGuardUpgradeableFilterer: ReentrancyGuardUpgradeableFilterer{contract: contract}}, nil
}

// NewReentrancyGuardUpgradeableCaller creates a new read-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardUpgradeableCaller, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableCaller{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableTransactor creates a new write-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardUpgradeableTransactor, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableTransactor{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableFilterer creates a new log filterer instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardUpgradeableFilterer, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableFilterer{contract: contract}, nil
}

// bindReentrancyGuardUpgradeable binds a generic wrapper to an already deployed contract.
func bindReentrancyGuardUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ReentrancyGuardUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ReentrancyGuardUpgradeable contract.
type ReentrancyGuardUpgradeableInitializedIterator struct {
	Event *ReentrancyGuardUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ReentrancyGuardUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReentrancyGuardUpgradeableInitialized)
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
		it.Event = new(ReentrancyGuardUpgradeableInitialized)
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
func (it *ReentrancyGuardUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReentrancyGuardUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReentrancyGuardUpgradeableInitialized represents a Initialized event raised by the ReentrancyGuardUpgradeable contract.
type ReentrancyGuardUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReentrancyGuardUpgradeableInitializedIterator, error) {

	logs, sub, err := _ReentrancyGuardUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableInitializedIterator{contract: _ReentrancyGuardUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReentrancyGuardUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ReentrancyGuardUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReentrancyGuardUpgradeableInitialized)
				if err := _ReentrancyGuardUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) ParseInitialized(log types.Log) (*ReentrancyGuardUpgradeableInitialized, error) {
	event := new(ReentrancyGuardUpgradeableInitialized)
	if err := _ReentrancyGuardUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// TimeoutManagementMetaData contains all meta data concerning the TimeoutManagement contract.
var TimeoutManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockOrigin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"stillAlive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101cd610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806315a945ce1461003a575b600080fd5b610054600480360381019061004f91906100be565b61006a565b6040516100619190610119565b60405180910390f35b60008183436100799190610163565b1015905092915050565b600080fd5b6000819050919050565b61009b81610088565b81146100a657600080fd5b50565b6000813590506100b881610092565b92915050565b600080604083850312156100d5576100d4610083565b5b60006100e3858286016100a9565b92505060206100f4858286016100a9565b9150509250929050565b60008115159050919050565b610113816100fe565b82525050565b600060208201905061012e600083018461010a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061016e82610088565b915061017983610088565b925082820390508181111561019157610190610134565b5b9291505056fea2646970667358221220a3176e05e69fdbabcf79a7a19c316fa4b4a9e288d5602022dad67d0597b0989264736f6c63430008110033",
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
	parsed, err := TimeoutManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ToolsMetaData contains all meta data concerning the Tools contract.
var ToolsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122074f65c426f0f1ee6dff9e92ed12d3385f20b7104afb679f6173a7db79f516ee964736f6c63430008110033",
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
	parsed, err := ConsoleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
