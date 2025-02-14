package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prasad89/devspace-api/initializers"
	"github.com/prasad89/devspace-api/models"
)

// GetDevSpaces retrieves all devspaces for the logged-in user
func GetDevSpaces(c *gin.Context) {
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var devspaces []models.DevSpace
	results := initializers.DB.Where("user=?", username).Find(&devspaces)

	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch devspaces"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"devspaces": devspaces})
}
