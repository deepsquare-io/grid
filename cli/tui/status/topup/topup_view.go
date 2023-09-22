package topup

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/internal/utils"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/tui/style"
)

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func (m model) loading() string {
	if m.isRunning {
		return "Topping up..."
	}
	return ""
}

func (m model) View() string {
	help := m.help.FullHelpView([][]key.Binding{
		{
			m.keyMap.NextInput,
			m.keyMap.PrevInput,
			m.keyMap.Exit,
		},
	})
	prices := `CPU pricing: ... credits/(CPU.min)
Memory pricing: ... credits/(MB.min)
GPU pricing: ... credits/(GPU.min)
`
	if m.prices != nil {
		prices = fmt.Sprintf(`CPU pricing: %s credits/(CPU.min)
Memory pricing: %s credits/(MB.min)
GPU pricing: %s credits/(GPU.min)
`,
			ether.FromWei(m.prices.CpuPricePerMin).String(),
			ether.FromWei(m.prices.MemPricePerMin).String(),
			ether.FromWei(m.prices.GpuPricePerMin).String(),
		)
	}
	definition := `Tasks: ...
CPU/task: ... (Total CPU: ...)
Mem(MB)/CPU: ... MB (Total Mem: ... MB)
GPU/task: ... (Total GPU: ...)
`
	if m.job != nil {
		definition = fmt.Sprintf(`Tasks: %d
CPU/task: %d (Total CPU: %d)
Mem(MB)/CPU: %d MB (Total Mem: %d MB)
GPU/task: %d (Total GPU: %d)
`,
			m.job.Definition.Ntasks,
			m.job.Definition.CpusPerTask,
			m.job.Definition.CpusPerTask*m.job.Definition.Ntasks,
			m.job.Definition.MemPerCpu,
			m.job.Definition.MemPerCpu*m.job.Definition.CpusPerTask*m.job.Definition.Ntasks,
			m.job.Definition.GpusPerTask,
			m.job.Definition.GpusPerTask*m.job.Definition.Ntasks,
		)
	}

	duration := "N/A"
	if m.job != nil && m.prices != nil {
		func() {
			value, err := strconv.ParseFloat(m.inputs[amountInput].Value(), 64)
			if err != nil {
				return
			}
			durationB := metascheduler.CreditToDuration(
				*m.prices,
				m.job.Definition,
				ether.ToWei(big.NewFloat(value)),
			)
			duration = durationB.String() + " minutes"
		}()

	}

	return fmt.Sprintf(`%s
%s
%s

%s
%s

%s
%s
%s

Expected duration gain: %s

%s
%s
%s`,
		style.Title1.Width(20).Render("Topping up job"),
		style.Title2.Width(20).Render("Prices"),
		indent(2, prices),
		style.Title2.Width(20).Render("Job Allocation"),
		indent(2, definition),
		style.Foreground.Render("Amount in credits (not in wei)"),
		m.inputs[amountInput].View(),
		style.Error.Render(utils.ErrorfOrEmpty("^^^%s", m.errors[amountInput])),
		duration,
		style.Error.Render(utils.ErrorfOrEmpty("Error: %s", m.err)),
		m.loading(),
		help,
	)
}
