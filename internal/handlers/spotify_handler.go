package handlers

import (
	"fmt"
	"net/http"
	"splitflap-backend/internal/lcd_display"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var SplitflapDeviceId = "splitflap"

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

type LoginRequest struct {
	DeviceId string `form:"device_id"`
}

func (a *Application) SpotifyLogin(c *gin.Context, request LoginRequest) {
	fmt.Println("login req", request)
	clientID := c.DefaultQuery("client_id", cfg.Spotify.ClientId)

	// Construct the redirect URL
	redirectURL := "https://accounts.spotify.com/authorize?" +
		"response_type=code" +
		"&scope=user-read-currently-playing" +
		"&client_id=" + clientID +
		"&redirect_uri=" + cfg.Spotify.RedirectUrl + "/" + request.DeviceId

	fmt.Println("Redirect URL: ", redirectURL)

	c.Redirect(307, redirectURL)
}

type CallbackRequest struct {
	Id   string // parses first param here
	Code string `form:"code"`
}

func (a *Application) SpotifyLoginCallback(c *gin.Context, request CallbackRequest) {

	fmt.Println("login callback req", request)
	auth := spotify.GetInitialAccessToken(request.Id, request.Code)
	if auth == nil {
		c.Redirect(307, "/")
		return
	}

	// Means this is for this splitflap application
	auth.Expiry = time.Now().Add(time.Second * 3600)
	tokenSource := spotify.CreateSpotifyTokenSource(*auth)

	tokenSrc := oauth2.ReuseTokenSourceWithExpiry(auth, &tokenSource, time.Minute*10)
	client := oauth2.NewClient(c, tokenSrc)
	newClient := spotify.NewSpotifyClient(client)

	if request.Id == SplitflapDeviceId {
		a.Spotify = newClient
	} else {
		// Means this is for a external LCD client
		a.LcdDisplays[request.Id] = lcd_display.NewLcdDisplay(request.Id, newClient)

		d := a.LcdDisplays[request.Id]
		go d.StartFetchLoop()
	}

	c.Redirect(307, "/?message=logged_in&device_id="+request.Id)
}

func (a *Application) ToggleSpotify(c *gin.Context) {
	isActive := a.Spotify.ToggleShouldUpdateSplitFlap()
	if !isActive {
		a.SetToIdleState("spotify toggle")
	}
	c.JSON(200, gin.H{"isActive": isActive})
}
