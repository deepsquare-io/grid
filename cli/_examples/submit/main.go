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

package main

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/sbatch"
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

	// Submit a fake job.
	var jobName [32]byte
	jobNameS := "test"
	copy(jobName[:], jobNameS)
	_, err = client.SubmitJob(
		ctx,
		&sbatch.Job{
			Resources: &sbatch.JobResources{
				Tasks:       1,
				CpusPerTask: 1,
				MemPerCPU:   100,
				GpusPerTask: 0,
			},
			Steps: []*sbatch.Step{
				{
					Run: &sbatch.StepRun{
						Command: "echo test",
					},
				},
			},
		},
		big.NewInt(100),
		jobName,
	)

	// We expect an error
	if err == nil {
		panic("received no error")
	}
	if !strings.Contains(err.Error(), "insufficient funds for transfer") {
		log.Fatalln(err.Error())
	}

	log.Printf("we've received an error '%s', but that's expected", err)
}
