// Package deepsquare defines APIs for interacting with the DeepSquare Grid.
package metascheduler

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNewRequestJobNotFound = errors.New("new request job event not found")
)

var (
	metaschedulerABI   *abi.ABI
	ierc20ABI          *abi.ABI
	newJobRequestEvent abi.Event
	jobTransitionEvent abi.Event
	transferEvent      abi.Event
	approvalEvent      abi.Event
	defaultChainID     = big.NewInt(179188)
)

func init() {
	var err error
	metaschedulerABI, err = metaschedulerabi.MetaSchedulerMetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to parse metascheduler contract ABI: %w", err))
	}
	ierc20ABI, err = metaschedulerabi.IERC20MetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to parse erc20 contract ABI: %w", err))
	}

	// Find the event signature dynamically
	var ok bool
	newJobRequestEvent, ok = metaschedulerABI.Events["NewJobRequestEvent"]
	if !ok {
		panic(fmt.Errorf("failed to get NewJobRequestEvent: %w", err))
	}

	jobTransitionEvent, ok = metaschedulerABI.Events["JobTransitionEvent"]
	if !ok {
		panic(fmt.Errorf("failed to get JobTransitionEvent: %w", err))
	}

	transferEvent, ok = ierc20ABI.Events["Transfer"]
	if !ok {
		panic(fmt.Errorf("failed to get Transfer: %w", err))
	}
	approvalEvent, ok = ierc20ABI.Events["Approval"]
	if !ok {
		panic(fmt.Errorf("failed to get Approval: %w", err))
	}
}

type EthereumBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

// Backend is a wrapper around the EthereumBackend used to interact with the Meta-Scheduler smart-contract.
type Backend struct {
	// EthereumBackend is the Ethereum Client.
	//
	// TODO: check if websocket or rpc.
	EthereumBackend
	// Address of the metascheduler smart-contract.
	MetaschedulerAddress common.Address
	// ChainID of the blockchain.
	ChainID *big.Int
	// PrivateKey of the user.
	UserPrivateKey *ecdsa.PrivateKey
}

// NewJobScheduler creates a JobScheduler.
func NewJobScheduler(b Backend, sbatch sbatch.Service) (client types.JobScheduler, err error) {
	b = b.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(b.MetaschedulerAddress, b.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Backend:       b,
		sbatch:        sbatch,
	}, err
}

func (b Backend) applyDefault() Backend {
	if b.ChainID == nil {
		b.ChainID = defaultChainID
	}
	return b
}

func NewJobFetcher(b Backend) (fetcher types.JobFetcher, err error) {
	b = b.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(b.MetaschedulerAddress, b.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Backend:       b,
	}, err
}

func NewEventSubscriber(b Backend) (types.EventSubscriber, error) {
	b = b.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(b.MetaschedulerAddress, b.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &wsClient{
		MetaScheduler: m,
		Backend:       b,
	}, err
}

func NewJobFilterer(b Backend) (watcher types.JobFilterer, err error) {
	b = b.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(b.MetaschedulerAddress, b.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &wsClient{
		MetaScheduler: m,
		Backend:       b,
	}, err
}

func NewCreditManager(ctx context.Context, b Backend) (credits types.CreditManager, err error) {
	b = b.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(b.MetaschedulerAddress, b.EthereumBackend)
	if err != nil {
		return nil, err
	}
	creditAddress, err := m.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, b)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Backend:       b,
		IERC20:        ierc20,
	}, err
}

func NewCreditFilterer(
	ctx context.Context,
	ws Backend,
	credit types.CreditManager,
) (watcher types.CreditFilterer, err error) {
	ws = ws.applyDefault()
	mWs, err := metaschedulerabi.NewMetaScheduler(ws.MetaschedulerAddress, ws.EthereumBackend)
	if err != nil {
		return nil, err
	}
	wsClient := wsClient{
		MetaScheduler: mWs,
		Backend:       ws,
	}
	creditAddress, err := mWs.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	ierc20Filterer, err := metaschedulerabi.NewIERC20Filterer(creditAddress, ws)
	if err != nil {
		return nil, err
	}
	return &creditFilterer{
		CreditManager:  credit,
		wsClient:       wsClient,
		IERC20Filterer: ierc20Filterer,
	}, err
}

func NewAllowanceManager(
	ctx context.Context,
	b Backend,
) (allowance types.AllowanceManager, err error) {
	b = b.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(b.MetaschedulerAddress, b.EthereumBackend)
	if err != nil {
		return nil, err
	}
	creditAddress, err := m.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, b)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Backend:       b,
		IERC20:        ierc20,
	}, err
}

func NewAllowanceFilterer(
	ctx context.Context,
	ws Backend,
	allowance types.AllowanceManager,
) (watcher types.AllowanceFilterer, err error) {
	ws = ws.applyDefault()
	mWs, err := metaschedulerabi.NewMetaScheduler(ws.MetaschedulerAddress, ws.EthereumBackend)
	if err != nil {
		return nil, err
	}
	wsClient := wsClient{
		MetaScheduler: mWs,
		Backend:       ws,
	}
	creditAddress, err := mWs.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	ierc20Filterer, err := metaschedulerabi.NewIERC20Filterer(creditAddress, ws)
	if err != nil {
		return nil, err
	}
	return &allowanceFilterer{
		AllowanceManager: allowance,
		wsClient:         wsClient,
		IERC20Filterer:   ierc20Filterer,
	}, err
}
