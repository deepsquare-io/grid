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

package transfer

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
	"github.com/ethereum/go-ethereum/common"
)

const (
	toInput = iota
	amountInput
)

type KeyMap struct {
	NextInput key.Binding
	PrevInput key.Binding
	Exit      key.Binding
}

type errorMsg error

// ExitMsg msg closes to editor model
type ExitMsg struct{}

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type transferDoneMsg struct{}

func (m *model) transfer(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		// Parse input
		if !common.IsHexAddress(m.inputs[toInput].Value()) {
			err := errors.New("field is not an hex address")
			m.errors[amountInput] = err
			return errorMsg(err)
		}
		to := common.HexToAddress(m.inputs[toInput].Value())
		in, ok := new(big.Float).SetString(m.inputs[amountInput].Value())
		if !ok {
			err := errors.New("couldn't parse amount value")
			m.errors[amountInput] = err
			return errorMsg(err)
		}
		amount := ether.ToWei(in)

		if err := m.client.Transfer(ctx, to, amount); err != nil {
			return errorMsg(err)
		}
		return transferDoneMsg{}
	}
}

type model struct {
	help help.Model

	client deepsquare.Client

	inputs  []textinput.Model
	errors  []error
	focused int

	keyMap KeyMap

	err error
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		inputCmd tea.Cmd
		cmds     = make([]tea.Cmd, 0)
	)
	switch msg := msg.(type) {
	case errorMsg:
		m.err = msg
		return m, nil
	case transferDoneMsg:
		cmds = append(cmds, emitExitMsg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, emitExitMsg)
		case key.Matches(msg, m.keyMap.NextInput):
			m.nextInput()
		case key.Matches(msg, m.keyMap.PrevInput):
			m.prevInput()
		case msg.String() == "enter" && m.focused == len(m.inputs)-1:
			cmds = append(cmds, m.transfer(context.TODO()))
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()
		cmds = append(cmds, textinput.Blink)
	}

	m.inputs[toInput], inputCmd = m.inputs[toInput].Update(msg)
	cmds = append(cmds, inputCmd)

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

func allowedHex(input string) error {
	for _, ch := range input {
		if !unicode.Is(unicode.Hex_Digit, ch) {
			return fmt.Errorf("character '%c' is not allowed", ch)
		}
	}
	return nil
}
