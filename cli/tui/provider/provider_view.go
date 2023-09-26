// Copyright (C) 2023 DeepSquare Asociation
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
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/grid/cli/tui/style"
)

var columns = []table.Column{
	{Title: "Provider Address", Width: 16},
	{Title: "Is Schedulable", Width: 14},
	{Title: "Is Waiting For Approval", Width: 23},
	{Title: "Is Banned", Width: 9},
	{Title: "Prices", Width: 6},
	{Title: "Specifications", Width: 14},
	{Title: "Labels", Width: 6},
}

func (m model) View() string {
	if m.detailsModel != nil {
		return m.detailsModel.View()
	}
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.ShowProviderDetails,
			m.keyMap.Exit,
		},
		{
			m.keyMap.TableKeyMap.LineUp,
			m.keyMap.TableKeyMap.LineDown,
		},
	})
	return lipgloss.JoinHorizontal(lipgloss.Center, style.Base.Render(m.table.View()), help)
}
