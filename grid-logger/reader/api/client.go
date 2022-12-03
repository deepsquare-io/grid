package api

import (
	"context"
	"fmt"
	"io"

	authv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/auth/v1alpha1"
	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	conn      *grpc.ClientConn
	loggerAPI loggerv1alpha1.LoggerAPIClient
	authAPI   authv1alpha1.AuthAPIClient
}

func New(conn *grpc.ClientConn) *Client {
	if conn == nil {
		logger.I.Panic("conn is nil")
	}
	return &Client{
		conn:      conn,
		authAPI:   authv1alpha1.NewAuthAPIClient(conn),
		loggerAPI: loggerv1alpha1.NewLoggerAPIClient(conn),
	}
}

func (c *Client) SignIn(
	ctx context.Context,
	address string,
	nonce []byte,
	sig []byte,
) (string, error) {
	logger.I.Debug(
		"SignIn",
		zap.String("address", address),
		zap.String("nonce", hexutil.Encode(nonce)),
		zap.String("sig", hexutil.Encode(sig)),
	)
	resp, err := c.authAPI.SignIn(ctx, &authv1alpha1.SignInRequest{
		Address: address,
		Nonce:   nonce,
		Sig:     sig,
	})
	if err != nil {
		return "", err
	}
	return resp.GetAccessToken(), nil
}

func (c *Client) Register(
	ctx context.Context,
	address string,
) error {
	logger.I.Debug("Register", zap.String("address", address))
	_, err := c.authAPI.Register(ctx, &authv1alpha1.RegisterRequest{
		Address: address,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetNonce(
	ctx context.Context,
	address string,
) ([]byte, error) {
	logger.I.Debug("GetNonce", zap.String("address", address))
	resp, err := c.authAPI.Nonce(ctx, &authv1alpha1.NonceRequest{
		Address: address,
	})
	if err != nil {
		return []byte{}, err
	}
	return resp.GetNonce(), nil
}

func (c *Client) ReadAndWatch(
	ctx context.Context,
	logName, user, token string,
) error {
	logger.I.Debug(
		"ReadAndWatch",
		zap.String("address", logName),
		zap.String("user", user),
		zap.String("token", token),
	)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	stream, err := c.loggerAPI.Read(ctx, &loggerv1alpha1.ReadRequest{
		LogName: logName,
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
