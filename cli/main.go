package main

import (
	"math/big"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/deepsquare/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/deepsquare/sbatch"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/deepsquare-io/the-grid/cli/tui/status"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	trace bool
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
		Value:       "https://logger.deepsquare.run",
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
		Name:        "trace",
		Usage:       "Trace logging",
		Destination: &trace,
		EnvVars:     []string{"TRACE"},
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
		address := common.HexToAddress(metaschedulerSmartContract)
		ethclientRPC, err := ethclient.Dial(ethEndpointRPC)
		if err != nil {
			return err
		}
		ethclientWS, err := ethclient.Dial(ethEndpointWS)
		if err != nil {
			return err
		}
		pk, err := crypto.HexToECDSA(ethHexPK)
		if err != nil {
			return err
		}
		sbatch := sbatch.NewService(http.DefaultClient, sbatchEndpoint)
		rpc, err := metascheduler.NewRPC(address, ethclientRPC, big.NewInt(179188), pk, sbatch)
		if err != nil {
			return err
		}
		ws, err := metascheduler.NewWS(address, ethclientWS, big.NewInt(179188), pk)
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(status.Status(ctx, rpc, ws), tea.WithContext(ctx)).Run()
		return err
	},
}

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
