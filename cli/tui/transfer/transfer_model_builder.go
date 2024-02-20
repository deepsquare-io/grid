// Copyright (C) 2024 DeepSquare Association
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
	"context"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/validator"
	"github.com/deepsquare-io/grid/cli/tui/style"
)

// ModelBuilder contains the dependencies to build the model for the transfer page.
type ModelBuilder struct {
	Client *deepsquare.Client
}

// Build the bubbletea model for the transfer page.
func (b *ModelBuilder) Build(ctx context.Context) tea.Model {
	if b.Client == nil {
		panic("Client is nil")
	}
	help := help.New()
	help.ShowAll = true

	inputs := make([]textinput.Model, inputsSize)
	inputs[toInput] = textinput.New()
	inputs[toInput].Focus()
	inputs[toInput].Placeholder = "example: 0x0000000000000000000000000000000000000000"
	inputs[toInput].Width = 64
	inputs[toInput].Prompt = style.Foreground().Render("❱ ")
	inputs[toInput].Validate = validator.AllowedHexChar

	inputs[amountInput] = textinput.New()
	inputs[amountInput].Placeholder = "example: 0.0"
	inputs[amountInput].Width = 64
	inputs[amountInput].Prompt = style.Foreground().Render("❱ ")
	inputs[amountInput].Validate = validator.AllowedNumberChar

	return &model{
		context: ctx,
		client:  b.Client,
		help:    help,
		inputs:  inputs,
		errors:  make([]error, inputsSize),
		keyMap: keyMap{
			Exit: key.NewBinding(
				key.WithKeys("esc", "ctrl+q"),
				key.WithHelp("esc/ctrl+q", "exit"),
			),
			NextInput: key.NewBinding(
				key.WithKeys("tab", "enter"),
				key.WithHelp("tab/enter", "next input/finish"),
			),
			PrevInput: key.NewBinding(
				key.WithKeys("shift+tab", "ctrl+p"),
				key.WithHelp("shift+tab", "prev input"),
			),
		},
	}
}
