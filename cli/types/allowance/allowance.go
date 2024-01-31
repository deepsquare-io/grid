// Package allowance describes the types and methods to interact with the
package allowance

import (
	"context"
	"math/big"

	"github.com/deepsquare-io/grid/cli/types"
)

// Manager set the allowed quantity of credit for smart-contract interactions.
type Manager interface {
	// Set the allowance for smart-contract interactions.
	SetAllowance(ctx context.Context, amount *big.Int) error

	// ClearAllowance is an alias to SetAllowance 0.
	ClearAllowance(ctx context.Context) error

	// Get the current allowance toward the contract.
	GetAllowance(ctx context.Context) (*big.Int, error)

	// ReduceToAllowance reduces a channel of approval into allowance.
	ReduceToAllowance(
		ctx context.Context,
		approvals <-chan types.Approval,
	) (<-chan *big.Int, error)
}
