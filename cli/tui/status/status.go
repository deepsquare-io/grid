package status

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/deepsquare/metascheduler"
	"github.com/ethereum/go-ethereum/common"
)

var (
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
)

const tableHeight = 10

var columns = []table.Column{
	{Title: "Job ID", Width: 6},
	{Title: "Job Name", Width: 30},
	{Title: "Status", Width: 14},
	{Title: "Start date", Width: 30},
}

func (m Model) View() string {
	s := baseStyle.Render(m.table.View()) + "\n"
	return s
}

func Status(
	ctx context.Context,
	rpc metascheduler.RPC,
	ws metascheduler.WS,
	userAddress common.Address,
) tea.Model {
	// Initialize rows
	rows, idToRow, it := initializeRows(ctx, rpc)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(tableHeight),
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

	return Model{
		table:   t,
		idToRow: idToRow,
		it:      it,
		help:    help,

		jobs:        make(chan deepsquare.Job, 100),
		fetcher:     rpc,
		watcher:     ws,
		userAddress: userAddress,
	}
}
