package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func Sign(pk *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), pk)
	if err != nil {
		return []byte{}, err
	}

	return signature, nil
}
