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
