// Copyright (C) 2023 DeepSquare
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
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/the-grid/cli/cmd/submit"
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
		Value:       deepsquare.DefaultRPCEndpoint,
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointRPC,
		EnvVars:     []string{"METASCHEDULER_RPC"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.ws",
		Value:       deepsquare.DefaultWSEndpoint,
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
	Commands: []*cli.Command{
		&submit.Command,
	},
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context
		pk, err := crypto.HexToECDSA(ethHexPK)
		if err != nil {
			return err
		}
		client, err := deepsquare.NewClient(ctx, &deepsquare.ClientConfig{
			MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
			RPCEndpoint:          ethEndpointRPC,
			SBatchEndpoint:       sbatchEndpoint,
			LoggerEndpoint:       loggerEndpoint,
			UserPrivateKey:       pk,
		})
		if err != nil {
			return err
		}
		defer func() {
			if err := client.Close(); err != nil {
				internallog.I.Error("failed to close client", zap.Error(err))
			}
		}()
		watcher, err := deepsquare.NewWatcher(ctx, &deepsquare.WatcherConfig{
			MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
			RPCEndpoint:          ethEndpointRPC,
			WSEndpoint:           ethEndpointWS,
			UserPrivateKey:       pk,
		})
		if err != nil {
			return err
		}
		defer func() {
			if err := watcher.Close(); err != nil {
				internallog.I.Error("failed to close watcher", zap.Error(err))
			}
		}()
		userAddress := crypto.PubkeyToAddress(pk.PublicKey)
		_, err = tea.NewProgram(
			nav.Model(
				ctx,
				userAddress,
				client,
				watcher,
				status.Model(
					ctx,
					client,
					watcher,
					userAddress,
				),
				log.ModelBuilder{
					Logger:      client,
					UserAddress: userAddress,
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
