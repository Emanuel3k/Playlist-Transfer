package services

import (
	"errors"
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/internal/domain/user/mock"
	"github.com/emanuel3k/playlist-transfer/internal/dtos"
	"github.com/emanuel3k/playlist-transfer/pkg/security"
	"github.com/emanuel3k/playlist-transfer/pkg/test"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

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

	hashedPassword := "Res"
	userDomain.Password = hashedPassword

	expectedUser := dtos.UserToResponse(*userDomain)

	testCases := []test.UserTestCase{
		{
			Name: "Create user successfully",
			HashedPasswordResponse: test.HashedPasswordResponse{
				Res: hashedPassword,
				Err: nil,
			},
			ExpectedResponse: test.ExpectedResponse{
				Res: expectedUser,
				Err: nil,
			},
			GetByEmailResponse: test.GetByEmailResponse{
				Res: nil,
				Err: nil,
			},
			CreateResponse: test.CreateResponse{
				Err: nil,
			},
		},
		{
			Name: "User already exists",
			HashedPasswordResponse: test.HashedPasswordResponse{
				Res: hashedPassword,
				Err: nil,
			},
			ExpectedResponse: test.ExpectedResponse{
				Res: nil,
				Err: errUserEmailAlreadyExists,
			},
			GetByEmailResponse: test.GetByEmailResponse{
				Res: userDomain,
				Err: nil,
			},
			CreateResponse: test.CreateResponse{
				Err: nil,
			},
		},
		{
			Name: "Error getting user by email",
			HashedPasswordResponse: test.HashedPasswordResponse{
				Res: hashedPassword,
				Err: nil,
			},
			ExpectedResponse: test.ExpectedResponse{
				Res: nil,
				Err: errUserEmailAlreadyExists,
			},
			GetByEmailResponse: test.GetByEmailResponse{
				Res: nil,
				Err: errUserEmailAlreadyExists,
			},
			CreateResponse: test.CreateResponse{
				Err: nil,
			},
		},
		{
			Name: "Error creating user",
			HashedPasswordResponse: test.HashedPasswordResponse{
				Res: hashedPassword,
				Err: nil,
			},
			ExpectedResponse: test.ExpectedResponse{
				Res: nil,
				Err: web.InternalServerError(errors.New("")),
			},
			GetByEmailResponse: test.GetByEmailResponse{
				Res: nil,
				Err: nil,
			},
			CreateResponse: test.CreateResponse{
				Err: web.InternalServerError(errors.New("")),
			},
		},
		{
			Name: "Error hashing password",
			HashedPasswordResponse: test.HashedPasswordResponse{
				Res: "",
				Err: web.InternalServerError(errors.New("")),
			},
			ExpectedResponse: test.ExpectedResponse{
				Res: nil,
				Err: web.InternalServerError(errors.New("")),
			},
			GetByEmailResponse: test.GetByEmailResponse{
				Res: nil,
				Err: nil,
			},
			CreateResponse: test.CreateResponse{
				Err: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			userRespository := user_mock.NewRepositoryInterfaceMock(t)
			userService := NewUserService(userRespository)

			userRespository.EXPECT().GetByEmail(newUser.Email).Return(tc.GetByEmailResponse.Res, tc.GetByEmailResponse.Err)

			originalHashPassword := security.HashPassword
			defer func() { security.HashPassword = originalHashPassword }()

			security.HashPassword = func(password string) (string, *web.AppError) {
				return tc.HashedPasswordResponse.Res, tc.HashedPasswordResponse.Err
			}

			if tc.GetByEmailResponse.Res == nil && tc.GetByEmailResponse.Err == nil && tc.HashedPasswordResponse.Err == nil {
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
					Return(tc.CreateResponse.Err)
			}

			result, err := userService.Create(newUser)

			assert.Equal(t, tc.ExpectedResponse.Res, result)
			assert.Equal(t, tc.ExpectedResponse.Err, err)
		})
	}
}
