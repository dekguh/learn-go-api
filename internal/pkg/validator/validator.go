package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := e.Field()
			switch e.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s is required", field))
			case "email":
				errors = append(errors, fmt.Sprintf("%s is not a valid email", field))
			case "min":
				errors = append(errors, fmt.Sprintf("%s must be at least %s characters long", field, e.Param()))
			case "max":
				errors = append(errors, fmt.Sprintf("%s must be at most %s characters long", field, e.Param()))
			default:
				errors = append(errors, fmt.Sprintf("%s is not valid", field))
			}
		}
	} else {
		errors = append(errors, err.Error())
	}

	return errors
}

func JoinErrorValidation(errors []string) string {
	return strings.Join(errors, ", ")
}
