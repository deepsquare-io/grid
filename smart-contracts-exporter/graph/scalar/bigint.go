package scalar

import (
	"fmt"
	"io"
	"math/big"

	"github.com/deepsquare-io/grid/smart-contracts-exporter/logger"
	"go.uber.org/zap"
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
	_, err := w.Write([]byte(fmt.Sprintf("\"%s\"", bi.Int.String())))
	if err != nil {
		logger.I.Error("failed to write string", zap.Error(err))
	}
}
