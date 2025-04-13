package dtos

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/google/uuid"
)

type CreateUserDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
}

func (u *CreateUserDTO) ToDomain() *domain.User {
	uid := uuid.NewString()
	return &domain.User{
		ID:        &uid,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	}
}

type UserResponseDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func UserToResponse(u domain.User) *UserResponseDTO {
	return &UserResponseDTO{
		ID:        *u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
