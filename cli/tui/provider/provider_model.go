package provider

import (
	"context"
	"encoding/json"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/internal/utils"
	"github.com/deepsquare-io/the-grid/cli/tui/provider/details"
	"github.com/deepsquare-io/the-grid/cli/types"
	"go.uber.org/zap"
)

type KeyMap struct {
	TableKeyMap         table.KeyMap
	ShowProviderDetails key.Binding
	Exit                key.Binding
}

// ExitMsg msg closes to transfer model
type ExitMsg struct{}

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type ShowProviderDetailsMsg types.ProviderDetail

func emitShowProviderDetailsMsg(p types.ProviderDetail) tea.Cmd {
	return func() tea.Msg {
		return ShowProviderDetailsMsg(p)
	}
}

type model struct {
	table  table.Model
	help   help.Model
	keyMap KeyMap

	client    types.ProviderManager
	providers map[string]types.ProviderDetail

	// detailsModel is nullable
	detailsModel tea.Model
}

func (m *model) initializeRows(ctx context.Context) []table.Row {
	pp, err := m.client.GetProviders(ctx)
	if err != nil {
		log.I.Fatal("failed to get providers", zap.Error(err))
	}

	rows := make([]table.Row, 0, len(pp))
	for _, p := range pp {
		prices, err := json.Marshal(struct {
			GPUPricePerMin  string
			CPUPricePerMin  string
			MemPricesPerMin string
		}{
			GPUPricePerMin:  ether.FromWei(p.ProviderPrices.GpuPricePerMin).String(),
			CPUPricePerMin:  ether.FromWei(p.ProviderPrices.CpuPricePerMin).String(),
			MemPricesPerMin: ether.FromWei(p.ProviderPrices.MemPricePerMin).String(),
		})
		if err != nil {
			panic("failed to marshal provider prices")
		}
		hardware, err := json.Marshal(p.ProviderHardware)
		if err != nil {
			panic("failed to marshal provider prices")
		}
		labels := utils.FormatLabels(p.Labels)
		yLabels, err := json.Marshal(labels)
		if err != nil {
			panic("failed to marshal provider prices")
		}

		row := table.Row{
			p.Addr.Hex(),
			utils.BoolToYN(p.IsValidForScheduling),
			utils.BoolToYN(p.IsWaitingForApproval),
			utils.BoolToYN(p.IsBanned),
			string(prices),
			string(hardware),
			string(yLabels),
		}

		rows = append(rows, row)
		m.providers[p.Addr.Hex()] = p
	}
	return rows
}

type initRowsMsg struct{}

func initRows() tea.Msg {
	return initRowsMsg{}
}

func (m model) Init() tea.Cmd {
	return initRows
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tbCmd   tea.Cmd
		pageCmd tea.Cmd
		cmds    = make([]tea.Cmd, 0)
	)

	switch msg := msg.(type) {
	case initRowsMsg:
		rows := m.initializeRows(context.TODO())
		m.table.SetRows(rows)
	case ShowProviderDetailsMsg:
		m.detailsModel = details.Model(types.ProviderDetail(msg))
		cmds = append(
			cmds,
			m.detailsModel.Init(),
		)
	case details.ExitMsg:
		_, _ = m.detailsModel.Update(msg)
		m.detailsModel = nil
	case tea.KeyMsg:
		switch {
		case m.detailsModel != nil:
			m.detailsModel, pageCmd = m.detailsModel.Update(msg)
			cmds = append(cmds, pageCmd)
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, emitExitMsg)
		case key.Matches(msg, m.keyMap.ShowProviderDetails):
			if len(m.table.SelectedRow()) > 0 {
				cmds = append(cmds, emitShowProviderDetailsMsg(m.providers[m.table.SelectedRow()[0]]))
			}
		}
	default:
		if m.detailsModel != nil {
			m.detailsModel, pageCmd = m.detailsModel.Update(msg)
			cmds = append(cmds, pageCmd)
		}
	}

	if m.detailsModel == nil {
		m.table, tbCmd = m.table.Update(msg)
		cmds = append(cmds, tbCmd)
	}

	return m, tea.Batch(cmds...)
}
