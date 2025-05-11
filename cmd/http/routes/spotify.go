package routes

import (
	"github.com/emanuel3k/playlist-transfer/cmd/http/handlers"
	"github.com/emanuel3k/playlist-transfer/cmd/http/middleware"
	"github.com/emanuel3k/playlist-transfer/config/redis"
	"github.com/emanuel3k/playlist-transfer/internal/provider"
	"github.com/emanuel3k/playlist-transfer/internal/repositories"
	"github.com/emanuel3k/playlist-transfer/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MapSpotifyRoutes() http.Handler {
	r := chi.NewRouter()

	rds := redis.Get()
	spotifyRepository := repositories.NewSpotifyRepository(rds)
	spotifyProvider := provider.NewSpotifyProvider()
	spotifyService := services.NewSpotifyService(spotifyRepository, spotifyProvider)
	spotifyHandler := handlers.NewSpotifyHandler(spotifyService)

	r.With(middleware.Authorization).Get("/auth", spotifyHandler.Auth)
	r.Get("/callback", spotifyHandler.Callback)

	return r
}
