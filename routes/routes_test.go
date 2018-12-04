package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetVersion)
	handler.ServeHTTP(recorder, req)
	assert.Equal(recorder.Code, http.StatusOK)

	assert.Equal(recorder.Body.String(), "1.0")
}
