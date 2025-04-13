package handlers

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain/user"
	"github.com/emanuel3k/playlist-transfer/internal/dtos"
	"github.com/emanuel3k/playlist-transfer/pkg/web/request"
	"github.com/emanuel3k/playlist-transfer/pkg/web/response"
	"net/http"
)

type UserHandler struct {
	userService user.ServiceInterface
}

func NewUserHandler(userServiceInterface user.ServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userServiceInterface,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body dtos.CreateUserDTO
	if err := request.Decode(r, &body); err != nil {
		response.Send(w, err.Code, err)
		return
	}

	if err := request.Validate(body); err != nil {
		response.Send(w, err.Code, err)
		return
	}

	res, err := h.userService.Create(body)
	if err != nil {
		response.Send(w, err.Code, err.Message)
		return
	}

	response.Send(w, http.StatusCreated, res)
}
