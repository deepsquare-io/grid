// Copyright (C) 2023 DeepSquare
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

package status

import (
	"context"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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
			m.keyMap.SubmitJob,
			m.keyMap.CancelJob,
			m.keyMap.TransferCredits,
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
	client deepsquare.Client,
	watcher deepsquare.Watcher,
	userAddress common.Address,
) tea.Model {
	// Initialize rows
	rows, idToRow, it := initializeRows(ctx, client)

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
			TransferCredits: key.NewBinding(
				key.WithKeys("t"),
				key.WithHelp("t", "tranfer credits"),
			),
			Exit: key.NewBinding(
				key.WithKeys("esc", "q"),
				key.WithHelp("esc/q", "exit"),
			),
		},
		scheduler: client,
		watchJobs: makeWatchJobsModel(
			ctx,
			userAddress,
			make(chan ethtypes.Log, 100),
			watcher,
			client,
		),
	}
}
