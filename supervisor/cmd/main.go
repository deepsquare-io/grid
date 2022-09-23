package main

import (
	"context"
	"os"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/deepsquare-io/the-grid/supervisor/server"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	var listenAddress string
	var tls bool
	var keyFile string
	var certFile string
	var metaschedulerGRPCEndpoint string
	var ethRPCEndpoint string
	var ethSmartContract string
	var sshAddress string

	app := &cli.App{
		Name:  "supervisor",
		Usage: "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "grpc.listen-address",
				Value:       ":3000",
				Usage:       "Address to listen on. Is used for receiving job status via the job completion plugin.",
				Destination: &listenAddress,
				EnvVars:     []string{"LISTEN_ADDRESS"},
			},
			&cli.BoolFlag{
				Name:        "tls",
				Value:       false,
				Destination: &tls,
				Usage:       "Enable TLS for GRPC.",
				EnvVars:     []string{"TLS_ENABLE"},
			},
			&cli.StringFlag{
				Name:        "tls.key-file",
				Value:       "",
				Destination: &keyFile,
				Usage:       "TLS Private Key file.",
				EnvVars:     []string{"TLS_KEY"},
			},
			&cli.StringFlag{
				Name:        "tls.cert-file",
				Value:       "",
				Destination: &certFile,
				Usage:       "TLS Certificate file.",
				EnvVars:     []string{"TLS_CERT"},
			},
			&cli.StringFlag{
				Name:        "metascheduler.eth.endpoint",
				Value:       "https://mainnet.infura.io",
				Usage:       "Metascheduler RPC endpoint.",
				Destination: &ethRPCEndpoint,
				EnvVars:     []string{"METASCHEDULER_RPC_ENDPOINT"},
			},
			&cli.StringFlag{
				Name:        "metascheduler.grpc.endpoint",
				Value:       "127.0.0.1:443",
				Usage:       "Metascheduler gRPC endpoint.",
				Destination: &metaschedulerGRPCEndpoint,
				EnvVars:     []string{"METASCHEDULER_GRPC_ENDPOINT"},
			},
			&cli.StringFlag{
				Name:        "metascheduler.eth.smart-contract",
				Value:       "0x",
				Usage:       "Metascheduler smart-contract address.",
				Destination: &ethSmartContract,
				EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
			},
			&cli.StringFlag{
				Name:        "slurm.ssh.address",
				Value:       "127.0.0.1:22",
				Usage:       "Address of the Slurm login node.",
				Destination: &sshAddress,
				EnvVars:     []string{"SLURM_SSH_ADDRESS"},
			},
			&cli.StringFlag{
				Name:     "slurm.ssh.private-key",
				Usage:    "Base64-encoded SSH private key used for impersonation. The public key must be inserted in the authorized_keys file of each user.",
				Required: true,
				EnvVars:  []string{"SLURM_SSH_PRIVATE_KEY"},
			},
			&cli.StringFlag{
				Name:    "slurm.batch",
				Value:   "/usr/bin/sbatch",
				Usage:   "Server-side SLURM sbatch path.",
				EnvVars: []string{"SLURM_SBATCH_PATH"},
			},
			&cli.StringFlag{
				Name:    "slurm.cancel",
				Value:   "/usr/bin/scancel",
				Usage:   "Server-side SLURM scancel path.",
				EnvVars: []string{"SLURM_SCANCEL_PATH"},
			},
		},
		Action: func(ctx *cli.Context) error {
			// TODO: Need two loops, the grpc and the ethereum listener

			// gRPC server
			if tls {
				return server.ListenAndServeTLS(listenAddress, keyFile, certFile)
			} else {
				return server.ListenAndServe(listenAddress)
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}

func WatchQueue(e eth.DataSource, s slurm.JobService) error {
	ctx := context.Background()
	resp := make(chan eth.ClaimJobResponse)
	done := make(chan error)
	for {
		func() {
			ctx, cancel := context.WithTimeout(ctx, time.Duration(60*time.Second))
			defer cancel()

			go e.ClaimJob(resp, done)

			select {
			case r := <-resp:
				// TODO: fetch sbatch here
				body := `#!/bin/sh

				srun hostname
				srun sleep infinity
				`
				job := eth.JobDefinitionMapToSlurm(r.JobDefinition, r.TimeLimit, body)
				req := &slurm.SubmitJobRequest{
					Name:          r.JobID,
					User:          r.User,
					JobDefinition: job,
				}
				slurmJobId, err := s.SubmitJob(req)
				if err != nil {
					logger.I.Error("slurm submit job failed", zap.Error(err))
				} else {
					logger.I.Info(
						"submitted a job successfully",
						zap.Int("JobID", slurmJobId),
						zap.Any("Req", req),
					)
				}
			case err := <-done:
				if err != nil {
					logger.I.Error("claimJob failed", zap.Error(err))
				}

			case <-ctx.Done():
				logger.I.Error("claimJob timed out", zap.Error(ctx.Err()))
			}
		}()

		// TODO: extract variable
		time.Sleep(time.Duration(10 * time.Second))
	}
}
