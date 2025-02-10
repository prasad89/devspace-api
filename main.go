package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prasad89/devspace-api/controllers"
	"github.com/prasad89/devspace-api/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	log.Println("✅ Database initialized successfully!")

	// Create Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "DevSpace API is running!"})
	})

	// Register routes
	r.POST("/login", controllers.Login)

	// Start API server
	log.Println("🚀 Starting server on port 8080...")
	log.Fatal(r.Run(":8080"))
}
