package validator

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidationIsCool(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
