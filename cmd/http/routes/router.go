package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type router struct {
}

func NewRouter() *router {
	return &router{}
}

func (router *router) InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/user", MapUserRoutes())
	})

	return r
}
