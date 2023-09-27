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
	"github.com/deepsquare-io/grid/cli/tui/style"
	"github.com/deepsquare-io/grid/cli/types"
)

// ModelBuilder contains the dependencies to build the model for the log page.
type ModelBuilder struct {
	Logger types.Logger
}

// Build the bubbletea model for the log page.
func (b *ModelBuilder) Build(ctx context.Context, jobID [32]byte) tea.Model {
	if b.Logger == nil {
		panic("Logger is nil")
	}
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
			b.Logger,
		),
		logs:     make([]logMsg, 0, 100),
		messages: messages,
		title:    fmt.Sprintf("Job %s", new(big.Int).SetBytes(jobID[:])),
	}
}
