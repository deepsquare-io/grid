package jobapi

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils/try"
	"go.uber.org/zap"
)

type JobHandler interface {
	SetJobStatus(
		ctx context.Context,
		jobID [32]byte,
		jobStatus eth.JobStatus,
		jobDuration uint64,
	) error
}

type jobAPIServer struct {
	supervisorv1alpha1.UnimplementedJobAPIServer
	jobHandler JobHandler
}

func New(
	jobHandler JobHandler,
) *jobAPIServer {
	if jobHandler == nil {
		logger.I.Fatal("jobHandler is nil")
	}
	return &jobAPIServer{
		jobHandler: jobHandler,
	}
}

var gRPCToEthJobStatus = map[supervisorv1alpha1.JobStatus]eth.JobStatus{
	supervisorv1alpha1.JobStatus_JOB_STATUS_PENDING:        eth.JobStatusPending,
	supervisorv1alpha1.JobStatus_JOB_STATUS_META_SCHEDULED: eth.JobStatusMetaScheduled,
	supervisorv1alpha1.JobStatus_JOB_STATUS_SCHEDULED:      eth.JobStatusScheduled,
	supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING:        eth.JobStatusRunning,
	supervisorv1alpha1.JobStatus_JOB_STATUS_CANCELLING:     eth.JobStatusCancelling,
	supervisorv1alpha1.JobStatus_JOB_STATUS_CANCELLED:      eth.JobStatusCancelled,
	supervisorv1alpha1.JobStatus_JOB_STATUS_FINISHED:       eth.JobStatusFinished,
	supervisorv1alpha1.JobStatus_JOB_STATUS_FAILED:         eth.JobStatusFailed,
	supervisorv1alpha1.JobStatus_JOB_STATUS_OUT_OF_CREDITS: eth.JobStatusOutOfCredits,
}

// SetJobStatus to the ethereum network
func (s *jobAPIServer) SetJobStatus(ctx context.Context, req *supervisorv1alpha1.SetJobStatusRequest) (*supervisorv1alpha1.SetJobStatusResponse, error) {
	logger.I.Info("grpc received job result", zap.Any("job_result", req))
	jobName, err := hex.DecodeString(req.Name)
	if err != nil {
		logger.I.Warn("SetJobStatus: DecodeString failed", zap.Error(err), zap.String("name", req.Name))
		return nil, err
	}
	var jobNameFixedLength [32]byte
	copy(jobNameFixedLength[:], jobName)

	if status, ok := gRPCToEthJobStatus[req.Status]; ok {
		if err = try.Do(func() error {
			return s.jobHandler.SetJobStatus(
				ctx,
				jobNameFixedLength,
				status,
				req.Duration/60,
			)
		}, 3, 3*time.Second); err != nil {
			logger.I.Error(
				"SetJobStatus failed",
				zap.Error(err),
				zap.String("status", req.Status.String()),
				zap.String("name", string(jobName)),
				zap.Uint64("duration", req.Duration/60),
			)
			return nil, err
		}
		return &supervisorv1alpha1.SetJobStatusResponse{}, nil
	} else {
		logger.I.Error("SetJobStatus unknown job status", zap.Error(err), zap.String("status", req.Status.String()))
		return nil, fmt.Errorf("unknown job status %s", req.Status.String())
	}

}
