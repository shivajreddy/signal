package main

import (
	"fmt"
	"log"

	"signal/config"
	"signal/models"
	"signal/tcp"

	// Routes
	"signal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Initialize mainRouter
	mainRouter := gin.Default()

	// Configure CORS
	mainRouter.Use(cors.New(cors.Config{
		AllowOrigins:     config.AppConfig.CORSConfig.AllowedOrigins,
		AllowMethods:     config.AppConfig.CORSConfig.AllowedMethods,
		AllowHeaders:     config.AppConfig.CORSConfig.AllowedHeaders,
		AllowCredentials: config.AppConfig.CORSConfig.AllowCredentials,
		MaxAge:           config.AppConfig.CORSConfig.MaxAge,
	}))

	// Try to connect to database
	if err := config.ConnectDB(); err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
	} else {
		// Auto migrate models only if database connection is successful
		if err := config.DB.AutoMigrate(&models.User{}); err != nil {
			log.Printf("Warning: Database migration failed: %v", err)
		}
	}

	// Register Routes
	routes.RegisterAPIRoutes(mainRouter)

	// Start TCP server in a goroutine
	go func() {
		fmt.Println("Starting TCP server...")
		tcp.ServerStart()
	}()

	// Start HTTP server
	log.Printf("HTTP Server starting on port %s...", config.AppConfig.Port)
	if err := mainRouter.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
