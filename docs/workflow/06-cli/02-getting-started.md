# Getting Started

## Install binaries

You can download static binaries in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=cli%2F&expanded=true).

To install it:

```shell
# Create a local directory to store executable.
mkdir -p ~/.local/bin
# Move executable to install directory.
mv dps* ~/.local/bin/dps
# Add executation permission.
chmod +x ~/.local/bin/dps
```

**Or, in one shot: `curl https://raw.githubusercontent.com/deepsquare-io/grid/main/cli/get-dps | bash`.**

For **Windows** users:

```bat
:: Create a local directory to store executable.
mkdir "%USERPROFILE%\.local\bin"
:: Move executable to install directory.
move dps* "%USERPROFILE%\.local\bin\dps"
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
   ```

   You can find the smart-contract address in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=smart-contracts&expanded=true).

2. Create a directory at `.dps` and put your wallet private key at `.dps/key`:

   ```shell
   mkdir -p .dps
   echo "0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff" > .dps/key
   chmod 600 .dps/key
   ```

   You can fetch your private key from MetaMask by following [this guide](https://support.metamask.io/hc/en-us/articles/360015289632-How-to-export-an-account-s-private-key#).

   If you want to generate a private key instead, you can run `dps init`.

   To run on DeepSquare, you need **credits**. You can fetch free credits by filling [this form](https://share-eu1.hsforms.com/1PVlRXYdMSdy-iBH_PXx_0wev6gi).

3. You can run CLI commands by executing:

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

4. You can run the TUI by simply running `dps`.

The next pages will explain the commands and their usages.
