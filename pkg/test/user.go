package test

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain"
	"github.com/emanuel3k/playlist-transfer/internal/dtos"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
)

type CreateResponse struct {
	Err *web.AppError
}
type GetByEmailResponse struct {
	Res *domain.User
	Err *web.AppError
}
type HashedPasswordResponse struct {
	Res string
	Err *web.AppError
}
type ExpectedResponse struct {
	Res *dtos.UserResponseDTO
	Err *web.AppError
}
type UserTestCase struct {
	Name                   string
	HashedPasswordResponse HashedPasswordResponse
	ExpectedResponse       ExpectedResponse
	GetByEmailResponse     GetByEmailResponse
	CreateResponse         CreateResponse
}
