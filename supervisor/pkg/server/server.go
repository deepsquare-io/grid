package server

import (
	"net"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/jobapi"
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
	jobFinisher jobapi.JobFinisher,
) *Server {
	var opts []grpc.ServerOption
	if tls {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logger.I.Fatal("failed to load certificates", zap.Error(err))
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	supervisorv1alpha1.RegisterJobAPIServer(grpcServer, jobapi.New(jobFinisher))

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
