// Copyright (C) 2023 DeepSquare Association
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

package osu

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseOSULog(reader io.Reader) (float64, error) {
	scanner := bufio.NewScanner(reader)

	// Initialize a variable to store the last value
	var lastValue float64

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by whitespace
		parts := strings.Fields(line)

		// Check if the line has at least two parts
		if len(parts) >= 2 {
			// Convert the second part to a float64
			value, err := strconv.ParseFloat(parts[1], 64)
			if err == nil {
				// Update the lastValue if successful
				lastValue = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lastValue, nil
}
