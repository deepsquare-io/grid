// Copyright (C) 2023 DeepSquare Asociation
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

package topup

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

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
	provider := fmt.Sprintf(`%s
  CPU pricing: ... credits/(CPU.min)
  Memory pricing: ... credits/(MB.min)
  GPU pricing: ... credits/(GPU.min)
`, style.Title2.Width(52).Render("Provider ..."))
	if m.provider != nil {
		provider = fmt.Sprintf(`%s
  CPU pricing: %s credits/(CPU.min)
  Memory pricing: %s credits/(MB.min)
  GPU pricing: %s credits/(GPU.min)
`,
			style.Title2.Width(52).Render("Provider", m.provider.Addr.Hex()),
			ether.FromWei(m.provider.ProviderPrices.CpuPricePerMin).String(),
			ether.FromWei(m.provider.ProviderPrices.MemPricePerMin).String(),
			ether.FromWei(m.provider.ProviderPrices.GpuPricePerMin).String(),
		)
	}
	definition := `Tasks: ...
CPU/task: ... (Total CPU: ...)
Mem(MB)/CPU: ... MB (Total Mem: ... MB)
GPU/task: ... (Total GPU: ...)
Credits Allocated: ... credits
Max Duration: ...
Elapsed: ...
`
	if m.job != nil {
		maxDuration := "..."
		if m.provider != nil {
			maxDuration = (time.Duration(metascheduler.CreditToDuration(
				m.provider.ProviderPrices,
				m.job.Definition,
				m.job.Cost.MaxCost,
			).Int64()) * time.Second).String()
		}
		var elapsed time.Duration
		start := m.job.Time.Start.Int64()
		end := m.job.Time.End.Int64()
		if end != 0 {
			elapsed = time.Duration(end-start) * time.Second
		} else {
			elapsed = time.Since(time.Unix(start, 0))
		}
		definition = fmt.Sprintf(
			`Tasks: %d
CPU/task: %d (Total CPU: %d)
Mem(MB)/CPU: %d MB (Total Mem: %d MB)
GPU/task: %d (Total GPU: %d)
Credits Allocated: %s credits
Max Duration: %s
Elapsed: %s
`,
			m.job.Definition.Ntasks,
			m.job.Definition.CpusPerTask,
			m.job.Definition.CpusPerTask*m.job.Definition.Ntasks,
			m.job.Definition.MemPerCpu,
			m.job.Definition.MemPerCpu*m.job.Definition.CpusPerTask*m.job.Definition.Ntasks,
			m.job.Definition.GpusPerTask,
			m.job.Definition.GpusPerTask*m.job.Definition.Ntasks,
			ether.FromWei(m.job.Cost.MaxCost).String(),
			maxDuration,
			elapsed,
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
			duration = (time.Duration(durationB.Int64()) * time.Minute).String()
		}()

	}

	return fmt.Sprintf(
		`%s
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
		style.Title2.Width(80).
			Render(fmt.Sprintf("Job %s", new(big.Int).SetBytes(m.jobID[:]).String())),
		indent(2, definition),
		provider,
		style.Foreground.Render("Amount in credits (not in wei)"),
		m.inputs[amountInput].View(),
		style.Error.Render(utils.ErrorfOrEmpty("^^^%s", m.errors[amountInput])),
		duration,
		style.Error.Render(utils.ErrorfOrEmpty("Error: %s", m.err)),
		m.loading(),
		help,
	)
}
