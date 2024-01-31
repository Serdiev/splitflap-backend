package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
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
		fmt.Println("here")
		return nil, nil
	}

	resp, err := sc.client.Get(cfg.Spotify.BaseUrl + "/me/player/currently-playing?additional_types=track,episode")
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

	// fmt.Println(string(bytes))
	if spotifyResp.IsPlaying {
		dto := mapToDto(spotifyResp)
		return dto, nil
	}

	return nil, nil
}

func mapToDto(resp SpotifyResponse) *models.SpotifyIsPlaying {
	secondsLeft := asSeconds(resp.Item.DurationMS - resp.ProgressMS)

	if resp.CurrentlyPlayingType == "episode" {
		return &models.SpotifyIsPlaying{
			Song:        utils.ReplaceDisallowedLetters(resp.Item.Name),
			Artist:      utils.ReplaceDisallowedLetters(resp.Item.Show.Name),
			ProgressMs:  asSeconds(resp.ProgressMS),
			DurationMs:  asSeconds(resp.Item.DurationMS),
			SecondsLeft: secondsLeft,
			TimeLeft:    formatSecondsToMMSS(secondsLeft),
		}
	}

	return &models.SpotifyIsPlaying{
		Song:        utils.ReplaceDisallowedLetters(resp.Item.Name),
		Artist:      utils.ReplaceDisallowedLetters(resp.Item.Artists[0].Name),
		ProgressMs:  asSeconds(resp.ProgressMS),
		DurationMs:  asSeconds(resp.Item.DurationMS),
		SecondsLeft: secondsLeft,
		TimeLeft:    formatSecondsToMMSS(secondsLeft),
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
