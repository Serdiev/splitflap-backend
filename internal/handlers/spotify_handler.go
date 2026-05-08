package handlers

import (
	"net/http"
	"time"

	config "splitflap-backend/configs"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (a *Application) GetCurrentlyPlaying(c *gin.Context) {
	client, exists := a.SpotifyClients[MainSpotifyAccountId]
	if !exists || !client.IsLoggedIn() {
		c.JSON(http.StatusOK, nil)
		return
	}

	playing, err := client.GetCurrentlyPlaying()
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
	RegisterHandler(name string, handler func(playing *models.SpotifyIsPlaying))
	DeleteHandler(name string)
}

type LoginRequest struct {
	DeviceId string `form:"device_id"`
}

func (a *Application) SpotifyLogin(c *gin.Context, request LoginRequest) {
	config := cfg.Spotify.SpotifyConfigurations[request.DeviceId]

	logger.Info().Str("device_id", request.DeviceId).Msg("Spotify: login request")

	redirectURL := "https://accounts.spotify.com/authorize?" +
		"response_type=code" +
		"&scope=user-read-currently-playing" +
		"&client_id=" + config.ClientId +
		"&redirect_uri=" + cfg.Spotify.RedirectUrl + "/" + request.DeviceId

	logger.Info().Str("redirect_url", redirectURL).Msg("Spotify: redirect URL")

	c.Redirect(307, redirectURL)
}

func (a *Application) IsLoggedIn(c *gin.Context) {
	client, exists := a.SpotifyClients[MainSpotifyAccountId]
	if !exists || !client.IsLoggedIn() {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusNoContent)
}

type CallbackRequest struct {
	Id   string // parses first param here
	Code string `form:"code"`
}

func (a *Application) SpotifyLoginCallback(c *gin.Context, request CallbackRequest) {
	authToken := spotify.GetInitialAccessToken(request.Id, request.Code)
	if authToken == nil {
		c.Redirect(307, "/")
		return
	}

	spotifyConfig := cfg.Spotify.SpotifyConfigurations[request.Id]
	newSpotifyClient := createSpotifyClient(c, authToken, spotifyConfig)

	a.handleSpotifyClient(SpotifyAccountId(request.Id), newSpotifyClient)

	c.Redirect(307, "/?message=logged_in&device_id="+request.Id)
}

func createSpotifyClient(c *gin.Context, auth *oauth2.Token, config config.SpotifyAccountConfig) *spotify.SpotifyClient {
	auth.Expiry = time.Now().Add(time.Second * 3600)
	tokenSource := spotify.CreateSpotifyTokenSource(*auth, config)

	tokenSrc := oauth2.ReuseTokenSourceWithExpiry(auth, &tokenSource, time.Minute*10)
	client := oauth2.NewClient(c, tokenSrc)
	newClient := spotify.NewSpotifyClient(client)
	return newClient
}

func (a *Application) ToggleSpotify(c *gin.Context) {
	a.SpotifyShouldUpdateSplitFlap = !a.SpotifyShouldUpdateSplitFlap
	if !a.SpotifyShouldUpdateSplitFlap {
		a.SetToIdleState("spotify toggle")
	}
	c.JSON(200, gin.H{"isActive": a.SpotifyShouldUpdateSplitFlap})
}

func (a *Application) handleSpotifyClient(spotifyAccountId SpotifyAccountId, client *spotify.SpotifyClient) {
	logger.Info().Str("id", string(spotifyAccountId)).Msg("Spotify: handling new client")
	spotifyClient, exists := a.SpotifyClients[spotifyAccountId]
	if exists {
		logger.Info().Str("id", string(spotifyAccountId)).Msg("Spotify: disposing old client")
		spotifyClient.Dispose()
		delete(a.SpotifyClients, spotifyAccountId)
	}

	a.SpotifyClients[spotifyAccountId] = client

	if spotifyAccountId == MainSpotifyAccountId {
		logger.Info().Str("id", string(spotifyAccountId)).Msg("Spotify: adding splitflap handler")
		client.RegisterHandler("update-splitflap", a.SendIsPlayingTextToSplitflap)
	}

	client.StartLoop()
}

func (a *Application) BroadcastSpotifyImage(playing *models.SpotifyIsPlaying) {
}
