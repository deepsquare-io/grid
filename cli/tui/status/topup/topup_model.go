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

package topup

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"unicode"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/tui/channel"
	"github.com/deepsquare-io/the-grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum/common"
)

const (
	amountInput = iota
)

type KeyMap struct {
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

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type topupProgressMsg struct{}

func emitTransferProgressMsg() tea.Msg {
	return topupProgressMsg{}
}

type topupDoneMsg struct{}

func (m *model) topup(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		// Parse input
		in, ok := new(big.Float).SetString(m.inputs[amountInput].Value())
		if !ok {
			err := errors.New("couldn't parse amount value")
			m.errors[amountInput] = err
			return errorMsg(err)
		}
		amount := ether.ToWei(in)

		if err := m.client.TopUpJob(ctx, m.jobID, amount); err != nil {
			return errorMsg(err)
		}
		return topupDoneMsg{}
	}
}

type updateProviderPricesMsg metaschedulerabi.ProviderPrices

func (m *model) loadProviderPrices(ctx context.Context, provider common.Address) tea.Cmd {
	return func() tea.Msg {
		p, err := m.client.GetProvider(ctx, m.job.ProviderAddr)
		if err != nil {
			return errorMsg(err)
		}
		return updateProviderPricesMsg(p.ProviderPrices)
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

	keyMap KeyMap

	job    types.Job
	prices *metaschedulerabi.ProviderPrices

	err       error
	isRunning bool
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
	case ExitMsg:
		cmds = append(cmds, m.watchJob.Dispose)
	case updateProviderPricesMsg:
		p := metaschedulerabi.ProviderPrices(msg)
		m.prices = &p
	case transitionMsg:
		m.job = types.Job(msg)

		// Fetch provider prices
		if (m.job.ProviderAddr != common.Address{}) {
			cmds = append(cmds, m.loadProviderPrices(context.TODO(), m.job.ProviderAddr))
		}
	case errorMsg:
		m.err = msg
		m.isRunning = false
		return m, nil
	case clearErrorsMsg:
		m.errors[amountInput] = nil
		m.err = nil
		m.isRunning = false
	case topupDoneMsg:
		m.isRunning = false
		cmds = append(cmds, emitExitMsg)
	case topupProgressMsg:
		m.isRunning = true
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, emitExitMsg)
		case msg.String() == "enter" && m.focused == len(m.inputs)-1:
			cmds = append(cmds, tea.Sequence(emitClearErrorsMsg, emitTransferProgressMsg, m.topup(context.TODO())))
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

	m.inputs[amountInput], inputCmd = m.inputs[amountInput].Update(msg)
	cmds = append(cmds, inputCmd)

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
