package validate

import (
	"fmt"
	"regexp"
)

const (
	regexUser     = `[[:alnum:]_.!~*\'()%\;:\&=+$,-@]+`
	regexRegistry = `[^#]+`
	regexImage    = `[[:lower:][:digit:]/._-]+`
	regexTag      = `[[:alnum:]._:-]+`
)

var (
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
	return regexContainerURL.MatchString(url)
}
