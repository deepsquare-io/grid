package api

import (
	"fmt"
	"io"
	"strings"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/deepsquare-io/the-grid/grid-logger/server/auth"
	"github.com/deepsquare-io/the-grid/grid-logger/server/db"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
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
	p, ok := peer.FromContext(ctx)
	if ok {
		logger.I.Info("writer connected", zap.String("IP", p.Addr.String()))
	} else {
		logger.I.Info("writer connected but was not identified")
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logger.I.Info("writer closed", zap.String("IP", p.Addr.String()))
			return stream.SendAndClose(&loggerv1alpha1.WriteResponse{})
		}
		if err != nil {
			return err
		}

		n, err := s.db.Append(req.GetLogName(), strings.ToLower(req.GetUser()), req.GetData())
		if err != nil {
			return err
		}
		logger.I.Debug("write", zap.Int("size", n), zap.Any("req", req))
	}
}

func (s *loggerAPIServer) Read(req *loggerv1alpha1.ReadRequest, stream loggerv1alpha1.LoggerAPI_ReadServer) error {
	ctx := stream.Context()
	if p, ok := peer.FromContext(ctx); ok {
		logger.I.Info("reader connected", zap.String("IP", p.Addr.String()), zap.Any("req", req))
	} else {
		logger.I.Info("reader connected but was not identified", zap.Any("req", req))
	}
	address := strings.ToLower(req.GetAddress())

	if err := auth.Verify(
		address,
		[]byte(fmt.Sprintf("read:%s/%s/%d", address, req.GetLogName(), req.GetTimestamp())),
		req.GetSignedHash(),
	); err != nil {
		return status.Errorf(codes.Unauthenticated, "failed to authenticate: %v", err)
	}
	logger.I.Info("reader authenticated", zap.String("user", address), zap.Any("req", req))

	logs := make(chan string)
	go func() {
		if err := s.db.ReadAndWatch(ctx, address, req.GetLogName(), logs); err != nil {
			logger.I.Error("read and watch failed", zap.Error(err))
		}
	}()
	for {
		select {
		case <-ctx.Done():
			logger.I.Info("reader closed", zap.String("user", address))
			return nil
		case log := <-logs:
			bytes := []byte(log)
			logger.I.Debug("reader send", zap.Int("size", len(bytes)))
			if err := stream.Send(&loggerv1alpha1.ReadResponse{
				Data: bytes,
			}); err != nil {
				return err
			}
		}
	}
}

func (s *loggerAPIServer) WatchList(req *loggerv1alpha1.WatchListRequest, stream loggerv1alpha1.LoggerAPI_WatchListServer) error {
	ctx := stream.Context()
	address := strings.ToLower(req.GetAddress())

	if err := auth.Verify(
		address,
		[]byte(fmt.Sprintf("watchList:%s/%d", address, req.GetTimestamp())),
		req.GetSignedHash(),
	); err != nil {
		return status.Errorf(codes.Unauthenticated, "failed to authenticate: %v", err)
	}

	lists := make(chan []string)
	go func() {
		if err := s.db.ListAndWatch(ctx, address, lists); err != nil {
			logger.I.Error("list and watch failed", zap.Error(err))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case list := <-lists:
			if err := stream.Send(&loggerv1alpha1.WatchListResponse{
				LogNames: list,
			}); err != nil {
				return err
			}
		}
	}
}
