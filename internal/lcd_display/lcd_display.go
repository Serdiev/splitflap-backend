package lcd_display

import (
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
)

type Esp32LcdDisplay struct {
	id          string
	isPlaying   *models.SpotifyIsPlaying
	latestImage *utils.Image
}

func NewLcdDisplay(id string) *Esp32LcdDisplay {
	c := &Esp32LcdDisplay{
		id: id,
		// client: client,
	}
	// c.client.RegisterHandler("update-lcd-image", c.handleIsPlaying)
	return c
}

func (d *Esp32LcdDisplay) GetImage() *utils.Image {
	return d.latestImage
}

func (d *Esp32LcdDisplay) SetImage(img *utils.Image) {
	d.latestImage = img
}

func (d *Esp32LcdDisplay) HandleIsPlaying(playing *models.SpotifyIsPlaying) {
	d.isPlaying = playing

	if d.isPlaying == nil {
		d.latestImage = nil
		return
	}

	if d.isPlaying.Image64PixelUrl != playing.Image64PixelUrl {
		d.latestImage = utils.ConvertUrlToImage(playing.Image64PixelUrl)
	}
}
