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

	"github.com/go-playground/validator/v10"
)

func Format(err validator.FieldError) string {
	switch err.Tag() {
	case "url":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid URL {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "ip":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid IP {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "cidr":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid CIDR {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "startswith":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must start with '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Param(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "endsnotwith":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must not end with '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Param(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "required":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is required {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "valid_envvar_name":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid environment variable {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "lt":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must be less than '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Param(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "gte":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must be greater than or equal to '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Param(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "ne":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must be not equal to '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Param(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "valid_container_image_url":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid container image url {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "oneof":
		return fmt.Sprintf(
			"Field '%s' (='%+v') must be one of '%s' {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Param(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "hostname":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid hostname {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "hostname_port":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a valid hostname:port {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "alphanum_underscore":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not a alphanumeric with or without underscore value {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	case "ip|fqdn":
		return fmt.Sprintf(
			"Field '%s' (='%+v') is not an IP address nor an FQDN {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	default:
		return fmt.Sprintf(
			"Field '%s' (='%+v') doesn't respect the '%s' rule {\"error\": \"%s\", \"field\": \"%s\", \"value\": \"%+v\"}",
			err.Field(),
			err.Value(),
			err.Tag(),
			err.Tag(),
			err.Field(),
			err.Value(),
		)
	}
}
