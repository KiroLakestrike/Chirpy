package handler

import (
	"net/http"
)

// Middleware increments the hits count and calls the next handler
func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileserverHits.Add(1) // increment counter safely
		next.ServeHTTP(w, r)      // serve the file or next handler
	})
}
