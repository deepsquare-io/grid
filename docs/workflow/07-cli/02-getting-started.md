# Getting Started

## Install binaries

You can download static binaries in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=cli%2F&expanded=true).

## Install from source

With Go, it is possible to install from source directly by running:

```shell
go install github.com/deepsquare-io/grid/cli/cmd/deepsquaretui@latest
```

## Usage

1. Prepare a `.env` file along the executable or load these environment variables:

   ```shell
   METASCHEDULER_SMART_CONTRACT=<0x address>
   ETH_PRIVATE_KEY=<hex key, without 0x>
   ```

   You can find the smart-contract address in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=smart-contracts&expanded=true).

2. You can run CLI commands by executing:

   ```shell
   deepsquaretui command [command options] [arguments...]
   ```

   The available commands are:

   ```shell
   allowance  Manage allowance.
   credit     Manage credits.
   init       Bootstrap a job workflow file.
   job        Manage jobs.
   provider   Manage providers (need to use an admin smart-contract).
   submit     Quickly submit a job.
   help, h    Shows a list of commands or help for one command
   ```

3. You can run the TUI by simply running `deepsquaretui`.

The next pages will explain the commands and their usages.
