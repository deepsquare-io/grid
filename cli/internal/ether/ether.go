// Copyright (C) 2023 DeepSquare Asociation
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
