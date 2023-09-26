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

package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
)

var (
	ethEndpointRPC             string
	metaschedulerSmartContract string
	ethHexPK                   string
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
		Value:       deepsquare.DefaultMetaSchedulerAddress.Hex(),
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
}

var authFlags = append(
	flags,
	&cli.StringFlag{
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
)

var Command = cli.Command{
	Name:  "provider",
	Usage: "Manage providers (need to use an admin smart-contract).",
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "List providers.",
			Flags: flags,
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
				providers, err := clientset.ProviderManager().GetProviders(ctx)
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
				if cCtx.NArg() < 1 {
					return errors.New("missing arguments")
				}
				ctx := cCtx.Context
				pk, err := crypto.HexToECDSA(ethHexPK)
				if err != nil {
					return err
				}
				providerAddress := common.HexToAddress(cCtx.Args().First())
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
					UserPrivateKey:       pk,
				})
				return clientset.ProviderManager().Approve(ctx, providerAddress)
			},
		},
		{
			Name:      "remove",
			Usage:     "Remove a provider.",
			ArgsUsage: "<0x>",
			Flags:     authFlags,
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() < 1 {
					return errors.New("missing arguments")
				}
				ctx := cCtx.Context
				pk, err := crypto.HexToECDSA(ethHexPK)
				if err != nil {
					return err
				}
				providerAddress := common.HexToAddress(cCtx.Args().First())
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
					UserPrivateKey:       pk,
				})
				return clientset.ProviderManager().Remove(ctx, providerAddress)
			},
		},
	},
}
