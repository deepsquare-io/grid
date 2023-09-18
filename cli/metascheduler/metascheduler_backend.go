// Package deepsquare defines APIs for interacting with the DeepSquare Grid.
package metascheduler

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrNewRequestJobNotFound = errors.New("new request job event not found")
)

var (
	metaschedulerABI   *abi.ABI
	ierc20ABI          *abi.ABI
	newJobRequestEvent abi.Event
	jobTransitionEvent abi.Event
	transferEvent      abi.Event
	approvalEvent      abi.Event
)

func init() {
	var err error
	metaschedulerABI, err = metaschedulerabi.MetaSchedulerMetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to parse metascheduler contract ABI: %w", err))
	}
	ierc20ABI, err = metaschedulerabi.IERC20MetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to parse erc20 contract ABI: %w", err))
	}

	// Find the event signature dynamically
	var ok bool
	newJobRequestEvent, ok = metaschedulerABI.Events["NewJobRequestEvent"]
	if !ok {
		panic(fmt.Errorf("failed to get NewJobRequestEvent: %w", err))
	}

	jobTransitionEvent, ok = metaschedulerABI.Events["JobTransitionEvent"]
	if !ok {
		panic(fmt.Errorf("failed to get JobTransitionEvent: %w", err))
	}

	transferEvent, ok = ierc20ABI.Events["Transfer"]
	if !ok {
		panic(fmt.Errorf("failed to get Transfer: %w", err))
	}
	approvalEvent, ok = ierc20ABI.Events["Approval"]
	if !ok {
		panic(fmt.Errorf("failed to get Approval: %w", err))
	}
}

type EthereumBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

// Backend is a wrapper around the EthereumBackend used to interact with the Meta-Scheduler smart-contract.
type Backend struct {
	MetaschedulerAddress common.Address
	// EthereumBackend is the Ethereum Client.
	//
	// TODO: check if websocket or rpc.
	EthereumBackend
	// ChainID of the blockchain.
	ChainID *big.Int
	// PrivateKey of the user.
	UserPrivateKey *ecdsa.PrivateKey
}

func (b *Backend) from() (addr common.Address) {
	if b.UserPrivateKey == nil {
		return addr
	}
	return crypto.PubkeyToAddress(b.UserPrivateKey.PublicKey)
}
