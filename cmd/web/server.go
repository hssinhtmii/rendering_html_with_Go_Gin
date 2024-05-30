package main

import (
	"net/http"
	hndlr "ws/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handler *hndlr.Handler
	router *gin.Engine
}

func (server *Server)server() http.Handler{
	router := gin.Default()

	router.GET("/", server.handler.Home)

	server.router = router
	
	router.Run(":8080")
	return router
}
