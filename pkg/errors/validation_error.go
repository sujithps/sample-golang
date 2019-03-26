package errors

import (
	"strings"
)

type ValidationError struct {
	validationErrors []error
	code             string
}

func NewValidationError(validationErrors []error) *ValidationError {
	return &ValidationError{
		code:             "validation_error",
		validationErrors: validationErrors,
	}
}

func (e *ValidationError) Error() string {
	var errorMessages []string
	for _, s := range e.validationErrors {
		errorMessages = append(errorMessages, s.Error())
	}
	return strings.Join(errorMessages, ",")
}

func (e *ValidationError) Errors() []error {
	return e.validationErrors
}
