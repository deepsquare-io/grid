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
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

func (c Client) JobScheduler(sbatch sbatch.Service) (client cli.JobScheduler, err error) {
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

type rpcClient struct {
	*metaschedulerabi.MetaScheduler
	*metaschedulerabi.IERC20
	Client
	sbatch sbatch.Service
}

func (c *rpcClient) from() (addr common.Address) {
	if c.UserPrivateKey == nil {
		return addr
	}
	return crypto.PubkeyToAddress(c.UserPrivateKey.PublicKey)
}

// authOpts generate transact options based on the network.
func (c *rpcClient) authOpts(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.PendingNonceAt(ctx, c.from())
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.UserPrivateKey, c.ChainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(0)
	auth.GasPrice = gasPrice
	auth.Context = ctx

	return auth, nil
}

func (c *rpcClient) SetAllowance(ctx context.Context, amount *big.Int) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed get auth options: %w", err)
	}
	tx, err := c.Approve(opts, c.MetaschedulerAddress, amount)
	if err != nil {
		return fmt.Errorf("failed to approve credit: %w", err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
}

func (c *rpcClient) ClearAllowance(ctx context.Context) error {
	return c.SetAllowance(ctx, big.NewInt(0))
}

func (c *rpcClient) GetAllowance(ctx context.Context) (*big.Int, error) {
	return c.Allowance(&bind.CallOpts{
		Context: ctx,
	}, c.from(), c.MetaschedulerAddress)
}

func (c *rpcClient) GetJob(ctx context.Context, id [32]byte) (*cli.Job, error) {
	job, err := c.Jobs(&bind.CallOpts{
		Context: ctx,
	}, id)
	if err != nil {
		return nil, err
	}
	return &cli.Job{
		JobID:            job.JobId,
		Status:           job.Status,
		CustomerAddr:     job.CustomerAddr,
		ProviderAddr:     job.ProviderAddr,
		Definition:       job.Definition,
		Valid:            job.Valid,
		Cost:             job.Cost,
		Time:             job.Time,
		JobName:          job.JobName,
		HasCancelRequest: job.HasCancelRequest,
	}, err
}

type jobIterator struct {
	*rpcClient
	array  [][32]byte
	length int
	index  int
	job    *cli.Job
}

func (it *jobIterator) Next(
	ctx context.Context,
) (next cli.JobLazyIterator, ok bool, err error) {
	if it.index+1 >= it.length {
		return nil, false, nil
	}
	job, err := it.GetJob(ctx, it.array[it.index+1])
	if err != nil {
		return nil, false, fmt.Errorf("failed to get job: %w", err)
	}

	return &jobIterator{
		rpcClient: it.rpcClient,
		array:     it.array,
		length:    it.length,
		index:     it.index + 1,
		job:       job,
	}, true, nil
}

func (it *jobIterator) Prev(
	ctx context.Context,
) (prev cli.JobLazyIterator, ok bool, err error) {
	if it.index-1 < 0 {
		return nil, false, nil
	}
	job, err := it.GetJob(ctx, it.array[it.index-1])
	if err != nil {
		return nil, false, fmt.Errorf("failed to get job: %w", err)
	}

	return &jobIterator{
		rpcClient: it.rpcClient,
		array:     it.array,
		length:    it.length,
		index:     it.index - 1,
		job:       job,
	}, true, nil
}

func (it *jobIterator) Current() *cli.Job {
	return it.job
}

func (c *rpcClient) GetJobs(ctx context.Context) (cli.JobLazyIterator, error) {
	jobIDs, err := c.MetaScheduler.GetJobs(&bind.CallOpts{
		Context: ctx,
	}, c.from())
	if err != nil {
		return nil, err
	}
	if len(jobIDs) == 0 {
		return nil, nil
	}
	// Reverse array, from new to old
	for i, j := 0, len(jobIDs)-1; i < j; i, j = i+1, j-1 {
		jobIDs[i], jobIDs[j] = jobIDs[j], jobIDs[i]
	}
	// Initialize first data
	job, err := c.GetJob(ctx, jobIDs[0])
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %w", err)
	}
	return &jobIterator{
		rpcClient: c,
		array:     jobIDs,
		length:    len(jobIDs),
		index:     0,
		job:       job,
	}, nil
}

func (c *rpcClient) requestNewJob(
	ctx context.Context,
	definition metaschedulerabi.JobDefinition,
	lockedCredits *big.Int,
	jobName [32]byte,
	delegateSpendingAuthority bool,
) (id [32]byte, err error) {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create auth options: %w", err)
	}
	tx, err := c.MetaScheduler.RequestNewJob(
		opts,
		definition,
		lockedCredits,
		jobName,
		delegateSpendingAuthority,
	)
	if err != nil {
		return [32]byte{}, err
	}

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create auth options: %w", err)
	}

	// Fetch the event to get the job ID
	for _, log := range receipt.Logs {
		if log.Topics[0].Hex() == newJobRequestEvent.ID.Hex() {
			event, err := c.ParseNewJobRequestEvent(*log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}
			return event.JobId, nil
		}
	}
	return [32]byte{}, ErrNewRequestJobNotFound
}

func (c *rpcClient) SubmitJob(
	ctx context.Context,
	content string,
	definition metaschedulerabi.JobDefinition,
	lockedAmount *big.Int,
	jobName [32]byte,
) ([32]byte, error) {
	hash, err := c.sbatch.Submit(ctx, content)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to submit job: %w", err)
	}
	definition.BatchLocationHash = hash
	return c.requestNewJob(
		ctx,
		definition,
		lockedAmount,
		jobName,
		false, // Set to false, we don't allow any third-party to auto top-up.
	)
}

func (c *rpcClient) CancelJob(ctx context.Context, id [32]byte) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create auth options: %w", err)
	}
	_, err = c.MetaScheduler.CancelJob(
		opts,
		id,
	)
	return err
}

func (c *rpcClient) TopUpJob(ctx context.Context, id [32]byte, amount *big.Int) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create auth options: %w", err)
	}
	_, err = c.MetaScheduler.TopUpJob(
		opts,
		id,
		amount,
	)
	return err
}

func (c *rpcClient) Balance(ctx context.Context) (*big.Int, error) {
	c.from()
	return c.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, c.from())
}

type wsClient struct {
	*metaschedulerabi.MetaScheduler
	Client
}

func (c *wsClient) from() (addr common.Address) {
	if c.UserPrivateKey == nil {
		return addr
	}
	return crypto.PubkeyToAddress(c.UserPrivateKey.PublicKey)
}

func (c *wsClient) SubscribeEvents(
	ctx context.Context,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	creditAddress, err := c.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			c.MetaschedulerAddress,
			creditAddress,
		},
		Topics: [][]common.Hash{
			{
				newJobRequestEvent.ID,
				jobTransitionEvent.ID,
				transferEvent.ID,
				approvalEvent.ID,
			},
		},
	}

	return c.SubscribeFilterLogs(ctx, query, ch)
}

func (c *wsClient) FilterNewJobRequests(
	ch <-chan types.Log,
) (filtered <-chan *metaschedulerabi.MetaSchedulerNewJobRequestEvent, rest <-chan types.Log) {
	fChan := make(chan *metaschedulerabi.MetaSchedulerNewJobRequestEvent)
	rChan := make(chan types.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == newJobRequestEvent.ID.Hex() {
				event, err := c.ParseNewJobRequestEvent(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

func (c *wsClient) FilterJobTransition(
	ch <-chan types.Log,
) (filtered <-chan *metaschedulerabi.MetaSchedulerJobTransitionEvent, rest <-chan types.Log) {
	fChan := make(chan *metaschedulerabi.MetaSchedulerJobTransitionEvent)
	rChan := make(chan types.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == jobTransitionEvent.ID.Hex() {
				event, err := c.ParseJobTransitionEvent(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

type creditFilterer struct {
	cli.CreditManager
	wsClient
	*metaschedulerabi.IERC20Filterer
}

func (c *creditFilterer) FilterTransfer(
	ctx context.Context,
	ch <-chan types.Log,
) (filtered <-chan *metaschedulerabi.IERC20Transfer, rest <-chan types.Log) {
	fChan := make(chan *metaschedulerabi.IERC20Transfer)
	rChan := make(chan types.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == transferEvent.ID.Hex() {
				event, err := c.ParseTransfer(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

func (c *creditFilterer) ReduceToBalance(
	ctx context.Context,
	transfers <-chan *metaschedulerabi.IERC20Transfer,
) (<-chan *big.Int, error) {
	rChan := make(chan *big.Int, 2)
	errChan := make(chan error, 1)

	// Fetch initial value
	value, err := c.CreditManager.Balance(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(rChan)
		defer close(errChan)

		// Send initial value
		rChan <- value

		// Track value
		for transfer := range transfers {
			// User is sending data
			if c.from() == transfer.From {
				value = new(big.Int).Sub(value, transfer.Value)
				rChan <- value
			} else if c.from() == transfer.To {
				value = new(big.Int).Add(value, transfer.Value)
				rChan <- value
			}
		}
	}()

	return rChan, nil
}

type allowanceFilterer struct {
	cli.AllowanceManager
	wsClient
	*metaschedulerabi.IERC20Filterer
}

func (c *allowanceFilterer) FilterApproval(
	ctx context.Context,
	ch <-chan types.Log,
) (filtered <-chan *metaschedulerabi.IERC20Approval, rest <-chan types.Log) {
	fChan := make(chan *metaschedulerabi.IERC20Approval)
	rChan := make(chan types.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == transferEvent.ID.Hex() {
				event, err := c.ParseApproval(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

func (c *allowanceFilterer) ReduceToAllowance(
	ctx context.Context,
	approvals <-chan *metaschedulerabi.IERC20Approval,
) (<-chan *big.Int, error) {
	rChan := make(chan *big.Int, 2)
	errChan := make(chan error, 1)

	// Fetch initial value
	value, err := c.AllowanceManager.GetAllowance(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(rChan)
		defer close(errChan)

		// Send initial value
		rChan <- value

		// Track value
		for approval := range approvals {
			if approval.Owner == c.from() && approval.Spender == c.MetaschedulerAddress {
				rChan <- approval.Value
			}
		}
	}()

	return rChan, nil
}
