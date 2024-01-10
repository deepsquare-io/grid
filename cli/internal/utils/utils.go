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

// Package utils provides utilities functions.
package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// BoolToYN converts a boolean to "yes" or "no".
func BoolToYN(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// YNToBool converts a yes/no to bool. Defaults to "no" on failure.
func YNToBool(b string) bool {
	return strings.Contains(b, "yes")
}

// ErrorfOrEmpty returns a message if the error is not nil.
func ErrorfOrEmpty(msg string, err error) string {
	if err != nil {
		return msg
	}
	return ""
}

// FormatErrorfOrEmpty returns a formatted message if the error is not nil.
func FormatErrorfOrEmpty(format string, err error, va ...any) string {
	if err != nil {
		a := append([]any{err}, va...)
		return fmt.Sprintf(format, a...)
	}
	return ""
}

// FormatLabels formats labels into a slice of "key: value".
func FormatLabels(labels []metaschedulerabi.Label) []string {
	out := make([]string, 0, len(labels))
	for _, l := range labels {
		out = append(out, fmt.Sprintf("%s: %s", l.Key, l.Value))
	}
	return out
}

// StringsToLabels converts a "key=value,key2=value2" into a slice of labels.
func StringsToLabels(input string) ([]types.Label, error) {
	if input == "" {
		// Empty string is an empty map
		return []types.Label{}, nil
	}

	items := strings.Split(input, ",")

	labels := make([]types.Label, 0, len(items))

	// Iterate over each item
	for _, item := range items {
		// Split the item by the equals sign separator
		parts := strings.Split(item, "=")

		// Check if the item has two parts
		if len(parts) == 0 {
			// Ignore enpty parts
			continue
		}
		if len(parts) != 2 {
			return []types.Label{}, fmt.Errorf(
				"invalid map: missing value for key %s",
				parts[0],
			)
		}

		labels = append(labels, types.Label{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	return labels, nil
}

// GetPrivateKey fetch or generates an ethereum private key.
func GetPrivateKey(ethHexPK, orPath string) (*ecdsa.PrivateKey, error) {
	var pk *ecdsa.PrivateKey
	if ethHexPK == "" {
		// Default dps path
		if orPath == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return pk, err
			}
			orPath = filepath.Join(home, "/.dps/key")
		}

		finfo, err := os.Stat(orPath)
		if errors.Is(err, fs.ErrNotExist) {
			internallog.I.Sugar().Errorf(
				"Ethereum private key not found at path %s.\n",
				orPath,
			)

			input := confirmation.New(
				fmt.Sprintf("Do you wish to generate a private key at `%s`?", orPath),
				confirmation.No,
			)
			ok, prompterr := input.RunPrompt()
			if prompterr != nil {
				return nil, prompterr
			}
			if !ok {
				return nil, err
			}
			key, err := crypto.GenerateKey()
			if err != nil {
				return nil, err
			}
			keyb := hexutil.Encode(crypto.FromECDSA(key))
			if err := os.MkdirAll(filepath.Dir(orPath), 0700); err != nil {
				panic(err)
			}
			if err := os.WriteFile(orPath, []byte(keyb), 0600); err != nil {
				return nil, err
			}
			finfo, err = os.Stat(orPath)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Private Key: %s\n", keyb)
			fmt.Printf("Public Address: %s\n", crypto.PubkeyToAddress(key.PublicKey))

		} else if err != nil {
			return nil, err
		}
		if !(finfo.Mode()&0700 > 0 && (finfo.Mode()&0070 == 0) && (finfo.Mode()&0007 == 0)) {
			internallog.I.Sugar().Errorf(
				"Permission of %s is insecure! Please `chmod 600 %s`.\n",
				orPath,
				orPath,
			)
			return nil, errors.New("insecure file permission")
		}
		b, err := os.ReadFile(orPath)
		if err != nil {
			return pk, err
		}
		ethHexPK = string(b)
	}
	kb, err := hexutil.Decode(ethHexPK)
	if errors.Is(err, hexutil.ErrMissingPrefix) {
		kb, err = hex.DecodeString(ethHexPK)
		if err != nil {
			return pk, err
		}
	} else if err != nil {
		return pk, err
	}
	pk, err = crypto.ToECDSA(kb)
	if err != nil {
		return pk, err
	}
	return pk, nil
}
