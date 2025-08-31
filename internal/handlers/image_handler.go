package handlers

import (
	"fmt"
	"net/http"
	"splitflap-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type GetImageRequest struct {
	Id   SpotifyAccountId
	Hash string
}

type GetImageResponse struct {
	Hash       string `json:"image_url"`
	ImageBytes []byte `json:"data"`
}

func (a *Application) FetchImage(ctx *gin.Context, request GetImageRequest) {
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

	if request.Hash == img.Hash {
		ctx.JSON(http.StatusAlreadyReported, gin.H{"error": "Image has already been sent."})
		return
	}

	imgBytes, err := img.ToBytes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert image to bytes."})
		return
	}

	res := GetImageResponse{
		Hash:       img.Hash,
		ImageBytes: imgBytes,
	}

	ctx.JSON(http.StatusOK, res)
}

type SetImageRequest struct {
	Id    SpotifyAccountId `json:"id"`
	Image [][]utils.Color  `json:"image"`
}

func (a *Application) SetImage(ctx *gin.Context, request SetImageRequest) {
	lcd, exists := a.LcdDisplays[request.Id]
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "LCD client not found. You must connect a ESP32."})
		return
	}

	if len(request.Image) != 64 || len(request.Image[0]) != 64 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image must be 64x64 pixels."})
		return
	}

	img := utils.NewImage(request.Image)
	lcd.SetImage(img)

	fmt.Println("img hash", img.Hash)

	ctx.Status(http.StatusNoContent)
}
