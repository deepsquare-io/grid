package main

import (
	"context"
	"net/http"
	"os"
	"time"

	cryptotls "crypto/tls"
	"crypto/x509"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/debug"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/gridlogger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/watcher"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/middleware"
	pkgsbatch "github.com/deepsquare-io/the-grid/supervisor/pkg/sbatch"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/ssh"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	version       string = "dev"
	listenAddress string
	publicAddress string

	tls      bool
	keyFile  string
	certFile string

	sbatchEndpoint             string
	sbatchTLS                  bool
	sbatchTLSInsecure          bool
	sbatchCAFile               string
	ethEndpointRPC             string
	ethEndpointWS              string
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

	trace bool
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "grpc.listen-address",
		Value:       ":3000",
		Usage:       "Address to listen on. Is used for receiving job status via the job completion plugin.",
		Destination: &listenAddress,
		EnvVars:     []string{"LISTEN_ADDRESS"},
	},
	&cli.StringFlag{
		Name:        "public-address",
		Value:       "supervisor.example.com:3000",
		Usage:       "Public address or address of the reverse proxy. Is used by the SLURL plugins to know where to report job statuses.",
		Destination: &publicAddress,
		EnvVars:     []string{"PUBLIC_ADDRESS"},
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
		Name:        "metascheduler.endpoint.rpc",
		Value:       "https://testnet.deepsquare.run/rpc",
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointRPC,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT_RPC"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.endpoint.ws",
		Value:       "wss://testnet.deepsquare.run/ws",
		Usage:       "Metascheduler Avalanche C-Chain WS endpoint.",
		Destination: &ethEndpointWS,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT_WS"},
	},
	&cli.StringFlag{
		Name:        "sbatch.endpoint",
		Value:       "127.0.0.1:443",
		Usage:       "SBatch API gRPC endpoint.",
		Destination: &sbatchEndpoint,
		EnvVars:     []string{"SBATCH_ENDPOINT"},
	},
	&cli.BoolFlag{
		Name:        "sbatch.tls",
		Value:       true,
		Usage:       "Enable TLS for the SBatch API.",
		Destination: &sbatchTLS,
		EnvVars:     []string{"SBATCH_TLS_ENABLE"},
	},
	&cli.BoolFlag{
		Name:        "sbatch.tls.insecure",
		Value:       false,
		Usage:       "Skip TLS verification. By enabling it, sbatch.tls.ca and sbatch.tls.server-host-override are ignored.",
		Destination: &sbatchTLSInsecure,
		EnvVars:     []string{"SBATCH_TLS_INSECURE"},
	},
	&cli.StringFlag{
		Name:        "sbatch.tls.ca",
		Value:       "",
		Usage:       "Path to CA certificate for TLS verification.",
		Destination: &sbatchCAFile,
		EnvVars:     []string{"SBATCH_CA"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0x",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
	},
	&cli.StringFlag{
		Name:        "metascheduler.private-key",
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
		Value:       "/usr/bin/squeue",
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
		Usage:       "Total number of Nodes reported by 'scontrol show partitions'",
		Destination: &nodes,
		Required:    true,
		EnvVars:     []string{"TOTAL_NODES"},
	},
	&cli.Uint64Flag{
		Name:        "res.cpus",
		Usage:       "Total number of CPUs reported by 'scontrol show partitions'",
		Destination: &cpus,
		Required:    true,
		EnvVars:     []string{"TOTAL_CPUS"},
	},
	&cli.Uint64Flag{
		Name:        "res.gpus",
		Usage:       "Total number of GPUs reported by 'scontrol show partitions'",
		Destination: &gpus,
		Required:    true,
		EnvVars:     []string{"TOTAL_GPUS"},
	},
	&cli.Uint64Flag{
		Name:        "res.mem",
		Usage:       "Total number of Memory (MB) reported by 'scontrol show partitions'",
		Destination: &mem,
		Required:    true,
		EnvVars:     []string{"TOTAL_MEMORY"},
	},
	&cli.BoolFlag{
		Name:        "trace",
		Usage:       "Trace logging",
		Destination: &trace,
		EnvVars:     []string{"TRACE"},
	},
}

// Container stores the instances for dependency injection.
type Container struct {
	server        *server.Server
	sbatchAPI     pkgsbatch.Client
	metascheduler *metascheduler.Client
	slurm         *scheduler.Slurm
	jobWatcher    *watcher.Watcher
}

func Init(ctx context.Context) *Container {
	var err error
	var sbatchOpts []grpc.DialOption
	var tlsConfig = &cryptotls.Config{}
	if sbatchTLS {
		if !sbatchTLSInsecure {
			// Fetch the CA
			func() {
				b, err := os.ReadFile(sbatchCAFile)
				if err != nil {
					logger.I.Warn("failed to read TLS CA file", zap.Error(err))
					return
				}
				cp := x509.NewCertPool()
				if !cp.AppendCertsFromPEM(b) {
					logger.I.Warn("failed to append certificates", zap.Error(err))
					return
				}
				tlsConfig.RootCAs = cp
			}()
		}
		tlsConfig.InsecureSkipVerify = sbatchTLSInsecure
		sbatchOpts = append(
			sbatchOpts,
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		)
	}

	sbatchClient := pkgsbatch.NewClient(
		sbatchEndpoint,
		sbatchOpts...,
	)

	var client *http.Client
	if trace {
		client = &http.Client{
			Transport: &middleware.LoggingTransport{
				Transport: http.DefaultTransport,
			},
		}
	} else {
		client = http.DefaultClient
	}

	rpcClient, err := rpc.DialOptions(ctx, ethEndpointRPC, rpc.WithHTTPClient(client))
	if err != nil {
		logger.I.Fatal("ethclientRPC dial failed", zap.Error(err))
	}
	ethClientRPC := ethclient.NewClient(rpcClient)
	wsClient, err := rpc.DialOptions(ctx, ethEndpointWS, rpc.WithHTTPClient(client))
	if err != nil {
		logger.I.Fatal("ethclientWS dial failed", zap.Error(err))
	}
	ethClientWS := ethclient.NewClient(wsClient)
	pk, err := crypto.HexToECDSA(ethHexPK)
	if err != nil {
		logger.I.Fatal("couldn't decode private key", zap.Error(err))
	}
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		logger.I.Fatal("couldn't fetch chainID", zap.Error(err))
	}
	metascheduler := metascheduler.NewClient(
		chainID,
		common.HexToAddress(metaschedulerSmartContract),
		ethClientRPC,
		ethClientRPC,
		ethClientWS,
		pk,
	)
	sshService := ssh.New(
		slurmSSHAddress,
		slurmSSHB64PK,
	)
	slurmJobService := scheduler.NewSlurm(
		sshService,
		slurmSSHAdminUser,
		scancel,
		sbatch,
		squeue,
		scontrol,
		publicAddress,
	)
	resourceManager := lock.NewResourceManager()

	// TODO: do not hardcode this
	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		caCertPool = x509.NewCertPool()
	}
	gridLoggerOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&cryptotls.Config{
			RootCAs: caCertPool,
		})),
	}
	watcher := watcher.New(
		metascheduler,
		slurmJobService,
		sbatchClient,
		time.Duration(5*time.Second),
		resourceManager,
		gridlogger.NewDialer(
			gridLoggerOpts...,
		),
	)

	opts := []grpc.ServerOption{}
	if tls {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logger.I.Fatal("failed to load certificates", zap.Error(err))
		}
		opts = append(opts, grpc.Creds(creds))
	}
	server := server.New(
		metascheduler,
		resourceManager,
		slurmSSHB64PK,
		opts...,
	)

	return &Container{
		sbatchAPI:     sbatchClient,
		metascheduler: metascheduler,
		slurm:         slurmJobService,
		jobWatcher:    watcher,
		server:        server,
	}
}

var app = &cli.App{
	Name:    "supervisor",
	Version: version,
	Usage:   "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
	Flags:   flags,
	Action: func(cCtx *cli.Context) error {
		ctx := cCtx.Context
		container := Init(ctx)

		// Register the cluster with the declared resources
		// TODO: automatically fetch the resources limit
		if err := container.metascheduler.Register(
			ctx,
			nodes,
			cpus,
			gpus,
			mem,
		); err != nil {
			return err
		}

		go debug.WatchGoRoutines(ctx)

		go func(ctx context.Context) {
			if err := container.jobWatcher.Watch(ctx); err != nil {
				logger.I.Fatal("jobWatcher crashed", zap.Error(err))
			}
		}(ctx)

		logger.I.Info(
			"listening",
			zap.String("address", listenAddress),
			zap.String("version", version),
		)

		// gRPC server
		return container.server.ListenAndServe(listenAddress)
	},
}

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
