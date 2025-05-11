package provider

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

type SpotifyProvider struct {
}

func NewSpotifyProvider() *SpotifyProvider {
	return &SpotifyProvider{}
}

var (
	spotifyRedirectURI  = "SPOTIFY_REDIRECT_URI"
	spotifyClientID     = "SPOTIFY_CLIENT_ID"
	scope               = "user-read-private user-read-email"
	spotifyBaseURL      = "https://accounts.spotify.com"
	spotifyClientSecret = "SPOTIFY_CLIENT_SECRET"
)

func (s *SpotifyProvider) GetAuthURI(userId string) string {
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

func (s *SpotifyProvider) GetAccessToken(code string) (*spotify.GetAccessTokenSpotifyAPIResponse, *web.AppError) {
	redirectUri := os.Getenv(spotifyRedirectURI)

	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	req, reqErr := http.NewRequest(http.MethodPost, spotifyBaseURL+"/api/token", strings.NewReader(data.Encode()))
	if reqErr != nil {
		return nil, web.InternalServerError(reqErr)
	}

	clientId := os.Getenv(spotifyClientID)
	clientsecret := os.Getenv(spotifyClientSecret)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientsecret))
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		return nil, web.InternalServerError(resErr)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, web.InternalServerError(fmt.Errorf("failed to exchange code for token: %d", res.StatusCode))
	}

	var tokenResponse spotify.GetAccessTokenSpotifyAPIResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenResponse); err != nil {
		return nil, web.InternalServerError(err)
	}

	return &tokenResponse, nil
}
