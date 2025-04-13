package routes

import (
	"github.com/emanuel3k/playlist-transfer/cmd/http/handlers"
	"github.com/emanuel3k/playlist-transfer/config/postgres"
	"github.com/emanuel3k/playlist-transfer/internal/repositories"
	"github.com/emanuel3k/playlist-transfer/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MapUserRoutes() http.Handler {
	db := postgres.GetDB()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()

	r.Post("/create", userHandler.Create)
	r.Post("/login", userHandler.Login)

	return r
}
