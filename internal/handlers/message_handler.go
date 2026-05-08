package handlers

import (
	"encoding/base64"
	"net/http"

	"splitflap-backend/internal/logger"

	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	Message string `json:"message"`
}

func (a *Application) SendMessage(c *gin.Context, request SendMessageRequest) {
	_ = a.Sender.SendMessage(request.Message, "sendmessage handler")
	c.JSON(http.StatusOK, request)
}

type CurrentTextResponse struct {
	CurrentText string `json:"currentText"`
}

func (a *Application) GetCurrentMessage(c *gin.Context) {
	c.JSON(http.StatusOK, CurrentTextResponse{CurrentText: a.CurrentSplitflapText})
}

type PlayingInfo struct {
	Song     string `json:"song"`
	Artist   string `json:"artist"`
	TimeLeft string `json:"timeLeft"`
}

type StatusResponse struct {
	CurrentText    string       `json:"currentText"`
	IsEnabled      bool         `json:"isEnabled"`
	IsSpotifyLogin bool         `json:"isSpotifyLogin"`
	Image          string       `json:"image"`
	Playing        *PlayingInfo `json:"playing"`
}

func (a *Application) GetStatus(c *gin.Context) {
	var imageBase64 string
	lcd, exists := a.LcdDisplays[MainSpotifyAccountId]
	if exists {
		img := lcd.GetImage()
		if img != nil {
			imgBytes, err := img.ToBytes()
			if err == nil {
				imageBase64 = base64.StdEncoding.EncodeToString(imgBytes)
			}
		}
	}

	client, spotifyExists := a.SpotifyClients[MainSpotifyAccountId]
	isLoggedIn := false
	var playing *PlayingInfo
	if spotifyExists && client.IsLoggedIn() {
		isLoggedIn = true
		currentlyPlaying, err := client.GetCurrentlyPlaying()
		if err == nil && currentlyPlaying != nil {
			playing = &PlayingInfo{
				Song:     currentlyPlaying.Song,
				Artist:   currentlyPlaying.Artist,
				TimeLeft: currentlyPlaying.TimeLeft,
			}
		}
	}

	c.JSON(http.StatusOK, StatusResponse{
		CurrentText:    a.CurrentSplitflapText,
		IsEnabled:      a.SpotifyShouldUpdateSplitFlap,
		IsSpotifyLogin: isLoggedIn,
		Image:          imageBase64,
		Playing:        playing,
	})
}

type LogRequest struct {
	Text string `json:"text"`
}

func (a *Application) LogMessage(c *gin.Context) {
	var request LogRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Error().Msg(request.Text)
	c.JSON(http.StatusOK, gin.H{"status": "logged"})
}

func (a *Application) GetWsClientCount(c *gin.Context) {
	count := a.Ws.ClientCount()
	c.JSON(http.StatusOK, gin.H{"count": count})
}
