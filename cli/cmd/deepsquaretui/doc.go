// Copyright (C) 2023 DeepSquare
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// DeepSquare TUI is a client to launch jobs on the DeepSquare Grid.
//
// Usage:
//
//	deepsquaretui [global options] command [command options] [arguments...]
//
// The flags are:
//
//	--metascheduler.rpc value
//		Metascheduler Avalanche C-Chain JSON-RPC endpoint.
//		(default: "https://testnet.deepsquare.run/rpc") [$METASCHEDULER_RPC]
//	--metascheduler.ws value
//		Metascheduler Avalanche C-Chain WS endpoint.
//		(default: "wss://testnet.deepsquare.run/ws") [$METASCHEDULER_WS]
//	--metascheduler.smart-contract value
//		Metascheduler smart-contract address.
//		(default: "0x3707aB457CF457275b7ec32e203c54df80C299d5") [$METASCHEDULER_SMART_CONTRACT]
//	--sbatch.endpoint value
//		SBatch Service GraphQL endpoint.
//		(default: "https://sbatch.deepsquare.run/graphql") [$SBATCH_ENDPOINT]
//	--logger.endpoint value
//		Grid Logger endpoint.
//		(default: "https://grid-logger.deepsquare.run") [$LOGGER_ENDPOINT]
//	--private-key value
//		An hexadecimal private key for ethereum transactions. [$ETH_PRIVATE_KEY]
//	--debug
//		Debug logging (default: false) [$DEBUG]
//	--help, -h
//		show help
//	--version, -v
//		print the version
//
// The DeepSquare Terminal User Interface (TUI) shows the job statuses, logs
// and is able to launch DeepSquare Workflows from the terminal via the
// meta-scheduler smart-contract deployed on a EVM blockchain.
package main
