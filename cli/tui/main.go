package main

import (
	"crypto/tls"
	"crypto/x509"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/v1/internal/log"
	"github.com/deepsquare-io/the-grid/cli/v1/logger"
	"github.com/deepsquare-io/the-grid/cli/v1/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/v1/sbatch"
	"github.com/deepsquare-io/the-grid/cli/v1/tui/nav"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	version string = "dev"

	sbatchEndpoint             string
	loggerEndpoint             string
	ethEndpointRPC             string
	ethEndpointWS              string
	ethHexPK                   string
	metaschedulerSmartContract string

	debug bool
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "metascheduler.rpc",
		Value:       "https://testnet.deepsquare.run/rpc",
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointRPC,
		EnvVars:     []string{"METASCHEDULER_RPC"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.ws",
		Value:       "wss://testnet.deepsquare.run/ws",
		Usage:       "Metascheduler Avalanche C-Chain WS endpoint.",
		Destination: &ethEndpointWS,
		EnvVars:     []string{"METASCHEDULER_WS"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0xc9AcB97F1132f0FB5dC9c5733B7b04F9079540f0",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "sbatch.endpoint",
		Value:       "https://sbatch.deepsquare.run/graphql",
		Usage:       "SBatch Service GraphQL endpoint.",
		Destination: &sbatchEndpoint,
		EnvVars:     []string{"SBATCH_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "logger.endpoint",
		Value:       "https://grid-logger.deepsquare.run",
		Usage:       "Grid Logger endpoint.",
		Destination: &loggerEndpoint,
		EnvVars:     []string{"LOGGER_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
	&cli.BoolFlag{
		Name:        "debug",
		Usage:       "Debug logging",
		Destination: &debug,
		Action: func(ctx *cli.Context, b bool) error {
			if b {
				log.EnableDebug()
			}
			return nil
		},
		EnvVars: []string{"DEBUG"},
	},
}

var app = &cli.App{
	Name:    "dps",
	Version: version,
	Usage:   "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
	Flags:   flags,
	Suggest: true,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context
		// Load the system CA certificates
		caCertPool, err := x509.SystemCertPool()
		if err != nil {
			log.I.Warn("failed to load system CA certificates", zap.Error(err))
			caCertPool = x509.NewCertPool()
		}
		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		address := common.HexToAddress(metaschedulerSmartContract)
		rpcClient, err := rpc.DialOptions(ctx, ethEndpointRPC, rpc.WithHTTPClient(client))
		if err != nil {
			return err
		}
		ethClientRPC := ethclient.NewClient(rpcClient)
		wsClient, err := rpc.DialOptions(ctx, ethEndpointWS, rpc.WithHTTPClient(client))
		if err != nil {
			return err
		}
		ethClientWS := ethclient.NewClient(wsClient)
		pk, err := crypto.HexToECDSA(ethHexPK)
		if err != nil {
			return err
		}
		sbatch := sbatch.NewService(http.DefaultClient, sbatchEndpoint)
		rpc, err := metascheduler.NewRPC(address, ethClientRPC, big.NewInt(179188), pk, sbatch)
		if err != nil {
			return err
		}
		ws, err := metascheduler.NewWS(address, ethClientWS, big.NewInt(179188), pk)
		if err != nil {
			return err
		}
		dialOptions := []grpc.DialOption{
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		}
		u, err := url.Parse(loggerEndpoint)
		if err != nil {
			log.I.Error("Failed to parse URL", zap.Error(err))
			return err
		}

		port := u.Port()
		if port == "" {
			// If the URL doesn't explicitly specify a port, use the default port for the scheme.
			switch strings.ToLower(u.Scheme) {
			case "http":
				port = "80"
			case "https":
				port = "443"
			default:
				log.I.Fatal("Unknown scheme for logger URL", zap.String("scheme", u.Scheme))
			}
		}
		l, conn, err := logger.DialContext(
			ctx,
			net.JoinHostPort(u.Hostname(), port),
			pk,
			dialOptions...,
		)
		if err != nil {
			return err
		}
		defer conn.Close()
		_, err = tea.NewProgram(
			nav.Model(
				ctx,
				crypto.PubkeyToAddress(pk.PublicKey),
				rpc,
				ws,
				l,
				version,
				metaschedulerSmartContract,
			),
			tea.WithContext(ctx),
			tea.WithAltScreen(),
		).Run()
		return err
	},
}

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")
	if err := app.Run(os.Args); err != nil {
		log.I.Fatal("app crashed", zap.Error(err))
	}
}
