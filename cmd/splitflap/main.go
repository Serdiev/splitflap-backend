package main

import (
	"context"
	"log"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/handlers"
	"splitflap-backend/internal/statemachine"
	"splitflap-backend/internal/utils"
)

var cfg = config.New()
var app handlers.Application

func main() {
	utils.SetTimeZone()

	c := context.Background()
	app = handlers.CreateService(c)

	r := handlers.SetupRouting(&app)

	go statemachine.Initiate(&app)

	// r.Run(":8080")

	server := &http.Server{
		Addr:    ":https", // Listen on HTTPS port 443
		Handler: r,
	}

	err := server.ListenAndServeTLS(cfg.General.CertFile, cfg.General.KeyFile)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
