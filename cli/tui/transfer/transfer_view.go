package transfer

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/deepsquare-io/the-grid/cli/internal/utils"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
)

func (m model) loading() string {
	if m.isRunning {
		return "Transferring..."
	}
	return ""
}

func (m model) View() string {
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.NextInput,
			m.keyMap.PrevInput,
			m.keyMap.Exit,
		},
	})
	return fmt.Sprintf(`%s
%s
%s
%s

%s
%s
%s

%s
%s
%s`,
		style.Title1.Width(20).Render("Transfer credits"),
		style.Foreground.Render("Send to"),
		m.inputs[toInput].View(),
		style.Error.Render(utils.ErrorfOrEmpty("^^^%s", m.errors[toInput])),
		style.Foreground.Render("Amount in credits (not in wei)"),
		m.inputs[amountInput].View(),
		style.Error.Render(utils.ErrorfOrEmpty("^^^%s", m.errors[amountInput])),
		style.Error.Render(utils.ErrorfOrEmpty("Error: %s", m.err)),
		m.loading(),
		help,
	)
}
