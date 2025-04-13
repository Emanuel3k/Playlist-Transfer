package user

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/internal/dtos"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
)

type ServiceInterface interface {
	Create(body dtos.CreateUserDTO) (*dtos.UserResponseDTO, *web.AppError)
	Login(body dtos.LoginDTO) (string, *web.AppError)
}

type RepositoryInterface interface {
	GetByEmail(email string) (*domain.User, *web.AppError)
	Create(user *domain.User) *web.AppError
}
