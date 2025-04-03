package services

import "github.com/emanuel3k/playlist-transfer/internal/domain"

type UserService struct {
	domain.UserRepositoryInterface
}

func NewUserService(userRepositoryInterface domain.UserRepositoryInterface) *UserService {
	return &UserService{
		UserRepositoryInterface: userRepositoryInterface,
	}
}
