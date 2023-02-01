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
