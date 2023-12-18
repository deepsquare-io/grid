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
Package job provides subcommands to manage the jobs of an user.

USAGE:

	dps job command [command options] [arguments...]

COMMANDS:

	get      Get job.
	panic    Panic a job (need a METASCHEDULER role).
	topup    Top up a job.
	help, h  Shows a list of commands or help for one command
*/
package job

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/internal/ether"
	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	ethEndpointRPC             string
	ethEndpointWS              string
	metaschedulerSmartContract string
	ethHexPK                   string
	panicReason                string
	loggerEndpoint             string

	wei         bool
	useTime     bool
	force       bool
	noTimestamp bool
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

var logsFlags = append(
	authFlags,
	&cli.StringFlag{
		Name:        "metascheduler.ws",
		Value:       deepsquare.DefaultWSEndpoint,
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointWS,
		EnvVars:     []string{"METASCHEDULER_WS"},
	},
	&cli.StringFlag{
		Name:        "logger.endpoint",
		Value:       deepsquare.DefaultLoggerEndpoint,
		Usage:       "Grid Logger endpoint.",
		Destination: &loggerEndpoint,
		EnvVars:     []string{"LOGGER_ENDPOINT"},
	},
	&cli.BoolFlag{
		Name:        "no-timestamp",
		Usage:       "Hide timestamp.",
		Aliases:     []string{"no-ts"},
		Category:    "Submit Settings:",
		Destination: &noTimestamp,
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
		Destination: &useTime,
	},
	&cli.BoolFlag{
		Name:        "force",
		Usage:       "Don't ask for confirmation.",
		Destination: &force,
	},
)

// Command is the job subcommand used to manage jobs.
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
			Name:      "logs",
			Usage:     "Watch job logs.",
			Flags:     logsFlags,
			ArgsUsage: "<job ID>",
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() < 1 {
					return errors.New("missing arguments")
				}
				kb, err := hexutil.Decode(ethHexPK)
				if errors.Is(err, hexutil.ErrMissingPrefix) {
					kb, err = hex.DecodeString(ethHexPK)
					if err != nil {
						return err
					}
				} else if err != nil {
					return err
				}
				pk, err := crypto.ToECDSA(kb)
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
				client, err := deepsquare.NewClient(ctx, &deepsquare.ClientConfig{
					MetaschedulerAddress: common.HexToAddress(metaschedulerSmartContract),
					RPCEndpoint:          ethEndpointRPC,
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
				transitions := make(chan types.JobTransition, 1)
				sub, err := watcher.SubscribeEvents(
					ctx,
					types.FilterJobTransition(transitions),
				)
				if err != nil {
					return err
				}
				defer sub.Unsubscribe()

				// Find job
				job, err := client.GetJob(ctx, jobID)
				if err != nil {
					return err
				}

				switch metascheduler.JobStatus(job.Status) {
				case metascheduler.JobStatusCancelled,
					metascheduler.JobStatusFailed,
					metascheduler.JobStatusFinished,
					metascheduler.JobStatusPanicked,
					metascheduler.JobStatusOutOfCredits,
					metascheduler.JobStatusRunning:
				default:
					_, err = waitUntilJobRunningOrFinished(sub, transitions, jobID)
					if err != nil {
						fmt.Printf("---Waiting for job running failed---\n%s\n", err)
						return err
					}
				}

				stream, err := client.WatchLogs(ctx, jobID)
				if err != nil {
					fmt.Printf("---Watching logs has unexpectedly failed---\n%s\n", err)
					return err
				}
				defer func() {
					_ = stream.CloseSend()
				}()
				for {
					req, err := stream.Recv()
					if err == io.EOF || errors.Is(err, context.Canceled) {
						fmt.Println("---Connection to logging server closed---")
						return nil
					}
					if err != nil {
						fmt.Printf(
							"---Connection to logging server closed unexpectedly---\n%s\n",
							err,
						)
						return err
					}
					clean := forbiddenReplacer.Replace(string(req.GetData()))
					if noTimestamp {
						fmt.Printf("%s\n", clean)
					} else {
						fmt.Printf("%s:\t%s\n", time.Unix(0, req.GetTimestamp()), clean)
					}
				}
			},
		},
		{
			Name:      "cancel",
			Usage:     "Cancel job.",
			Flags:     authFlags,
			ArgsUsage: "<job ID>",
			Action: func(cCtx *cli.Context) error {
				if cCtx.NArg() < 1 {
					return errors.New("missing arguments")
				}
				kb, err := hexutil.Decode(ethHexPK)
				if errors.Is(err, hexutil.ErrMissingPrefix) {
					kb, err = hex.DecodeString(ethHexPK)
					if err != nil {
						return err
					}
				} else if err != nil {
					return err
				}
				pk, err := crypto.ToECDSA(kb)
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
				if err := clientset.JobScheduler(nil).CancelJob(ctx, jobID); err != nil {
					return err
				}
				fmt.Println("Done.")
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
				kb, err := hexutil.Decode(ethHexPK)
				if errors.Is(err, hexutil.ErrMissingPrefix) {
					kb, err = hex.DecodeString(ethHexPK)
					if err != nil {
						return err
					}
				} else if err != nil {
					return err
				}
				pk, err := crypto.ToECDSA(kb)
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
				kb, err := hexutil.Decode(ethHexPK)
				if errors.Is(err, hexutil.ErrMissingPrefix) {
					kb, err = hex.DecodeString(ethHexPK)
					if err != nil {
						return err
					}
				} else if err != nil {
					return err
				}
				pk, err := crypto.ToECDSA(kb)
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
				if !useTime {
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
					jobIDBig := new(big.Int).SetBytes(jobID[:])
					msg := fmt.Sprintf(
						"Confirm topup of %s credits (%s wei) to job %s?",
						credits.String(),
						creditsWei.String(),
						jobIDBig.String(),
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

				// Add allowance
				curr, err := clientset.AllowanceManager().GetAllowance(ctx)
				if err != nil {
					return err
				}
				if err = clientset.AllowanceManager().SetAllowance(ctx, curr.Add(curr, creditsWei)); err != nil {
					return err
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

func waitUntilJobRunningOrFinished(
	sub ethereum.Subscription,
	ch <-chan types.JobTransition,
	jobID [32]byte,
) (metascheduler.JobStatus, error) {
	for {
		select {
		case tr := <-ch:
			if bytes.Equal(jobID[:], tr.JobId[:]) {
				fmt.Printf("(Job is %s)\n", metascheduler.JobStatus(tr.To))
				switch metascheduler.JobStatus(tr.To) {
				case metascheduler.JobStatusCancelled,
					metascheduler.JobStatusFailed,
					metascheduler.JobStatusFinished,
					metascheduler.JobStatusPanicked,
					metascheduler.JobStatusOutOfCredits,
					metascheduler.JobStatusRunning:
					return metascheduler.JobStatus(tr.To), nil
				}
			}
		case err := <-sub.Err():
			return metascheduler.JobStatusUnknown, err
		}
	}
}

var forbiddenReplacer = strings.NewReplacer(
	"\x1b[A", "", // Move Up
	"\x1b[B", "", // Move Down
	"\x1b[C", "", // Move Forward (Right)
	"\x1b[D", "", // Move Backward (Left)
	"\x1b[G", "", // Move to Beginning of Line
	"\x1b[H", "", // Move to Specific Position
	"\x1b[f", "", // Move to Specific Position (alternative)
	"\x1b[s", "", // Save Cursor Position
	"\x1b[u", "", // Restore Cursor Position
	"\r\n", "\n",
	"\r", "\n",
)
