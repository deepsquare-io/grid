package validate

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var LDAPUserIllegalCharacters = []rune{
	'\\', '#', '+', '<', '>', ',', ';', '"', '=',
}

var I *validator.Validate

func init() {
	I = validator.New()
}

func LDAPUserIsValid(user string) string {
	for _, r := range LDAPUserIllegalCharacters {
		if strings.ContainsRune(user, r) {
			return fmt.Sprintf("contains illegal character: %c", r)
		}
	}
	if user[0] == ' ' {
		return "contains leading spaces"
	}
	if user[len(user)] == ' ' {
		return "contains trailing spaces"
	}

	return ""
}
