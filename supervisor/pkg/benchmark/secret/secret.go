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
