package services

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain/user"
	"github.com/emanuel3k/playlist-transfer/internal/dtos"
	"github.com/emanuel3k/playlist-transfer/pkg/security"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
)

var (
	errUserEmailAlredyExists = web.ConflictError("User with this email already exists")
	InvalidEmailOrPassword   = web.UnauthorizedError("Invalid email or password")
)

type UserService struct {
	userRepository user.RepositoryInterface
}

func NewUserService(userRepositoryInterface user.RepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepositoryInterface,
	}
}

func (s *UserService) Create(body dtos.CreateUserDTO) (*dtos.UserResponseDTO, *web.AppError) {
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

	return dtos.UserToResponse(*newUser), nil
}

func (s *UserService) Login(body dtos.LoginDTO) (string, *web.AppError) {
	u, err := s.userRepository.GetByEmail(body.Email)
	if err != nil {
		return "", err
	}

	if u == nil {
		return "", InvalidEmailOrPassword
	}

	if err = security.ComparePassword(u.Password, body.Password); err != nil {
		return "", err
	}

	token, err := u.GenerateToken()
	if err != nil {
		return "", err
	}

	return token, nil
}
