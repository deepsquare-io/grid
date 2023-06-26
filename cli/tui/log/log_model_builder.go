package log

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
)

type ModelBuilder struct {
	LoggerDialer types.LoggerDialer
	UserAddress  common.Address
}

func (b *ModelBuilder) Build(ctx context.Context, jobID [32]byte) tea.Model {
	return Model(ctx, b.LoggerDialer, b.UserAddress, jobID)
}
