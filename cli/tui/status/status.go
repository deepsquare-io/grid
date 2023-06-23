package status

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var columns = []table.Column{
	{Title: "Job ID", Width: 6},
	{Title: "Job Name", Width: 30},
	{Title: "Status", Width: 14},
	{Title: "Start date", Width: 30},
}

func (m model) View() string {
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.OpenLogs,
			m.keyMap.CancelJob,
			m.keyMap.SubmitJob,
			m.keyMap.Exit,
		},
		{
			m.keyMap.TableKeyMap.LineUp,
			m.keyMap.TableKeyMap.LineDown,
		},
	})
	return lipgloss.JoinHorizontal(lipgloss.Center, style.Base.Render(m.table.View()), help)
}

func Model(
	ctx context.Context,
	eventSubscriber cli.EventSubscriber,
	jobFetcher cli.JobFetcher,
	jobFilterer cli.JobFilterer,
	scheduler cli.JobScheduler,
	userAddress common.Address,
) tea.Model {
	// Initialize rows
	rows, idToRow, it := initializeRows(ctx, jobFetcher)

	tableKeymap := table.DefaultKeyMap()
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(style.StandardHeight),
		table.WithKeyMap(tableKeymap),
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

	return &model{
		table:   t,
		idToRow: idToRow,
		it:      it,
		help:    help,
		keyMap: KeyMap{
			TableKeyMap: tableKeymap,
			OpenLogs: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "show job logs"),
			),
			CancelJob: key.NewBinding(
				key.WithKeys("c"),
				key.WithHelp("c", "cancel job"),
			),
			SubmitJob: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "submit job"),
			),
			Exit: key.NewBinding(
				key.WithKeys("esc", "q"),
				key.WithHelp("esc/q", "exit"),
			),
		},
		scheduler: scheduler,
		watchJobs: makeWatchJobsModel(
			ctx,
			userAddress,
			make(chan types.Log, 100),
			eventSubscriber,
			jobFilterer,
			jobFetcher,
		),
	}
}
