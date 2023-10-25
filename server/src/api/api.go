// api/api.go
package api

import (
	"github.com/Deathfireofdoom/terraxcel/server/src/api/extension"
	"github.com/Deathfireofdoom/terraxcel/server/src/api/health"
	"github.com/Deathfireofdoom/terraxcel/server/src/api/workbook"
	"github.com/Deathfireofdoom/terraxcel/server/src/middleware"
	"github.com/go-chi/chi"
)

// NewRouter initializes all the api routes and returns the router.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/health", health.NewRouter())
	r.Mount("/extension", extension.NewRouter())

	r.Group(
		func(r chi.Router) {
			r.Use(middleware.Authentication)
			r.Mount("/workbook", workbook.NewRouter())
		})
	return r
}
