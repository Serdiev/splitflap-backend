package spotify_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/spotify"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func Test_MapToDto_Track(t *testing.T) {
	resp := spotify.SpotifyResponse{
		IsPlaying: true,
		ProgressMS: 30000,
		CurrentlyPlayingType: "track",
		Item: spotify.Item{
			Name: "Test Song",
			DurationMS: 180000,
			Artists: []spotify.Artist{{Name: "Test Artist"}},
			Album: spotify.Album{
				Images: []spotify.Image{
					{URL: "https://example.com/64.jpg", Height: 64, Width: 64},
					{URL: "https://example.com/300.jpg", Height: 300, Width: 300},
				},
			},
			Images: []spotify.Image{
				{URL: "https://example.com/alt64.jpg", Height: 64, Width: 64},
			},
		},
	}

	dto := spotify.MapToDto(&resp)

	assert.Equal(t, "Test S0ng", dto.Song)
	assert.Equal(t, "Test Artist", dto.Artist)
	assert.Equal(t, 30, dto.ProgressMs)
	assert.Equal(t, 180, dto.DurationMs)
	assert.Equal(t, 150, dto.SecondsLeft)
	assert.Equal(t, "02:30", dto.TimeLeft)
	assert.Equal(t, "https://example.com/64.jpg", dto.Image64PixelUrl)
}

func Test_MapToDto_PodcastEpisode(t *testing.T) {
	resp := spotify.SpotifyResponse{
		IsPlaying: true,
		ProgressMS: 60000,
		CurrentlyPlayingType: "episode",
		Item: spotify.Item{
			Name: "Test Episode #42",
			DurationMS: 3600000,
			Images: []spotify.Image{
				{URL: "https://example.com/episode64.jpg", Height: 64, Width: 64},
			},
			Show: spotify.Show{Name: "Test Show"},
		},
	}

	dto := spotify.MapToDto(&resp)

	assert.Equal(t, "Test Epis0de 42", dto.Song)
	assert.Equal(t, "Test Sh0w", dto.Artist)
	assert.Equal(t, 60, dto.ProgressMs)
	assert.Equal(t, 3600, dto.DurationMs)
	assert.Equal(t, 3540, dto.SecondsLeft)
	assert.Equal(t, "59:00", dto.TimeLeft)
	assert.Equal(t, "https://example.com/episode64.jpg", dto.Image64PixelUrl)
}

func Test_MapToDto_NoMatchingImage(t *testing.T) {
	resp := spotify.SpotifyResponse{
		IsPlaying: true,
		ProgressMS: 0,
		CurrentlyPlayingType: "track",
		Item: spotify.Item{
			Name: "No Image Song",
			DurationMS: 1000,
			Artists: []spotify.Artist{{Name: "Unknown"}},
			Album: spotify.Album{
				Images: []spotify.Image{
					{URL: "https://example.com/300.jpg", Height: 300, Width: 300},
				},
			},
		},
	}

	dto := spotify.MapToDto(&resp)

	assert.Equal(t, "", dto.Image64PixelUrl)
}

func Test_FormatSecondsToMMSS(t *testing.T) {
	tests := []struct {
		seconds  int
		expected string
	}{
		{0, "00:00"},
		{30, "00:30"},
		{60, "01:00"},
		{90, "01:30"},
		{125, "02:05"},
		{600, "10:00"},
		{3599, "59:59"},
		{3600, "60:00"},
		{3661, "61:01"},
	}

	for _, tt := range tests {
		result := spotify.FormatSecondsToMMSS(tt.seconds)
		assert.Equal(t, tt.expected, result)
	}
}

func Test_AsSeconds(t *testing.T) {
	tests := []struct {
		ms       int64
		expected int
	}{
		{0, 0},
		{500, 0},
		{999, 0},
		{1000, 1},
		{1500, 1},
		{60000, 60},
		{180000, 180},
		{3600000, 3600},
	}

	for _, tt := range tests {
		result := spotify.AsSeconds(tt.ms)
		assert.Equal(t, tt.expected, result)
	}
}

func Test_NoopSpotifyClient_IsLoggedIn(t *testing.T) {
	client := spotify.NewNoopSpotifyClient()
	assert.Equal(t, false, client.IsLoggedIn())
}

func Test_SpotifyClient_RegisterHandler(t *testing.T) {
	client := spotify.NewNoopSpotifyClient()

	handler := func(playing *models.SpotifyIsPlaying) {}

	client.RegisterHandler("test-handler", handler)

	assert.Equal(t, 1, len(client.Handlers()))
}

func Test_SpotifyClient_DeleteHandler(t *testing.T) {
	client := spotify.NewNoopSpotifyClient()

	handler := func(playing *models.SpotifyIsPlaying) {}
	client.RegisterHandler("test-handler", handler)

	client.DeleteHandler("test-handler")

	assert.Equal(t, 0, len(client.Handlers()))
}

func Test_SpotifyClient_Handlers(t *testing.T) {
	client := spotify.NewNoopSpotifyClient()

	handler1 := func(playing *models.SpotifyIsPlaying) {}
	handler2 := func(playing *models.SpotifyIsPlaying) {}
	handler3 := func(playing *models.SpotifyIsPlaying) {}

	client.RegisterHandler("handler-1", handler1)
	client.RegisterHandler("handler-2", handler2)
	client.RegisterHandler("handler-3", handler3)

	handlers := client.Handlers()
	assert.Equal(t, 3, len(handlers))
}

func Test_SpotifyClient_GetCurrentlyPlaying_NotLoggedIn(t *testing.T) {
	client := spotify.NewNoopSpotifyClient()

	result, err := client.GetCurrentlyPlaying()

	assert.Equal(t, nil, result)
	assert.Equal(t, nil, err)
}

func Test_SpotifyClient_GetCurrentlyPlaying_204NoContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := spotify.NewSpotifyClientWithBaseUrl(server.URL, &http.Client{})
	result, err := client.GetCurrentlyPlaying()

	assert.Equal(t, nil, result)
	assert.Equal(t, nil, err)
}

func Test_SpotifyClient_GetCurrentlyPlaying_PlayingTrack(t *testing.T) {
	spotifyResp := spotify.SpotifyResponse{
		IsPlaying:            true,
		ProgressMS:           30000,
		CurrentlyPlayingType: "track",
		Item: spotify.Item{
			Name:       "Test Song",
			DurationMS: 180000,
			Artists:    []spotify.Artist{{Name: "Test Artist"}},
			Album: spotify.Album{
				Images: []spotify.Image{
					{URL: "https://example.com/64.jpg", Height: 64, Width: 64},
				},
			},
		},
	}
	respBytes, _ := json.Marshal(spotifyResp)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	}))
	defer server.Close()

	client := spotify.NewSpotifyClientWithBaseUrl(server.URL, &http.Client{})
	result, err := client.GetCurrentlyPlaying()

	require.NoError(t, err)
	assert.Equal(t, "Test S0ng", result.Song)
	assert.Equal(t, "Test Artist", result.Artist)
}

func Test_SpotifyClient_GetCurrentlyPlaying_Paused(t *testing.T) {
	spotifyResp := spotify.SpotifyResponse{
		IsPlaying:            false,
		ProgressMS:           30000,
		CurrentlyPlayingType: "track",
		Item: spotify.Item{
			Name:       "Paused Song",
			DurationMS: 180000,
			Artists:    []spotify.Artist{{Name: "Test Artist"}},
			Album: spotify.Album{
				Images: []spotify.Image{},
			},
		},
	}
	respBytes, _ := json.Marshal(spotifyResp)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
	}))
	defer server.Close()

	client := spotify.NewSpotifyClientWithBaseUrl(server.URL, &http.Client{})
	result, err := client.GetCurrentlyPlaying()

	require.NoError(t, err)
	assert.Equal(t, nil, result)
}

func Test_SpotifyClient_GetCurrentlyPlaying_HttpError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := spotify.NewSpotifyClientWithBaseUrl(server.URL, &http.Client{})
	result, err := client.GetCurrentlyPlaying()

	assert.Equal(t, nil, result)
	require.NotNil(t, err)
	assert.Equal(t, true, strings.Contains(err.Error(), "failed to get the current song"))
}

func Test_SpotifyIsPlaying_PercentageLeft(t *testing.T) {
	playing := &models.SpotifyIsPlaying{
		ProgressMs: 30,
		DurationMs: 180,
	}

	percentage := playing.PercentageLeft()

	assert.Equal(t, 16, percentage)
}

func Test_SpotifyIsPlaying_PercentageLeft_ZeroDuration(t *testing.T) {
	playing := &models.SpotifyIsPlaying{
		ProgressMs: 0,
		DurationMs: 0,
	}

	percentage := playing.PercentageLeft()

	assert.Equal(t, 0, percentage)
}