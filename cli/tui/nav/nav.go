package nav

import (
	"context"
	"fmt"

	_ "embed"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/v1"
	"github.com/deepsquare-io/the-grid/cli/v1/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/log"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/status"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/style"
	"github.com/ethereum/go-ethereum/common"
)

type model struct {
	showLog              bool
	isFocusedOnLogs      bool
	version              string
	metaschedulerAddress string

	help help.Model
	// logModel is nullable
	logModel    tea.Model
	statusModel tea.Model

	logger      cli.Logger
	userAddress common.Address
	keymap      keymap
}

type keymap = struct {
	next key.Binding
}

func (m model) Init() tea.Cmd {
	return nil
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
		m.logModel = log.Model(m.logger, m.userAddress, msg)
		cmds = append(cmds, m.logModel.Init())
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
	info := style.Foreground.Render(
		fmt.Sprintf(`Version: %s
User Address: %s
Smart-Contract Address: %s`,
			m.version,
			m.userAddress,
			m.metaschedulerAddress,
		),
	)
	return style.Foreground.Render(titlePixelArt) + "\n" + info + "\n" + navView + "\n\n" + help
}

func Model(
	ctx context.Context,
	userAddress common.Address,
	rpc metascheduler.RPC,
	ws metascheduler.WS,
	logger cli.Logger,
	version string,
	metaschedulerAddress string,
) tea.Model {
	return model{
		statusModel: status.Model(ctx, rpc, ws, userAddress),
		logger:      logger,
		userAddress: userAddress,
		help:        help.New(),
		keymap: keymap{
			next: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "change focus"),
			),
		},
		version:              version,
		metaschedulerAddress: metaschedulerAddress,
	}
}
