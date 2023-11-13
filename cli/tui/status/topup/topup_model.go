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
	"errors"
	"math/big"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/ether"
	"github.com/deepsquare-io/grid/cli/tui/channel"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum/common"
)

const (
	amountInput = iota

	inputsSize
)

type keyMap struct {
	NextInput key.Binding
	PrevInput key.Binding
	Exit      key.Binding
}

type errorMsg error

type clearErrorsMsg struct{}

func emitClearErrorsMsg() tea.Msg {
	return clearErrorsMsg{}
}

// ExitMsg msg closes to topup model
type ExitMsg struct{}

func (m *model) emitExitMsg() tea.Cmd {
	return tea.Sequence(
		m.watchJob.Dispose,
		func() tea.Msg {
			return ExitMsg{}
		},
	)
}

type topupProgressMsg bool

func (m *model) topup(ctx context.Context) tea.Cmd {
	return tea.Batch(
		tea.Sequence(
			func() tea.Msg {
				// Parse input
				in, ok := new(big.Float).SetString(m.inputs[amountInput].Value())
				if !ok {
					err := errors.New("couldn't parse amount value")
					m.errors[amountInput] = err
					return errorMsg(err)
				}
				amount := ether.ToWei(in)

				curr, err := m.client.GetAllowance(ctx)
				if err != nil {
					return err
				}
				if err = m.client.SetAllowance(ctx, curr.Add(curr, amount)); err != nil {
					return err
				}

				if err := m.client.TopUpJob(ctx, m.jobID, amount); err != nil {
					return errorMsg(err)
				}
				return m.emitExitMsg()()
			},
			func() tea.Msg {
				return topupProgressMsg(false)
			},
		),
		func() tea.Msg {
			return topupProgressMsg(true)
		},
	)
}

type updateProviderMsg metaschedulerabi.Provider

func (m *model) loadProvider(ctx context.Context, provider common.Address) tea.Cmd {
	return func() tea.Msg {
		p, err := m.client.GetProvider(ctx, provider)
		if err != nil {
			return errorMsg(err)
		}
		return updateProviderMsg(p.Provider)
	}
}

type model struct {
	help help.Model

	client deepsquare.Client

	jobID   [32]byte
	inputs  []textinput.Model
	errors  []error
	focused int

	watchJob channel.Model[transitionMsg]

	keyMap keyMap

	job      types.Job
	provider *metaschedulerabi.Provider

	err         error
	isToppingUp bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.watchJob.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		inputCmd tea.Cmd
		wCmd     tea.Cmd
		cmds     = make([]tea.Cmd, 0)
	)

	m.watchJob, wCmd = m.watchJob.Update(msg)
	if wCmd != nil {
		cmds = append(cmds, wCmd)
	}

	switch msg := msg.(type) {
	case updateProviderMsg:
		p := metaschedulerabi.Provider(msg)
		m.provider = &p
	case transitionMsg:
		m.job = types.Job(msg)

		// Fetch provider prices
		if (m.job.ProviderAddr != common.Address{}) {
			cmds = append(cmds, m.loadProvider(context.TODO(), m.job.ProviderAddr))
		}
	case errorMsg:
		m.err = msg
		return m, nil
	case clearErrorsMsg:
		m.errors[amountInput] = nil
		m.err = nil
	case topupProgressMsg:
		m.isToppingUp = bool(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, m.emitExitMsg())
		case msg.String() == "enter" && m.focused == len(m.inputs)-1:
			cmds = append(cmds, tea.Sequence(emitClearErrorsMsg, m.topup(context.TODO())))
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
