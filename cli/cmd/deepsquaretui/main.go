package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	internallog "github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/deepsquare-io/the-grid/cli/tui/editor"
	"github.com/deepsquare-io/the-grid/cli/tui/log"
	"github.com/deepsquare-io/the-grid/cli/tui/nav"
	"github.com/deepsquare-io/the-grid/cli/tui/status"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
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
		Value:       deepsquare.DefaultEndpointRPC,
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointRPC,
		EnvVars:     []string{"METASCHEDULER_RPC"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.ws",
		Value:       deepsquare.DefaultEndpointWS,
		Usage:       "Metascheduler Avalanche C-Chain WS endpoint.",
		Destination: &ethEndpointWS,
		EnvVars:     []string{"METASCHEDULER_WS"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       deepsquare.DefaultMetaSchedulerAddress.Hex(),
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "sbatch.endpoint",
		Value:       deepsquare.DefaultSBatchEndpoint,
		Usage:       "SBatch Service GraphQL endpoint.",
		Destination: &sbatchEndpoint,
		EnvVars:     []string{"SBATCH_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "logger.endpoint",
		Value:       deepsquare.DefaultLoggerEndpoint,
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
				internallog.EnableDebug()
			}
			return nil
		},
		EnvVars: []string{"DEBUG"},
	},
}

var app = &cli.App{
	Name:    "deepsquaretui",
	Version: version,
	Usage:   "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
	Flags:   flags,
	Suggest: true,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context
		pk, err := crypto.HexToECDSA(ethHexPK)
		if err != nil {
			return err
		}
		client, err := deepsquare.NewClient(ctx, &deepsquare.ClientConfig{
			MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
			EndpointRPC:          ethEndpointRPC,
			SBatchEndpoint:       sbatchEndpoint,
			LoggerEndpoint:       loggerEndpoint,
			UserPrivateKey:       pk,
		})
		if err != nil {
			return err
		}
		watcher, err := deepsquare.NewWatcher(ctx, &deepsquare.WatcherConfig{
			MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
			EndpointRPC:          ethEndpointRPC,
			EndpointWS:           ethEndpointWS,
			UserPrivateKey:       pk,
		})
		if err != nil {
			return err
		}
		userAddress := crypto.PubkeyToAddress(pk.PublicKey)
		_, err = tea.NewProgram(
			nav.Model(
				ctx,
				userAddress,
				watcher,
				status.Model(
					ctx,
					client,
					watcher,
					userAddress,
				),
				log.ModelBuilder{
					LoggerDialer: client,
					UserAddress:  userAddress,
				},
				editor.ModelBuilder{
					Client: client,
				},
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
		internallog.I.Fatal("app crashed", zap.Error(err))
	}
}
