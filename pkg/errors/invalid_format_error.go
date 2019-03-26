package errors

import "fmt"

type InvalidFormatError struct {
	field   string
	message string
	cause   string
}

func NewInvalidFormatError(field, message, cause string) *InvalidFormatError {
	return &InvalidFormatError{
		field:   field,
		message: message,
		cause:   cause,
	}
}

func (e *InvalidFormatError) Error() string {
	return fmt.Sprintf("Field %s is invalid, %s : %s ", e.field, e.message, e.cause)
}
