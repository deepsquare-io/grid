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

// Package channel provides utilities to create an Bubbletea model compatible with Go channels.
package channel

import (
	tea "github.com/charmbracelet/bubbletea"
)

type initMsg struct {
	Cancel func() error
}

// ErrorMsg is a signal used to close the channel of the Model.
type ErrorMsg[T any] error

// Model is the object used to implement a bubbletea model with a channel.
type Model[T any] struct {
	Channel chan T
	// OnInit is used to insert data into the channel and initialize the dispose method.
	OnInit func(chan T) func() error

	dispose func() error
}

func (m *Model[T]) init() tea.Msg {
	dispose := m.OnInit(m.Channel)
	if dispose == nil {
		panic("missing dispose function")
	}

	return initMsg{
		Cancel: dispose,
	}
}

func (m *Model[T]) tick() tea.Msg {
	return <-m.Channel
}

// Init emit a signal used to start listening the channel.
func (m Model[T]) Init() tea.Cmd {
	return tea.Batch(
		m.init, m.tick,
	)
}

// Update prepare the dispose method and emit a tick for listening the channel.
func (m Model[T]) Update(msg tea.Msg) (Model[T], tea.Cmd) {
	switch msg := msg.(type) {
	case initMsg:
		m.dispose = msg.Cancel
	case T:
		return m, m.tick
	case ErrorMsg[T]:
		if msg != nil {
			return m, tea.Println(msg)
		}
	}
	return m, nil
}

// Dispose executes the dispose function.
func (m *Model[T]) Dispose() tea.Msg {
	if err := m.dispose(); err != nil {
		return ErrorMsg[T](err)
	}
	return nil
}
