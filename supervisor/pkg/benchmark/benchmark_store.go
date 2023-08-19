package benchmark

import (
	"context"
	"sync"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

var DefaultStore = &Store{}

type Data struct {
	UploadBandwidth           uint64
	DownloadBandwidth         uint64
	GFLOPS                    float64
	P2PBidirectionalBandwidth float64
	AllToAllCollectiveLatency float64
	P2PLatency                float64
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
		s.data.P2PLatency != 0)
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
