package handlers

import (
	"image"
	"image/color"
	"net/http"
	"splitflap-backend/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConvertCustomImageToRGBA(customImage utils.Image) *image.RGBA {
	height := len(customImage.Image)
	width := len(customImage.Image[0])

	// Create an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with your custom RGB data
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgb := customImage.Image[y][x]
			img.Set(x, y, color.RGBA{rgb.R, rgb.G, rgb.B, 255}) // Set pixel with full opacity
		}
	}

	return img
}

func SetupRouting(a *Application) *gin.Engine {

	r := gin.Default()
	config := cors.Config{
		AllowOrigins: []string{"https://fdev.store", "https://fdevc.store"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{"*"},
	}

	r.Use(cors.New(config))

	r.Group("/api")
	{
		r.GET("/login", a.SpotifyLogin)
		r.GET("/api/callback", a.SpotifyLoginCallback)
		r.GET("/api/playing", a.GetCurrentlyPlaying)
		r.POST("/api/toggle", a.ToggleSpotify)

		r.GET("/api/message", a.GetCurrentMessage)
		r.POST("/api/message", utils.ValidateRequest(a.SendMessage))

		r.POST("/api/actions", utils.ValidateRequest(a.PostAction))
		r.GET("/api/actions", a.GetActions)

		r.GET("/api/ws", a.Ws.HandleWebSocket)

		r.POST("/api/ip", utils.ValidateRequest(a.UpdateESP32IPAddress))
	}

	// host webpage to interact
	r.LoadHTMLGlob("html/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "web.html", gin.H{
			"title": "Welcome to the Home Page",
		})
	})

	r.GET("/_/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
