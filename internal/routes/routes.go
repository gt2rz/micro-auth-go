package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/gt2rz/micro-auth/internal/servers"
)

func InitRoutes(server *servers.HttpServer, router *mux.Router) {
	fmt.Println("Initializing routes...")
	// Home routes
	HomeRoutes(server, router)

	// Auth routes
	AuthRoutes(server, router)
}
