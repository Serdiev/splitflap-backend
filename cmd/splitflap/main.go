package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/handlers"
	"splitflap-backend/internal/logger"
	"splitflap-backend/internal/statemachine"
	"splitflap-backend/internal/utils"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var cfg = config.New()
var app *handlers.Application

func main() {
	utils.SetTimeZone()

	closeFile := logger.InitiateLoggerToFile()
	defer closeFile()

	c := context.Background()
	app = handlers.CreateService(c)

	r := handlers.SetupRouting(app)

	go statemachine.Initiate(app)

	server := startServer(r)

	handleGracefulShutdown(server)
	log.Info().Msg("Exiting main")
}

func startServer(r *gin.Engine) *http.Server {
	if cfg.General.IsLocal {
		err := r.Run(":443")
		if err != nil {
			panic(err)
		}

		log.Info().Msg("Starting local server on 8080")
		return nil
	}

	server := &http.Server{
		Addr:    ":https", // Listen on HTTPS port 443
		Handler: r,
	}

	go func() {
		err := server.ListenAndServeTLS(cfg.General.CertFile, cfg.General.KeyFile)
		if err != nil && err != http.ErrServerClosed {
			logger.Error().Msgf("ListenAndServeTLS: %s", err.Error())
		}
	}()

	return server
}

func handleGracefulShutdown(server *http.Server) {
	if server == nil {
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-quit

	// Gracefully shut down the server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error().Msgf("Server Shutdown: %s", err.Error())
	} else {
		logger.Info().Msg("Server exited gracefully")
	}
}
