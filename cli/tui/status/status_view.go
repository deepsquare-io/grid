// Copyright (C) 2023 DeepSquare Association
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
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/tui/components/table"
	"github.com/deepsquare-io/grid/cli/tui/style"
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
			m.keyMap.TopupJob,
			m.keyMap.CancelJob,
			m.keyMap.TransferCredits,
			m.keyMap.ViewProviders,
			m.keyMap.Exit,
		},
		{
			m.keyMap.TableKeyMap.LineUp,
			m.keyMap.TableKeyMap.LineDown,
		},
	})
	var status string
	if m.isCancelling {
		status += "Cancelling..."
	}
	status += style.Error.Width(30).Render(utils.ErrorfOrEmpty("Error: %s", m.err))
	right := fmt.Sprintf(
		"%s\n%s",
		status,
		help,
	)
	return lipgloss.JoinHorizontal(lipgloss.Center, style.Base.Render(m.table.View()), right)
}
