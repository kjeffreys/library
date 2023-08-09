package tests_test

import (
	"bytes"
	"library/middlewares"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(log.Writer())
	}()

	// Create a dummy HTTP request and response writer
	req, err := http.NewRequest("GET", "/dummy-path", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.RequestURI = "/dummy-path" // Manually setting the RequestURI
	rr := httptest.NewRecorder()

	// Dummy handler to simulate the next handler in the chain
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test response"))
	})

	handler := middlewares.LoggingMiddleware(nextHandler)
	handler.ServeHTTP(rr, req)

	// Check the captured log output
	logOutput := buf.String()
	if !strings.Contains(logOutput, "GET") || !strings.Contains(logOutput, "/dummy-path") {
		t.Errorf("Expected log to contain method and path, got %s", logOutput)
	}
}
