package utils

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidatePath[T any](next func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query T

		// Bind JSON body if present (ignore errors for empty body on GET)
		_ = c.ShouldBindJSON(&query)

		// Reflect over struct fields
		v := reflect.ValueOf(&query).Elem()
		t := v.Type()

		// Loop through path params and try to map them to struct fields
		for _, paramName := range c.Params {
			field := findFieldByName(t, paramName.Key)
			if field != nil && v.FieldByIndex(field.Index).CanSet() {
				val := c.Param(paramName.Key)
				fv := v.FieldByIndex(field.Index)
				if fv.Kind() == reflect.String {
					fv.SetString(val)
				}
				// TODO: extend for ints, bools, etc. if you want more flexibility
			}
		}

		next(c, query)
	}
}

func findFieldByName(t reflect.Type, name string) *reflect.StructField {
	nameLower := strings.ToLower(name)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if strings.ToLower(f.Name) == nameLower {
			return &f
		}
	}
	return nil
}

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
