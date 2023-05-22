package handler

import (
	"fmt"
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

//------------------------------------------------------------------------

// Link returns an HTML page with a short list of links to useful(?) server URLs.
//
// The links are:
//   - /ping
//   - /exit
//
// The server must be configured with these routes or the links won't work.
//
// Deprecated: this is probably not ever going to be used and may be removed
// in a future major release.
func Link(c *gin.Context) {
	Links(c, LinkTableStyleSheet,
		LinkDef{"/ping", "Ping", "server existence"},
		LinkDef{"/exit", "Exit", "graceful shut down"})
}

// LinkTableStyleSheet is the default style sheet for use in Links().
const LinkTableStyleSheet = `
    table.links {
      border-collapse: collapse;
      margin-left: auto;
      margin-right: auto;
    }

    table.links td.link {
        text-align: right;
    }

    table.links td.description {
        font-style: italic;
    }

    table.links td.spacer {
        width: 1em;
    }
`

// LinkDef defines a link for use in the Links() function.
type LinkDef struct {
	Path        string
	Name        string
	Description string
}

// Links returns an HTML page with a short list of links to useful server URLs.
// The links are provided by argument and are displayed in a simple table.
// A stylesheet may be provided (and may be "") or the default LinkTableStyleSheet is used.
//
// This function can't be passed directly as a handler as it has too many arguments.
// Surround it with another function that has a single *gin.Context argument
// and no return values.
//
// The server must be configured with all defined links.
func Links(c *gin.Context, styleSheet string, links ...LinkDef) {
	var page strings.Builder
	page.WriteString("<head>\n  <title>Links</title>")
	if styleSheet == "" {
		styleSheet = LinkTableStyleSheet
	}
	page.WriteString("  <style>\n")
	page.WriteString(styleSheet)
	page.WriteString("  </style>\n")
	page.WriteString("</head>\n<body>\n")
	page.WriteString("  <table class=\"links\">\n")
	for _, link := range links {
		_, _ = fmt.Fprintf(&page,
			"<tr>"+
				"<td class=\"link\"><a href=\"%s\">%s</a></td>"+
				"<td class=\"spacer\"></td>"+
				"<td class=\"descr\">%s</td>"+
				"</tr>\n",
			link.Path, link.Name, link.Description)
	}
	page.WriteString("  </table>\n</body>\n")
	writePage(c.Writer, "links", page.String())
}

//------------------------------------------------------------------------

// Ping returns a simple JSON object containing the message "pong".
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//////////////////////////////////////////////////////////////////////////

const htmlCenteredText = `
<head><title>[TITLE]</title></head>
<body><center>[TEXT]</center></body>
`

// centeredText returns an HTML page with the specified title and
// the specified text centered on the page.
func centeredText(title string, text string) string {
	return strings.Replace(strings.Replace(htmlCenteredText,
		"[TITLE]", title, 1),
		"[TEXT]", text, 1)
}

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
