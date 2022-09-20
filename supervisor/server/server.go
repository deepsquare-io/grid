package server

import (
	"context"

	supervisorv1alpha1 "github.com/deepsquare/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare/supervisor/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/peer"
)

type RouteGuideServer struct {
	supervisorv1alpha1.UnimplementedRouteGuideAPIServer
}

func New() *RouteGuideServer {
	return &RouteGuideServer{}
}

func (s *RouteGuideServer) GetFeature(ctx context.Context, req *supervisorv1alpha1.GetFeatureRequest) (*supervisorv1alpha1.GetFeatureResponse, error) {
	p, _ := peer.FromContext(ctx)
	logger.I.Info(
		"recv",
		zap.String("body", req.String()),
		zap.String("address", p.Addr.String()),
	)

	return &supervisorv1alpha1.GetFeatureResponse{}, nil
}
