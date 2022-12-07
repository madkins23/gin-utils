/*
demo of ginzero package to support using [zerolog] with [gin].

Once running the application responds to http://:55555/ping with:

	{"message":"pong"}

[gin]: https://gin-gonic.com/docs/
[zerolog]: https://github.com/rs/zerolog
*/
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/madkins23/gin-utils/pkg/ginzero"
	"github.com/madkins23/gin-utils/pkg/handler"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
	gin.DefaultWriter = ginzero.NewWriter(zerolog.InfoLevel)
	gin.DefaultErrorWriter = ginzero.NewWriter(zerolog.ErrorLevel)
	router := gin.New()
	router.Use(ginzero.Logger())
	router.GET("/ping", handler.Ping)
	log.Info().Str("link", "http://localhost:55555/ping").Msg("Ping")
	if err := router.Run(":55555"); err != nil {
		log.Error().Err(err).Msg("Server failure")
	}
}
