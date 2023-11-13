package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	server := setupRoute()

	req, _ := http.NewRequest("GET", "/hello", nil)
	req.Header.Set("Authorization", "Bearer testToken")
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, w.Code)

	t.Logf("Body: %s", w.Body.String())
}
