package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/emanuel3k/playlist-transfer/internal/domain/spotify"
	"github.com/emanuel3k/playlist-transfer/pkg/web"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	spotifyRedirectURI  = "SPOTIFY_REDIRECT_URI"
	spotifyClientID     = "SPOTIFY_CLIENT_ID"
	scope               = "user-read-private user-read-email"
	spotifyBaseURL      = "https://accounts.spotify.com"
	spotifyClientSecret = "SPOTIFY_CLIENT_SECRET"
)

type SpotifyService struct {
	spotifyRepository spotify.RepositoryInterface
}

func NewSpotifyService(spotifyRepository spotify.RepositoryInterface) *SpotifyService {
	return &SpotifyService{
		spotifyRepository: spotifyRepository,
	}
}

func (s *SpotifyService) GetRedirectURI(userId string) string {
	RedirectUri := os.Getenv(spotifyRedirectURI)
	clientId := os.Getenv(spotifyClientID)

	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("response_type", "code")
	params.Add("redirect_uri", RedirectUri)
	params.Add("scope", scope)
	params.Add("state", userId)

	spotifyAuthURL := spotifyBaseURL + "/authorize?" + params.Encode()

	return spotifyAuthURL
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
		return "", web.BadRequestErrorWithCauses("Invalid state", nil)
	}

	redirectUri := os.Getenv(spotifyRedirectURI)

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	req, reqErr := http.NewRequest(http.MethodPost, spotifyBaseURL+"/api/token", strings.NewReader(data.Encode()))
	if reqErr != nil {
		return "", web.InternalServerError(reqErr)
	}

	clientId := os.Getenv(spotifyClientID)
	clientsecret := os.Getenv(spotifyClientSecret)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientsecret))
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		return "", web.InternalServerError(resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", web.InternalServerError(fmt.Errorf("failed to exchange code for token: %d", res.StatusCode))
	}

	var tokenResponse spotify.GetAccessTokenSpotifyAPIResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
		return "", web.InternalServerError(err)
	}

	setStateErr := s.spotifyRepository.SetAccessToken(tokenResponse.AccessToken, tokenResponse.ExpiresIn)
	if setStateErr != nil {
		return "", setStateErr
	}

	return tokenResponse.AccessToken, nil
}
