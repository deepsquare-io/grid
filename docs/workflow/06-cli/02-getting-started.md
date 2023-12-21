# Getting Started

## Install binaries

You can download static binaries in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=cli%2F&expanded=true).

To install it:

```shell
mkdir -p ~/.local/bin
# Move local install dir
mv dps* ~/.local/bin/dps
# Add executation permission
chmod +x ~/.local/bin/dps
```

For Windows users:

```bat
mkdir "%USERPROFILE%\.local\bin"
move dps* "%USERPROFILE%\.local\bin\dps"
cd "%USERPROFILE%\.local\bin\dps"
```

Make sure that `$HOME/.local/bin` (`%USERPROFILE%\.local\bin` for Windows) is added to the `$PATH` (`%PATH%`) environment variable.

[Help: Adding to PATH](https://www.java.com/en/download/help/path.html).

## Install from source

With Go, it is possible to install from source directly by running:

```shell
go install github.com/deepsquare-io/grid/cli/cmd/dps@latest
```

Make sure that `$HOME/go/bin` (`%HOME%\go\bin` for Windows) is added to the `$PATH` (`%PATH%`) environment variable.

[Help: Adding to PATH](https://www.java.com/en/download/help/path.html).

## Usage

1. Prepare a `.env` file along the executable or load these environment variables:

   ```shell
   METASCHEDULER_SMART_CONTRACT=<0x address>
   ETH_PRIVATE_KEY=<hex key, without 0x>
   ```

   You can find the smart-contract address in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=smart-contracts&expanded=true).

2. You can run CLI commands by executing:

   ```shell
   dps command [command options] [arguments...]
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

3. You can run the TUI by simply running `dps`.

The next pages will explain the commands and their usages.
