package ether

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func FromWei(amount *big.Int) *big.Float {
	if amount == nil {
		return new(big.Float)
	}
	wei := new(big.Float).SetInt(amount)
	return new(big.Float).Quo(wei, big.NewFloat(params.Ether))
}

func ToWei(value *big.Float) *big.Int {
	if value == nil {
		return new(big.Int)
	}
	amount := new(big.Float).Mul(value, big.NewFloat(params.Ether))
	wei := new(big.Int)
	amount.Int(wei)
	return wei
}
