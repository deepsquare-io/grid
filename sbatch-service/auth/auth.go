// Package auth defines the authentication layer of the application.
package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/deepsquare-io/grid/sbatch-service/storage"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
)

const expirationDuration = 10 * time.Minute

const nonceLength = 16 // You can adjust the length as needed

func generateNonce() (string, error) {
	// Create a byte slice to store the random nonce
	nonce := make([]byte, nonceLength)

	// Use the crypto/rand package to generate random bytes
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}

	nonceString := base64.StdEncoding.EncodeToString(nonce)

	return nonceString, nil
}

// Auth is a service that provides HTTP handlers and middlewares used for authentication.
//
// It uses a time-based nonce. The nonce is encrypted with the private key.
type Auth struct {
	storage storage.Storage
}

// NewAuth builds the Auth struct.
func NewAuth(
	storage storage.Storage,
) *Auth {
	return &Auth{
		storage: storage,
	}
}

type challengeMessage struct {
	Message string `json:"message"`
	Nonce   string `json:"nonce"`
}

// Challenge returns a message with a nonce.
func (a *Auth) Challenge(ctx context.Context, message string) []byte {
	nonce, err := generateNonce()
	if err != nil {
		panic(err)
	}
	if err = a.storage.Set(ctx, "nonce-"+nonce, nonce, expirationDuration); err != nil {
		panic(err)
	}
	dat, err := json.Marshal(challengeMessage{
		Message: message,
		Nonce:   nonce,
	})
	if err != nil {
		panic(err)
	}
	return dat
}

// Verify checks the signature and nonce.
//
// This is a time-based nonce. In production, it is preferable to use a true nonce (random number) which is stored in a database.
func (a *Auth) Verify(ctx context.Context, address string, data []byte, sig []byte) error {
	var hash []byte
	if sig[ethcrypto.RecoveryIDOffset] > 1 {
		// Legacy Keccak256
		// Transform yellow paper V from 27/28 to 0/1
		sig[ethcrypto.RecoveryIDOffset] -= 27
	}
	hash = accounts.TextHash(data)

	// Verify signature
	sigPublicKey, err := ethcrypto.SigToPub(hash, sig)
	if err != nil {
		log.Err(err).
			Str("hash", hexutil.Encode(hash)).
			Str("sig", hexutil.Encode(sig)).
			Msg("SigToPub failed")
		return err
	}
	sigAddr := ethcrypto.PubkeyToAddress(*sigPublicKey)

	// Verify public key
	if !strings.EqualFold(address, sigAddr.Hex()) {
		log.Error().
			Str("sig.Address", sigAddr.Hex()).
			Str("address", address).
			Str("sig", hexutil.Encode(sig)).
			Str("expected hash", hexutil.Encode(hash)).
			Msg("addresses are not equal")
		return errors.New("authentication error: addresses are not equal")
	}

	// Verify message
	var msg challengeMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Err(err).
			Str("data", string(data)).
			Msg("invalid msg")
		return fmt.Errorf("authentication error: invalid msg: %w", err)
	}
	nonce, err := a.storage.Get(ctx, "nonce-"+msg.Nonce)
	if err != nil {
		log.Err(err).
			Str("data", string(data)).
			Msg("nonce failure")
		return fmt.Errorf("authentication error: nonce failure: %w", err)
	}
	if nonce != msg.Nonce {
		log.Error().
			Str("data", string(data)).
			Msg("nonce failed verification")
		return errors.New("authentication error: nonce failed verification")
	}

	return nil
}
