package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Format(err validator.FieldError) string {
	switch err.Tag() {
	case "url":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid URL {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	case "ip":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid IP {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	case "cidr":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid CIDR {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	case "startswith":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must start with '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Param(), err.Tag(), err.Field(), err.Value(),
		)
	case "endsnotwith":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must not end with '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Param(), err.Tag(), err.Field(), err.Value(),
		)
	case "required":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is required {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	case "valid_envvar_name":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid environment variable {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	case "lt":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must be less than '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Param(), err.Tag(), err.Field(), err.Value(),
		)
	case "gte":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must be greater than or equal to '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Param(), err.Tag(), err.Field(), err.Value(),
		)
	case "valid_container_image_url":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid container image url {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	case "oneof":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must be one of '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Param(), err.Tag(), err.Field(), err.Value(),
		)
	case "hostname":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid hostname {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	case "alphanum_underscore":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a alphanumeric with underscore value {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(), err.Value(), err.Tag(), err.Field(), err.Value(),
		)
	}
	return err.Error()
}
