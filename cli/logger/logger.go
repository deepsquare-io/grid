package logger

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"strings"
	"time"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/cli/internal/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
)

type dialer struct {
	pk       *ecdsa.PrivateKey
	endpoint string
	opts     []grpc.DialOption
}

func NewDialer(endpoint string, pk *ecdsa.PrivateKey, opts ...grpc.DialOption) types.LoggerDialer {
	return &dialer{
		pk:       pk,
		endpoint: endpoint,
		opts:     opts,
	}
}

type gridlogger struct {
	loggerv1alpha1.LoggerAPIClient
	pk *ecdsa.PrivateKey
}

func (d *dialer) DialContext(
	ctx context.Context,
) (l types.Logger, conn *grpc.ClientConn, err error) {
	conn, err = grpc.DialContext(ctx,
		d.endpoint,
		d.opts...,
	)
	if err != nil {
		return nil, conn, err
	}
	l = &gridlogger{
		LoggerAPIClient: loggerv1alpha1.NewLoggerAPIClient(conn),
		pk:              d.pk,
	}
	return l, conn, nil
}

func (l *gridlogger) from() common.Address {
	return crypto.PubkeyToAddress(l.pk.PublicKey)
}

func (l *gridlogger) WatchLogs(
	ctx context.Context,
	jobID [32]byte,
) (types.LogStream, error) {
	address := l.from().Hex()
	timestamp := uint64(time.Now().Unix())
	logName := strings.ToLower(hexutil.Encode(jobID[:]))
	data := []byte(
		fmt.Sprintf(
			"read:%s/%s/%d",
			strings.ToLower(address),
			logName,
			timestamp,
		),
	)
	hash := accounts.TextHash(data)

	signedHash, err := crypto.Sign(hash, l.pk)
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash: %v", err)
	}

	return l.Read(ctx, &loggerv1alpha1.ReadRequest{
		LogName:    logName,
		Address:    address,
		Timestamp:  timestamp,
		SignedHash: signedHash,
	})
}
