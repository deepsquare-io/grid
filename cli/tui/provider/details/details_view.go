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

package details

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/tui/style"
)

func (m model) headerView() string {
	return style.LogTitle.Render(m.ProviderDetail.Addr.Hex())
}

func (m model) View() string {
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.ViewPortKeyMap.Up,
			m.keyMap.ViewPortKeyMap.Down,
		},
		{
			m.keyMap.NextInput,
			m.keyMap.PrevInput,
			m.keyMap.Exit,
		},
	})
	emptyLine := style.Foreground.Render(
		strings.Repeat(" ", max(0, m.viewport.Width/2)),
	)
	view := fmt.Sprintf(
		"%s\n%s\n%s",
		m.headerView(), m.viewport.View(), emptyLine,
	)
	form := fmt.Sprintf(
		`%s
%s %s %s
%s %s %s
%s %s %s
%s %s %s
%s %s %s

%s %s
%s
%s`,
		style.LogTitle.Render("Cost Calculator"),
		style.Foreground.Render("Tasks:"),
		m.inputs[tasksInput].View(),
		style.Error.Width(m.viewport.Width/2).
			Render(utils.ErrorfOrEmpty("\n^^^%s", m.errors[tasksInput])),
		style.Foreground.Render("CPUs per Task:"),
		m.inputs[cpusPerTaskInput].View(),
		style.Error.Width(m.viewport.Width/2).
			Render(utils.ErrorfOrEmpty("\n^^^%s", m.errors[cpusPerTaskInput])),
		style.Foreground.Render("Memory (MB) per CPU:"),
		m.inputs[memPerCPUInput].View(),
		style.Error.Width(m.viewport.Width/2).
			Render(utils.ErrorfOrEmpty("\n^^^%s", m.errors[memPerCPUInput])),
		style.Foreground.Render("GPUs Per Task:"),
		m.inputs[gpusPerTaskInput].View(),
		style.Error.Width(m.viewport.Width/2).
			Render(utils.ErrorfOrEmpty("\n^^^%s", m.errors[gpusPerTaskInput])),
		style.Foreground.Render("Allocated credits:"),
		m.inputs[creditsInput].View(),
		style.Error.Width(m.viewport.Width/2).
			Render(utils.ErrorfOrEmpty("\n^^^%s", m.errors[creditsInput])),
		"Expected Max Duration: ",
		m.duration,
		style.Error.Width(20).Render(utils.ErrorfOrEmpty("%s", m.err)),
		emptyLine,
	)
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		style.Box.Render(view),
		style.Box.Render(form),
		help,
	)
}
