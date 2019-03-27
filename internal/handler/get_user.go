package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"git.thoughtworks.net/mahadeva/sample-golang/internal/service/userservice"
	"git.thoughtworks.net/mahadeva/sample-golang/pkg/contract"
)

func GetUser(userService userservice.Client) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) (*contract.Response, error) {
		vars := mux.Vars(r)
		userID := vars["id"]

		user, err := userService.GetByUserID(r.Context(), userID)
		if err != nil {
			return nil, err
		}

		return contract.NewSuccessResponse(contract.NewUser(user)), nil
	}
}
