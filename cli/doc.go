// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

/*
Package cli is the root of the Go client for interacting with the DeepSquare Grid.

# The Terminal User Interface

The TUI can be run by simply running the [github.com/deepsquare-io/grid/cli/cmd/dps] package.
Its implementation is in the [github.com/deepsquare-io/grid/cli/tui] package and uses the [bubbletea framework] to build the terminal application.

# The Command Line Interface

The CLI can be run by running subcommands defined in the [github.com/deepsquare-io/grid/cli/cmd] package.
Compared to the TUI, the CLI contains all the commands of the DeepSquare Client.

# The DeepSquare Client

A DeepSquare client is available in the [github.com/deepsquare-io/grid/cli/deepsquare] package. It implements
all of the features needed to manage jobs, providers, credits and job logs.

Example:

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
	curr, err := client.GetAllowance(ctx)
	if err != nil {
		// ...
	}
	err = client.SetAllowance(ctx, curr.Add(curr, lockedAmount))
	if err != nil {
		// ...
	}
	jobID, err := client.SubmitJob(ctx, job, lockedAmount, jobName, types.WithAffinity(affinities))

	// Initialize client for streaming RPCs
	watcher, err := deepsquare.NewWatcher(ctx, &deepsquare.WatcherConfig{
		MetaschedulerAddress: common.HexToAddress("0x..."),
		RPCEndpoint:          "https://testnet.deepsquare.run/rpc",  // Optional
		WSEndpoint:           "https://testnet.deepsquare.run/ws",  // Optional
		UserPrivateKey:       pk,  // Optional, but needed for authenticated requests
	})
	defer watcher.Close()

	// Example of watching job transition
	transitions := make(chan types.JobTransition, 1)
	sub, err := watcher.SubscribeEvents(ctx, types.FilterJobTransition(transitions))
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

See the [github.com/deepsquare-io/grid/cli/deepsquare] package documentation for more information.

# Fine-grained RPC or WebSocket client

If the DeepSquare client is too much "coupled", a more fine-grained client is
available in the [github.com/deepsquare-io/grid/cli/metascheduler] package.

Example:

	// Initialize RPC client
	ethClientRPC := ethclient.Dial("https://testnet.deepsquare.run/rpc")
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	rpcBackend := metascheduler.Backend{
		EthereumBackend:      ethClientRPC,
		MetaschedulerAddress: common.HexToAddress("0x..."),
		ChainID:              chainID,
		UserPrivateKey:       pk,
	}
	clientset := metascheduler.NewRPCClientSet(rpcBackend)

	// Example of job submit
	curr, err := clientset.AllowanceManager().GetAllowance(ctx)
	if err != nil {
		// ...
	}
	err = clientset.AllowanceManager().SetAllowance(ctx, curr.Add(curr, lockedAmount))
	if err != nil {
		// ...
	}

	sbatch := sbatch.NewService(http.DefaultClient, "https://sbatch.deepsquare.run/graphql")
	jobID, err := clientset.JobScheduler(sbatch).SubmitJob(ctx, job, lockedAmount, jobName, types.WithAffinity(affinities))

	// Initialize WebSocket client
	ethClientWS := ethclient.Dial("https://testnet.deepsquare.run/ws")
	defer ethClientWS.Close()
	wsBackend := metascheduler.Backend{
		MetaschedulerAddress: common.HexToAddress("0x..."),
		EthereumBackend:      ethClientWS,
		ChainID:              chainID,
		UserPrivateKey:       pk,
	}
	es := metascheduler.NewEventSubscriber(rpcBackend, wsBackend)

	// Example of watching job transition
	transitions := make(chan types.JobTransition, 1)
	sub, err := es.SubscribeEvents(ctx, types.FilterJobTransition(transitions))
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

See the [github.com/deepsquare-io/grid/cli/metascheduler] package documentation for more information.

[bubbletea framework]: https://github.com/charmbracelet/bubbletea
*/
package cli
