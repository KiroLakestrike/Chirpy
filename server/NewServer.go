package server

import (
	"net/http"
	"strconv"
)

type Route struct {
	Pattern string // e.g., "GET /api/healthz" or just "/api/healthz"
	Handler http.Handler
}

func NewServer(port int, routes []Route) *http.Server {
	mux := http.NewServeMux()
	for _, r := range routes {
		mux.Handle(r.Pattern, r.Handler)
	}

	return &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}
}
