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

package allowance

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"

	"github.com/erikgeiser/promptkit/confirmation"
)

var (
	ethEndpointRPC             string
	ethHexPK                   string
	metaschedulerSmartContract string

	credits    *big.Float
	creditsWei *big.Int
	force      bool
)

// "Get" flags
var (
	wei bool
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
	&cli.StringFlag{
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
	&cli.BoolFlag{
		Name:        "force",
		Usage:       "Do not prompt",
		Destination: &force,
		Aliases:     []string{"f"},
		EnvVars:     []string{"FORCE"},
	},
	&cli.BoolFlag{
		Name:        "wei",
		Usage:       "Show in wei.",
		Destination: &wei,
	},
}

var Command = cli.Command{
	Name:  "allowance",
	Usage: "Manage allowance.",
	Subcommands: []*cli.Command{
		{
			Name:      "set",
			Usage:     "Set credits allowance.",
			ArgsUsage: "<amount in credits (use --wei for wei)>",
			Flags:     flags,
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() != 1 {
					return errors.New("missing arguments")
				}
				if wei {
					c, ok := new(big.Int).SetString(cCtx.Args().First(), 10)
					if !ok {
						return errors.New("couldn't parse amount")
					}
					creditsWei = c
					credits = ether.FromWei(creditsWei)
				} else {
					c, ok := new(big.Float).SetString(cCtx.Args().First())
					if !ok {
						return errors.New("couldn't parse amount")
					}
					credits = c
					creditsWei = ether.ToWei(credits)
				}
				ctx := cCtx.Context
				clientset, err := initClient(ctx)
				if err != nil {
					return err
				}

				if !force {
					msg := fmt.Sprintf(
						"Set allowance of %s credits (%s wei) to %s?",
						credits.String(),
						creditsWei.String(),
						metaschedulerSmartContract,
					)
					input := confirmation.New(msg, confirmation.No)
					ok, err := input.RunPrompt()
					if err != nil {
						return err
					}
					if !ok {
						fmt.Println("Cancelled.")
						return nil
					}
				}

				if err := clientset.AllowanceManager().SetAllowance(ctx, creditsWei); err != nil {
					return err
				}

				fmt.Println("Done.")
				return nil
			},
		},
		{
			Name:  "get",
			Usage: "Get the allowance.",
			Flags: flags,
			Action: func(cCtx *cli.Context) error {
				ctx := cCtx.Context
				clientset, err := initClient(ctx)
				if err != nil {
					return err
				}

				amount, err := clientset.AllowanceManager().GetAllowance(ctx)
				if err != nil {
					return err
				}

				if wei {
					fmt.Println(amount.String())
				} else {
					fmt.Println(ether.FromWei(amount).String())
				}

				return nil
			},
		},
	},
}

func initClient(ctx context.Context) (*metascheduler.RPCClientSet, error) {
	pk, err := crypto.HexToECDSA(ethHexPK)
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
	clientset := metascheduler.NewRPCClientSet(metascheduler.Backend{
		EthereumBackend:      ethClientRPC,
		MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
		ChainID:              chainID,
		UserPrivateKey:       pk,
	})
	return clientset, nil
}
