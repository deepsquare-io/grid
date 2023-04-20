package main

import (
	"net"
	"os"

	healthv1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/grpc/health/v1"
	loggerv1alpha1 "github.com/deepsquare-io/the-grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/deepsquare-io/the-grid/grid-logger/server/api"
	"github.com/deepsquare-io/the-grid/grid-logger/server/api/health"
	"github.com/deepsquare-io/the-grid/grid-logger/server/db"
	"github.com/deepsquare-io/the-grid/grid-logger/server/debug"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	listenAddress string

	tls      bool
	keyFile  string
	certFile string

	storagePath string

	secret []byte
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "grpc.listen-address",
		Value:       ":3000",
		Usage:       "Address to listen on. Is used for receiving job status via the job completion plugin.",
		Destination: &listenAddress,
		EnvVars:     []string{"LISTEN_ADDRESS"},
	},
	&cli.BoolFlag{
		Name:        "tls",
		Value:       false,
		Destination: &tls,
		Usage:       "Enable TLS for GRPC.",
		EnvVars:     []string{"TLS_ENABLE"},
	},
	&cli.StringFlag{
		Name:        "tls.key-file",
		Value:       "",
		Destination: &keyFile,
		Usage:       "TLS Private Key file.",
		EnvVars:     []string{"TLS_KEY"},
	},
	&cli.StringFlag{
		Name:        "tls.cert-file",
		Value:       "",
		Destination: &certFile,
		Usage:       "TLS Certificate file.",
		EnvVars:     []string{"TLS_CERT"},
	},
	&cli.StringFlag{
		Name:        "storage.path",
		Usage:       "Directory path to store logs.",
		Value:       "./db",
		Destination: &storagePath,
		EnvVars:     []string{"STORAGE_PATH"},
	},
	&cli.StringFlag{
		Name:     "secret-path",
		Usage:    "Path to a 32 bytes AES-256 secret used to encrypt logs. (use openssl rand -out secret.key 32)",
		Required: true,
		EnvVars:  []string{"SECRET_PATH"},
		Action: func(ctx *cli.Context, s string) (err error) {
			secret, err = os.ReadFile(s)
			if err != nil {
				return err
			}
			return nil
		},
	},
	&cli.BoolFlag{
		Name:    "debug",
		EnvVars: []string{"DEBUG"},
		Value:   false,
		Action: func(ctx *cli.Context, s bool) error {
			if s {
				logger.EnableDebug()
				go debug.WatchGoRoutines(ctx.Context)
			}
			return nil
		},
	},
}

var app = &cli.App{
	Name:    "grid-logger-server",
	Usage:   "Receives log and stores it",
	Flags:   flags,
	Suggest: true,
	Action: func(cCtx *cli.Context) error {
		lis, err := net.Listen("tcp", listenAddress)
		if err != nil {
			logger.I.Error("listen failed", zap.Error(err))
			return err
		}

		opts := []grpc.ServerOption{}
		if tls {
			creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
			if err != nil {
				logger.I.Fatal("failed to load certificates", zap.Error(err))
			}
			opts = append(opts, grpc.Creds(creds))
		}

		server := grpc.NewServer(opts...)
		loggerv1alpha1.RegisterLoggerAPIServer(
			server,
			api.NewLoggerAPIServer(
				db.NewFileDB(storagePath, secret),
			),
		)
		healthv1.RegisterHealthServer(
			server,
			health.New(),
		)

		logger.I.Info("listening")

		return server.Serve(lis)
	},
}

func main() {
	_ = godotenv.Load(".server.env", ".server.env.local")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
