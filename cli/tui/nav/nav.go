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
	"github.com/deepsquare-io/the-grid/cli/tui/log"
	"github.com/deepsquare-io/the-grid/cli/tui/status"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/deepsquare-io/the-grid/cli/tui/util"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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
	statusModel tea.Model

	logModelBuilder    log.ModelBuilder
	editorModelBuilder editor.ModelBuilder

	watcher     deepsquare.Watcher
	userAddress common.Address
}

type balanceMsg *big.Int
type allowanceMsg *big.Int

func (m *model) watchEvents(
	ctx context.Context,
) tea.Cmd {
	return func() tea.Msg {
		sub, err := m.watcher.SubscribeEvents(ctx, m.logs)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		defer sub.Unsubscribe()
		transfers, rest := m.watcher.FilterTransfer(ctx, m.logs)
		approvals, rest := m.watcher.FilterApproval(ctx, rest)

		balances, err := m.watcher.ReduceToBalance(ctx, transfers)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		allowances, err := m.watcher.ReduceToAllowance(ctx, approvals)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		go util.IgnoreElements(rest)

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
	case log.ExitMsg:
		_, _ = m.logModel.Update(msg)
		m.logModel = nil
	case editor.ExitMsg:
		_, _ = m.editorModel.Update(msg)
		m.editorModel = nil
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
	default:
		navView = m.statusModel.View()
	}

	values := fmt.Sprintf(`%s
%s
%s
%s creds (%s wei)
%s creds (%s wei)`,
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

func Model(
	ctx context.Context,
	userAddress common.Address,
	watcher deepsquare.Watcher,
	statusModel tea.Model,
	logModelBuilder log.ModelBuilder,
	editorModelBuilder editor.ModelBuilder,
	version string,
	metaschedulerAddress string,
) *model {
	return &model{
		logs:          make(chan ethtypes.Log, 100),
		balanceChan:   make(chan *big.Int, 10),
		balance:       new(big.Int),
		allowanceChan: make(chan *big.Int, 10),
		allowance:     new(big.Int),

		statusModel:        statusModel,
		logModelBuilder:    logModelBuilder,
		editorModelBuilder: editorModelBuilder,

		userAddress:          userAddress,
		watcher:              watcher,
		version:              version,
		metaschedulerAddress: metaschedulerAddress,
	}
}
