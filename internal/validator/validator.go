package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(data interface{}) error {
	if err := v.validate.Struct(data); err != nil {
		return v.formatValidationError(err)
	}
	return nil
}

func (v *Validator) formatValidationError(err error) error {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()

		var message string
		switch tag {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters", field, err.Param())
		case "oneof":
			message = fmt.Sprintf("%s must be one of: %s", field, err.Param())
		default:
			message = fmt.Sprintf("%s is invalid", field)
		}

		errors = append(errors, message)
	}

	return fmt.Errorf(strings.Join(errors, ", "))
}
