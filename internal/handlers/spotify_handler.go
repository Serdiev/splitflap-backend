package handlers

import (
	"net/http"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (a *Application) GetCurrentlyPlaying(c *gin.Context) {
	playing, err := a.Spotify.GetCurrentlyPlaying()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	logger.Info().Interface("playing", playing).Msg("playing")
	c.JSON(http.StatusOK, playing)
}

type SpotifyClient interface {
	IsLoggedIn() bool
	GetCurrentlyPlaying() (*models.SpotifyIsPlaying, error)
	ShouldUpdateSplitFlap() bool
	ToggleShouldUpdateSplitFlap() bool
}

func (a *Application) SpotifyLogin(c *gin.Context) {
	responseType := c.DefaultQuery("response_type", "code")
	clientID := c.DefaultQuery("client_id", cfg.Spotify.ClientId)
	scope := c.DefaultQuery("scope", "user-read-currently-playing")
	redirectURI := c.DefaultQuery("redirect_uri", cfg.Spotify.RedirectUrl)

	// Construct the redirect URL
	redirectURL := "https://accounts.spotify.com/authorize?" +
		"response_type=" + responseType +
		"&client_id=" + clientID +
		"&scope=" + scope +
		"&redirect_uri=" + redirectURI

	c.Redirect(307, redirectURL)
}

func (a *Application) SpotifyLoginCallback(c *gin.Context) {
	code := c.Query("code")
	auth := spotify.GetInitialAccessToken(code)
	if auth == nil {
		c.Redirect(307, "/")
		return
	}

	auth.Expiry = time.Now().Add(time.Second * 3600)
	tokenSource := spotify.CreateSpotifyTokenSource(*auth)

	tokenSrc := oauth2.ReuseTokenSourceWithExpiry(auth, &tokenSource, time.Minute*10)
	client := oauth2.NewClient(c, tokenSrc)
	a.Spotify = spotify.NewSpotifyClient(client)

	c.Redirect(307, "/?message=logged_in")
}

func (a *Application) ToggleSpotify(c *gin.Context) {
	c.JSON(200, gin.H{"isActive": a.Spotify.ToggleShouldUpdateSplitFlap()})
}
