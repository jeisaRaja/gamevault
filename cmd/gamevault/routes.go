package main

import (
	"fmt"

	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (app *application) routes(r *chi.Mux) {

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", app.config.port))))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Route("/games", func(r chi.Router) {
			r.Post("/", app.createGameHandler)
			r.Get("/{id}", app.showGameHandler)
      r.Put("/{id}", app.updateGameHandler)
		})
	})

	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.HandleFunc("/*", app.notFoundResponse)
}
