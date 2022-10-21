package eth

import (
	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
)

func JobDefinitionMapToSlurm(j metascheduler.JobDefinition, t uint64, body string) slurm.JobDefinition {
	return slurm.JobDefinition{
		NTasks:        j.Ntasks,
		GPUsPerNode:   j.GpuPerNode,
		CPUsPerTask:   j.CpuPerTask,
		TimeLimit:     t,
		Nodes:         j.Nodes,
		MemoryPerNode: j.MemPerNode,
		Body:          body,
	}
}
