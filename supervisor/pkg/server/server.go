package server

import (
	"net"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/grpc/health/v1"
	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/health"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/jobapi"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/sshapi"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	grpc *grpc.Server
}

func New(
	tls bool,
	keyFile string,
	certFile string,
	jobHandler jobapi.JobHandler,
	slurm *slurm.Service,
	pkB64 string,
) *Server {
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger.I),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger.I),
		)),
	}
	if tls {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logger.I.Fatal("failed to load certificates", zap.Error(err))
		}
		opts = append(opts, grpc.Creds(creds))
	}

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
		health.New(slurm),
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
