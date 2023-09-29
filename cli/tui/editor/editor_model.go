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

package editor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"unicode"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/ether"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/internal/validator"
	"github.com/deepsquare-io/grid/cli/sbatch"
	"github.com/deepsquare-io/grid/cli/tui/channel"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/mistakenelf/teacup/code"
	"gopkg.in/yaml.v3"
)

const (
	creditsLockingInput = iota
	usesInput
	jobNameInput
)

type keyMap struct {
	EditAgain      key.Binding
	Exit           key.Binding
	NextInput      key.Binding
	PrevInput      key.Binding
	ViewPortKeymap viewport.KeyMap
}

// ExitMsg msg closes to editor model
type ExitMsg struct {
	JobID [32]byte
}

func (m *model) emitExitMsg(jobID [32]byte) tea.Cmd {
	return tea.Sequence(
		m.watchFileChanges.Dispose,
		func() tea.Msg {
			return ExitMsg{
				JobID: jobID,
			}
		},
	)
}

type submitProgressMsg bool

type errorMsg error

type clearErrorsMsg struct{}

func emitClearErrorsMsg() tea.Msg {
	return clearErrorsMsg{}
}

type model struct {
	// Form
	inputs  []textinput.Model
	errors  []error
	focused int

	// Code
	watchFileChanges channel.Model[fileChangedMsg]
	code             code.Model
	jobPath          string

	err error

	help   help.Model
	keyMap keyMap

	client deepsquare.Client

	isSubmitting bool
}

type editorDone struct {
	err     error
	jobPath string
}

func (m *model) openEditor(
	ctx context.Context,
	jobPath string,
) tea.Cmd {
	return tea.ExecProcess(Command(ctx, jobPath), func(err error) tea.Msg {
		return editorDone{
			jobPath: jobPath,
			err:     err,
		}
	})
}

func (m *model) submitJob(ctx context.Context, jobPath string) tea.Cmd {
	return tea.Batch(
		tea.Sequence(
			func() tea.Msg {
				// Read file
				dat, err := os.ReadFile(jobPath)
				if err != nil {
					return errorMsg(err)
				}
				var job sbatch.Job
				if err := yaml.Unmarshal(dat, &job); err != nil {
					return errorMsg(err)
				}

				// Validate input
				if m.inputs[jobNameInput].Value() == "" {
					err := errors.New("empty value is not allowed")
					m.errors[jobNameInput] = err
					return fmt.Errorf("job name field: %w", err)
				}
				if err := validator.IsNumber(m.inputs[creditsLockingInput].Value()); err != nil {
					m.errors[creditsLockingInput] = err
					return fmt.Errorf("allocate credits field: %w", err)
				}
				if err := validator.IsMap(m.inputs[usesInput].Value()); err != nil {
					m.errors[usesInput] = err
					return fmt.Errorf("use flags field: %w", err)
				}

				// Parse input
				allocatedCredits := m.inputs[creditsLockingInput].Value()
				allocatedCreditsF, err := strconv.ParseFloat(allocatedCredits, 64)
				if err != nil {
					return errorMsg(err)
				}
				allocatedCreditsBigF := new(big.Float).SetFloat64(allocatedCreditsF)
				curr, err := m.client.GetAllowance(ctx)
				if err != nil {
					return errorMsg(err)
				}
				labels, err := utils.StringsToLabels(m.inputs[usesInput].Value())
				if err != nil {
					return errorMsg(err)
				}

				// Mutate
				allocatedCreditsBigI := ether.ToWei(allocatedCreditsBigF)
				if err = m.client.SetAllowance(ctx, curr.Add(curr, allocatedCreditsBigI)); err != nil {
					return errorMsg(err)
				}

				var jobName [32]byte
				copy(jobName[:], m.inputs[jobNameInput].Value())
				jobID, err := m.client.SubmitJob(
					ctx,
					&job,
					allocatedCreditsBigI,
					jobName,
					types.WithUse(labels...),
				)
				if err != nil {
					return errorMsg(err)
				}
				return m.emitExitMsg(jobID)()
			},
			func() tea.Msg {
				return submitProgressMsg(false)
			},
		),
		func() tea.Msg {
			return submitProgressMsg(true)
		},
	)
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.openEditor(context.TODO(), m.jobPath),
		m.watchFileChanges.Init(),
		textinput.Blink,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		codeCmd  tea.Cmd
		inputCmd tea.Cmd
		wCmd     tea.Cmd
		cmds     = make([]tea.Cmd, len(m.inputs))
	)

	m.watchFileChanges, wCmd = m.watchFileChanges.Update(msg)
	if wCmd != nil {
		cmds = append(cmds, wCmd)
	}

switchmsg:
	switch msg := msg.(type) {
	case editorDone:
		if msg.err != nil {
			m.err = msg.err
			break switchmsg
		}
		m.jobPath = msg.jobPath
		cmds = append(cmds, m.code.SetFileName(m.jobPath))
	case fileChangedMsg:
		cmds = append(cmds, m.code.SetFileName(m.jobPath))
	case submitProgressMsg:
		m.isSubmitting = bool(msg)
	case errorMsg:
		m.err = msg
		return m, nil
	case tea.KeyMsg:
		if m.isSubmitting {
			break switchmsg
		}
		switch {
		case key.Matches(msg, m.keyMap.EditAgain):
			cmds = append(cmds, m.openEditor(context.TODO(), m.jobPath))
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, m.emitExitMsg([32]byte{}))
		case msg.String() == "enter" && m.focused == len(m.inputs)-1:
			cmds = append(cmds, tea.Sequence(emitClearErrorsMsg, m.submitJob(context.TODO(), m.jobPath)))
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

	m.inputs[creditsLockingInput], inputCmd = m.inputs[creditsLockingInput].Update(msg)
	cmds = append(cmds, inputCmd)

	m.inputs[usesInput], inputCmd = m.inputs[usesInput].Update(msg)
	cmds = append(cmds, inputCmd)

	m.inputs[jobNameInput], inputCmd = m.inputs[jobNameInput].Update(msg)
	cmds = append(cmds, inputCmd)

	m.code, codeCmd = m.code.Update(msg)
	if codeCmd != nil {
		cmds = append(cmds, codeCmd)
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

func isNumberCharacter(ch rune) bool {
	return unicode.IsDigit(ch) || ch == 'e' || ch == '.'
}

func allowedNumber(input string) error {
	for _, ch := range input {
		if !isNumberCharacter(ch) {
			return fmt.Errorf("character '%c' is not allowed", ch)
		}
	}
	return nil
}
