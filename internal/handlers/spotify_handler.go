package handlers

import (
	"fmt"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/lcd_display"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"
	"time"

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

	fmt.Println("Using config: ", config)
	fmt.Println("request: ", request)

	// Construct the redirect URL
	redirectURL := "https://accounts.spotify.com/authorize?" +
		"response_type=code" +
		"&scope=user-read-currently-playing" +
		"&client_id=" + config.ClientId +
		"&redirect_uri=" + cfg.Spotify.RedirectUrl + "/" + request.DeviceId

	fmt.Println("Redirect URL: ", redirectURL)

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
	fmt.Println("Handling new spotify client for id", spotifyAccountId)
	// Delete old client
	spotifyClient, exists := a.SpotifyClients[spotifyAccountId]
	if exists {
		fmt.Println("Disposing old client", spotifyAccountId)
		spotifyClient.Dispose()
		delete(a.SpotifyClients, spotifyAccountId)
	}

	// Save new client
	a.SpotifyClients[spotifyAccountId] = client

	// Delete old lcd display
	delete(a.LcdDisplays, spotifyAccountId)

	// Create new lcd display
	newLcd := lcd_display.NewLcdDisplay(string(spotifyAccountId))
	a.LcdDisplays[spotifyAccountId] = newLcd

	client.RegisterHandler("update-lcd-image", newLcd.HandleIsPlaying)
	if spotifyAccountId == MainSpotifyAccountId {
		fmt.Println("Adding splitflap handler", spotifyAccountId)
		client.RegisterHandler("update-splitflap", a.HandleIsPlaying)
	}

	client.StartLoop()
}
