# Getting Started

## Install binaries

You can download static binaries in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=cli%2F&expanded=true).

To install it:

```shell
## One Shot Method:
curl https://raw.githubusercontent.com/deepsquare-io/grid/main/cli/get-dps | bash

## Manual Method:
# Create a local directory to store executable.
mkdir -p ~/.local/bin
# Move executable to install directory.
mv <path/to/dps> ~/.local/bin/dps
# Add executation permission.
chmod +x ~/.local/bin/dps
```

For **Windows** users, open `cmd.exe` (don't use PowerShell) and execute:

```bat
:: Create a local directory to store executable.
mkdir "%USERPROFILE%\.local\bin"
:: Move executable to install directory.
move <path/to/dps> "%USERPROFILE%\.local\bin\dps.exe"
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

1. Edit the`.bashrc` (or `.zshrc`) file and load this environment variable:

   ```shell
   export METASCHEDULER_SMART_CONTRACT=<0x address>
   ```

   You can find the smart-contract address (`MetaSchedulerAddr`) in the [Releases tab of the Grid git repository](https://github.com/deepsquare-io/grid/releases?q=smart-contracts&expanded=true).

2. Create a directory at `~/.dps` and put your wallet private key at `~/.dps/key`:

   ```shell
   mkdir -p ~/.dps
   echo "0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff" > ~/.dps/key
   # Ensure the permission of the key excludes world and group permissions.
   chmod 600 ~/.dps/key
   ```

   For **Windows** users:

   ```bat
   mkdir "%USERPROFILE%\.dps"
   echo "0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff" > "%USERPROFILE%\.dps\key"
   icacls "%USERPROFILE%\.dps\key" /inheritance:r
   icacls "%USERPROFILE%\.dps\key" /grant:r "%username%":"(R)"
   ```

   :::note

   **If you want to typing these steps manually, or want to generate a private key instead, you can run `dps init`.**

   :::

   You can fetch your private key from MetaMask by following [this guide](https://support.metamask.io/hc/en-us/articles/360015289632-How-to-export-an-account-s-private-key#).

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
