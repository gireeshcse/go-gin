package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateUsername checks whether it is a valid username or not
func ValidateUsername(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "user")
}
