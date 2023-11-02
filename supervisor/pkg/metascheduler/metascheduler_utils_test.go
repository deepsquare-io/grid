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

package metascheduler_test

import (
	"testing"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/stretchr/testify/require"
)

func TestProcessLabels(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"my_invalid-label123!", "my_invalid-label123"},
		{"validLabel123", "validlabel123"},
		{".a@b#c$d%e^f&g.f", ".a-b-c-d-e-f-g.f"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := metascheduler.ProcessLabel(metaschedulerabi.Label{
				Key:   "",
				Value: tc.input,
			})
			require.Equal(t, tc.expected, result.Value)
		})
	}
}

func TestLabelsContains(t *testing.T) {
	// Define some test cases
	testCases := []struct {
		name     string
		sliceA   []metaschedulerabi.Label
		sliceB   []metaschedulerabi.Label
		expected bool
	}{
		{
			name: "Included case",
			sliceA: []metaschedulerabi.Label{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
			},
			sliceB: []metaschedulerabi.Label{
				{Key: "key1", Value: "value1"},
				{Key: "key3", Value: "value3"},
				{Key: "key2", Value: "value2"},
			},
			expected: true,
		},
		{
			name: "Included case 2",
			sliceA: []metaschedulerabi.Label{
				{Key: "name", Value: "basel-1"},
				{Key: "region", Value: "ch-basel"},
				{Key: "zone", Value: "ch-basel-1"},
			},
			sliceB: []metaschedulerabi.Label{
				{Key: "storage.scratch.write.bw.mibps", Value: "66.15"},
				{Key: "storage.shared-world-tmp.read.iops", Value: "54.99"},
				{Key: "os", Value: "linux"},
				{Key: "gpu", Value: "nvidia-geforce-rtx-3090"},
				{Key: "compute.gflops", Value: "94690.00"},
				{Key: "network.p2p.bw.mbps", Value: "2314.03"},
				{Key: "storage.shared-tmp.write.bw.mibps", Value: "67.18"},
				{Key: "storage.scratch.read.bw.mibps", Value: "109.73"},
				{Key: "storage.scratch.read.iops", Value: "54.88"},
				{Key: "storage.disk-tmp.read.bw.mibps", Value: "88006.11"},
				{Key: "arch", Value: "amd64"},
				{Key: "storage.shared-world-tmp.write.bw.mibps", Value: "66.20"},
				{Key: "storage.shared-tmp.read.iops", Value: "54.92"},
				{Key: "storage.disk-world-tmp.read.iops", Value: "44810.94"},
				{Key: "network.p2p.latency.us", Value: "3635.51"},
				{Key: "storage.shared-tmp.read.bw.mibps", Value: "109.82"},
				{Key: "storage.disk-world-tmp.write.bw.mibps", Value: "928.51"},
				{Key: "storage.disk-tmp.read.iops", Value: "44570.06"},
				{Key: "storage.disk-tmp.write.iops", Value: "463.15"},
				{Key: "cpu", Value: "amd-epyc-7302-16-core-processor"},
				{Key: "cpu.microarch", Value: "zen2"},
				{Key: "network.all-to-all.latency.us", Value: "2831.69"},
				{Key: "storage.shared-world-tmp.write.iops", Value: "33.11"},
				{Key: "name", Value: "basel-1"},
				{Key: "network.upload.bw.mbps", Value: "1171.54"},
				{Key: "storage.scratch.write.iops", Value: "33.08"},
				{Key: "storage.disk-world-tmp.read.bw.mibps", Value: "89579.60"},
				{Key: "storage.disk-tmp.write.bw.mibps", Value: "926.27"},
				{Key: "region", Value: "ch-basel"},
				{Key: "zone", Value: "ch-basel-1"},
				{Key: "network.download.bw.mbps", Value: "1157.54"},
				{Key: "storage.shared-world-tmp.read.bw.mibps", Value: "109.97"},
				{Key: "storage.shared-tmp.write.iops", Value: "33.60"},
				{Key: "storage.disk-world-tmp.write.iops", Value: "464.27"},
			},
			expected: true,
		},
		{
			name: "Not included case",
			sliceA: []metaschedulerabi.Label{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
				{Key: "key4", Value: "value4"},
			},
			sliceB: []metaschedulerabi.Label{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: "value2"},
				{Key: "key3", Value: "value3"},
			},
			expected: false,
		},
	}

	// Iterate over the test cases and run the tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := metascheduler.LabelsContains(tc.sliceA, tc.sliceB)
			require.Equal(t, tc.expected, result)
		})
	}
}
