package main

import (
	"log"

	"signal/config"
	"signal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Connect to database
	config.ConnectDB()

	// Auto migrate models
	config.DB.AutoMigrate(&models.User{})

	// Initialize router
	router := gin.Default()

	// Basic health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"env":    config.AppConfig.Env,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start server
	log.Printf("Server starting on port %s...", config.AppConfig.Port)
	router.Run(":" + config.AppConfig.Port)
}
