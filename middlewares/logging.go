// middlewares/logging.go

package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Proceed to the next handler (or the actual route handler)
		next.ServeHTTP(w, r)

		// Calculate the request duration
		duration := time.Since(startTime)

		// Log details about the request
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, duration)
	})
}
