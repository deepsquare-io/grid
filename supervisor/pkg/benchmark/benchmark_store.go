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
	"fmt"
	"sync"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/ior"
	"go.uber.org/zap"
)

var keys = []string{
	cpuKey, cpuMicroArchKey, osKey, archKey, gpuKey, computeKey,
	networkUploadBWKey, networkDownloadBWKey, networkP2PBWKey,
	networkP2PLatencyKey, networkAllToAllLatencyKey,
	storageScratchReadBWKey, storageScratchReadIOPSKey,
	storageScratchWriteBWKey, storageScratchWriteIOPSKey,
	storageSharedWorldTmpReadBWKey, storageSharedWorldTmpReadIOPSKey,
	storageSharedWorldTmpWriteBWKey, storageSharedWorldTmpWriteIOPSKey,
	storageSharedTmpReadBWKey, storageSharedTmpReadIOPSKey,
	storageSharedTmpWriteBWKey, storageSharedTmpWriteIOPSKey,
	storageDiskWorldTmpReadBWKey, storageDiskWorldTmpReadIOPSKey,
	storageDiskWorldTmpWriteBWKey, storageDiskWorldTmpWriteIOPSKey,
	storageDiskTmpReadBWKey, storageDiskTmpReadIOPSKey,
	storageDiskTmpWriteBWKey, storageDiskTmpWriteIOPSKey,
}

const (
	cpuKey                            = "cpu"
	cpuMicroArchKey                   = "cpu.microarch"
	osKey                             = "os"
	archKey                           = "arch"
	gpuKey                            = "gpu"
	computeKey                        = "compute.gflops"
	networkUploadBWKey                = "network.upload.bw.mbps"
	networkDownloadBWKey              = "network.download.bw.mbps"
	networkP2PBWKey                   = "network.p2p.bw.mbps"
	networkP2PLatencyKey              = "network.p2p.latency.us"
	networkAllToAllLatencyKey         = "network.all-to-all.latency.us"
	storageScratchReadBWKey           = "storage.scratch.read.bw.mibps"
	storageScratchReadIOPSKey         = "storage.scratch.read.iops"
	storageScratchWriteBWKey          = "storage.scratch.write.bw.mibps"
	storageScratchWriteIOPSKey        = "storage.scratch.write.iops"
	storageSharedWorldTmpReadBWKey    = "storage.shared-world-tmp.read.bw.mibps"
	storageSharedWorldTmpReadIOPSKey  = "storage.shared-world-tmp.read.iops"
	storageSharedWorldTmpWriteBWKey   = "storage.shared-world-tmp.write.bw.mibps"
	storageSharedWorldTmpWriteIOPSKey = "storage.shared-world-tmp.write.iops"
	storageSharedTmpReadBWKey         = "storage.shared-tmp.read.bw.mibps"
	storageSharedTmpReadIOPSKey       = "storage.shared-tmp.read.iops"
	storageSharedTmpWriteBWKey        = "storage.shared-tmp.write.bw.mibps"
	storageSharedTmpWriteIOPSKey      = "storage.shared-tmp.write.iops"
	storageDiskWorldTmpReadBWKey      = "storage.disk-world-tmp.read.bw.mibps"
	storageDiskWorldTmpReadIOPSKey    = "storage.disk-world-tmp.read.iops"
	storageDiskWorldTmpWriteBWKey     = "storage.disk-world-tmp.write.bw.mibps"
	storageDiskWorldTmpWriteIOPSKey   = "storage.disk-world-tmp.write.iops"
	storageDiskTmpReadBWKey           = "storage.disk-tmp.read.bw.mibps"
	storageDiskTmpReadIOPSKey         = "storage.disk-tmp.read.iops"
	storageDiskTmpWriteBWKey          = "storage.disk-tmp.write.bw.mibps"
	storageDiskTmpWriteIOPSKey        = "storage.disk-tmp.write.iops"
)

func isKeyValid(key string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

var DefaultStore = NewStore()

// Store is a simple store with hard-coded keys to store benchmark results.
type Store struct {
	mu sync.RWMutex // Mutex for thread safety

	data map[string]string

	refresh chan struct{}
	failure chan error
}

func NewStore() *Store {
	data := make(map[string]string)
	for _, k := range keys {
		data[k] = ""
	}
	return &Store{
		refresh: make(chan struct{}, 1),
		failure: make(chan error, 10),
		data:    data,
	}
}

func (s *Store) SetUploadBandwidth(upload uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[networkUploadBWKey] = fmt.Sprintf("%.2f", float64(upload)/1e6)
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Uint64("upload-bandwidth", upload))
}

func (s *Store) SetDownloadBandwidth(download uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[networkDownloadBWKey] = fmt.Sprintf("%.2f", float64(download)/1e6)
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Uint64("download-bandwidth", download))
}

func (s *Store) SetGFLOPS(gflops float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[computeKey] = fmt.Sprintf("%.2f", gflops)
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("gflops", gflops))
}

func (s *Store) SetP2PBidirectionalBandwidth(bandwidth float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[networkP2PBWKey] = fmt.Sprintf("%.2f", bandwidth)
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("p2p-bibw", bandwidth))
}

func (s *Store) SetAllToAllCollectiveLatency(latency float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[networkAllToAllLatencyKey] = fmt.Sprintf("%.2f", latency)
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("alltoall-latency", latency))
}

func (s *Store) SetP2PLatency(latency float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[networkP2PLatencyKey] = fmt.Sprintf("%.2f", latency)
	select {
	case s.refresh <- struct{}{}:
	default:
	}
	logger.I.Info("stored benchmark result", zap.Float64("p2p-latency", latency))
}

func (s *Store) SetScratchResult(avgr, avgw *ior.Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[storageScratchReadBWKey] = fmt.Sprintf("%.2f", avgr.Bandwidth)
	s.data[storageScratchReadIOPSKey] = fmt.Sprintf("%.2f", avgr.IOPS)
	s.data[storageScratchWriteBWKey] = fmt.Sprintf("%.2f", avgw.Bandwidth)
	s.data[storageScratchWriteIOPSKey] = fmt.Sprintf("%.2f", avgw.IOPS)
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
	s.data[storageSharedWorldTmpReadBWKey] = fmt.Sprintf("%.2f", avgr.Bandwidth)
	s.data[storageSharedWorldTmpReadIOPSKey] = fmt.Sprintf("%.2f", avgr.IOPS)
	s.data[storageSharedWorldTmpWriteBWKey] = fmt.Sprintf("%.2f", avgw.Bandwidth)
	s.data[storageSharedWorldTmpWriteIOPSKey] = fmt.Sprintf("%.2f", avgw.IOPS)
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
	s.data[storageSharedTmpReadBWKey] = fmt.Sprintf("%.2f", avgr.Bandwidth)
	s.data[storageSharedTmpReadIOPSKey] = fmt.Sprintf("%.2f", avgr.IOPS)
	s.data[storageSharedTmpWriteBWKey] = fmt.Sprintf("%.2f", avgw.Bandwidth)
	s.data[storageSharedTmpWriteIOPSKey] = fmt.Sprintf("%.2f", avgw.IOPS)
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
	s.data[storageDiskWorldTmpReadBWKey] = fmt.Sprintf("%.2f", avgr.Bandwidth)
	s.data[storageDiskWorldTmpReadIOPSKey] = fmt.Sprintf("%.2f", avgr.IOPS)
	s.data[storageDiskWorldTmpWriteBWKey] = fmt.Sprintf("%.2f", avgw.Bandwidth)
	s.data[storageDiskWorldTmpWriteIOPSKey] = fmt.Sprintf("%.2f", avgw.IOPS)
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
	s.data[storageDiskTmpReadBWKey] = fmt.Sprintf("%.2f", avgr.Bandwidth)
	s.data[storageDiskTmpReadIOPSKey] = fmt.Sprintf("%.2f", avgr.IOPS)
	s.data[storageDiskTmpWriteBWKey] = fmt.Sprintf("%.2f", avgw.Bandwidth)
	s.data[storageDiskTmpWriteIOPSKey] = fmt.Sprintf("%.2f", avgw.IOPS)
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
	s.data[cpuKey] = spec.CPU
	s.data[cpuMicroArchKey] = spec.MicroArch
	s.data[osKey] = spec.OS
	s.data[archKey] = spec.Arch
	s.data[gpuKey] = spec.CPU
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

func cloneMap(originalMap map[string]string) map[string]string {
	clonedMap := make(map[string]string)

	for key, value := range originalMap {
		clonedMap[key] = value
	}

	return clonedMap
}

func (s *Store) Dump() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return cloneMap(s.data)
}

func (s *Store) ImportFromLabels(labels []metaschedulerabi.Label) {
	for _, l := range labels {
		if isKeyValid(l.Key) {
			s.data[l.Key] = l.Value
		}
	}
}

func (s *Store) DumpAsLabels() []metaschedulerabi.Label {
	ll := make([]metaschedulerabi.Label, 0, len(s.data))
	for k, v := range s.data {
		ll = append(ll, metaschedulerabi.Label{
			Key:   k,
			Value: v,
		})
	}
	return ll
}

func (s *Store) IsComplete() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, value := range s.data {
		if value == "" {
			return false
		}
	}
	return true
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
