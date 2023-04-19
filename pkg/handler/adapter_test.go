package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ginHandlerFunc = gin.HandlerFunc(ghf)

func ghf(_ *gin.Context) {}

func TestAdapt(t *testing.T) {
	hdlr := &AdaptedHandler{}
	require.NotNil(t, hdlr)
	adapted := Adapt(hdlr)
	require.NotNil(t, adapted)
	assert.IsType(t, ginHandlerFunc, adapted)
}

func TestAdaptFunc(t *testing.T) {
	adapted := AdaptFunc(AdaptedFunc)
	require.NotNil(t, adapted)
	assert.IsType(t, ginHandlerFunc, adapted)
}
