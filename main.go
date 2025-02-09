package main

import (
	"log"

	"github.com/gin-gonic/gin"
	initializer "github.com/prasad89/devspace-api/initializers"
)

func main() {
	// Connect to the database
	_, err := initializer.ConnectDB()
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}
	log.Println("✅ Database initialized successfully!")

	// Create Gin router
	r := gin.Default()

	// Example health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "DevSpace API is running!"})
	})

	// Start API server
	log.Println("🚀 Starting server on port 8080...")
	log.Fatal(r.Run(":8080"))
}
