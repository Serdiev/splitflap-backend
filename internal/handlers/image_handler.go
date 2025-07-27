package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	Id       string
	ImageUrl string `json:"image_url"`
}

type ImageResponse struct {
	ImageUrl   string `json:"image_url"`
	ImageBytes []byte `json:"data"`
}

func (a *Application) FetchImage(ctx *gin.Context, request ImageRequest) {
	lcd, exists := a.LcdDisplays[request.Id]
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "LCD client not found. You must login first."})
		return
	}

	img := lcd.GetImage()
	if img == nil {
		// This is pure optimization so ESP32 doesn't have to parse a new image
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Don't have image to return."})
		return
	}

	if request.ImageUrl == img.Url {
		ctx.JSON(http.StatusAlreadyReported, gin.H{"error": "Image has already been sent."})
		return
	}

	imgBytes, err := img.ToBytes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert image to bytes."})
		return
	}

	res := ImageResponse{
		ImageUrl:   img.Url,
		ImageBytes: imgBytes,
	}

	ctx.JSON(http.StatusOK, res)
}
