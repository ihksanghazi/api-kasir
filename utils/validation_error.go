package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) []string {
	var errors []string
	for _, fieldErr := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("%s Must Be %s", fieldErr.Field(), fieldErr.Tag()))
	}

	return errors
}
