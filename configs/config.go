package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MQTT       MQTTConfig
	Spotify    SpotifyConfig
	Splitflap  SplitFlapConfig
	IP_ADDRESS string
}

type MQTTConfig struct {
	Enabled   bool
	BrokerUrl string
	Topic     string
}

type SplitFlapConfig struct {
	ModuleCount         int
	DriverCount         int
	AlphabetOffset      string
	AlphabetCustomOrder string
	AlphabetESP32Order  string
}

type SpotifyConfig struct {
	BaseUrl      string
	TokenUrl     string
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

func (c *Configuration) GetRowLength() int {
	return c.Splitflap.ModuleCount / 2
}

var cfg *Configuration

func New() Configuration {
	if cfg != nil {
		return *cfg
	}

	godotenv.Load(".env")

	count, err := strconv.Atoi(os.Getenv("SPLITFLAP_MODULE_COUNT"))
	if err != nil {
		// panic("need module count")
		count = 6
	}

	cfg = &Configuration{
		MQTT: MQTTConfig{
			Enabled:   os.Getenv("MQTT_ENABLED") == "true",
			BrokerUrl: os.Getenv("MQTT_BROKER"),
			Topic:     os.Getenv("MQTT_TOPIC"),
		},
		Spotify: SpotifyConfig{
			BaseUrl:      os.Getenv("SPOTIFY_URL"),
			TokenUrl:     os.Getenv("SPOTIFY_TOKEN_URL"),
			ClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			RedirectUrl:  os.Getenv("SPOTIFY_REDIRECT_URL"),
		},
		Splitflap: SplitFlapConfig{
			ModuleCount:         count,
			DriverCount:         count / 6,
			AlphabetOffset:      os.Getenv("ALPHABET_OFFSET_UPPER") + os.Getenv("ALPHABET_OFFSET_LOWER"),
			AlphabetCustomOrder: os.Getenv("ALPHABET_CUSTOM_ORDER"),
			AlphabetESP32Order:  os.Getenv("ALPHABET_ARDUIN_ORDER"),
		},
	}

	if len(cfg.Splitflap.AlphabetOffset) != cfg.Splitflap.ModuleCount {
		panic("Offset setup doesnt match number of modules")
	} else if len(cfg.Splitflap.AlphabetESP32Order) != len(cfg.Splitflap.AlphabetCustomOrder) {
		panic("Mismatch between custom alphabet and ESP32 alphabet length")
	}

	return *cfg
}
