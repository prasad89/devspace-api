package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prasad89/devspace-api/models"
)

// SecretKey to sign JWT
var SecretKey = []byte("secret") // Currently hardcoded but in k8s environment, it should be stored in a secret and accessed from there

// Auth struct to hold login credentials
type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login authenticates the user
func Login(c *gin.Context) {
	var auth Auth
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := models.GetByUsername(auth.Username)
	if err != nil || user.Password != auth.Password {
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
