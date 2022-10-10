package jobapi

import (
	"context"
	"encoding/hex"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

type JobHandler interface {
	FinishJob(
		ctx context.Context,
		jobID [32]byte,
		jobDuration uint64,
	) error
	FailJob(
		ctx context.Context,
		jobID [32]byte,
	) error
	StartJob(
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
	err = s.jobHandler.FinishJob(ctx, jobNameFixedLength, req.JobDuration)
	if err != nil {
		return nil, err
	}
	return &supervisorv1alpha1.SendJobResultResponse{}, nil
}

// SendJobFail to the ethereum network
func (s *jobAPIServer) SendJobFail(ctx context.Context, req *supervisorv1alpha1.SendJobFailRequest) (*supervisorv1alpha1.SendJobFailResponse, error) {
	logger.I.Info("grpc received job fail", zap.Any("job_fail", req))
	jobName, err := hex.DecodeString(req.JobName)
	if err != nil {
		return nil, err
	}
	var jobNameFixedLength [32]byte
	copy(jobNameFixedLength[:], jobName)
	err = s.jobHandler.FailJob(ctx, jobNameFixedLength)
	if err != nil {
		return nil, err
	}
	return &supervisorv1alpha1.SendJobFailResponse{}, nil
}

// SendJobStart to the ethereum network
func (s *jobAPIServer) SendJobStart(ctx context.Context, req *supervisorv1alpha1.SendJobStartRequest) (*supervisorv1alpha1.SendJobStartResponse, error) {
	logger.I.Info("grpc received job start", zap.Any("job_start", req))
	jobName, err := hex.DecodeString(req.JobName)
	if err != nil {
		return nil, err
	}
	var jobNameFixedLength [32]byte
	copy(jobNameFixedLength[:], jobName)
	err = s.jobHandler.FailJob(ctx, jobNameFixedLength)
	if err != nil {
		return nil, err
	}
	return &supervisorv1alpha1.SendJobStartResponse{}, nil
}
