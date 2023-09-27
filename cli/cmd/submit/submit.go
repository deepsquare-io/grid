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
Submit permits the submission of a job to the DeepSquare Grid.

USAGE:

	deepsquaretui submit [command options] <job.yaml>

OPTIONS:

Submit Settings:

	--affinities key<value [ --affinities key<value ]  Affinities flag. Used to filter the clusters. Format: key<value, `key<=value`, `key=value`, `key>=value`, `key>value`, `key!=value`


	--credits value                                    Allocated a number of credits. Unit is 1e18. Is a float and is not precise. (default: 0)


	--credits-wei value                                Allocated a number of credits. Unit is wei. Is a big int.


	--exit-on-job-exit, -e                             Exit the job after the job has finished and throw on error. (default: false)
	--job-name value                                   The job name.
	--no-timestamp, --no-ts                            Hide timestamp. (default: false)


	--uses key=value [ --uses key=value ]              Uses flag. Used to filter the clusters. Format: key=value


	--watch, -w                                        Watch logs after submitting the job (default: false)
*/
package submit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/sbatch"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var (
	sbatchEndpoint             string
	ethEndpointRPC             string
	ethEndpointWS              string
	ethHexPK                   string
	metaschedulerSmartContract string
	loggerEndpoint             string
	watch                      bool
	exitOnJobExit              bool
	noTimestamp                bool

	credits         *big.Int
	jobName         string
	uses            cli.StringSlice
	affinitiesSlice cli.StringSlice
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "metascheduler.rpc",
		Value:       deepsquare.DefaultRPCEndpoint,
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointRPC,
		Category:    "DeepSquare Settings:",
		EnvVars:     []string{"METASCHEDULER_RPC"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.ws",
		Value:       deepsquare.DefaultWSEndpoint,
		Usage:       "Metascheduler Avalanche C-Chain WS endpoint.",
		Destination: &ethEndpointWS,
		Category:    "DeepSquare Settings:",
		EnvVars:     []string{"METASCHEDULER_WS"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Required:    true,
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		Category:    "DeepSquare Settings:",
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "sbatch.endpoint",
		Value:       deepsquare.DefaultSBatchEndpoint,
		Usage:       "SBatch Service GraphQL endpoint.",
		Destination: &sbatchEndpoint,
		Category:    "DeepSquare Settings:",
		EnvVars:     []string{"SBATCH_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "logger.endpoint",
		Value:       deepsquare.DefaultLoggerEndpoint,
		Usage:       "Grid Logger endpoint.",
		Destination: &loggerEndpoint,
		Category:    "DeepSquare Settings:",
		EnvVars:     []string{"LOGGER_ENDPOINT"},
	},
	&cli.BoolFlag{
		Name:        "watch",
		Usage:       "Watch logs after submitting the job",
		Category:    "Submit Settings:",
		Aliases:     []string{"w"},
		Destination: &watch,
	},
	&cli.BoolFlag{
		Name:        "exit-on-job-exit",
		Usage:       "Exit the job after the job has finished and throw on error.",
		Aliases:     []string{"e"},
		Category:    "Submit Settings:",
		Destination: &exitOnJobExit,
	},
	&cli.BoolFlag{
		Name:        "no-timestamp",
		Usage:       "Hide timestamp.",
		Aliases:     []string{"no-ts"},
		Category:    "Submit Settings:",
		Destination: &noTimestamp,
	},
	&cli.StringFlag{
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		Category:    "DeepSquare Settings:",
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
	&cli.StringFlag{
		Name:        "job-name",
		Usage:       "The job name.",
		Required:    true,
		Destination: &jobName,
		Category:    "Submit Settings:",
	},
	&cli.StringSliceFlag{
		Name:        "uses",
		Usage:       "Uses flag. Used to filter the clusters. Format: `key=value`",
		Destination: &uses,
		Category:    "Submit Settings:",
	},
	&cli.StringSliceFlag{
		Name:        "affinities",
		Usage:       "Affinities flag. Used to filter the clusters. Format: `key<value`, `key<=value`, `key=value`, `key>=value`, `key>value`, `key!=value`",
		Destination: &affinitiesSlice,
		Category:    "Submit Settings:",
	},
	&cli.StringFlag{
		Name:  "credits-wei",
		Usage: "Allocated a number of credits. Unit is wei. Is a big int.",
		Action: func(ctx *cli.Context, s string) error {
			c, ok := new(big.Int).SetString(s, 10)
			if !ok {
				return errors.New("failed to parse credits")
			}
			credits = c
			return nil
		},
		Category: "Submit Settings:",
	},
	&cli.Float64Flag{
		Name:  "credits",
		Usage: "Allocated a number of credits. Unit is 1e18. Is a float and is not precise.",
		Action: func(ctx *cli.Context, val float64) error {
			c := new(big.Float).SetFloat64(val)

			coin := new(big.Float)
			coin.SetInt(big.NewInt(1000000000000000000))

			c.Mul(c, coin)

			result := new(big.Int)
			c.Int(result)
			credits = result
			return nil
		},
		Category: "Submit Settings:",
	},
}

var keyValueOperatorRegex = regexp.MustCompile(
	`^([^<>=!]+)\s*(==|<=|>=|<|>|!=|=)\s*([^<>=!]+)$`,
)

// parseKeyValueOperator parses a string in the format "key operator value" and returns the key, value, and operator.
func parseKeyValueOperator(input string) (key, value, op string, err error) {
	// Find the submatches in the input string.
	matches := keyValueOperatorRegex.FindStringSubmatch(input)

	// Check if the regex pattern matched.
	if len(matches) != 4 {
		return "", "", "", fmt.Errorf("invalid input format")
	}

	key = matches[1]
	op = matches[2]
	value = matches[3]

	return key, value, op, nil
}

var Command = cli.Command{
	Name:      "submit",
	Usage:     "Quickly submit a job.",
	ArgsUsage: "<job.yaml>",
	Flags:     flags,
	Action: func(cCtx *cli.Context) error {
		if cCtx.NArg() < 1 {
			return errors.New("missing arguments")
		}
		if credits == nil {
			return errors.New("missing --credits or --credits-wei parameter")
		}
		jobPath := cCtx.Args().First()
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
		dat, err := os.ReadFile(jobPath)
		if err != nil {
			return err
		}
		var job sbatch.Job
		if err := yaml.Unmarshal(dat, &job); err != nil {
			return err
		}

		var jobNameB [32]byte
		copy(jobNameB[:], jobName)

		// Map slices to label
		usesLabels := make([]types.Label, 0, len(uses.Value()))
		for _, use := range uses.Value() {
			if key, value, ok := strings.Cut(use, "="); ok {
				usesLabels = append(usesLabels, types.Label{
					Key:   key,
					Value: value,
				})
			}
		}

		// Map slices to affinites
		affinities := make([]types.Affinity, 0, len(affinitiesSlice.Value()))
		for _, affinity := range affinitiesSlice.Value() {
			if k, v, op, err := parseKeyValueOperator(affinity); err != nil {
				internallog.I.Error(
					"failed to parse",
					zap.String("affinity", affinity),
					zap.Error(err),
				)
				return err
			} else {
				var opB [2]byte
				copy(opB[:], op)
				affinities = append(affinities, types.Affinity{
					Label: metaschedulerabi.Label{
						Key:   k,
						Value: v,
					},
					Op: opB,
				})
			}
		}

		// Set allowance
		curr, err := client.GetAllowance(ctx)
		if err != nil {
			return err
		}
		if err = client.SetAllowance(ctx, curr.Add(curr, credits)); err != nil {
			return err
		}

		// Quick submit logic
		if !watch {
			jobID, err := client.SubmitJob(
				ctx,
				&job,
				credits,
				jobNameB,
				types.WithUse(usesLabels...),
				types.WithAffinity(affinities...),
			)
			if err != nil {
				return err
			}

			fmt.Printf("job %s submitted\n", hexutil.Encode(jobID[:]))
			return nil
		}

		// Watch submit logic
		transitions := make(chan types.JobTransition, 1)
		sub, err := watcher.SubscribeEvents(ctx, types.FilterJobTransition(transitions))
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()

		jobID, err := client.SubmitJob(ctx, &job, credits, jobNameB, types.WithUse(usesLabels...))
		if err != nil {
			return err
		}

		fmt.Println("---Waiting for job to be running...---")
		_, err = waitUntilJobRunningOrFinished(sub, transitions, jobID)
		if err != nil {
			fmt.Printf("---Watching transitions has unexpectedly closed---\n%s\n", err)
			return err
		}

		if exitOnJobExit {
			go func() {
				status, err := waitUntilJobFinished(sub, transitions, jobID)
				if err != nil {
					fmt.Printf("---Watching transitions has unexpectedly closed---\n%s\n", err)
					os.Exit(1)
				}
				switch status {
				case metascheduler.JobStatusFinished:
					os.Exit(0)
				case metascheduler.JobStatusCancelled:
					os.Exit(130)
				case metascheduler.JobStatusFailed, metascheduler.JobStatusPanicked:
					os.Exit(1)
				case metascheduler.JobStatusOutOfCredits:
					os.Exit(143)
				}
			}()
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
				fmt.Printf("---Connection to logging server closed unexpectedly---\n%s\n", err)
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
}

func waitUntilJobRunningOrFinished(
	sub ethereum.Subscription,
	ch <-chan types.JobTransition,
	jobID [32]byte,
) (metascheduler.JobStatus, error) {
	for {
		select {
		case tr := <-ch:
			if bytes.EqualFold(jobID[:], tr.JobId[:]) {
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

func waitUntilJobFinished(
	sub ethereum.Subscription,
	ch <-chan types.JobTransition,
	jobID [32]byte,
) (metascheduler.JobStatus, error) {
	for {
		select {
		case tr := <-ch:
			if bytes.EqualFold(jobID[:], tr.JobId[:]) {
				switch metascheduler.JobStatus(tr.To) {
				case metascheduler.JobStatusCancelled,
					metascheduler.JobStatusFailed,
					metascheduler.JobStatusFinished,
					metascheduler.JobStatusPanicked,
					metascheduler.JobStatusOutOfCredits:
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
