package benchmark_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/stretchr/testify/require"
)

func TestWaitForCompletion(t *testing.T) {
	go func() {
		benchmark.DefaultStore.SetAllToAllCollectiveLatency(1)
		benchmark.DefaultStore.SetDownloadBandwidth(2)
		benchmark.DefaultStore.SetGFLOPS(3)
		benchmark.DefaultStore.SetP2PBidirectionalBandwidth(4)
		benchmark.DefaultStore.SetP2PLatency(5)
		benchmark.DefaultStore.SetUploadBandwidth(6)
		fmt.Println("set")
	}()

	<-benchmark.DefaultStore.WaitForCompletion(context.Background())

	require.Equal(t, benchmark.Data{
		AllToAllCollectiveLatency: 1,
		DownloadBandwidth:         2,
		GFLOPS:                    3,
		P2PBidirectionalBandwidth: 4,
		P2PLatency:                5,
		UploadBandwidth:           6,
	}, benchmark.DefaultStore.Dump())
}
