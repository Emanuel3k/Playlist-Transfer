package handlers

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain/spotify"
	"github.com/emanuel3k/playlist-transfer/pkg/web/response"
	"net/http"
	"net/url"
	"os"
)

var (
	spotifyRedirectURI = "SPOTIFY_REDIRECT_URI"
	spotifyClientID    = "SPOTIFY_CLIENT_ID"
)

type SpotifyHandler struct {
	spotifyService spotify.ServiceInterface
}

func NewSpotifyHandler(spotifyService spotify.ServiceInterface) *SpotifyHandler {
	return &SpotifyHandler{
		spotifyService: spotifyService,
	}
}

func (h *SpotifyHandler) Auth(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("user_id")

	RedirectUri := os.Getenv(spotifyRedirectURI)
	clientId := os.Getenv(spotifyClientID)
	scope := "user-read-private user-read-email"

	spotifyURL := "https://accounts.spotify.com/authorize"
	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("response_type", "code")
	params.Add("redirect_uri", RedirectUri)
	params.Add("scope", scope)
	params.Add("state", userId)

	if apiErr := h.spotifyService.SetState(userId, userId); apiErr != nil {
		response.Send(w, apiErr.Code, apiErr.Message)
		return
	}

	spotifyAuthURL := spotifyURL + "?" + params.Encode()

	http.Redirect(w, r, spotifyAuthURL, http.StatusFound)
}
