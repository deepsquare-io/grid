package ether

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func FromWei(amount *big.Int) *big.Float {
	if amount == nil {
		return new(big.Float)
	}
	return new(big.Float).Quo(new(big.Float).SetInt(amount), big.NewFloat(params.Ether))
}
