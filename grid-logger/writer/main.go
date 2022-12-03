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
	"github.com/joho/godotenv"
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
	logName  string
	user     string
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
		Destination: &user,
		EnvVars:     []string{"OWNER"},
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
	Name:    "grid-logger-writer",
	Usage:   "Send logs from pipe",
	Flags:   flags,
	Suggest: true,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context

		// Trap cleanup
		c := make(chan os.Signal, 1)
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

		// Open pipe
		_ = os.Remove(pipeFile)
		if err := syscall.Mkfifo(pipeFile, 0666); err != nil {
			logger.I.Error("mkfifo failed", zap.Error(err))
			return err
		}

		pipe, err := os.OpenFile(pipeFile, os.O_RDWR, os.ModeNamedPipe)
		if err != nil {
			logger.I.Error("pipe open failed", zap.Error(err))
			return err
		}

		// Open grpc stream
		stream, err := client.Write(ctx)
		if err != nil {
			logger.I.Fatal("grpc open failed", zap.Error(err))
		}

		logger.I.Info("reading")

		reader := bufio.NewReader(pipe)
		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				_, err := stream.CloseAndRecv()
				if err != nil {
					logger.I.Error("grpc close failed", zap.Error(err))
					return err
				}
				logger.I.Info("pipe EOF, exiting...")
				return nil
			}
			if err != nil {
				logger.I.Error("pipe read failed", zap.Error(err))
				return err
			}
			logger.I.Debug(
				"pipe recv",
				zap.String("data", string(line)),
			)
			if err := stream.Send(&loggerv1alpha1.WriteRequest{
				LogName: logName,
				Data:    line,
				User:    user,
			}); err != nil {
				logger.I.Error("grpc write failed", zap.Error(err))
			}
		}
	},
}

func main() {
	_ = godotenv.Load(".writer.env", ".writer.env.local")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
