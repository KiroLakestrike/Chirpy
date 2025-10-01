package main

import "net/http"

// Middleware increments the hits count and calls the next handler
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1) // increment counter safely
		next.ServeHTTP(w, r)      // serve the file or next handler
	})
}
