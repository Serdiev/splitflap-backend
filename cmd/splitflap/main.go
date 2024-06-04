package main

import (
	"context"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/handlers"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/statemachine"
	"splitflap-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

var cfg = config.New()
var app handlers.Application

func main() {
	utils.SetTimeZone()

	closeFile := logger.InitiateLoggerToFile()
	defer closeFile()

	c := context.Background()
	app = handlers.CreateService(c)

	r := handlers.SetupRouting(&app)

	go statemachine.Initiate(&app)

	startServer(r)
}

func startServer(r *gin.Engine) {
	if cfg.General.IsLocal {
		r.Run(":8080")
		return
	}

	server := &http.Server{
		Addr:    ":https", // Listen on HTTPS port 443
		Handler: r,
	}

	err := server.ListenAndServeTLS(cfg.General.CertFile, cfg.General.KeyFile)
	if err != nil {
		logger.Error().Msgf("ListenAndServe: %s", err.Error())
	}
}
