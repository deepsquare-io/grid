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

package topup

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/deepsquare-io/grid/cli/internal/ether"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/tui/style"
)

func indent(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func (m model) loading() string {
	if m.isToppingUp {
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
`, style.Title2().Width(52).Render("Provider ..."))
	if m.provider != nil {
		provider = fmt.Sprintf(`%s
  CPU pricing: %s credits/(CPU.min)
  Memory pricing: %s credits/(MB.min)
  GPU pricing: %s credits/(GPU.min)
`,
			style.Title2().Width(52).Render("Provider", m.provider.Addr.Hex()),
			ether.FromWei(m.provider.ProviderPrices.CpuPricePerMin).String(),
			ether.FromWei(m.provider.ProviderPrices.MemPricePerMin).String(),
			ether.FromWei(m.provider.ProviderPrices.GpuPricePerMin).String(),
		)
	}
	definition := `Name: ...
Status: ...
LastError: ...
Tasks: ...
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
			maxDurationB, err := metascheduler.CreditToDuration(
				m.provider.ProviderPrices,
				m.job.Definition,
				m.job.Cost.MaxCost,
			)
			if err != nil {
				maxDuration = "NaN"
			} else {
				maxDuration = (time.Duration(maxDurationB.Int64()) * time.Minute).String()
			}
		}
		var elapsed time.Duration
		start := m.job.Time.Start.Int64()
		end := m.job.Time.End.Int64()
		if end >= start || isFinished(metascheduler.JobStatus(m.job.Status)) {
			elapsed = time.Duration(end-start) * time.Second
		} else {
			elapsed = time.Since(time.Unix(start, 0)).Truncate(time.Second)
		}
		definition = fmt.Sprintf(
			`Name: %s
Status: %s
LastError: %s
Tasks: %d
CPU/task: %d (Total CPU: %d)
Mem(MB)/CPU: %d MB (Total Mem: %d MB)
GPUs: %d
Credits Allocated: %s credits
Max Duration: %s
Elapsed: %s
`,
			string(m.job.JobName[:]),
			style.JobStatusStyle(metascheduler.JobStatus(m.job.Status).String()).
				Render(metascheduler.JobStatus(m.job.Status).String()),
			renderOnError(style.Error().Render(m.job.LastError), m.job.LastError),
			m.job.Definition.Ntasks,
			m.job.Definition.CpusPerTask,
			m.job.Definition.CpusPerTask*m.job.Definition.Ntasks,
			m.job.Definition.MemPerCpu,
			m.job.Definition.MemPerCpu*m.job.Definition.CpusPerTask*m.job.Definition.Ntasks,
			m.job.Definition.Gpus,
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
			durationB, err := metascheduler.CreditToDuration(
				m.provider.ProviderPrices,
				m.job.Definition,
				ether.ToWei(big.NewFloat(value)),
			)
			if err != nil {
				duration = "NaN"
				return
			}
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
		style.Title1().Width(20).Render("Topping up job"),
		style.Title2().Width(80).
			Render(fmt.Sprintf("Job %s", new(big.Int).SetBytes(m.jobID[:]).String())),
		indent(2, definition),
		provider,
		style.Foreground().Render("Amount in credits (not in wei)"),
		m.inputs[amountInput].View(),
		style.Error().Width(50).Render(utils.FormatErrorfOrEmpty("^^^%s", m.errors[amountInput])),
		duration,
		style.Error().Width(50).Render(utils.FormatErrorfOrEmpty("Error: %s", m.err)),
		m.loading(),
		help,
	)
}

func isFinished(s metascheduler.JobStatus) bool {
	return s == metascheduler.JobStatusCancelled ||
		s == metascheduler.JobStatusFinished ||
		s == metascheduler.JobStatusOutOfCredits ||
		s == metascheduler.JobStatusPanicked ||
		s == metascheduler.JobStatusFailed
}

func renderOnError(v string, err string) string {
	if err == "" {
		return "no error reported"
	}
	return v
}
