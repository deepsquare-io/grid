// Copyright (C) 2023 DeepSquare Asociation
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

package utils

import (
	"fmt"
	"strings"

	"github.com/deepsquare-io/the-grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
)

func FormatLabels(labels []metaschedulerabi.Label) []string {
	out := make([]string, 0, len(labels))
	for _, l := range labels {
		out = append(out, fmt.Sprintf("%s: %s", l.Key, l.Value))
	}
	return out
}

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
