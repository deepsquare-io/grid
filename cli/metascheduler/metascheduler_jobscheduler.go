package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/deepsquare-io/the-grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
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
	lockedAmount *big.Int,
	jobName [32]byte,
	opts ...types.SubmitJobOption,
) ([32]byte, error) {
	var o types.SubmitJobOptions
	for _, opt := range opts {
		opt(&o)
	}
	hash, err := c.Submit(ctx, job)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to submit job: %w", err)
	}

	msUses := make([]metaschedulerabi.Label, 0, len(o.Uses))
	for _, u := range o.Uses {
		msUses = append(msUses, metaschedulerabi.Label{
			Key:   u.Key,
			Value: u.Value,
		})
	}

	msAffinities := make([]metaschedulerabi.Affinity, 0, len(o.Affinities))
	for _, a := range o.Affinities {
		msAffinities = append(msAffinities, metaschedulerabi.Affinity{
			Label: a.Label,
			Op:    a.Op,
		})
	}

	definition := metaschedulerabi.JobDefinition{
		Ntasks:            uint64(job.Resources.Tasks),
		GpusPerTask:       uint64(job.Resources.GpusPerTask),
		MemPerCpu:         uint64(job.Resources.MemPerCPU),
		CpusPerTask:       uint64(job.Resources.CpusPerTask),
		StorageType:       0,
		BatchLocationHash: hash,
		Uses:              msUses,
		Affinity:          msAffinities,
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

func (c *jobScheduler) PanicJob(ctx context.Context, id [32]byte, reason string) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create auth options: %w", err)
	}
	_, err = c.MetaScheduler.PanicJob(
		opts,
		id,
		reason,
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
