// Package health contains health-related functionalities.
package health

import (
	"github.com/go-chi/chi"
)

// NewRouter initializes a new routing instance and registers the health check route.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Registers a GET endpoint for health checks.
	r.Get("/", HealthCheckHandler)

	return r
}
