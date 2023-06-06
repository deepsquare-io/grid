package log

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (m model) View() string {
	if len(m.logs) == 0 {
		return fmt.Sprintf(
			"%s\n\n",
			m.spinner.View(),
		)
	}
	return fmt.Sprintf(
		"%s\n\n",
		m.viewport.View(),
	)
}

func Model(logger deepsquare.Logger, userAddress common.Address, jobID [32]byte) tea.Model {
	vp := viewport.New(80, 10)
	vp.SetContent(fmt.Sprintf("Waiting for the logs of job %s:", hexutil.Encode(jobID[:])))
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		viewport: vp,
		spinner:  s,
		logs:     make([]logMsg, 0, 100),
		logsChan: make(chan logMsg, 100),
		messages: make([]string, 0, 100),

		logger:      logger,
		userAddress: userAddress,
		jobID:       jobID,
	}
}
