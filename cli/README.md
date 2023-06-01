# The Grid CLI

Submit directly jobs to the DeepSquare Grid!

## Usage

**`grid submit`**:

```shell
Usage:
  grid submit <path to script> [flags]

Flags:
  -c, --cpus-per-task uint                    Allocated CPUs per task.
  -t, --credits string                        Amount of credits locked for the job, which is equivalent to the time limit.
      --eth.private-key string                An hexadecimal private key for ethereum transactions. (env: ETH_PRIVATE_KEY)
      --gpus uint                             Allocated GPUs per node.
  -h, --help                                  help for submit
      --mem uint                              Allocated memory per node (MB).
      --metascheduler.endpoint string         Metascheduler RPC endpoint. (env: METASCHEDULER_ENDPOINT) (default "https://testnet.deepsquare.run/rpc")
      --metascheduler.smart-contract string   Metascheduler smart-contract address. Must have the prefix 0x. (env: METASCHEDULER_SMART_CONTRACT)
  -N, --nodes uint                            Allocated nodes. (default 1)
  -n, --tasks uint                            Run the same script in parallel if tasks > 1. (default 1)

Global Flags:
  -d, --debug   Show debug logging.
```

## Download

You can download static binaries in the [Releases tab](https://github.com/deepsquare-io/the-grid/releases?q=cli&expanded=true).

## Build

The `main` is stored in the `cmd` package. You can build the CLI using the following:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o grid ./cmd
```

More examples are stored in the [`scripts/999.build-all.sh` script](scripts/999.build-all.sh).
