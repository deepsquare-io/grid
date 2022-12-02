package api

import (
	"io"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/deepsquare-io/the-grid/grid-logger/server/db"
	"go.uber.org/zap"
	"google.golang.org/grpc/peer"
)

type loggerAPIServer struct {
	loggerv1alpha1.UnimplementedLoggerAPIServer
	db *db.File
}

func NewLoggerAPIServer(db *db.File) *loggerAPIServer {
	if db == nil {
		logger.I.Panic("db is nil")
	}
	return &loggerAPIServer{
		db: db,
	}
}

func (s *loggerAPIServer) Write(stream loggerv1alpha1.LoggerAPI_WriteServer) error {
	ctx := stream.Context()
	if p, ok := peer.FromContext(ctx); ok {
		logger.I.Info("writer connected", zap.String("IP", p.Addr.String()))
	} else {
		logger.I.Info("writer connected but was not identified")
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&loggerv1alpha1.WriteResponse{})
		}
		if err != nil {
			return err
		}

		n, err := s.db.Append(req.LogName, req.User, req.Data)
		if err != nil {
			return err
		}
		logger.I.Info("write", zap.Int("size", n), zap.String("log-name", req.LogName))
	}
}

func (s *loggerAPIServer) Read(req *loggerv1alpha1.ReadRequest, stream loggerv1alpha1.LoggerAPI_ReadServer) error {
	ctx := stream.Context()
	if p, ok := peer.FromContext(ctx); ok {
		logger.I.Info("reader connected", zap.String("IP", p.Addr.String()))
	} else {
		logger.I.Info("reader connected but was not identified")
	}

	logs := make(chan string)
	go func() {
		if err := s.db.ReadAndWatch(req.LogName, req.User, logs); err != nil {
			logger.I.Error("read and watch failed", zap.Error(err))
		}
	}()
	for log := range logs {
		if err := stream.Send(&loggerv1alpha1.ReadResponse{
			Data: []byte(log),
		}); err != nil {
			return err
		}
	}
	return nil
}
