// Package event describes the types and methods to interact with the events.
package event

import (
	"context"

	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum"
)

// SubscriptionOptions contains the channels used to pass events.
type SubscriptionOptions struct {
	NewJobRequestChan chan<- types.NewJobRequest
	JobTransitionChan chan<- types.JobTransition
	TransferChan      chan<- types.Transfer
	ApprovalChan      chan<- types.Approval
}

// SubscriptionOption applies default and optional parameters to the SubscribeEvents method.
type SubscriptionOption func(*SubscriptionOptions)

// FilterTransfer allows taking the Transfer events from the subscription.
func FilterTransfer(filtered chan<- types.Transfer) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.TransferChan = filtered
	}
}

// FilterApproval allows taking the Approval events from the subscription.
func FilterApproval(filtered chan<- types.Approval) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.ApprovalChan = filtered
	}
}

// FilterNewJobRequest allows taking the NewJobRequest events from the subscription.
func FilterNewJobRequest(
	filtered chan<- types.NewJobRequest,
) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.NewJobRequestChan = filtered
	}
}

// FilterJobTransition allows taking the JobTransition events from the subscription.
func FilterJobTransition(
	filtered chan<- types.JobTransition,
) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.JobTransitionChan = filtered
	}
}

// Subscriber watches smart-contract events.
type Subscriber interface {
	// Subscribe to metascheduler events.
	SubscribeEvents(
		ctx context.Context,
		opts ...SubscriptionOption,
	) (ethereum.Subscription, error)
}
