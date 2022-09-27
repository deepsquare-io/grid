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
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var claimNextJobSig = []byte("ClaimNextJobEvent(address,bytes32,uint256,(uint64,uint64,uint64,uint64,uint64,string))")
var claimNextJobSigHash = crypto.Keccak256Hash(claimNextJobSig)

// DataSource handles communications with the smart contract.
type DataSource struct {
	client          *ethclient.Client
	metascheduler   *metascheduler.MetaScheduler
	providerManager *metascheduler.ProviderManager
	pk              *ecdsa.PrivateKey
	pub             *ecdsa.PublicKey
	fromAddress     common.Address
}

func New(
	rpcEndpoint string,
	hexPK string,
	metaschedulerAddress string,
	providerManagerAddress string,
) *DataSource {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		logger.I.Fatal("ethclient dial failed", zap.Error(err))
	}

	ms, err := metascheduler.NewMetaScheduler(common.HexToAddress(metaschedulerAddress), client)
	if err != nil {
		logger.I.Fatal("metascheduler dial failed", zap.Error(err))
	}
	pm, err := metascheduler.NewProviderManager(common.HexToAddress(providerManagerAddress), client)
	if err != nil {
		logger.I.Fatal("providerManager dial failed", zap.Error(err))
	}
	pk, err := crypto.HexToECDSA(hexPK)
	if err != nil {
		logger.I.Fatal("couldn't decode private key", zap.Error(err))
	}
	publicKey := pk.Public()
	pubECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.I.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*pubECDSA)

	return &DataSource{
		client:          client,
		metascheduler:   ms,
		providerManager: pm,
		pk:              pk,
		pub:             pubECDSA,
		fromAddress:     fromAddress,
	}
}

func (s *DataSource) auth(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := s.client.PendingNonceAt(ctx, s.fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	chainID, err := s.client.ChainID(ctx)
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
func (s *DataSource) Claim(ctx context.Context) (*metascheduler.MetaSchedulerClaimNextJobEvent, error) {
	auth, err := s.auth(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := s.metascheduler.ClaimNextJob(auth)
	if err != nil {
		return nil, err
	}
	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return nil, err
	}

	for _, log := range receipt.Logs {
		switch log.Topics[0].Hex() {
		case claimNextJobSigHash.Hex():
			// TODO: maybe filter by cluster
			r, err := s.metascheduler.ParseClaimNextJobEvent(*log)
			if err != nil {
				return nil, err
			}
			return r, nil
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
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.providerManager.Register(
		auth,
		cpus,
		gpus,
		mem,
		nodes,
	)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(ctx, s.client, tx)
	return err
}
