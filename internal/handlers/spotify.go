package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

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
	if code == "" {
		// c.JSON(http.StatusOK, "could not login, did not receive code")
		c.Redirect(307, "/")
		return
	}

	auth := GetInitialAccessToken(code)
	if auth == nil {
		c.Redirect(307, "/")
		return
	}

	tokenSource := CreateSpotifyTokenSource(*auth)

	tokenSrc := oauth2.ReuseTokenSourceWithExpiry(auth, &tokenSource, time.Second*30)
	client := oauth2.NewClient(c, tokenSrc)
	a.Spotify = spotify.NewSpotifyClient(client)

	aa, b := a.Spotify.GetCurrentlyPlaying()

	fmt.Println(aa, b)
	c.Redirect(307, "/?message=logged_in")
}

func GetInitialAccessToken(code string) *oauth2.Token {
	formData := url.Values{}
	formData.Set("code", code)
	formData.Set("redirect_uri", cfg.Spotify.RedirectUrl)
	formData.Set("grant_type", "authorization_code")

	// Create the HTTP request
	req, err := http.NewRequest("POST", cfg.Spotify.TokenUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// Set the headers
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(cfg.Spotify.ClientId+":"+cfg.Spotify.ClientSecret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authHeader)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return nil
	}

	auth := oauth2.Token{}

	err = json.Unmarshal(respBody, &auth)

	if err != nil {
		return nil
	}

	return &auth
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

	// If the token is still valid, return it
	if sts.token.Valid() {
		return sts.token, nil
	}

	// Otherwise, refresh the token
	formData := url.Values{}
	formData.Set("refresh_token", sts.token.RefreshToken)
	formData.Set("grant_type", "refresh_token")

	// Create the HTTP request
	req, err := http.NewRequest("POST", cfg.Spotify.TokenUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch access token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response status code: %d", resp.StatusCode)
	}

	var token oauth2.Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil, fmt.Errorf("failed to decode access token response: %w", err)
	}

	// Update the token
	sts.token = &token

	return &token, nil
}
