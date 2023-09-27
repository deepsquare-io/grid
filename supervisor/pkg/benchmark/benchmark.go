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

package benchmark

import (
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils/hash"
	"go.uber.org/zap"
)

//go:embed templates/*.tmpl
var templates embed.FS

const (
	GBtoMB           = 1000
	jobNameFormat    = "benchmark-%s-%s"
	DefaultTimeLimit = 24 * time.Hour
)

type Benchmark scheduler.JobDefinition

type options struct {
	nodes       uint64
	cpusPerNode uint64
	gpusPerNode uint64
	memPerNode  uint64

	phase                   string
	image                   string
	secret                  string
	supervisorPublicAddress string

	ucx          bool
	ucxAffinity  string
	ucxTransport string
	// trace enables benchmark trace logging.
	trace bool
}

type Option func(*options)

func WithClusterSpecs(
	nodes uint64,
	cpusPerNode uint64,
	gpusPerNode uint64,
	memPerNode uint64,
) Option {
	return func(o *options) {
		o.nodes = nodes
		o.cpusPerNode = cpusPerNode
		o.gpusPerNode = gpusPerNode
		o.memPerNode = memPerNode
	}
}

func WithImage(
	image string,
) Option {
	return func(o *options) {
		o.image = image
	}
}

func WithTrace() Option {
	return func(o *options) {
		o.trace = true
	}
}

func WithSupervisorPublicAddress(supervisorPublicAddress string) Option {
	return func(o *options) {
		o.supervisorPublicAddress = supervisorPublicAddress
	}
}

func WithUCX(affinity string, transport string) Option {
	return func(o *options) {
		o.ucx = true
		o.ucxTransport = transport
		o.ucxAffinity = affinity
	}
}

type Launcher interface {
	// Cancel cancels all running benchmark
	Cancel(ctx context.Context, name string) error
	// Get the generated job name for benchmarks.
	GetJobName(name string) string
	Launch(
		ctx context.Context,
		name string,
		benchmark *Benchmark,
	) error
}

type launcher struct {
	supervisorPublicAddress string
	user                    string
	scheduler               scheduler.Scheduler
	timeLimit               time.Duration
	wait                    bool
}

type LauncherOption func(*launcher)

func WithNoWait() LauncherOption {
	return func(l *launcher) {
		l.wait = false
	}
}

func WithTimeLimit(timeLimit time.Duration) LauncherOption {
	return func(l *launcher) {
		l.timeLimit = timeLimit
	}
}

func NewLauncher(
	user string,
	supervisorPublicAddress string,
	scheduler scheduler.Scheduler,
	opts ...LauncherOption,
) Launcher {
	l := &launcher{
		scheduler:               scheduler,
		user:                    user,
		supervisorPublicAddress: supervisorPublicAddress,
		timeLimit:               DefaultTimeLimit,
		wait:                    true,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *launcher) GetJobName(name string) string {
	hash := hash.GenerateAlphanumeric(l.supervisorPublicAddress)
	return fmt.Sprintf(jobNameFormat, name, hash)
}

func (l *launcher) Cancel(ctx context.Context, name string) error {
	return l.scheduler.CancelJob(ctx, l.GetJobName(name), l.user)
}

func (l *launcher) Launch(
	ctx context.Context,
	name string,
	benchmark *Benchmark,
) error {
	benchmark.Wait = l.wait
	benchmark.TimeLimit = uint64(l.timeLimit.Minutes())
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	errC := make(chan error, 1)
	go func() {
		_, err := l.scheduler.Submit(ctx, &scheduler.SubmitRequest{
			Name:          l.GetJobName(name),
			User:          l.user,
			Prefix:        "benchmark",
			JobDefinition: (*scheduler.JobDefinition)(benchmark),
		})
		errC <- err
	}()

	logger.I.Info("benchmark started", zap.String("name", name))

	for {
		select {
		case err := <-errC:
			if err != nil {
				logger.I.Error("benchmark failed", zap.String("name", name), zap.Error(err))
			} else {
				logger.I.Info("benchmark succeeded", zap.String("name", name))
			}
			return err
		case <-ticker.C:
			logger.I.Info("benchmark is still running", zap.String("name", name))
		}
	}
}
