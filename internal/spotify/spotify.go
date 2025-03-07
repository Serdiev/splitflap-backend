package spotify

import (
	"errors"
	"fmt"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
	"splitflap-backend/pkg/fluent"
	"strings"
	"time"
)

var cfg = config.New()

type SpotifyClient struct {
	client   *http.Client
	loggedIn bool
	isActive bool
}

func NewNoopSpotifyClient() *SpotifyClient {
	return &SpotifyClient{
		loggedIn: false,
		isActive: false,
	}
}

func NewSpotifyClient(client *http.Client) *SpotifyClient {
	return &SpotifyClient{
		client:   client,
		loggedIn: true,
		isActive: true,
	}
}

func (sc *SpotifyClient) ToggleShouldUpdateSplitFlap() bool {
	sc.isActive = !sc.isActive
	return sc.isActive
}

func (sc *SpotifyClient) ShouldUpdateSplitFlap() bool {
	return sc.isActive
}

func (sc *SpotifyClient) IsLoggedIn() bool {
	return sc.loggedIn
}

func (sc *SpotifyClient) GetCurrentlyPlaying() (*models.SpotifyIsPlaying, error) {
	if !sc.IsLoggedIn() {
		return nil, nil
	}

	var spotifyResp *SpotifyResponse

	err := fluent.Get(cfg.Spotify.BaseUrl+"/me/player/currently-playing?additional_types=track,episode").
		WithClient(sc.client).
		OnStatusCode(http.StatusNoContent, func(_ []byte) error {
			return nil
		}).
		OnSuccess(func(payload []byte) error {
			res, err := fluent.Deserialize[SpotifyResponse](payload)
			if err != nil {
				return fmt.Errorf("failed to parse API response: %w", err)
			}
			spotifyResp = res

			return nil
		}).
		OnError(func(payload []byte) error {
			return fmt.Errorf("failed to get the current song. err:%s", string(payload))
		}).
		Execute()

	if err != nil || spotifyResp == nil {
		logger.Error().Err(err).Msg("failed to get the current song")
		return nil, errors.New("failed to get the current song")
	}

	if spotifyResp.IsPlaying {
		return mapToDto(spotifyResp), nil
	}

	return nil, nil
}

func mapToDto(resp *SpotifyResponse) *models.SpotifyIsPlaying {
	secondsLeft := asSeconds(resp.Item.DurationMS - resp.ProgressMS)
	imageUrl := ""
	for _, image := range resp.Item.Album.Images {
		if image.Height == 64 && image.Width == 64 {
			imageUrl = image.URL
		}
	}
	if imageUrl == "" {
		for _, image := range resp.Item.Images {
			if image.Height == 64 && image.Width == 64 {
				imageUrl = image.URL
			}
		}
	}

	if resp.CurrentlyPlayingType == "episode" {
		return &models.SpotifyIsPlaying{
			Song:            utils.ReplaceDisallowedLetters(strings.Replace(resp.Item.Name, "#", "", 1)),
			Artist:          utils.ReplaceDisallowedLetters(resp.Item.Show.Name),
			ProgressMs:      asSeconds(resp.ProgressMS),
			DurationMs:      asSeconds(resp.Item.DurationMS),
			SecondsLeft:     secondsLeft,
			TimeLeft:        formatSecondsToMMSS(secondsLeft),
			Image64PixelUrl: imageUrl,
		}
	}

	return &models.SpotifyIsPlaying{
		Song:            utils.ReplaceDisallowedLetters(resp.Item.Name),
		Artist:          utils.ReplaceDisallowedLetters(resp.Item.Artists[0].Name),
		ProgressMs:      asSeconds(resp.ProgressMS),
		DurationMs:      asSeconds(resp.Item.DurationMS),
		SecondsLeft:     secondsLeft,
		TimeLeft:        formatSecondsToMMSS(secondsLeft),
		Image64PixelUrl: imageUrl,
	}
}

func asSeconds(ms int64) int {
	return int(ms / 1000)
}

func formatSecondsToMMSS(seconds int) string {
	duration := time.Second * time.Duration(seconds)
	minutes := int(duration.Minutes())
	remainingSeconds := seconds - (minutes * 60)
	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}
