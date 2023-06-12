// Package deepsquare defines APIs for interacting with the DeepSquare Grid.
package metascheduler

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/deepsquare-io/the-grid/cli"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNewRequestJobNotFound = errors.New("new request job event not found")
)

var (
	MetaschedulerABI   *abi.ABI
	IERC20ABI          *abi.ABI
	newJobRequestEvent abi.Event
	jobTransitionEvent abi.Event
	transferEvent      abi.Event
	approvalEvent      abi.Event
	defaultChainID     = big.NewInt(179188)
)

func init() {
	var err error
	MetaschedulerABI, err = metaschedulerabi.MetaSchedulerMetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to parse metascheduler contract ABI: %w", err))
	}
	IERC20ABI, err = metaschedulerabi.IERC20MetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to parse erc20 contract ABI: %w", err))
	}

	// Find the event signature dynamically
	var ok bool
	newJobRequestEvent, ok = MetaschedulerABI.Events["NewJobRequestEvent"]
	if !ok {
		panic(fmt.Errorf("failed to get NewJobRequestEvent: %w", err))
	}

	jobTransitionEvent, ok = MetaschedulerABI.Events["JobTransitionEvent"]
	if !ok {
		panic(fmt.Errorf("failed to get JobTransitionEvent: %w", err))
	}

	transferEvent, ok = IERC20ABI.Events["Transfer"]
	if !ok {
		panic(fmt.Errorf("failed to get Transfer: %w", err))
	}
	approvalEvent, ok = IERC20ABI.Events["Approval"]
	if !ok {
		panic(fmt.Errorf("failed to get Approval: %w", err))
	}
}

type EthereumBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

type Client struct {
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

func NewJobScheduler(c Client, sbatch sbatch.Service) (client cli.JobScheduler, err error) {
	c = c.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Client:        c,
		sbatch:        sbatch,
	}, err
}

func (c Client) applyDefault() Client {
	if c.ChainID == nil {
		c.ChainID = defaultChainID
	}
	return c
}

func NewJobFetcher(c Client) (fetcher cli.JobFetcher, err error) {
	c = c.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Client:        c,
	}, err
}

func NewEventSubscriber(c Client) (cli.EventSubscriber, error) {
	c = c.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &wsClient{
		MetaScheduler: m,
		Client:        c,
	}, err
}

func NewJobFilterer(c Client) (watcher cli.JobFilterer, err error) {
	c = c.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c.EthereumBackend)
	if err != nil {
		return nil, err
	}
	return &wsClient{
		MetaScheduler: m,
		Client:        c,
	}, err
}

func NewCreditManager(ctx context.Context, c Client) (credits cli.CreditManager, err error) {
	c = c.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c.EthereumBackend)
	if err != nil {
		return nil, err
	}
	creditAddress, err := m.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, c)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Client:        c,
		IERC20:        ierc20,
	}, err
}

func NewCreditFilterer(
	ctx context.Context,
	ws Client,
	credit cli.CreditManager,
) (watcher cli.CreditFilterer, err error) {
	ws = ws.applyDefault()
	mWs, err := metaschedulerabi.NewMetaScheduler(ws.MetaschedulerAddress, ws.EthereumBackend)
	if err != nil {
		return nil, err
	}
	wsClient := wsClient{
		MetaScheduler: mWs,
		Client:        ws,
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
	c Client,
) (allowance cli.AllowanceManager, err error) {
	c = c.applyDefault()
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c.EthereumBackend)
	if err != nil {
		return nil, err
	}
	creditAddress, err := m.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, c)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler: m,
		Client:        c,
		IERC20:        ierc20,
	}, err
}

func NewAllowanceFilterer(
	ctx context.Context,
	ws Client,
	allowance cli.AllowanceManager,
) (watcher cli.AllowanceFilterer, err error) {
	ws = ws.applyDefault()
	mWs, err := metaschedulerabi.NewMetaScheduler(ws.MetaschedulerAddress, ws.EthereumBackend)
	if err != nil {
		return nil, err
	}
	wsClient := wsClient{
		MetaScheduler: mWs,
		Client:        ws,
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
