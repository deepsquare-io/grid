// go:build client

package main

import (
	"bufio"
	"crypto/tls"
	"io"
	"os"
	"os/signal"
	"syscall"

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

	pipeFile string
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
		Name:        "pipe-file",
		Value:       "pipe",
		Destination: &pipeFile,
		Usage:       "FIFO Pipe file.",
	},
}

var app = &cli.App{
	Name:    "grid-logger-client",
	Usage:   "Send log from pipe",
	Flags:   flags,
	Suggest: true,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context

		// Trap cleanup
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-c
			logger.I.Info("cleaning up")
			_ = os.Remove(pipeFile)
			os.Exit(0)
		}()

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
			logger.I.Error("grpc dial failed:", err)
			return err
		}
		defer func() {
			if err := conn.Close(); err != nil {
				if err != io.EOF {
					logger.I.Error("grpc close failed:", err)
				}
			}
		}()
		client := loggerv1alpha1.NewLoggerAPIClient(conn)

		// Open pipe
		_ = os.Remove(pipeFile)
		if err := syscall.Mkfifo(pipeFile, 0666); err != nil {
			logger.I.Error("mkfifo failed:", err)
			return err
		}

		pipe, err := os.OpenFile(pipeFile, os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			logger.I.Error("pipe open failed:", err)
			return err
		}

		// Open grpc stream
		stream, err := client.Write(ctx)
		if err != nil {
			logger.I.Fatal("grpc open failed:", err)
		}

		logger.I.Info("reading")

		reader := bufio.NewReader(pipe)
		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				_, err := stream.CloseAndRecv()
				if err != nil {
					logger.I.Error("grpc close failed:", err)
					return err
				}
				logger.I.Info("pipe EOF, exiting...")
				return nil
			}
			if err != nil {
				logger.I.Error("pipe read failed:", err)
				return err
			}
			logger.I.Debug("recv", line)
			if err := stream.Send(&loggerv1alpha1.WriteRequest{
				Data: line,
			}); err != nil {
				logger.I.Error("grpc write failed:", err)
			}
		}
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
