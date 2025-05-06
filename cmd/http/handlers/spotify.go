package handlers

import (
	"net/http"
	"net/url"
	"os"
)

var (
	spotifyRedirectURI = "SPOTIFY_REDIRECT_URI"
	spotifyClientID    = "SPOTIFY_CLIENT_ID"
)

type SpotifyHandler struct {
}

func NewSpotifyHandler() *SpotifyHandler {
	return &SpotifyHandler{}
}

func (h *SpotifyHandler) Auth(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("user_id")

	spotifyRedirectURL := os.Getenv(spotifyRedirectURI)
	spotifyClientID := os.Getenv(spotifyClientID)
	scope := "user-read-private user-read-email"

	spotifyURL := "https://accounts.spotify.com/authorize"
	params := url.Values{}
	params.Add("client_id", spotifyClientID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", spotifyRedirectURL)
	params.Add("scope", scope)
	params.Add("state", userId)

	// todo: add redis to store the state

	spotifyAuthURL := spotifyURL + "?" + params.Encode()

	http.Redirect(w, r, spotifyAuthURL, http.StatusFound)
}
