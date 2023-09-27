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

package log

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/tui/channel"
)

// ExitMsg msg closes to log model
type ExitMsg struct{}

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type model struct {
	viewport  viewport.Model
	spinner   spinner.Model
	messages  *strings.Builder
	logs      []logMsg
	watchLogs channel.Model[logMsg]
	title     string

	showTimestamp bool
}

func (m model) Init() tea.Cmd {
	// TODO: handle termination
	return tea.Batch(
		m.watchLogs.Init(),
		m.spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd    tea.Cmd
		sCmd     tea.Cmd
		lChanCmd tea.Cmd
		cmds     = make([]tea.Cmd, 0)
	)

	m.viewport, vpCmd = m.viewport.Update(msg)
	if vpCmd != nil {
		cmds = append(cmds, vpCmd)
	}
	m.watchLogs, lChanCmd = m.watchLogs.Update(msg)
	if lChanCmd != nil {
		cmds = append(cmds, lChanCmd)
	}

	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, sCmd = m.spinner.Update(msg)
		if sCmd != nil {
			cmds = append(cmds, sCmd)
		}
	case ExitMsg:
		cmds = append(cmds, m.watchLogs.Dispose)
	case logMsg:
		m.logs = append(m.logs, msg)
		if m.showTimestamp {
			m.messages.WriteString(fmt.Sprintf("\n%s: %s", msg.timestamp, msg.message))
		} else {
			m.messages.WriteString("\n" + msg.message)
		}

		m.viewport.SetContent(m.messages.String())
		m.viewport.GotoBottom()
	case tea.KeyMsg:
		switch {
		case msg.Type == tea.KeyEscape, msg.String() == "q":
			cmds = append(cmds, emitExitMsg)
		}
	}
	return m, tea.Batch(cmds...)
}
