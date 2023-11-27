package router

import (
	"github.com/gin-gonic/gin"
	"test-chat/internal/client"
)

var r *gin.Engine

func InitRouter(clientHandler *client.Handler) {
	r = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/ws/connect", clientHandler.Connect)
}

func Start(addr string) error {
	return r.Run(addr)
}
