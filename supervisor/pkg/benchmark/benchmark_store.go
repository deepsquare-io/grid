package benchmark

import (
	"context"
	"sync"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/ior"
	"go.uber.org/zap"
)

var DefaultStore = NewStore()

type Data struct {
	// UploadBandwidth is in bps
	UploadBandwidth uint64
	// DownloadBandwidth is in bps
	DownloadBandwidth uint64
	GFLOPS            float64
	// P2PBidirectionalBandwidth is in MB/s
	P2PBidirectionalBandwidth float64
	// P2PBidirectionalBandwidth is in us
	AllToAllCollectiveLatency float64
	// P2PLatency is in us
	P2PLatency             float64
	ScratchAvgRead         *ior.Result
	ScratchAvgWrite        *ior.Result
	SharedWorldTmpAvgRead  *ior.Result
	SharedWorldTmpAvgWrite *ior.Result
	SharedTmpAvgRead       *ior.Result
	SharedTmpAvgWrite      *ior.Result
	DiskWorldTmpAvgRead    *ior.Result
	DiskWorldTmpAvgWrite   *ior.Result
	DiskTmpAvgRead         *ior.Result
	DiskTmpAvgWrite        *ior.Result
}

// Store is a simple store with hard-coded keys to store benchmark results.
type Store struct {
	mu sync.RWMutex // Mutex for thread safety

	data Data

	refresh chan struct{}
}

func NewStore() *Store {
	return &Store{
		refresh: make(chan struct{}, 1),
	}
}

func (s *Store) SetUploadBandwidth(upload uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.UploadBandwidth = upload
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Uint64("upload-bandwidth", upload))
}

func (s *Store) SetDownloadBandwidth(download uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.DownloadBandwidth = download
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Uint64("download-bandwidth", download))
}

func (s *Store) SetGFLOPS(gflops float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.GFLOPS = gflops
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("gflops", gflops))
}

func (s *Store) SetP2PBidirectionalBandwidth(bandwidth float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.P2PBidirectionalBandwidth = bandwidth
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("p2p-bibw", bandwidth))
}

func (s *Store) SetAllToAllCollectiveLatency(latency float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.AllToAllCollectiveLatency = latency
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("alltoall-latency", latency))
}

func (s *Store) SetP2PLatency(latency float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.P2PLatency = latency
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("p2p-latency", latency))
}

func (s *Store) SetScratchResult(avgr, avgw *ior.Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.ScratchAvgRead = avgr
	s.data.ScratchAvgWrite = avgw
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info(
		"stored benchmark result",
		zap.Any("scratch-read", avgr),
		zap.Any("scratch-write", avgw),
	)
}
func (s *Store) SetSharedWorldTmpResult(avgr, avgw *ior.Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.SharedWorldTmpAvgRead = avgr
	s.data.SharedWorldTmpAvgWrite = avgw
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info(
		"stored benchmark result",
		zap.Any("shared-world-tmp-read", avgr),
		zap.Any("shared-world-tmp-write", avgw),
	)
}
func (s *Store) SetSharedTmpResult(avgr, avgw *ior.Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.SharedTmpAvgRead = avgr
	s.data.SharedTmpAvgWrite = avgw
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info(
		"stored benchmark result",
		zap.Any("shared-tmp-read", avgr),
		zap.Any("shared-tmp-write", avgw),
	)
}
func (s *Store) SetDiskWorldTmpResult(avgr, avgw *ior.Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.DiskWorldTmpAvgRead = avgr
	s.data.DiskWorldTmpAvgWrite = avgw
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info(
		"stored benchmark result",
		zap.Any("disk-world-tmp-read", avgr),
		zap.Any("disk-world-tmp-write", avgw),
	)
}
func (s *Store) SetDiskTmpResult(avgr, avgw *ior.Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.DiskTmpAvgRead = avgr
	s.data.DiskTmpAvgWrite = avgw
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info(
		"stored benchmark result",
		zap.Any("disk-tmp-read", avgr),
		zap.Any("disk-tmp-write", avgw),
	)
}

func (s *Store) Dump() Data {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data
}

func (s *Store) IsComplete() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return (s.data.UploadBandwidth != 0 &&
		s.data.DownloadBandwidth != 0 &&
		s.data.GFLOPS != 0 &&
		s.data.P2PBidirectionalBandwidth != 0 &&
		s.data.AllToAllCollectiveLatency != 0 &&
		s.data.P2PLatency != 0 &&
		s.data.ScratchAvgRead != nil &&
		s.data.ScratchAvgWrite != nil &&
		s.data.SharedWorldTmpAvgRead != nil &&
		s.data.SharedWorldTmpAvgWrite != nil &&
		s.data.SharedTmpAvgRead != nil &&
		s.data.SharedTmpAvgWrite != nil &&
		s.data.DiskTmpAvgRead != nil &&
		s.data.DiskTmpAvgWrite != nil &&
		s.data.DiskWorldTmpAvgRead != nil &&
		s.data.DiskWorldTmpAvgWrite != nil)
}

func (s *Store) WaitForCompletion(ctx context.Context) chan struct{} {
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-s.refresh:
				if s.IsComplete() {
					done <- struct{}{}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return done
}
