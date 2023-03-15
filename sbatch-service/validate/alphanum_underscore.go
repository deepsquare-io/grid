package validate

import (
	"regexp"
)

var (
	regexAlphaNumUnderscore = regexp.MustCompilePOSIX(`^[[:alnum:]_]+$`)
)

func AlphaNumUnderscoreValidator(value string) bool {
	return regexAlphaNumUnderscore.MatchString(value)
}
