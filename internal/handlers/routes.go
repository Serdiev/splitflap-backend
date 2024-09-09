package handlers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouting(a *Application) *gin.Engine {

	r := gin.Default()
	config := cors.Config{
		AllowOrigins: []string{"https://fdev.store"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{"*"},
	}

	r.Use(cors.New(config))

	r.GET("/login", a.SpotifyLogin)
	r.GET("/callback", a.SpotifyLoginCallback)
	r.GET("/playing", a.GetCurrentlyPlaying)
	r.POST("/toggle", a.ToggleSpotify)

	r.POST("/message", a.SendMessage)

	r.POST("/actions", a.ExecuteAction)
	r.GET("/actions", a.GetActions)

	// host webpage to interact
	r.LoadHTMLGlob("html/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "web.html", gin.H{
			"title": "Welcome to the Home Page",
		})
	})

	return r
}
