package nav

import (
	"context"
	"fmt"
	"math/big"

	_ "embed"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	internallog "github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/tui/log"
	"github.com/deepsquare-io/the-grid/cli/tui/status"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/deepsquare-io/the-grid/cli/tui/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type model struct {
	version              string
	metaschedulerAddress string
	logs                 chan types.Log
	balance              *big.Int
	balanceChan          chan *big.Int
	allowance            *big.Int
	allowanceChan        chan *big.Int

	// logModel is nullable
	logModel    tea.Model
	statusModel tea.Model

	eventSubscriber   cli.EventSubscriber
	creditFilterer    cli.CreditFilterer
	allowanceFilterer cli.AllowanceFilterer
	logDialer         logger.Dialer
	userAddress       common.Address
	keymap            keymap
}

type keymap = struct {
}

type balanceMsg *big.Int
type allowanceMsg *big.Int

func (m *model) watchEvents(
	ctx context.Context,
) tea.Cmd {
	return func() tea.Msg {
		sub, err := m.eventSubscriber.SubscribeEvents(ctx, m.logs)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		defer sub.Unsubscribe()
		transfers, rest := m.creditFilterer.FilterTransfer(ctx, m.logs)
		approvals, rest := m.allowanceFilterer.FilterApproval(ctx, rest)

		balances, err := m.creditFilterer.ReduceToBalance(ctx, transfers)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		allowances, err := m.allowanceFilterer.ReduceToAllowance(ctx, approvals)
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
		lCmd tea.Cmd
		sCmd tea.Cmd
		cmds = make([]tea.Cmd, 0)
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.logModel != nil {
			m.logModel, lCmd = m.logModel.Update(msg)
			cmds = append(cmds, lCmd)
		} else {
			m.statusModel, sCmd = m.statusModel.Update(msg)
			cmds = append(cmds, sCmd)
		}
		switch {
		case msg.Type == tea.KeyCtrlC:
			return m, tea.Quit
		}
	case status.SelectJobMsg:
		// Render job logs
		m.logModel = log.Model(context.TODO(), m.logDialer, m.userAddress, msg)
		cmds = append(cmds, m.logModel.Init())
	case log.ExitMsg:
		_, _ = m.logModel.Update(msg)
		m.logModel = nil
	case balanceMsg:
		m.balance = msg
		cmds = append(cmds, m.tick)
	case allowanceMsg:
		m.allowance = msg
		cmds = append(cmds, m.tick)
	default:
		if m.logModel != nil {
			m.logModel, lCmd = m.logModel.Update(msg)
			cmds = append(cmds, lCmd)
		}
		m.statusModel, sCmd = m.statusModel.Update(msg)
		cmds = append(cmds, sCmd)
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
	if m.logModel != nil {
		navView = m.logModel.View()
	} else {
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
	eventSubscriber cli.EventSubscriber,
	jobFetcher cli.JobFetcher,
	jobFilterer cli.JobFilterer,
	creditFilterer cli.CreditFilterer,
	allowanceFilterer cli.AllowanceFilterer,
	logDialer logger.Dialer,
	version string,
	metaschedulerAddress string,
) tea.Model {
	return model{
		logs:          make(chan types.Log, 100),
		balanceChan:   make(chan *big.Int, 10),
		balance:       new(big.Int),
		allowanceChan: make(chan *big.Int, 10),
		allowance:     new(big.Int),

		statusModel: status.Model(
			ctx,
			eventSubscriber,
			jobFetcher,
			jobFilterer,
			userAddress,
		),
		logDialer:            logDialer,
		userAddress:          userAddress,
		keymap:               keymap{},
		eventSubscriber:      eventSubscriber,
		creditFilterer:       creditFilterer,
		allowanceFilterer:    allowanceFilterer,
		version:              version,
		metaschedulerAddress: metaschedulerAddress,
	}
}
