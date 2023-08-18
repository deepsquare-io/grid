# Deepsquare Grid Supervisor

The middleman between the job scheduler and the DeepSquare Meta-scheduler.

## Usage

### Summary

```shell
NAME:
   supervisor - Overwatch the job scheduling and register the compute to the Deepsquare Grid.

USAGE:
   supervisor [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

   Benchmark:

   --benchmark.disable                Disable benchmark (and registering). (default: false) [$BENCHMARK_DISABLE]
   --benchmark.hpl.image value        Docker image used for HPL benchmark (default: "registry-1.deepsquare.run#library/hpc-benchmarks:23.5") [$BENCHMARK_HPL_IMAGE]
   --benchmark.hpl.single-node        Force single node benchmark for HPL. (default: false) [$BENCHMARK_HPL_SINGLE_NODE]
   --benchmark.include-unresponsive   Force benchmark on unresponsive nodes (sinfo --responding --partition=<partition>). (default: false) [$BENCHMARK_UNRESPONSIVE]
   --benchmark.osu.image value        Docker image used for OSU benchmark (default: "registry-1.deepsquare.run#library/osu-benchmarks:latest") [$BENCHMARK_OSU_IMAGE]
   --benchmark.run-as value           User used for benchmark (default: "root") [$BENCHMARK_RUN_AS]
   --benchmark.speedtest.image value  Docker image used for SpeedTest benchmark (default: "registry-1.docker.io#gists/speedtest-cli:1.2.0") [$BENCHMARK_SPEEDTEST_IMAGE]
   --benchmark.time-limit value       Time limit (syntax is golang duration style). (default: 24h0m0s) [$BENCHMARK_TIME_LIMIT]
   --benchmark.trace                  Enables benchmark trace logging. Very verbose. (default: false) [$BENCHMARK_TRACE]
   --benchmark.ucx                    Use UCX transport for MPI. Choose this for RDMA. Do not for TCP. (default: false) [$BENCHMARK_UCX]
   --benchmark.ucx.affinity value     UCX Affinity. Select the devices with the format devices_for_node_1|devices_for_node_2|...

See 'ucx_info -bd' to see available devices.

Examples:
  mlx5_0:1|mlx5_0:1 means that cn1 will use mlx5_0 port 1 and cn2 will use mlx5_0 port 1
  mlx5_0:1,mlx5_0:1|mlx5_0:1 means that cn1 will use mlx5_0 port 1 or mlx5_0 port 1, and cn2 will use mlx5_0 port 1 [$BENCHMARK_UCX_AFFINITY]
   --benchmark.ucx.transport value  UCX Tranport. Select the common transport.

See 'ucx_info -bd' to see available tranports.

Value is often: sm,self,rc (shared memory, self, rdma reliable connected). Set to empty to set automatically.

Note that TCP is not supported at the moment. [$BENCHMARK_UCX_TRANSPORT]

   MetaScheduler:

   --metascheduler.endpoint.rpc value    Metascheduler Avalanche C-Chain JSON-RPC endpoint. (default: "https://testnet.deepsquare.run/rpc") [$METASCHEDULER_ENDPOINT_RPC]
   --metascheduler.endpoint.ws value     Metascheduler Avalanche C-Chain WS endpoint. (default: "wss://testnet.deepsquare.run/ws") [$METASCHEDULER_ENDPOINT_WS]
   --metascheduler.private-key value     An hexadecimal private key for ethereum transactions. [$ETH_PRIVATE_KEY]
   --metascheduler.smart-contract value  Metascheduler smart-contract address. (default: "0x") [$METASCHEDULER_SMART_CONTRACT]

   Miscellaneous:

   --nvidia-smi value  Server-side nvidia-smi path. (default: "nvidia-smi") [$NVIDIA_SMI_PATH]
   --trace             Trace logging (default: false) [$TRACE]

   Network:

   --grpc.listen-address value  Address to listen on. Is used for receiving job status via the job completion plugin. (default: ":3000") [$LISTEN_ADDRESS]
   --public-address value       Public address or address of the reverse proxy. Is used by the SLURM plugins to know where to report job statuses. Must be protected with TLS. (default: "supervisor.example.com:3000") [$PUBLIC_ADDRESS]

   SBatch API:

   --sbatch.endpoint value  SBatch API gRPC endpoint. (default: "127.0.0.1:443") [$SBATCH_ENDPOINT]
   --sbatch.tls             Enable TLS for the SBatch API. (default: true) [$SBATCH_TLS_ENABLE]
   --sbatch.tls.ca value    Path to CA certificate for TLS verification. [$SBATCH_CA]
   --sbatch.tls.insecure    Skip TLS verification. By enabling it, sbatch.tls.ca is ignored. (default: false) [$SBATCH_TLS_INSECURE]

   Secure Transport:

   --tls                  Enable TLS for HTTP. (default: false) [$TLS_ENABLE]
   --tls.cert-file value  TLS Certificate file. [$TLS_CERT]
   --tls.key-file value   TLS Private Key file. [$TLS_KEY]

   Slurm:

   --slurm.batch value      Server-side SLURM sbatch path. (default: "/usr/bin/sbatch") [$SLURM_SBATCH_PATH]
   --slurm.cancel value     Server-side SLURM scancel path. (default: "/usr/bin/scancel") [$SLURM_SCANCEL_PATH]
   --slurm.control value    Server-side SLURM scontrol path. (default: "/usr/bin/scontrol") [$SLURM_SCONTROL_PATH]
   --slurm.info value       Server-side SLURM info path. (default: "/usr/bin/sinfo") [$SLURM_SINFO_PATH]
   --slurm.partition value  Slurm partition used for jobs and registering.

All the specifications returned by 'scontrol show partition' will be registered to the blockchain. (default: "main") [$SLURM_PARTITION]
   --slurm.squeue value             Server-side SLURM squeue path. (default: "/usr/bin/squeue") [$SLURM_SQUEUE_PATH]
   --slurm.ssh.address value        Address of the Slurm login node. [$SLURM_SSH_ADDRESS]
   --slurm.ssh.admin-user scontrol  SLURM admin user used for calling scontrol commands. [$SLURM_SSH_ADMIN_USER]
   --slurm.ssh.private-key value    Base64-encoded one line SSH private key used for impersonation. The public key must be inserted in the authorized_keys file of each user. [$SLURM_SSH_PRIVATE_KEY]
```

## Docker

You can pull the image with:

```shell
docker pull ghcr.io/deepsquare-io/supervisor:<version>
```

More details [here](https://github.com/deepsquare-io/the-grid/pkgs/container/supervisor).

## Build

The `main` function is stored in the `cmd` package. To build the supervisor, do:

```shell
make
```
