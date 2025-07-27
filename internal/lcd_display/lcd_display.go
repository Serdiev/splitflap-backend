package lcd_display

import (
	"fmt"
	"math"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/spotify"
	"splitflap-backend/internal/utils"
	"time"
)

type LcdDisplay struct {
	id                    string
	client                *spotify.SpotifyClient
	playing               *LatestPlaying
	isFetchingFromSpotify bool
}

type LatestPlaying struct {
	url string
	img utils.Image
}

func NewLcdDisplay(id string, client *spotify.SpotifyClient) *LcdDisplay {
	return &LcdDisplay{
		id:     id,
		client: client,
	}
}

func (d *LcdDisplay) FetchImage(client *spotify.SpotifyClient) *utils.Image {
	if d.playing == nil {
		return nil
	}

	return &d.playing.img
}

func (d *LcdDisplay) GetImage() *utils.Image {
	if d.playing == nil {
		return nil
	}

	return &d.playing.img
}

func (d *LcdDisplay) StartFetchLoop() {
	fmt.Println("Starting fetch loop for LCD display", d.id)
	if d.isFetchingFromSpotify {
		logger.Info().Msg("Already fetching from Spotify, skipping new fetch loop")
		return
	}

	d.isFetchingFromSpotify = true
	backoffSeconds := 2

	for {
		if !d.isFetchingFromSpotify {
			return
		}

		playing, err := d.client.GetCurrentlyPlaying()

		if err != nil || playing == nil {
			d.playing = nil
			backoffSeconds = int(math.Min(float64(backoffSeconds*2), 32))
			time.Sleep(time.Duration(backoffSeconds) * time.Second)
			continue
		}

		backoffSeconds = 2

		if d.playing == nil || d.playing.url != playing.Image64PixelUrl {
			img := utils.ConvertUrlToImage(playing.Image64PixelUrl)

			if img != nil {
				fmt.Println("setting img")
				d.playing = &LatestPlaying{
					url: playing.Image64PixelUrl,
					img: *img,
				}
			}
		}

		time.Sleep(time.Duration(backoffSeconds) * time.Second)
	}
}
