// Package credit describes the types and methods to interact with the credits.
package credit

import (
	"context"
	"math/big"

	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
)

// Manager handles the credits of the user.
type Manager interface {
	// Balance fetches the current balance of credits.
	Balance(ctx context.Context) (*big.Int, error)

	// Balance fetches the current balance of credits.
	BalanceOf(ctx context.Context, address common.Address) (*big.Int, error)

	// Transfer tranfers credits from one address to another.
	Transfer(ctx context.Context, to common.Address, amount *big.Int) error

	// ReduceToBalance reduces a channel of transfers into balance.
	ReduceToBalance(
		ctx context.Context,
		transfers <-chan types.Transfer,
	) (<-chan *big.Int, error)
}
