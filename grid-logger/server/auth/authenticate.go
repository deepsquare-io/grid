package auth

import (
	"errors"
	"strings"

	"github.com/deepsquare-io/grid/grid-logger/logger"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

var (
	ErrAuthError = errors.New("authentication error")
)

func Verify(address string, data []byte, sig []byte) error {
	var hash []byte
	if sig[crypto.RecoveryIDOffset] > 1 {
		// Legacy Keccak256
		// Transform yellow paper V from 27/28 to 0/1
		sig[crypto.RecoveryIDOffset] -= 27
	}
	hash = accounts.TextHash(data)

	// Verify signature
	sigPublicKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		logger.I.Error("Authenticate.SigToPub",
			zap.Error(err),
			zap.String("hash", hexutil.Encode(hash)),
			zap.String("sig", hexutil.Encode(sig)),
		)
		return err
	}
	sigAddr := crypto.PubkeyToAddress(*sigPublicKey)

	if !strings.EqualFold(address, sigAddr.Hex()) {
		logger.I.Error(
			"Authenticate: addresses are not equal",
			zap.String("sig.Address", sigAddr.Hex()),
			zap.String("address", address),
			zap.String("sig", hexutil.Encode(sig)),
			zap.String("expected hash", hexutil.Encode(hash)),
		)
		return ErrAuthError
	}

	return nil
}
