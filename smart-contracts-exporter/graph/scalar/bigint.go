package scalar

import (
	"fmt"
	"io"
	"math/big"
)

type BigInt struct {
	*big.Int
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (bi *BigInt) UnmarshalGQL(v interface{}) error {
	biString, ok := v.(string)
	if !ok {
		return fmt.Errorf("BigInt must be a string")
	}

	bInt, ok := new(big.Int).SetString(biString, 10)
	if !ok {
		return fmt.Errorf("BigInt failed to parse string (wrong base?)")
	}

	bi.Int = bInt

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (bi BigInt) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("\"%s\"", bi.Int.String())))
}
