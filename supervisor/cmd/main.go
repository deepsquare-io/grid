package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	cryptotls "crypto/tls"
	"crypto/x509"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
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

	scancel   string
	sbatch    string
	squeue    string
	scontrol  string
	sinfo     string
	nvidiaSMI string
	partition string

	benchmarkImage        string
	benchmarkSingleNode   bool
	benchmarkDisable      bool
	benchmarkRunAs        string
	benchmarkUnresponsive bool
	benchmarkTimeLimit    time.Duration

	trace bool
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "grpc.listen-address",
		Value:       ":3000",
		Usage:       "Address to listen on. Is used for receiving job status via the job completion plugin.",
		Destination: &listenAddress,
		EnvVars:     []string{"LISTEN_ADDRESS"},
		Category:    "Network:",
	},
	&cli.StringFlag{
		Name:        "public-address",
		Value:       "supervisor.example.com:3000",
		Usage:       "Public address or address of the reverse proxy. Is used by the SLURM plugins to know where to report job statuses. Must be protected with TLS.",
		Destination: &publicAddress,
		EnvVars:     []string{"PUBLIC_ADDRESS"},
		Category:    "Network:",
	},
	&cli.BoolFlag{
		Name:        "tls",
		Value:       false,
		Destination: &tls,
		Usage:       "Enable TLS for HTTP.",
		EnvVars:     []string{"TLS_ENABLE"},
		Category:    "Secure Transport:",
	},
	&cli.StringFlag{
		Name:        "tls.key-file",
		Value:       "",
		Destination: &keyFile,
		Usage:       "TLS Private Key file.",
		EnvVars:     []string{"TLS_KEY"},
		Category:    "Secure Transport:",
	},
	&cli.StringFlag{
		Name:        "tls.cert-file",
		Value:       "",
		Destination: &certFile,
		Usage:       "TLS Certificate file.",
		EnvVars:     []string{"TLS_CERT"},
		Category:    "Secure Transport:",
	},
	&cli.StringFlag{
		Name:        "sbatch.endpoint",
		Value:       "127.0.0.1:443",
		Usage:       "SBatch API gRPC endpoint.",
		Destination: &sbatchEndpoint,
		EnvVars:     []string{"SBATCH_ENDPOINT"},
		Category:    "SBatch API:",
	},
	&cli.BoolFlag{
		Name:        "sbatch.tls",
		Value:       true,
		Usage:       "Enable TLS for the SBatch API.",
		Destination: &sbatchTLS,
		EnvVars:     []string{"SBATCH_TLS_ENABLE"},
		Category:    "SBatch API:",
	},
	&cli.BoolFlag{
		Name:        "sbatch.tls.insecure",
		Value:       false,
		Usage:       "Skip TLS verification. By enabling it, sbatch.tls.ca is ignored.",
		Destination: &sbatchTLSInsecure,
		EnvVars:     []string{"SBATCH_TLS_INSECURE"},
		Category:    "SBatch API:",
	},
	&cli.StringFlag{
		Name:        "sbatch.tls.ca",
		Value:       "",
		Usage:       "Path to CA certificate for TLS verification.",
		Destination: &sbatchCAFile,
		EnvVars:     []string{"SBATCH_CA"},
		Category:    "SBatch API:",
	},
	&cli.StringFlag{
		Name:        "metascheduler.endpoint.rpc",
		Value:       "https://testnet.deepsquare.run/rpc",
		Usage:       "Metascheduler Avalanche C-Chain JSON-RPC endpoint.",
		Destination: &ethEndpointRPC,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT_RPC"},
		Category:    "MetaScheduler:",
	},
	&cli.StringFlag{
		Name:        "metascheduler.endpoint.ws",
		Value:       "wss://testnet.deepsquare.run/ws",
		Usage:       "Metascheduler Avalanche C-Chain WS endpoint.",
		Destination: &ethEndpointWS,
		EnvVars:     []string{"METASCHEDULER_ENDPOINT_WS"},
		Category:    "MetaScheduler:",
	},
	&cli.StringFlag{
		Name:        "metascheduler.smart-contract",
		Value:       "0x",
		Usage:       "Metascheduler smart-contract address.",
		Destination: &metaschedulerSmartContract,
		EnvVars:     []string{"METASCHEDULER_SMART_CONTRACT"},
		Category:    "MetaScheduler:",
	},
	&cli.StringFlag{
		Name:        "metascheduler.private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Required:    true,
		Destination: &ethHexPK,
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
		Category:    "MetaScheduler:",
	},
	&cli.StringFlag{
		Name:        "slurm.ssh.address",
		Usage:       "Address of the Slurm login node.",
		Required:    true,
		Destination: &slurmSSHAddress,
		EnvVars:     []string{"SLURM_SSH_ADDRESS"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "slurm.ssh.admin-user",
		Usage:       "SLURM admin user used for calling `scontrol` commands.",
		Required:    true,
		Destination: &slurmSSHAdminUser,
		EnvVars:     []string{"SLURM_SSH_ADMIN_USER"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "slurm.ssh.private-key",
		Usage:       "Base64-encoded one line SSH private key used for impersonation. The public key must be inserted in the authorized_keys file of each user.",
		Required:    true,
		Destination: &slurmSSHB64PK,
		EnvVars:     []string{"SLURM_SSH_PRIVATE_KEY"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "slurm.batch",
		Value:       "/usr/bin/sbatch",
		Usage:       "Server-side SLURM sbatch path.",
		Destination: &sbatch,
		EnvVars:     []string{"SLURM_SBATCH_PATH"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "slurm.cancel",
		Value:       "/usr/bin/scancel",
		Usage:       "Server-side SLURM scancel path.",
		Destination: &scancel,
		EnvVars:     []string{"SLURM_SCANCEL_PATH"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "slurm.squeue",
		Value:       "/usr/bin/squeue",
		Usage:       "Server-side SLURM squeue path.",
		Destination: &squeue,
		EnvVars:     []string{"SLURM_SQUEUE_PATH"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "slurm.control",
		Value:       "/usr/bin/scontrol",
		Usage:       "Server-side SLURM scontrol path.",
		Destination: &scontrol,
		EnvVars:     []string{"SLURM_SCONTROL_PATH"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "slurm.info",
		Value:       "/usr/bin/sinfo",
		Usage:       "Server-side SLURM info path.",
		Destination: &sinfo,
		EnvVars:     []string{"SLURM_SINFO_PATH"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name: "slurm.partition",
		Usage: `Slurm partition used for jobs and registering.

All the specifications returned by 'scontrol show partition' will be registered to the blockchain.`,
		Destination: &partition,
		Value:       "main",
		EnvVars:     []string{"SLURM_PARTITION"},
		Category:    "Slurm:",
	},
	&cli.StringFlag{
		Name:        "nvidia-smi",
		Value:       "nvidia-smi",
		Usage:       "Server-side nvidia-smi path.",
		Destination: &nvidiaSMI,
		EnvVars:     []string{"NVIDIA_SMI_PATH"},
		Category:    "Miscellaneous:",
	},
	&cli.StringFlag{
		Name:        "benchmark.image",
		Usage:       "Docker image used for benchmark",
		Destination: &benchmarkImage,
		Value:       "registry-1.deepsquare.run#library/hpc-benchmarks:21.4-hpl",
		EnvVars:     []string{"BENCHMARK_IMAGE"},
		Category:    "Benchmark:",
	},
	&cli.StringFlag{
		Name:        "benchmark.run-as",
		Usage:       "User used for benchmark",
		Destination: &benchmarkRunAs,
		Value:       "root",
		EnvVars:     []string{"BENCHMARK_RUN_AS"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "benchmark.single-node",
		Usage:       "Force single node benchmark.",
		Destination: &benchmarkSingleNode,
		Value:       false,
		EnvVars:     []string{"BENCHMARK_SINGLE_NODE"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "benchmark.include-unresponsive",
		Usage:       "Force benchmark on unresponsive nodes (sinfo --responding --partition=<partition>).",
		Destination: &benchmarkUnresponsive,
		Value:       false,
		EnvVars:     []string{"BENCHMARK_UNRESPONSIVE"},
		Category:    "Benchmark:",
	},
	&cli.DurationFlag{
		Name:        "benchmark.time-limit",
		Usage:       "Time limit (syntax is golang duration style).",
		Destination: &benchmarkTimeLimit,
		Value:       24 * time.Hour,
		EnvVars:     []string{"BENCHMARK_DURATION"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "benchmark.disable",
		Usage:       "Disable benchmark (and registering).",
		Destination: &benchmarkDisable,
		Value:       false,
		EnvVars:     []string{"BENCHMARK_DISABLE"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "trace",
		Usage:       "Trace logging",
		Destination: &trace,
		EnvVars:     []string{"TRACE"},
		Category:    "Miscellaneous:",
	},
}

// Container stores the instances for dependency injection.
type Container struct {
	server        *server.Server
	sbatchAPI     pkgsbatch.Client
	metascheduler metascheduler.MetaScheduler
	scheduler     scheduler.Scheduler
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
	metaScheduler := metascheduler.NewClient(
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
	slurmScheduler := scheduler.NewSlurm(
		sshService,
		slurmSSHAdminUser,
		publicAddress,
		"main",
		scheduler.WithSBatch(sbatch),
		scheduler.WithSCancel(scancel),
		scheduler.WithSControl(scontrol),
		scheduler.WithSQueue(squeue),
		scheduler.WithSInfo(sinfo),
		scheduler.WithNVidiaSMI(nvidiaSMI),
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
		metaScheduler,
		slurmScheduler,
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
	if err := slurmScheduler.HealthCheck(ctx); err != nil {
		logger.I.Fatal("failed to check slurm health", zap.Error(err))
	}
	server := server.New(
		metaScheduler,
		resourceManager,
		slurmSSHB64PK,
		opts...,
	)

	return &Container{
		sbatchAPI:     sbatchClient,
		metascheduler: metaScheduler,
		scheduler:     slurmScheduler,
		jobWatcher:    watcher,
		server:        server,
	}
}

var app = &cli.App{
	Name:    "supervisor",
	Version: version,
	Usage:   "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
	Flags:   flags,
	Action: func(cCtx *cli.Context) (err error) {
		ctx := cCtx.Context
		container := Init(ctx)

		go func() {
			logger.I.Info(
				"listening",
				zap.String("address", listenAddress),
				zap.String("version", version),
			)

			// gRPC server
			lis, err := net.Listen("tcp", listenAddress)
			if err != nil {
				logger.I.Fatal("failed to listen", zap.Error(err))
			}

			if err := container.server.Serve(lis); err != nil {
				logger.I.Fatal("http server crashed", zap.Error(err))
			}
		}()

		logger.I.Info("searching for cluster specs...")
		var benchmarkOpts []scheduler.FindSpecOption
		if !benchmarkUnresponsive {
			benchmarkOpts = append(benchmarkOpts, scheduler.WithOnlyResponding())
		}
		var benchmarkNodes uint64
		if benchmarkSingleNode {
			benchmarkNodes = 1
		} else {
			benchmarkNodes, err = container.scheduler.FindTotalNodes(ctx, benchmarkOpts...)
			if err != nil {
				logger.I.Fatal("failed to check total number of nodes", zap.Error(err))
			}
		}
		cpusPerNode, err := container.scheduler.FindCPUsPerNode(ctx, benchmarkOpts...)
		if err != nil {
			logger.I.Fatal("failed to check cpus per node", zap.Error(err))
		}
		gpusPerNode, err := container.scheduler.FindGPUsPerNode(ctx, benchmarkOpts...)
		if err != nil {
			logger.I.Fatal("failed to check gpus per node", zap.Error(err))
		}
		memPerNode, err := container.scheduler.FindMemPerNode(ctx, benchmarkOpts...)
		if err != nil {
			logger.I.Fatal("failed to check mem per node", zap.Error(err))
		}
		bl := benchmark.NewLauncher(
			benchmarkImage,
			benchmarkRunAs,
			publicAddress,
			container.scheduler,
			benchmarkNodes,
			cpusPerNode,
			memPerNode,
			gpusPerNode,
			benchmarkTimeLimit,
		)
		container.server.AddBenchmarkRoutes(
			container.metascheduler,
			container.scheduler,
			bl,
		)

		// Launch benchmark which will register the node
		if !benchmarkDisable {
			go func() {
				logger.I.Info("cancelling old benchmarks")
				if err := bl.Cancel(ctx); err != nil {
					logger.I.Warn("failed to cancel old benchmarks", zap.Error(err))
				}
				logger.I.Info("launching new benchmarks")

				if err := bl.RunPhase1(ctx); err != nil {
					logger.I.Fatal(
						"phase1 benchmark failed or failed to be tracked",
						zap.Error(err),
					)
				}
			}()
		} else {
			logger.I.Warn("benchmark disabled, will not register to the smart-contract")
		}

		return container.jobWatcher.Watch(ctx)
	},
}

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
