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

package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/deepsquare-io/grid/cli/types/event"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func clientExample(ctx context.Context) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Initialize the client
	client, err := deepsquare.NewClient(ctx, &deepsquare.ClientConfig{
		MetaschedulerAddress: common.HexToAddress("0x48af46ee836514551886bbC3b5920Eba81126F62"),
		UserPrivateKey:       pk, // Optional, but needed for authenticated requests
		// RPCEndpoint:          "https://testnet.deepsquare.run/rpc",    // Optional
		// SBatchEndpoint:       "https://sbatch.deepsquare.run/graphql", // Optional
		// LoggerEndpoint:       "https://grid-logger.deepsquare.run",    // Optional
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Test the client by checking the balance
	value, err := client.Balance(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if new(big.Int).Cmp(value) != 0 {
		panic("the balance is not zero ?!")
	}

	fmt.Println("balance:", value.String())
}

func watcherExample(ctx context.Context) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Initialize the client
	watcher, err := deepsquare.NewWatcher(ctx, &deepsquare.WatcherConfig{
		MetaschedulerAddress: common.HexToAddress("0x48af46ee836514551886bbC3b5920Eba81126F62"),
		UserPrivateKey:       pk, // Optional, but needed for authenticated requests
		// RPCEndpoint:          "https://testnet.deepsquare.run/rpc",    // Optional
		// SBatchEndpoint:       "https://sbatch.deepsquare.run/graphql", // Optional
		// LoggerEndpoint:       "https://grid-logger.deepsquare.run",    // Optional
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer watcher.Close()

	// Test the watcher
	ch := make(chan types.NewJobRequest, 1)
	defer close(ch)
	sub, err := watcher.SubscribeEvents(ctx, event.FilterNewJobRequest(ch))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer sub.Unsubscribe()

	// TODO: handle events
}

func main() {
	ctx := context.Background()
	clientExample(ctx)
	watcherExample(ctx)
}
