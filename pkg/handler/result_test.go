package handler

import (
	"encoding/json"
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
	assert.Equal(t, hdrContentTypeTextValue, hdrType[0])
	hdrOpts := header[hdrContentTypeOptions]
	assert.NotNil(t, hdrOpts)
	assert.Len(t, hdrOpts, 1)
	assert.Equal(t, hdrContentTypeOptionsValue, hdrOpts[0])
	body := rec.Body.String()
	assert.Contains(t, body, "Bad")
	assert.Contains(t, body, "doggy")
}

// Note: Use float64 instead of int because JSON marshal/unmarshal
// will convert all numbers to float64 and it makes assert.Equal just work.

type obj struct {
	Text string
	OK   bool
}

type thing struct {
	Text   string
	Number float64
	List   []string
	Dict   map[string]any
	Sub    obj
}

func TestJSONResult(t *testing.T) {
	rec := httptest.NewRecorder()
	require.NotNil(t, rec)
	anything := thing{
		Text:   "something",
		Number: 23,
		List: []string{
			"alpha",
			"bravo",
			"charile",
		},
		Dict: map[string]any{
			"first":  float64(13),
			"second": "hand",
			"third":  false,
		},
		Sub: obj{
			Text: "within",
			OK:   true,
		},
	}
	JSONResult(rec, anything)
	assert.Equal(t, http.StatusOK, rec.Code)
	result := rec.Result()
	require.NotNil(t, result)
	header := result.Header
	require.NotNil(t, header)
	hdrType := header[hdrContentType]
	assert.NotNil(t, hdrType)
	assert.Len(t, hdrType, 1)
	assert.Equal(t, hdrContentTypeJSONValue, hdrType[0])
	var received thing
	body := rec.Body
	require.NoError(t, json.Unmarshal(body.Bytes(), &received))
	assert.Equal(t, anything, received)
}
