package main

import (
	"github.com/evergreenies/go-api-tdd/pkg/domain"
	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
	store  domain.Store
}

func newServer(store domain.Store) *server {
	return &server{
		store: store,
	}
}

func (s *server) routes() *gin.Engine {
	if s.router == nil {
		s.setupRoutes()
	}

	return s.router
}

func (s *server) run(address string) error {
	return s.router.Run(address)
}
