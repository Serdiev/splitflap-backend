package main

import (
	"context"
	"log"
	"net/http"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/handlers"
	"splitflap-backend/internal/routes"
	"splitflap-backend/internal/statemachine"
)

var cfg = config.New()
var app handlers.Application

func main() {
	c := context.Background()
	app = handlers.CreateService(c)

	r := routes.SetupRouting(&app)

	go statemachine.Initiate(&app)

	// r.Run(":8080")

	certFile := "/etc/letsencrypt/live/fdev.store/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/fdev.store/privkey.pem"

	server := &http.Server{
		Addr:    ":https", // Listen on HTTPS port 443
		Handler: r,
	}

	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
