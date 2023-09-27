// Copyright (C) 2023 DeepSquare Association
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

package crypto_test

import (
	"bytes"
	"testing"

	"github.com/deepsquare-io/grid/grid-logger/server/crypto"
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
