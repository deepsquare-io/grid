package nav

import (
	"context"
	"fmt"
	"math/big"

	_ "embed"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/v1"
	"github.com/deepsquare-io/the-grid/cli/v1/internal/ether"
	internallog "github.com/deepsquare-io/the-grid/cli/v1/internal/log"
	"github.com/deepsquare-io/the-grid/cli/v1/logger"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/log"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/status"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/style"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type model struct {
	showLog              bool
	isFocusedOnLogs      bool
	version              string
	metaschedulerAddress string
	balance              *big.Int
	balanceChan          chan *big.Int
	allowance            *big.Int
	allowanceChan        chan *big.Int

	help help.Model
	// logModel is nullable
	logModel    tea.Model
	statusModel tea.Model

	creditWatcher    cli.CreditWatcher
	allowanceWatcher cli.AllowanceWatcher
	logDialer        logger.Dialer
	userAddress      common.Address
	keymap           keymap
}

type keymap = struct {
	next key.Binding
}

type balanceMsg *big.Int
type allowanceMsg *big.Int

func (m *model) watchEvents(
	ctx context.Context,
) tea.Cmd {
	return func() tea.Msg {
		logs := make(chan types.Log, 100)
		sub, err := m.creditWatcher.SubscribeEvents(ctx, logs)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		defer sub.Unsubscribe()
		transfers, rest := m.creditWatcher.FilterTransfer(ctx, logs)
		approvals, rest := m.allowanceWatcher.FilterApproval(ctx, rest)

		balances, err := m.creditWatcher.Balance(ctx, transfers)
		if err != nil {
			internallog.I.Fatal(err.Error())
		}
		allowances, err := m.allowanceWatcher.WatchAllowance(ctx, approvals)
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

func (m *model) tickBalance() tea.Msg {
	return balanceMsg(<-m.balanceChan)
}

func (m *model) tickAllowance() tea.Msg {
	return allowanceMsg(<-m.allowanceChan)
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.statusModel.Init(),
		m.watchEvents(context.Background()),
		m.tickBalance,
		m.tickAllowance,
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
		if m.isFocusedOnLogs && m.logModel != nil {
			m.logModel, lCmd = m.logModel.Update(msg)
			cmds = append(cmds, lCmd)
		} else {
			m.statusModel, sCmd = m.statusModel.Update(msg)
			cmds = append(cmds, sCmd)
		}
		switch {
		case key.Matches(msg, m.keymap.next):
			if m.logModel != nil {
				m.isFocusedOnLogs = !m.isFocusedOnLogs
			}
		case msg.Type == tea.KeyCtrlC, msg.Type == tea.KeyEsc, msg.String() == "q":
			return m, tea.Quit
		}
	case status.SelectJobMsg:
		m.showLog = true

		// Render job logs
		m.logModel = log.Model(m.logDialer, m.userAddress, msg)
		cmds = append(cmds, m.logModel.Init())
	case balanceMsg:
		m.balance = msg
		cmds = append(cmds, m.tickBalance)
	case allowanceMsg:
		m.allowance = msg
		cmds = append(cmds, m.tickAllowance)
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

func (m model) View() string {
	help := m.help.ShortHelpView([]key.Binding{
		m.keymap.next,
	})

	var navView string
	if m.showLog && m.logModel != nil {
		navView = lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.statusModel.View(),
			style.LeftVerticalSeparator.Render(m.logModel.View()),
		)
	} else {
		navView = m.statusModel.View()
	}
	info := style.Box.Render(
		fmt.Sprintf(`Version: %s
User Address: %s
Smart-Contract Address: %s
Current balance: %s creds (%s wei)
Allowance: %s creds (%s wei)`,
			m.version,
			m.userAddress,
			m.metaschedulerAddress,
			ether.FromWei(m.balance).String(),
			m.balance,
			ether.FromWei(m.allowance).String(),
			m.allowance,
		),
	)

	return style.Foreground.Render(titlePixelArt) + "\n" + info + "\n" + navView + "\n\n" + help
}

func Model(
	ctx context.Context,
	userAddress common.Address,
	fetcher cli.JobFetcher,
	watcher cli.JobWatcher,
	credit cli.CreditWatcher,
	allowance cli.AllowanceWatcher,
	logDialer logger.Dialer,
	version string,
	metaschedulerAddress string,
) tea.Model {
	return model{
		balanceChan:   make(chan *big.Int, 10),
		balance:       new(big.Int),
		allowanceChan: make(chan *big.Int, 10),
		allowance:     new(big.Int),

		statusModel: status.Model(ctx, fetcher, watcher, userAddress),
		logDialer:   logDialer,
		userAddress: userAddress,
		help:        help.New(),
		keymap: keymap{
			next: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "change focus"),
			),
		},
		creditWatcher:        credit,
		allowanceWatcher:     allowance,
		version:              version,
		metaschedulerAddress: metaschedulerAddress,
	}
}
