package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	logUtils "github.com/madkins23/go-utils/log"

	"gin-utils/pkg/ginzero"
	"gin-utils/pkg/shutdown"
)

const appName = "template"

var port uint

func main() {
	flags := flag.NewFlagSet(appName, flag.ContinueOnError)
	flags.UintVar(&port, "port", 8080, "specify server port with leading colon")

	cof := logUtils.ConsoleOrFile{}
	cof.AddFlagsToSet(flags, "/tmp/console-or-file.log")
	if err := flags.Parse(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Printf("Error parsing command line flags: %s", err)
		}
		return
	}
	if err := cof.Setup(); err != nil {
		fmt.Printf("Log file creation error: %s", err)
		return
	}
	defer cof.CloseForDefer()

	// TODO: check port number

	// Initialize for graceful shutdown.
	graceful := &shutdown.Graceful{}
	graceful.Initialize()
	defer graceful.Close()

	gin.DefaultWriter = ginzero.NewWriter(zerolog.InfoLevel)
	gin.DefaultErrorWriter = ginzero.NewWriter(zerolog.ErrorLevel)
	router := gin.New() // not gin.Default()
	router.Use(ginzero.Logger())

	router.GET("/exit", func(c *gin.Context) {
		graceful.Interrupt()
		c.JSON(http.StatusOK, gin.H{
			"message": "exiting",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Logger.Info().Msgf("Application %s starting", appName)
	log.Logger.Info().Msgf("> http://localhost:%d/ping", port)
	log.Logger.Info().Msgf("> http://localhost:%d/exit", port)
	defer log.Logger.Info().Msgf("Application %s finished", appName)

	if err := graceful.Serve(router, port); err != nil {
		log.Fatal().Err(err).Msg("Running gin server")
	}

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	graceful.Close()

}
