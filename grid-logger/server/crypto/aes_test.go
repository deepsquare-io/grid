package crypto_test

import (
	"bytes"
	"testing"

	"github.com/deepsquare-io/the-grid/grid-logger/server/crypto"
)

func TestEncryptDecrypt(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("error generating key: %v", err)
	}
	plaintext := []byte("test plaintext")
	ciphertext, err := crypto.Encrypt(key, plaintext)
	if err != nil {
		t.Errorf("error encrypting plaintext: %v", err)
	}
	decryptedText, err := crypto.Decrypt(key, ciphertext)
	if err != nil {
		t.Errorf("error decrypting ciphertext: %v", err)
	}
	if !bytes.Equal(decryptedText, plaintext) {
		t.Errorf("decrypted text (%s) does not match plaintext (%s)", decryptedText, plaintext)
	}
}
