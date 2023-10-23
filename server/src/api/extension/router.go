package extension

import (
	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", ReadExtensionHandler)

	return r
}
