package details

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/internal/utils"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/deepsquare-io/the-grid/cli/types"
	"gopkg.in/yaml.v3"
)

var titleStyle = style.Title1.Width(10)

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func Model(p types.ProviderDetail) tea.Model {
	vp := viewport.New(118, style.StandardHeight)
	prices := fmt.Sprintf(`CPU pricing: %s creds/(CPU.min)
Memory pricing: %s creds/(MB.min)
GPU pricing: %s creds/(GPU.min)
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

	return &model{
		ProviderDetail: p,
		help:           help,
		viewport:       vp,
		keyMap: KeyMap{
			ViewPortKeyMap: vp.KeyMap,
			Exit: key.NewBinding(
				key.WithKeys("esc", "q"),
				key.WithHelp("esc/q", "exit"),
			),
		},
	}
}
