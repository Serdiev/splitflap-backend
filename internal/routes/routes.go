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

	r.GET("/login", a.SpotifyLogin)
	r.GET("/callback", a.SpotifyLoginCallback)
	r.GET("/playing", a.GetCurrentlyPlaying)

	r.POST("/message", a.SendMessage)

	// host webpage to interact
	r.LoadHTMLGlob("html/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "web.html", gin.H{
			"title": "Welcome to the Home Page",
		})
	})

	return r
}
