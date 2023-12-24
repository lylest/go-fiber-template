package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validate(modelInterface any) (string, error) {
	validate := validator.New()
	validationErrors := validate.Struct(modelInterface)

	if validationErrors != nil {
		var formattedErrors validator.ValidationErrors
		errors.As(validationErrors, &formattedErrors)

		var firstValidationError string
		for _, errObject := range formattedErrors {
			firstValidationError = fmt.Sprintf("%s: %s", errObject.Field(), errObject.Tag())
			break
		}

		return firstValidationError, formattedErrors
	}

	return "", nil
}
