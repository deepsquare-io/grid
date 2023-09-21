package provider

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/deepsquare-io/the-grid/cli/types"
)

type ModelBuilder struct {
	Client deepsquare.Client
}

func (b *ModelBuilder) Build() tea.Model {
	if b.Client == nil {
		panic("Client is nil")
	}

	tableKeymap := table.DefaultKeyMap()
	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{
			{
				"Loading...",
				"Loading...",
				"Loading...",
				"Loading...",
				"Loading...",
				"Loading...",
				"Loading...",
			},
		}),
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
		table: t,
		help:  help,
		keyMap: KeyMap{
			TableKeyMap: tableKeymap,
			ShowProviderDetails: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "show details"),
			),
			Exit: key.NewBinding(
				key.WithKeys("esc", "q"),
				key.WithHelp("esc/q", "exit"),
			),
		},
		providers: make(map[string]types.ProviderDetail),
		client:    b.Client,
	}
}
