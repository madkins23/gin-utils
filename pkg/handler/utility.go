package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/madkins23/go-utils/server"
)

// Exit returns a handler function that will interrupt the current process via SIGINT.
// A JSON message "exiting" is sent back as the response.
//
// Since the signal goes to the current process there is no need to pass in an http.Server.
//
// Note that executing this function is likely suicide for the parent process.
func Exit(c *gin.Context) {
	msg := gin.H{"message": "exiting"}
	if err := server.Interrupt(); err != nil {
		log.Logger.Error().Err(err).Msg("Unable to interrupt server")
		msg["error"] = err.Error()
	}
	c.JSON(http.StatusOK, msg)
}

// Link returns an HTML page with a list of links to useful server URLs.
//
// These include:
//   - /ping
//   - /exit
//
// The server must be configured with these routes or the links won't work.
func Link(c *gin.Context) {
	if _, err := c.Writer.WriteString(htmlLinks); err != nil {
		c.Status(http.StatusInternalServerError)
		log.Logger.Error().Err(err).Msg("Writing Links page")
		_, _ = c.Writer.WriteString("Error: " + err.Error())
	} else {
		c.Status(http.StatusOK)
	}

}

const htmlLinks = `
<head><title>Links</title></head>
<body><ul>
  <li><a href="/ping">Ping</a> server existence</li>
  <li><a href="/exit">Exit</a> graceful shut down</li>
</ul></body>
`

// Ping returns a simple JSON object containing the message "pong".
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
