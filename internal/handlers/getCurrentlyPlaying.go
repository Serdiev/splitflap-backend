package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Application) GetCurrentlyPlaying(c *gin.Context) {
	playing, err := a.Spotify.GetCurrentlyPlaying()

	fmt.Println(playing)
	fmt.Println(err)

	c.JSON(http.StatusOK, playing)
}
