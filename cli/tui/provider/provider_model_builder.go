// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package provider

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/tui/components/table"
	"github.com/deepsquare-io/grid/cli/tui/style"
	"github.com/deepsquare-io/grid/cli/types/provider"
)

// ModelBuilder contains the dependencies used to build the bubbletea Model for the provider page.
type ModelBuilder struct {
	Client deepsquare.Client
}

// Build the bubbletea Model for the provider page.
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
	s.RenderCell = func(model table.Model, value string, rawValue string, position table.CellPosition) string {
		if strings.Contains(rawValue, "Loading") {
			return value
		}
		switch position.Column {
		// Is Schedulable
		case 1:
			if utils.YNToBool(rawValue) {
				return style.NoError().Render(value)
			}
			return style.Error().Render(value)
		case 2, 3:
			if utils.YNToBool(rawValue) {
				return style.Error().Render(value)
			}
			return style.NoError().Render(value)
		}
		return value
	}
	t.SetStyles(s)

	help := help.New()
	help.ShowAll = true

	return &model{
		table: t,
		help:  help,
		keyMap: keyMap{
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
		providers: make(map[string]provider.Detail),
		client:    b.Client,
	}
}
