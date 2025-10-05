package server

import (
	"net/http"
)

// SetupRoutes creates all application routes with the given configuration
func SetupRoutes(cfg RouteConfig) []Route {
	// Static file server with middleware
	fileServer := http.StripPrefix("/app", http.FileServer(http.Dir(cfg.FileRoot)))
	wrappedFileServer := cfg.APIConfig.MiddlewareMetricsInc(fileServer)

	// ... existing code ...

	routes := []Route{
		// Static files
		{Pattern: "/app/", Handler: wrappedFileServer},

		// Simple handlers
		{Pattern: "/", Handler: cfg.HelloHandler},
		{Pattern: "/about", Handler: cfg.AboutHandler},

		// Admin routes
		{Pattern: "GET /admin/metrics", Handler: http.HandlerFunc(cfg.APIConfig.MetricsHandler)},
		{Pattern: "POST /admin/reset", Handler: http.HandlerFunc(cfg.APIConfig.ResetHandler)},

		// API routes
		{Pattern: "GET /api/healthz", Handler: cfg.HealthzHandler},
		{Pattern: "POST /api/users", Handler: http.HandlerFunc(cfg.APIConfig.NewUser)},
		{Pattern: "POST /api/chirps", Handler: http.HandlerFunc(cfg.APIConfig.Chirps)},
		{Pattern: "GET /api/chirps", Handler: http.HandlerFunc(cfg.APIConfig.GetAllChirps)},
		{Pattern: "GET /api/chirps/{id}", Handler: http.HandlerFunc(cfg.APIConfig.GetChirp)},
		{Pattern: "POST /api/login", Handler: http.HandlerFunc(cfg.APIConfig.Login)},
	}

	// ... existing code ...

	return routes
}

// RouteConfig holds all the dependencies needed for route setup
type RouteConfig struct {
	APIConfig interface {
		MiddlewareMetricsInc(http.Handler) http.Handler
		MetricsHandler(w http.ResponseWriter, r *http.Request)
		ResetHandler(w http.ResponseWriter, r *http.Request)
		NewUser(w http.ResponseWriter, r *http.Request)
		Chirps(w http.ResponseWriter, r *http.Request)
		GetAllChirps(w http.ResponseWriter, r *http.Request)
		GetChirp(w http.ResponseWriter, r *http.Request)
		Login(w http.ResponseWriter, r *http.Request)
	}
	FileRoot       string
	HelloHandler   http.HandlerFunc
	AboutHandler   http.HandlerFunc
	HealthzHandler http.HandlerFunc
	ChirpsHandler  http.HandlerFunc
}
