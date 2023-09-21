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
