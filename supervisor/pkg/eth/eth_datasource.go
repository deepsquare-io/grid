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

// TODO: unit test to avoid runtime errors

var ClaimNextJobSig = []byte("ClaimNextJobEvent(address,address,bytes32,uint64,(uint64,uint64,uint64,uint64,string))")
var ClaimNextJobSigHash = crypto.Keccak256Hash(ClaimNextJobSig)

// DataSource handles communications with the smart contract.
type DataSource struct {
	authenticator EthereumAuthenticator
	metascheduler MetaScheduler
	pk            *ecdsa.PrivateKey
	pub           *ecdsa.PublicKey
	fromAddress   common.Address
}

func New(
	a EthereumAuthenticator,
	ms MetaScheduler,
	pk *ecdsa.PrivateKey,
) *DataSource {
	if a == nil {
		logger.I.Fatal("EthereumAuthenticator is nil")
	}
	if ms == nil {
		logger.I.Fatal("MetaScheduler is nil")
	}
	publicKey := pk.Public()
	pubECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.I.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*pubECDSA)
	return &DataSource{
		authenticator: a,
		metascheduler: ms,
		pk:            pk,
		pub:           pubECDSA,
		fromAddress:   fromAddress,
	}
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
// If the queue is not empty, it will claim the job and send it to the SLURM cluster.
// Else, it will return an error.
func (s *DataSource) Claim(ctx context.Context) error {
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.metascheduler.ClaimNextJob(auth)
	if err != nil {
		return err
	}
	logger.I.Debug("called claimnextjob", zap.String("tx", tx.Hash().String()))

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

	// tx, err := s.metascheduler.Register(
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

// StartJob reports the RUNNING state to the metascheduler.
func (s *DataSource) StartJob(
	ctx context.Context,
	jobID [32]byte,
) error {
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}
	tx, err := s.metascheduler.StartJob(
		auth,
		jobID,
	)
	if err != nil {
		return err
	}
	logger.Debug(
		"called start job",
		zap.String("tx", tx.Hash().String()),
	)
	return err
}

// FinishJob sends the invoice to the metascheduler
func (s *DataSource) FinishJob(
	ctx context.Context,
	jobID [32]byte,
	jobDuration uint64,
) error {
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}
	tx, err := s.metascheduler.FinishJob(
		auth,
		jobID,
		jobDuration,
	)
	if err != nil {
		return err
	}
	logger.Debug(
		"called finish job",
		zap.String("tx", tx.Hash().String()),
	)
	return err
}

// FailJob reports the FAILED state to the metascheduler.
func (s *DataSource) FailJob(
	ctx context.Context,
	jobID [32]byte,
) error {
	logger := logger.I.With(zap.String("jobID", common.Bytes2Hex(jobID[:])))
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}
	tx, err := s.metascheduler.TriggerFailedJob(
		auth,
		jobID,
	)
	if err != nil {
		return err
	}
	logger.Debug("called failed job", zap.String("tx", tx.Hash().String()))
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
	tx, err := s.metascheduler.RefuseJob(
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
	return s.metascheduler.WatchClaimNextJobEvent(&bind.WatchOpts{
		Context: ctx,
	}, sink)
}
