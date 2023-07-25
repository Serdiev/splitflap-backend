package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/models"
	"time"
)

var cfg = config.New()
var spotifyAuthorizeCode = ""
var refreshToken = ""

type SpotifyClient struct {
	client   *http.Client
	loggedIn bool
}

func NewNoopSpotifyClient() SpotifyClient {
	return SpotifyClient{
		loggedIn: false,
	}
}

func NewSpotifyClient(client *http.Client) SpotifyClient {
	return SpotifyClient{
		client:   client,
		loggedIn: true,
	}
}

func (sc SpotifyClient) IsLoggedIn() bool {
	return sc.loggedIn
}

func (sc SpotifyClient) GetCurrentlyPlaying() (*models.SpotifyIsPlaying, error) {
	if !sc.IsLoggedIn() {
		return nil, nil
	}

	resp, err := sc.client.Get(cfg.SPOTIFY_URL + "me/player/currently-playing")
	if err != nil {
		return nil, fmt.Errorf("failed to get the current song: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get the current song, status code: %d", resp.StatusCode)
	}

	var spotifyResp SpotifyResponse

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read bytes")
	}

	err = json.Unmarshal(bytes, &spotifyResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	if spotifyResp.IsPlaying {
		return mapToDto(spotifyResp), nil
	}

	return nil, nil
}

func mapToDto(resp SpotifyResponse) *models.SpotifyIsPlaying {
	secondsLeft := int((resp.Item.DurationMS - resp.ProgressMS) / 1000)
	return &models.SpotifyIsPlaying{
		Song:        resp.Item.Name,
		Artist:      resp.Item.Artists[0].Name,
		SecondsLeft: secondsLeft,
		TimeLeft:    formatSecondsToMMSS(secondsLeft),
	}
}

func formatSecondsToMMSS(seconds int) string {
	duration := time.Second * time.Duration(seconds)
	minutes := int(duration.Minutes())
	remainingSeconds := seconds - (minutes * 60)
	return fmt.Sprintf("%02d:%02d", minutes, remainingSeconds)
}
