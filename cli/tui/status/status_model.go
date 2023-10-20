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

package status

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/ether"
	"github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/tui/channel"
	"github.com/deepsquare-io/grid/cli/tui/components/table"
	"github.com/deepsquare-io/grid/cli/tui/components/ticker"
	"github.com/deepsquare-io/grid/cli/tui/style"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type keyMap struct {
	TableKeyMap     table.KeyMap
	OpenLogs        key.Binding
	TopupJob        key.Binding
	CancelJob       key.Binding
	SubmitJob       key.Binding
	TransferCredits key.Binding
	ViewProviders   key.Binding
	Exit            key.Binding
}

type model struct {
	table table.Model
	// Index jobID to row
	idToRow map[[32]byte]table.Row
	idToJob map[[32]byte]types.Job
	// Index jobID to running job
	runningIDs map[[32]byte]bool
	it         types.JobLazyIterator
	help       help.Model
	watchJobs  channel.Model[transitionMsg]
	scheduler  types.JobScheduler
	keyMap     keyMap

	ticker ticker.Model

	isCancelling bool

	err error
}

var zero = big.NewInt(0)

type cancellingProgressMsg bool

type errorMsg error

type clearErrorsMsg struct{}

func emitClearErrorsMsg() tea.Msg {
	return clearErrorsMsg{}
}

// SelectJobMsg is a public msg used to indicate that the user selected a job.
type SelectJobMsg [32]byte

// EmitSelectJobMsg emits a [SelectJobMsg].
func EmitSelectJobMsg(msg [32]byte) tea.Cmd {
	return func() tea.Msg {
		return SelectJobMsg(msg)
	}
}

// SubmitJobMsg is a public msg used to indicate that the user want to submit a job.
type SubmitJobMsg struct{}

func emitSubmitJobMsg() tea.Msg {
	return SubmitJobMsg{}
}

// TransferCreditsMsg is a public msg used to indicate that the user want to transfer credits.
type TransferCreditsMsg struct{}

func emitTransferCreditsMsg() tea.Msg {
	return TransferCreditsMsg{}
}

// ViewProvidersMsg is a public msg used to indicate that the user want to transfer credits.
type ViewProvidersMsg struct{}

func emitViewProvidersMsg() tea.Msg {
	return ViewProvidersMsg{}
}

// TopupJobMsg is a public msg used to indicate that the user top up a job.
type TopupJobMsg [32]byte

func emitTopupJobMsg(msg [32]byte) tea.Cmd {
	return func() tea.Msg {
		return TopupJobMsg(msg)
	}
}

func jobToRow(job types.Job) table.Row {
	startDate := "---"
	if job.Time.Start.Cmp(zero) > 0 {
		startDate = (time.UnixMilli(job.Time.Start.Int64() * 1000)).Format("15:04 _2 Jan 2006")
	}
	duration := "---"
	if job.Time.End.Cmp(zero) > 0 && job.Time.End.Cmp(job.Time.Start) >= 0 {
		durationB := new(big.Int).Sub(job.Time.End, job.Time.Start)
		duration = (time.Duration(durationB.Int64()) * time.Second).String()
	}
	return table.Row{
		new(big.Int).SetBytes(job.JobId[:]).String(),
		string(job.JobName[:]),
		metascheduler.JobStatus(job.Status).String(),
		startDate,
		duration,
		ether.FromWei(job.Cost.FinalCost).String(),
		strconv.Itoa(int(job.ExitCode / 256)),
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
) (
	rows []table.Row,
	idToRow map[[32]byte]table.Row,
	idToJob map[[32]byte]types.Job,
	runningIDs map[[32]byte]bool,
	it types.JobLazyIterator,
) {
	it, err := fetcher.GetJobs(ctx)
	if err != nil {
		log.I.Error("failed to get jobs", zap.Error(err))
		return nil, nil, nil, nil, it
	}
	if it == nil {
		return nil, nil, nil, nil, it
	}
	rows = make([]table.Row, 0, style.StandardHeight)
	idToRow = make(map[[32]byte]table.Row)
	idToJob = make(map[[32]byte]types.Job)
	runningIDs = make(map[[32]byte]bool)
	for i := 0; i < style.StandardHeight; i++ {
		if !it.Next(ctx) {
			if it.Error() != nil {
				log.I.Error("failed to get next job, ignoring...", zap.Error(err))
			}
			break
		}
		job := it.Current()
		row := jobToRow(job)
		idToRow[job.JobId] = row
		idToJob[job.JobId] = job
		if metascheduler.JobStatus(job.Status) == metascheduler.JobStatusRunning {
			runningIDs[job.JobId] = true
		}
		rows = append(rows, row)
	}
	return rows, idToRow, idToJob, runningIDs, it
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
	var err error
	for i := 0; i < style.StandardHeight; i++ {
		if !m.it.Next(ctx) {
			if m.it.Error() != nil {
				log.I.Error("failed to get next job, ignoring...", zap.Error(err))
			}
			break
		}
		job := m.it.Current()
		row := jobToRow(job)
		m.idToRow[job.JobId] = row
		rows = append(rows, row)
	}
	m.table.SetRows(rows)
}

func (m model) CancelJob(ctx context.Context, jobID [32]byte) tea.Cmd {
	return tea.Batch(tea.Sequence(func() tea.Msg {
		if err := m.scheduler.CancelJob(ctx, jobID); err != nil {
			return errorMsg(err)
		}
		return nil
	}, func() tea.Msg {
		return cancellingProgressMsg(false)
	}), func() tea.Msg {
		return cancellingProgressMsg(true)
	})
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.watchJobs.Init(), m.ticker.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tbCmd tea.Cmd
		wCmd  tea.Cmd
		tCmd  tea.Cmd
		cmds  = make([]tea.Cmd, 0)
	)
	m.watchJobs, wCmd = m.watchJobs.Update(msg)
	if wCmd != nil {
		cmds = append(cmds, wCmd)
	}

	m.ticker, tCmd = m.ticker.Update(msg)
	if tCmd != nil {
		// Update duration column
		var needRefresh bool
		for id := range m.runningIDs {
			job := m.idToJob[id]
			row := m.idToRow[id]
			// Duration column
			elapsed := time.Since(time.Unix(job.Time.Start.Int64(), 0)).Truncate(time.Second)
			row[4] = elapsed.String()
			needRefresh = true
		}
		if needRefresh {
			m.table.SetRows(m.table.Rows())
		}
		cmds = append(cmds, tCmd)
	}

	switch msg := msg.(type) {
	case errorMsg:
		m.err = msg
		return m, nil
	case cancellingProgressMsg:
		m.isCancelling = bool(msg)
	case clearErrorsMsg:
		m.err = nil
	case transitionMsg:
		rows := m.table.Rows()
		m.idToJob[msg.JobId] = types.Job(msg)
		row, ok := m.idToRow[msg.JobId]
		if ok {
			copy(row[:], jobToRow(types.Job(msg)))
		} else {
			row = jobToRow(types.Job(msg))
			m.idToRow[msg.JobId] = row
			rows = append(rows, table.Row{})
			copy(rows[1:], rows)
			rows[0] = row
		}
		if msg.Status == uint8(metascheduler.JobStatusRunning) {
			m.runningIDs[msg.JobId] = true
		} else {
			if m.runningIDs[msg.JobId] {
				delete(m.runningIDs, msg.JobId)
			}
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
				cmds = append(cmds, tea.Sequence(emitClearErrorsMsg, m.CancelJob(context.TODO(), rowToJobID(m.table.SelectedRow()))))
			}
		case key.Matches(msg, m.keyMap.OpenLogs):
			if len(m.table.SelectedRow()) > 0 {
				cmds = append(cmds, EmitSelectJobMsg(rowToJobID(m.table.SelectedRow())))
			}
		case key.Matches(msg, m.keyMap.TopupJob):
			if len(m.table.SelectedRow()) > 0 {
				cmds = append(cmds, emitTopupJobMsg(rowToJobID(m.table.SelectedRow())))
			}
		case key.Matches(msg, m.keyMap.SubmitJob):
			cmds = append(cmds, emitSubmitJobMsg)
		case key.Matches(msg, m.keyMap.TransferCredits):
			cmds = append(cmds, emitTransferCreditsMsg)
		case key.Matches(msg, m.keyMap.ViewProviders):
			cmds = append(cmds, emitViewProvidersMsg)
		case key.Matches(msg, m.keyMap.Exit):
			return m, tea.Sequence(
				m.watchJobs.Dispose,
				tea.Quit,
			)
		}
	}
	m.table, tbCmd = m.table.Update(msg)
	cmds = append(cmds, tbCmd)
	return m, tea.Batch(cmds...)
}

// Model builds the bubbletea model for the status page.
func Model(
	ctx context.Context,
	client deepsquare.Client,
	watcher deepsquare.Watcher,
	userAddress common.Address,
) tea.Model {
	if client == nil {
		panic("client is nil")
	}
	if watcher == nil {
		panic("watcher is nil")
	}
	// Initialize rows
	rows, idToRow, idToJob, runningIDs, it := initializeRows(ctx, client)

	tableKeymap := table.DefaultKeyMap()
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(style.StandardHeight),
		table.WithKeyMap(tableKeymap),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	s.RenderCell = func(model table.Model, value string, rawValue string, position table.CellPosition) string {
		// Job Status
		if position.Column == 2 {
			return style.JobStatusStyle(rawValue).Render(value)
		}
		return value
	}
	t.SetStyles(s)

	help := help.New()
	help.ShowAll = true

	return &model{
		table:      t,
		idToRow:    idToRow,
		it:         it,
		help:       help,
		runningIDs: runningIDs,
		idToJob:    idToJob,
		ticker: ticker.Model{
			Ticker: time.NewTicker(time.Second),
		},
		keyMap: keyMap{
			TableKeyMap: tableKeymap,
			OpenLogs: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "show job logs"),
			),
			TopupJob: key.NewBinding(
				key.WithKeys("t", "v"),
				key.WithHelp("t/v", "topup/view job"),
			),
			CancelJob: key.NewBinding(
				key.WithKeys("c"),
				key.WithHelp("c", "cancel job"),
			),
			SubmitJob: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "submit job"),
			),
			TransferCredits: key.NewBinding(
				key.WithKeys("ctrl+t"),
				key.WithHelp("ctrl+t", "transfer credits"),
			),
			ViewProviders: key.NewBinding(
				key.WithKeys("p"),
				key.WithHelp("p", "view providers"),
			),
			Exit: key.NewBinding(
				key.WithKeys("esc", "q"),
				key.WithHelp("esc/q", "exit"),
			),
		},
		scheduler: client,
		watchJobs: makeWatchJobsModel(
			ctx,
			userAddress,
			watcher,
			client,
		),
	}
}
