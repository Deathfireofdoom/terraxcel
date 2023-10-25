package sheet

import (
	"github.com/Deathfireofdoom/terraxcel/server/src/api/cell"
	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// CRUD operations for "sheet"
	r.Post("/", CreateSheetHandler)            // Create
	r.Get("/{sheetID}", GetSheetHandler)       // Read
	r.Put("/{sheetID}", UpdateSheetHandler)    // Update
	r.Delete("/{sheetID}", DeleteSheetHandler) // Delete

	r.Route("/{sheetID}/cell", func(r chi.Router) {
		r.Mount("/", cell.NewRouter())
	})

	return r
}
