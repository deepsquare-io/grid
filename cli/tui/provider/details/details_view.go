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

package details

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
)

func (m model) headerView() string {
	return style.LogTitle.Render(m.ProviderDetail.Addr.Hex())
}

func (m model) View() string {
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.Exit,
		},
		{
			m.keyMap.ViewPortKeyMap.Up,
			m.keyMap.ViewPortKeyMap.Down,
		},
	})
	emptyLine := style.Foreground.Render(
		strings.Repeat(" ", max(0, m.viewport.Width)),
	)
	view := fmt.Sprintf(
		"%s\n%s\n%s",
		m.headerView(), m.viewport.View(), emptyLine,
	)
	return lipgloss.JoinHorizontal(lipgloss.Center, style.Box.Render(view), help)
}
