package shutdown

import (
	"context"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/madkins23/go-utils/log"
	"github.com/rs/zerolog"
)

type Graceful struct {
	ctxt   context.Context
	stop   context.CancelFunc
	logger zerolog.Logger
	server *http.Server
	closed bool
}

// Initialize configures the Graceful object.
func (g *Graceful) Initialize() {
	// Create context that listens for the interrupt signal from the OS.
	// NOTE: this code assumes we're running on Linux, it won't work for Apple or Windows.
	g.ctxt, g.stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	g.logger = log.Logger().With().Str("sys", "graceful").Logger()
}

// Serve executes the gin service as defined.
// Service is done in a separate goroutine but this method waits until service is done.
func (g *Graceful) Serve(router *gin.Engine, port uint) error {
	// Build http.Server object manually, don't use gin.Run().
	g.server = &http.Server{
		Addr:    ":" + strconv.Itoa(int(port)),
		Handler: router,
	}

	// Start server in goroutine so shutdown code can run.
	go func() {
		if err := g.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			g.logger.Fatal().Err(err).Msg("Running gin server")
		}
	}()

	// Listen for the interrupt signal.
	<-g.ctxt.Done()

	return nil
}

// Close the Graceful object, stopping signal capture.
func (g *Graceful) Close() {
	if !g.closed {
		// Return signal behavior to initial state.
		g.stop()
		g.logger.Info().Msg("Shutting down gracefully, press Ctrl+C again to force exit.")

		if g.server == nil {
			g.logger.Warn().Msg("No server during Graceful.Close()")
		} else {
			// This context is used to inform the server it has 5 seconds to finish
			// the request it is currently handling
			ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := g.server.Shutdown(ctxt); err != nil {
				g.logger.Error().Err(err).Msg("Server forced to shutdown")
			}

		}

		g.closed = true
	}
}
