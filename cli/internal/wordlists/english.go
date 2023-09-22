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
