package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Devspace struct to hold devspace details
type Devspace struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Owner string `gorm:"uniqueIndex:owner_name_unique"`
	Name  string `gorm:"uniqueIndex:owner_name_unique"`
}

// Devspace struct to hold devspace response details
type DevspaceResponse struct {
	Name string `json:"name"`
}

// GetDevspacesByOwner retrieves all devspaces for a given user
func GetDevspacesByOwner(db *gorm.DB, c *gin.Context, username string) ([]DevspaceResponse, error) {
	var devspaces []Devspace
	results := db.Select("name").Where("owner=?", username).Find(&devspaces)

	if results.Error != nil {
		return nil, results.Error
	}

	var devspaceResponses []DevspaceResponse
	for _, devspace := range devspaces {
		devspaceResponses = append(devspaceResponses, DevspaceResponse{
			Name: devspace.Name,
		})
	}

	return devspaceResponses, nil
}
