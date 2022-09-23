// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package grid

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

// GridMetaData contains all meta data concerning the Grid contract.
var GridMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212206c1bc47b3d2850012833f02283c062f5c663e202b52ba07600ab93fa49e8ad2964736f6c63430008110033",
}

// GridABI is the input ABI used to generate the binding from.
// Deprecated: Use GridMetaData.ABI instead.
var GridABI = GridMetaData.ABI

// GridBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GridMetaData.Bin instead.
var GridBin = GridMetaData.Bin

// DeployGrid deploys a new Ethereum contract, binding an instance of Grid to it.
func DeployGrid(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Grid, error) {
	parsed, err := GridMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GridBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Grid{GridCaller: GridCaller{contract: contract}, GridTransactor: GridTransactor{contract: contract}, GridFilterer: GridFilterer{contract: contract}}, nil
}

// Grid is an auto generated Go binding around an Ethereum contract.
type Grid struct {
	GridCaller     // Read-only binding to the contract
	GridTransactor // Write-only binding to the contract
	GridFilterer   // Log filterer for contract events
}

// GridCaller is an auto generated read-only Go binding around an Ethereum contract.
type GridCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GridTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GridTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GridFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GridFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GridSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GridSession struct {
	Contract     *Grid             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GridCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GridCallerSession struct {
	Contract *GridCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GridTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GridTransactorSession struct {
	Contract     *GridTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GridRaw is an auto generated low-level Go binding around an Ethereum contract.
type GridRaw struct {
	Contract *Grid // Generic contract binding to access the raw methods on
}

// GridCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GridCallerRaw struct {
	Contract *GridCaller // Generic read-only contract binding to access the raw methods on
}

// GridTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GridTransactorRaw struct {
	Contract *GridTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGrid creates a new instance of Grid, bound to a specific deployed contract.
func NewGrid(address common.Address, backend bind.ContractBackend) (*Grid, error) {
	contract, err := bindGrid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Grid{GridCaller: GridCaller{contract: contract}, GridTransactor: GridTransactor{contract: contract}, GridFilterer: GridFilterer{contract: contract}}, nil
}

// NewGridCaller creates a new read-only instance of Grid, bound to a specific deployed contract.
func NewGridCaller(address common.Address, caller bind.ContractCaller) (*GridCaller, error) {
	contract, err := bindGrid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GridCaller{contract: contract}, nil
}

// NewGridTransactor creates a new write-only instance of Grid, bound to a specific deployed contract.
func NewGridTransactor(address common.Address, transactor bind.ContractTransactor) (*GridTransactor, error) {
	contract, err := bindGrid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GridTransactor{contract: contract}, nil
}

// NewGridFilterer creates a new log filterer instance of Grid, bound to a specific deployed contract.
func NewGridFilterer(address common.Address, filterer bind.ContractFilterer) (*GridFilterer, error) {
	contract, err := bindGrid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GridFilterer{contract: contract}, nil
}

// bindGrid binds a generic wrapper to an already deployed contract.
func bindGrid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GridABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Grid *GridRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Grid.Contract.GridCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Grid *GridRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Grid.Contract.GridTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Grid *GridRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Grid.Contract.GridTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Grid *GridCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Grid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Grid *GridTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Grid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Grid *GridTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Grid.Contract.contract.Transact(opts, method, params...)
}
