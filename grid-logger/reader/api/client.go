package api

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"strings"
	"time"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Client struct {
	conn      *grpc.ClientConn
	pk        *ecdsa.PrivateKey
	loggerAPI loggerv1alpha1.LoggerAPIClient
}

func New(conn *grpc.ClientConn, pk *ecdsa.PrivateKey) *Client {
	if conn == nil {
		logger.I.Panic("conn is nil")
	}
	return &Client{
		conn:      conn,
		loggerAPI: loggerv1alpha1.NewLoggerAPIClient(conn),
		pk:        pk,
	}
}

func (c *Client) ReadAndWatch(
	ctx context.Context,
	address string,
	logName string,
) error {
	logger.I.Debug(
		"ReadAndWatch",
		zap.String("address", address),
		zap.String("logName", logName),
	)
	timestamp := uint64(time.Now().Unix())
	data := []byte(fmt.Sprintf("read:%s/%s/%d", strings.ToLower(address), logName, timestamp))
	hash := accounts.TextHash(data)

	signedHash, err := crypto.Sign(hash, c.pk)
	if err != nil {
		return err
	}
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
		fmt.Printf("%d ", log.Timestamp)
		fmt.Println(string(log.Data))
	}
	return nil
}

func (c *Client) List(
	ctx context.Context,
	address string,
) error {
	logger.I.Debug(
		"List",
		zap.String("address", address),
	)
	timestamp := uint64(time.Now().Unix())
	data := []byte(fmt.Sprintf("watchList:%s/%d", strings.ToLower(address), timestamp))
	hash := accounts.TextHash(data)

	signedHash, err := crypto.Sign(hash, c.pk)
	if err != nil {
		return err
	}
	stream, err := c.loggerAPI.WatchList(ctx, &loggerv1alpha1.WatchListRequest{
		Address:    address,
		Timestamp:  timestamp,
		SignedHash: signedHash,
	})
	if err != nil {
		return err
	}

	resp, err := stream.Recv()
	if err != nil {
		return err
	}

	for _, name := range resp.LogNames {
		fmt.Println(name)
	}

	return nil
}
