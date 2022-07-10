package main

import (
	"github.com/gin-gonic/gin"
	"golang-realworld-example-app/internal/handler"
)

const BaseURL = "localhost:8080"

func main() {
	server := gin.Default()

	root := server.Group("/api")

	handler.PingModule(root.Group(""))

	server.Run(BaseURL)
}
