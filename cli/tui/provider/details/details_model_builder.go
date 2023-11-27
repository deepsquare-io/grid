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

package details

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/internal/ether"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/internal/validator"
	"github.com/deepsquare-io/grid/cli/tui/style"
	"github.com/deepsquare-io/grid/cli/types"
	"gopkg.in/yaml.v3"
)

var titleStyle = style.Title1().Width(10)

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

// Model is used to build the bubbletea Model for showing the provider details.
func Model(p types.ProviderDetail) tea.Model {
	vp := viewport.New(118, style.StandardHeight)
	prices := fmt.Sprintf(`CPU pricing: %s credits/(CPU.min)
Memory pricing: %s credits/(MB.min)
GPU pricing: %s credits/(GPU.min)
`,
		ether.FromWei(p.ProviderPrices.CpuPricePerMin).String(),
		ether.FromWei(p.ProviderPrices.MemPricePerMin).String(),
		ether.FromWei(p.ProviderPrices.GpuPricePerMin).String(),
	)
	hardware := fmt.Sprintf(`Nodes: %d
CPU per node: %v
Mem(MB) per node: %v
GPU per node: %v
`,
		p.ProviderHardware.Nodes,
		p.ProviderHardware.CpusPerNode,
		p.ProviderHardware.MemPerNode,
		p.ProviderHardware.GpusPerNode,
	)
	labels, err := yaml.Marshal(utils.FormatLabels(p.Labels))
	if err != nil {
		panic(fmt.Sprintf("failed to marshal labels: %s", err.Error()))
	}
	statuses := fmt.Sprintf(`Is valid for scheduling: %s
Is waiting for approval: %s
Is banned: %s
`,
		style.BoolToYN(p.IsValidForScheduling),
		style.BoolToYNColorReverted(p.IsWaitingForApproval),
		style.BoolToYNColorReverted(p.IsBanned))
	vp.SetContent(fmt.Sprintf(`%s
%s
%s
%s
%s
%s
%s
%s
`,
		titleStyle.Render("Statuses"),
		indent(2, statuses),
		titleStyle.Render("Pricing"),
		indent(2, prices),
		titleStyle.Render("Hardware"),
		indent(2, hardware),
		titleStyle.Render("Labels"),
		indent(2, string(labels)),
	))

	help := help.New()
	help.ShowAll = true

	inputs := make([]textinput.Model, inputsSize)
	inputs[tasksInput] = textinput.New()
	inputs[tasksInput].Focus()
	inputs[tasksInput].Placeholder = "example: 1"
	inputs[tasksInput].Width = 10
	inputs[tasksInput].Prompt = style.Foreground().Render("❱ ")
	inputs[tasksInput].Validate = validator.AllowedNumberChar
	inputs[tasksInput].SetValue("1")

	inputs[cpusPerTaskInput] = textinput.New()
	inputs[cpusPerTaskInput].Placeholder = "example: 1"
	inputs[cpusPerTaskInput].Width = 10
	inputs[cpusPerTaskInput].Prompt = style.Foreground().Render("❱ ")
	inputs[cpusPerTaskInput].Validate = validator.AllowedNumberChar
	inputs[cpusPerTaskInput].SetValue("1")

	inputs[memPerCPUInput] = textinput.New()
	inputs[memPerCPUInput].Placeholder = "example: 200"
	inputs[memPerCPUInput].Width = 10
	inputs[memPerCPUInput].Prompt = style.Foreground().Render("❱ ")
	inputs[memPerCPUInput].Validate = validator.AllowedNumberChar
	inputs[memPerCPUInput].SetValue("200")

	inputs[gpusInput] = textinput.New()
	inputs[gpusInput].Placeholder = "example: 1"
	inputs[gpusInput].Width = 10
	inputs[gpusInput].Prompt = style.Foreground().Render("❱ ")
	inputs[gpusInput].Validate = validator.AllowedNumberChar
	inputs[gpusInput].SetValue("0")

	inputs[creditsInput] = textinput.New()
	inputs[creditsInput].Placeholder = "example: 100"
	inputs[creditsInput].Width = 10
	inputs[creditsInput].Prompt = style.Foreground().Render("❱ ")
	inputs[creditsInput].Validate = validator.AllowedNumberChar
	inputs[creditsInput].SetValue("100")

	return &model{
		ProviderDetail: p,
		help:           help,
		viewport:       vp,
		inputs:         inputs,
		credits:        new(big.Float),
		errors:         make([]error, inputsSize),
		keyMap: keyMap{
			ViewPortKeyMap: vp.KeyMap,
			Exit: key.NewBinding(
				key.WithKeys("esc", "ctrl+q"),
				key.WithHelp("esc/ctrl+q", "exit"),
			),
			NextInput: key.NewBinding(
				key.WithKeys("tab", "ctrl+n", "enter"),
				key.WithHelp("tab/enter", "next input"),
			),
			PrevInput: key.NewBinding(
				key.WithKeys("shift+tab", "ctrl+p"),
				key.WithHelp("shift+tab", "prev input"),
			),
		},
	}
}
