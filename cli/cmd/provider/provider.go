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

/*
Package provider provides subcommands to manage the providers of the Grid.

USAGE:

	dps provider command [command options] [arguments...]

COMMANDS:

	list     List providers.
	approve  Approve a provider.
	remove   Remove a provider.
	help, h  Shows a list of commands or help for one command
*/
package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
)

var (
	ethEndpointRPC             string
	metaschedulerSmartContract string
	ethHexPK                   string
	ethHexPKPath               string
	proposal                   bool
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
		Name:        "metascheduler.smart-contract",
		Required:    true,
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
}

var listFlags = append(
	flags,
	&cli.BoolFlag{
		Name:        "proposal",
		Usage:       "See the proposal of the provider.",
		Destination: &proposal,
		Value:       false,
	},
)

var authFlags = append(
	flags,
	&cli.StringFlag{
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Value:       "",
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
	&cli.StringFlag{
		Name:        "private-key.path",
		Usage:       "Path to an hexadecimal private key for ethereum transactions.",
		Destination: &ethHexPKPath,
		EnvVars:     []string{"ETH_PRIVATE_KEY_PATH"},
		Value:       "",
	},
)

// Command is the provider command used to manage providers.
var Command = cli.Command{
	Name:  "provider",
	Usage: "Manage providers (need to use an admin smart-contract).",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "List providers.",
			Flags: listFlags,
			Action: func(cCtx *cli.Context) error {
				ctx := cCtx.Context
				rpcClient, err := rpc.DialOptions(
					ctx,
					ethEndpointRPC,
					rpc.WithHTTPClient(http.DefaultClient),
				)
				if err != nil {
					return err
				}
				defer rpcClient.Close()
				ethClientRPC := ethclient.NewClient(rpcClient)
				chainID, err := ethClientRPC.ChainID(ctx)
				if err != nil {
					return err
				}
				clientset := metascheduler.NewRPCClientSet(metascheduler.Backend{
					EthereumBackend:      ethClientRPC,
					MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
					ChainID:              chainID,
				})
				opts := make([]types.GetProviderOption, 0)
				if proposal {
					opts = append(opts, types.WithProposal())
				}
				providers, err := clientset.ProviderManager().GetProviders(ctx, opts...)
				if err != nil {
					return err
				}
				providersJSON, err := json.MarshalIndent(providers, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(providersJSON))
				return nil
			},
		},
		{
			Name:      "approve",
			Usage:     "Approve a provider.",
			ArgsUsage: "<0x>",
			Flags:     authFlags,
			Action: func(cCtx *cli.Context) error {
				clientset, err := initMutableClientSet(cCtx)
				if err != nil {
					return err
				}
				providerAddress := common.HexToAddress(cCtx.Args().First())
				return clientset.ProviderManager().Approve(cCtx.Context, providerAddress)
			},
		},
		{
			Name:      "remove",
			Usage:     "Remove a provider.",
			ArgsUsage: "<0x>",
			Flags:     authFlags,
			Action: func(cCtx *cli.Context) error {
				clientset, err := initMutableClientSet(cCtx)
				if err != nil {
					return err
				}
				providerAddress := common.HexToAddress(cCtx.Args().First())
				return clientset.ProviderManager().Remove(cCtx.Context, providerAddress)
			},
		},
	},
}

func initMutableClientSet(cCtx *cli.Context) (*metascheduler.RPCClientSet, error) {
	if cCtx.NArg() < 1 {
		return nil, errors.New("missing arguments")
	}
	ctx := cCtx.Context
	pk, err := utils.GetPrivateKey(ethHexPK, ethHexPKPath)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpc.DialOptions(
		ctx,
		ethEndpointRPC,
		rpc.WithHTTPClient(http.DefaultClient),
	)
	if err != nil {
		return nil, err
	}
	defer rpcClient.Close()
	ethClientRPC := ethclient.NewClient(rpcClient)
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	return metascheduler.NewRPCClientSet(metascheduler.Backend{
		EthereumBackend:      ethClientRPC,
		MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
		ChainID:              chainID,
		UserPrivateKey:       pk,
	}), nil
}
