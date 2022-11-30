package shutdown

import (
	"context"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/madkins23/go-utils/server"

	"gin-utils/pkg/ginzero"
	"gin-utils/pkg/handler"
)

const (
	port    = 8080
	timeout = 100 * time.Millisecond
	url     = "http://localhost:8080/ping"
)

func TestGraceful_Initialize(t *testing.T) {
	g := initialized(t)
	proc, err := os.FindProcess(os.Getpid())
	require.NotNil(t, proc)
	require.NoError(t, err)
	require.NoError(t, proc.Signal(syscall.SIGINT))
	done := false
	select {
	case <-g.ctxt.Done():
		done = true
	case <-time.After(1 * time.Second):
		require.Fail(t, "timeout waiting for SIGINT")
	}
	require.True(t, done)
}

func TestGraceful_Interrupt(t *testing.T) {
	g := initialized(t)
	require.NoError(t, server.Interrupt())
	done := false
	select {
	case <-g.ctxt.Done():
		done = true
	case <-time.After(1 * time.Second):
		require.Fail(t, "timeout waiting for SIGINT")
	}
	require.True(t, done)
}

func TestGraceful_Serve(t *testing.T) {
	g := initialized(t)
	require.NotNil(t, g)

	// Simple router with /ping
	router := gin.New() // not gin.Default()
	router.Use(ginzero.Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	go func() { require.NoError(t, g.Serve(router, 8080)) }()

	// Wait for router to respond properly.
	require.NoError(t, server.WaitFor(url, timeout))

	ctxt, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := g.server.Shutdown(ctxt); err != nil {
		g.logger.Error().Err(err).Msg("Server forced to shutdown")
	}

	require.True(t, true)
}

func TestGraceful_Close(t *testing.T) {
	g := initialized(t)
	require.False(t, g.closed)
	g.Close()
	require.True(t, g.closed)
}

func initialized(t *testing.T) *Graceful {
	g := &Graceful{}
	g.Initialize()
	require.NotNil(t, g.ctxt)
	require.NotNil(t, g.stop)
	return g
}

//////////////////////////////////////////////////////////////////////////

func ExampleGraceful() {
	graceful := &Graceful{}
	graceful.Initialize()
	defer graceful.Close()

	router := gin.New() // not gin.Default()
	router.Use(ginzero.Logger())
	router.GET("/ping", handler.Ping)

	go func() {
		if err := server.WaitFor(url, timeout); err != nil {
			log.Logger.Error().Err(err).Msg("Waiting for server")
		}
		time.Sleep(25 * time.Millisecond)
		if err := server.Interrupt(); err != nil {
			log.Logger.Error().Err(err).Msg("Unable to interrupt server")
		}
	}()

	if err := graceful.Serve(router, port); err != nil {
		log.Logger.Fatal().Err(err).Msg("Running gin server example")
	}

	// Output:
}
