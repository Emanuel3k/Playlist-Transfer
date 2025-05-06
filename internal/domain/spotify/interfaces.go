package spotify

import "github.com/emanuel3k/playlist-transfer/pkg/web"

type ServiceInterface interface {
	SetState(userID string, scope string) *web.AppError
}

type RepositoryInterface interface {
	SetState(userID string, scope string) *web.AppError
}
