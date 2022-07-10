package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingModule(router *gin.RouterGroup) {
	router.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
