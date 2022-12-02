package handler

import (
	"net/http"
	"strings"

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
	if err := server.Interrupt(); err != nil {
		log.Logger.Error().Err(err).Msg("Unable to interrupt server")
	}
	writeCenteredText(c.Writer, "Exit", "Server shutting down via SIGINT")
}

// Link returns an HTML page with a short list of links to useful server URLs.
//
// These include:
//   - /ping
//   - /exit
//
// The server must be configured with these routes or the links won't work.
func Link(c *gin.Context) {
	writePage(c.Writer, "link", htmlLinks)
}

const htmlLinks = `
<head><title>Links</title></head>
<body><ul>
  <li><a href="/ping">Ping</a> server existence</li>
  <li><a href="/adapted">AdaptedHandler</a> demo http.Handler adapter</li>
  <li><a href="/adaptFn">AdaptedFunc</a> demo http.HandlerFunc adapter</li>
  <li><a href="/exit">Exit</a> graceful shut down</li>
</ul></body>
`

// Ping returns a simple JSON object containing the message "pong".
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//////////////////////////////////////////////////////////////////////////
// Pages that test adapting http.Handler and http.HandlerFunc to gin.HandlerFunc.

// Make certain that AdaptedHandler implements the http.Handler interface.
var _ = http.Handler(&AdaptedHandler{})

// AdaptedHandler is an http.Handler that returns a page demonstrating the use of handler.Adapter.
type AdaptedHandler struct{}

// ServeHTTP method defined so that AdaptedHandler implements the http.Handler interface.
func (hh *AdaptedHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writeCenteredText(writer, "Adapted Handler", "Adapt http.Handler to gin.HandlerFunc")
}

// AdaptedFunc returns a page demonstrating the use of handler.AdapterFunc.
func AdaptedFunc(writer http.ResponseWriter, _ *http.Request) {
	writeCenteredText(writer, "Adapted HandlerFunc", "Adapt http.HandlerFunc to gin.HandlerFunc")
}

//////////////////////////////////////////////////////////////////////////

// centeredText returns an HTML page with the specified title and
// the specified text centered on the page.
func centeredText(title string, text string) string {
	return strings.Replace(strings.Replace(htmlCenteredText,
		"[TITLE]", title, 1),
		"[TEXT]", text, 1)
}

const htmlCenteredText = `
<head><title>[TITLE]</title></head>
<body><center>[TEXT]</center></body>
`

// writePage writes an HTML page.
func writePage(writer http.ResponseWriter, pageName string, html string) {
	if _, err := writer.Write([]byte(html)); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Logger.Error().Err(err).Str("page", pageName).Msg("Writing page")
		_, _ = writer.Write([]byte("Error: " + err.Error()))
	}
}

// writeCenteredText writes an HTML page with centered text.
func writeCenteredText(writer http.ResponseWriter, title, text string) {
	writePage(writer, title, centeredText(title, text))
}
