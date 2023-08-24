# The Grid CLI

Submit directly jobs to the DeepSquare Grid!

## Usage

```shell
NAME:
   deepsquaretui - Overwatch the job scheduling and register the compute to the Deepsquare Grid.

USAGE:
   deepsquaretui [global options] command [command options] [arguments...]

COMMANDS:
   submit   Quickly submit a job.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --metascheduler.rpc value             Metascheduler Avalanche C-Chain JSON-RPC endpoint. (default: "https://testnet.deepsquare.run/rpc") [$METASCHEDULER_RPC]
   --metascheduler.ws value              Metascheduler Avalanche C-Chain WS endpoint. (default: "wss://testnet.deepsquare.run/ws") [$METASCHEDULER_WS]
   --metascheduler.smart-contract value  Metascheduler smart-contract address. (default: "0xc9AcB97F1132f0FB5dC9c5733B7b04F9079540f0") [$METASCHEDULER_SMART_CONTRACT]
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
   --metascheduler.rpc value              Metascheduler Avalanche C-Chain JSON-RPC endpoint. (default: "https://testnet.deepsquare.run/rpc") [$METASCHEDULER_RPC]
   --metascheduler.smart-contract value   Metascheduler smart-contract address. (default: "0xc9AcB97F1132f0FB5dC9c5733B7b04F9079540f0") [$METASCHEDULER_SMART_CONTRACT]
   --sbatch.endpoint value                SBatch Service GraphQL endpoint. (default: "https://sbatch.deepsquare.run/graphql") [$SBATCH_ENDPOINT]
   --private-key value                    An hexadecimal private key for ethereum transactions. [$ETH_PRIVATE_KEY]
   --job-name value                       The job name.
   --uses key=value [ --uses key=value ]  Uses flag. Used to filter the clusters. Format: key=value
   --credits-wei value                    Allocated a number of credits. Unit is wei. Is a big int.
   --credits value                        Allocated a number of credits. Unit is 1e18. Is a float and is not precise. (default: 0)
   --help, -h                             show help
```

## Download

You can download static binaries in the [Releases tab](https://github.com/deepsquare-io/the-grid/releases?q=cli&expanded=true).

## Build

The `main` is stored in the `cmd` package. You can build the CLI using the following:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o grid ./cmd
```

More examples are stored in the [`scripts/999.build-all.sh` script](scripts/999.build-all.sh).

## Licence

The DeepSquare CLI library (i.e. all code outside of the `cmd` and `tui` directories) is licensed under the GNU Lesser General Public License v3.0, also included in our repository in the COPYING.LGPL3 file.

The DeepSquare CLI binaries (i.e. all code inside of the `cmd` and `tui` directories) are licensed under the GNU General Public License v3.0, also included in our repository in the COPYING.GPL3 file.
