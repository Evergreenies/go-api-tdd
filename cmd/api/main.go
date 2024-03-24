package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func databaseConnection() (*sql.DB, error) {
	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		return nil, errors.New("Database connection does not exist in enviroment variables.")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		return nil, errors.New("Cannot connect as database driver does not exists in enviroment.")
	}

	db, err := sql.Open(dbDriver, dbUri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	godotenv.Load(".env")

	_, err := databaseConnection()
	if err != nil {
		log.Fatal(err)
	}
}
