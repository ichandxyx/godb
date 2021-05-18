package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/ichandxyx/godb/pkg/store"
)

type API struct {
	store  *store.Store
	logger log.Logger
}

func New(s *store.Store, logger log.Logger) *API {
	//  var a API
	//  a.store=s
	//  a.logger=logger
	//  return &a
	return &API{
		store:  s,
		logger: logger,
	}
}
func (a *API) Register(r chi.Router) {
	a.registerRoutes(r)
}
