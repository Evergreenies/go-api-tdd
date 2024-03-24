package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	godotenv.Load("../../../../.env")
	dbUri := os.Getenv("TEST_DB_URI")
	if dbUri == "" {
		log.Fatal("Database connection string does not exist in enviroment variables.")
	}

	dbDriver := os.Getenv("TEST_DB_DRIVER")
	if dbDriver == "" {
		log.Fatal("Cannot connect as database driver does not exists in enviroment.")
	}

	dbConn, err := sql.Open(dbDriver, dbUri)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	testDB = dbConn
	code := m.Run()
	err = testDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
