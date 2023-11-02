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

package metascheduler

import (
	"regexp"
	"slices"
	"strings"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
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

func ProviderPricesEqual(
	a metaschedulerabi.ProviderPrices,
	b metaschedulerabi.ProviderPrices,
) bool {
	return a.CpuPricePerMin.Cmp(b.CpuPricePerMin) == 0 &&
		a.GpuPricePerMin.Cmp(b.GpuPricePerMin) == 0 &&
		a.MemPricePerMin.Cmp(b.MemPricePerMin) == 0
}

// LabelsContains returns true if a is included in b
func LabelsContains(a, b []metaschedulerabi.Label) bool {
	for _, labelA := range a {
		found := false
		for _, labelB := range b {
			if labelA == labelB {
				found = true
				break
			}
		}
		if !found {
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
