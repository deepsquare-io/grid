package jobapi

import (
	"context"
	"encoding/hex"
	"math/big"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

type JobFinisher interface {
	FinishJob(
		ctx context.Context,
		jobID [32]byte,
		jobDuration *big.Int,
	) error
}

type jobAPIServer struct {
	supervisorv1alpha1.UnimplementedJobAPIServer
	finisher JobFinisher
}

func New(
	finisher JobFinisher,
) *jobAPIServer {
	if finisher == nil {
		logger.I.Fatal("finisher is nil")
	}
	return &jobAPIServer{
		finisher: finisher,
	}
}

// SendJobResult to the ethereum network
func (s *jobAPIServer) SendJobResult(ctx context.Context, req *supervisorv1alpha1.SendJobResultRequest) (*supervisorv1alpha1.SendJobResultResponse, error) {
	logger.I.Info("Received job result", zap.Any("job_result", req))
	res := req.GetJobResult()
	jobName, err := hex.DecodeString(res.JobName)
	if err != nil {
		return nil, err
	}
	var jobNameFixedLength [32]byte
	copy(jobNameFixedLength[:], jobName)
	s.finisher.FinishJob(ctx, jobNameFixedLength, new(big.Int).SetUint64(res.JobDuration))
	return &supervisorv1alpha1.SendJobResultResponse{}, nil
}
