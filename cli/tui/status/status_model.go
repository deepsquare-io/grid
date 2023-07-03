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

package status

import (
	"context"
	"math/big"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/tui/channel"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
	"github.com/deepsquare-io/the-grid/cli/types"
	"go.uber.org/zap"
)

type KeyMap struct {
	TableKeyMap table.KeyMap
	OpenLogs    key.Binding
	CancelJob   key.Binding
	SubmitJob   key.Binding
	Exit        key.Binding
}

type model struct {
	table     table.Model
	idToRow   map[[32]byte]table.Row
	it        types.JobLazyIterator
	help      help.Model
	watchJobs channel.Model[transitionMsg]
	scheduler types.JobScheduler
	keyMap    KeyMap
}

// SelectJobMsg is a public msg used to indicate that the user selected a job.
type SelectJobMsg [32]byte

func emitSelectJobMsg(msg [32]byte) tea.Cmd {
	return func() tea.Msg {
		return SelectJobMsg(msg)
	}
}

// SubmitJobMsg is a public msg used to indicate that the user want to submit a job.
type SubmitJobMsg struct{}

func emitSubmitJobMsg() tea.Msg {
	return SubmitJobMsg{}
}

func jobToRow(job types.Job) table.Row {
	return table.Row{
		new(big.Int).SetBytes(job.JobID[:]).String(),
		string(job.JobName[:]),
		metascheduler.JobStatus(job.Status).String(),
		(time.UnixMilli(job.Time.Start.Int64() * 1000)).Format(time.UnixDate),
	}
}

func rowToJobID(row table.Row) [32]byte {
	jobIDStr := row[0]
	jobIDBig, _ := new(big.Int).SetString(jobIDStr, 10)
	var jobID [32]byte
	jobIDBig.FillBytes(jobID[:])
	return jobID
}

// initializeRows fetches the last "x" jobs and store an iterator in the model
// for lazy loading.
//
// This method is executed before the loading of the page for SSR.
func initializeRows(
	ctx context.Context,
	fetcher types.JobFetcher,
) (rows []table.Row, idToRow map[[32]byte]table.Row, it types.JobLazyIterator) {
	it, err := fetcher.GetJobs(ctx)
	if err != nil {
		log.I.Error("failed to get jobs", zap.Error(err))
		return nil, nil, it
	}
	if it == nil {
		return nil, nil, it
	}
	rows = make([]table.Row, 0, style.StandardHeight)
	idToRow = make(map[[32]byte]table.Row)
	var ok bool
	for i := 0; i < style.StandardHeight; i++ {
		job := it.Current()
		row := jobToRow(*job)
		idToRow[job.JobID] = row
		rows = append(rows, row)
		it, ok, err = it.Next(ctx)
		if err != nil {
			log.I.Error("failed to get next job, ignoring...", zap.Error(err))
			break
		}
		if !ok {
			break
		}
	}
	return rows, idToRow, it
}

// addMoreRows fetches the last "x" jobs and add it as row.
func (m *model) addMoreRows(ctx context.Context) {
	// TODO: handle termination
	oldRows := m.table.Rows()
	rows := make([]table.Row, 0, len(oldRows)+style.StandardHeight)
	if m.it == nil {
		return
	}
	rows = append(rows, m.table.Rows()...)
	var ok bool
	var err error
	for i := 0; i < style.StandardHeight; i++ {
		job := m.it.Current()
		row := jobToRow(*job)
		m.idToRow[job.JobID] = row
		rows = append(rows, row)
		m.it, ok, err = m.it.Next(ctx)
		if err != nil {
			log.I.Error("failed to get next job, ignoring...", zap.Error(err))
			break
		}
		if !ok {
			break
		}
	}
	m.table.SetRows(rows)
}

func (m model) CancelJob(ctx context.Context, jobID [32]byte) tea.Cmd {
	return func() tea.Msg {
		if err := m.scheduler.CancelJob(ctx, jobID); err != nil {
			log.I.Error(err.Error())
		}
		return nil
	}
}

func (m model) Init() tea.Cmd {
	return m.watchJobs.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tbCmd tea.Cmd
		wCmd  tea.Cmd
		cmds  = make([]tea.Cmd, 0)
	)
	m.watchJobs, wCmd = m.watchJobs.Update(msg)
	if wCmd != nil {
		cmds = append(cmds, wCmd)
	}

	switch msg := msg.(type) {
	case transitionMsg:
		rows := m.table.Rows()
		if row, ok := m.idToRow[msg.JobID]; ok {
			row[2] = metascheduler.JobStatus(msg.Status).String()
		} else {
			job := types.Job(msg)
			row = jobToRow(job)
			m.idToRow[msg.JobID] = row
			rows = append(rows, table.Row{})
			copy(rows[1:], rows)
			rows[0] = row
		}
		m.table.SetRows(rows)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.TableKeyMap.LineDown):
			if m.table.Cursor() == len(m.table.Rows())-1 {
				// TODO: handle termination
				m.addMoreRows(context.TODO())
			}

		case key.Matches(msg, m.keyMap.CancelJob):
			if len(m.table.SelectedRow()) > 0 {
				cmds = append(cmds, m.CancelJob(context.TODO(), rowToJobID(m.table.SelectedRow())))
			}
		case key.Matches(msg, m.keyMap.OpenLogs):
			if len(m.table.SelectedRow()) > 0 {
				cmds = append(cmds, emitSelectJobMsg(rowToJobID(m.table.SelectedRow())))
			}
		case key.Matches(msg, m.keyMap.SubmitJob):
			cmds = append(cmds, emitSubmitJobMsg)
		case key.Matches(msg, m.keyMap.Exit):
			return m, tea.Batch(
				m.watchJobs.Dispose,
				tea.Quit,
			)
		}
	}
	m.table, tbCmd = m.table.Update(msg)
	cmds = append(cmds, tbCmd)
	return m, tea.Batch(cmds...)
}
