package home

import (
	"encoding/json"
	"net/http"

	"github.com/gt2rz/micro-auth/internal/servers"
)

func HomeHandler(server *servers.HttpServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Welcome to the home page!")
	}
}
