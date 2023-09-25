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
	"net/http"

	"github.com/deepsquare-io/the-grid/cli/deepsquare"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
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

var Command = cli.Command{
	Name:  "job",
	Usage: "Manage jobs.",
	Subcommands: []*cli.Command{
		{
			Name:      "get",
			Usage:     "Get job.",
			Flags:     flags,
			ArgsUsage: "<0x job ID (32 bytes hex)>",
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() != 1 {
					return errors.New("missing arguments")
				}
				var jobID [32]byte
				jobIDB, err := hexutil.Decode(cCtx.Args().First())
				if err != nil {
					return err
				}
				copy(jobID[:], jobIDB)
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
			ArgsUsage: "<0x job ID (32 bytes hex)>",
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() != 1 {
					return errors.New("missing arguments")
				}
				pk, err := crypto.HexToECDSA(ethHexPK)
				if err != nil {
					return err
				}
				var jobID [32]byte
				jobIDB, err := hexutil.Decode(cCtx.Args().First())
				if err != nil {
					return err
				}
				copy(jobID[:], jobIDB)
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
				fmt.Println("done")
				return nil
			},
		},
	},
}
