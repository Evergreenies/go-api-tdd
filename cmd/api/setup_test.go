package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/evergreenies/go-api-tdd/pkg/domain"
	"github.com/evergreenies/go-api-tdd/pkg/store/sqlstore/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var testStore domain.Store

type testServer struct {
	*httptest.Server
}

func newTestServer(h http.Handler) *testServer {
	return &testServer{Server: httptest.NewServer(h)}
}

func TestMain(m *testing.M) {
	godotenv.Load("../../.env")
	gin.SetMode(gin.TestMode)

	db, err := databaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	testStore = postgres.NewPostgresStore(db)
	code := m.Run()

	// cleanup here
	_ = testStore.DeleteAllUsers()

	os.Exit(code)
}
