package servers

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gt2rz/micro-auth/internal/database"
	"github.com/gt2rz/micro-auth/internal/repositories"
	"github.com/rs/cors"
)

// HttpServer is the struct for the http server
type HttpServer struct {
	port           string
	router         *mux.Router
	db             *sql.DB
	UserRepository repositories.UserRepository
}

// NewHttpServer creates a new http server
func NewHttpServer(ctx context.Context) (*HttpServer, error) {

	// Get a database handle.
	db, err := database.GetDbSqlConnection()
	if err != nil {
		return nil, err
	}

	// Set port server if not set
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}

	// Get port server
	port := os.Getenv("PORT")

	userRepository, _ := repositories.NewUserRepository(db)

	// Create new server
	return &HttpServer{
		port:           port,
		router:         mux.NewRouter(),
		db:             db,
		UserRepository: userRepository,
	}, nil
}

// Start starts the http server
func (server *HttpServer) Start(router func(s *HttpServer, r *mux.Router)) {

	// Add routes to server router
	router(server, server.router)

	// Add CORS middleware to server router
	handlerRoutes := cors.Default().Handler(server.router)

	// Set port server if not set
	if os.Getenv("APP_ENV") != "production" {
		port, _ := strconv.Atoi(server.port)

		for {
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err == nil {
				listener.Close()
				break
			}
			port++
		}

		server.port = strconv.Itoa(port)
	}

	fmt.Println("Server listening on port", server.port)

	// Start server on port anf handler routes
	err := http.ListenAndServe(fmt.Sprintf(":%s", server.port), handlerRoutes)
	if err != nil {
		panic(err.Error())
	}
}
