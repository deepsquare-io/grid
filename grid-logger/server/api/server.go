// go:build server

package api

import (
	"fmt"
	"io"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
)

type LoggerAPIServer struct {
	loggerv1alpha1.UnimplementedLoggerAPIServer
}

func (s *LoggerAPIServer) Write(stream loggerv1alpha1.LoggerAPI_WriteServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&loggerv1alpha1.WriteResponse{})
		}
		if err != nil {
			return err
		}
		fmt.Print(string(req.Data))
	}
}
