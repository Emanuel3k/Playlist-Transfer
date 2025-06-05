package services

import (
	"errors"
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/internal/domain/user/mock"
	"github.com/emanuel3k/playlist-transfer/internal/dtos"
	"github.com/emanuel3k/playlist-transfer/pkg/security"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type createResponse struct {
	err *web.AppError
}
type getByEmailResponse struct {
	res *domain.User
	err *web.AppError
}
type hashedPasswordResponse struct {
	res string
	err *web.AppError
}
type expectedResponse struct {
	res *dtos.UserResponseDTO
	err *web.AppError
}
type testCase struct {
	name                   string
	hashedPasswordResponse hashedPasswordResponse
	expectedResponse       expectedResponse
	getByEmailResponse     getByEmailResponse
	createResponse         createResponse
}

func TestUserService_Create(t *testing.T) {
	newUser := dtos.CreateUserDTO{
		FirstName: "Jhon",
		LastName:  "Doe",
		Email:     "test@email.com",
		Password:  "12345678",
	}

	userDomain := newUser.ToDomain()

	userId := "0"
	userDomain.ID = &userId

	hashedPassword := "res"
	userDomain.Password = hashedPassword

	expectedUser := dtos.UserToResponse(*userDomain)

	testCases := []testCase{
		{
			name: "Create user sucessfully",
			hashedPasswordResponse: hashedPasswordResponse{
				res: hashedPassword,
				err: nil,
			},
			expectedResponse: expectedResponse{
				res: expectedUser,
				err: nil,
			},
			getByEmailResponse: getByEmailResponse{
				res: nil,
				err: nil,
			},
			createResponse: createResponse{
				err: nil,
			},
		},
		{
			name: "User already exists",
			hashedPasswordResponse: hashedPasswordResponse{
				res: hashedPassword,
				err: nil,
			},
			expectedResponse: expectedResponse{
				res: nil,
				err: errUserEmailAlredyExists,
			},
			getByEmailResponse: getByEmailResponse{
				res: userDomain,
				err: nil,
			},
			createResponse: createResponse{
				err: nil,
			},
		},
		{
			name: "Error getting user by email",
			hashedPasswordResponse: hashedPasswordResponse{
				res: hashedPassword,
				err: nil,
			},
			expectedResponse: expectedResponse{
				res: nil,
				err: errUserEmailAlredyExists,
			},
			getByEmailResponse: getByEmailResponse{
				res: nil,
				err: errUserEmailAlredyExists,
			},
			createResponse: createResponse{
				err: nil,
			},
		},
		{
			name: "Error creating user",
			hashedPasswordResponse: hashedPasswordResponse{
				res: hashedPassword,
				err: nil,
			},
			expectedResponse: expectedResponse{
				res: nil,
				err: web.InternalServerError(errors.New("")),
			},
			getByEmailResponse: getByEmailResponse{
				res: nil,
				err: nil,
			},
			createResponse: createResponse{
				err: web.InternalServerError(errors.New("")),
			},
		},
		{
			name: "Error hashing password",
			hashedPasswordResponse: hashedPasswordResponse{
				res: "",
				err: web.InternalServerError(errors.New("")),
			},
			expectedResponse: expectedResponse{
				res: nil,
				err: web.InternalServerError(errors.New("")),
			},
			getByEmailResponse: getByEmailResponse{
				res: nil,
				err: nil,
			},
			createResponse: createResponse{
				err: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRespository := user_mock.NewRepositoryInterfaceMock(t)
			userService := NewUserService(userRespository)

			userRespository.EXPECT().GetByEmail(newUser.Email).Return(tc.getByEmailResponse.res, tc.getByEmailResponse.err)

			originalHashPassword := security.HashPassword
			defer func() { security.HashPassword = originalHashPassword }()

			security.HashPassword = func(password string) (string, *web.AppError) {
				return tc.hashedPasswordResponse.res, tc.hashedPasswordResponse.err
			}

			if tc.getByEmailResponse.res == nil && tc.getByEmailResponse.err == nil && tc.hashedPasswordResponse.err == nil {
				userRespository.EXPECT().
					Create(mock.MatchedBy(func(u *domain.User) bool {
						return u.FirstName == userDomain.FirstName &&
							u.LastName == userDomain.LastName &&
							u.Email == userDomain.Email &&
							u.Password == userDomain.Password
					})).
					Run(func(u *domain.User) {
						u.ID = &userId
					}).
					Return(tc.createResponse.err)
			}

			result, err := userService.Create(newUser)

			assert.Equal(t, result, tc.expectedResponse.res)
			assert.Equal(t, err, tc.expectedResponse.err)
		})
	}
}
