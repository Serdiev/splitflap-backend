package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MQTT_ENABLED    bool
	MQTT_BROKER_URL string
	MQTT_TOPIC      string

	SPOTIFY_URL           string
	SPOTIFY_TOKEN_URL     string
	SPOTIFY_CLIENT_ID     string
	SPOTIFY_CLIENT_SECRET string
	SPOTIFY_REDIRECT_URL  string

	ALPHA_VANTAGE_URL     string
	ALPHA_VANTAGE_API_KEY string

	SPLITFLAP_MODULE_COUNT int
	DRIVER_COUNT           int
	ALPHABET_OFFSET        string
	ALPHABET_CUSTOM_ORDER  string
	ALPHABET_ARDUIN_ORDER  string
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
		MQTT_ENABLED:    os.Getenv("MQTT_ENABLED") == "true",
		MQTT_BROKER_URL: os.Getenv("MQTT_BROKER"),
		MQTT_TOPIC:      os.Getenv("MQTT_TOPIC"),

		SPOTIFY_URL:           os.Getenv("SPOTIFY_URL"),
		SPOTIFY_TOKEN_URL:     os.Getenv("SPOTIFY_TOKEN_URL"),
		SPOTIFY_CLIENT_ID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SPOTIFY_CLIENT_SECRET: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		SPOTIFY_REDIRECT_URL:  os.Getenv("SPOTIFY_REDIRECT_URL"),

		ALPHA_VANTAGE_URL:     os.Getenv("ALPHA_VANTAGE_URL"),
		ALPHA_VANTAGE_API_KEY: os.Getenv("ALPHA_VANTAGE_API_KEY"),

		SPLITFLAP_MODULE_COUNT: count,
		DRIVER_COUNT:           count / 6,
		ALPHABET_OFFSET:        os.Getenv("ALPHABET_OFFSET_UPPER") + os.Getenv("ALPHABET_OFFSET_LOWER"),
		ALPHABET_CUSTOM_ORDER:  os.Getenv("ALPHABET_CUSTOM_ORDER"),
		ALPHABET_ARDUIN_ORDER:  os.Getenv("ALPHABET_ARDUIN_ORDER"),
	}
	// fmt.Println(os.Getenv("ALPHABET_OFFSET_UPPER"))
	// fmt.Println(os.Getenv("ALPHABET_OFFSET_LOWER"))
	// fmt.Println(cfg.ALPHABET_OFFSET)

	if len(cfg.ALPHABET_OFFSET) != cfg.SPLITFLAP_MODULE_COUNT {
		panic("Offset setup doesnt match number of modules")
	}
	if len(cfg.ALPHABET_ARDUIN_ORDER) != len(cfg.ALPHABET_CUSTOM_ORDER) {
		panic("mismatch between custom and arduino length")
	}

	return *cfg
}
