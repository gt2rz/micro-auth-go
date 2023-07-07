package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gt2rz/micro-auth/internal/constants"
	"github.com/gt2rz/micro-auth/internal/servers"
	"github.com/gt2rz/micro-auth/internal/utils"

	"github.com/gorilla/mux"
)

type ResetPasswordRequest struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (r *ResetPasswordRequest) Validate() error {
	if !(len(strings.TrimSpace(r.Password)) > 0) {
		return utils.ErrInvalidRequest
	}

	if !(len(strings.TrimSpace(r.PasswordConfirmation)) > 0) {
		return utils.ErrInvalidRequest
	}

	if r.Password != r.PasswordConfirmation {
		return utils.ErrInvalidRequest
	}

	return nil
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func ResetPasswordHandler(s *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get the reset token from the request
		params := mux.Vars(r)
		resetToken := params["resetToken"]

		// Check if the reset token is empty
		if !(len(strings.TrimSpace(resetToken)) > 0) {
			utils.SendHttpResponseError(w, utils.ErrInvalidRequest, http.StatusBadRequest)
			return
		}

		// Get the user by reset token
		user, err := s.UserRepository.GetUserByResetToken(context.Background(), resetToken)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrAnErrorOccurred, http.StatusInternalServerError)
			return
		}

		// Check if the reset token is expired
		if time.Until(user.PasswordResetTokenAt).Minutes() < 0 {
			utils.SendHttpResponseError(w, constants.ErrResetTokenExpired, http.StatusUnauthorized)
			return
		}

		// Check if user is not found
		if user == nil {
			utils.SendHttpResponseError(w, constants.ErrInvalidCredentials, http.StatusUnauthorized)
			return
		}

		// Decode the request body into the struct and failed if any error occur
		var request = ResetPasswordRequest{}
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrInvalidRequest, http.StatusBadRequest)
			return
		}

		// Validate the request
		err = request.Validate()
		if err != nil {
			utils.SendHttpResponseError(w, err, http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(request.Password)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrAnErrorOccurred, http.StatusInternalServerError)
			return
		}

		// Update the user's password
		err = s.UserRepository.UpdatePassword(context.Background(), user.Id, hashedPassword)
		if err != nil {
			utils.SendHttpResponseError(w, utils.ErrAnErrorOccurred, http.StatusInternalServerError)
			return
		}

		// Return the response
		utils.SendHttpResponse(w, ResetPasswordResponse{
			Message: "Reset password successfully",
			Status:  true,
		}, http.StatusOK)
	}
}
