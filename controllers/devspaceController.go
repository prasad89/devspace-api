package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prasad89/devspace-api/initializers"
	"github.com/prasad89/devspace-api/models"
)

// GetDevspaces retrieves all devspaces for the logged-in user
func GetDevspaces(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	devspaces, err := models.GetDevspacesByOwner(initializers.DB, c, username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get devspaces"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"devspaces": devspaces})
}

// // CreateDevspace creates a new DevSpace instance
func CreateDevspace(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	devspace := models.Devspace{
		Owner: username.(string),
		Name:  req.Name,
	}

	if err := initializers.DB.Create(&devspace).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Devspace with same name already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Devspace created successfully",
	})
}
