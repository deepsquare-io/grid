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
	var fakeIORResult ior.Result
	var fakeMachineSpec benchmark.MachineSpec
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

	require.Equal(t, benchmark.Data{
		AllToAllCollectiveLatency: 1,
		DownloadBandwidth:         2,
		GFLOPS:                    3,
		P2PBidirectionalBandwidth: 4,
		P2PLatency:                5,
		UploadBandwidth:           6,
		ScratchAvgRead:            &fakeIORResult,
		ScratchAvgWrite:           &fakeIORResult,
		SharedWorldTmpAvgRead:     &fakeIORResult,
		SharedWorldTmpAvgWrite:    &fakeIORResult,
		SharedTmpAvgRead:          &fakeIORResult,
		SharedTmpAvgWrite:         &fakeIORResult,
		DiskWorldTmpAvgRead:       &fakeIORResult,
		DiskWorldTmpAvgWrite:      &fakeIORResult,
		DiskTmpAvgRead:            &fakeIORResult,
		DiskTmpAvgWrite:           &fakeIORResult,
		MachineSpec:               &fakeMachineSpec,
	}, store.Dump())
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
