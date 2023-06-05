package status

import (
	"context"
	"math/big"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/deepsquare/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/ethereum/go-ethereum/core/types"
)

type model struct {
	table   table.Model
	idToRow map[[32]byte]table.Row
	it      deepsquare.JobLazyIterator
	help    help.Model

	fetcher deepsquare.JobFetcher
	watcher deepsquare.JobWatcher
}

func jobToRow(job deepsquare.Job) table.Row {
	return table.Row{
		new(big.Int).SetBytes(job.JobId[:]).String(),
		string(job.JobName[:]),
		metascheduler.JobStatus(job.Status).String(),
		(time.UnixMilli(job.Time.Start.Int64() * 1000)).String(),
	}
}

func initializeRows(
	ctx context.Context,
	fetcher deepsquare.JobFetcher,
) (rows []table.Row, idToRow map[[32]byte]table.Row, it deepsquare.JobLazyIterator) {
	it, err := fetcher.GetJobs(ctx)
	if err != nil {
		logger.I.Error(err.Error())
		return nil, nil, it
	}
	if it == nil {
		return nil, nil, it
	}
	rows = make([]table.Row, 0, tableHeight)
	idToRow = make(map[[32]byte]table.Row)
	var ok bool
	for i := 0; i < tableHeight; i++ {
		job := it.Current()
		row := jobToRow(*job)
		idToRow[job.JobId] = row
		rows = append(rows, row)
		it, ok, err = it.Next(ctx)
		if err != nil {
			logger.I.Error(err.Error())
			break
		}
		if !ok {
			break
		}
	}
	return rows, idToRow, it
}

func watchTransition(
	ctx context.Context,
	watcher deepsquare.JobWatcher,
	fetcher deepsquare.JobFetcher,
) tea.Cmd {
	out := make(chan deepsquare.Job)
	logs := make(chan types.Log, 100)
	sub, err := watcher.SubscribeEvents(ctx, logs)
	if err != nil {
		logger.I.Fatal(err.Error())
		return nil
	}

	go func() {
		defer sub.Unsubscribe()
		transitions := watcher.FilterJobTransition(logs)
		newJobs := watcher.FilterNewJobRequests(logs)
		for {
			select {
			case transition := <-transitions:
				job, err := fetcher.GetJob(ctx, transition.JobId)
				if err != nil {
					logger.I.Fatal(err.Error())
					return
				}
				out <- *job
			case newJob := <-newJobs:
				job, err := fetcher.GetJob(ctx, newJob.JobId)
				if err != nil {
					logger.I.Fatal(err.Error())
					return
				}
				out <- *job
			}
		}
	}()

	return func() tea.Msg {
		return transitionMsg{<-out, out}
	}
}

func (m *model) addMoreRows() {
	ctx := context.Background()
	oldRows := m.table.Rows()
	rows := make([]table.Row, 0, len(oldRows)+tableHeight)
	if m.it == nil {
		return
	}
	rows = append(rows, m.table.Rows()...)
	var ok bool
	var err error
	for i := 0; i < tableHeight; i++ {
		job := m.it.Current()
		row := jobToRow(*job)
		m.idToRow[job.JobId] = row
		rows = append(rows, row)
		m.it, ok, err = m.it.Next(ctx)
		if err != nil {
			logger.I.Error(err.Error())
			break
		}
		if !ok {
			break
		}
	}
	m.table.SetRows(rows)
}

func (m model) Init() tea.Cmd {
	return watchTransition(context.Background(), m.watcher, m.fetcher)
}

func (m transitionMsg) awaitNext() transitionMsg {
	return transitionMsg{current: <-m.next, next: m.next}
}

type transitionMsg struct {
	current deepsquare.Job
	next    <-chan deepsquare.Job
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case transitionMsg:
		rows := m.table.Rows()
		if row, ok := m.idToRow[msg.current.JobId]; ok {
			row[2] = metascheduler.JobStatus(msg.current.Status).String()
		} else {
			job := deepsquare.Job(msg.current)
			row = jobToRow(job)
			m.idToRow[msg.current.JobId] = row
			rows = append(rows, table.Row{})
			copy(rows[1:], rows)
			rows[0] = row
		}
		m.table.SetRows(rows)
		cmds = append(cmds, func() tea.Msg { return msg.awaitNext() })
	case tea.KeyMsg:
		switch {
		case msg.String() == "q", msg.String() == "ctrl+c":
			return m, tea.Quit
		case msg.String() == "enter":
			return m, tea.Printf("len(rows) %v!", len(m.table.Rows()))
		case key.Matches(msg, m.table.KeyMap.LineDown):
			if m.table.Cursor() == len(m.table.Rows())-1 {
				m.addMoreRows()
			}
		}
	}
	table, cmd := m.table.Update(msg)
	m.table = table
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
