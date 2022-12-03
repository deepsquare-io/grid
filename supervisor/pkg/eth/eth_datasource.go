package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"go.uber.org/zap"
)

// DataSource handles communications with the smart contract.
type DataSource struct {
	authenticator    EthereumAuthenticator
	metaschedulerRPC MetaSchedulerRPC
	metaschedulerWS  MetaSchedulerWS
	pk               *ecdsa.PrivateKey
	fromAddress      common.Address
}

func New(
	a EthereumAuthenticator,
	msRPC MetaSchedulerRPC,
	msWS MetaSchedulerWS,
	pk *ecdsa.PrivateKey,
) *DataSource {
	if a == nil {
		logger.I.Fatal("EthereumAuthenticator is nil")
	}
	if msRPC == nil {
		logger.I.Fatal("MetaSchedulerRPC is nil")
	}
	if msWS == nil {
		logger.I.Fatal("metaschedulerWS is nil")
	}

	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	return &DataSource{
		authenticator:    a,
		metaschedulerRPC: msRPC,
		metaschedulerWS:  msWS,
		pk:               pk,
		fromAddress:      fromAddress,
	}
}

func (s *DataSource) GetProviderAddress() common.Address {
	return s.fromAddress
}

func (s *DataSource) auth(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := s.authenticator.PendingNonceAt(ctx, s.fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.authenticator.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	chainID, err := s.authenticator.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(s.pk, chainID)
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
	ok, err := s.metaschedulerRPC.HasNextJob(&bind.CallOpts{
		Context: ctx,
	}, s.fromAddress)
	if err != nil {
		logger.I.Error("HasNextJob failed", zap.Error(err))
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
		return err
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
		return err
	}
	logger.Debug(
		"called set job status",
		zap.String("tx", tx.Hash().String()),
		zap.Uint8("status", uint8(status)),
	)
	return err
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
		return err
	}
	logger.Debug("called refuse job", zap.String("tx", tx.Hash().String()))
	return err
}

// WatchClaimNextJobEvent observes the incoming ClaimNextJobEvents.
func (s *DataSource) WatchClaimNextJobEvent(
	ctx context.Context,
	sink chan<- *metascheduler.MetaSchedulerClaimNextJobEvent,
) (event.Subscription, error) {
	return s.metaschedulerWS.WatchClaimNextJobEvent(&bind.WatchOpts{
		Context: ctx,
	}, sink)
}

// WatchClaimNextCancellingJobEvent observes the incoming JobCanceledEvents.
func (s *DataSource) WatchClaimNextCancellingJobEvent(
	ctx context.Context,
	sink chan<- *metascheduler.MetaSchedulerClaimNextCancellingJobEvent,
) (event.Subscription, error) {
	return s.metaschedulerWS.WatchClaimNextCancellingJobEvent(&bind.WatchOpts{
		Context: ctx,
	}, sink)
}

// GetJobStatus fetches the job status.
func (s *DataSource) GetJobStatus(ctx context.Context, jobID [32]byte) (JobStatus, error) {
	status, err := s.metaschedulerRPC.GetJobStatus(&bind.CallOpts{
		Context: ctx,
	}, jobID)
	if err != nil {
		return 0, err
	}
	return JobStatus(status), nil
}

// ClaimCancelling a cancelling call.
//
// If the queue is not empty, it will claim the cancelling call and return true.
// Else, it will return false.
// Else, it will return false and an error.
func (s *DataSource) ClaimCancelling(ctx context.Context) error {
	ok, err := s.metaschedulerRPC.HasCancellingJob(&bind.CallOpts{
		Context: ctx,
	}, s.fromAddress)
	if err != nil {
		logger.I.Error("HasCancellingJob failed", zap.Error(err))
		return err
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
		return err
	}
	logger.I.Debug("called ClaimNextCancellingJob", zap.String("tx", tx.Hash().String()))

	return nil
}
