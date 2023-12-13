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

package log

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/grid/cli/tui/style"
)

func (m model) headerView() string {
	return style.LogTitle().Render(m.title)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m model) footerView() string {
	info := style.LogInfo().Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := style.Foreground().Render(
		strings.Repeat(" ", max(0, m.viewport.Width-lipgloss.Width(info))),
	)
	return lipgloss.JoinHorizontal(lipgloss.Bottom, line, info)
}

func (m model) View() string {
	var view string
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.Exit,
		},
		{
			m.keyMap.ViewPort.Up,
			m.keyMap.ViewPort.Down,
		},
	})
	if len(m.logs) == 0 {
		view = fmt.Sprintf(
			"%s\nWaiting for logs... %s%s\n%s",
			m.headerView(),
			m.spinner.View(),
			m.viewport.View(),
			m.footerView(),
		)
	} else {
		view = fmt.Sprintf(
			"%s\n%s\n%s",
			m.headerView(),
			m.viewport.View(),
			m.footerView(),
		)
	}
	return style.Box().
		BorderTop(true).
		BorderBottom(true).
		BorderLeft(false).
		BorderRight(false).
		Render(view) +
		"\n" + help
}
