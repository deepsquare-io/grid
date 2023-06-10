package status

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/ethereum/go-ethereum/common"
)

var columns = []table.Column{
	{Title: "Job ID", Width: 6},
	{Title: "Job Name", Width: 30},
	{Title: "Status", Width: 14},
	{Title: "Start date", Width: 30},
}

func (m model) View() string {
	return style.Base.Render(m.table.View())
}

func Model(
	ctx context.Context,
	fetcher cli.JobFetcher,
	watcher cli.JobWatcher,
	userAddress common.Address,
) tea.Model {
	// Initialize rows
	rows, idToRow, it := initializeRows(ctx, fetcher)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(style.StandardHeight),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	help := help.New()
	help.ShowAll = true

	return model{
		table:   t,
		idToRow: idToRow,
		it:      it,
		help:    help,

		jobs:        make(chan cli.Job, 100),
		fetcher:     fetcher,
		watcher:     watcher,
		userAddress: userAddress,
	}
}
