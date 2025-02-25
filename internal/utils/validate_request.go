package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequest[T any](next func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request = new(T)

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		next(c, *request)
	}
}
