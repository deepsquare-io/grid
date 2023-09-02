package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/graph"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/logger"
	metricsv1 "github.com/deepsquare-io/the-grid/smart-contracts-exporter/metrics/v1"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/watcher"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	listenAddress        string
	metaschedulerAddress common.Address
	prometheusURL        string

	avaxEndpointWS  string
	avaxEndpointRPC string

	checkpointFile string

	version string = "dev"
)

var app = &cli.App{
	Name:    "smart-contracts-exporter",
	Version: version,
	Usage:   "Prometheus exporter for DeepSquare smart contracts.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "http.listen-address",
			Value:       ":3000",
			Usage:       "Address to listen on. Is used for receiving job status via the job completion plugin.",
			Destination: &listenAddress,
			EnvVars:     []string{"LISTEN_ADDRESS"},
		},
		&cli.StringFlag{
			Name:       "metascheduler.smart-contract",
			Usage:      "Metascheduler smart-contract address.",
			Value:      "0x3707aB457CF457275b7ec32e203c54df80C299d5",
			HasBeenSet: true,
			Action: func(ctx *cli.Context, s string) error {
				metaschedulerAddress = common.HexToAddress(s)
				return nil
			},
			EnvVars: []string{"METASCHEDULER_SMART_CONTRACT"},
		},
		&cli.StringFlag{
			Name:        "prometheus",
			Usage:       "Prometheus url",
			Required:    false,
			Destination: &prometheusURL,
			EnvVars:     []string{"PROMETHEUS_URL"},
		},
		&cli.PathFlag{
			Name:        "checkpoint.file",
			Usage:       "Checkpoint file to avoid re-reading the blockchain.",
			Value:       "checkpoint.bin",
			Destination: &checkpointFile,
			EnvVars:     []string{"CHECKPOINT_FILE"},
		},
		&cli.StringFlag{
			Name:        "avax.endpoint.rpc",
			Value:       "https://testnet.deepsquare.run/rpc",
			Usage:       "Avalanche C-Chain JSON-RPC endpoint.",
			Destination: &avaxEndpointRPC,
			EnvVars:     []string{"AVAX_ENDPOINT_RPC"},
		},
		&cli.StringFlag{
			Name:        "avax.endpoint.ws",
			Value:       "wss://testnet.deepsquare.run/ws",
			Usage:       "Avalanche C-Chain WS endpoint.",
			Destination: &avaxEndpointWS,
			EnvVars:     []string{"AVAX_ENDPOINT_WS"},
		},
		&cli.BoolFlag{
			Name:    "debug",
			EnvVars: []string{"DEBUG"},
			Aliases: []string{"d"},
			Value:   false,
			Action: func(ctx *cli.Context, s bool) error {
				if s {
					logger.EnableDebug()
				}
				return nil
			},
		},
	},
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context
		errChan := make(chan error, 10)
		var wg sync.WaitGroup
		// Prometheus configuration
		metricsv1.Init(metaschedulerAddress.Hex(), checkpointFile, version)

		if err := metricsv1.Load(); err != nil {
			logger.I.Error("failed to load checkpoint", zap.Error(err))
		}

		// Ethereum configuration
		ethClientRPC, err := ethclient.Dial(avaxEndpointRPC)
		if err != nil {
			logger.I.Fatal("ethclientRPC dial failed", zap.Error(err))
		}
		ethClientWS, err := ethclient.Dial(avaxEndpointWS)
		if err != nil {
			logger.I.Fatal("ethClientWS dial failed", zap.Error(err))
		}

		msRPC, err := metascheduler.NewMetaScheduler(metaschedulerAddress, ethClientRPC)
		if err != nil {
			logger.I.Fatal("msRPC dial failed", zap.Error(err))
		}
		msWS, err := metascheduler.NewMetaScheduler(metaschedulerAddress, ethClientWS)
		if err != nil {
			logger.I.Fatal("msWS dial failed", zap.Error(err))
		}

		// Watch
		readyChan := make(chan struct{})
		w := watcher.New(
			ethClientRPC,
			ethClientWS,
			msRPC,
			msWS,
			metaschedulerAddress,
		)
		wg.Add(1)
		go func() {
			errChan <- w.WatchNewEvents(ctx, readyChan)
			wg.Done()
		}()
		select {
		case <-readyChan:

		case err := <-errChan:
			return err

		case <-ctx.Done():
			return ctx.Err()
		}

		// Server configuration
		srv := handler.New(graph.NewExecutableSchema(
			graph.Config{
				Resolvers: graph.NewResolver(prometheusURL, metaschedulerAddress.Hex()),
			},
		))
		srv.AddTransport(&transport.Websocket{
			Upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		})
		srv.AddTransport(transport.Options{})
		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})
		srv.AddTransport(transport.MultipartForm{})
		srv.Use(extension.Introspection{})
		logger.I.Info("listening", zap.String("listeningAddress", listenAddress))
		r := chi.NewRouter()
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Content-Type", "X-Requested-With"},
			AllowCredentials: true,
			MaxAge:           300,
		}))

		r.HandleFunc("/", playground.ApolloSandboxHandler("GraphQL Playground", "/graphql"))
		r.Handle("/graphql", srv)
		r.Handle("/metrics", promhttp.Handler())
		server := &http.Server{
			Addr:    listenAddress,
			Handler: r,
		}

		wg.Add(1)
		go func() {
			errChan <- server.ListenAndServe()
			wg.Done()
		}()

		err = <-errChan

		if shutdownErr := server.Shutdown(ctx); shutdownErr != nil {
			logger.I.Error("server failed to shutdown", zap.Error(err))
		}
		wg.Wait()

		return err
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
