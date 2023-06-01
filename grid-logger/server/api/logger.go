package api

import (
	"context"
	"errors"
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
		if err == io.EOF || errors.Is(err, context.Canceled) {
			logger.I.Info("writer closed", zap.String("IP", p.Addr.String()))
			_ = stream.SendAndClose(&loggerv1alpha1.WriteResponse{})
			return nil
		}
		if err != nil {
			logger.I.Error(
				"writer closed with error",
				zap.String("IP", p.Addr.String()),
				zap.Error(err),
			)
			return err
		}

		_, err = s.db.Append(req)
		if err != nil {
			logger.I.Error(
				"writer failed to write",
				zap.Any("req", req),
				zap.String("IP", p.Addr.String()),
				zap.Error(err),
			)
			_ = stream.SendAndClose(&loggerv1alpha1.WriteResponse{})
			return err
		}
	}
}

func (s *loggerAPIServer) Read(
	req *loggerv1alpha1.ReadRequest,
	stream loggerv1alpha1.LoggerAPI_ReadServer,
) error {
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

	logs := make(chan *loggerv1alpha1.ReadResponse, 100)
	go func() {
		defer close(logs)
		if err := s.db.ReadAndWatch(ctx, address, req.GetLogName(), logs); err != nil {
			logger.I.Error("read and watch failed", zap.Error(err))
		}
	}()
	for {
		select {
		case <-ctx.Done():
			logger.I.Info("reader closed", zap.String("user", address), zap.Error(ctx.Err()))
			return nil
		case log, ok := <-logs:
			if !ok {
				logger.I.Error(
					"logs closed (read and watch the database might have closed)",
					zap.String("user", address),
					zap.Error(ctx.Err()),
				)
				return nil
			}
			if err := stream.Send(log); err != nil {
				return err
			}
		}
	}
}

func (s *loggerAPIServer) WatchList(
	req *loggerv1alpha1.WatchListRequest,
	stream loggerv1alpha1.LoggerAPI_WatchListServer,
) error {
	ctx := stream.Context()
	address := strings.ToLower(req.GetAddress())

	if err := auth.Verify(
		address,
		[]byte(fmt.Sprintf("watchList:%s/%d", address, req.GetTimestamp())),
		req.GetSignedHash(),
	); err != nil {
		return status.Errorf(codes.Unauthenticated, "failed to authenticate: %v", err)
	}

	lists := make(chan []string, 100)
	defer close(lists)
	go func() {
		if err := s.db.ListAndWatch(ctx, address, lists); err != nil {
			logger.I.Error("list and watch failed", zap.Error(err))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case list, ok := <-lists:
			if !ok {
				logger.I.Error(
					"logs closed (read and watch the database might have closed)",
					zap.String("user", address),
					zap.Error(ctx.Err()),
				)
				return nil
			}
			if err := stream.Send(&loggerv1alpha1.WatchListResponse{
				LogNames: list,
			}); err != nil {
				return err
			}
		}
	}
}
