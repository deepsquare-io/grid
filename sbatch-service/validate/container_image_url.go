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

package validate

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	regexImage = `[[:lower:][:digit:]/._-]+`
	regexTag   = `[[:alnum:]._:-]+`
)

var (
	// regexContainerURL is a matcher from https://github.com/NVIDIA/enroot/blob/master/src/docker.sh
	regexContainerURL = regexp.MustCompilePOSIX(
		fmt.Sprintf(
			"^(%s)(:(%s))?$",
			regexImage,
			regexTag,
		),
	)
)

func ContainerURLValidator(url string) bool {
	return strings.HasPrefix(url, "/") || regexContainerURL.MatchString(url)
}
