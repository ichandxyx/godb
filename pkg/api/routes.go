package api

import (
	"github.com/go-chi/chi/v5"
)

func (a *API) registerRoutes(r chi.Router) {
	r.Get("/user", a.handleGetusers)
	r.Post("/user",a.handleAddUsers)
}
