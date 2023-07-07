package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gt2rz/micro-auth/internal/constants"
	"github.com/gt2rz/micro-auth/internal/servers"
	"github.com/gt2rz/micro-auth/internal/utils"
)

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
}

func ForgotPasswordHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = ForgotPasswordRequest{}

		// Decode the request body into the struct and failed if any error occur
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrInvalidRequest, http.StatusBadRequest)
			return
		}

		// Get the user by email
		user, err := s.UserRepository.GetUserByEmail(context.Background(), request.Email)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrAnErrorOccurred, http.StatusInternalServerError)
			return
		}

		// Check if user is not found
		if user == nil {
			utils.SendHttpResponseError(w, constants.ErrInvalidCredentials, http.StatusUnauthorized)
			return
		}

		// Generate a reset token
		resetToken, err := s.UserRepository.GenerateResetToken(context.Background(), user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(resetToken)
		// Return the response
		utils.SendHttpResponse(w, &ForgotPasswordResponse{
			Message: "You will receive a reset email if user with that email exist",
		}, http.StatusOK)
	}
}
