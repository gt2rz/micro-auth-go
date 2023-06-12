package servers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/gt2rz/micro-auth/internal/database"
)

type HttpServer struct {
	port   string
	router *mux.Router
	db	 *sql.Db
}

func NewHttpServer(ctx context.Context) (*HttpServer, error) {

	// Get a database handle.
	db, err := database.GetDbSqlConnection()
	if err != nil {
		return nil, err
	}

	// Set port server if not set
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", ":3000")
	}

	// Get port server
	port := os.Getenv("PORT")

	// Create new server
	return &HttpServer{
		port:   port,
		router: mux.NewRouter(),
	}, nil
}

func (server *HttpServer) Start(router func(s *HttpServer, r *mux.Router)) {

	// Add routes to server router
	router(server, server.router)

	// Add CORS middleware to server router
	handlerRoutes := cors.Default().Handler(server.router)

	fmt.Println("Server listening on port", server.port)

	// Start server on port anf handler routes
	err := http.ListenAndServe(server.port, handlerRoutes)
	if err != nil {
		panic(err.Error())
	}
}
