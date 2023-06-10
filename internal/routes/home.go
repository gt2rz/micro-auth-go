package routes

import (
	"github.com/gorilla/mux"
	"github.com/gt2rz/micro-auth/internal/handlers/home"
	"github.com/gt2rz/micro-auth/internal/servers"
)

func HomeRoutes(server *servers.HttpServer, route *mux.Router) {
	route.HandleFunc("/", home.HomeHandler(server)).Methods("GET")
}
