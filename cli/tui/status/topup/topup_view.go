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
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	provider := `Provider address: ...
CPU pricing: ... credits/(CPU.min)
Memory pricing: ... credits/(MB.min)
GPU pricing: ... credits/(GPU.min)
`
	if m.provider != nil {
		provider = fmt.Sprintf(`Provider address: %s
CPU pricing: %s credits/(CPU.min)
Memory pricing: %s credits/(MB.min)
GPU pricing: %s credits/(GPU.min)
`,
			m.provider.Addr.Hex(),
			ether.FromWei(m.provider.ProviderPrices.CpuPricePerMin).String(),
			ether.FromWei(m.provider.ProviderPrices.MemPricePerMin).String(),
			ether.FromWei(m.provider.ProviderPrices.GpuPricePerMin).String(),
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
	if m.job != nil && m.provider != nil {
		func() {
			value, err := strconv.ParseFloat(m.inputs[amountInput].Value(), 64)
			if err != nil {
				return
			}
			durationB := metascheduler.CreditToDuration(
				m.provider.ProviderPrices,
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
		style.Title2.Width(80).Render(fmt.Sprintf("Job ID: %s", hexutil.Encode(m.jobID[:]))),
		indent(2, definition),
		style.Title2.Render("Provider"),
		indent(2, provider),
		style.Foreground.Render("Amount in credits (not in wei)"),
		m.inputs[amountInput].View(),
		style.Error.Render(utils.ErrorfOrEmpty("^^^%s", m.errors[amountInput])),
		duration,
		style.Error.Render(utils.ErrorfOrEmpty("Error: %s", m.err)),
		m.loading(),
		help,
	)
}
