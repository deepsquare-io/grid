// Copyright (C) 2023 DeepSquare Asociation
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

package hash

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateAlphanumeric(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	// Convert the hash bytes to base64-encoded string
	hashBase64 := base64.URLEncoding.EncodeToString(hashBytes)

	// Remove non-alphanumeric characters from the base64 string
	var alphanumericHash string
	for _, char := range hashBase64 {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') {
			alphanumericHash += string(char)
		}
	}

	return alphanumericHash
}
