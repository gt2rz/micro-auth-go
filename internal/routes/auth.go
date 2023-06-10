package routes

import (
	"github.com/gorilla/mux"
	"github.com/gt2rz/micro-auth/internal/servers"
	"github.com/gt2rz/micro-auth/internal/handlers/auth"
)

func AuthRoutes(server *servers.HttpServer, router *mux.Router) {
	router.HandleFunc("/auth/signup", auth.SignupHandler(server)).Methods("POST")
}
