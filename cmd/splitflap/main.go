package main

import (
	"context"
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

	r.Run(":8080")
}
