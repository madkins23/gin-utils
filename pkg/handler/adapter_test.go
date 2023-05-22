package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ginHandlerFunc = gin.HandlerFunc(ghf)

func ghf(_ *gin.Context) {}

func TestAdapt(t *testing.T) {
	hdlr := &adaptedHandler{}
	require.NotNil(t, hdlr)
	adapted := Adapt(hdlr)
	require.NotNil(t, adapted)
	assert.IsType(t, ginHandlerFunc, adapted)
}

func TestAdaptFunc(t *testing.T) {
	adapted := AdaptFunc(adaptedFunc)
	require.NotNil(t, adapted)
	assert.IsType(t, ginHandlerFunc, adapted)
}

func TestAdapted(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	ctx, eng := gin.CreateTestContext(rec)
	require.NotNil(t, ctx)
	require.NotNil(t, eng)
	Adapt(&adaptedHandler{})(ctx)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestAdaptedFunc(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	ctx, eng := gin.CreateTestContext(rec)
	require.NotNil(t, ctx)
	require.NotNil(t, eng)
	AdaptFunc(adaptedFunc)(ctx)
	assert.Equal(t, http.StatusOK, rec.Code)
}

//////////////////////////////////////////////////////////////////////////
// Pages that test adapting http.Handler and http.HandlerFunc to gin.HandlerFunc.

// Make certain that adaptedHandler implements the http.Handler interface.
var _ = http.Handler(&adaptedHandler{})

// adaptedHandler is an http.Handler that returns a page demonstrating the use of handler.Adapter.
type adaptedHandler struct{}

// ServeHTTP method defined so that adaptedHandler implements the http.Handler interface.
func (hh *adaptedHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writeCenteredText(writer, "Adapted Handler", "Adapt http.Handler to gin.HandlerFunc")
}

// adaptedFunc returns a page demonstrating the use of handler.AdapterFunc.
func adaptedFunc(writer http.ResponseWriter, _ *http.Request) {
	writeCenteredText(writer, "Adapted HandlerFunc", "Adapt http.HandlerFunc to gin.HandlerFunc")
}
