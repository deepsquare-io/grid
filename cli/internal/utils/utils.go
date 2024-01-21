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
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
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
			orPath = filepath.Join(home, ".dps", "key")
		}

		finfo, err := os.Stat(orPath)
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Hey! It looks like you didn't set the private key of your wallet.")
			fmt.Printf(
				"You can fetch your MetaMask private key by following this guide:\nhttps://support.metamask.io/hc/en-us/articles/360015289632-How-to-export-an-account-s-private-key .\n\n",
			)
			fmt.Printf(
				"After that, you can put the Hexadecimal-encoded private key at `%s`.\nAnd make sure the key has the right permissions (no world nor group permissions: chmod 600 %s)!\n\n",
				orPath,
				orPath,
			)
			input := confirmation.New(
				fmt.Sprintf("Or, do you prefer to generate a private key at `%s`?", orPath),
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
			fmt.Println(renderOutput(keyb, crypto.PubkeyToAddress(key.PublicKey).String()))
			fmt.Println(
				"You can fetch free credits by filling this form:\nhttps://share-eu1.hsforms.com/1PVlRXYdMSdy-iBH_PXx_0wev6gi",
			)

		} else if err != nil {
			return nil, err
		}
		// Check os is not windows
		if os.PathSeparator != '\\' {
			if !(finfo.Mode()&0700 > 0 && (finfo.Mode()&0070 == 0) && (finfo.Mode()&0007 == 0)) {
				if err := os.Chmod(orPath, 0600); err != nil {
					internallog.I.Error("Couldn't chmod to read-only.", zap.Error(err))
					internallog.I.Sugar().Errorf(
						"Permission of %s is insecure! Please `chmod 600 %s`.\n",
						orPath,
						orPath,
					)
				}

				return nil, errors.New("insecure file permission")
			}
		} else {
			internallog.I.Sugar().Warnf("You are running on Windows. Please make sure the file permission of %s is secure.\n", orPath)
		}
		b, err := os.ReadFile(orPath)
		if err != nil {
			return pk, err
		}
		b = bytes.TrimSpace(b)
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

var primaryColor = lipgloss.Color("#9202de")
var borderStyle = lipgloss.NewStyle().Foreground(primaryColor)
var cellStyle = lipgloss.NewStyle().PaddingRight(1).PaddingLeft(1).Bold(true)

func renderOutput(hexkey string, publicAddress string) string {
	rows := [][]string{
		{"Private Key", hexkey},
		{"Public Address", publicAddress},
	}

	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		BorderStyle(borderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			return cellStyle
		}).
		Rows(rows...).
		Render() +
		"\n"
}
