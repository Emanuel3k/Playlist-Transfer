package services

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain/spotify"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
)

type SpotifyService struct {
	spotifyRepository spotify.RepositoryInterface
}

func NewSpotifyService(spotifyRepository spotify.RepositoryInterface) *SpotifyService {
	return &SpotifyService{
		spotifyRepository: spotifyRepository,
	}
}

func (s *SpotifyService) SetState(userID string, scope string) *web.AppError {
	err := s.spotifyRepository.SetState(userID, scope)
	if err != nil {
		return err
	}

	return nil
}
