package main

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/deepsquare-io/the-grid/sbatch-service/graph"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	listenAddress string

	redisAddress      string
	redisTLS          bool
	redisTLSInsecure  bool
	redisCAFile       string
	redisHostOverride string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "http.listen-address",
		Value:       ":3000",
		Usage:       "Address to listen on. Is used for receiving job status via the job completion plugin.",
		Destination: &listenAddress,
		EnvVars:     []string{"LISTEN_ADDRESS"},
	},
	&cli.StringFlag{
		Name:        "redis.url",
		Usage:       "Redis storage connection string.",
		Value:       "./db",
		Destination: &redisAddress,
		EnvVars:     []string{"REDIS_ADDRESS"},
	},
	&cli.BoolFlag{
		Name:        "redis.tls",
		Value:       true,
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
	Name:    "sbatch-service",
	Usage:   "sbatch script hosting service",
	Flags:   flags,
	Suggest: true,
	Action: func(cCtx *cli.Context) error {
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

		srv := handler.NewDefaultServer(graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					RedisClient: rdb,
				},
			},
		))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)
		http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
			_, _ = w.Write([]byte("ok"))
		})

		logger.I.Info("listening", zap.String("listeningAddress", listenAddress))

		return http.ListenAndServe(listenAddress, nil)
	},
}

func main() {
	_ = godotenv.Load(".server.env", ".server.env.local")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
