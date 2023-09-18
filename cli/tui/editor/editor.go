// Copyright (C) 2023 DeepSquare
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
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/utils"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/mistakenelf/teacup/code"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m model) viewportFooterView() string {
	info := style.LogInfo.Render(fmt.Sprintf("%3.f%%", m.code.Viewport.ScrollPercent()*100))
	line := style.Foreground.Render(
		strings.Repeat(" ", max(0, m.code.Viewport.Width-lipgloss.Width(info))),
	)
	return lipgloss.JoinHorizontal(lipgloss.Bottom, line, info)
}

func (m model) formView() string {
	return fmt.Sprintf(`%s
%s
%s

%s
%s
%s

%s
%s
%s
%s
`,
		style.Foreground.Render("Allocate Credits"),
		m.inputs[creditsLockingInput].View(),
		style.Error.Render(utils.ErrorOrEmpty(m.errors[creditsLockingInput])),
		style.Foreground.Render("Use flags"),
		m.inputs[usesInput].View(),
		style.Error.Render(utils.ErrorOrEmpty(m.errors[usesInput])),
		style.Foreground.Render("Job Name"),
		m.inputs[jobNameInput].View(),
		style.Error.Render(utils.ErrorOrEmpty(m.errors[jobNameInput])),
		style.Error.Render(utils.ErrorOrEmpty(m.err)),
	)
}

func (m model) View() string {
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.ViewPortKeymap.Up,
			m.keyMap.ViewPortKeymap.Down,
		},
		{
			m.keyMap.NextInput,
			m.keyMap.PrevInput,
			m.keyMap.EditAgain,
			m.keyMap.Exit,
		},
	})
	leftView := fmt.Sprintf(
		"%s\n%s",
		m.code.View(), m.viewportFooterView(),
	)
	rightView := m.formView() + "\n" + help
	mainView := lipgloss.JoinHorizontal(lipgloss.Center, style.Box.Render(leftView), rightView)

	return mainView
}

func Model(
	client deepsquare.Client,
) tea.Model {
	code := code.New(true, true, lipgloss.AdaptiveColor{Light: "#000000", Dark: "#ffffff"})
	code.SetSize(118, style.StandardHeight)

	help := help.New()
	help.ShowAll = true

	inputs := make([]textinput.Model, 3)
	inputs[creditsLockingInput] = textinput.New()
	inputs[creditsLockingInput].Placeholder = "100"
	inputs[creditsLockingInput].Focus()
	inputs[creditsLockingInput].Width = 32
	inputs[creditsLockingInput].Prompt = ""
	inputs[creditsLockingInput].Validate = allowedNumber

	inputs[usesInput] = textinput.New()
	inputs[usesInput].Placeholder = "os=linux,arch=amd64"
	inputs[usesInput].Width = 32
	inputs[usesInput].Prompt = ""

	inputs[jobNameInput] = textinput.New()
	inputs[jobNameInput].Width = 32
	inputs[jobNameInput].Prompt = ""

	return &model{
		code:   code,
		inputs: inputs,
		errors: make([]error, 3),
		keyMap: KeyMap{
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
		help:   help,
		client: client,
	}
}
