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
	_ "log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	healthv1 "github.com/deepsquare-io/grid/grid-logger/gen/go/grpc/health/v1"
	loggerv1alpha1 "github.com/deepsquare-io/grid/grid-logger/gen/go/logger/v1alpha1"
	"github.com/deepsquare-io/grid/grid-logger/logger"
	"github.com/deepsquare-io/grid/grid-logger/server/api"
	"github.com/deepsquare-io/grid/grid-logger/server/api/health"
	"github.com/deepsquare-io/grid/grid-logger/server/db"
	"github.com/deepsquare-io/grid/grid-logger/server/debug"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

var (
	listenAddress string

	tls      bool
	keyFile  string
	certFile string

	storagePath string

	secret []byte

	pprofListenAddress string
	pprofEnabled       bool
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
	&cli.BoolFlag{
		Name:        "pprof",
		Usage:       "Enable pprof",
		Value:       false,
		Destination: &pprofEnabled,
		EnvVars:     []string{"PPROF_ENABLE"},
	},
	&cli.StringFlag{
		Name:        "pprof.listen-address",
		Usage:       "Address to listen on for pprof",
		Value:       ":9000",
		Destination: &pprofListenAddress,
		EnvVars:     []string{"PPROF_LISTEN_ADDRESS"},
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
	Name:                 "grid-logger-server",
	Usage:                "Receives log and stores it",
	Flags:                flags,
	Suggest:              true,
	EnableBashCompletion: true,
	Copyright: `grid-logger-server  Copyright (C) 2023 DeepSquare Association
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.
See the GNU General Public License for more details.`,
	Action: func(cCtx *cli.Context) error {
		lis, err := net.Listen("tcp", listenAddress)
		if err != nil {
			logger.I.Error("listen failed", zap.Error(err))
			return err
		}

		opts := []grpc.ServerOption{
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle:     60 * time.Minute, // If a client is idle for given duration, send a GOAWAY.
				MaxConnectionAge:      60 * time.Minute, // If any connection is alive for more than given duration, send a GOAWAY.
				MaxConnectionAgeGrace: 10 * time.Second, // Allow given duration for pending RPCs to complete before forcibly closing connections
				Time:                  10 * time.Second, // Ping the client if it is idle for given duration to ensure the connection is still active.
				Timeout:               20 * time.Second, // Wait given duration for the ping ack before assuming the connection is dead.
			}),
		}
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

		if pprofEnabled {
			go func() {
				if err := http.ListenAndServe(pprofListenAddress, nil); err != nil {
					logger.I.Warn("pprof crashed", zap.Error(err))
				}
			}()
		}

		return server.Serve(lis)
	},
}

func main() {
	_ = godotenv.Load(".server.env", ".server.env.local")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
