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

// Package ticker contains the Model and Update function to work with a Go Ticker.
package ticker

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Msg is the ticking message that must be passed to the Init or Update function.
type Msg time.Time

// Model is the bubbletea model storing the data of the Ticker.
type Model struct {
	Ticker *time.Ticker
}

// Tick passes the Msg when the ticker is triggered.
func (m Model) Tick() tea.Msg {
	return Msg(<-m.Ticker.C)
}

// Update prepare the dispose method and emit a tick for listening the channel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if _, ok := msg.(Msg); ok {
		return m, m.Tick
	}
	return m, nil
}
