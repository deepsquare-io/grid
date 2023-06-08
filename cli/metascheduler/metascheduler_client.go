// Package deepsquare defines APIs for interacting with the DeepSquare Grid.
package metascheduler

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/deepsquare-io/the-grid/cli/v1"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/v1/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/v1/sbatch"
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
	newJobRequestEvent abi.Event
	jobTransitionEvent abi.Event
)

func init() {
	var err error
	MetaschedulerABI, err = metaschedulerabi.MetaSchedulerMetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to parse contract ABI: %w", err))
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
}

// RPC client for metascheduler.
type RPC interface {
	cli.JobFetcher
	cli.AllowanceManager
	cli.JobScheduler
}

type EthereumBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

type rpcClient struct {
	*metaschedulerabi.MetaScheduler
	EthereumBackend
	metaschedulerAddress common.Address
	chainID              *big.Int
	pk                   *ecdsa.PrivateKey
	sbatch               sbatch.Service
}

func NewRPC(
	address common.Address,
	ethereumBackend EthereumBackend,
	chainID *big.Int,
	pk *ecdsa.PrivateKey,
	sbatch sbatch.Service,
) (client RPC, err error) {
	m, err := metaschedulerabi.NewMetaScheduler(address, ethereumBackend)
	if err != nil {
		return nil, err
	}
	return &rpcClient{
		MetaScheduler:        m,
		EthereumBackend:      ethereumBackend,
		metaschedulerAddress: address,
		chainID:              chainID,
		pk:                   pk,
		sbatch:               sbatch,
	}, err
}

// credit fetches the smart-contract Credit.
func (c *rpcClient) credit(ctx context.Context) (*metaschedulerabi.IERC20, error) {
	address, err := c.MetaScheduler.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	return metaschedulerabi.NewIERC20(address, c)
}

func (c *rpcClient) from() common.Address {
	return crypto.PubkeyToAddress(c.pk.PublicKey)
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

	auth, err := bind.NewKeyedTransactorWithChainID(c.pk, c.chainID)
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
	credit, err := c.credit(ctx)
	if err != nil {
		return fmt.Errorf("failed to get Credit contract: %w", err)
	}
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed get auth options: %w", err)
	}
	tx, err := credit.Approve(opts, c.metaschedulerAddress, amount)
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
	credit, err := c.credit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Credit contract: %w", err)
	}
	return credit.Allowance(&bind.CallOpts{
		Context: ctx,
	}, c.from(), c.metaschedulerAddress)
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

// WS client for metascheduler.
type WS interface {
	cli.JobWatcher
}

type wsClient struct {
	*metaschedulerabi.MetaScheduler
	EthereumBackend
	metaschedulerAddress common.Address
	chainID              *big.Int
	pk                   *ecdsa.PrivateKey
}

func NewWS(
	address common.Address,
	ethereumBackend EthereumBackend,
	chainID *big.Int,
	pk *ecdsa.PrivateKey,
) (c WS, err error) {
	m, err := metaschedulerabi.NewMetaScheduler(address, ethereumBackend)
	if err != nil {
		return nil, err
	}
	return &wsClient{
		MetaScheduler:        m,
		EthereumBackend:      ethereumBackend,
		metaschedulerAddress: address,
		chainID:              chainID,
		pk:                   pk,
	}, err
}

func (c *wsClient) SubscribeEvents(
	ctx context.Context,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.metaschedulerAddress},
		Topics: [][]common.Hash{
			{
				newJobRequestEvent.ID,
				jobTransitionEvent.ID,
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
