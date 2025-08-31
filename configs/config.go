package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

type SpotifyAccountConfig struct {
	Id           string
	ClientId     string
	ClientSecret string
}

type SplitFlapConfig struct {
	ModuleCount         int
	DriverCount         int
	AlphabetOffset      string
	AlphabetCustomOrder string
	AlphabetESP32Order  string
}

type SpotifyConfig struct {
	BaseUrl               string
	TokenUrl              string
	RedirectUrl           string
	SpotifyConfigurations map[string]SpotifyAccountConfig
}

func (c *Configuration) GetRowLength() int {
	return c.Splitflap.ModuleCount / 2
}

var cfg *Configuration

func New() Configuration {
	if cfg != nil {
		return *cfg
	}

	_ = godotenv.Load(".env")

	count, err := strconv.Atoi(GetVar("SPLITFLAP_MODULE_COUNT", "24"))
	if err != nil {
		panic("count not valid number")
	}

	cfg = &Configuration{
		Spotify: SpotifyConfig{
			BaseUrl:               GetVar("SPOTIFY_URL", ""),
			TokenUrl:              GetVar("SPOTIFY_TOKEN_URL", ""),
			RedirectUrl:           GetVar("SPOTIFY_REDIRECT_URL", ""),
			SpotifyConfigurations: map[string]SpotifyAccountConfig{},
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

	spotifyOwnerIds := GetVar("SPOTIFY_OWNER_IDS", "")
	spotifyClientIds := GetVar("SPOTIFY_CLIENT_IDS", "")
	spotifyClientSecrets := GetVar("SPOTIFY_CLIENT_SECRETS", "")

	if spotifyClientIds != "" && spotifyClientSecrets != "" {
		ownerIds := strings.Split(spotifyOwnerIds, ",")
		clientIds := strings.Split(spotifyClientIds, ",")
		secrets := strings.Split(spotifyClientSecrets, ",")
		if len(ownerIds) != len(clientIds) && len(clientIds) != len(secrets) {
			panic("LCD mismatch between number of Ids, ClientIds, Secrets")
		}

		for i := 0; i < len(ownerIds); i++ {
			cfg.Spotify.SpotifyConfigurations[ownerIds[i]] = SpotifyAccountConfig{
				Id:           ownerIds[i],
				ClientId:     clientIds[i],
				ClientSecret: secrets[i],
			}
		}
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
		fmt.Printf("missing env var: %s\n", str)
		return fallback
	}

	return v
}
