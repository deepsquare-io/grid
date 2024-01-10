// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package wordlists

import (
	"fmt"
	"hash/crc32"
	"strings"

	_ "embed"
)

func init() {
	// Ensure word list is correct
	// $ wget https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt
	// $ crc32 english.txt
	// c1dbd296
	checksum := crc32.ChecksumIEEE([]byte(english))
	if fmt.Sprintf("%x", checksum) != "c1dbd296" {
		panic("english checksum invalid")
	}
}

// English is a slice of mnemonic words taken from the bip39 specification
// https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt
var English = strings.Split(strings.TrimSpace(english), "\n")

//go:embed english.txt
var english string
