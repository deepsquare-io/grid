package metascheduler

import (
	"slices"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
)

func ProviderHardwareEqual(
	a metaschedulerabi.ProviderHardware,
	b metaschedulerabi.ProviderHardware,
) bool {
	return slices.Equal(a.CpusPerNode, b.CpusPerNode) &&
		slices.Equal(a.GpusPerNode, b.GpusPerNode) &&
		slices.Equal(a.MemPerNode, b.MemPerNode) &&
		a.Nodes == b.Nodes
}

func LabelsEqual(
	a []metaschedulerabi.Label,
	b []metaschedulerabi.Label,
) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
