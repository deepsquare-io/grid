package deepsquare

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	loggerv1alpha1 "github.com/deepsquare-io/the-grid/cli/deepsquare/generated/logger/v1alpha1"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
)

type logger struct {
	loggerv1alpha1.LoggerAPIClient
	pk *ecdsa.PrivateKey
}

type DialOption *grpc.DialOption

func DialContext(
	ctx context.Context,
	target string,
	pk *ecdsa.PrivateKey,
	opts ...grpc.DialOption,
) (l deepsquare.Logger, conn grpc.ClientConnInterface, err error) {
	conn, err = grpc.DialContext(ctx,
		target,
		opts...,
	)
	if err != nil {
		return nil, conn, err
	}
	l = &logger{
		LoggerAPIClient: loggerv1alpha1.NewLoggerAPIClient(conn),
		pk:              pk,
	}
	return l, conn, nil
}

func (l *logger) from() common.Address {
	return crypto.PubkeyToAddress(l.pk.PublicKey)
}

func (l *logger) WatchLogs(
	ctx context.Context,
	jobName string,
) (deepsquare.LogStream, error) {
	address := l.from().Hex()
	timestamp := uint64(time.Now().Unix())
	data := []byte(
		fmt.Sprintf("read:%s/%s/%d", strings.ToLower(address), jobName, timestamp),
	)
	hash := accounts.TextHash(data)

	signedHash, err := crypto.Sign(hash, l.pk)
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash: %v", err)
	}

	return l.Read(ctx, &loggerv1alpha1.ReadRequest{
		LogName:    jobName,
		Address:    address,
		Timestamp:  timestamp,
		SignedHash: signedHash,
	})
}
