package shutdown

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"gin-utils/pkg/ginzero"
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
	g.Interrupt()
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
	require.NoError(t, waitForServer())

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

const (
	serverURL  = "http://localhost:8080/ping"
	serverWait = 3 * time.Second
)

// waitForServer pings server until it wakes up.
// TODO: should this be an exported method for Graceful?
func waitForServer() error {
	tooLate := time.Now().Add(serverWait)
	for time.Now().Before(tooLate) {
		time.Sleep(100 * time.Millisecond)
		if resp, err := http.Get(serverURL); err != nil {
			// No response yet.
			continue
		} else if err = resp.Body.Close(); err != nil {
			// Shouldn't ever happen.
			return fmt.Errorf("close response body: %w", err)
		} else if resp.StatusCode == http.StatusOK {
			// Server is up and running.
			return nil
		}
	}

	return errors.New("server not ready")
}
