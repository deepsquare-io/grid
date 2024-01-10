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

// Package validator providers utilities for user input validation.
package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var isMapRegex = regexp.MustCompile(
	`^([a-zA-Z0-9.-]+=[a-zA-Z0-9.-]+,)*([a-zA-Z0-9.-]+=[a-zA-Z0-9.-]+)$`,
)

// IsNumber checks if the string is a number.
func IsNumber(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid number: %s", s)
	}
	return nil
}

// IsMap checks if the string can be parsed as a map.
//
// Example: "key=value,key2=value2" is a map.
func IsMap(input string) error {
	if input == "" {
		// Empty string is an empty map
		return nil
	}

	// Check if the input matches the pattern
	if !isMapRegex.MatchString(input) {
		return fmt.Errorf("invalid map: illegal characters, only a-zA-Z0-9.- are accepted")
	}

	// Split the input string by comma separator
	items := strings.Split(input, ",")

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
			return fmt.Errorf("invalid map: missing value for key %s", parts[0])
		}
	}

	return nil
}

// IsNumberChar returns true when the character is allowed when writing a number.
func IsNumberChar(ch rune) bool {
	return unicode.IsDigit(ch) || ch == 'e' || ch == '.' || ch == '-'
}

// AllowedNumberChar checks if a string contains allowed number characters.
func AllowedNumberChar(input string) error {
	for _, ch := range input {
		if !IsNumberChar(ch) {
			return fmt.Errorf("character '%c' is not allowed", ch)
		}
	}
	return nil
}

// AllowedHexChar checks if a string contains allowed hex characters.
func AllowedHexChar(input string) error {
	for _, ch := range input {
		if !unicode.Is(unicode.Hex_Digit, ch) && ch != 'x' {
			return fmt.Errorf("character '%c' is not allowed", ch)
		}
	}
	return nil
}
