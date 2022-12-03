package auth

import (
	"bytes"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func Authenticate(storage *MemStorage, address string, nonce []byte, sig []byte) (User, error) {
	user, err := storage.Get(address)
	if err != nil {
		return user, err
	}

	// Verify nonce
	if !bytes.Equal(user.Nonce, nonce) {
		return user, ErrAuthError
	}

	// Verify signature
	hash := crypto.Keccak256Hash(nonce)
	sigPublicKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}
	sigAddr := crypto.PubkeyToAddress(*sigPublicKey)

	if !strings.EqualFold(user.Address, sigAddr.Hex()) {
		return user, ErrAuthError
	}

	// update the nonce here so that the signature cannot be resused
	nonce, err = GetNonce()
	if err != nil {
		return user, err
	}
	user.Nonce = nonce
	if err := storage.Update(user); err != nil {
		return user, err
	}

	return user, nil
}
