package database

import (
	"database/sql"
	"errors"
	"fmt"
	"microtwo/database/drivers"
	"os"
)

var ErrNoDatabaseTypeSpecified = errors.New("no database type specified")

func GetDbSqlConnection() (*sql.DB, error) {
	var db *sql.DB

	switch os.Getenv("DB_TYPE") {
	case "postgres":
		db = drivers.NewPostgresConnection()
		
	default:
		fmt.Println(constants.ErrNoDatabaseTypeSpecified.Error())
		return nil, constants.ErrNoDatabaseTypeSpecified
	}

	fmt.Println(constants.ConnectedToDatabaseType + ": " + os.Getenv("DB_TYPE"))
	return db, nil
}
