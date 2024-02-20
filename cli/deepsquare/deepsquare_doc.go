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

To submit jobs, you need to set the allowance first to allow the meta-scheduler to take you some credits:

	lockedCredits, _ := new(big.Int).SetString("100000000000000000000", 10)

	// Set allowance
	curr, err := client.AllowanceManager.GetAllowance(ctx)
	if err != nil {
		return err
	}
	if err = client.AllowanceManager.SetAllowance(ctx, curr.Add(curr, lockedCredits)); err != nil {
		return err
	}

You can set a high number to allow auto-topup:

	// Set allowance
	limit, _ := new(big.Int).SetString("10000000000000000000000000000000000000", 10)
	if err = client.AllowanceManager.SetAllowance(ctx, curr.Add(curr, limit)); err != nil {
		return err
	}

The credits will be used to finance the project and the infrastructure providers.

After settings the allowance, you can submit a job:

	_, err = client.SubmitJob(
		ctx,
		&sbatch.Job{
			Resources: &sbatch.JobResources{
				Tasks:       1,
				CPUsPerTask: 1,
				MemPerCPU:   100,
				GPUs: 0,
			},
			Steps: []*sbatch.Step{
				{
					Run: &sbatch.StepRun{
						Command: "echo test",
					},
				},
			},
		},
		lockedCredits,
		jobName,
	)

The object [sbatch.Job] has `json` and `yaml` tags, so it's possible to unmarshall a JSON object to an [sbatch.Job].

# Managing Jobs

Get a job:

	job, err := client.GetJob(ctx, jobID)

Get your jobs:

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

Top-up a job:

	err := client.TopUpJob(ctx, jobID, amount)

Panicking a job (only for admins):

	err := client.PanicJob(ctx, jobID, reason)

# Managing Allowance

To get, do:

	allowance, err := client.AllowanceManager.GetAllowance(ctx)

To set, do:

	err := client.AllowanceManager.SetAllowance(ctx, amount)

To watch, do:

	approvals := make(chan types.Approval, 1)
		sub, err := m.watcher.SubscribeEvents(
			ctx,
			event.FilterApproval(approvals),
		)
		if err != nil {
			// ...
		}
		defer sub.Unsubscribe()

		allowances, err := m.client.AllowanceManager.ReduceToAllowance(ctx, approvals)

		for {
			select {
			case allowance := <-allowances:
				// Handle allowance

			case err := <-sub.Err():
				// Handle error
			}
		}

# Managing Credits

To get, do:

	credits, err := client.CreditManager.Balance(ctx)
	// Or
	credits, err := client.CreditManager.BalanceOf(ctx, address)

To transfer, do:

	err := client.CreditManager.Transfer(ctx, address, amount)

To watch, do:

	transfers := make(chan types.Transfer, 1)
	sub, err := m.watcher.SubscribeEvents(
		ctx,
		event.FilterTransfer(transfers),
	)
	if err != nil {
		// ...
	}
	defer sub.Unsubscribe()

	balances, err := m.client.CreditManager.ReduceToBalance(ctx, transfers)

	for {
		select {
		case balance := <-balances:
			// Handle balance

		case err := <-sub.Err():
			// Handle error
		}
	}

# Managing Providers (for admins)

To get, do:

	provider, err := client.ProviderManager.GetProvider(ctx, providerAddress)

To list all providers, do:

	providers, err := client.ProviderManager.GetProviders(ctx, providerAddress)

To approve, do:

	err := client.ProviderManager.Approve(ctx, providerAddress)

To remove, do:

	err := client.ProviderManager.Remove(ctx, providerAddress)

# Watching events

To watch the events, the watcher has a `SubscribeEvents` method in which you can pass filters:

	transfers := make(chan types.Transfer, 1)
	sub, err := m.watcher.SubscribeEvents(
		ctx,
		event.FilterTransfer(transfers),
	)
	if err != nil {
		// ...
	}
	defer sub.Unsubscribe()

A thread will be handling the passing of data to the channels.

Call `sub.Unsubscribe()` to stop the subscription, and use `<-sub.Err()` to fetch the error.
*/
package deepsquare
