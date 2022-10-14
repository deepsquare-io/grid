package eth

import (
	"context"
	"math/big"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthereumAuthenticator interface {
	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// ChainID retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)
}

type MetaScheduler interface {
	ClaimNextJob(opts *bind.TransactOpts) (*types.Transaction, error)
	FinishJob(opts *bind.TransactOpts, _jobID [32]byte, actualJobDurationMinute uint64) (*types.Transaction, error)
	StartJob(opts *bind.TransactOpts, _jobID [32]byte) (*types.Transaction, error)
	RefuseJob(opts *bind.TransactOpts, _jobID [32]byte) (*types.Transaction, error)
	TriggerFailedJob(opts *bind.TransactOpts, _jobID [32]byte) (*types.Transaction, error)
	ParseClaimNextJobEvent(log types.Log) (*metascheduler.MetaSchedulerClaimNextJobEvent, error)
}
