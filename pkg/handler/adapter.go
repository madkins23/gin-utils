package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Adapt wraps a gin.HandlerFunc around an http.Handler.
// Provides a way to quickly reuse existing http.Handler objects with Gin.
//
// Note: the gin.Context object is not passed into the http.Handler
// so there is no way to access path variables stored on the context.
//
// Deprecated: Use gin.WrapH instead, it does the same thing.
func Adapt(handler http.Handler) gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		handler.ServeHTTP(ctxt.Writer, ctxt.Request)
	}
}

// AdaptFunc wraps a gin.HandlerFunc around an http.HandlerFunc.
// Provides a way to quickly reuse existing http.HandlerFunc functions with Gin.
//
// Note: the gin.Context object is not passed into the http.Handler
// so there is no way to access path variables stored on the context.
//
// Deprecated: Use gin.WrapF instead, it does the same thing.
func AdaptFunc(hfn http.HandlerFunc) gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		hfn(ctxt.Writer, ctxt.Request)
	}
}
