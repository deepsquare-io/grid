package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
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
	MetaschedulerABI            *abi.ABI
	claimNextTopUpJobEvent      abi.Event
	claimNextCancellingJobEvent abi.Event
	claimJobEvent               abi.Event
)

func init() {
	var err error
	MetaschedulerABI, err = metascheduler.MetaSchedulerMetaData.GetAbi()
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

// DataSource handles communications with the smart contract.
type DataSource struct {
	bind.DeployBackend
	chainID              *big.Int
	metaschedulerAddress common.Address
	contractBackendRPC   bind.ContractBackend
	metaschedulerRPC     *metascheduler.MetaScheduler
	contractBackendWS    bind.ContractBackend
	metaschedulerWS      *metascheduler.MetaScheduler
	jobQueues            *metascheduler.IProviderJobQueues
	pk                   *ecdsa.PrivateKey
	fromAddress          common.Address
}

func New(
	chainID *big.Int,
	metaschedulerAddress common.Address,
	deployBackend bind.DeployBackend,
	contractBackendRPC bind.ContractBackend,
	contractBackendWS bind.ContractBackend,
	msRPC *metascheduler.MetaScheduler,
	msWS *metascheduler.MetaScheduler,
	pk *ecdsa.PrivateKey,
) *DataSource {
	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)

	address, err := msRPC.ProviderJobQueues(&bind.CallOpts{})
	if err != nil {
		logger.I.Panic("failed to fetch provider job queues smart-contract address", zap.Error(err))
	}
	jobQueues, err := metascheduler.NewIProviderJobQueues(address, contractBackendRPC)
	if err != nil {
		logger.I.Panic(
			"failed to instanciate provider job queues smart-contract address",
			zap.Error(err),
			zap.String("address", address.Hex()),
		)
	}
	return &DataSource{
		DeployBackend:        deployBackend,
		metaschedulerAddress: metaschedulerAddress,
		contractBackendRPC:   contractBackendRPC,
		contractBackendWS:    contractBackendWS,
		chainID:              chainID,
		metaschedulerRPC:     msRPC,
		metaschedulerWS:      msWS,
		jobQueues:            jobQueues,
		pk:                   pk,
		fromAddress:          fromAddress,
	}
}

func (s *DataSource) GetProviderAddress() common.Address {
	return s.fromAddress
}

func (s *DataSource) auth(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := s.contractBackendRPC.PendingNonceAt(ctx, s.fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.contractBackendRPC.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(s.pk, s.chainID)
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
func (s *DataSource) Claim(ctx context.Context) error {
	ok, err := s.jobQueues.HasNextClaimableJob(&bind.CallOpts{
		Context: ctx,
	}, s.fromAddress)
	if err != nil {
		logger.I.Error("HasNextClaimableJob failed", zap.Error(err))
		return err
	}
	if !ok {
		logger.I.Debug("No available job")
		return nil
	}

	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.metaschedulerRPC.ClaimNextJob(auth)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Debug("called ClaimNextJob", zap.String("tx", tx.Hash().String()))

	return nil
}

// Register a cluster
//
// Will send a transaction to register the cluster.
func (s *DataSource) Register(
	ctx context.Context,
	nodes uint64,
	cpus uint64,
	gpus uint64,
	mem uint64,
) error {
	// TODO: implements
	// auth, err := s.auth(ctx)
	// if err != nil {
	// 	return err
	// }

	// tx, err := s.metaschedulerRPC.Register(
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
	// _, err = bind.WaitMined(ctx, s.client, tx)
	// logger.I.Info("register mined", zap.String("tx", tx.Hash().String()))
	// return err
	return nil
}

// SetJobStatus reports the [State] state to the metascheduler.
func (s *DataSource) SetJobStatus(
	ctx context.Context,
	jobID [32]byte,
	status JobStatus,
	jobDurationMinute uint64,
) error {
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}
	tx, err := s.metaschedulerRPC.ProviderSetJobStatus(
		auth,
		jobID,
		uint8(status),
		jobDurationMinute,
	)
	if err != nil {
		return WrapError(err)
	}
	logger.Debug(
		"called set job status, waiting for transaction",
		zap.String("tx", tx.Hash().String()),
		zap.Uint8("status", uint8(status)),
	)
	_, err = bind.WaitMined(ctx, s.DeployBackend, tx)
	if err != nil {
		logger.Error("failed to wait mined", zap.Error(err))
		return WrapError(err)
	}
	logger.Debug(
		"set job status success",
		zap.String("tx", tx.Hash().String()),
		zap.Uint8("status", uint8(status)),
	)
	return nil
}

// RefuseJob rejects a job from the metascheduler.
func (s *DataSource) RefuseJob(
	ctx context.Context,
	jobID [32]byte,
) error {
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	logger.Warn("calling refuse job")
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}
	tx, err := s.metaschedulerRPC.RefuseJob(
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
func (s *DataSource) WatchEvents(
	ctx context.Context,
	claimNextTopUpJobEvents chan<- *metascheduler.MetaSchedulerClaimNextTopUpJobEvent,
	claimNextCancellingJobEvents chan<- *metascheduler.MetaSchedulerClaimNextCancellingJobEvent,
	claimJobEvents chan<- *metascheduler.MetaSchedulerClaimJobEvent,
) (event.Subscription, error) {
	logs := make(chan types.Log)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{s.metaschedulerAddress},
		Topics: [][]common.Hash{
			{
				claimNextTopUpJobEvent.ID,
				claimNextCancellingJobEvent.ID,
				claimJobEvent.ID,
			},
		},
	}

	sub, err := s.contractBackendWS.SubscribeFilterLogs(ctx, query, logs)
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
				switch log.Topics[0].Hex() {
				case claimNextTopUpJobEvent.ID.Hex():
					event, err := s.metaschedulerRPC.ParseClaimNextTopUpJobEvent(log)
					if err != nil {
						logger.I.Panic("failed to parse event", zap.Error(err))
					}

					select {
					case claimNextTopUpJobEvents <- event:
					case err := <-sub.Err():
						logger.I.Error("subscription thrown an error", zap.Error(err))
						return
					}

				case claimNextCancellingJobEvent.ID.Hex():
					event, err := s.metaschedulerRPC.ParseClaimNextCancellingJobEvent(log)
					if err != nil {
						logger.I.Panic("failed to parse event", zap.Error(err))
					}

					select {
					case claimNextCancellingJobEvents <- event:
					case err := <-sub.Err():
						logger.I.Error("subscription thrown an error", zap.Error(err))
						return
					}

				case claimJobEvent.ID.Hex():
					event, err := s.metaschedulerRPC.ParseClaimJobEvent(log)
					if err != nil {
						logger.I.Panic("failed to parse event", zap.Error(err))
					}

					select {
					case claimJobEvents <- event:
					case err := <-sub.Err():
						logger.I.Error("subscription thrown an error", zap.Error(err))
						return
					}
				}

			case err := <-sub.Err():
				logger.I.Error("subscription thrown an error", zap.Error(err))
				return
			}
		}
	}()

	return sub, nil
}

// GetJobStatus fetches the job status.
func (s *DataSource) GetJobStatus(ctx context.Context, jobID [32]byte) (JobStatus, error) {
	status, err := s.metaschedulerRPC.Jobs(&bind.CallOpts{
		Context: ctx,
	}, jobID)
	if err != nil {
		return 0, WrapError(err)
	}
	return JobStatus(status.Status), nil
}

// ClaimCancelling a cancelling call.
func (s *DataSource) ClaimCancelling(ctx context.Context) error {
	ok, err := s.jobQueues.HasCancellingJob(&bind.CallOpts{
		Context: ctx,
	}, s.fromAddress)
	if err != nil {
		logger.I.Error("HasCancellingJob failed", zap.Error(err))
		return WrapError(err)
	}
	if !ok {
		logger.I.Debug("No cancelling call")
		return nil
	}

	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.metaschedulerRPC.ClaimNextCancellingJob(auth)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Debug("called ClaimNextCancellingJob", zap.String("tx", tx.Hash().String()))

	return nil
}

// ClaimTopUp a top up call.
func (s *DataSource) ClaimTopUp(ctx context.Context) error {
	ok, err := s.jobQueues.HasTopUpJob(&bind.CallOpts{
		Context: ctx,
	}, s.fromAddress)
	if err != nil {
		logger.I.Error("HasTopUpJob failed", zap.Error(err))
		return WrapError(err)
	}
	if !ok {
		logger.I.Debug("No top up call")
		return nil
	}

	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.metaschedulerRPC.ClaimNextTopUpJob(auth)
	if err != nil {
		return WrapError(err)
	}
	logger.I.Debug("called ClaimNextTopUpJob", zap.String("tx", tx.Hash().String()))

	return nil
}
