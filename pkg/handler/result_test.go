package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorResult(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	ErrorResult(rec, http.StatusInternalServerError, "Bad", "doggy!")
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	result := rec.Result()
	require.NotNil(t, result)
	header := result.Header
	require.NotNil(t, header)
	hdrType := header[hdrContentType]
	assert.NotNil(t, hdrType)
	assert.Len(t, hdrType, 1)
	assert.Equal(t, hdrContentTypeValue, hdrType[0])
	hdrOpts := header[hdrContentTypeOptions]
	assert.NotNil(t, hdrOpts)
	assert.Len(t, hdrOpts, 1)
	assert.Equal(t, hdrContentTypeOptionsValue, hdrOpts[0])
	body := rec.Body.String()
	assert.Contains(t, body, "Bad")
	assert.Contains(t, body, "doggy")
}
