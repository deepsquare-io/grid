package types

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

type EthereumBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}
