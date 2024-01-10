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

/*
DeepSquare TUI is a client to launch jobs on the DeepSquare Grid.

The DeepSquare Terminal User Interface (TUI) shows the job statuses, logs
and is able to launch DeepSquare Workflows from the terminal via the
meta-scheduler smart-contract deployed on a EVM blockchain.

USAGE:

To run the TUI:

	dps

To run the CLI commands:

	dps [global options] command [command options] [arguments...]

COMMANDS:

	allowance  Manage allowance.
	credit     Manage credits.
	init       Bootstrap a job workflow file.
	job        Manage jobs.
	provider   Manage providers (need to use an admin smart-contract).
	submit     Quickly submit a job.

GLOBAL OPTIONS:

	--metascheduler.rpc value             Metascheduler Avalanche C-Chain JSON-RPC endpoint. (default: "https://testnet.deepsquare.run/rpc") [$METASCHEDULER_RPC]


	--metascheduler.ws value              Metascheduler Avalanche C-Chain WS endpoint. (default: "wss://testnet.deepsquare.run/ws") [$METASCHEDULER_WS]


	--metascheduler.smart-contract value  Metascheduler smart-contract address. (default: "0x3707aB457CF457275b7ec32e203c54df80C299d5") [$METASCHEDULER_SMART_CONTRACT]


	--sbatch.endpoint value               SBatch Service GraphQL endpoint. (default: "https://sbatch.deepsquare.run/graphql") [$SBATCH_ENDPOINT]


	--logger.endpoint value               Grid Logger endpoint. (default: "https://grid-logger.deepsquare.run") [$LOGGER_ENDPOINT]


	--private-key value                   An hexadecimal private key for ethereum transactions. [$ETH_PRIVATE_KEY]

The environment variables must be initialized for proper usage.
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/deepsquare-io/grid/cli/cmd/allowance"
	"github.com/deepsquare-io/grid/cli/cmd/credit"
	"github.com/deepsquare-io/grid/cli/cmd/initc"
	"github.com/deepsquare-io/grid/cli/cmd/job"
	"github.com/deepsquare-io/grid/cli/cmd/provider"
	"github.com/deepsquare-io/grid/cli/cmd/submit"
	"github.com/deepsquare-io/grid/cli/cmd/validate"
	"github.com/deepsquare-io/grid/cli/deepsquare"
	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/tui/nav"
	versionpkg "github.com/deepsquare-io/grid/cli/version"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/mod/semver"
)

var (
	version          = "dev"
	availableVersion string

	sbatchEndpoint              string
	loggerEndpoint              string
	ethEndpointRPC              string
	ethEndpointWS               string
	ethHexPK                    string
	ethHexPKPath                string
	metaschedulerSmartContract  string
	metaschedulerOracleEndpoint string

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
		Name:        "metascheduler.oracle",
		Value:       metascheduler.DefaultOracleURL,
		Usage:       "Metascheduler Oracle endpoint.",
		Destination: &metaschedulerOracleEndpoint,
		EnvVars:     []string{"METASCHEDULER_ORACLE"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
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
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
		Value:       "",
	},
	&cli.StringFlag{
		Name:        "private-key.path",
		Usage:       "Path to an hexadecimal private key for ethereum transactions.",
		Destination: &ethHexPKPath,
		EnvVars:     []string{"ETH_PRIVATE_KEY_PATH"},
		Value:       "",
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
	Name:                 "dps",
	Version:              version,
	Usage:                "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
	Flags:                flags,
	Suggest:              true,
	EnableBashCompletion: true,
	Copyright: `dps  Copyright (C) 2023 DeepSquare Association
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.
See the GNU General Public License for more details.`,
	Commands: []*cli.Command{
		&allowance.Command,
		&credit.Command,
		&initc.Command,
		&job.Command,
		&provider.Command,
		&submit.Command,
		&validate.Command,
	},
	Action: func(cCtx *cli.Context) (err error) {
		ctx := cCtx.Context
		pk, err := utils.GetPrivateKey(ethHexPK, ethHexPKPath)
		if err != nil {
			return err
		}

		if metaschedulerSmartContract == "" {
			return errors.New("metascheduler smart-contract address not set")
		}

		client, err := deepsquare.NewClient(ctx, &deepsquare.ClientConfig{
			MetaschedulerAddress:        common.HexToAddress(metaschedulerSmartContract),
			RPCEndpoint:                 ethEndpointRPC,
			SBatchEndpoint:              sbatchEndpoint,
			LoggerEndpoint:              loggerEndpoint,
			UserPrivateKey:              pk,
			MetaschedulerOracleEndpoint: metaschedulerOracleEndpoint,
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

		nav := nav.NewModelBuilder().
			WithUserAddress(userAddress).
			WithClient(client).
			WithWatcher(watcher).
			WithVersion(version).
			WithAvailableVersion(availableVersion).
			WithMetaschedulerAddress(metaschedulerSmartContract).
			Build(ctx)

		_, err = tea.NewProgram(
			nav,
			tea.WithContext(ctx),
			tea.WithAltScreen(),
		).Run()
		return err
	},
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	v, err := versionpkg.CheckLatest(ctx)
	if err != nil {
		internallog.I.Warn("failed to check version", zap.Error(err))
	}
	availableVersion = v
	if semver.Compare(availableVersion, version) > 0 {
		app.Version = fmt.Sprintf("%s (new version available: %s)", version, availableVersion)
	}
}

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")
	if err := app.Run(os.Args); err != nil {
		internallog.I.Fatal("app crashed", zap.Error(err))
	}
}
