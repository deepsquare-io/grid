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

package topup

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

// ModelBuilder contains the dependencies used to build the bubbletea Model for the topup page.
type ModelBuilder struct {
	Client  deepsquare.Client
	Watcher deepsquare.Watcher
}

// Build the bubbletea Model for the topup page.
func (b *ModelBuilder) Build(jobID [32]byte) tea.Model {
	if b.Client == nil {
		panic("Client is nil")
	}
	if b.Watcher == nil {
		panic("Watcher is nil")
	}
	help := help.New()
	help.ShowAll = true

	inputs := make([]textinput.Model, inputsSize)
	inputs[amountInput] = textinput.New()
	inputs[amountInput].Focus()
	inputs[amountInput].Placeholder = "example: 0.0"
	inputs[amountInput].Width = 64
	inputs[amountInput].Prompt = style.Foreground.Render("❱ ")
	inputs[amountInput].Validate = validator.AllowedNumberChar

	return &model{
		client:   b.Client,
		help:     help,
		inputs:   inputs,
		jobID:    jobID,
		watchJob: makeWatchJobModel(context.TODO(), jobID, b.Watcher, b.Client),
		errors:   make([]error, inputsSize),
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
