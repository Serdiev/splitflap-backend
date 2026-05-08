package lcd_display

import (
	"fmt"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/utils"
	ws "splitflap-backend/internal/websocket"
)

type Esp32LcdDisplay struct {
	id          string
	isPlaying   *models.SpotifyIsPlaying
	latestImage *utils.Image
	Ws          ws.WebSocket
}

func NewLcdDisplay(id string, allowedOrigins []string) *Esp32LcdDisplay {
	c := &Esp32LcdDisplay{
		id: id,
		Ws: *ws.NewWebsocket(allowedOrigins),
	}
	return c
}

func (d *Esp32LcdDisplay) GetImage() *utils.Image {
	return d.latestImage
}

func (d *Esp32LcdDisplay) SetImage(img *utils.Image) {
	d.latestImage = img

	if d.latestImage == nil {
		d.Ws.BroadcastMessageAsText([]byte("no-image"))
	} else {
		bts, err := d.latestImage.ToBytes()
		if err == nil {
			d.Ws.BroadcastMessageAsBinary(bts)
		}
	}
}

func (d *Esp32LcdDisplay) HandleIsPlaying(playing *models.SpotifyIsPlaying) {
	if playing == nil && d.isPlaying != nil {
		// broadcast that we don't have an image anymore
		fmt.Println("send no-image")
		d.latestImage = nil
		d.Ws.BroadcastMessageAsText([]byte("no-image"))
	} else if (playing != nil && d.isPlaying == nil) ||
		(playing != nil && d.isPlaying != nil && playing.Image64PixelUrl != d.isPlaying.Image64PixelUrl) {
		fmt.Println("update image to ", playing.Image64PixelUrl)
		// we have a new image to show
		d.latestImage = utils.ConvertUrlToImage(playing.Image64PixelUrl)
		bts, err := d.latestImage.ToBytes()
		if err == nil {
			d.Ws.BroadcastMessageAsBinary(bts)
		}
	}

	d.isPlaying = playing
}
