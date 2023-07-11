package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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

type SignupCreateEvent struct {
	Type      string `json:"type"` // signup.create
	Email     string `json:"email"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Phone     string `json:"phone"`
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

		// Generate a new UUID
		id, err := uuid.NewRandom()
		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrSignupFailed, http.StatusInternalServerError)
			return
		}

		newUser := models.User{
			Id:        id.String(),
			Email:     request.Email,
			Password:  string(hashedPassword),
			Firstname: request.Firstname,
			Lastname:  request.Lastname,
			Phone:     request.Phone,
		}

		// Create a new user
		err = server.UserRepository.SaveUser(context.Background(), newUser)

		if err != nil {
			utils.SendHttpResponseError(w, constants.ErrSignupFailed, http.StatusBadRequest)
			return
		}

		// Publish the event
		payload := SignupCreateEvent{
			Type:      constants.Event_signup_create_success,
			Email:     newUser.Email,
			Firstname: newUser.Firstname,
			Lastname:  newUser.Lastname,
			Phone:     newUser.Phone,
		}
		newUserString := utils.StructToJsonString(payload)
		server.Publisher.Publish(context.Background(), constants.Channel_signup_create, newUserString)

		// Send the response
		utils.SendHttpResponse(w, SignupResponse{
			Status:  constants.Success,
			Id:      1,
			Message: constants.UserCreatedSuccessfully,
		}, http.StatusCreated)
	}
}
