package lcd_display

import (
	config "splitflap-backend/configs"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/models"
	"splitflap-backend/internal/sender"
	"splitflap-backend/internal/utils"
	ws "splitflap-backend/internal/websocket"
)

var cfg = config.New()

type Esp32LcdDisplay struct {
	id          string
	isPlaying   *models.SpotifyIsPlaying
	latestImage *utils.Image
	Ws          ws.WebSocket
	mqtt        *sender.MQTTSender
}

func NewLcdDisplay(id string, allowedOrigins []string) *Esp32LcdDisplay {
	c := &Esp32LcdDisplay{
		id:   id,
		Ws:   *ws.NewWebsocket(allowedOrigins),
		mqtt: sender.NewMQTTSender(cfg.MQTT.LcdTopic + "/" + id),
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
		logger.Info().Msg("LCD: sending no-image")
		d.latestImage = nil
		d.Ws.BroadcastMessageAsText([]byte("no-image"))
		d.mqtt.SendMessage("")
	} else if (playing != nil && d.isPlaying == nil) ||
		(playing != nil && d.isPlaying != nil && playing.Image64PixelUrl != d.isPlaying.Image64PixelUrl) {

		logger.Info().Str("url", playing.Image64PixelUrl).Msg("LCD: updating image")
		mqttErr := d.mqtt.SendMessage(playing.Image64PixelUrl)
		if mqttErr != nil {
			logger.Error().Err(mqttErr).Msg("LCD: failed to send MQTT message")
		}

		d.latestImage = utils.ConvertUrlToImage(playing.Image64PixelUrl)
		bts, err := d.latestImage.ToBytes()
		if err == nil {
			d.Ws.BroadcastMessageAsBinary(bts)
		}
	}

	d.isPlaying = playing
}

// mosquitto_pub -h 192.168.1.92 -t mqtt-lcd-display -m "hello world" -u esp32 -P lcddisplay
