package transfer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
)

type ModelBuilder struct {
	Client deepsquare.Client
}

func (b *ModelBuilder) Build() tea.Model {
	return Model(b.Client)
}
