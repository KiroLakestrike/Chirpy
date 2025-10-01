package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	cfg := &apiConfig{}

	// Wrap the file server handler with middleware to increment hits
	fileServer := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))

	// Register Handlers
	mux.Handle("/app/", cfg.middlewareMetricsInc(fileServer))
	mux.HandleFunc("GET /admin/metrics", cfg.metricsHandler)
	mux.HandleFunc("POST /admin/reset", cfg.resetHandler)
	mux.HandleFunc("GET /api/healthz", healthz)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
