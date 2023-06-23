package utils

import (
	"fmt"
	"strings"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
)

func StringsToLabels(input string) ([]metaschedulerabi.Label, error) {
	if input == "" {
		// Empty string is an empty map
		return []metaschedulerabi.Label{}, nil
	}

	items := strings.Split(input, ",")

	labels := make([]metaschedulerabi.Label, 0, len(items))

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
			return []metaschedulerabi.Label{}, fmt.Errorf(
				"invalid map: missing value for key %s",
				parts[0],
			)
		}

		labels = append(labels, metaschedulerabi.Label{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	return labels, nil
}
