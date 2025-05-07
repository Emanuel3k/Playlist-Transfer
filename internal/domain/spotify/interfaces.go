package spotify

import "github.com/emanuel3k/playlist-transfer/pkg/web"

type ServiceInterface interface {
	GetRedirectURI(userId string) string
	SetState(userID string, scope string) *web.AppError
	ExchangeCodeForToken(code string, state string) (string, *web.AppError)
}

type RepositoryInterface interface {
	SetState(userID string, scope string) *web.AppError
	GetState(userID string) (string, *web.AppError)
	SetAccessToken(accessToken string, ttl int) *web.AppError
}
