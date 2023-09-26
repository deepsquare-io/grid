package gridlogger

import (
	"context"

	loggerv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/logger/v1alpha1"
	"google.golang.org/grpc"
)

type Dialer interface {
	DialContext(
		ctx context.Context,
		endpoint string,
	) (c loggerv1alpha1.LoggerAPI_WriteClient, close func() error, err error)
}

func NewDialer(opts ...grpc.DialOption) Dialer {
	return &dialer{
		Options: opts,
	}
}

type dialer struct {
	Options []grpc.DialOption
}

func (d *dialer) DialContext(
	ctx context.Context,
	endpoint string,
) (c loggerv1alpha1.LoggerAPI_WriteClient, close func() error, err error) {
	conn, err := grpc.Dial(endpoint, d.Options...)
	if err != nil {
		return nil, nil, err
	}
	client := loggerv1alpha1.NewLoggerAPIClient(conn)
	stream, err := client.Write(ctx)
	if err != nil {
		return nil, nil, err
	}
	return stream, conn.Close, nil
}
