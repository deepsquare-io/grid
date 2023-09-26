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
