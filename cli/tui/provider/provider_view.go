package provider

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
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
