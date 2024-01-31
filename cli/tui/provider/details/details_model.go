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

package details

import (
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/internal/ether"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/deepsquare-io/grid/cli/types/provider"
)

const (
	tasksInput = iota
	cpusPerTaskInput
	memPerCPUInput
	gpusInput
	creditsInput

	inputsSize
)

type keyMap struct {
	ViewPortKeyMap viewport.KeyMap
	Exit           key.Binding
	NextInput      key.Binding
	PrevInput      key.Binding
}

// ExitMsg msg closes to details model
type ExitMsg struct{}

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type model struct {
	help     help.Model
	viewport viewport.Model
	keyMap   keyMap

	// Form
	inputs  []textinput.Model
	errors  []error
	focused int
	err     error

	// Intermediary variables
	tasks       uint64
	cpusPerTask uint64
	memPerCPU   uint64
	gpus        uint64
	credits     *big.Float

	// Form result
	duration string

	provider.Detail
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd    tea.Cmd
		inputCmd tea.Cmd
		cmds     = make([]tea.Cmd, 0)
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
		case key.Matches(msg, m.keyMap.NextInput):
			m.nextInput()
		case key.Matches(msg, m.keyMap.PrevInput):
			m.prevInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
		cmds = append(cmds, textinput.Blink)
	}

	for i := range m.inputs {
		m.inputs[i], inputCmd = m.inputs[i].Update(msg)
		cmds = append(cmds, inputCmd)
	}

	var err error
	m.cpusPerTask, err = strconv.ParseUint(m.inputs[cpusPerTaskInput].Value(), 10, 64)
	m.errors[cpusPerTaskInput] = err
	if m.cpusPerTask <= 0 && m.errors[cpusPerTaskInput] == nil {
		m.errors[cpusPerTaskInput] = errors.New("must be greater than 0")
	}
	m.memPerCPU, err = strconv.ParseUint(m.inputs[memPerCPUInput].Value(), 10, 64)
	m.errors[memPerCPUInput] = err
	if m.memPerCPU <= 0 && m.errors[memPerCPUInput] == nil {
		m.errors[memPerCPUInput] = errors.New("must be greater than 0")
	}
	m.gpus, err = strconv.ParseUint(m.inputs[gpusInput].Value(), 10, 64)
	m.errors[gpusInput] = err
	m.tasks, err = strconv.ParseUint(m.inputs[tasksInput].Value(), 10, 64)
	m.errors[tasksInput] = err
	if m.tasks <= 0 && m.errors[tasksInput] == nil {
		m.errors[tasksInput] = errors.New("must be greater than 0")
	}
	var ok bool
	_, ok = m.credits.SetString(m.inputs[creditsInput].Value())
	if !ok {
		m.errors[creditsInput] = errors.New("value couldn't be parsed")
	} else {
		m.errors[creditsInput] = nil
	}
	durationB, err := metascheduler.CreditToDuration(
		m.ProviderPrices,
		metaschedulerabi.JobDefinition{
			Gpus:        m.gpus,
			MemPerCpu:   m.memPerCPU,
			CpusPerTask: m.cpusPerTask,
			Ntasks:      m.tasks,
		},
		ether.ToWei(m.credits),
	)
	m.err = err
	if err != nil {
		m.duration = "NaN"
	} else {
		m.duration = (time.Duration(durationB.Int64()) * time.Minute).String()
	}

	return m, tea.Batch(cmds...)
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
