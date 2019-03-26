package errors

type InvalidRequestError struct {
	message string
}

func NewInvalidRequestError(message string) *InvalidRequestError {
	return &InvalidRequestError{
		message: message,
	}
}

func (e *InvalidRequestError) Error() string {
	return e.message
}
