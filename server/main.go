package main

import (
	"fmt"
	"log"
	"net/http"

	"signal/config"
	"signal/models"
	"signal/tcp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Initialize router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.AppConfig.CORSConfig.AllowedOrigins,
		AllowMethods:     config.AppConfig.CORSConfig.AllowedMethods,
		AllowHeaders:     config.AppConfig.CORSConfig.AllowedHeaders,
		AllowCredentials: config.AppConfig.CORSConfig.AllowCredentials,
		MaxAge:           config.AppConfig.CORSConfig.MaxAge,
	}))

	// Basic health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"env":    config.AppConfig.Env,
		})
	})

	// Try to connect to database
	if err := config.ConnectDB(); err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
	} else {
		// Auto migrate models only if database connection is successful
		if err := config.DB.AutoMigrate(&models.User{}); err != nil {
			log.Printf("Warning: Database migration failed: %v", err)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	type Message struct {
		Message string `json:"message" binding:"required"`
	}
	router.POST("/send", func(c *gin.Context) {
		var msg Message

		// Bind the JSON request body to the Message struct
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Process the message
		// For example, just echo it back with a success status
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "received Message is :" + msg.Message,
		})
	})

	// Start TCP server in a goroutine
	go func() {
		fmt.Println("Starting TCP server...")
		tcp.ServerStart()
	}()

	// Start HTTP server
	log.Printf("HTTP Server starting on port %s...", config.AppConfig.Port)
	if err := router.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
