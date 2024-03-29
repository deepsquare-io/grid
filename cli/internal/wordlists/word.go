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

// Package wordlists providers utilities to generate random words.
package wordlists

import "math/rand"

// GetRandomWords generates random words from the bip39.
func GetRandomWords(length int) []string {
	words := make([]string, length)
	for i := 0; i < length; i++ {
		words[i] = English[rand.Intn(len(English))]
	}
	return words
}
