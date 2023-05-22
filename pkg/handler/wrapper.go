package handler

import (
	"github.com/gin-gonic/gin"
)

// CanServe is the base handler that will be wrapped and served.
// This object can be arbitrarily complex and instantiated with specific parameters.
type CanServe interface {
	ServeHTTP(ctx *gin.Context)
}

// Options that can be specified to the Wrapper.
type Options struct {
	// Allow Cross-Origin Resource Sharing (CORS).
	// This will probably only be true for testing.
	Cors bool
}

// Wrapper composed of a CanServe instance and Options.
// Can provide a gin.HandlerFunc for use with one or more links.
type Wrapper struct {
	wrapped CanServe
	options Options
}

// NewWrapped returns a new wrapped handler.
func NewWrapped(wrapped CanServe, options Options) *Wrapper {
	return &Wrapper{
		wrapped: wrapped,
		options: options,
	}
}

// HandlerFunc returns a gin.HandlerFunc for use with one or more links.
func (h *Wrapper) HandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if h.options.Cors {
			// CORS is necessary for testing locally but should not be there in production.
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		h.wrapped.ServeHTTP(ctx)
	}
}
