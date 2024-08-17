package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	database = "api_db"
)

func ConnectDB() (*sql.DB, error) {
	dbdata := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database,
	)

	db, error := sql.Open("postgres", dbdata)

	// error = db.Ping()
	if error != nil {
		panic(error)
	}

	return db, nil
}
