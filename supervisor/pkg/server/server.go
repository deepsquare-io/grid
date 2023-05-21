package server

import (
	"net"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/generated/grpc/health/v1"
	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/health"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/jobapi"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/sshapi"
	"google.golang.org/grpc"
)

type Server struct {
	grpc *grpc.Server
}

func New(
	jobHandler jobapi.JobHandler,
	pkB64 string,
	opts ...grpc.ServerOption,
) *Server {
	grpcServer := grpc.NewServer(opts...)
	supervisorv1alpha1.RegisterJobAPIServer(
		grpcServer,
		jobapi.New(jobHandler),
	)
	supervisorv1alpha1.RegisterSshAPIServer(
		grpcServer,
		sshapi.New(pkB64),
	)
	healthv1.RegisterHealthServer(
		grpcServer,
		health.New(),
	)

	return &Server{
		grpc: grpcServer,
	}
}

func (s *Server) ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return s.grpc.Serve(lis)
}
