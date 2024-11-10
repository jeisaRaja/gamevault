package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Route("/games", func(r chi.Router) {
			r.Post("/", app.createGameHandler)
			r.Get("/{id}", app.showGameHandler)
		})
	})

	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.HandleFunc("/*", app.notFoundResponse)

	return r
}
