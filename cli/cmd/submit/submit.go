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
Package submit permits the submission of a job to the DeepSquare Grid.

USAGE:

	dps submit [command options] <job.yaml>

OPTIONS:

Submit Settings:

	--affinities key<value [ --affinities key<value ]  Affinities flag. Used to filter the clusters. Format: key<value, `key<=value`, `key=value`, `key>=value`, `key>value`, `key!=value` or `key:in:value`.


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
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/sbatch"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var (
	sbatchEndpoint              string
	ethEndpointRPC              string
	ethEndpointWS               string
	ethHexPK                    string
	metaschedulerSmartContract  string
	metaschedulerOracleEndpoint string
	loggerEndpoint              string
	watch                       bool
	exitOnJobExit               bool
	noTimestamp                 bool

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
		Name:        "metascheduler.oracle",
		Value:       metascheduler.DefaultOracleURL,
		Usage:       "Metascheduler Oracle endpoint.",
		Destination: &metaschedulerOracleEndpoint,
		EnvVars:     []string{"METASCHEDULER_ORACLE"},
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
		Usage:       "Affinities flag. Used to filter the clusters. Format: `key<value`, `key<=value`, `key=value`, `key>=value`, `key>value`, `key!=value` or `key:in:value`",
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
	`^([^<>=!in:]+)\s*(==|<=|>=|<|>|!=|=|:in:)\s*([^<>=!in:]+)$`,
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

// Command is the submit subcommand used to submit jobs.
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
			MetaschedulerAddress:        common.HexToAddress(metaschedulerSmartContract),
			RPCEndpoint:                 ethEndpointRPC,
			SBatchEndpoint:              sbatchEndpoint,
			LoggerEndpoint:              loggerEndpoint,
			MetaschedulerOracleEndpoint: metaschedulerOracleEndpoint,
			UserPrivateKey:              pk,
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
			k, v, op, err := parseKeyValueOperator(affinity)
			if err != nil {
				internallog.I.Error(
					"failed to parse",
					zap.String("affinity", affinity),
					zap.Error(err),
				)
				return err
			}

			var opB [2]byte
			switch op {
			case ":in:":
				opB = [2]byte{'i', 'n'}
			case "=", "==":
				opB = [2]byte{'=', '='}
			default:
				copy(opB[:], op)
			}
			affinities = append(affinities, types.Affinity{
				Label: metaschedulerabi.Label{
					Key:   k,
					Value: v,
				},
				Op: opB,
			})
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

			jobIDBig := new(big.Int).SetBytes(jobID[:])
			fmt.Printf("job %s submitted\n", jobIDBig.String())
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

		jobIDBig := new(big.Int).SetBytes(jobID[:])

		fmt.Printf("---Waiting for job %s to be running...---\n", jobIDBig.String())
		var finished = false
		var allocatedProviderAddress common.Address
		var provider types.ProviderDetail
		msOrSchedLen, runningLen := int64(0), int64(0)
		// Wait for finished or running
	loop:
		for {
			select {
			case tr := <-transitions:
				if (allocatedProviderAddress != common.Address{}) {
					jobs, err := client.GetJobsByProvider(ctx, allocatedProviderAddress)
					if err != nil {
						internallog.I.Warn("failed to fetch running jobs info", zap.Error(err))
					}
					msLen, rLen := reduceJobsIntoRunningOrScheduledLens(jobs)
					if len(jobs) > 1 && msOrSchedLen > 1 && (msOrSchedLen != msLen || runningLen != rLen) {
						waitingTime, err := computeWaitingTime(jobID, provider, jobs)
						if err != nil {
							internallog.I.Fatal("failed to compute waiting time", zap.Error(err))
						}
						fmt.Printf("(%d jobs in provider queue: %d waiting, %d running, wait ~%s)\n", len(jobs), msLen, rLen, waitingTime)
					}
					msOrSchedLen, runningLen = msLen, rLen
				}

				if bytes.Equal(jobID[:], tr.JobId[:]) {
					fmt.Printf("(Job is %s)\n", metascheduler.JobStatus(tr.To))
					switch metascheduler.JobStatus(tr.To) {
					case metascheduler.JobStatusMetaScheduled,
						metascheduler.JobStatusScheduled:
						// Print job position in the queue
						job, err := client.GetJob(ctx, jobID)
						if err != nil {
							internallog.I.Fatal("failed to fetch job info", zap.Error(err))
						}
						allocatedProviderAddress = job.ProviderAddr
						jobs, err := client.GetJobsByProvider(ctx, allocatedProviderAddress)
						if err != nil {
							internallog.I.Warn("failed to fetch running jobs info", zap.Error(err))
						}
						p, err := client.GetProvider(ctx, allocatedProviderAddress)
						if err != nil {
							internallog.I.Fatal("failed to get provider info", zap.Error(err))
						}
						provider = p
						msLen, rLen := reduceJobsIntoRunningOrScheduledLens(jobs)
						if len(jobs) > 1 && msOrSchedLen > 1 && (msOrSchedLen != msLen || runningLen != rLen) {
							waitingTime, err := computeWaitingTime(jobID, provider, jobs)
							if err != nil {
								internallog.I.Fatal("failed to compute waiting time", zap.Error(err))
							}
							fmt.Printf("(%d jobs in provider queue: %d waiting, %d running, wait ~%s)\n", len(jobs), msLen, rLen, waitingTime)
						}
						msOrSchedLen, runningLen = msLen, rLen

					case metascheduler.JobStatusCancelled,
						metascheduler.JobStatusFailed,
						metascheduler.JobStatusFinished,
						metascheduler.JobStatusPanicked,
						metascheduler.JobStatusOutOfCredits:
						finished = true
						break loop
					case metascheduler.JobStatusRunning:
						finished = false
						break loop
					}
				}
			case err := <-sub.Err():
				fmt.Printf("---Watching transitions has unexpectedly closed---\n%s\n", err)
				return err
			}
		}

		if exitOnJobExit {
			go func() {
				if !finished {
					cleanChan := make(chan os.Signal, 1)
					signal.Notify(cleanChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
					go func() {
						<-cleanChan
						fmt.Printf("\nWARNING: Your job %s is still running.\n", jobIDBig.String())
						os.Exit(1)
					}()

					_, err := waitUntilJobFinished(sub, transitions, jobID)
					if err != nil {
						fmt.Printf("---Watching transitions has unexpectedly closed---\n%s\n", err)
						os.Exit(1)
					}
				}

				job, err := client.GetJob(ctx, jobID)
				if err != nil {
					internallog.I.Fatal("failed to fetch job info", zap.Error(err))
				}
				os.Exit(int(job.ExitCode / 256))
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

func waitUntilJobFinished(
	sub ethereum.Subscription,
	ch <-chan types.JobTransition,
	jobID [32]byte,
) (metascheduler.JobStatus, error) {
	for {
		select {
		case tr := <-ch:
			if bytes.Equal(jobID[:], tr.JobId[:]) {
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

// computeWaitingTime returns min(running) + sum(waiting)
func computeWaitingTime(
	jobID [32]byte,
	provider types.ProviderDetail,
	jobs []types.Job,
) (time.Duration, error) {
	var waiting, running time.Duration
	for _, job := range jobs {
		if bytes.Equal(job.JobId[:], jobID[:]) {
			continue
		}
		switch metascheduler.JobStatus(job.Status) {
		case metascheduler.JobStatusRunning:
			durationB, err := metascheduler.CreditToDuration(
				provider.ProviderPrices,
				job.Definition,
				job.Cost.MaxCost,
			)
			if err != nil {
				return 0, err
			}
			startTime := time.Unix(job.Time.Start.Int64(), 0)
			duration := (time.Duration(durationB.Int64())*time.Second - time.Since(startTime)).Truncate(
				time.Second,
			)
			if running > duration || running == 0 {
				running = duration
			}

		case metascheduler.JobStatusMetaScheduled, metascheduler.JobStatusScheduled:
			durationB, err := metascheduler.CreditToDuration(
				provider.ProviderPrices,
				job.Definition,
				job.Cost.MaxCost,
			)
			if err != nil {
				return 0, err
			}
			waiting += time.Duration(durationB.Int64()) * time.Second
		}
	}
	return running + waiting, nil
}

func reduceJobsIntoRunningOrScheduledLens(
	jobs []types.Job,
) (metascheduledOrScheduled int64, running int64) {
	for _, job := range jobs {
		switch metascheduler.JobStatus(job.Status) {
		case metascheduler.JobStatusRunning:
			running++
		case metascheduler.JobStatusMetaScheduled, metascheduler.JobStatusScheduled:
			metascheduledOrScheduled++
		}
	}
	return
}
