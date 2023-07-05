package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gt2rz/micro-auth/internal/constants"
	"github.com/gt2rz/micro-auth/internal/models"
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

		// Check if the user already exists
		_, err = server.UserRepository.GetUserByEmail(context.Background(), request.Email)
		if err != constants.ErrUserNotFound {
			utils.SendHttpResponseError(w, constants.ErrUserAlreadyExists, http.StatusBadRequest)
			return
		}

		// Generate from the password with a predefined cost
		hashedPassword, err := utils.HashPassword(request.Password)
		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrSignupFailed, http.StatusInternalServerError)
			return
		}

		// Create a new user
		err = server.UserRepository.SaveUser(context.Background(), models.User{
			Email:     request.Email,
			Password:  string(hashedPassword),
			Firstname: request.Firstname,
			Lastname:  request.Lastname,
			Phone:     request.Phone,
		})

		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrSignupFailed, http.StatusBadRequest)
			return
		}

		// Send the response
		utils.SendHttpResponse(w, SignupResponse{
			Status:  constants.Success,
			Id:      1,
			Message: constants.UserCreatedSuccessfully,
		}, http.StatusCreated)
	}
}
