package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverEndpoint     string
	serverTLS          bool
	serverTLSInsecure  bool
	serverCAFile       string
	serverHostOverride string

	logName string
	user    string
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
		Name:        "log-name",
		Usage:       "Name of the log. Used as a key in the database.",
		Required:    true,
		Destination: &logName,
	},
	&cli.StringFlag{
		Name:        "user",
		Usage:       "User/Owner of the log. Used for authentication.",
		Required:    true,
		Destination: &user,
	},
}

var app = &cli.App{
	Name:    "grid-logger-reader",
	Usage:   "Read logs",
	Flags:   flags,
	Suggest: true,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context

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
		conn, err := grpc.Dial(serverEndpoint, opts...)
		if err != nil {
			logger.I.Error("grpc dial failed", zap.Error(err))
			return err
		}
		defer func() {
			if err := conn.Close(); err != nil {
				if err != io.EOF {
					logger.I.Error("grpc close failed", zap.Error(err))
				}
			}
		}()
		client := loggerv1alpha1.NewLoggerAPIClient(conn)

		stream, err := client.Read(ctx, &loggerv1alpha1.ReadRequest{
			LogName: logName,
			User:    user,
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
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
