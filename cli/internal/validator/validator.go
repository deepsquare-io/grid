package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var isMapRegex = regexp.MustCompile(
	`^([a-zA-Z0-9.-]+=[a-zA-Z0-9.-]+,)*([a-zA-Z0-9.-]+=[a-zA-Z0-9.-]+)$`,
)

func IsNumber(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid number: %s", s)
	}
	return nil
}

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
