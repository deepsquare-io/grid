// Copyright (C) 2023 DeepSquare Asociation
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

// Package logger providers an implementation of the Grid Logger reader.
package logger

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"strings"
	"time"

	loggerv1alpha1 "github.com/deepsquare-io/grid/cli/internal/logger/v1alpha1"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
)

type gridlogger struct {
	loggerv1alpha1.LoggerAPIClient
	pk *ecdsa.PrivateKey
}

func DialContext(
	ctx context.Context,
	endpoint string,
	privateKey *ecdsa.PrivateKey,
	opts ...grpc.DialOption,
) (l types.Logger, conn *grpc.ClientConn, err error) {
	conn, err = grpc.DialContext(ctx,
		endpoint,
		opts...,
	)
	if err != nil {
		return nil, conn, err
	}
	l = &gridlogger{
		LoggerAPIClient: loggerv1alpha1.NewLoggerAPIClient(conn),
		pk:              privateKey,
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
