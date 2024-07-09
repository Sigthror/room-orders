package http

import (
	"net/http"
	"top-selection-test/internal/http/middleware"

	"github.com/go-chi/chi/v5"
)

type orderHandlers interface {
	Create(w http.ResponseWriter, r *http.Request) error
}

func NewRouter(
	oh orderHandlers,
) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", healthHandler)
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.UUID, middleware.Logger)
		r.Route("/v1", func(r chi.Router) {
			r.Route("/order", func(r chi.Router) {
				r.Post("/", newEndpoint(oh.Create))
			})
		})
	})

	return r
}
