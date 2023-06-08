package log

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type logMsg struct {
	timestamp time.Time
	message   string
}

type model struct {
	viewport viewport.Model
	spinner  spinner.Model
	messages []string
	logs     []logMsg
	logsChan chan logMsg
	title    string

	showTimestamp bool

	logger      deepsquare.Logger
	userAddress common.Address
	jobID       [32]byte
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m *model) watchLogs(
	ctx context.Context,
) tea.Cmd {
	return func() tea.Msg {
		stream, err := m.logger.WatchLogs(ctx, m.jobID)
		if err != nil {
			logger.I.Error("failed to get logs", zap.Error(err))
			return nil
		}
		defer stream.CloseSend()
		for {
			req, err := stream.Recv()
			if err == io.EOF || errors.Is(err, context.Canceled) {
				// TODO: handle closure
				logger.I.Info("logs closed", zap.Error(err))
				return nil
			}
			if err != nil {
				logger.I.Error("failed to get logs", zap.Error(err))
				return nil
			}
			select {
			case m.logsChan <- logMsg{
				timestamp: time.Unix(0, req.GetTimestamp()),
				message:   string(req.GetData()),
			}:
			case <-ctx.Done():
				// Context canceled. This is not an error.
				return nil
			}
		}
	}
}

func (m *model) tickLog() tea.Msg {
	return logMsg(<-m.logsChan)
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.watchLogs(context.Background()),
		m.tickLog,
		m.spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd tea.Cmd
		sCmd  tea.Cmd
		cmds  = make([]tea.Cmd, 0)
	)

	m.viewport, vpCmd = m.viewport.Update(msg)
	cmds = append(cmds, vpCmd)
	m.spinner, sCmd = m.spinner.Update(msg)
	cmds = append(cmds, sCmd)

	switch msg := msg.(type) {
	case logMsg:
		m.logs = append(m.logs, msg)
		if m.showTimestamp {
			m.messages = append(m.messages, fmt.Sprintf("%s: %s", msg.timestamp, msg.message))
		} else {
			m.messages = append(m.messages, msg.message)
		}

		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()
		cmds = append(cmds, m.tickLog)
	}
	return m, tea.Batch(cmds...)
}
