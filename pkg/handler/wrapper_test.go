package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	hdrCORS = "Access-Control-Allow-Origin"
	valCORS = "*"
)

type wrapped struct {
	title, text string
}

func (w *wrapped) ServeHTTP(ctx *gin.Context) {
	writeCenteredText(ctx.Writer, w.title, w.text)
}

func TestWrapped(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	ctx, eng := gin.CreateTestContext(rec)
	require.NotNil(t, ctx)
	require.NotNil(t, eng)
	canServe := &wrapped{title: "Greetings", text: "Hello, world!"}
	require.NotNil(t, canServe)
	wrapper := NewWrapped(canServe, Options{})
	require.NotNil(t, wrapper)
	hdlrFunc := wrapper.HandlerFunc()
	require.NotNil(t, hdlrFunc)
	hdlrFunc(ctx)
	assert.Equal(t, http.StatusOK, rec.Code)
	result := rec.Result()
	require.NotNil(t, result)
	header := result.Header
	require.NotNil(t, header)
	corsHdr := header[hdrCORS]
	assert.Nil(t, corsHdr)
	body := rec.Body.String()
	assert.Contains(t, body, "<title>"+canServe.title+"</title>")
	assert.Contains(t, body, canServe.text)
}

func TestWrapped_Cors(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	ctx, eng := gin.CreateTestContext(rec)
	require.NotNil(t, ctx)
	require.NotNil(t, eng)
	canServe := &wrapped{title: "Greetings", text: "Hello, world!"}
	require.NotNil(t, canServe)
	wrapper := NewWrapped(canServe, Options{Cors: true})
	require.NotNil(t, wrapper)
	hdlrFunc := wrapper.HandlerFunc()
	require.NotNil(t, hdlrFunc)
	hdlrFunc(ctx)
	assert.Equal(t, http.StatusOK, rec.Code)
	result := rec.Result()
	require.NotNil(t, result)
	header := result.Header
	require.NotNil(t, header)
	corsHdr := header[hdrCORS]
	assert.NotNil(t, corsHdr)
	assert.Len(t, corsHdr, 1)
	assert.Equal(t, valCORS, corsHdr[0])
}
