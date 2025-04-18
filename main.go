package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prasad89/devspace-api/controllers"
	"github.com/prasad89/devspace-api/initializers"
	"github.com/prasad89/devspace-api/middlewares"
)

// Initialize database connection and run migration
func init() {
	initializers.ConnectDB()
	initializers.MigrateDB()
	initializers.InitDevspaceClient()
}

func main() {
	// Create Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "DevSpace API is running!"})
	})

	// Public routes
	r.POST("/login", controllers.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/devspaces", controllers.GetDevspaces)
	protected.POST("/devspace", controllers.CreateDevspace)

	// Start API server
	log.Println("🚀 Started server on port 8080...")
	log.Fatal(r.Run(":8080"))
}
