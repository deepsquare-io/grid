package jobapi

import (
	"context"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

type jobAPIServer struct {
	supervisorv1alpha1.UnimplementedJobAPIServer
}

func New() *jobAPIServer {
	return &jobAPIServer{}
}

func (s *jobAPIServer) SendJobResult(ctx context.Context, req *supervisorv1alpha1.SendJobResultRequest) (*supervisorv1alpha1.SendJobResultResponse, error) {
	logger.I.Info("Received job result", zap.Any("job_result", req))
	// TODO: send job result to ethereum
	return &supervisorv1alpha1.SendJobResultResponse{}, nil
}
