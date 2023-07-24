package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MQTT_BROKER_URL string
	MQTT_TOPIC      string

	SPOTIFY_URL           string
	SPOTIFY_TOKEN_URL     string
	SPOTIFY_CLIENT_ID     string
	SPOTIFY_CLIENT_SECRET string
	SPOTIFY_REDIRECT_URL  string

	SPLITFLAP_MODULE_COUNT int
}

var cfg *Configuration

func New() Configuration {
	if cfg != nil {
		return *cfg
	}

	godotenv.Load(".env")

	count, err := strconv.Atoi(os.Getenv("SPLITFLAP_MODULE_COUNT"))
	if err != nil {
		panic("need module count")
	}

	cfg = &Configuration{
		MQTT_BROKER_URL:        os.Getenv("MQTT_BROKER"),
		MQTT_TOPIC:             os.Getenv("MQTT_TOPIC"),
		SPOTIFY_URL:            os.Getenv("SPOTIFY_URL"),
		SPOTIFY_TOKEN_URL:      os.Getenv("SPOTIFY_TOKEN_URL"),
		SPOTIFY_CLIENT_ID:      os.Getenv("SPOTIFY_CLIENT_ID"),
		SPOTIFY_CLIENT_SECRET:  os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SPOTIFY_REDIRECT_URL:   os.Getenv("SPOTIFY_REDIRECT_URL"),
		SPLITFLAP_MODULE_COUNT: count,
	}

	return *cfg
}