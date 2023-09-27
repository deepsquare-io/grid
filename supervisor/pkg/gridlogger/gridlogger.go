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
