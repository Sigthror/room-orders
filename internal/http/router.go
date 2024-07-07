package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type orderHandlers interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewRouter(
	oh orderHandlers,
	hFunc http.HandlerFunc,
) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", hFunc)
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Route("/v1", func(r chi.Router) {
			r.Route("/order", func(r chi.Router) {
				r.Get("/", oh.Get)
				r.Post("/", oh.Create)
				r.Put("/", oh.Update)
				r.Delete("/", oh.Delete)
			})
		})
	})

	return r
}
