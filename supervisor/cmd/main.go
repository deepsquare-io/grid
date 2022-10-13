package main

import (
	"context"
	"os"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/customer"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	listenAddress string

	tls      bool
	keyFile  string
	certFile string

	customerEndpoint           string
	customerTLS                bool
	customerTLSInsecure        bool
	customerCAFile             string
	customerServerHostOverride string
	ethEndpoint                string
	ethHexPK                   string
	metaschedulerSmartContract string

	slurmSSHAddress   string
	slurmSSHB64PK     string
	slurmSSHAdminUser string

	scancel  string
	sbatch   string
	squeue   string
	scontrol string

	nodes uint64
	cpus  uint64
	gpus  uint64
	mem   uint64
)

var flags = []cli.Flag{
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
		Name:        "metascheduler.endpoint",
		Value:       "https://testnet.deepsquare.run/rpc",
		Usage:       "Metascheduler RPC endpoint.",
		Destination: &ethEndpoint,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT"},
	},
	&cli.StringFlag{
		Name:        "customer.endpoint",
		Value:       "127.0.0.1:443",
		Usage:       "Oracle gRPC endpoint.",
		Destination: &customerEndpoint,
		EnvVars:     []string{"CUSTOMER_ENDPOINT"},
	},
	&cli.BoolFlag{
		Name:        "customer.tls",
		Value:       true,
		Usage:       "Enable TLS for the Customer API.",
		Destination: &customerTLS,
		EnvVars:     []string{"CUSTOMER_TLS_ENABLE"},
	},
	&cli.BoolFlag{
		Name:        "customer.tls.insecure",
		Value:       false,
		Usage:       "Skip TLS verification. By enabling it, customer.tls.ca and customer.tls.server-host-override are ignored.",
		Destination: &customerTLSInsecure,
		EnvVars:     []string{"CUSTOMER_TLS_INSECURE"},
	},
	&cli.StringFlag{
		Name:        "customer.tls.ca",
		Value:       "",
		Usage:       "Path to CA certificate for TLS verification.",
		Destination: &customerCAFile,
		EnvVars:     []string{"CUSTOMER_CA"},
	},
	&cli.StringFlag{
		Name:        "customer.tls.server-host-override",
		Value:       "customer.deepsquare.io",
		Usage:       "The server name used to verify the hostname returned by the TLS handshake.",
		Destination: &customerServerHostOverride,
		EnvVars:     []string{"CUSTOMER_SERVER_HOST_OVERRIDE"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0x",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
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
		Usage:       "Address of the Slurm login node.",
		Required:    true,
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
	&cli.StringFlag{
		Name:        "slurm.squeue",
		Value:       "/usr/bin/ssqueue",
		Usage:       "Server-side SLURM squeue path.",
		Destination: &squeue,
		EnvVars:     []string{"SLURM_SQUEUE_PATH"},
	},
	&cli.StringFlag{
		Name:        "slurm.control",
		Value:       "/usr/bin/scontrol",
		Usage:       "Server-side SLURM scontrol path.",
		Destination: &scontrol,
		EnvVars:     []string{"SLURM_SCONTROL_PATH"},
	},
	&cli.Uint64Flag{
		Name:        "res.nodes",
		Usage:       "Total number of Nodes reported by `scontrol show partitions`",
		Destination: &nodes,
		Required:    true,
		EnvVars:     []string{"TOTAL_NODES"},
	},
	&cli.Uint64Flag{
		Name:        "res.cpus",
		Usage:       "Total number of CPUs reported by `scontrol show partitions`",
		Destination: &cpus,
		Required:    true,
		EnvVars:     []string{"TOTAL_CPUS"},
	},
	&cli.Uint64Flag{
		Name:        "res.gpus",
		Usage:       "Total number of GPUs reported by `scontrol show partitions`",
		Destination: &gpus,
		Required:    true,
		EnvVars:     []string{"TOTAL_GPUS"},
	},
	&cli.Uint64Flag{
		Name:        "res.mem",
		Usage:       "Total number of Memory (MB) reported by `scontrol show partitions`",
		Destination: &mem,
		Required:    true,
		EnvVars:     []string{"TOTAL_MEMORY"},
	},
}

// Container stores the instances for dependency injection.
type Container struct {
	server     *server.Server
	customer   *customer.FakeDataSource
	eth        *eth.DataSource
	slurm      *slurm.Service
	jobWatcher *job.Watcher
}

func Init() *Container {
	customer := &customer.FakeDataSource{}
	ethDataSource := eth.New(
		ethEndpoint,
		ethHexPK,
		metaschedulerSmartContract,
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
	watcher := job.New(
		ethDataSource,
		slurmJobService,
		customer,
	)
	server := server.New(
		tls,
		keyFile,
		certFile,
		ethDataSource,
		slurmSSHB64PK,
	)

	return &Container{
		customer:   customer,
		eth:        ethDataSource,
		slurm:      slurmJobService,
		jobWatcher: watcher,
		server:     server,
	}
}

var app = &cli.App{
	Name:  "supervisor",
	Usage: "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
	Flags: flags,
	Action: func(ctx *cli.Context) error {
		c := ctx.Context
		container := Init()

		// Register the cluster with the declared resources
		// TODO: automatically fetch the resources limit
		if err := container.eth.Register(
			c,
			nodes,
			cpus,
			gpus,
			mem,
		); err != nil {
			return err
		}

		go func(c context.Context) {
			if err := container.jobWatcher.Watch(c); err != nil {
				logger.I.Fatal("app crashed", zap.Error(err))
			}
		}(c)

		logger.I.Info("listening", zap.String("address", listenAddress))

		// gRPC server
		return container.server.ListenAndServe(listenAddress)
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
