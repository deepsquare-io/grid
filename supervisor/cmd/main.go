package main

import (
	"os"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/deepsquare-io/the-grid/supervisor/server"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	listenAddress                string
	tls                          bool
	keyFile                      string
	certFile                     string
	metaschedulerGRPCEndpoint    string
	ethRPCEndpoint               string
	ethHexPK                     string
	metaschedulerSmartContract   string
	providerManagerSmartContract string
	slurmSSHAddress              string
	slurmSSHB64PK                string
	slurmSSHAdminUser            string
	scancel                      string
	sbatch                       string
	squeue                       string
	scontrol                     string
)

// Container stores the instances for dependency injection.
type Container struct {
	ethDataSource   *eth.DataSource
	slurmJobService *slurm.Service
}

func Init() *Container {
	ethDataSource := eth.NewDataSource(
		ethRPCEndpoint,
		metaschedulerGRPCEndpoint,
		providerManagerSmartContract,
		ethHexPK,
	)
	slurmJobService := slurm.New(
		slurmSSHAddress,
		slurmSSHB64PK,
		slurmSSHAdminUser,
		scancel,
		sbatch,
		squeue,
		scontrol,
	)

	return &Container{
		ethDataSource:   ethDataSource,
		slurmJobService: slurmJobService,
	}
}

func main() {

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
				Name:        "eth.metascheduler.endpoint",
				Value:       "https://mainnet.infura.io",
				Usage:       "Metascheduler RPC endpoint.",
				Destination: &ethRPCEndpoint,
				EnvVars:     []string{"METASCHEDULER_RPC_ENDPOINT"},
			},
			&cli.StringFlag{
				Name:        "grpc.metascheduler.endpoint",
				Value:       "127.0.0.1:443",
				Usage:       "Metascheduler gRPC endpoint.",
				Destination: &metaschedulerGRPCEndpoint,
				EnvVars:     []string{"METASCHEDULER_GRPC_ENDPOINT"},
			},
			&cli.StringFlag{
				Name:        "eth.metascheduler.smart-contract",
				Value:       "0x",
				Usage:       "Metascheduler smart-contract address.",
				Destination: &metaschedulerSmartContract,
				EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
			},
			&cli.StringFlag{
				Name:        "eth.provider-manager.smart-contract",
				Value:       "0x",
				Usage:       "Provider Manager smart-contract address.",
				Destination: &providerManagerSmartContract,
				EnvVars:     []string{"PROVIDER_MANAGER_SMART_CONTRACT"},
			},
			&cli.StringFlag{
				Name:        "eth.private-key",
				Usage:       "An hexadecimal private key for ethereum transactions.",
				Required:    true,
				Destination: &ethHexPK,
				EnvVars:     []string{"ETH_PRIVATE_KEY"},
			},
			&cli.StringFlag{
				Name:        "slurm.ssh.address",
				Value:       "127.0.0.1:22",
				Usage:       "Address of the Slurm login node.",
				Destination: &slurmSSHAddress,
				EnvVars:     []string{"SLURM_SSH_ADDRESS"},
			},
			&cli.StringFlag{
				Name:        "slurm.ssh.admin-user",
				Usage:       "SLURM admin user used for calling `scontrol` commands.",
				Required:    true,
				Destination: &slurmSSHAdminUser,
				EnvVars:     []string{"SLURM_SSH_ADMIN_USER"},
			},
			&cli.StringFlag{
				Name:        "slurm.ssh.private-key",
				Usage:       "Base64-encoded one line SSH private key used for impersonation. The public key must be inserted in the authorized_keys file of each user.",
				Required:    true,
				Destination: &slurmSSHB64PK,
				EnvVars:     []string{"SLURM_SSH_PRIVATE_KEY"},
			},
			&cli.StringFlag{
				Name:        "slurm.batch",
				Value:       "/usr/bin/sbatch",
				Usage:       "Server-side SLURM sbatch path.",
				Destination: &sbatch,
				EnvVars:     []string{"SLURM_SBATCH_PATH"},
			},
			&cli.StringFlag{
				Name:        "slurm.cancel",
				Value:       "/usr/bin/scancel",
				Usage:       "Server-side SLURM scancel path.",
				Destination: &scancel,
				EnvVars:     []string{"SLURM_SCANCEL_PATH"},
			},
		},
		Action: func(ctx *cli.Context) error {
			c := ctx.Context
			container := Init()
			// TODO: Need two loops, the grpc and the ethereum listener
			go job.Watch(c, container.ethDataSource, container.slurmJobService)

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
