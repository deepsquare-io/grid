package log

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/ethereum/go-ethereum/common"
)

func (m model) headerView() string {
	return style.LogTitle.Render(m.title)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m model) footerView() string {
	info := style.LogInfo.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := style.Foreground.Render(
		strings.Repeat(" ", max(0, m.viewport.Width-lipgloss.Width(info))),
	)
	return lipgloss.JoinHorizontal(lipgloss.Bottom, line, info)
}

func (m model) View() string {
	var view string
	if len(m.logs) == 0 {
		view = fmt.Sprintf(
			"%s\nWaiting for logs... %s%s\n%s",
			m.headerView(),
			m.spinner.View(),
			strings.Repeat("\n", max(0, m.viewport.Height-1)),
			m.footerView(),
		)
	} else {
		view = fmt.Sprintf(
			"%s\n%s\n%s",
			m.headerView(), m.viewport.View(), m.footerView(),
		)
	}
	return style.Box.Render(view)
}

func Model(
	ctx context.Context,
	loggerDialer logger.Dialer,
	userAddress common.Address,
	jobID [32]byte,
) tea.Model {
	vp := viewport.New(118, style.StandardHeight-4)
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	messages := &strings.Builder{}
	messages.Grow(1048576)
	return &model{
		viewport: vp,
		spinner:  s,
		watchLogs: makeWatchLogsModel(
			ctx,
			jobID,
			loggerDialer,
		),
		logs:     make([]logMsg, 0, 100),
		messages: messages,
		title:    fmt.Sprintf("Job %s", new(big.Int).SetBytes(jobID[:])),
	}
}
