package handlers

import (
	"net/http"
	"slices"
	"splitflap-backend/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouting(a *Application) *gin.Engine {

	r := gin.Default()
	validOrigins := []string{"http://fdevc.com", "https://fdevc.com", "https://github.com/gilmaimon/TinyWebsockets"}
	config := cors.Config{
		// AllowOrigins: validOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return slices.Contains(validOrigins, origin)
		},
	}

	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		// related to login in on spotify
		api.GET("/login", utils.ValidateQuery(a.SpotifyLogin))
		api.GET("/callback/:id", utils.ValidateQuery(a.SpotifyLoginCallback))

		api.GET("/logged-in", a.IsLoggedIn)
		api.POST("/toggle", a.ToggleSpotify)
		api.GET("/status", a.GetStatus)

		api.GET("/message", a.GetCurrentMessage)
		api.POST("/message", utils.ValidateRequest(a.SendMessage))

		api.POST("/actions", utils.ValidateRequest(a.PostAction))
		api.GET("/actions", a.GetActions)

		api.GET("/ws", a.Ws.HandleWebSocket)
		api.GET("/ws-img/:id", utils.ValidatePath(a.LcdWebsocketHandler))

		api.POST("/ip", utils.ValidateRequest(a.UpdateESP32IPAddress))

		api.GET("/image/:id/:hash", utils.ValidatePath(a.FetchImage))
		api.POST("/image/:id", utils.ValidateRequest(a.SetImage))
		api.DELETE("/image/:id", utils.ValidatePath(a.DeleteImage))
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

	r.GET("/favicon.svg", func(c *gin.Context) {
		c.File("html/favicon.svg")
	})

	r.GET("/_/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
