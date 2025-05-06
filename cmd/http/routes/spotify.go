package routes

import (
	"github.com/emanuel3k/playlist-transfer/cmd/http/handlers"
	"github.com/emanuel3k/playlist-transfer/cmd/http/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MapSpotifyRoutes() http.Handler {
	r := chi.NewRouter()

	spotifyHandler := handlers.NewSpotifyHandler()

	r.With(middleware.Authorization).Get("/auth", spotifyHandler.Auth)

	return r
}
