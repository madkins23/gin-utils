package handler

import (
	"fmt"
	"net/http"
)

const (
	hdrContentType             = "Content-Type"
	hdrContentTypeValue        = "text/plain; charset=utf-8"
	hdrContentTypeOptions      = "X-Content-Type-Options"
	hdrContentTypeOptionsValue = "nosniff"
)

// ErrorResult returns a basic error result object as a convenience.
// Substitute something specific to the containing application.
func ErrorResult(writer http.ResponseWriter, code int, text ...string) {
	// Note: copied from http.Error() source.
	writer.Header().Set(hdrContentType, hdrContentTypeValue)
	writer.Header().Set(hdrContentTypeOptions, hdrContentTypeOptionsValue)
	writer.WriteHeader(code)
	_, _ = fmt.Fprintln(writer, http.StatusText(code))
	for _, t := range text {
		_, _ = fmt.Fprintln(writer, t)
	}
}
