package cell

import (
	"github.com/go-chi/chi"
)

// NewRouter initializes a new router for cell-related routes.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// CRUD operations for "cell"
	r.Post("/", CreateCellHandler)           // Create
	r.Get("/{cellID}", GetCellHandler)       // Read
	r.Put("/{cellID}", UpdateCellHandler)    // Update
	r.Delete("/{cellID}", DeleteCellHandler) // Delete

	return r
}
