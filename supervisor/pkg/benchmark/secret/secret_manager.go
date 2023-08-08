package secret

import (
	"bytes"
	"crypto/rand"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
)

type Manager interface {
	Get() []byte
	Validate([]byte) bool
}

type manager struct {
	secret []byte
}

func generateSecret() []byte {
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	if err != nil {
		logger.I.Panic("failed to create secret")
	}
	return secret
}

func NewManager() Manager {
	return &manager{
		secret: generateSecret(),
	}
}

func (m *manager) Get() []byte {
	return m.secret
}

func (m *manager) Validate(data []byte) bool {
	return bytes.Equal(m.secret, data)
}
