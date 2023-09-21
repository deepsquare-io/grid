package transfer

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/deepsquare-io/the-grid/cli/internal/utils"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
)

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
%s`,
		style.Foreground.Render("Send to"),
		m.inputs[toInput].View(),
		style.Error.Render(utils.ErrorOrEmpty(m.errors[toInput])),
		style.Foreground.Render("Amount"),
		m.inputs[amountInput].View(),
		style.Error.Render(utils.ErrorOrEmpty(m.errors[amountInput])),
		style.Error.Render(utils.ErrorOrEmpty(m.err)),
		help,
	)
}
