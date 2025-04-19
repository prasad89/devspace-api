package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prasad89/devspace-api/initializers"
	"github.com/prasad89/devspace-api/models"
)

// SecretKey to sign JWT
var SecretKey = []byte("secret") // In production, store this in a secure environment (e.g., K8s Secret)

// Credentials struct to hold login credentials
type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login authenticates the user
func Login(c *gin.Context) {
	var Credentials Credentials
	if err := c.ShouldBindJSON(&Credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	user, err := models.GetByUsername(initializers.DB, Credentials.Username)
	if err != nil || user.Password != Credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
