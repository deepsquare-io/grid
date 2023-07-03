// Copyright (C) 2023 DeepSquare
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
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/deepsquare-io/the-grid/cli/types"
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
	logger types.Logger,
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
			logger,
		),
		logs:     make([]logMsg, 0, 100),
		messages: messages,
		title:    fmt.Sprintf("Job %s", new(big.Int).SetBytes(jobID[:])),
	}
}
