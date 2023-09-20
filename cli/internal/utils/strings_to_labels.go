package utils

import (
	"fmt"
	"strings"

	"github.com/deepsquare-io/the-grid/cli/types"
)

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
