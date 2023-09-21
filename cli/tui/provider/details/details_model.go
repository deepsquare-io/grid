package details

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/types"
)

type KeyMap struct {
	ViewPortKeyMap viewport.KeyMap
	Exit           key.Binding
}

// ExitMsg msg closes to details model
type ExitMsg struct{}

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type model struct {
	help     help.Model
	viewport viewport.Model
	keyMap   KeyMap

	types.ProviderDetail
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd tea.Cmd
		cmds  = make([]tea.Cmd, 0)
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
		}
	}
	return m, tea.Batch(cmds...)
}
