package drivers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() *sql.DB {
	// Capture connection properties
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST_POSTGRES"),
		os.Getenv("DB_PORT_POSTGRES"),
		os.Getenv("DB_USERNAME_POSTGRES"),
		os.Getenv("DB_PASSWORD_POSTGRES"),
		os.Getenv("DB_NAME_POSTGRES"),
	)

	// Get a database handle.
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	// Ping to check the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
