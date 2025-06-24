package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (router *Router) Routes() chi.Router {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/url", router.PostCreateURL)

	return mux
}
