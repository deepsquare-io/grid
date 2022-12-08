package validate

import (
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/go-playground/validator/v10"
)

var I *validator.Validate

func init() {
	I = validator.New()
	if err := I.RegisterValidation("valid_container_image_url", func(fl validator.FieldLevel) bool {
		url := fl.Field().String()
		return ContainerURLValidator(url)
	}); err != nil {
		logger.I.Panic("failed to register valid_container_image_url")
	}
	if err := I.RegisterValidation("valid_envvar_name", func(fl validator.FieldLevel) bool {
		url := fl.Field().String()
		return EnvVarNameValidator(url)
	}); err != nil {
		logger.I.Panic("failed to register valid_envvar_name")
	}
}
