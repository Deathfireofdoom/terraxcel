package workbook

import (
	"github.com/Deathfireofdoom/terraxcel/server/src/api/sheet"
	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// CRUD for workbook
	r.Post("/", CreateWorkbookHandler)
	r.Get("/{id}", GetWorkbookHandler)
	r.Put("/{id}", UpdateWorkbookHandler)
	r.Delete("/{id}", DeleteWorkbookHandler)

	// CRUD for sheet within workbook
	r.Route("/{workbookID}/sheet", func(r chi.Router) {
		r.Mount("/", sheet.NewRouter()) // Mount the sheet router here
	})

	return r
}
