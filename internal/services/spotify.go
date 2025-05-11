package services

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain/spotify"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
)

type SpotifyService struct {
	spotifyRepository spotify.RepositoryInterface
	spotifyProvider   spotify.ProviderInterface
}

func NewSpotifyService(spotifyRepository spotify.RepositoryInterface, spotifyProvider spotify.ProviderInterface) *SpotifyService {
	return &SpotifyService{
		spotifyRepository: spotifyRepository,
		spotifyProvider:   spotifyProvider,
	}
}

func (s *SpotifyService) GetAuthURI(userId string) string {
	return s.spotifyProvider.GetAuthURI(userId)
}

func (s *SpotifyService) SetState(userID string, scope string) *web.AppError {
	err := s.spotifyRepository.SetState(userID, scope)
	if err != nil {
		return err
	}

	return nil
}

func (s *SpotifyService) ExchangeCodeForToken(code string, state string) (string, *web.AppError) {
	cachedState, appErr := s.spotifyRepository.GetState(state)
	if appErr != nil {
		return "", appErr
	}

	if cachedState != state {
		return "", web.BadRequestErrorWithCauses("Invalid or missing state", nil)
	}

	tokenResponse, apiErr := s.spotifyProvider.GetAccessToken(code)
	if apiErr != nil {

	}

	setStateErr := s.spotifyRepository.SetAccessToken(tokenResponse.AccessToken, tokenResponse.ExpiresIn)
	if setStateErr != nil {
		return "", setStateErr
	}

	return tokenResponse.AccessToken, nil
}
