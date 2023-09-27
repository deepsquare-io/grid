// Copyright (C) 2023 DeepSquare Asociation
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
	"github.com/deepsquare-io/grid/sbatch-service/logger"
	"github.com/go-playground/validator/v10"
)

var I *validator.Validate

func init() {
	I = validator.New()
	if err := I.RegisterValidation("valid_container_image_url", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return ContainerURLValidator(value)
	}); err != nil {
		logger.I.Panic("failed to register valid_container_image_url")
	}
	if err := I.RegisterValidation("valid_envvar_name", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return EnvVarNameValidator(value)
	}); err != nil {
		logger.I.Panic("failed to register valid_envvar_name")
	}
	if err := I.RegisterValidation("alphanum_underscore", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return AlphaNumUnderscoreValidator(value)
	}); err != nil {
		logger.I.Panic("failed to register alphanum_underscore")
	}
}
