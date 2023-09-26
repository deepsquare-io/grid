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

package job

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/internal/ether"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
)

var (
	ethEndpointRPC             string
	metaschedulerSmartContract string
	ethHexPK                   string
	panicReason                string

	wei   bool
	time  bool
	force bool
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

var panicFlags = append(
	authFlags,
	&cli.StringFlag{
		Name:        "reason",
		Usage:       "Add reason for panicking.",
		Value:       "An admin has forced the job to stop.",
		Destination: &panicReason,
	},
)

var topupFlags = append(
	authFlags,
	&cli.BoolFlag{
		Name:        "wei",
		Usage:       "Use wei.",
		Destination: &wei,
	},
	&cli.BoolFlag{
		Name:        "time",
		Usage:       "Use duration instead of credits.",
		Destination: &time,
	},
	&cli.BoolFlag{
		Name:        "force",
		Usage:       "Don't ask for confirmation.",
		Destination: &force,
	},
)

var Command = cli.Command{
	Name:  "job",
	Usage: "Manage jobs.",
	Subcommands: []*cli.Command{
		{
			Name:      "get",
			Usage:     "Get job.",
			Flags:     flags,
			ArgsUsage: "<job ID>",
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() < 1 {
					return errors.New("missing arguments")
				}
				jobIDBig, ok := new(big.Int).SetString(cCtx.Args().First(), 10)
				if !ok {
					return errors.New("failed to parse job ID")
				}
				var jobID [32]byte
				jobIDBig.FillBytes(jobID[:])
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
				job, err := clientset.JobFetcher().GetJob(ctx, jobID)
				if err != nil {
					return err
				}
				jobJSON, err := json.MarshalIndent(job, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(jobJSON))
				return nil
			},
		},
		{
			Name:      "panic",
			Usage:     "Panic a job (need a METASCHEDULER role).",
			Flags:     panicFlags,
			ArgsUsage: "<job ID>",
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() < 1 {
					return errors.New("missing arguments")
				}
				pk, err := crypto.HexToECDSA(ethHexPK)
				if err != nil {
					return err
				}
				jobIDBig, ok := new(big.Int).SetString(cCtx.Args().First(), 10)
				if !ok {
					return errors.New("failed to parse job ID")
				}
				var jobID [32]byte
				jobIDBig.FillBytes(jobID[:])
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
					UserPrivateKey:       pk,
				})
				if err := clientset.JobScheduler(nil).PanicJob(ctx, jobID, panicReason); err != nil {
					return err
				}
				fmt.Println("Done.")
				return nil
			},
		},
		{
			Name:      "topup",
			Usage:     "Top up a job.",
			Flags:     topupFlags,
			ArgsUsage: "<job ID> <amount (use --time to topup with a duration)>",
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() < 2 {
					return errors.New("missing arguments")
				}
				pk, err := crypto.HexToECDSA(ethHexPK)
				if err != nil {
					return err
				}
				jobIDBig, ok := new(big.Int).SetString(cCtx.Args().First(), 10)
				if !ok {
					return fmt.Errorf("couldn't parse job ID: %s", cCtx.Args().First())
				}
				var jobID [32]byte
				jobIDBig.FillBytes(jobID[:])

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
					UserPrivateKey:       pk,
				})

				var creditsWei *big.Int
				var credits *big.Float
				if !time {
					if wei {
						c, ok := new(big.Int).SetString(cCtx.Args().Get(1), 10)
						if !ok {
							return fmt.Errorf("couldn't parse amount: %s", cCtx.Args().Get(1))
						}
						creditsWei = c
						credits = ether.FromWei(creditsWei)
					} else {
						c, ok := new(big.Float).SetString(cCtx.Args().Get(1))
						if !ok {
							return fmt.Errorf("couldn't parse amount: %s", cCtx.Args().Get(1))
						}
						credits = c
						creditsWei = ether.ToWei(credits)
					}
				} else {
					c, ok := new(big.Int).SetString(cCtx.Args().Get(1), 10)
					if !ok {
						return errors.New("couldn't parse duration")
					}
					job, err := clientset.JobFetcher().GetJob(ctx, jobID)
					if err != nil {
						return err
					}
					p, err := clientset.ProviderManager().GetProvider(ctx, job.ProviderAddr)
					if err != nil {
						return err
					}
					creditsWei = metascheduler.DurationToCredit(p.ProviderPrices, job.Definition, c)
					credits = ether.FromWei(creditsWei)
				}

				if !force {
					msg := fmt.Sprintf(
						"Confirm topup of %s credits (%s wei) to job %s?",
						credits.String(),
						creditsWei.String(),
						hexutil.Encode(jobID[:]),
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

				if err := clientset.JobScheduler(nil).TopUpJob(ctx, jobID, creditsWei); err != nil {
					return err
				}
				fmt.Println("done")
				return nil
			},
		},
	},
}
