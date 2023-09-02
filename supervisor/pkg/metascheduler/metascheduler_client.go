package metascheduler

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
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
			"failed to instanciate provider job queues smart-contract address",
			zap.Error(err),
			zap.String("address", jobQueuesAddress.Hex()),
		)
	}
	jobsAddress, err := msRPC.Jobs(&bind.CallOpts{})
	if err != nil {
		logger.I.Panic("failed to fetch provider job queues smart-contract address", zap.Error(err))
	}
	jobs, err := metaschedulerabi.NewIJobRepository(jobsAddress, rpc)
	if err != nil {
		logger.I.Panic(
			"failed to instanciate job repository smart-contract address",
			zap.Error(err),
			zap.String("address", jobsAddress.Hex()),
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
	}
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
		logger.I.Debug("No available job")
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

// Register a cluster
//
// Will send a transaction to register the cluster.
func (c *Client) Register(
	ctx context.Context,
	nodes uint64,
	cpus uint64,
	gpus uint64,
	mem uint64,
	gflops float64,
) error {
	logger.I.Info(
		"called register",
		zap.Uint64("nodes", nodes),
		zap.Uint64("cpus", cpus),
		zap.Uint64("mem", mem),
		zap.Uint64("gpus", gpus),
		zap.Float64("gflops", gflops),
	)
	// TODO: implements
	// auth, err := c.auth(ctx)
	// if err != nil {
	// 	return err
	// }

	// tx, err := c.metaschedulerRPC.Register(
	// 	auth,
	// 	cpus,
	// 	gpus,
	// 	mem,
	// 	nodes,
	// )
	// if err != nil {
	// 	return err
	// }
	// logger.I.Info("called register", zap.String("tx", tx.Hash().String()))
	// _, err = bind.WaitMined(ctx, c.client, tx)
	// logger.I.Info("register mined", zap.String("tx", tx.Hash().String()))
	// return err
	return nil
}

// SetJobStatus reports the [State] state to the metascheduler.
func (c *Client) SetJobStatus(
	ctx context.Context,
	jobID [32]byte,
	status JobStatus,
	jobDurationMinute uint64,
) error {
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	auth, err := c.auth(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contractRPC.ProviderSetJobStatus(
		auth,
		jobID,
		uint8(status),
		jobDurationMinute,
		"",
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
		logger.I.Debug("No cancelling call")
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
		logger.I.Debug("No top up call")
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
		if it.Event.ProviderAddr == c.fromAddress {
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
