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

package editor

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/validator"
	"github.com/deepsquare-io/grid/cli/internal/wordlists"
	"github.com/deepsquare-io/grid/cli/tui/style"
	"github.com/mistakenelf/teacup/code"
)

const jobSchemaPath = "https://raw.githubusercontent.com/deepsquare-io/grid/main/cli/job.schema.json"

// ModelBuilder contains the dependencies used to build the bubbletea Model for the text editor.
type ModelBuilder struct {
	Client *deepsquare.Client
}

func prepareFiles() (words string, jobPath string, err error) {
	words = strings.Join(wordlists.GetRandomWords(3), "-")
	jobPath = fmt.Sprintf("job.%s.yaml", words)

	// Insert the yaml-language-server parameter
	template := []byte(
		fmt.Sprintf(templateFormat, jobSchemaPath, template),
	)

	if err := os.WriteFile(jobPath, template, 0644); err != nil {
		return "", "", fmt.Errorf("fail to write %s: %w", jobPath, err)
	}

	return words, jobPath, nil
}

// Build the bubbletea Model for the text editor.
func (b *ModelBuilder) Build(ctx context.Context) tea.Model {
	if b.Client == nil {
		panic("Client is nil")
	}

	code := code.New(true, true, lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"})
	code.SetSize(80, style.StandardHeight+3)

	help := help.New()
	help.ShowAll = true

	inputs := make([]textinput.Model, inputsSize)
	inputs[creditsLockingInput] = textinput.New()
	inputs[creditsLockingInput].Placeholder = "example: 100"
	inputs[creditsLockingInput].Focus()
	inputs[creditsLockingInput].Width = 32
	inputs[creditsLockingInput].Prompt = style.Foreground().Render("❱ ")
	inputs[creditsLockingInput].Validate = validator.AllowedNumberChar

	inputs[usesInput] = textinput.New()
	inputs[usesInput].Placeholder = "example: os=linux,arch=amd64"
	inputs[usesInput].Width = 32
	inputs[usesInput].Prompt = style.Foreground().Render("❱ ")

	inputs[jobNameInput] = textinput.New()
	inputs[jobNameInput].Width = 32
	inputs[jobNameInput].Prompt = style.Foreground().Render("❱ ")

	jobName, jobPath, err := prepareFiles()
	if err != nil {
		panic(err.Error())
	}
	inputs[jobNameInput].SetValue(jobName)

	return &model{
		context: ctx,
		code:    code,
		inputs:  inputs,
		errors:  make([]error, inputsSize),
		keyMap: keyMap{
			EditAgain: key.NewBinding(
				key.WithKeys("ctrl+e"),
				key.WithHelp("ctrl+e", "edit job"),
			),
			Exit: key.NewBinding(
				key.WithKeys("esc", "ctrl+q"),
				key.WithHelp("esc/ctrl+q", "exit"),
			),
			NextInput: key.NewBinding(
				key.WithKeys("tab", "ctrl+n", "enter"),
				key.WithHelp("tab/enter", "next input/finish"),
			),
			PrevInput: key.NewBinding(
				key.WithKeys("shift+tab", "ctrl+p"),
				key.WithHelp("shift+tab", "prev input"),
			),
			ViewPortKeymap: code.Viewport.KeyMap,
		},
		watchFileChanges: makeWatchFileChangeModel(jobPath),
		jobPath:          jobPath,
		help:             help,
		client:           b.Client,
	}
}
