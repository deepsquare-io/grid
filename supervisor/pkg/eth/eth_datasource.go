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
	"go.uber.org/zap"
)

// TODO: unit test to avoid runtime errors

var ClaimNextJobSig = []byte("ClaimNextJobEvent(address,address,bytes32,uint64,(uint64,uint64,uint64,uint64,string))")
var ClaimNextJobSigHash = crypto.Keccak256Hash(ClaimNextJobSig)

// DataSource handles communications with the smart contract.
type DataSource struct {
	authenticator EthereumAuthenticator
	deployBackend bind.DeployBackend
	metascheduler MetaScheduler
	pk            *ecdsa.PrivateKey
	pub           *ecdsa.PublicKey
	fromAddress   common.Address
}

func New(
	a EthereumAuthenticator,
	b bind.DeployBackend,
	ms MetaScheduler,
	pk *ecdsa.PrivateKey,
) *DataSource {
	if a == nil {
		logger.I.Fatal("EthereumAuthenticator is nil")
	}
	if b == nil {
		logger.I.Fatal("DeployBackend is nil")
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
		deployBackend: b,
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
//
// Can return nil if no event found.
func (s *DataSource) Claim(ctx context.Context) (*metascheduler.MetaSchedulerClaimNextJobEvent, error) {
	auth, err := s.auth(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := s.metascheduler.ClaimNextJob(auth)
	if err != nil {
		return nil, err
	}
	logger.I.Debug("called claimnextjob", zap.String("tx", tx.Hash().String()))
	receipt, err := bind.WaitMined(ctx, s.deployBackend, tx)
	if err != nil {
		return nil, err
	}
	logger.I.Debug("claimnextjob has been mined", zap.String("tx", tx.Hash().String()))

	// TODO: handle in an independent listener
	for _, log := range receipt.Logs {
		logger.I.Debug("claimnextjob found event", zap.Any("event", log))
		switch log.Topics[0].Hex() {
		case ClaimNextJobSigHash.Hex():
			r, err := s.metascheduler.ParseClaimNextJobEvent(*log)
			if err != nil {
				return nil, err
			}
			if r.ProviderAddr.Hex() == s.fromAddress.Hex() {
				logger.I.Debug("claimnextjob selected event", zap.Any("event", log), zap.Any("content", r))
				return r, nil
			}
			logger.I.Debug("claimnextjob skipped event", zap.Any("event", log), zap.Any("content", r))
		}
	}

	return nil, nil
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
	// We need to wait to make sure the job is accepted by the metascheduler and avoid race conditions
	_, err = bind.WaitMined(ctx, s.deployBackend, tx)
	if err != nil {
		return err
	}
	logger.Debug("start job has been mined", zap.String("tx", tx.Hash().String()))
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
	_, err = bind.WaitMined(ctx, s.deployBackend, tx)
	if err != nil {
		return err
	}
	logger.Debug("refuse job has been mined", zap.String("tx", tx.Hash().String()))
	return err
}
