// Package provider defines the interface for managing providers.
package provider

import (
	"context"

	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum/common"
)

// Detail contains all the specs and statuses of a Provider.
type Detail struct {
	metaschedulerabi.Provider
	IsWaitingForApproval bool
	IsValidForScheduling bool
	JobCount             uint64
}

// GetProviderOption is an interface for options for the GetProvider method.
type GetProviderOption func(*GetProviderOptions)

// GetProviderOptions is the structure for holding options for the GetProvider method.
type GetProviderOptions struct {
	Proposal   bool
	Affinities []types.Affinity
}

// WithProposal enables GetProvider to show the proposal of a provider.
func WithProposal() GetProviderOption {
	return func(gpo *GetProviderOptions) {
		gpo.Proposal = true
	}
}

// WithAffinity adds key-value filters with operators to the job, which filters the available clusters.
func WithAffinity(affinities ...types.Affinity) GetProviderOption {
	return func(gpo *GetProviderOptions) {
		gpo.Affinities = affinities
	}
}

// Manager manages admin operation of providers
type Manager interface {
	ApproveProvider(ctx context.Context, provider common.Address) error
	RemoveProvider(ctx context.Context, provider common.Address) error
	GetProvider(
		ctx context.Context,
		address common.Address,
		opts ...GetProviderOption,
	) (provider Detail, err error)
	GetProviders(
		ctx context.Context,
		opts ...GetProviderOption,
	) (providers []Detail, err error)
}
