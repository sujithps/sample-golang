package errors

import "fmt"

type NotFoundError struct {
	entityName  string
	entityValue string
}

func NewNotFoundError(name, value string) *NotFoundError {
	return &NotFoundError{
		entityName:  name,
		entityValue: value,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s %s not found.", e.entityName, e.entityValue)
}

func IsNotFoundError(t interface{}) bool {
	switch t.(type) {
	case *NotFoundError:
		return true
	default:
		return false
	}
}
