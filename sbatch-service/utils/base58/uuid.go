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

func (enc FakeEncoder) Encode(u uuid.UUID) string {
	return "BhNwmz1fC9zVZ8im94bLbw"
}

func (enc FakeEncoder) Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode("BhNwmz1fC9zVZ8im94bLbw"))
}
