# The Grid CLI

Submit directly jobs to the DeepSquare Grid!

![tui-demo](./assets/tui-demo.svg)

![submit-demo](./assets/submit-demo.svg)

## Usage

```shell
NAME:
   deepsquaretui - Overwatch the job scheduling and register the compute to the Deepsquare Grid.

USAGE:
   deepsquaretui [global options] command [command options] [arguments...]

COMMANDS:
   allowance  Manage allowance.
   credit     Manage credits.
   init       Bootstrap a job workflow file.
   job        Manage jobs.
   provider   Manage providers (need to use an admin smart-contract).
   submit     Quickly submit a job.
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --metascheduler.rpc value             Metascheduler Avalanche C-Chain JSON-RPC endpoint. (default: "https://testnet.deepsquare.run/rpc") [$METASCHEDULER_RPC]
   --metascheduler.ws value              Metascheduler Avalanche C-Chain WS endpoint. (default: "wss://testnet.deepsquare.run/ws") [$METASCHEDULER_WS]
   --metascheduler.smart-contract value  Metascheduler smart-contract address. (default: "0x3707aB457CF457275b7ec32e203c54df80C299d5") [$METASCHEDULER_SMART_CONTRACT]
   --sbatch.endpoint value               SBatch Service GraphQL endpoint. (default: "https://sbatch.deepsquare.run/graphql") [$SBATCH_ENDPOINT]
   --logger.endpoint value               Grid Logger endpoint. (default: "https://grid-logger.deepsquare.run") [$LOGGER_ENDPOINT]
   --private-key value                   An hexadecimal private key for ethereum transactions. [$ETH_PRIVATE_KEY]
   --debug                               Debug logging (default: false) [$DEBUG]
   --help, -h                            show help
   --version, -v                         print the version
```

**`submit`**:

```shell
NAME:
   deepsquaretui submit - Quickly submit a job.

USAGE:
   deepsquaretui submit [command options] <job.yaml>

OPTIONS:
   DeepSquare Settings:

   --logger.endpoint value               Grid Logger endpoint. (default: "https://grid-logger.deepsquare.run") [$LOGGER_ENDPOINT]
   --metascheduler.rpc value             Metascheduler Avalanche C-Chain JSON-RPC endpoint. (default: "https://testnet.deepsquare.run/rpc") [$METASCHEDULER_RPC]
   --metascheduler.smart-contract value  Metascheduler smart-contract address. (default: "0x3707aB457CF457275b7ec32e203c54df80C299d5") [$METASCHEDULER_SMART_CONTRACT]
   --metascheduler.ws value              Metascheduler Avalanche C-Chain WS endpoint. (default: "wss://testnet.deepsquare.run/ws") [$METASCHEDULER_WS]
   --private-key value                   An hexadecimal private key for ethereum transactions. [$ETH_PRIVATE_KEY]
   --sbatch.endpoint value               SBatch Service GraphQL endpoint. (default: "https://sbatch.deepsquare.run/graphql") [$SBATCH_ENDPOINT]

   Submit Settings:

   --affinities key<value [ --affinities key<value ]  Affinities flag. Used to filter the clusters. Format: key<value, `key<=value`, `key=value`, `key>=value`, `key>value`, `key!=value`
   --credits value                                    Allocated a number of credits. Unit is 1e18. Is a float and is not precise. (default: 0)
   --credits-wei value                                Allocated a number of credits. Unit is wei. Is a big int.
   --exit-on-job-exit, -e                             Exit the job after the job has finished and throw on error. (default: false)
   --job-name value                                   The job name.
   --no-timestamp, --no-ts                            Hide timestamp. (default: false)
   --uses key=value [ --uses key=value ]              Uses flag. Used to filter the clusters. Format: key=value
   --watch, -w                                        Watch logs after submitting the job (default: false)

```

## Install binraries

You can download static binaries in the [Releases tab](https://github.com/deepsquare-io/grid/releases?q=cli&expanded=true).

## Install from source

Install Go and run:

```shell
go install github.com/deepsquare-io/grid/cli/cmd/deepsquaretui@latest
```

## Build

The `main` is stored in the `cmd/deepsquaretui` package. You can build the CLI using the following:

```sh
make
# Defaults to: make bin/deepsquaretui
```

Available `make`` commands:

```shell
# Build
bin/deepsquaretui
bin/deepsquaretui-darwin-amd64
bin/deepsquaretui-darwin-arm64
bin/deepsquaretui-freebsd-amd64
bin/deepsquaretui-freebsd-arm64
bin/deepsquaretui-linux-amd64
bin/deepsquaretui-linux-arm64
bin/deepsquaretui-linux-mips64
bin/deepsquaretui-linux-mips64le
bin/deepsquaretui-linux-ppc64
bin/deepsquaretui-linux-ppc64le
bin/deepsquaretui-linux-riscv64
bin/deepsquaretui-linux-s390x
bin/deepsquaretui-windows-amd64.exe
build-all

# Checksums
bin/checksums.md
bin/checksums.txt

# Clean builds
clean

# Generate code
generate

# Generate license
license

# Lint code
lint

# Run unit tests
unit

# Print version that will be used for the next build.
version
```

## Licence

The DeepSquare CLI library (i.e. all code outside of the `cmd` and `tui` directories) is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the COPYING file.

The DeepSquare CLI binaries (i.e. all code inside of the `cmd` and `tui` directories) are licensed under the GNU General Public License v3.0, also included in our repository in the COPYING.GPL3 file.
