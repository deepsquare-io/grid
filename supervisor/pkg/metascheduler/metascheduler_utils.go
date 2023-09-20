package metascheduler

import (
	"regexp"
	"slices"
	"strings"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
)

var labelRegex = regexp.MustCompile(`[a-zA-Z0-9\.\_]([-a-zA-Z0-9\.\_]*[a-zA-Z0-9\.\_])?`)

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

// MergeLabels b on a.
//
// b has priority over a.
func MergeLabels(
	a []metaschedulerabi.Label,
	b []metaschedulerabi.Label,
) []metaschedulerabi.Label {
	mergedMap := make(map[string]string)

	for _, kv := range a {
		mergedMap[kv.Key] = kv.Value
	}

	// Merge values from the second array, override conflicts
	for _, kv := range b {
		mergedMap[kv.Key] = kv.Value
	}

	// Convert the map back to an array of KeyValue structs
	mergedArray := make([]metaschedulerabi.Label, 0, len(mergedMap))
	for key, value := range mergedMap {
		mergedArray = append(mergedArray, metaschedulerabi.Label{Key: key, Value: value})
	}

	return mergedArray
}

func ProcessLabel(l metaschedulerabi.Label) metaschedulerabi.Label {
	return metaschedulerabi.Label{
		Key:   strings.ToLower(strings.Join(labelRegex.FindAllString(l.Key, -1), "-")),
		Value: strings.ToLower(strings.Join(labelRegex.FindAllString(l.Value, -1), "-")),
	}
}

func ProcessLabels(ll []metaschedulerabi.Label) []metaschedulerabi.Label {
	o := make([]metaschedulerabi.Label, 0, len(ll))

	for _, l := range ll {
		o = append(o, ProcessLabel(l))
	}

	return o
}
