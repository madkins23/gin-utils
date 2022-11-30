package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	ctx, eng := gin.CreateTestContext(rec)
	require.NotNil(t, ctx)
	require.NotNil(t, eng)
	Ping(ctx)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestExit(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	ctx, eng := gin.CreateTestContext(rec)
	require.NotNil(t, ctx)
	require.NotNil(t, eng)
	ctxt, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	require.NotNil(t, ctxt)
	require.NotNil(t, stop)
	defer stop()
	Exit(ctx)
	var done bool
	select {
	case <-ctxt.Done():
		done = true
	case <-time.After(1 * time.Second):
		require.Fail(t, "timeout waiting for SIGINT")
	}
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, done)
}
