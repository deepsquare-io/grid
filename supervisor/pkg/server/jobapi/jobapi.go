package jobapi

import (
	"context"
	"encoding/hex"
	"math/big"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

type JobHandler interface {
	FinishJob(
		ctx context.Context,
		jobID [32]byte,
		jobDuration *big.Int,
	) error
	FailedJob(
		ctx context.Context,
		jobID [32]byte,
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

// SendJobResult to the ethereum network
func (s *jobAPIServer) SendJobResult(ctx context.Context, req *supervisorv1alpha1.SendJobResultRequest) (*supervisorv1alpha1.SendJobResultResponse, error) {
	logger.I.Info("grpc received job result", zap.Any("job_result", req))
	jobName, err := hex.DecodeString(req.JobName)
	if err != nil {
		return nil, err
	}
	var jobNameFixedLength [32]byte
	copy(jobNameFixedLength[:], jobName)
	err = s.jobHandler.FinishJob(ctx, jobNameFixedLength, new(big.Int).SetUint64(req.JobDuration))
	if err != nil {
		return nil, err
	}
	return &supervisorv1alpha1.SendJobResultResponse{}, nil
}

// SendJobFailed to the ethereum network
func (s *jobAPIServer) SendJobFailed(ctx context.Context, req *supervisorv1alpha1.SendJobFailedRequest) (*supervisorv1alpha1.SendJobFailedResponse, error) {
	logger.I.Info("grpc received job failed", zap.Any("job_failed", req))
	jobName, err := hex.DecodeString(req.JobName)
	if err != nil {
		return nil, err
	}
	var jobNameFixedLength [32]byte
	copy(jobNameFixedLength[:], jobName)
	err = s.jobHandler.FailedJob(ctx, jobNameFixedLength)
	if err != nil {
		return nil, err
	}
	return &supervisorv1alpha1.SendJobFailedResponse{}, nil
}
