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

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	ctx := context.Background()
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

	// Get our jobs
	jobs, err := client.GetJobs(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Iterate
	for jobs.Next(ctx) {
		fmt.Println(jobs.Current())
	}

	// Handle error
	if jobs.Error() != nil {
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
