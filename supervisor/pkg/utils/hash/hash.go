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
