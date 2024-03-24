package main

import (
	"github.com/gin-gonic/gin"
)

func (s *server) setupRoutes() {
	mux := gin.Default()

	v1 := mux.Group("/api/v1")
	v1.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, "OK.")
	})
	v1.POST("/users/create", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})

	s.router = mux
}
