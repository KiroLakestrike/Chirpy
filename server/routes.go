package server

import (
	"net/http"
)

// SetupRoutes creates all application routes with the given configuration
func SetupRoutes(cfg RouteConfig) []Route {
	// Static file server with middleware
	fileServer := http.StripPrefix("/app", http.FileServer(http.Dir(cfg.FileRoot)))
	wrappedFileServer := cfg.APIConfig.MiddlewareMetricsInc(fileServer)

	routes := []Route{
		// Static files
		{Pattern: "/app/", Handler: wrappedFileServer},

		// Simple handlers
		{Pattern: "/", Handler: http.HandlerFunc(cfg.HelloHandler)},
		{Pattern: "/about", Handler: http.HandlerFunc(cfg.AboutHandler)},

		// Admin routes
		{Pattern: "GET /admin/metrics", Handler: http.HandlerFunc(cfg.APIConfig.MetricsHandler)},
		{Pattern: "POST /admin/reset", Handler: http.HandlerFunc(cfg.APIConfig.ResetHandler)},

		// API routes
		{Pattern: "GET /api/healthz", Handler: http.HandlerFunc(cfg.HealthzHandler)},
		{Pattern: "POST /api/validate_chirp", Handler: http.HandlerFunc(cfg.ValidateChirpHandler)},
		{Pattern: "POST /api/users", Handler: http.HandlerFunc(cfg.APIConfig.NewUser)},
	}

	return routes
}

// RouteConfig holds all the dependencies needed for route setup
type RouteConfig struct {
	APIConfig interface {
		MiddlewareMetricsInc(http.Handler) http.Handler
		MetricsHandler(w http.ResponseWriter, r *http.Request)
		ResetHandler(w http.ResponseWriter, r *http.Request)
		NewUser(w http.ResponseWriter, r *http.Request)
	}
	FileRoot             string
	HelloHandler         http.HandlerFunc
	AboutHandler         http.HandlerFunc
	HealthzHandler       http.HandlerFunc
	ValidateChirpHandler http.HandlerFunc
}
