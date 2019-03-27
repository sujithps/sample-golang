package handler

import (
	"encoding/json"
	"fmt"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/constant"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/contract"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/errors"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/logger"
	"io/ioutil"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) (*contract.Response, error)

func Wrap(handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := handler(w, r)
		if err != nil {
			switch err.(type) {
			case *errors.NotFoundError:
				createNotFoundResponse(w, r, err)
				return
			case *errors.RequiredFieldMisingError:
				createInvalidRequest(w, r, err)
				return
			case *errors.InvalidFormatError:
				createInvalidRequest(w, r, err)
				return
			case *errors.ValidationError:
				validationError := err.(*errors.ValidationError)
				createValidationErrorResponse(w, r, validationError)
				return
			case *errors.MarshallingError:
				createInvalidRequest(w, r, err)
				return
			case *errors.InvalidRequestError:
				createInvalidRequest(w, r, err)
				return
			default:
				createInternalServerErrorResponse(w, r, err)
				return
			}
		}
		if response != nil {
			resp, err := json.Marshal(response)
			if err != nil {
				logger.NewContextLogger(r.Context()).Error("ResponseMarshal", "handler: unable to marshal json response", err)
			}
			w.Write(resp)
		}
	}
}

func UnmarshalRequestBody(request *http.Request, object interface{}) error {
	if request.Body == nil {
		return errors.NewInvalidRequestError(constant.RequestBodyEmpty)
	}
	defer request.Body.Close()

	if request.ContentLength == 0 {
		return errors.NewInvalidRequestError(constant.RequestBodyEmpty)
	}

	buf, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.NewInvalidRequestError(constant.InvalidRequestBody)
	}

	err = json.Unmarshal(buf, object)
	if err != nil {
		return errors.NewInvalidRequestError(constant.ErrorWhileUnMarshalling)
	}
	return nil
}

func createInternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log := logger.NewContextLogger(r.Context())
	log.Error("Handler.Wrap", "Internal server error", err)
	resp := contract.NewInternalServerErrorResponse()
	writeResponse(w, r, resp)
}

func createNotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log := logger.NewContextLogger(r.Context())
	log.Error("Handler.Wrap", "entity not found", err)
	resp := contract.NewNotFoundResponse(err.Error())
	writeResponse(w, r, resp)
}

func createInvalidRequest(w http.ResponseWriter, r *http.Request, err error) {
	log := logger.NewContextLogger(r.Context())
	resp := contract.NewInvalidRequest(err.Error())
	log.Error("Handler.Wrap", "Invalid Request", err)
	writeResponse(w, r, resp)
}

func createValidationErrorResponse(w http.ResponseWriter, r *http.Request, validationError *errors.ValidationError) {
	log := logger.NewContextLogger(r.Context())
	resp := contract.NewValidationErrorResponse(validationError)
	log.Warn("Handler.Wrap", "Validation Error", validationError)
	writeResponse(w, r, resp)
}

func writeResponse(w http.ResponseWriter, r *http.Request, resp contract.Response) {
	fmt.Sprintln("**********")
	fmt.Println(resp.StatusCode)
	w.WriteHeader(resp.StatusCode)
	response, err := json.Marshal(resp)
	if err != nil {
		logger.NewContextLogger(r.Context()).Error("Handler.Wrap", "unable to marshal json response", err)
	}
	w.Write(response)
}
