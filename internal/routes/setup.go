package routes

import (
	config "splitflap-backend/configs"
	"splitflap-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

var cfg = config.New()

func SetupRouting(a *handlers.Application) *gin.Engine {

	r := gin.Default()

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

	r.GET("/callback", a.LoginCallback)
	r.GET("/playing", a.GetCurrentlyPlaying)

	return r
}
