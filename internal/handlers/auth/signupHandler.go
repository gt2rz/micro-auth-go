package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gt2rz/micro-auth/internal/constants"
	"github.com/gt2rz/micro-auth/internal/servers"
	"github.com/gt2rz/micro-auth/internal/utils"
)

type SignupRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type SignupResponse struct {
	Status  bool   `json:"status"`
	Id      int    `json:"id"`
	Message string `json:"message"`
}

func SignupHandler(server *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignupRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrSignupFailed, http.StatusBadRequest)
			return
		}

		utils.SendHttpResponse(w, SignupResponse{
			Status:  constants.Success,
			Id:      1,
			Message: constants.UserCreatedSuccessfully,
		}, http.StatusCreated)
	}
}
