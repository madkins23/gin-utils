package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/madkins23/go-utils/server"
)

// Ping returns a simple JSON object containing the message "pong".
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Exit returns a handler function that will interrupt the current process via SIGINT.
// Since the signal goes to the current process there is no need to pass in an http.Server.
func Exit(c *gin.Context) {
	msg := gin.H{"message": "exiting"}
	if err := server.Interrupt(); err != nil {
		log.Logger.Error().Err(err).Msg("Unable to interrupt server")
		msg["error"] = err.Error()
	}
	c.JSON(http.StatusOK, msg)
}
