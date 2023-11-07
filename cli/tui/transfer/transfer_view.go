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

package transfer

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/tui/style"
)

func (m model) loading() string {
	if m.isTransferring {
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
		style.Title1().Width(20).Render("Transfer credits"),
		style.Foreground().Render("Send to"),
		m.inputs[toInput].View(),
		style.Error().Width(50).Render(utils.FormatErrorfOrEmpty("^^^%s", m.errors[toInput])),
		style.Foreground().Render("Amount in credits (not in wei)"),
		m.inputs[amountInput].View(),
		style.Error().Width(50).Render(utils.FormatErrorfOrEmpty("^^^%s", m.errors[amountInput])),
		style.Error().Width(50).Render(utils.FormatErrorfOrEmpty("Error: %s", m.err)),
		m.loading(),
		help,
	)
}
