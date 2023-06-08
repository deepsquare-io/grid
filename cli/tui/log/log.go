package log

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/ethereum/go-ethereum/common"
)

func (m model) headerView() string {
	title := style.LogTitle.Render(m.title)
	line := style.Foreground.Render(
		strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title))),
	)
	return lipgloss.JoinHorizontal(lipgloss.Top, title, line)
}

func (m model) footerView() string {
	info := style.LogInfo.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := style.Foreground.Render(
		strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info))),
	)
	return lipgloss.JoinHorizontal(lipgloss.Bottom, line, info)
}

func (m model) View() string {
	if len(m.logs) == 0 {
		return fmt.Sprintf(
			"%s\nWaiting for logs... %s%s\n%s",
			m.headerView(),
			m.spinner.View(),
			strings.Repeat("\n", max(0, m.viewport.Height-1)),
			m.footerView(),
		)
	}
	return fmt.Sprintf(
		"%s\n%s\n%s",
		m.headerView(), m.viewport.View(), m.footerView(),
	)
}

func Model(logger deepsquare.Logger, userAddress common.Address, jobID [32]byte) tea.Model {
	vp := viewport.New(80, style.StandardHeight-2)
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		viewport: vp,
		spinner:  s,
		logs:     make([]logMsg, 0, 100),
		logsChan: make(chan logMsg, 100),
		messages: make([]string, 0, 100),
		title:    fmt.Sprintf("Job %s", new(big.Int).SetBytes(jobID[:])),

		logger:      logger,
		userAddress: userAddress,
		jobID:       jobID,
	}
}
