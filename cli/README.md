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

## Install binaries

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

Available `make` commands:

```shell
# Build
make bin/deepsquaretui
make bin/deepsquaretui-darwin-amd64
make bin/deepsquaretui-darwin-arm64
make bin/deepsquaretui-freebsd-amd64
make bin/deepsquaretui-freebsd-arm64
make bin/deepsquaretui-linux-amd64
make bin/deepsquaretui-linux-arm64
make bin/deepsquaretui-linux-mips64
make bin/deepsquaretui-linux-mips64le
make bin/deepsquaretui-linux-ppc64
make bin/deepsquaretui-linux-ppc64le
make bin/deepsquaretui-linux-riscv64
make bin/deepsquaretui-linux-s390x
make bin/deepsquaretui-windows-amd64.exe
make build-all

# Checksums
make bin/checksums.md
make bin/checksums.txt

# Clean builds
make clean

# Generate code
make generate

# Generate license
make license

# Lint code
make lint

# Run unit tests
make unit

# Print version that will be used for the next build.
make version

# Preview doc
make doc
```

## Use as a library

The `deepsquare`, `sbatch`, `logger` and `metascheduler` packages can be used as a library. For example, to submit a job:

```go
// Parse private key
pk, err := crypto.HexToECDSA(ethHexPK)
if err != nil {
   // ...
}

// Initialize client for simple RPCs
client, err := deepsquare.NewClient(ctx, &deepsquare.ClientConfig{
   MetaschedulerAddress: common.HexToAddress("0x..."),
   RPCEndpoint:          "https://testnet.deepsquare.run/rpc",  // Optional
   SBatchEndpoint:       "https://sbatch.deepsquare.run/graphql",  // Optional
   LoggerEndpoint:       "https://grid-logger.deepsquare.run",  // Optional
   UserPrivateKey:       pk,  // Optional, but needed for authenticated requests
})

// Example of job submit
jobID, err = client.SubmitJob(
   ctx,
   &sbatch.Job{
      Resources: &sbatch.JobResources{
         Tasks:       1,
         CpusPerTask: 1,
         MemPerCPU:   100,
         GpusPerTask: 0,
      },
      Steps: []*sbatch.Step{
         {
            Run: &sbatch.StepRun{
               Command: "echo test",
            },
         },
      },
   },
   big.NewInt(100),
   jobName,
)
```

For more information, check out:

- [The Go Package Documentation: Library documentation](https://pkg.go.dev/github.com/deepsquare-io/grid/cli)
- [The Go Package Documentation: DeepSquare Client](https://pkg.go.dev/github.com/deepsquare-io/grid/cli/deepsquare)
- [Examples](https://github.com/deepsquare-io/grid/tree/main/cli/_examples)

## Licence

The DeepSquare CLI library (i.e. all code outside the `cmd` and `tui` directories) is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the COPYING file.

The DeepSquare CLI binaries (i.e. all code inside the `cmd` and `tui` directories) are licensed under the GNU General Public License v3.0, also included in our repository in the COPYING.GPL3 file.
