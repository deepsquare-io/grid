package validate

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	regexUser     = `[[:alnum:]_.!~*\'()%\;:\&=+$,-@]+`
	regexRegistry = `[^#]+`
	regexImage    = `[[:lower:][:digit:]/._-]+`
	regexTag      = `[[:alnum:]._:-]+`
)

var (
	// regexContainerURL is a matcher from https://github.com/NVIDIA/enroot/blob/master/src/docker.sh
	regexContainerURL = regexp.MustCompilePOSIX(
		fmt.Sprintf(
			"^((%s)@)?((%s)#)?(%s)(:(%s))?$",
			regexUser,
			regexRegistry,
			regexImage,
			regexTag,
		),
	)
)

func ContainerURLValidator(url string) bool {
	return strings.HasPrefix(url, "/") || regexContainerURL.MatchString(url)
}
