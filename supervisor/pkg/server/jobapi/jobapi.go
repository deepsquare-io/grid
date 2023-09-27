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

package jobapi

import (
	"context"
	"errors"
	"fmt"
	"time"

	supervisorv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils/try"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
)

type Server struct {
	supervisorv1alpha1.UnimplementedJobAPIServer
	ms      metascheduler.MetaScheduler
	Timeout time.Duration
	// Delay between tries
	Delay           time.Duration
	resourceManager *lock.ResourceManager
}

func New(
	ms metascheduler.MetaScheduler,
	resourceManager *lock.ResourceManager,
) *Server {
	if ms == nil {
		logger.I.Fatal("ms is nil")
	}
	if resourceManager == nil {
		logger.I.Fatal("resourceManager is nil")
	}
	return &Server{
		ms:              ms,
		Timeout:         15 * time.Second,
		Delay:           3 * time.Second,
		resourceManager: resourceManager,
	}
}

var gRPCToEthJobStatus = map[supervisorv1alpha1.JobStatus]metascheduler.JobStatus{
	supervisorv1alpha1.JobStatus_JOB_STATUS_PENDING:        metascheduler.JobStatusPending,
	supervisorv1alpha1.JobStatus_JOB_STATUS_META_SCHEDULED: metascheduler.JobStatusMetaScheduled,
	supervisorv1alpha1.JobStatus_JOB_STATUS_SCHEDULED:      metascheduler.JobStatusScheduled,
	supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING:        metascheduler.JobStatusRunning,
	supervisorv1alpha1.JobStatus_JOB_STATUS_CANCELLING:     metascheduler.JobStatusUnknown,
	supervisorv1alpha1.JobStatus_JOB_STATUS_CANCELLED:      metascheduler.JobStatusCancelled,
	supervisorv1alpha1.JobStatus_JOB_STATUS_FINISHED:       metascheduler.JobStatusFinished,
	supervisorv1alpha1.JobStatus_JOB_STATUS_FAILED:         metascheduler.JobStatusFailed,
	supervisorv1alpha1.JobStatus_JOB_STATUS_OUT_OF_CREDITS: metascheduler.JobStatusOutOfCredits,
}

// SetJobStatus to the ethereum network
func (s *Server) SetJobStatus(
	ctx context.Context,
	req *supervisorv1alpha1.SetJobStatusRequest,
) (*supervisorv1alpha1.SetJobStatusResponse, error) {
	logger.I.Info("grpc received job result", zap.Any("job_result", req))
	go func() {
		logger.I.Info("launched setJobStatusTask", zap.Any("job_result", req))
		if err := s.setJobStatusTask(context.Background(), req); err != nil {
			logger.I.Error("setJobStatusTask failed", zap.Error(err))
		}
	}()
	return &supervisorv1alpha1.SetJobStatusResponse{}, nil
}

func (s *Server) setJobStatusTask(
	ctx context.Context,
	req *supervisorv1alpha1.SetJobStatusRequest,
) error {

	jobName, err := hexutil.Decode(req.Name)
	if err != nil {
		logger.I.Warn(
			"SetJobStatus: DecodeString failed",
			zap.Error(err),
			zap.String("name", req.Name),
		)
		return err
	}
	var jobNameFixedLength [32]byte
	copy(jobNameFixedLength[:], jobName)

	// Lock the job: avoid any mutation of the job until a setjob is perfectly sent
	s.resourceManager.Lock(req.Name)
	defer s.resourceManager.Unlock(req.Name)

	if status, ok := gRPCToEthJobStatus[req.Status]; ok {
		// Ignore unknown status transition, this is for backward compatilibility
		if status == metascheduler.JobStatusUnknown {
			logger.I.Warn(
				"status unknown (if the status is deprecated, ignore this warning)",
				zap.Error(err),
				zap.String("status", req.Status.String()),
				zap.String("name", req.Name),
				zap.Uint64("duration", req.Duration/60),
			)
			return nil
		}

		// Do set job status
		if err = try.DoWithContextTimeout(
			ctx,
			3, s.Delay, s.Timeout,
			func(ctx context.Context, _ int) error {
				err := s.ms.SetJobStatus(
					ctx,
					jobNameFixedLength,
					status,
					req.Duration/60,
				)
				if err != nil {
					if errors.Is(err, &metascheduler.SameStatusError{}) {
						logger.I.Warn(
							"Cannot change status to itself",
							zap.Error(err),
							zap.String("status", req.Status.String()),
							zap.String("name", req.Name),
							zap.Uint64("duration", req.Duration/60),
						)
						return nil
					}
					var customErr *metascheduler.InvalidTransition
					if ok := errors.As(err, &customErr); ok && customErr.From == metascheduler.JobStatusScheduled {
						logger.I.Warn(
							"Invalid state transition from SCHEDULED.",
							zap.Error(err),
							zap.String("status", req.Status.String()),
							zap.String("name", req.Name),
							zap.Uint64("duration", req.Duration/60),
						)
						if err := s.ms.SetJobStatus(
							ctx,
							jobNameFixedLength,
							metascheduler.JobStatusRunning,
							req.Duration/60,
						); err != nil {
							logger.I.Error(
								"Failed to put the job in RUNNING",
								zap.Error(err),
								zap.String("status", req.Status.String()),
								zap.String("name", req.Name),
								zap.Uint64("duration", req.Duration/60),
							)
							return err
						}
					}
				}

				return err

			}); err != nil {
			logger.I.Error(
				"SetJobStatus failed",
				zap.Error(err),
				zap.String("status", req.Status.String()),
				zap.String("name", req.Name),
				zap.Uint64("duration", req.Duration/60),
			)
			return err
		}
		return nil
	} else {
		logger.I.Error("SetJobStatus unknown job status", zap.Error(err), zap.String("status", req.Status.String()))
		return fmt.Errorf("unknown job status %s", req.Status.String())
	}
}
