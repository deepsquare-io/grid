package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type rpcClient struct {
	*metaschedulerabi.MetaScheduler
	*metaschedulerabi.IERC20
	Backend
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
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}
	return nil
}

func (c *rpcClient) ClearAllowance(ctx context.Context) error {
	return c.SetAllowance(ctx, big.NewInt(0))
}

func (c *rpcClient) GetAllowance(ctx context.Context) (*big.Int, error) {
	return c.Allowance(&bind.CallOpts{
		Context: ctx,
	}, c.from(), c.MetaschedulerAddress)
}

func (c *rpcClient) GetJob(ctx context.Context, id [32]byte) (*types.Job, error) {
	job, err := c.Jobs(&bind.CallOpts{
		Context: ctx,
	}, id)
	if err != nil {
		return nil, WrapError(err)
	}
	return &types.Job{
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
	job    *types.Job
}

func (it *jobIterator) Next(
	ctx context.Context,
) (next types.JobLazyIterator, ok bool, err error) {
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
) (prev types.JobLazyIterator, ok bool, err error) {
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

func (it *jobIterator) Current() *types.Job {
	return it.job
}

func (c *rpcClient) GetJobs(ctx context.Context) (types.JobLazyIterator, error) {
	jobIDs, err := c.MetaScheduler.GetJobs(&bind.CallOpts{
		Context: ctx,
	}, c.from())
	if err != nil {
		return nil, WrapError(err)
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
		return [32]byte{}, WrapError(err)
	}

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to wait transaction to be mined: %w", err)
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
	job *sbatch.Job,
	uses []metaschedulerabi.Label,
	lockedAmount *big.Int,
	jobName [32]byte,
) ([32]byte, error) {
	hash, err := c.sbatch.Submit(ctx, job)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to submit job: %w", err)
	}

	definition := metaschedulerabi.JobDefinition{
		Ntasks:            uint64(job.Resources.Tasks),
		GpuPerTask:        uint64(job.Resources.GpusPerTask),
		MemPerCpu:         uint64(job.Resources.MemPerCPU),
		CpuPerTask:        uint64(job.Resources.CpusPerTask),
		StorageType:       0,
		BatchLocationHash: hash,
		Uses:              uses,
	}
	id, err := c.requestNewJob(
		ctx,
		definition,
		lockedAmount,
		jobName,
		false, // Set to false, we don't allow any third-party to auto top-up.
	)
	return id, WrapError(err)
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
	return WrapError(err)
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
	return WrapError(err)
}

func (c *rpcClient) Balance(ctx context.Context) (*big.Int, error) {
	c.from()
	balance, err := c.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, c.from())
	return balance, WrapError(err)
}
