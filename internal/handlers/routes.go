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

	api := r.Group("/api")
	{
		// related to login in on spotify
		api.GET("/login", utils.ValidateQuery(a.SpotifyLogin))
		api.GET("/callback/:id", utils.ValidateQuery(a.SpotifyLoginCallback))

		api.GET("/logged-in", a.IsLoggedIn)
		api.GET("/playing", a.GetCurrentlyPlaying)
		api.POST("/toggle", a.ToggleSpotify)

		api.GET("/message", a.GetCurrentMessage)
		api.POST("/message", utils.ValidateRequest(a.SendMessage))

		api.POST("/actions", utils.ValidateRequest(a.PostAction))
		api.GET("/actions", a.GetActions)

		api.GET("/ws", a.Ws.HandleWebSocket)

		api.POST("/ip", utils.ValidateRequest(a.UpdateESP32IPAddress))

		api.GET("/image/:id/:hash", utils.ValidatePath(a.FetchImage))
		api.POST("/image/:id", utils.ValidateRequest(a.SetImage))
	}

	// host webpage to interact
	r.LoadHTMLGlob("html/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "web.html", gin.H{
			"title": "Welcome to the Home Page",
		})
	})

	r.GET("/webcam", func(c *gin.Context) {
		c.HTML(http.StatusOK, "webcam.html", gin.H{
			"title": "Welcome to the Home Page",
		})
	})

	r.GET("/_/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
