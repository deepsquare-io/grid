package channel

import tea "github.com/charmbracelet/bubbletea"

type initMsg struct {
	Cancel func() error
}

type DisposeMsg[T any] struct{}

type Model[T any] struct {
	Channel chan T
	// OnInit is used to insert data into the channel and initialize the dispose method.
	OnInit func(chan T) func() error

	dispose func() error
}

func (m *Model[T]) init() tea.Msg {
	dispose := m.OnInit(m.Channel)

	return initMsg{
		Cancel: dispose,
	}
}

func (m *Model[T]) tick() tea.Msg {
	return <-m.Channel
}

func (m Model[T]) Init() tea.Cmd {
	return tea.Batch(
		m.init, m.tick,
	)
}

func (m Model[T]) Update(msg tea.Msg) (Model[T], tea.Cmd) {
	switch msg := msg.(type) {
	case initMsg:
		m.dispose = msg.Cancel
	case T:
		return m, m.tick
	case DisposeMsg[T]:
		if m.dispose != nil {
			if err := m.dispose(); err != nil {
				return m, tea.Println(err)
			}
		}
	}
	return m, nil
}

func (m *Model[T]) Dispose() tea.Msg {
	return DisposeMsg[T]{}
}
