package services

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/pkg/security"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
)

var (
	errUserEmailAlredyExists = web.ConflictError("User with this email already exists")
)

type UserService struct {
	userRepository domain.UserRepositoryInterface
}

func NewUserService(userRepositoryInterface domain.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepositoryInterface,
	}
}

func (s *UserService) Create(body domain.CreateUserDTO) (*domain.UserResponseDTO, *web.AppError) {
	existsWithEmail, err := s.userRepository.GetByEmail(body.Email)
	if err != nil {
		return nil, err
	}

	if existsWithEmail != nil {
		return nil, errUserEmailAlredyExists
	}

	newUser := body.ToDomain()

	hashedPassword, err := security.HashPassword(body.Password)
	if err != nil {
		return nil, err
	}

	newUser.Password = hashedPassword

	if err = s.userRepository.Create(newUser); err != nil {
		return nil, err
	}

	return newUser.ToResponse(), nil
}
