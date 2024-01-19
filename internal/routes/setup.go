package routes

import (
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var cfg = config.New()

func SetupRouting(a *handlers.Application) *gin.Engine {

	r := gin.Default()
	config := cors.Config{}
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"*"} // Add "OPTIONS" to the allowed methods
	config.AllowHeaders = []string{"*"} // Allow the necessary headers

	r.Use(cors.New(config))

	r.GET("/login", func(c *gin.Context) {
		responseType := c.DefaultQuery("response_type", "code")
		clientID := c.DefaultQuery("client_id", cfg.SPOTIFY_CLIENT_ID)
		scope := c.DefaultQuery("scope", "user-read-currently-playing")
		redirectURI := c.DefaultQuery("redirect_uri", "http://localhost:8080/callback")

		// Construct the redirect URL
		redirectURL := "https://accounts.spotify.com/authorize?" +
			"response_type=" + responseType +
			"&client_id=" + clientID +
			"&scope=" + scope +
			"&redirect_uri=" + redirectURI

		c.Redirect(307, redirectURL)
	})

	// spotify
	r.GET("/callback", a.LoginCallback)
	r.GET("/playing", a.GetCurrentlyPlaying)

	r.POST("/message", a.SendMessage)

	// host webpage to interact
	r.LoadHTMLGlob("html/*.html")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "web.html", gin.H{
			"title": "Welcome to the Home Page",
		})
	})

	return r
}
