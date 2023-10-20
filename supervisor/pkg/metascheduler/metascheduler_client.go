// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package metascheduler

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"go.uber.org/zap"
)

var (
	MetaschedulerABI *abi.ABI

	claimNextTopUpJobEvent      abi.Event
	claimNextCancellingJobEvent abi.Event
	claimJobEvent               abi.Event
)

func init() {
	var err error
	MetaschedulerABI, err = metaschedulerabi.MetaSchedulerMetaData.GetAbi()
	if err != nil {
		logger.I.Fatal("failed to parse contract ABI", zap.Error(err))
	}

	// Find the event signature dynamically
	var ok bool
	claimNextTopUpJobEvent, ok = MetaschedulerABI.Events["ClaimNextTopUpJobEvent"]
	if !ok {
		logger.I.Fatal("failed to find ClaimNextTopUpJobEvent in contract ABI")
	}

	claimNextCancellingJobEvent, ok = MetaschedulerABI.Events["ClaimNextCancellingJobEvent"]
	if !ok {
		logger.I.Fatal("failed to find ClaimNextCancellingJobEvent in contract ABI")
	}

	claimJobEvent, ok = MetaschedulerABI.Events["ClaimJobEvent"]
	if !ok {
		logger.I.Fatal("failed to find ClaimJobEvent in contract ABI")
	}
}

// Client handles communications with the smart contract.
type Client struct {
	bind.DeployBackend
	chainID              *big.Int
	metaschedulerAddress common.Address
	rpc                  bind.ContractBackend
	contractRPC          *metaschedulerabi.MetaScheduler
	ws                   bind.ContractBackend
	contractWS           *metaschedulerabi.MetaScheduler
	jobQueues            *metaschedulerabi.IProviderJobQueues
	jobs                 *metaschedulerabi.IJobRepository
	providerManager      *metaschedulerabi.IProviderManager
	pk                   *ecdsa.PrivateKey
	fromAddress          common.Address
}

func NewClient(
	chainID *big.Int,
	metaschedulerAddress common.Address,
	deployBackend bind.DeployBackend,
	rpc bind.ContractBackend,
	ws bind.ContractBackend,
	pk *ecdsa.PrivateKey,
) MetaScheduler {
	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)

	msRPC, err := metaschedulerabi.NewMetaScheduler(metaschedulerAddress, rpc)
	if err != nil {
		logger.I.Fatal("failed to create metascheduler with rpc", zap.Error(err))
	}
	msWS, err := metaschedulerabi.NewMetaScheduler(metaschedulerAddress, ws)
	if err != nil {
		logger.I.Fatal("failed to create metascheduler with rpc", zap.Error(err))
	}
	jobQueuesAddress, err := msRPC.ProviderJobQueues(&bind.CallOpts{})
	if err != nil {
		logger.I.Panic("failed to fetch provider job queues smart-contract address", zap.Error(err))
	}
	jobQueues, err := metaschedulerabi.NewIProviderJobQueues(jobQueuesAddress, rpc)
	if err != nil {
		logger.I.Panic(
			"failed to instanciate provider job queues",
			zap.Error(err),
			zap.String("address", jobQueuesAddress.Hex()),
		)
	}
	jobsAddress, err := msRPC.Jobs(&bind.CallOpts{})
	if err != nil {
		logger.I.Panic("failed to fetch jobs smart-contract address", zap.Error(err))
	}
	jobs, err := metaschedulerabi.NewIJobRepository(jobsAddress, rpc)
	if err != nil {
		logger.I.Panic(
			"failed to instanciate job repository",
			zap.Error(err),
			zap.String("address", jobsAddress.Hex()),
		)
	}
	providerManagerAddress, err := msRPC.ProviderManager(&bind.CallOpts{})
	if err != nil {
		logger.I.Panic("failed to fetch provider manager smart-contract address", zap.Error(err))
	}
	providerManager, err := metaschedulerabi.NewIProviderManager(providerManagerAddress, rpc)
	if err != nil {
		logger.I.Panic(
			"failed to instanciate provider manager",
			zap.Error(err),
			zap.String("address", providerManagerAddress.Hex()),
		)
	}
	return &Client{
		DeployBackend:        deployBackend,
		metaschedulerAddress: metaschedulerAddress,
		rpc:                  rpc,
		ws:                   ws,
		chainID:              chainID,
		contractRPC:          msRPC,
		contractWS:           msWS,
		jobQueues:            jobQueues,
		pk:                   pk,
		fromAddress:          fromAddress,
		jobs:                 jobs,
		providerManager:      providerManager,
	}
}

func (c *Client) IsRequestNewJobEnabled(ctx context.Context) (bool, error) {
	v, err := c.contractRPC.EnableRequestNewJob(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return false, WrapError(err)
	}
	return v, nil
}

func (c *Client) GetProviderAddress() common.Address {
	return c.fromAddress
}

func (c *Client) auth(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.rpc.PendingNonceAt(ctx, c.fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.rpc.SuggestGasPrice(ctx)
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

	return auth, nil
}

// Claim a job.
//
// If the queue is not empty, it will claim the job and return true.
// Else, it will return false.
// Else, it will return false and an error.
func (c *Client) Claim(ctx context.Context) error {
	ok, err := c.jobQueues.HasNextClaimableJob(&bind.CallOpts{
		Context: ctx,
	}, c.fromAddress)
	if err != nil {
		logger.I.Error("HasNextClaimableJob failed", zap.Error(err))
		return err
	}
	if !ok {
		return nil
	}

	auth, err := c.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := c.contractRPC.ClaimNextJob(auth)
	if err != nil {
		return WrapError(err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Debug("called ClaimNextJob", zap.String("tx", tx.Hash().String()))

	return nil
}

func (c *Client) GetOldInfo(ctx context.Context) (*metaschedulerabi.Provider, error) {
	p, err := c.providerManager.GetProvider(&bind.CallOpts{Context: ctx}, c.fromAddress)
	if err != nil {
		return nil, WrapError(err)
	}
	return &p, nil
}

// Register a cluster
//
// Will send a transaction to register the cluster.
func (c *Client) Register(
	ctx context.Context,
	hardware metaschedulerabi.ProviderHardware,
	prices metaschedulerabi.ProviderPrices,
	labels []metaschedulerabi.Label,
) error {
	logger.I.Info(
		"called register",
		zap.Any("hardware", hardware),
		zap.Any("prices", prices),
		zap.Any("labels", labels),
	)
	auth, err := c.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := c.providerManager.Register(
		auth,
		hardware,
		prices,
		labels,
	)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Info("called register", zap.String("tx", tx.Hash().String()))
	_, err = bind.WaitMined(ctx, c.DeployBackend, tx)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Info("register mined", zap.String("tx", tx.Hash().String()))
	return nil
}

// SetJobStatus reports the [State] state to the metascheduler.
func (c *Client) SetJobStatus(
	ctx context.Context,
	jobID [32]byte,
	status JobStatus,
	jobDurationMinute uint64,
	opts ...SetJobStatusOption,
) error {
	o := applySetJobStatusOptions(opts)
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	auth, err := c.auth(ctx)
	if err != nil {
		return err
	}
	var errMsg string
	if o.err != nil {
		errMsg = o.err.Error()
	}
	tx, err := c.contractRPC.ProviderSetJobStatus(
		auth,
		jobID,
		uint8(status),
		jobDurationMinute,
		errMsg,
		o.exitCode,
	)
	if err != nil {
		return WrapError(err)
	}
	logger.Debug(
		"called set job status, waiting for transaction",
		zap.String("tx", tx.Hash().String()),
		zap.Uint8("status", uint8(status)),
	)
	_, err = bind.WaitMined(ctx, c.DeployBackend, tx)
	if err != nil {
		logger.Error("failed to wait mined", zap.Error(err))
		return WrapError(err)
	}
	logger.Debug(
		"set job status",
		zap.String("tx", tx.Hash().String()),
		zap.Uint8("status", uint8(status)),
	)
	return nil
}

// RefuseJob rejects a job from the metascheduler.
func (c *Client) RefuseJob(
	ctx context.Context,
	jobID [32]byte,
) error {
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	logger.Warn("calling refuse job")
	auth, err := c.auth(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contractRPC.RefuseJob(
		auth,
		jobID,
	)
	if err != nil {
		return WrapError(err)
	}
	logger.Debug("called refuse job", zap.String("tx", tx.Hash().String()))
	return nil
}

// WatchEvents observes the incoming ClaimNextTopUpJobEvent, ClaimNextCancellingJobEvent and ClaimJobEvent.
func (c *Client) WatchEvents(
	ctx context.Context,
	claimNextTopUpJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent,
	claimNextCancellingJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
	claimJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent,
) (event.Subscription, error) {
	logs := make(chan types.Log, 100)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.metaschedulerAddress},
		Topics: [][]common.Hash{
			{
				claimNextTopUpJobEvent.ID,
				claimNextCancellingJobEvent.ID,
				claimJobEvent.ID,
			},
		},
	}

	sub, err := c.ws.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		logger.I.Error("failed to subscribe", zap.Error(err))
		return nil, WrapError(err)
	}

	go func() {
		defer close(logs)
		defer sub.Unsubscribe()
		for {
			select {
			case log, ok := <-logs:
				if !ok {
					return
				}
				if len(log.Topics) == 0 {
					return
				}
				switch log.Topics[0].Hex() {
				case claimNextTopUpJobEvent.ID.Hex():
					event, err := c.contractRPC.ParseClaimNextTopUpJobEvent(log)
					if err != nil {
						logger.I.Panic("failed to parse event", zap.Error(err))
					}

					select {
					case claimNextTopUpJobEvents <- event:
					case <-ctx.Done():
						logger.I.Error("subscription canceled", zap.Error(ctx.Err()))
						return
					}

				case claimNextCancellingJobEvent.ID.Hex():
					event, err := c.contractRPC.ParseClaimNextCancellingJobEvent(log)
					if err != nil {
						logger.I.Panic("failed to parse event", zap.Error(err))
					}

					select {
					case claimNextCancellingJobEvents <- event:
					case <-ctx.Done():
						logger.I.Error("subscription canceled", zap.Error(ctx.Err()))
						return
					}

				case claimJobEvent.ID.Hex():
					event, err := c.contractRPC.ParseClaimJobEvent(log)
					if err != nil {
						logger.I.Panic("failed to parse event", zap.Error(err))
					}

					select {
					case claimJobEvents <- event:
					case <-ctx.Done():
						logger.I.Error("subscription canceled", zap.Error(ctx.Err()))
						return
					}
				}
			case <-ctx.Done():
				logger.I.Error("subscription canceled", zap.Error(ctx.Err()))
				return
			}
		}
	}()

	return sub, nil
}

// GetJobStatus fetches the job status.
func (c *Client) GetJobStatus(ctx context.Context, jobID [32]byte) (JobStatus, error) {
	job, err := c.jobs.Get(&bind.CallOpts{
		Context: ctx,
	}, jobID)
	if err != nil {
		return 0, WrapError(err)
	}
	return JobStatus(job.Status), nil
}

// ClaimCancelling a cancelling call.
func (c *Client) ClaimCancelling(ctx context.Context) error {
	ok, err := c.jobQueues.HasCancellingJob(&bind.CallOpts{
		Context: ctx,
	}, c.fromAddress)
	if err != nil {
		logger.I.Error("HasCancellingJob failed", zap.Error(err))
		return WrapError(err)
	}
	if !ok {
		return nil
	}

	auth, err := c.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := c.contractRPC.ClaimNextCancellingJob(auth)
	if err != nil {
		return WrapError(err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Debug("called ClaimNextCancellingJob", zap.String("tx", tx.Hash().String()))

	return nil
}

// ClaimTopUp a top up call.
func (c *Client) ClaimTopUp(ctx context.Context) error {
	ok, err := c.jobQueues.HasTopUpJob(&bind.CallOpts{
		Context: ctx,
	}, c.fromAddress)
	if err != nil {
		logger.I.Error("HasTopUpJob failed", zap.Error(err))
		return WrapError(err)
	}
	if !ok {
		return nil
	}

	auth, err := c.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := c.contractRPC.ClaimNextTopUpJob(auth)
	if err != nil {
		return WrapError(err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Debug("called ClaimNextTopUpJob", zap.String("tx", tx.Hash().String()))

	return nil
}

func (c *Client) GetJob(ctx context.Context, jobID [32]byte) (*Job, error) {
	job, err := c.jobs.Get(&bind.CallOpts{Context: ctx}, jobID)
	if err != nil {
		return &Job{}, WrapError(err)
	}
	return FromStructToJob(job), nil
}

func (c *Client) GetJobs(ctx context.Context) (*ProviderJobIterator, error) {
	it, err := c.contractRPC.FilterClaimJobEvent(&bind.FilterOpts{
		Context: ctx,
	})
	if err != nil {
		logger.I.Error("FilterClaimJobEvent failed", zap.Error(err))
		return nil, WrapError(err)
	}

	// Find a job for the provider
	for it.Next() {
		// Filter case
		if bytes.EqualFold(it.Event.ProviderAddr[:], c.fromAddress[:]) {
			job, err := c.GetJob(ctx, it.Event.JobId)
			if err != nil {
				logger.I.Error("GetJob failed", zap.Error(err))
				return nil, err
			}

			return &ProviderJobIterator{
				client:                             c,
				Job:                                job,
				MetaSchedulerClaimJobEventIterator: it,
				providerAddress:                    c.fromAddress,
			}, nil
		}
	}

	// Not found case
	return &ProviderJobIterator{
		MetaSchedulerClaimJobEventIterator: it,
		providerAddress:                    c.fromAddress,
	}, nil
}
