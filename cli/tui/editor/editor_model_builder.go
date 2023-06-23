package editor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli"
)

type ModelBuilder struct {
	AllowanceManager cli.AllowanceManager
	JobScheduler     cli.JobScheduler
}

func (b *ModelBuilder) Build() tea.Model {
	return Model(b.AllowanceManager, b.JobScheduler)
}
