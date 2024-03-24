package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/evergreenies/go-api-tdd/pkg/store/sqlstore/postgres"
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

func setup() (*server, error) {
	db, err := databaseConnection()
	if err != nil {
		return nil, err
	}

	pStore := postgres.NewPostgresStore(db)
	serv := newServer(pStore)
	serv.setupRoutes()

	return serv, nil
}

func main() {
	godotenv.Load(".env")

	serv, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf("%s:%s", host, port)

	if err = serv.run(address); err != nil {
		log.Fatal(err)
	}
}
