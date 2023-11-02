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

package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	cryptotls "crypto/tls"
	"crypto/x509"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/grid/supervisor/pkg/gc"
	"github.com/deepsquare-io/grid/supervisor/pkg/gridlogger"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/watcher"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/middleware"
	pkgsbatch "github.com/deepsquare-io/grid/supervisor/pkg/sbatch"
	"github.com/deepsquare-io/grid/supervisor/pkg/server"
	"github.com/deepsquare-io/grid/supervisor/pkg/ssh"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils/try"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	version       string = "dev"
	listenAddress string
	publicAddress string

	commonBenchmarkOpts = []benchmark.Option{}

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

	cpuPricePerMin *big.Int
	gpuPricePerMin *big.Int
	memPricePerMin *big.Int
	labels         []metaschedulerabi.Label

	benchmarkHPLImage       string
	benchmarkHPLSingleNode  bool
	benchmarkSpeedTestImage string
	benchmarkOSUImage       string
	benchmarkIORImage       string
	benchmarkIORSingleNode  bool
	benchmarkDisable        bool
	benchmarkRunAs          string
	benchmarkUnresponsive   bool
	benchmarkTimeLimit      time.Duration
	benchmarkTrace          bool

	benchmarkUCX          bool
	benchmarkUCXAffinity  string
	benchmarkUCXTransport string

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
		Value:       "0x3707aB457CF457275b7ec32e203c54df80C299d5",
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
		Name:       "metascheduler.gpu-price-per-min",
		Usage:      "Price of the GPU per min. Reference for a rtx3090: 8500000000000000000 (8.5 creds/(CPU.min)).",
		HasBeenSet: true,
		Value:      "8500000000000000000",
		Action: func(ctx *cli.Context, s string) error {
			val, ok := new(big.Int).SetString(s, 10)
			if !ok {
				return errors.New("invalid gpu-price-per-min value")
			}
			gpuPricePerMin = val
			return nil
		},
		EnvVars:  []string{"METASCHEDULER_GPU_PRICE_PER_MIN"},
		Category: "MetaScheduler:",
	},
	&cli.StringFlag{
		Name:       "metascheduler.cpu-price-per-min",
		Usage:      "Price of the CPU per min. Reference for a zen2: 950000000000000000 (0.95 creds/(GPU.min)).",
		HasBeenSet: true,
		Value:      "950000000000000000",
		Action: func(ctx *cli.Context, s string) error {
			val, ok := new(big.Int).SetString(s, 10)
			if !ok {
				return errors.New("invalid cpu-price-per-min value")
			}
			cpuPricePerMin = val
			return nil
		},
		EnvVars:  []string{"METASCHEDULER_CPU_PRICE_PER_MIN"},
		Category: "MetaScheduler:",
	},
	&cli.StringFlag{
		Name:       "metascheduler.mem-price-per-min",
		Usage:      "Price of the Mem (MB) per min. Reference: 80000000000000 (0.00008 creds/(MB.min)).",
		HasBeenSet: true,
		Value:      "80000000000000",
		Action: func(ctx *cli.Context, s string) error {
			val, ok := new(big.Int).SetString(s, 10)
			if !ok {
				return errors.New("invalid mem-price-per-min value")
			}
			memPricePerMin = val
			return nil
		},
		EnvVars:  []string{"METASCHEDULER_MEM_PRICE_PER_MIN"},
		Category: "MetaScheduler:",
	},
	&cli.StringSliceFlag{
		Name:  "metascheduler.label",
		Usage: "Additional `key=value` label for registration\n(recommended os=linux,arch=amd64,gpu=rtx3090,cpu=amd-epyc-7302,name=my-cluster,zone=fr-paris-1,region=fr-paris).",
		Action: func(ctx *cli.Context, slabels []string) error {
			labels = make([]metaschedulerabi.Label, 0, len(slabels))
			for _, sl := range slabels {
				k, v, ok := strings.Cut(sl, "=")
				if !ok {
					return fmt.Errorf("label %q is missing a '=' key-value separator", sl)
				}
				labels = append(labels, metaschedulerabi.Label{
					Key:   k,
					Value: v,
				})
			}

			labels = metascheduler.ProcessLabels(labels)
			return nil
		},
		Aliases:  []string{"l", "label"},
		EnvVars:  []string{"METASCHEDULER_LABELS"},
		Category: "MetaScheduler:",
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
		Name:        "benchmark.speedtest.image",
		Usage:       "Docker image used for SpeedTest benchmark",
		Destination: &benchmarkSpeedTestImage,
		Value:       benchmark.DefaultSpeedTestImage,
		EnvVars:     []string{"BENCHMARK_SPEEDTEST_IMAGE"},
		Category:    "Benchmark:",
	},
	&cli.StringFlag{
		Name:        "benchmark.osu.image",
		Usage:       "Docker image used for OSU benchmark",
		Destination: &benchmarkOSUImage,
		Value:       benchmark.DefaultOSUImage,
		EnvVars:     []string{"BENCHMARK_OSU_IMAGE"},
		Category:    "Benchmark:",
	},
	&cli.StringFlag{
		Name:        "benchmark.ior.image",
		Usage:       "Docker image used for IOR benchmark",
		Destination: &benchmarkIORImage,
		Value:       benchmark.DefaultIORImage,
		EnvVars:     []string{"BENCHMARK_IOR_IMAGE"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "benchmark.ior.single-node",
		Usage:       "Force single node benchmark for IOR.",
		Destination: &benchmarkIORSingleNode,
		Value:       false,
		EnvVars:     []string{"BENCHMARK_IOR_SINGLE_NODE"},
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
	&cli.StringFlag{
		Name:        "benchmark.hpl.image",
		Usage:       "Docker image used for HPL benchmark",
		Destination: &benchmarkHPLImage,
		Value:       benchmark.DefaultHPLImage,
		EnvVars:     []string{"BENCHMARK_HPL_IMAGE"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "benchmark.hpl.single-node",
		Usage:       "Force single node benchmark for HPL.",
		Destination: &benchmarkHPLSingleNode,
		Value:       false,
		EnvVars:     []string{"BENCHMARK_HPL_SINGLE_NODE"},
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
		Value:       benchmark.DefaultTimeLimit,
		EnvVars:     []string{"BENCHMARK_TIME_LIMIT"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "benchmark.ucx",
		Usage:       "Use UCX transport for MPI. Choose this for RDMA. Do not for TCP.",
		Destination: &benchmarkUCX,
		Value:       false,
		EnvVars:     []string{"BENCHMARK_UCX"},
		Category:    "Benchmark:",
	},
	&cli.StringFlag{
		Name: "benchmark.ucx.affinity",
		Usage: `UCX Affinity for each node. Select the network devices with the format devices_for_node_1|devices_for_node_2|...

See 'ucx_info -bd' to see available devices.

Examples:
  mlx5_0:1|mlx5_0:1 means that cn1 will use mlx5_0 port 1 and cn2 will use mlx5_0 port 1
  mlx5_0:1,mlx5_0:1|mlx5_0:1 means that cn1 will use mlx5_0 port 1 or mlx5_0 port 1, and cn2 will use mlx5_0 port 1

`,

		Destination: &benchmarkUCXAffinity,
		Required:    false,
		EnvVars:     []string{"BENCHMARK_UCX_AFFINITY"},
		Category:    "Benchmark:",
	},
	&cli.StringFlag{
		Name: "benchmark.ucx.transport",
		Usage: `UCX Tranport. Select the common transport.

See 'ucx_info -bd' to see available tranports.

Value is often: sm,self,rc (shared memory, self, rdma reliable connected). Set to empty to set automatically.

Note that TCP is not supported at the moment.

`,

		Destination: &benchmarkUCXTransport,
		EnvVars:     []string{"BENCHMARK_UCX_TRANSPORT"},
		Category:    "Benchmark:",
	},
	&cli.BoolFlag{
		Name:        "benchmark.trace",
		Usage:       `Enables benchmark trace logging. Very verbose.`,
		Destination: &benchmarkTrace,
		EnvVars:     []string{"BENCHMARK_TRACE"},
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
	server            *http.Server
	sbatchAPI         pkgsbatch.Client
	metascheduler     metascheduler.MetaScheduler
	scheduler         scheduler.Scheduler
	jobWatcher        *watcher.Watcher
	gc                *gc.GC
	benchmarkLauncher benchmark.Launcher
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
	bl := benchmark.NewLauncher(
		benchmarkRunAs,
		publicAddress,
		slurmScheduler,
		benchmark.WithTimeLimit(benchmarkTimeLimit),
	)
	hplOpts := append(
		commonBenchmarkOpts,
		benchmark.WithImage(benchmarkHPLImage),
	)
	server := server.New(
		metaScheduler,
		resourceManager,
		bl,
		slurmSSHB64PK,
		hplOpts,
		opts...,
	)
	gc := gc.NewGC(metaScheduler, slurmScheduler)

	return &Container{
		sbatchAPI:         sbatchClient,
		metascheduler:     metaScheduler,
		scheduler:         slurmScheduler,
		jobWatcher:        watcher,
		benchmarkLauncher: bl,
		server:            server,
		gc:                gc,
	}
}

var app = &cli.App{
	Name:                 "supervisor",
	Version:              version,
	Usage:                "Overwatch the job scheduling and register the compute to the Deepsquare Grid.",
	Flags:                flags,
	Suggest:              true,
	EnableBashCompletion: true,
	Copyright: `supervisor  Copyright (C) 2023 DeepSquare Association
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.
See the GNU General Public License for more details.`,
	Action: func(cCtx *cli.Context) (err error) {
		ctx := cCtx.Context
		commonBenchmarkOpts = append(
			commonBenchmarkOpts,
			benchmark.WithSupervisorPublicAddress(publicAddress),
		)
		if benchmarkUCX {
			commonBenchmarkOpts = append(
				commonBenchmarkOpts,
				benchmark.WithUCX(benchmarkUCXAffinity, benchmarkUCXTransport),
			)
		}
		if benchmarkTrace {
			commonBenchmarkOpts = append(
				commonBenchmarkOpts,
				benchmark.WithTrace(),
			)
		}
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

		// Launch benchmark which will register the node
		if !benchmarkDisable {
			go func() {
				logger.I.Info("initial slurm healthcheck...")
				if err := try.Do(10, 10*time.Second, func(try int) error {
					return container.scheduler.HealthCheck(ctx)
				}); err != nil {
					logger.I.Fatal("healthcheck failed", zap.Error(err))
				}

				logger.I.Info("cancelling old benchmarks")
				if err := container.benchmarkLauncher.Cancel(ctx, "osu"); err != nil {
					logger.I.Warn("failed to cancel osu benchmark", zap.Error(err))
				}
				if err := container.benchmarkLauncher.Cancel(ctx, "speedtest"); err != nil {
					logger.I.Warn("failed to cancel speedtest benchmark", zap.Error(err))
				}
				if err := container.benchmarkLauncher.Cancel(ctx, "hpl-phase1"); err != nil {
					logger.I.Warn("failed to cancel hpl-phase1 benchmark", zap.Error(err))
				}
				if err := container.benchmarkLauncher.Cancel(ctx, "ior"); err != nil {
					logger.I.Warn("failed to cancel ior benchmark", zap.Error(err))
				}
				logger.I.Info("launching new benchmarks")

				logger.I.Info("searching for cluster specs...")
				var findOpts []scheduler.FindSpecOption
				if !benchmarkUnresponsive {
					findOpts = append(findOpts, scheduler.WithOnlyResponding())
				}
				nodes, err := container.scheduler.FindTotalNodes(ctx, findOpts...)
				if err != nil {
					logger.I.Fatal("failed to check total number of nodes", zap.Error(err))
				}
				if nodes == 0 {
					logger.I.Fatal("no nodes available, check sinfo")
				}
				var hplNodes uint64
				if benchmarkHPLSingleNode {
					hplNodes = 1
				} else {
					hplNodes = nodes
				}
				var iorNodes uint64
				if benchmarkIORSingleNode {
					iorNodes = 1
				} else {
					iorNodes = nodes
				}
				cpusPerNode, err := container.scheduler.FindCPUsPerNode(ctx, findOpts...)
				if err != nil {
					logger.I.Fatal("failed to check cpus per node", zap.Error(err))
				}
				var minCPUsPerNode uint64
				if len(cpusPerNode) == 0 {
					minCPUsPerNode = 0
				} else {
					minCPUsPerNode = slices.Min(cpusPerNode)
				}
				gpusPerNode, err := container.scheduler.FindGPUsPerNode(ctx, findOpts...)
				if err != nil {
					logger.I.Fatal("failed to check gpus per node", zap.Error(err))
				}
				var minGPUsPerNode uint64
				if len(gpusPerNode) == 0 {
					minGPUsPerNode = 0
				} else {
					minGPUsPerNode = slices.Min(gpusPerNode)
				}
				memPerNode, err := container.scheduler.FindMemPerNode(ctx, findOpts...)
				if err != nil {
					logger.I.Fatal("failed to check mem per node", zap.Error(err))
				}
				var minMemPerNode uint64
				if len(memPerNode) == 0 {
					minMemPerNode = 0
				} else {
					minMemPerNode = slices.Min(memPerNode)
				}

				hplOpts := append(
					commonBenchmarkOpts,
					benchmark.WithClusterSpecs(
						hplNodes,
						minCPUsPerNode,
						minGPUsPerNode,
						minMemPerNode,
					),
					benchmark.WithImage(benchmarkHPLImage),
				)

				osuOpts := append(
					commonBenchmarkOpts,
					benchmark.WithClusterSpecs(
						nodes,
						minCPUsPerNode,
						minGPUsPerNode,
						minMemPerNode,
					),
					benchmark.WithImage(benchmarkOSUImage),
				)

				iorOpts := append(
					commonBenchmarkOpts,
					benchmark.WithClusterSpecs(
						iorNodes,
						minCPUsPerNode,
						minGPUsPerNode,
						minMemPerNode,
					),
					benchmark.WithImage(benchmarkIORImage),
				)

				speedtestOpts := append(
					commonBenchmarkOpts,
					benchmark.WithClusterSpecs(
						nodes,
						minCPUsPerNode,
						minGPUsPerNode,
						minMemPerNode,
					),
					benchmark.WithImage(benchmarkSpeedTestImage),
				)

				logger.I.Info("checking if it is necessary to re-run the benchmark")
				oldInfo, err := container.metascheduler.GetOldInfo(ctx)
				if err != nil {
					logger.I.Warn("failed to fetch old info, running benchmark...", zap.Error(err))
				}
				hardware := metaschedulerabi.ProviderHardware{
					Nodes:       nodes,
					GpusPerNode: gpusPerNode,
					CpusPerNode: cpusPerNode,
					MemPerNode:  memPerNode,
				}
				prices := metaschedulerabi.ProviderPrices{
					GpuPricePerMin: gpuPricePerMin,
					CpuPricePerMin: cpuPricePerMin,
					MemPricePerMin: memPricePerMin,
				}
				if metascheduler.ProviderHardwareEqual(oldInfo.ProviderHardware, hardware) &&
					oldInfo.ProviderPrices == prices &&
					metascheduler.LabelsContains(labels, oldInfo.Labels) {
					logger.I.Info(
						"hardware, prices and labels are the same, no need to run a benchmark",
						zap.Any("info", oldInfo),
					)
					// Add old benchmark labels
					labels = oldInfo.Labels
				} else {
					logger.I.Info(
						"need to run a new benchmark",
						zap.Any("oldInfo", oldInfo),
						zap.Any("expectedHardware", hardware),
						zap.Any("expectedPrices", prices),
						zap.Any("expectedLabels", labels),
					)
					launchBenchmarks(
						ctx,
						container.benchmarkLauncher,
						hplOpts,
						osuOpts,
						speedtestOpts,
						iorOpts,
					)

					result := benchmark.DefaultStore.Dump()

					labels = metascheduler.MergeLabels(labels, []metaschedulerabi.Label{
						{
							Key:   "cpu",
							Value: result.MachineSpec.CPU,
						},
						{
							Key:   "cpu.microarch",
							Value: result.MachineSpec.MicroArch,
						},
						{
							Key:   "os",
							Value: result.MachineSpec.OS,
						},
						{
							Key:   "arch",
							Value: result.MachineSpec.Arch,
						},
						{
							Key:   "gpu",
							Value: result.MachineSpec.GPU,
						},
						{
							Key:   "compute.gflops",
							Value: fmt.Sprintf("%.2f", result.GFLOPS),
						},
						{
							Key:   "network.upload.bw.mbps",
							Value: fmt.Sprintf("%.2f", float64(result.UploadBandwidth)/1e6),
						},
						{
							Key:   "network.download.bw.mbps",
							Value: fmt.Sprintf("%.2f", float64(result.DownloadBandwidth)/1e6),
						},
						{
							Key:   "network.p2p.bw.mbps",
							Value: fmt.Sprintf("%.2f", result.P2PBidirectionalBandwidth),
						},
						{
							Key:   "network.p2p.latency.us",
							Value: fmt.Sprintf("%.2f", result.P2PLatency),
						},
						{
							Key:   "network.all-to-all.latency.us",
							Value: fmt.Sprintf("%.2f", result.AllToAllCollectiveLatency),
						},
						{
							Key:   "storage.scratch.read.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.ScratchAvgRead.Bandwidth),
						},
						{
							Key:   "storage.scratch.read.iops",
							Value: fmt.Sprintf("%.2f", result.ScratchAvgRead.IOPS),
						},
						{
							Key:   "storage.scratch.write.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.ScratchAvgWrite.Bandwidth),
						},
						{
							Key:   "storage.scratch.write.iops",
							Value: fmt.Sprintf("%.2f", result.ScratchAvgWrite.IOPS),
						},
						{
							Key:   "storage.shared-world-tmp.read.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.SharedWorldTmpAvgRead.Bandwidth),
						},
						{
							Key:   "storage.shared-world-tmp.read.iops",
							Value: fmt.Sprintf("%.2f", result.SharedWorldTmpAvgRead.IOPS),
						},
						{
							Key:   "storage.shared-world-tmp.write.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.SharedWorldTmpAvgWrite.Bandwidth),
						},
						{
							Key:   "storage.shared-world-tmp.write.iops",
							Value: fmt.Sprintf("%.2f", result.SharedWorldTmpAvgWrite.IOPS),
						},
						{
							Key:   "storage.shared-tmp.read.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.SharedTmpAvgRead.Bandwidth),
						},
						{
							Key:   "storage.shared-tmp.read.iops",
							Value: fmt.Sprintf("%.2f", result.SharedTmpAvgRead.IOPS),
						},
						{
							Key:   "storage.shared-tmp.write.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.SharedTmpAvgWrite.Bandwidth),
						},
						{
							Key:   "storage.shared-tmp.write.iops",
							Value: fmt.Sprintf("%.2f", result.SharedTmpAvgWrite.IOPS),
						},
						{
							Key:   "storage.disk-world-tmp.read.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.DiskWorldTmpAvgRead.Bandwidth),
						},
						{
							Key:   "storage.disk-world-tmp.read.iops",
							Value: fmt.Sprintf("%.2f", result.DiskWorldTmpAvgRead.IOPS),
						},
						{
							Key:   "storage.disk-world-tmp.write.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.DiskWorldTmpAvgWrite.Bandwidth),
						},
						{
							Key:   "storage.disk-world-tmp.write.iops",
							Value: fmt.Sprintf("%.2f", result.DiskWorldTmpAvgWrite.IOPS),
						},
						{
							Key:   "storage.disk-tmp.read.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.DiskTmpAvgRead.Bandwidth),
						},
						{
							Key:   "storage.disk-tmp.read.iops",
							Value: fmt.Sprintf("%.2f", result.DiskTmpAvgRead.IOPS),
						},
						{
							Key:   "storage.disk-tmp.write.bw.mibps",
							Value: fmt.Sprintf("%.2f", result.DiskTmpAvgWrite.Bandwidth),
						},
						{
							Key:   "storage.disk-tmp.write.iops",
							Value: fmt.Sprintf("%.2f", result.DiskTmpAvgWrite.IOPS),
						},
					})
				}

				if (oldInfo.Addr == common.Address{}) {
					labels = metascheduler.ProcessLabels(labels)
					logger.I.Info("trying to register since we are not in the grid")
					if err := container.metascheduler.Register(
						ctx,
						hardware,
						prices,
						labels,
					); err != nil {
						logger.I.Fatal("supervisor failed to register", zap.Error(err))
					}
				}
			}()
		} else {
			logger.I.Warn("benchmark disabled, will not register to the smart-contract")
		}

		go func() {
			if err := container.gc.Loop(ctx); err != nil {
				logger.I.Fatal("gc stopped", zap.Error(err))
			}
		}()

		return container.jobWatcher.Watch(ctx)
	},
}

func launchBenchmarks(
	parent context.Context,
	benchmarkLauncher benchmark.Launcher,
	hplOpts []benchmark.Option,
	osuOpts []benchmark.Option,
	speedtestOpts []benchmark.Option,
	iorOpts []benchmark.Option,
) {
	start := time.Now()
	g, ctx := errgroup.WithContext(parent)

	g.Go(func() error {
		b, err := benchmark.GenerateOSUBenchmark(osuOpts...)
		if err != nil {
			logger.I.Error("failed to generate osu benchmark", zap.Error(err))
			return err
		}

		if err := benchmarkLauncher.Launch(ctx, "osu", b); err != nil {
			logger.I.Error(
				"osu benchmark failed or failed to be tracked",
				zap.Error(err),
			)
			return err
		}
		return nil
	})

	g.Go(func() error {
		b, err := benchmark.GenerateSpeedTestBenchmark(speedtestOpts...)
		if err != nil {
			logger.I.Error("failed to generate speedtest benchmark", zap.Error(err))
			return err
		}

		if err := benchmarkLauncher.Launch(ctx, "speedtest", b); err != nil {
			logger.I.Error(
				"speedtest benchmark failed or failed to be tracked",
				zap.Error(err),
			)
			return err
		}
		return nil
	})

	g.Go(func() error {
		b, err := benchmark.GeneratePhase1HPLBenchmark(hplOpts...)
		if err != nil {
			logger.I.Error("failed to generate hpl phase 1 benchmark", zap.Error(err))
			return err
		}

		if err := benchmarkLauncher.Launch(ctx, "hpl-phase1", b); err != nil {
			logger.I.Error(
				"hpl-phase1 benchmark failed or failed to be tracked",
				zap.Error(err),
			)
			return err
		}
		return nil
	})

	g.Go(func() error {
		b, err := benchmark.GenerateIORBenchmark(iorOpts...)
		if err != nil {
			logger.I.Error("failed to generate ior benchmark", zap.Error(err))
			return err
		}

		if err := benchmarkLauncher.Launch(ctx, "ior", b); err != nil {
			logger.I.Error(
				"ior benchmark failed or failed to be tracked",
				zap.Error(err),
			)
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		logger.I.Warn("cancelling benchmarks due to failure", zap.Error(err))
		if err := benchmarkLauncher.Cancel(parent, "osu"); err != nil {
			logger.I.Warn("failed to cancel osu benchmark", zap.Error(err))
		}
		if err := benchmarkLauncher.Cancel(parent, "speedtest"); err != nil {
			logger.I.Warn("failed to cancel speedtest benchmark", zap.Error(err))
		}
		if err := benchmarkLauncher.Cancel(parent, "hpl-phase1"); err != nil {
			logger.I.Warn("failed to cancel hpl-phase1 benchmark", zap.Error(err))
		}
		if err := benchmarkLauncher.Cancel(parent, "ior"); err != nil {
			logger.I.Warn("failed to cancel ior benchmark", zap.Error(err))
		}
		logger.I.Fatal("benchmarks failed", zap.Error(err))
	}

	done, errc := benchmark.DefaultStore.WaitForCompletion(parent)
	select {
	case <-done:
	case err := <-errc:
		logger.I.Fatal("benchmark has failed",
			zap.Error(err),
			zap.Any("results", benchmark.DefaultStore.Dump()),
		)
	}

	logger.I.Info("benchmark has finished",
		zap.Any("results", benchmark.DefaultStore.Dump()),
		zap.Duration("duration", time.Since(start)),
	)
}

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
