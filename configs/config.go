package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MQTT      MQTTConfig
	Spotify   SpotifyConfig
	Splitflap SplitFlapConfig
	General   General
}

type General struct {
	TimeZone string
	CertFile string
	KeyFile  string
	IsLocal  bool
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

	count, err := strconv.Atoi(GetVar("SPLITFLAP_MODULE_COUNT", "24"))
	if err != nil {
		panic("count not valid number")
	}

	cfg = &Configuration{
		// MQTT: MQTTConfig{
		// 	Enabled:   GetVar("MQTT_ENABLED") == "true",
		// 	BrokerUrl: GetVar("MQTT_BROKER"),
		// 	Topic:     GetVar("MQTT_TOPIC"),
		// },
		Spotify: SpotifyConfig{
			BaseUrl:      GetVar("SPOTIFY_URL", ""),
			TokenUrl:     GetVar("SPOTIFY_TOKEN_URL", ""),
			ClientId:     GetVar("SPOTIFY_CLIENT_ID", ""),
			ClientSecret: GetVar("SPOTIFY_CLIENT_SECRET", ""),
			RedirectUrl:  GetVar("SPOTIFY_REDIRECT_URL", ""),
		},
		Splitflap: SplitFlapConfig{
			ModuleCount:         count,
			DriverCount:         count / 6,
			AlphabetOffset:      GetVar("ALPHABET_OFFSET_UPPER", ",+%%9%::,%::") + GetVar("ALPHABET_OFFSET_LOWER", ",%:::::,:,,:"),
			AlphabetCustomOrder: GetVar("ALPHABET_CUSTOM_ORDER", " abcdefghijklmnpqrstuvwxy0123456789+-%,:"),
			AlphabetESP32Order:  GetVar("ALPHABET_ESP32_ORDER", " abcdefghijklmnopqrstuvwxyz0123456789.,'"),
		},
		General: General{
			TimeZone: GetVar("TIMEZONE", "Europe/Stockholm"),
			CertFile: GetVar("LETSENCRYPT_CERTFILE", ""),
			KeyFile:  GetVar("LETSENCRYPT_KEYFILE", ""),
			IsLocal:  GetVar("IS_LOCAL", "f") == "TRUE",
		},
	}

	if len(cfg.Splitflap.AlphabetOffset) != cfg.Splitflap.ModuleCount {
		panic("offset setup doesnt match number of modules")
	} else if len(cfg.Splitflap.AlphabetESP32Order) != len(cfg.Splitflap.AlphabetCustomOrder) {
		panic("mismatch between custom alphabet and ESP32 alphabet length")
	}

	return *cfg
}

func GetVar(str string, fallback string) string {
	v := os.Getenv(str)
	if v == "" {
		fmt.Println(fmt.Sprintf("missing env var: %s", str))
		return fallback
	}

	return v
}
