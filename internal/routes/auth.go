package routes

import (
	"github.com/gorilla/mux"
	"github.com/gt2rz/micro-auth/internal/handlers/auth"
	"github.com/gt2rz/micro-auth/internal/servers"
)

func AuthRoutes(server *servers.HttpServer, router *mux.Router) {
	router.HandleFunc("/auth/signup", auth.SignupHandler(server)).Methods("POST")
	router.HandleFunc("/auth/login", auth.LoginHandler(server)).Methods("POST")
	router.HandleFunc("/auth/forgotpassword", auth.ForgotPasswordHandler(server)).Methods("POST")
	router.HandleFunc("/auth/resetpassword/{resetToken}", auth.ResetPasswordHandler(server)).Methods("PATCH")

}
