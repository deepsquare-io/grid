// Copyright (C) 2024 DeepSquare Association
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

package base58

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

type Encoder struct{}

func (enc Encoder) Encode(u uuid.UUID) string {
	return base58.Encode(u[:])
}

func (enc Encoder) Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}

type FakeEncoder struct{}

func (enc FakeEncoder) Encode(_ uuid.UUID) string {
	return "BhNwmz1fC9zVZ8im94bLbw"
}

func (enc FakeEncoder) Decode(_ string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode("BhNwmz1fC9zVZ8im94bLbw"))
}
