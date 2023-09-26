package secret

import (
	"bytes"
	"crypto/rand"

	"github.com/deepsquare-io/grid/supervisor/logger"
)

var secret = generateSecret()

func generateSecret() []byte {
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	if err != nil {
		logger.I.Panic("failed to create secret")
	}
	return secret
}

func Get() []byte {
	return secret
}

func Validate(data []byte) bool {
	return bytes.Equal(secret, data)
}
