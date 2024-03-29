// Copyright (C) 2024 DeepSquare Association
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
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	metaschedulerabi "github.com/deepsquare-io/grid/sbatch-service/abi/metascheduler"
	internalauth "github.com/deepsquare-io/grid/sbatch-service/auth"
	"github.com/deepsquare-io/grid/sbatch-service/cmd"
	sbatchapiv1alpha1 "github.com/deepsquare-io/grid/sbatch-service/gen/go/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/grid/sbatch-service/graph"
	"github.com/deepsquare-io/grid/sbatch-service/graph/playground"
	"github.com/deepsquare-io/grid/sbatch-service/grpc/auth"
	"github.com/deepsquare-io/grid/sbatch-service/grpc/sbatch"
	"github.com/deepsquare-io/grid/sbatch-service/logger"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/deepsquare-io/grid/sbatch-service/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

var (
	listenAddress string

	redisDisable      bool
	redisAddress      string
	redisTLS          bool
	redisTLSInsecure  bool
	redisCAFile       string
	redisHostOverride string

	loggerEndpoint string
	loggerPath     string

	preJobScript  string
	postJobScript string

	metaschedulerEndpointRPC   string
	metaschedulerSmartContract string

	version = "dev"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "http.listen-address",
		Value:       ":3000",
		Usage:       "Address to listen on. Is used for receiving job status via the job completion plugin.",
		Destination: &listenAddress,
		EnvVars:     []string{"LISTEN_ADDRESS"},
	},
	&cli.BoolFlag{
		Name:        "redis.disable",
		Usage:       "Disable Redis and use internal map.",
		Value:       false,
		Destination: &redisDisable,
		EnvVars:     []string{"REDIS_DISABLE"},
	},
	&cli.StringFlag{
		Name:        "redis.url",
		Usage:       "Redis storage connection string.",
		Value:       "redis://redis:6379/0",
		Destination: &redisAddress,
		EnvVars:     []string{"REDIS_ADDRESS"},
	},
	&cli.BoolFlag{
		Name:        "redis.tls",
		Value:       false,
		Usage:       "Enable TLS for Redis.",
		Destination: &redisTLS,
		EnvVars:     []string{"REDIS_TLS_ENABLE"},
	},
	&cli.BoolFlag{
		Name:        "redis.tls.insecure",
		Value:       false,
		Usage:       "Skip TLS verification. By enabling it, redis.tls.ca and redis.tls.server-host-override are ignored.",
		Destination: &redisTLSInsecure,
		EnvVars:     []string{"REDIS_TLS_INSECURE"},
	},
	&cli.StringFlag{
		Name:        "redis.tls.ca",
		Value:       "",
		Usage:       "Path to CA certificate for TLS verification.",
		Destination: &redisCAFile,
		EnvVars:     []string{"REDIS_CA"},
	},
	&cli.StringFlag{
		Name:        "redis.tls.host-override",
		Value:       "",
		Usage:       "The server name used to verify the hostname returned by the TLS handshake.",
		Destination: &redisHostOverride,
		EnvVars:     []string{"REDIS_HOST_OVERRIDE"},
	},
	&cli.StringFlag{
		Name:        "logger.endpoint",
		Value:       "",
		Usage:       "The grid logger URL endpoint. (ex: logger.example.com:443)",
		Destination: &loggerEndpoint,
		EnvVars:     []string{"LOGGER_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "logger.writer.path",
		Value:       "/usr/local/bin/grid-logger-writer",
		Usage:       "Path of the grid logger on the compute nodes.",
		Destination: &loggerPath,
		EnvVars:     []string{"LOGGER_WRITER_PATH"},
	},
	&cli.StringFlag{
		Name:        "hook.pre.path",
		Usage:       "Path to a prescript which will be embedded in the job. Will be run with bash at the very beginning (before logging).",
		Required:    false,
		Destination: &preJobScript,
		EnvVars:     []string{"HOOK_PRE_PATH"},
	},
	&cli.StringFlag{
		Name:        "hook.post.path",
		Usage:       "Path to a postscript which will be embedded in the job. Will be run with bash at the very end (after logging).",
		Required:    false,
		Destination: &postJobScript,
		EnvVars:     []string{"HOOK_POST_PATH"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.endpoint.rpc",
		Value:       "https://testnet.deepsquare.run/rpc",
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &metaschedulerEndpointRPC,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT_RPC"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0x3707aB457CF457275b7ec32e203c54df80C299d5",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.BoolFlag{
		Name:    "debug",
		EnvVars: []string{"DEBUG"},
		Value:   false,
		Action: func(_ *cli.Context, s bool) error {
			if s {
				logger.EnableDebug()
			}
			return nil
		},
	},
}

var app = &cli.App{
	Name:                 "sbatch-service",
	Usage:                "sbatch script hosting service",
	Flags:                flags,
	Suggest:              true,
	EnableBashCompletion: true,
	Copyright: `sbatch-service  Copyright (C) 2023 DeepSquare Association
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.
See the GNU General Public License for more details.`,
	Version: version,
	Commands: []*cli.Command{
		&cmd.RenderCmd,
	},
	Action: func(cCtx *cli.Context) error {
		logger.I.Info("running", zap.String("version", version))
		var stor storage.Storage
		// Redis connection
		if !redisDisable {
			opt, err := redis.ParseURL(redisAddress)
			if err != nil {
				return err
			}
			if redisTLS {
				var tlsConfig tls.Config
				if redisTLSInsecure {
					tlsConfig = tls.Config{
						InsecureSkipVerify: true,
					}
				} else {
					certs, err := x509.SystemCertPool()
					if err != nil {
						logger.I.Warn("failed to load system certs pool")
						certs = x509.NewCertPool()
					}
					if redisCAFile != "" {
						pem, err := os.ReadFile(redisCAFile)
						if err != nil {
							return err
						}

						certs.AppendCertsFromPEM(pem)
					}
					tlsConfig = tls.Config{
						MinVersion: tls.VersionTLS12,
						RootCAs:    certs,
					}
				}
				opt.TLSConfig = &tlsConfig
			}
			rdb := redis.NewClient(opt)
			stor = storage.NewRedisStorage(rdb)
		} else {
			stor = storage.NewInMemoryStorage()
		}

		jobRenderer := renderer.NewJobRenderer(
			loggerEndpoint,
			loggerPath,
			renderer.WithPostscript(postJobScript),
			renderer.WithPrescript(preJobScript),
		)

		// GraphQL server
		c := graph.Config{
			Resolvers: graph.NewResolver(stor, jobRenderer),
		}
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
		r := chi.NewRouter()
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           300,
		}))
		r.Handle("/", playground.ApolloSandboxHandler("GraphQL playground", "/graphql"))
		r.Handle("/graphql", srv)
		r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("ok"))
		})
		if cCtx.Bool("debug") {
			r.HandleFunc("/job/{jobID}", func(w http.ResponseWriter, r *http.Request) {
				jobID := chi.URLParam(r, "jobID")
				logger.I.Info("get", zap.String("batchLocationHash", jobID))
				resp, err := stor.Get(r.Context(), jobID)
				if err != nil {
					logger.I.Error("get failed", zap.Error(err))
					http.Error(w, http.StatusText(404), 404)
					return
				}
				_, err = w.Write([]byte(resp))
				if err != nil {
					logger.I.Error("get failed: write", zap.Error(err))
					http.Error(w, http.StatusText(500), 500)
				}
			})
		}

		client, err := ethclient.DialContext(cCtx.Context, metaschedulerEndpointRPC)
		if err != nil {
			return err
		}
		ms, err := metaschedulerabi.NewMetaScheduler(
			common.HexToAddress(metaschedulerSmartContract),
			client,
		)
		if err != nil {
			return err
		}
		jobsAddr, err := ms.Jobs(&bind.CallOpts{
			Context: cCtx.Context,
		})
		if err != nil {
			return err
		}
		jobs, err := metaschedulerabi.NewIJobRepository(jobsAddr, client)
		if err != nil {
			return err
		}

		// gRPC server
		g := grpc.NewServer()
		a := internalauth.NewAuth(stor)
		sbatchapiv1alpha1.RegisterSBatchAPIServer(g, sbatch.NewAPI(a, stor, jobs, loggerEndpoint))
		sbatchapiv1alpha1.RegisterAuthAPIServer(g, auth.NewAPI(a))

		rg := mixedHandler(r, g)

		http2Server := &http2.Server{}
		http1Server := &http.Server{Handler: h2c.NewHandler(rg, http2Server)}

		// Listener
		lis, err := net.Listen("tcp", listenAddress)
		if err != nil {
			return err
		}

		logger.I.Info("listening", zap.String("listeningAddress", listenAddress))

		return http1Server.Serve(lis)
	},
}

func mixedHandler(httpHand http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 &&
			strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
			return
		}
		httpHand.ServeHTTP(w, r)
	})
}

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
