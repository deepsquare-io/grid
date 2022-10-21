package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/deepsquare-io/the-grid/cli/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// DataSource handles communications with the smart contract.
type DataSource struct {
	client               *ethclient.Client
	metaschedulerAddress common.Address
	metascheduler        *metascheduler.MetaScheduler
	credit               *metascheduler.IERC20Metadata
	pk                   *ecdsa.PrivateKey
	pub                  *ecdsa.PublicKey
	fromAddress          common.Address
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

	ma := common.HexToAddress(metaschedulerAddress)
	ms, err := metascheduler.NewMetaScheduler(ma, client)
	if err != nil {
		logger.I.Fatal("metascheduler dial failed", zap.Error(err))
	}
	creditAddr, err := ms.Credit(nil)
	if err != nil {
		logger.I.Fatal("metascheduler failed to fetch credit address", zap.Error(err))
	}
	ierc20, err := metascheduler.NewIERC20Metadata(creditAddr, client)
	if err != nil {
		logger.I.Fatal("IERC20Metadata dial failed", zap.Error(err))
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
		client:               client,
		metaschedulerAddress: ma,
		metascheduler:        ms,
		credit:               ierc20,
		pk:                   pk,
		pub:                  pubECDSA,
		fromAddress:          fromAddress,
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

// RequestNewJob submits job.
func (s *DataSource) RequestNewJob(ctx context.Context, jobDefinition metascheduler.JobDefinition, amountLocked *big.Int) (*types.Transaction, error) {
	if err := s.approve(ctx, amountLocked); err != nil {
		return nil, err
	}

	auth, err := s.auth(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := s.metascheduler.RequestNewJob(auth, jobDefinition, amountLocked)
	if err != nil {
		return nil, err
	}
	logger.I.Debug("called RequestNewJob", zap.String("tx", tx.Hash().String()))

	return tx, err
}

// Register a provider.
func (s *DataSource) Register(ctx context.Context, providerAddress common.Address, providerDefinition metascheduler.ProviderDefinition) (*types.Transaction, error) {
	auth, err := s.auth(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := s.metascheduler.ProviderRegister(auth, providerAddress, providerDefinition)
	if err != nil {
		return nil, err
	}
	logger.I.Debug("called ProviderRegister", zap.String("tx", tx.Hash().String()))
	return tx, err
}

// Approve
func (s *DataSource) approve(ctx context.Context, amount *big.Int) error {
	auth, err := s.auth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.credit.Approve(auth, s.metaschedulerAddress, amount)
	if err != nil {
		return err
	}
	logger.I.Debug("called Approve", zap.String("tx", tx.Hash().String()))
	_, err = bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return err
	}
	logger.I.Debug("mined Approve", zap.String("tx", tx.Hash().String()))
	return nil
}
