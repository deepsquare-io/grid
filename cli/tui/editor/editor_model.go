package editor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"unicode"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/internal/utils"
	"github.com/deepsquare-io/the-grid/cli/internal/validator"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
	"github.com/knipferrc/teacup/code"
	"gopkg.in/yaml.v3"
)

const (
	creditsLockingInput = iota
	usesInput
	jobNameInput
)

type KeyMap struct {
	EditAgain      key.Binding
	Exit           key.Binding
	NextInput      key.Binding
	PrevInput      key.Binding
	ViewPortKeymap viewport.KeyMap
}

// ExitMsg msg closes to editor model
type ExitMsg struct{}

func emitExitMsg() tea.Msg {
	return ExitMsg{}
}

type submitProgressMsg struct{}

func emitSubmitProgressMsg() tea.Msg {
	return submitProgressMsg{}
}

type submitDoneMsg struct{}

type errorMsg error

type model struct {
	// Form
	inputs  []textinput.Model
	errors  []error
	focused int

	// Code
	code          code.Model
	jobSchemaPath string
	jobPath       string

	err error

	help   help.Model
	keyMap KeyMap

	allowanceManager cli.AllowanceManager
	jobScheduler     cli.JobScheduler

	isRunning bool
}

type editorDone struct {
	err           error
	jobSchemaPath string
	jobPath       string
}

func (m *model) openEditor(ctx context.Context, jobSchemaPath string, jobPath string) tea.Cmd {
	return tea.ExecProcess(Command(ctx, jobPath), func(err error) tea.Msg {
		return editorDone{
			jobPath:       jobPath,
			jobSchemaPath: jobSchemaPath,
			err:           err,
		}
	})
}

func (m *model) submitJob(ctx context.Context, jobPath string) tea.Cmd {
	return func() tea.Msg {
		// Read file
		dat, err := os.ReadFile(jobPath)
		if err != nil {
			return errorMsg(err)
		}
		var job sbatch.Job
		if err := yaml.Unmarshal(dat, &job); err != nil {
			return errorMsg(err)
		}

		// Validate input
		if m.inputs[jobNameInput].Value() == "" {
			err := errors.New("empty value is not allowed")
			m.errors[jobNameInput] = err
			return fmt.Errorf("job name field: %w", err)
		}
		if err := validator.IsNumber(m.inputs[creditsLockingInput].Value()); err != nil {
			m.errors[creditsLockingInput] = err
			return fmt.Errorf("allocate credits field: %w", err)
		}
		if err := validator.IsMap(m.inputs[usesInput].Value()); err != nil {
			m.errors[usesInput] = err
			return fmt.Errorf("use flags field: %w", err)
		}

		// Parse input
		allocatedCredits := m.inputs[creditsLockingInput].Value()
		allocatedCreditsF, err := strconv.ParseFloat(allocatedCredits, 64)
		if err != nil {
			return errorMsg(err)
		}
		allocatedCreditsBigF := new(big.Float).SetFloat64(allocatedCreditsF)
		curr, err := m.allowanceManager.GetAllowance(ctx)
		if err != nil {
			return errorMsg(err)
		}
		labels, err := utils.StringsToLabels(m.inputs[usesInput].Value())
		if err != nil {
			return errorMsg(err)
		}

		// Mutate
		allocatedCreditsBigI := ether.ToWei(allocatedCreditsBigF)
		if err = m.allowanceManager.SetAllowance(ctx, curr.Add(curr, allocatedCreditsBigI)); err != nil {
			return errorMsg(err)
		}

		var jobName [32]byte
		copy(jobName[:], m.inputs[jobNameInput].Value())
		_, err = m.jobScheduler.SubmitJob(
			ctx,
			&job,
			labels,
			allocatedCreditsBigI,
			jobName,
		)
		if err != nil {
			return errorMsg(err)
		}
		return submitDoneMsg{}
	}
}

func (m *model) initEditor(ctx context.Context) tea.Cmd {
	jobSchemaPath, jobPath, err := PrepareFiles()
	if err != nil {
		panic(err.Error())
	}
	return m.openEditor(ctx, jobSchemaPath, jobPath)
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.initEditor(context.TODO()), textinput.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		codeCmd  tea.Cmd
		inputCmd tea.Cmd
		cmds     = make([]tea.Cmd, len(m.inputs))
	)

switchmsg:
	switch msg := msg.(type) {
	case ExitMsg:
		Clean(m.jobSchemaPath, m.jobPath)
	case editorDone:
		if msg.err != nil {
			m.err = msg.err
			break switchmsg
		}
		m.jobPath = msg.jobPath
		m.jobSchemaPath = msg.jobSchemaPath
		setFile := m.code.SetFileName(m.jobPath)
		cmds = append(cmds, setFile)
	case submitProgressMsg:
		m.isRunning = true
	case submitDoneMsg:
		m.isRunning = false
		cmds = append(cmds, emitExitMsg)
	case errorMsg:
		m.isRunning = false
		m.err = msg
		return m, nil
	case tea.KeyMsg:
		if m.isRunning {
			break switchmsg
		}
		switch {
		case key.Matches(msg, m.keyMap.EditAgain):
			cmds = append(cmds, m.openEditor(context.TODO(), m.jobSchemaPath, m.jobPath))
		case key.Matches(msg, m.keyMap.Exit):
			cmds = append(cmds, emitExitMsg)
		case msg.String() == "enter" && m.focused == len(m.inputs)-1:
			cmds = append(cmds, m.submitJob(context.TODO(), m.jobPath), emitSubmitProgressMsg)
		case key.Matches(msg, m.keyMap.NextInput):
			m.nextInput()
		case key.Matches(msg, m.keyMap.PrevInput):
			m.prevInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

		m.inputs[creditsLockingInput], inputCmd = m.inputs[creditsLockingInput].Update(msg)
		cmds = append(cmds, inputCmd)

		m.inputs[usesInput], inputCmd = m.inputs[usesInput].Update(msg)
		cmds = append(cmds, inputCmd)

		m.inputs[jobNameInput], inputCmd = m.inputs[jobNameInput].Update(msg)
		cmds = append(cmds, inputCmd)
	}

	m.code, codeCmd = m.code.Update(msg)
	if codeCmd != nil {
		cmds = append(cmds, codeCmd)
	}

	return m, tea.Batch(cmds...)
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func isNumberCharacter(ch rune) bool {
	return unicode.IsDigit(ch) || ch == 'e' || ch == '.'
}

func allowedNumber(input string) error {
	for _, ch := range input {
		if !isNumberCharacter(ch) {
			return fmt.Errorf("character '%c' is not allowed", ch)
		}
	}
	return nil
}
