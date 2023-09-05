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

func MergeLabels(
	a []metaschedulerabi.Label,
	b []metaschedulerabi.Label,
) []metaschedulerabi.Label {
	// Create a map to store the merged values while avoiding conflicts
	mergedMap := make(map[string]string)

	for _, kv := range a {
		mergedMap[kv.Key] = kv.Value
	}

	// Merge values from the second array, avoiding conflicts
	for _, kv := range b {
		if _, exists := mergedMap[kv.Key]; !exists {
			mergedMap[kv.Key] = kv.Value
		}
	}

	// Convert the map back to an array of KeyValue structs
	mergedArray := make([]metaschedulerabi.Label, 0, len(mergedMap))
	for key, value := range mergedMap {
		mergedArray = append(mergedArray, metaschedulerabi.Label{Key: key, Value: value})
	}

	return mergedArray
}
