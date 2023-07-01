package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gt2rz/micro-auth/internal/constants"
	"github.com/gt2rz/micro-auth/internal/database/drivers"
)

func GetDbSqlConnection() (*sql.DB, error) {
	var db *sql.DB

	switch os.Getenv("DB_TYPE") {
	case "postgres":
		db = drivers.NewPostgresConnection()

	default:
		fmt.Println(constants.ErrNotDatabaseTypeSpecified.Error())
		return nil, constants.ErrNotDatabaseTypeSpecified
	}

	fmt.Println(constants.ConnectedToDatabaseType + ": " + os.Getenv("DB_TYPE"))
	return db, nil
}
