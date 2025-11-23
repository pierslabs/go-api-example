package main

import (
	"fmt"
	"log"
	"simple-go-api/internal/infrastructure/config"
	"simple-go-api/internal/infrastructure/container"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize dependency injection container
	app, err := container.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Setup routes
	app.Router.SetupRoutes()

	// Start server
	port := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server is running on port %s\n", cfg.Server.Port)

	if err := app.Engine.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
