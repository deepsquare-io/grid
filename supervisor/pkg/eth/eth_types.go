package eth

import (
	"context"
	"math/big"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

type MetaSchedulerRPC interface {
	ClaimNextJob(opts *bind.TransactOpts) (*types.Transaction, error)
	HasNextJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error)
	ClaimNextCancellingJob(opts *bind.TransactOpts) (*types.Transaction, error)
	HasCancellingJob(opts *bind.CallOpts, _providerAddr common.Address) (bool, error)
	RefuseJob(opts *bind.TransactOpts, _jobID [32]byte) (*types.Transaction, error)
	ParseClaimJobEvent(log types.Log) (*metascheduler.MetaSchedulerClaimJobEvent, error)
	ProviderSetJobStatus(opts *bind.TransactOpts, _jobID [32]byte, _jobStatus uint8, jobDurationMinute uint64) (*types.Transaction, error)
	Jobs(opts *bind.CallOpts, arg0 [32]byte) (struct {
		JobId            [32]byte
		Status           uint8
		CustomerAddr     common.Address
		ProviderAddr     common.Address
		Definition       metascheduler.JobDefinition
		Valid            bool
		Cost             metascheduler.JobCost
		Time             metascheduler.JobTime
		JobName          [32]byte
		HasCancelRequest bool
	}, error)
}

type MetaSchedulerWS interface {
	WatchClaimJobEvent(opts *bind.WatchOpts, sink chan<- *metascheduler.MetaSchedulerClaimJobEvent) (event.Subscription, error)
	WatchClaimNextCancellingJobEvent(opts *bind.WatchOpts, sink chan<- *metascheduler.MetaSchedulerClaimNextCancellingJobEvent) (event.Subscription, error)
}
