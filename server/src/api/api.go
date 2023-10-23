// api/api.go
package api

import (
	"github.com/Deathfireofdoom/excel-infra-service/api/extension"
	"github.com/Deathfireofdoom/excel-infra-service/api/health"
	"github.com/Deathfireofdoom/excel-infra-service/api/workbook"
	"github.com/Deathfireofdoom/excel-infra-service/middleware"
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
