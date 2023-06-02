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

// DoubleEndedQueueMetaData contains all meta data concerning the DoubleEndedQueue contract.
var DoubleEndedQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202032d65aab7c0771139195a236ce214264c0f6d13b16e771b658e1b449bbd21664736f6c63430008110033",
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

// ErrorContractMetaData contains all meta data concerning the ErrorContract contract.
var ErrorContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Banned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreditAddrIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CustomerMetaSchedulerProviderOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"CustomerOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DivisionByZeroError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJobDefinition\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"InvalidNCpu\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"InvalidNMem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"InvalidNNodes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTransitionFromMetascheduled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTransitionFromPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTransitionFromRunning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTransitionFromScheduled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"JobColdStatusOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"JobHotStatusOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"JobProviderOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"JobProviderThisOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"MetaScheduledScheduledStatusOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MetashedulerProviderOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoJob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpendingAuthority\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"OwnerOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderAddrIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNotJoined\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"RemainingTimeAboveLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"RunningColdStatusOnly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"RunningScheduledStatusOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameStatusError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Uninitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WaitingApprovalOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThrowBanned\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowCreditAddrIsZero\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowCustomerMetaSchedulerProviderOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"ThrowCustomerOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowDivisionByZeroError\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowEmpty\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"ThrowInsufficientFunds\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidJob\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidJobDefinition\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"ThrowInvalidNCpu\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"ThrowInvalidNMem\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"}],\"name\":\"ThrowInvalidNNodes\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidTransitionFromMetascheduled\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidTransitionFromPending\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidTransitionFromRunning\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowInvalidTransitionFromScheduled\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowJobColdStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowJobHotStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"ThrowJobProviderOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"ThrowJobProviderThisOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowMetaScheduledScheduledStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowMetashedulerProviderOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowNoJob\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowNoProvider\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowNoSpendingAuthority\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowOutOfBounds\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"ThrowOwnerOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowProviderAddrIsZero\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowProviderNotJoined\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"ThrowRemainingTimeAboveLimit\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowRunningColdStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumJobStatus\",\"name\":\"current\",\"type\":\"uint8\"}],\"name\":\"ThrowRunningScheduledStatusOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowSameStatusError\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowUninitialized\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ThrowWaitingApprovalOnly\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506108be806100206000396000f3fe608060405234801561001057600080fd5b50600436106102105760003560e01c8063693a229711610125578063be7c7510116100ad578063d0980b701161007c578063d0980b701461038b578063db32d27d14610393578063de736e40146103a6578063f2f22fab146103ae578063f522f385146103c157600080fd5b8063be7c751014610355578063c33400ae1461035d578063c8ba784514610370578063cbfbbf8a1461037857600080fd5b806393514ad5116100f457806393514ad51461030c57806398f0bb4e146103145780639c46b0061461031c578063a6b419581461032f578063bbe7c6071461034257600080fd5b8063693a2297146102e1578063818902a7146102e95780638d578081146102f157806392234b11146102f957600080fd5b806346f87e01116101a85780635acf443e116101775780635acf443e146102b95780635ad71c45146102c15780635c71386a146102c957806361cb8370146102d157806361fdbb96146102d957600080fd5b806346f87e011461028e57806356d7498214610296578063573cc84f1461029e57806358b28b69146102a657600080fd5b806319e396dd116101e457806319e396dd14610258578063310891461461026b5780633e3c2769146102735780634094eedf1461028657600080fd5b80628d98001461021557806301d72a591461022a57806310efeaea1461023d57806317d0a33f14610245575b600080fd5b6102286102233660046107ae565b6103c9565b005b6102286102383660046107c7565b6103ea565b61022861040d565b6102286102533660046107e9565b610426565b6102286102663660046107ae565b610441565b61022861045d565b6102286102813660046107e9565b610479565b610228610494565b6102286104ad565b6102286104c6565b6102286104df565b6102286102b43660046107e9565b6104f8565b610228610513565b61022861052c565b610228610545565b61022861055e565b610228610577565b610228610590565b6102286105a9565b6102286105c2565b6102286103073660046107e9565b6105db565b6102286105f6565b61022861060f565b61022861032a3660046107c7565b610628565b61022861033d36600461082d565b61064a565b61022861035036600461082d565b610676565b6102286106a2565b61022861036b36600461082d565b6106bb565b6102286106e7565b61022861038636600461082d565b610700565b61022861072c565b6102286103a13660046107ae565b610745565b610228610761565b6102286103bc3660046107e9565b61077a565b610228610795565b60405163e5c0669560e01b8152600481018290526024015b60405180910390fd5b604051634801db4560e11b815260048101839052602481018290526044016103e1565b60405163b63bcc0960e01b815260040160405180910390fd5b8060405163ed3e2aad60e01b81526004016103e19190610860565b6040516352357b3f60e11b8152600481018290526024016103e1565b604051600162031a2960e01b0319815260040160405180910390fd5b80604051634634126160e11b81526004016103e19190610860565b604051630d208e5960e41b815260040160405180910390fd5b604051630f6a35a360e41b815260040160405180910390fd5b60405163038e47a360e51b815260040160405180910390fd5b604051632a9126eb60e01b815260040160405180910390fd5b80604051631024fb8560e31b81526004016103e19190610860565b604051630317892360e11b815260040160405180910390fd5b604051637064572b60e01b815260040160405180910390fd5b60405163ef341f6d60e01b815260040160405180910390fd5b604051632a856fc960e01b815260040160405180910390fd5b604051639773692760e01b815260040160405180910390fd5b6040516342f9901960e01b815260040160405180910390fd5b604051637b1afa6760e11b815260040160405180910390fd5b6040516330d0592960e01b815260040160405180910390fd5b806040516314475eb760e01b81526004016103e19190610860565b604051637897ef6d60e01b815260040160405180910390fd5b60405163071cbeb560e21b815260040160405180910390fd5b60405162fae2d560e21b815260048101839052602481018290526044016103e1565b604051630cb8c19760e21b81526001600160a01b038084166004830152821660248201526044016103e1565b604051633efce78d60e01b81526001600160a01b038084166004830152821660248201526044016103e1565b6040516329e460df60e21b815260040160405180910390fd5b60405163521eb56d60e11b81526001600160a01b038084166004830152821660248201526044016103e1565b604051631ed9509560e11b815260040160405180910390fd5b604051638942331960e01b81526001600160a01b038084166004830152821660248201526044016103e1565b604051633a43ca4160e01b815260040160405180910390fd5b604051631b8e63e960e31b8152600481018290526024016103e1565b604051633a48939560e01b815260040160405180910390fd5b8060405163048389ff60e11b81526004016103e19190610860565b604051632d0483c560e21b815260040160405180910390fd5b6000602082840312156107c057600080fd5b5035919050565b600080604083850312156107da57600080fd5b50508035926020909101359150565b6000602082840312156107fb57600080fd5b81356008811061080a57600080fd5b9392505050565b80356001600160a01b038116811461082857600080fd5b919050565b6000806040838503121561084057600080fd5b61084983610811565b915061085760208401610811565b90509250929050565b602081016008831061088257634e487b7160e01b600052602160045260246000fd5b9190529056fea2646970667358221220e54a82efe85da0dba76c5f9dacd5612271b431d5922dd21aa9ea5a17eea769a264736f6c63430008110033",
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

// ThrowCreditAddrIsZero is a free data retrieval call binding the contract method 0xde736e40.
//
// Solidity: function ThrowCreditAddrIsZero() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowCreditAddrIsZero(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowCreditAddrIsZero")

	if err != nil {
		return err
	}

	return err

}

// ThrowCreditAddrIsZero is a free data retrieval call binding the contract method 0xde736e40.
//
// Solidity: function ThrowCreditAddrIsZero() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowCreditAddrIsZero() error {
	return _ErrorContract.Contract.ThrowCreditAddrIsZero(&_ErrorContract.CallOpts)
}

// ThrowCreditAddrIsZero is a free data retrieval call binding the contract method 0xde736e40.
//
// Solidity: function ThrowCreditAddrIsZero() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowCreditAddrIsZero() error {
	return _ErrorContract.Contract.ThrowCreditAddrIsZero(&_ErrorContract.CallOpts)
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

// ThrowDivisionByZeroError is a free data retrieval call binding the contract method 0xbe7c7510.
//
// Solidity: function ThrowDivisionByZeroError() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowDivisionByZeroError(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowDivisionByZeroError")

	if err != nil {
		return err
	}

	return err

}

// ThrowDivisionByZeroError is a free data retrieval call binding the contract method 0xbe7c7510.
//
// Solidity: function ThrowDivisionByZeroError() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowDivisionByZeroError() error {
	return _ErrorContract.Contract.ThrowDivisionByZeroError(&_ErrorContract.CallOpts)
}

// ThrowDivisionByZeroError is a free data retrieval call binding the contract method 0xbe7c7510.
//
// Solidity: function ThrowDivisionByZeroError() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowDivisionByZeroError() error {
	return _ErrorContract.Contract.ThrowDivisionByZeroError(&_ErrorContract.CallOpts)
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

// ThrowInvalidNCpu is a free data retrieval call binding the contract method 0x19e396dd.
//
// Solidity: function ThrowInvalidNCpu(uint256 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidNCpu(opts *bind.CallOpts, current *big.Int) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidNCpu", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidNCpu is a free data retrieval call binding the contract method 0x19e396dd.
//
// Solidity: function ThrowInvalidNCpu(uint256 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidNCpu(current *big.Int) error {
	return _ErrorContract.Contract.ThrowInvalidNCpu(&_ErrorContract.CallOpts, current)
}

// ThrowInvalidNCpu is a free data retrieval call binding the contract method 0x19e396dd.
//
// Solidity: function ThrowInvalidNCpu(uint256 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidNCpu(current *big.Int) error {
	return _ErrorContract.Contract.ThrowInvalidNCpu(&_ErrorContract.CallOpts, current)
}

// ThrowInvalidNMem is a free data retrieval call binding the contract method 0x008d9800.
//
// Solidity: function ThrowInvalidNMem(uint256 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidNMem(opts *bind.CallOpts, current *big.Int) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidNMem", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidNMem is a free data retrieval call binding the contract method 0x008d9800.
//
// Solidity: function ThrowInvalidNMem(uint256 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidNMem(current *big.Int) error {
	return _ErrorContract.Contract.ThrowInvalidNMem(&_ErrorContract.CallOpts, current)
}

// ThrowInvalidNMem is a free data retrieval call binding the contract method 0x008d9800.
//
// Solidity: function ThrowInvalidNMem(uint256 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidNMem(current *big.Int) error {
	return _ErrorContract.Contract.ThrowInvalidNMem(&_ErrorContract.CallOpts, current)
}

// ThrowInvalidNNodes is a free data retrieval call binding the contract method 0xdb32d27d.
//
// Solidity: function ThrowInvalidNNodes(uint256 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidNNodes(opts *bind.CallOpts, current *big.Int) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidNNodes", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidNNodes is a free data retrieval call binding the contract method 0xdb32d27d.
//
// Solidity: function ThrowInvalidNNodes(uint256 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidNNodes(current *big.Int) error {
	return _ErrorContract.Contract.ThrowInvalidNNodes(&_ErrorContract.CallOpts, current)
}

// ThrowInvalidNNodes is a free data retrieval call binding the contract method 0xdb32d27d.
//
// Solidity: function ThrowInvalidNNodes(uint256 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidNNodes(current *big.Int) error {
	return _ErrorContract.Contract.ThrowInvalidNNodes(&_ErrorContract.CallOpts, current)
}

// ThrowInvalidTransitionFromMetascheduled is a free data retrieval call binding the contract method 0x31089146.
//
// Solidity: function ThrowInvalidTransitionFromMetascheduled() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidTransitionFromMetascheduled(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidTransitionFromMetascheduled")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidTransitionFromMetascheduled is a free data retrieval call binding the contract method 0x31089146.
//
// Solidity: function ThrowInvalidTransitionFromMetascheduled() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidTransitionFromMetascheduled() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromMetascheduled(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransitionFromMetascheduled is a free data retrieval call binding the contract method 0x31089146.
//
// Solidity: function ThrowInvalidTransitionFromMetascheduled() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidTransitionFromMetascheduled() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromMetascheduled(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransitionFromPending is a free data retrieval call binding the contract method 0x818902a7.
//
// Solidity: function ThrowInvalidTransitionFromPending() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidTransitionFromPending(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidTransitionFromPending")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidTransitionFromPending is a free data retrieval call binding the contract method 0x818902a7.
//
// Solidity: function ThrowInvalidTransitionFromPending() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidTransitionFromPending() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromPending(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransitionFromPending is a free data retrieval call binding the contract method 0x818902a7.
//
// Solidity: function ThrowInvalidTransitionFromPending() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidTransitionFromPending() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromPending(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransitionFromRunning is a free data retrieval call binding the contract method 0x8d578081.
//
// Solidity: function ThrowInvalidTransitionFromRunning() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidTransitionFromRunning(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidTransitionFromRunning")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidTransitionFromRunning is a free data retrieval call binding the contract method 0x8d578081.
//
// Solidity: function ThrowInvalidTransitionFromRunning() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidTransitionFromRunning() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromRunning(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransitionFromRunning is a free data retrieval call binding the contract method 0x8d578081.
//
// Solidity: function ThrowInvalidTransitionFromRunning() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidTransitionFromRunning() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromRunning(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransitionFromScheduled is a free data retrieval call binding the contract method 0x10efeaea.
//
// Solidity: function ThrowInvalidTransitionFromScheduled() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowInvalidTransitionFromScheduled(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowInvalidTransitionFromScheduled")

	if err != nil {
		return err
	}

	return err

}

// ThrowInvalidTransitionFromScheduled is a free data retrieval call binding the contract method 0x10efeaea.
//
// Solidity: function ThrowInvalidTransitionFromScheduled() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowInvalidTransitionFromScheduled() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromScheduled(&_ErrorContract.CallOpts)
}

// ThrowInvalidTransitionFromScheduled is a free data retrieval call binding the contract method 0x10efeaea.
//
// Solidity: function ThrowInvalidTransitionFromScheduled() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowInvalidTransitionFromScheduled() error {
	return _ErrorContract.Contract.ThrowInvalidTransitionFromScheduled(&_ErrorContract.CallOpts)
}

// ThrowJobColdStatusOnly is a free data retrieval call binding the contract method 0x58b28b69.
//
// Solidity: function ThrowJobColdStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowJobColdStatusOnly(opts *bind.CallOpts, current uint8) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowJobColdStatusOnly", current)

	if err != nil {
		return err
	}

	return err

}

// ThrowJobColdStatusOnly is a free data retrieval call binding the contract method 0x58b28b69.
//
// Solidity: function ThrowJobColdStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowJobColdStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowJobColdStatusOnly(&_ErrorContract.CallOpts, current)
}

// ThrowJobColdStatusOnly is a free data retrieval call binding the contract method 0x58b28b69.
//
// Solidity: function ThrowJobColdStatusOnly(uint8 current) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowJobColdStatusOnly(current uint8) error {
	return _ErrorContract.Contract.ThrowJobColdStatusOnly(&_ErrorContract.CallOpts, current)
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

// ThrowJobProviderThisOnly is a free data retrieval call binding the contract method 0xbbe7c607.
//
// Solidity: function ThrowJobProviderThisOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowJobProviderThisOnly(opts *bind.CallOpts, current common.Address, expected common.Address) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowJobProviderThisOnly", current, expected)

	if err != nil {
		return err
	}

	return err

}

// ThrowJobProviderThisOnly is a free data retrieval call binding the contract method 0xbbe7c607.
//
// Solidity: function ThrowJobProviderThisOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowJobProviderThisOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowJobProviderThisOnly(&_ErrorContract.CallOpts, current, expected)
}

// ThrowJobProviderThisOnly is a free data retrieval call binding the contract method 0xbbe7c607.
//
// Solidity: function ThrowJobProviderThisOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowJobProviderThisOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowJobProviderThisOnly(&_ErrorContract.CallOpts, current, expected)
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

// ThrowMetashedulerProviderOnly is a free data retrieval call binding the contract method 0x5acf443e.
//
// Solidity: function ThrowMetashedulerProviderOnly() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowMetashedulerProviderOnly(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowMetashedulerProviderOnly")

	if err != nil {
		return err
	}

	return err

}

// ThrowMetashedulerProviderOnly is a free data retrieval call binding the contract method 0x5acf443e.
//
// Solidity: function ThrowMetashedulerProviderOnly() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowMetashedulerProviderOnly() error {
	return _ErrorContract.Contract.ThrowMetashedulerProviderOnly(&_ErrorContract.CallOpts)
}

// ThrowMetashedulerProviderOnly is a free data retrieval call binding the contract method 0x5acf443e.
//
// Solidity: function ThrowMetashedulerProviderOnly() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowMetashedulerProviderOnly() error {
	return _ErrorContract.Contract.ThrowMetashedulerProviderOnly(&_ErrorContract.CallOpts)
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

// ThrowOwnerOnly is a free data retrieval call binding the contract method 0xc33400ae.
//
// Solidity: function ThrowOwnerOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowOwnerOnly(opts *bind.CallOpts, current common.Address, expected common.Address) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowOwnerOnly", current, expected)

	if err != nil {
		return err
	}

	return err

}

// ThrowOwnerOnly is a free data retrieval call binding the contract method 0xc33400ae.
//
// Solidity: function ThrowOwnerOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractSession) ThrowOwnerOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowOwnerOnly(&_ErrorContract.CallOpts, current, expected)
}

// ThrowOwnerOnly is a free data retrieval call binding the contract method 0xc33400ae.
//
// Solidity: function ThrowOwnerOnly(address current, address expected) pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowOwnerOnly(current common.Address, expected common.Address) error {
	return _ErrorContract.Contract.ThrowOwnerOnly(&_ErrorContract.CallOpts, current, expected)
}

// ThrowProviderAddrIsZero is a free data retrieval call binding the contract method 0x46f87e01.
//
// Solidity: function ThrowProviderAddrIsZero() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowProviderAddrIsZero(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowProviderAddrIsZero")

	if err != nil {
		return err
	}

	return err

}

// ThrowProviderAddrIsZero is a free data retrieval call binding the contract method 0x46f87e01.
//
// Solidity: function ThrowProviderAddrIsZero() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowProviderAddrIsZero() error {
	return _ErrorContract.Contract.ThrowProviderAddrIsZero(&_ErrorContract.CallOpts)
}

// ThrowProviderAddrIsZero is a free data retrieval call binding the contract method 0x46f87e01.
//
// Solidity: function ThrowProviderAddrIsZero() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowProviderAddrIsZero() error {
	return _ErrorContract.Contract.ThrowProviderAddrIsZero(&_ErrorContract.CallOpts)
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

// ThrowUninitialized is a free data retrieval call binding the contract method 0x98f0bb4e.
//
// Solidity: function ThrowUninitialized() pure returns()
func (_ErrorContract *ErrorContractCaller) ThrowUninitialized(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ErrorContract.contract.Call(opts, &out, "ThrowUninitialized")

	if err != nil {
		return err
	}

	return err

}

// ThrowUninitialized is a free data retrieval call binding the contract method 0x98f0bb4e.
//
// Solidity: function ThrowUninitialized() pure returns()
func (_ErrorContract *ErrorContractSession) ThrowUninitialized() error {
	return _ErrorContract.Contract.ThrowUninitialized(&_ErrorContract.CallOpts)
}

// ThrowUninitialized is a free data retrieval call binding the contract method 0x98f0bb4e.
//
// Solidity: function ThrowUninitialized() pure returns()
func (_ErrorContract *ErrorContractCallerSession) ThrowUninitialized() error {
	return _ErrorContract.Contract.ThrowUninitialized(&_ErrorContract.CallOpts)
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
