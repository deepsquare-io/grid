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

type Model struct {
	table   table.Model
	idToRow map[[32]byte]table.Row
	it      deepsquare.JobLazyIterator
	help    help.Model
	jobs    chan deepsquare.Job

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

// initializeRows fetches the last "x" jobs and store an iterator in the model
// for lazy loading.
//
// This method is executed before the loading of the page for SSR.
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

// addMoreRows fetches the last "x" jobs and add it as row.
func (m *Model) addMoreRows() {
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

// watchTransition send new jobs object by watching events
func (m *Model) watchTransition(
	ctx context.Context,
) tea.Cmd {
	return func() tea.Msg {
		logs := make(chan types.Log, 100)
		sub, err := m.watcher.SubscribeEvents(ctx, logs)
		if err != nil {
			logger.I.Fatal(err.Error())
			return nil
		}
		transitions := m.watcher.FilterJobTransition(logs)
		newJobs := m.watcher.FilterNewJobRequests(logs)

		defer sub.Unsubscribe()
		for {
			select {
			case transition := <-transitions:
				job, err := m.fetcher.GetJob(ctx, transition.JobId)
				if err != nil {
					logger.I.Fatal(err.Error())
					return nil
				}
				m.jobs <- *job
			case newJob := <-newJobs:
				job, err := m.fetcher.GetJob(ctx, newJob.JobId)
				if err != nil {
					logger.I.Fatal(err.Error())
					return nil
				}
				m.jobs <- *job
			}
		}
	}
}

type transitionMsg deepsquare.Job

func (m *Model) handleTransition() tea.Cmd {
	return func() tea.Msg {
		if job, ok := <-m.jobs; ok {
			return transitionMsg(job)
		}
		return nil
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.watchTransition(context.Background()),
		m.handleTransition(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case transitionMsg:
		rows := m.table.Rows()
		if row, ok := m.idToRow[msg.JobId]; ok {
			row[2] = metascheduler.JobStatus(msg.Status).String()
		} else {
			job := deepsquare.Job(msg)
			row = jobToRow(job)
			m.idToRow[msg.JobId] = row
			rows = append(rows, table.Row{})
			copy(rows[1:], rows)
			rows[0] = row
		}
		m.table.SetRows(rows)
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
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}
