package spotify

import (
	"encoding/json"
	"errors"
	"net/url"
	"splitflap-backend/internal/logger"
	"splitflap-backend/pkg/fluent"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

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

	newToken := GetRefreshedToken(sts.token.RefreshToken)
	if newToken == nil {
		return nil, errors.New("failed to get token")
	}

	sts.token.AccessToken = newToken.AccessToken
	return newToken, nil
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

	return &token
}

func GetRefreshedToken(refreshToken string) *oauth2.Token {
	formData := url.Values{}
	formData.Set("client_id", cfg.Spotify.ClientId)
	formData.Set("refresh_token", refreshToken)
	formData.Set("grant_type", "refresh_token")

	var token *oauth2.Token
	err := fluent.Post(cfg.Spotify.TokenUrl, []byte(formData.Encode())).
		WithContentType("application/x-www-form-urlencoded").
		WithClientCredentials(cfg.Spotify.ClientId, cfg.Spotify.ClientSecret).
		OnSuccess(func(bytes []byte) error {
			innerToken, innerErr := fluent.Deserialize[oauth2.Token](bytes)
			token = innerToken
			return innerErr
		}).
		Execute()

	if err != nil {
		logger.Error().Err(err).Msg("failed to refresh token")
		return nil
	}

	// manually set sincy spotify uses expires_in 3600 instead of expiry in ouath2.token.
	token.Expiry = time.Now().Add(time.Second * 3600)
	return token
}
