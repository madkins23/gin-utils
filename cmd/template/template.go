package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	logUtils "github.com/madkins23/go-utils/log"

	"github.com/madkins23/gin-utils/pkg/ginzero"
	"github.com/madkins23/gin-utils/pkg/handler"
	"github.com/madkins23/gin-utils/pkg/shutdown"
	"github.com/madkins23/gin-utils/pkg/system"
)

// appName is the name of this application.
const appName = "template"

// Config collects all configuration information.
// Use json and yaml struct attributes as needed in any embedded structs.
type Config struct {
	// Configuration object for gin.
	Gin system.Config `json:"gin" yaml:"gin"`

	// Configuration object for zerolog.
	Log logUtils.ConsoleOrFile `json:"log" yaml:"log"`

	// Other configuration items may be added as required.
}

func main() {
	flags := flag.NewFlagSet(appName, flag.ContinueOnError)

	var config Config
	config.Gin.AddFlagsToSet(flags)
	config.Log.AddFlagsToSet(flags, "/tmp/console-or-file.log")
	if err := flags.Parse(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Printf("Error parsing command line flags: %s", err)
		}
		return
	}
	if err := config.Log.Setup(); err != nil {
		fmt.Printf("Log file creation error: %s", err)
		return
	}
	defer config.Log.CloseForDefer()

	// Initialize for graceful shutdown.
	graceful := &shutdown.Graceful{Config: config.Gin}
	graceful.Initialize()
	defer graceful.Close()

	gin.DefaultWriter = ginzero.NewWriter(zerolog.InfoLevel)
	gin.DefaultErrorWriter = ginzero.NewWriter(zerolog.ErrorLevel)
	router := gin.New() // not gin.Default()
	router.Use(ginzero.Logger())

	router.GET("/link", handler.Link)
	router.GET("/ping", handler.Ping)
	router.GET("/adaptFn", handler.AdaptFunc(handler.AdaptedFunc))
	router.GET("/adapted", handler.Adapt(&handler.AdaptedHandler{}))
	router.GET("/exit", handler.Exit)

	log.Logger.Info().Msgf("Application %s starting", appName)
	log.Logger.Info().Msgf("> http://localhost:%d/link", config.Gin.Port)
	log.Logger.Info().Msgf("> http://localhost:%d/exit", config.Gin.Port)
	defer log.Logger.Info().Msgf("Application %s finished", appName)

	if err := graceful.Serve(router, 0); err != nil {
		log.Fatal().Err(err).Msg("Running gin server")
	}
}
