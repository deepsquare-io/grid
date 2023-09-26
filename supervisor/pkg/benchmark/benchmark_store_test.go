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
