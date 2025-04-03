package handlers

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"net/http"
)

type UserHandler struct {
	userService domain.UserServiceInterface
}

func NewUserHandler(userServiceInterface domain.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userServiceInterface,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// todo: implement user creation logic
}
