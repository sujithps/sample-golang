package errors

import "fmt"

type MarshallingError struct {
	err error
}

func NewMarshallingError(err error) *MarshallingError {
	return &MarshallingError{
		err: err,
	}
}

func (e *MarshallingError) Error() string {
	return fmt.Sprintf("Failed to marshall the object. %s", e.err.Error())
}
