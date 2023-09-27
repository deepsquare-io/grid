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

package benchmark

import (
	"context"
	"sync"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/ior"
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

	MachineSpec *MachineSpec
}

// Store is a simple store with hard-coded keys to store benchmark results.
type Store struct {
	mu sync.RWMutex // Mutex for thread safety

	data Data

	refresh chan struct{}
	failure chan error
}

func NewStore() *Store {
	return &Store{
		refresh: make(chan struct{}, 1),
		failure: make(chan error, 10),
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

func (s *Store) SetMachineSpec(spec *MachineSpec) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data.MachineSpec = spec
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info(
		"stored benchmark result",
		zap.Any("spec", spec),
	)
}

func (s *Store) SetFailure(err error) {
	s.failure <- err
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
		s.data.DiskWorldTmpAvgWrite != nil &&
		s.data.MachineSpec != nil)
}

func (s *Store) WaitForCompletion(ctx context.Context) (done chan struct{}, errc chan error) {
	done = make(chan struct{}, 1)

	go func() {
		defer close(done)
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

	return done, s.failure
}
