package auth

import (
	"bytes"
	"strings"

	"github.com/deepsquare-io/the-grid/grid-logger/logger"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

func Authenticate(storage *MemStorage, address string, nonce []byte, sig []byte) (User, error) {
	user, err := storage.Get(address)
	if err != nil {
		return user, err
	}

	// Verify nonce
	if !bytes.Equal(user.Nonce, nonce) {
		logger.I.Error(
			"Authenticate: nonces are not equal",
			zap.String("recv nonce", hexutil.Encode(nonce)),
			zap.String("user.nonce", hexutil.Encode(user.Nonce)),
		)
		return user, ErrAuthError
	}

	var hash []byte
	if sig[crypto.RecoveryIDOffset] > 1 {
		// Legacy Keccak256
		// Transform yellow paper V from 27/28 to 0/1
		sig[crypto.RecoveryIDOffset] -= 27
		hash = accounts.TextHash(nonce)
	} else {
		hash = crypto.Keccak256(nonce)
	}

	// Verify signature
	sigPublicKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		logger.I.Error("Authenticate.SigToPub",
			zap.Error(err),
			zap.String("hash", hexutil.Encode(hash)),
			zap.String("sig", hexutil.Encode(sig)),
		)
		return user, err
	}
	sigAddr := crypto.PubkeyToAddress(*sigPublicKey)

	if !strings.EqualFold(user.Address, sigAddr.Hex()) {
		logger.I.Error(
			"Authenticate: addresses are not equal",
			zap.String("sig.Address", sigAddr.Hex()),
			zap.String("user.Address", user.Address),
		)
		return user, ErrAuthError
	}

	// update the nonce here so that the signature cannot be resused
	nonce, err = GetNonce()
	if err != nil {
		logger.I.Error("Authenticate.GetNonce", zap.Error(err))
		return user, err
	}
	user.Nonce = nonce
	logger.I.Debug("New nonce for user", zap.String("nonce", hexutil.Encode(nonce)), zap.String("user", user.Address))
	if err := storage.Update(user); err != nil {
		logger.I.Error("Authenticate.Update", zap.Error(err))
		return user, err
	}

	return user, nil
}
