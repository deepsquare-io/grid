package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/deepsquare-io/the-grid/cli/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// DataSource handles communications with the smart contract.
type DataSource struct {
	client        *ethclient.Client
	metascheduler *metascheduler.MetaScheduler
	pk            *ecdsa.PrivateKey
	pub           *ecdsa.PublicKey
	fromAddress   common.Address
}

func New(
	rpcEndpoint string,
	hexPK string,
	metaschedulerAddress string,
) *DataSource {
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		logger.I.Fatal("ethclient dial failed", zap.Error(err))
	}

	ms, err := metascheduler.NewMetaScheduler(common.HexToAddress(metaschedulerAddress), client)
	if err != nil {
		logger.I.Fatal("metascheduler dial failed", zap.Error(err))
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
		client:        client,
		metascheduler: ms,
		pk:            pk,
		pub:           pubECDSA,
		fromAddress:   fromAddress,
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

// Request a job.
func (s *DataSource) RequestNewJob(ctx context.Context, jobDefinition metascheduler.JobDefinition, amountLocked *big.Int) error {
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.metascheduler.RequestNewJob(auth, jobDefinition, amountLocked)
	if err != nil {
		return err
	}
	logger.I.Debug("called RequestNewJob", zap.String("tx", tx.Hash().String()))

	return nil
}
