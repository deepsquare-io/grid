// Copyright (C) 2023 DeepSquare Association
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
Package deepsquare provides a all-in-one Client that implements all of the DeepSquare services.

See more examples: https://github.com/deepsquare-io/grid/tree/main/_examples

# Initializing the clients

There are two types of client: the Client and the Watcher.
The Watcher uses WebSockets to monitor events, where the endpoint may be different from the RPC endpoint.

To initialize the client, do:

	// Initialize client for simple RPCs
	client, err := deepsquare.NewClient(ctx, &deepsquare.ClientConfig{
		MetaschedulerAddress: common.HexToAddress("0x..."),
		RPCEndpoint:          "https://testnet.deepsquare.run/rpc",  // Optional
		SBatchEndpoint:       "https://sbatch.deepsquare.run/graphql",  // Optional
		LoggerEndpoint:       "https://grid-logger.deepsquare.run",  // Optional
		UserPrivateKey:       pk,  // Optional, but needed for authenticated requests
	})

To initialize the watcher, do:

	// Initialize client for streaming RPCs
	watcher, err := deepsquare.NewWatcher(ctx, &deepsquare.WatcherConfig{
		MetaschedulerAddress: common.HexToAddress("0x..."),
		RPCEndpoint:          "https://testnet.deepsquare.run/rpc",  // Optional
		WSEndpoint:           "https://testnet.deepsquare.run/ws",  // Optional
		UserPrivateKey:       pk,  // Optional, but needed for authenticated requests
	})
	defer watcher.Close()

The private key can be parsed with the `go-ethereum` package:

	import (
		"github.com/ethereum/go-ethereum/common"
		"github.com/ethereum/go-ethereum/crypto"
	)

	// Parse private key
	pk, err := crypto.HexToECDSA(ethHexPK)
	if err != nil {
		// ...
	}

# Submitting Jobs

# Managing Jobs

# Managing Allowance

# Managing Credits

# Managing Providers

# Watching events
*/
package deepsquare
