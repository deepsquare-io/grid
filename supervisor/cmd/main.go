package main

import (
	"os"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/server"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	var listenAddress string
	var tls bool
	var keyFile string
	var certFile string
	var metaschedulerAddress string
	var ethAddress string
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
				Name:        "metascheduler.address",
				Value:       "127.0.0.1:443",
				Usage:       "Metascheduler address.",
				Destination: &metaschedulerAddress,
				EnvVars:     []string{"METASCHEDULER_ADDRESS"},
			},
			&cli.StringFlag{
				Name:        "eth.address",
				Value:       "https://mainnet.infura.io",
				Usage:       "Ethereum net RPC address (can also be an IPC endpoint file).",
				Destination: &ethAddress,
				EnvVars:     []string{"ETH_ADDRESS"},
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
