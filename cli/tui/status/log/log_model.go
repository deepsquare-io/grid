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
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/tui/channel"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

// ExitMsg msg closes to log model
type ExitMsg struct{}

func (m *model) emitExitMsg() tea.Cmd {
	return tea.Sequence(
		m.watchLogs.Dispose,
		m.transitions.Dispose,
		func() tea.Msg {
			return ExitMsg{}
		},
	)
}

type keyMap struct {
	ViewPort viewport.KeyMap
	Exit     key.Binding
}

type model struct {
	viewport    viewport.Model
	spinner     spinner.Model
	messages    *strings.Builder
	logs        []logMsg
	watchLogs   channel.Model[logMsg]
	transitions channel.Model[transitionMsg]
	title       string
	client      deepsquare.Client

	help   help.Model
	keyMap keyMap

	jobID                    [32]byte
	allocatedProviderAddress common.Address
	provider                 types.ProviderDetail
	msOrSchedLen             int64
	runningLen               int64
	status                   metascheduler.JobStatus

	showTimestamp bool
}

func (m model) Init() tea.Cmd {
	// TODO: handle termination
	return tea.Batch(
		m.watchLogs.Init(),
		m.transitions.Init(),
		m.spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd    tea.Cmd
		sCmd     tea.Cmd
		lChanCmd tea.Cmd
		tChanCmd tea.Cmd
		cmds     = make([]tea.Cmd, 0)
	)

	m.viewport, vpCmd = m.viewport.Update(msg)
	if vpCmd != nil {
		cmds = append(cmds, vpCmd)
	}
	m.watchLogs, lChanCmd = m.watchLogs.Update(msg)
	if lChanCmd != nil {
		cmds = append(cmds, lChanCmd)
	}
	m.transitions, tChanCmd = m.transitions.Update(msg)
	if tChanCmd != nil {
		cmds = append(cmds, tChanCmd)
	}

	switch msg := msg.(type) {
	case spinner.TickMsg:
		m.spinner, sCmd = m.spinner.Update(msg)
		if sCmd != nil {
			cmds = append(cmds, sCmd)
		}
	case transitionMsg:
		if (m.allocatedProviderAddress != common.Address{} && !isRunningOrFinished(m.status)) {
			jobs, err := m.client.GetJobsByProvider(context.Background(), m.allocatedProviderAddress)
			if err != nil {
				internallog.I.Warn("failed to fetch running jobs info", zap.Error(err))
			}
			msLen, rLen := reduceJobsIntoRunningOrScheduledLens(jobs)
			if len(jobs) > 1 && m.msOrSchedLen > 1 && (m.msOrSchedLen != msLen || m.runningLen != rLen) {
				waitingTime, err := computeWaitingTime(m.jobID, m.provider, jobs)
				if err != nil {
					internallog.I.Fatal("failed to compute waiting time", zap.Error(err))
				}
				fmt.Printf("(%d jobs in provider queue: %d waiting, %d running, wait ~%s)\n", len(jobs), msLen, rLen, waitingTime)
			}
			m.msOrSchedLen, m.runningLen = msLen, rLen
		}

		if bytes.Equal(m.jobID[:], msg.JobId[:]) {
			m.status = metascheduler.JobStatus(msg.To)
			switch metascheduler.JobStatus(msg.To) {
			case metascheduler.JobStatusMetaScheduled,
				metascheduler.JobStatusScheduled:
				m.messages.WriteString(fmt.Sprintf("\n(Job is %s)", metascheduler.JobStatus(msg.To)))
				job, err := m.client.GetJob(context.Background(), m.jobID)
				if err != nil {
					internallog.I.Fatal("failed to fetch job info", zap.Error(err))
				}
				m.allocatedProviderAddress = job.ProviderAddr
				p, err := m.client.GetProvider(context.Background(), m.allocatedProviderAddress)
				if err != nil {
					internallog.I.Fatal("failed to get provider info", zap.Error(err))
				}
				m.provider = p
				jobs, err := m.client.GetJobsByProvider(context.Background(), m.allocatedProviderAddress)
				if err != nil {
					internallog.I.Warn("failed to fetch running jobs info", zap.Error(err))
				}
				msLen, rLen := reduceJobsIntoRunningOrScheduledLens(jobs)
				if len(jobs) > 1 && m.msOrSchedLen > 1 && (m.msOrSchedLen != msLen || m.runningLen != rLen) {
					waitingTime, err := computeWaitingTime(m.jobID, m.provider, jobs)
					if err != nil {
						internallog.I.Fatal("failed to compute waiting time", zap.Error(err))
					}
					fmt.Printf("(%d jobs in provider queue: %d waiting, %d running, wait ~%s)\n", len(jobs), msLen, rLen, waitingTime)
				}
				m.msOrSchedLen, m.runningLen = msLen, rLen
			case metascheduler.JobStatusCancelled,
				metascheduler.JobStatusFailed,
				metascheduler.JobStatusFinished,
				metascheduler.JobStatusPanicked,
				metascheduler.JobStatusOutOfCredits:
				// Do nothing
			default:
				m.messages.WriteString(fmt.Sprintf("\n(Job is %s)", metascheduler.JobStatus(msg.To)))
			}
		}
		m.viewport.SetContent(m.messages.String())
		m.viewport.GotoBottom()
	case logMsg:
		m.logs = append(m.logs, msg)
		if m.showTimestamp {
			m.messages.WriteString(fmt.Sprintf("\n%s: %s", msg.timestamp, msg.message))
		} else {
			m.messages.WriteString("\n" + msg.message)
		}

		m.viewport.SetContent(m.messages.String())
		m.viewport.GotoBottom()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, m.emitExitMsg())
		}
	}
	return m, tea.Batch(cmds...)
}

func reduceJobsIntoRunningOrScheduledLens(
	jobs []types.Job,
) (metascheduledOrScheduled int64, running int64) {
	for _, job := range jobs {
		switch metascheduler.JobStatus(job.Status) {
		case metascheduler.JobStatusRunning:
			running++
		case metascheduler.JobStatusMetaScheduled, metascheduler.JobStatusScheduled:
			metascheduledOrScheduled++
		}
	}
	return
}

// computeWaitingTime returns min(running) + sum(waiting)
func computeWaitingTime(
	jobID [32]byte,
	provider types.ProviderDetail,
	jobs []types.Job,
) (time.Duration, error) {
	var waiting, running time.Duration
	for _, job := range jobs {
		if bytes.Equal(job.JobId[:], jobID[:]) {
			continue
		}
		switch metascheduler.JobStatus(job.Status) {
		case metascheduler.JobStatusRunning:
			durationB, err := metascheduler.CreditToDuration(
				provider.ProviderPrices,
				job.Definition,
				job.Cost.MaxCost,
			)
			if err != nil {
				return 0, err
			}
			startTime := time.Unix(job.Time.Start.Int64(), 0)
			duration := (time.Duration(durationB.Int64())*time.Second - time.Since(startTime)).Truncate(
				time.Second,
			)
			if running > duration || running == 0 {
				running = duration
			}

		case metascheduler.JobStatusMetaScheduled, metascheduler.JobStatusScheduled:
			durationB, err := metascheduler.CreditToDuration(
				provider.ProviderPrices,
				job.Definition,
				job.Cost.MaxCost,
			)
			if err != nil {
				return 0, err
			}
			waiting += time.Duration(durationB.Int64()) * time.Second
		}
	}
	return running + waiting, nil
}

func isRunningOrFinished(status metascheduler.JobStatus) bool {
	switch status {
	case metascheduler.JobStatusRunning,
		metascheduler.JobStatusCancelled,
		metascheduler.JobStatusFailed,
		metascheduler.JobStatusFinished,
		metascheduler.JobStatusPanicked,
		metascheduler.JobStatusOutOfCredits:
		return true
	default:
		return false
	}
}
