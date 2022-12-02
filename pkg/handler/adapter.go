package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Adapt wraps a gin.HandlerFunc around an http.Handler.
// Provides a way to quickly reuse existing http.Handler objects with Gin.
func Adapt(handler http.Handler) gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		handler.ServeHTTP(ctxt.Writer, ctxt.Request)
	}
}

// AdaptFunc wraps a gin.HandlerFunc around an http.HandlerFunc.
// Provides a way to quickly reuse existing http.HandlerFunc functions with Gin.
func AdaptFunc(hfn http.HandlerFunc) gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		hfn(ctxt.Writer, ctxt.Request)
	}
}
