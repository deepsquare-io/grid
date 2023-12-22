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

package benchmark_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/ior"
	"github.com/stretchr/testify/require"
)

func TestWaitForCompletion(t *testing.T) {
	fakeIORResult := ior.Result{
		IOPS:      1,
		Bandwidth: 2,
	}
	fakeMachineSpec := benchmark.MachineSpec{
		MicroArch: "a",
		OS:        "b",
		CPU:       "c",
		Arch:      "d",
		GPU:       "e",
	}
	store := benchmark.NewStore()
	go func() {
		store.SetAllToAllCollectiveLatency(1)
		store.SetDownloadBandwidth(2)
		store.SetGFLOPS(3)
		store.SetP2PBidirectionalBandwidth(4)
		store.SetP2PLatency(5)
		store.SetUploadBandwidth(6)
		store.SetScratchResult(&fakeIORResult, &fakeIORResult)
		store.SetSharedWorldTmpResult(&fakeIORResult, &fakeIORResult)
		store.SetSharedTmpResult(&fakeIORResult, &fakeIORResult)
		store.SetDiskTmpResult(&fakeIORResult, &fakeIORResult)
		store.SetDiskWorldTmpResult(&fakeIORResult, &fakeIORResult)
		store.SetMachineSpec(&fakeMachineSpec)
		fmt.Println("set")
	}()

	done, errc := store.WaitForCompletion(context.Background())
	select {
	case <-done:
	case err := <-errc:
		require.NoError(t, err)
	}

	expected := map[string]string{
		"arch":                                    "d",
		"compute.gflops":                          "3.00",
		"cpu":                                     "c",
		"cpu.microarch":                           "a",
		"gpu":                                     "c",
		"network.all-to-all.latency.us":           "1.00",
		"network.download.bw.mbps":                "0.00",
		"network.p2p.bw.mbps":                     "4.00",
		"network.p2p.latency.us":                  "5.00",
		"network.upload.bw.mbps":                  "0.00",
		"os":                                      "b",
		"storage.disk-tmp.read.bw.mibps":          "2.00",
		"storage.disk-tmp.read.iops":              "1.00",
		"storage.disk-tmp.write.bw.mibps":         "2.00",
		"storage.disk-tmp.write.iops":             "1.00",
		"storage.disk-world-tmp.read.bw.mibps":    "2.00",
		"storage.disk-world-tmp.read.iops":        "1.00",
		"storage.disk-world-tmp.write.bw.mibps":   "2.00",
		"storage.disk-world-tmp.write.iops":       "1.00",
		"storage.scratch.read.bw.mibps":           "2.00",
		"storage.scratch.read.iops":               "1.00",
		"storage.scratch.write.bw.mibps":          "2.00",
		"storage.scratch.write.iops":              "1.00",
		"storage.shared-tmp.read.bw.mibps":        "2.00",
		"storage.shared-tmp.read.iops":            "1.00",
		"storage.shared-tmp.write.bw.mibps":       "2.00",
		"storage.shared-tmp.write.iops":           "1.00",
		"storage.shared-world-tmp.read.bw.mibps":  "2.00",
		"storage.shared-world-tmp.read.iops":      "1.00",
		"storage.shared-world-tmp.write.bw.mibps": "2.00",
		"storage.shared-world-tmp.write.iops":     "1.00",
	}
	actual := store.Dump()
	for k, v := range expected {
		require.Equal(t, v, actual[k])
	}
}

func TestWaitForCompletionFailure(t *testing.T) {
	store := benchmark.NewStore()
	expectErr := errors.New("fail")
	go func() {
		store.SetAllToAllCollectiveLatency(1)
		store.SetDownloadBandwidth(2)
		store.SetGFLOPS(3)
		store.SetP2PBidirectionalBandwidth(4)
		store.SetFailure(expectErr)
		store.SetP2PLatency(5)
	}()

	done, errc := store.WaitForCompletion(context.Background())
	select {
	case <-done:
		t.FailNow()
	case err := <-errc:
		require.EqualError(t, err, expectErr.Error())
	}
}
