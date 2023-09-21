package transfer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
)

type ModelBuilder struct {
	Client deepsquare.Client
}

func (b *ModelBuilder) Build() tea.Model {
	if b.Client == nil {
		panic("Client is nil")
	}
	help := help.New()
	help.ShowAll = true

	inputs := make([]textinput.Model, 2)
	inputs[toInput] = textinput.New()
	inputs[toInput].Focus()
	inputs[toInput].Placeholder = "0x"
	inputs[toInput].Width = 64
	inputs[toInput].Prompt = ""
	inputs[toInput].Validate = allowedHex

	inputs[amountInput] = textinput.New()
	inputs[amountInput].Placeholder = "0.0"
	inputs[amountInput].Width = 64
	inputs[amountInput].Prompt = ""
	inputs[amountInput].Validate = allowedNumber

	return &model{
		client: b.Client,
		help:   help,
		inputs: inputs,
		errors: make([]error, 2),
		keyMap: KeyMap{
			Exit: key.NewBinding(
				key.WithKeys("esc", "ctrl+q"),
				key.WithHelp("esc/ctrl+q", "exit"),
			),
			NextInput: key.NewBinding(
				key.WithKeys("tab", "ctrl+n", "enter"),
				key.WithHelp("tab/enter", "next input/finish"),
			),
			PrevInput: key.NewBinding(
				key.WithKeys("shift+tab", "ctrl+p"),
				key.WithHelp("shift+tab", "prev input"),
			),
		},
	}
}
