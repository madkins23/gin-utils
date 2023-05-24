package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	hdrContentType             = "Content-Type"
	hdrContentTypeTextValue    = "text/plain; charset=utf-8"
	hdrContentTypeJSONValue    = "application/json; charset=utf-8"
	hdrContentTypeOptions      = "X-Content-Type-Options"
	hdrContentTypeOptionsValue = "nosniff"
)

// ErrorResult returns a basic error result object as a convenience.
// Substitute something specific to the containing application.
func ErrorResult(writer http.ResponseWriter, code int, text ...string) {
	// Note: copied from http.Error() source.
	writer.Header().Set(hdrContentType, hdrContentTypeTextValue)
	writer.Header().Set(hdrContentTypeOptions, hdrContentTypeOptionsValue)
	writer.WriteHeader(code)
	_, _ = fmt.Fprintln(writer, http.StatusText(code))
	for _, t := range text {
		_, _ = fmt.Fprintln(writer, t)
	}
}

// JSONResult returns a JSON message representing the specified object
func JSONResult(writer http.ResponseWriter, object any) {
	if bytes, err := json.Marshal(object); err != nil {
		ErrorResult(writer, http.StatusInternalServerError, err.Error())
	} else {
		writer.Header().Set(hdrContentType, hdrContentTypeJSONValue)
		writer.WriteHeader(http.StatusOK)
		if _, err = writer.Write(bytes); err != nil {
			ErrorResult(writer, http.StatusInternalServerError, err.Error())
		}
	}
}
