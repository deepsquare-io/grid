package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type jobScheduler struct {
	*RPCClientSet
	*metaschedulerabi.MetaScheduler
	sbatch.Service
}

func (c *jobScheduler) requestNewJob(
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

func (c *jobScheduler) SubmitJob(
	ctx context.Context,
	job *sbatch.Job,
	uses []metaschedulerabi.Label,
	lockedAmount *big.Int,
	jobName [32]byte,
) ([32]byte, error) {
	hash, err := c.Submit(ctx, job)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to submit job: %w", err)
	}

	definition := metaschedulerabi.JobDefinition{
		Ntasks:            uint64(job.Resources.Tasks),
		GpusPerTask:       uint64(job.Resources.GpusPerTask),
		MemPerCpu:         uint64(job.Resources.MemPerCPU),
		CpusPerTask:       uint64(job.Resources.CpusPerTask),
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

func (c *jobScheduler) CancelJob(ctx context.Context, id [32]byte) error {
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

func (c *jobScheduler) TopUpJob(ctx context.Context, id [32]byte, amount *big.Int) error {
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
