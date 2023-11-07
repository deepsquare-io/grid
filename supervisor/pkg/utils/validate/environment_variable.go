package validate

import "regexp"

var (
	// regexEnvVarName is a matcher from the POSIX standard
	regexEnvVarName = regexp.MustCompilePOSIX(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
)

func EnvVarNameValidator(env string) bool {
	return regexEnvVarName.MatchString(env)
}
