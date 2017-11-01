package app

import (
	"log"
	"net/http"
	"time"
)

// Logger func
func Logger(innerLogger http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		start := time.Now()

		innerLogger.ServeHTTP(w, request)

		log.Printf(
			"%s\t%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			name,
			time.Since(start),
		)
	})
}
