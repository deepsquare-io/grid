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
	MaxCost                   *big.Int
	FinalCost                 *big.Int
	PendingTopUp              *big.Int
	DelegateSpendingAuthority bool
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_billingAmount\",\"type\":\"uint256\"}],\"name\":\"BilledTooMuchEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"ClaimNextCancellingJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"}],\"name\":\"ClaimNextTopUpJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumJobStatus\",\"name\":\"_from\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumJobStatus\",\"name\":\"_to\",\"type\":\"uint8\"}],\"name\":\"JobTransitionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextCancellingJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasCancellingJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasNextJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_jobStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_jobDurationMinute\",\"type\":\"uint64\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"delegateSpendingAuthority\",\"type\":\"bool\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"topUpJobDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool delegateSpendingAuthority) returns(bytes32)
func (_IMetaScheduler *IMetaSchedulerTransactor) RequestNewJob(opts *bind.TransactOpts, _definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "requestNewJob", _definition, _maxCost, _jobName, delegateSpendingAuthority)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool delegateSpendingAuthority) returns(bytes32)
func (_IMetaScheduler *IMetaSchedulerSession) RequestNewJob(_definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.RequestNewJob(&_IMetaScheduler.TransactOpts, _definition, _maxCost, _jobName, delegateSpendingAuthority)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _maxCost, bytes32 _jobName, bool delegateSpendingAuthority) returns(bytes32)
func (_IMetaScheduler *IMetaSchedulerTransactorSession) RequestNewJob(_definition JobDefinition, _maxCost *big.Int, _jobName [32]byte, delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.RequestNewJob(&_IMetaScheduler.TransactOpts, _definition, _maxCost, _jobName, delegateSpendingAuthority)
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

// TopUpJobDelegate is a paid mutator transaction binding the contract method 0x69ee1bf9.
//
// Solidity: function topUpJobDelegate(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactor) TopUpJobDelegate(opts *bind.TransactOpts, _jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.contract.Transact(opts, "topUpJobDelegate", _jobId)
}

// TopUpJobDelegate is a paid mutator transaction binding the contract method 0x69ee1bf9.
//
// Solidity: function topUpJobDelegate(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerSession) TopUpJobDelegate(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.TopUpJobDelegate(&_IMetaScheduler.TransactOpts, _jobId)
}

// TopUpJobDelegate is a paid mutator transaction binding the contract method 0x69ee1bf9.
//
// Solidity: function topUpJobDelegate(bytes32 _jobId) returns()
func (_IMetaScheduler *IMetaSchedulerTransactorSession) TopUpJobDelegate(_jobId [32]byte) (*types.Transaction, error) {
	return _IMetaScheduler.Contract.TopUpJobDelegate(&_IMetaScheduler.TransactOpts, _jobId)
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

// IMetaSchedulerBilledTooMuchEventIterator is returned from FilterBilledTooMuchEvent and is used to iterate over the raw logs and unpacked data for BilledTooMuchEvent events raised by the IMetaScheduler contract.
type IMetaSchedulerBilledTooMuchEventIterator struct {
	Event *IMetaSchedulerBilledTooMuchEvent // Event containing the contract specifics and raw log

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
func (it *IMetaSchedulerBilledTooMuchEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMetaSchedulerBilledTooMuchEvent)
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
		it.Event = new(IMetaSchedulerBilledTooMuchEvent)
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
func (it *IMetaSchedulerBilledTooMuchEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMetaSchedulerBilledTooMuchEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMetaSchedulerBilledTooMuchEvent represents a BilledTooMuchEvent event raised by the IMetaScheduler contract.
type IMetaSchedulerBilledTooMuchEvent struct {
	JobId         [32]byte
	ProviderAddr  common.Address
	BillingAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBilledTooMuchEvent is a free log retrieval operation binding the contract event 0x17e65314b087df225f56701d0a66a3f7d9ce0f26077307b4b765a19c60a36d44.
//
// Solidity: event BilledTooMuchEvent(bytes32 _jobId, address _providerAddr, uint256 _billingAmount)
func (_IMetaScheduler *IMetaSchedulerFilterer) FilterBilledTooMuchEvent(opts *bind.FilterOpts) (*IMetaSchedulerBilledTooMuchEventIterator, error) {

	logs, sub, err := _IMetaScheduler.contract.FilterLogs(opts, "BilledTooMuchEvent")
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerBilledTooMuchEventIterator{contract: _IMetaScheduler.contract, event: "BilledTooMuchEvent", logs: logs, sub: sub}, nil
}

// WatchBilledTooMuchEvent is a free log subscription operation binding the contract event 0x17e65314b087df225f56701d0a66a3f7d9ce0f26077307b4b765a19c60a36d44.
//
// Solidity: event BilledTooMuchEvent(bytes32 _jobId, address _providerAddr, uint256 _billingAmount)
func (_IMetaScheduler *IMetaSchedulerFilterer) WatchBilledTooMuchEvent(opts *bind.WatchOpts, sink chan<- *IMetaSchedulerBilledTooMuchEvent) (event.Subscription, error) {

	logs, sub, err := _IMetaScheduler.contract.WatchLogs(opts, "BilledTooMuchEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMetaSchedulerBilledTooMuchEvent)
				if err := _IMetaScheduler.contract.UnpackLog(event, "BilledTooMuchEvent", log); err != nil {
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
func (_IMetaScheduler *IMetaSchedulerFilterer) ParseBilledTooMuchEvent(log types.Log) (*IMetaSchedulerBilledTooMuchEvent, error) {
	event := new(IMetaSchedulerBilledTooMuchEvent)
	if err := _IMetaScheduler.contract.UnpackLog(event, "BilledTooMuchEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// IMetaSchedulerClaimNextTopUpJobEventIterator is returned from FilterClaimNextTopUpJobEvent and is used to iterate over the raw logs and unpacked data for ClaimNextTopUpJobEvent events raised by the IMetaScheduler contract.
type IMetaSchedulerClaimNextTopUpJobEventIterator struct {
	Event *IMetaSchedulerClaimNextTopUpJobEvent // Event containing the contract specifics and raw log

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
func (it *IMetaSchedulerClaimNextTopUpJobEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMetaSchedulerClaimNextTopUpJobEvent)
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
		it.Event = new(IMetaSchedulerClaimNextTopUpJobEvent)
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
func (it *IMetaSchedulerClaimNextTopUpJobEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMetaSchedulerClaimNextTopUpJobEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMetaSchedulerClaimNextTopUpJobEvent represents a ClaimNextTopUpJobEvent event raised by the IMetaScheduler contract.
type IMetaSchedulerClaimNextTopUpJobEvent struct {
	JobId             [32]byte
	ProviderAddr      common.Address
	MaxDurationMinute uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimNextTopUpJobEvent is a free log retrieval operation binding the contract event 0xa42f2b4a7ee7f91857a4c98fc71fc48546a284d5db48dd77b7ab81030a494470.
//
// Solidity: event ClaimNextTopUpJobEvent(bytes32 _jobId, address _providerAddr, uint64 maxDurationMinute)
func (_IMetaScheduler *IMetaSchedulerFilterer) FilterClaimNextTopUpJobEvent(opts *bind.FilterOpts) (*IMetaSchedulerClaimNextTopUpJobEventIterator, error) {

	logs, sub, err := _IMetaScheduler.contract.FilterLogs(opts, "ClaimNextTopUpJobEvent")
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerClaimNextTopUpJobEventIterator{contract: _IMetaScheduler.contract, event: "ClaimNextTopUpJobEvent", logs: logs, sub: sub}, nil
}

// WatchClaimNextTopUpJobEvent is a free log subscription operation binding the contract event 0xa42f2b4a7ee7f91857a4c98fc71fc48546a284d5db48dd77b7ab81030a494470.
//
// Solidity: event ClaimNextTopUpJobEvent(bytes32 _jobId, address _providerAddr, uint64 maxDurationMinute)
func (_IMetaScheduler *IMetaSchedulerFilterer) WatchClaimNextTopUpJobEvent(opts *bind.WatchOpts, sink chan<- *IMetaSchedulerClaimNextTopUpJobEvent) (event.Subscription, error) {

	logs, sub, err := _IMetaScheduler.contract.WatchLogs(opts, "ClaimNextTopUpJobEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMetaSchedulerClaimNextTopUpJobEvent)
				if err := _IMetaScheduler.contract.UnpackLog(event, "ClaimNextTopUpJobEvent", log); err != nil {
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
func (_IMetaScheduler *IMetaSchedulerFilterer) ParseClaimNextTopUpJobEvent(log types.Log) (*IMetaSchedulerClaimNextTopUpJobEvent, error) {
	event := new(IMetaSchedulerClaimNextTopUpJobEvent)
	if err := _IMetaScheduler.contract.UnpackLog(event, "ClaimNextTopUpJobEvent", log); err != nil {
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

// IMetaSchedulerJobTransitionEventIterator is returned from FilterJobTransitionEvent and is used to iterate over the raw logs and unpacked data for JobTransitionEvent events raised by the IMetaScheduler contract.
type IMetaSchedulerJobTransitionEventIterator struct {
	Event *IMetaSchedulerJobTransitionEvent // Event containing the contract specifics and raw log

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
func (it *IMetaSchedulerJobTransitionEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMetaSchedulerJobTransitionEvent)
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
		it.Event = new(IMetaSchedulerJobTransitionEvent)
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
func (it *IMetaSchedulerJobTransitionEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMetaSchedulerJobTransitionEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMetaSchedulerJobTransitionEvent represents a JobTransitionEvent event raised by the IMetaScheduler contract.
type IMetaSchedulerJobTransitionEvent struct {
	JobId [32]byte
	From  uint8
	To    uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJobTransitionEvent is a free log retrieval operation binding the contract event 0x0bba917f0a1e0fc0d51a75273e7088a4dfecb010699e60ac9c58526429f6c37f.
//
// Solidity: event JobTransitionEvent(bytes32 _jobId, uint8 _from, uint8 _to)
func (_IMetaScheduler *IMetaSchedulerFilterer) FilterJobTransitionEvent(opts *bind.FilterOpts) (*IMetaSchedulerJobTransitionEventIterator, error) {

	logs, sub, err := _IMetaScheduler.contract.FilterLogs(opts, "JobTransitionEvent")
	if err != nil {
		return nil, err
	}
	return &IMetaSchedulerJobTransitionEventIterator{contract: _IMetaScheduler.contract, event: "JobTransitionEvent", logs: logs, sub: sub}, nil
}

// WatchJobTransitionEvent is a free log subscription operation binding the contract event 0x0bba917f0a1e0fc0d51a75273e7088a4dfecb010699e60ac9c58526429f6c37f.
//
// Solidity: event JobTransitionEvent(bytes32 _jobId, uint8 _from, uint8 _to)
func (_IMetaScheduler *IMetaSchedulerFilterer) WatchJobTransitionEvent(opts *bind.WatchOpts, sink chan<- *IMetaSchedulerJobTransitionEvent) (event.Subscription, error) {

	logs, sub, err := _IMetaScheduler.contract.WatchLogs(opts, "JobTransitionEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMetaSchedulerJobTransitionEvent)
				if err := _IMetaScheduler.contract.UnpackLog(event, "JobTransitionEvent", log); err != nil {
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
func (_IMetaScheduler *IMetaSchedulerFilterer) ParseJobTransitionEvent(log types.Log) (*IMetaSchedulerJobTransitionEvent, error) {
	event := new(IMetaSchedulerJobTransitionEvent)
	if err := _IMetaScheduler.contract.UnpackLog(event, "JobTransitionEvent", log); err != nil {
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
	ABI: "[{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_billingAmount\",\"type\":\"uint256\"}],\"name\":\"BilledTooMuchEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structJobDefinition\",\"name\":\"jobDefinition\",\"type\":\"tuple\"}],\"name\":\"ClaimJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"}],\"name\":\"ClaimNextCancellingJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxDurationMinute\",\"type\":\"uint64\"}],\"name\":\"ClaimNextTopUpJobEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"JobRefusedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumJobStatus\",\"name\":\"_from\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumJobStatus\",\"name\":\"_to\",\"type\":\"uint8\"}],\"name\":\"JobTransitionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customerAddr\",\"type\":\"address\"}],\"name\":\"NewJobRequestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BILL_DURATION_DELTA_MINUTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BILL_TIME_CONTROL_DELTA_S\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCELLATION_FEE_MINUTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEEPSQUARE_CUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"METASCHEDULER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOP_UP_SLICE_DURATION_MIN\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"claimJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimJobTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextCancellingJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimNextTopUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasCancellingJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasNextJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"hasTopUpJob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotJobList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_credit\",\"type\":\"address\"},{\"internalType\":\"contractIProviderManager\",\"name\":\"_providerManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"customerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"internalType\":\"structJobDefinition\",\"name\":\"definition\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingTopUp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateSpendingAuthority\",\"type\":\"bool\"}],\"internalType\":\"structJobCost\",\"name\":\"cost\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelRequestTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumberStateChange\",\"type\":\"uint256\"}],\"internalType\":\"structJobTime\",\"name\":\"time\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"hasCancelRequest\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_providerAddr\",\"type\":\"address\"}],\"name\":\"metaSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerCancellingJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerClaimableJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerManager\",\"outputs\":[{\"internalType\":\"contractIProviderManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"enumJobStatus\",\"name\":\"_nextJobStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_jobDurationMinute\",\"type\":\"uint64\"}],\"name\":\"providerSetJobStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerTimeoutJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providerTopUpJobsQueues\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"_begin\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"_end\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"refuseJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"gpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"memPerCpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpuPerTask\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ntasks\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"batchLocationHash\",\"type\":\"string\"},{\"internalType\":\"enumStorageType\",\"name\":\"storageType\",\"type\":\"uint8\"}],\"internalType\":\"structJobDefinition\",\"name\":\"_definition\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_lockedCredits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_jobName\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_delegateSpendingAuthority\",\"type\":\"bool\"}],\"name\":\"requestNewJob\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_delegateSpendingAuthority\",\"type\":\"bool\"}],\"name\":\"setDelegateSpendingAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"topUpJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_jobId\",\"type\":\"bytes32\"}],\"name\":\"topUpJobDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallet2JobId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061a16480620000226000396000f3fe608060405234801561001057600080fd5b50600436106102695760003560e01c8063485cc95511610151578063aef3276f116100c3578063d547741f11610087578063d547741f1461078f578063d6aa37a6146107ab578063d77836ce146107c9578063e052888c146107e5578063e88fe8ca14610803578063ebd4bf001461083357610269565b8063aef3276f146106c4578063bc8f18fb146106f4578063c51067db14610725578063c7070e2c14610743578063d1cee5461461077357610269565b806369ee1bf91161011557806369ee1bf9146106025780637a9d0df71461061e5780638fb70f631461063c57806391d1485414610658578063a06d083c14610688578063a217fddf146106a657610269565b8063485cc9551461059a57806348841b9c146105b65780635d3a7180146105d25780635e1b2d65146105dc5780635fae1450146105e657610269565b8063248a9ca3116101ea578063329af326116101ae578063329af326146104aa57806336568abe146104c6578063374c6e0e146104e257806338ed7cfc146105005780633c2fb3da14610539578063407969ae1461056957610269565b8063248a9ca3146103f3578063257d9bb8146104235780632f2ff15d146104415780632fecc4f61461045d57806331c3b8741461047957610269565b80631a3cbef4116102315780631a3cbef41461033b5780631a91c3e91461036b5780631df44a3c146103895780631f92a63f146103a757806320a5f919146103c357610269565b806301ffc9a71461026e5780630797094e1461029e5780630f5a11ca146102ce578063110e87a6146102ec57806313151ec91461031d575b600080fd5b61028860048036038101906102839190617932565b61083d565b604051610295919061797a565b60405180910390f35b6102b860048036038101906102b391906179f3565b6108b7565b6040516102c5919061797a565b60405180910390f35b6102d6610909565b6040516102e39190617a39565b60405180910390f35b610306600480360381019061030191906179f3565b61090f565b604051610314929190617a70565b60405180910390f35b61032561094d565b6040516103329190617af8565b60405180910390f35b610355600480360381019061035091906179f3565b610973565b6040516103629190617bdb565b60405180910390f35b610373610a0a565b6040516103809190617c20565b60405180910390f35b610391610a0f565b60405161039e9190617a39565b60405180910390f35b6103c160048036038101906103bc9190617c67565b610a14565b005b6103dd60048036038101906103d891906179f3565b610ce4565b6040516103ea919061797a565b60405180910390f35b61040d60048036038101906104089190617c67565b610d34565b60405161041a9190617ca3565b60405180910390f35b61042b610d53565b6040516104389190617a39565b60405180910390f35b61045b60048036038101906104569190617cbe565b610d60565b005b61047760048036038101906104729190617d2a565b610d81565b005b610493600480360381019061048e91906179f3565b610d90565b6040516104a1929190617a70565b60405180910390f35b6104c460048036038101906104bf9190617d96565b610dce565b005b6104e060048036038101906104db9190617cbe565b610f79565b005b6104ea610ffc565b6040516104f79190617a39565b60405180910390f35b61051a60048036038101906105159190617c67565b611001565b6040516105309a99989796959493929190618094565b60405180910390f35b610553600480360381019061054e9190618384565b6112d0565b6040516105609190617ca3565b60405180910390f35b610583600480360381019061057e91906179f3565b611855565b604051610591929190617a70565b60405180910390f35b6105b460048036038101906105af9190618483565b611893565b005b6105d060048036038101906105cb91906184e8565b611b72565b005b6105da611e39565b005b6105e4611f70565b005b61060060048036038101906105fb9190617c67565b61210c565b005b61061c60048036038101906106179190617c67565b6123b0565b005b610626612c8a565b6040516106339190617a39565b60405180910390f35b61065660048036038101906106519190617cbe565b612c8f565b005b610672600480360381019061066d9190617cbe565b612f28565b60405161067f919061797a565b60405180910390f35b610690612f92565b60405161069d919061855c565b60405180910390f35b6106ae612fb8565b6040516106bb9190617ca3565b60405180910390f35b6106de60048036038101906106d99190618577565b612fbf565b6040516106eb9190617ca3565b60405180910390f35b61070e600480360381019061070991906179f3565b612fe3565b60405161071c929190617a70565b60405180910390f35b61072d613021565b60405161073a9190617a39565b60405180910390f35b61075d600480360381019061075891906179f3565b613026565b60405161076a919061797a565b60405180910390f35b61078d60048036038101906107889190617cbe565b613076565b005b6107a960048036038101906107a49190617cbe565b613247565b005b6107b3613268565b6040516107c09190617c20565b60405180910390f35b6107e360048036038101906107de9190618577565b613282565b005b6107ed61334d565b6040516107fa9190617ca3565b60405180910390f35b61081d600480360381019061081891906185a4565b613371565b60405161082a9190617ca3565b60405180910390f35b61083b6133a2565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806108b057506108af826137d5565b5b9050919050565b60006109026001603a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061383f565b9050919050565b60365481565b603a6020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060603960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156109fe57602002820191906000526020600020905b8154815260200190600101908083116109ea575b50505050509050919050565b601e81565b601481565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401610a7091906185e4565b602060405180830381865afa158015610a8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab19190618614565b610af0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae79061869e565b60405180910390fd5b816038600082815260200190815260200160002060060160009054906101000a900460ff16610b54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4b9061870a565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166038600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bf8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bef90618776565b60405180910390fd5b60026007811115610c0c57610c0b617dd6565b5b6038600085815260200190815260200160002060010160009054906101000a900460ff166007811115610c4257610c41617dd6565b5b1480610c95575060016007811115610c5d57610c5c617dd6565b5b6038600085815260200190815260200160002060010160009054906101000a900460ff166007811115610c9357610c92617dd6565b5b145b610cd4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ccb90618808565b60405180910390fd5b610cdf836000613c49565b505050565b6000610d2d603c60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614821565b9050919050565b6000806000838152602001908152602001600020600101549050919050565b68056bc75e2d6310000081565b610d6982610d34565b610d7281614c06565b610d7c8383614c1a565b505050565b610d8c823383614cfa565b5050565b603b6020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b816038600082815260200190815260200160002060060160009054906101000a900460ff16610e32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e299061870a565b60405180910390fd5b82610e5f6038600083815260200190815260200160002060010160009054906101000a900460ff16614ff5565b610e9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9590618874565b60405180910390fd5b6038600085815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f39906188e0565b60405180910390fd5b826038600086815260200190815260200160002060070160030160006101000a81548160ff02191690831515021790555050505050565b610f81615008565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610fee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe590618972565b60405180910390fd5b610ff88282615010565b5050565b600581565b60386020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054611161906189c1565b80601f016020809104026020016040519081016040528092919081815260200182805461118d906189c1565b80156111da5780601f106111af576101008083540402835291602001916111da565b820191906000526020600020905b8154815290600101906020018083116111bd57829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561120857611207617dd6565b5b600481111561121a57611219617dd6565b5b81525050908060060160009054906101000a900460ff1690806007016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250509080600b016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250509080600f0154908060100160009054906101000a900460ff1690508a565b600068056bc75e2d6310000084101561131e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161131590618a3e565b60405180910390fd5b611327856150f1565b611366576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161135d90618aaa565b60405180910390fd5b60006036600081548092919061137b90618af9565b9190505560001b9050604051806101400160405280828152602001600060078111156113aa576113a9617dd6565b5b81526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001878152602001600115158152602001604051806080016040528088815260200160008152602001600081526020018615158152508152602001604051806080016040528042815260200142815260200142815260200143815250815260200185815260200160001515815250603860008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff0219169083600781111561149c5761149b617dd6565b5b021790555060408201518160010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010190816116099190618ce3565b5060a08201518160020160006101000a81548160ff0219169083600481111561163557611634617dd6565b5b0217905550505060a08201518160060160006101000a81548160ff02191690831515021790555060c08201518160070160008201518160000155602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550505060e082015181600b0160008201518160000155602082015181600101556040820151816002015560608201518160030155505061010082015181600f01556101208201518160100160006101000a81548160ff021916908315150217905550905050603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff1660e01b815260040161176693929190618db5565b6020604051808303816000875af1158015611785573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117a99190618614565b50603960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150557f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada38133604051611841929190618dec565b60405180910390a180915050949350505050565b603c6020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b600060018054906101000a900460ff161590508080156118c4575060018060009054906101000a900460ff1660ff16105b806118f257506118d33061513f565b1580156118f1575060018060009054906101000a900460ff1660ff16145b5b611931576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161192890618e87565b60405180910390fd5b60018060006101000a81548160ff021916908360ff160217905550801561196d5760018060016101000a81548160ff0219169083151502179055505b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036119dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d390618ef3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a4290618f5f565b60405180910390fd5b611a586000801b33614c1a565b82603460086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600a603460006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600160368190555060006037819055508015611b6d5760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051611b649190618fc7565b60405180910390a15b505050565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401611bce91906185e4565b602060405180830381865afa158015611beb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c0f9190618614565b611c4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c459061869e565b60405180910390fd5b836038600082815260200190815260200160002060060160009054906101000a900460ff16611cb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ca99061870a565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611d56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d4d90619054565b60405180910390fd5b60036007811115611d6a57611d69617dd6565b5b846007811115611d7d57611d7c617dd6565b5b1480611d8e5750611d8d84615162565b5b611dcd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dc4906190c0565b60405180910390fd5b60036007811115611de157611de0617dd6565b5b6038600087815260200190815260200160002060010160009054906101000a900460ff166007811115611e1757611e16617dd6565b5b03611e2857611e2785338561521f565b5b611e328585613c49565b5050505050565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401611e9591906185e4565b602060405180830381865afa158015611eb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ed69190618614565b611f15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f0c9061869e565b60405180910390fd5b6000611f606001603a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061580d565b9050611f6c8133612c8f565b5050565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401611fcc91906185e4565b602060405180830381865afa158015611fe9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061200d9190618614565b61204c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120439061869e565b60405180910390fd5b6000612095603c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615ca9565b90507f290fa751f58fe2a1f5758b401eb3110dbbb71b68540282856c0dcdcc7011e07d6038600083815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff163383604051612100939291906190e0565b60405180910390a15050565b806038600082815260200190815260200160002060060160009054906101000a900460ff16612170576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121679061870a565b60405180910390fd5b8161219d6038600083815260200190815260200160002060010160009054906101000a900460ff16614ff5565b6121dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121d390618874565b60405180910390fd5b6038600084815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161227790619189565b60405180910390fd5b60016038600085815260200190815260200160002060100160006101000a81548160ff0219169083151502179055504260386000858152602001908152602001600020600b0160020181905550600360078111156122e1576122e0617dd6565b5b6038600085815260200190815260200160002060010160009054906101000a900460ff16600781111561231757612316617dd6565b5b0361239f5761239a603c60006038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846160fa565b6123ab565b6123aa836004613c49565b5b505050565b6038600082815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061248157506038600082815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806124b257506124b17f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e33612f28565b5b6124f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124e8906191f5565b60405180910390fd5b6038600082815260200190815260200160002060070160030160009054906101000a900460ff16612557576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161254e90619261565b60405180910390fd5b6125836038600083815260200190815260200160002060010160009054906101000a900460ff16616176565b6125c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125b9906192cd565b60405180910390fd5b6000603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355f21eb76038600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040161265591906185e4565b6101a060405180830381865afa158015612673573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061269791906194bf565b602001519050601e67ffffffffffffffff16612a346038600085815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff1660078111156126fa576126f9617dd6565b5b600781111561270c5761270b617dd6565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546128a4906189c1565b80601f01602080910402602001604051908101604052809291908181526020018280546128d0906189c1565b801561291d5780601f106128f25761010080835404028352916020019161291d565b820191906000526020600020905b81548152906001019060200180831161290057829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561294b5761294a617dd6565b5b600481111561295d5761295c617dd6565b5b8152505081526020016006820160009054906101000a900460ff16151515158152602001600782016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250508152602001600b82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600f82015481526020016010820160009054906101000a900460ff161515151581525050836161d6565b67ffffffffffffffff161115612a7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a7690619539565b60405180910390fd5b6000612c42603860008581526020019081526020016000206003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054612b7d906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054612ba9906189c1565b8015612bf65780601f10612bcb57610100808354040283529160200191612bf6565b820191906000526020600020905b815481529060010190602001808311612bd957829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115612c2457612c23617dd6565b5b6004811115612c3657612c35617dd6565b5b8152505083601e616236565b9050612c85836038600086815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683614cfa565b505050565b600f81565b80603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b8152600401612ceb91906185e4565b602060405180830381865afa158015612d08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d2c9190618614565b612d6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d629061869e565b60405180910390fd5b826038600082815260200190815260200160002060060160009054906101000a900460ff16612dcf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dc69061870a565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480612e3457503073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b612e73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e6a906195a5565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166038600086815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612f17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f0e90618776565b60405180910390fd5b612f22846002613c49565b50505050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000801b81565b603e8181548110612fcf57600080fd5b906000526020600020016000915090505481565b603d6020528060005260406000206000915090508060000160009054906101000a9004600f0b908060000160109054906101000a9004600f0b905082565b600f81565b600061306f603d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020616324565b9050919050565b80603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b81526004016130d291906185e4565b602060405180830381865afa1580156130ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131139190618614565b613152576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131499061869e565b60405180910390fd5b826038600082815260200190815260200160002060060160009054906101000a900460ff166131b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131ad9061870a565b60405180910390fd5b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e6131e081614c06565b836038600087815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613240856001613c49565b5050505050565b61325082610d34565b61325981614c06565b6132638383615010565b505050565b603460009054906101000a900467ffffffffffffffff1681565b6000801b61328f81614c06565b81603760008282546132a191906195c5565b92505081905550603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016133059291906195f9565b6020604051808303816000875af1158015613324573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133489190618614565b505050565b7f34fe770ac2473ec704bda003df1f7ec520ba6602bc5ebb22f4d41610283d996e81565b6039602052816000526040600020818154811061338d57600080fd5b90600052602060002001600091509150505481565b33603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663877f4e12826040518263ffffffff1660e01b81526004016133fe91906185e4565b602060405180830381865afa15801561341b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061343f9190618614565b61347e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134759061869e565b60405180910390fd5b60006134c7603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061670f565b90507fa42f2b4a7ee7f91857a4c98fc71fc48546a284d5db48dd77b7ab81030a494470813361379b603860008681526020019081526020016000206003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546135e8906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054613614906189c1565b80156136615780601f1061363657610100808354040283529160200191613661565b820191906000526020600020905b81548152906001019060200180831161364457829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561368f5761368e617dd6565b5b60048111156136a1576136a0617dd6565b5b815250506038600087815260200190815260200160002060070160020154603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355f21eb7603860008a815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040161375091906185e4565b6101a060405180830381865afa15801561376e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061379291906194bf565b60200151616b6b565b6040516137aa93929190619622565b60405180910390a1600060386000838152602001908152602001600020600701600201819055505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008061384a617757565b60005b61385685616c97565b811015613c3b576138678582616ccb565b92506038600084815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff1660078111156138b7576138b6617dd6565b5b60078111156138c9576138c8617dd6565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054613a61906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054613a8d906189c1565b8015613ada5780601f10613aaf57610100808354040283529160200191613ada565b820191906000526020600020905b815481529060010190602001808311613abd57829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115613b0857613b07617dd6565b5b6004811115613b1a57613b19617dd6565b5b8152505081526020016006820160009054906101000a900460ff16151515158152602001600782016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250508152602001600b82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600f82015481526020016010820160009054906101000a900460ff1615151515815250509150856007811115613c0057613bff617dd6565b5b82602001516007811115613c1757613c16617dd6565b5b03613c285760019350505050613c43565b8080613c3390618af9565b91505061384d565b506000925050505b92915050565b816038600082815260200190815260200160002060060160009054906101000a900460ff16613cad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ca49061870a565b60405180910390fd5b613cb5616d75565b613ce26038600085815260200190815260200160002060010160009054906101000a900460ff1683616dc2565b7f0bba917f0a1e0fc0d51a75273e7088a4dfecb010699e60ac9c58526429f6c37f836038600086815260200190815260200160002060010160009054906101000a900460ff1684604051613d3893929190619659565b60405180910390a1816038600085815260200190815260200160002060010160006101000a81548160ff02191690836007811115613d7957613d78617dd6565b5b02179055504360386000858152602001908152602001600020600b016003018190555060006007811115613db057613daf617dd6565b5b826007811115613dc357613dc2617dd6565b5b03613f63577f5d0260cf2f490cac7a98928e721dcc1c49f1bcc33458b3103755adfd1c1eada0836038600086815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051613e2f929190618dec565b60405180910390a17f1b831e9023e41b1f2ae42f1cb9a173ca2de2eb05475bf206d3762717a826ada3836038600086815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051613e9e929190618dec565b60405180910390a160006038600085815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180608001604052804281526020014281526020014281526020014381525060386000858152602001908152602001600020600b0160008201518160000155602082015181600101556040820151816002015560608201518160030155905050614814565b60016007811115613f7757613f76617dd6565b5b826007811115613f8a57613f89617dd6565b5b0361403b57603e839080600181540180825580915050600190039060005260206000200160009091909190915055614036603a60006038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846160fa565b614813565b6002600781111561404f5761404e617dd6565b5b82600781111561406257614061617dd6565b5b0361445e576140e5603b60006038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846160fa565b6000603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355f21eb76038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040161417891906185e4565b6101a060405180830381865afa158015614196573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141ba91906194bf565b90507fc3037ff6238c842f0908a76f68fa0ec2490f1096e61e371d93fe9edca33c3c396038600086815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686614429603860008a81526020019081526020016000206003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054614348906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054614374906189c1565b80156143c15780601f10614396576101008083540402835291602001916143c1565b820191906000526020600020905b8154815290600101906020018083116143a457829003601f168201915b505050505081526020016002820160009054906101000a900460ff1660048111156143ef576143ee617dd6565b5b600481111561440157614400617dd6565b5b81525050603860008b8152602001908152602001600020600701600001548760200151616b6b565b603860008a8152602001908152602001600020600301604051614450959493929190619897565b60405180910390a150614812565b6003600781111561447257614471617dd6565b5b82600781111561448557614484617dd6565b5b03614570574260386000858152602001908152602001600020600b0160000181905550603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f6edb5f6038600086815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040161453991906185e4565b600060405180830381600087803b15801561455357600080fd5b505af1158015614567573d6000803e3d6000fd5b50505050614811565b4260386000858152602001908152602001600020600b016001018190555060006064603860008681526020019081526020016000206007016001015460146145b891906198f1565b6145c29190619962565b905080603760008282546145d69190619993565b925050819055506000603860008681526020019081526020016000206007016001015411156146fb57603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6038600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683603860008981526020019081526020016000206007016001015461469991906195c5565b6040518363ffffffff1660e01b81526004016146b69291906195f9565b6020604051808303816000875af11580156146d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146f99190618614565b505b603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6038600087815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16603860008881526020019081526020016000206007016001015460386000898152602001908152602001600020600701600001546147ae91906195c5565b6040518363ffffffff1660e01b81526004016147cb9291906195f9565b6020604051808303816000875af11580156147ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061480e9190618614565b50505b5b5b5b61481c617279565b505050565b60008061482c617757565b60005b61483885616c97565b811015614bf9576148498582616ccb565b92506038600084815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff16600781111561489957614898617dd6565b5b60078111156148ab576148aa617dd6565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054614a43906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054614a6f906189c1565b8015614abc5780601f10614a9157610100808354040283529160200191614abc565b820191906000526020600020905b815481529060010190602001808311614a9f57829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115614aea57614ae9617dd6565b5b6004811115614afc57614afb617dd6565b5b8152505081526020016006820160009054906101000a900460ff16151515158152602001600782016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250508152602001600b82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600f82015481526020016010820160009054906101000a900460ff161515151581525050915081610120015115614be65760019350505050614c01565b8080614bf190618af9565b91505061482f565b506000925050505b919050565b614c1781614c12615008565b617283565b50565b614c248282612f28565b614cf657600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550614c9b615008565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b826038600082815260200190815260200160002060060160009054906101000a900460ff16614d5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d559061870a565b60405180910390fd5b83614d8b6038600083815260200190815260200160002060010160009054906101000a900460ff16614ff5565b614dca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614dc190618874565b60405180910390fd5b82603860008781526020019081526020016000206007016000016000828254614df39190619993565b9250508190555060026007811115614e0e57614e0d617dd6565b5b6038600087815260200190815260200160002060010160009054906101000a900460ff166007811115614e4457614e43617dd6565b5b1480614e97575060036007811115614e5f57614e5e617dd6565b5b6038600087815260200190815260200160002060010160009054906101000a900460ff166007811115614e9557614e94617dd6565b5b145b15614f4b5782603860008781526020019081526020016000206007016002016000828254614ec59190619993565b92505081905550614f4a603d60006038600089815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020866160fa565b5b603460089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8530866040518463ffffffff1660e01b8152600401614faa93929190618db5565b6020604051808303816000875af1158015614fc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614fed9190618614565b505050505050565b600061500082615162565b159050919050565b600033905090565b61501a8282612f28565b156150ed57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550615092615008565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600080826060015167ffffffffffffffff1611801561511e57506000826040015167ffffffffffffffff16115b801561513857506000826020015167ffffffffffffffff16115b9050919050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60006004600781111561517857615177617dd6565b5b82600781111561518b5761518a617dd6565b5b14806151bb5750600560078111156151a6576151a5617dd6565b5b8260078111156151b9576151b8617dd6565b5b145b806151e957506007808111156151d4576151d3617dd6565b5b8260078111156151e7576151e6617dd6565b5b145b8061521857506006600781111561520357615202617dd6565b5b82600781111561521657615215617dd6565b5b145b9050919050565b6000603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166355f21eb7846040518263ffffffff1660e01b815260040161527c91906185e4565b6101a060405180830381865afa15801561529a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906152be91906194bf565b905060008267ffffffffffffffff16036152d757600191505b600f603c60386000878152602001908152602001600020600b01600001544261530091906195c5565b61530a9190619962565b6153149190619993565b8267ffffffffffffffff161115615394577f17e65314b087df225f56701d0a66a3f7d9ce0f26077307b4b765a19c60a36d44848484604051615358939291906199f8565b60405180910390a1603c60386000868152602001908152602001600020600b01600001544261538791906195c5565b6153919190619962565b91505b600061555a603860008781526020019081526020016000206003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054615492906189c1565b80601f01602080910402602001604051908101604052809291908181526020018280546154be906189c1565b801561550b5780601f106154e05761010080835404028352916020019161550b565b820191906000526020600020905b8154815290600101906020018083116154ee57829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561553957615538617dd6565b5b600481111561554b5761554a617dd6565b5b81525050836020015185616236565b90506038600086815260200190815260200160002060100160009054906101000a900460ff16156157aa576157a7816157a2603860008981526020019081526020016000206003016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054615685906189c1565b80601f01602080910402602001604051908101604052809291908181526020018280546156b1906189c1565b80156156fe5780601f106156d3576101008083540402835291602001916156fe565b820191906000526020600020905b8154815290600101906020018083116156e157829003601f168201915b505050505081526020016002820160009054906101000a900460ff16600481111561572c5761572b617dd6565b5b600481111561573e5761573d617dd6565b5b8152505085602001516005603c603860008d8152602001908152602001600020600b0160000154603860008e8152602001908152602001600020600b016002015461578991906195c5565b6157939190619962565b61579d9190619993565b616236565b617308565b90505b60386000868152602001908152602001600020600701600001548111156157e857603860008681526020019081526020016000206007016000015490505b8060386000878152602001908152602001600020600701600101819055505050505050565b600061581882617321565b15615858576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161584f90619a7b565b60405180910390fd5b6000615862617757565b5b61586c84617356565b91506038600083815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff1660078111156158bc576158bb617dd6565b5b60078111156158ce576158cd617dd6565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054615a66906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054615a92906189c1565b8015615adf5780601f10615ab457610100808354040283529160200191615adf565b820191906000526020600020905b815481529060010190602001808311615ac257829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115615b0d57615b0c617dd6565b5b6004811115615b1f57615b1e617dd6565b5b8152505081526020016006820160009054906101000a900460ff16151515158152602001600782016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250508152602001600b82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600f82015481526020016010820160009054906101000a900460ff1615151515815250509050615bfb84617321565b158015615c305750846007811115615c1657615c15617dd6565b5b81602001516007811115615c2d57615c2c617dd6565b5b14155b61586357846007811115615c4757615c46617dd6565b5b81602001516007811115615c5e57615c5d617dd6565b5b14615c9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615c9590619a7b565b60405180910390fd5b819250505092915050565b6000615cb482617321565b15615cf4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615ceb90619a7b565b60405180910390fd5b6000615cfe617757565b5b615d0884617356565b91506038600083815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff166007811115615d5857615d57617dd6565b5b6007811115615d6a57615d69617dd6565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054615f02906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054615f2e906189c1565b8015615f7b5780601f10615f5057610100808354040283529160200191615f7b565b820191906000526020600020905b815481529060010190602001808311615f5e57829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115615fa957615fa8617dd6565b5b6004811115615fbb57615fba617dd6565b5b8152505081526020016006820160009054906101000a900460ff16151515158152602001600782016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250508152602001600b82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600f82015481526020016010820160009054906101000a900460ff161515151581525050905061609784617321565b1580156160a75750806101200151155b615cff578061012001516160f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016160e790619a7b565b60405180910390fd5b8192505050919050565b60008260000160109054906101000a9004600f0b90508183600101600083600f0b600f0b815260200190815260200160002081905550600181018360000160106101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff160217905550505050565b60006003600781111561618c5761618b617dd6565b5b82600781111561619f5761619e617dd6565b5b14806161cf5750600260078111156161ba576161b9617dd6565b5b8260078111156161cd576161cc617dd6565b5b145b9050919050565b6000806161f084608001518560c001516000015185616b6b565b90506000603c8560e00151600001514261620a91906195c5565b6162149190619962565b9050808267ffffffffffffffff1661622c91906195c5565b9250505092915050565b600064e8d4a51000846040015167ffffffffffffffff16846080015167ffffffffffffffff1661626691906198f1565b856040015167ffffffffffffffff16866020015167ffffffffffffffff168660c0015167ffffffffffffffff1661629d91906198f1565b6162a791906198f1565b866000015167ffffffffffffffff16866040015167ffffffffffffffff166162cf91906198f1565b6162d99190619993565b6162e39190619993565b856060015167ffffffffffffffff168467ffffffffffffffff1661630791906198f1565b61631191906198f1565b61631b91906198f1565b90509392505050565b60008061632f617757565b60005b61633b85616c97565b8110156167025761634c8582616ccb565b92506038600084815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff16600781111561639c5761639b617dd6565b5b60078111156163ae576163ad617dd6565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054616546906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054616572906189c1565b80156165bf5780601f10616594576101008083540402835291602001916165bf565b820191906000526020600020905b8154815290600101906020018083116165a257829003601f168201915b505050505081526020016002820160009054906101000a900460ff1660048111156165ed576165ec617dd6565b5b60048111156165ff576165fe617dd6565b5b8152505081526020016006820160009054906101000a900460ff16151515158152602001600782016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250508152602001600b82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600f82015481526020016010820160009054906101000a900460ff161515151581525050915060008260c001516040015111156166ef576001935050505061670a565b80806166fa90618af9565b915050616332565b506000925050505b919050565b600061671a82617321565b1561675a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161675190619a7b565b60405180910390fd5b6000616764617757565b5b61676e84617356565b91506038600083815260200190815260200160002060405180610140016040529081600082015481526020016001820160009054906101000a900460ff1660078111156167be576167bd617dd6565b5b60078111156167d0576167cf617dd6565b5b81526020016001820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382016040518060c00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054616968906189c1565b80601f0160208091040260200160405190810160405280929190818152602001828054616994906189c1565b80156169e15780601f106169b6576101008083540402835291602001916169e1565b820191906000526020600020905b8154815290600101906020018083116169c457829003601f168201915b505050505081526020016002820160009054906101000a900460ff166004811115616a0f57616a0e617dd6565b5b6004811115616a2157616a20617dd6565b5b8152505081526020016006820160009054906101000a900460ff16151515158152602001600782016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250508152602001600b82016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250508152602001600f82015481526020016010820160009054906101000a900460ff1615151515815250509050616afd84617321565b158015616b12575060008160c0015160400151145b6167655760008160c001516040015103616b61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616b5890619a7b565b60405180910390fd5b8192505050919050565b60008064e8d4a51000856040015167ffffffffffffffff16846080015167ffffffffffffffff16616b9c91906198f1565b866040015167ffffffffffffffff16876020015167ffffffffffffffff168660c0015167ffffffffffffffff16616bd391906198f1565b616bdd91906198f1565b876000015167ffffffffffffffff16866040015167ffffffffffffffff16616c0591906198f1565b616c0f9190619993565b616c199190619993565b866060015167ffffffffffffffff16616c3291906198f1565b616c3c91906198f1565b905060008103616c81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616c7890619ae7565b60405180910390fd5b8084616c8d9190619962565b9150509392505050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b039050919050565b600080616cfe616cda84617432565b8560000160009054906101000a9004600f0b600f0b616cf99190619b11565b61749f565b90508360000160109054906101000a9004600f0b600f0b81600f0b12616d50576040517fb4120f1400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600101600082600f0b600f0b81526020019081526020016000205491505092915050565b6002805403616db9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616db090619ba1565b60405180910390fd5b60028081905550565b806007811115616dd557616dd4617dd6565b5b826007811115616de857616de7617dd6565b5b03616e28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616e1f90619c0d565b60405180910390fd5b616e3182614ff5565b616e70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616e6790618874565b60405180910390fd5b60006007811115616e8457616e83617dd6565b5b826007811115616e9757616e96617dd6565b5b03616f375760016007811115616eb057616eaf617dd6565b5b816007811115616ec357616ec2617dd6565b5b1480616ef3575060046007811115616ede57616edd617dd6565b5b816007811115616ef157616ef0617dd6565b5b145b616f32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616f2990619c79565b60405180910390fd5b617275565b60016007811115616f4b57616f4a617dd6565b5b826007811115616f5e57616f5d617dd6565b5b0361702d5760006007811115616f7757616f76617dd6565b5b816007811115616f8a57616f89617dd6565b5b1480616fba575060026007811115616fa557616fa4617dd6565b5b816007811115616fb857616fb7617dd6565b5b145b80616fe9575060046007811115616fd457616fd3617dd6565b5b816007811115616fe757616fe6617dd6565b5b145b617028576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161701f90619ce5565b60405180910390fd5b617274565b6002600781111561704157617040617dd6565b5b82600781111561705457617053617dd6565b5b03617152576000600781111561706d5761706c617dd6565b5b8160078111156170805761707f617dd6565b5b14806170b057506003600781111561709b5761709a617dd6565b5b8160078111156170ae576170ad617dd6565b5b145b806170df5750600460078111156170ca576170c9617dd6565b5b8160078111156170dd576170dc617dd6565b5b145b8061710e5750600660078111156170f9576170f8617dd6565b5b81600781111561710c5761710b617dd6565b5b145b61714d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161714490619d51565b60405180910390fd5b617273565b6003600781111561716657617165617dd6565b5b82600781111561717957617178617dd6565b5b03617272576005600781111561719257617191617dd6565b5b8160078111156171a5576171a4617dd6565b5b14806171d55750600660078111156171c0576171bf617dd6565b5b8160078111156171d3576171d2617dd6565b5b145b806172045750600460078111156171ef576171ee617dd6565b5b81600781111561720257617201617dd6565b5b145b80617232575060078081111561721d5761721c617dd6565b5b8160078111156172305761722f617dd6565b5b145b617271576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161726890619dbd565b60405180910390fd5b5b5b5b5b5050565b6001600281905550565b61728d8282612f28565b6173045761729a816174ee565b6172a88360001c602061751b565b6040516020016172b9929190619eb1565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016172fb9190619f24565b60405180910390fd5b5050565b60008183106173175781617319565b825b905092915050565b60008160000160009054906101000a9004600f0b600f0b8260000160109054906101000a9004600f0b600f0b13159050919050565b600061736182617321565b15617398576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260000160009054906101000a9004600f0b905082600101600082600f0b600f0b815260200190815260200160002054915082600101600082600f0b600f0b815260200190815260200160002060009055600181018360000160006101000a8154816fffffffffffffffffffffffffffffffff0219169083600f0b6fffffffffffffffffffffffffffffffff16021790555050919050565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115617497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161748e90619fb8565b60405180910390fd5b819050919050565b60008190508181600f0b146174e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016174e09061a04a565b60405180910390fd5b919050565b60606175148273ffffffffffffffffffffffffffffffffffffffff16601460ff1661751b565b9050919050565b60606000600283600261752e91906198f1565b6175389190619993565b67ffffffffffffffff8111156175515761755061813d565b5b6040519080825280601f01601f1916602001820160405280156175835781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106175bb576175ba61a06a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061761f5761761e61a06a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000600184600261765f91906198f1565b6176699190619993565b90505b6001811115617709577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106176ab576176aa61a06a565b5b1a60f81b8282815181106176c2576176c161a06a565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806177029061a099565b905061766c565b506000841461774d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016177449061a10e565b60405180910390fd5b8091505092915050565b604051806101400160405280600080191681526020016000600781111561778157617780617dd6565b5b8152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016177c8617804565b81526020016000151581526020016177de617874565b81526020016177eb61789e565b8152602001600080191681526020016000151581525090565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016000600481111561786e5761786d617dd6565b5b81525090565b60405180608001604052806000815260200160008152602001600081526020016000151581525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61790f816178da565b811461791a57600080fd5b50565b60008135905061792c81617906565b92915050565b600060208284031215617948576179476178d0565b5b60006179568482850161791d565b91505092915050565b60008115159050919050565b6179748161795f565b82525050565b600060208201905061798f600083018461796b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006179c082617995565b9050919050565b6179d0816179b5565b81146179db57600080fd5b50565b6000813590506179ed816179c7565b92915050565b600060208284031215617a0957617a086178d0565b5b6000617a17848285016179de565b91505092915050565b6000819050919050565b617a3381617a20565b82525050565b6000602082019050617a4e6000830184617a2a565b92915050565b600081600f0b9050919050565b617a6a81617a54565b82525050565b6000604082019050617a856000830185617a61565b617a926020830184617a61565b9392505050565b6000819050919050565b6000617abe617ab9617ab484617995565b617a99565b617995565b9050919050565b6000617ad082617aa3565b9050919050565b6000617ae282617ac5565b9050919050565b617af281617ad7565b82525050565b6000602082019050617b0d6000830184617ae9565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b617b5281617b3f565b82525050565b6000617b648383617b49565b60208301905092915050565b6000602082019050919050565b6000617b8882617b13565b617b928185617b1e565b9350617b9d83617b2f565b8060005b83811015617bce578151617bb58882617b58565b9750617bc083617b70565b925050600181019050617ba1565b5085935050505092915050565b60006020820190508181036000830152617bf58184617b7d565b905092915050565b600067ffffffffffffffff82169050919050565b617c1a81617bfd565b82525050565b6000602082019050617c356000830184617c11565b92915050565b617c4481617b3f565b8114617c4f57600080fd5b50565b600081359050617c6181617c3b565b92915050565b600060208284031215617c7d57617c7c6178d0565b5b6000617c8b84828501617c52565b91505092915050565b617c9d81617b3f565b82525050565b6000602082019050617cb86000830184617c94565b92915050565b60008060408385031215617cd557617cd46178d0565b5b6000617ce385828601617c52565b9250506020617cf4858286016179de565b9150509250929050565b617d0781617a20565b8114617d1257600080fd5b50565b600081359050617d2481617cfe565b92915050565b60008060408385031215617d4157617d406178d0565b5b6000617d4f85828601617c52565b9250506020617d6085828601617d15565b9150509250929050565b617d738161795f565b8114617d7e57600080fd5b50565b600081359050617d9081617d6a565b92915050565b60008060408385031215617dad57617dac6178d0565b5b6000617dbb85828601617c52565b9250506020617dcc85828601617d81565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60088110617e1657617e15617dd6565b5b50565b6000819050617e2782617e05565b919050565b6000617e3782617e19565b9050919050565b617e4781617e2c565b82525050565b617e56816179b5565b82525050565b617e6581617bfd565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015617ea5578082015181840152602081019050617e8a565b60008484015250505050565b6000601f19601f8301169050919050565b6000617ecd82617e6b565b617ed78185617e76565b9350617ee7818560208601617e87565b617ef081617eb1565b840191505092915050565b60058110617f0c57617f0b617dd6565b5b50565b6000819050617f1d82617efb565b919050565b6000617f2d82617f0f565b9050919050565b617f3d81617f22565b82525050565b600060c083016000830151617f5b6000860182617e5c565b506020830151617f6e6020860182617e5c565b506040830151617f816040860182617e5c565b506060830151617f946060860182617e5c565b5060808301518482036080860152617fac8282617ec2565b91505060a0830151617fc160a0860182617f34565b508091505092915050565b617fd581617a20565b82525050565b617fe48161795f565b82525050565b6080820160008201516180006000850182617fcc565b5060208201516180136020850182617fcc565b5060408201516180266040850182617fcc565b5060608201516180396060850182617fdb565b50505050565b6080820160008201516180556000850182617fcc565b5060208201516180686020850182617fcc565b50604082015161807b6040850182617fcc565b50606082015161808e6060850182617fcc565b50505050565b6000610200820190506180aa600083018d617c94565b6180b7602083018c617e3e565b6180c4604083018b617e4d565b6180d1606083018a617e4d565b81810360808301526180e38189617f43565b90506180f260a083018861796b565b6180ff60c0830187617fea565b61810d61014083018661803f565b61811b6101c0830185617c94565b6181296101e083018461796b565b9b9a5050505050505050505050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61817582617eb1565b810181811067ffffffffffffffff821117156181945761819361813d565b5b80604052505050565b60006181a76178c6565b90506181b3828261816c565b919050565b600080fd5b6181c681617bfd565b81146181d157600080fd5b50565b6000813590506181e3816181bd565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff82111561820e5761820d61813d565b5b61821782617eb1565b9050602081019050919050565b82818337600083830152505050565b6000618246618241846181f3565b61819d565b905082815260208101848484011115618262576182616181ee565b5b61826d848285618224565b509392505050565b600082601f83011261828a576182896181e9565b5b813561829a848260208601618233565b91505092915050565b600581106182b057600080fd5b50565b6000813590506182c2816182a3565b92915050565b600060c082840312156182de576182dd618138565b5b6182e860c061819d565b905060006182f8848285016181d4565b600083015250602061830c848285016181d4565b6020830152506040618320848285016181d4565b6040830152506060618334848285016181d4565b606083015250608082013567ffffffffffffffff811115618358576183576181b8565b5b61836484828501618275565b60808301525060a0618378848285016182b3565b60a08301525092915050565b6000806000806080858703121561839e5761839d6178d0565b5b600085013567ffffffffffffffff8111156183bc576183bb6178d5565b5b6183c8878288016182c8565b94505060206183d987828801617d15565b93505060406183ea87828801617c52565b92505060606183fb87828801617d81565b91505092959194509250565b6000618412826179b5565b9050919050565b61842281618407565b811461842d57600080fd5b50565b60008135905061843f81618419565b92915050565b6000618450826179b5565b9050919050565b61846081618445565b811461846b57600080fd5b50565b60008135905061847d81618457565b92915050565b6000806040838503121561849a576184996178d0565b5b60006184a885828601618430565b92505060206184b98582860161846e565b9150509250929050565b600881106184d057600080fd5b50565b6000813590506184e2816184c3565b92915050565b600080600060608486031215618501576185006178d0565b5b600061850f86828701617c52565b9350506020618520868287016184d3565b9250506040618531868287016181d4565b9150509250925092565b600061854682617ac5565b9050919050565b6185568161853b565b82525050565b6000602082019050618571600083018461854d565b92915050565b60006020828403121561858d5761858c6178d0565b5b600061859b84828501617d15565b91505092915050565b600080604083850312156185bb576185ba6178d0565b5b60006185c9858286016179de565b92505060206185da85828601617d15565b9150509250929050565b60006020820190506185f96000830184617e4d565b92915050565b60008151905061860e81617d6a565b92915050565b60006020828403121561862a576186296178d0565b5b6000618638848285016185ff565b91505092915050565b600082825260208201905092915050565b7f4a4f494e45440000000000000000000000000000000000000000000000000000600082015250565b6000618688600683618641565b915061869382618652565b602082019050919050565b600060208201905081810360008301526186b78161867b565b9050919050565b7f4a6f62206e6f742076616c696421000000000000000000000000000000000000600082015250565b60006186f4600e83618641565b91506186ff826186be565b602082019050919050565b60006020820190508181036000830152618723816186e7565b9050919050565b7f4a42380000000000000000000000000000000000000000000000000000000000600082015250565b6000618760600383618641565b915061876b8261872a565b602082019050919050565b6000602082019050818103600083015261878f81618753565b9050919050565b7f43616e277420726566757365206966206e6f7420696e204d4554415f5343484560008201527f44554c4544206f72205343484544554c45442073746174650000000000000000602082015250565b60006187f2603883618641565b91506187fd82618796565b604082019050919050565b60006020820190508181036000830152618821816187e5565b9050919050565b7f484f540000000000000000000000000000000000000000000000000000000000600082015250565b600061885e600383618641565b915061886982618828565b602082019050919050565b6000602082019050818103600083015261888d81618851565b9050919050565b7f5045524d31000000000000000000000000000000000000000000000000000000600082015250565b60006188ca600583618641565b91506188d582618894565b602082019050919050565b600060208201905081810360008301526188f9816188bd565b9050919050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b600061895c602f83618641565b915061896782618900565b604082019050919050565b6000602082019050818103600083015261898b8161894f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806189d957607f821691505b6020821081036189ec576189eb618992565b5b50919050565b7f42414c3100000000000000000000000000000000000000000000000000000000600082015250565b6000618a28600483618641565b9150618a33826189f2565b602082019050919050565b60006020820190508181036000830152618a5781618a1b565b9050919050565b7f4a44454600000000000000000000000000000000000000000000000000000000600082015250565b6000618a94600483618641565b9150618a9f82618a5e565b602082019050919050565b60006020820190508181036000830152618ac381618a87565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000618b0482617a20565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203618b3657618b35618aca565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302618ba37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82618b66565b618bad8683618b66565b95508019841693508086168417925050509392505050565b6000618be0618bdb618bd684617a20565b617a99565b617a20565b9050919050565b6000819050919050565b618bfa83618bc5565b618c0e618c0682618be7565b848454618b73565b825550505050565b600090565b618c23618c16565b618c2e818484618bf1565b505050565b5b81811015618c5257618c47600082618c1b565b600181019050618c34565b5050565b601f821115618c9757618c6881618b41565b618c7184618b56565b81016020851015618c80578190505b618c94618c8c85618b56565b830182618c33565b50505b505050565b600082821c905092915050565b6000618cba60001984600802618c9c565b1980831691505092915050565b6000618cd38383618ca9565b9150826002028217905092915050565b618cec82617e6b565b67ffffffffffffffff811115618d0557618d0461813d565b5b618d0f82546189c1565b618d1a828285618c56565b600060209050601f831160018114618d4d5760008415618d3b578287015190505b618d458582618cc7565b865550618dad565b601f198416618d5b86618b41565b60005b82811015618d8357848901518255600182019150602085019450602081019050618d5e565b86831015618da05784890151618d9c601f891682618ca9565b8355505b6001600288020188555050505b505050505050565b6000606082019050618dca6000830186617e4d565b618dd76020830185617e4d565b618de46040830184617a2a565b949350505050565b6000604082019050618e016000830185617c94565b618e0e6020830184617e4d565b9392505050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000618e71602e83618641565b9150618e7c82618e15565b604082019050919050565b60006020820190508181036000830152618ea081618e64565b9050919050565b7f4372656469742061646472206973203000000000000000000000000000000000600082015250565b6000618edd601083618641565b9150618ee882618ea7565b602082019050919050565b60006020820190508181036000830152618f0c81618ed0565b9050919050565b7f50726f76696465724d616e616765722061646472206973203000000000000000600082015250565b6000618f49601983618641565b9150618f5482618f13565b602082019050919050565b60006020820190508181036000830152618f7881618f3c565b9050919050565b6000819050919050565b600060ff82169050919050565b6000618fb1618fac618fa784618f7f565b617a99565b618f89565b9050919050565b618fc181618f96565b82525050565b6000602082019050618fdc6000830184618fb8565b92915050565b7f4f6e6c7920746865206a6f622070726f76696465722063616e206368616e676560008201527f2069747320737461747573000000000000000000000000000000000000000000602082015250565b600061903e602b83618641565b915061904982618fe2565b604082019050919050565b6000602082019050818103600083015261906d81619031565b9050919050565b7f4a42313300000000000000000000000000000000000000000000000000000000600082015250565b60006190aa600483618641565b91506190b582619074565b602082019050919050565b600060208201905081810360008301526190d98161909d565b9050919050565b60006060820190506190f56000830186617e4d565b6191026020830185617e4d565b61910f6040830184617c94565b949350505050565b7f4f6e6c79206a6f62206f776e6572732063616e2063616e63656c20746865697260008201527f206a6f6273000000000000000000000000000000000000000000000000000000602082015250565b6000619173602583618641565b915061917e82619117565b604082019050919050565b600060208201905081810360008301526191a281619166565b9050919050565b7f5045524d32000000000000000000000000000000000000000000000000000000600082015250565b60006191df600583618641565b91506191ea826191a9565b602082019050919050565b6000602082019050818103600083015261920e816191d2565b9050919050565b7f4a42323600000000000000000000000000000000000000000000000000000000600082015250565b600061924b600483618641565b915061925682619215565b602082019050919050565b6000602082019050818103600083015261927a8161923e565b9050919050565b7f534348454452554e000000000000000000000000000000000000000000000000600082015250565b60006192b7600883618641565b91506192c282619281565b602082019050919050565b600060208201905081810360008301526192e6816192aa565b9050919050565b6000815190506192fc816179c7565b92915050565b600081519050619311816181bd565b92915050565b600060e0828403121561932d5761932c618138565b5b61933760e061819d565b9050600061934784828501619302565b600083015250602061935b84828501619302565b602083015250604061936f84828501619302565b604083015250606061938384828501619302565b606083015250608061939784828501619302565b60808301525060a06193ab84828501619302565b60a08301525060c06193bf84828501619302565b60c08301525092915050565b600381106193d857600080fd5b50565b6000815190506193ea816193cb565b92915050565b6000815190506193ff81617cfe565b92915050565b60006101a0828403121561941c5761941b618138565b5b61942660e061819d565b90506000619436848285016192ed565b600083015250602061944a84828501619317565b60208301525061010061945f848285016193db565b604083015250610120619474848285016185ff565b60608301525061014061948984828501619302565b60808301525061016061949e848285016193f0565b60a0830152506101806194b3848285016193f0565b60c08301525092915050565b60006101a082840312156194d6576194d56178d0565b5b60006194e484828501619405565b91505092915050565b7f54494d4531000000000000000000000000000000000000000000000000000000600082015250565b6000619523600583618641565b915061952e826194ed565b602082019050919050565b6000602082019050818103600083015261955281619516565b9050919050565b7f4f6e6c792070726f7669646572206f7220746869732063616e2063616c6c0000600082015250565b600061958f601e83618641565b915061959a82619559565b602082019050919050565b600060208201905081810360008301526195be81619582565b9050919050565b60006195d082617a20565b91506195db83617a20565b92508282039050818111156195f3576195f2618aca565b5b92915050565b600060408201905061960e6000830185617e4d565b61961b6020830184617a2a565b9392505050565b60006060820190506196376000830186617c94565b6196446020830185617e4d565b6196516040830184617c11565b949350505050565b600060608201905061966e6000830186617c94565b61967b6020830185617e3e565b6196886040830184617e3e565b949350505050565b60008160001c9050919050565b600067ffffffffffffffff82169050919050565b60006196c46196bf83619690565b61969d565b9050919050565b60008160401c9050919050565b60006196eb6196e6836196cb565b61969d565b9050919050565b60008160801c9050919050565b600061971261970d836196f2565b61969d565b9050919050565b60008160c01c9050919050565b600061973961973483619719565b61969d565b9050919050565b6000815461974d816189c1565b6197578186617e76565b945060018216600081146197725760018114619788576197bb565b60ff1983168652811515602002860193506197bb565b61979185618b41565b60005b838110156197b357815481890152600182019150602081019050619794565b808801955050505b50505092915050565b600060ff82169050919050565b60006197e46197df83619690565b6197c4565b9050919050565b600060c083016000808401549050619802816196b1565b61980f6000870182617e5c565b50619819816196d8565b6198266020870182617e5c565b50619830816196ff565b61983d6040870182617e5c565b5061984781619726565b6198546060870182617e5c565b5060018401858303608087015261986b8382619740565b9250506002840154905061987e816197d1565b61988b60a0870182617f34565b50819250505092915050565b600060a0820190506198ac6000830188617e4d565b6198b96020830187617e4d565b6198c66040830186617c94565b6198d36060830185617c11565b81810360808301526198e581846197eb565b90509695505050505050565b60006198fc82617a20565b915061990783617a20565b925082820261991581617a20565b9150828204841483151761992c5761992b618aca565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061996d82617a20565b915061997883617a20565b92508261998857619987619933565b5b828204905092915050565b600061999e82617a20565b91506199a983617a20565b92508282019050808211156199c1576199c0618aca565b5b92915050565b60006199e26199dd6199d884617bfd565b617a99565b617a20565b9050919050565b6199f2816199c7565b82525050565b6000606082019050619a0d6000830186617c94565b619a1a6020830185617e4d565b619a2760408301846199e9565b949350505050565b7f4a42313000000000000000000000000000000000000000000000000000000000600082015250565b6000619a65600483618641565b9150619a7082619a2f565b602082019050919050565b60006020820190508181036000830152619a9481619a58565b9050919050565b7f43616e6e6f742064697669646520627920300000000000000000000000000000600082015250565b6000619ad1601283618641565b9150619adc82619a9b565b602082019050919050565b60006020820190508181036000830152619b0081619ac4565b9050919050565b6000819050919050565b6000619b1c82619b07565b9150619b2783619b07565b925082820190508281121560008312168382126000841215161715619b4f57619b4e618aca565b5b92915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000619b8b601f83618641565b9150619b9682619b55565b602082019050919050565b60006020820190508181036000830152619bba81619b7e565b9050919050565b7f7472616e73300000000000000000000000000000000000000000000000000000600082015250565b6000619bf7600683618641565b9150619c0282619bc1565b602082019050919050565b60006020820190508181036000830152619c2681619bea565b9050919050565b7f7472616e73310000000000000000000000000000000000000000000000000000600082015250565b6000619c63600683618641565b9150619c6e82619c2d565b602082019050919050565b60006020820190508181036000830152619c9281619c56565b9050919050565b7f7472616e73320000000000000000000000000000000000000000000000000000600082015250565b6000619ccf600683618641565b9150619cda82619c99565b602082019050919050565b60006020820190508181036000830152619cfe81619cc2565b9050919050565b7f7472616e73330000000000000000000000000000000000000000000000000000600082015250565b6000619d3b600683618641565b9150619d4682619d05565b602082019050919050565b60006020820190508181036000830152619d6a81619d2e565b9050919050565b7f7472616e73340000000000000000000000000000000000000000000000000000600082015250565b6000619da7600683618641565b9150619db282619d71565b602082019050919050565b60006020820190508181036000830152619dd681619d9a565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000619e1e601783619ddd565b9150619e2982619de8565b601782019050919050565b6000619e3f82617e6b565b619e498185619ddd565b9350619e59818560208601617e87565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000619e9b601183619ddd565b9150619ea682619e65565b601182019050919050565b6000619ebc82619e11565b9150619ec88285619e34565b9150619ed382619e8e565b9150619edf8284619e34565b91508190509392505050565b6000619ef682617e6b565b619f008185618641565b9350619f10818560208601617e87565b619f1981617eb1565b840191505092915050565b60006020820190508181036000830152619f3e8184619eeb565b905092915050565b7f53616665436173743a2076616c756520646f65736e27742066697420696e206160008201527f6e20696e74323536000000000000000000000000000000000000000000000000602082015250565b6000619fa2602883618641565b9150619fad82619f46565b604082019050919050565b60006020820190508181036000830152619fd181619f95565b9050919050565b7f53616665436173743a2076616c756520646f65736e27742066697420696e203160008201527f3238206269747300000000000000000000000000000000000000000000000000602082015250565b600061a034602783618641565b915061a03f82619fd8565b604082019050919050565b6000602082019050818103600083015261a0638161a027565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061a0a482617a20565b91506000820361a0b75761a0b6618aca565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b600061a0f8602083618641565b915061a1038261a0c2565b602082019050919050565b6000602082019050818103600083015261a1278161a0eb565b905091905056fea26469706673582212208472e9ab6939a06aeef230782150d821b3cd9a25af9ce8463c9459dd8aba2cca64736f6c63430008110033",
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

// HasTopUpJob is a free data retrieval call binding the contract method 0xc7070e2c.
//
// Solidity: function hasTopUpJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerCaller) HasTopUpJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "hasTopUpJob", _providerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasTopUpJob is a free data retrieval call binding the contract method 0xc7070e2c.
//
// Solidity: function hasTopUpJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerSession) HasTopUpJob(_providerAddr common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasTopUpJob(&_MetaScheduler.CallOpts, _providerAddr)
}

// HasTopUpJob is a free data retrieval call binding the contract method 0xc7070e2c.
//
// Solidity: function hasTopUpJob(address _providerAddr) view returns(bool)
func (_MetaScheduler *MetaSchedulerCallerSession) HasTopUpJob(_providerAddr common.Address) (bool, error) {
	return _MetaScheduler.Contract.HasTopUpJob(&_MetaScheduler.CallOpts, _providerAddr)
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
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string,uint8) definition, bool valid, (uint256,uint256,uint256,bool) cost, (uint256,uint256,uint256,uint256) time, bytes32 jobName, bool hasCancelRequest)
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
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string,uint8) definition, bool valid, (uint256,uint256,uint256,bool) cost, (uint256,uint256,uint256,uint256) time, bytes32 jobName, bool hasCancelRequest)
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
// Solidity: function jobs(bytes32 ) view returns(bytes32 jobId, uint8 status, address customerAddr, address providerAddr, (uint64,uint64,uint64,uint64,string,uint8) definition, bool valid, (uint256,uint256,uint256,bool) cost, (uint256,uint256,uint256,uint256) time, bytes32 jobName, bool hasCancelRequest)
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

// ProviderTopUpJobsQueues is a free data retrieval call binding the contract method 0xbc8f18fb.
//
// Solidity: function providerTopUpJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCaller) ProviderTopUpJobsQueues(opts *bind.CallOpts, arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _MetaScheduler.contract.Call(opts, &out, "providerTopUpJobsQueues", arg0)

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

// ProviderTopUpJobsQueues is a free data retrieval call binding the contract method 0xbc8f18fb.
//
// Solidity: function providerTopUpJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerSession) ProviderTopUpJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderTopUpJobsQueues(&_MetaScheduler.CallOpts, arg0)
}

// ProviderTopUpJobsQueues is a free data retrieval call binding the contract method 0xbc8f18fb.
//
// Solidity: function providerTopUpJobsQueues(address ) view returns(int128 _begin, int128 _end)
func (_MetaScheduler *MetaSchedulerCallerSession) ProviderTopUpJobsQueues(arg0 common.Address) (struct {
	Begin *big.Int
	End   *big.Int
}, error) {
	return _MetaScheduler.Contract.ProviderTopUpJobsQueues(&_MetaScheduler.CallOpts, arg0)
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
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _lockedCredits, bytes32 _jobName, bool _delegateSpendingAuthority) returns(bytes32)
func (_MetaScheduler *MetaSchedulerTransactor) RequestNewJob(opts *bind.TransactOpts, _definition JobDefinition, _lockedCredits *big.Int, _jobName [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.contract.Transact(opts, "requestNewJob", _definition, _lockedCredits, _jobName, _delegateSpendingAuthority)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _lockedCredits, bytes32 _jobName, bool _delegateSpendingAuthority) returns(bytes32)
func (_MetaScheduler *MetaSchedulerSession) RequestNewJob(_definition JobDefinition, _lockedCredits *big.Int, _jobName [32]byte, _delegateSpendingAuthority bool) (*types.Transaction, error) {
	return _MetaScheduler.Contract.RequestNewJob(&_MetaScheduler.TransactOpts, _definition, _lockedCredits, _jobName, _delegateSpendingAuthority)
}

// RequestNewJob is a paid mutator transaction binding the contract method 0x3c2fb3da.
//
// Solidity: function requestNewJob((uint64,uint64,uint64,uint64,string,uint8) _definition, uint256 _lockedCredits, bytes32 _jobName, bool _delegateSpendingAuthority) returns(bytes32)
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
	Bin: "0x6101cd610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806315a945ce1461003a575b600080fd5b610054600480360381019061004f91906100be565b61006a565b6040516100619190610119565b60405180910390f35b60008183436100799190610163565b1015905092915050565b600080fd5b6000819050919050565b61009b81610088565b81146100a657600080fd5b50565b6000813590506100b881610092565b92915050565b600080604083850312156100d5576100d4610083565b5b60006100e3858286016100a9565b92505060206100f4858286016100a9565b9150509250929050565b60008115159050919050565b610113816100fe565b82525050565b600060208201905061012e600083018461010a565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061016e82610088565b915061017983610088565b925082820390508181111561019157610190610134565b5b9291505056fea26469706673582212202f54b5e7bd26299a0b11bfe60ce232e73f930804f294d4d57183c6dabf53289c64736f6c63430008110033",
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
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220391d8dec75077b928a0396b6517eaa811bad2ae0354dc8e64363ae5e3594158664736f6c63430008110033",
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
