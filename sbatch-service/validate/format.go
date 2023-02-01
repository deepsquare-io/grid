package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Format(err validator.FieldError) string {
	switch err.Tag() {
	case "url":
		return fmt.Sprintf("'%s' is not a valid URL (error: '%s')", err.Field(), err.Tag())
	case "startswith":
		return fmt.Sprintf("'%s' must start with '%s' (error: '%s')", err.Field(), err.Value(), err.Tag())
	case "endsnotwith":
		return fmt.Sprintf("'%s' must not end with '%s' (error: '%s')", err.Field(), err.Value(), err.Tag())
	case "required":
		return fmt.Sprintf("'%s' is required (error: '%s')", err.Field(), err.Tag())
	case "valid_envvar_name":
		return fmt.Sprintf("'%s' is not a valid environment variable (error: '%s')", err.Field(), err.Tag())
	case "lt":
		return fmt.Sprintf("'%s' must be less than '%s' (error: '%s')", err.Field(), err.Value(), err.Tag())
	case "gte":
		return fmt.Sprintf("'%s' must be greater than or equal to '%s' (error: '%s')", err.Field(), err.Value(), err.Tag())
	case "valid_container_image_url":
		return fmt.Sprintf("'%s' is not a valid container image url (error: '%s')", err.Field(), err.Tag())
	case "oneof":
		return fmt.Sprintf("'%s' must be one of '%s' (error: '%s')", err.Field(), err.Value(), err.Tag())
	case "hostname":
		return fmt.Sprintf("'%s' is not a valid hostname (error: '%s')", err.Field(), err.Tag())
	}
	return err.Error()
}
