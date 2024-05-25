package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/przb/rcv-server/src/api"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/poll", api.PollCreate)
	r.GET("/poll", api.PollGetResults)
	r.POST("/vote", api.VoteSubmit)
	r.GET("/vote", api.VoteGetOptions)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
