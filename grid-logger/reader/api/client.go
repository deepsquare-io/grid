package api

import (
	"context"
	"fmt"
	"io"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Client struct {
	conn      *grpc.ClientConn
	loggerAPI loggerv1alpha1.LoggerAPIClient
}

func New(conn *grpc.ClientConn) *Client {
	if conn == nil {
		logger.I.Panic("conn is nil")
	}
	return &Client{
		conn:      conn,
		loggerAPI: loggerv1alpha1.NewLoggerAPIClient(conn),
	}
}

func (c *Client) ReadAndWatch(
	ctx context.Context,
	address string,
	logName string,
	timestamp uint64,
	signedHash []byte,
) error {
	logger.I.Debug(
		"ReadAndWatch",
		zap.String("address", address),
		zap.String("logName", logName),
		zap.String("signedHash", hexutil.Encode(signedHash)),
	)
	stream, err := c.loggerAPI.Read(ctx, &loggerv1alpha1.ReadRequest{
		LogName:    logName,
		Address:    address,
		Timestamp:  timestamp,
		SignedHash: signedHash,
	})
	if err != nil {
		return err
	}

	for {
		log, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(string(log.Data))
	}
	return nil
}
