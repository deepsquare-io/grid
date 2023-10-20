package ticker

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Msg time.Time

type Model struct {
	Ticker *time.Ticker
}

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
