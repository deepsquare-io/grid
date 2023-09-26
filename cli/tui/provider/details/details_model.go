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
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/types"
)

type KeyMap struct {
	ViewPortKeyMap viewport.KeyMap
	Exit           key.Binding
}

// ExitMsg msg closes to details model
type ExitMsg struct{}

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type model struct {
	help     help.Model
	viewport viewport.Model
	keyMap   KeyMap

	types.ProviderDetail
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd tea.Cmd
		cmds  = make([]tea.Cmd, 0)
	)

	m.viewport, vpCmd = m.viewport.Update(msg)
	if vpCmd != nil {
		cmds = append(cmds, vpCmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, emitExitMsg)
		}
	}
	return m, tea.Batch(cmds...)
}
