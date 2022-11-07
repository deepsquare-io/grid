package eth

import (
	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
)

func JobDefinitionMapToSlurm(j metascheduler.JobDefinition, t uint64, body string) slurm.JobDefinition {
	return slurm.JobDefinition{
		NTasks:       j.Ntasks,
		GPUsPerTask:  j.GpuPerTask,
		CPUsPerTask:  j.CpuPerTask,
		TimeLimit:   t,
		MemoryPerCPU: j.MemPerCpu,
		Body:         body,
	}
}
