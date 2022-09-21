package server

import (
	"net"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/server/jobapi"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func listenAndServe(addr string, opts []grpc.ServerOption) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(opts...)
	supervisorv1alpha1.RegisterJobAPIServer(grpcServer, jobapi.New())

	return grpcServer.Serve(lis)
}

func ListenAndServe(addr string) error {
	opts := []grpc.ServerOption{}
	return listenAndServe(addr, opts)
}

func ListenAndServeTLS(addr string, keyFile string, certFile string) error {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		logger.I.Fatal("failed to load certificates", zap.Error(err))
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	return listenAndServe(addr, opts)
}
