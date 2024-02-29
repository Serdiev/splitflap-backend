package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"
	"splitflap-backend/pkg/fluent"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

func (a *Application) GetCurrentlyPlaying(c *gin.Context) {
	playing, err := a.Spotify.GetCurrentlyPlaying()

	fmt.Println(playing)
	fmt.Println(err)

	c.JSON(http.StatusOK, playing)
}

type SpotifyClient interface {
	IsLoggedIn() bool
	GetCurrentlyPlaying() (*models.SpotifyIsPlaying, error)
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
	auth := GetInitialAccessToken(code)
	if auth == nil {
		c.Redirect(307, "/")
		return
	}

	auth.Expiry = time.Now().Add(time.Second * 3600)
	tokenSource := CreateSpotifyTokenSource(*auth)

	tokenSrc := oauth2.ReuseTokenSourceWithExpiry(auth, &tokenSource, time.Minute*10)
	client := oauth2.NewClient(c, tokenSrc)
	a.Spotify = spotify.NewSpotifyClient(client)

	c.Redirect(307, "/?message=logged_in")
}

func GetInitialAccessToken(code string) *oauth2.Token {
	formData := url.Values{}
	formData.Set("code", code)
	formData.Set("redirect_uri", cfg.Spotify.RedirectUrl)
	formData.Set("grant_type", "authorization_code")

	token := oauth2.Token{}

	err := fluent.Post(cfg.Spotify.TokenUrl, []byte(formData.Encode())).
		WithClientCredentials(cfg.Spotify.ClientId, cfg.Spotify.ClientSecret).
		WithContentType("application/x-www-form-urlencoded").
		OnSuccess(func(bytes []byte) error {
			innerErr := json.Unmarshal(bytes, &token)

			if innerErr != nil {
				return innerErr
			}

			return nil
		}).
		Execute()

	if err != nil {
		return nil
	}

	fmt.Println("okay", token)
	return &token
}

func CreateSpotifyTokenSource(token oauth2.Token) SpotifyTokenSource {
	return SpotifyTokenSource{
		mutex: sync.Mutex{},
		token: &token,
	}
}

type SpotifyTokenSource struct {
	mutex sync.Mutex
	token *oauth2.Token
}

func (sts *SpotifyTokenSource) Token() (*oauth2.Token, error) {
	sts.mutex.Lock()
	defer sts.mutex.Unlock()

	if sts.token.Valid() {
		return sts.token, nil
	}

	newToken := getNewToken(sts.token.RefreshToken)
	if newToken == nil {
		return nil, errors.New("failed to get token")
	}

	sts.token.AccessToken = newToken.AccessToken
	return newToken, nil
}

func getNewToken(refreshToken string) *oauth2.Token {
	formData := url.Values{}
	formData.Set("client_id", cfg.Spotify.ClientId)
	formData.Set("refresh_token", refreshToken)
	formData.Set("grant_type", "refresh_token")

	var token oauth2.Token
	err := fluent.Post(cfg.Spotify.TokenUrl, []byte(formData.Encode())).
		WithContentType("application/x-www-form-urlencoded").
		WithClientCredentials(cfg.Spotify.ClientId, cfg.Spotify.ClientSecret).
		OnSuccess(func(bytes []byte) error {
			innerToken, innerErr := fluent.BytesToStruct[oauth2.Token](bytes)
			token = innerToken
			return innerErr
		}).
		Execute()

	if err != nil {
		log.Error().Err(err).Msg("failed ")
		return nil
	}

	// manually set sincy spotify uses expires_in 3600 instead of expiry in ouath2.token.
	token.Expiry = time.Now().Add(time.Second * 3600)
	return &token
}
