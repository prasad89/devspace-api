package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prasad89/devspace-api/initializers"
	"github.com/prasad89/devspace-api/models"
	devspacev1alpha1 "github.com/prasad89/devspace-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// CreateDevspace creates a new DevSpace instance
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

	devspaceCR := &devspacev1alpha1.DevSpace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: username.(string),
		},
		Spec: devspacev1alpha1.DevSpaceSpec{
			Owner: username.(string),
		},
	}

	_, err := initializers.DevspaceClient.
		ApiV1alpha1().
		DevSpaces(username.(string)).
		Create(context.TODO(), devspaceCR, metav1.CreateOptions{})


	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create DevSpace in Kubernetes"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "DevSpace created successfully",
	})
}
