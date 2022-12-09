package validate

import "regexp"

var (
	regexEnvVarName = regexp.MustCompilePOSIX(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
)

func EnvVarNameValidator(url string) bool {
	return regexEnvVarName.MatchString(url)
}
