package errors

import "fmt"

type RequiredFieldMisingError struct {
	fieldName string
}

func NewRequiredFieldMisingError(fieldName string) *RequiredFieldMisingError {
	return &RequiredFieldMisingError{fieldName: fieldName}
}

func (e *RequiredFieldMisingError) Error() string {
	return fmt.Sprintf("Field %s cannot be empty.", e.fieldName)
}
