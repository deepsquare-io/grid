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

package nav

import (
	"context"
	"fmt"
	"math/big"

	_ "embed"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	internallog "github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/tui/editor"
	"github.com/deepsquare-io/the-grid/cli/tui/provider"
	"github.com/deepsquare-io/the-grid/cli/tui/status"
	"github.com/deepsquare-io/the-grid/cli/tui/status/log"
	"github.com/deepsquare-io/the-grid/cli/tui/status/topup"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/deepsquare-io/the-grid/cli/tui/transfer"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type model struct {
	version              string
	metaschedulerAddress string
	logs                 chan ethtypes.Log
	balance              *big.Int
	balanceChan          chan *big.Int
	allowance            *big.Int
	allowanceChan        chan *big.Int

	// logModel is nullable
	logModel tea.Model
	// editorModel is nullable
	editorModel tea.Model
	// transferModel is nullable
	transferModel tea.Model
	// providerModel is nullable
	providerModel tea.Model
	// topupModel is nullable
	topupModel  tea.Model
	statusModel tea.Model

	logModelBuilder      log.ModelBuilder
	editorModelBuilder   editor.ModelBuilder
	transferModelBuilder transfer.ModelBuilder
	providerModelBuilder provider.ModelBuilder
	topupModelBuilder    topup.ModelBuilder

	client      deepsquare.Client
	watcher     deepsquare.Watcher
	userAddress common.Address
}

type balanceMsg *big.Int
type allowanceMsg *big.Int

func (m *model) watchEvents(
	ctx context.Context,
) tea.Cmd {
	return func() tea.Msg {
		approvals := make(chan types.Approval, 1)
		transfers := make(chan types.Transfer, 1)
		sub, err := m.watcher.SubscribeEvents(
			ctx,
			types.FilterApproval(approvals),
			types.FilterTransfer(transfers),
		)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		defer sub.Unsubscribe()

		allowances, err := m.client.ReduceToAllowance(ctx, approvals)
		if err != nil {
			internallog.I.Fatal("failed to watch allowances", zap.Error(err))
		}
		balances, err := m.client.ReduceToBalance(ctx, transfers)
		if err != nil {
			internallog.I.Fatal("failed to watch balances", zap.Error(err))
		}

		for {
			select {
			case balance := <-balances:
				select {
				case m.balanceChan <- balance:
				case <-ctx.Done():
					return nil
				}
			case allowance := <-allowances:
				select {
				case m.allowanceChan <- allowance:
				case <-ctx.Done():
					return nil
				}

			case <-ctx.Done():
				return nil
			}
		}
	}
}

func (m *model) tick() tea.Msg {
	select {
	case balance := <-m.balanceChan:
		return balanceMsg(balance)
	case allowance := <-m.allowanceChan:
		return allowanceMsg(allowance)
	}
}

func (m model) Init() tea.Cmd {
	// TODO: handle termination
	return tea.Batch(
		m.statusModel.Init(),
		m.watchEvents(context.TODO()),
		m.tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		pageCmd tea.Cmd
		cmds    = make([]tea.Cmd, 0)
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case m.logModel != nil:
			m.logModel, pageCmd = m.logModel.Update(msg)
		case m.editorModel != nil:
			m.editorModel, pageCmd = m.editorModel.Update(msg)
		case m.transferModel != nil:
			m.transferModel, pageCmd = m.transferModel.Update(msg)
		case m.providerModel != nil:
			m.providerModel, pageCmd = m.providerModel.Update(msg)
		case m.topupModel != nil:
			m.topupModel, pageCmd = m.topupModel.Update(msg)
		default:
			m.statusModel, pageCmd = m.statusModel.Update(msg)
		}
		cmds = append(cmds, pageCmd)
	case status.SelectJobMsg:
		// Render job logs
		m.logModel = m.logModelBuilder.Build(context.TODO(), msg)
		cmds = append(cmds, m.logModel.Init())
	case status.SubmitJobMsg:
		m.editorModel = m.editorModelBuilder.Build()
		cmds = append(
			cmds,
			m.editorModel.Init(),
		)
	case status.TransferCreditsMsg:
		m.transferModel = m.transferModelBuilder.Build()
		cmds = append(
			cmds,
			m.transferModel.Init(),
		)
	case status.ViewProvidersMsg:
		m.providerModel = m.providerModelBuilder.Build()
		cmds = append(
			cmds,
			m.providerModel.Init(),
		)
	case status.TopupJobMsg:
		m.topupModel = m.topupModelBuilder.Build(msg)
		cmds = append(
			cmds,
			m.topupModel.Init(),
		)
	case log.ExitMsg:
		_, _ = m.logModel.Update(msg)
		m.logModel = nil
	case editor.ExitMsg:
		_, _ = m.editorModel.Update(msg)
		m.editorModel = nil
		if msg.JobID != [32]byte{} {
			cmds = append(cmds, status.EmitSelectJobMsg(msg.JobID))
		}
	case transfer.ExitMsg:
		_, _ = m.transferModel.Update(msg)
		m.transferModel = nil
	case provider.ExitMsg:
		_, _ = m.providerModel.Update(msg)
		m.providerModel = nil
	case topup.ExitMsg:
		_, _ = m.topupModel.Update(msg)
		m.topupModel = nil
	case balanceMsg:
		m.balance = msg
		cmds = append(cmds, m.tick)
	case allowanceMsg:
		m.allowance = msg
		cmds = append(cmds, m.tick)
	default:
		switch {
		case m.logModel != nil:
			m.logModel, pageCmd = m.logModel.Update(msg)
		case m.editorModel != nil:
			m.editorModel, pageCmd = m.editorModel.Update(msg)
		case m.transferModel != nil:
			m.transferModel, pageCmd = m.transferModel.Update(msg)
		case m.providerModel != nil:
			m.providerModel, pageCmd = m.providerModel.Update(msg)
		case m.topupModel != nil:
			m.topupModel, pageCmd = m.topupModel.Update(msg)
		default:
			m.statusModel, pageCmd = m.statusModel.Update(msg)
		}
		cmds = append(cmds, pageCmd)
	}

	return m, tea.Batch(cmds...)
}

//go:embed title.txt
var titlePixelArt string

var labels = lipgloss.NewStyle().PaddingRight(2).Render(`Version:
User Address:
Smart-Contract Address:
Current balance:
Allowance:`)

func (m model) View() string {
	var navView string
	switch {
	case m.logModel != nil:
		navView = m.logModel.View()
	case m.editorModel != nil:
		navView = m.editorModel.View()
	case m.transferModel != nil:
		navView = m.transferModel.View()
	case m.providerModel != nil:
		navView = m.providerModel.View()
	case m.topupModel != nil:
		navView = m.topupModel.View()
	default:
		navView = m.statusModel.View()
	}

	values := fmt.Sprintf(`%s
%s
%s
%s credits (%s wei)
%s credits (%s wei)`,
		m.version,
		m.userAddress,
		m.metaschedulerAddress,
		ether.FromWei(m.balance).String(),
		m.balance,
		ether.FromWei(m.allowance).String(),
		m.allowance,
	)
	info := style.Box.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			labels,
			values,
		),
	)

	return style.Foreground.Render(titlePixelArt) + "\n" + info + "\n" + navView
}
