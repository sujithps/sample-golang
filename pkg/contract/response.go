package contract

import (
	"github.com/sujithps/sample-golang/pkg/errors"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
	Success    bool        `json:"success"`
	Errors     []Error     `json:"errors,omitempty"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		StatusCode: http.StatusAccepted,
		Data:       data,
		Success:    true,
		Errors:     []Error{},
	}
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewError(code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func NewInternalServerErrorResponse() Response {
	return Response{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Errors: []Error{
			{Code: "internal_error", Message: "server error"},
		},
	}
}

func NewNotFoundResponse(msg string) Response {
	return Response{
		StatusCode: http.StatusNotFound,
		Success:    false,
		Errors: []Error{
			{Code: "not_found", Message: "not_found"},
		},
	}
}

func NewInvalidRequest(msg string) Response {
	return Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
		Errors: []Error{
			{Code: "invalid_request", Message: "msg"},
		},
	}
}

func NewValidationErrorResponse(validationError *errors.ValidationError) Response {
	var errorResponse []Error
	for _, validationError := range validationError.Errors() {
		errorResponse = append(errorResponse, *NewError("909", validationError.Error()))
	}
	return Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
		Errors:     errorResponse,
	}
}
