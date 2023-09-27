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

package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"syscall"
	"time"

	loggerv1alpha1 "github.com/deepsquare-io/grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/grid/grid-logger/logger"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	serverEndpoint     string
	serverTLS          bool
	serverTLSInsecure  bool
	serverCAFile       string
	serverHostOverride string

	pipeFile   string
	logName    string
	userString string

	uidVerify bool
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "server.endpoint",
		Value:       "127.0.0.1:3000",
		Usage:       "Server to send logs.",
		Destination: &serverEndpoint,
		EnvVars:     []string{"SERVER_ENDPOINT"},
	},
	&cli.BoolFlag{
		Name:        "server.tls",
		Value:       true,
		Usage:       "Enable TLS for the Customer API.",
		Destination: &serverTLS,
		EnvVars:     []string{"SERVER_TLS_ENABLE"},
	},
	&cli.BoolFlag{
		Name:        "server.tls.insecure",
		Value:       false,
		Usage:       "Skip TLS verification. By enabling it, server.tls.ca and server.tls.server-host-override are ignored.",
		Destination: &serverTLSInsecure,
		EnvVars:     []string{"SERVER_TLS_INSECURE"},
	},
	&cli.StringFlag{
		Name:        "server.tls.ca",
		Value:       "",
		Usage:       "Path to CA certificate for TLS verification.",
		Destination: &serverCAFile,
		EnvVars:     []string{"SERVER_CA"},
	},
	&cli.StringFlag{
		Name:        "server.tls.server-host-override",
		Value:       "logging.deepsquare.io",
		Usage:       "The server name used to verify the hostname returned by the TLS handshake.",
		Destination: &serverHostOverride,
		EnvVars:     []string{"SERVER_HOST_OVERRIDE"},
	},

	&cli.StringFlag{
		Name:        "pipe.path",
		Value:       "pipe",
		Destination: &pipeFile,
		Usage:       "FIFO Pipe file.",
		EnvVars:     []string{"PIPE_PATH"},
	},
	&cli.StringFlag{
		Name:        "log-name",
		Usage:       "Name of the log. Used as a key in the database.",
		Required:    true,
		Destination: &logName,
		EnvVars:     []string{"LOG_NAME"},
	},
	&cli.StringFlag{
		Name:        "user",
		Usage:       "User/Owner of the log. Used for authentication.",
		Required:    true,
		Destination: &userString,
		EnvVars:     []string{"OWNER"},
	},
	&cli.BoolFlag{
		Name:        "uid-verify",
		Usage:       "Verify that the uid and the user field matches.",
		Destination: &uidVerify,
		EnvVars:     []string{"UID_VERIFY"},
	},
	&cli.BoolFlag{
		Name:    "debug",
		EnvVars: []string{"DEBUG"},
		Value:   false,
		Action: func(ctx *cli.Context, s bool) error {
			if s {
				logger.EnableDebug()
			}
			return nil
		},
	},
}

var app = &cli.App{
	Name:                 "grid-logger-writer",
	Usage:                "Send logs from pipe",
	Flags:                flags,
	Suggest:              true,
	EnableBashCompletion: true,
	Copyright: `grid-logger-writer  Copyright (C) 2023 DeepSquare Association
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.
See the GNU General Public License for more details.`,
	Action: func(cCtx *cli.Context) (err error) {
		ctx := cCtx.Context

		// Check UNIX username with user string
		if uidVerify {
			currentUser, err := user.Current()
			if err != nil {
				return err
			}
			if !strings.EqualFold(currentUser.Username, userString) {
				return errors.New("UNIX username does not match with user address")
			}
		}

		// Trap cleanup
		cleanChan := make(chan os.Signal, 1)
		signal.Notify(cleanChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		// Dial server
		opts := []grpc.DialOption{}
		if serverTLS {
			if !serverTLSInsecure {
				creds, err := credentials.NewClientTLSFromFile(serverCAFile, serverHostOverride)
				if err != nil {
					logger.I.Fatal("Failed to create TLS credentials", zap.Error(err))
				}
				opts = append(opts, grpc.WithTransportCredentials(creds))
			} else {
				tlsConfig := &tls.Config{
					InsecureSkipVerify: true,
				}
				opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
			}
		} else {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		// Open pipe
		_ = os.Remove(pipeFile)
		if err := syscall.Mkfifo(pipeFile, 0666); err != nil {
			logger.I.Error("mkfifo failed", zap.Error(err))
			return err
		}

		// Open grpc stream
		stream, conn, err := openGRPCConn(ctx, opts)
		if err != nil {
			logger.I.Error("openGRPCConn failed", zap.Error(err))
			return err
		}

		readerChan := make(chan *loggerv1alpha1.WriteRequest, 100)
		readerErrChan := make(chan error, 1)
		pipe, err := os.OpenFile(pipeFile, os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			logger.I.Error("pipe open failed", zap.Error(err))
			return err
		}

		// Pipe reading thread
		go func(ctx context.Context) {
			logger.I.Info("reading")
			reader := bufio.NewReader(pipe)
			for {
				line, err := reader.ReadBytes('\n')
				if err != nil {
					readerErrChan <- err
					logger.I.Error(
						"reader thread receive a reading error, exiting...",
						zap.Error(err),
					)
					return
				}
				req := &loggerv1alpha1.WriteRequest{
					LogName:   logName,
					Data:      line,
					User:      userString,
					Timestamp: time.Now().UnixNano(),
				}
				readerChan <- req

				if ctx.Err() != nil {
					logger.I.Warn(
						"reader thread receive a context error, exiting...",
						zap.Error(err),
					)
					return
				}
			}
		}(ctx)

		// gRPC sending thread
		for {
			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					logger.I.Error("context closed", zap.Error(err))
					return err
				}
			case data := <-readerChan:
				// Connection lost retry
				if err = stream.Send(data); err != nil {
					logger.I.Error("grpc write failed", zap.Error(err))
					_ = conn.Close()
					time.Sleep(time.Second)
					stream, conn, err = openGRPCConn(ctx, opts)
					if err != nil {
						logger.I.Error("failed to reconnect on EOF error", zap.Error(err))
						return err
					}
					_ = stream.Send(&loggerv1alpha1.WriteRequest{
						LogName: logName,
						Data: []byte(
							"Logger was disconnected temporarly, some logs may be lost.",
						),
						User:      userString,
						Timestamp: time.Now().Unix(),
					})
					logger.I.Info("successfully reconnected")
				}
			case err := <-readerErrChan:
				if err == io.EOF {
					err := stream.CloseSend()
					status, ok := status.FromError(err)
					if ok {
						if ok && status.Code() == codes.Canceled {
							logger.I.Info("gRPC channel already closing", zap.Error(err))
							return nil
						}
					}
					if err != nil && err != io.EOF {
						logger.I.Error("gRPC close failed", zap.Error(err))
						return err
					}
					logger.I.Info("pipe EOF, exiting...")
					return nil
				}
				if err != nil {
					logger.I.Error("pipe read failed", zap.Error(err))
					return err
				}
			// Graceful exit
			case <-cleanChan:
				logger.I.Info("gracefully exiting...")
				_ = stream.CloseSend()
				_ = conn.Close()
				_ = pipe.Close()
				_ = os.Remove(pipeFile)
				return nil
			}
		}
	},
}

func openGRPCConn(
	ctx context.Context,
	opts []grpc.DialOption,
) (loggerv1alpha1.LoggerAPI_WriteClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(serverEndpoint, opts...)
	if err != nil {
		logger.I.Error("grpc dial failed", zap.Error(err))
		return nil, nil, err
	}
	client := loggerv1alpha1.NewLoggerAPIClient(conn)
	stream, err := client.Write(ctx)
	if err != nil {
		logger.I.Error("grpc open failed", zap.Error(err))
		return nil, nil, err
	}
	return stream, conn, nil
}

func main() {
	_ = godotenv.Load(".writer.env", ".writer.env.local")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
