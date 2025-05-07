package handlers

import (
	"github.com/emanuel3k/playlist-transfer/internal/domain/spotify"
	"github.com/emanuel3k/playlist-transfer/pkg/web/response"
	"net/http"
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

	redirectURI := h.spotifyService.GetRedirectURI(userId)

	if apiErr := h.spotifyService.SetState(userId, userId); apiErr != nil {
		response.Send(w, apiErr.Code, apiErr.Message)
		return
	}

	http.Redirect(w, r, redirectURI, http.StatusFound)
}

func (h *SpotifyHandler) Callback(w http.ResponseWriter, r *http.Request) {
	spotifyErr := r.URL.Query().Get("error")
	if spotifyErr != "" {
		response.Send(w, http.StatusBadRequest, "error: "+spotifyErr)
		return
	}

	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")
	if code == "" || state == "" {
		response.Send(w, http.StatusBadRequest, "missing code or state")
		return
	}

	accessToken, err := h.spotifyService.ExchangeCodeForToken(code, state)
	if err != nil {
		response.Send(w, err.Code, err.Message)
		return
	}

	w.Header().Set("Spotify-Access-Token", accessToken)
	response.Send(w, http.StatusOK, nil)
}
