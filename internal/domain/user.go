package domain

import "github.com/emanuel3k/playlist-transfer/pkg/web"

type User struct {
	ID        *string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u *User) ToResponse() *UserResponseDTO {
	return &UserResponseDTO{
		ID:        *u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}

type CreateUserDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
}

func (u *CreateUserDTO) ToDomain() *User {
	return &User{
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

type UserServiceInterface interface {
	Create(body CreateUserDTO) (*UserResponseDTO, *web.AppError)
}

type UserRepositoryInterface interface {
	GetByEmail(email string) (*User, *web.AppError)
	Create(user *User) *web.AppError
}
