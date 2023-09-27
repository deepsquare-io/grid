// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	loggerv1alpha1 "github.com/deepsquare-io/grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/grid/grid-logger/logger"
	"github.com/deepsquare-io/grid/grid-logger/server/auth"
	"github.com/deepsquare-io/grid/grid-logger/server/db"
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
	log := logger.I
	p, ok := peer.FromContext(ctx)
	if ok {
		log = log.With(zap.String("IP", p.Addr.String()))
		log.Info("writer connected")
	} else {
		log.Info("writer connected but was not identified")
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF || errors.Is(err, context.Canceled) {
			log.Info("writer closed")
			_ = stream.SendAndClose(&loggerv1alpha1.WriteResponse{})
			return nil
		}
		if err != nil {
			log.Error(
				"writer closed with error",
				zap.Error(err),
			)
			return err
		}

		_, err = s.db.Append(req)
		if err != nil {
			log.Error(
				"writer failed to write",
				zap.Any("req", req),
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
	log := logger.I.With(zap.Any("req", req))
	p, ok := peer.FromContext(ctx)
	if ok {
		log = log.With(zap.String("IP", p.Addr.String()))
		log.Info("reader connected")
	} else {
		log.Info("reader connected but was not identified")
	}
	address := strings.ToLower(req.GetAddress())

	if err := auth.Verify(
		address,
		[]byte(fmt.Sprintf("read:%s/%s/%d", address, req.GetLogName(), req.GetTimestamp())),
		req.GetSignedHash(),
	); err != nil {
		return status.Errorf(codes.Unauthenticated, "failed to authenticate: %v", err)
	}
	log = log.With(zap.String("user", address))
	log.Info("reader authenticated")

	logs := make(chan *loggerv1alpha1.ReadResponse, 100)
	go func() {
		defer close(logs)
		if err := s.db.ReadAndWatch(ctx, address, req.GetLogName(), logs); err != nil {
			log.Error("read and watch failed", zap.Error(err))
		}
	}()
	for {
		select {
		case <-ctx.Done():
			log.Info("reader closed", zap.Error(ctx.Err()))
			return nil
		case l, ok := <-logs:
			if !ok {
				log.Error(
					"logs closed (read and watch the database might have closed)",
				)
				return nil
			}
			if err := stream.Send(l); err != nil {
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
