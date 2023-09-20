package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) []string {
	var errors []string
	for _, fieldErr := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldErr.Field(), fieldErr.Tag()))
	}

	return errors
}
