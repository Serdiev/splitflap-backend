package utils

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ValidateRequest[T any](next func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query T

		if err := c.ShouldBindJSON(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Use reflection to inject path param "some_id" into struct if it exists
		v := reflect.ValueOf(&query).Elem()
		if field := v.FieldByName("Id"); field.IsValid() && field.CanSet() {
			if id := c.Param("id"); id != "" {
				if field.Kind() == reflect.String {
					field.SetString(id)
				}
			}
		}

		next(c, query)
	}
}

func ValidateQuery[T any](next func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query T

		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Use reflection to inject path param "some_id" into struct if it exists
		v := reflect.ValueOf(&query).Elem()
		if field := v.FieldByName("Id"); field.IsValid() && field.CanSet() {
			if id := c.Param("id"); id != "" {
				if field.Kind() == reflect.String {
					field.SetString(id)
				}
			}
		}

		next(c, query)
	}
}
